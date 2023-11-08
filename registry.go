package protomongo_google

import (
	structpb_bson "github.com/pepper-iot/protomongo-google/structpb-bson"
	timestamppb_bson "github.com/pepper-iot/protomongo-google/timestamppb-bson"
	wrappers_bson "github.com/pepper-iot/protomongo-google/wrappers-bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
)

func RegisterAll(rb *bsoncodec.Registry) *bsoncodec.Registry {

	rb = structpb_bson.RegisterRegistry(rb)
	rb = timestamppb_bson.RegisterRegistry(rb)
	rb = wrappers_bson.RegisterRegistry(rb)

	return rb

}
