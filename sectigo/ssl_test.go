package sectigo

import (
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/hm-edu/sectigo-client/sectigo/ssl"
	"github.com/stretchr/testify/assert"

	"github.com/jarcoal/httpmock"
)

func TestSslService_Revoke(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", "https://cert-manager.com/api/ssl/v1/revoke/serial/1234", httpmock.ResponderFromResponse(&http.Response{
		Status:        strconv.Itoa(201),
		StatusCode:    201,
		Header:        http.Header{},
		ContentLength: -1,
	}))

	c := NewClient(http.DefaultClient, "", "", "")
	err := c.SslService.Revoke("1234", "test")
	assert.Nil(t, err)
}

func TestSslService_List_Unauthorized(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://cert-manager.com/api/ssl/v1", httpmock.NewStringResponder(401, ``))

	c := NewClient(http.DefaultClient, "", "", "")
	list, err := c.SslService.List()

	assert.NotNil(t, err)
	assert.Nil(t, list)
}

func TestSslService_List(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://cert-manager.com/api/ssl/v1", httpmock.NewStringResponder(200, `[{"sslId":206,"commonName":"ccmqa.com"}]`))

	c := NewClient(http.DefaultClient, "", "", "")
	list, err := c.SslService.List()

	assert.Nil(t, err)
	assert.Equal(t, 1, len(*list))
}

func TestSslService_Details(t *testing.T) {
	x := SslCertificateDetails{
		Issuer:                  "CN=Sectigo RSA Organization Validation Secure Server CA,O=Sectigo Limited,L=Salford,ST=Greater Manchester,C=GB",
		Subject:                 "CN=dev.dummy.edu,O=Dummy Hochschule,ST=Bayern,C=DE",
		SubjectAlternativeNames: "dNSName=dev.dummy.edu, dNSName=www.dev.dummy.edu",
		Md5Hash:                 "1234",
		Sha1Hash:                "1234",
	}
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://cert-manager.com/api/ssl/v1/1234", httpmock.NewStringResponder(200, ` { "commonName": "dev.dummy.edu", "sslId": 1234, "id": 1234, "orgId": 1234, "status": "Revoked", "orderNumber": 1234, "backendCertId": "1234", "vendor": "Sectigo Limited", "term": 365, "expires": "02/16/2023", "serialNumber": "AB:CD:EF:01:23:45:10:2A:EB:1F:65:E7:27:F1:34:9F", "signatureAlg": "SHA256WITHRSA", "keyAlgorithm": "RSA", "keySize": 2048, "keyType": "RSA - 2048", "subjectAlternativeNames": [ "www.dev.dummy.edu" ], "certificateDetails": { "issuer": "CN=Sectigo RSA Organization Validation Secure Server CA,O=Sectigo Limited,L=Salford,ST=Greater Manchester,C=GB", "subject": "CN=dev.dummy.edu,O=Dummy Hochschule,ST=Bayern,C=DE", "subjectAltNames": "dNSName=dev.dummy.edu, dNSName=www.dev.dummy.edu", "md5Hash": "1234", "sha1Hash": "1234" } }`))

	c := NewClient(http.DefaultClient, "", "", "")
	details, err := c.SslService.Details(1234)

	assert.Nil(t, err)
	assert.Equal(t, x, details.CertificateDetails)
	assert.Equal(t, time.Date(2023, 02, 16, 0, 0, 0, 0, time.UTC), details.Expires.Time)
	assert.Equal(t, ssl.Revoked, details.Status)
	assert.Equal(t, 1, httpmock.GetTotalCallCount())
}

func TestSslService_Profiles(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://cert-manager.com/api/ssl/v1/types", httpmock.NewStringResponder(200, ` [{"id":1999,"name":"SSL SASP 41301020","description":"SSL SASP -1498892847","terms":[365],"keyTypes":{"RSA":["2048"]},"useSecondaryOrgName":false}]`))

	c := NewClient(http.DefaultClient, "", "", "")
	profiles, err := c.SslService.Profiles()

	assert.Nil(t, err)
	assert.Equal(t, 1, len(*profiles))
	assert.Equal(t, []int{365}, (*profiles)[0].Terms)
}
