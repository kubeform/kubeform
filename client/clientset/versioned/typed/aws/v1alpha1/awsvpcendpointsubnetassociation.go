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

// AwsVpcEndpointSubnetAssociationsGetter has a method to return a AwsVpcEndpointSubnetAssociationInterface.
// A group's client should implement this interface.
type AwsVpcEndpointSubnetAssociationsGetter interface {
	AwsVpcEndpointSubnetAssociations() AwsVpcEndpointSubnetAssociationInterface
}

// AwsVpcEndpointSubnetAssociationInterface has methods to work with AwsVpcEndpointSubnetAssociation resources.
type AwsVpcEndpointSubnetAssociationInterface interface {
	Create(*v1alpha1.AwsVpcEndpointSubnetAssociation) (*v1alpha1.AwsVpcEndpointSubnetAssociation, error)
	Update(*v1alpha1.AwsVpcEndpointSubnetAssociation) (*v1alpha1.AwsVpcEndpointSubnetAssociation, error)
	UpdateStatus(*v1alpha1.AwsVpcEndpointSubnetAssociation) (*v1alpha1.AwsVpcEndpointSubnetAssociation, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsVpcEndpointSubnetAssociation, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsVpcEndpointSubnetAssociationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsVpcEndpointSubnetAssociation, err error)
	AwsVpcEndpointSubnetAssociationExpansion
}

// awsVpcEndpointSubnetAssociations implements AwsVpcEndpointSubnetAssociationInterface
type awsVpcEndpointSubnetAssociations struct {
	client rest.Interface
}

// newAwsVpcEndpointSubnetAssociations returns a AwsVpcEndpointSubnetAssociations
func newAwsVpcEndpointSubnetAssociations(c *AwsV1alpha1Client) *awsVpcEndpointSubnetAssociations {
	return &awsVpcEndpointSubnetAssociations{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsVpcEndpointSubnetAssociation, and returns the corresponding awsVpcEndpointSubnetAssociation object, and an error if there is any.
func (c *awsVpcEndpointSubnetAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsVpcEndpointSubnetAssociation, err error) {
	result = &v1alpha1.AwsVpcEndpointSubnetAssociation{}
	err = c.client.Get().
		Resource("awsvpcendpointsubnetassociations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsVpcEndpointSubnetAssociations that match those selectors.
func (c *awsVpcEndpointSubnetAssociations) List(opts v1.ListOptions) (result *v1alpha1.AwsVpcEndpointSubnetAssociationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsVpcEndpointSubnetAssociationList{}
	err = c.client.Get().
		Resource("awsvpcendpointsubnetassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsVpcEndpointSubnetAssociations.
func (c *awsVpcEndpointSubnetAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsvpcendpointsubnetassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsVpcEndpointSubnetAssociation and creates it.  Returns the server's representation of the awsVpcEndpointSubnetAssociation, and an error, if there is any.
func (c *awsVpcEndpointSubnetAssociations) Create(awsVpcEndpointSubnetAssociation *v1alpha1.AwsVpcEndpointSubnetAssociation) (result *v1alpha1.AwsVpcEndpointSubnetAssociation, err error) {
	result = &v1alpha1.AwsVpcEndpointSubnetAssociation{}
	err = c.client.Post().
		Resource("awsvpcendpointsubnetassociations").
		Body(awsVpcEndpointSubnetAssociation).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsVpcEndpointSubnetAssociation and updates it. Returns the server's representation of the awsVpcEndpointSubnetAssociation, and an error, if there is any.
func (c *awsVpcEndpointSubnetAssociations) Update(awsVpcEndpointSubnetAssociation *v1alpha1.AwsVpcEndpointSubnetAssociation) (result *v1alpha1.AwsVpcEndpointSubnetAssociation, err error) {
	result = &v1alpha1.AwsVpcEndpointSubnetAssociation{}
	err = c.client.Put().
		Resource("awsvpcendpointsubnetassociations").
		Name(awsVpcEndpointSubnetAssociation.Name).
		Body(awsVpcEndpointSubnetAssociation).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsVpcEndpointSubnetAssociations) UpdateStatus(awsVpcEndpointSubnetAssociation *v1alpha1.AwsVpcEndpointSubnetAssociation) (result *v1alpha1.AwsVpcEndpointSubnetAssociation, err error) {
	result = &v1alpha1.AwsVpcEndpointSubnetAssociation{}
	err = c.client.Put().
		Resource("awsvpcendpointsubnetassociations").
		Name(awsVpcEndpointSubnetAssociation.Name).
		SubResource("status").
		Body(awsVpcEndpointSubnetAssociation).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsVpcEndpointSubnetAssociation and deletes it. Returns an error if one occurs.
func (c *awsVpcEndpointSubnetAssociations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsvpcendpointsubnetassociations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsVpcEndpointSubnetAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsvpcendpointsubnetassociations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsVpcEndpointSubnetAssociation.
func (c *awsVpcEndpointSubnetAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsVpcEndpointSubnetAssociation, err error) {
	result = &v1alpha1.AwsVpcEndpointSubnetAssociation{}
	err = c.client.Patch(pt).
		Resource("awsvpcendpointsubnetassociations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
