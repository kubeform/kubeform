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

// GoogleComputeRegionAutoscalersGetter has a method to return a GoogleComputeRegionAutoscalerInterface.
// A group's client should implement this interface.
type GoogleComputeRegionAutoscalersGetter interface {
	GoogleComputeRegionAutoscalers() GoogleComputeRegionAutoscalerInterface
}

// GoogleComputeRegionAutoscalerInterface has methods to work with GoogleComputeRegionAutoscaler resources.
type GoogleComputeRegionAutoscalerInterface interface {
	Create(*v1alpha1.GoogleComputeRegionAutoscaler) (*v1alpha1.GoogleComputeRegionAutoscaler, error)
	Update(*v1alpha1.GoogleComputeRegionAutoscaler) (*v1alpha1.GoogleComputeRegionAutoscaler, error)
	UpdateStatus(*v1alpha1.GoogleComputeRegionAutoscaler) (*v1alpha1.GoogleComputeRegionAutoscaler, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GoogleComputeRegionAutoscaler, error)
	List(opts v1.ListOptions) (*v1alpha1.GoogleComputeRegionAutoscalerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleComputeRegionAutoscaler, err error)
	GoogleComputeRegionAutoscalerExpansion
}

// googleComputeRegionAutoscalers implements GoogleComputeRegionAutoscalerInterface
type googleComputeRegionAutoscalers struct {
	client rest.Interface
}

// newGoogleComputeRegionAutoscalers returns a GoogleComputeRegionAutoscalers
func newGoogleComputeRegionAutoscalers(c *GoogleV1alpha1Client) *googleComputeRegionAutoscalers {
	return &googleComputeRegionAutoscalers{
		client: c.RESTClient(),
	}
}

// Get takes name of the googleComputeRegionAutoscaler, and returns the corresponding googleComputeRegionAutoscaler object, and an error if there is any.
func (c *googleComputeRegionAutoscalers) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleComputeRegionAutoscaler, err error) {
	result = &v1alpha1.GoogleComputeRegionAutoscaler{}
	err = c.client.Get().
		Resource("googlecomputeregionautoscalers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GoogleComputeRegionAutoscalers that match those selectors.
func (c *googleComputeRegionAutoscalers) List(opts v1.ListOptions) (result *v1alpha1.GoogleComputeRegionAutoscalerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GoogleComputeRegionAutoscalerList{}
	err = c.client.Get().
		Resource("googlecomputeregionautoscalers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested googleComputeRegionAutoscalers.
func (c *googleComputeRegionAutoscalers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("googlecomputeregionautoscalers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a googleComputeRegionAutoscaler and creates it.  Returns the server's representation of the googleComputeRegionAutoscaler, and an error, if there is any.
func (c *googleComputeRegionAutoscalers) Create(googleComputeRegionAutoscaler *v1alpha1.GoogleComputeRegionAutoscaler) (result *v1alpha1.GoogleComputeRegionAutoscaler, err error) {
	result = &v1alpha1.GoogleComputeRegionAutoscaler{}
	err = c.client.Post().
		Resource("googlecomputeregionautoscalers").
		Body(googleComputeRegionAutoscaler).
		Do().
		Into(result)
	return
}

// Update takes the representation of a googleComputeRegionAutoscaler and updates it. Returns the server's representation of the googleComputeRegionAutoscaler, and an error, if there is any.
func (c *googleComputeRegionAutoscalers) Update(googleComputeRegionAutoscaler *v1alpha1.GoogleComputeRegionAutoscaler) (result *v1alpha1.GoogleComputeRegionAutoscaler, err error) {
	result = &v1alpha1.GoogleComputeRegionAutoscaler{}
	err = c.client.Put().
		Resource("googlecomputeregionautoscalers").
		Name(googleComputeRegionAutoscaler.Name).
		Body(googleComputeRegionAutoscaler).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *googleComputeRegionAutoscalers) UpdateStatus(googleComputeRegionAutoscaler *v1alpha1.GoogleComputeRegionAutoscaler) (result *v1alpha1.GoogleComputeRegionAutoscaler, err error) {
	result = &v1alpha1.GoogleComputeRegionAutoscaler{}
	err = c.client.Put().
		Resource("googlecomputeregionautoscalers").
		Name(googleComputeRegionAutoscaler.Name).
		SubResource("status").
		Body(googleComputeRegionAutoscaler).
		Do().
		Into(result)
	return
}

// Delete takes name of the googleComputeRegionAutoscaler and deletes it. Returns an error if one occurs.
func (c *googleComputeRegionAutoscalers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("googlecomputeregionautoscalers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *googleComputeRegionAutoscalers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("googlecomputeregionautoscalers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched googleComputeRegionAutoscaler.
func (c *googleComputeRegionAutoscalers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleComputeRegionAutoscaler, err error) {
	result = &v1alpha1.GoogleComputeRegionAutoscaler{}
	err = c.client.Patch(pt).
		Resource("googlecomputeregionautoscalers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
