// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package mocker

import (
	"math/rand"
	"net"
	"net/netip"

	petname "github.com/dustinkirkland/golang-petname"

	"github.com/cilium/cilium/pkg/identity"
	"github.com/cilium/cilium/pkg/labels"
	"github.com/cilium/cilium/pkg/lock"
)

type random struct {
	nodeIP4, nodeIP6 addr
	podIP4, podIP6   addr
	svcIP4, svcIP6   addr
	cidr4, cidr6     prefix
}

func newRandom() *random {
	return &random{
		nodeIP4: addr{addr: netip.MustParseAddr("172.16.0.0")},
		nodeIP6: addr{addr: netip.MustParseAddr("fc00::0")},
		podIP4:  addr{addr: netip.MustParseAddr("10.0.0.0")},
		podIP6:  addr{addr: netip.MustParseAddr("fd00::0")},
		svcIP4:  addr{addr: netip.MustParseAddr("172.252.0.0")},
		svcIP6:  addr{addr: netip.MustParseAddr("fdff::0")},
		cidr4:   prefix{pfx: netip.MustParsePrefix("10.0.0.0/28")},
		cidr6:   prefix{pfx: netip.MustParsePrefix("fd00::0/120")},
	}
}

func (r *random) Name() string      { return petname.Generate(2, "-") }
func (r *random) Namespace() string { return petname.Name() }

func (r *random) NodeIP4() net.IP { return r.nodeIP4.Next() }
func (r *random) NodeIP6() net.IP { return r.nodeIP6.Next() }

func (r *random) PodIP4() net.IP { return r.podIP4.Next() }
func (r *random) PodIP6() net.IP { return r.podIP6.Next() }

func (r *random) ServiceIP4() net.IP { return r.svcIP4.Next() }
func (r *random) ServiceIP6() net.IP { return r.svcIP6.Next() }

func (r *random) PodIP() net.IP {
	if rand.Intn(2) == 1 {
		return r.PodIP4()
	}

	return r.PodIP6()
}

func (r *random) CIDR4() *net.IPNet { return r.cidr4.Next() }
func (r *random) CIDR6() *net.IPNet { return r.cidr6.Next() }

func (r *random) Index(length int) int       { return rand.Intn(length) }
func (r *random) ShouldUpdateUnlikely() bool { return rand.Intn(5) == 0 }
func (r *random) ShouldUpdateLikely() bool   { return rand.Intn(100) != 0 }
func (r *random) ShouldRemove() bool         { return rand.Intn(2) == 1 }

func (r *random) Identity(cluster uint32) identity.NumericIdentity {
	return identity.NumericIdentity(cluster<<16 + uint32(rand.Intn(65536-256)+256))
}

func (r *random) IdentityLabels(cluster string) labels.LabelArray {
	n := rand.Intn(8) + 1
	lbls := make(labels.LabelArray, 0, n+4)

	ns := r.Namespace()
	lbls = append(lbls, labels.NewLabel("io.kubernetes.pod.namespace", ns, labels.LabelSourceK8s))
	lbls = append(lbls, labels.NewLabel("io.cilium.k8s.namespace.labels.kubernetes.io/metadata.name", ns, labels.LabelSourceK8s))
	lbls = append(lbls, labels.NewLabel("io.cilium.k8s.policy.serviceaccount", petname.Name(), labels.LabelSourceK8s))
	lbls = append(lbls, labels.NewLabel("io.cilium.k8s.policy.cluster", cluster, labels.LabelSourceK8s))

	for len(lbls) <= n {
		lbls = append(lbls, labels.NewLabel(petname.Generate(3, "."), petname.Adjective(), labels.LabelSourceK8s))
	}

	return lbls
}

func (r *random) ServiceBackends() int { return rand.Intn(50) }
func (r *random) ServiceLabels() map[string]string {
	n := rand.Intn(6) + 1
	lbls := make(map[string]string, n)

	for len(lbls) <= n {
		lbls[petname.Generate(3, ".")] = petname.Adjective()
	}

	return lbls
}

type addr struct {
	addr netip.Addr
	mu   lock.Mutex
}

func (a *addr) Next() net.IP {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.addr = a.addr.Next()
	return a.addr.AsSlice()
}

type prefix struct {
	pfx netip.Prefix
	mu  lock.Mutex
}

func (p *prefix) Next() *net.IPNet {
	p.mu.Lock()
	defer p.mu.Unlock()

	next := p.pfx.Addr()
	increment := 1 << (p.pfx.Addr().BitLen() - p.pfx.Bits())
	for i := 0; i < increment; i++ {
		next = next.Next()
	}

	p.pfx = netip.PrefixFrom(next, p.pfx.Bits())
	return &net.IPNet{IP: p.pfx.Addr().AsSlice(), Mask: net.CIDRMask(p.pfx.Bits(), next.BitLen())}
}
