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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// AzurermSharedImageGalleriesGetter has a method to return a AzurermSharedImageGalleryInterface.
// A group's client should implement this interface.
type AzurermSharedImageGalleriesGetter interface {
	AzurermSharedImageGalleries() AzurermSharedImageGalleryInterface
}

// AzurermSharedImageGalleryInterface has methods to work with AzurermSharedImageGallery resources.
type AzurermSharedImageGalleryInterface interface {
	Create(*v1alpha1.AzurermSharedImageGallery) (*v1alpha1.AzurermSharedImageGallery, error)
	Update(*v1alpha1.AzurermSharedImageGallery) (*v1alpha1.AzurermSharedImageGallery, error)
	UpdateStatus(*v1alpha1.AzurermSharedImageGallery) (*v1alpha1.AzurermSharedImageGallery, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermSharedImageGallery, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermSharedImageGalleryList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermSharedImageGallery, err error)
	AzurermSharedImageGalleryExpansion
}

// azurermSharedImageGalleries implements AzurermSharedImageGalleryInterface
type azurermSharedImageGalleries struct {
	client rest.Interface
}

// newAzurermSharedImageGalleries returns a AzurermSharedImageGalleries
func newAzurermSharedImageGalleries(c *AzurermV1alpha1Client) *azurermSharedImageGalleries {
	return &azurermSharedImageGalleries{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermSharedImageGallery, and returns the corresponding azurermSharedImageGallery object, and an error if there is any.
func (c *azurermSharedImageGalleries) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermSharedImageGallery, err error) {
	result = &v1alpha1.AzurermSharedImageGallery{}
	err = c.client.Get().
		Resource("azurermsharedimagegalleries").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermSharedImageGalleries that match those selectors.
func (c *azurermSharedImageGalleries) List(opts v1.ListOptions) (result *v1alpha1.AzurermSharedImageGalleryList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermSharedImageGalleryList{}
	err = c.client.Get().
		Resource("azurermsharedimagegalleries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermSharedImageGalleries.
func (c *azurermSharedImageGalleries) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermsharedimagegalleries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermSharedImageGallery and creates it.  Returns the server's representation of the azurermSharedImageGallery, and an error, if there is any.
func (c *azurermSharedImageGalleries) Create(azurermSharedImageGallery *v1alpha1.AzurermSharedImageGallery) (result *v1alpha1.AzurermSharedImageGallery, err error) {
	result = &v1alpha1.AzurermSharedImageGallery{}
	err = c.client.Post().
		Resource("azurermsharedimagegalleries").
		Body(azurermSharedImageGallery).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermSharedImageGallery and updates it. Returns the server's representation of the azurermSharedImageGallery, and an error, if there is any.
func (c *azurermSharedImageGalleries) Update(azurermSharedImageGallery *v1alpha1.AzurermSharedImageGallery) (result *v1alpha1.AzurermSharedImageGallery, err error) {
	result = &v1alpha1.AzurermSharedImageGallery{}
	err = c.client.Put().
		Resource("azurermsharedimagegalleries").
		Name(azurermSharedImageGallery.Name).
		Body(azurermSharedImageGallery).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermSharedImageGalleries) UpdateStatus(azurermSharedImageGallery *v1alpha1.AzurermSharedImageGallery) (result *v1alpha1.AzurermSharedImageGallery, err error) {
	result = &v1alpha1.AzurermSharedImageGallery{}
	err = c.client.Put().
		Resource("azurermsharedimagegalleries").
		Name(azurermSharedImageGallery.Name).
		SubResource("status").
		Body(azurermSharedImageGallery).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermSharedImageGallery and deletes it. Returns an error if one occurs.
func (c *azurermSharedImageGalleries) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermsharedimagegalleries").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermSharedImageGalleries) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermsharedimagegalleries").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermSharedImageGallery.
func (c *azurermSharedImageGalleries) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermSharedImageGallery, err error) {
	result = &v1alpha1.AzurermSharedImageGallery{}
	err = c.client.Patch(pt).
		Resource("azurermsharedimagegalleries").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
