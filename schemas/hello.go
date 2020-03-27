package schemas

// Schema schemasss
var Schema = `
schema {
	query: Query
	mutation: Mutation
}

type User {
	userID: ID!
	username: String!
	emoji: String!
	notes: [Note!]!
	email: String!
	token: String!
	expiresat: String!
}

type Note {
	userID: ID!
	data: String!
}

type Query {
	users: [User!]!
	user(userID: ID!): User!
	notes(userID: ID!): [Note!]!
	note(userID: ID!): Note!
}

input NoteInput {
	data: String!
}

type Mutation {
	createNote(userID: ID!, note: NoteInput!): Note!
	createUser(
		userID: ID!
		username: String!
		emoji: String!
		password: String!
		notes: [NoteInput!]
		email: String!
	): User!
	login(
		email: String!
		password: String!
	): User!
	registerUser(
		password: String!
		email: String!
	): User!
}

`
