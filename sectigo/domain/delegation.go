package domain

type DelegationStatus string

const (
	Requested  DelegationStatus = "REQUESTED"
	Active     DelegationStatus = "ACTIVE"
	AwaitingMe DelegationStatus = "AWAITING_ME"
)
