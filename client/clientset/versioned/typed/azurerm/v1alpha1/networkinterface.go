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

// NetworkInterfacesGetter has a method to return a NetworkInterfaceInterface.
// A group's client should implement this interface.
type NetworkInterfacesGetter interface {
	NetworkInterfaces(namespace string) NetworkInterfaceInterface
}

// NetworkInterfaceInterface has methods to work with NetworkInterface resources.
type NetworkInterfaceInterface interface {
	Create(*v1alpha1.NetworkInterface) (*v1alpha1.NetworkInterface, error)
	Update(*v1alpha1.NetworkInterface) (*v1alpha1.NetworkInterface, error)
	UpdateStatus(*v1alpha1.NetworkInterface) (*v1alpha1.NetworkInterface, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.NetworkInterface, error)
	List(opts v1.ListOptions) (*v1alpha1.NetworkInterfaceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetworkInterface, err error)
	NetworkInterfaceExpansion
}

// networkInterfaces implements NetworkInterfaceInterface
type networkInterfaces struct {
	client rest.Interface
	ns     string
}

// newNetworkInterfaces returns a NetworkInterfaces
func newNetworkInterfaces(c *AzurermV1alpha1Client, namespace string) *networkInterfaces {
	return &networkInterfaces{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the networkInterface, and returns the corresponding networkInterface object, and an error if there is any.
func (c *networkInterfaces) Get(name string, options v1.GetOptions) (result *v1alpha1.NetworkInterface, err error) {
	result = &v1alpha1.NetworkInterface{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkinterfaces").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NetworkInterfaces that match those selectors.
func (c *networkInterfaces) List(opts v1.ListOptions) (result *v1alpha1.NetworkInterfaceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NetworkInterfaceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkinterfaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested networkInterfaces.
func (c *networkInterfaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("networkinterfaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a networkInterface and creates it.  Returns the server's representation of the networkInterface, and an error, if there is any.
func (c *networkInterfaces) Create(networkInterface *v1alpha1.NetworkInterface) (result *v1alpha1.NetworkInterface, err error) {
	result = &v1alpha1.NetworkInterface{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("networkinterfaces").
		Body(networkInterface).
		Do().
		Into(result)
	return
}

// Update takes the representation of a networkInterface and updates it. Returns the server's representation of the networkInterface, and an error, if there is any.
func (c *networkInterfaces) Update(networkInterface *v1alpha1.NetworkInterface) (result *v1alpha1.NetworkInterface, err error) {
	result = &v1alpha1.NetworkInterface{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networkinterfaces").
		Name(networkInterface.Name).
		Body(networkInterface).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *networkInterfaces) UpdateStatus(networkInterface *v1alpha1.NetworkInterface) (result *v1alpha1.NetworkInterface, err error) {
	result = &v1alpha1.NetworkInterface{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networkinterfaces").
		Name(networkInterface.Name).
		SubResource("status").
		Body(networkInterface).
		Do().
		Into(result)
	return
}

// Delete takes name of the networkInterface and deletes it. Returns an error if one occurs.
func (c *networkInterfaces) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkinterfaces").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *networkInterfaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkinterfaces").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched networkInterface.
func (c *networkInterfaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetworkInterface, err error) {
	result = &v1alpha1.NetworkInterface{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("networkinterfaces").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}