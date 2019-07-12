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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// GoogleComputeSslPoliciesGetter has a method to return a GoogleComputeSslPolicyInterface.
// A group's client should implement this interface.
type GoogleComputeSslPoliciesGetter interface {
	GoogleComputeSslPolicies() GoogleComputeSslPolicyInterface
}

// GoogleComputeSslPolicyInterface has methods to work with GoogleComputeSslPolicy resources.
type GoogleComputeSslPolicyInterface interface {
	Create(*v1alpha1.GoogleComputeSslPolicy) (*v1alpha1.GoogleComputeSslPolicy, error)
	Update(*v1alpha1.GoogleComputeSslPolicy) (*v1alpha1.GoogleComputeSslPolicy, error)
	UpdateStatus(*v1alpha1.GoogleComputeSslPolicy) (*v1alpha1.GoogleComputeSslPolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GoogleComputeSslPolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.GoogleComputeSslPolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleComputeSslPolicy, err error)
	GoogleComputeSslPolicyExpansion
}

// googleComputeSslPolicies implements GoogleComputeSslPolicyInterface
type googleComputeSslPolicies struct {
	client rest.Interface
}

// newGoogleComputeSslPolicies returns a GoogleComputeSslPolicies
func newGoogleComputeSslPolicies(c *GoogleV1alpha1Client) *googleComputeSslPolicies {
	return &googleComputeSslPolicies{
		client: c.RESTClient(),
	}
}

// Get takes name of the googleComputeSslPolicy, and returns the corresponding googleComputeSslPolicy object, and an error if there is any.
func (c *googleComputeSslPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleComputeSslPolicy, err error) {
	result = &v1alpha1.GoogleComputeSslPolicy{}
	err = c.client.Get().
		Resource("googlecomputesslpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GoogleComputeSslPolicies that match those selectors.
func (c *googleComputeSslPolicies) List(opts v1.ListOptions) (result *v1alpha1.GoogleComputeSslPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GoogleComputeSslPolicyList{}
	err = c.client.Get().
		Resource("googlecomputesslpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested googleComputeSslPolicies.
func (c *googleComputeSslPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("googlecomputesslpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a googleComputeSslPolicy and creates it.  Returns the server's representation of the googleComputeSslPolicy, and an error, if there is any.
func (c *googleComputeSslPolicies) Create(googleComputeSslPolicy *v1alpha1.GoogleComputeSslPolicy) (result *v1alpha1.GoogleComputeSslPolicy, err error) {
	result = &v1alpha1.GoogleComputeSslPolicy{}
	err = c.client.Post().
		Resource("googlecomputesslpolicies").
		Body(googleComputeSslPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a googleComputeSslPolicy and updates it. Returns the server's representation of the googleComputeSslPolicy, and an error, if there is any.
func (c *googleComputeSslPolicies) Update(googleComputeSslPolicy *v1alpha1.GoogleComputeSslPolicy) (result *v1alpha1.GoogleComputeSslPolicy, err error) {
	result = &v1alpha1.GoogleComputeSslPolicy{}
	err = c.client.Put().
		Resource("googlecomputesslpolicies").
		Name(googleComputeSslPolicy.Name).
		Body(googleComputeSslPolicy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *googleComputeSslPolicies) UpdateStatus(googleComputeSslPolicy *v1alpha1.GoogleComputeSslPolicy) (result *v1alpha1.GoogleComputeSslPolicy, err error) {
	result = &v1alpha1.GoogleComputeSslPolicy{}
	err = c.client.Put().
		Resource("googlecomputesslpolicies").
		Name(googleComputeSslPolicy.Name).
		SubResource("status").
		Body(googleComputeSslPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the googleComputeSslPolicy and deletes it. Returns an error if one occurs.
func (c *googleComputeSslPolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("googlecomputesslpolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *googleComputeSslPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("googlecomputesslpolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched googleComputeSslPolicy.
func (c *googleComputeSslPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleComputeSslPolicy, err error) {
	result = &v1alpha1.GoogleComputeSslPolicy{}
	err = c.client.Patch(pt).
		Resource("googlecomputesslpolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
