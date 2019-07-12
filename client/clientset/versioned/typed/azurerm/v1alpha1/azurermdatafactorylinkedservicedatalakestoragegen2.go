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

// AzurermDataFactoryLinkedServiceDataLakeStorageGen2sGetter has a method to return a AzurermDataFactoryLinkedServiceDataLakeStorageGen2Interface.
// A group's client should implement this interface.
type AzurermDataFactoryLinkedServiceDataLakeStorageGen2sGetter interface {
	AzurermDataFactoryLinkedServiceDataLakeStorageGen2s() AzurermDataFactoryLinkedServiceDataLakeStorageGen2Interface
}

// AzurermDataFactoryLinkedServiceDataLakeStorageGen2Interface has methods to work with AzurermDataFactoryLinkedServiceDataLakeStorageGen2 resources.
type AzurermDataFactoryLinkedServiceDataLakeStorageGen2Interface interface {
	Create(*v1alpha1.AzurermDataFactoryLinkedServiceDataLakeStorageGen2) (*v1alpha1.AzurermDataFactoryLinkedServiceDataLakeStorageGen2, error)
	Update(*v1alpha1.AzurermDataFactoryLinkedServiceDataLakeStorageGen2) (*v1alpha1.AzurermDataFactoryLinkedServiceDataLakeStorageGen2, error)
	UpdateStatus(*v1alpha1.AzurermDataFactoryLinkedServiceDataLakeStorageGen2) (*v1alpha1.AzurermDataFactoryLinkedServiceDataLakeStorageGen2, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermDataFactoryLinkedServiceDataLakeStorageGen2, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermDataFactoryLinkedServiceDataLakeStorageGen2List, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermDataFactoryLinkedServiceDataLakeStorageGen2, err error)
	AzurermDataFactoryLinkedServiceDataLakeStorageGen2Expansion
}

// azurermDataFactoryLinkedServiceDataLakeStorageGen2s implements AzurermDataFactoryLinkedServiceDataLakeStorageGen2Interface
type azurermDataFactoryLinkedServiceDataLakeStorageGen2s struct {
	client rest.Interface
}

// newAzurermDataFactoryLinkedServiceDataLakeStorageGen2s returns a AzurermDataFactoryLinkedServiceDataLakeStorageGen2s
func newAzurermDataFactoryLinkedServiceDataLakeStorageGen2s(c *AzurermV1alpha1Client) *azurermDataFactoryLinkedServiceDataLakeStorageGen2s {
	return &azurermDataFactoryLinkedServiceDataLakeStorageGen2s{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermDataFactoryLinkedServiceDataLakeStorageGen2, and returns the corresponding azurermDataFactoryLinkedServiceDataLakeStorageGen2 object, and an error if there is any.
func (c *azurermDataFactoryLinkedServiceDataLakeStorageGen2s) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermDataFactoryLinkedServiceDataLakeStorageGen2, err error) {
	result = &v1alpha1.AzurermDataFactoryLinkedServiceDataLakeStorageGen2{}
	err = c.client.Get().
		Resource("azurermdatafactorylinkedservicedatalakestoragegen2s").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermDataFactoryLinkedServiceDataLakeStorageGen2s that match those selectors.
func (c *azurermDataFactoryLinkedServiceDataLakeStorageGen2s) List(opts v1.ListOptions) (result *v1alpha1.AzurermDataFactoryLinkedServiceDataLakeStorageGen2List, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermDataFactoryLinkedServiceDataLakeStorageGen2List{}
	err = c.client.Get().
		Resource("azurermdatafactorylinkedservicedatalakestoragegen2s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermDataFactoryLinkedServiceDataLakeStorageGen2s.
func (c *azurermDataFactoryLinkedServiceDataLakeStorageGen2s) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermdatafactorylinkedservicedatalakestoragegen2s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermDataFactoryLinkedServiceDataLakeStorageGen2 and creates it.  Returns the server's representation of the azurermDataFactoryLinkedServiceDataLakeStorageGen2, and an error, if there is any.
func (c *azurermDataFactoryLinkedServiceDataLakeStorageGen2s) Create(azurermDataFactoryLinkedServiceDataLakeStorageGen2 *v1alpha1.AzurermDataFactoryLinkedServiceDataLakeStorageGen2) (result *v1alpha1.AzurermDataFactoryLinkedServiceDataLakeStorageGen2, err error) {
	result = &v1alpha1.AzurermDataFactoryLinkedServiceDataLakeStorageGen2{}
	err = c.client.Post().
		Resource("azurermdatafactorylinkedservicedatalakestoragegen2s").
		Body(azurermDataFactoryLinkedServiceDataLakeStorageGen2).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermDataFactoryLinkedServiceDataLakeStorageGen2 and updates it. Returns the server's representation of the azurermDataFactoryLinkedServiceDataLakeStorageGen2, and an error, if there is any.
func (c *azurermDataFactoryLinkedServiceDataLakeStorageGen2s) Update(azurermDataFactoryLinkedServiceDataLakeStorageGen2 *v1alpha1.AzurermDataFactoryLinkedServiceDataLakeStorageGen2) (result *v1alpha1.AzurermDataFactoryLinkedServiceDataLakeStorageGen2, err error) {
	result = &v1alpha1.AzurermDataFactoryLinkedServiceDataLakeStorageGen2{}
	err = c.client.Put().
		Resource("azurermdatafactorylinkedservicedatalakestoragegen2s").
		Name(azurermDataFactoryLinkedServiceDataLakeStorageGen2.Name).
		Body(azurermDataFactoryLinkedServiceDataLakeStorageGen2).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermDataFactoryLinkedServiceDataLakeStorageGen2s) UpdateStatus(azurermDataFactoryLinkedServiceDataLakeStorageGen2 *v1alpha1.AzurermDataFactoryLinkedServiceDataLakeStorageGen2) (result *v1alpha1.AzurermDataFactoryLinkedServiceDataLakeStorageGen2, err error) {
	result = &v1alpha1.AzurermDataFactoryLinkedServiceDataLakeStorageGen2{}
	err = c.client.Put().
		Resource("azurermdatafactorylinkedservicedatalakestoragegen2s").
		Name(azurermDataFactoryLinkedServiceDataLakeStorageGen2.Name).
		SubResource("status").
		Body(azurermDataFactoryLinkedServiceDataLakeStorageGen2).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermDataFactoryLinkedServiceDataLakeStorageGen2 and deletes it. Returns an error if one occurs.
func (c *azurermDataFactoryLinkedServiceDataLakeStorageGen2s) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermdatafactorylinkedservicedatalakestoragegen2s").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermDataFactoryLinkedServiceDataLakeStorageGen2s) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermdatafactorylinkedservicedatalakestoragegen2s").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermDataFactoryLinkedServiceDataLakeStorageGen2.
func (c *azurermDataFactoryLinkedServiceDataLakeStorageGen2s) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermDataFactoryLinkedServiceDataLakeStorageGen2, err error) {
	result = &v1alpha1.AzurermDataFactoryLinkedServiceDataLakeStorageGen2{}
	err = c.client.Patch(pt).
		Resource("azurermdatafactorylinkedservicedatalakestoragegen2s").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
