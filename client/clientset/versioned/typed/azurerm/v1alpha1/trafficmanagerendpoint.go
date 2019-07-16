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

// TrafficManagerEndpointsGetter has a method to return a TrafficManagerEndpointInterface.
// A group's client should implement this interface.
type TrafficManagerEndpointsGetter interface {
	TrafficManagerEndpoints() TrafficManagerEndpointInterface
}

// TrafficManagerEndpointInterface has methods to work with TrafficManagerEndpoint resources.
type TrafficManagerEndpointInterface interface {
	Create(*v1alpha1.TrafficManagerEndpoint) (*v1alpha1.TrafficManagerEndpoint, error)
	Update(*v1alpha1.TrafficManagerEndpoint) (*v1alpha1.TrafficManagerEndpoint, error)
	UpdateStatus(*v1alpha1.TrafficManagerEndpoint) (*v1alpha1.TrafficManagerEndpoint, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.TrafficManagerEndpoint, error)
	List(opts v1.ListOptions) (*v1alpha1.TrafficManagerEndpointList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TrafficManagerEndpoint, err error)
	TrafficManagerEndpointExpansion
}

// trafficManagerEndpoints implements TrafficManagerEndpointInterface
type trafficManagerEndpoints struct {
	client rest.Interface
}

// newTrafficManagerEndpoints returns a TrafficManagerEndpoints
func newTrafficManagerEndpoints(c *AzurermV1alpha1Client) *trafficManagerEndpoints {
	return &trafficManagerEndpoints{
		client: c.RESTClient(),
	}
}

// Get takes name of the trafficManagerEndpoint, and returns the corresponding trafficManagerEndpoint object, and an error if there is any.
func (c *trafficManagerEndpoints) Get(name string, options v1.GetOptions) (result *v1alpha1.TrafficManagerEndpoint, err error) {
	result = &v1alpha1.TrafficManagerEndpoint{}
	err = c.client.Get().
		Resource("trafficmanagerendpoints").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TrafficManagerEndpoints that match those selectors.
func (c *trafficManagerEndpoints) List(opts v1.ListOptions) (result *v1alpha1.TrafficManagerEndpointList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.TrafficManagerEndpointList{}
	err = c.client.Get().
		Resource("trafficmanagerendpoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested trafficManagerEndpoints.
func (c *trafficManagerEndpoints) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("trafficmanagerendpoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a trafficManagerEndpoint and creates it.  Returns the server's representation of the trafficManagerEndpoint, and an error, if there is any.
func (c *trafficManagerEndpoints) Create(trafficManagerEndpoint *v1alpha1.TrafficManagerEndpoint) (result *v1alpha1.TrafficManagerEndpoint, err error) {
	result = &v1alpha1.TrafficManagerEndpoint{}
	err = c.client.Post().
		Resource("trafficmanagerendpoints").
		Body(trafficManagerEndpoint).
		Do().
		Into(result)
	return
}

// Update takes the representation of a trafficManagerEndpoint and updates it. Returns the server's representation of the trafficManagerEndpoint, and an error, if there is any.
func (c *trafficManagerEndpoints) Update(trafficManagerEndpoint *v1alpha1.TrafficManagerEndpoint) (result *v1alpha1.TrafficManagerEndpoint, err error) {
	result = &v1alpha1.TrafficManagerEndpoint{}
	err = c.client.Put().
		Resource("trafficmanagerendpoints").
		Name(trafficManagerEndpoint.Name).
		Body(trafficManagerEndpoint).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *trafficManagerEndpoints) UpdateStatus(trafficManagerEndpoint *v1alpha1.TrafficManagerEndpoint) (result *v1alpha1.TrafficManagerEndpoint, err error) {
	result = &v1alpha1.TrafficManagerEndpoint{}
	err = c.client.Put().
		Resource("trafficmanagerendpoints").
		Name(trafficManagerEndpoint.Name).
		SubResource("status").
		Body(trafficManagerEndpoint).
		Do().
		Into(result)
	return
}

// Delete takes name of the trafficManagerEndpoint and deletes it. Returns an error if one occurs.
func (c *trafficManagerEndpoints) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("trafficmanagerendpoints").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *trafficManagerEndpoints) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("trafficmanagerendpoints").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched trafficManagerEndpoint.
func (c *trafficManagerEndpoints) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TrafficManagerEndpoint, err error) {
	result = &v1alpha1.TrafficManagerEndpoint{}
	err = c.client.Patch(pt).
		Resource("trafficmanagerendpoints").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
