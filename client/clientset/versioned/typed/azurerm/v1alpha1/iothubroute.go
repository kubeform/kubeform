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

// IothubRoutesGetter has a method to return a IothubRouteInterface.
// A group's client should implement this interface.
type IothubRoutesGetter interface {
	IothubRoutes(namespace string) IothubRouteInterface
}

// IothubRouteInterface has methods to work with IothubRoute resources.
type IothubRouteInterface interface {
	Create(*v1alpha1.IothubRoute) (*v1alpha1.IothubRoute, error)
	Update(*v1alpha1.IothubRoute) (*v1alpha1.IothubRoute, error)
	UpdateStatus(*v1alpha1.IothubRoute) (*v1alpha1.IothubRoute, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.IothubRoute, error)
	List(opts v1.ListOptions) (*v1alpha1.IothubRouteList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IothubRoute, err error)
	IothubRouteExpansion
}

// iothubRoutes implements IothubRouteInterface
type iothubRoutes struct {
	client rest.Interface
	ns     string
}

// newIothubRoutes returns a IothubRoutes
func newIothubRoutes(c *AzurermV1alpha1Client, namespace string) *iothubRoutes {
	return &iothubRoutes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the iothubRoute, and returns the corresponding iothubRoute object, and an error if there is any.
func (c *iothubRoutes) Get(name string, options v1.GetOptions) (result *v1alpha1.IothubRoute, err error) {
	result = &v1alpha1.IothubRoute{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iothubroutes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IothubRoutes that match those selectors.
func (c *iothubRoutes) List(opts v1.ListOptions) (result *v1alpha1.IothubRouteList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.IothubRouteList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iothubroutes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iothubRoutes.
func (c *iothubRoutes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("iothubroutes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a iothubRoute and creates it.  Returns the server's representation of the iothubRoute, and an error, if there is any.
func (c *iothubRoutes) Create(iothubRoute *v1alpha1.IothubRoute) (result *v1alpha1.IothubRoute, err error) {
	result = &v1alpha1.IothubRoute{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("iothubroutes").
		Body(iothubRoute).
		Do().
		Into(result)
	return
}

// Update takes the representation of a iothubRoute and updates it. Returns the server's representation of the iothubRoute, and an error, if there is any.
func (c *iothubRoutes) Update(iothubRoute *v1alpha1.IothubRoute) (result *v1alpha1.IothubRoute, err error) {
	result = &v1alpha1.IothubRoute{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iothubroutes").
		Name(iothubRoute.Name).
		Body(iothubRoute).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *iothubRoutes) UpdateStatus(iothubRoute *v1alpha1.IothubRoute) (result *v1alpha1.IothubRoute, err error) {
	result = &v1alpha1.IothubRoute{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iothubroutes").
		Name(iothubRoute.Name).
		SubResource("status").
		Body(iothubRoute).
		Do().
		Into(result)
	return
}

// Delete takes name of the iothubRoute and deletes it. Returns an error if one occurs.
func (c *iothubRoutes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iothubroutes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *iothubRoutes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iothubroutes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched iothubRoute.
func (c *iothubRoutes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IothubRoute, err error) {
	result = &v1alpha1.IothubRoute{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("iothubroutes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
