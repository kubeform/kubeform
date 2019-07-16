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

// DefaultNetworkAclsGetter has a method to return a DefaultNetworkAclInterface.
// A group's client should implement this interface.
type DefaultNetworkAclsGetter interface {
	DefaultNetworkAcls() DefaultNetworkAclInterface
}

// DefaultNetworkAclInterface has methods to work with DefaultNetworkAcl resources.
type DefaultNetworkAclInterface interface {
	Create(*v1alpha1.DefaultNetworkAcl) (*v1alpha1.DefaultNetworkAcl, error)
	Update(*v1alpha1.DefaultNetworkAcl) (*v1alpha1.DefaultNetworkAcl, error)
	UpdateStatus(*v1alpha1.DefaultNetworkAcl) (*v1alpha1.DefaultNetworkAcl, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DefaultNetworkAcl, error)
	List(opts v1.ListOptions) (*v1alpha1.DefaultNetworkAclList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DefaultNetworkAcl, err error)
	DefaultNetworkAclExpansion
}

// defaultNetworkAcls implements DefaultNetworkAclInterface
type defaultNetworkAcls struct {
	client rest.Interface
}

// newDefaultNetworkAcls returns a DefaultNetworkAcls
func newDefaultNetworkAcls(c *AwsV1alpha1Client) *defaultNetworkAcls {
	return &defaultNetworkAcls{
		client: c.RESTClient(),
	}
}

// Get takes name of the defaultNetworkAcl, and returns the corresponding defaultNetworkAcl object, and an error if there is any.
func (c *defaultNetworkAcls) Get(name string, options v1.GetOptions) (result *v1alpha1.DefaultNetworkAcl, err error) {
	result = &v1alpha1.DefaultNetworkAcl{}
	err = c.client.Get().
		Resource("defaultnetworkacls").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DefaultNetworkAcls that match those selectors.
func (c *defaultNetworkAcls) List(opts v1.ListOptions) (result *v1alpha1.DefaultNetworkAclList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DefaultNetworkAclList{}
	err = c.client.Get().
		Resource("defaultnetworkacls").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested defaultNetworkAcls.
func (c *defaultNetworkAcls) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("defaultnetworkacls").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a defaultNetworkAcl and creates it.  Returns the server's representation of the defaultNetworkAcl, and an error, if there is any.
func (c *defaultNetworkAcls) Create(defaultNetworkAcl *v1alpha1.DefaultNetworkAcl) (result *v1alpha1.DefaultNetworkAcl, err error) {
	result = &v1alpha1.DefaultNetworkAcl{}
	err = c.client.Post().
		Resource("defaultnetworkacls").
		Body(defaultNetworkAcl).
		Do().
		Into(result)
	return
}

// Update takes the representation of a defaultNetworkAcl and updates it. Returns the server's representation of the defaultNetworkAcl, and an error, if there is any.
func (c *defaultNetworkAcls) Update(defaultNetworkAcl *v1alpha1.DefaultNetworkAcl) (result *v1alpha1.DefaultNetworkAcl, err error) {
	result = &v1alpha1.DefaultNetworkAcl{}
	err = c.client.Put().
		Resource("defaultnetworkacls").
		Name(defaultNetworkAcl.Name).
		Body(defaultNetworkAcl).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *defaultNetworkAcls) UpdateStatus(defaultNetworkAcl *v1alpha1.DefaultNetworkAcl) (result *v1alpha1.DefaultNetworkAcl, err error) {
	result = &v1alpha1.DefaultNetworkAcl{}
	err = c.client.Put().
		Resource("defaultnetworkacls").
		Name(defaultNetworkAcl.Name).
		SubResource("status").
		Body(defaultNetworkAcl).
		Do().
		Into(result)
	return
}

// Delete takes name of the defaultNetworkAcl and deletes it. Returns an error if one occurs.
func (c *defaultNetworkAcls) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("defaultnetworkacls").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *defaultNetworkAcls) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("defaultnetworkacls").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched defaultNetworkAcl.
func (c *defaultNetworkAcls) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DefaultNetworkAcl, err error) {
	result = &v1alpha1.DefaultNetworkAcl{}
	err = c.client.Patch(pt).
		Resource("defaultnetworkacls").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
