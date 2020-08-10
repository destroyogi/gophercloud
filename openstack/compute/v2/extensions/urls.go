package extensions

import "github.com/yogeshwargnanasekaran/gophercloud"

func ActionURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "action")
}
