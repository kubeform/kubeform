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

// AzurermDnsTxtRecordsGetter has a method to return a AzurermDnsTxtRecordInterface.
// A group's client should implement this interface.
type AzurermDnsTxtRecordsGetter interface {
	AzurermDnsTxtRecords() AzurermDnsTxtRecordInterface
}

// AzurermDnsTxtRecordInterface has methods to work with AzurermDnsTxtRecord resources.
type AzurermDnsTxtRecordInterface interface {
	Create(*v1alpha1.AzurermDnsTxtRecord) (*v1alpha1.AzurermDnsTxtRecord, error)
	Update(*v1alpha1.AzurermDnsTxtRecord) (*v1alpha1.AzurermDnsTxtRecord, error)
	UpdateStatus(*v1alpha1.AzurermDnsTxtRecord) (*v1alpha1.AzurermDnsTxtRecord, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermDnsTxtRecord, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermDnsTxtRecordList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermDnsTxtRecord, err error)
	AzurermDnsTxtRecordExpansion
}

// azurermDnsTxtRecords implements AzurermDnsTxtRecordInterface
type azurermDnsTxtRecords struct {
	client rest.Interface
}

// newAzurermDnsTxtRecords returns a AzurermDnsTxtRecords
func newAzurermDnsTxtRecords(c *AzurermV1alpha1Client) *azurermDnsTxtRecords {
	return &azurermDnsTxtRecords{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermDnsTxtRecord, and returns the corresponding azurermDnsTxtRecord object, and an error if there is any.
func (c *azurermDnsTxtRecords) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermDnsTxtRecord, err error) {
	result = &v1alpha1.AzurermDnsTxtRecord{}
	err = c.client.Get().
		Resource("azurermdnstxtrecords").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermDnsTxtRecords that match those selectors.
func (c *azurermDnsTxtRecords) List(opts v1.ListOptions) (result *v1alpha1.AzurermDnsTxtRecordList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermDnsTxtRecordList{}
	err = c.client.Get().
		Resource("azurermdnstxtrecords").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermDnsTxtRecords.
func (c *azurermDnsTxtRecords) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermdnstxtrecords").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermDnsTxtRecord and creates it.  Returns the server's representation of the azurermDnsTxtRecord, and an error, if there is any.
func (c *azurermDnsTxtRecords) Create(azurermDnsTxtRecord *v1alpha1.AzurermDnsTxtRecord) (result *v1alpha1.AzurermDnsTxtRecord, err error) {
	result = &v1alpha1.AzurermDnsTxtRecord{}
	err = c.client.Post().
		Resource("azurermdnstxtrecords").
		Body(azurermDnsTxtRecord).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermDnsTxtRecord and updates it. Returns the server's representation of the azurermDnsTxtRecord, and an error, if there is any.
func (c *azurermDnsTxtRecords) Update(azurermDnsTxtRecord *v1alpha1.AzurermDnsTxtRecord) (result *v1alpha1.AzurermDnsTxtRecord, err error) {
	result = &v1alpha1.AzurermDnsTxtRecord{}
	err = c.client.Put().
		Resource("azurermdnstxtrecords").
		Name(azurermDnsTxtRecord.Name).
		Body(azurermDnsTxtRecord).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermDnsTxtRecords) UpdateStatus(azurermDnsTxtRecord *v1alpha1.AzurermDnsTxtRecord) (result *v1alpha1.AzurermDnsTxtRecord, err error) {
	result = &v1alpha1.AzurermDnsTxtRecord{}
	err = c.client.Put().
		Resource("azurermdnstxtrecords").
		Name(azurermDnsTxtRecord.Name).
		SubResource("status").
		Body(azurermDnsTxtRecord).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermDnsTxtRecord and deletes it. Returns an error if one occurs.
func (c *azurermDnsTxtRecords) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermdnstxtrecords").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermDnsTxtRecords) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermdnstxtrecords").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermDnsTxtRecord.
func (c *azurermDnsTxtRecords) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermDnsTxtRecord, err error) {
	result = &v1alpha1.AzurermDnsTxtRecord{}
	err = c.client.Patch(pt).
		Resource("azurermdnstxtrecords").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
