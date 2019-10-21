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

// EventgridDomainsGetter has a method to return a EventgridDomainInterface.
// A group's client should implement this interface.
type EventgridDomainsGetter interface {
	EventgridDomains(namespace string) EventgridDomainInterface
}

// EventgridDomainInterface has methods to work with EventgridDomain resources.
type EventgridDomainInterface interface {
	Create(*v1alpha1.EventgridDomain) (*v1alpha1.EventgridDomain, error)
	Update(*v1alpha1.EventgridDomain) (*v1alpha1.EventgridDomain, error)
	UpdateStatus(*v1alpha1.EventgridDomain) (*v1alpha1.EventgridDomain, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.EventgridDomain, error)
	List(opts v1.ListOptions) (*v1alpha1.EventgridDomainList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EventgridDomain, err error)
	EventgridDomainExpansion
}

// eventgridDomains implements EventgridDomainInterface
type eventgridDomains struct {
	client rest.Interface
	ns     string
}

// newEventgridDomains returns a EventgridDomains
func newEventgridDomains(c *AzurermV1alpha1Client, namespace string) *eventgridDomains {
	return &eventgridDomains{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the eventgridDomain, and returns the corresponding eventgridDomain object, and an error if there is any.
func (c *eventgridDomains) Get(name string, options v1.GetOptions) (result *v1alpha1.EventgridDomain, err error) {
	result = &v1alpha1.EventgridDomain{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("eventgriddomains").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EventgridDomains that match those selectors.
func (c *eventgridDomains) List(opts v1.ListOptions) (result *v1alpha1.EventgridDomainList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.EventgridDomainList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("eventgriddomains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested eventgridDomains.
func (c *eventgridDomains) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("eventgriddomains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a eventgridDomain and creates it.  Returns the server's representation of the eventgridDomain, and an error, if there is any.
func (c *eventgridDomains) Create(eventgridDomain *v1alpha1.EventgridDomain) (result *v1alpha1.EventgridDomain, err error) {
	result = &v1alpha1.EventgridDomain{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("eventgriddomains").
		Body(eventgridDomain).
		Do().
		Into(result)
	return
}

// Update takes the representation of a eventgridDomain and updates it. Returns the server's representation of the eventgridDomain, and an error, if there is any.
func (c *eventgridDomains) Update(eventgridDomain *v1alpha1.EventgridDomain) (result *v1alpha1.EventgridDomain, err error) {
	result = &v1alpha1.EventgridDomain{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("eventgriddomains").
		Name(eventgridDomain.Name).
		Body(eventgridDomain).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *eventgridDomains) UpdateStatus(eventgridDomain *v1alpha1.EventgridDomain) (result *v1alpha1.EventgridDomain, err error) {
	result = &v1alpha1.EventgridDomain{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("eventgriddomains").
		Name(eventgridDomain.Name).
		SubResource("status").
		Body(eventgridDomain).
		Do().
		Into(result)
	return
}

// Delete takes name of the eventgridDomain and deletes it. Returns an error if one occurs.
func (c *eventgridDomains) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("eventgriddomains").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *eventgridDomains) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("eventgriddomains").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched eventgridDomain.
func (c *eventgridDomains) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EventgridDomain, err error) {
	result = &v1alpha1.EventgridDomain{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("eventgriddomains").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}