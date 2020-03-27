package resolver

func (r *UserResolver) Token() string {
	return r.u.Token
}
