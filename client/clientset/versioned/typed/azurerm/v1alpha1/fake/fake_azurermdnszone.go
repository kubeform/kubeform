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

// FakeAzurermDnsZones implements AzurermDnsZoneInterface
type FakeAzurermDnsZones struct {
	Fake *FakeAzurermV1alpha1
}

var azurermdnszonesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermdnszones"}

var azurermdnszonesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermDnsZone"}

// Get takes name of the azurermDnsZone, and returns the corresponding azurermDnsZone object, and an error if there is any.
func (c *FakeAzurermDnsZones) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermDnsZone, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermdnszonesResource, name), &v1alpha1.AzurermDnsZone{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDnsZone), err
}

// List takes label and field selectors, and returns the list of AzurermDnsZones that match those selectors.
func (c *FakeAzurermDnsZones) List(opts v1.ListOptions) (result *v1alpha1.AzurermDnsZoneList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermdnszonesResource, azurermdnszonesKind, opts), &v1alpha1.AzurermDnsZoneList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermDnsZoneList{ListMeta: obj.(*v1alpha1.AzurermDnsZoneList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermDnsZoneList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermDnsZones.
func (c *FakeAzurermDnsZones) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermdnszonesResource, opts))
}

// Create takes the representation of a azurermDnsZone and creates it.  Returns the server's representation of the azurermDnsZone, and an error, if there is any.
func (c *FakeAzurermDnsZones) Create(azurermDnsZone *v1alpha1.AzurermDnsZone) (result *v1alpha1.AzurermDnsZone, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermdnszonesResource, azurermDnsZone), &v1alpha1.AzurermDnsZone{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDnsZone), err
}

// Update takes the representation of a azurermDnsZone and updates it. Returns the server's representation of the azurermDnsZone, and an error, if there is any.
func (c *FakeAzurermDnsZones) Update(azurermDnsZone *v1alpha1.AzurermDnsZone) (result *v1alpha1.AzurermDnsZone, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermdnszonesResource, azurermDnsZone), &v1alpha1.AzurermDnsZone{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDnsZone), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermDnsZones) UpdateStatus(azurermDnsZone *v1alpha1.AzurermDnsZone) (*v1alpha1.AzurermDnsZone, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermdnszonesResource, "status", azurermDnsZone), &v1alpha1.AzurermDnsZone{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDnsZone), err
}

// Delete takes name of the azurermDnsZone and deletes it. Returns an error if one occurs.
func (c *FakeAzurermDnsZones) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermdnszonesResource, name), &v1alpha1.AzurermDnsZone{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermDnsZones) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermdnszonesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermDnsZoneList{})
	return err
}

// Patch applies the patch and returns the patched azurermDnsZone.
func (c *FakeAzurermDnsZones) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermDnsZone, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermdnszonesResource, name, pt, data, subresources...), &v1alpha1.AzurermDnsZone{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDnsZone), err
}
