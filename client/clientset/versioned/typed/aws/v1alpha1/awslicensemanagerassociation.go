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

// AwsLicensemanagerAssociationsGetter has a method to return a AwsLicensemanagerAssociationInterface.
// A group's client should implement this interface.
type AwsLicensemanagerAssociationsGetter interface {
	AwsLicensemanagerAssociations() AwsLicensemanagerAssociationInterface
}

// AwsLicensemanagerAssociationInterface has methods to work with AwsLicensemanagerAssociation resources.
type AwsLicensemanagerAssociationInterface interface {
	Create(*v1alpha1.AwsLicensemanagerAssociation) (*v1alpha1.AwsLicensemanagerAssociation, error)
	Update(*v1alpha1.AwsLicensemanagerAssociation) (*v1alpha1.AwsLicensemanagerAssociation, error)
	UpdateStatus(*v1alpha1.AwsLicensemanagerAssociation) (*v1alpha1.AwsLicensemanagerAssociation, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsLicensemanagerAssociation, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsLicensemanagerAssociationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsLicensemanagerAssociation, err error)
	AwsLicensemanagerAssociationExpansion
}

// awsLicensemanagerAssociations implements AwsLicensemanagerAssociationInterface
type awsLicensemanagerAssociations struct {
	client rest.Interface
}

// newAwsLicensemanagerAssociations returns a AwsLicensemanagerAssociations
func newAwsLicensemanagerAssociations(c *AwsV1alpha1Client) *awsLicensemanagerAssociations {
	return &awsLicensemanagerAssociations{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsLicensemanagerAssociation, and returns the corresponding awsLicensemanagerAssociation object, and an error if there is any.
func (c *awsLicensemanagerAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsLicensemanagerAssociation, err error) {
	result = &v1alpha1.AwsLicensemanagerAssociation{}
	err = c.client.Get().
		Resource("awslicensemanagerassociations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsLicensemanagerAssociations that match those selectors.
func (c *awsLicensemanagerAssociations) List(opts v1.ListOptions) (result *v1alpha1.AwsLicensemanagerAssociationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsLicensemanagerAssociationList{}
	err = c.client.Get().
		Resource("awslicensemanagerassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsLicensemanagerAssociations.
func (c *awsLicensemanagerAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awslicensemanagerassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsLicensemanagerAssociation and creates it.  Returns the server's representation of the awsLicensemanagerAssociation, and an error, if there is any.
func (c *awsLicensemanagerAssociations) Create(awsLicensemanagerAssociation *v1alpha1.AwsLicensemanagerAssociation) (result *v1alpha1.AwsLicensemanagerAssociation, err error) {
	result = &v1alpha1.AwsLicensemanagerAssociation{}
	err = c.client.Post().
		Resource("awslicensemanagerassociations").
		Body(awsLicensemanagerAssociation).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsLicensemanagerAssociation and updates it. Returns the server's representation of the awsLicensemanagerAssociation, and an error, if there is any.
func (c *awsLicensemanagerAssociations) Update(awsLicensemanagerAssociation *v1alpha1.AwsLicensemanagerAssociation) (result *v1alpha1.AwsLicensemanagerAssociation, err error) {
	result = &v1alpha1.AwsLicensemanagerAssociation{}
	err = c.client.Put().
		Resource("awslicensemanagerassociations").
		Name(awsLicensemanagerAssociation.Name).
		Body(awsLicensemanagerAssociation).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsLicensemanagerAssociations) UpdateStatus(awsLicensemanagerAssociation *v1alpha1.AwsLicensemanagerAssociation) (result *v1alpha1.AwsLicensemanagerAssociation, err error) {
	result = &v1alpha1.AwsLicensemanagerAssociation{}
	err = c.client.Put().
		Resource("awslicensemanagerassociations").
		Name(awsLicensemanagerAssociation.Name).
		SubResource("status").
		Body(awsLicensemanagerAssociation).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsLicensemanagerAssociation and deletes it. Returns an error if one occurs.
func (c *awsLicensemanagerAssociations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awslicensemanagerassociations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsLicensemanagerAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awslicensemanagerassociations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsLicensemanagerAssociation.
func (c *awsLicensemanagerAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsLicensemanagerAssociation, err error) {
	result = &v1alpha1.AwsLicensemanagerAssociation{}
	err = c.client.Patch(pt).
		Resource("awslicensemanagerassociations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
