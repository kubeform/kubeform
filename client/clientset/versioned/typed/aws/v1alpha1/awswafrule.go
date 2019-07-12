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

// AwsWafRulesGetter has a method to return a AwsWafRuleInterface.
// A group's client should implement this interface.
type AwsWafRulesGetter interface {
	AwsWafRules() AwsWafRuleInterface
}

// AwsWafRuleInterface has methods to work with AwsWafRule resources.
type AwsWafRuleInterface interface {
	Create(*v1alpha1.AwsWafRule) (*v1alpha1.AwsWafRule, error)
	Update(*v1alpha1.AwsWafRule) (*v1alpha1.AwsWafRule, error)
	UpdateStatus(*v1alpha1.AwsWafRule) (*v1alpha1.AwsWafRule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsWafRule, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsWafRuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafRule, err error)
	AwsWafRuleExpansion
}

// awsWafRules implements AwsWafRuleInterface
type awsWafRules struct {
	client rest.Interface
}

// newAwsWafRules returns a AwsWafRules
func newAwsWafRules(c *AwsV1alpha1Client) *awsWafRules {
	return &awsWafRules{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsWafRule, and returns the corresponding awsWafRule object, and an error if there is any.
func (c *awsWafRules) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsWafRule, err error) {
	result = &v1alpha1.AwsWafRule{}
	err = c.client.Get().
		Resource("awswafrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsWafRules that match those selectors.
func (c *awsWafRules) List(opts v1.ListOptions) (result *v1alpha1.AwsWafRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsWafRuleList{}
	err = c.client.Get().
		Resource("awswafrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsWafRules.
func (c *awsWafRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awswafrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsWafRule and creates it.  Returns the server's representation of the awsWafRule, and an error, if there is any.
func (c *awsWafRules) Create(awsWafRule *v1alpha1.AwsWafRule) (result *v1alpha1.AwsWafRule, err error) {
	result = &v1alpha1.AwsWafRule{}
	err = c.client.Post().
		Resource("awswafrules").
		Body(awsWafRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsWafRule and updates it. Returns the server's representation of the awsWafRule, and an error, if there is any.
func (c *awsWafRules) Update(awsWafRule *v1alpha1.AwsWafRule) (result *v1alpha1.AwsWafRule, err error) {
	result = &v1alpha1.AwsWafRule{}
	err = c.client.Put().
		Resource("awswafrules").
		Name(awsWafRule.Name).
		Body(awsWafRule).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsWafRules) UpdateStatus(awsWafRule *v1alpha1.AwsWafRule) (result *v1alpha1.AwsWafRule, err error) {
	result = &v1alpha1.AwsWafRule{}
	err = c.client.Put().
		Resource("awswafrules").
		Name(awsWafRule.Name).
		SubResource("status").
		Body(awsWafRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsWafRule and deletes it. Returns an error if one occurs.
func (c *awsWafRules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awswafrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsWafRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awswafrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsWafRule.
func (c *awsWafRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafRule, err error) {
	result = &v1alpha1.AwsWafRule{}
	err = c.client.Patch(pt).
		Resource("awswafrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
