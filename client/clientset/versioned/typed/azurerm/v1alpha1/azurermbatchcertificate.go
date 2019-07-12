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

// AzurermBatchCertificatesGetter has a method to return a AzurermBatchCertificateInterface.
// A group's client should implement this interface.
type AzurermBatchCertificatesGetter interface {
	AzurermBatchCertificates() AzurermBatchCertificateInterface
}

// AzurermBatchCertificateInterface has methods to work with AzurermBatchCertificate resources.
type AzurermBatchCertificateInterface interface {
	Create(*v1alpha1.AzurermBatchCertificate) (*v1alpha1.AzurermBatchCertificate, error)
	Update(*v1alpha1.AzurermBatchCertificate) (*v1alpha1.AzurermBatchCertificate, error)
	UpdateStatus(*v1alpha1.AzurermBatchCertificate) (*v1alpha1.AzurermBatchCertificate, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermBatchCertificate, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermBatchCertificateList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermBatchCertificate, err error)
	AzurermBatchCertificateExpansion
}

// azurermBatchCertificates implements AzurermBatchCertificateInterface
type azurermBatchCertificates struct {
	client rest.Interface
}

// newAzurermBatchCertificates returns a AzurermBatchCertificates
func newAzurermBatchCertificates(c *AzurermV1alpha1Client) *azurermBatchCertificates {
	return &azurermBatchCertificates{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermBatchCertificate, and returns the corresponding azurermBatchCertificate object, and an error if there is any.
func (c *azurermBatchCertificates) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermBatchCertificate, err error) {
	result = &v1alpha1.AzurermBatchCertificate{}
	err = c.client.Get().
		Resource("azurermbatchcertificates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermBatchCertificates that match those selectors.
func (c *azurermBatchCertificates) List(opts v1.ListOptions) (result *v1alpha1.AzurermBatchCertificateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermBatchCertificateList{}
	err = c.client.Get().
		Resource("azurermbatchcertificates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermBatchCertificates.
func (c *azurermBatchCertificates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermbatchcertificates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermBatchCertificate and creates it.  Returns the server's representation of the azurermBatchCertificate, and an error, if there is any.
func (c *azurermBatchCertificates) Create(azurermBatchCertificate *v1alpha1.AzurermBatchCertificate) (result *v1alpha1.AzurermBatchCertificate, err error) {
	result = &v1alpha1.AzurermBatchCertificate{}
	err = c.client.Post().
		Resource("azurermbatchcertificates").
		Body(azurermBatchCertificate).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermBatchCertificate and updates it. Returns the server's representation of the azurermBatchCertificate, and an error, if there is any.
func (c *azurermBatchCertificates) Update(azurermBatchCertificate *v1alpha1.AzurermBatchCertificate) (result *v1alpha1.AzurermBatchCertificate, err error) {
	result = &v1alpha1.AzurermBatchCertificate{}
	err = c.client.Put().
		Resource("azurermbatchcertificates").
		Name(azurermBatchCertificate.Name).
		Body(azurermBatchCertificate).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermBatchCertificates) UpdateStatus(azurermBatchCertificate *v1alpha1.AzurermBatchCertificate) (result *v1alpha1.AzurermBatchCertificate, err error) {
	result = &v1alpha1.AzurermBatchCertificate{}
	err = c.client.Put().
		Resource("azurermbatchcertificates").
		Name(azurermBatchCertificate.Name).
		SubResource("status").
		Body(azurermBatchCertificate).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermBatchCertificate and deletes it. Returns an error if one occurs.
func (c *azurermBatchCertificates) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermbatchcertificates").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermBatchCertificates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermbatchcertificates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermBatchCertificate.
func (c *azurermBatchCertificates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermBatchCertificate, err error) {
	result = &v1alpha1.AzurermBatchCertificate{}
	err = c.client.Patch(pt).
		Resource("azurermbatchcertificates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
