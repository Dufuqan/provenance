package handlers

import (
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/provenance-io/provenance/internal/antewrapper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	msgbasedfeetypes "github.com/provenance-io/provenance/x/msgfees/types"
)

type MsgBasedFeeInvoker struct {
	msgBasedFeeKeeper msgbasedfeetypes.MsgBasedFeeKeeper
	bankKeeper        msgbasedfeetypes.BankKeeper
	accountKeeper     msgbasedfeetypes.AccountKeeper
	feegrantKeeper    msgbasedfeetypes.FeegrantKeeper
	txDecoder         sdk.TxDecoder
}

// NewMsgBasedFeeInvoker concrete impl of how to charge Msg Based Fees
func NewMsgBasedFeeInvoker(bankKeeper msgbasedfeetypes.BankKeeper, accountKeeper msgbasedfeetypes.AccountKeeper,
	feegrantKeeper msgbasedfeetypes.FeegrantKeeper, msgBasedFeeKeeper msgbasedfeetypes.MsgBasedFeeKeeper, decoder sdk.TxDecoder) MsgBasedFeeInvoker {
	return MsgBasedFeeInvoker{
		msgBasedFeeKeeper,
		bankKeeper,
		accountKeeper,
		feegrantKeeper,
		decoder,
	}
}

func (afd MsgBasedFeeInvoker) Invoke(ctx sdk.Context, simulate bool) (coins sdk.Coins, events sdk.Events, err error) {
	chargedFees := sdk.Coins{}
	eventsToReturn := sdk.Events{}

	if ctx.TxBytes() != nil && len(ctx.TxBytes()) != 0 {
		originalGasMeter := ctx.GasMeter()

		tx, err := afd.txDecoder(ctx.TxBytes())
		if err != nil {
			panic(fmt.Errorf("error in chargeFees() while getting txBytes: %w", err))
		}

		// cast to FeeTx
		feeTx, ok := tx.(sdk.FeeTx)
		// only charge additional fee if of type FeeTx since it should give fee payer.
		// for provenance should be a FeeTx since antehandler should enforce it, but
		// not adding complexity here
		if !ok {
			panic("Provenance only supports feeTx for now")
		}
		feePayer := feeTx.FeePayer()
		feeGranter := feeTx.FeeGranter()
		deductFeesFrom := feePayer
		// if fee granter set deduct fee from feegranter account.
		// this works with only when feegrant enabled.

		feeGasMeter, ok := ctx.GasMeter().(*antewrapper.FeeGasMeter)
		if !ok {
			// all provenance tx's should have this set
			panic("GasMeter is not of type FeeGasMeter")
		}
		chargedFees = feeGasMeter.FeeConsumed()

		// check chargedFees is not nil && is not all zero(IsZero returns true if there are no coins or all coins are zero.)
		if chargedFees != nil && !chargedFees.IsZero() {
			// there should not be any negative coins, just to be very sure here
			if chargedFees.IsAnyNegative() {
				return nil, nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "charged fees %v are negative, which should not be possible, aborting", chargedFees)
			}
			// eat up the gas cost for charging fees. (This one is on us, Cheers!, mainly because we don't want to fail at this step, imo, but we can remove this is f necessary)
			ctx = ctx.WithGasMeter(sdk.NewInfiniteGasMeter())
			// if feegranter set deduct fee from feegranter account.
			// this works with only when feegrant enabled.
			if feeGranter != nil {
				if afd.feegrantKeeper == nil {
					return nil, nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "fee grants are not enabled")
				} else if !feeGranter.Equals(feePayer) {
					err = afd.feegrantKeeper.UseGrantedFees(ctx, feeGranter, feePayer, chargedFees, tx.GetMsgs())
					if err != nil {
						return nil, nil, sdkerrors.Wrapf(err, "%s not allowed to pay fees from %s", feeGranter, feePayer)
					}
				}
				deductFeesFrom = feeGranter
			}
			deductFeesFromAcc := afd.accountKeeper.GetAccount(ctx, deductFeesFrom)
			if deductFeesFromAcc == nil {
				return nil, nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "fee payer address: %s does not exist", deductFeesFrom)
			}

			ctx.Logger().Debug(fmt.Sprintf("The Fee consumed by message types : %v", feeGasMeter.FeeConsumedByMsg()))

			baseFeeConsumedAtAnteHandler := feeGasMeter.BaseFeeConsumed()

			var isNeg bool
			// this sweeps all extra fees too, 1. keeps current behavior 2. accounts for priority mempool
			chargedFees, isNeg = feeTx.GetFee().SafeSub(baseFeeConsumedAtAnteHandler)

			if isNeg {
				return nil, nil, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "This should never happen, because fee charged in ante handlers should never have been more than fee promised without error'ing out.")
			}

			if len(chargedFees) > 0 && chargedFees.IsAllPositive() {
				err = afd.msgBasedFeeKeeper.DeductFees(afd.bankKeeper, ctx, deductFeesFromAcc, chargedFees)
				if err != nil {
					return nil, nil, err
				}
			}
			eventsToReturn = sdk.Events{
				sdk.NewEvent(sdk.EventTypeTx,
					sdk.NewAttribute(antewrapper.AttributeKeyAdditionalFee, feeGasMeter.FeeConsumed().String()),
				),
				sdk.NewEvent(sdk.EventTypeTx,
					sdk.NewAttribute(antewrapper.AttributeKeyBaseFee, feeGasMeter.BaseFeeConsumed().Add(chargedFees...).Sub(feeGasMeter.FeeConsumed()).String()),
				)}
		}

		// set back the original gasMeter
		ctx = ctx.WithGasMeter(originalGasMeter)
	}

	return chargedFees, eventsToReturn, nil
}
