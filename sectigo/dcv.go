package sectigo

import (
	"context"

	"github.com/hm-edu/sectigo-client/sectigo/dcv"
)

// DomainValidationService provides some methods handling sectigo domain validation actions.
type DomainValidationService struct {
	Client *Client
}

// List enumerates all domain validation requests.
func (c *DomainValidationService) List() (*[]dcv.ListItem, error) {
	data, _, err := Get[[]dcv.ListItem](context.Background(), c.Client, "/dcv/v1/validation")
	return data, err
}

// Status queries the status of a single domain.
func (c *DomainValidationService) Status(domain string) (*dcv.StatusResponse, error) {
	data, _, err := Post[dcv.StatusResponse](context.Background(), c.Client, "/dcv/v2/validation/status", dcv.Request{
		Domain: domain,
	})
	return data, err
}

// Clear resets the status of the DCV request.
func (c *DomainValidationService) Clear(domain string) (*dcv.ClearResponse, error) {
	data, _, err := Post[dcv.ClearResponse](context.Background(), c.Client, "/dcv/v1/validation/clear", dcv.Request{
		Domain: domain,
	})
	return data, err
}

// StartCNAME starts the validation using the CNAME method.
func (c *DomainValidationService) StartCNAME(domain string) (*dcv.StartCNAMEResponse, error) {
	data, _, err := Post[dcv.StartCNAMEResponse](context.Background(), c.Client, "/dcv/v1/validation/start/domain/cname", dcv.Request{
		Domain: domain,
	})
	return data, err
}

// SubmitCNAME submits the completion of the validation using the CNAME method.
func (c *DomainValidationService) SubmitCNAME(domain string) (*dcv.SubmitCNAMEResponse, error) {
	data, _, err := Post[dcv.SubmitCNAMEResponse](context.Background(), c.Client, "/dcv/v1/validation/submit/domain/cname", dcv.Request{
		Domain: domain,
	})
	return data, err
}
