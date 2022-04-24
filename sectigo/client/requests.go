package client

import (
	"github.com/hm-edu/sectigo-client/sectigo/misc"
)

// CertificateDetails represents the details about a single client certificate.
type CertificateDetails struct {
	Issuer                  string `json:"issuer,omitempty"`
	Subject                 string `json:"subject,omitempty"`
	SubjectAlternativeNames string `json:"subjectAltNames,omitempty"`
	Md5Hash                 string `json:"md5Hash,omitempty"`
	Sha1Hash                string `json:"sha1Hash,omitempty"`
}

// ListItem represents a single item returned by the ClientService.List method.
type ListItem struct {
	ID                 int                `json:"id,omitempty"`
	CertificateDetails CertificateDetails `json:"certificateDetails,omitempty"`
	SerialNumber       string             `json:"serialNumber,omitempty"`
	Expires            misc.JSONDate      `json:"expires,omitempty"`
	State              State              `json:"state,omitempty"`
	OrderNumber        int                `json:"orderNumber,omitempty"`
}

// ListProfileItem represents a single item returned by the ClientService.Profiles method.
type ListProfileItem struct {
	ID                  int                 `json:"id,omitempty"`
	Name                string              `json:"name,omitempty"`
	Description         string              `json:"description,omitempty"`
	Terms               []int               `json:"terms,omitempty"`
	KeyTypes            map[string][]string `json:"keyTypes,omitempty"`
	UseSecondaryOrgName bool                `json:"useSecondaryOrgName,omitempty"`
}

// EnrollmentRequest represents the information required for enrolling a single client certificate using a CSR.
type EnrollmentRequest struct {
	OrgID           int      `json:"orgId"`
	FirstName       string   `json:"firstName"`
	MiddleName      string   `json:"middleName"`
	CommonName      string   `json:"commonName"`
	LastName        string   `json:"lastName"`
	Email           string   `json:"email"`
	Phone           string   `json:"phone"`
	SecondaryEmails []string `json:"secondaryEmails"`
	CSR             string   `json:"csr"`
	CertType        int      `json:"certType"`
	Term            int      `json:"term"`
	Eppn            string   `json:"eppn"`
}

// RevokeBySerialRequest represents the information required for revoking a certificate by a serial.
type RevokeBySerialRequest struct {
	Serial string `json:"email"`
	Reason string `json:"reason"`
}

// RevokeByEmailRequest represents the information required for revoking a certificate by an email.
type RevokeByEmailRequest struct {
	Email  string `json:"email"`
	Reason string `json:"reason"`
}

// EnrollmentResponse represents the information returned after enrolling a certificate using a CSR.
type EnrollmentResponse struct {
	OrderNumber   int    `json:"orderNumber"`
	BackendCertID string `json:"backendCertId"`
}

// ListClientRequest provides the possible filters that can be passed to the ClientService.List method.
type ListClientRequest struct {
	Size          int    `url:"size,omitempty"`
	Position      int    `url:"position,omitempty"`
	PersonID      int    `url:"personId,omitempty"`
	State         State  `url:"state,omitempty"`
	CertTypeID    int    `url:"certTypeId,omitempty"`
	SerialNumber  string `url:"serialNumber,omitempty"`
	BackendCertID int    `url:"backendCertId,omitempty"`
	Email         string `url:"email,omitempty"`
}
