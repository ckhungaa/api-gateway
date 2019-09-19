package resolver

import (
	"context"
	"github.com/ckhungaa/proto/proto"
	"github.com/graph-gophers/graphql-go"
)

type QueryResolver struct {
	ContactClient proto.ContactServerClient
}

func (r *QueryResolver) Contact(ctx context.Context, args struct{ ID graphql.ID }) (*ContactResolver, error) {
	id := string(args.ID)
	contactRsp, err := r.ContactClient.FindContactById(ctx, &proto.FindContactByIdRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return &ContactResolver{contact:contactRsp}, nil
}