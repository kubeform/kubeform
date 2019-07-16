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

// LoggingOrganizationSinksGetter has a method to return a LoggingOrganizationSinkInterface.
// A group's client should implement this interface.
type LoggingOrganizationSinksGetter interface {
	LoggingOrganizationSinks() LoggingOrganizationSinkInterface
}

// LoggingOrganizationSinkInterface has methods to work with LoggingOrganizationSink resources.
type LoggingOrganizationSinkInterface interface {
	Create(*v1alpha1.LoggingOrganizationSink) (*v1alpha1.LoggingOrganizationSink, error)
	Update(*v1alpha1.LoggingOrganizationSink) (*v1alpha1.LoggingOrganizationSink, error)
	UpdateStatus(*v1alpha1.LoggingOrganizationSink) (*v1alpha1.LoggingOrganizationSink, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.LoggingOrganizationSink, error)
	List(opts v1.ListOptions) (*v1alpha1.LoggingOrganizationSinkList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LoggingOrganizationSink, err error)
	LoggingOrganizationSinkExpansion
}

// loggingOrganizationSinks implements LoggingOrganizationSinkInterface
type loggingOrganizationSinks struct {
	client rest.Interface
}

// newLoggingOrganizationSinks returns a LoggingOrganizationSinks
func newLoggingOrganizationSinks(c *GoogleV1alpha1Client) *loggingOrganizationSinks {
	return &loggingOrganizationSinks{
		client: c.RESTClient(),
	}
}

// Get takes name of the loggingOrganizationSink, and returns the corresponding loggingOrganizationSink object, and an error if there is any.
func (c *loggingOrganizationSinks) Get(name string, options v1.GetOptions) (result *v1alpha1.LoggingOrganizationSink, err error) {
	result = &v1alpha1.LoggingOrganizationSink{}
	err = c.client.Get().
		Resource("loggingorganizationsinks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LoggingOrganizationSinks that match those selectors.
func (c *loggingOrganizationSinks) List(opts v1.ListOptions) (result *v1alpha1.LoggingOrganizationSinkList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LoggingOrganizationSinkList{}
	err = c.client.Get().
		Resource("loggingorganizationsinks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested loggingOrganizationSinks.
func (c *loggingOrganizationSinks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("loggingorganizationsinks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a loggingOrganizationSink and creates it.  Returns the server's representation of the loggingOrganizationSink, and an error, if there is any.
func (c *loggingOrganizationSinks) Create(loggingOrganizationSink *v1alpha1.LoggingOrganizationSink) (result *v1alpha1.LoggingOrganizationSink, err error) {
	result = &v1alpha1.LoggingOrganizationSink{}
	err = c.client.Post().
		Resource("loggingorganizationsinks").
		Body(loggingOrganizationSink).
		Do().
		Into(result)
	return
}

// Update takes the representation of a loggingOrganizationSink and updates it. Returns the server's representation of the loggingOrganizationSink, and an error, if there is any.
func (c *loggingOrganizationSinks) Update(loggingOrganizationSink *v1alpha1.LoggingOrganizationSink) (result *v1alpha1.LoggingOrganizationSink, err error) {
	result = &v1alpha1.LoggingOrganizationSink{}
	err = c.client.Put().
		Resource("loggingorganizationsinks").
		Name(loggingOrganizationSink.Name).
		Body(loggingOrganizationSink).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *loggingOrganizationSinks) UpdateStatus(loggingOrganizationSink *v1alpha1.LoggingOrganizationSink) (result *v1alpha1.LoggingOrganizationSink, err error) {
	result = &v1alpha1.LoggingOrganizationSink{}
	err = c.client.Put().
		Resource("loggingorganizationsinks").
		Name(loggingOrganizationSink.Name).
		SubResource("status").
		Body(loggingOrganizationSink).
		Do().
		Into(result)
	return
}

// Delete takes name of the loggingOrganizationSink and deletes it. Returns an error if one occurs.
func (c *loggingOrganizationSinks) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("loggingorganizationsinks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *loggingOrganizationSinks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("loggingorganizationsinks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched loggingOrganizationSink.
func (c *loggingOrganizationSinks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LoggingOrganizationSink, err error) {
	result = &v1alpha1.LoggingOrganizationSink{}
	err = c.client.Patch(pt).
		Resource("loggingorganizationsinks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
