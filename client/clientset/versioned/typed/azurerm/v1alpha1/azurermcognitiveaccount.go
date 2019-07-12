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

// AzurermCognitiveAccountsGetter has a method to return a AzurermCognitiveAccountInterface.
// A group's client should implement this interface.
type AzurermCognitiveAccountsGetter interface {
	AzurermCognitiveAccounts() AzurermCognitiveAccountInterface
}

// AzurermCognitiveAccountInterface has methods to work with AzurermCognitiveAccount resources.
type AzurermCognitiveAccountInterface interface {
	Create(*v1alpha1.AzurermCognitiveAccount) (*v1alpha1.AzurermCognitiveAccount, error)
	Update(*v1alpha1.AzurermCognitiveAccount) (*v1alpha1.AzurermCognitiveAccount, error)
	UpdateStatus(*v1alpha1.AzurermCognitiveAccount) (*v1alpha1.AzurermCognitiveAccount, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermCognitiveAccount, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermCognitiveAccountList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermCognitiveAccount, err error)
	AzurermCognitiveAccountExpansion
}

// azurermCognitiveAccounts implements AzurermCognitiveAccountInterface
type azurermCognitiveAccounts struct {
	client rest.Interface
}

// newAzurermCognitiveAccounts returns a AzurermCognitiveAccounts
func newAzurermCognitiveAccounts(c *AzurermV1alpha1Client) *azurermCognitiveAccounts {
	return &azurermCognitiveAccounts{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermCognitiveAccount, and returns the corresponding azurermCognitiveAccount object, and an error if there is any.
func (c *azurermCognitiveAccounts) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermCognitiveAccount, err error) {
	result = &v1alpha1.AzurermCognitiveAccount{}
	err = c.client.Get().
		Resource("azurermcognitiveaccounts").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermCognitiveAccounts that match those selectors.
func (c *azurermCognitiveAccounts) List(opts v1.ListOptions) (result *v1alpha1.AzurermCognitiveAccountList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermCognitiveAccountList{}
	err = c.client.Get().
		Resource("azurermcognitiveaccounts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermCognitiveAccounts.
func (c *azurermCognitiveAccounts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermcognitiveaccounts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermCognitiveAccount and creates it.  Returns the server's representation of the azurermCognitiveAccount, and an error, if there is any.
func (c *azurermCognitiveAccounts) Create(azurermCognitiveAccount *v1alpha1.AzurermCognitiveAccount) (result *v1alpha1.AzurermCognitiveAccount, err error) {
	result = &v1alpha1.AzurermCognitiveAccount{}
	err = c.client.Post().
		Resource("azurermcognitiveaccounts").
		Body(azurermCognitiveAccount).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermCognitiveAccount and updates it. Returns the server's representation of the azurermCognitiveAccount, and an error, if there is any.
func (c *azurermCognitiveAccounts) Update(azurermCognitiveAccount *v1alpha1.AzurermCognitiveAccount) (result *v1alpha1.AzurermCognitiveAccount, err error) {
	result = &v1alpha1.AzurermCognitiveAccount{}
	err = c.client.Put().
		Resource("azurermcognitiveaccounts").
		Name(azurermCognitiveAccount.Name).
		Body(azurermCognitiveAccount).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermCognitiveAccounts) UpdateStatus(azurermCognitiveAccount *v1alpha1.AzurermCognitiveAccount) (result *v1alpha1.AzurermCognitiveAccount, err error) {
	result = &v1alpha1.AzurermCognitiveAccount{}
	err = c.client.Put().
		Resource("azurermcognitiveaccounts").
		Name(azurermCognitiveAccount.Name).
		SubResource("status").
		Body(azurermCognitiveAccount).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermCognitiveAccount and deletes it. Returns an error if one occurs.
func (c *azurermCognitiveAccounts) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermcognitiveaccounts").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermCognitiveAccounts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermcognitiveaccounts").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermCognitiveAccount.
func (c *azurermCognitiveAccounts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermCognitiveAccount, err error) {
	result = &v1alpha1.AzurermCognitiveAccount{}
	err = c.client.Patch(pt).
		Resource("azurermcognitiveaccounts").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
