# IBC-hooks

This module is copied from [osmosis](https://github.com/osmosis-labs/osmosis) and changed to execute evm contract with ICS-20 token transfer calls.

## EVM Hooks

The evm hook is an IBC middleware which is used to allow ICS-20 token transfers to initiate contract calls.
This allows cross-chain contract calls, that involve token evmment.
This is useful for a variety of usecases.
One of primary importance is cross-chain swaps, which is an extremely powerful primitive.

The mechanism enabling this is a `memo` field on every ICS20 transfer packet as of [IBC v3.4.0](https://medium.com/the-interchain-foundation/moving-beyond-simple-token-transfers-d42b2b1dc29b).
EVM hooks is an IBC middleware that parses an ICS20 transfer, and if the `memo` field is of a particular form, executes a evm contract call. We now detail the `memo` format for `evm` contract calls, and the execution guarantees provided.

### EVM Contract Execution Format

Before we dive into the IBC metadata format, we show the evm execute message format, so the reader has a sense of what are the fields we need to be setting in.

```go
// MsgCall is a message to call an Ethereum contract.
type MsgCall struct {
 // Sender is the that actor that signed the messages
 Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
 // ContractAddr is the contract address to be executed.
 // It can be cosmos address or hex encoded address.
 ContractAddr string `protobuf:"bytes,2,opt,name=contract_addr,json=contractAddr,proto3" json:"contract_addr,omitempty"`
 // Execution input bytes.
 Input []byte `protobuf:"bytes,3,opt,name=input,proto3" json:"input,omitempty"`
}
```

So we detail where we want to get each of these fields from:

- Sender: We cannot trust the sender of an IBC packet, the counter-party chain has full ability to lie about it.
  We cannot risk this sender being confused for a particular user or module address on Initia.
  So we replace the sender with an account to represent the sender prefixed by the channel and a evm module prefix.
  This is done by setting the sender to `Bech32(Hash(Hash("ibc-hook-intermediary") + channelID/sender))`, where the channelId is the channel id on the local chain.
- ModuleAddress: This field should be directly obtained from the ICS-20 packet metadata
- ModuleName: This field should be directly obtained from the ICS-20 packet metadata
- FunctionName: This field should be directly obtained from the ICS-20 packet metadata
- TypeArgs: This field should be directly obtained from the ICS-20 packet metadata
- Args: This field should be directly obtained from the ICS-20 packet metadata.

So our constructed evm message that we execute will look like:

```go
msg := MsgCall{
 // Sender is the that actor that signed the messages
 Sender: "init1-hash-of-channel-and-sender",
 // ContractAddr is the address of the contract
 ContractAddr: packet.data.memo["evm"]["contract_addr"],
    // Input is the input bytes
 Input: packet.data.memo["evm"]["input"],
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
        "contract_addr": "0x1",
        "input": "base64 encoded bytes",
      }
    }
  }
}
```

An ICS20 packet is formatted correctly for evmhooks iff the following all hold:

- `memo` is not blank
- `memo` is valid JSON
- `memo` has at least one key, with value `"evm"`
- `memo["evm"]` has exactly two entries, `"contract_addr"` and `"input"`
- `receiver` == "" || `receiver` == "contract_addr"

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

# Testing strategy

See go tests.
