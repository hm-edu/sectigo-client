package misc

// CertType represents different certificate types.
type CertType string

const (
	// SSL certificate.
	SSL CertType = "SSL"
	// SMIME certificate.
	SMIME CertType = "SMIME"
	// CodeSign certificate.
	CodeSign CertType = "CodeSign"
)
