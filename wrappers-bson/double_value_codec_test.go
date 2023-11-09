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

func TestDoubleValueCodec_Encode(t *testing.T) {
	type testCase struct {
		name     string
		val      *wrapperspb.DoubleValue
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
			"double",
			wrapperspb.Double(3.1415926535),
			&bsonrwtest.ValueReaderWriter{BSONType: bson.TypeDouble},
			bsonrwtest.WriteDouble,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := NewDoubleValueCodec()
			v := reflect.ValueOf(tc.val)
			err := c.EncodeValue(bsoncodec.EncodeContext{}, tc.vw, v)
			assert.NilError(t, err)
			assert.DeepEqual(t, tc.expected, tc.vw.Invoked)
		})
	}
}
func TestDoubleValueCodec_Decode(t *testing.T) {
	type testCase struct {
		name     string
		vr       *bsonrwtest.ValueReaderWriter
		expected *wrapperspb.DoubleValue
	}
	testCases := []testCase{
		{
			"double",
			&bsonrwtest.ValueReaderWriter{
				BSONType: bson.TypeDouble,
				Return:   float64(3.1415926535),
			},
			wrapperspb.Double(3.1415926535),
		},
		{
			"string",
			&bsonrwtest.ValueReaderWriter{
				BSONType: bson.TypeString,
				Return:   "3.1415926535",
			},
			wrapperspb.Double(3.1415926535),
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
			&wrapperspb.DoubleValue{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := NewDoubleValueCodec()
			got := reflect.New(reflect.TypeOf(tc.expected)).Elem()
			err := c.DecodeValue(bsoncodec.DecodeContext{}, tc.vr, got)
			assert.NilError(t, err)
			assert.DeepEqual(t, tc.expected, got.Interface(), protocmp.Transform())
		})
	}
}
