package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/1eedaegon/golang-app-oauth2-system-practice/db/ent"
	"github.com/1eedaegon/golang-app-oauth2-system-practice/db/ent/image"
	"github.com/1eedaegon/golang-app-oauth2-system-practice/db/ent/tenant"
)

func main() {
	var dsn string
	flag.StringVar(&dsn, "dsn", "", "database DSN")
	flag.Parse()

	// Instantiate the Ent client.
	client, err := ent.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("failed connecting to PostgresDB: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	if !client.Tenant.Query().Where(tenant.Name("resource-zone")).ExistX(ctx) {
		if err := seed(ctx, client); err != nil {
			log.Fatalf("failed seeding the database: %v", err)
		}
	}
}

func seed(ctx context.Context, client *ent.Client) error {
	// Check if the user "rotemtam" already exists.
	r, err := client.Image.Query().
		Where(
			image.Name("docker"),
		).
		Only(ctx)
	switch {
	// If not, create the user.
	case ent.IsNotFound(err):
		r, err = client.Image.Create().
			SetName("docker").
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed creating user: %v", err)
		}
	case err != nil:
		return fmt.Errorf("failed querying user: %v", err)
	}

	return client.Tenant.Create().
		SetName("resource-zone").
		Exec(ctx)
}
