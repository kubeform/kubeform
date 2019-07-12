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

// AzurermCosmosdbAccountsGetter has a method to return a AzurermCosmosdbAccountInterface.
// A group's client should implement this interface.
type AzurermCosmosdbAccountsGetter interface {
	AzurermCosmosdbAccounts() AzurermCosmosdbAccountInterface
}

// AzurermCosmosdbAccountInterface has methods to work with AzurermCosmosdbAccount resources.
type AzurermCosmosdbAccountInterface interface {
	Create(*v1alpha1.AzurermCosmosdbAccount) (*v1alpha1.AzurermCosmosdbAccount, error)
	Update(*v1alpha1.AzurermCosmosdbAccount) (*v1alpha1.AzurermCosmosdbAccount, error)
	UpdateStatus(*v1alpha1.AzurermCosmosdbAccount) (*v1alpha1.AzurermCosmosdbAccount, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermCosmosdbAccount, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermCosmosdbAccountList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermCosmosdbAccount, err error)
	AzurermCosmosdbAccountExpansion
}

// azurermCosmosdbAccounts implements AzurermCosmosdbAccountInterface
type azurermCosmosdbAccounts struct {
	client rest.Interface
}

// newAzurermCosmosdbAccounts returns a AzurermCosmosdbAccounts
func newAzurermCosmosdbAccounts(c *AzurermV1alpha1Client) *azurermCosmosdbAccounts {
	return &azurermCosmosdbAccounts{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermCosmosdbAccount, and returns the corresponding azurermCosmosdbAccount object, and an error if there is any.
func (c *azurermCosmosdbAccounts) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermCosmosdbAccount, err error) {
	result = &v1alpha1.AzurermCosmosdbAccount{}
	err = c.client.Get().
		Resource("azurermcosmosdbaccounts").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermCosmosdbAccounts that match those selectors.
func (c *azurermCosmosdbAccounts) List(opts v1.ListOptions) (result *v1alpha1.AzurermCosmosdbAccountList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermCosmosdbAccountList{}
	err = c.client.Get().
		Resource("azurermcosmosdbaccounts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermCosmosdbAccounts.
func (c *azurermCosmosdbAccounts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermcosmosdbaccounts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermCosmosdbAccount and creates it.  Returns the server's representation of the azurermCosmosdbAccount, and an error, if there is any.
func (c *azurermCosmosdbAccounts) Create(azurermCosmosdbAccount *v1alpha1.AzurermCosmosdbAccount) (result *v1alpha1.AzurermCosmosdbAccount, err error) {
	result = &v1alpha1.AzurermCosmosdbAccount{}
	err = c.client.Post().
		Resource("azurermcosmosdbaccounts").
		Body(azurermCosmosdbAccount).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermCosmosdbAccount and updates it. Returns the server's representation of the azurermCosmosdbAccount, and an error, if there is any.
func (c *azurermCosmosdbAccounts) Update(azurermCosmosdbAccount *v1alpha1.AzurermCosmosdbAccount) (result *v1alpha1.AzurermCosmosdbAccount, err error) {
	result = &v1alpha1.AzurermCosmosdbAccount{}
	err = c.client.Put().
		Resource("azurermcosmosdbaccounts").
		Name(azurermCosmosdbAccount.Name).
		Body(azurermCosmosdbAccount).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermCosmosdbAccounts) UpdateStatus(azurermCosmosdbAccount *v1alpha1.AzurermCosmosdbAccount) (result *v1alpha1.AzurermCosmosdbAccount, err error) {
	result = &v1alpha1.AzurermCosmosdbAccount{}
	err = c.client.Put().
		Resource("azurermcosmosdbaccounts").
		Name(azurermCosmosdbAccount.Name).
		SubResource("status").
		Body(azurermCosmosdbAccount).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermCosmosdbAccount and deletes it. Returns an error if one occurs.
func (c *azurermCosmosdbAccounts) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermcosmosdbaccounts").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermCosmosdbAccounts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermcosmosdbaccounts").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermCosmosdbAccount.
func (c *azurermCosmosdbAccounts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermCosmosdbAccount, err error) {
	result = &v1alpha1.AzurermCosmosdbAccount{}
	err = c.client.Patch(pt).
		Resource("azurermcosmosdbaccounts").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
