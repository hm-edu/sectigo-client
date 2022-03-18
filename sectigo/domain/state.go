package domain

type State string

const (
	DomainActive    State = "ACTIVE"
	DomainSuspended State = "SUSPENDED"
)
