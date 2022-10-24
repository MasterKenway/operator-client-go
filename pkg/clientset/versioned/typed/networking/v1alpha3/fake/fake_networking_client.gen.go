// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha3 "istio.io/client-go/pkg/clientset/versioned/typed/networking/v1alpha3"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeNetworkingV1alpha3 struct {
	*testing.Fake
}

func (c *FakeNetworkingV1alpha3) DestinationRules(namespace string) v1alpha3.DestinationRuleInterface {
	return &FakeDestinationRules{c, namespace}
}

func (c *FakeNetworkingV1alpha3) EnvoyFilters(namespace string) v1alpha3.EnvoyFilterInterface {
	return &FakeEnvoyFilters{c, namespace}
}

func (c *FakeNetworkingV1alpha3) Gateways(namespace string) v1alpha3.GatewayInterface {
	return &FakeGateways{c, namespace}
}

func (c *FakeNetworkingV1alpha3) ServiceEntries(namespace string) v1alpha3.ServiceEntryInterface {
	return &FakeServiceEntries{c, namespace}
}

func (c *FakeNetworkingV1alpha3) Sidecars(namespace string) v1alpha3.SidecarInterface {
	return &FakeSidecars{c, namespace}
}

func (c *FakeNetworkingV1alpha3) VirtualServices(namespace string) v1alpha3.VirtualServiceInterface {
	return &FakeVirtualServices{c, namespace}
}

func (c *FakeNetworkingV1alpha3) WorkloadEntries(namespace string) v1alpha3.WorkloadEntryInterface {
	return &FakeWorkloadEntries{c, namespace}
}

func (c *FakeNetworkingV1alpha3) WorkloadGroups(namespace string) v1alpha3.WorkloadGroupInterface {
	return &FakeWorkloadGroups{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeNetworkingV1alpha3) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
