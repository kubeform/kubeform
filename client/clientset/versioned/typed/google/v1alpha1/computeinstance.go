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

// ComputeInstancesGetter has a method to return a ComputeInstanceInterface.
// A group's client should implement this interface.
type ComputeInstancesGetter interface {
	ComputeInstances() ComputeInstanceInterface
}

// ComputeInstanceInterface has methods to work with ComputeInstance resources.
type ComputeInstanceInterface interface {
	Create(*v1alpha1.ComputeInstance) (*v1alpha1.ComputeInstance, error)
	Update(*v1alpha1.ComputeInstance) (*v1alpha1.ComputeInstance, error)
	UpdateStatus(*v1alpha1.ComputeInstance) (*v1alpha1.ComputeInstance, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ComputeInstance, error)
	List(opts v1.ListOptions) (*v1alpha1.ComputeInstanceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeInstance, err error)
	ComputeInstanceExpansion
}

// computeInstances implements ComputeInstanceInterface
type computeInstances struct {
	client rest.Interface
}

// newComputeInstances returns a ComputeInstances
func newComputeInstances(c *GoogleV1alpha1Client) *computeInstances {
	return &computeInstances{
		client: c.RESTClient(),
	}
}

// Get takes name of the computeInstance, and returns the corresponding computeInstance object, and an error if there is any.
func (c *computeInstances) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeInstance, err error) {
	result = &v1alpha1.ComputeInstance{}
	err = c.client.Get().
		Resource("computeinstances").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComputeInstances that match those selectors.
func (c *computeInstances) List(opts v1.ListOptions) (result *v1alpha1.ComputeInstanceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ComputeInstanceList{}
	err = c.client.Get().
		Resource("computeinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested computeInstances.
func (c *computeInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("computeinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a computeInstance and creates it.  Returns the server's representation of the computeInstance, and an error, if there is any.
func (c *computeInstances) Create(computeInstance *v1alpha1.ComputeInstance) (result *v1alpha1.ComputeInstance, err error) {
	result = &v1alpha1.ComputeInstance{}
	err = c.client.Post().
		Resource("computeinstances").
		Body(computeInstance).
		Do().
		Into(result)
	return
}

// Update takes the representation of a computeInstance and updates it. Returns the server's representation of the computeInstance, and an error, if there is any.
func (c *computeInstances) Update(computeInstance *v1alpha1.ComputeInstance) (result *v1alpha1.ComputeInstance, err error) {
	result = &v1alpha1.ComputeInstance{}
	err = c.client.Put().
		Resource("computeinstances").
		Name(computeInstance.Name).
		Body(computeInstance).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *computeInstances) UpdateStatus(computeInstance *v1alpha1.ComputeInstance) (result *v1alpha1.ComputeInstance, err error) {
	result = &v1alpha1.ComputeInstance{}
	err = c.client.Put().
		Resource("computeinstances").
		Name(computeInstance.Name).
		SubResource("status").
		Body(computeInstance).
		Do().
		Into(result)
	return
}

// Delete takes name of the computeInstance and deletes it. Returns an error if one occurs.
func (c *computeInstances) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("computeinstances").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *computeInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("computeinstances").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched computeInstance.
func (c *computeInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeInstance, err error) {
	result = &v1alpha1.ComputeInstance{}
	err = c.client.Patch(pt).
		Resource("computeinstances").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
