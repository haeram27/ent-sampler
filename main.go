package entsampler

import (
	"context"
	"entsampler/entx/ent"
	"fmt"
	"log"

	_ "github.com/lib/pq" // SHOULD import postgresql driver
)

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func main() {
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
}
