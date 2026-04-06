package types

import (
	"github.com/cosmos/gogoproto/proto"
	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ExecuteRequest struct {
	Msg    sdk.Msg
	Caller common.Address

	// options
	AllowFailure bool
	CallbackId   uint64

	// supplied gas
	GasLimit uint64
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
