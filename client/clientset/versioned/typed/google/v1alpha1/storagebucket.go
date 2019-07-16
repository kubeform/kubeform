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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// StorageBucketsGetter has a method to return a StorageBucketInterface.
// A group's client should implement this interface.
type StorageBucketsGetter interface {
	StorageBuckets() StorageBucketInterface
}

// StorageBucketInterface has methods to work with StorageBucket resources.
type StorageBucketInterface interface {
	Create(*v1alpha1.StorageBucket) (*v1alpha1.StorageBucket, error)
	Update(*v1alpha1.StorageBucket) (*v1alpha1.StorageBucket, error)
	UpdateStatus(*v1alpha1.StorageBucket) (*v1alpha1.StorageBucket, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.StorageBucket, error)
	List(opts v1.ListOptions) (*v1alpha1.StorageBucketList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StorageBucket, err error)
	StorageBucketExpansion
}

// storageBuckets implements StorageBucketInterface
type storageBuckets struct {
	client rest.Interface
}

// newStorageBuckets returns a StorageBuckets
func newStorageBuckets(c *GoogleV1alpha1Client) *storageBuckets {
	return &storageBuckets{
		client: c.RESTClient(),
	}
}

// Get takes name of the storageBucket, and returns the corresponding storageBucket object, and an error if there is any.
func (c *storageBuckets) Get(name string, options v1.GetOptions) (result *v1alpha1.StorageBucket, err error) {
	result = &v1alpha1.StorageBucket{}
	err = c.client.Get().
		Resource("storagebuckets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of StorageBuckets that match those selectors.
func (c *storageBuckets) List(opts v1.ListOptions) (result *v1alpha1.StorageBucketList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.StorageBucketList{}
	err = c.client.Get().
		Resource("storagebuckets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested storageBuckets.
func (c *storageBuckets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("storagebuckets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a storageBucket and creates it.  Returns the server's representation of the storageBucket, and an error, if there is any.
func (c *storageBuckets) Create(storageBucket *v1alpha1.StorageBucket) (result *v1alpha1.StorageBucket, err error) {
	result = &v1alpha1.StorageBucket{}
	err = c.client.Post().
		Resource("storagebuckets").
		Body(storageBucket).
		Do().
		Into(result)
	return
}

// Update takes the representation of a storageBucket and updates it. Returns the server's representation of the storageBucket, and an error, if there is any.
func (c *storageBuckets) Update(storageBucket *v1alpha1.StorageBucket) (result *v1alpha1.StorageBucket, err error) {
	result = &v1alpha1.StorageBucket{}
	err = c.client.Put().
		Resource("storagebuckets").
		Name(storageBucket.Name).
		Body(storageBucket).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *storageBuckets) UpdateStatus(storageBucket *v1alpha1.StorageBucket) (result *v1alpha1.StorageBucket, err error) {
	result = &v1alpha1.StorageBucket{}
	err = c.client.Put().
		Resource("storagebuckets").
		Name(storageBucket.Name).
		SubResource("status").
		Body(storageBucket).
		Do().
		Into(result)
	return
}

// Delete takes name of the storageBucket and deletes it. Returns an error if one occurs.
func (c *storageBuckets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("storagebuckets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *storageBuckets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("storagebuckets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched storageBucket.
func (c *storageBuckets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StorageBucket, err error) {
	result = &v1alpha1.StorageBucket{}
	err = c.client.Patch(pt).
		Resource("storagebuckets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
