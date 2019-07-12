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

// AzurermExpressRouteCircuitAuthorizationsGetter has a method to return a AzurermExpressRouteCircuitAuthorizationInterface.
// A group's client should implement this interface.
type AzurermExpressRouteCircuitAuthorizationsGetter interface {
	AzurermExpressRouteCircuitAuthorizations() AzurermExpressRouteCircuitAuthorizationInterface
}

// AzurermExpressRouteCircuitAuthorizationInterface has methods to work with AzurermExpressRouteCircuitAuthorization resources.
type AzurermExpressRouteCircuitAuthorizationInterface interface {
	Create(*v1alpha1.AzurermExpressRouteCircuitAuthorization) (*v1alpha1.AzurermExpressRouteCircuitAuthorization, error)
	Update(*v1alpha1.AzurermExpressRouteCircuitAuthorization) (*v1alpha1.AzurermExpressRouteCircuitAuthorization, error)
	UpdateStatus(*v1alpha1.AzurermExpressRouteCircuitAuthorization) (*v1alpha1.AzurermExpressRouteCircuitAuthorization, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermExpressRouteCircuitAuthorization, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermExpressRouteCircuitAuthorizationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermExpressRouteCircuitAuthorization, err error)
	AzurermExpressRouteCircuitAuthorizationExpansion
}

// azurermExpressRouteCircuitAuthorizations implements AzurermExpressRouteCircuitAuthorizationInterface
type azurermExpressRouteCircuitAuthorizations struct {
	client rest.Interface
}

// newAzurermExpressRouteCircuitAuthorizations returns a AzurermExpressRouteCircuitAuthorizations
func newAzurermExpressRouteCircuitAuthorizations(c *AzurermV1alpha1Client) *azurermExpressRouteCircuitAuthorizations {
	return &azurermExpressRouteCircuitAuthorizations{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermExpressRouteCircuitAuthorization, and returns the corresponding azurermExpressRouteCircuitAuthorization object, and an error if there is any.
func (c *azurermExpressRouteCircuitAuthorizations) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermExpressRouteCircuitAuthorization, err error) {
	result = &v1alpha1.AzurermExpressRouteCircuitAuthorization{}
	err = c.client.Get().
		Resource("azurermexpressroutecircuitauthorizations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermExpressRouteCircuitAuthorizations that match those selectors.
func (c *azurermExpressRouteCircuitAuthorizations) List(opts v1.ListOptions) (result *v1alpha1.AzurermExpressRouteCircuitAuthorizationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermExpressRouteCircuitAuthorizationList{}
	err = c.client.Get().
		Resource("azurermexpressroutecircuitauthorizations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermExpressRouteCircuitAuthorizations.
func (c *azurermExpressRouteCircuitAuthorizations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermexpressroutecircuitauthorizations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermExpressRouteCircuitAuthorization and creates it.  Returns the server's representation of the azurermExpressRouteCircuitAuthorization, and an error, if there is any.
func (c *azurermExpressRouteCircuitAuthorizations) Create(azurermExpressRouteCircuitAuthorization *v1alpha1.AzurermExpressRouteCircuitAuthorization) (result *v1alpha1.AzurermExpressRouteCircuitAuthorization, err error) {
	result = &v1alpha1.AzurermExpressRouteCircuitAuthorization{}
	err = c.client.Post().
		Resource("azurermexpressroutecircuitauthorizations").
		Body(azurermExpressRouteCircuitAuthorization).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermExpressRouteCircuitAuthorization and updates it. Returns the server's representation of the azurermExpressRouteCircuitAuthorization, and an error, if there is any.
func (c *azurermExpressRouteCircuitAuthorizations) Update(azurermExpressRouteCircuitAuthorization *v1alpha1.AzurermExpressRouteCircuitAuthorization) (result *v1alpha1.AzurermExpressRouteCircuitAuthorization, err error) {
	result = &v1alpha1.AzurermExpressRouteCircuitAuthorization{}
	err = c.client.Put().
		Resource("azurermexpressroutecircuitauthorizations").
		Name(azurermExpressRouteCircuitAuthorization.Name).
		Body(azurermExpressRouteCircuitAuthorization).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermExpressRouteCircuitAuthorizations) UpdateStatus(azurermExpressRouteCircuitAuthorization *v1alpha1.AzurermExpressRouteCircuitAuthorization) (result *v1alpha1.AzurermExpressRouteCircuitAuthorization, err error) {
	result = &v1alpha1.AzurermExpressRouteCircuitAuthorization{}
	err = c.client.Put().
		Resource("azurermexpressroutecircuitauthorizations").
		Name(azurermExpressRouteCircuitAuthorization.Name).
		SubResource("status").
		Body(azurermExpressRouteCircuitAuthorization).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermExpressRouteCircuitAuthorization and deletes it. Returns an error if one occurs.
func (c *azurermExpressRouteCircuitAuthorizations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermexpressroutecircuitauthorizations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermExpressRouteCircuitAuthorizations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermexpressroutecircuitauthorizations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermExpressRouteCircuitAuthorization.
func (c *azurermExpressRouteCircuitAuthorizations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermExpressRouteCircuitAuthorization, err error) {
	result = &v1alpha1.AzurermExpressRouteCircuitAuthorization{}
	err = c.client.Patch(pt).
		Resource("azurermexpressroutecircuitauthorizations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
