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

// AwsAutoscalingPoliciesGetter has a method to return a AwsAutoscalingPolicyInterface.
// A group's client should implement this interface.
type AwsAutoscalingPoliciesGetter interface {
	AwsAutoscalingPolicies() AwsAutoscalingPolicyInterface
}

// AwsAutoscalingPolicyInterface has methods to work with AwsAutoscalingPolicy resources.
type AwsAutoscalingPolicyInterface interface {
	Create(*v1alpha1.AwsAutoscalingPolicy) (*v1alpha1.AwsAutoscalingPolicy, error)
	Update(*v1alpha1.AwsAutoscalingPolicy) (*v1alpha1.AwsAutoscalingPolicy, error)
	UpdateStatus(*v1alpha1.AwsAutoscalingPolicy) (*v1alpha1.AwsAutoscalingPolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsAutoscalingPolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsAutoscalingPolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAutoscalingPolicy, err error)
	AwsAutoscalingPolicyExpansion
}

// awsAutoscalingPolicies implements AwsAutoscalingPolicyInterface
type awsAutoscalingPolicies struct {
	client rest.Interface
}

// newAwsAutoscalingPolicies returns a AwsAutoscalingPolicies
func newAwsAutoscalingPolicies(c *AwsV1alpha1Client) *awsAutoscalingPolicies {
	return &awsAutoscalingPolicies{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsAutoscalingPolicy, and returns the corresponding awsAutoscalingPolicy object, and an error if there is any.
func (c *awsAutoscalingPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsAutoscalingPolicy, err error) {
	result = &v1alpha1.AwsAutoscalingPolicy{}
	err = c.client.Get().
		Resource("awsautoscalingpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsAutoscalingPolicies that match those selectors.
func (c *awsAutoscalingPolicies) List(opts v1.ListOptions) (result *v1alpha1.AwsAutoscalingPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsAutoscalingPolicyList{}
	err = c.client.Get().
		Resource("awsautoscalingpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsAutoscalingPolicies.
func (c *awsAutoscalingPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsautoscalingpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsAutoscalingPolicy and creates it.  Returns the server's representation of the awsAutoscalingPolicy, and an error, if there is any.
func (c *awsAutoscalingPolicies) Create(awsAutoscalingPolicy *v1alpha1.AwsAutoscalingPolicy) (result *v1alpha1.AwsAutoscalingPolicy, err error) {
	result = &v1alpha1.AwsAutoscalingPolicy{}
	err = c.client.Post().
		Resource("awsautoscalingpolicies").
		Body(awsAutoscalingPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsAutoscalingPolicy and updates it. Returns the server's representation of the awsAutoscalingPolicy, and an error, if there is any.
func (c *awsAutoscalingPolicies) Update(awsAutoscalingPolicy *v1alpha1.AwsAutoscalingPolicy) (result *v1alpha1.AwsAutoscalingPolicy, err error) {
	result = &v1alpha1.AwsAutoscalingPolicy{}
	err = c.client.Put().
		Resource("awsautoscalingpolicies").
		Name(awsAutoscalingPolicy.Name).
		Body(awsAutoscalingPolicy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsAutoscalingPolicies) UpdateStatus(awsAutoscalingPolicy *v1alpha1.AwsAutoscalingPolicy) (result *v1alpha1.AwsAutoscalingPolicy, err error) {
	result = &v1alpha1.AwsAutoscalingPolicy{}
	err = c.client.Put().
		Resource("awsautoscalingpolicies").
		Name(awsAutoscalingPolicy.Name).
		SubResource("status").
		Body(awsAutoscalingPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsAutoscalingPolicy and deletes it. Returns an error if one occurs.
func (c *awsAutoscalingPolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsautoscalingpolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsAutoscalingPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsautoscalingpolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsAutoscalingPolicy.
func (c *awsAutoscalingPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAutoscalingPolicy, err error) {
	result = &v1alpha1.AwsAutoscalingPolicy{}
	err = c.client.Patch(pt).
		Resource("awsautoscalingpolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
