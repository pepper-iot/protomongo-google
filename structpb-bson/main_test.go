package structpb_bson

import (
	"github.com/gogo/protobuf/jsonpb"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/types/known/structpb"
	"testing"
)

func Test_Convert(t *testing.T) {

	rb := bson.NewRegistryBuilder()
	rb.RegisterCodec(ProtoStructType, StructCodec{})
	rb.RegisterCodec(ProtoValueType, ValueCodec{})
	rb.RegisterCodec(ProtoListValueType, ListCodec{})
	bson.DefaultRegistry = rb.Build()

	fixture := &structpb.Struct{Fields: map[string]*structpb.Value{
		"nullValue": nil,
		"sliceValue": {
			Kind: &structpb.Value_ListValue{
				ListValue: &structpb.ListValue{
					Values: []*structpb.Value{
						{
							Kind: &structpb.Value_ListValue{},
						},
						{
							Kind: &structpb.Value_ListValue{ListValue: nil},
						},
						{
							Kind: &structpb.Value_ListValue{ListValue: &structpb.ListValue{Values: []*structpb.Value{}}},
						},
						{
							Kind: &structpb.Value_ListValue{
								ListValue: &structpb.ListValue{
									Values: []*structpb.Value{
										{
											Kind: &structpb.Value_StructValue{StructValue: &structpb.Struct{Fields: nil}},
										},
										{
											Kind: &structpb.Value_StructValue{StructValue: &structpb.Struct{Fields: map[string]*structpb.Value{}}},
										},
										{
											Kind: &structpb.Value_StructValue{
												StructValue: &structpb.Struct{
													Fields: map[string]*structpb.Value{
														"nullValue": nil,
														"zeroValue": {
															Kind: &structpb.Value_StructValue{},
														},
														"emptyValue": {
															Kind: &structpb.Value_StructValue{StructValue: &structpb.Struct{}},
														},
														"structValue": {Kind: &structpb.Value_StructValue{
															StructValue: &structpb.Struct{
																Fields: map[string]*structpb.Value{
																	"string":  {Kind: &structpb.Value_StringValue{StringValue: "str"}},
																	"number":  {Kind: &structpb.Value_NumberValue{NumberValue: 1234}},
																	"boolean": {Kind: &structpb.Value_BoolValue{BoolValue: true}},
																	"null":    {Kind: &structpb.Value_NullValue{NullValue: 0}},
																}},
														}},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}}

	bs, err := bson.Marshal(fixture)
	if err != nil {
		t.Fatal(err)
	}

	target := &structpb.Struct{}
	if err := bson.Unmarshal(bs, target); err != nil {
		t.Fatal(err)
	}

	jsonMarshaler := jsonpb.Marshaler{}

	jb, err := jsonMarshaler.MarshalToString(fixture)
	assert.NoErrorf(t, err, "jsonpb.Marshaler.MarshalToString(fixture) failed: %v", err)

	jt, err := jsonMarshaler.MarshalToString(target)
	assert.NoErrorf(t, err, "jsonpb.Marshaler.MarshalToString(target) failed: %v", err)

	assert.Equal(t, jb, jt)

}
