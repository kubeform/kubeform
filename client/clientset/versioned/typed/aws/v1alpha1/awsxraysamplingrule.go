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

// AwsXraySamplingRulesGetter has a method to return a AwsXraySamplingRuleInterface.
// A group's client should implement this interface.
type AwsXraySamplingRulesGetter interface {
	AwsXraySamplingRules() AwsXraySamplingRuleInterface
}

// AwsXraySamplingRuleInterface has methods to work with AwsXraySamplingRule resources.
type AwsXraySamplingRuleInterface interface {
	Create(*v1alpha1.AwsXraySamplingRule) (*v1alpha1.AwsXraySamplingRule, error)
	Update(*v1alpha1.AwsXraySamplingRule) (*v1alpha1.AwsXraySamplingRule, error)
	UpdateStatus(*v1alpha1.AwsXraySamplingRule) (*v1alpha1.AwsXraySamplingRule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsXraySamplingRule, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsXraySamplingRuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsXraySamplingRule, err error)
	AwsXraySamplingRuleExpansion
}

// awsXraySamplingRules implements AwsXraySamplingRuleInterface
type awsXraySamplingRules struct {
	client rest.Interface
}

// newAwsXraySamplingRules returns a AwsXraySamplingRules
func newAwsXraySamplingRules(c *AwsV1alpha1Client) *awsXraySamplingRules {
	return &awsXraySamplingRules{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsXraySamplingRule, and returns the corresponding awsXraySamplingRule object, and an error if there is any.
func (c *awsXraySamplingRules) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsXraySamplingRule, err error) {
	result = &v1alpha1.AwsXraySamplingRule{}
	err = c.client.Get().
		Resource("awsxraysamplingrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsXraySamplingRules that match those selectors.
func (c *awsXraySamplingRules) List(opts v1.ListOptions) (result *v1alpha1.AwsXraySamplingRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsXraySamplingRuleList{}
	err = c.client.Get().
		Resource("awsxraysamplingrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsXraySamplingRules.
func (c *awsXraySamplingRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsxraysamplingrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsXraySamplingRule and creates it.  Returns the server's representation of the awsXraySamplingRule, and an error, if there is any.
func (c *awsXraySamplingRules) Create(awsXraySamplingRule *v1alpha1.AwsXraySamplingRule) (result *v1alpha1.AwsXraySamplingRule, err error) {
	result = &v1alpha1.AwsXraySamplingRule{}
	err = c.client.Post().
		Resource("awsxraysamplingrules").
		Body(awsXraySamplingRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsXraySamplingRule and updates it. Returns the server's representation of the awsXraySamplingRule, and an error, if there is any.
func (c *awsXraySamplingRules) Update(awsXraySamplingRule *v1alpha1.AwsXraySamplingRule) (result *v1alpha1.AwsXraySamplingRule, err error) {
	result = &v1alpha1.AwsXraySamplingRule{}
	err = c.client.Put().
		Resource("awsxraysamplingrules").
		Name(awsXraySamplingRule.Name).
		Body(awsXraySamplingRule).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsXraySamplingRules) UpdateStatus(awsXraySamplingRule *v1alpha1.AwsXraySamplingRule) (result *v1alpha1.AwsXraySamplingRule, err error) {
	result = &v1alpha1.AwsXraySamplingRule{}
	err = c.client.Put().
		Resource("awsxraysamplingrules").
		Name(awsXraySamplingRule.Name).
		SubResource("status").
		Body(awsXraySamplingRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsXraySamplingRule and deletes it. Returns an error if one occurs.
func (c *awsXraySamplingRules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsxraysamplingrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsXraySamplingRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsxraysamplingrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsXraySamplingRule.
func (c *awsXraySamplingRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsXraySamplingRule, err error) {
	result = &v1alpha1.AwsXraySamplingRule{}
	err = c.client.Patch(pt).
		Resource("awsxraysamplingrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
