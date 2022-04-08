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

// Enroll submits an CSR to the server.
func (c *SSLService) Enroll(req ssl.EnrollmentRequest) (*ssl.EnrollmentResponse, error) {
	data, _, err := Post[ssl.EnrollmentResponse](context.Background(), c.Client, "/ssl/v1/enroll", req)
	return data, err
}

// Collect downloads the certificate for the given sslId and format.
//
// Possible formats are:
// 'x509' - for Certificate (w/ chain), PEM encoded,
// 'x509CO' - for Certificate only, PEM encoded,
// 'base64' - for PKCS#7, PEM encoded,
// 'bin' - for PKCS#7
// 'x509IO' - for Root/Intermediate(s) only, PEM encoded,
// 'x509IOR' - for Intermediate(s)/Root only, PEM encoded,
// 'pem' - for Certificate (w/ chain), PEM encoded,
// 'pemco' - for Certificate only, PEM encoded,
// 'pemia' - for Certificate (w/ issuer after), PEM encoded,
// 'x509' - for Certificate (w/ chain), PEM encoded,
// 'pkcs12' - for Certificate and Private key, PKCS#12
//
// Depending on configuration at sectigo some options could be unavailable (e.g. pkcs12 requires access to the private key).
func (c *SSLService) Collect(sslID int, format string) (*string, error) {
	resp, err := GetWithoutJSONResponse(context.Background(), c.Client, fmt.Sprintf("/ssl/v1/collect/%d?format=%s", sslID, format))
	bodyString, err := stringFromResponse(err, resp)
	return bodyString, err
}
