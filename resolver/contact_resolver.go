package resolver

import (
	"context"
	"github.com/ckhungaa/proto/proto"
	"github.com/graph-gophers/graphql-go"
)

var(
	contactTypeMap = map[proto.ContactType]string{
		proto.ContactType_AdminContactType: "Admin",
		proto.ContactType_SubscriberContactType: "Subscriber",
		proto.ContactType_FriendContactType: "Friend",
	}
)

type ContactResolver struct {
	contact *proto.ContactResponse
}

func (r *ContactResolver) ID(ctx context.Context) graphql.ID  {
	id := graphql.ID(r.contact.Audit.Id)
	return id
}

func (r *ContactResolver) OwnerID(ctx context.Context) string  {
	return r.contact.OwnerId
}

func (r *ContactResolver) UserID(ctx context.Context) string  {
	return r.contact.UserId
}

func (r *ContactResolver) UserName(ctx context.Context) string {
	return r.contact.UserName
}

func (r *ContactResolver) UserAlias(ctx context.Context) string  {
	return r.contact.UserAlias
}

func (r *ContactResolver) AvatarId(ctx context.Context) string {
	return r.contact.AvatarId
}

func (r *ContactResolver) PhoneNumber(ctx context.Context) string {
	return r.contact.PhoneNumber
}

func (r *ContactResolver) ContactType(ctx context.Context) string {
	return contactTypeMap[r.contact.ContactType]
}