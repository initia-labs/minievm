package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"

	"github.com/ethereum/go-ethereum/core/vm"
)

type ExecuteRequest struct {
	Msg    sdk.Msg
	Caller vm.ContractRef

	// options
	AllowFailure bool
	CallbackId   uint64
}

// ExtractLogsFromResponse extracts msg response from the data
func ExtractLogsFromResponse(data []byte, msgTypeURL string) (logs Logs, err error) {
	switch msgTypeURL {
	case sdk.MsgTypeURL(&MsgCall{}):
		var resp MsgCallResponse
		err = proto.Unmarshal(data, &resp)
		if err != nil {
			return
		}

		logs = resp.Logs
	case sdk.MsgTypeURL(&MsgCreate{}):
		var resp MsgCreateResponse
		err = proto.Unmarshal(data, &resp)
		if err != nil {
			return
		}

		logs = resp.Logs
	case sdk.MsgTypeURL(&MsgCreate2{}):
		var resp MsgCreate2Response
		err = proto.Unmarshal(data, &resp)
		if err != nil {
			return
		}

		logs = resp.Logs
	default:
		return
	}

	return
}
