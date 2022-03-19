package sectigo

import (
	"context"

	"github.com/hm-edu/sectigo-client/sectigo/domain"
)

// DomainValidationService provides some methods handling sectigo domain validation actions.
type DomainValidationService struct {
	Client *Client
}

// DomainValidationOrderStatus represents the state of a DCV request.
type DomainValidationOrderStatus string

const (
	// Submitted means that DCV request was submitted.
	Submitted DomainValidationOrderStatus = "SUBMITTED"
	// NotInitiated means the DCV request was not initiated.
	NotInitiated DomainValidationOrderStatus = "NOT_INITIATED"
)

// StatusResponse represents the information about the DCV status.
type StatusResponse struct {
	Status         domain.ValidationStatus     `json:"status"`
	OrderStatus    DomainValidationOrderStatus `json:"orderStatus"`
	ExpirationDate JSONDate                    `json:"expirationDate"`
}

// ListDCVItem represents a single item returned by the DomainValidationService.List method.
type ListDCVItem struct {
	Domain         string                      `json:"domain"`
	DCVStatus      domain.ValidationStatus     `json:"dcvStatus"`
	DCVOrderStatus DomainValidationOrderStatus `json:"dcvOrderStatus"`
	ExpirationDate JSONDate                    `json:"expirationDate"`
	DCVMethod      string                      `json:"dcvMethod"`
}

// DCVRequest represents the information required for the DCV operations.
type DCVRequest struct {
	Domain string `json:"domain"`
}

// StartCNAMEResponse represents the information returned after starting the DCV using the CNAME method.
type StartCNAMEResponse struct {
	Host  string `json:"host"`
	Point string `json:"point"`
}

// SubmitCNAMEResponse represents the information after submitting the DCV using the CNAME method.
type SubmitCNAMEResponse struct {
	Status      domain.ValidationStatus     `json:"status"`
	OrderStatus DomainValidationOrderStatus `json:"orderStatus"`
	Message     string                      `json:"message"`
}

// List enumerates all domain validation requests.
func (c *DomainValidationService) List() (*[]ListDCVItem, error) {
	data, _, err := Get[[]ListDCVItem](context.Background(), c.Client, "/dcv/v1/validation")
	return data, err
}

// Status queries the status of a single domain.
func (c *DomainValidationService) Status(domain string) (*StatusResponse, error) {
	data, _, err := Post[StatusResponse](context.Background(), c.Client, "/dcv/v2/validation/status", DCVRequest{
		Domain: domain,
	})
	return data, err
}

// StartCNAME starts the validation using the CNAME method.
func (c *DomainValidationService) StartCNAME(domain string) (*StartCNAMEResponse, error) {
	data, _, err := Post[StartCNAMEResponse](context.Background(), c.Client, "/dcv/v1/validation/start/domain/cname", DCVRequest{
		Domain: domain,
	})
	return data, err
}

// SubmitCNAME submits the completion of the validation using the CNAME method.
func (c *DomainValidationService) SubmitCNAME(domain string) (*SubmitCNAMEResponse, error) {
	data, _, err := Post[SubmitCNAMEResponse](context.Background(), c.Client, "/dcv/v1/validation/submit/domain/cname", DCVRequest{
		Domain: domain,
	})
	return data, err
}
