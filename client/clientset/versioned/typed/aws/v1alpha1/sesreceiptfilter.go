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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// SesReceiptFiltersGetter has a method to return a SesReceiptFilterInterface.
// A group's client should implement this interface.
type SesReceiptFiltersGetter interface {
	SesReceiptFilters() SesReceiptFilterInterface
}

// SesReceiptFilterInterface has methods to work with SesReceiptFilter resources.
type SesReceiptFilterInterface interface {
	Create(*v1alpha1.SesReceiptFilter) (*v1alpha1.SesReceiptFilter, error)
	Update(*v1alpha1.SesReceiptFilter) (*v1alpha1.SesReceiptFilter, error)
	UpdateStatus(*v1alpha1.SesReceiptFilter) (*v1alpha1.SesReceiptFilter, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SesReceiptFilter, error)
	List(opts v1.ListOptions) (*v1alpha1.SesReceiptFilterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SesReceiptFilter, err error)
	SesReceiptFilterExpansion
}

// sesReceiptFilters implements SesReceiptFilterInterface
type sesReceiptFilters struct {
	client rest.Interface
}

// newSesReceiptFilters returns a SesReceiptFilters
func newSesReceiptFilters(c *AwsV1alpha1Client) *sesReceiptFilters {
	return &sesReceiptFilters{
		client: c.RESTClient(),
	}
}

// Get takes name of the sesReceiptFilter, and returns the corresponding sesReceiptFilter object, and an error if there is any.
func (c *sesReceiptFilters) Get(name string, options v1.GetOptions) (result *v1alpha1.SesReceiptFilter, err error) {
	result = &v1alpha1.SesReceiptFilter{}
	err = c.client.Get().
		Resource("sesreceiptfilters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SesReceiptFilters that match those selectors.
func (c *sesReceiptFilters) List(opts v1.ListOptions) (result *v1alpha1.SesReceiptFilterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SesReceiptFilterList{}
	err = c.client.Get().
		Resource("sesreceiptfilters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sesReceiptFilters.
func (c *sesReceiptFilters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("sesreceiptfilters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a sesReceiptFilter and creates it.  Returns the server's representation of the sesReceiptFilter, and an error, if there is any.
func (c *sesReceiptFilters) Create(sesReceiptFilter *v1alpha1.SesReceiptFilter) (result *v1alpha1.SesReceiptFilter, err error) {
	result = &v1alpha1.SesReceiptFilter{}
	err = c.client.Post().
		Resource("sesreceiptfilters").
		Body(sesReceiptFilter).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sesReceiptFilter and updates it. Returns the server's representation of the sesReceiptFilter, and an error, if there is any.
func (c *sesReceiptFilters) Update(sesReceiptFilter *v1alpha1.SesReceiptFilter) (result *v1alpha1.SesReceiptFilter, err error) {
	result = &v1alpha1.SesReceiptFilter{}
	err = c.client.Put().
		Resource("sesreceiptfilters").
		Name(sesReceiptFilter.Name).
		Body(sesReceiptFilter).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *sesReceiptFilters) UpdateStatus(sesReceiptFilter *v1alpha1.SesReceiptFilter) (result *v1alpha1.SesReceiptFilter, err error) {
	result = &v1alpha1.SesReceiptFilter{}
	err = c.client.Put().
		Resource("sesreceiptfilters").
		Name(sesReceiptFilter.Name).
		SubResource("status").
		Body(sesReceiptFilter).
		Do().
		Into(result)
	return
}

// Delete takes name of the sesReceiptFilter and deletes it. Returns an error if one occurs.
func (c *sesReceiptFilters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("sesreceiptfilters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sesReceiptFilters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("sesreceiptfilters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sesReceiptFilter.
func (c *sesReceiptFilters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SesReceiptFilter, err error) {
	result = &v1alpha1.SesReceiptFilter{}
	err = c.client.Patch(pt).
		Resource("sesreceiptfilters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
