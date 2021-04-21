package types_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	markertypes "github.com/provenance-io/provenance/x/marker/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/provenance-io/provenance/x/authz/exported"
	"github.com/provenance-io/provenance/x/authz/types"
)

var (
	coinsPos = sdk.NewCoins(sdk.NewInt64Coin("steak", 100))
	granter  = sdk.AccAddress("_______granter______")
	grantee  = sdk.AccAddress("_______grantee______")
)

func TestMsgExecAuthorized(t *testing.T) {
	tests := []struct {
		title      string
		grantee    sdk.AccAddress
		msgs       []sdk.ServiceMsg
		expectPass bool
	}{
		{"nil grantee address", nil, []sdk.ServiceMsg{}, false},
		{"zero-messages test: should fail", grantee, []sdk.ServiceMsg{}, false},
		{"valid test: msg type", grantee, []sdk.ServiceMsg{
			{
				MethodName: markertypes.MarkerSendAuthorization{}.MethodName(),
				Request: &markertypes.MsgTransferRequest{
					Amount:      sdk.NewCoin("steak", sdk.NewInt(2)),
					FromAddress: granter.String(),
					ToAddress:   grantee.String(),
					Administrator: grantee.String(),
				},
			},
		}, true},
	}
	for i, tc := range tests {
		msg := types.NewMsgExecAuthorized(tc.grantee, tc.msgs)
		if tc.expectPass {
			require.NoError(t, msg.ValidateBasic(), "test: %v", i)
		} else {
			require.Error(t, msg.ValidateBasic(), "test: %v", i)
		}
	}
}
func TestMsgRevokeAuthorization(t *testing.T) {
	tests := []struct {
		title            string
		granter, grantee sdk.AccAddress
		msgType          string
		expectPass       bool
	}{
		{"nil Granter address", nil, grantee, "hello", false},
		{"nil Grantee address", granter, nil, "hello", false},
		{"nil Granter and Grantee address", nil, nil, "hello", false},
		{"valid test case", granter, grantee, "hello", true},
	}
	for i, tc := range tests {
		msg := types.NewMsgRevokeAuthorization(tc.granter, tc.grantee, tc.msgType)
		if tc.expectPass {
			require.NoError(t, msg.ValidateBasic(), "test: %v", i)
		} else {
			require.Error(t, msg.ValidateBasic(), "test: %v", i)
		}
	}
}

func TestMsgGrantAuthorization(t *testing.T) {
	tests := []struct {
		title            string
		granter, grantee sdk.AccAddress
		authorization    exported.Authorization
		expiration       time.Time
		expectErr        bool
		expectPass       bool
	}{
		{"nil granter address", nil, grantee, &markertypes.MarkerSendAuthorization{SpendLimit: coinsPos}, time.Now(), false, false},
		{"nil grantee address", granter, nil, &markertypes.MarkerSendAuthorization{SpendLimit: coinsPos}, time.Now(), false, false},
		{"nil granter and grantee address", nil, nil, &markertypes.MarkerSendAuthorization{SpendLimit: coinsPos}, time.Now(), false, false},
		{"nil authorization", granter, grantee, nil, time.Now(), true, false},
		{"valid test case", granter, grantee, &markertypes.MarkerSendAuthorization{SpendLimit: coinsPos}, time.Now().AddDate(0, 1, 0), false, true},
		{"past time", granter, grantee, &markertypes.MarkerSendAuthorization{SpendLimit: coinsPos}, time.Now().AddDate(0, 0, -1), false, false},
	}
	for i, tc := range tests {
		msg, err := types.NewMsgGrantAuthorization(
			tc.granter, tc.grantee, tc.authorization, tc.expiration,
		)
		if !tc.expectErr {
			require.NoError(t, err)
		} else {
			continue
		}
		if tc.expectPass {
			require.NoError(t, msg.ValidateBasic(), "test: %v", i)
		} else {
			require.Error(t, msg.ValidateBasic(), "test: %v", i)
		}
	}
}
