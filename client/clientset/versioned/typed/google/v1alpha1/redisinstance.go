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

// RedisInstancesGetter has a method to return a RedisInstanceInterface.
// A group's client should implement this interface.
type RedisInstancesGetter interface {
	RedisInstances() RedisInstanceInterface
}

// RedisInstanceInterface has methods to work with RedisInstance resources.
type RedisInstanceInterface interface {
	Create(*v1alpha1.RedisInstance) (*v1alpha1.RedisInstance, error)
	Update(*v1alpha1.RedisInstance) (*v1alpha1.RedisInstance, error)
	UpdateStatus(*v1alpha1.RedisInstance) (*v1alpha1.RedisInstance, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.RedisInstance, error)
	List(opts v1.ListOptions) (*v1alpha1.RedisInstanceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RedisInstance, err error)
	RedisInstanceExpansion
}

// redisInstances implements RedisInstanceInterface
type redisInstances struct {
	client rest.Interface
}

// newRedisInstances returns a RedisInstances
func newRedisInstances(c *GoogleV1alpha1Client) *redisInstances {
	return &redisInstances{
		client: c.RESTClient(),
	}
}

// Get takes name of the redisInstance, and returns the corresponding redisInstance object, and an error if there is any.
func (c *redisInstances) Get(name string, options v1.GetOptions) (result *v1alpha1.RedisInstance, err error) {
	result = &v1alpha1.RedisInstance{}
	err = c.client.Get().
		Resource("redisinstances").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RedisInstances that match those selectors.
func (c *redisInstances) List(opts v1.ListOptions) (result *v1alpha1.RedisInstanceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.RedisInstanceList{}
	err = c.client.Get().
		Resource("redisinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested redisInstances.
func (c *redisInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("redisinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a redisInstance and creates it.  Returns the server's representation of the redisInstance, and an error, if there is any.
func (c *redisInstances) Create(redisInstance *v1alpha1.RedisInstance) (result *v1alpha1.RedisInstance, err error) {
	result = &v1alpha1.RedisInstance{}
	err = c.client.Post().
		Resource("redisinstances").
		Body(redisInstance).
		Do().
		Into(result)
	return
}

// Update takes the representation of a redisInstance and updates it. Returns the server's representation of the redisInstance, and an error, if there is any.
func (c *redisInstances) Update(redisInstance *v1alpha1.RedisInstance) (result *v1alpha1.RedisInstance, err error) {
	result = &v1alpha1.RedisInstance{}
	err = c.client.Put().
		Resource("redisinstances").
		Name(redisInstance.Name).
		Body(redisInstance).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *redisInstances) UpdateStatus(redisInstance *v1alpha1.RedisInstance) (result *v1alpha1.RedisInstance, err error) {
	result = &v1alpha1.RedisInstance{}
	err = c.client.Put().
		Resource("redisinstances").
		Name(redisInstance.Name).
		SubResource("status").
		Body(redisInstance).
		Do().
		Into(result)
	return
}

// Delete takes name of the redisInstance and deletes it. Returns an error if one occurs.
func (c *redisInstances) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("redisinstances").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *redisInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("redisinstances").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched redisInstance.
func (c *redisInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RedisInstance, err error) {
	result = &v1alpha1.RedisInstance{}
	err = c.client.Patch(pt).
		Resource("redisinstances").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
