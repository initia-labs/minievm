# IBC-hooks

This module is copied from [osmosis](https://github.com/osmosis-labs/osmosis) and changed to execute evm contract with ICS-20 token transfer calls.

## EVM Hooks

The evm hook is an IBC middleware which is used to allow ICS-20 token transfers to initiate contract calls.
This allows cross-chain contract calls, that involve token movement.
This is useful for a variety of use cases.
One of primary importance is cross-chain swaps, which is an extremely powerful primitive.

The mechanism enabling this is a `memo` field on every ICS20 and ICS721 transfer packet as of [IBC v3.4.0](https://medium.com/the-interchain-foundation/moving-beyond-simple-token-transfers-d42b2b1dc29b).
Move hooks is an IBC middleware that parses an ICS20 transfer, and if the `memo` field is of a particular form, executes a evm contract call. We now detail the `memo` format for `evm` contract calls, and the execution guarantees provided.

### EVM Contract Execution Format

Before we dive into the IBC metadata format, we show the hook data format, so the reader has a sense of what are the fields we need to be setting in.
The evm `MsgCall` is defined [here](../../x/evm/types/tx.pb.go) and other types are defined [here](./message.go) as the following type:

```go
// HookData defines a wrapper for evm execute message
// and async callback.
type HookData struct {
 // Message is a evm execute message which will be executed
 // at `OnRecvPacket` of receiver chain.
 Message evmtypes.MsgCall `json:"message"`

 // AsyncCallback is a callback message which will be executed
 // at `OnTimeoutPacket` and `OnAcknowledgementPacket` of
 // sender chain.
 AsyncCallback *AsyncCallback `json:"async_callback,omitempty"`
}

// AsyncCallback is data wrapper which is required
// when we implement async callback.
type AsyncCallback struct {
 // callback id should be issued form the executor contract
 Id            uint64 `json:"id"`
 ContractAddr  string `json:"contract_addr"`
}

// MsgCall is a message to call an Ethereum contract.
type MsgCall struct {
 // Sender is the that actor that signed the messages
 Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
 // ContractAddr is the contract address to be executed.
 // It can be cosmos address or hex encoded address.
 ContractAddr string `protobuf:"bytes,2,opt,name=contract_addr,json=contractAddr,proto3" json:"contract_addr,omitempty"`
 // Hex encoded execution input bytes.
 Input string `protobuf:"bytes,3,opt,name=input,proto3" json:"input,omitempty"`
 // Value is the amount of fee denom token to transfer to the contract.
 Value string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
 // AccessList is a predefined list of Ethereum addresses and their corresponding storage slots that a transaction will interact with during its execution. can be none
 AccessList []AccessTuple `protobuf:"bytes,5,rep,name=access_list,json=accessList,proto3" json:"access_list"`
}
```

So we detail where we want to get each of these fields from:

- Sender: We cannot trust the sender of an IBC packet, the counter-party chain has full ability to lie about it.
  We cannot risk this sender being confused for a particular user or module address on Initia.
  So we replace the sender with an account to represent the sender prefixed by the channel and a evm module prefix.
  This is done by setting the sender to `Bech32(Hash(Hash("ibc-evm-hook-intermediary") + channelID/sender))`, where the channelId is the channel id on the local chain.
- ContractAddr: This field should be directly obtained from the ICS-20 packet metadata
- Input: This field should be directly obtained from the ICS-20 packet metadata.

So our constructed evm call message that we execute will look like:

```go
msg := MsgCall{
 // Sender is the that actor that signed the messages
 Sender: "init1-hash-of-channel-and-sender",
 // ContractAddr is the contract address to be executed.
 // It can be cosmos address or hex encoded address.
 ContractAddr: packet.data.memo["evm"]["message"]["contract_addr"],
 // Hex encoded execution input bytes.
 Input: packet.data.memo["evm"]["message"]["input"],
 // Value is the amount of fee denom token to transfer to the contract.
 Value: packet.data.memo["evm"]["message"]["value"]
 // Value is the amount of fee denom token to transfer to the contract.
 AccessList: packet.data.memo["evm"]["message"]["access_list"]
}
```

### ICS20 packet structure

So given the details above, we propogate the implied ICS20 packet data structure.
ICS20 is JSON native, so we use JSON for the memo format.

```json
{
  //... other ibc fields that we don't care about
  "data": {
    "denom": "denom on counterparty chain (e.g. uatom)", // will be transformed to the local denom (ibc/...)
    "amount": "1000",
    "sender": "addr on counterparty chain", // will be transformed
    "receiver": "ModuleAddr::ModuleName::FunctionName",
    "memo": {
      "evm": {
        // execute message on receive packet
        "message": {
          "contract_addr": "0x1",
          "input": "hex encoded byte string",
          "value": "0"
        },
        // optional field to get async callback (ack and timeout)
        "async_callback": {
          "id": 1,
          "contract_addr": "0x1"
        }
      }
    }
  }
}
```

An ICS20 packet is formatted correctly for evmhooks iff the following all hold:

- `memo` is not blank
- `memo` is valid JSON
- `memo` has at least one key, with value `"evm"`
- `memo["evm"]["message"]` has exactly 3 entries, `"contract_addr"`, `"input"`, `"value"`
- `receiver` == "" || `receiver` == "module_address::module_name::function_name"

We consider an ICS20 packet as directed towards evmhooks iff all of the following hold:

- `memo` is not blank
- `memo` is valid JSON
- `memo` has at least one key, with name `"evm"`

If an ICS20 packet is not directed towards evmhooks, evmhooks doesn't do anything.
If an ICS20 packet is directed towards evmhooks, and is formatted incorrectly, then evmhooks returns an error.

### Execution flow

Pre evm hooks:

- Ensure the incoming IBC packet is cryptogaphically valid
- Ensure the incoming IBC packet is not timed out.

In evm hooks, pre packet execution:

- Ensure the packet is correctly formatted (as defined above)
- Edit the receiver to be the hardcoded IBC module account

In evm hooks, post packet execution:

- Construct evm message as defined before
- Execute evm message
- if evm message has error, return ErrAck
- otherwise continue through middleware

### Async Callback

A contract that sends an IBC transfer, may need to listen for the ACK from that packet.
To allow contracts to listen on the ack of specific packets, we provide Ack callbacks.
The contract, which wants to receive ack callback, have to implement two functions.

- ibc_ack
- ibc_timeout

```solidity
interface IIBCAsyncCallback {
    function ibc_ack(uint64 callback_id, bool success) external;
    function ibc_timeout(uint64 callback_id) external;
}
```

Also when a contract make IBC transfer request, it should provide async callback data through memo field.

- `memo['evm']['async_callback']['id']`: the async callback id is assigned from the contract. so later it will be passed as argument of `ibc_ack` and `ibc_timeout`.
- `memo['evm']['async_callback']['contract_addr']`: The address of contract which defines the callback function.

### IBC Transfer using ERC20Wrapper

`src -> dst`: Execute the ERC20Wrapper contract to wrap and do ibc-transfer

`dst -> src`: ibc-trasfer and execute the ERC20Wrapper contract via ibc-hook

- data example

```json
{
  //... other ibc fields that we don't care about
  "data": {
    "denom": "wrapped token denom", // will be transformed to the local denom (ibc/...)
    "amount": "1000",
    "sender": "addr on counterparty chain", // will be transformed
    "receiver": "0xcontractaddr",
    "memo": {
      "evm": {
        // execute message on receive packet
        "message": {
          "contract_addr": "0xerc20_wrapper_contract", // should query erc20 wrapper contract addr
          "input": "pack(unwrap, denom, recipient, amount)", // function selector(fc078758) + abiCoder.encode([string,address,address],denom,recipient,amount) ref) https://docs.ethers.org/v6/api/abi/abi-coder/#AbiCoder-encode
          "value": "0",
          "access_list": {
            "address" : "...", // contract address
            "storage_keys":  ["...","..."] // storage keys of contract
          }
        }
      }
    }
  }
}
```
