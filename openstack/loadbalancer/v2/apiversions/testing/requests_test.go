package testing

import (
	"testing"

	"github.com/yogeshwargnanasekaran/gophercloud/openstack/loadbalancer/v2/apiversions"
	th "github.com/yogeshwargnanasekaran/gophercloud/testhelper"
	"github.com/yogeshwargnanasekaran/gophercloud/testhelper/client"
)

func TestListVersions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockListResponse(t)

	allVersions, err := apiversions.List(client.ServiceClient()).AllPages()
	th.AssertNoErr(t, err)

	actual, err := apiversions.ExtractAPIVersions(allVersions)
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, OctaviaAllAPIVersionResults, actual)
}
