package client

// State represents the state of a client certificate.
type State string

const (
	// Blank means that the client certificate is Blank.
	Blank State = "blank"
	// Created means that the client certificate was Created.
	Created State = "created"
	// Requested means that the client certificate was Requested.
	Requested State = "requested"
	// Issued means that the client certificate was Issued.
	Issued State = "issued"
	// Downloaded means that the client certificate was Downloaded.
	Downloaded State = "downloaded"
	// Expired means that the client certificate is Expired.
	Expired State = "expired"
	// Revoked means that the client certificate was Revoked.
	Revoked State = "revoked"
	// Rejected means that the client certificate was Rejected.
	Rejected State = "rejected"
	// PreRevoked means that the client certificate is PreRevoked.
	PreRevoked State = "pre_revoked"
)
