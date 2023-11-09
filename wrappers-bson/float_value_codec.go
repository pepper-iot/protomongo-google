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

var FloatValueType = reflect.TypeOf((*wrapperspb.FloatValue)(nil))

// FloatValueCodec is the Codec used for *wrapperspb.FloatValue values.
type FloatValueCodec struct{}

// EncodeValue is the ValueEncoderFunc for *wrapperspb.FloatValue.
func (c *FloatValueCodec) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, v reflect.Value) error {
	if !v.IsValid() || v.Type() != FloatValueType {
		return bsoncodec.ValueEncoderError{
			Name:     "FloatValueCodec.EncodeValue",
			Types:    []reflect.Type{FloatValueType},
			Received: v,
		}
	}
	val := v.Interface().(*wrapperspb.FloatValue)
	if val == nil {
		return vw.WriteNull()
	}
	return vw.WriteDouble(float64(val.Value))
}

// DecodeValue is the ValueDecoderFunc for *wrapperspb.FloatValue.
func (c *FloatValueCodec) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, v reflect.Value) error {
	if !v.CanSet() || v.Type() != FloatValueType {
		return bsoncodec.ValueDecoderError{
			Name:     "FloatValueCodec.DecodeValue",
			Types:    []reflect.Type{FloatValueType},
			Received: v,
		}
	}
	var val *wrapperspb.FloatValue
	switch bsonTyp := vr.Type(); bsonTyp {
	case bson.TypeDouble:
		v, err := vr.ReadDouble()
		if err != nil {
			return err
		}
		val = wrapperspb.Float(float32(v))
	case bson.TypeString:
		s, err := vr.ReadString()
		if err != nil {
			return err
		}
		v, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return err
		}
		val = wrapperspb.Float(float32(v))
	case bson.TypeNull:
		if err := vr.ReadNull(); err != nil {
			return err
		}
		val = nil
	case bson.TypeUndefined:
		if err := vr.ReadUndefined(); err != nil {
			return err
		}
		val = &wrapperspb.FloatValue{}
	default:
		return fmt.Errorf("cannot decode %v into a *wrapperspb.FloatValue", bsonTyp)
	}
	v.Set(reflect.ValueOf(val))
	return nil
}

// NewFloatValueCodec returns a FloatValueCodec.
func NewFloatValueCodec() *FloatValueCodec {
	return &FloatValueCodec{}
}
