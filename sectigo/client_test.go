package sectigo

import (
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/hm-edu/sectigo-client/sectigo/client"
	"go.uber.org/zap"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestClientService_RevokeByEmail(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", "https://cert-manager.com/api/smime/v1/revoke", httpmock.ResponderFromResponse(&http.Response{
		Status:        strconv.Itoa(204),
		StatusCode:    204,
		Header:        http.Header{},
		ContentLength: -1,
	}))

	logger, _ := zap.NewProduction()
	c := NewClient(http.DefaultClient, logger, "", "", "")
	err := c.ClientService.RevokeByEmail(client.RevokeByEmailRequest{Email: "test@online.de", Reason: "Compromised"})
	assert.Nil(t, err)
}

func TestClientService_RevokeBySerial(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", "https://cert-manager.com/api/smime/v1/revoke/serial/test",
		func(r *http.Request) (*http.Response, error) {
			b, err := io.ReadAll(r.Body)
			if err != nil {
				assert.Fail(t, err.Error())
			}
			data := strings.TrimSpace(string(b))
			assert.Equal(t, `{"reason":"Compromised"}`, data)
			return &http.Response{
				Status:        strconv.Itoa(204),
				StatusCode:    204,
				Header:        http.Header{},
				ContentLength: -1,
			}, nil
		})

	logger, _ := zap.NewProduction()
	c := NewClient(http.DefaultClient, logger, "", "", "")
	err := c.ClientService.RevokeBySerial(client.RevokeBySerialRequest{Serial: "test", Reason: "Compromised"})
	assert.Nil(t, err)
}
func TestClientService_Enroll(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", "https://cert-manager.com/api/smime/v1/enroll", httpmock.NewStringResponder(200, `{"orderNumber":123,"backendCertId":"123"}`))

	logger, _ := zap.NewProduction()
	c := NewClient(http.DefaultClient, logger, "", "", "")
	enroll, err := c.ClientService.Enroll(client.EnrollmentRequest{})
	if err != nil {
		return
	}
	assert.Nil(t, err)
	assert.Equal(t, client.EnrollmentResponse{OrderNumber: 123, BackendCertID: "123"}, *enroll)
}

func TestClientService_Profiles(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://cert-manager.com/api/smime/v1/types", httpmock.NewStringResponder(200, `[{"id":1938,"name":"Client cert SASP 221376727","description":"Client cert SASP -1425069207","terms":[365],"keyTypes":{"RSA":["2048"]},"useSecondaryOrgName":false}]`))

	logger, _ := zap.NewProduction()
	c := NewClient(http.DefaultClient, logger, "", "", "")
	profiles, err := c.ClientService.Profiles()

	assert.Nil(t, err)
	assert.Equal(t, 1, len(*profiles))
	assert.Equal(t, []int{365}, (*profiles)[0].Terms)
}

func TestClientService_Collect(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://cert-manager.com/api/smime/v1/collect/1234?format=x509", httpmock.NewStringResponder(200, `Test`))

	logger, _ := zap.NewProduction()
	c := NewClient(http.DefaultClient, logger, "", "", "")
	cert, err := c.ClientService.Collect(1234, "x509")
	assert.Nil(t, err)
	assert.Equal(t, "Test", *cert)
}

func TestClientService_ListByEmailNil(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://cert-manager.com/api/smime/v2/byPersonEmail/foobar%40test.de", httpmock.NewStringResponder(200, `[{  "id": 1,  "state": "rejected",  "certificateDetails": {}, "serialNumber": "",  "orderNumber": 0,  "backendCertId": null,  "expires": null}]`))
	logger, _ := zap.NewProduction()
	c := NewClient(http.DefaultClient, logger, "", "", "")
	list, err := c.ClientService.ListByEmail("foobar@test.de")
	assert.Nil(t, err)
	assert.Equal(t, 1, httpmock.GetTotalCallCount())
	assert.Len(t, *list, 1)
}
func TestClientService_ListByEmail(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://cert-manager.com/api/smime/v2/byPersonEmail/foobar%40test.de", httpmock.NewStringResponder(200, `[{"id":1,"state":"issued","certificateDetails":{"subject":"S/MIME Subject string"},"serialNumber":"C3:DB:6F:88:E7:20:DF:99:71:70:59:FB:D0:2D:29:B0","orderNumber":12345,"backendCertId":"12345","expires":"2345-06-07"}]`))

	logger, _ := zap.NewProduction()
	c := NewClient(http.DefaultClient, logger, "", "", "")
	list, err := c.ClientService.ListByEmail("foobar@test.de")
	assert.Nil(t, err)
	assert.Equal(t, 1, httpmock.GetTotalCallCount())
	assert.Len(t, *list, 1)
}
func TestClientService_List(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://cert-manager.com/api/smime/v2", httpmock.NewStringResponder(200, `[{"id":1,"state":"issued","certificateDetails":{"subject":"S/MIME Subject string"},"serialNumber":"C3:DB:6F:88:E7:20:DF:99:71:70:59:FB:D0:2D:29:B0","orderNumber":12345,"backendCertId":"12345","expires":"2345-06-07"}]`))

	logger, _ := zap.NewProduction()
	c := NewClient(http.DefaultClient, logger, "", "", "")
	list, err := c.ClientService.List(nil)
	assert.Nil(t, err)
	assert.Equal(t, 1, httpmock.GetTotalCallCount())
	assert.Len(t, *list, 1)
}

func TestClientService_ListFiltered(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://cert-manager.com/api/smime/v2?email=foobar%40test.de", httpmock.NewStringResponder(200, `[{"id":1,"state":"issued","certificateDetails":{"subject":"S/MIME Subject string"},"serialNumber":"C3:DB:6F:88:E7:20:DF:99:71:70:59:FB:D0:2D:29:B0","orderNumber":12345,"backendCertId":"12345","expires":"2345-06-07"}]`))

	logger, _ := zap.NewProduction()
	c := NewClient(http.DefaultClient, logger, "", "", "")
	list, err := c.ClientService.List(&client.ListClientRequest{Email: "foobar@test.de"})
	assert.Nil(t, err)
	assert.Equal(t, 1, httpmock.GetTotalCallCount())
	assert.Len(t, *list, 1)
}
