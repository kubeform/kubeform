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

// ApiManagementProductPoliciesGetter has a method to return a ApiManagementProductPolicyInterface.
// A group's client should implement this interface.
type ApiManagementProductPoliciesGetter interface {
	ApiManagementProductPolicies() ApiManagementProductPolicyInterface
}

// ApiManagementProductPolicyInterface has methods to work with ApiManagementProductPolicy resources.
type ApiManagementProductPolicyInterface interface {
	Create(*v1alpha1.ApiManagementProductPolicy) (*v1alpha1.ApiManagementProductPolicy, error)
	Update(*v1alpha1.ApiManagementProductPolicy) (*v1alpha1.ApiManagementProductPolicy, error)
	UpdateStatus(*v1alpha1.ApiManagementProductPolicy) (*v1alpha1.ApiManagementProductPolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ApiManagementProductPolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.ApiManagementProductPolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiManagementProductPolicy, err error)
	ApiManagementProductPolicyExpansion
}

// apiManagementProductPolicies implements ApiManagementProductPolicyInterface
type apiManagementProductPolicies struct {
	client rest.Interface
}

// newApiManagementProductPolicies returns a ApiManagementProductPolicies
func newApiManagementProductPolicies(c *AzurermV1alpha1Client) *apiManagementProductPolicies {
	return &apiManagementProductPolicies{
		client: c.RESTClient(),
	}
}

// Get takes name of the apiManagementProductPolicy, and returns the corresponding apiManagementProductPolicy object, and an error if there is any.
func (c *apiManagementProductPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.ApiManagementProductPolicy, err error) {
	result = &v1alpha1.ApiManagementProductPolicy{}
	err = c.client.Get().
		Resource("apimanagementproductpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApiManagementProductPolicies that match those selectors.
func (c *apiManagementProductPolicies) List(opts v1.ListOptions) (result *v1alpha1.ApiManagementProductPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ApiManagementProductPolicyList{}
	err = c.client.Get().
		Resource("apimanagementproductpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested apiManagementProductPolicies.
func (c *apiManagementProductPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("apimanagementproductpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a apiManagementProductPolicy and creates it.  Returns the server's representation of the apiManagementProductPolicy, and an error, if there is any.
func (c *apiManagementProductPolicies) Create(apiManagementProductPolicy *v1alpha1.ApiManagementProductPolicy) (result *v1alpha1.ApiManagementProductPolicy, err error) {
	result = &v1alpha1.ApiManagementProductPolicy{}
	err = c.client.Post().
		Resource("apimanagementproductpolicies").
		Body(apiManagementProductPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a apiManagementProductPolicy and updates it. Returns the server's representation of the apiManagementProductPolicy, and an error, if there is any.
func (c *apiManagementProductPolicies) Update(apiManagementProductPolicy *v1alpha1.ApiManagementProductPolicy) (result *v1alpha1.ApiManagementProductPolicy, err error) {
	result = &v1alpha1.ApiManagementProductPolicy{}
	err = c.client.Put().
		Resource("apimanagementproductpolicies").
		Name(apiManagementProductPolicy.Name).
		Body(apiManagementProductPolicy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *apiManagementProductPolicies) UpdateStatus(apiManagementProductPolicy *v1alpha1.ApiManagementProductPolicy) (result *v1alpha1.ApiManagementProductPolicy, err error) {
	result = &v1alpha1.ApiManagementProductPolicy{}
	err = c.client.Put().
		Resource("apimanagementproductpolicies").
		Name(apiManagementProductPolicy.Name).
		SubResource("status").
		Body(apiManagementProductPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the apiManagementProductPolicy and deletes it. Returns an error if one occurs.
func (c *apiManagementProductPolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("apimanagementproductpolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *apiManagementProductPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("apimanagementproductpolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched apiManagementProductPolicy.
func (c *apiManagementProductPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiManagementProductPolicy, err error) {
	result = &v1alpha1.ApiManagementProductPolicy{}
	err = c.client.Patch(pt).
		Resource("apimanagementproductpolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
