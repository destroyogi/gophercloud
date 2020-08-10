package testing

import (
	"github.com/yogeshwargnanasekaran/gophercloud"
	"github.com/yogeshwargnanasekaran/gophercloud/testhelper"
)

func createClient() *gophercloud.ServiceClient {
	return &gophercloud.ServiceClient{
		ProviderClient: &gophercloud.ProviderClient{TokenID: "abc123"},
		Endpoint:       testhelper.Endpoint(),
	}
}
