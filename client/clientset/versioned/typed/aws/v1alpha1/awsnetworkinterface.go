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

// AwsNetworkInterfacesGetter has a method to return a AwsNetworkInterfaceInterface.
// A group's client should implement this interface.
type AwsNetworkInterfacesGetter interface {
	AwsNetworkInterfaces() AwsNetworkInterfaceInterface
}

// AwsNetworkInterfaceInterface has methods to work with AwsNetworkInterface resources.
type AwsNetworkInterfaceInterface interface {
	Create(*v1alpha1.AwsNetworkInterface) (*v1alpha1.AwsNetworkInterface, error)
	Update(*v1alpha1.AwsNetworkInterface) (*v1alpha1.AwsNetworkInterface, error)
	UpdateStatus(*v1alpha1.AwsNetworkInterface) (*v1alpha1.AwsNetworkInterface, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsNetworkInterface, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsNetworkInterfaceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsNetworkInterface, err error)
	AwsNetworkInterfaceExpansion
}

// awsNetworkInterfaces implements AwsNetworkInterfaceInterface
type awsNetworkInterfaces struct {
	client rest.Interface
}

// newAwsNetworkInterfaces returns a AwsNetworkInterfaces
func newAwsNetworkInterfaces(c *AwsV1alpha1Client) *awsNetworkInterfaces {
	return &awsNetworkInterfaces{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsNetworkInterface, and returns the corresponding awsNetworkInterface object, and an error if there is any.
func (c *awsNetworkInterfaces) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsNetworkInterface, err error) {
	result = &v1alpha1.AwsNetworkInterface{}
	err = c.client.Get().
		Resource("awsnetworkinterfaces").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsNetworkInterfaces that match those selectors.
func (c *awsNetworkInterfaces) List(opts v1.ListOptions) (result *v1alpha1.AwsNetworkInterfaceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsNetworkInterfaceList{}
	err = c.client.Get().
		Resource("awsnetworkinterfaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsNetworkInterfaces.
func (c *awsNetworkInterfaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsnetworkinterfaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsNetworkInterface and creates it.  Returns the server's representation of the awsNetworkInterface, and an error, if there is any.
func (c *awsNetworkInterfaces) Create(awsNetworkInterface *v1alpha1.AwsNetworkInterface) (result *v1alpha1.AwsNetworkInterface, err error) {
	result = &v1alpha1.AwsNetworkInterface{}
	err = c.client.Post().
		Resource("awsnetworkinterfaces").
		Body(awsNetworkInterface).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsNetworkInterface and updates it. Returns the server's representation of the awsNetworkInterface, and an error, if there is any.
func (c *awsNetworkInterfaces) Update(awsNetworkInterface *v1alpha1.AwsNetworkInterface) (result *v1alpha1.AwsNetworkInterface, err error) {
	result = &v1alpha1.AwsNetworkInterface{}
	err = c.client.Put().
		Resource("awsnetworkinterfaces").
		Name(awsNetworkInterface.Name).
		Body(awsNetworkInterface).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsNetworkInterfaces) UpdateStatus(awsNetworkInterface *v1alpha1.AwsNetworkInterface) (result *v1alpha1.AwsNetworkInterface, err error) {
	result = &v1alpha1.AwsNetworkInterface{}
	err = c.client.Put().
		Resource("awsnetworkinterfaces").
		Name(awsNetworkInterface.Name).
		SubResource("status").
		Body(awsNetworkInterface).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsNetworkInterface and deletes it. Returns an error if one occurs.
func (c *awsNetworkInterfaces) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsnetworkinterfaces").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsNetworkInterfaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsnetworkinterfaces").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsNetworkInterface.
func (c *awsNetworkInterfaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsNetworkInterface, err error) {
	result = &v1alpha1.AwsNetworkInterface{}
	err = c.client.Patch(pt).
		Resource("awsnetworkinterfaces").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
