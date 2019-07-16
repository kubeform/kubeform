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

// SnsTopicPoliciesGetter has a method to return a SnsTopicPolicyInterface.
// A group's client should implement this interface.
type SnsTopicPoliciesGetter interface {
	SnsTopicPolicies() SnsTopicPolicyInterface
}

// SnsTopicPolicyInterface has methods to work with SnsTopicPolicy resources.
type SnsTopicPolicyInterface interface {
	Create(*v1alpha1.SnsTopicPolicy) (*v1alpha1.SnsTopicPolicy, error)
	Update(*v1alpha1.SnsTopicPolicy) (*v1alpha1.SnsTopicPolicy, error)
	UpdateStatus(*v1alpha1.SnsTopicPolicy) (*v1alpha1.SnsTopicPolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SnsTopicPolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.SnsTopicPolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SnsTopicPolicy, err error)
	SnsTopicPolicyExpansion
}

// snsTopicPolicies implements SnsTopicPolicyInterface
type snsTopicPolicies struct {
	client rest.Interface
}

// newSnsTopicPolicies returns a SnsTopicPolicies
func newSnsTopicPolicies(c *AwsV1alpha1Client) *snsTopicPolicies {
	return &snsTopicPolicies{
		client: c.RESTClient(),
	}
}

// Get takes name of the snsTopicPolicy, and returns the corresponding snsTopicPolicy object, and an error if there is any.
func (c *snsTopicPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.SnsTopicPolicy, err error) {
	result = &v1alpha1.SnsTopicPolicy{}
	err = c.client.Get().
		Resource("snstopicpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SnsTopicPolicies that match those selectors.
func (c *snsTopicPolicies) List(opts v1.ListOptions) (result *v1alpha1.SnsTopicPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SnsTopicPolicyList{}
	err = c.client.Get().
		Resource("snstopicpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested snsTopicPolicies.
func (c *snsTopicPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("snstopicpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a snsTopicPolicy and creates it.  Returns the server's representation of the snsTopicPolicy, and an error, if there is any.
func (c *snsTopicPolicies) Create(snsTopicPolicy *v1alpha1.SnsTopicPolicy) (result *v1alpha1.SnsTopicPolicy, err error) {
	result = &v1alpha1.SnsTopicPolicy{}
	err = c.client.Post().
		Resource("snstopicpolicies").
		Body(snsTopicPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a snsTopicPolicy and updates it. Returns the server's representation of the snsTopicPolicy, and an error, if there is any.
func (c *snsTopicPolicies) Update(snsTopicPolicy *v1alpha1.SnsTopicPolicy) (result *v1alpha1.SnsTopicPolicy, err error) {
	result = &v1alpha1.SnsTopicPolicy{}
	err = c.client.Put().
		Resource("snstopicpolicies").
		Name(snsTopicPolicy.Name).
		Body(snsTopicPolicy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *snsTopicPolicies) UpdateStatus(snsTopicPolicy *v1alpha1.SnsTopicPolicy) (result *v1alpha1.SnsTopicPolicy, err error) {
	result = &v1alpha1.SnsTopicPolicy{}
	err = c.client.Put().
		Resource("snstopicpolicies").
		Name(snsTopicPolicy.Name).
		SubResource("status").
		Body(snsTopicPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the snsTopicPolicy and deletes it. Returns an error if one occurs.
func (c *snsTopicPolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("snstopicpolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *snsTopicPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("snstopicpolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched snsTopicPolicy.
func (c *snsTopicPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SnsTopicPolicy, err error) {
	result = &v1alpha1.SnsTopicPolicy{}
	err = c.client.Patch(pt).
		Resource("snstopicpolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
