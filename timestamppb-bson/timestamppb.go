package timestamppb_bson

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"google.golang.org/protobuf/types/known/timestamppb"
	"reflect"
	"time"
)

var (
	TimestampType = reflect.TypeOf((*timestamppb.Timestamp)(nil))
)

// TimestampCodec is codec for Protobuf Timestamp
type TimestampCodec struct{}

// EncodeValue is the ValueEncoderFunc for *timestamppb.Timestamp.
func (c *TimestampCodec) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, v reflect.Value) error {
	if !v.IsValid() || v.Type() != TimestampType {
		return bsoncodec.ValueEncoderError{
			Name:     "TimestampCodec.EncodeValue",
			Types:    []reflect.Type{TimestampType},
			Received: v,
		}
	}
	ts := v.Interface().(*timestamp.Timestamp)
	if ts == nil {
		return vw.WriteNull()
	}
	return vw.WriteDateTime(ts.AsTime().UnixMilli())
}

// DecodeValue is the ValueDecoderFunc for *timestamppb.Timestamp.
func (c *TimestampCodec) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, v reflect.Value) error {
	if !v.CanSet() || v.Type() != TimestampType {
		return bsoncodec.ValueDecoderError{
			Name:     "TimestampCodec.DecodeValue",
			Types:    []reflect.Type{TimestampType},
			Received: v,
		}
	}
	var ts *timestamppb.Timestamp
	switch bsonTyp := vr.Type(); bsonTyp {
	case bson.TypeDateTime:
		msec, err := vr.ReadDateTime()
		if err != nil {
			return err
		}
		ts = timestamppb.New(time.UnixMilli(msec))
	case bson.TypeInt64:
		msec, err := vr.ReadInt64()
		if err != nil {
			return err
		}
		ts = timestamppb.New(time.UnixMilli(msec))
	case bson.TypeString:
		s, err := vr.ReadString()
		if err != nil {
			return err
		}
		t, err := time.Parse(time.RFC3339Nano, s)
		if err != nil {
			return err
		}
		ts = timestamppb.New(t)
	case bson.TypeNull:
		if err := vr.ReadNull(); err != nil {
			return err
		}
		ts = nil
	case bson.TypeUndefined:
		if err := vr.ReadUndefined(); err != nil {
			return err
		}
		ts = &timestamppb.Timestamp{}
	default:
		return fmt.Errorf("cannot decode %v into a *timestamppb.Timestamp", bsonTyp)
	}
	v.Set(reflect.ValueOf(ts))
	return nil
}
