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

// AzurermFunctionAppsGetter has a method to return a AzurermFunctionAppInterface.
// A group's client should implement this interface.
type AzurermFunctionAppsGetter interface {
	AzurermFunctionApps() AzurermFunctionAppInterface
}

// AzurermFunctionAppInterface has methods to work with AzurermFunctionApp resources.
type AzurermFunctionAppInterface interface {
	Create(*v1alpha1.AzurermFunctionApp) (*v1alpha1.AzurermFunctionApp, error)
	Update(*v1alpha1.AzurermFunctionApp) (*v1alpha1.AzurermFunctionApp, error)
	UpdateStatus(*v1alpha1.AzurermFunctionApp) (*v1alpha1.AzurermFunctionApp, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermFunctionApp, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermFunctionAppList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermFunctionApp, err error)
	AzurermFunctionAppExpansion
}

// azurermFunctionApps implements AzurermFunctionAppInterface
type azurermFunctionApps struct {
	client rest.Interface
}

// newAzurermFunctionApps returns a AzurermFunctionApps
func newAzurermFunctionApps(c *AzurermV1alpha1Client) *azurermFunctionApps {
	return &azurermFunctionApps{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermFunctionApp, and returns the corresponding azurermFunctionApp object, and an error if there is any.
func (c *azurermFunctionApps) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermFunctionApp, err error) {
	result = &v1alpha1.AzurermFunctionApp{}
	err = c.client.Get().
		Resource("azurermfunctionapps").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermFunctionApps that match those selectors.
func (c *azurermFunctionApps) List(opts v1.ListOptions) (result *v1alpha1.AzurermFunctionAppList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermFunctionAppList{}
	err = c.client.Get().
		Resource("azurermfunctionapps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermFunctionApps.
func (c *azurermFunctionApps) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermfunctionapps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermFunctionApp and creates it.  Returns the server's representation of the azurermFunctionApp, and an error, if there is any.
func (c *azurermFunctionApps) Create(azurermFunctionApp *v1alpha1.AzurermFunctionApp) (result *v1alpha1.AzurermFunctionApp, err error) {
	result = &v1alpha1.AzurermFunctionApp{}
	err = c.client.Post().
		Resource("azurermfunctionapps").
		Body(azurermFunctionApp).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermFunctionApp and updates it. Returns the server's representation of the azurermFunctionApp, and an error, if there is any.
func (c *azurermFunctionApps) Update(azurermFunctionApp *v1alpha1.AzurermFunctionApp) (result *v1alpha1.AzurermFunctionApp, err error) {
	result = &v1alpha1.AzurermFunctionApp{}
	err = c.client.Put().
		Resource("azurermfunctionapps").
		Name(azurermFunctionApp.Name).
		Body(azurermFunctionApp).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermFunctionApps) UpdateStatus(azurermFunctionApp *v1alpha1.AzurermFunctionApp) (result *v1alpha1.AzurermFunctionApp, err error) {
	result = &v1alpha1.AzurermFunctionApp{}
	err = c.client.Put().
		Resource("azurermfunctionapps").
		Name(azurermFunctionApp.Name).
		SubResource("status").
		Body(azurermFunctionApp).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermFunctionApp and deletes it. Returns an error if one occurs.
func (c *azurermFunctionApps) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermfunctionapps").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermFunctionApps) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermfunctionapps").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermFunctionApp.
func (c *azurermFunctionApps) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermFunctionApp, err error) {
	result = &v1alpha1.AzurermFunctionApp{}
	err = c.client.Patch(pt).
		Resource("azurermfunctionapps").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
