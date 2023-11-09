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

var UInt32ValueType = reflect.TypeOf((*wrapperspb.UInt32Value)(nil))

// UInt32ValueCodec is the Codec used for *wrapperspb.UInt32Value values.
type UInt32ValueCodec struct{}

// EncodeValue is the ValueEncoderFunc for *wrapperspb.UInt32Value.
func (vc *UInt32ValueCodec) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, v reflect.Value) error {
	if !v.IsValid() || v.Type() != UInt32ValueType {
		return bsoncodec.ValueEncoderError{
			Name:     "UInt32ValueCodec.EncodeValue",
			Types:    []reflect.Type{UInt32ValueType},
			Received: v,
		}
	}
	val := v.Interface().(*wrapperspb.UInt32Value)
	if val == nil {
		return vw.WriteNull()
	}
	return vw.WriteInt32(int32(val.Value))
}

// DecodeValue is the ValueDecoderFunc for *wrapperspb.UInt32Value.
func (vc *UInt32ValueCodec) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, v reflect.Value) error {
	if !v.CanSet() || v.Type() != UInt32ValueType {
		return bsoncodec.ValueDecoderError{
			Name:     "UInt32ValueCodec.DecodeValue",
			Types:    []reflect.Type{UInt32ValueType},
			Received: v,
		}
	}
	var val *wrapperspb.UInt32Value
	switch bsonTyp := vr.Type(); bsonTyp {
	case bson.TypeInt32:
		v, err := vr.ReadInt32()
		if err != nil {
			return err
		}
		val = wrapperspb.UInt32(uint32(v))
	case bson.TypeString:
		s, err := vr.ReadString()
		if err != nil {
			return err
		}
		v, err := strconv.ParseUint(s, 10, 32)
		if err != nil {
			return err
		}
		val = wrapperspb.UInt32(uint32(v))
	case bson.TypeNull:
		if err := vr.ReadNull(); err != nil {
			return err
		}
		val = nil
	case bson.TypeUndefined:
		if err := vr.ReadUndefined(); err != nil {
			return err
		}
		val = &wrapperspb.UInt32Value{}
	default:
		return fmt.Errorf("cannot decode %v into a *wrapperspb.UInt32Value", bsonTyp)
	}
	v.Set(reflect.ValueOf(val))
	return nil
}

// NewUInt32ValueCodec returns a UInt32ValueCodec.
func NewUInt32ValueCodec() *UInt32ValueCodec {
	return &UInt32ValueCodec{}
}
