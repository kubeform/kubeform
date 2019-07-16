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

// DevTestVirtualNetworksGetter has a method to return a DevTestVirtualNetworkInterface.
// A group's client should implement this interface.
type DevTestVirtualNetworksGetter interface {
	DevTestVirtualNetworks() DevTestVirtualNetworkInterface
}

// DevTestVirtualNetworkInterface has methods to work with DevTestVirtualNetwork resources.
type DevTestVirtualNetworkInterface interface {
	Create(*v1alpha1.DevTestVirtualNetwork) (*v1alpha1.DevTestVirtualNetwork, error)
	Update(*v1alpha1.DevTestVirtualNetwork) (*v1alpha1.DevTestVirtualNetwork, error)
	UpdateStatus(*v1alpha1.DevTestVirtualNetwork) (*v1alpha1.DevTestVirtualNetwork, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DevTestVirtualNetwork, error)
	List(opts v1.ListOptions) (*v1alpha1.DevTestVirtualNetworkList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DevTestVirtualNetwork, err error)
	DevTestVirtualNetworkExpansion
}

// devTestVirtualNetworks implements DevTestVirtualNetworkInterface
type devTestVirtualNetworks struct {
	client rest.Interface
}

// newDevTestVirtualNetworks returns a DevTestVirtualNetworks
func newDevTestVirtualNetworks(c *AzurermV1alpha1Client) *devTestVirtualNetworks {
	return &devTestVirtualNetworks{
		client: c.RESTClient(),
	}
}

// Get takes name of the devTestVirtualNetwork, and returns the corresponding devTestVirtualNetwork object, and an error if there is any.
func (c *devTestVirtualNetworks) Get(name string, options v1.GetOptions) (result *v1alpha1.DevTestVirtualNetwork, err error) {
	result = &v1alpha1.DevTestVirtualNetwork{}
	err = c.client.Get().
		Resource("devtestvirtualnetworks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DevTestVirtualNetworks that match those selectors.
func (c *devTestVirtualNetworks) List(opts v1.ListOptions) (result *v1alpha1.DevTestVirtualNetworkList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DevTestVirtualNetworkList{}
	err = c.client.Get().
		Resource("devtestvirtualnetworks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested devTestVirtualNetworks.
func (c *devTestVirtualNetworks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("devtestvirtualnetworks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a devTestVirtualNetwork and creates it.  Returns the server's representation of the devTestVirtualNetwork, and an error, if there is any.
func (c *devTestVirtualNetworks) Create(devTestVirtualNetwork *v1alpha1.DevTestVirtualNetwork) (result *v1alpha1.DevTestVirtualNetwork, err error) {
	result = &v1alpha1.DevTestVirtualNetwork{}
	err = c.client.Post().
		Resource("devtestvirtualnetworks").
		Body(devTestVirtualNetwork).
		Do().
		Into(result)
	return
}

// Update takes the representation of a devTestVirtualNetwork and updates it. Returns the server's representation of the devTestVirtualNetwork, and an error, if there is any.
func (c *devTestVirtualNetworks) Update(devTestVirtualNetwork *v1alpha1.DevTestVirtualNetwork) (result *v1alpha1.DevTestVirtualNetwork, err error) {
	result = &v1alpha1.DevTestVirtualNetwork{}
	err = c.client.Put().
		Resource("devtestvirtualnetworks").
		Name(devTestVirtualNetwork.Name).
		Body(devTestVirtualNetwork).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *devTestVirtualNetworks) UpdateStatus(devTestVirtualNetwork *v1alpha1.DevTestVirtualNetwork) (result *v1alpha1.DevTestVirtualNetwork, err error) {
	result = &v1alpha1.DevTestVirtualNetwork{}
	err = c.client.Put().
		Resource("devtestvirtualnetworks").
		Name(devTestVirtualNetwork.Name).
		SubResource("status").
		Body(devTestVirtualNetwork).
		Do().
		Into(result)
	return
}

// Delete takes name of the devTestVirtualNetwork and deletes it. Returns an error if one occurs.
func (c *devTestVirtualNetworks) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("devtestvirtualnetworks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *devTestVirtualNetworks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("devtestvirtualnetworks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched devTestVirtualNetwork.
func (c *devTestVirtualNetworks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DevTestVirtualNetwork, err error) {
	result = &v1alpha1.DevTestVirtualNetwork{}
	err = c.client.Patch(pt).
		Resource("devtestvirtualnetworks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
