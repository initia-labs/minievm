// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "../i_cosmos/ICosmos.sol";
import "../i_jsonutils/IJSONUtils.sol";

contract ConnectOracle {
    struct Price {
        uint256 price;
        uint256 timestamp;
        uint64 height;
        uint64 nonce;
        uint64 decimal;
        uint64 id;
    }

    constructor() {}

    function get_price(string memory pair_id) external returns (Price memory) {
        string memory path = "/connect.oracle.v2.Query/GetPrice";
        string memory queryRes = COSMOS_CONTRACT.query_cosmos(path, string(
            abi.encodePacked(
                '{"currency_pair": ',
                JSONUTILS_CONTRACT.stringify_json(pair_id),
                '}'
            )
        ));

        IJSONUtils.JSONObject memory obj = JSONUTILS_CONTRACT
            .unmarshal_to_object(bytes(queryRes));

        return unmarshal_price(obj);
    }

    function get_prices(
        string[] memory pair_ids
    ) external returns (Price[] memory) {
        string memory path = "/connect.oracle.v2.Query/GetPrices";
        string memory req = string(
            abi.encodePacked(
                '{"currency_pair_ids":',
                marshal_string_array(pair_ids),
                '}'
            )
        );

        string memory queryRes = COSMOS_CONTRACT.query_cosmos(path, req);
        IJSONUtils.JSONObject memory prices_obj = JSONUTILS_CONTRACT
            .unmarshal_to_object(bytes(queryRes));
        bytes[] memory prices = JSONUTILS_CONTRACT.unmarshal_to_array(
            prices_obj.elements[0].value
        );

        Price[] memory response = new Price[](prices.length);
        for (uint256 i = 0; i < prices.length; i++) {
            response[i] = unmarshal_price(
                JSONUTILS_CONTRACT.unmarshal_to_object(prices[i])
            );
        }

        return response;
    }

    function marshal_string_array(
        string[] memory strArray
    ) internal view returns (string memory) {
        uint256 len = strArray.length;
        string memory res = string.concat('[', JSONUTILS_CONTRACT.stringify_json(strArray[0]));
        for (uint256 i = 1; i < len; i++) {
            res = string.concat(res, ",");
            res = string.concat(res, JSONUTILS_CONTRACT.stringify_json(strArray[i]));
        }
        res = string.concat(res, ']');

        return res;
    }

    function unmarshal_price(
        IJSONUtils.JSONObject memory obj
    ) internal view returns (Price memory) {
        IJSONUtils.JSONObject memory price_obj = JSONUTILS_CONTRACT
            .unmarshal_to_object(obj.elements[3].value);
        uint256 price = JSONUTILS_CONTRACT.unmarshal_to_uint(
            price_obj.elements[2].value
        );
        uint256 timestamp = JSONUTILS_CONTRACT.unmarshal_iso_to_unix(
            price_obj.elements[1].value
        );
        uint64 height = uint64(
            JSONUTILS_CONTRACT.unmarshal_to_uint(price_obj.elements[0].value)
        );
        uint64 nonce = uint64(
            JSONUTILS_CONTRACT.unmarshal_to_uint(obj.elements[2].value)
        );
        uint64 decimal = uint64(
            JSONUTILS_CONTRACT.unmarshal_to_uint(obj.elements[0].value)
        );
        uint64 id = uint64(
            JSONUTILS_CONTRACT.unmarshal_to_uint(obj.elements[1].value)
        );

        return
            Price({
                price: price,
                timestamp: timestamp,
                height: height,
                nonce: nonce,
                decimal: decimal,
                id: id
            });
    }
}
