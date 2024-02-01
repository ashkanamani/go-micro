package data

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func New(mongo *mongo.Client) Models {
	client = mongo
	return Models{
		LogEntry: LogEntry{},
	}
}

type Models struct {
	LogEntry LogEntry
}

type LogEntry struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string    `bson:"name" json:"name"`
	Data      string    `bson:"data" json:"data"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

func (l *LogEntry) Insert(enrty LogEntry) error {
	collection := client.Database("logs").Collection("logs")
	
	_, err := collection.InsertOne(context.TODO(), LogEntry{
		Name: enrty.Name,
		Data: enrty.Data,
		CreatedAt: enrty.CreatedAt,
		UpdatedAt: enrty.UpdatedAt,
	})
	if err != nil {
		log.Println("error inserting into logs: ", err)
		return err
	}
	return nil
}


func (l *LogEntry) All() ([]*LogEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()

	collection := client.Database("logs").Collection("logs")

	opts := options.Find()
	opts.SetSort(bson.D{{"created_at", -1}})

	cursor, err := collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		log.Println("finding all docs error")
		return nil, err
	}

	defer cursor.Close(ctx)

	var logs []*LogEntry

	for cursor.Next(ctx) {
		var item LogEntry

		err := cursor.Decode(&item)
		if err != nil {
			log.Println("error decoding log into slice:", err)
			return nil, err 
		}
		logs = append(logs, &item)
	}
	return logs, nil

}