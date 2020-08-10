package availabilityzones

import "github.com/yogeshwargnanasekaran/gophercloud"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-availability-zone")
}
