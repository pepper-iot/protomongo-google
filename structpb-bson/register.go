package structpb_bson

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"google.golang.org/protobuf/types/known/structpb"
	"reflect"
)

var (
	ProtoStructType       = reflect.TypeOf(structpb.Struct{})
	ProtoValueType        = reflect.TypeOf(structpb.Value{})
	ProtoListValueType    = reflect.TypeOf(structpb.ListValue{})
	ProtoValueStructType  = reflect.TypeOf(structpb.Value_StructValue{})
	ProtoValueNumberType  = reflect.TypeOf(structpb.Value_NumberValue{})
	ProtoValueStringType  = reflect.TypeOf(structpb.Value_StringValue{})
	ProtoValueBoolType    = reflect.TypeOf(structpb.Value_BoolValue{})
	ProtoValueNullType    = reflect.TypeOf(structpb.Value_NullValue{})
	ProtoValueListType    = reflect.TypeOf(structpb.Value_ListValue{})
	ProtoValueNullPtrType = reflect.TypeOf(&structpb.Value_NullValue{})
	nullValueType         = reflect.TypeOf(bson.TypeNull)
	undefinedValueType    = reflect.TypeOf(bson.TypeUndefined)
)

func RegisterRegistry(rb *bsoncodec.Registry) *bsoncodec.Registry {

	// Decoders
	rb.RegisterTypeDecoder(ProtoStructType, StructCodec{})
	rb.RegisterTypeDecoder(ProtoValueType, ValueCodec{})
	rb.RegisterTypeDecoder(ProtoListValueType, ListCodec{})
	rb.RegisterTypeDecoder(ProtoValueNullType, ValueCodec{})
	rb.RegisterTypeDecoder(nullValueType, ValueCodec{})
	rb.RegisterTypeDecoder(undefinedValueType, ValueCodec{})

	// Encoders
	rb.RegisterTypeEncoder(ProtoStructType, StructCodec{})
	rb.RegisterTypeEncoder(ProtoValueType, ValueCodec{})
	rb.RegisterTypeEncoder(ProtoListValueType, ListCodec{})
	rb.RegisterTypeEncoder(ProtoValueNullType, ValueCodec{})
	rb.RegisterTypeEncoder(nullValueType, ValueCodec{})
	rb.RegisterTypeEncoder(undefinedValueType, ValueCodec{})

	return rb

}
