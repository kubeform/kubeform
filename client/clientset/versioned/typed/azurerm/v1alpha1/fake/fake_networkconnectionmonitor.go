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

// FakeNetworkConnectionMonitors implements NetworkConnectionMonitorInterface
type FakeNetworkConnectionMonitors struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var networkconnectionmonitorsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "networkconnectionmonitors"}

var networkconnectionmonitorsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "NetworkConnectionMonitor"}

// Get takes name of the networkConnectionMonitor, and returns the corresponding networkConnectionMonitor object, and an error if there is any.
func (c *FakeNetworkConnectionMonitors) Get(name string, options v1.GetOptions) (result *v1alpha1.NetworkConnectionMonitor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkconnectionmonitorsResource, c.ns, name), &v1alpha1.NetworkConnectionMonitor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkConnectionMonitor), err
}

// List takes label and field selectors, and returns the list of NetworkConnectionMonitors that match those selectors.
func (c *FakeNetworkConnectionMonitors) List(opts v1.ListOptions) (result *v1alpha1.NetworkConnectionMonitorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkconnectionmonitorsResource, networkconnectionmonitorsKind, c.ns, opts), &v1alpha1.NetworkConnectionMonitorList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NetworkConnectionMonitorList{ListMeta: obj.(*v1alpha1.NetworkConnectionMonitorList).ListMeta}
	for _, item := range obj.(*v1alpha1.NetworkConnectionMonitorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkConnectionMonitors.
func (c *FakeNetworkConnectionMonitors) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkconnectionmonitorsResource, c.ns, opts))

}

// Create takes the representation of a networkConnectionMonitor and creates it.  Returns the server's representation of the networkConnectionMonitor, and an error, if there is any.
func (c *FakeNetworkConnectionMonitors) Create(networkConnectionMonitor *v1alpha1.NetworkConnectionMonitor) (result *v1alpha1.NetworkConnectionMonitor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkconnectionmonitorsResource, c.ns, networkConnectionMonitor), &v1alpha1.NetworkConnectionMonitor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkConnectionMonitor), err
}

// Update takes the representation of a networkConnectionMonitor and updates it. Returns the server's representation of the networkConnectionMonitor, and an error, if there is any.
func (c *FakeNetworkConnectionMonitors) Update(networkConnectionMonitor *v1alpha1.NetworkConnectionMonitor) (result *v1alpha1.NetworkConnectionMonitor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkconnectionmonitorsResource, c.ns, networkConnectionMonitor), &v1alpha1.NetworkConnectionMonitor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkConnectionMonitor), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkConnectionMonitors) UpdateStatus(networkConnectionMonitor *v1alpha1.NetworkConnectionMonitor) (*v1alpha1.NetworkConnectionMonitor, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networkconnectionmonitorsResource, "status", c.ns, networkConnectionMonitor), &v1alpha1.NetworkConnectionMonitor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkConnectionMonitor), err
}

// Delete takes name of the networkConnectionMonitor and deletes it. Returns an error if one occurs.
func (c *FakeNetworkConnectionMonitors) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(networkconnectionmonitorsResource, c.ns, name), &v1alpha1.NetworkConnectionMonitor{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkConnectionMonitors) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkconnectionmonitorsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.NetworkConnectionMonitorList{})
	return err
}

// Patch applies the patch and returns the patched networkConnectionMonitor.
func (c *FakeNetworkConnectionMonitors) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetworkConnectionMonitor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkconnectionmonitorsResource, c.ns, name, pt, data, subresources...), &v1alpha1.NetworkConnectionMonitor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkConnectionMonitor), err
}