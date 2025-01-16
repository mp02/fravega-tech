package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/mp02/fravega-tech/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoProductRepository struct {
	Collection *mongo.Collection
}

func NewMongoProductRepository(collection *mongo.Collection) *MongoProductRepository {
	return &MongoProductRepository{Collection: collection}
}

func (r *MongoProductRepository) Create(product *domain.Product) (domain.Product, error) {
	productInserted, err := r.Collection.InsertOne(context.Background(), product)
	if err != nil {
		return domain.Product{}, err
	}
	if oid, ok := productInserted.InsertedID.(primitive.ObjectID); ok {
		product.ID = oid.Hex()
	}
	return *product, nil
}

func (r *MongoProductRepository) GetAll() ([]domain.Product, error) {
	filter := bson.M{"is_deleted": false}

	cursor, err := r.Collection.Find(context.Background(), filter)
	if err != nil {
		fmt.Println("error en el find", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var products []domain.Product
	if err := cursor.All(context.Background(), &products); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *MongoProductRepository) Delete(productID string) (domain.Product, error) {
	// Convertir el productID de string a ObjectID
	objID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return domain.Product{}, err
	}

	filter := bson.M{"_id": objID}

	update := bson.M{"$set": bson.M{"is_deleted": true, "updated_at": time.Now()}}

	var updatedProduct domain.Product

	err = r.Collection.FindOneAndUpdate(
		context.Background(),
		filter,
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&updatedProduct)

	if err != nil {
		return domain.Product{}, err
	}

	return updatedProduct, nil
}

func (r *MongoProductRepository) Update(productID string, updates *domain.UpdateProduct) (domain.Product, error) {
	// Convertir el productID de string a ObjectID
	objID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return domain.Product{}, errors.New("invalid product ID format")
	}

	updateData := bson.M{}
	if updates.Name != nil {
		updateData["name"] = *updates.Name
	}
	if updates.Description != nil {
		updateData["description"] = *updates.Description
	}
	if updates.Price != nil {
		updateData["price"] = *updates.Price
	}
	if updates.Categories != nil {
		updateData["categories"] = *updates.Categories
	}
	if updates.IsDeleted != nil {
		updateData["is_deleted"] = *updates.IsDeleted
	}

	if len(updateData) == 0 {
		return domain.Product{}, errors.New("no fields to update")
	}

	updateData["updated_at"] = time.Now()

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updateData}

	result := r.Collection.FindOneAndUpdate(context.Background(), filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After))

	// Decodificar el documento actualizado
	var updatedProduct domain.Product
	if err := result.Decode(&updatedProduct); err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.Product{}, err
		}
		return domain.Product{}, err
	}

	return updatedProduct, nil
}

func (r *MongoProductRepository) GetProductsByFilters(filters domain.ProductFilters) ([]domain.Product, error) {
	// Construir el filtro MongoDB basado en los filtros proporcionados
	filter := bson.M{}
	if filters.Name != nil {
		filter["name"] = bson.M{"$regex": *filters.Name, "$options": "i"} // Filtro insensible a mayúsculas/minúsculas
	}
	if filters.MinPrice != nil {
		filter["price"] = bson.M{"$gte": *filters.MinPrice}
	}
	if filters.MaxPrice != nil {
		filter["price"] = bson.M{"$lte": *filters.MaxPrice}
	}
	if filters.Categories != nil && len(*filters.Categories) > 0 {
		filter["categories"] = bson.M{"$all": filters.Categories}
	}

	if filters.IsDeleted == nil {
		filter["is_deleted"] = bson.M{"$eq": false}
	} else {
		filter["is_deleted"] = bson.M{"$eq": filters.IsDeleted}
	}

	// Consultar la base de datos con los filtros
	cursor, err := r.Collection.Find(context.Background(), filter)
	if err != nil {
		fmt.Println("find", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var products []domain.Product
	if err := cursor.All(context.Background(), &products); err != nil {
		fmt.Println("cursor", err)
		return nil, err
	}
	return products, nil
}
