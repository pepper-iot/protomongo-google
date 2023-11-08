package structpb_bson

import (
	"fmt"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"google.golang.org/protobuf/types/known/structpb"
)

var DefaultValueCodec = ValueCodec{}

type ValueCodec struct{}

func (c ValueCodec) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != ProtoValueType {
		return bsoncodec.ValueEncoderError{Name: "ValueCodec.EncodeValue", Types: []reflect.Type{ProtoValueType}, Received: val}
	}

	//val = reflect.Indirect(val)
	kindField := val.FieldByName("Kind") // the 'Kind' field
	if kindField.IsNil() {
		return vw.WriteNull()
	}

	kv := kindField.Elem()

	var xv reflect.Value
	switch kv.Type() {
	case ProtoValueNullPtrType:
		return vw.WriteNull()
	default:
		xv = kv.Elem().Field(0) // the 'XXXValue' field
	}

	encoder, err := ec.LookupEncoder(xv.Type())
	if err != nil {
		return err
	}
	return encoder.EncodeValue(ec, vw, xv)
}

func (c ValueCodec) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.IsValid() || val.Type() != ProtoValueType {
		return bsoncodec.ValueDecoderError{Name: "ValueCodec.DecodeValue", Types: []reflect.Type{ProtoValueType}, Received: val}
	}

	kindField := val.FieldByName("Kind") // the 'Kind' field

	switch vr.Type() {
	case bson.TypeNull:
		kindField.Set(reflect.ValueOf(&structpb.Value_NullValue{}))
		return vr.ReadNull()
	case bson.TypeUndefined:
		kindField.Set(reflect.ValueOf(&structpb.Value_NullValue{}))
		return vr.ReadUndefined()
	case bsontype.Type(0):
		kindField.Set(reflect.ValueOf(&structpb.Value_NullValue{}))
		return nil
	case bson.TypeEmbeddedDocument:
		value := &structpb.Value_StructValue{StructValue: &structpb.Struct{}}
		if err := DefaultStructCodec.DecodeValue(dc, vr, reflect.ValueOf(value.StructValue).Elem()); err != nil {
			return err
		}
		kindField.Set(reflect.ValueOf(value))
	case bson.TypeArray:
		list := &structpb.Value_ListValue{ListValue: &structpb.ListValue{}}
		if err := DefaultListCodec.DecodeValue(dc, vr, reflect.ValueOf(list.ListValue).Elem()); err != nil {
			return err
		}
		kindField.Set(reflect.ValueOf(list))
	case bson.TypeDouble:
		v, err := vr.ReadDouble()
		if err != nil {
			return err
		}
		kindField.Set(reflect.ValueOf(&structpb.Value_NumberValue{NumberValue: v}))
	case bson.TypeInt32:
		v, err := vr.ReadInt32()
		if err != nil {
			return err
		}
		kindField.Set(reflect.ValueOf(&structpb.Value_NumberValue{NumberValue: float64(v)}))
	case bson.TypeInt64:
		v, err := vr.ReadInt64()
		if err != nil {
			return err
		}
		kindField.Set(reflect.ValueOf(&structpb.Value_NumberValue{NumberValue: float64(v)}))
	case bson.TypeString:
		v, err := vr.ReadString()
		if err != nil {
			return err
		}
		kindField.Set(reflect.ValueOf(&structpb.Value_StringValue{StringValue: v}))
	case bson.TypeBoolean:
		v, err := vr.ReadBoolean()
		if err != nil {
			return err
		}
		kindField.Set(reflect.ValueOf(&structpb.Value_BoolValue{BoolValue: v}))
	default:
		return fmt.Errorf("cannot decode %v into a %s", vr.Type(), val.Type())
	}
	return nil
}
