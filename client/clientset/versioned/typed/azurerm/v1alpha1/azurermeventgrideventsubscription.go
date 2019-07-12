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

// AzurermEventgridEventSubscriptionsGetter has a method to return a AzurermEventgridEventSubscriptionInterface.
// A group's client should implement this interface.
type AzurermEventgridEventSubscriptionsGetter interface {
	AzurermEventgridEventSubscriptions() AzurermEventgridEventSubscriptionInterface
}

// AzurermEventgridEventSubscriptionInterface has methods to work with AzurermEventgridEventSubscription resources.
type AzurermEventgridEventSubscriptionInterface interface {
	Create(*v1alpha1.AzurermEventgridEventSubscription) (*v1alpha1.AzurermEventgridEventSubscription, error)
	Update(*v1alpha1.AzurermEventgridEventSubscription) (*v1alpha1.AzurermEventgridEventSubscription, error)
	UpdateStatus(*v1alpha1.AzurermEventgridEventSubscription) (*v1alpha1.AzurermEventgridEventSubscription, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermEventgridEventSubscription, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermEventgridEventSubscriptionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermEventgridEventSubscription, err error)
	AzurermEventgridEventSubscriptionExpansion
}

// azurermEventgridEventSubscriptions implements AzurermEventgridEventSubscriptionInterface
type azurermEventgridEventSubscriptions struct {
	client rest.Interface
}

// newAzurermEventgridEventSubscriptions returns a AzurermEventgridEventSubscriptions
func newAzurermEventgridEventSubscriptions(c *AzurermV1alpha1Client) *azurermEventgridEventSubscriptions {
	return &azurermEventgridEventSubscriptions{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermEventgridEventSubscription, and returns the corresponding azurermEventgridEventSubscription object, and an error if there is any.
func (c *azurermEventgridEventSubscriptions) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermEventgridEventSubscription, err error) {
	result = &v1alpha1.AzurermEventgridEventSubscription{}
	err = c.client.Get().
		Resource("azurermeventgrideventsubscriptions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermEventgridEventSubscriptions that match those selectors.
func (c *azurermEventgridEventSubscriptions) List(opts v1.ListOptions) (result *v1alpha1.AzurermEventgridEventSubscriptionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermEventgridEventSubscriptionList{}
	err = c.client.Get().
		Resource("azurermeventgrideventsubscriptions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermEventgridEventSubscriptions.
func (c *azurermEventgridEventSubscriptions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermeventgrideventsubscriptions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermEventgridEventSubscription and creates it.  Returns the server's representation of the azurermEventgridEventSubscription, and an error, if there is any.
func (c *azurermEventgridEventSubscriptions) Create(azurermEventgridEventSubscription *v1alpha1.AzurermEventgridEventSubscription) (result *v1alpha1.AzurermEventgridEventSubscription, err error) {
	result = &v1alpha1.AzurermEventgridEventSubscription{}
	err = c.client.Post().
		Resource("azurermeventgrideventsubscriptions").
		Body(azurermEventgridEventSubscription).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermEventgridEventSubscription and updates it. Returns the server's representation of the azurermEventgridEventSubscription, and an error, if there is any.
func (c *azurermEventgridEventSubscriptions) Update(azurermEventgridEventSubscription *v1alpha1.AzurermEventgridEventSubscription) (result *v1alpha1.AzurermEventgridEventSubscription, err error) {
	result = &v1alpha1.AzurermEventgridEventSubscription{}
	err = c.client.Put().
		Resource("azurermeventgrideventsubscriptions").
		Name(azurermEventgridEventSubscription.Name).
		Body(azurermEventgridEventSubscription).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermEventgridEventSubscriptions) UpdateStatus(azurermEventgridEventSubscription *v1alpha1.AzurermEventgridEventSubscription) (result *v1alpha1.AzurermEventgridEventSubscription, err error) {
	result = &v1alpha1.AzurermEventgridEventSubscription{}
	err = c.client.Put().
		Resource("azurermeventgrideventsubscriptions").
		Name(azurermEventgridEventSubscription.Name).
		SubResource("status").
		Body(azurermEventgridEventSubscription).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermEventgridEventSubscription and deletes it. Returns an error if one occurs.
func (c *azurermEventgridEventSubscriptions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermeventgrideventsubscriptions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermEventgridEventSubscriptions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermeventgrideventsubscriptions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermEventgridEventSubscription.
func (c *azurermEventgridEventSubscriptions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermEventgridEventSubscription, err error) {
	result = &v1alpha1.AzurermEventgridEventSubscription{}
	err = c.client.Patch(pt).
		Resource("azurermeventgrideventsubscriptions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
