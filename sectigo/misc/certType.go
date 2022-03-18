package misc

type CertType string

const (
	SSL      CertType = "SSL"
	SMIME    CertType = "SMIME"
	CodeSign CertType = "CodeSign"
)
