package gql

import (
	"context"
	"errors"
	"fmt"
	"github.com/abhayprakashtiwari/estatebidding/mongo"
	"github.com/mongodb/mongo-go-driver/bson"
	"log"
	"math/rand"
	"sort"
	"time"
)

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateEstate(ctx context.Context, input NewEstate) (*Estate, error) {
	log.Print("Request received to create estate")
	newEstate := &Estate{
		ID: fmt.Sprintf("%d", rand.Int()),
		Name: input.Name,
		Description: &input.Description,
		RegisteredAt: time.Now().UTC(),
		OpenForBidTill: time.Now().AddDate(0, 0, 5).UTC(),
		Bids: make([]Bid, 0),
	}
	_, errs:= mongo.EstateCollection.InsertOne(ctx, newEstate)
	log.Printf("Estate created with id %s", newEstate.ID)
	return newEstate,errs
}
func (r *mutationResolver) UpdateEstate(ctx context.Context, input ChangedEstate) (*Estate, error) {
	log.Printf("Request to update id %s", input.ID)
	newEstate := &Estate{
		ID: input.ID,
		Name: *input.Name,
		Description: input.Description,
		RegisteredAt: time.Now().UTC(),
		OpenForBidTill: time.Now().AddDate(0, 0, 5).UTC(),
		Bids: make([]Bid, 0),
	}
	match := bson.M{
		"id": input.ID,
	}
	//On modification, bids and timestamps will reset
	change := bson.M{
		"$set": newEstate,
	}
	updatedResult, err :=mongo.EstateCollection.UpdateOne(ctx, match, change)
	log.Printf("Estate update with id %s and matched %d", newEstate.ID, updatedResult.MatchedCount)
	return newEstate, err
}
func (r *mutationResolver) DeleteEstate(ctx context.Context, input DeleteEstate) (string, error) {
	log.Printf("Request to delete Estate id %s", input.ID)
	match := bson.M{"id": input.ID}
	result, err := mongo.EstateCollection.DeleteOne(ctx, match)
	if err != nil || result.DeletedCount != 1 {
		return "", errors.New("could_not_delete_estate")
	}
	log.Printf("Estate deleted with id %s , match count %d", input.ID, result.DeletedCount)
	return "Deleted estate", err
}
func (r *mutationResolver) CreateBid(ctx context.Context, input NewBid) (*Bid, error) {
	log.Printf("Request to create Bid for Estate  with id %s", input.EstateID)
	newBid := &Bid{
		ID: fmt.Sprintf("%d", rand.Int()),
		Amount: input.Amount,
		Bidder: input.Bidder,
		CreatedAt: time.Now().UTC(),
	}
	change := bson.M{
		"$push": bson.M{"bids": newBid},
	}
	estateID := bson.M{
		"id": input.EstateID,
	}
	_, err := mongo.EstateCollection.UpdateOne(ctx, estateID, change)
	if err != nil {
		print(err.Error())
	}
	log.Printf("Bid created with id %s", newBid.ID)
	return newBid, err
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Estate(ctx context.Context, id string) (*Estate, error) {
	log.Printf("Request to read Estate  with id %s", id)
	match := bson.M{
		"id": id,
	}
	estate := mongo.EstateCollection.FindOne(ctx, match)
	if estate  == nil {
		 return nil, errors.New("estate_not_found")
	}
	var result Estate
	err := estate.Decode(&result)
	log.Printf("Estate fetched with id %s", id)
	return &result, err
}
func (r *queryResolver) TopBid(ctx context.Context, estateID string) (*Bid, error) {
	log.Printf("Request to fetch top bid for Estate  with id %s", estateID)
	match := bson.M{
		"id":estateID,
	}
	estate := mongo.EstateCollection.FindOne(ctx, match)
	if estate == nil {
		return nil , errors.New("estate_not_found")
	}
	var result Estate
	err := estate.Decode(&result)
	if len(result.Bids) < 1 {
		return  nil, errors.New("no_bids_for_estate")
	}
	//In second price auction the highest bidder wins but the price is the slight greater than the second bidder
	sort.Slice(result.Bids[:],
		func(i, j int) bool {
			if result.Bids[i].Amount == result.Bids[j].Amount {
				//In case both bid prices are same oldest bidder wins
				return result.Bids[i].CreatedAt.Before(result.Bids[j].CreatedAt)
			}
			return result.Bids[i].Amount > result.Bids[j].Amount
		})
	var winningBid Bid
	if len(result.Bids) == 1 {
		log.Printf("Top Bid fetched for Estate with id %s with bid id %s", estateID, result.Bids[0].ID)
		return &result.Bids[0], err
	}
	secondPriceAmt := result.Bids[1].Amount
	winningBid = result.Bids[0]
	winningBid.Amount = secondPriceAmt + 0.01
	log.Printf("Top Bid fetched for Estate with id %s with bid id %s", estateID, winningBid.ID)
	return &winningBid, err
}
