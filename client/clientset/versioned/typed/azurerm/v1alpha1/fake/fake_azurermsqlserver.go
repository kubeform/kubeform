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

// FakeAzurermSqlServers implements AzurermSqlServerInterface
type FakeAzurermSqlServers struct {
	Fake *FakeAzurermV1alpha1
}

var azurermsqlserversResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermsqlservers"}

var azurermsqlserversKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermSqlServer"}

// Get takes name of the azurermSqlServer, and returns the corresponding azurermSqlServer object, and an error if there is any.
func (c *FakeAzurermSqlServers) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermSqlServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermsqlserversResource, name), &v1alpha1.AzurermSqlServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermSqlServer), err
}

// List takes label and field selectors, and returns the list of AzurermSqlServers that match those selectors.
func (c *FakeAzurermSqlServers) List(opts v1.ListOptions) (result *v1alpha1.AzurermSqlServerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermsqlserversResource, azurermsqlserversKind, opts), &v1alpha1.AzurermSqlServerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermSqlServerList{ListMeta: obj.(*v1alpha1.AzurermSqlServerList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermSqlServerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermSqlServers.
func (c *FakeAzurermSqlServers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermsqlserversResource, opts))
}

// Create takes the representation of a azurermSqlServer and creates it.  Returns the server's representation of the azurermSqlServer, and an error, if there is any.
func (c *FakeAzurermSqlServers) Create(azurermSqlServer *v1alpha1.AzurermSqlServer) (result *v1alpha1.AzurermSqlServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermsqlserversResource, azurermSqlServer), &v1alpha1.AzurermSqlServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermSqlServer), err
}

// Update takes the representation of a azurermSqlServer and updates it. Returns the server's representation of the azurermSqlServer, and an error, if there is any.
func (c *FakeAzurermSqlServers) Update(azurermSqlServer *v1alpha1.AzurermSqlServer) (result *v1alpha1.AzurermSqlServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermsqlserversResource, azurermSqlServer), &v1alpha1.AzurermSqlServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermSqlServer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermSqlServers) UpdateStatus(azurermSqlServer *v1alpha1.AzurermSqlServer) (*v1alpha1.AzurermSqlServer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermsqlserversResource, "status", azurermSqlServer), &v1alpha1.AzurermSqlServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermSqlServer), err
}

// Delete takes name of the azurermSqlServer and deletes it. Returns an error if one occurs.
func (c *FakeAzurermSqlServers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermsqlserversResource, name), &v1alpha1.AzurermSqlServer{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermSqlServers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermsqlserversResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermSqlServerList{})
	return err
}

// Patch applies the patch and returns the patched azurermSqlServer.
func (c *FakeAzurermSqlServers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermSqlServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermsqlserversResource, name, pt, data, subresources...), &v1alpha1.AzurermSqlServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermSqlServer), err
}
