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

// FakeAzurermIotDpses implements AzurermIotDpsInterface
type FakeAzurermIotDpses struct {
	Fake *FakeAzurermV1alpha1
}

var azurermiotdpsesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermiotdpses"}

var azurermiotdpsesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermIotDps"}

// Get takes name of the azurermIotDps, and returns the corresponding azurermIotDps object, and an error if there is any.
func (c *FakeAzurermIotDpses) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermIotDps, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermiotdpsesResource, name), &v1alpha1.AzurermIotDps{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermIotDps), err
}

// List takes label and field selectors, and returns the list of AzurermIotDpses that match those selectors.
func (c *FakeAzurermIotDpses) List(opts v1.ListOptions) (result *v1alpha1.AzurermIotDpsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermiotdpsesResource, azurermiotdpsesKind, opts), &v1alpha1.AzurermIotDpsList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermIotDpsList{ListMeta: obj.(*v1alpha1.AzurermIotDpsList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermIotDpsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermIotDpses.
func (c *FakeAzurermIotDpses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermiotdpsesResource, opts))
}

// Create takes the representation of a azurermIotDps and creates it.  Returns the server's representation of the azurermIotDps, and an error, if there is any.
func (c *FakeAzurermIotDpses) Create(azurermIotDps *v1alpha1.AzurermIotDps) (result *v1alpha1.AzurermIotDps, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermiotdpsesResource, azurermIotDps), &v1alpha1.AzurermIotDps{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermIotDps), err
}

// Update takes the representation of a azurermIotDps and updates it. Returns the server's representation of the azurermIotDps, and an error, if there is any.
func (c *FakeAzurermIotDpses) Update(azurermIotDps *v1alpha1.AzurermIotDps) (result *v1alpha1.AzurermIotDps, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermiotdpsesResource, azurermIotDps), &v1alpha1.AzurermIotDps{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermIotDps), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermIotDpses) UpdateStatus(azurermIotDps *v1alpha1.AzurermIotDps) (*v1alpha1.AzurermIotDps, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermiotdpsesResource, "status", azurermIotDps), &v1alpha1.AzurermIotDps{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermIotDps), err
}

// Delete takes name of the azurermIotDps and deletes it. Returns an error if one occurs.
func (c *FakeAzurermIotDpses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermiotdpsesResource, name), &v1alpha1.AzurermIotDps{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermIotDpses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermiotdpsesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermIotDpsList{})
	return err
}

// Patch applies the patch and returns the patched azurermIotDps.
func (c *FakeAzurermIotDpses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermIotDps, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermiotdpsesResource, name, pt, data, subresources...), &v1alpha1.AzurermIotDps{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermIotDps), err
}
