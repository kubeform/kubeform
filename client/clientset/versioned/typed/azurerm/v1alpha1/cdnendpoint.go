/*
Copyright The Kubeform Authors.

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

	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CdnEndpointsGetter has a method to return a CdnEndpointInterface.
// A group's client should implement this interface.
type CdnEndpointsGetter interface {
	CdnEndpoints(namespace string) CdnEndpointInterface
}

// CdnEndpointInterface has methods to work with CdnEndpoint resources.
type CdnEndpointInterface interface {
	Create(*v1alpha1.CdnEndpoint) (*v1alpha1.CdnEndpoint, error)
	Update(*v1alpha1.CdnEndpoint) (*v1alpha1.CdnEndpoint, error)
	UpdateStatus(*v1alpha1.CdnEndpoint) (*v1alpha1.CdnEndpoint, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CdnEndpoint, error)
	List(opts v1.ListOptions) (*v1alpha1.CdnEndpointList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CdnEndpoint, err error)
	CdnEndpointExpansion
}

// cdnEndpoints implements CdnEndpointInterface
type cdnEndpoints struct {
	client rest.Interface
	ns     string
}

// newCdnEndpoints returns a CdnEndpoints
func newCdnEndpoints(c *AzurermV1alpha1Client, namespace string) *cdnEndpoints {
	return &cdnEndpoints{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cdnEndpoint, and returns the corresponding cdnEndpoint object, and an error if there is any.
func (c *cdnEndpoints) Get(name string, options v1.GetOptions) (result *v1alpha1.CdnEndpoint, err error) {
	result = &v1alpha1.CdnEndpoint{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cdnendpoints").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CdnEndpoints that match those selectors.
func (c *cdnEndpoints) List(opts v1.ListOptions) (result *v1alpha1.CdnEndpointList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CdnEndpointList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cdnendpoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cdnEndpoints.
func (c *cdnEndpoints) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cdnendpoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cdnEndpoint and creates it.  Returns the server's representation of the cdnEndpoint, and an error, if there is any.
func (c *cdnEndpoints) Create(cdnEndpoint *v1alpha1.CdnEndpoint) (result *v1alpha1.CdnEndpoint, err error) {
	result = &v1alpha1.CdnEndpoint{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cdnendpoints").
		Body(cdnEndpoint).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cdnEndpoint and updates it. Returns the server's representation of the cdnEndpoint, and an error, if there is any.
func (c *cdnEndpoints) Update(cdnEndpoint *v1alpha1.CdnEndpoint) (result *v1alpha1.CdnEndpoint, err error) {
	result = &v1alpha1.CdnEndpoint{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cdnendpoints").
		Name(cdnEndpoint.Name).
		Body(cdnEndpoint).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *cdnEndpoints) UpdateStatus(cdnEndpoint *v1alpha1.CdnEndpoint) (result *v1alpha1.CdnEndpoint, err error) {
	result = &v1alpha1.CdnEndpoint{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cdnendpoints").
		Name(cdnEndpoint.Name).
		SubResource("status").
		Body(cdnEndpoint).
		Do().
		Into(result)
	return
}

// Delete takes name of the cdnEndpoint and deletes it. Returns an error if one occurs.
func (c *cdnEndpoints) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cdnendpoints").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cdnEndpoints) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cdnendpoints").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cdnEndpoint.
func (c *cdnEndpoints) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CdnEndpoint, err error) {
	result = &v1alpha1.CdnEndpoint{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cdnendpoints").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
