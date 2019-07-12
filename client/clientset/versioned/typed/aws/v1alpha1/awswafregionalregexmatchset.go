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

// AwsWafregionalRegexMatchSetsGetter has a method to return a AwsWafregionalRegexMatchSetInterface.
// A group's client should implement this interface.
type AwsWafregionalRegexMatchSetsGetter interface {
	AwsWafregionalRegexMatchSets() AwsWafregionalRegexMatchSetInterface
}

// AwsWafregionalRegexMatchSetInterface has methods to work with AwsWafregionalRegexMatchSet resources.
type AwsWafregionalRegexMatchSetInterface interface {
	Create(*v1alpha1.AwsWafregionalRegexMatchSet) (*v1alpha1.AwsWafregionalRegexMatchSet, error)
	Update(*v1alpha1.AwsWafregionalRegexMatchSet) (*v1alpha1.AwsWafregionalRegexMatchSet, error)
	UpdateStatus(*v1alpha1.AwsWafregionalRegexMatchSet) (*v1alpha1.AwsWafregionalRegexMatchSet, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsWafregionalRegexMatchSet, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsWafregionalRegexMatchSetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafregionalRegexMatchSet, err error)
	AwsWafregionalRegexMatchSetExpansion
}

// awsWafregionalRegexMatchSets implements AwsWafregionalRegexMatchSetInterface
type awsWafregionalRegexMatchSets struct {
	client rest.Interface
}

// newAwsWafregionalRegexMatchSets returns a AwsWafregionalRegexMatchSets
func newAwsWafregionalRegexMatchSets(c *AwsV1alpha1Client) *awsWafregionalRegexMatchSets {
	return &awsWafregionalRegexMatchSets{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsWafregionalRegexMatchSet, and returns the corresponding awsWafregionalRegexMatchSet object, and an error if there is any.
func (c *awsWafregionalRegexMatchSets) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsWafregionalRegexMatchSet, err error) {
	result = &v1alpha1.AwsWafregionalRegexMatchSet{}
	err = c.client.Get().
		Resource("awswafregionalregexmatchsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsWafregionalRegexMatchSets that match those selectors.
func (c *awsWafregionalRegexMatchSets) List(opts v1.ListOptions) (result *v1alpha1.AwsWafregionalRegexMatchSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsWafregionalRegexMatchSetList{}
	err = c.client.Get().
		Resource("awswafregionalregexmatchsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsWafregionalRegexMatchSets.
func (c *awsWafregionalRegexMatchSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awswafregionalregexmatchsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsWafregionalRegexMatchSet and creates it.  Returns the server's representation of the awsWafregionalRegexMatchSet, and an error, if there is any.
func (c *awsWafregionalRegexMatchSets) Create(awsWafregionalRegexMatchSet *v1alpha1.AwsWafregionalRegexMatchSet) (result *v1alpha1.AwsWafregionalRegexMatchSet, err error) {
	result = &v1alpha1.AwsWafregionalRegexMatchSet{}
	err = c.client.Post().
		Resource("awswafregionalregexmatchsets").
		Body(awsWafregionalRegexMatchSet).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsWafregionalRegexMatchSet and updates it. Returns the server's representation of the awsWafregionalRegexMatchSet, and an error, if there is any.
func (c *awsWafregionalRegexMatchSets) Update(awsWafregionalRegexMatchSet *v1alpha1.AwsWafregionalRegexMatchSet) (result *v1alpha1.AwsWafregionalRegexMatchSet, err error) {
	result = &v1alpha1.AwsWafregionalRegexMatchSet{}
	err = c.client.Put().
		Resource("awswafregionalregexmatchsets").
		Name(awsWafregionalRegexMatchSet.Name).
		Body(awsWafregionalRegexMatchSet).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsWafregionalRegexMatchSets) UpdateStatus(awsWafregionalRegexMatchSet *v1alpha1.AwsWafregionalRegexMatchSet) (result *v1alpha1.AwsWafregionalRegexMatchSet, err error) {
	result = &v1alpha1.AwsWafregionalRegexMatchSet{}
	err = c.client.Put().
		Resource("awswafregionalregexmatchsets").
		Name(awsWafregionalRegexMatchSet.Name).
		SubResource("status").
		Body(awsWafregionalRegexMatchSet).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsWafregionalRegexMatchSet and deletes it. Returns an error if one occurs.
func (c *awsWafregionalRegexMatchSets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awswafregionalregexmatchsets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsWafregionalRegexMatchSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awswafregionalregexmatchsets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsWafregionalRegexMatchSet.
func (c *awsWafregionalRegexMatchSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafregionalRegexMatchSet, err error) {
	result = &v1alpha1.AwsWafregionalRegexMatchSet{}
	err = c.client.Patch(pt).
		Resource("awswafregionalregexmatchsets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
