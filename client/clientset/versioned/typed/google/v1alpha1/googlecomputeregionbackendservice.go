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

// GoogleComputeRegionBackendServicesGetter has a method to return a GoogleComputeRegionBackendServiceInterface.
// A group's client should implement this interface.
type GoogleComputeRegionBackendServicesGetter interface {
	GoogleComputeRegionBackendServices() GoogleComputeRegionBackendServiceInterface
}

// GoogleComputeRegionBackendServiceInterface has methods to work with GoogleComputeRegionBackendService resources.
type GoogleComputeRegionBackendServiceInterface interface {
	Create(*v1alpha1.GoogleComputeRegionBackendService) (*v1alpha1.GoogleComputeRegionBackendService, error)
	Update(*v1alpha1.GoogleComputeRegionBackendService) (*v1alpha1.GoogleComputeRegionBackendService, error)
	UpdateStatus(*v1alpha1.GoogleComputeRegionBackendService) (*v1alpha1.GoogleComputeRegionBackendService, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GoogleComputeRegionBackendService, error)
	List(opts v1.ListOptions) (*v1alpha1.GoogleComputeRegionBackendServiceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleComputeRegionBackendService, err error)
	GoogleComputeRegionBackendServiceExpansion
}

// googleComputeRegionBackendServices implements GoogleComputeRegionBackendServiceInterface
type googleComputeRegionBackendServices struct {
	client rest.Interface
}

// newGoogleComputeRegionBackendServices returns a GoogleComputeRegionBackendServices
func newGoogleComputeRegionBackendServices(c *GoogleV1alpha1Client) *googleComputeRegionBackendServices {
	return &googleComputeRegionBackendServices{
		client: c.RESTClient(),
	}
}

// Get takes name of the googleComputeRegionBackendService, and returns the corresponding googleComputeRegionBackendService object, and an error if there is any.
func (c *googleComputeRegionBackendServices) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleComputeRegionBackendService, err error) {
	result = &v1alpha1.GoogleComputeRegionBackendService{}
	err = c.client.Get().
		Resource("googlecomputeregionbackendservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GoogleComputeRegionBackendServices that match those selectors.
func (c *googleComputeRegionBackendServices) List(opts v1.ListOptions) (result *v1alpha1.GoogleComputeRegionBackendServiceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GoogleComputeRegionBackendServiceList{}
	err = c.client.Get().
		Resource("googlecomputeregionbackendservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested googleComputeRegionBackendServices.
func (c *googleComputeRegionBackendServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("googlecomputeregionbackendservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a googleComputeRegionBackendService and creates it.  Returns the server's representation of the googleComputeRegionBackendService, and an error, if there is any.
func (c *googleComputeRegionBackendServices) Create(googleComputeRegionBackendService *v1alpha1.GoogleComputeRegionBackendService) (result *v1alpha1.GoogleComputeRegionBackendService, err error) {
	result = &v1alpha1.GoogleComputeRegionBackendService{}
	err = c.client.Post().
		Resource("googlecomputeregionbackendservices").
		Body(googleComputeRegionBackendService).
		Do().
		Into(result)
	return
}

// Update takes the representation of a googleComputeRegionBackendService and updates it. Returns the server's representation of the googleComputeRegionBackendService, and an error, if there is any.
func (c *googleComputeRegionBackendServices) Update(googleComputeRegionBackendService *v1alpha1.GoogleComputeRegionBackendService) (result *v1alpha1.GoogleComputeRegionBackendService, err error) {
	result = &v1alpha1.GoogleComputeRegionBackendService{}
	err = c.client.Put().
		Resource("googlecomputeregionbackendservices").
		Name(googleComputeRegionBackendService.Name).
		Body(googleComputeRegionBackendService).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *googleComputeRegionBackendServices) UpdateStatus(googleComputeRegionBackendService *v1alpha1.GoogleComputeRegionBackendService) (result *v1alpha1.GoogleComputeRegionBackendService, err error) {
	result = &v1alpha1.GoogleComputeRegionBackendService{}
	err = c.client.Put().
		Resource("googlecomputeregionbackendservices").
		Name(googleComputeRegionBackendService.Name).
		SubResource("status").
		Body(googleComputeRegionBackendService).
		Do().
		Into(result)
	return
}

// Delete takes name of the googleComputeRegionBackendService and deletes it. Returns an error if one occurs.
func (c *googleComputeRegionBackendServices) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("googlecomputeregionbackendservices").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *googleComputeRegionBackendServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("googlecomputeregionbackendservices").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched googleComputeRegionBackendService.
func (c *googleComputeRegionBackendServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleComputeRegionBackendService, err error) {
	result = &v1alpha1.GoogleComputeRegionBackendService{}
	err = c.client.Patch(pt).
		Resource("googlecomputeregionbackendservices").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
