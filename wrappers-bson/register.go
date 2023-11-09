package wrappers_bson

import (
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
)

func RegisterRegistry(rb *bsoncodec.Registry) *bsoncodec.Registry {

	// Decoders
	rb.RegisterTypeDecoder(BoolValueType, NewBoolValueCodec())
	rb.RegisterTypeDecoder(BytesValueType, NewBytesValueCodec())
	rb.RegisterTypeDecoder(DoubleValueType, NewDoubleValueCodec())
	rb.RegisterTypeDecoder(DurationType, NewDurationCodec())
	rb.RegisterTypeDecoder(FloatValueType, NewFloatValueCodec())
	rb.RegisterTypeDecoder(Int32ValueType, NewInt32ValueCodec())
	rb.RegisterTypeDecoder(Int64ValueType, NewInt64ValueCodec())
	rb.RegisterTypeDecoder(StringValueType, NewStringValueCodec())
	rb.RegisterTypeDecoder(UInt32ValueType, NewUInt32ValueCodec())
	rb.RegisterTypeDecoder(UInt64ValueType, NewUInt64ValueCodec())

	// Encoders
	rb.RegisterTypeEncoder(BoolValueType, NewBoolValueCodec())
	rb.RegisterTypeEncoder(BytesValueType, NewBytesValueCodec())
	rb.RegisterTypeEncoder(DoubleValueType, NewDoubleValueCodec())
	rb.RegisterTypeEncoder(DurationType, NewDurationCodec())
	rb.RegisterTypeEncoder(FloatValueType, NewFloatValueCodec())
	rb.RegisterTypeEncoder(Int32ValueType, NewInt32ValueCodec())
	rb.RegisterTypeEncoder(Int64ValueType, NewInt64ValueCodec())
	rb.RegisterTypeEncoder(StringValueType, NewStringValueCodec())
	rb.RegisterTypeEncoder(UInt32ValueType, NewUInt32ValueCodec())
	rb.RegisterTypeEncoder(UInt64ValueType, NewUInt64ValueCodec())

	return rb

}
