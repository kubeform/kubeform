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

// AzurermFirewallApplicationRuleCollectionsGetter has a method to return a AzurermFirewallApplicationRuleCollectionInterface.
// A group's client should implement this interface.
type AzurermFirewallApplicationRuleCollectionsGetter interface {
	AzurermFirewallApplicationRuleCollections() AzurermFirewallApplicationRuleCollectionInterface
}

// AzurermFirewallApplicationRuleCollectionInterface has methods to work with AzurermFirewallApplicationRuleCollection resources.
type AzurermFirewallApplicationRuleCollectionInterface interface {
	Create(*v1alpha1.AzurermFirewallApplicationRuleCollection) (*v1alpha1.AzurermFirewallApplicationRuleCollection, error)
	Update(*v1alpha1.AzurermFirewallApplicationRuleCollection) (*v1alpha1.AzurermFirewallApplicationRuleCollection, error)
	UpdateStatus(*v1alpha1.AzurermFirewallApplicationRuleCollection) (*v1alpha1.AzurermFirewallApplicationRuleCollection, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermFirewallApplicationRuleCollection, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermFirewallApplicationRuleCollectionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermFirewallApplicationRuleCollection, err error)
	AzurermFirewallApplicationRuleCollectionExpansion
}

// azurermFirewallApplicationRuleCollections implements AzurermFirewallApplicationRuleCollectionInterface
type azurermFirewallApplicationRuleCollections struct {
	client rest.Interface
}

// newAzurermFirewallApplicationRuleCollections returns a AzurermFirewallApplicationRuleCollections
func newAzurermFirewallApplicationRuleCollections(c *AzurermV1alpha1Client) *azurermFirewallApplicationRuleCollections {
	return &azurermFirewallApplicationRuleCollections{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermFirewallApplicationRuleCollection, and returns the corresponding azurermFirewallApplicationRuleCollection object, and an error if there is any.
func (c *azurermFirewallApplicationRuleCollections) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermFirewallApplicationRuleCollection, err error) {
	result = &v1alpha1.AzurermFirewallApplicationRuleCollection{}
	err = c.client.Get().
		Resource("azurermfirewallapplicationrulecollections").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermFirewallApplicationRuleCollections that match those selectors.
func (c *azurermFirewallApplicationRuleCollections) List(opts v1.ListOptions) (result *v1alpha1.AzurermFirewallApplicationRuleCollectionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermFirewallApplicationRuleCollectionList{}
	err = c.client.Get().
		Resource("azurermfirewallapplicationrulecollections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermFirewallApplicationRuleCollections.
func (c *azurermFirewallApplicationRuleCollections) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermfirewallapplicationrulecollections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermFirewallApplicationRuleCollection and creates it.  Returns the server's representation of the azurermFirewallApplicationRuleCollection, and an error, if there is any.
func (c *azurermFirewallApplicationRuleCollections) Create(azurermFirewallApplicationRuleCollection *v1alpha1.AzurermFirewallApplicationRuleCollection) (result *v1alpha1.AzurermFirewallApplicationRuleCollection, err error) {
	result = &v1alpha1.AzurermFirewallApplicationRuleCollection{}
	err = c.client.Post().
		Resource("azurermfirewallapplicationrulecollections").
		Body(azurermFirewallApplicationRuleCollection).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermFirewallApplicationRuleCollection and updates it. Returns the server's representation of the azurermFirewallApplicationRuleCollection, and an error, if there is any.
func (c *azurermFirewallApplicationRuleCollections) Update(azurermFirewallApplicationRuleCollection *v1alpha1.AzurermFirewallApplicationRuleCollection) (result *v1alpha1.AzurermFirewallApplicationRuleCollection, err error) {
	result = &v1alpha1.AzurermFirewallApplicationRuleCollection{}
	err = c.client.Put().
		Resource("azurermfirewallapplicationrulecollections").
		Name(azurermFirewallApplicationRuleCollection.Name).
		Body(azurermFirewallApplicationRuleCollection).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermFirewallApplicationRuleCollections) UpdateStatus(azurermFirewallApplicationRuleCollection *v1alpha1.AzurermFirewallApplicationRuleCollection) (result *v1alpha1.AzurermFirewallApplicationRuleCollection, err error) {
	result = &v1alpha1.AzurermFirewallApplicationRuleCollection{}
	err = c.client.Put().
		Resource("azurermfirewallapplicationrulecollections").
		Name(azurermFirewallApplicationRuleCollection.Name).
		SubResource("status").
		Body(azurermFirewallApplicationRuleCollection).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermFirewallApplicationRuleCollection and deletes it. Returns an error if one occurs.
func (c *azurermFirewallApplicationRuleCollections) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermfirewallapplicationrulecollections").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermFirewallApplicationRuleCollections) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermfirewallapplicationrulecollections").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermFirewallApplicationRuleCollection.
func (c *azurermFirewallApplicationRuleCollections) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermFirewallApplicationRuleCollection, err error) {
	result = &v1alpha1.AzurermFirewallApplicationRuleCollection{}
	err = c.client.Patch(pt).
		Resource("azurermfirewallapplicationrulecollections").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
