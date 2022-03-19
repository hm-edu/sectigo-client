package ssl

import (
	"github.com/hm-edu/sectigo-client/sectigo/misc"
)

// ListItem represents a single item returned by the SSLService.List method.
type ListItem struct {
	SslID                   int      `json:"sslId"`
	CommonName              string   `json:"commonName"`
	SubjectAlternativeNames []string `json:"subjectAlternativeNames"`
	SerialNumber            string   `json:"serialNumber"`
}

//Details represents information about a single ssl certificate.
type Details struct {
	CommonName              string             `json:"commonName"`
	SslID                   int                `json:"sslId"`
	ID                      int                `json:"id"`
	OrgID                   int                `json:"orgId"`
	Status                  CertificateStatus  `json:"status"`
	OrderNumber             int                `json:"orderNumber"`
	CertType                misc.CertType      `json:"certType"`
	Term                    int                `json:"term"`
	Expires                 JSONDate           `json:"expires"`
	SubjectAlternativeNames []string           `json:"subjectAlternativeNames"`
	SerialNumber            string             `json:"serialNumber"`
	CertificateDetails      CertificateDetails `json:"certificateDetails"`
	KeyType                 string             `json:"keyType"`
}

//CertificateDetails represents information about the details of a ssl certificate.
type CertificateDetails struct {
	Issuer                  string `json:"issuer,omitempty"`
	Subject                 string `json:"subject,omitempty"`
	SubjectAlternativeNames string `json:"subjectAltNames,omitempty"`
	Md5Hash                 string `json:"md5Hash,omitempty"`
	Sha1Hash                string `json:"sha1Hash,omitempty"`
}

// ListProfileItem represents a single item returned by the SSLService.Profiles method.
type ListProfileItem struct {
	ID                  int                 `json:"id,omitempty"`
	Name                string              `json:"name,omitempty"`
	Description         string              `json:"description,omitempty"`
	Terms               []int               `json:"terms,omitempty"`
	KeyTypes            map[string][]string `json:"keyTypes,omitempty"`
	UseSecondaryOrgName bool                `json:"useSecondaryOrgName,omitempty"`
}
