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

// AzurermApiManagementApiPoliciesGetter has a method to return a AzurermApiManagementApiPolicyInterface.
// A group's client should implement this interface.
type AzurermApiManagementApiPoliciesGetter interface {
	AzurermApiManagementApiPolicies() AzurermApiManagementApiPolicyInterface
}

// AzurermApiManagementApiPolicyInterface has methods to work with AzurermApiManagementApiPolicy resources.
type AzurermApiManagementApiPolicyInterface interface {
	Create(*v1alpha1.AzurermApiManagementApiPolicy) (*v1alpha1.AzurermApiManagementApiPolicy, error)
	Update(*v1alpha1.AzurermApiManagementApiPolicy) (*v1alpha1.AzurermApiManagementApiPolicy, error)
	UpdateStatus(*v1alpha1.AzurermApiManagementApiPolicy) (*v1alpha1.AzurermApiManagementApiPolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermApiManagementApiPolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermApiManagementApiPolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermApiManagementApiPolicy, err error)
	AzurermApiManagementApiPolicyExpansion
}

// azurermApiManagementApiPolicies implements AzurermApiManagementApiPolicyInterface
type azurermApiManagementApiPolicies struct {
	client rest.Interface
}

// newAzurermApiManagementApiPolicies returns a AzurermApiManagementApiPolicies
func newAzurermApiManagementApiPolicies(c *AzurermV1alpha1Client) *azurermApiManagementApiPolicies {
	return &azurermApiManagementApiPolicies{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermApiManagementApiPolicy, and returns the corresponding azurermApiManagementApiPolicy object, and an error if there is any.
func (c *azurermApiManagementApiPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermApiManagementApiPolicy, err error) {
	result = &v1alpha1.AzurermApiManagementApiPolicy{}
	err = c.client.Get().
		Resource("azurermapimanagementapipolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermApiManagementApiPolicies that match those selectors.
func (c *azurermApiManagementApiPolicies) List(opts v1.ListOptions) (result *v1alpha1.AzurermApiManagementApiPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermApiManagementApiPolicyList{}
	err = c.client.Get().
		Resource("azurermapimanagementapipolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermApiManagementApiPolicies.
func (c *azurermApiManagementApiPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermapimanagementapipolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermApiManagementApiPolicy and creates it.  Returns the server's representation of the azurermApiManagementApiPolicy, and an error, if there is any.
func (c *azurermApiManagementApiPolicies) Create(azurermApiManagementApiPolicy *v1alpha1.AzurermApiManagementApiPolicy) (result *v1alpha1.AzurermApiManagementApiPolicy, err error) {
	result = &v1alpha1.AzurermApiManagementApiPolicy{}
	err = c.client.Post().
		Resource("azurermapimanagementapipolicies").
		Body(azurermApiManagementApiPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermApiManagementApiPolicy and updates it. Returns the server's representation of the azurermApiManagementApiPolicy, and an error, if there is any.
func (c *azurermApiManagementApiPolicies) Update(azurermApiManagementApiPolicy *v1alpha1.AzurermApiManagementApiPolicy) (result *v1alpha1.AzurermApiManagementApiPolicy, err error) {
	result = &v1alpha1.AzurermApiManagementApiPolicy{}
	err = c.client.Put().
		Resource("azurermapimanagementapipolicies").
		Name(azurermApiManagementApiPolicy.Name).
		Body(azurermApiManagementApiPolicy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermApiManagementApiPolicies) UpdateStatus(azurermApiManagementApiPolicy *v1alpha1.AzurermApiManagementApiPolicy) (result *v1alpha1.AzurermApiManagementApiPolicy, err error) {
	result = &v1alpha1.AzurermApiManagementApiPolicy{}
	err = c.client.Put().
		Resource("azurermapimanagementapipolicies").
		Name(azurermApiManagementApiPolicy.Name).
		SubResource("status").
		Body(azurermApiManagementApiPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermApiManagementApiPolicy and deletes it. Returns an error if one occurs.
func (c *azurermApiManagementApiPolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermapimanagementapipolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermApiManagementApiPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermapimanagementapipolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermApiManagementApiPolicy.
func (c *azurermApiManagementApiPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermApiManagementApiPolicy, err error) {
	result = &v1alpha1.AzurermApiManagementApiPolicy{}
	err = c.client.Patch(pt).
		Resource("azurermapimanagementapipolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
