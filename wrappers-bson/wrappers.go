package wrappers_bson

import (
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"reflect"
)

// WrapperValueCodec is codec for Protobuf type wrappers
type WrapperValueCodec struct{}

// EncodeValue encodes Protobuf type wrapper value to BSON value
func (e *WrapperValueCodec) EncodeValue(ectx bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	val = val.FieldByName("Value")
	enc, err := ectx.LookupEncoder(val.Type())
	if err != nil {
		return err
	}
	return enc.EncodeValue(ectx, vw, val)
}

// DecodeValue decodes BSON value to Protobuf type wrapper value
func (e *WrapperValueCodec) DecodeValue(ectx bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	val = val.FieldByName("Value")
	enc, err := ectx.LookupDecoder(val.Type())
	if err != nil {
		return err
	}
	return enc.DecodeValue(ectx, vr, val)
}
