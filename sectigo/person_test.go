package sectigo

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"

	"github.com/jarcoal/httpmock"
)

func TestPersonService_CreatePerson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", "https://cert-manager.com/api/person/v1", httpmock.ResponderFromResponse(&http.Response{
		Status:        strconv.Itoa(201),
		StatusCode:    201,
		Header:        http.Header{},
		ContentLength: -1,
	}))

	c := NewClient(http.DefaultClient, "", "", "")
	err := c.PersonService.CreatePerson(CreatePersonRequest{})
	if err != nil {
		log.Warn().Err(err)
		t.Fatalf("err: %v", err)
	}
	if httpmock.GetTotalCallCount() != 1 {
		t.Fail()
	}

}

func TestPersonService_DeletePerson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", "https://cert-manager.com/api/person/v1/111", httpmock.ResponderFromResponse(&http.Response{
		Status:        strconv.Itoa(200),
		StatusCode:    200,
		Header:        http.Header{},
		ContentLength: -1,
	}))

	c := NewClient(http.DefaultClient, "", "", "")
	err := c.PersonService.DeletePerson(111)
	if err != nil {
		log.Warn().Err(err)
		t.Fatalf("err: %v", err)
	}
	if httpmock.GetTotalCallCount() != 1 {
		t.Fail()
	}
}

func TestPersonService_List(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://cert-manager.com/api/person/v1", httpmock.NewStringResponder(200, `[{"id":301,"firstName":"Tester","middleName":"","lastName":"","email":"18259_nobody@nobody.comodo.od.ua","organizationId":3105,"validationType":"STANDARD","phone":"123456789","secondaryEmails":["321nobody@nobody.comodo.od.ua","100500admin@nobody.comodo.od.ua"],"commonName":"Tester","eppn":"","upn":""}]`))

	c := NewClient(http.DefaultClient, "", "", "")
	list, err := c.PersonService.List(nil)
	assert.Nil(t, err)
	assert.Len(t, *list, 1)
}

func TestPersonService_ListFiltered(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://cert-manager.com/api/person/v1?organizationId=3105", httpmock.NewStringResponder(200, `[{"id":301,"firstName":"Tester","middleName":"","lastName":"","email":"18259_nobody@nobody.comodo.od.ua","organizationId":3105,"validationType":"STANDARD","phone":"123456789","secondaryEmails":["321nobody@nobody.comodo.od.ua","100500admin@nobody.comodo.od.ua"],"commonName":"Tester","eppn":"","upn":""}]`))

	c := NewClient(http.DefaultClient, "", "", "")
	list, err := c.PersonService.List(&ListPersonRequest{OrganizationID: 3105})
	assert.Nil(t, err)
	assert.Equal(t, 1, httpmock.GetTotalCallCount())
	assert.Len(t, *list, 1)
}
