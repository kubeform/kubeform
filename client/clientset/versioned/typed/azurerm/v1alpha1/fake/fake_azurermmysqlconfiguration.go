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

// FakeAzurermMysqlConfigurations implements AzurermMysqlConfigurationInterface
type FakeAzurermMysqlConfigurations struct {
	Fake *FakeAzurermV1alpha1
}

var azurermmysqlconfigurationsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermmysqlconfigurations"}

var azurermmysqlconfigurationsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermMysqlConfiguration"}

// Get takes name of the azurermMysqlConfiguration, and returns the corresponding azurermMysqlConfiguration object, and an error if there is any.
func (c *FakeAzurermMysqlConfigurations) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermMysqlConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermmysqlconfigurationsResource, name), &v1alpha1.AzurermMysqlConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMysqlConfiguration), err
}

// List takes label and field selectors, and returns the list of AzurermMysqlConfigurations that match those selectors.
func (c *FakeAzurermMysqlConfigurations) List(opts v1.ListOptions) (result *v1alpha1.AzurermMysqlConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermmysqlconfigurationsResource, azurermmysqlconfigurationsKind, opts), &v1alpha1.AzurermMysqlConfigurationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermMysqlConfigurationList{ListMeta: obj.(*v1alpha1.AzurermMysqlConfigurationList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermMysqlConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermMysqlConfigurations.
func (c *FakeAzurermMysqlConfigurations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermmysqlconfigurationsResource, opts))
}

// Create takes the representation of a azurermMysqlConfiguration and creates it.  Returns the server's representation of the azurermMysqlConfiguration, and an error, if there is any.
func (c *FakeAzurermMysqlConfigurations) Create(azurermMysqlConfiguration *v1alpha1.AzurermMysqlConfiguration) (result *v1alpha1.AzurermMysqlConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermmysqlconfigurationsResource, azurermMysqlConfiguration), &v1alpha1.AzurermMysqlConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMysqlConfiguration), err
}

// Update takes the representation of a azurermMysqlConfiguration and updates it. Returns the server's representation of the azurermMysqlConfiguration, and an error, if there is any.
func (c *FakeAzurermMysqlConfigurations) Update(azurermMysqlConfiguration *v1alpha1.AzurermMysqlConfiguration) (result *v1alpha1.AzurermMysqlConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermmysqlconfigurationsResource, azurermMysqlConfiguration), &v1alpha1.AzurermMysqlConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMysqlConfiguration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermMysqlConfigurations) UpdateStatus(azurermMysqlConfiguration *v1alpha1.AzurermMysqlConfiguration) (*v1alpha1.AzurermMysqlConfiguration, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermmysqlconfigurationsResource, "status", azurermMysqlConfiguration), &v1alpha1.AzurermMysqlConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMysqlConfiguration), err
}

// Delete takes name of the azurermMysqlConfiguration and deletes it. Returns an error if one occurs.
func (c *FakeAzurermMysqlConfigurations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermmysqlconfigurationsResource, name), &v1alpha1.AzurermMysqlConfiguration{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermMysqlConfigurations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermmysqlconfigurationsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermMysqlConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched azurermMysqlConfiguration.
func (c *FakeAzurermMysqlConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermMysqlConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermmysqlconfigurationsResource, name, pt, data, subresources...), &v1alpha1.AzurermMysqlConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMysqlConfiguration), err
}
