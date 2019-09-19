package main

import (
	"context"
	"github.com/ckhungaa/api-gateway/resolver"
	"github.com/ckhungaa/api-gateway/schema"
	"github.com/ckhungaa/proto/proto"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {
	con, err := grpc.Dial("localhost:28080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to contact server")
	}
	contactClient := proto.NewContactServerClient(con)

	ctx := context.TODO()
	schemaDef, err := schema.String(ctx)
	if err != nil {
		log.Fatalf("failed to read schema: %v", err)
	}

	queryResolver := resolver.QueryResolver{ContactClient: contactClient}

	contactSchema := graphql.MustParseSchema( schemaDef, &queryResolver, graphql.UseStringDescriptions())

	http.Handle("/query", &relay.Handler{Schema: contactSchema})
	log.Fatal(http.ListenAndServe(":28081", nil))
}
