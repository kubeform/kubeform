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

// AwsRoute53DelegationSetsGetter has a method to return a AwsRoute53DelegationSetInterface.
// A group's client should implement this interface.
type AwsRoute53DelegationSetsGetter interface {
	AwsRoute53DelegationSets() AwsRoute53DelegationSetInterface
}

// AwsRoute53DelegationSetInterface has methods to work with AwsRoute53DelegationSet resources.
type AwsRoute53DelegationSetInterface interface {
	Create(*v1alpha1.AwsRoute53DelegationSet) (*v1alpha1.AwsRoute53DelegationSet, error)
	Update(*v1alpha1.AwsRoute53DelegationSet) (*v1alpha1.AwsRoute53DelegationSet, error)
	UpdateStatus(*v1alpha1.AwsRoute53DelegationSet) (*v1alpha1.AwsRoute53DelegationSet, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsRoute53DelegationSet, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsRoute53DelegationSetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsRoute53DelegationSet, err error)
	AwsRoute53DelegationSetExpansion
}

// awsRoute53DelegationSets implements AwsRoute53DelegationSetInterface
type awsRoute53DelegationSets struct {
	client rest.Interface
}

// newAwsRoute53DelegationSets returns a AwsRoute53DelegationSets
func newAwsRoute53DelegationSets(c *AwsV1alpha1Client) *awsRoute53DelegationSets {
	return &awsRoute53DelegationSets{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsRoute53DelegationSet, and returns the corresponding awsRoute53DelegationSet object, and an error if there is any.
func (c *awsRoute53DelegationSets) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsRoute53DelegationSet, err error) {
	result = &v1alpha1.AwsRoute53DelegationSet{}
	err = c.client.Get().
		Resource("awsroute53delegationsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsRoute53DelegationSets that match those selectors.
func (c *awsRoute53DelegationSets) List(opts v1.ListOptions) (result *v1alpha1.AwsRoute53DelegationSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsRoute53DelegationSetList{}
	err = c.client.Get().
		Resource("awsroute53delegationsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsRoute53DelegationSets.
func (c *awsRoute53DelegationSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsroute53delegationsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsRoute53DelegationSet and creates it.  Returns the server's representation of the awsRoute53DelegationSet, and an error, if there is any.
func (c *awsRoute53DelegationSets) Create(awsRoute53DelegationSet *v1alpha1.AwsRoute53DelegationSet) (result *v1alpha1.AwsRoute53DelegationSet, err error) {
	result = &v1alpha1.AwsRoute53DelegationSet{}
	err = c.client.Post().
		Resource("awsroute53delegationsets").
		Body(awsRoute53DelegationSet).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsRoute53DelegationSet and updates it. Returns the server's representation of the awsRoute53DelegationSet, and an error, if there is any.
func (c *awsRoute53DelegationSets) Update(awsRoute53DelegationSet *v1alpha1.AwsRoute53DelegationSet) (result *v1alpha1.AwsRoute53DelegationSet, err error) {
	result = &v1alpha1.AwsRoute53DelegationSet{}
	err = c.client.Put().
		Resource("awsroute53delegationsets").
		Name(awsRoute53DelegationSet.Name).
		Body(awsRoute53DelegationSet).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsRoute53DelegationSets) UpdateStatus(awsRoute53DelegationSet *v1alpha1.AwsRoute53DelegationSet) (result *v1alpha1.AwsRoute53DelegationSet, err error) {
	result = &v1alpha1.AwsRoute53DelegationSet{}
	err = c.client.Put().
		Resource("awsroute53delegationsets").
		Name(awsRoute53DelegationSet.Name).
		SubResource("status").
		Body(awsRoute53DelegationSet).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsRoute53DelegationSet and deletes it. Returns an error if one occurs.
func (c *awsRoute53DelegationSets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsroute53delegationsets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsRoute53DelegationSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsroute53delegationsets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsRoute53DelegationSet.
func (c *awsRoute53DelegationSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsRoute53DelegationSet, err error) {
	result = &v1alpha1.AwsRoute53DelegationSet{}
	err = c.client.Patch(pt).
		Resource("awsroute53delegationsets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
