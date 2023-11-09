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

var DoubleValueType = reflect.TypeOf((*wrapperspb.DoubleValue)(nil))

// DoubleValueCodec is the Codec used for *wrapperspb.DoubleValue values.
type DoubleValueCodec struct{}

// EncodeValue is the ValueEncoderFunc for *wrapperspb.DoubleValue.
func (c *DoubleValueCodec) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, v reflect.Value) error {
	if !v.IsValid() || v.Type() != DoubleValueType {
		return bsoncodec.ValueEncoderError{
			Name:     "DoubleValueCodec.EncodeValue",
			Types:    []reflect.Type{DoubleValueType},
			Received: v,
		}
	}
	val := v.Interface().(*wrapperspb.DoubleValue)
	if val == nil {
		return vw.WriteNull()
	}
	return vw.WriteDouble(val.Value)
}

// DecodeValue is the ValueDecoderFunc for *wrapperspb.DoubleValue.
func (c *DoubleValueCodec) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, v reflect.Value) error {
	if !v.CanSet() || v.Type() != DoubleValueType {
		return bsoncodec.ValueDecoderError{
			Name:     "DoubleValueCodec.DecodeValue",
			Types:    []reflect.Type{DoubleValueType},
			Received: v,
		}
	}
	var val *wrapperspb.DoubleValue
	switch bsonTyp := vr.Type(); bsonTyp {
	case bson.TypeDouble:
		v, err := vr.ReadDouble()
		if err != nil {
			return err
		}
		val = wrapperspb.Double(v)
	case bson.TypeString:
		s, err := vr.ReadString()
		if err != nil {
			return err
		}
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}
		val = wrapperspb.Double(v)
	case bson.TypeNull:
		if err := vr.ReadNull(); err != nil {
			return err
		}
		val = nil
	case bson.TypeUndefined:
		if err := vr.ReadUndefined(); err != nil {
			return err
		}
		val = &wrapperspb.DoubleValue{}
	default:
		return fmt.Errorf("cannot decode %v into a *wrapperspb.DoubleValue", bsonTyp)
	}
	v.Set(reflect.ValueOf(val))
	return nil
}

// NewDoubleValueCodec returns a DoubleValueCodec.
func NewDoubleValueCodec() *DoubleValueCodec {
	return &DoubleValueCodec{}
}
