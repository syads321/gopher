package types

type ClientQuery struct {
	OpName    string
	Query     string
	Variables map[string]interface{}
}
