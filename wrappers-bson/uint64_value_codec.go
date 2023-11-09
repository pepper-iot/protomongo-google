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

var UInt64ValueType = reflect.TypeOf((*wrapperspb.UInt64Value)(nil))

// UInt64ValueCodec is the Codec used for *wrapperspb.UInt64Value values.
type UInt64ValueCodec struct{}

// EncodeValue is the ValueEncoderFunc for *wrapperspb.UInt64Value.
func (c *UInt64ValueCodec) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, v reflect.Value) error {
	if !v.IsValid() || v.Type() != UInt64ValueType {
		return bsoncodec.ValueEncoderError{
			Name:     "UInt64ValueCodec.EncodeValue",
			Types:    []reflect.Type{UInt64ValueType},
			Received: v,
		}
	}
	val := v.Interface().(*wrapperspb.UInt64Value)
	if val == nil {
		return vw.WriteNull()
	}
	return vw.WriteInt64(int64(val.Value))
}

// DecodeValue is the ValueDecoderFunc for *wrapperspb.UInt64Value.
func (c *UInt64ValueCodec) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, v reflect.Value) error {
	if !v.CanSet() || v.Type() != UInt64ValueType {
		return bsoncodec.ValueDecoderError{
			Name:     "UInt64ValueCodec.DecodeValue",
			Types:    []reflect.Type{UInt64ValueType},
			Received: v,
		}
	}
	var val *wrapperspb.UInt64Value
	switch bsonTyp := vr.Type(); bsonTyp {
	case bson.TypeInt64:
		v, err := vr.ReadInt64()
		if err != nil {
			return err
		}
		val = wrapperspb.UInt64(uint64(v))
	case bson.TypeInt32:
		v, err := vr.ReadInt32()
		if err != nil {
			return err
		}
		val = wrapperspb.UInt64(uint64(v))
	case bson.TypeString:
		s, err := vr.ReadString()
		if err != nil {
			return err
		}
		v, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return err
		}
		val = wrapperspb.UInt64(v)
	case bson.TypeNull:
		if err := vr.ReadNull(); err != nil {
			return err
		}
		val = nil
	case bson.TypeUndefined:
		if err := vr.ReadUndefined(); err != nil {
			return err
		}
		val = &wrapperspb.UInt64Value{}
	default:
		return fmt.Errorf("cannot decode %v into a *wrapperspb.UInt64Value", bsonTyp)
	}
	v.Set(reflect.ValueOf(val))
	return nil
}

// NewUInt64ValueCodec returns a UInt64ValueCodec.
func NewUInt64ValueCodec() *UInt64ValueCodec {
	return &UInt64ValueCodec{}
}
