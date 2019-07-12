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

// AzurermNetworkProfilesGetter has a method to return a AzurermNetworkProfileInterface.
// A group's client should implement this interface.
type AzurermNetworkProfilesGetter interface {
	AzurermNetworkProfiles() AzurermNetworkProfileInterface
}

// AzurermNetworkProfileInterface has methods to work with AzurermNetworkProfile resources.
type AzurermNetworkProfileInterface interface {
	Create(*v1alpha1.AzurermNetworkProfile) (*v1alpha1.AzurermNetworkProfile, error)
	Update(*v1alpha1.AzurermNetworkProfile) (*v1alpha1.AzurermNetworkProfile, error)
	UpdateStatus(*v1alpha1.AzurermNetworkProfile) (*v1alpha1.AzurermNetworkProfile, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermNetworkProfile, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermNetworkProfileList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermNetworkProfile, err error)
	AzurermNetworkProfileExpansion
}

// azurermNetworkProfiles implements AzurermNetworkProfileInterface
type azurermNetworkProfiles struct {
	client rest.Interface
}

// newAzurermNetworkProfiles returns a AzurermNetworkProfiles
func newAzurermNetworkProfiles(c *AzurermV1alpha1Client) *azurermNetworkProfiles {
	return &azurermNetworkProfiles{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermNetworkProfile, and returns the corresponding azurermNetworkProfile object, and an error if there is any.
func (c *azurermNetworkProfiles) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermNetworkProfile, err error) {
	result = &v1alpha1.AzurermNetworkProfile{}
	err = c.client.Get().
		Resource("azurermnetworkprofiles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermNetworkProfiles that match those selectors.
func (c *azurermNetworkProfiles) List(opts v1.ListOptions) (result *v1alpha1.AzurermNetworkProfileList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermNetworkProfileList{}
	err = c.client.Get().
		Resource("azurermnetworkprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermNetworkProfiles.
func (c *azurermNetworkProfiles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermnetworkprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermNetworkProfile and creates it.  Returns the server's representation of the azurermNetworkProfile, and an error, if there is any.
func (c *azurermNetworkProfiles) Create(azurermNetworkProfile *v1alpha1.AzurermNetworkProfile) (result *v1alpha1.AzurermNetworkProfile, err error) {
	result = &v1alpha1.AzurermNetworkProfile{}
	err = c.client.Post().
		Resource("azurermnetworkprofiles").
		Body(azurermNetworkProfile).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermNetworkProfile and updates it. Returns the server's representation of the azurermNetworkProfile, and an error, if there is any.
func (c *azurermNetworkProfiles) Update(azurermNetworkProfile *v1alpha1.AzurermNetworkProfile) (result *v1alpha1.AzurermNetworkProfile, err error) {
	result = &v1alpha1.AzurermNetworkProfile{}
	err = c.client.Put().
		Resource("azurermnetworkprofiles").
		Name(azurermNetworkProfile.Name).
		Body(azurermNetworkProfile).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermNetworkProfiles) UpdateStatus(azurermNetworkProfile *v1alpha1.AzurermNetworkProfile) (result *v1alpha1.AzurermNetworkProfile, err error) {
	result = &v1alpha1.AzurermNetworkProfile{}
	err = c.client.Put().
		Resource("azurermnetworkprofiles").
		Name(azurermNetworkProfile.Name).
		SubResource("status").
		Body(azurermNetworkProfile).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermNetworkProfile and deletes it. Returns an error if one occurs.
func (c *azurermNetworkProfiles) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermnetworkprofiles").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermNetworkProfiles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermnetworkprofiles").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermNetworkProfile.
func (c *azurermNetworkProfiles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermNetworkProfile, err error) {
	result = &v1alpha1.AzurermNetworkProfile{}
	err = c.client.Patch(pt).
		Resource("azurermnetworkprofiles").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
