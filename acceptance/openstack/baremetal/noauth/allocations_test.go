// +build acceptance baremetal allocations

package noauth

import (
	"testing"

	"github.com/yogeshwargnanasekaran/gophercloud/acceptance/clients"
	v1 "github.com/yogeshwargnanasekaran/gophercloud/acceptance/openstack/baremetal/v1"
	"github.com/yogeshwargnanasekaran/gophercloud/openstack/baremetal/v1/allocations"
	"github.com/yogeshwargnanasekaran/gophercloud/pagination"
	th "github.com/yogeshwargnanasekaran/gophercloud/testhelper"
)

func TestAllocationsCreateDestroy(t *testing.T) {
	clients.RequireLong(t)

	client, err := clients.NewBareMetalV1NoAuthClient()
	th.AssertNoErr(t, err)

	client.Microversion = "1.52"

	allocation, err := v1.CreateAllocation(t, client)
	th.AssertNoErr(t, err)
	defer v1.DeleteAllocation(t, client, allocation)

	found := false
	err = allocations.List(client, allocations.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		allocationList, err := allocations.ExtractAllocations(page)
		if err != nil {
			return false, err
		}

		for _, a := range allocationList {
			if a.UUID == allocation.UUID {
				found = true
				return true, nil
			}
		}

		return false, nil
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, found, true)
}
