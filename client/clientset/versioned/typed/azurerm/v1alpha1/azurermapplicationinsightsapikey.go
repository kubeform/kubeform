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

// AzurermApplicationInsightsApiKeysGetter has a method to return a AzurermApplicationInsightsApiKeyInterface.
// A group's client should implement this interface.
type AzurermApplicationInsightsApiKeysGetter interface {
	AzurermApplicationInsightsApiKeys() AzurermApplicationInsightsApiKeyInterface
}

// AzurermApplicationInsightsApiKeyInterface has methods to work with AzurermApplicationInsightsApiKey resources.
type AzurermApplicationInsightsApiKeyInterface interface {
	Create(*v1alpha1.AzurermApplicationInsightsApiKey) (*v1alpha1.AzurermApplicationInsightsApiKey, error)
	Update(*v1alpha1.AzurermApplicationInsightsApiKey) (*v1alpha1.AzurermApplicationInsightsApiKey, error)
	UpdateStatus(*v1alpha1.AzurermApplicationInsightsApiKey) (*v1alpha1.AzurermApplicationInsightsApiKey, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermApplicationInsightsApiKey, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermApplicationInsightsApiKeyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermApplicationInsightsApiKey, err error)
	AzurermApplicationInsightsApiKeyExpansion
}

// azurermApplicationInsightsApiKeys implements AzurermApplicationInsightsApiKeyInterface
type azurermApplicationInsightsApiKeys struct {
	client rest.Interface
}

// newAzurermApplicationInsightsApiKeys returns a AzurermApplicationInsightsApiKeys
func newAzurermApplicationInsightsApiKeys(c *AzurermV1alpha1Client) *azurermApplicationInsightsApiKeys {
	return &azurermApplicationInsightsApiKeys{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermApplicationInsightsApiKey, and returns the corresponding azurermApplicationInsightsApiKey object, and an error if there is any.
func (c *azurermApplicationInsightsApiKeys) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermApplicationInsightsApiKey, err error) {
	result = &v1alpha1.AzurermApplicationInsightsApiKey{}
	err = c.client.Get().
		Resource("azurermapplicationinsightsapikeys").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermApplicationInsightsApiKeys that match those selectors.
func (c *azurermApplicationInsightsApiKeys) List(opts v1.ListOptions) (result *v1alpha1.AzurermApplicationInsightsApiKeyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermApplicationInsightsApiKeyList{}
	err = c.client.Get().
		Resource("azurermapplicationinsightsapikeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermApplicationInsightsApiKeys.
func (c *azurermApplicationInsightsApiKeys) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermapplicationinsightsapikeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermApplicationInsightsApiKey and creates it.  Returns the server's representation of the azurermApplicationInsightsApiKey, and an error, if there is any.
func (c *azurermApplicationInsightsApiKeys) Create(azurermApplicationInsightsApiKey *v1alpha1.AzurermApplicationInsightsApiKey) (result *v1alpha1.AzurermApplicationInsightsApiKey, err error) {
	result = &v1alpha1.AzurermApplicationInsightsApiKey{}
	err = c.client.Post().
		Resource("azurermapplicationinsightsapikeys").
		Body(azurermApplicationInsightsApiKey).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermApplicationInsightsApiKey and updates it. Returns the server's representation of the azurermApplicationInsightsApiKey, and an error, if there is any.
func (c *azurermApplicationInsightsApiKeys) Update(azurermApplicationInsightsApiKey *v1alpha1.AzurermApplicationInsightsApiKey) (result *v1alpha1.AzurermApplicationInsightsApiKey, err error) {
	result = &v1alpha1.AzurermApplicationInsightsApiKey{}
	err = c.client.Put().
		Resource("azurermapplicationinsightsapikeys").
		Name(azurermApplicationInsightsApiKey.Name).
		Body(azurermApplicationInsightsApiKey).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermApplicationInsightsApiKeys) UpdateStatus(azurermApplicationInsightsApiKey *v1alpha1.AzurermApplicationInsightsApiKey) (result *v1alpha1.AzurermApplicationInsightsApiKey, err error) {
	result = &v1alpha1.AzurermApplicationInsightsApiKey{}
	err = c.client.Put().
		Resource("azurermapplicationinsightsapikeys").
		Name(azurermApplicationInsightsApiKey.Name).
		SubResource("status").
		Body(azurermApplicationInsightsApiKey).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermApplicationInsightsApiKey and deletes it. Returns an error if one occurs.
func (c *azurermApplicationInsightsApiKeys) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermapplicationinsightsapikeys").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermApplicationInsightsApiKeys) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermapplicationinsightsapikeys").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermApplicationInsightsApiKey.
func (c *azurermApplicationInsightsApiKeys) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermApplicationInsightsApiKey, err error) {
	result = &v1alpha1.AzurermApplicationInsightsApiKey{}
	err = c.client.Patch(pt).
		Resource("azurermapplicationinsightsapikeys").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
