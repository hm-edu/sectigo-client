package domain

// DelegationStatus represents the state of a delegation request.
type DelegationStatus string

const (
	// Requested means that the delegation was requested and is awaiting some action.
	Requested DelegationStatus = "REQUESTED"
	// Active means that the delegation is active.
	Active DelegationStatus = "ACTIVE"
	// AwaitingMe means that the delegation is awaiting some action by this account.
	AwaitingMe DelegationStatus = "AWAITING_ME"
)
