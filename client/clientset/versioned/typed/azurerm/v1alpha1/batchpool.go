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

// BatchPoolsGetter has a method to return a BatchPoolInterface.
// A group's client should implement this interface.
type BatchPoolsGetter interface {
	BatchPools(namespace string) BatchPoolInterface
}

// BatchPoolInterface has methods to work with BatchPool resources.
type BatchPoolInterface interface {
	Create(*v1alpha1.BatchPool) (*v1alpha1.BatchPool, error)
	Update(*v1alpha1.BatchPool) (*v1alpha1.BatchPool, error)
	UpdateStatus(*v1alpha1.BatchPool) (*v1alpha1.BatchPool, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.BatchPool, error)
	List(opts v1.ListOptions) (*v1alpha1.BatchPoolList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BatchPool, err error)
	BatchPoolExpansion
}

// batchPools implements BatchPoolInterface
type batchPools struct {
	client rest.Interface
	ns     string
}

// newBatchPools returns a BatchPools
func newBatchPools(c *AzurermV1alpha1Client, namespace string) *batchPools {
	return &batchPools{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the batchPool, and returns the corresponding batchPool object, and an error if there is any.
func (c *batchPools) Get(name string, options v1.GetOptions) (result *v1alpha1.BatchPool, err error) {
	result = &v1alpha1.BatchPool{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("batchpools").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BatchPools that match those selectors.
func (c *batchPools) List(opts v1.ListOptions) (result *v1alpha1.BatchPoolList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BatchPoolList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("batchpools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested batchPools.
func (c *batchPools) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("batchpools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a batchPool and creates it.  Returns the server's representation of the batchPool, and an error, if there is any.
func (c *batchPools) Create(batchPool *v1alpha1.BatchPool) (result *v1alpha1.BatchPool, err error) {
	result = &v1alpha1.BatchPool{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("batchpools").
		Body(batchPool).
		Do().
		Into(result)
	return
}

// Update takes the representation of a batchPool and updates it. Returns the server's representation of the batchPool, and an error, if there is any.
func (c *batchPools) Update(batchPool *v1alpha1.BatchPool) (result *v1alpha1.BatchPool, err error) {
	result = &v1alpha1.BatchPool{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("batchpools").
		Name(batchPool.Name).
		Body(batchPool).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *batchPools) UpdateStatus(batchPool *v1alpha1.BatchPool) (result *v1alpha1.BatchPool, err error) {
	result = &v1alpha1.BatchPool{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("batchpools").
		Name(batchPool.Name).
		SubResource("status").
		Body(batchPool).
		Do().
		Into(result)
	return
}

// Delete takes name of the batchPool and deletes it. Returns an error if one occurs.
func (c *batchPools) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("batchpools").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *batchPools) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("batchpools").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched batchPool.
func (c *batchPools) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BatchPool, err error) {
	result = &v1alpha1.BatchPool{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("batchpools").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
