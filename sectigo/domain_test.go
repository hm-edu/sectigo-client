package sectigo

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/hm-edu/sectigo-client/sectigo/domain"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/jarcoal/httpmock"
)

func TestDomainService_List(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://cert-manager.com/api/domain/v1", httpmock.NewStringResponder(200, `[{"id":0,"name":"example0.com"},{"id":1,"name":"example1.com"},{"id":2,"name":"example2.com"},{"id":3,"name":"example3.com"}]`))

	logger, _ := zap.NewProduction()
	c := NewClient(http.DefaultClient, logger, "", "", "")
	list, err := c.DomainService.List()
	assert.Nil(t, err)
	assert.Equal(t, 1, httpmock.GetTotalCallCount())
	assert.Equal(t, []domain.ListItem{{ID: 0, Name: "example0.com"}, {ID: 1, Name: "example1.com"}, {ID: 2, Name: "example2.com"}, {ID: 3, Name: "example3.com"}}, *list)
}

func TestDomainService_Infos(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://cert-manager.com/api/domain/v1/50", httpmock.NewStringResponder(200, `{"id":50,"name":"testdomain.com","delegationStatus":"ACTIVE","state":"ACTIVE","validationStatus":"VALIDATED","dcvExpiration":"2020-08-08","delegations":[{"orgId":2623,"certTypes":["SSL"],"status":"ACTIVE"}]}`))

	logger, _ := zap.NewProduction()
	c := NewClient(http.DefaultClient, logger, "", "", "")
	list, err := c.DomainService.Infos(50)
	assert.Nil(t, err)
	assert.Equal(t, 1, httpmock.GetTotalCallCount())
	assert.Equal(t, domain.Active, list.DelegationStatus)
}

func TestDomainService_CreateDomain(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", "https://cert-manager.com/api/domain/v1", httpmock.ResponderFromResponse(&http.Response{
		Status:        strconv.Itoa(201),
		StatusCode:    201,
		Header:        http.Header{},
		ContentLength: -1,
	}))

	logger, _ := zap.NewProduction()
	c := NewClient(http.DefaultClient, logger, "", "", "")
	err := c.DomainService.CreateDomain(domain.CreateRequest{})
	assert.Nil(t, err)
	assert.Equal(t, 1, httpmock.GetTotalCallCount())
}
