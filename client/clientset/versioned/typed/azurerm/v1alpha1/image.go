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

	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ImagesGetter has a method to return a ImageInterface.
// A group's client should implement this interface.
type ImagesGetter interface {
	Images(namespace string) ImageInterface
}

// ImageInterface has methods to work with Image resources.
type ImageInterface interface {
	Create(*v1alpha1.Image) (*v1alpha1.Image, error)
	Update(*v1alpha1.Image) (*v1alpha1.Image, error)
	UpdateStatus(*v1alpha1.Image) (*v1alpha1.Image, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Image, error)
	List(opts v1.ListOptions) (*v1alpha1.ImageList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Image, err error)
	ImageExpansion
}

// images implements ImageInterface
type images struct {
	client rest.Interface
	ns     string
}

// newImages returns a Images
func newImages(c *AzurermV1alpha1Client, namespace string) *images {
	return &images{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the image, and returns the corresponding image object, and an error if there is any.
func (c *images) Get(name string, options v1.GetOptions) (result *v1alpha1.Image, err error) {
	result = &v1alpha1.Image{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("images").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Images that match those selectors.
func (c *images) List(opts v1.ListOptions) (result *v1alpha1.ImageList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ImageList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("images").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested images.
func (c *images) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("images").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a image and creates it.  Returns the server's representation of the image, and an error, if there is any.
func (c *images) Create(image *v1alpha1.Image) (result *v1alpha1.Image, err error) {
	result = &v1alpha1.Image{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("images").
		Body(image).
		Do().
		Into(result)
	return
}

// Update takes the representation of a image and updates it. Returns the server's representation of the image, and an error, if there is any.
func (c *images) Update(image *v1alpha1.Image) (result *v1alpha1.Image, err error) {
	result = &v1alpha1.Image{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("images").
		Name(image.Name).
		Body(image).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *images) UpdateStatus(image *v1alpha1.Image) (result *v1alpha1.Image, err error) {
	result = &v1alpha1.Image{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("images").
		Name(image.Name).
		SubResource("status").
		Body(image).
		Do().
		Into(result)
	return
}

// Delete takes name of the image and deletes it. Returns an error if one occurs.
func (c *images) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("images").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *images) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("images").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched image.
func (c *images) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Image, err error) {
	result = &v1alpha1.Image{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("images").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
