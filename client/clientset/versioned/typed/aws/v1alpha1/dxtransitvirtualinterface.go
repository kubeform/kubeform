/*
Copyright The Kubeform Authors.

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

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DxTransitVirtualInterfacesGetter has a method to return a DxTransitVirtualInterfaceInterface.
// A group's client should implement this interface.
type DxTransitVirtualInterfacesGetter interface {
	DxTransitVirtualInterfaces(namespace string) DxTransitVirtualInterfaceInterface
}

// DxTransitVirtualInterfaceInterface has methods to work with DxTransitVirtualInterface resources.
type DxTransitVirtualInterfaceInterface interface {
	Create(*v1alpha1.DxTransitVirtualInterface) (*v1alpha1.DxTransitVirtualInterface, error)
	Update(*v1alpha1.DxTransitVirtualInterface) (*v1alpha1.DxTransitVirtualInterface, error)
	UpdateStatus(*v1alpha1.DxTransitVirtualInterface) (*v1alpha1.DxTransitVirtualInterface, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DxTransitVirtualInterface, error)
	List(opts v1.ListOptions) (*v1alpha1.DxTransitVirtualInterfaceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DxTransitVirtualInterface, err error)
	DxTransitVirtualInterfaceExpansion
}

// dxTransitVirtualInterfaces implements DxTransitVirtualInterfaceInterface
type dxTransitVirtualInterfaces struct {
	client rest.Interface
	ns     string
}

// newDxTransitVirtualInterfaces returns a DxTransitVirtualInterfaces
func newDxTransitVirtualInterfaces(c *AwsV1alpha1Client, namespace string) *dxTransitVirtualInterfaces {
	return &dxTransitVirtualInterfaces{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dxTransitVirtualInterface, and returns the corresponding dxTransitVirtualInterface object, and an error if there is any.
func (c *dxTransitVirtualInterfaces) Get(name string, options v1.GetOptions) (result *v1alpha1.DxTransitVirtualInterface, err error) {
	result = &v1alpha1.DxTransitVirtualInterface{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dxtransitvirtualinterfaces").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DxTransitVirtualInterfaces that match those selectors.
func (c *dxTransitVirtualInterfaces) List(opts v1.ListOptions) (result *v1alpha1.DxTransitVirtualInterfaceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DxTransitVirtualInterfaceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dxtransitvirtualinterfaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dxTransitVirtualInterfaces.
func (c *dxTransitVirtualInterfaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dxtransitvirtualinterfaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a dxTransitVirtualInterface and creates it.  Returns the server's representation of the dxTransitVirtualInterface, and an error, if there is any.
func (c *dxTransitVirtualInterfaces) Create(dxTransitVirtualInterface *v1alpha1.DxTransitVirtualInterface) (result *v1alpha1.DxTransitVirtualInterface, err error) {
	result = &v1alpha1.DxTransitVirtualInterface{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dxtransitvirtualinterfaces").
		Body(dxTransitVirtualInterface).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dxTransitVirtualInterface and updates it. Returns the server's representation of the dxTransitVirtualInterface, and an error, if there is any.
func (c *dxTransitVirtualInterfaces) Update(dxTransitVirtualInterface *v1alpha1.DxTransitVirtualInterface) (result *v1alpha1.DxTransitVirtualInterface, err error) {
	result = &v1alpha1.DxTransitVirtualInterface{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dxtransitvirtualinterfaces").
		Name(dxTransitVirtualInterface.Name).
		Body(dxTransitVirtualInterface).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *dxTransitVirtualInterfaces) UpdateStatus(dxTransitVirtualInterface *v1alpha1.DxTransitVirtualInterface) (result *v1alpha1.DxTransitVirtualInterface, err error) {
	result = &v1alpha1.DxTransitVirtualInterface{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dxtransitvirtualinterfaces").
		Name(dxTransitVirtualInterface.Name).
		SubResource("status").
		Body(dxTransitVirtualInterface).
		Do().
		Into(result)
	return
}

// Delete takes name of the dxTransitVirtualInterface and deletes it. Returns an error if one occurs.
func (c *dxTransitVirtualInterfaces) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dxtransitvirtualinterfaces").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dxTransitVirtualInterfaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dxtransitvirtualinterfaces").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dxTransitVirtualInterface.
func (c *dxTransitVirtualInterfaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DxTransitVirtualInterface, err error) {
	result = &v1alpha1.DxTransitVirtualInterface{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dxtransitvirtualinterfaces").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
