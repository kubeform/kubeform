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

// AwsNetworkAclRulesGetter has a method to return a AwsNetworkAclRuleInterface.
// A group's client should implement this interface.
type AwsNetworkAclRulesGetter interface {
	AwsNetworkAclRules() AwsNetworkAclRuleInterface
}

// AwsNetworkAclRuleInterface has methods to work with AwsNetworkAclRule resources.
type AwsNetworkAclRuleInterface interface {
	Create(*v1alpha1.AwsNetworkAclRule) (*v1alpha1.AwsNetworkAclRule, error)
	Update(*v1alpha1.AwsNetworkAclRule) (*v1alpha1.AwsNetworkAclRule, error)
	UpdateStatus(*v1alpha1.AwsNetworkAclRule) (*v1alpha1.AwsNetworkAclRule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsNetworkAclRule, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsNetworkAclRuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsNetworkAclRule, err error)
	AwsNetworkAclRuleExpansion
}

// awsNetworkAclRules implements AwsNetworkAclRuleInterface
type awsNetworkAclRules struct {
	client rest.Interface
}

// newAwsNetworkAclRules returns a AwsNetworkAclRules
func newAwsNetworkAclRules(c *AwsV1alpha1Client) *awsNetworkAclRules {
	return &awsNetworkAclRules{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsNetworkAclRule, and returns the corresponding awsNetworkAclRule object, and an error if there is any.
func (c *awsNetworkAclRules) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsNetworkAclRule, err error) {
	result = &v1alpha1.AwsNetworkAclRule{}
	err = c.client.Get().
		Resource("awsnetworkaclrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsNetworkAclRules that match those selectors.
func (c *awsNetworkAclRules) List(opts v1.ListOptions) (result *v1alpha1.AwsNetworkAclRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsNetworkAclRuleList{}
	err = c.client.Get().
		Resource("awsnetworkaclrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsNetworkAclRules.
func (c *awsNetworkAclRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsnetworkaclrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsNetworkAclRule and creates it.  Returns the server's representation of the awsNetworkAclRule, and an error, if there is any.
func (c *awsNetworkAclRules) Create(awsNetworkAclRule *v1alpha1.AwsNetworkAclRule) (result *v1alpha1.AwsNetworkAclRule, err error) {
	result = &v1alpha1.AwsNetworkAclRule{}
	err = c.client.Post().
		Resource("awsnetworkaclrules").
		Body(awsNetworkAclRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsNetworkAclRule and updates it. Returns the server's representation of the awsNetworkAclRule, and an error, if there is any.
func (c *awsNetworkAclRules) Update(awsNetworkAclRule *v1alpha1.AwsNetworkAclRule) (result *v1alpha1.AwsNetworkAclRule, err error) {
	result = &v1alpha1.AwsNetworkAclRule{}
	err = c.client.Put().
		Resource("awsnetworkaclrules").
		Name(awsNetworkAclRule.Name).
		Body(awsNetworkAclRule).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsNetworkAclRules) UpdateStatus(awsNetworkAclRule *v1alpha1.AwsNetworkAclRule) (result *v1alpha1.AwsNetworkAclRule, err error) {
	result = &v1alpha1.AwsNetworkAclRule{}
	err = c.client.Put().
		Resource("awsnetworkaclrules").
		Name(awsNetworkAclRule.Name).
		SubResource("status").
		Body(awsNetworkAclRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsNetworkAclRule and deletes it. Returns an error if one occurs.
func (c *awsNetworkAclRules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsnetworkaclrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsNetworkAclRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsnetworkaclrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsNetworkAclRule.
func (c *awsNetworkAclRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsNetworkAclRule, err error) {
	result = &v1alpha1.AwsNetworkAclRule{}
	err = c.client.Patch(pt).
		Resource("awsnetworkaclrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
