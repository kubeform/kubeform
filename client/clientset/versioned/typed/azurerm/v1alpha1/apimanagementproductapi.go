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

	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ApiManagementProductAPIsGetter has a method to return a ApiManagementProductAPIInterface.
// A group's client should implement this interface.
type ApiManagementProductAPIsGetter interface {
	ApiManagementProductAPIs(namespace string) ApiManagementProductAPIInterface
}

// ApiManagementProductAPIInterface has methods to work with ApiManagementProductAPI resources.
type ApiManagementProductAPIInterface interface {
	Create(*v1alpha1.ApiManagementProductAPI) (*v1alpha1.ApiManagementProductAPI, error)
	Update(*v1alpha1.ApiManagementProductAPI) (*v1alpha1.ApiManagementProductAPI, error)
	UpdateStatus(*v1alpha1.ApiManagementProductAPI) (*v1alpha1.ApiManagementProductAPI, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ApiManagementProductAPI, error)
	List(opts v1.ListOptions) (*v1alpha1.ApiManagementProductAPIList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiManagementProductAPI, err error)
	ApiManagementProductAPIExpansion
}

// apiManagementProductAPIs implements ApiManagementProductAPIInterface
type apiManagementProductAPIs struct {
	client rest.Interface
	ns     string
}

// newApiManagementProductAPIs returns a ApiManagementProductAPIs
func newApiManagementProductAPIs(c *AzurermV1alpha1Client, namespace string) *apiManagementProductAPIs {
	return &apiManagementProductAPIs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the apiManagementProductAPI, and returns the corresponding apiManagementProductAPI object, and an error if there is any.
func (c *apiManagementProductAPIs) Get(name string, options v1.GetOptions) (result *v1alpha1.ApiManagementProductAPI, err error) {
	result = &v1alpha1.ApiManagementProductAPI{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apimanagementproductapis").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApiManagementProductAPIs that match those selectors.
func (c *apiManagementProductAPIs) List(opts v1.ListOptions) (result *v1alpha1.ApiManagementProductAPIList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ApiManagementProductAPIList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apimanagementproductapis").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested apiManagementProductAPIs.
func (c *apiManagementProductAPIs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("apimanagementproductapis").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a apiManagementProductAPI and creates it.  Returns the server's representation of the apiManagementProductAPI, and an error, if there is any.
func (c *apiManagementProductAPIs) Create(apiManagementProductAPI *v1alpha1.ApiManagementProductAPI) (result *v1alpha1.ApiManagementProductAPI, err error) {
	result = &v1alpha1.ApiManagementProductAPI{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("apimanagementproductapis").
		Body(apiManagementProductAPI).
		Do().
		Into(result)
	return
}

// Update takes the representation of a apiManagementProductAPI and updates it. Returns the server's representation of the apiManagementProductAPI, and an error, if there is any.
func (c *apiManagementProductAPIs) Update(apiManagementProductAPI *v1alpha1.ApiManagementProductAPI) (result *v1alpha1.ApiManagementProductAPI, err error) {
	result = &v1alpha1.ApiManagementProductAPI{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apimanagementproductapis").
		Name(apiManagementProductAPI.Name).
		Body(apiManagementProductAPI).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *apiManagementProductAPIs) UpdateStatus(apiManagementProductAPI *v1alpha1.ApiManagementProductAPI) (result *v1alpha1.ApiManagementProductAPI, err error) {
	result = &v1alpha1.ApiManagementProductAPI{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apimanagementproductapis").
		Name(apiManagementProductAPI.Name).
		SubResource("status").
		Body(apiManagementProductAPI).
		Do().
		Into(result)
	return
}

// Delete takes name of the apiManagementProductAPI and deletes it. Returns an error if one occurs.
func (c *apiManagementProductAPIs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apimanagementproductapis").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *apiManagementProductAPIs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apimanagementproductapis").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched apiManagementProductAPI.
func (c *apiManagementProductAPIs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiManagementProductAPI, err error) {
	result = &v1alpha1.ApiManagementProductAPI{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("apimanagementproductapis").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
