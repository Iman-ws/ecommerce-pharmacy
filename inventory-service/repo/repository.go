package repo

import (
	"context"
	"github.com/Iman-ws/ecommerce-pharmacy/inventory-service/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	CreateProduct(p *model.Product) error
	GetProductByID(id int) (*model.Product, error)
	UpdateProduct(p *model.Product) error
	DeleteProduct(id int) error
	ListProducts() ([]model.Product, error)
}

type MongoRepo struct {
	collection *mongo.Collection
}

func NewMongoRepo(client *mongo.Client) *MongoRepo {
	collection := client.Database("pharmacy").Collection("products")
	return &MongoRepo{collection: collection}
}

func (r *MongoRepo) CreateProduct(p *model.Product) error {
	count, _ := r.collection.CountDocuments(context.Background(), bson.D{})
	p.ID = int(count + 1)
	_, err := r.collection.InsertOne(context.Background(), p)
	return err
}

func (r *MongoRepo) GetProductByID(id int) (*model.Product, error) {
	filter := bson.D{{"id", id}}
	var p model.Product
	err := r.collection.FindOne(context.Background(), filter).Decode(&p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *MongoRepo) UpdateProduct(p *model.Product) error {
	filter := bson.D{{"id", p.ID}}
	update := bson.D{{"$set", p}}
	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	return err
}

func (r *MongoRepo) DeleteProduct(id int) error {
	filter := bson.D{{"id", id}}
	_, err := r.collection.DeleteOne(context.Background(), filter)
	return err
}

func (r *MongoRepo) ListProducts() ([]model.Product, error) {
	cursor, err := r.collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var products []model.Product
	for cursor.Next(context.Background()) {
		var p model.Product
		cursor.Decode(&p)
		products = append(products, p)
	}
	return products, nil
}
