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

// PrivateLinkServicesGetter has a method to return a PrivateLinkServiceInterface.
// A group's client should implement this interface.
type PrivateLinkServicesGetter interface {
	PrivateLinkServices(namespace string) PrivateLinkServiceInterface
}

// PrivateLinkServiceInterface has methods to work with PrivateLinkService resources.
type PrivateLinkServiceInterface interface {
	Create(*v1alpha1.PrivateLinkService) (*v1alpha1.PrivateLinkService, error)
	Update(*v1alpha1.PrivateLinkService) (*v1alpha1.PrivateLinkService, error)
	UpdateStatus(*v1alpha1.PrivateLinkService) (*v1alpha1.PrivateLinkService, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.PrivateLinkService, error)
	List(opts v1.ListOptions) (*v1alpha1.PrivateLinkServiceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PrivateLinkService, err error)
	PrivateLinkServiceExpansion
}

// privateLinkServices implements PrivateLinkServiceInterface
type privateLinkServices struct {
	client rest.Interface
	ns     string
}

// newPrivateLinkServices returns a PrivateLinkServices
func newPrivateLinkServices(c *AzurermV1alpha1Client, namespace string) *privateLinkServices {
	return &privateLinkServices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the privateLinkService, and returns the corresponding privateLinkService object, and an error if there is any.
func (c *privateLinkServices) Get(name string, options v1.GetOptions) (result *v1alpha1.PrivateLinkService, err error) {
	result = &v1alpha1.PrivateLinkService{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("privatelinkservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PrivateLinkServices that match those selectors.
func (c *privateLinkServices) List(opts v1.ListOptions) (result *v1alpha1.PrivateLinkServiceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PrivateLinkServiceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("privatelinkservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested privateLinkServices.
func (c *privateLinkServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("privatelinkservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a privateLinkService and creates it.  Returns the server's representation of the privateLinkService, and an error, if there is any.
func (c *privateLinkServices) Create(privateLinkService *v1alpha1.PrivateLinkService) (result *v1alpha1.PrivateLinkService, err error) {
	result = &v1alpha1.PrivateLinkService{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("privatelinkservices").
		Body(privateLinkService).
		Do().
		Into(result)
	return
}

// Update takes the representation of a privateLinkService and updates it. Returns the server's representation of the privateLinkService, and an error, if there is any.
func (c *privateLinkServices) Update(privateLinkService *v1alpha1.PrivateLinkService) (result *v1alpha1.PrivateLinkService, err error) {
	result = &v1alpha1.PrivateLinkService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("privatelinkservices").
		Name(privateLinkService.Name).
		Body(privateLinkService).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *privateLinkServices) UpdateStatus(privateLinkService *v1alpha1.PrivateLinkService) (result *v1alpha1.PrivateLinkService, err error) {
	result = &v1alpha1.PrivateLinkService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("privatelinkservices").
		Name(privateLinkService.Name).
		SubResource("status").
		Body(privateLinkService).
		Do().
		Into(result)
	return
}

// Delete takes name of the privateLinkService and deletes it. Returns an error if one occurs.
func (c *privateLinkServices) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("privatelinkservices").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *privateLinkServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("privatelinkservices").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched privateLinkService.
func (c *privateLinkServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PrivateLinkService, err error) {
	result = &v1alpha1.PrivateLinkService{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("privatelinkservices").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
