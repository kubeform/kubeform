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

// AzurermDnsSrvRecordsGetter has a method to return a AzurermDnsSrvRecordInterface.
// A group's client should implement this interface.
type AzurermDnsSrvRecordsGetter interface {
	AzurermDnsSrvRecords() AzurermDnsSrvRecordInterface
}

// AzurermDnsSrvRecordInterface has methods to work with AzurermDnsSrvRecord resources.
type AzurermDnsSrvRecordInterface interface {
	Create(*v1alpha1.AzurermDnsSrvRecord) (*v1alpha1.AzurermDnsSrvRecord, error)
	Update(*v1alpha1.AzurermDnsSrvRecord) (*v1alpha1.AzurermDnsSrvRecord, error)
	UpdateStatus(*v1alpha1.AzurermDnsSrvRecord) (*v1alpha1.AzurermDnsSrvRecord, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermDnsSrvRecord, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermDnsSrvRecordList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermDnsSrvRecord, err error)
	AzurermDnsSrvRecordExpansion
}

// azurermDnsSrvRecords implements AzurermDnsSrvRecordInterface
type azurermDnsSrvRecords struct {
	client rest.Interface
}

// newAzurermDnsSrvRecords returns a AzurermDnsSrvRecords
func newAzurermDnsSrvRecords(c *AzurermV1alpha1Client) *azurermDnsSrvRecords {
	return &azurermDnsSrvRecords{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermDnsSrvRecord, and returns the corresponding azurermDnsSrvRecord object, and an error if there is any.
func (c *azurermDnsSrvRecords) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermDnsSrvRecord, err error) {
	result = &v1alpha1.AzurermDnsSrvRecord{}
	err = c.client.Get().
		Resource("azurermdnssrvrecords").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermDnsSrvRecords that match those selectors.
func (c *azurermDnsSrvRecords) List(opts v1.ListOptions) (result *v1alpha1.AzurermDnsSrvRecordList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermDnsSrvRecordList{}
	err = c.client.Get().
		Resource("azurermdnssrvrecords").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermDnsSrvRecords.
func (c *azurermDnsSrvRecords) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermdnssrvrecords").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermDnsSrvRecord and creates it.  Returns the server's representation of the azurermDnsSrvRecord, and an error, if there is any.
func (c *azurermDnsSrvRecords) Create(azurermDnsSrvRecord *v1alpha1.AzurermDnsSrvRecord) (result *v1alpha1.AzurermDnsSrvRecord, err error) {
	result = &v1alpha1.AzurermDnsSrvRecord{}
	err = c.client.Post().
		Resource("azurermdnssrvrecords").
		Body(azurermDnsSrvRecord).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermDnsSrvRecord and updates it. Returns the server's representation of the azurermDnsSrvRecord, and an error, if there is any.
func (c *azurermDnsSrvRecords) Update(azurermDnsSrvRecord *v1alpha1.AzurermDnsSrvRecord) (result *v1alpha1.AzurermDnsSrvRecord, err error) {
	result = &v1alpha1.AzurermDnsSrvRecord{}
	err = c.client.Put().
		Resource("azurermdnssrvrecords").
		Name(azurermDnsSrvRecord.Name).
		Body(azurermDnsSrvRecord).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermDnsSrvRecords) UpdateStatus(azurermDnsSrvRecord *v1alpha1.AzurermDnsSrvRecord) (result *v1alpha1.AzurermDnsSrvRecord, err error) {
	result = &v1alpha1.AzurermDnsSrvRecord{}
	err = c.client.Put().
		Resource("azurermdnssrvrecords").
		Name(azurermDnsSrvRecord.Name).
		SubResource("status").
		Body(azurermDnsSrvRecord).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermDnsSrvRecord and deletes it. Returns an error if one occurs.
func (c *azurermDnsSrvRecords) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermdnssrvrecords").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermDnsSrvRecords) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermdnssrvrecords").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermDnsSrvRecord.
func (c *azurermDnsSrvRecords) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermDnsSrvRecord, err error) {
	result = &v1alpha1.AzurermDnsSrvRecord{}
	err = c.client.Patch(pt).
		Resource("azurermdnssrvrecords").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
