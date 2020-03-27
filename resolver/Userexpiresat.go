package resolver

func (r *UserResolver) ExpiresAt() string {
	return r.u.ExpiresAt
}
