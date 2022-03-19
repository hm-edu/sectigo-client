package sectigo

import (
	"context"
	"fmt"
	"github.com/hm-edu/sectigo-client/sectigo/client"
	"net/url"
)

type ClientService struct {
	Client *Client
}

type ClientCertificateDetails struct {
	Issuer                  string `json:"issuer,omitempty"`
	Subject                 string `json:"subject,omitempty"`
	SubjectAlternativeNames string `json:"subjectAltNames,omitempty"`
	Md5Hash                 string `json:"md5Hash,omitempty"`
	Sha1Hash                string `json:"sha1Hash,omitempty"`
}

type ListClientItem struct {
	Id                 int                      `json:"id,omitempty"`
	CertificateDetails ClientCertificateDetails `json:"certificateDetails,omitempty"`
	SerialNumber       string                   `json:"serialNumber,omitempty"`
	Expires            JsonDate                 `json:"expires,omitempty"`
	State              client.State             `json:"state,omitempty"`
	OrderNumber        int                      `json:"orderNumber,omitempty"`
}

type ListClientProfilesItem struct {
	ID                  int                 `json:"id,omitempty"`
	Name                string              `json:"name,omitempty"`
	Description         string              `json:"description,omitempty"`
	Terms               []int               `json:"terms,omitempty"`
	KeyTypes            map[string][]string `json:"keyTypes,omitempty"`
	UseSecondaryOrgName bool                `json:"useSecondaryOrgName,omitempty"`
}

type ClientEnrollmentRequest struct {
	OrgID           int           `json:"orgId"`
	FirstName       string        `json:"firstName"`
	MiddleName      string        `json:"middleName"`
	CommonName      string        `json:"commonName"`
	LastName        string        `json:"lastName"`
	Email           string        `json:"email"`
	Phone           string        `json:"phone"`
	SecondaryEmails []string      `json:"secondaryEmails"`
	Csr             string        `json:"csr"`
	CertType        int           `json:"certType"`
	Term            int           `json:"term"`
	CustomFields    []interface{} `json:"customFields"`
	Eppn            string        `json:"eppn"`
}

type RevokeByEmailRequest struct {
	Email  string `json:"email"`
	Reason string `json:"reason"`
}

type ClientEnrollmentResponse struct {
	OrderNumber   int    `json:"orderNumber"`
	BackendCertID string `json:"backendCertId"`
}

type ListClientRequest struct {
	Size          int          `url:"size,omitempty"`
	Position      int          `url:"position,omitempty"`
	PersonID      int          `url:"personId,omitempty"`
	State         client.State `url:"state,omitempty"`
	CertTypeID    int          `url:"certTypeId,omitempty"`
	SerialNumber  string       `url:"serialNumber,omitempty"`
	BackendCertID int          `url:"backendCertId,omitempty"`
	Email         string       `url:"email,omitempty"`
}

// List enumerates all existing certificates
func (c *ClientService) List(q *ListClientRequest) (*[]ListClientItem, error) {
	params, err := formatParams(q)
	if err != nil {
		return nil, err
	}
	data, _, err := Get[[]ListClientItem](context.Background(), c.Client, fmt.Sprintf("/smime/v2%v", params))
	return data, err
}

// ListByEmail enumerates all existing certificates owned by one user
func (c *ClientService) ListByEmail(email string) (*[]ListClientItem, error) {
	data, _, err := Get[[]ListClientItem](context.Background(), c.Client, fmt.Sprintf("/smime/v2/byPersonEmail/%v", url.QueryEscape(email)))
	return data, err
}

// Profiles enumerates all client certificate profiles
func (c *ClientService) Profiles() (*[]ListClientProfilesItem, error) {
	data, _, err := Get[[]ListClientProfilesItem](context.Background(), c.Client, "/smime/v1/types")
	return data, err
}

// Enroll submits an CSR to the server
func (c *ClientService) Enroll(req ClientEnrollmentRequest) (*ClientEnrollmentResponse, error) {
	data, _, err := Post[ClientEnrollmentResponse](context.Background(), c.Client, "/smime/v1/enroll", req)
	return data, err
}

// RevokeByEmail revokes all certificates associated with an email
func (c *ClientService) RevokeByEmail(req RevokeByEmailRequest) error {
	_, err := PostWithoutJSONResponse(context.Background(), c.Client, "/smime/v1/revoke", req)
	return err
}

// Collect downloads the certificate for the given id
func (c *ClientService) Collect(id int) (*string, error) {
	resp, err := GetWithoutJSONResponse(context.Background(), c.Client, fmt.Sprintf("/smime/v1/collect/%v?format=x509", id))
	bodyString, err := stringFromResponse(err, resp)
	return bodyString, err
}
