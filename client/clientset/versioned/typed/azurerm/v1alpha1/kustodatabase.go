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

// KustoDatabasesGetter has a method to return a KustoDatabaseInterface.
// A group's client should implement this interface.
type KustoDatabasesGetter interface {
	KustoDatabases(namespace string) KustoDatabaseInterface
}

// KustoDatabaseInterface has methods to work with KustoDatabase resources.
type KustoDatabaseInterface interface {
	Create(*v1alpha1.KustoDatabase) (*v1alpha1.KustoDatabase, error)
	Update(*v1alpha1.KustoDatabase) (*v1alpha1.KustoDatabase, error)
	UpdateStatus(*v1alpha1.KustoDatabase) (*v1alpha1.KustoDatabase, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.KustoDatabase, error)
	List(opts v1.ListOptions) (*v1alpha1.KustoDatabaseList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KustoDatabase, err error)
	KustoDatabaseExpansion
}

// kustoDatabases implements KustoDatabaseInterface
type kustoDatabases struct {
	client rest.Interface
	ns     string
}

// newKustoDatabases returns a KustoDatabases
func newKustoDatabases(c *AzurermV1alpha1Client, namespace string) *kustoDatabases {
	return &kustoDatabases{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kustoDatabase, and returns the corresponding kustoDatabase object, and an error if there is any.
func (c *kustoDatabases) Get(name string, options v1.GetOptions) (result *v1alpha1.KustoDatabase, err error) {
	result = &v1alpha1.KustoDatabase{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kustodatabases").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KustoDatabases that match those selectors.
func (c *kustoDatabases) List(opts v1.ListOptions) (result *v1alpha1.KustoDatabaseList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.KustoDatabaseList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kustodatabases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kustoDatabases.
func (c *kustoDatabases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kustodatabases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a kustoDatabase and creates it.  Returns the server's representation of the kustoDatabase, and an error, if there is any.
func (c *kustoDatabases) Create(kustoDatabase *v1alpha1.KustoDatabase) (result *v1alpha1.KustoDatabase, err error) {
	result = &v1alpha1.KustoDatabase{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kustodatabases").
		Body(kustoDatabase).
		Do().
		Into(result)
	return
}

// Update takes the representation of a kustoDatabase and updates it. Returns the server's representation of the kustoDatabase, and an error, if there is any.
func (c *kustoDatabases) Update(kustoDatabase *v1alpha1.KustoDatabase) (result *v1alpha1.KustoDatabase, err error) {
	result = &v1alpha1.KustoDatabase{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kustodatabases").
		Name(kustoDatabase.Name).
		Body(kustoDatabase).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *kustoDatabases) UpdateStatus(kustoDatabase *v1alpha1.KustoDatabase) (result *v1alpha1.KustoDatabase, err error) {
	result = &v1alpha1.KustoDatabase{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kustodatabases").
		Name(kustoDatabase.Name).
		SubResource("status").
		Body(kustoDatabase).
		Do().
		Into(result)
	return
}

// Delete takes name of the kustoDatabase and deletes it. Returns an error if one occurs.
func (c *kustoDatabases) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kustodatabases").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kustoDatabases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kustodatabases").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched kustoDatabase.
func (c *kustoDatabases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KustoDatabase, err error) {
	result = &v1alpha1.KustoDatabase{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kustodatabases").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
