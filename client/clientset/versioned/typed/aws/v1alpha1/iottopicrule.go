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

// IotTopicRulesGetter has a method to return a IotTopicRuleInterface.
// A group's client should implement this interface.
type IotTopicRulesGetter interface {
	IotTopicRules(namespace string) IotTopicRuleInterface
}

// IotTopicRuleInterface has methods to work with IotTopicRule resources.
type IotTopicRuleInterface interface {
	Create(*v1alpha1.IotTopicRule) (*v1alpha1.IotTopicRule, error)
	Update(*v1alpha1.IotTopicRule) (*v1alpha1.IotTopicRule, error)
	UpdateStatus(*v1alpha1.IotTopicRule) (*v1alpha1.IotTopicRule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.IotTopicRule, error)
	List(opts v1.ListOptions) (*v1alpha1.IotTopicRuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IotTopicRule, err error)
	IotTopicRuleExpansion
}

// iotTopicRules implements IotTopicRuleInterface
type iotTopicRules struct {
	client rest.Interface
	ns     string
}

// newIotTopicRules returns a IotTopicRules
func newIotTopicRules(c *AwsV1alpha1Client, namespace string) *iotTopicRules {
	return &iotTopicRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the iotTopicRule, and returns the corresponding iotTopicRule object, and an error if there is any.
func (c *iotTopicRules) Get(name string, options v1.GetOptions) (result *v1alpha1.IotTopicRule, err error) {
	result = &v1alpha1.IotTopicRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iottopicrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IotTopicRules that match those selectors.
func (c *iotTopicRules) List(opts v1.ListOptions) (result *v1alpha1.IotTopicRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.IotTopicRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iottopicrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iotTopicRules.
func (c *iotTopicRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("iottopicrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a iotTopicRule and creates it.  Returns the server's representation of the iotTopicRule, and an error, if there is any.
func (c *iotTopicRules) Create(iotTopicRule *v1alpha1.IotTopicRule) (result *v1alpha1.IotTopicRule, err error) {
	result = &v1alpha1.IotTopicRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("iottopicrules").
		Body(iotTopicRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a iotTopicRule and updates it. Returns the server's representation of the iotTopicRule, and an error, if there is any.
func (c *iotTopicRules) Update(iotTopicRule *v1alpha1.IotTopicRule) (result *v1alpha1.IotTopicRule, err error) {
	result = &v1alpha1.IotTopicRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iottopicrules").
		Name(iotTopicRule.Name).
		Body(iotTopicRule).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *iotTopicRules) UpdateStatus(iotTopicRule *v1alpha1.IotTopicRule) (result *v1alpha1.IotTopicRule, err error) {
	result = &v1alpha1.IotTopicRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iottopicrules").
		Name(iotTopicRule.Name).
		SubResource("status").
		Body(iotTopicRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the iotTopicRule and deletes it. Returns an error if one occurs.
func (c *iotTopicRules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iottopicrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *iotTopicRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iottopicrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched iotTopicRule.
func (c *iotTopicRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IotTopicRule, err error) {
	result = &v1alpha1.IotTopicRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("iottopicrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}