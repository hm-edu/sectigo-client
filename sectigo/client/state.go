package client

type State string

const (
	Blank      State = "blank"
	Created    State = "created"
	Requested  State = "requested"
	Issued     State = "issued"
	Downloaded State = "downloaded"
	Expired    State = "expired"
	Revoked    State = "revoked"
	Rejected   State = "rejected"
	PreRevoked State = "pre_revoked"
)
