package mongo

import (
	"context"
	"time"

	"github.com/Marlliton/ddd-golang/aggregate"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	db       *mongo.Database
	customer *mongo.Collection
}

// mongocustomer is a internal type that is used to store a
// CustomerAggregate inside this repository

type mongoCustomer struct {
	ID   uuid.UUID `bson:"id"`
	Name string    `bson:"name"`
}

func New(ctx context.Context, connectionString string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	db := client.Database("ddd")
	customer := db.Collection("customers")
	return &MongoRepository{
		db,
		customer,
	}, nil
}

func NewFromCustomer(c aggregate.Customer) mongoCustomer {
	return mongoCustomer{
		ID:   c.GetId(),
		Name: c.GetName(),
	}
}

func (m mongoCustomer) ToAggregate() aggregate.Customer {
	c := aggregate.Customer{}

	c.SetId(m.ID)
	c.SetName(m.Name)

	return c
}

func (mr *MongoRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := mr.customer.FindOne(ctx, bson.D{{Key: "id", Value: id}})

	var c mongoCustomer
	if err := result.Decode(&c); err != nil {
		return aggregate.Customer{}, err
	}

	return c.ToAggregate(), nil
}

func (mr *MongoRepository) Add(addCustomer aggregate.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	doc := NewFromCustomer(addCustomer)
	_, err := mr.customer.InsertOne(ctx, doc)
	if err != nil {
		return err
	}

	return nil
}

func (mr *MongoRepository) Update(upCustomer aggregate.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	doc := NewFromCustomer(upCustomer)
	_, err := mr.customer.UpdateOne(
		ctx,
		bson.D{{Key: "id", Value: upCustomer.GetId()}},
		doc,
	)
	if err != nil {
		return err
	}

	return nil
}
