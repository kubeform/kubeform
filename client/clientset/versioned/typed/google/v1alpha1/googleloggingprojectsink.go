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

// GoogleLoggingProjectSinksGetter has a method to return a GoogleLoggingProjectSinkInterface.
// A group's client should implement this interface.
type GoogleLoggingProjectSinksGetter interface {
	GoogleLoggingProjectSinks() GoogleLoggingProjectSinkInterface
}

// GoogleLoggingProjectSinkInterface has methods to work with GoogleLoggingProjectSink resources.
type GoogleLoggingProjectSinkInterface interface {
	Create(*v1alpha1.GoogleLoggingProjectSink) (*v1alpha1.GoogleLoggingProjectSink, error)
	Update(*v1alpha1.GoogleLoggingProjectSink) (*v1alpha1.GoogleLoggingProjectSink, error)
	UpdateStatus(*v1alpha1.GoogleLoggingProjectSink) (*v1alpha1.GoogleLoggingProjectSink, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GoogleLoggingProjectSink, error)
	List(opts v1.ListOptions) (*v1alpha1.GoogleLoggingProjectSinkList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleLoggingProjectSink, err error)
	GoogleLoggingProjectSinkExpansion
}

// googleLoggingProjectSinks implements GoogleLoggingProjectSinkInterface
type googleLoggingProjectSinks struct {
	client rest.Interface
}

// newGoogleLoggingProjectSinks returns a GoogleLoggingProjectSinks
func newGoogleLoggingProjectSinks(c *GoogleV1alpha1Client) *googleLoggingProjectSinks {
	return &googleLoggingProjectSinks{
		client: c.RESTClient(),
	}
}

// Get takes name of the googleLoggingProjectSink, and returns the corresponding googleLoggingProjectSink object, and an error if there is any.
func (c *googleLoggingProjectSinks) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleLoggingProjectSink, err error) {
	result = &v1alpha1.GoogleLoggingProjectSink{}
	err = c.client.Get().
		Resource("googleloggingprojectsinks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GoogleLoggingProjectSinks that match those selectors.
func (c *googleLoggingProjectSinks) List(opts v1.ListOptions) (result *v1alpha1.GoogleLoggingProjectSinkList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GoogleLoggingProjectSinkList{}
	err = c.client.Get().
		Resource("googleloggingprojectsinks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested googleLoggingProjectSinks.
func (c *googleLoggingProjectSinks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("googleloggingprojectsinks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a googleLoggingProjectSink and creates it.  Returns the server's representation of the googleLoggingProjectSink, and an error, if there is any.
func (c *googleLoggingProjectSinks) Create(googleLoggingProjectSink *v1alpha1.GoogleLoggingProjectSink) (result *v1alpha1.GoogleLoggingProjectSink, err error) {
	result = &v1alpha1.GoogleLoggingProjectSink{}
	err = c.client.Post().
		Resource("googleloggingprojectsinks").
		Body(googleLoggingProjectSink).
		Do().
		Into(result)
	return
}

// Update takes the representation of a googleLoggingProjectSink and updates it. Returns the server's representation of the googleLoggingProjectSink, and an error, if there is any.
func (c *googleLoggingProjectSinks) Update(googleLoggingProjectSink *v1alpha1.GoogleLoggingProjectSink) (result *v1alpha1.GoogleLoggingProjectSink, err error) {
	result = &v1alpha1.GoogleLoggingProjectSink{}
	err = c.client.Put().
		Resource("googleloggingprojectsinks").
		Name(googleLoggingProjectSink.Name).
		Body(googleLoggingProjectSink).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *googleLoggingProjectSinks) UpdateStatus(googleLoggingProjectSink *v1alpha1.GoogleLoggingProjectSink) (result *v1alpha1.GoogleLoggingProjectSink, err error) {
	result = &v1alpha1.GoogleLoggingProjectSink{}
	err = c.client.Put().
		Resource("googleloggingprojectsinks").
		Name(googleLoggingProjectSink.Name).
		SubResource("status").
		Body(googleLoggingProjectSink).
		Do().
		Into(result)
	return
}

// Delete takes name of the googleLoggingProjectSink and deletes it. Returns an error if one occurs.
func (c *googleLoggingProjectSinks) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("googleloggingprojectsinks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *googleLoggingProjectSinks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("googleloggingprojectsinks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched googleLoggingProjectSink.
func (c *googleLoggingProjectSinks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleLoggingProjectSink, err error) {
	result = &v1alpha1.GoogleLoggingProjectSink{}
	err = c.client.Patch(pt).
		Resource("googleloggingprojectsinks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
