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

// GoogleBigtableTablesGetter has a method to return a GoogleBigtableTableInterface.
// A group's client should implement this interface.
type GoogleBigtableTablesGetter interface {
	GoogleBigtableTables() GoogleBigtableTableInterface
}

// GoogleBigtableTableInterface has methods to work with GoogleBigtableTable resources.
type GoogleBigtableTableInterface interface {
	Create(*v1alpha1.GoogleBigtableTable) (*v1alpha1.GoogleBigtableTable, error)
	Update(*v1alpha1.GoogleBigtableTable) (*v1alpha1.GoogleBigtableTable, error)
	UpdateStatus(*v1alpha1.GoogleBigtableTable) (*v1alpha1.GoogleBigtableTable, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GoogleBigtableTable, error)
	List(opts v1.ListOptions) (*v1alpha1.GoogleBigtableTableList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleBigtableTable, err error)
	GoogleBigtableTableExpansion
}

// googleBigtableTables implements GoogleBigtableTableInterface
type googleBigtableTables struct {
	client rest.Interface
}

// newGoogleBigtableTables returns a GoogleBigtableTables
func newGoogleBigtableTables(c *GoogleV1alpha1Client) *googleBigtableTables {
	return &googleBigtableTables{
		client: c.RESTClient(),
	}
}

// Get takes name of the googleBigtableTable, and returns the corresponding googleBigtableTable object, and an error if there is any.
func (c *googleBigtableTables) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleBigtableTable, err error) {
	result = &v1alpha1.GoogleBigtableTable{}
	err = c.client.Get().
		Resource("googlebigtabletables").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GoogleBigtableTables that match those selectors.
func (c *googleBigtableTables) List(opts v1.ListOptions) (result *v1alpha1.GoogleBigtableTableList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GoogleBigtableTableList{}
	err = c.client.Get().
		Resource("googlebigtabletables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested googleBigtableTables.
func (c *googleBigtableTables) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("googlebigtabletables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a googleBigtableTable and creates it.  Returns the server's representation of the googleBigtableTable, and an error, if there is any.
func (c *googleBigtableTables) Create(googleBigtableTable *v1alpha1.GoogleBigtableTable) (result *v1alpha1.GoogleBigtableTable, err error) {
	result = &v1alpha1.GoogleBigtableTable{}
	err = c.client.Post().
		Resource("googlebigtabletables").
		Body(googleBigtableTable).
		Do().
		Into(result)
	return
}

// Update takes the representation of a googleBigtableTable and updates it. Returns the server's representation of the googleBigtableTable, and an error, if there is any.
func (c *googleBigtableTables) Update(googleBigtableTable *v1alpha1.GoogleBigtableTable) (result *v1alpha1.GoogleBigtableTable, err error) {
	result = &v1alpha1.GoogleBigtableTable{}
	err = c.client.Put().
		Resource("googlebigtabletables").
		Name(googleBigtableTable.Name).
		Body(googleBigtableTable).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *googleBigtableTables) UpdateStatus(googleBigtableTable *v1alpha1.GoogleBigtableTable) (result *v1alpha1.GoogleBigtableTable, err error) {
	result = &v1alpha1.GoogleBigtableTable{}
	err = c.client.Put().
		Resource("googlebigtabletables").
		Name(googleBigtableTable.Name).
		SubResource("status").
		Body(googleBigtableTable).
		Do().
		Into(result)
	return
}

// Delete takes name of the googleBigtableTable and deletes it. Returns an error if one occurs.
func (c *googleBigtableTables) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("googlebigtabletables").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *googleBigtableTables) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("googlebigtabletables").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched googleBigtableTable.
func (c *googleBigtableTables) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleBigtableTable, err error) {
	result = &v1alpha1.GoogleBigtableTable{}
	err = c.client.Patch(pt).
		Resource("googlebigtabletables").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
