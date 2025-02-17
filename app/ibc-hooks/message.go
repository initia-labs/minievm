package evm_hooks

import (
	"encoding/json"

	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

// A contract that sends an IBC transfer, may need to listen for the ACK from that packet.
// To allow contracts to listen on the ack of specific packets, we provide Ack callbacks.
//
// The contract, which wants to receive ack callback, have to implement two functions
// - ibc_ack
// - ibc_timeout
//
// interface IIBCAsyncCallback {
//     function ibc_ack(uint64 callback_id, bool success) external;
//     function ibc_timeout(uint64 callback_id) external;
// }
//

const (
	// The memo key is used to parse ics-20 or ics-712 memo fields.
	evmHookMemoKey = "evm"

	functionNameAck     = "ibc_ack"
	functionNameTimeout = "ibc_timeout"
)

// AsyncCallback is data wrapper which is required
// when we implement async callback.
type AsyncCallback struct {
	// callback id should be issued form the executor contract
	Id              uint64 `json:"id"`
	ContractAddress string `json:"contract_address"`
}

// HookData defines a wrapper for evm execute message
// and async callback.
type HookData struct {
	// Message is a evm execute message which will be executed
	// at `OnRecvPacket` of receiver chain.
	Message *evmtypes.MsgCall `json:"message,omitempty"`

	// AsyncCallback is a contract address
	AsyncCallback *AsyncCallback `json:"async_callback,omitempty"`
}

// asyncCallback is same as AsyncCallback.
type asyncCallback struct {
	// callback id should be issued from the executor contract
	Id              uint64 `json:"id"`
	ContractAddress string `json:"contract_address"`
}

// asyncCallbackStringID is same as AsyncCallback but
// it has Id as string.
type asyncCallbackStringID struct {
	// callback id should be issued form the executor contract
	Id              uint64 `json:"id,string"`
	ContractAddress string `json:"contract_address"`
}

// UnmarshalJSON implements the json unmarshaler interface.
// custom unmarshaler is required because we have to handle
// id as string and uint64.
func (a *AsyncCallback) UnmarshalJSON(bz []byte) error {
	var ac asyncCallback
	err := json.Unmarshal(bz, &ac)
	if err != nil {
		var acStr asyncCallbackStringID
		err := json.Unmarshal(bz, &acStr)
		if err != nil {
			return err
		}

		a.Id = acStr.Id
		a.ContractAddress = acStr.ContractAddress
		return nil
	}

	a.Id = ac.Id
	a.ContractAddress = ac.ContractAddress
	return nil
}
