package wrappers_bson

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var BoolValueType = reflect.TypeOf((*wrapperspb.BoolValue)(nil))

// BoolValueCodec is the Codec used for *wrapperspb.BoolValue values.
type BoolValueCodec struct{}

// EncodeValue is the ValueEncoderFunc for *wrapperspb.BoolValue.
func (c *BoolValueCodec) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, v reflect.Value) error {
	if !v.IsValid() || v.Type() != BoolValueType {
		return bsoncodec.ValueEncoderError{
			Name:     "BoolValueCodec.EncodeValue",
			Types:    []reflect.Type{BoolValueType},
			Received: v,
		}
	}
	val := v.Interface().(*wrapperspb.BoolValue)
	if val == nil {
		return vw.WriteNull()
	}
	return vw.WriteBoolean(val.Value)
}

// DecodeValue is the ValueDecoderFunc for *wrapperspb.BoolValue.
func (c *BoolValueCodec) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, v reflect.Value) error {
	if !v.CanSet() || v.Type() != BoolValueType {
		return bsoncodec.ValueDecoderError{
			Name:     "BoolValueCodec.DecodeValue",
			Types:    []reflect.Type{BoolValueType},
			Received: v,
		}
	}
	var val *wrapperspb.BoolValue
	switch bsonTyp := vr.Type(); bsonTyp {
	case bson.TypeBoolean:
		v, err := vr.ReadBoolean()
		if err != nil {
			return err
		}
		val = wrapperspb.Bool(v)
	case bson.TypeString:
		s, err := vr.ReadString()
		if err != nil {
			return err
		}
		v, err := strconv.ParseBool(s)
		if err != nil {
			return err
		}
		val = wrapperspb.Bool(v)
	case bson.TypeNull:
		if err := vr.ReadNull(); err != nil {
			return err
		}
		val = nil
	case bson.TypeUndefined:
		if err := vr.ReadUndefined(); err != nil {
			return err
		}
		val = &wrapperspb.BoolValue{}
	default:
		return fmt.Errorf("cannot decode %v into a *wrapperspb.BoolValue", bsonTyp)
	}
	v.Set(reflect.ValueOf(val))
	return nil
}

// NewBoolValueCodec returns a BoolValueCodec.
func NewBoolValueCodec() *BoolValueCodec {
	return &BoolValueCodec{}
}
