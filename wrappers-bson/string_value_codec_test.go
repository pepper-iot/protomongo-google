package wrappers_bson

import (
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw/bsonrwtest"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"gotest.tools/v3/assert"
)

func TestStringValueCodec_Encode(t *testing.T) {
	type testCase struct {
		name     string
		val      *wrapperspb.StringValue
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
			"string binary",
			wrapperspb.String("Hello, World!"),
			&bsonrwtest.ValueReaderWriter{BSONType: bson.TypeString},
			bsonrwtest.WriteString,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := NewStringValueCodec()
			v := reflect.ValueOf(tc.val)
			err := c.EncodeValue(bsoncodec.EncodeContext{}, tc.vw, v)
			assert.NilError(t, err)
			assert.DeepEqual(t, tc.expected, tc.vw.Invoked)
		})
	}
}
func TestStringValueCodec_Decode(t *testing.T) {
	type testCase struct {
		name     string
		vr       *bsonrwtest.ValueReaderWriter
		expected *wrapperspb.StringValue
	}

	testCases := []testCase{
		{
			"string",
			&bsonrwtest.ValueReaderWriter{
				BSONType: bson.TypeString,
				Return:   "Hello, World!",
			},
			wrapperspb.String("Hello, World!"),
		},
		{
			"binary",
			&bsonrwtest.ValueReaderWriter{
				BSONType: bson.TypeBinary,
				Return: bsoncore.Value{
					Type: bson.TypeBinary,
					Data: bsoncore.AppendBinary(nil, 0x00, []byte("Hello, World!")),
				},
			},
			wrapperspb.String("Hello, World!"),
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
			&wrapperspb.StringValue{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := NewStringValueCodec()
			got := reflect.New(reflect.TypeOf(tc.expected)).Elem()
			err := c.DecodeValue(bsoncodec.DecodeContext{}, tc.vr, got)
			assert.NilError(t, err)
			assert.DeepEqual(t, tc.expected, got.Interface(), protocmp.Transform())
		})
	}
}
