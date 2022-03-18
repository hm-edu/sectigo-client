package ssl

type CertificateStatus string

const (
	Invalid    CertificateStatus = "Invalid"
	Requested  CertificateStatus = "Requested"
	Approved   CertificateStatus = "Approved"
	Declined   CertificateStatus = "Declined"
	Applied    CertificateStatus = "Applied"
	Issued     CertificateStatus = "Issued"
	Revoked    CertificateStatus = "Revoked"
	Expired    CertificateStatus = "Expired"
	Replaced   CertificateStatus = "Replaced"
	Rejected   CertificateStatus = "Rejected"
	Unmanaged  CertificateStatus = "Unmanaged"
	SAApproved CertificateStatus = "SAApproved"
	Init       CertificateStatus = "Init"
)
