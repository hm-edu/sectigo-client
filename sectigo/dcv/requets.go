package dcv

import (
	"github.com/hm-edu/sectigo-client/sectigo/domain"
	"github.com/hm-edu/sectigo-client/sectigo/misc"
)

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
	ExpirationDate misc.JSONDate               `json:"expirationDate"`
}

// ListItem represents a single item returned by the DomainValidationService.List method.
type ListItem struct {
	Domain         string                      `json:"domain"`
	DCVStatus      domain.ValidationStatus     `json:"dcvStatus"`
	DCVOrderStatus DomainValidationOrderStatus `json:"dcvOrderStatus"`
	ExpirationDate misc.JSONDate               `json:"expirationDate"`
	DCVMethod      string                      `json:"dcvMethod"`
}

// Request represents the information required for the DCV operations.
type Request struct {
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

// ClearResponse represents the information after clearing the DCV.
type ClearResponse struct {
	Status      domain.ValidationStatus     `json:"status"`
	OrderStatus DomainValidationOrderStatus `json:"orderStatus"`
	Message     string                      `json:"message"`
}
