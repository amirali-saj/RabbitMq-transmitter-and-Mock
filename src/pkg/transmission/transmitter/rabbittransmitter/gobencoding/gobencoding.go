package gobencoding

import (
	"bytes"
	"encoding/gob"
)

func Marshal(object interface{}) []byte {
	b := new(bytes.Buffer)

	err := gob.NewEncoder(b).Encode(object)
	if err != nil {
		panic(err)
	}
	return b.Bytes()
}

func Unmarshal(data []byte, object interface{}) error {
	b := bytes.NewBuffer(data)
	err := gob.NewDecoder(b).Decode(object)

	return err
}
