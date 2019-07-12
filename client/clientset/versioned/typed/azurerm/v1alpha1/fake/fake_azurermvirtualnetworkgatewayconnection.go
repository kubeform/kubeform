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

// FakeAzurermVirtualNetworkGatewayConnections implements AzurermVirtualNetworkGatewayConnectionInterface
type FakeAzurermVirtualNetworkGatewayConnections struct {
	Fake *FakeAzurermV1alpha1
}

var azurermvirtualnetworkgatewayconnectionsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermvirtualnetworkgatewayconnections"}

var azurermvirtualnetworkgatewayconnectionsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermVirtualNetworkGatewayConnection"}

// Get takes name of the azurermVirtualNetworkGatewayConnection, and returns the corresponding azurermVirtualNetworkGatewayConnection object, and an error if there is any.
func (c *FakeAzurermVirtualNetworkGatewayConnections) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermVirtualNetworkGatewayConnection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermvirtualnetworkgatewayconnectionsResource, name), &v1alpha1.AzurermVirtualNetworkGatewayConnection{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermVirtualNetworkGatewayConnection), err
}

// List takes label and field selectors, and returns the list of AzurermVirtualNetworkGatewayConnections that match those selectors.
func (c *FakeAzurermVirtualNetworkGatewayConnections) List(opts v1.ListOptions) (result *v1alpha1.AzurermVirtualNetworkGatewayConnectionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermvirtualnetworkgatewayconnectionsResource, azurermvirtualnetworkgatewayconnectionsKind, opts), &v1alpha1.AzurermVirtualNetworkGatewayConnectionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermVirtualNetworkGatewayConnectionList{ListMeta: obj.(*v1alpha1.AzurermVirtualNetworkGatewayConnectionList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermVirtualNetworkGatewayConnectionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermVirtualNetworkGatewayConnections.
func (c *FakeAzurermVirtualNetworkGatewayConnections) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermvirtualnetworkgatewayconnectionsResource, opts))
}

// Create takes the representation of a azurermVirtualNetworkGatewayConnection and creates it.  Returns the server's representation of the azurermVirtualNetworkGatewayConnection, and an error, if there is any.
func (c *FakeAzurermVirtualNetworkGatewayConnections) Create(azurermVirtualNetworkGatewayConnection *v1alpha1.AzurermVirtualNetworkGatewayConnection) (result *v1alpha1.AzurermVirtualNetworkGatewayConnection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermvirtualnetworkgatewayconnectionsResource, azurermVirtualNetworkGatewayConnection), &v1alpha1.AzurermVirtualNetworkGatewayConnection{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermVirtualNetworkGatewayConnection), err
}

// Update takes the representation of a azurermVirtualNetworkGatewayConnection and updates it. Returns the server's representation of the azurermVirtualNetworkGatewayConnection, and an error, if there is any.
func (c *FakeAzurermVirtualNetworkGatewayConnections) Update(azurermVirtualNetworkGatewayConnection *v1alpha1.AzurermVirtualNetworkGatewayConnection) (result *v1alpha1.AzurermVirtualNetworkGatewayConnection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermvirtualnetworkgatewayconnectionsResource, azurermVirtualNetworkGatewayConnection), &v1alpha1.AzurermVirtualNetworkGatewayConnection{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermVirtualNetworkGatewayConnection), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermVirtualNetworkGatewayConnections) UpdateStatus(azurermVirtualNetworkGatewayConnection *v1alpha1.AzurermVirtualNetworkGatewayConnection) (*v1alpha1.AzurermVirtualNetworkGatewayConnection, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermvirtualnetworkgatewayconnectionsResource, "status", azurermVirtualNetworkGatewayConnection), &v1alpha1.AzurermVirtualNetworkGatewayConnection{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermVirtualNetworkGatewayConnection), err
}

// Delete takes name of the azurermVirtualNetworkGatewayConnection and deletes it. Returns an error if one occurs.
func (c *FakeAzurermVirtualNetworkGatewayConnections) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermvirtualnetworkgatewayconnectionsResource, name), &v1alpha1.AzurermVirtualNetworkGatewayConnection{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermVirtualNetworkGatewayConnections) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermvirtualnetworkgatewayconnectionsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermVirtualNetworkGatewayConnectionList{})
	return err
}

// Patch applies the patch and returns the patched azurermVirtualNetworkGatewayConnection.
func (c *FakeAzurermVirtualNetworkGatewayConnections) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermVirtualNetworkGatewayConnection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermvirtualnetworkgatewayconnectionsResource, name, pt, data, subresources...), &v1alpha1.AzurermVirtualNetworkGatewayConnection{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermVirtualNetworkGatewayConnection), err
}
