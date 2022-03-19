package domain

// State represents the state of a domain.
type State string

const (
	// DomainActive means that the domain is active.
	DomainActive State = "ACTIVE"
	// DomainSuspended means that the domain is suspended.
	DomainSuspended State = "SUSPENDED"
)
