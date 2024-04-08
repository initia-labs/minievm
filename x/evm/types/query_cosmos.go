package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	proto "github.com/cosmos/gogoproto/proto"

	oracletypes "github.com/skip-mev/slinky/x/oracle/types"
)

type ProtoSet struct {
	Request  proto.Message
	Response proto.Message
}

type QueryCosmosWhitelist map[string]ProtoSet

func DefaultQueryCosmosWhitelist() QueryCosmosWhitelist {
	res := make(QueryCosmosWhitelist)
	res["/slinky.oracle.v1.Query/GetPrices"] = ProtoSet{
		Request:  &oracletypes.GetPricesRequest{},
		Response: &oracletypes.GetPricesResponse{},
	}
	res["/slinky.oracle.v1.Query/GetPrice"] = ProtoSet{
		Request:  &oracletypes.GetPriceRequest{},
		Response: &oracletypes.GetPriceResponse{},
	}
	res["/slinky.oracle.v1.Query/GetAllCurrencyPairs"] = ProtoSet{
		Request:  &oracletypes.GetAllCurrencyPairsRequest{},
		Response: &oracletypes.GetAllCurrencyPairsResponse{},
	}

	return res
}

// ConvertProtoToJSON unmarshal the given bytes into a proto message and then marshals it to json.
func ConvertProtoToJSON(cdc codec.Codec, protoResponse proto.Message, bz []byte) ([]byte, error) {
	// unmarshal binary into stargate response data structure
	err := cdc.Unmarshal(bz, protoResponse)
	if err != nil {
		return nil, err
	}

	bz, err = cdc.MarshalJSON(protoResponse)
	if err != nil {
		return nil, err
	}

	protoResponse.Reset()
	return bz, nil
}

func ConvertJSONToProto(cdc codec.Codec, protoRequest proto.Message, bz []byte) ([]byte, error) {
	// unmarshal binary into stargate response data structure
	err := cdc.UnmarshalJSON(bz, protoRequest)
	if err != nil {
		return nil, err
	}

	bz, err = cdc.Marshal(protoRequest)
	if err != nil {
		return nil, err
	}

	protoRequest.Reset()
	return bz, nil
}
