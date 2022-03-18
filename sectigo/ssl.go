package sectigo

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hm-edu/sectigo-client/sectigo/misc"
	"github.com/hm-edu/sectigo-client/sectigo/ssl"
)

type SslService struct {
	Client *Client
}

type SslJSONDate struct {
	time.Time
}

func (t *SslJSONDate) UnmarshalJSON(buf []byte) error {
	val := strings.Trim(string(buf), `"`)
	tt, err := time.Parse("01/02/2006", val)
	if err != nil {
		return err
	}
	t.Time = tt
	return nil
}

type ListSslItem struct {
	SslID                   int      `json:"sslId"`
	CommonName              string   `json:"commonName"`
	SubjectAlternativeNames []string `json:"subjectAlternativeNames"`
	SerialNumber            string   `json:"serialNumber"`
}

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

type SslCertificateDetails struct {
	Issuer                  string `json:"issuer,omitempty"`
	Subject                 string `json:"subject,omitempty"`
	SubjectAlternativeNames string `json:"subjectAltNames,omitempty"`
	Md5Hash                 string `json:"md5Hash,omitempty"`
	Sha1Hash                string `json:"sha1Hash,omitempty"`
}

type ListSslProfilesItem struct {
	ID                  int                 `json:"id,omitempty"`
	Name                string              `json:"name,omitempty"`
	Description         string              `json:"description,omitempty"`
	Terms               []int               `json:"terms,omitempty"`
	KeyTypes            map[string][]string `json:"keyTypes,omitempty"`
	UseSecondaryOrgName bool                `json:"useSecondaryOrgName,omitempty"`
}

type RevokeRequest struct {
	Reason string `json:"reason"`
}

func (c *SslService) List() (*[]ListSslItem, error) {
	data, _, err := Get[[]ListSslItem](context.Background(), c.Client, "/ssl/v1")
	return data, err
}

func (c *SslService) Details(id int) (*SslDetails, error) {
	data, _, err := Get[SslDetails](context.Background(), c.Client, fmt.Sprintf("/ssl/v1/%v", id))
	return data, err
}

func (c *SslService) Revoke(serial, reason string) error {
	_, err := PostWithoutJSONResponse(context.Background(), c.Client, fmt.Sprintf("/ssl/v1/revoke/serial/%v", serial), RevokeRequest{
		Reason: reason,
	})
	return err
}

func (c *SslService) Profiles() (*[]ListSslProfilesItem, error) {
	data, _, err := Get[[]ListSslProfilesItem](context.Background(), c.Client, "/ssl/v1/types")
	return data, err
}
