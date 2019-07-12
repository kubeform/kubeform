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

// AwsVpcDhcpOptionsesGetter has a method to return a AwsVpcDhcpOptionsInterface.
// A group's client should implement this interface.
type AwsVpcDhcpOptionsesGetter interface {
	AwsVpcDhcpOptionses() AwsVpcDhcpOptionsInterface
}

// AwsVpcDhcpOptionsInterface has methods to work with AwsVpcDhcpOptions resources.
type AwsVpcDhcpOptionsInterface interface {
	Create(*v1alpha1.AwsVpcDhcpOptions) (*v1alpha1.AwsVpcDhcpOptions, error)
	Update(*v1alpha1.AwsVpcDhcpOptions) (*v1alpha1.AwsVpcDhcpOptions, error)
	UpdateStatus(*v1alpha1.AwsVpcDhcpOptions) (*v1alpha1.AwsVpcDhcpOptions, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsVpcDhcpOptions, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsVpcDhcpOptionsList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsVpcDhcpOptions, err error)
	AwsVpcDhcpOptionsExpansion
}

// awsVpcDhcpOptionses implements AwsVpcDhcpOptionsInterface
type awsVpcDhcpOptionses struct {
	client rest.Interface
}

// newAwsVpcDhcpOptionses returns a AwsVpcDhcpOptionses
func newAwsVpcDhcpOptionses(c *AwsV1alpha1Client) *awsVpcDhcpOptionses {
	return &awsVpcDhcpOptionses{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsVpcDhcpOptions, and returns the corresponding awsVpcDhcpOptions object, and an error if there is any.
func (c *awsVpcDhcpOptionses) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsVpcDhcpOptions, err error) {
	result = &v1alpha1.AwsVpcDhcpOptions{}
	err = c.client.Get().
		Resource("awsvpcdhcpoptionses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsVpcDhcpOptionses that match those selectors.
func (c *awsVpcDhcpOptionses) List(opts v1.ListOptions) (result *v1alpha1.AwsVpcDhcpOptionsList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsVpcDhcpOptionsList{}
	err = c.client.Get().
		Resource("awsvpcdhcpoptionses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsVpcDhcpOptionses.
func (c *awsVpcDhcpOptionses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsvpcdhcpoptionses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsVpcDhcpOptions and creates it.  Returns the server's representation of the awsVpcDhcpOptions, and an error, if there is any.
func (c *awsVpcDhcpOptionses) Create(awsVpcDhcpOptions *v1alpha1.AwsVpcDhcpOptions) (result *v1alpha1.AwsVpcDhcpOptions, err error) {
	result = &v1alpha1.AwsVpcDhcpOptions{}
	err = c.client.Post().
		Resource("awsvpcdhcpoptionses").
		Body(awsVpcDhcpOptions).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsVpcDhcpOptions and updates it. Returns the server's representation of the awsVpcDhcpOptions, and an error, if there is any.
func (c *awsVpcDhcpOptionses) Update(awsVpcDhcpOptions *v1alpha1.AwsVpcDhcpOptions) (result *v1alpha1.AwsVpcDhcpOptions, err error) {
	result = &v1alpha1.AwsVpcDhcpOptions{}
	err = c.client.Put().
		Resource("awsvpcdhcpoptionses").
		Name(awsVpcDhcpOptions.Name).
		Body(awsVpcDhcpOptions).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsVpcDhcpOptionses) UpdateStatus(awsVpcDhcpOptions *v1alpha1.AwsVpcDhcpOptions) (result *v1alpha1.AwsVpcDhcpOptions, err error) {
	result = &v1alpha1.AwsVpcDhcpOptions{}
	err = c.client.Put().
		Resource("awsvpcdhcpoptionses").
		Name(awsVpcDhcpOptions.Name).
		SubResource("status").
		Body(awsVpcDhcpOptions).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsVpcDhcpOptions and deletes it. Returns an error if one occurs.
func (c *awsVpcDhcpOptionses) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsvpcdhcpoptionses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsVpcDhcpOptionses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsvpcdhcpoptionses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsVpcDhcpOptions.
func (c *awsVpcDhcpOptionses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsVpcDhcpOptions, err error) {
	result = &v1alpha1.AwsVpcDhcpOptions{}
	err = c.client.Patch(pt).
		Resource("awsvpcdhcpoptionses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
