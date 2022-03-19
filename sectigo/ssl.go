package sectigo

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hm-edu/sectigo-client/sectigo/misc"
	"github.com/hm-edu/sectigo-client/sectigo/ssl"
)

// SslService provides some methods handling sectigo ssl certificate actions.
type SslService struct {
	Client *Client
}

// SslJSONDate is a wrapper arround the time struct with a customized implementation of the json.Unmarshaler interface.
type SslJSONDate struct {
	time.Time
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in the format MM/DD/YYYY.
func (t *SslJSONDate) UnmarshalJSON(buf []byte) error {
	val := strings.Trim(string(buf), `"`)
	tt, err := time.Parse("01/02/2006", val)
	if err != nil {
		return err
	}
	t.Time = tt
	return nil
}

// ListSslItem represents a single item returned by the SslService.List method.
type ListSslItem struct {
	SslID                   int      `json:"sslId"`
	CommonName              string   `json:"commonName"`
	SubjectAlternativeNames []string `json:"subjectAlternativeNames"`
	SerialNumber            string   `json:"serialNumber"`
}

//SslDetails represents information about a single ssl certificate.
type SslDetails struct {
	CommonName              string                `json:"commonName"`
	SslID                   int                   `json:"sslId"`
	ID                      int                   `json:"id"`
	OrgID                   int                   `json:"orgId"`
	Status                  ssl.CertificateStatus `json:"status"`
	OrderNumber             int                   `json:"orderNumber"`
	CertType                misc.CertType         `json:"certType"`
	Term                    int                   `json:"term"`
	Expires                 SslJSONDate           `json:"expires"`
	SubjectAlternativeNames []string              `json:"subjectAlternativeNames"`
	SerialNumber            string                `json:"serialNumber"`
	CertificateDetails      SslCertificateDetails `json:"certificateDetails"`
	KeyType                 string                `json:"keyType"`
}

//SslCertificateDetails represents information about the details of a ssl certificate.
type SslCertificateDetails struct {
	Issuer                  string `json:"issuer,omitempty"`
	Subject                 string `json:"subject,omitempty"`
	SubjectAlternativeNames string `json:"subjectAltNames,omitempty"`
	Md5Hash                 string `json:"md5Hash,omitempty"`
	Sha1Hash                string `json:"sha1Hash,omitempty"`
}

// ListSslProfilesItem represents a single item returned by the SslService.Profiles method.
type ListSslProfilesItem struct {
	ID                  int                 `json:"id,omitempty"`
	Name                string              `json:"name,omitempty"`
	Description         string              `json:"description,omitempty"`
	Terms               []int               `json:"terms,omitempty"`
	KeyTypes            map[string][]string `json:"keyTypes,omitempty"`
	UseSecondaryOrgName bool                `json:"useSecondaryOrgName,omitempty"`
}

// RevokeRequest represents the information on a certificate revocation.
type RevokeRequest struct {
	Reason string `json:"reason"`
}

// List enumerates all ssl certificates.
func (c *SslService) List() (*[]ListSslItem, error) {
	data, _, err := Get[[]ListSslItem](context.Background(), c.Client, "/ssl/v1")
	return data, err
}

// Details queries details on a single certificate using the (internal) sectigo ID.
func (c *SslService) Details(id int) (*SslDetails, error) {
	data, _, err := Get[SslDetails](context.Background(), c.Client, fmt.Sprintf("/ssl/v1/%v", id))
	return data, err
}

// Revoke revokes a single certificate by the serial and the reason.
func (c *SslService) Revoke(serial, reason string) error {
	_, err := PostWithoutJSONResponse(context.Background(), c.Client, fmt.Sprintf("/ssl/v1/revoke/serial/%v", serial), RevokeRequest{
		Reason: reason,
	})
	return err
}

// Profiles enumerates all ssl certificate profiles.
func (c *SslService) Profiles() (*[]ListSslProfilesItem, error) {
	data, _, err := Get[[]ListSslProfilesItem](context.Background(), c.Client, "/ssl/v1/types")
	return data, err
}
