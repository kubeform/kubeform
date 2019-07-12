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

// AzurermDataFactoryDatasetSqlServerTablesGetter has a method to return a AzurermDataFactoryDatasetSqlServerTableInterface.
// A group's client should implement this interface.
type AzurermDataFactoryDatasetSqlServerTablesGetter interface {
	AzurermDataFactoryDatasetSqlServerTables() AzurermDataFactoryDatasetSqlServerTableInterface
}

// AzurermDataFactoryDatasetSqlServerTableInterface has methods to work with AzurermDataFactoryDatasetSqlServerTable resources.
type AzurermDataFactoryDatasetSqlServerTableInterface interface {
	Create(*v1alpha1.AzurermDataFactoryDatasetSqlServerTable) (*v1alpha1.AzurermDataFactoryDatasetSqlServerTable, error)
	Update(*v1alpha1.AzurermDataFactoryDatasetSqlServerTable) (*v1alpha1.AzurermDataFactoryDatasetSqlServerTable, error)
	UpdateStatus(*v1alpha1.AzurermDataFactoryDatasetSqlServerTable) (*v1alpha1.AzurermDataFactoryDatasetSqlServerTable, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermDataFactoryDatasetSqlServerTable, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermDataFactoryDatasetSqlServerTableList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermDataFactoryDatasetSqlServerTable, err error)
	AzurermDataFactoryDatasetSqlServerTableExpansion
}

// azurermDataFactoryDatasetSqlServerTables implements AzurermDataFactoryDatasetSqlServerTableInterface
type azurermDataFactoryDatasetSqlServerTables struct {
	client rest.Interface
}

// newAzurermDataFactoryDatasetSqlServerTables returns a AzurermDataFactoryDatasetSqlServerTables
func newAzurermDataFactoryDatasetSqlServerTables(c *AzurermV1alpha1Client) *azurermDataFactoryDatasetSqlServerTables {
	return &azurermDataFactoryDatasetSqlServerTables{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermDataFactoryDatasetSqlServerTable, and returns the corresponding azurermDataFactoryDatasetSqlServerTable object, and an error if there is any.
func (c *azurermDataFactoryDatasetSqlServerTables) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermDataFactoryDatasetSqlServerTable, err error) {
	result = &v1alpha1.AzurermDataFactoryDatasetSqlServerTable{}
	err = c.client.Get().
		Resource("azurermdatafactorydatasetsqlservertables").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermDataFactoryDatasetSqlServerTables that match those selectors.
func (c *azurermDataFactoryDatasetSqlServerTables) List(opts v1.ListOptions) (result *v1alpha1.AzurermDataFactoryDatasetSqlServerTableList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermDataFactoryDatasetSqlServerTableList{}
	err = c.client.Get().
		Resource("azurermdatafactorydatasetsqlservertables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermDataFactoryDatasetSqlServerTables.
func (c *azurermDataFactoryDatasetSqlServerTables) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermdatafactorydatasetsqlservertables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermDataFactoryDatasetSqlServerTable and creates it.  Returns the server's representation of the azurermDataFactoryDatasetSqlServerTable, and an error, if there is any.
func (c *azurermDataFactoryDatasetSqlServerTables) Create(azurermDataFactoryDatasetSqlServerTable *v1alpha1.AzurermDataFactoryDatasetSqlServerTable) (result *v1alpha1.AzurermDataFactoryDatasetSqlServerTable, err error) {
	result = &v1alpha1.AzurermDataFactoryDatasetSqlServerTable{}
	err = c.client.Post().
		Resource("azurermdatafactorydatasetsqlservertables").
		Body(azurermDataFactoryDatasetSqlServerTable).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermDataFactoryDatasetSqlServerTable and updates it. Returns the server's representation of the azurermDataFactoryDatasetSqlServerTable, and an error, if there is any.
func (c *azurermDataFactoryDatasetSqlServerTables) Update(azurermDataFactoryDatasetSqlServerTable *v1alpha1.AzurermDataFactoryDatasetSqlServerTable) (result *v1alpha1.AzurermDataFactoryDatasetSqlServerTable, err error) {
	result = &v1alpha1.AzurermDataFactoryDatasetSqlServerTable{}
	err = c.client.Put().
		Resource("azurermdatafactorydatasetsqlservertables").
		Name(azurermDataFactoryDatasetSqlServerTable.Name).
		Body(azurermDataFactoryDatasetSqlServerTable).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermDataFactoryDatasetSqlServerTables) UpdateStatus(azurermDataFactoryDatasetSqlServerTable *v1alpha1.AzurermDataFactoryDatasetSqlServerTable) (result *v1alpha1.AzurermDataFactoryDatasetSqlServerTable, err error) {
	result = &v1alpha1.AzurermDataFactoryDatasetSqlServerTable{}
	err = c.client.Put().
		Resource("azurermdatafactorydatasetsqlservertables").
		Name(azurermDataFactoryDatasetSqlServerTable.Name).
		SubResource("status").
		Body(azurermDataFactoryDatasetSqlServerTable).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermDataFactoryDatasetSqlServerTable and deletes it. Returns an error if one occurs.
func (c *azurermDataFactoryDatasetSqlServerTables) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermdatafactorydatasetsqlservertables").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermDataFactoryDatasetSqlServerTables) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermdatafactorydatasetsqlservertables").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermDataFactoryDatasetSqlServerTable.
func (c *azurermDataFactoryDatasetSqlServerTables) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermDataFactoryDatasetSqlServerTable, err error) {
	result = &v1alpha1.AzurermDataFactoryDatasetSqlServerTable{}
	err = c.client.Patch(pt).
		Resource("azurermdatafactorydatasetsqlservertables").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
