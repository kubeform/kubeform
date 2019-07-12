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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// AzurermDataLakeStoresGetter has a method to return a AzurermDataLakeStoreInterface.
// A group's client should implement this interface.
type AzurermDataLakeStoresGetter interface {
	AzurermDataLakeStores() AzurermDataLakeStoreInterface
}

// AzurermDataLakeStoreInterface has methods to work with AzurermDataLakeStore resources.
type AzurermDataLakeStoreInterface interface {
	Create(*v1alpha1.AzurermDataLakeStore) (*v1alpha1.AzurermDataLakeStore, error)
	Update(*v1alpha1.AzurermDataLakeStore) (*v1alpha1.AzurermDataLakeStore, error)
	UpdateStatus(*v1alpha1.AzurermDataLakeStore) (*v1alpha1.AzurermDataLakeStore, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermDataLakeStore, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermDataLakeStoreList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermDataLakeStore, err error)
	AzurermDataLakeStoreExpansion
}

// azurermDataLakeStores implements AzurermDataLakeStoreInterface
type azurermDataLakeStores struct {
	client rest.Interface
}

// newAzurermDataLakeStores returns a AzurermDataLakeStores
func newAzurermDataLakeStores(c *AzurermV1alpha1Client) *azurermDataLakeStores {
	return &azurermDataLakeStores{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermDataLakeStore, and returns the corresponding azurermDataLakeStore object, and an error if there is any.
func (c *azurermDataLakeStores) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermDataLakeStore, err error) {
	result = &v1alpha1.AzurermDataLakeStore{}
	err = c.client.Get().
		Resource("azurermdatalakestores").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermDataLakeStores that match those selectors.
func (c *azurermDataLakeStores) List(opts v1.ListOptions) (result *v1alpha1.AzurermDataLakeStoreList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermDataLakeStoreList{}
	err = c.client.Get().
		Resource("azurermdatalakestores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermDataLakeStores.
func (c *azurermDataLakeStores) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermdatalakestores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermDataLakeStore and creates it.  Returns the server's representation of the azurermDataLakeStore, and an error, if there is any.
func (c *azurermDataLakeStores) Create(azurermDataLakeStore *v1alpha1.AzurermDataLakeStore) (result *v1alpha1.AzurermDataLakeStore, err error) {
	result = &v1alpha1.AzurermDataLakeStore{}
	err = c.client.Post().
		Resource("azurermdatalakestores").
		Body(azurermDataLakeStore).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermDataLakeStore and updates it. Returns the server's representation of the azurermDataLakeStore, and an error, if there is any.
func (c *azurermDataLakeStores) Update(azurermDataLakeStore *v1alpha1.AzurermDataLakeStore) (result *v1alpha1.AzurermDataLakeStore, err error) {
	result = &v1alpha1.AzurermDataLakeStore{}
	err = c.client.Put().
		Resource("azurermdatalakestores").
		Name(azurermDataLakeStore.Name).
		Body(azurermDataLakeStore).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermDataLakeStores) UpdateStatus(azurermDataLakeStore *v1alpha1.AzurermDataLakeStore) (result *v1alpha1.AzurermDataLakeStore, err error) {
	result = &v1alpha1.AzurermDataLakeStore{}
	err = c.client.Put().
		Resource("azurermdatalakestores").
		Name(azurermDataLakeStore.Name).
		SubResource("status").
		Body(azurermDataLakeStore).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermDataLakeStore and deletes it. Returns an error if one occurs.
func (c *azurermDataLakeStores) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermdatalakestores").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermDataLakeStores) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermdatalakestores").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermDataLakeStore.
func (c *azurermDataLakeStores) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermDataLakeStore, err error) {
	result = &v1alpha1.AzurermDataLakeStore{}
	err = c.client.Patch(pt).
		Resource("azurermdatalakestores").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
