/*
Copyright 2019 The KubeDB Authors.

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

// GoogleEndpointsServicesGetter has a method to return a GoogleEndpointsServiceInterface.
// A group's client should implement this interface.
type GoogleEndpointsServicesGetter interface {
	GoogleEndpointsServices() GoogleEndpointsServiceInterface
}

// GoogleEndpointsServiceInterface has methods to work with GoogleEndpointsService resources.
type GoogleEndpointsServiceInterface interface {
	Create(*v1alpha1.GoogleEndpointsService) (*v1alpha1.GoogleEndpointsService, error)
	Update(*v1alpha1.GoogleEndpointsService) (*v1alpha1.GoogleEndpointsService, error)
	UpdateStatus(*v1alpha1.GoogleEndpointsService) (*v1alpha1.GoogleEndpointsService, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GoogleEndpointsService, error)
	List(opts v1.ListOptions) (*v1alpha1.GoogleEndpointsServiceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleEndpointsService, err error)
	GoogleEndpointsServiceExpansion
}

// googleEndpointsServices implements GoogleEndpointsServiceInterface
type googleEndpointsServices struct {
	client rest.Interface
}

// newGoogleEndpointsServices returns a GoogleEndpointsServices
func newGoogleEndpointsServices(c *GoogleV1alpha1Client) *googleEndpointsServices {
	return &googleEndpointsServices{
		client: c.RESTClient(),
	}
}

// Get takes name of the googleEndpointsService, and returns the corresponding googleEndpointsService object, and an error if there is any.
func (c *googleEndpointsServices) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleEndpointsService, err error) {
	result = &v1alpha1.GoogleEndpointsService{}
	err = c.client.Get().
		Resource("googleendpointsservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GoogleEndpointsServices that match those selectors.
func (c *googleEndpointsServices) List(opts v1.ListOptions) (result *v1alpha1.GoogleEndpointsServiceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GoogleEndpointsServiceList{}
	err = c.client.Get().
		Resource("googleendpointsservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested googleEndpointsServices.
func (c *googleEndpointsServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("googleendpointsservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a googleEndpointsService and creates it.  Returns the server's representation of the googleEndpointsService, and an error, if there is any.
func (c *googleEndpointsServices) Create(googleEndpointsService *v1alpha1.GoogleEndpointsService) (result *v1alpha1.GoogleEndpointsService, err error) {
	result = &v1alpha1.GoogleEndpointsService{}
	err = c.client.Post().
		Resource("googleendpointsservices").
		Body(googleEndpointsService).
		Do().
		Into(result)
	return
}

// Update takes the representation of a googleEndpointsService and updates it. Returns the server's representation of the googleEndpointsService, and an error, if there is any.
func (c *googleEndpointsServices) Update(googleEndpointsService *v1alpha1.GoogleEndpointsService) (result *v1alpha1.GoogleEndpointsService, err error) {
	result = &v1alpha1.GoogleEndpointsService{}
	err = c.client.Put().
		Resource("googleendpointsservices").
		Name(googleEndpointsService.Name).
		Body(googleEndpointsService).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *googleEndpointsServices) UpdateStatus(googleEndpointsService *v1alpha1.GoogleEndpointsService) (result *v1alpha1.GoogleEndpointsService, err error) {
	result = &v1alpha1.GoogleEndpointsService{}
	err = c.client.Put().
		Resource("googleendpointsservices").
		Name(googleEndpointsService.Name).
		SubResource("status").
		Body(googleEndpointsService).
		Do().
		Into(result)
	return
}

// Delete takes name of the googleEndpointsService and deletes it. Returns an error if one occurs.
func (c *googleEndpointsServices) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("googleendpointsservices").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *googleEndpointsServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("googleendpointsservices").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched googleEndpointsService.
func (c *googleEndpointsServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleEndpointsService, err error) {
	result = &v1alpha1.GoogleEndpointsService{}
	err = c.client.Patch(pt).
		Resource("googleendpointsservices").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
