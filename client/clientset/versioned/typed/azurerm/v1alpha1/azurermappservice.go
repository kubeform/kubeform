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

// AzurermAppServicesGetter has a method to return a AzurermAppServiceInterface.
// A group's client should implement this interface.
type AzurermAppServicesGetter interface {
	AzurermAppServices() AzurermAppServiceInterface
}

// AzurermAppServiceInterface has methods to work with AzurermAppService resources.
type AzurermAppServiceInterface interface {
	Create(*v1alpha1.AzurermAppService) (*v1alpha1.AzurermAppService, error)
	Update(*v1alpha1.AzurermAppService) (*v1alpha1.AzurermAppService, error)
	UpdateStatus(*v1alpha1.AzurermAppService) (*v1alpha1.AzurermAppService, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermAppService, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermAppServiceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermAppService, err error)
	AzurermAppServiceExpansion
}

// azurermAppServices implements AzurermAppServiceInterface
type azurermAppServices struct {
	client rest.Interface
}

// newAzurermAppServices returns a AzurermAppServices
func newAzurermAppServices(c *AzurermV1alpha1Client) *azurermAppServices {
	return &azurermAppServices{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermAppService, and returns the corresponding azurermAppService object, and an error if there is any.
func (c *azurermAppServices) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermAppService, err error) {
	result = &v1alpha1.AzurermAppService{}
	err = c.client.Get().
		Resource("azurermappservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermAppServices that match those selectors.
func (c *azurermAppServices) List(opts v1.ListOptions) (result *v1alpha1.AzurermAppServiceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermAppServiceList{}
	err = c.client.Get().
		Resource("azurermappservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermAppServices.
func (c *azurermAppServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermappservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermAppService and creates it.  Returns the server's representation of the azurermAppService, and an error, if there is any.
func (c *azurermAppServices) Create(azurermAppService *v1alpha1.AzurermAppService) (result *v1alpha1.AzurermAppService, err error) {
	result = &v1alpha1.AzurermAppService{}
	err = c.client.Post().
		Resource("azurermappservices").
		Body(azurermAppService).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermAppService and updates it. Returns the server's representation of the azurermAppService, and an error, if there is any.
func (c *azurermAppServices) Update(azurermAppService *v1alpha1.AzurermAppService) (result *v1alpha1.AzurermAppService, err error) {
	result = &v1alpha1.AzurermAppService{}
	err = c.client.Put().
		Resource("azurermappservices").
		Name(azurermAppService.Name).
		Body(azurermAppService).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermAppServices) UpdateStatus(azurermAppService *v1alpha1.AzurermAppService) (result *v1alpha1.AzurermAppService, err error) {
	result = &v1alpha1.AzurermAppService{}
	err = c.client.Put().
		Resource("azurermappservices").
		Name(azurermAppService.Name).
		SubResource("status").
		Body(azurermAppService).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermAppService and deletes it. Returns an error if one occurs.
func (c *azurermAppServices) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermappservices").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermAppServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermappservices").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermAppService.
func (c *azurermAppServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermAppService, err error) {
	result = &v1alpha1.AzurermAppService{}
	err = c.client.Patch(pt).
		Resource("azurermappservices").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
