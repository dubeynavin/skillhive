package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"freelance-connect-backend/config"
	"freelance-connect-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var teamCollection = config.GetCollection("team")

// GET all members
func GetTeamMembers(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := teamCollection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Error fetching team members", http.StatusInternalServerError)
		return
	}

	var members []models.TeamMember
	if err = cursor.All(ctx, &members); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(members)
}

// GET one member by ID
func GetTeamMemberByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	objID, _ := primitive.ObjectIDFromHex(id)

	var member models.TeamMember
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := teamCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&member)
	if err != nil {
		http.Error(w, "Member not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(member)
}

// POST add new member
func CreateTeamMember(w http.ResponseWriter, r *http.Request) {
	var member models.TeamMember
	json.NewDecoder(r.Body).Decode(&member)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := teamCollection.InsertOne(ctx, member)
	if err != nil {
		http.Error(w, "Error inserting member", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Team member added successfully ✅"})
}
