package sectigo

import (
	"context"
	"fmt"

	"github.com/hm-edu/sectigo-client/sectigo/ssl"
)

// SSLService provides some methods handling sectigo ssl certificate actions.
type SSLService struct {
	Client *Client
}

// RevokeRequest represents the information on a certificate revocation.
type RevokeRequest struct {
	Reason string `json:"reason"`
}

// List enumerates all ssl certificates.
func (c *SSLService) List() (*[]ssl.ListItem, error) {
	data, _, err := Get[[]ssl.ListItem](context.Background(), c.Client, "/ssl/v1")
	return data, err
}

// Details queries details on a single certificate using the (internal) sectigo ID.
func (c *SSLService) Details(id int) (*ssl.Details, error) {
	data, _, err := Get[ssl.Details](context.Background(), c.Client, fmt.Sprintf("/ssl/v1/%v", id))
	return data, err
}

// Revoke revokes a single certificate by the serial and the reason.
func (c *SSLService) Revoke(serial, reason string) error {
	_, err := PostWithoutJSONResponse(context.Background(), c.Client, fmt.Sprintf("/ssl/v1/revoke/serial/%v", serial), RevokeRequest{
		Reason: reason,
	})
	return err
}

// Profiles enumerates all ssl certificate profiles.
func (c *SSLService) Profiles() (*[]ssl.ListProfileItem, error) {
	data, _, err := Get[[]ssl.ListProfileItem](context.Background(), c.Client, "/ssl/v1/types")
	return data, err
}
