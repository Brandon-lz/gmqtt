package utils

import (
	"encoding/json"
	"fmt"
	"log/slog"
)

func DeserializeData[T interface{}](source any, target *T) T { // target必须为指针类型
    var jsonData []byte
    var err error

    if sourceString, isString := source.(string); isString {
        jsonData = []byte(sourceString)
    } else if sourceBytes, isBytes := source.([]byte); isBytes {
        jsonData = sourceBytes
    } else {
        jsonData, err = json.Marshal(source)
        if err != nil {
			slog.Error(fmt.Sprintf("JSON序列化失败: %s", WrapErrorLocation(err)))
            panic(err)
        }
    }

    err = json.Unmarshal(jsonData, target)
    if err != nil {
		slog.Error(fmt.Sprintf("JSON反序列化失败: %s", WrapErrorLocation(err)))
        panic(err)
    }
    return *target
}
