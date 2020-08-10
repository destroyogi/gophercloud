package bootfromvolume

import "github.com/yogeshwargnanasekaran/gophercloud"

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("servers")
}
