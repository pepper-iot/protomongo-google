package timestamppb_bson

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw/bsonrwtest"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"
	"reflect"
	"testing"
	"time"
)

func Test_TimestampCodecEncodeValue(t *testing.T) {
	type testCase struct {
		name     string
		ts       *timestamppb.Timestamp
		vw       *bsonrwtest.ValueReaderWriter
		expected bsonrwtest.Invoked
	}

	testCases := []testCase{
		{
			name:     "nil",
			ts:       nil,
			vw:       &bsonrwtest.ValueReaderWriter{BSONType: bson.TypeNull},
			expected: bsonrwtest.WriteNull,
		},
		{
			name:     "now",
			ts:       timestamppb.Now(),
			vw:       &bsonrwtest.ValueReaderWriter{BSONType: bson.TypeDateTime},
			expected: bsonrwtest.WriteDateTime,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := NewTimestampCodec()
			v := reflect.ValueOf(tc.ts)
			err := c.EncodeValue(bsoncodec.EncodeContext{}, tc.vw, v)
			assert.NilError(t, err)
			assert.DeepEqual(t, tc.expected, tc.vw.Invoked)
		})
	}
}

func Test_TimestampCodecDecodeValue(t *testing.T) {
	type testCase struct {
		name     string
		vr       *bsonrwtest.ValueReaderWriter
		expected *timestamppb.Timestamp
	}

	ts := timestamppb.New(time.Date(2021, 3, 31, 0, 21, 0, 0, time.UTC))
	testCases := []testCase{
		{
			name: "datetime",
			vr: &bsonrwtest.ValueReaderWriter{
				BSONType: bson.TypeDateTime,
				Return:   int64(1617150060000),
			},
			expected: ts,
		},
		{
			name: "int64",
			vr: &bsonrwtest.ValueReaderWriter{
				BSONType: bson.TypeInt64,
				Return:   int64(1617150060000),
			},
			expected: ts,
		},
		{
			name: "string",
			vr: &bsonrwtest.ValueReaderWriter{
				BSONType: bson.TypeString,
				Return:   "2021-03-31T00:21:00Z",
			},
			expected: ts,
		},
		{
			name: "null",
			vr: &bsonrwtest.ValueReaderWriter{
				BSONType: bson.TypeNull,
			},
			expected: nil,
		},
		{
			name: "undefined",
			vr: &bsonrwtest.ValueReaderWriter{
				BSONType: bson.TypeUndefined,
			},
			expected: &timestamppb.Timestamp{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := NewTimestampCodec()
			got := reflect.New(reflect.TypeOf(tc.expected)).Elem()
			err := c.DecodeValue(bsoncodec.DecodeContext{}, tc.vr, got)
			assert.NilError(t, err)
			assert.DeepEqual(t, tc.expected, got.Interface(), protocmp.Transform())
		})
	}
}
