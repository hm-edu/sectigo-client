package sectigo

import (
	"fmt"
	"github.com/hm-edu/sectigo-client/sectigo/dcv"
	"github.com/hm-edu/sectigo-client/sectigo/misc"
	"net/http"
	"testing"
	"time"

	"github.com/hm-edu/sectigo-client/sectigo/domain"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestDomainValidationService_NoPermission(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://cert-manager.com/api/dcv/v1/validation", httpmock.NewStringResponder(400, `{"code":-3,"description":"You are not authorized to perform DCV validation."}`))

	c := NewClient(http.DefaultClient, "", "", "")
	validation, err := c.DomainValidationService.List()
	assert.Nil(t, validation)
	assert.Equal(t, "GET https://cert-manager.com/api/dcv/v1/validation: 400 -3 You are not authorized to perform DCV validation.", fmt.Sprintf("%v", err))
}
func TestDomainValidationService_List(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://cert-manager.com/api/dcv/v1/validation", httpmock.NewStringResponder(200, `[{"domain":"ccmqa.com","dcvStatus":"VALIDATED","dcvOrderStatus":"NOT_INITIATED","dcvMethod":"EMAIL","expirationDate":"2022-01-22"},{"domain":"www.ccmqa.com","dcvStatus":"VALIDATED","dcvOrderStatus":"NOT_INITIATED","dcvMethod":"EMAIL","expirationDate":"2022-01-22"}]`))

	c := NewClient(http.DefaultClient, "", "", "")
	validation, err := c.DomainValidationService.List()
	assert.Nil(t, err)
	assert.Equal(t, dcv.ListItem{Domain: "ccmqa.com", DCVStatus: domain.Validated, DCVOrderStatus: dcv.NotInitiated, ExpirationDate: misc.JSONDate{time.Date(2022, time.January, 22, 0, 0, 0, 0, time.UTC)}, DCVMethod: "EMAIL"}, (*validation)[0])
}

func TestDomainValidationService_Status(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", "https://cert-manager.com/api/dcv/v2/validation/status", httpmock.NewStringResponder(200, `{"status":"EXPIRED","orderStatus":"SUBMITTED","expirationDate":"2022-01-20"}`))

	c := NewClient(http.DefaultClient, "", "", "")
	validation, err := c.DomainValidationService.Status("ccmdev.com")
	assert.Nil(t, err)
	assert.Equal(t, dcv.StatusResponse{Status: domain.Expired, OrderStatus: dcv.Submitted, ExpirationDate: misc.JSONDate{Time: time.Date(2022, time.January, 20, 0, 0, 0, 0, time.UTC)}}, *validation)
}

func TestDomainValidationService_StartCname(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", "https://cert-manager.com/api/dcv/v1/validation/start/domain/cname", httpmock.NewStringResponder(200, `{"host":"_fc69541f3cb60467c4cf674225e89931.ccmqa.com.","point":"2e48bb4e8a04ec6ff6396052bfa7e3df.7756667c4e96769d2101d67de72584dd.sectigo.com."}`))

	c := NewClient(http.DefaultClient, "", "", "")
	data, err := c.DomainValidationService.StartCNAME("ccmdev.com")
	assert.Nil(t, err)
	assert.Equal(t, dcv.StartCNAMEResponse{Host: "_fc69541f3cb60467c4cf674225e89931.ccmqa.com.", Point: "2e48bb4e8a04ec6ff6396052bfa7e3df.7756667c4e96769d2101d67de72584dd.sectigo.com."}, *data)
}

func TestDomainValidationService_SubmitCname(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", "https://cert-manager.com/api/dcv/v1/validation/submit/domain/cname", httpmock.NewStringResponder(200, `{"status":"VALIDATED","orderStatus":"SUBMITTED","message":"Submitted successfully"}`))

	c := NewClient(http.DefaultClient, "", "", "")
	data, err := c.DomainValidationService.SubmitCNAME("ccmdev.com")
	assert.Nil(t, err)
	assert.Equal(t, dcv.SubmitCNAMEResponse{Status: domain.Validated, OrderStatus: dcv.Submitted, Message: "Submitted successfully"}, *data)

}
