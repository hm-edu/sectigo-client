package sectigo

import (
	"context"
	"fmt"
	"github.com/hm-edu/sectigo-client/sectigo/client"
	"net/url"
)

// ClientService provides some methods handling sectigo client certificate actions.
type ClientService struct {
	Client *Client
}

// List enumerates all existing certificates.
func (c *ClientService) List(q *client.ListClientRequest) (*[]client.ListItem, error) {
	params, err := formatParams(q)
	if err != nil {
		return nil, err
	}
	data, _, err := Get[[]client.ListItem](context.Background(), c.Client, fmt.Sprintf("/smime/v2%v", params))
	return data, err
}

// ListByEmail enumerates all existing certificates owned by one user.
func (c *ClientService) ListByEmail(email string) (*[]client.ListItem, error) {
	data, _, err := Get[[]client.ListItem](context.Background(), c.Client, fmt.Sprintf("/smime/v2/byPersonEmail/%v", url.QueryEscape(email)))
	return data, err
}

// Profiles enumerates all client certificate profiles.
func (c *ClientService) Profiles() (*[]client.ListProfileItem, error) {
	data, _, err := Get[[]client.ListProfileItem](context.Background(), c.Client, "/smime/v1/types")
	return data, err
}

// Enroll submits an CSR to the server.
func (c *ClientService) Enroll(req client.EnrollmentRequest) (*client.EnrollmentResponse, error) {
	data, _, err := Post[client.EnrollmentResponse](context.Background(), c.Client, "/smime/v1/enroll", req)
	return data, err
}

// RevokeByEmail revokes all certificates associated with an email.
func (c *ClientService) RevokeByEmail(req client.RevokeByEmailRequest) error {
	_, err := PostWithoutJSONResponse(context.Background(), c.Client, "/smime/v1/revoke", req)
	return err
}

// Collect downloads the certificate for the given id.
func (c *ClientService) Collect(id int) (*string, error) {
	resp, err := GetWithoutJSONResponse(context.Background(), c.Client, fmt.Sprintf("/smime/v1/collect/%v?format=x509", id))
	bodyString, err := stringFromResponse(err, resp)
	return bodyString, err
}
