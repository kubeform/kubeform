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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// AwsLbCookieStickinessPoliciesGetter has a method to return a AwsLbCookieStickinessPolicyInterface.
// A group's client should implement this interface.
type AwsLbCookieStickinessPoliciesGetter interface {
	AwsLbCookieStickinessPolicies() AwsLbCookieStickinessPolicyInterface
}

// AwsLbCookieStickinessPolicyInterface has methods to work with AwsLbCookieStickinessPolicy resources.
type AwsLbCookieStickinessPolicyInterface interface {
	Create(*v1alpha1.AwsLbCookieStickinessPolicy) (*v1alpha1.AwsLbCookieStickinessPolicy, error)
	Update(*v1alpha1.AwsLbCookieStickinessPolicy) (*v1alpha1.AwsLbCookieStickinessPolicy, error)
	UpdateStatus(*v1alpha1.AwsLbCookieStickinessPolicy) (*v1alpha1.AwsLbCookieStickinessPolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsLbCookieStickinessPolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsLbCookieStickinessPolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsLbCookieStickinessPolicy, err error)
	AwsLbCookieStickinessPolicyExpansion
}

// awsLbCookieStickinessPolicies implements AwsLbCookieStickinessPolicyInterface
type awsLbCookieStickinessPolicies struct {
	client rest.Interface
}

// newAwsLbCookieStickinessPolicies returns a AwsLbCookieStickinessPolicies
func newAwsLbCookieStickinessPolicies(c *AwsV1alpha1Client) *awsLbCookieStickinessPolicies {
	return &awsLbCookieStickinessPolicies{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsLbCookieStickinessPolicy, and returns the corresponding awsLbCookieStickinessPolicy object, and an error if there is any.
func (c *awsLbCookieStickinessPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsLbCookieStickinessPolicy, err error) {
	result = &v1alpha1.AwsLbCookieStickinessPolicy{}
	err = c.client.Get().
		Resource("awslbcookiestickinesspolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsLbCookieStickinessPolicies that match those selectors.
func (c *awsLbCookieStickinessPolicies) List(opts v1.ListOptions) (result *v1alpha1.AwsLbCookieStickinessPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsLbCookieStickinessPolicyList{}
	err = c.client.Get().
		Resource("awslbcookiestickinesspolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsLbCookieStickinessPolicies.
func (c *awsLbCookieStickinessPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awslbcookiestickinesspolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsLbCookieStickinessPolicy and creates it.  Returns the server's representation of the awsLbCookieStickinessPolicy, and an error, if there is any.
func (c *awsLbCookieStickinessPolicies) Create(awsLbCookieStickinessPolicy *v1alpha1.AwsLbCookieStickinessPolicy) (result *v1alpha1.AwsLbCookieStickinessPolicy, err error) {
	result = &v1alpha1.AwsLbCookieStickinessPolicy{}
	err = c.client.Post().
		Resource("awslbcookiestickinesspolicies").
		Body(awsLbCookieStickinessPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsLbCookieStickinessPolicy and updates it. Returns the server's representation of the awsLbCookieStickinessPolicy, and an error, if there is any.
func (c *awsLbCookieStickinessPolicies) Update(awsLbCookieStickinessPolicy *v1alpha1.AwsLbCookieStickinessPolicy) (result *v1alpha1.AwsLbCookieStickinessPolicy, err error) {
	result = &v1alpha1.AwsLbCookieStickinessPolicy{}
	err = c.client.Put().
		Resource("awslbcookiestickinesspolicies").
		Name(awsLbCookieStickinessPolicy.Name).
		Body(awsLbCookieStickinessPolicy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsLbCookieStickinessPolicies) UpdateStatus(awsLbCookieStickinessPolicy *v1alpha1.AwsLbCookieStickinessPolicy) (result *v1alpha1.AwsLbCookieStickinessPolicy, err error) {
	result = &v1alpha1.AwsLbCookieStickinessPolicy{}
	err = c.client.Put().
		Resource("awslbcookiestickinesspolicies").
		Name(awsLbCookieStickinessPolicy.Name).
		SubResource("status").
		Body(awsLbCookieStickinessPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsLbCookieStickinessPolicy and deletes it. Returns an error if one occurs.
func (c *awsLbCookieStickinessPolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awslbcookiestickinesspolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsLbCookieStickinessPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awslbcookiestickinesspolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsLbCookieStickinessPolicy.
func (c *awsLbCookieStickinessPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsLbCookieStickinessPolicy, err error) {
	result = &v1alpha1.AwsLbCookieStickinessPolicy{}
	err = c.client.Patch(pt).
		Resource("awslbcookiestickinesspolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
