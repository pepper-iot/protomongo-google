package wrappers_bson

import (
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw/bsonrwtest"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"gotest.tools/v3/assert"
)

func TestInt64ValueCodec(t *testing.T) {
	type testCase struct {
		name     string
		val      *wrapperspb.Int64Value
		vw       *bsonrwtest.ValueReaderWriter
		expected bsonrwtest.Invoked
	}
	testCases := []testCase{
		{
			"null",
			nil,
			&bsonrwtest.ValueReaderWriter{BSONType: bson.TypeNull},
			bsonrwtest.WriteNull,
		},
		{
			"int64",
			wrapperspb.Int64(42),
			&bsonrwtest.ValueReaderWriter{BSONType: bson.TypeInt64},
			bsonrwtest.WriteInt64,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := NewInt64ValueCodec()
			v := reflect.ValueOf(tc.val)
			err := c.EncodeValue(bsoncodec.EncodeContext{}, tc.vw, v)
			assert.NilError(t, err)
			assert.DeepEqual(t, tc.expected, tc.vw.Invoked)
		})
	}
}
func TestInt64ValueCodec_Decode(t *testing.T) {
	type testCase struct {
		name     string
		vr       *bsonrwtest.ValueReaderWriter
		expected *wrapperspb.Int64Value
	}
	testCases := []testCase{
		{
			"int64",
			&bsonrwtest.ValueReaderWriter{
				BSONType: bson.TypeInt64,
				Return:   int64(42),
			},
			wrapperspb.Int64(42),
		},
		{
			"int32",
			&bsonrwtest.ValueReaderWriter{
				BSONType: bson.TypeInt32,
				Return:   int32(42),
			},
			wrapperspb.Int64(42),
		},
		{
			"string",
			&bsonrwtest.ValueReaderWriter{
				BSONType: bson.TypeString,
				Return:   "42",
			},
			wrapperspb.Int64(42),
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
			&wrapperspb.Int64Value{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := NewInt64ValueCodec()
			got := reflect.New(reflect.TypeOf(tc.expected)).Elem()
			err := c.DecodeValue(bsoncodec.DecodeContext{}, tc.vr, got)
			assert.NilError(t, err)
			assert.DeepEqual(t, tc.expected, got.Interface(), protocmp.Transform())
		})
	}
}
