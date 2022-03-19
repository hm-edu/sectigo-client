package ssl

// CertificateStatus represents the status of a SSL certificate.
type CertificateStatus string

const (
	// Invalid certificate state.
	Invalid CertificateStatus = "Invalid"
	// Requested certificate state.
	Requested CertificateStatus = "Requested"
	// Approved certificate state.
	Approved CertificateStatus = "Approved"
	// Declined certificate state.
	Declined CertificateStatus = "Declined"
	// Applied certificate state.
	Applied CertificateStatus = "Applied"
	// Issued certificate state.
	Issued CertificateStatus = "Issued"
	// Revoked certificate state.
	Revoked CertificateStatus = "Revoked"
	// Expired certificate state.
	Expired CertificateStatus = "Expired"
	// Replaced certificate state.
	Replaced CertificateStatus = "Replaced"
	// Rejected certificate state.
	Rejected CertificateStatus = "Rejected"
	// Unmanaged certificate state.
	Unmanaged CertificateStatus = "Unmanaged"
	// SAApproved certificate state.
	SAApproved CertificateStatus = "SAApproved"
	// Init certificate state.
	Init CertificateStatus = "Init"
)
