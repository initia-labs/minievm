package jsonutils

import (
	"encoding/json"
	"reflect"

	"github.com/pkg/errors"
)

// MergeJSON recursively merges the src and dst maps. Key conflicts are resolved by
// preferring src, or recursively descending, if both src and dst are maps.
func MergeJSON(dstStr, srcStr string) (string, error) {
	if len(dstStr) == 0 || len(srcStr) == 0 {
		return "", errors.New("empty JSON string")
	}
	if len(dstStr) > MaxJSONSize || len(srcStr) > MaxJSONSize {
		return "", errors.New("JSON string too large")
	}

	var dstMap, srcMap map[string]interface{}
	err := json.Unmarshal([]byte(dstStr), &dstMap)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal([]byte(srcStr), &srcMap)
	if err != nil {
		return "", err
	}

	res, err := merge(dstMap, srcMap, 0)
	if err != nil {
		return "", err
	}

	resBz, err := json.Marshal(res)
	if err != nil {
		return "", err
	}

	return string(resBz), nil
}

func merge(dst, src map[string]interface{}, depth int) (map[string]interface{}, error) {
	if depth > MaxDepth {
		return nil, errors.New("max recursion depth reached")
	}

	for key, srcVal := range src {
		if dstVal, ok := dst[key]; ok {
			srcMap, srcMapOk := mapify(srcVal)
			dstMap, dstMapOk := mapify(dstVal)
			if srcMapOk && dstMapOk {
				var err error
				if srcVal, err = merge(dstMap, srcMap, depth+1); err != nil {
					return nil, err
				}
			}
		}

		dst[key] = srcVal
	}

	return dst, nil
}

func mapify(i interface{}) (map[string]interface{}, bool) {
	value := reflect.ValueOf(i)
	if value.Kind() == reflect.Map {
		m := map[string]interface{}{}
		for _, k := range value.MapKeys() {
			m[k.String()] = value.MapIndex(k).Interface()
		}

		return m, true
	}

	return map[string]interface{}{}, false
}
