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

func TestUInt32ValueCodec_Encode(t *testing.T) {
	type testCase struct {
		name     string
		val      *wrapperspb.UInt32Value
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
			"uint32",
			wrapperspb.UInt32(42),
			&bsonrwtest.ValueReaderWriter{BSONType: bson.TypeInt32},
			bsonrwtest.WriteInt32,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := NewUInt32ValueCodec()
			v := reflect.ValueOf(tc.val)
			err := c.EncodeValue(bsoncodec.EncodeContext{}, tc.vw, v)
			assert.NilError(t, err)
			assert.DeepEqual(t, tc.expected, tc.vw.Invoked)
		})
	}
}
func TestUInt32ValueCodec_Decode(t *testing.T) {
	type testCase struct {
		name     string
		vr       *bsonrwtest.ValueReaderWriter
		expected *wrapperspb.UInt32Value
	}
	testCases := []testCase{
		{
			"int32",
			&bsonrwtest.ValueReaderWriter{
				BSONType: bson.TypeInt32,
				Return:   int32(42),
			},
			wrapperspb.UInt32(42),
		},
		{
			"string",
			&bsonrwtest.ValueReaderWriter{
				BSONType: bson.TypeString,
				Return:   "42",
			},
			wrapperspb.UInt32(42),
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
			&wrapperspb.UInt32Value{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := NewUInt32ValueCodec()
			got := reflect.New(reflect.TypeOf(tc.expected)).Elem()
			err := c.DecodeValue(bsoncodec.DecodeContext{}, tc.vr, got)
			assert.NilError(t, err)
			assert.DeepEqual(t, tc.expected, got.Interface(), protocmp.Transform())
		})
	}
}
