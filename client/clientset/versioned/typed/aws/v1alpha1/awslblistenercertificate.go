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

// AwsLbListenerCertificatesGetter has a method to return a AwsLbListenerCertificateInterface.
// A group's client should implement this interface.
type AwsLbListenerCertificatesGetter interface {
	AwsLbListenerCertificates() AwsLbListenerCertificateInterface
}

// AwsLbListenerCertificateInterface has methods to work with AwsLbListenerCertificate resources.
type AwsLbListenerCertificateInterface interface {
	Create(*v1alpha1.AwsLbListenerCertificate) (*v1alpha1.AwsLbListenerCertificate, error)
	Update(*v1alpha1.AwsLbListenerCertificate) (*v1alpha1.AwsLbListenerCertificate, error)
	UpdateStatus(*v1alpha1.AwsLbListenerCertificate) (*v1alpha1.AwsLbListenerCertificate, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsLbListenerCertificate, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsLbListenerCertificateList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsLbListenerCertificate, err error)
	AwsLbListenerCertificateExpansion
}

// awsLbListenerCertificates implements AwsLbListenerCertificateInterface
type awsLbListenerCertificates struct {
	client rest.Interface
}

// newAwsLbListenerCertificates returns a AwsLbListenerCertificates
func newAwsLbListenerCertificates(c *AwsV1alpha1Client) *awsLbListenerCertificates {
	return &awsLbListenerCertificates{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsLbListenerCertificate, and returns the corresponding awsLbListenerCertificate object, and an error if there is any.
func (c *awsLbListenerCertificates) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsLbListenerCertificate, err error) {
	result = &v1alpha1.AwsLbListenerCertificate{}
	err = c.client.Get().
		Resource("awslblistenercertificates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsLbListenerCertificates that match those selectors.
func (c *awsLbListenerCertificates) List(opts v1.ListOptions) (result *v1alpha1.AwsLbListenerCertificateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsLbListenerCertificateList{}
	err = c.client.Get().
		Resource("awslblistenercertificates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsLbListenerCertificates.
func (c *awsLbListenerCertificates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awslblistenercertificates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsLbListenerCertificate and creates it.  Returns the server's representation of the awsLbListenerCertificate, and an error, if there is any.
func (c *awsLbListenerCertificates) Create(awsLbListenerCertificate *v1alpha1.AwsLbListenerCertificate) (result *v1alpha1.AwsLbListenerCertificate, err error) {
	result = &v1alpha1.AwsLbListenerCertificate{}
	err = c.client.Post().
		Resource("awslblistenercertificates").
		Body(awsLbListenerCertificate).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsLbListenerCertificate and updates it. Returns the server's representation of the awsLbListenerCertificate, and an error, if there is any.
func (c *awsLbListenerCertificates) Update(awsLbListenerCertificate *v1alpha1.AwsLbListenerCertificate) (result *v1alpha1.AwsLbListenerCertificate, err error) {
	result = &v1alpha1.AwsLbListenerCertificate{}
	err = c.client.Put().
		Resource("awslblistenercertificates").
		Name(awsLbListenerCertificate.Name).
		Body(awsLbListenerCertificate).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsLbListenerCertificates) UpdateStatus(awsLbListenerCertificate *v1alpha1.AwsLbListenerCertificate) (result *v1alpha1.AwsLbListenerCertificate, err error) {
	result = &v1alpha1.AwsLbListenerCertificate{}
	err = c.client.Put().
		Resource("awslblistenercertificates").
		Name(awsLbListenerCertificate.Name).
		SubResource("status").
		Body(awsLbListenerCertificate).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsLbListenerCertificate and deletes it. Returns an error if one occurs.
func (c *awsLbListenerCertificates) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awslblistenercertificates").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsLbListenerCertificates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awslblistenercertificates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsLbListenerCertificate.
func (c *awsLbListenerCertificates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsLbListenerCertificate, err error) {
	result = &v1alpha1.AwsLbListenerCertificate{}
	err = c.client.Patch(pt).
		Resource("awslblistenercertificates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
