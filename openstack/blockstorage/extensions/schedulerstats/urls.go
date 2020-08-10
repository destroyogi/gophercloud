package schedulerstats

import "github.com/yogeshwargnanasekaran/gophercloud"

func storagePoolsListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("scheduler-stats", "get_pools")
}
