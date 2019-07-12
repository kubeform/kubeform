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

package v1alpha1

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// AzurermNetworkConnectionMonitorsGetter has a method to return a AzurermNetworkConnectionMonitorInterface.
// A group's client should implement this interface.
type AzurermNetworkConnectionMonitorsGetter interface {
	AzurermNetworkConnectionMonitors() AzurermNetworkConnectionMonitorInterface
}

// AzurermNetworkConnectionMonitorInterface has methods to work with AzurermNetworkConnectionMonitor resources.
type AzurermNetworkConnectionMonitorInterface interface {
	Create(*v1alpha1.AzurermNetworkConnectionMonitor) (*v1alpha1.AzurermNetworkConnectionMonitor, error)
	Update(*v1alpha1.AzurermNetworkConnectionMonitor) (*v1alpha1.AzurermNetworkConnectionMonitor, error)
	UpdateStatus(*v1alpha1.AzurermNetworkConnectionMonitor) (*v1alpha1.AzurermNetworkConnectionMonitor, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermNetworkConnectionMonitor, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermNetworkConnectionMonitorList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermNetworkConnectionMonitor, err error)
	AzurermNetworkConnectionMonitorExpansion
}

// azurermNetworkConnectionMonitors implements AzurermNetworkConnectionMonitorInterface
type azurermNetworkConnectionMonitors struct {
	client rest.Interface
}

// newAzurermNetworkConnectionMonitors returns a AzurermNetworkConnectionMonitors
func newAzurermNetworkConnectionMonitors(c *AzurermV1alpha1Client) *azurermNetworkConnectionMonitors {
	return &azurermNetworkConnectionMonitors{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermNetworkConnectionMonitor, and returns the corresponding azurermNetworkConnectionMonitor object, and an error if there is any.
func (c *azurermNetworkConnectionMonitors) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermNetworkConnectionMonitor, err error) {
	result = &v1alpha1.AzurermNetworkConnectionMonitor{}
	err = c.client.Get().
		Resource("azurermnetworkconnectionmonitors").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermNetworkConnectionMonitors that match those selectors.
func (c *azurermNetworkConnectionMonitors) List(opts v1.ListOptions) (result *v1alpha1.AzurermNetworkConnectionMonitorList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermNetworkConnectionMonitorList{}
	err = c.client.Get().
		Resource("azurermnetworkconnectionmonitors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermNetworkConnectionMonitors.
func (c *azurermNetworkConnectionMonitors) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermnetworkconnectionmonitors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermNetworkConnectionMonitor and creates it.  Returns the server's representation of the azurermNetworkConnectionMonitor, and an error, if there is any.
func (c *azurermNetworkConnectionMonitors) Create(azurermNetworkConnectionMonitor *v1alpha1.AzurermNetworkConnectionMonitor) (result *v1alpha1.AzurermNetworkConnectionMonitor, err error) {
	result = &v1alpha1.AzurermNetworkConnectionMonitor{}
	err = c.client.Post().
		Resource("azurermnetworkconnectionmonitors").
		Body(azurermNetworkConnectionMonitor).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermNetworkConnectionMonitor and updates it. Returns the server's representation of the azurermNetworkConnectionMonitor, and an error, if there is any.
func (c *azurermNetworkConnectionMonitors) Update(azurermNetworkConnectionMonitor *v1alpha1.AzurermNetworkConnectionMonitor) (result *v1alpha1.AzurermNetworkConnectionMonitor, err error) {
	result = &v1alpha1.AzurermNetworkConnectionMonitor{}
	err = c.client.Put().
		Resource("azurermnetworkconnectionmonitors").
		Name(azurermNetworkConnectionMonitor.Name).
		Body(azurermNetworkConnectionMonitor).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermNetworkConnectionMonitors) UpdateStatus(azurermNetworkConnectionMonitor *v1alpha1.AzurermNetworkConnectionMonitor) (result *v1alpha1.AzurermNetworkConnectionMonitor, err error) {
	result = &v1alpha1.AzurermNetworkConnectionMonitor{}
	err = c.client.Put().
		Resource("azurermnetworkconnectionmonitors").
		Name(azurermNetworkConnectionMonitor.Name).
		SubResource("status").
		Body(azurermNetworkConnectionMonitor).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermNetworkConnectionMonitor and deletes it. Returns an error if one occurs.
func (c *azurermNetworkConnectionMonitors) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermnetworkconnectionmonitors").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermNetworkConnectionMonitors) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermnetworkconnectionmonitors").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermNetworkConnectionMonitor.
func (c *azurermNetworkConnectionMonitors) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermNetworkConnectionMonitor, err error) {
	result = &v1alpha1.AzurermNetworkConnectionMonitor{}
	err = c.client.Patch(pt).
		Resource("azurermnetworkconnectionmonitors").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
