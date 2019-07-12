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

// AzurermApiManagementOpenidConnectProvidersGetter has a method to return a AzurermApiManagementOpenidConnectProviderInterface.
// A group's client should implement this interface.
type AzurermApiManagementOpenidConnectProvidersGetter interface {
	AzurermApiManagementOpenidConnectProviders() AzurermApiManagementOpenidConnectProviderInterface
}

// AzurermApiManagementOpenidConnectProviderInterface has methods to work with AzurermApiManagementOpenidConnectProvider resources.
type AzurermApiManagementOpenidConnectProviderInterface interface {
	Create(*v1alpha1.AzurermApiManagementOpenidConnectProvider) (*v1alpha1.AzurermApiManagementOpenidConnectProvider, error)
	Update(*v1alpha1.AzurermApiManagementOpenidConnectProvider) (*v1alpha1.AzurermApiManagementOpenidConnectProvider, error)
	UpdateStatus(*v1alpha1.AzurermApiManagementOpenidConnectProvider) (*v1alpha1.AzurermApiManagementOpenidConnectProvider, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermApiManagementOpenidConnectProvider, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermApiManagementOpenidConnectProviderList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermApiManagementOpenidConnectProvider, err error)
	AzurermApiManagementOpenidConnectProviderExpansion
}

// azurermApiManagementOpenidConnectProviders implements AzurermApiManagementOpenidConnectProviderInterface
type azurermApiManagementOpenidConnectProviders struct {
	client rest.Interface
}

// newAzurermApiManagementOpenidConnectProviders returns a AzurermApiManagementOpenidConnectProviders
func newAzurermApiManagementOpenidConnectProviders(c *AzurermV1alpha1Client) *azurermApiManagementOpenidConnectProviders {
	return &azurermApiManagementOpenidConnectProviders{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermApiManagementOpenidConnectProvider, and returns the corresponding azurermApiManagementOpenidConnectProvider object, and an error if there is any.
func (c *azurermApiManagementOpenidConnectProviders) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermApiManagementOpenidConnectProvider, err error) {
	result = &v1alpha1.AzurermApiManagementOpenidConnectProvider{}
	err = c.client.Get().
		Resource("azurermapimanagementopenidconnectproviders").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermApiManagementOpenidConnectProviders that match those selectors.
func (c *azurermApiManagementOpenidConnectProviders) List(opts v1.ListOptions) (result *v1alpha1.AzurermApiManagementOpenidConnectProviderList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermApiManagementOpenidConnectProviderList{}
	err = c.client.Get().
		Resource("azurermapimanagementopenidconnectproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermApiManagementOpenidConnectProviders.
func (c *azurermApiManagementOpenidConnectProviders) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermapimanagementopenidconnectproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermApiManagementOpenidConnectProvider and creates it.  Returns the server's representation of the azurermApiManagementOpenidConnectProvider, and an error, if there is any.
func (c *azurermApiManagementOpenidConnectProviders) Create(azurermApiManagementOpenidConnectProvider *v1alpha1.AzurermApiManagementOpenidConnectProvider) (result *v1alpha1.AzurermApiManagementOpenidConnectProvider, err error) {
	result = &v1alpha1.AzurermApiManagementOpenidConnectProvider{}
	err = c.client.Post().
		Resource("azurermapimanagementopenidconnectproviders").
		Body(azurermApiManagementOpenidConnectProvider).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermApiManagementOpenidConnectProvider and updates it. Returns the server's representation of the azurermApiManagementOpenidConnectProvider, and an error, if there is any.
func (c *azurermApiManagementOpenidConnectProviders) Update(azurermApiManagementOpenidConnectProvider *v1alpha1.AzurermApiManagementOpenidConnectProvider) (result *v1alpha1.AzurermApiManagementOpenidConnectProvider, err error) {
	result = &v1alpha1.AzurermApiManagementOpenidConnectProvider{}
	err = c.client.Put().
		Resource("azurermapimanagementopenidconnectproviders").
		Name(azurermApiManagementOpenidConnectProvider.Name).
		Body(azurermApiManagementOpenidConnectProvider).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermApiManagementOpenidConnectProviders) UpdateStatus(azurermApiManagementOpenidConnectProvider *v1alpha1.AzurermApiManagementOpenidConnectProvider) (result *v1alpha1.AzurermApiManagementOpenidConnectProvider, err error) {
	result = &v1alpha1.AzurermApiManagementOpenidConnectProvider{}
	err = c.client.Put().
		Resource("azurermapimanagementopenidconnectproviders").
		Name(azurermApiManagementOpenidConnectProvider.Name).
		SubResource("status").
		Body(azurermApiManagementOpenidConnectProvider).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermApiManagementOpenidConnectProvider and deletes it. Returns an error if one occurs.
func (c *azurermApiManagementOpenidConnectProviders) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermapimanagementopenidconnectproviders").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermApiManagementOpenidConnectProviders) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermapimanagementopenidconnectproviders").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermApiManagementOpenidConnectProvider.
func (c *azurermApiManagementOpenidConnectProviders) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermApiManagementOpenidConnectProvider, err error) {
	result = &v1alpha1.AzurermApiManagementOpenidConnectProvider{}
	err = c.client.Patch(pt).
		Resource("azurermapimanagementopenidconnectproviders").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
