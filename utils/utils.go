package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

func HandleErr(err error){
	if err != nil { log.Panic(err)}
}

func ToBytes(i interface{}) []byte {
	var aBuffer bytes.Buffer
	
	encoder := gob.NewEncoder(&aBuffer)
	HandleErr(encoder.Encode(i))

	return aBuffer.Bytes()
}

func FromByte(i interface{}, data []byte){
	decoder := gob.NewDecoder(bytes.NewReader(data))
	HandleErr(decoder.Decode(i))
}

func Hash(i interface{}) string {
	s := fmt.Sprintf("%v", i)
	return fmt.Sprintf("%x",  sha256.Sum256([]byte(s)))
}