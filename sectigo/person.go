package sectigo

import (
	"context"
	"fmt"

	"github.com/google/go-querystring/query"
)

type PersonService struct {
	Client *Client
}
type CreatePersonRequest struct {
	FirstName       string   `json:"firstName"`
	MiddleName      string   `json:"middleName"`
	LastName        string   `json:"lastName"`
	Email           string   `json:"email"`
	OrganizationID  int      `json:"organizationId"`
	ValidationType  string   `json:"validationType"`
	Phone           string   `json:"phone"`
	SecondaryEmails []string `json:"secondaryEmails"`
	CommonName      string   `json:"commonName"`
	Eppn            string   `json:"eppn"`
	Upn             string   `json:"upn"`
}

type ListPersonRequest struct {
	Position       int    `url:"position,omitempty"`
	Size           int    `url:"size,omitempty"`
	Name           string `url:"name,omitempty"`
	Email          string `url:"email,omitempty"`
	CommonName     string `url:"commonName,omitempty"`
	Phone          string `url:"phone,omitempty"`
	SecondaryEmail string `url:"secondaryEmail,omitempty"`
	OrganizationID int    `url:"organizationId,omitempty"`
}

type ListPersonItem struct {
	ID              int      `json:"id"`
	FirstName       string   `json:"firstName"`
	MiddleName      string   `json:"middleName"`
	LastName        string   `json:"lastName"`
	Email           string   `json:"email"`
	OrganizationID  int      `json:"organizationId"`
	ValidationType  string   `json:"validationType"`
	Phone           string   `json:"phone"`
	SecondaryEmails []string `json:"secondaryEmails"`
	CommonName      string   `json:"commonName"`
	Eppn            string   `json:"eppn"`
	Upn             string   `json:"upn"`
}

func (c *PersonService) List(q *ListPersonRequest) (*[]ListPersonItem, error) {
	params, err := formatParams(q)
	if err != nil {
		return nil, err
	}
	data, _, err := Get[[]ListPersonItem](context.Background(), c.Client, fmt.Sprintf("/person/v1%v", params))
	return data, err
}

func (c *PersonService) CreatePerson(q CreatePersonRequest) error {
	_, err := PostWithoutJSONResponse(context.Background(), c.Client, "/person/v1", q)
	return err
}

func (c *PersonService) DeletePerson(id int) error {
	_, err := Delete(context.Background(), c.Client, fmt.Sprintf("/person/v1/%v", id))
	return err
}

func formatParams[T any](q *T) (string, error) {
	params := ""
	if q != nil {
		values, err := query.Values(q)
		if err != nil {
			return "", err
		}
		params = fmt.Sprintf("?%v", values.Encode())
	}
	return params, nil
}
