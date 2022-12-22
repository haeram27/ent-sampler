package test

import (
	"context"
	"log"
	"testing"

	"entsampler/entx/ent" // generated package for ent

	_ "github.com/lib/pq" // SHOULD import postgresql driver
)

func getClient() *ent.Client {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=agents password=agents dbname=agentdb sslmode=disable")
	if err != nil {
		log.Fatalf("failed connecting to postgresql: %v", err)
	}
	defer client.Close()

	// Run the auto migration tool.
	ctx := context.Background()
	if err = client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	client := getClient()

	u, err := client.User.
		Create().
		SetName("a8m").
		SetEmail(("a8m@alab.com")).
		Save(ctx)

	if err != nil {
		t.Error("failed creating user: ", err)
	}
	t.Log("user was created: ", u)
}
