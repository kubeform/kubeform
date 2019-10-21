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

// XraySamplingRulesGetter has a method to return a XraySamplingRuleInterface.
// A group's client should implement this interface.
type XraySamplingRulesGetter interface {
	XraySamplingRules(namespace string) XraySamplingRuleInterface
}

// XraySamplingRuleInterface has methods to work with XraySamplingRule resources.
type XraySamplingRuleInterface interface {
	Create(*v1alpha1.XraySamplingRule) (*v1alpha1.XraySamplingRule, error)
	Update(*v1alpha1.XraySamplingRule) (*v1alpha1.XraySamplingRule, error)
	UpdateStatus(*v1alpha1.XraySamplingRule) (*v1alpha1.XraySamplingRule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.XraySamplingRule, error)
	List(opts v1.ListOptions) (*v1alpha1.XraySamplingRuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.XraySamplingRule, err error)
	XraySamplingRuleExpansion
}

// xraySamplingRules implements XraySamplingRuleInterface
type xraySamplingRules struct {
	client rest.Interface
	ns     string
}

// newXraySamplingRules returns a XraySamplingRules
func newXraySamplingRules(c *AwsV1alpha1Client, namespace string) *xraySamplingRules {
	return &xraySamplingRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the xraySamplingRule, and returns the corresponding xraySamplingRule object, and an error if there is any.
func (c *xraySamplingRules) Get(name string, options v1.GetOptions) (result *v1alpha1.XraySamplingRule, err error) {
	result = &v1alpha1.XraySamplingRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("xraysamplingrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of XraySamplingRules that match those selectors.
func (c *xraySamplingRules) List(opts v1.ListOptions) (result *v1alpha1.XraySamplingRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.XraySamplingRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("xraysamplingrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested xraySamplingRules.
func (c *xraySamplingRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("xraysamplingrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a xraySamplingRule and creates it.  Returns the server's representation of the xraySamplingRule, and an error, if there is any.
func (c *xraySamplingRules) Create(xraySamplingRule *v1alpha1.XraySamplingRule) (result *v1alpha1.XraySamplingRule, err error) {
	result = &v1alpha1.XraySamplingRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("xraysamplingrules").
		Body(xraySamplingRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a xraySamplingRule and updates it. Returns the server's representation of the xraySamplingRule, and an error, if there is any.
func (c *xraySamplingRules) Update(xraySamplingRule *v1alpha1.XraySamplingRule) (result *v1alpha1.XraySamplingRule, err error) {
	result = &v1alpha1.XraySamplingRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("xraysamplingrules").
		Name(xraySamplingRule.Name).
		Body(xraySamplingRule).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *xraySamplingRules) UpdateStatus(xraySamplingRule *v1alpha1.XraySamplingRule) (result *v1alpha1.XraySamplingRule, err error) {
	result = &v1alpha1.XraySamplingRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("xraysamplingrules").
		Name(xraySamplingRule.Name).
		SubResource("status").
		Body(xraySamplingRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the xraySamplingRule and deletes it. Returns an error if one occurs.
func (c *xraySamplingRules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("xraysamplingrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *xraySamplingRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("xraysamplingrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched xraySamplingRule.
func (c *xraySamplingRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.XraySamplingRule, err error) {
	result = &v1alpha1.XraySamplingRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("xraysamplingrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}