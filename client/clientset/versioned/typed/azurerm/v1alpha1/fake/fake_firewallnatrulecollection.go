/*
Copyright The Kubeform Authors.

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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFirewallNATRuleCollections implements FirewallNATRuleCollectionInterface
type FakeFirewallNATRuleCollections struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var firewallnatrulecollectionsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "firewallnatrulecollections"}

var firewallnatrulecollectionsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "FirewallNATRuleCollection"}

// Get takes name of the firewallNATRuleCollection, and returns the corresponding firewallNATRuleCollection object, and an error if there is any.
func (c *FakeFirewallNATRuleCollections) Get(name string, options v1.GetOptions) (result *v1alpha1.FirewallNATRuleCollection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(firewallnatrulecollectionsResource, c.ns, name), &v1alpha1.FirewallNATRuleCollection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirewallNATRuleCollection), err
}

// List takes label and field selectors, and returns the list of FirewallNATRuleCollections that match those selectors.
func (c *FakeFirewallNATRuleCollections) List(opts v1.ListOptions) (result *v1alpha1.FirewallNATRuleCollectionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(firewallnatrulecollectionsResource, firewallnatrulecollectionsKind, c.ns, opts), &v1alpha1.FirewallNATRuleCollectionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FirewallNATRuleCollectionList{ListMeta: obj.(*v1alpha1.FirewallNATRuleCollectionList).ListMeta}
	for _, item := range obj.(*v1alpha1.FirewallNATRuleCollectionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested firewallNATRuleCollections.
func (c *FakeFirewallNATRuleCollections) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(firewallnatrulecollectionsResource, c.ns, opts))

}

// Create takes the representation of a firewallNATRuleCollection and creates it.  Returns the server's representation of the firewallNATRuleCollection, and an error, if there is any.
func (c *FakeFirewallNATRuleCollections) Create(firewallNATRuleCollection *v1alpha1.FirewallNATRuleCollection) (result *v1alpha1.FirewallNATRuleCollection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(firewallnatrulecollectionsResource, c.ns, firewallNATRuleCollection), &v1alpha1.FirewallNATRuleCollection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirewallNATRuleCollection), err
}

// Update takes the representation of a firewallNATRuleCollection and updates it. Returns the server's representation of the firewallNATRuleCollection, and an error, if there is any.
func (c *FakeFirewallNATRuleCollections) Update(firewallNATRuleCollection *v1alpha1.FirewallNATRuleCollection) (result *v1alpha1.FirewallNATRuleCollection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(firewallnatrulecollectionsResource, c.ns, firewallNATRuleCollection), &v1alpha1.FirewallNATRuleCollection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirewallNATRuleCollection), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFirewallNATRuleCollections) UpdateStatus(firewallNATRuleCollection *v1alpha1.FirewallNATRuleCollection) (*v1alpha1.FirewallNATRuleCollection, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(firewallnatrulecollectionsResource, "status", c.ns, firewallNATRuleCollection), &v1alpha1.FirewallNATRuleCollection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirewallNATRuleCollection), err
}

// Delete takes name of the firewallNATRuleCollection and deletes it. Returns an error if one occurs.
func (c *FakeFirewallNATRuleCollections) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(firewallnatrulecollectionsResource, c.ns, name), &v1alpha1.FirewallNATRuleCollection{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFirewallNATRuleCollections) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(firewallnatrulecollectionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.FirewallNATRuleCollectionList{})
	return err
}

// Patch applies the patch and returns the patched firewallNATRuleCollection.
func (c *FakeFirewallNATRuleCollections) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FirewallNATRuleCollection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(firewallnatrulecollectionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.FirewallNATRuleCollection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirewallNATRuleCollection), err
}
