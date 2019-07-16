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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// StorageSharesGetter has a method to return a StorageShareInterface.
// A group's client should implement this interface.
type StorageSharesGetter interface {
	StorageShares() StorageShareInterface
}

// StorageShareInterface has methods to work with StorageShare resources.
type StorageShareInterface interface {
	Create(*v1alpha1.StorageShare) (*v1alpha1.StorageShare, error)
	Update(*v1alpha1.StorageShare) (*v1alpha1.StorageShare, error)
	UpdateStatus(*v1alpha1.StorageShare) (*v1alpha1.StorageShare, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.StorageShare, error)
	List(opts v1.ListOptions) (*v1alpha1.StorageShareList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StorageShare, err error)
	StorageShareExpansion
}

// storageShares implements StorageShareInterface
type storageShares struct {
	client rest.Interface
}

// newStorageShares returns a StorageShares
func newStorageShares(c *AzurermV1alpha1Client) *storageShares {
	return &storageShares{
		client: c.RESTClient(),
	}
}

// Get takes name of the storageShare, and returns the corresponding storageShare object, and an error if there is any.
func (c *storageShares) Get(name string, options v1.GetOptions) (result *v1alpha1.StorageShare, err error) {
	result = &v1alpha1.StorageShare{}
	err = c.client.Get().
		Resource("storageshares").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of StorageShares that match those selectors.
func (c *storageShares) List(opts v1.ListOptions) (result *v1alpha1.StorageShareList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.StorageShareList{}
	err = c.client.Get().
		Resource("storageshares").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested storageShares.
func (c *storageShares) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("storageshares").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a storageShare and creates it.  Returns the server's representation of the storageShare, and an error, if there is any.
func (c *storageShares) Create(storageShare *v1alpha1.StorageShare) (result *v1alpha1.StorageShare, err error) {
	result = &v1alpha1.StorageShare{}
	err = c.client.Post().
		Resource("storageshares").
		Body(storageShare).
		Do().
		Into(result)
	return
}

// Update takes the representation of a storageShare and updates it. Returns the server's representation of the storageShare, and an error, if there is any.
func (c *storageShares) Update(storageShare *v1alpha1.StorageShare) (result *v1alpha1.StorageShare, err error) {
	result = &v1alpha1.StorageShare{}
	err = c.client.Put().
		Resource("storageshares").
		Name(storageShare.Name).
		Body(storageShare).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *storageShares) UpdateStatus(storageShare *v1alpha1.StorageShare) (result *v1alpha1.StorageShare, err error) {
	result = &v1alpha1.StorageShare{}
	err = c.client.Put().
		Resource("storageshares").
		Name(storageShare.Name).
		SubResource("status").
		Body(storageShare).
		Do().
		Into(result)
	return
}

// Delete takes name of the storageShare and deletes it. Returns an error if one occurs.
func (c *storageShares) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("storageshares").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *storageShares) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("storageshares").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched storageShare.
func (c *storageShares) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StorageShare, err error) {
	result = &v1alpha1.StorageShare{}
	err = c.client.Patch(pt).
		Resource("storageshares").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
