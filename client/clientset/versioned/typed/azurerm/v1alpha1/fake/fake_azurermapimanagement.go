/*
Copyright 2019 The KubeDB Authors.

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

// FakeAzurermApiManagements implements AzurermApiManagementInterface
type FakeAzurermApiManagements struct {
	Fake *FakeAzurermV1alpha1
}

var azurermapimanagementsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermapimanagements"}

var azurermapimanagementsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermApiManagement"}

// Get takes name of the azurermApiManagement, and returns the corresponding azurermApiManagement object, and an error if there is any.
func (c *FakeAzurermApiManagements) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermApiManagement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermapimanagementsResource, name), &v1alpha1.AzurermApiManagement{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApiManagement), err
}

// List takes label and field selectors, and returns the list of AzurermApiManagements that match those selectors.
func (c *FakeAzurermApiManagements) List(opts v1.ListOptions) (result *v1alpha1.AzurermApiManagementList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermapimanagementsResource, azurermapimanagementsKind, opts), &v1alpha1.AzurermApiManagementList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermApiManagementList{ListMeta: obj.(*v1alpha1.AzurermApiManagementList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermApiManagementList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermApiManagements.
func (c *FakeAzurermApiManagements) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermapimanagementsResource, opts))
}

// Create takes the representation of a azurermApiManagement and creates it.  Returns the server's representation of the azurermApiManagement, and an error, if there is any.
func (c *FakeAzurermApiManagements) Create(azurermApiManagement *v1alpha1.AzurermApiManagement) (result *v1alpha1.AzurermApiManagement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermapimanagementsResource, azurermApiManagement), &v1alpha1.AzurermApiManagement{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApiManagement), err
}

// Update takes the representation of a azurermApiManagement and updates it. Returns the server's representation of the azurermApiManagement, and an error, if there is any.
func (c *FakeAzurermApiManagements) Update(azurermApiManagement *v1alpha1.AzurermApiManagement) (result *v1alpha1.AzurermApiManagement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermapimanagementsResource, azurermApiManagement), &v1alpha1.AzurermApiManagement{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApiManagement), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermApiManagements) UpdateStatus(azurermApiManagement *v1alpha1.AzurermApiManagement) (*v1alpha1.AzurermApiManagement, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermapimanagementsResource, "status", azurermApiManagement), &v1alpha1.AzurermApiManagement{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApiManagement), err
}

// Delete takes name of the azurermApiManagement and deletes it. Returns an error if one occurs.
func (c *FakeAzurermApiManagements) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermapimanagementsResource, name), &v1alpha1.AzurermApiManagement{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermApiManagements) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermapimanagementsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermApiManagementList{})
	return err
}

// Patch applies the patch and returns the patched azurermApiManagement.
func (c *FakeAzurermApiManagements) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermApiManagement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermapimanagementsResource, name, pt, data, subresources...), &v1alpha1.AzurermApiManagement{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApiManagement), err
}
