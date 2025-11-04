package types

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"
)

func TestRegisterLegacyAminoCodec(t *testing.T) {
	cdc := codec.NewLegacyAmino()

	// Test that registration doesn't panic
	require.NotPanics(t, func() {
		RegisterLegacyAminoCodec(cdc)
	})

	// Test that concrete types can be marshaled and unmarshaled
	concreteTypes := []struct {
		name string
		obj  interface{}
	}{
		{
			name: "ContractAccount",
			obj: &ContractAccount{
				BaseAccount: &authtypes.BaseAccount{
					Address: "init1test",
				},
			},
		},
		{
			name: "ShorthandAccount",
			obj: &ShorthandAccount{
				BaseAccount: &authtypes.BaseAccount{
					Address: "init1test",
				},
				OriginalAddress: "0x1234567890123456789012345678901234567890",
			},
		},
		{
			name: "CallAuthorization",
			obj: &CallAuthorization{
				Contracts: []string{"0x1234567890123456789012345678901234567890"},
			},
		},
	}

	for _, tc := range concreteTypes {
		t.Run(tc.name+"_amino_registered", func(t *testing.T) {
			// Try to marshal the type using amino codec
			bz, err := cdc.Marshal(tc.obj)
			require.NoError(t, err)
			require.NotEmpty(t, bz)

			// Try to unmarshal back into a new object of the same type
			switch tc.name {
			case "ContractAccount":
				var newObj ContractAccount
				err = cdc.Unmarshal(bz, &newObj)
				require.NoError(t, err)
				require.Equal(t, tc.obj.(*ContractAccount).Address, newObj.Address)
			case "ShorthandAccount":
				var newObj ShorthandAccount
				err = cdc.Unmarshal(bz, &newObj)
				require.NoError(t, err)
				require.Equal(t, tc.obj.(*ShorthandAccount).Address, newObj.Address)
				require.Equal(t, tc.obj.(*ShorthandAccount).OriginalAddress, newObj.OriginalAddress)
			case "CallAuthorization":
				var newObj CallAuthorization
				err = cdc.Unmarshal(bz, &newObj)
				require.NoError(t, err)
				require.Equal(t, tc.obj.(*CallAuthorization).Contracts, newObj.Contracts)
			}
		})
	}
}

func TestCodecIntegration(t *testing.T) {
	// Test that both registrations work together
	cdc := codec.NewLegacyAmino()
	registry := cdctypes.NewInterfaceRegistry()

	require.NotPanics(t, func() {
		RegisterLegacyAminoCodec(cdc)
		RegisterInterfaces(registry)
	})

	// Test that we can create a codec with both registrations
	interfaceRegistry := cdctypes.NewInterfaceRegistry()
	RegisterInterfaces(interfaceRegistry)

	protoCodec := codec.NewProtoCodec(interfaceRegistry)
	require.NotNil(t, protoCodec)

	// Test that we can marshal and unmarshal a message
	msg := &MsgCall{
		Sender:       "0x1234567890123456789012345678901234567890",
		ContractAddr: "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd",
		Input:        "0x12345678",
	}

	// Test proto marshaling
	protoData, err := protoCodec.Marshal(msg)
	require.NoError(t, err)
	require.NotEmpty(t, protoData)

	// Test proto unmarshaling
	var unmarshaledMsg MsgCall
	err = protoCodec.Unmarshal(protoData, &unmarshaledMsg)
	require.NoError(t, err)
	require.Equal(t, msg.Sender, unmarshaledMsg.Sender)
	require.Equal(t, msg.ContractAddr, unmarshaledMsg.ContractAddr)
	require.Equal(t, msg.Input, unmarshaledMsg.Input)
}
