package buildinfo

import "github.com/yogeshwargnanasekaran/gophercloud"

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("build_info")
}
