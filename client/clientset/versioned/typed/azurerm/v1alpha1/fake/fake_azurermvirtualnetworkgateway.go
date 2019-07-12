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

// FakeAzurermVirtualNetworkGateways implements AzurermVirtualNetworkGatewayInterface
type FakeAzurermVirtualNetworkGateways struct {
	Fake *FakeAzurermV1alpha1
}

var azurermvirtualnetworkgatewaysResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermvirtualnetworkgateways"}

var azurermvirtualnetworkgatewaysKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermVirtualNetworkGateway"}

// Get takes name of the azurermVirtualNetworkGateway, and returns the corresponding azurermVirtualNetworkGateway object, and an error if there is any.
func (c *FakeAzurermVirtualNetworkGateways) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermVirtualNetworkGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermvirtualnetworkgatewaysResource, name), &v1alpha1.AzurermVirtualNetworkGateway{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermVirtualNetworkGateway), err
}

// List takes label and field selectors, and returns the list of AzurermVirtualNetworkGateways that match those selectors.
func (c *FakeAzurermVirtualNetworkGateways) List(opts v1.ListOptions) (result *v1alpha1.AzurermVirtualNetworkGatewayList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermvirtualnetworkgatewaysResource, azurermvirtualnetworkgatewaysKind, opts), &v1alpha1.AzurermVirtualNetworkGatewayList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermVirtualNetworkGatewayList{ListMeta: obj.(*v1alpha1.AzurermVirtualNetworkGatewayList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermVirtualNetworkGatewayList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermVirtualNetworkGateways.
func (c *FakeAzurermVirtualNetworkGateways) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermvirtualnetworkgatewaysResource, opts))
}

// Create takes the representation of a azurermVirtualNetworkGateway and creates it.  Returns the server's representation of the azurermVirtualNetworkGateway, and an error, if there is any.
func (c *FakeAzurermVirtualNetworkGateways) Create(azurermVirtualNetworkGateway *v1alpha1.AzurermVirtualNetworkGateway) (result *v1alpha1.AzurermVirtualNetworkGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermvirtualnetworkgatewaysResource, azurermVirtualNetworkGateway), &v1alpha1.AzurermVirtualNetworkGateway{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermVirtualNetworkGateway), err
}

// Update takes the representation of a azurermVirtualNetworkGateway and updates it. Returns the server's representation of the azurermVirtualNetworkGateway, and an error, if there is any.
func (c *FakeAzurermVirtualNetworkGateways) Update(azurermVirtualNetworkGateway *v1alpha1.AzurermVirtualNetworkGateway) (result *v1alpha1.AzurermVirtualNetworkGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermvirtualnetworkgatewaysResource, azurermVirtualNetworkGateway), &v1alpha1.AzurermVirtualNetworkGateway{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermVirtualNetworkGateway), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermVirtualNetworkGateways) UpdateStatus(azurermVirtualNetworkGateway *v1alpha1.AzurermVirtualNetworkGateway) (*v1alpha1.AzurermVirtualNetworkGateway, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermvirtualnetworkgatewaysResource, "status", azurermVirtualNetworkGateway), &v1alpha1.AzurermVirtualNetworkGateway{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermVirtualNetworkGateway), err
}

// Delete takes name of the azurermVirtualNetworkGateway and deletes it. Returns an error if one occurs.
func (c *FakeAzurermVirtualNetworkGateways) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermvirtualnetworkgatewaysResource, name), &v1alpha1.AzurermVirtualNetworkGateway{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermVirtualNetworkGateways) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermvirtualnetworkgatewaysResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermVirtualNetworkGatewayList{})
	return err
}

// Patch applies the patch and returns the patched azurermVirtualNetworkGateway.
func (c *FakeAzurermVirtualNetworkGateways) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermVirtualNetworkGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermvirtualnetworkgatewaysResource, name, pt, data, subresources...), &v1alpha1.AzurermVirtualNetworkGateway{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermVirtualNetworkGateway), err
}
