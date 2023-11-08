package wrappers_bson

import (
	"github.com/golang/protobuf/ptypes/wrappers"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"reflect"
)

var (
	boolValueType   = reflect.TypeOf(wrappers.BoolValue{})
	bytesValueType  = reflect.TypeOf(wrappers.BytesValue{})
	doubleValueType = reflect.TypeOf(wrappers.DoubleValue{})
	floatValueType  = reflect.TypeOf(wrappers.FloatValue{})
	int32ValueType  = reflect.TypeOf(wrappers.Int32Value{})
	int64ValueType  = reflect.TypeOf(wrappers.Int64Value{})
	stringValueType = reflect.TypeOf(wrappers.StringValue{})
	uint32ValueType = reflect.TypeOf(wrappers.UInt32Value{})
	uint64ValueType = reflect.TypeOf(wrappers.UInt64Value{})
)

func RegisterRegistry(rb *bsoncodec.Registry) *bsoncodec.Registry {

	// Decoders
	rb.RegisterTypeDecoder(boolValueType, &WrapperValueCodec{})
	rb.RegisterTypeDecoder(bytesValueType, &WrapperValueCodec{})
	rb.RegisterTypeDecoder(doubleValueType, &WrapperValueCodec{})
	rb.RegisterTypeDecoder(floatValueType, &WrapperValueCodec{})
	rb.RegisterTypeDecoder(int32ValueType, &WrapperValueCodec{})
	rb.RegisterTypeDecoder(int64ValueType, &WrapperValueCodec{})
	rb.RegisterTypeDecoder(stringValueType, &WrapperValueCodec{})
	rb.RegisterTypeDecoder(uint32ValueType, &WrapperValueCodec{})
	rb.RegisterTypeDecoder(uint64ValueType, &WrapperValueCodec{})

	// Encoders
	rb.RegisterTypeEncoder(boolValueType, &WrapperValueCodec{})
	rb.RegisterTypeEncoder(bytesValueType, &WrapperValueCodec{})
	rb.RegisterTypeEncoder(doubleValueType, &WrapperValueCodec{})
	rb.RegisterTypeEncoder(floatValueType, &WrapperValueCodec{})
	rb.RegisterTypeEncoder(int32ValueType, &WrapperValueCodec{})
	rb.RegisterTypeEncoder(int64ValueType, &WrapperValueCodec{})
	rb.RegisterTypeEncoder(stringValueType, &WrapperValueCodec{})
	rb.RegisterTypeEncoder(uint32ValueType, &WrapperValueCodec{})
	rb.RegisterTypeEncoder(uint64ValueType, &WrapperValueCodec{})

	return rb

}
