package store

import (
	"context"
	"fmt"
	"os"

	"github.com/jamolpe/invitational-generator/internal/invitational"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func createInvitationalCollection(database *mongo.Database) *mongo.Collection {
	invitationalCollectionName := os.Getenv("INVITATION_COLLECTION")
	invitationalCollection := database.Collection(invitationalCollectionName)
	return invitationalCollection
}

func (r *repository) SaveInvitation(invitation invitational.Invitation) {
	insertResult, err := r.invitationalCollection.InsertOne(context.TODO(), invitation)
	if err != nil && insertResult == nil {
		fmt.Println("[Error - repository] an error ocurred creating a invitation in the database")
	}
	fmt.Println("[OK] - repository] saved in database")
}
func (r *repository) GetInvitations() ([]invitational.Invitation, error) {
	var dbinvitations []invitational.Invitation
	cursor, err := r.invitationalCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("[Error - repository] an error ocurred loading invitations")
	}
	for cursor.Next(context.TODO()) {
		var invitation invitational.Invitation
		err := cursor.Decode(&invitation)
		if err != nil {
			fmt.Println("[ERROR - repository]: an error ocurred loading business from owner")
			break
		}
		dbinvitations = append(dbinvitations, invitation)
	}

	cursor.Close(context.TODO())
	return dbinvitations, nil
}
