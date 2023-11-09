package timestamppb_bson

import (
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
)

func RegisterRegistry(rb *bsoncodec.Registry) *bsoncodec.Registry {

	// Decoders
	rb.RegisterTypeDecoder(TimestampType, &TimestampCodec{})

	// Encoders
	rb.RegisterTypeEncoder(TimestampType, &TimestampCodec{})

	return rb

}
