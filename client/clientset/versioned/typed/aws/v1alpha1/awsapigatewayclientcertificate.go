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

// AwsApiGatewayClientCertificatesGetter has a method to return a AwsApiGatewayClientCertificateInterface.
// A group's client should implement this interface.
type AwsApiGatewayClientCertificatesGetter interface {
	AwsApiGatewayClientCertificates() AwsApiGatewayClientCertificateInterface
}

// AwsApiGatewayClientCertificateInterface has methods to work with AwsApiGatewayClientCertificate resources.
type AwsApiGatewayClientCertificateInterface interface {
	Create(*v1alpha1.AwsApiGatewayClientCertificate) (*v1alpha1.AwsApiGatewayClientCertificate, error)
	Update(*v1alpha1.AwsApiGatewayClientCertificate) (*v1alpha1.AwsApiGatewayClientCertificate, error)
	UpdateStatus(*v1alpha1.AwsApiGatewayClientCertificate) (*v1alpha1.AwsApiGatewayClientCertificate, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsApiGatewayClientCertificate, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsApiGatewayClientCertificateList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayClientCertificate, err error)
	AwsApiGatewayClientCertificateExpansion
}

// awsApiGatewayClientCertificates implements AwsApiGatewayClientCertificateInterface
type awsApiGatewayClientCertificates struct {
	client rest.Interface
}

// newAwsApiGatewayClientCertificates returns a AwsApiGatewayClientCertificates
func newAwsApiGatewayClientCertificates(c *AwsV1alpha1Client) *awsApiGatewayClientCertificates {
	return &awsApiGatewayClientCertificates{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsApiGatewayClientCertificate, and returns the corresponding awsApiGatewayClientCertificate object, and an error if there is any.
func (c *awsApiGatewayClientCertificates) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsApiGatewayClientCertificate, err error) {
	result = &v1alpha1.AwsApiGatewayClientCertificate{}
	err = c.client.Get().
		Resource("awsapigatewayclientcertificates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsApiGatewayClientCertificates that match those selectors.
func (c *awsApiGatewayClientCertificates) List(opts v1.ListOptions) (result *v1alpha1.AwsApiGatewayClientCertificateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsApiGatewayClientCertificateList{}
	err = c.client.Get().
		Resource("awsapigatewayclientcertificates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsApiGatewayClientCertificates.
func (c *awsApiGatewayClientCertificates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsapigatewayclientcertificates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsApiGatewayClientCertificate and creates it.  Returns the server's representation of the awsApiGatewayClientCertificate, and an error, if there is any.
func (c *awsApiGatewayClientCertificates) Create(awsApiGatewayClientCertificate *v1alpha1.AwsApiGatewayClientCertificate) (result *v1alpha1.AwsApiGatewayClientCertificate, err error) {
	result = &v1alpha1.AwsApiGatewayClientCertificate{}
	err = c.client.Post().
		Resource("awsapigatewayclientcertificates").
		Body(awsApiGatewayClientCertificate).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsApiGatewayClientCertificate and updates it. Returns the server's representation of the awsApiGatewayClientCertificate, and an error, if there is any.
func (c *awsApiGatewayClientCertificates) Update(awsApiGatewayClientCertificate *v1alpha1.AwsApiGatewayClientCertificate) (result *v1alpha1.AwsApiGatewayClientCertificate, err error) {
	result = &v1alpha1.AwsApiGatewayClientCertificate{}
	err = c.client.Put().
		Resource("awsapigatewayclientcertificates").
		Name(awsApiGatewayClientCertificate.Name).
		Body(awsApiGatewayClientCertificate).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsApiGatewayClientCertificates) UpdateStatus(awsApiGatewayClientCertificate *v1alpha1.AwsApiGatewayClientCertificate) (result *v1alpha1.AwsApiGatewayClientCertificate, err error) {
	result = &v1alpha1.AwsApiGatewayClientCertificate{}
	err = c.client.Put().
		Resource("awsapigatewayclientcertificates").
		Name(awsApiGatewayClientCertificate.Name).
		SubResource("status").
		Body(awsApiGatewayClientCertificate).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsApiGatewayClientCertificate and deletes it. Returns an error if one occurs.
func (c *awsApiGatewayClientCertificates) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsapigatewayclientcertificates").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsApiGatewayClientCertificates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsapigatewayclientcertificates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsApiGatewayClientCertificate.
func (c *awsApiGatewayClientCertificates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayClientCertificate, err error) {
	result = &v1alpha1.AwsApiGatewayClientCertificate{}
	err = c.client.Patch(pt).
		Resource("awsapigatewayclientcertificates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
