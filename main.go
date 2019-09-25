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
	log.Printf("api-gateway start begin 3")
	con, err := grpc.Dial("doki-contact.app:8888", grpc.WithInsecure())
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

	http.Handle("/", loggingFilter(nil))
	http.Handle("/query", loggingFilter(&relay.Handler{Schema: contactSchema}))
	log.Printf("api-gateway start end")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loggingFilter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("receive request:%s\n",r.URL.String())
		if next != nil {
			next.ServeHTTP(w, r)
		}
	})
}

type loggingHandler struct {
}

func (h *loggingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("receive request:%s\n",r.URL.String())
}
