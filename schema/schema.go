package schema

import (
	"context"
	"io/ioutil"
)

func String(ctx context.Context) (string, error) {
	schemaBytes, err := ioutil.ReadFile("./schema/schema.graphql")
	return string(schemaBytes), err
}