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

// AwsLoadBalancerPoliciesGetter has a method to return a AwsLoadBalancerPolicyInterface.
// A group's client should implement this interface.
type AwsLoadBalancerPoliciesGetter interface {
	AwsLoadBalancerPolicies() AwsLoadBalancerPolicyInterface
}

// AwsLoadBalancerPolicyInterface has methods to work with AwsLoadBalancerPolicy resources.
type AwsLoadBalancerPolicyInterface interface {
	Create(*v1alpha1.AwsLoadBalancerPolicy) (*v1alpha1.AwsLoadBalancerPolicy, error)
	Update(*v1alpha1.AwsLoadBalancerPolicy) (*v1alpha1.AwsLoadBalancerPolicy, error)
	UpdateStatus(*v1alpha1.AwsLoadBalancerPolicy) (*v1alpha1.AwsLoadBalancerPolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsLoadBalancerPolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsLoadBalancerPolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsLoadBalancerPolicy, err error)
	AwsLoadBalancerPolicyExpansion
}

// awsLoadBalancerPolicies implements AwsLoadBalancerPolicyInterface
type awsLoadBalancerPolicies struct {
	client rest.Interface
}

// newAwsLoadBalancerPolicies returns a AwsLoadBalancerPolicies
func newAwsLoadBalancerPolicies(c *AwsV1alpha1Client) *awsLoadBalancerPolicies {
	return &awsLoadBalancerPolicies{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsLoadBalancerPolicy, and returns the corresponding awsLoadBalancerPolicy object, and an error if there is any.
func (c *awsLoadBalancerPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsLoadBalancerPolicy, err error) {
	result = &v1alpha1.AwsLoadBalancerPolicy{}
	err = c.client.Get().
		Resource("awsloadbalancerpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsLoadBalancerPolicies that match those selectors.
func (c *awsLoadBalancerPolicies) List(opts v1.ListOptions) (result *v1alpha1.AwsLoadBalancerPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsLoadBalancerPolicyList{}
	err = c.client.Get().
		Resource("awsloadbalancerpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsLoadBalancerPolicies.
func (c *awsLoadBalancerPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsloadbalancerpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsLoadBalancerPolicy and creates it.  Returns the server's representation of the awsLoadBalancerPolicy, and an error, if there is any.
func (c *awsLoadBalancerPolicies) Create(awsLoadBalancerPolicy *v1alpha1.AwsLoadBalancerPolicy) (result *v1alpha1.AwsLoadBalancerPolicy, err error) {
	result = &v1alpha1.AwsLoadBalancerPolicy{}
	err = c.client.Post().
		Resource("awsloadbalancerpolicies").
		Body(awsLoadBalancerPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsLoadBalancerPolicy and updates it. Returns the server's representation of the awsLoadBalancerPolicy, and an error, if there is any.
func (c *awsLoadBalancerPolicies) Update(awsLoadBalancerPolicy *v1alpha1.AwsLoadBalancerPolicy) (result *v1alpha1.AwsLoadBalancerPolicy, err error) {
	result = &v1alpha1.AwsLoadBalancerPolicy{}
	err = c.client.Put().
		Resource("awsloadbalancerpolicies").
		Name(awsLoadBalancerPolicy.Name).
		Body(awsLoadBalancerPolicy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsLoadBalancerPolicies) UpdateStatus(awsLoadBalancerPolicy *v1alpha1.AwsLoadBalancerPolicy) (result *v1alpha1.AwsLoadBalancerPolicy, err error) {
	result = &v1alpha1.AwsLoadBalancerPolicy{}
	err = c.client.Put().
		Resource("awsloadbalancerpolicies").
		Name(awsLoadBalancerPolicy.Name).
		SubResource("status").
		Body(awsLoadBalancerPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsLoadBalancerPolicy and deletes it. Returns an error if one occurs.
func (c *awsLoadBalancerPolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsloadbalancerpolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsLoadBalancerPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsloadbalancerpolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsLoadBalancerPolicy.
func (c *awsLoadBalancerPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsLoadBalancerPolicy, err error) {
	result = &v1alpha1.AwsLoadBalancerPolicy{}
	err = c.client.Patch(pt).
		Resource("awsloadbalancerpolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
