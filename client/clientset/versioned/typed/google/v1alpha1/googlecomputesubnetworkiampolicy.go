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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// GoogleComputeSubnetworkIamPoliciesGetter has a method to return a GoogleComputeSubnetworkIamPolicyInterface.
// A group's client should implement this interface.
type GoogleComputeSubnetworkIamPoliciesGetter interface {
	GoogleComputeSubnetworkIamPolicies() GoogleComputeSubnetworkIamPolicyInterface
}

// GoogleComputeSubnetworkIamPolicyInterface has methods to work with GoogleComputeSubnetworkIamPolicy resources.
type GoogleComputeSubnetworkIamPolicyInterface interface {
	Create(*v1alpha1.GoogleComputeSubnetworkIamPolicy) (*v1alpha1.GoogleComputeSubnetworkIamPolicy, error)
	Update(*v1alpha1.GoogleComputeSubnetworkIamPolicy) (*v1alpha1.GoogleComputeSubnetworkIamPolicy, error)
	UpdateStatus(*v1alpha1.GoogleComputeSubnetworkIamPolicy) (*v1alpha1.GoogleComputeSubnetworkIamPolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GoogleComputeSubnetworkIamPolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.GoogleComputeSubnetworkIamPolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleComputeSubnetworkIamPolicy, err error)
	GoogleComputeSubnetworkIamPolicyExpansion
}

// googleComputeSubnetworkIamPolicies implements GoogleComputeSubnetworkIamPolicyInterface
type googleComputeSubnetworkIamPolicies struct {
	client rest.Interface
}

// newGoogleComputeSubnetworkIamPolicies returns a GoogleComputeSubnetworkIamPolicies
func newGoogleComputeSubnetworkIamPolicies(c *GoogleV1alpha1Client) *googleComputeSubnetworkIamPolicies {
	return &googleComputeSubnetworkIamPolicies{
		client: c.RESTClient(),
	}
}

// Get takes name of the googleComputeSubnetworkIamPolicy, and returns the corresponding googleComputeSubnetworkIamPolicy object, and an error if there is any.
func (c *googleComputeSubnetworkIamPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleComputeSubnetworkIamPolicy, err error) {
	result = &v1alpha1.GoogleComputeSubnetworkIamPolicy{}
	err = c.client.Get().
		Resource("googlecomputesubnetworkiampolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GoogleComputeSubnetworkIamPolicies that match those selectors.
func (c *googleComputeSubnetworkIamPolicies) List(opts v1.ListOptions) (result *v1alpha1.GoogleComputeSubnetworkIamPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GoogleComputeSubnetworkIamPolicyList{}
	err = c.client.Get().
		Resource("googlecomputesubnetworkiampolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested googleComputeSubnetworkIamPolicies.
func (c *googleComputeSubnetworkIamPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("googlecomputesubnetworkiampolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a googleComputeSubnetworkIamPolicy and creates it.  Returns the server's representation of the googleComputeSubnetworkIamPolicy, and an error, if there is any.
func (c *googleComputeSubnetworkIamPolicies) Create(googleComputeSubnetworkIamPolicy *v1alpha1.GoogleComputeSubnetworkIamPolicy) (result *v1alpha1.GoogleComputeSubnetworkIamPolicy, err error) {
	result = &v1alpha1.GoogleComputeSubnetworkIamPolicy{}
	err = c.client.Post().
		Resource("googlecomputesubnetworkiampolicies").
		Body(googleComputeSubnetworkIamPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a googleComputeSubnetworkIamPolicy and updates it. Returns the server's representation of the googleComputeSubnetworkIamPolicy, and an error, if there is any.
func (c *googleComputeSubnetworkIamPolicies) Update(googleComputeSubnetworkIamPolicy *v1alpha1.GoogleComputeSubnetworkIamPolicy) (result *v1alpha1.GoogleComputeSubnetworkIamPolicy, err error) {
	result = &v1alpha1.GoogleComputeSubnetworkIamPolicy{}
	err = c.client.Put().
		Resource("googlecomputesubnetworkiampolicies").
		Name(googleComputeSubnetworkIamPolicy.Name).
		Body(googleComputeSubnetworkIamPolicy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *googleComputeSubnetworkIamPolicies) UpdateStatus(googleComputeSubnetworkIamPolicy *v1alpha1.GoogleComputeSubnetworkIamPolicy) (result *v1alpha1.GoogleComputeSubnetworkIamPolicy, err error) {
	result = &v1alpha1.GoogleComputeSubnetworkIamPolicy{}
	err = c.client.Put().
		Resource("googlecomputesubnetworkiampolicies").
		Name(googleComputeSubnetworkIamPolicy.Name).
		SubResource("status").
		Body(googleComputeSubnetworkIamPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the googleComputeSubnetworkIamPolicy and deletes it. Returns an error if one occurs.
func (c *googleComputeSubnetworkIamPolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("googlecomputesubnetworkiampolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *googleComputeSubnetworkIamPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("googlecomputesubnetworkiampolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched googleComputeSubnetworkIamPolicy.
func (c *googleComputeSubnetworkIamPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleComputeSubnetworkIamPolicy, err error) {
	result = &v1alpha1.GoogleComputeSubnetworkIamPolicy{}
	err = c.client.Patch(pt).
		Resource("googlecomputesubnetworkiampolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
