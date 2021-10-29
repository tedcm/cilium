package ipam

import (
	"fmt"
	"time"

	"github.com/cilium/cilium/pkg/addressing"
	"github.com/cilium/cilium/pkg/checker"
	"github.com/cilium/cilium/pkg/cidr"
	"github.com/cilium/cilium/pkg/datapath/fake"
	ipamOption "github.com/cilium/cilium/pkg/ipam/option"
	ipamTypes "github.com/cilium/cilium/pkg/ipam/types"
	"github.com/cilium/cilium/pkg/trigger"

	. "gopkg.in/check.v1"
)

type testConfigurationCRD struct{}

func (t *testConfigurationCRD) IPv4Enabled() bool                        { return true }
func (t *testConfigurationCRD) IPv6Enabled() bool                        { return false }
func (t *testConfigurationCRD) HealthCheckingEnabled() bool              { return true }
func (t *testConfigurationCRD) IPAMMode() string                         { return ipamOption.IPAMCRD }
func (t *testConfigurationCRD) BlacklistConflictingRoutesEnabled() bool  { return false }
func (t *testConfigurationCRD) SetIPv4NativeRoutingCIDR(cidr *cidr.CIDR) {}
func (t *testConfigurationCRD) IPv4NativeRoutingCIDR() *cidr.CIDR        { return nil }

func newFakeNodeStore(conf Configuration, c *C) *nodeStore {
	t, err := trigger.NewTrigger(trigger.Parameters{
		Name:        "fake-crd-allocator-node-refresher",
		MinInterval: 3 * time.Second,
		TriggerFunc: func(reasons []string) {},
	})
	if err != nil {
		log.WithError(err).Fatal("Unable to initialize CiliumNode synchronization trigger")
	}
	store := &nodeStore{
		allocators:         []*crdAllocator{},
		allocationPoolSize: map[Family]int{},
		conf:               conf,
		refreshTrigger:     t,
	}
	return store
}

func (s *IPAMSuite) TestMarkForReleaseNoAllocate(c *C) {
	cn := newCiliumNode("node1", 4, 4, 0)
	dummyResource := ipamTypes.AllocationIP{Resource: "foo"}
	for i := 1; i <= 4; i++ {
		cn.Spec.IPAM.Pool[fmt.Sprintf("1.1.1.%d", i)] = dummyResource
	}

	fakeAddressing := fake.NewNodeAddressing()
	conf := &testConfigurationCRD{}
	initNodeStore.Do(func() {
		sharedNodeStore = newFakeNodeStore(conf, c)
		sharedNodeStore.ownNode = cn
	})
	ipam := NewIPAM(fakeAddressing, conf, &ownerMock{}, &ownerMock{})
	sharedNodeStore.updateLocalNodeResource(cn)

	// Allocate the first 3 IPs
	for i := 1; i <= 3; i++ {
		epipv4, _ := addressing.NewCiliumIPv4(fmt.Sprintf("1.1.1.%d", i))
		_, err := ipam.IPv4Allocator.Allocate(epipv4.IP(), fmt.Sprintf("test%d", i))
		c.Assert(err, IsNil)
	}

	// Update 1.1.1.4 as marked for release like operator would.
	cn.Status.IPAM.ReleaseIps["1.1.1.4"] = ipamOption.IPAMMarkForRelease
	// Attempts to allocate 1.1.1.4 should fail, since it's already marked for release
	epipv4, _ := addressing.NewCiliumIPv4("1.1.1.4")
	_, err := ipam.IPv4Allocator.Allocate(epipv4.IP(), "test")
	c.Assert(err, NotNil)
	// Call agent's CRD update function. status for 1.1.1.4 should change from marked for release to ready for release
	sharedNodeStore.updateLocalNodeResource(cn)
	c.Assert(cn.Status.IPAM.ReleaseIps["1.1.1.4"], checker.Equals, ipamOption.IPAMReadyForRelease)

	// Verify that 1.1.1.3 is denied for release, since it's already in use
	cn.Status.IPAM.ReleaseIps["1.1.1.3"] = ipamOption.IPAMMarkForRelease
	sharedNodeStore.updateLocalNodeResource(cn)
	c.Assert(cn.Status.IPAM.ReleaseIps["1.1.1.3"], checker.Equals, ipamOption.IPAMDoNotRelease)
}
