package wrappers_bson

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"google.golang.org/protobuf/types/known/durationpb"
)

var DurationType = reflect.TypeOf((*durationpb.Duration)(nil))

// DurationCodec is the Codec used for *durationpb.Duration values.
type DurationCodec struct{}

// EncodeValue is the ValueEncoderFunc for *durationpb.Duration.
func (c *DurationCodec) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, v reflect.Value) error {
	if !v.IsValid() || v.Type() != DurationType {
		return bsoncodec.ValueEncoderError{
			Name:     "DurationCodec.EncodeValue",
			Types:    []reflect.Type{DurationType},
			Received: v,
		}
	}
	dur := v.Interface().(*durationpb.Duration)
	if dur == nil {
		return vw.WriteNull()
	}
	return vw.WriteInt64(int64(dur.AsDuration()))
}

// DecodeValue is the ValueDecoderFunc for *durationpb.Duration.
func (c *DurationCodec) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, v reflect.Value) error {
	if !v.CanSet() || v.Type() != DurationType {
		return bsoncodec.ValueDecoderError{
			Name:     "DurationCodec.DecodeValue",
			Types:    []reflect.Type{DurationType},
			Received: v,
		}
	}
	var dur *durationpb.Duration
	switch bsonTyp := vr.Type(); bsonTyp {
	case bson.TypeInt64:
		nsec, err := vr.ReadInt64()
		if err != nil {
			return err
		}
		dur = durationpb.New(time.Duration(nsec))
	case bson.TypeString:
		s, err := vr.ReadString()
		if err != nil {
			return err
		}
		d, err := time.ParseDuration(s)
		if err != nil {
			return err
		}
		dur = durationpb.New(d)
	case bson.TypeNull:
		if err := vr.ReadNull(); err != nil {
			return err
		}
		dur = nil
	case bson.TypeUndefined:
		if err := vr.ReadUndefined(); err != nil {
			return err
		}
		dur = &durationpb.Duration{}
	default:
		return fmt.Errorf("cannot decode %v into a *durationpb.Duration", bsonTyp)
	}
	v.Set(reflect.ValueOf(dur))
	return nil
}

// NewDurationCodec returns a DurationCodec.
func NewDurationCodec() *DurationCodec {
	return &DurationCodec{}
}
