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

// AwsRoute53ResolverRuleAssociationsGetter has a method to return a AwsRoute53ResolverRuleAssociationInterface.
// A group's client should implement this interface.
type AwsRoute53ResolverRuleAssociationsGetter interface {
	AwsRoute53ResolverRuleAssociations() AwsRoute53ResolverRuleAssociationInterface
}

// AwsRoute53ResolverRuleAssociationInterface has methods to work with AwsRoute53ResolverRuleAssociation resources.
type AwsRoute53ResolverRuleAssociationInterface interface {
	Create(*v1alpha1.AwsRoute53ResolverRuleAssociation) (*v1alpha1.AwsRoute53ResolverRuleAssociation, error)
	Update(*v1alpha1.AwsRoute53ResolverRuleAssociation) (*v1alpha1.AwsRoute53ResolverRuleAssociation, error)
	UpdateStatus(*v1alpha1.AwsRoute53ResolverRuleAssociation) (*v1alpha1.AwsRoute53ResolverRuleAssociation, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsRoute53ResolverRuleAssociation, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsRoute53ResolverRuleAssociationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsRoute53ResolverRuleAssociation, err error)
	AwsRoute53ResolverRuleAssociationExpansion
}

// awsRoute53ResolverRuleAssociations implements AwsRoute53ResolverRuleAssociationInterface
type awsRoute53ResolverRuleAssociations struct {
	client rest.Interface
}

// newAwsRoute53ResolverRuleAssociations returns a AwsRoute53ResolverRuleAssociations
func newAwsRoute53ResolverRuleAssociations(c *AwsV1alpha1Client) *awsRoute53ResolverRuleAssociations {
	return &awsRoute53ResolverRuleAssociations{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsRoute53ResolverRuleAssociation, and returns the corresponding awsRoute53ResolverRuleAssociation object, and an error if there is any.
func (c *awsRoute53ResolverRuleAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsRoute53ResolverRuleAssociation, err error) {
	result = &v1alpha1.AwsRoute53ResolverRuleAssociation{}
	err = c.client.Get().
		Resource("awsroute53resolverruleassociations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsRoute53ResolverRuleAssociations that match those selectors.
func (c *awsRoute53ResolverRuleAssociations) List(opts v1.ListOptions) (result *v1alpha1.AwsRoute53ResolverRuleAssociationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsRoute53ResolverRuleAssociationList{}
	err = c.client.Get().
		Resource("awsroute53resolverruleassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsRoute53ResolverRuleAssociations.
func (c *awsRoute53ResolverRuleAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsroute53resolverruleassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsRoute53ResolverRuleAssociation and creates it.  Returns the server's representation of the awsRoute53ResolverRuleAssociation, and an error, if there is any.
func (c *awsRoute53ResolverRuleAssociations) Create(awsRoute53ResolverRuleAssociation *v1alpha1.AwsRoute53ResolverRuleAssociation) (result *v1alpha1.AwsRoute53ResolverRuleAssociation, err error) {
	result = &v1alpha1.AwsRoute53ResolverRuleAssociation{}
	err = c.client.Post().
		Resource("awsroute53resolverruleassociations").
		Body(awsRoute53ResolverRuleAssociation).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsRoute53ResolverRuleAssociation and updates it. Returns the server's representation of the awsRoute53ResolverRuleAssociation, and an error, if there is any.
func (c *awsRoute53ResolverRuleAssociations) Update(awsRoute53ResolverRuleAssociation *v1alpha1.AwsRoute53ResolverRuleAssociation) (result *v1alpha1.AwsRoute53ResolverRuleAssociation, err error) {
	result = &v1alpha1.AwsRoute53ResolverRuleAssociation{}
	err = c.client.Put().
		Resource("awsroute53resolverruleassociations").
		Name(awsRoute53ResolverRuleAssociation.Name).
		Body(awsRoute53ResolverRuleAssociation).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsRoute53ResolverRuleAssociations) UpdateStatus(awsRoute53ResolverRuleAssociation *v1alpha1.AwsRoute53ResolverRuleAssociation) (result *v1alpha1.AwsRoute53ResolverRuleAssociation, err error) {
	result = &v1alpha1.AwsRoute53ResolverRuleAssociation{}
	err = c.client.Put().
		Resource("awsroute53resolverruleassociations").
		Name(awsRoute53ResolverRuleAssociation.Name).
		SubResource("status").
		Body(awsRoute53ResolverRuleAssociation).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsRoute53ResolverRuleAssociation and deletes it. Returns an error if one occurs.
func (c *awsRoute53ResolverRuleAssociations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsroute53resolverruleassociations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsRoute53ResolverRuleAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsroute53resolverruleassociations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsRoute53ResolverRuleAssociation.
func (c *awsRoute53ResolverRuleAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsRoute53ResolverRuleAssociation, err error) {
	result = &v1alpha1.AwsRoute53ResolverRuleAssociation{}
	err = c.client.Patch(pt).
		Resource("awsroute53resolverruleassociations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
