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

// AzurermNotificationHubsGetter has a method to return a AzurermNotificationHubInterface.
// A group's client should implement this interface.
type AzurermNotificationHubsGetter interface {
	AzurermNotificationHubs() AzurermNotificationHubInterface
}

// AzurermNotificationHubInterface has methods to work with AzurermNotificationHub resources.
type AzurermNotificationHubInterface interface {
	Create(*v1alpha1.AzurermNotificationHub) (*v1alpha1.AzurermNotificationHub, error)
	Update(*v1alpha1.AzurermNotificationHub) (*v1alpha1.AzurermNotificationHub, error)
	UpdateStatus(*v1alpha1.AzurermNotificationHub) (*v1alpha1.AzurermNotificationHub, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermNotificationHub, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermNotificationHubList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermNotificationHub, err error)
	AzurermNotificationHubExpansion
}

// azurermNotificationHubs implements AzurermNotificationHubInterface
type azurermNotificationHubs struct {
	client rest.Interface
}

// newAzurermNotificationHubs returns a AzurermNotificationHubs
func newAzurermNotificationHubs(c *AzurermV1alpha1Client) *azurermNotificationHubs {
	return &azurermNotificationHubs{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermNotificationHub, and returns the corresponding azurermNotificationHub object, and an error if there is any.
func (c *azurermNotificationHubs) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermNotificationHub, err error) {
	result = &v1alpha1.AzurermNotificationHub{}
	err = c.client.Get().
		Resource("azurermnotificationhubs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermNotificationHubs that match those selectors.
func (c *azurermNotificationHubs) List(opts v1.ListOptions) (result *v1alpha1.AzurermNotificationHubList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermNotificationHubList{}
	err = c.client.Get().
		Resource("azurermnotificationhubs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermNotificationHubs.
func (c *azurermNotificationHubs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermnotificationhubs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermNotificationHub and creates it.  Returns the server's representation of the azurermNotificationHub, and an error, if there is any.
func (c *azurermNotificationHubs) Create(azurermNotificationHub *v1alpha1.AzurermNotificationHub) (result *v1alpha1.AzurermNotificationHub, err error) {
	result = &v1alpha1.AzurermNotificationHub{}
	err = c.client.Post().
		Resource("azurermnotificationhubs").
		Body(azurermNotificationHub).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermNotificationHub and updates it. Returns the server's representation of the azurermNotificationHub, and an error, if there is any.
func (c *azurermNotificationHubs) Update(azurermNotificationHub *v1alpha1.AzurermNotificationHub) (result *v1alpha1.AzurermNotificationHub, err error) {
	result = &v1alpha1.AzurermNotificationHub{}
	err = c.client.Put().
		Resource("azurermnotificationhubs").
		Name(azurermNotificationHub.Name).
		Body(azurermNotificationHub).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermNotificationHubs) UpdateStatus(azurermNotificationHub *v1alpha1.AzurermNotificationHub) (result *v1alpha1.AzurermNotificationHub, err error) {
	result = &v1alpha1.AzurermNotificationHub{}
	err = c.client.Put().
		Resource("azurermnotificationhubs").
		Name(azurermNotificationHub.Name).
		SubResource("status").
		Body(azurermNotificationHub).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermNotificationHub and deletes it. Returns an error if one occurs.
func (c *azurermNotificationHubs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermnotificationhubs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermNotificationHubs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermnotificationhubs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermNotificationHub.
func (c *azurermNotificationHubs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermNotificationHub, err error) {
	result = &v1alpha1.AzurermNotificationHub{}
	err = c.client.Patch(pt).
		Resource("azurermnotificationhubs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
