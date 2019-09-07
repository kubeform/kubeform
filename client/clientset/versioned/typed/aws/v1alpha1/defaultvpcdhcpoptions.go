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

// DefaultVpcDHCPOptionsesGetter has a method to return a DefaultVpcDHCPOptionsInterface.
// A group's client should implement this interface.
type DefaultVpcDHCPOptionsesGetter interface {
	DefaultVpcDHCPOptionses(namespace string) DefaultVpcDHCPOptionsInterface
}

// DefaultVpcDHCPOptionsInterface has methods to work with DefaultVpcDHCPOptions resources.
type DefaultVpcDHCPOptionsInterface interface {
	Create(*v1alpha1.DefaultVpcDHCPOptions) (*v1alpha1.DefaultVpcDHCPOptions, error)
	Update(*v1alpha1.DefaultVpcDHCPOptions) (*v1alpha1.DefaultVpcDHCPOptions, error)
	UpdateStatus(*v1alpha1.DefaultVpcDHCPOptions) (*v1alpha1.DefaultVpcDHCPOptions, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DefaultVpcDHCPOptions, error)
	List(opts v1.ListOptions) (*v1alpha1.DefaultVpcDHCPOptionsList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DefaultVpcDHCPOptions, err error)
	DefaultVpcDHCPOptionsExpansion
}

// defaultVpcDHCPOptionses implements DefaultVpcDHCPOptionsInterface
type defaultVpcDHCPOptionses struct {
	client rest.Interface
	ns     string
}

// newDefaultVpcDHCPOptionses returns a DefaultVpcDHCPOptionses
func newDefaultVpcDHCPOptionses(c *AwsV1alpha1Client, namespace string) *defaultVpcDHCPOptionses {
	return &defaultVpcDHCPOptionses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the defaultVpcDHCPOptions, and returns the corresponding defaultVpcDHCPOptions object, and an error if there is any.
func (c *defaultVpcDHCPOptionses) Get(name string, options v1.GetOptions) (result *v1alpha1.DefaultVpcDHCPOptions, err error) {
	result = &v1alpha1.DefaultVpcDHCPOptions{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("defaultvpcdhcpoptionses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DefaultVpcDHCPOptionses that match those selectors.
func (c *defaultVpcDHCPOptionses) List(opts v1.ListOptions) (result *v1alpha1.DefaultVpcDHCPOptionsList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DefaultVpcDHCPOptionsList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("defaultvpcdhcpoptionses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested defaultVpcDHCPOptionses.
func (c *defaultVpcDHCPOptionses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("defaultvpcdhcpoptionses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a defaultVpcDHCPOptions and creates it.  Returns the server's representation of the defaultVpcDHCPOptions, and an error, if there is any.
func (c *defaultVpcDHCPOptionses) Create(defaultVpcDHCPOptions *v1alpha1.DefaultVpcDHCPOptions) (result *v1alpha1.DefaultVpcDHCPOptions, err error) {
	result = &v1alpha1.DefaultVpcDHCPOptions{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("defaultvpcdhcpoptionses").
		Body(defaultVpcDHCPOptions).
		Do().
		Into(result)
	return
}

// Update takes the representation of a defaultVpcDHCPOptions and updates it. Returns the server's representation of the defaultVpcDHCPOptions, and an error, if there is any.
func (c *defaultVpcDHCPOptionses) Update(defaultVpcDHCPOptions *v1alpha1.DefaultVpcDHCPOptions) (result *v1alpha1.DefaultVpcDHCPOptions, err error) {
	result = &v1alpha1.DefaultVpcDHCPOptions{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("defaultvpcdhcpoptionses").
		Name(defaultVpcDHCPOptions.Name).
		Body(defaultVpcDHCPOptions).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *defaultVpcDHCPOptionses) UpdateStatus(defaultVpcDHCPOptions *v1alpha1.DefaultVpcDHCPOptions) (result *v1alpha1.DefaultVpcDHCPOptions, err error) {
	result = &v1alpha1.DefaultVpcDHCPOptions{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("defaultvpcdhcpoptionses").
		Name(defaultVpcDHCPOptions.Name).
		SubResource("status").
		Body(defaultVpcDHCPOptions).
		Do().
		Into(result)
	return
}

// Delete takes name of the defaultVpcDHCPOptions and deletes it. Returns an error if one occurs.
func (c *defaultVpcDHCPOptionses) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("defaultvpcdhcpoptionses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *defaultVpcDHCPOptionses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("defaultvpcdhcpoptionses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched defaultVpcDHCPOptions.
func (c *defaultVpcDHCPOptionses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DefaultVpcDHCPOptions, err error) {
	result = &v1alpha1.DefaultVpcDHCPOptions{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("defaultvpcdhcpoptionses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}