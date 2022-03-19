package sectigo

import (
	"context"
	"fmt"
)

// PersonService provides some methods handling sectigo person actions.
type PersonService struct {
	Client *Client
}

// CreatePersonRequest represents the information required for creating a new person.
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

// ListPersonRequest provides the possible filters that can be passed to the PersonService.List method.
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

// ListPersonItem represents a single item returned by the PersonService.List method.
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

// List enumerates all persons using the provided (optional) filters.
func (c *PersonService) List(q *ListPersonRequest) (*[]ListPersonItem, error) {
	params, err := formatParams(q)
	if err != nil {
		return nil, err
	}
	data, _, err := Get[[]ListPersonItem](context.Background(), c.Client, fmt.Sprintf("/person/v1%v", params))
	return data, err
}

// CreatePerson creates a new person using the provided information.
func (c *PersonService) CreatePerson(q CreatePersonRequest) error {
	_, err := PostWithoutJSONResponse(context.Background(), c.Client, "/person/v1", q)
	return err
}

// DeletePerson deletes a person using the provided id.
func (c *PersonService) DeletePerson(id int) error {
	_, err := Delete(context.Background(), c.Client, fmt.Sprintf("/person/v1/%v", id))
	return err
}
