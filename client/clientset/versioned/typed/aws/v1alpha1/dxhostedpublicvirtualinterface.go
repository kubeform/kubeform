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

// DxHostedPublicVirtualInterfacesGetter has a method to return a DxHostedPublicVirtualInterfaceInterface.
// A group's client should implement this interface.
type DxHostedPublicVirtualInterfacesGetter interface {
	DxHostedPublicVirtualInterfaces(namespace string) DxHostedPublicVirtualInterfaceInterface
}

// DxHostedPublicVirtualInterfaceInterface has methods to work with DxHostedPublicVirtualInterface resources.
type DxHostedPublicVirtualInterfaceInterface interface {
	Create(*v1alpha1.DxHostedPublicVirtualInterface) (*v1alpha1.DxHostedPublicVirtualInterface, error)
	Update(*v1alpha1.DxHostedPublicVirtualInterface) (*v1alpha1.DxHostedPublicVirtualInterface, error)
	UpdateStatus(*v1alpha1.DxHostedPublicVirtualInterface) (*v1alpha1.DxHostedPublicVirtualInterface, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DxHostedPublicVirtualInterface, error)
	List(opts v1.ListOptions) (*v1alpha1.DxHostedPublicVirtualInterfaceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DxHostedPublicVirtualInterface, err error)
	DxHostedPublicVirtualInterfaceExpansion
}

// dxHostedPublicVirtualInterfaces implements DxHostedPublicVirtualInterfaceInterface
type dxHostedPublicVirtualInterfaces struct {
	client rest.Interface
	ns     string
}

// newDxHostedPublicVirtualInterfaces returns a DxHostedPublicVirtualInterfaces
func newDxHostedPublicVirtualInterfaces(c *AwsV1alpha1Client, namespace string) *dxHostedPublicVirtualInterfaces {
	return &dxHostedPublicVirtualInterfaces{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dxHostedPublicVirtualInterface, and returns the corresponding dxHostedPublicVirtualInterface object, and an error if there is any.
func (c *dxHostedPublicVirtualInterfaces) Get(name string, options v1.GetOptions) (result *v1alpha1.DxHostedPublicVirtualInterface, err error) {
	result = &v1alpha1.DxHostedPublicVirtualInterface{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dxhostedpublicvirtualinterfaces").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DxHostedPublicVirtualInterfaces that match those selectors.
func (c *dxHostedPublicVirtualInterfaces) List(opts v1.ListOptions) (result *v1alpha1.DxHostedPublicVirtualInterfaceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DxHostedPublicVirtualInterfaceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dxhostedpublicvirtualinterfaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dxHostedPublicVirtualInterfaces.
func (c *dxHostedPublicVirtualInterfaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dxhostedpublicvirtualinterfaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a dxHostedPublicVirtualInterface and creates it.  Returns the server's representation of the dxHostedPublicVirtualInterface, and an error, if there is any.
func (c *dxHostedPublicVirtualInterfaces) Create(dxHostedPublicVirtualInterface *v1alpha1.DxHostedPublicVirtualInterface) (result *v1alpha1.DxHostedPublicVirtualInterface, err error) {
	result = &v1alpha1.DxHostedPublicVirtualInterface{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dxhostedpublicvirtualinterfaces").
		Body(dxHostedPublicVirtualInterface).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dxHostedPublicVirtualInterface and updates it. Returns the server's representation of the dxHostedPublicVirtualInterface, and an error, if there is any.
func (c *dxHostedPublicVirtualInterfaces) Update(dxHostedPublicVirtualInterface *v1alpha1.DxHostedPublicVirtualInterface) (result *v1alpha1.DxHostedPublicVirtualInterface, err error) {
	result = &v1alpha1.DxHostedPublicVirtualInterface{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dxhostedpublicvirtualinterfaces").
		Name(dxHostedPublicVirtualInterface.Name).
		Body(dxHostedPublicVirtualInterface).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *dxHostedPublicVirtualInterfaces) UpdateStatus(dxHostedPublicVirtualInterface *v1alpha1.DxHostedPublicVirtualInterface) (result *v1alpha1.DxHostedPublicVirtualInterface, err error) {
	result = &v1alpha1.DxHostedPublicVirtualInterface{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dxhostedpublicvirtualinterfaces").
		Name(dxHostedPublicVirtualInterface.Name).
		SubResource("status").
		Body(dxHostedPublicVirtualInterface).
		Do().
		Into(result)
	return
}

// Delete takes name of the dxHostedPublicVirtualInterface and deletes it. Returns an error if one occurs.
func (c *dxHostedPublicVirtualInterfaces) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dxhostedpublicvirtualinterfaces").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dxHostedPublicVirtualInterfaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dxhostedpublicvirtualinterfaces").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dxHostedPublicVirtualInterface.
func (c *dxHostedPublicVirtualInterfaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DxHostedPublicVirtualInterface, err error) {
	result = &v1alpha1.DxHostedPublicVirtualInterface{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dxhostedpublicvirtualinterfaces").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
