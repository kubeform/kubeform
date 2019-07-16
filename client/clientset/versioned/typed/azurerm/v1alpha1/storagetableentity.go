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

// StorageTableEntitiesGetter has a method to return a StorageTableEntityInterface.
// A group's client should implement this interface.
type StorageTableEntitiesGetter interface {
	StorageTableEntities() StorageTableEntityInterface
}

// StorageTableEntityInterface has methods to work with StorageTableEntity resources.
type StorageTableEntityInterface interface {
	Create(*v1alpha1.StorageTableEntity) (*v1alpha1.StorageTableEntity, error)
	Update(*v1alpha1.StorageTableEntity) (*v1alpha1.StorageTableEntity, error)
	UpdateStatus(*v1alpha1.StorageTableEntity) (*v1alpha1.StorageTableEntity, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.StorageTableEntity, error)
	List(opts v1.ListOptions) (*v1alpha1.StorageTableEntityList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StorageTableEntity, err error)
	StorageTableEntityExpansion
}

// storageTableEntities implements StorageTableEntityInterface
type storageTableEntities struct {
	client rest.Interface
}

// newStorageTableEntities returns a StorageTableEntities
func newStorageTableEntities(c *AzurermV1alpha1Client) *storageTableEntities {
	return &storageTableEntities{
		client: c.RESTClient(),
	}
}

// Get takes name of the storageTableEntity, and returns the corresponding storageTableEntity object, and an error if there is any.
func (c *storageTableEntities) Get(name string, options v1.GetOptions) (result *v1alpha1.StorageTableEntity, err error) {
	result = &v1alpha1.StorageTableEntity{}
	err = c.client.Get().
		Resource("storagetableentities").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of StorageTableEntities that match those selectors.
func (c *storageTableEntities) List(opts v1.ListOptions) (result *v1alpha1.StorageTableEntityList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.StorageTableEntityList{}
	err = c.client.Get().
		Resource("storagetableentities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested storageTableEntities.
func (c *storageTableEntities) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("storagetableentities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a storageTableEntity and creates it.  Returns the server's representation of the storageTableEntity, and an error, if there is any.
func (c *storageTableEntities) Create(storageTableEntity *v1alpha1.StorageTableEntity) (result *v1alpha1.StorageTableEntity, err error) {
	result = &v1alpha1.StorageTableEntity{}
	err = c.client.Post().
		Resource("storagetableentities").
		Body(storageTableEntity).
		Do().
		Into(result)
	return
}

// Update takes the representation of a storageTableEntity and updates it. Returns the server's representation of the storageTableEntity, and an error, if there is any.
func (c *storageTableEntities) Update(storageTableEntity *v1alpha1.StorageTableEntity) (result *v1alpha1.StorageTableEntity, err error) {
	result = &v1alpha1.StorageTableEntity{}
	err = c.client.Put().
		Resource("storagetableentities").
		Name(storageTableEntity.Name).
		Body(storageTableEntity).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *storageTableEntities) UpdateStatus(storageTableEntity *v1alpha1.StorageTableEntity) (result *v1alpha1.StorageTableEntity, err error) {
	result = &v1alpha1.StorageTableEntity{}
	err = c.client.Put().
		Resource("storagetableentities").
		Name(storageTableEntity.Name).
		SubResource("status").
		Body(storageTableEntity).
		Do().
		Into(result)
	return
}

// Delete takes name of the storageTableEntity and deletes it. Returns an error if one occurs.
func (c *storageTableEntities) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("storagetableentities").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *storageTableEntities) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("storagetableentities").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched storageTableEntity.
func (c *storageTableEntities) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StorageTableEntity, err error) {
	result = &v1alpha1.StorageTableEntity{}
	err = c.client.Patch(pt).
		Resource("storagetableentities").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
