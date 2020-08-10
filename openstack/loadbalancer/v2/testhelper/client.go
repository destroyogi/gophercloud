package common

import (
	"github.com/yogeshwargnanasekaran/gophercloud"
	"github.com/yogeshwargnanasekaran/gophercloud/testhelper/client"
)

const TokenID = client.TokenID

func ServiceClient() *gophercloud.ServiceClient {
	sc := client.ServiceClient()
	sc.ResourceBase = sc.Endpoint + "v2.0/"
	return sc
}
