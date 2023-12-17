package utils

import (
	"net/http"
	"reflect"
)

// ReturnJsonResponse function for returning movies data in JSON format
func ReturnJsonResponse(res http.ResponseWriter, httpCode int, resMessage []byte) {
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(httpCode)
	res.Write(resMessage)
}

func GetMapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func GetLargestValue(num []int) int {
	var res = 0
	for _, n := range num {
		if n > res {
			res = n
		}
	}
	return res
}

func SetIfPresent[K any](value K, defValue K) K {
	if !reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface()) {
		return value
	}

	return defValue
}
