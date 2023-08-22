package sectigo

import (
	"context"
	"fmt"

	"github.com/hm-edu/sectigo-client/sectigo/person"
)

// PersonService provides some methods handling sectigo person actions.
type PersonService struct {
	Client *Client
}

// List enumerates all persons using the provided (optional) filters.
func (c *PersonService) List(q *person.ListParams) (*[]person.ListItem, error) {
	params, err := formatParams(q)
	if err != nil {
		return nil, err
	}
	data, _, err := Get[[]person.ListItem](context.Background(), c.Client, fmt.Sprintf("/person/v1%v", params))
	return data, err
}

// CreatePerson creates a new person using the provided information.
func (c *PersonService) CreatePerson(q person.CreateRequest) error {
	_, err := PostWithoutJSONResponse(context.Background(), c.Client, "/person/v1", q)
	return err
}

// DeletePerson deletes a person using the provided id.
func (c *PersonService) DeletePerson(id int) error {
	_, err := DeleteWithoutJSONResponse(context.Background(), c.Client, fmt.Sprintf("/person/v1/%v", id))
	return err
}

// UpdatePerson updates a person using the provided id and information.
func (c *PersonService) UpdatePerson(id int, q person.UpdateRequest) error {
	_, err := PutWithoutJSONResponse(context.Background(), c.Client, fmt.Sprintf("/person/v1/%v", id), q)
	return err
}
