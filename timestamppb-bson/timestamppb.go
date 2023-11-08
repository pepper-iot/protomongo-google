package timestamppb_bson

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"reflect"
	"time"
)

var (
	// Time type
	timeType = reflect.TypeOf(time.Time{})
	// Protobuf Timestamp type
	TimestampType = reflect.TypeOf(timestamp.Timestamp{})
)

// TimestampCodec is codec for Protobuf Timestamp
type TimestampCodec struct{}

// EncodeValue encodes Protobuf Timestamp value to BSON value
func (e *TimestampCodec) EncodeValue(ectx bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	v := val.Interface().(timestamp.Timestamp)
	t, err := ptypes.Timestamp(&v)
	if err != nil {
		return err
	}
	enc, err := ectx.LookupEncoder(timeType)
	if err != nil {
		return err
	}
	return enc.EncodeValue(ectx, vw, reflect.ValueOf(t.In(time.UTC)))
}

// DecodeValue decodes BSON value to Timestamp value
func (e *TimestampCodec) DecodeValue(ectx bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	enc, err := ectx.LookupDecoder(timeType)
	if err != nil {
		return err
	}
	var t time.Time
	if err = enc.DecodeValue(ectx, vr, reflect.ValueOf(&t).Elem()); err != nil {
		return err
	}
	ts, err := ptypes.TimestampProto(t.In(time.UTC))
	if err != nil {
		return err
	}
	val.Set(reflect.ValueOf(*ts))
	return nil
}
