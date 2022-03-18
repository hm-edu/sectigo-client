package sectigo

import (
	"context"

	"github.com/hm-edu/sectigo-client/sectigo/domain"
)

type DomainValidationService struct {
	Client *Client
}

type DomainValidationOrderStatus string

const (
	Submitted    DomainValidationOrderStatus = "SUBMITTED"
	NotInitiated DomainValidationOrderStatus = "NOT_INITIATED"
)

type StatusResponse struct {
	Status         domain.ValidationStatus     `json:"status"`
	OrderStatus    DomainValidationOrderStatus `json:"orderStatus"`
	ExpirationDate JsonDate                    `json:"expirationDate"`
}

type ListDcvItem struct {
	Domain         string                      `json:"domain"`
	DcvStatus      domain.ValidationStatus     `json:"dcvStatus"`
	DcvOrderStatus DomainValidationOrderStatus `json:"dcvOrderStatus"`
	ExpirationDate JsonDate                    `json:"expirationDate"`
	DcvMethod      string                      `json:"dcvMethod"`
}

type DcvRequest struct {
	Domain string `json:"domain"`
}

type StartCnameResponse struct {
	Host  string `json:"host"`
	Point string `json:"point"`
}

type SubmitCnameResponse struct {
	Status      domain.ValidationStatus     `json:"status"`
	OrderStatus DomainValidationOrderStatus `json:"orderStatus"`
	Message     string                      `json:"message"`
}

func (c *DomainValidationService) List() (*[]ListDcvItem, error) {
	data, _, err := Get[[]ListDcvItem](context.Background(), c.Client, "/dcv/v1/validation")
	return data, err
}

func (c *DomainValidationService) Status(domain string) (*StatusResponse, error) {
	data, _, err := Post[StatusResponse](context.Background(), c.Client, "/dcv/v2/validation/status", DcvRequest{
		Domain: domain,
	})
	return data, err
}

func (c *DomainValidationService) StartCname(domain string) (*StartCnameResponse, error) {
	data, _, err := Post[StartCnameResponse](context.Background(), c.Client, "/dcv/v1/validation/start/domain/cname", DcvRequest{
		Domain: domain,
	})
	return data, err
}

func (c *DomainValidationService) SubmitCname(domain string) (*SubmitCnameResponse, error) {
	data, _, err := Post[SubmitCnameResponse](context.Background(), c.Client, "/dcv/v1/validation/submit/domain/cname", DcvRequest{
		Domain: domain,
	})
	return data, err
}
