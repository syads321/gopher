package resolver

import (
	types "eaciit/gopher/types"
	graphql "github.com/graph-gophers/graphql-go"
)

var users = []*types.User{
	{
		UserID:   graphql.ID("u-001"),
		Username: "nyxerys",
		Emoji:    "🇵🇹",
		// Notes: []*types.Note{
		// 	{UserID: "n-001", Data: "Olá Mundo!"},
		// 	{UserID: "n-002", Data: "Olá novamente, mundo!"},
		// 	{UserID: "n-003", Data: "Olá, escuridão!"},
		// },
	}, {
		UserID:   graphql.ID("u-002"),
		Username: "rdnkta",
		Emoji:    "🇺🇦",
		// Notes: []*types.Note{
		// 	{UserID: "n-004", Data: "Привіт Світ!"},
		// 	{UserID: "n-005", Data: "Привіт ще раз, світ!"},
		// 	{UserID: "n-006", Data: "Привіт, темрява!"},
		// },
	}, {
		UserID:   graphql.ID("u-003"),
		Username: "username_ZAYDEK",
		Emoji:    "🇺🇸",
		// Notes: []*types.Note{
		// 	{UserID: "n-007", Data: "Hello, world!"},
		// 	{UserID: "n-008", Data: "Hello again, world!"},
		// 	{UserID: "n-009", Data: "Hello, darkness!"},
		// },
	},
}

// RootResolver jdjdj
type RootResolver struct {
	Session string
}

// UserResolver jsdkjfd
type UserResolver struct{ u *types.User }

type NoteResolver struct {
	n *types.Note
}

type ResolverError interface {
	error
	Extensions() map[string]interface{}
}

// Opt to return []*NoteResolver instead of []*Note:
func (r *UserResolver) Notes() []*NoteResolver {
	var noteRxs []*NoteResolver
	for _, note := range r.u.Notes {
		noteRxs = append(noteRxs, &NoteResolver{note})
	}
	return noteRxs
}

func (r *RootResolver) Notes(args struct{ UserID graphql.ID }) ([]*NoteResolver, error) {
	// Find user to find notes:
	user, err := r.User(args)
	if user == nil || err != nil {
		// Didn’t find user:
		return nil, err
	}
	// Found user; return notes:
	return user.Notes(), nil // We can reuse resolvers on resolvers, oh my.
}

func (r *RootResolver) Note(args struct{ UserID graphql.ID }) (*NoteResolver, error) {
	// Find note:
	for _, user := range users {
		for _, note := range user.Notes {
			if args.UserID == user.UserID {
				// Found note:
				return &NoteResolver{note}, nil
			}
		}
	}
	// Didn’t find note:
	return nil, nil
}

func (r *NoteResolver) UserID() graphql.ID {
	return r.n.UserID
}

func (r *NoteResolver) Data() string {
	return r.n.Data
}
