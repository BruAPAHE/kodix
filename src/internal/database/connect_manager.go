package database

import (
	"context"
	"errors"
	"fmt"
	"time"

	"kodix/src/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConnectManager struct {
	config    config.Database
	mongoPool map[string]*mongo.Database
}

func NewConnectManager(config config.Database) *ConnectManager {
	return &ConnectManager{
		config:    config,
		mongoPool: make(map[string]*mongo.Database),
	}
}

func (c *ConnectManager) InitConnections() []error {
	var errs []error

	// create Mongo instances ...
	pool, err := c.createMongoInstance(c.config.Mongo)
	if err != nil {
		errs = append(errs, err)
	}
	c.append(c.config.Mongo.Name, pool)

	return errs
}

func (c ConnectManager) ConnectMongo(key string) (*mongo.Database, error) {
	connect, exist := c.mongoPool[key]
	if !exist {
		return nil, errors.New(fmt.Sprintf("connect mongo `%s` not found", key))
	}
	return connect, nil
}

func (c *ConnectManager) createMongoInstance(config config.Mongo) (*mongo.Database, error) {
	clientOptions := options.Client().
		SetMaxPoolSize(100).
		SetConnectTimeout(time.Minute).
		ApplyURI(config.Dsn)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	instance := client.Database(config.Database)

	return instance, nil
}

func (c *ConnectManager) append(key string, connect *mongo.Database) {
	c.mongoPool[key] = connect
}
