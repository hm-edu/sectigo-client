package sectigo

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/hm-edu/sectigo-client/sectigo/acme"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestACMEService_ListError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://cert-manager.com/api/acme/v1/account", httpmock.NewStringResponder(400, `{ "code": -1001,  "description": "Value of 'organizationId' must not be null." }`))
	logger, _ := zap.NewProduction()
	c := NewClient(http.DefaultClient, logger, "", "", "")
	list, err := c.AcmeService.List(acme.ListRequest{})
	assert.NotNil(t, err)
	assert.Nil(t, list)
	assert.Equal(t, "GET https://cert-manager.com/api/acme/v1/account: 400 -1001 Value of 'organizationId' must not be null.", err.Error())
}

func TestACMEService_List(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://cert-manager.com/api/acme/v1/account?organizationId=1", httpmock.NewStringResponder(200, `[{"id":81,"name":"OV ACME Account","status":"Pending","macKey":"6687b6b1-e6cd-4388-9ac2-5742381b9519","macId":"b60f9263-9fd3-4c53-a919-c1ff3c4f5cbd","acmeServer":"OV ACME Server","organizationId":1988,"certValidationType":"OV","accountId":"b60f9263-9fd3-4c53-a919-c1ff3c4f5cbd","ovOrderNumber":1946394478,"contacts":"","evDetails":{},"domains":[{"name":"domain.ccmqa.com"},{"name":"sub.domain.ccmqa.com"}]}]`))
	logger, _ := zap.NewProduction()
	c := NewClient(http.DefaultClient, logger, "", "", "")
	list, err := c.AcmeService.List(acme.ListRequest{OrganizationID: 1})
	assert.Nil(t, err)
	assert.Equal(t, []acme.ListACMEItem{{ID: 81, Name: "OV ACME Account", Status: "Pending", MacKey: "6687b6b1-e6cd-4388-9ac2-5742381b9519", MacID: "b60f9263-9fd3-4c53-a919-c1ff3c4f5cbd", AcmeServer: "OV ACME Server", OrganizationID: 1988, CertValidationType: "OV", AccountID: "b60f9263-9fd3-4c53-a919-c1ff3c4f5cbd", OvOrderNumber: 1946394478, Contacts: "", Domains: []acme.Domain{{Name: "domain.ccmqa.com"}, {Name: "sub.domain.ccmqa.com"}}}}, *list)
}

func TestACMEService_ListServers(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://cert-manager.com/api/acme/v1/server", httpmock.NewStringResponder(200, `[{"active":true,"url":"https:/acmeserverfortest-OV","caId":40485,"name":"OV ACME Server","singleProductId":66362,"multiProductId":23234,"wcProductId":14608,"certValidationType":"OV"}]`))

	logger, _ := zap.NewProduction()
	c := NewClient(http.DefaultClient, logger, "", "", "")
	list, err := c.AcmeService.ListServers()
	assert.Nil(t, err)
	assert.Equal(t, []acme.ListACMEServerItem{{Active: true, URL: "https:/acmeserverfortest-OV", CaID: 40485, Name: "OV ACME Server", SingleProductID: 66362, MultiProductID: 23234, WcProductID: 14608, CertValidationType: "OV"}}, *list)
}

func TestACMEService_CreateAccountErr(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", "https://cert-manager.com/api/acme/v1/account", httpmock.ResponderFromResponse(&http.Response{
		Status:        strconv.Itoa(400),
		StatusCode:    400,
		Header:        http.Header{},
		ContentLength: -1,
	}))
	logger, _ := zap.NewProduction()
	c := NewClient(http.DefaultClient, logger, "", "", "")
	list, err := c.AcmeService.CreateAccount(acme.CreateACMERequest{})
	assert.NotNil(t, err)
	assert.Nil(t, list)

	httpmock.RegisterResponder("POST", "https://cert-manager.com/api/acme/v1/account", httpmock.ResponderFromResponse(&http.Response{
		Status:        strconv.Itoa(201),
		StatusCode:    201,
		Header:        http.Header{},
		ContentLength: -1,
	}))
	c = NewClient(http.DefaultClient, logger, "", "", "")
	list, err = c.AcmeService.CreateAccount(acme.CreateACMERequest{})
	assert.NotNil(t, err)
	assert.Nil(t, list)

	httpmock.Reset()

	httpmock.RegisterResponder("POST", "https://cert-manager.com/api/acme/v1/account", httpmock.ResponderFromResponse(&http.Response{
		Status:        strconv.Itoa(201),
		StatusCode:    201,
		Header:        http.Header{},
		ContentLength: -1,
	}))

	list, err = c.AcmeService.CreateAccount(acme.CreateACMERequest{})
	assert.NotNil(t, err)
	assert.Nil(t, list)

	httpmock.Reset()

	httpmock.RegisterResponder("POST", "https://cert-manager.com/api/acme/v1/account", httpmock.ResponderFromResponse(&http.Response{
		Status:        strconv.Itoa(201),
		StatusCode:    201,
		Header:        http.Header{"Location": []string{"https://cert-manager.com/api/acme/v1/account/foo"}},
		ContentLength: -1,
	}))

	list, err = c.AcmeService.CreateAccount(acme.CreateACMERequest{})
	assert.NotNil(t, err)
	assert.Nil(t, list)
}

func TestACMEService_CreateAccount(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", "https://cert-manager.com/api/acme/v1/account", httpmock.ResponderFromResponse(&http.Response{
		Status:        strconv.Itoa(201),
		StatusCode:    201,
		Header:        http.Header{"Location": []string{"https://cert-manager.com/api/acme/v1/account/1"}},
		ContentLength: -1,
	}))
	logger, _ := zap.NewProduction()
	c := NewClient(http.DefaultClient, logger, "", "", "")
	list, err := c.AcmeService.CreateAccount(acme.CreateACMERequest{})
	assert.Nil(t, err)
	assert.Equal(t, 1, *list)
}

func TestACMEService_AddDomains(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", "https://cert-manager.com/api/acme/v1/account/1/domains", httpmock.NewStringResponder(200, `{"notAddedDomains":["domain.ccmqa.com.ua"]}`))

	logger, _ := zap.NewProduction()
	c := NewClient(http.DefaultClient, logger, "", "", "")
	list, err := c.AcmeService.AddDomains(acme.AddOrRemoveDomainsRequest{Domains: []acme.Domain{{Name: "test.de"}}}, 1)
	assert.Nil(t, err)
	assert.Equal(t, acme.AddDomainsResponse{NotAddedDomains: []string{"domain.ccmqa.com.ua"}}, *list)
}

func TestACMEService_DeleteDomains(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", "https://cert-manager.com/api/acme/v1/account/1/domains", httpmock.NewStringResponder(200, `{"notRemovedDomains":["domain.ccmqa.com.ua"]}`))

	logger, _ := zap.NewProduction()
	c := NewClient(http.DefaultClient, logger, "", "", "")
	list, err := c.AcmeService.DeleteDomains(acme.AddOrRemoveDomainsRequest{Domains: []acme.Domain{{Name: "test.de"}}}, 1)
	assert.Nil(t, err)
	assert.Equal(t, acme.RemoveDomainsResponse{NotRemovedDomains: []string{"domain.ccmqa.com.ua"}}, *list)
}

func TestACMEService_DeleteAccount(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", "https://cert-manager.com/api/acme/v1/account/1", httpmock.ResponderFromResponse(&http.Response{
		Status:        strconv.Itoa(200),
		StatusCode:    200,
		Header:        http.Header{},
		ContentLength: -1,
	}))
	logger, _ := zap.NewProduction()
	c := NewClient(http.DefaultClient, logger, "", "", "")
	err := c.AcmeService.DeleteAccount(1)
	assert.Nil(t, err)
}
