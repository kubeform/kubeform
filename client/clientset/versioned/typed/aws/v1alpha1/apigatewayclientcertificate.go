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

// ApiGatewayClientCertificatesGetter has a method to return a ApiGatewayClientCertificateInterface.
// A group's client should implement this interface.
type ApiGatewayClientCertificatesGetter interface {
	ApiGatewayClientCertificates() ApiGatewayClientCertificateInterface
}

// ApiGatewayClientCertificateInterface has methods to work with ApiGatewayClientCertificate resources.
type ApiGatewayClientCertificateInterface interface {
	Create(*v1alpha1.ApiGatewayClientCertificate) (*v1alpha1.ApiGatewayClientCertificate, error)
	Update(*v1alpha1.ApiGatewayClientCertificate) (*v1alpha1.ApiGatewayClientCertificate, error)
	UpdateStatus(*v1alpha1.ApiGatewayClientCertificate) (*v1alpha1.ApiGatewayClientCertificate, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ApiGatewayClientCertificate, error)
	List(opts v1.ListOptions) (*v1alpha1.ApiGatewayClientCertificateList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiGatewayClientCertificate, err error)
	ApiGatewayClientCertificateExpansion
}

// apiGatewayClientCertificates implements ApiGatewayClientCertificateInterface
type apiGatewayClientCertificates struct {
	client rest.Interface
}

// newApiGatewayClientCertificates returns a ApiGatewayClientCertificates
func newApiGatewayClientCertificates(c *AwsV1alpha1Client) *apiGatewayClientCertificates {
	return &apiGatewayClientCertificates{
		client: c.RESTClient(),
	}
}

// Get takes name of the apiGatewayClientCertificate, and returns the corresponding apiGatewayClientCertificate object, and an error if there is any.
func (c *apiGatewayClientCertificates) Get(name string, options v1.GetOptions) (result *v1alpha1.ApiGatewayClientCertificate, err error) {
	result = &v1alpha1.ApiGatewayClientCertificate{}
	err = c.client.Get().
		Resource("apigatewayclientcertificates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApiGatewayClientCertificates that match those selectors.
func (c *apiGatewayClientCertificates) List(opts v1.ListOptions) (result *v1alpha1.ApiGatewayClientCertificateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ApiGatewayClientCertificateList{}
	err = c.client.Get().
		Resource("apigatewayclientcertificates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested apiGatewayClientCertificates.
func (c *apiGatewayClientCertificates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("apigatewayclientcertificates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a apiGatewayClientCertificate and creates it.  Returns the server's representation of the apiGatewayClientCertificate, and an error, if there is any.
func (c *apiGatewayClientCertificates) Create(apiGatewayClientCertificate *v1alpha1.ApiGatewayClientCertificate) (result *v1alpha1.ApiGatewayClientCertificate, err error) {
	result = &v1alpha1.ApiGatewayClientCertificate{}
	err = c.client.Post().
		Resource("apigatewayclientcertificates").
		Body(apiGatewayClientCertificate).
		Do().
		Into(result)
	return
}

// Update takes the representation of a apiGatewayClientCertificate and updates it. Returns the server's representation of the apiGatewayClientCertificate, and an error, if there is any.
func (c *apiGatewayClientCertificates) Update(apiGatewayClientCertificate *v1alpha1.ApiGatewayClientCertificate) (result *v1alpha1.ApiGatewayClientCertificate, err error) {
	result = &v1alpha1.ApiGatewayClientCertificate{}
	err = c.client.Put().
		Resource("apigatewayclientcertificates").
		Name(apiGatewayClientCertificate.Name).
		Body(apiGatewayClientCertificate).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *apiGatewayClientCertificates) UpdateStatus(apiGatewayClientCertificate *v1alpha1.ApiGatewayClientCertificate) (result *v1alpha1.ApiGatewayClientCertificate, err error) {
	result = &v1alpha1.ApiGatewayClientCertificate{}
	err = c.client.Put().
		Resource("apigatewayclientcertificates").
		Name(apiGatewayClientCertificate.Name).
		SubResource("status").
		Body(apiGatewayClientCertificate).
		Do().
		Into(result)
	return
}

// Delete takes name of the apiGatewayClientCertificate and deletes it. Returns an error if one occurs.
func (c *apiGatewayClientCertificates) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("apigatewayclientcertificates").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *apiGatewayClientCertificates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("apigatewayclientcertificates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched apiGatewayClientCertificate.
func (c *apiGatewayClientCertificates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiGatewayClientCertificate, err error) {
	result = &v1alpha1.ApiGatewayClientCertificate{}
	err = c.client.Patch(pt).
		Resource("apigatewayclientcertificates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
