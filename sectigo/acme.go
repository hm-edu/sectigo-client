package sectigo

import (
	"context"
	"fmt"
	"path"
	"strconv"

	"github.com/hm-edu/sectigo-client/sectigo/acme"
)

// ACMEService provides some methods handling sectigo ACME actions.
type ACMEService struct {
	Client *Client
}

// List enumerates all acme accounts.
func (c *ACMEService) List(request acme.ListRequest) (*[]acme.ListACMEItem, error) {
	params, err := formatParams(&request)
	if err != nil {
		return nil, err
	}
	data, _, err := Get[[]acme.ListACMEItem](context.Background(), c.Client, fmt.Sprintf("/acme/v1/account%v", params))
	return data, err
}

// ListServers enumerates all acme accounts.
func (c *ACMEService) ListServers() (*[]acme.ListACMEServerItem, error) {
	data, _, err := Get[[]acme.ListACMEServerItem](context.Background(), c.Client, "/acme/v1/server")
	return data, err
}

//AddDomains adds new domains to an existing acme account.
func (c *ACMEService) AddDomains(request acme.AddOrRemoveDomainsRequest, id int) (*acme.AddDomainsResponse, error) {
	data, _, err := Post[acme.AddDomainsResponse](context.Background(), c.Client, fmt.Sprintf("/acme/v1/account/%v/domains", id), request)
	return data, err
}

//DeleteDomains deletes domains from an existing acme account.
func (c *ACMEService) DeleteDomains(request acme.AddOrRemoveDomainsRequest, id int) (*acme.RemoveDomainsResponse, error) {
	data, _, err := Delete[acme.RemoveDomainsResponse](context.Background(), c.Client, fmt.Sprintf("/acme/v1/account/%v/domains", id), request)
	return data, err
}

//DeleteAccount deletes domains from an existing acme account.
func (c *ACMEService) DeleteAccount(id int) error {
	_, err := DeleteWithoutJSONResponse(context.Background(), c.Client, fmt.Sprintf("/acme/v1/account/%v", id))
	return err
}

//CreateAccount creates a new ACME account
func (c *ACMEService) CreateAccount(request acme.CreateACMERequest) (*int, error) {
	resp, err := PostWithoutJSONResponse(context.Background(), c.Client, "/acme/v1/account", request)
	if err != nil {
		return nil, err
	}
	loc, err := resp.Location()
	if err != nil {
		return nil, err
	}
	id, err := strconv.Atoi(path.Base(loc.Path))
	if err != nil {
		return nil, err
	}
	return &id, err
}
