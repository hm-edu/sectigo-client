package sectigo

import (
	"context"
	"fmt"
	"net/url"

	"github.com/hm-edu/sectigo-client/sectigo/client"
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

// Collect downloads the certificate for the given ordernumber and format.
// Possible formats are:
//
// - 'x509' for Certificate (w/ chain), PEM encoded
// - 'x509CO' for Certificate only, PEM encoded
// - 'base64' for PKCS#7, PEM encoded
// - 'bin' for PKCS#7
// - 'x509IO' for Root/Intermediate(s) only, PEM encoded
// - 'x509IOR' for Intermediate(s)/Root only, PEM encoded
// - 'pem' for Certificate (w/ chain), PEM encoded
// - 'pemco' for Certificate only, PEM encoded
// - 'pemia' for Certificate (w/ issuer after), PEM encoded
// - 'pkcs12' for Certificate and Private key, PKCS#12
//
// Depending on configuration at sectigo some options could be unavailable (e.g. pkcs12 requires access to the private key).
func (c *ClientService) Collect(orderNumber int, format string) (*string, error) {
	resp, err := GetWithoutJSONResponse(context.Background(), c.Client, fmt.Sprintf("/smime/v1/collect/%d?format=%s", orderNumber, format))
	bodyString, err := stringFromResponse(err, resp)
	return bodyString, err
}
