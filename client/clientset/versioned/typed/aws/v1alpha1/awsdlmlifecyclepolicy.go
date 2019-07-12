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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// AwsDlmLifecyclePoliciesGetter has a method to return a AwsDlmLifecyclePolicyInterface.
// A group's client should implement this interface.
type AwsDlmLifecyclePoliciesGetter interface {
	AwsDlmLifecyclePolicies() AwsDlmLifecyclePolicyInterface
}

// AwsDlmLifecyclePolicyInterface has methods to work with AwsDlmLifecyclePolicy resources.
type AwsDlmLifecyclePolicyInterface interface {
	Create(*v1alpha1.AwsDlmLifecyclePolicy) (*v1alpha1.AwsDlmLifecyclePolicy, error)
	Update(*v1alpha1.AwsDlmLifecyclePolicy) (*v1alpha1.AwsDlmLifecyclePolicy, error)
	UpdateStatus(*v1alpha1.AwsDlmLifecyclePolicy) (*v1alpha1.AwsDlmLifecyclePolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsDlmLifecyclePolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsDlmLifecyclePolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDlmLifecyclePolicy, err error)
	AwsDlmLifecyclePolicyExpansion
}

// awsDlmLifecyclePolicies implements AwsDlmLifecyclePolicyInterface
type awsDlmLifecyclePolicies struct {
	client rest.Interface
}

// newAwsDlmLifecyclePolicies returns a AwsDlmLifecyclePolicies
func newAwsDlmLifecyclePolicies(c *AwsV1alpha1Client) *awsDlmLifecyclePolicies {
	return &awsDlmLifecyclePolicies{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsDlmLifecyclePolicy, and returns the corresponding awsDlmLifecyclePolicy object, and an error if there is any.
func (c *awsDlmLifecyclePolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDlmLifecyclePolicy, err error) {
	result = &v1alpha1.AwsDlmLifecyclePolicy{}
	err = c.client.Get().
		Resource("awsdlmlifecyclepolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsDlmLifecyclePolicies that match those selectors.
func (c *awsDlmLifecyclePolicies) List(opts v1.ListOptions) (result *v1alpha1.AwsDlmLifecyclePolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsDlmLifecyclePolicyList{}
	err = c.client.Get().
		Resource("awsdlmlifecyclepolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsDlmLifecyclePolicies.
func (c *awsDlmLifecyclePolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsdlmlifecyclepolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsDlmLifecyclePolicy and creates it.  Returns the server's representation of the awsDlmLifecyclePolicy, and an error, if there is any.
func (c *awsDlmLifecyclePolicies) Create(awsDlmLifecyclePolicy *v1alpha1.AwsDlmLifecyclePolicy) (result *v1alpha1.AwsDlmLifecyclePolicy, err error) {
	result = &v1alpha1.AwsDlmLifecyclePolicy{}
	err = c.client.Post().
		Resource("awsdlmlifecyclepolicies").
		Body(awsDlmLifecyclePolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsDlmLifecyclePolicy and updates it. Returns the server's representation of the awsDlmLifecyclePolicy, and an error, if there is any.
func (c *awsDlmLifecyclePolicies) Update(awsDlmLifecyclePolicy *v1alpha1.AwsDlmLifecyclePolicy) (result *v1alpha1.AwsDlmLifecyclePolicy, err error) {
	result = &v1alpha1.AwsDlmLifecyclePolicy{}
	err = c.client.Put().
		Resource("awsdlmlifecyclepolicies").
		Name(awsDlmLifecyclePolicy.Name).
		Body(awsDlmLifecyclePolicy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsDlmLifecyclePolicies) UpdateStatus(awsDlmLifecyclePolicy *v1alpha1.AwsDlmLifecyclePolicy) (result *v1alpha1.AwsDlmLifecyclePolicy, err error) {
	result = &v1alpha1.AwsDlmLifecyclePolicy{}
	err = c.client.Put().
		Resource("awsdlmlifecyclepolicies").
		Name(awsDlmLifecyclePolicy.Name).
		SubResource("status").
		Body(awsDlmLifecyclePolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsDlmLifecyclePolicy and deletes it. Returns an error if one occurs.
func (c *awsDlmLifecyclePolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsdlmlifecyclepolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsDlmLifecyclePolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsdlmlifecyclepolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsDlmLifecyclePolicy.
func (c *awsDlmLifecyclePolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDlmLifecyclePolicy, err error) {
	result = &v1alpha1.AwsDlmLifecyclePolicy{}
	err = c.client.Patch(pt).
		Resource("awsdlmlifecyclepolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
