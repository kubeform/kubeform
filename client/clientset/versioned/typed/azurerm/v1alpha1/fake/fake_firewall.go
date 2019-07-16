/*
Copyright 2019 The Kubeform Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeFirewalls implements FirewallInterface
type FakeFirewalls struct {
	Fake *FakeAzurermV1alpha1
}

var firewallsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "firewalls"}

var firewallsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "Firewall"}

// Get takes name of the firewall, and returns the corresponding firewall object, and an error if there is any.
func (c *FakeFirewalls) Get(name string, options v1.GetOptions) (result *v1alpha1.Firewall, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(firewallsResource, name), &v1alpha1.Firewall{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Firewall), err
}

// List takes label and field selectors, and returns the list of Firewalls that match those selectors.
func (c *FakeFirewalls) List(opts v1.ListOptions) (result *v1alpha1.FirewallList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(firewallsResource, firewallsKind, opts), &v1alpha1.FirewallList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FirewallList{ListMeta: obj.(*v1alpha1.FirewallList).ListMeta}
	for _, item := range obj.(*v1alpha1.FirewallList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested firewalls.
func (c *FakeFirewalls) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(firewallsResource, opts))
}

// Create takes the representation of a firewall and creates it.  Returns the server's representation of the firewall, and an error, if there is any.
func (c *FakeFirewalls) Create(firewall *v1alpha1.Firewall) (result *v1alpha1.Firewall, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(firewallsResource, firewall), &v1alpha1.Firewall{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Firewall), err
}

// Update takes the representation of a firewall and updates it. Returns the server's representation of the firewall, and an error, if there is any.
func (c *FakeFirewalls) Update(firewall *v1alpha1.Firewall) (result *v1alpha1.Firewall, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(firewallsResource, firewall), &v1alpha1.Firewall{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Firewall), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFirewalls) UpdateStatus(firewall *v1alpha1.Firewall) (*v1alpha1.Firewall, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(firewallsResource, "status", firewall), &v1alpha1.Firewall{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Firewall), err
}

// Delete takes name of the firewall and deletes it. Returns an error if one occurs.
func (c *FakeFirewalls) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(firewallsResource, name), &v1alpha1.Firewall{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFirewalls) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(firewallsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.FirewallList{})
	return err
}

// Patch applies the patch and returns the patched firewall.
func (c *FakeFirewalls) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Firewall, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(firewallsResource, name, pt, data, subresources...), &v1alpha1.Firewall{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Firewall), err
}
