package wrappers_bson

import (
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw/bsonrwtest"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/durationpb"
	"gotest.tools/v3/assert"
)

func TestDurationCodec_Encode(t *testing.T) {
	elapsed := time.Now().Sub(time.Date(2021, 3, 31, 0, 21, 0, 0, time.UTC))
	type testCase struct {
		name     string
		dur      *durationpb.Duration
		vw       *bsonrwtest.ValueReaderWriter
		expected bsonrwtest.Invoked
	}
	testCases := []testCase{
		{
			"nil",
			nil,
			&bsonrwtest.ValueReaderWriter{BSONType: bson.TypeNull},
			bsonrwtest.WriteNull,
		},
		{
			"int64",
			durationpb.New(elapsed),
			&bsonrwtest.ValueReaderWriter{BSONType: bson.TypeInt64},
			bsonrwtest.WriteInt64,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := NewDurationCodec()
			v := reflect.ValueOf(tc.dur)
			err := c.EncodeValue(bsoncodec.EncodeContext{}, tc.vw, v)
			assert.NilError(t, err)
			assert.DeepEqual(t, tc.expected, tc.vw.Invoked)
		})
	}
}

func TestDurationCodec_Decode(t *testing.T) {
	elapsed := time.Now().Sub(time.Date(2021, 3, 31, 0, 21, 0, 0, time.UTC))
	type testCase struct {
		name     string
		vr       *bsonrwtest.ValueReaderWriter
		expected *durationpb.Duration
	}
	testCases := []testCase{
		{
			"int64",
			&bsonrwtest.ValueReaderWriter{
				BSONType: bson.TypeInt64,
				Return:   int64(elapsed),
			},
			durationpb.New(elapsed),
		},
		{
			"string",
			&bsonrwtest.ValueReaderWriter{
				BSONType: bson.TypeString,
				Return:   elapsed.String(),
			},
			durationpb.New(elapsed),
		},
		{
			"null",
			&bsonrwtest.ValueReaderWriter{
				BSONType: bson.TypeNull,
			},
			nil,
		},
		{
			"undefined",
			&bsonrwtest.ValueReaderWriter{
				BSONType: bson.TypeUndefined,
			},
			&durationpb.Duration{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := NewDurationCodec()
			got := reflect.New(reflect.TypeOf(tc.expected)).Elem()
			err := c.DecodeValue(bsoncodec.DecodeContext{}, tc.vr, got)
			assert.NilError(t, err)
			assert.DeepEqual(t, tc.expected, got.Interface(), protocmp.Transform())
		})
	}
}
