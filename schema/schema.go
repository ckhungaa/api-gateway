package schema

import (
	"context"
)

const fileContact = `
schema {
    query: Query
}

type Contact {
    id: ID!
    ownerId: String!
    userId: String!
    userName: String!
    userAlias: String!
    avatarId: String!
    phoneNumber: String!
    contactType: ContactType!
}
enum ContactType {
    Admin,
    Subscriber
    Friend
}

type Query {
    contact(id: ID!): Contact
}
`

func String(ctx context.Context) (string, error) {
	return fileContact, nil
}