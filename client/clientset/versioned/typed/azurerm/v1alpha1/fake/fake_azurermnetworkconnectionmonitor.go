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

// FakeAzurermNetworkConnectionMonitors implements AzurermNetworkConnectionMonitorInterface
type FakeAzurermNetworkConnectionMonitors struct {
	Fake *FakeAzurermV1alpha1
}

var azurermnetworkconnectionmonitorsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermnetworkconnectionmonitors"}

var azurermnetworkconnectionmonitorsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermNetworkConnectionMonitor"}

// Get takes name of the azurermNetworkConnectionMonitor, and returns the corresponding azurermNetworkConnectionMonitor object, and an error if there is any.
func (c *FakeAzurermNetworkConnectionMonitors) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermNetworkConnectionMonitor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermnetworkconnectionmonitorsResource, name), &v1alpha1.AzurermNetworkConnectionMonitor{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermNetworkConnectionMonitor), err
}

// List takes label and field selectors, and returns the list of AzurermNetworkConnectionMonitors that match those selectors.
func (c *FakeAzurermNetworkConnectionMonitors) List(opts v1.ListOptions) (result *v1alpha1.AzurermNetworkConnectionMonitorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermnetworkconnectionmonitorsResource, azurermnetworkconnectionmonitorsKind, opts), &v1alpha1.AzurermNetworkConnectionMonitorList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermNetworkConnectionMonitorList{ListMeta: obj.(*v1alpha1.AzurermNetworkConnectionMonitorList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermNetworkConnectionMonitorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermNetworkConnectionMonitors.
func (c *FakeAzurermNetworkConnectionMonitors) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermnetworkconnectionmonitorsResource, opts))
}

// Create takes the representation of a azurermNetworkConnectionMonitor and creates it.  Returns the server's representation of the azurermNetworkConnectionMonitor, and an error, if there is any.
func (c *FakeAzurermNetworkConnectionMonitors) Create(azurermNetworkConnectionMonitor *v1alpha1.AzurermNetworkConnectionMonitor) (result *v1alpha1.AzurermNetworkConnectionMonitor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermnetworkconnectionmonitorsResource, azurermNetworkConnectionMonitor), &v1alpha1.AzurermNetworkConnectionMonitor{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermNetworkConnectionMonitor), err
}

// Update takes the representation of a azurermNetworkConnectionMonitor and updates it. Returns the server's representation of the azurermNetworkConnectionMonitor, and an error, if there is any.
func (c *FakeAzurermNetworkConnectionMonitors) Update(azurermNetworkConnectionMonitor *v1alpha1.AzurermNetworkConnectionMonitor) (result *v1alpha1.AzurermNetworkConnectionMonitor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermnetworkconnectionmonitorsResource, azurermNetworkConnectionMonitor), &v1alpha1.AzurermNetworkConnectionMonitor{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermNetworkConnectionMonitor), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermNetworkConnectionMonitors) UpdateStatus(azurermNetworkConnectionMonitor *v1alpha1.AzurermNetworkConnectionMonitor) (*v1alpha1.AzurermNetworkConnectionMonitor, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermnetworkconnectionmonitorsResource, "status", azurermNetworkConnectionMonitor), &v1alpha1.AzurermNetworkConnectionMonitor{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermNetworkConnectionMonitor), err
}

// Delete takes name of the azurermNetworkConnectionMonitor and deletes it. Returns an error if one occurs.
func (c *FakeAzurermNetworkConnectionMonitors) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermnetworkconnectionmonitorsResource, name), &v1alpha1.AzurermNetworkConnectionMonitor{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermNetworkConnectionMonitors) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermnetworkconnectionmonitorsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermNetworkConnectionMonitorList{})
	return err
}

// Patch applies the patch and returns the patched azurermNetworkConnectionMonitor.
func (c *FakeAzurermNetworkConnectionMonitors) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermNetworkConnectionMonitor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermnetworkconnectionmonitorsResource, name, pt, data, subresources...), &v1alpha1.AzurermNetworkConnectionMonitor{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermNetworkConnectionMonitor), err
}
