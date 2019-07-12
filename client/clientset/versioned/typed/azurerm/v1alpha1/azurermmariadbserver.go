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

// AzurermMariadbServersGetter has a method to return a AzurermMariadbServerInterface.
// A group's client should implement this interface.
type AzurermMariadbServersGetter interface {
	AzurermMariadbServers() AzurermMariadbServerInterface
}

// AzurermMariadbServerInterface has methods to work with AzurermMariadbServer resources.
type AzurermMariadbServerInterface interface {
	Create(*v1alpha1.AzurermMariadbServer) (*v1alpha1.AzurermMariadbServer, error)
	Update(*v1alpha1.AzurermMariadbServer) (*v1alpha1.AzurermMariadbServer, error)
	UpdateStatus(*v1alpha1.AzurermMariadbServer) (*v1alpha1.AzurermMariadbServer, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermMariadbServer, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermMariadbServerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermMariadbServer, err error)
	AzurermMariadbServerExpansion
}

// azurermMariadbServers implements AzurermMariadbServerInterface
type azurermMariadbServers struct {
	client rest.Interface
}

// newAzurermMariadbServers returns a AzurermMariadbServers
func newAzurermMariadbServers(c *AzurermV1alpha1Client) *azurermMariadbServers {
	return &azurermMariadbServers{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermMariadbServer, and returns the corresponding azurermMariadbServer object, and an error if there is any.
func (c *azurermMariadbServers) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermMariadbServer, err error) {
	result = &v1alpha1.AzurermMariadbServer{}
	err = c.client.Get().
		Resource("azurermmariadbservers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermMariadbServers that match those selectors.
func (c *azurermMariadbServers) List(opts v1.ListOptions) (result *v1alpha1.AzurermMariadbServerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermMariadbServerList{}
	err = c.client.Get().
		Resource("azurermmariadbservers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermMariadbServers.
func (c *azurermMariadbServers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermmariadbservers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermMariadbServer and creates it.  Returns the server's representation of the azurermMariadbServer, and an error, if there is any.
func (c *azurermMariadbServers) Create(azurermMariadbServer *v1alpha1.AzurermMariadbServer) (result *v1alpha1.AzurermMariadbServer, err error) {
	result = &v1alpha1.AzurermMariadbServer{}
	err = c.client.Post().
		Resource("azurermmariadbservers").
		Body(azurermMariadbServer).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermMariadbServer and updates it. Returns the server's representation of the azurermMariadbServer, and an error, if there is any.
func (c *azurermMariadbServers) Update(azurermMariadbServer *v1alpha1.AzurermMariadbServer) (result *v1alpha1.AzurermMariadbServer, err error) {
	result = &v1alpha1.AzurermMariadbServer{}
	err = c.client.Put().
		Resource("azurermmariadbservers").
		Name(azurermMariadbServer.Name).
		Body(azurermMariadbServer).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermMariadbServers) UpdateStatus(azurermMariadbServer *v1alpha1.AzurermMariadbServer) (result *v1alpha1.AzurermMariadbServer, err error) {
	result = &v1alpha1.AzurermMariadbServer{}
	err = c.client.Put().
		Resource("azurermmariadbservers").
		Name(azurermMariadbServer.Name).
		SubResource("status").
		Body(azurermMariadbServer).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermMariadbServer and deletes it. Returns an error if one occurs.
func (c *azurermMariadbServers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermmariadbservers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermMariadbServers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermmariadbservers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermMariadbServer.
func (c *azurermMariadbServers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermMariadbServer, err error) {
	result = &v1alpha1.AzurermMariadbServer{}
	err = c.client.Patch(pt).
		Resource("azurermmariadbservers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
