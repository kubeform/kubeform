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

// AzurermEventhubAuthorizationRulesGetter has a method to return a AzurermEventhubAuthorizationRuleInterface.
// A group's client should implement this interface.
type AzurermEventhubAuthorizationRulesGetter interface {
	AzurermEventhubAuthorizationRules() AzurermEventhubAuthorizationRuleInterface
}

// AzurermEventhubAuthorizationRuleInterface has methods to work with AzurermEventhubAuthorizationRule resources.
type AzurermEventhubAuthorizationRuleInterface interface {
	Create(*v1alpha1.AzurermEventhubAuthorizationRule) (*v1alpha1.AzurermEventhubAuthorizationRule, error)
	Update(*v1alpha1.AzurermEventhubAuthorizationRule) (*v1alpha1.AzurermEventhubAuthorizationRule, error)
	UpdateStatus(*v1alpha1.AzurermEventhubAuthorizationRule) (*v1alpha1.AzurermEventhubAuthorizationRule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermEventhubAuthorizationRule, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermEventhubAuthorizationRuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermEventhubAuthorizationRule, err error)
	AzurermEventhubAuthorizationRuleExpansion
}

// azurermEventhubAuthorizationRules implements AzurermEventhubAuthorizationRuleInterface
type azurermEventhubAuthorizationRules struct {
	client rest.Interface
}

// newAzurermEventhubAuthorizationRules returns a AzurermEventhubAuthorizationRules
func newAzurermEventhubAuthorizationRules(c *AzurermV1alpha1Client) *azurermEventhubAuthorizationRules {
	return &azurermEventhubAuthorizationRules{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermEventhubAuthorizationRule, and returns the corresponding azurermEventhubAuthorizationRule object, and an error if there is any.
func (c *azurermEventhubAuthorizationRules) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermEventhubAuthorizationRule, err error) {
	result = &v1alpha1.AzurermEventhubAuthorizationRule{}
	err = c.client.Get().
		Resource("azurermeventhubauthorizationrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermEventhubAuthorizationRules that match those selectors.
func (c *azurermEventhubAuthorizationRules) List(opts v1.ListOptions) (result *v1alpha1.AzurermEventhubAuthorizationRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermEventhubAuthorizationRuleList{}
	err = c.client.Get().
		Resource("azurermeventhubauthorizationrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermEventhubAuthorizationRules.
func (c *azurermEventhubAuthorizationRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermeventhubauthorizationrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermEventhubAuthorizationRule and creates it.  Returns the server's representation of the azurermEventhubAuthorizationRule, and an error, if there is any.
func (c *azurermEventhubAuthorizationRules) Create(azurermEventhubAuthorizationRule *v1alpha1.AzurermEventhubAuthorizationRule) (result *v1alpha1.AzurermEventhubAuthorizationRule, err error) {
	result = &v1alpha1.AzurermEventhubAuthorizationRule{}
	err = c.client.Post().
		Resource("azurermeventhubauthorizationrules").
		Body(azurermEventhubAuthorizationRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermEventhubAuthorizationRule and updates it. Returns the server's representation of the azurermEventhubAuthorizationRule, and an error, if there is any.
func (c *azurermEventhubAuthorizationRules) Update(azurermEventhubAuthorizationRule *v1alpha1.AzurermEventhubAuthorizationRule) (result *v1alpha1.AzurermEventhubAuthorizationRule, err error) {
	result = &v1alpha1.AzurermEventhubAuthorizationRule{}
	err = c.client.Put().
		Resource("azurermeventhubauthorizationrules").
		Name(azurermEventhubAuthorizationRule.Name).
		Body(azurermEventhubAuthorizationRule).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermEventhubAuthorizationRules) UpdateStatus(azurermEventhubAuthorizationRule *v1alpha1.AzurermEventhubAuthorizationRule) (result *v1alpha1.AzurermEventhubAuthorizationRule, err error) {
	result = &v1alpha1.AzurermEventhubAuthorizationRule{}
	err = c.client.Put().
		Resource("azurermeventhubauthorizationrules").
		Name(azurermEventhubAuthorizationRule.Name).
		SubResource("status").
		Body(azurermEventhubAuthorizationRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermEventhubAuthorizationRule and deletes it. Returns an error if one occurs.
func (c *azurermEventhubAuthorizationRules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermeventhubauthorizationrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermEventhubAuthorizationRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermeventhubauthorizationrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermEventhubAuthorizationRule.
func (c *azurermEventhubAuthorizationRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermEventhubAuthorizationRule, err error) {
	result = &v1alpha1.AzurermEventhubAuthorizationRule{}
	err = c.client.Patch(pt).
		Resource("azurermeventhubauthorizationrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
