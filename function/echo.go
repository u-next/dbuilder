package function

// "Fake" function used to propagate custom predicate filters from the client side.
type Echo struct {
	echo string
}

func NewEcho(echo string) Echo {
	return Echo{echo: echo}
}

func (e Echo) String() string {
	return e.echo
}
