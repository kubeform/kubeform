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

// FakeAzurermMariadbServers implements AzurermMariadbServerInterface
type FakeAzurermMariadbServers struct {
	Fake *FakeAzurermV1alpha1
}

var azurermmariadbserversResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermmariadbservers"}

var azurermmariadbserversKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermMariadbServer"}

// Get takes name of the azurermMariadbServer, and returns the corresponding azurermMariadbServer object, and an error if there is any.
func (c *FakeAzurermMariadbServers) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermMariadbServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermmariadbserversResource, name), &v1alpha1.AzurermMariadbServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMariadbServer), err
}

// List takes label and field selectors, and returns the list of AzurermMariadbServers that match those selectors.
func (c *FakeAzurermMariadbServers) List(opts v1.ListOptions) (result *v1alpha1.AzurermMariadbServerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermmariadbserversResource, azurermmariadbserversKind, opts), &v1alpha1.AzurermMariadbServerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermMariadbServerList{ListMeta: obj.(*v1alpha1.AzurermMariadbServerList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermMariadbServerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermMariadbServers.
func (c *FakeAzurermMariadbServers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermmariadbserversResource, opts))
}

// Create takes the representation of a azurermMariadbServer and creates it.  Returns the server's representation of the azurermMariadbServer, and an error, if there is any.
func (c *FakeAzurermMariadbServers) Create(azurermMariadbServer *v1alpha1.AzurermMariadbServer) (result *v1alpha1.AzurermMariadbServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermmariadbserversResource, azurermMariadbServer), &v1alpha1.AzurermMariadbServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMariadbServer), err
}

// Update takes the representation of a azurermMariadbServer and updates it. Returns the server's representation of the azurermMariadbServer, and an error, if there is any.
func (c *FakeAzurermMariadbServers) Update(azurermMariadbServer *v1alpha1.AzurermMariadbServer) (result *v1alpha1.AzurermMariadbServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermmariadbserversResource, azurermMariadbServer), &v1alpha1.AzurermMariadbServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMariadbServer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermMariadbServers) UpdateStatus(azurermMariadbServer *v1alpha1.AzurermMariadbServer) (*v1alpha1.AzurermMariadbServer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermmariadbserversResource, "status", azurermMariadbServer), &v1alpha1.AzurermMariadbServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMariadbServer), err
}

// Delete takes name of the azurermMariadbServer and deletes it. Returns an error if one occurs.
func (c *FakeAzurermMariadbServers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermmariadbserversResource, name), &v1alpha1.AzurermMariadbServer{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermMariadbServers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermmariadbserversResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermMariadbServerList{})
	return err
}

// Patch applies the patch and returns the patched azurermMariadbServer.
func (c *FakeAzurermMariadbServers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermMariadbServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermmariadbserversResource, name, pt, data, subresources...), &v1alpha1.AzurermMariadbServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMariadbServer), err
}
