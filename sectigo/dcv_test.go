package sectigo

import (
	"fmt"
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
	assert.Equal(t, ListDcvItem{Domain: "ccmqa.com", DcvStatus: domain.Validated, DcvOrderStatus: NotInitiated, ExpirationDate: JsonDate{time.Date(2022, time.January, 22, 0, 0, 0, 0, time.UTC)}, DcvMethod: "EMAIL"}, (*validation)[0])
}

func TestDomainValidationService_Status(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", "https://cert-manager.com/api/dcv/v2/validation/status", httpmock.NewStringResponder(200, `{"status":"EXPIRED","orderStatus":"SUBMITTED","expirationDate":"2022-01-20"}`))

	c := NewClient(http.DefaultClient, "", "", "")
	validation, err := c.DomainValidationService.Status("ccmdev.com")
	assert.Nil(t, err)
	assert.Equal(t, StatusResponse{Status: domain.Expired, OrderStatus: Submitted, ExpirationDate: JsonDate{time.Date(2022, time.January, 20, 0, 0, 0, 0, time.UTC)}}, *validation)
}

func TestDomainValidationService_StartCname(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", "https://cert-manager.com/api/dcv/v1/validation/start/domain/cname", httpmock.NewStringResponder(200, `{"host":"_fc69541f3cb60467c4cf674225e89931.ccmqa.com.","point":"2e48bb4e8a04ec6ff6396052bfa7e3df.7756667c4e96769d2101d67de72584dd.sectigo.com."}`))

	c := NewClient(http.DefaultClient, "", "", "")
	data, err := c.DomainValidationService.StartCname("ccmdev.com")
	assert.Nil(t, err)
	assert.Equal(t, StartCnameResponse{Host: "_fc69541f3cb60467c4cf674225e89931.ccmqa.com.", Point: "2e48bb4e8a04ec6ff6396052bfa7e3df.7756667c4e96769d2101d67de72584dd.sectigo.com."}, *data)
}

func TestDomainValidationService_SubmitCname(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", "https://cert-manager.com/api/dcv/v1/validation/submit/domain/cname", httpmock.NewStringResponder(200, `{"status":"VALIDATED","orderStatus":"SUBMITTED","message":"Submitted successfully"}`))

	c := NewClient(http.DefaultClient, "", "", "")
	data, err := c.DomainValidationService.SubmitCname("ccmdev.com")
	assert.Nil(t, err)
	assert.Equal(t, SubmitCnameResponse{Status: domain.Validated, OrderStatus: Submitted, Message: "Submitted successfully"}, *data)

}
