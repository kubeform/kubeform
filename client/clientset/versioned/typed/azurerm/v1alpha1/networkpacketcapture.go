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

// NetworkPacketCapturesGetter has a method to return a NetworkPacketCaptureInterface.
// A group's client should implement this interface.
type NetworkPacketCapturesGetter interface {
	NetworkPacketCaptures() NetworkPacketCaptureInterface
}

// NetworkPacketCaptureInterface has methods to work with NetworkPacketCapture resources.
type NetworkPacketCaptureInterface interface {
	Create(*v1alpha1.NetworkPacketCapture) (*v1alpha1.NetworkPacketCapture, error)
	Update(*v1alpha1.NetworkPacketCapture) (*v1alpha1.NetworkPacketCapture, error)
	UpdateStatus(*v1alpha1.NetworkPacketCapture) (*v1alpha1.NetworkPacketCapture, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.NetworkPacketCapture, error)
	List(opts v1.ListOptions) (*v1alpha1.NetworkPacketCaptureList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetworkPacketCapture, err error)
	NetworkPacketCaptureExpansion
}

// networkPacketCaptures implements NetworkPacketCaptureInterface
type networkPacketCaptures struct {
	client rest.Interface
}

// newNetworkPacketCaptures returns a NetworkPacketCaptures
func newNetworkPacketCaptures(c *AzurermV1alpha1Client) *networkPacketCaptures {
	return &networkPacketCaptures{
		client: c.RESTClient(),
	}
}

// Get takes name of the networkPacketCapture, and returns the corresponding networkPacketCapture object, and an error if there is any.
func (c *networkPacketCaptures) Get(name string, options v1.GetOptions) (result *v1alpha1.NetworkPacketCapture, err error) {
	result = &v1alpha1.NetworkPacketCapture{}
	err = c.client.Get().
		Resource("networkpacketcaptures").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NetworkPacketCaptures that match those selectors.
func (c *networkPacketCaptures) List(opts v1.ListOptions) (result *v1alpha1.NetworkPacketCaptureList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NetworkPacketCaptureList{}
	err = c.client.Get().
		Resource("networkpacketcaptures").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested networkPacketCaptures.
func (c *networkPacketCaptures) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("networkpacketcaptures").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a networkPacketCapture and creates it.  Returns the server's representation of the networkPacketCapture, and an error, if there is any.
func (c *networkPacketCaptures) Create(networkPacketCapture *v1alpha1.NetworkPacketCapture) (result *v1alpha1.NetworkPacketCapture, err error) {
	result = &v1alpha1.NetworkPacketCapture{}
	err = c.client.Post().
		Resource("networkpacketcaptures").
		Body(networkPacketCapture).
		Do().
		Into(result)
	return
}

// Update takes the representation of a networkPacketCapture and updates it. Returns the server's representation of the networkPacketCapture, and an error, if there is any.
func (c *networkPacketCaptures) Update(networkPacketCapture *v1alpha1.NetworkPacketCapture) (result *v1alpha1.NetworkPacketCapture, err error) {
	result = &v1alpha1.NetworkPacketCapture{}
	err = c.client.Put().
		Resource("networkpacketcaptures").
		Name(networkPacketCapture.Name).
		Body(networkPacketCapture).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *networkPacketCaptures) UpdateStatus(networkPacketCapture *v1alpha1.NetworkPacketCapture) (result *v1alpha1.NetworkPacketCapture, err error) {
	result = &v1alpha1.NetworkPacketCapture{}
	err = c.client.Put().
		Resource("networkpacketcaptures").
		Name(networkPacketCapture.Name).
		SubResource("status").
		Body(networkPacketCapture).
		Do().
		Into(result)
	return
}

// Delete takes name of the networkPacketCapture and deletes it. Returns an error if one occurs.
func (c *networkPacketCaptures) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("networkpacketcaptures").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *networkPacketCaptures) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("networkpacketcaptures").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched networkPacketCapture.
func (c *networkPacketCaptures) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetworkPacketCapture, err error) {
	result = &v1alpha1.NetworkPacketCapture{}
	err = c.client.Patch(pt).
		Resource("networkpacketcaptures").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
