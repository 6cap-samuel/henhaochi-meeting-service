package configurations

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"meeting-service/exceptions"
	"os"
	"strconv"
	"time"
)

func NewMongoDatabase() *mongo.Database {
	ctx, cancel := NewMongoContext()
	defer cancel()

	mongoPoolMin, err := strconv.Atoi(os.Getenv("MONGO_POOL_MIN"))
	exceptions.PanicIfNeeded(err)

	mongoPoolMax, err := strconv.Atoi(os.Getenv("MONGO_POOL_MAX"))
	exceptions.PanicIfNeeded(err)

	mongoMaxIdleTime, err := strconv.Atoi(os.Getenv("MONGO_MAX_IDLE_TIME_SECOND"))
	exceptions.PanicIfNeeded(err)

	option := options.Client().
		ApplyURI(os.Getenv("MONGO_URI")).
		SetMinPoolSize(uint64(mongoPoolMin)).
		SetMaxPoolSize(uint64(mongoPoolMax)).
		SetMaxConnIdleTime(time.Duration(mongoMaxIdleTime) * time.Second)

	client, err := mongo.NewClient(option)
	exceptions.PanicIfNeeded(err)

	err = client.Connect(ctx)
	exceptions.PanicIfNeeded(err)

	database := client.Database(os.Getenv("MONGO_DATABASE"))
	return database
}

func NewMongoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
