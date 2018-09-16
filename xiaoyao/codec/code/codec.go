package code

import (
	"bytes"

	"github.com/cuixin/go/codec"
)

var h = new(codec.JsonHandle)

func Encode(obj interface{}) ([]byte, error) {
	buf := make([]byte, 0, 2048)
	writer := bytes.NewBuffer(buf)
	var encoder = codec.NewEncoder(writer, h)
	err := encoder.Encode(obj)
	return writer.Bytes(), err
	//var result bytes.Buffer

	//encoder := gob.NewEncoder(&result)

	//err := encoder.Encode(obj)
	///if err != nil {
	//	log.Panic(err)
	//}

	//return result.Bytes()
}
func Decoder(data []byte, v interface{}) error {
	reader := bytes.NewBuffer(data)
	decoder := codec.NewDecoder(reader, h)
	err := decoder.Decode(v)
	return err
}
