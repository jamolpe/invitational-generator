package store

import (
	"os"

	"github.com/jamolpe/invitational-generator/internal/invitational"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	Repository interface {
		invitational.InvitationalRepository
	}
	repository struct {
		client                 *mongo.Client
		database               *mongo.Database
		invitationalCollection *mongo.Collection
	}
)

func New(client *mongo.Client) Repository {
	database := client.Database(os.Getenv("DATABASE"))
	invitationalCollection := createInvitationalCollection(database)
	return &repository{client, database, invitationalCollection}
}
