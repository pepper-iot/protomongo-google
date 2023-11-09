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

func TestBytesValueCodec_Encode(t *testing.T) {
	type testCase struct {
		name     string
		val      *wrapperspb.BytesValue
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
			wrapperspb.Bytes([]byte("Hello, World!")),
			&bsonrwtest.ValueReaderWriter{BSONType: bson.TypeBinary},
			bsonrwtest.WriteBinary,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := NewBytesValueCodec()
			v := reflect.ValueOf(tc.val)
			err := c.EncodeValue(bsoncodec.EncodeContext{}, tc.vw, v)
			assert.NilError(t, err)
			assert.DeepEqual(t, tc.expected, tc.vw.Invoked)
		})
	}
}

func TestBytesValueCodec_Decode(t *testing.T) {
	type testCase struct {
		name     string
		vr       *bsonrwtest.ValueReaderWriter
		expected *wrapperspb.BytesValue
	}
	testCases := []testCase{
		{
			"string binary",
			&bsonrwtest.ValueReaderWriter{
				BSONType: bson.TypeBinary,
				Return: bsoncore.Value{
					Type: bson.TypeBinary,
					Data: bsoncore.AppendBinary(nil, 0x00, []byte("Hello, World!")),
				},
			},
			wrapperspb.Bytes([]byte("Hello, World!")),
		},
		{
			"string",
			&bsonrwtest.ValueReaderWriter{
				BSONType: bson.TypeString,
				Return:   "Hello, World!",
			},
			wrapperspb.Bytes([]byte("Hello, World!")),
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
			&wrapperspb.BytesValue{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := NewBytesValueCodec()
			got := reflect.New(reflect.TypeOf(tc.expected)).Elem()
			err := c.DecodeValue(bsoncodec.DecodeContext{}, tc.vr, got)
			assert.NilError(t, err)
			assert.DeepEqual(t, tc.expected, got.Interface(), protocmp.Transform())
		})
	}
}
