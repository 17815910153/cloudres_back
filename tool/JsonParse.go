package tool

import (
	"encoding/json"
	"io"
)

type JsonParse struct {

}

func Decode(io io.ReadCloser, v interface{}) error  {
	//将json字符串进行反序列化
	return json.NewDecoder(io).Decode(v)

}
