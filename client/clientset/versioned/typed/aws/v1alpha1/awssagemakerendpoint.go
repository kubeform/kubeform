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

// AwsSagemakerEndpointsGetter has a method to return a AwsSagemakerEndpointInterface.
// A group's client should implement this interface.
type AwsSagemakerEndpointsGetter interface {
	AwsSagemakerEndpoints() AwsSagemakerEndpointInterface
}

// AwsSagemakerEndpointInterface has methods to work with AwsSagemakerEndpoint resources.
type AwsSagemakerEndpointInterface interface {
	Create(*v1alpha1.AwsSagemakerEndpoint) (*v1alpha1.AwsSagemakerEndpoint, error)
	Update(*v1alpha1.AwsSagemakerEndpoint) (*v1alpha1.AwsSagemakerEndpoint, error)
	UpdateStatus(*v1alpha1.AwsSagemakerEndpoint) (*v1alpha1.AwsSagemakerEndpoint, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsSagemakerEndpoint, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsSagemakerEndpointList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSagemakerEndpoint, err error)
	AwsSagemakerEndpointExpansion
}

// awsSagemakerEndpoints implements AwsSagemakerEndpointInterface
type awsSagemakerEndpoints struct {
	client rest.Interface
}

// newAwsSagemakerEndpoints returns a AwsSagemakerEndpoints
func newAwsSagemakerEndpoints(c *AwsV1alpha1Client) *awsSagemakerEndpoints {
	return &awsSagemakerEndpoints{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsSagemakerEndpoint, and returns the corresponding awsSagemakerEndpoint object, and an error if there is any.
func (c *awsSagemakerEndpoints) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSagemakerEndpoint, err error) {
	result = &v1alpha1.AwsSagemakerEndpoint{}
	err = c.client.Get().
		Resource("awssagemakerendpoints").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsSagemakerEndpoints that match those selectors.
func (c *awsSagemakerEndpoints) List(opts v1.ListOptions) (result *v1alpha1.AwsSagemakerEndpointList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsSagemakerEndpointList{}
	err = c.client.Get().
		Resource("awssagemakerendpoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsSagemakerEndpoints.
func (c *awsSagemakerEndpoints) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awssagemakerendpoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsSagemakerEndpoint and creates it.  Returns the server's representation of the awsSagemakerEndpoint, and an error, if there is any.
func (c *awsSagemakerEndpoints) Create(awsSagemakerEndpoint *v1alpha1.AwsSagemakerEndpoint) (result *v1alpha1.AwsSagemakerEndpoint, err error) {
	result = &v1alpha1.AwsSagemakerEndpoint{}
	err = c.client.Post().
		Resource("awssagemakerendpoints").
		Body(awsSagemakerEndpoint).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsSagemakerEndpoint and updates it. Returns the server's representation of the awsSagemakerEndpoint, and an error, if there is any.
func (c *awsSagemakerEndpoints) Update(awsSagemakerEndpoint *v1alpha1.AwsSagemakerEndpoint) (result *v1alpha1.AwsSagemakerEndpoint, err error) {
	result = &v1alpha1.AwsSagemakerEndpoint{}
	err = c.client.Put().
		Resource("awssagemakerendpoints").
		Name(awsSagemakerEndpoint.Name).
		Body(awsSagemakerEndpoint).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsSagemakerEndpoints) UpdateStatus(awsSagemakerEndpoint *v1alpha1.AwsSagemakerEndpoint) (result *v1alpha1.AwsSagemakerEndpoint, err error) {
	result = &v1alpha1.AwsSagemakerEndpoint{}
	err = c.client.Put().
		Resource("awssagemakerendpoints").
		Name(awsSagemakerEndpoint.Name).
		SubResource("status").
		Body(awsSagemakerEndpoint).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsSagemakerEndpoint and deletes it. Returns an error if one occurs.
func (c *awsSagemakerEndpoints) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awssagemakerendpoints").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsSagemakerEndpoints) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awssagemakerendpoints").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsSagemakerEndpoint.
func (c *awsSagemakerEndpoints) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSagemakerEndpoint, err error) {
	result = &v1alpha1.AwsSagemakerEndpoint{}
	err = c.client.Patch(pt).
		Resource("awssagemakerendpoints").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
