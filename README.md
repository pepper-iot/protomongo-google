# protomongo-google

Codecs for Google Protobuf types.


## Usage

To register all the codecs for Google Protobuf types, use the `RegisterAll` function.

```go
package main

import (
    "context"
    "fmt"
    "time"

    codecRegistry "github.com/pepper-iot/protomongo-google"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func connectToMongo()(*mongo.Client, error){
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    endpoint := "mongodb://localhost:27017"
    reg := codecRegistry.RegisterAll(bson.NewRegistry())

    client, err := mongo.Connect(ctx, o.ApplyURI(endpoint).SetRegistry(reg))
    if err != nil {
        cancel()
        return nil, fmt.Errorf("error connecting to mongodb: %w", err)
    }

    return client, nil
}
```

Each Package also has its own `RegisterRegistry` function that can be used to register only the codecs for that package.

```go
package main

import (
    "context"
    "fmt"
    "time"

    structpbCodecRegistry "github.com/pepper-iot/protomongo-google/structpb-bson"
	timestampbCodecRegistry "github.com/pepper-iot/protomongo-google/timestamppb-bson"
	wrappersCodecRegistry "github.com/pepper-iot/protomongo-google/wrappers-bson"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func connectToMongo()(*mongo.Client, error){
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    endpoint := "mongodb://localhost:27017"
    reg := structpbCodecRegistry.RegisterRegistry(bson.NewRegistry())
	reg = timestampbCodecRegistry.RegisterRegistry(reg)
	reg = wrappersCodecRegistry.RegisterRegistry(reg)

    client, err := mongo.Connect(ctx, o.ApplyURI(endpoint).SetRegistry(reg))
    if err != nil {
        cancel()
        return nil, fmt.Errorf("error connecting to mongodb: %w", err)
    }

    return client, nil
}
```

