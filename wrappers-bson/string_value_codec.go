package wrappers_bson

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var StringValueType = reflect.TypeOf((*wrapperspb.StringValue)(nil))

// StringValueCodec is the Codec used for *wrapperspb.StringValue values.
type StringValueCodec struct{}

// EncodeValue is the ValueEncoderFunc for *wrapperspb.StringValue.
func (c *StringValueCodec) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, v reflect.Value) error {
	if !v.IsValid() || v.Type() != StringValueType {
		return bsoncodec.ValueEncoderError{
			Name:     "StringValueCodec.EncodeValue",
			Types:    []reflect.Type{StringValueType},
			Received: v,
		}
	}
	val := v.Interface().(*wrapperspb.StringValue)
	if val == nil {
		return vw.WriteNull()
	}
	return vw.WriteString(val.Value)
}

// DecodeValue is the ValueDecoderFunc for *wrapperspb.StringValue.
func (c *StringValueCodec) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, v reflect.Value) error {
	if !v.CanSet() || v.Type() != StringValueType {
		return bsoncodec.ValueDecoderError{
			Name:     "StringValueCodec.DecodeValue",
			Types:    []reflect.Type{StringValueType},
			Received: v,
		}
	}
	var val *wrapperspb.StringValue
	switch bsonTyp := vr.Type(); bsonTyp {
	case bson.TypeString:
		v, err := vr.ReadString()
		if err != nil {
			return err
		}
		val = wrapperspb.String(v)
	case bson.TypeBinary:
		v, _, err := vr.ReadBinary()
		if err != nil {
			return err
		}
		val = wrapperspb.String(string(v))
	case bson.TypeNull:
		if err := vr.ReadNull(); err != nil {
			return err
		}
		val = nil
	case bson.TypeUndefined:
		if err := vr.ReadUndefined(); err != nil {
			return err
		}
		val = &wrapperspb.StringValue{}
	default:
		return fmt.Errorf("cannot decode %v into a *wrapperspb.StringValue", bsonTyp)
	}
	v.Set(reflect.ValueOf(val))
	return nil
}

// NewStringValueCodec returns a StringValueCodec.
func NewStringValueCodec() *StringValueCodec {
	return &StringValueCodec{}
}
