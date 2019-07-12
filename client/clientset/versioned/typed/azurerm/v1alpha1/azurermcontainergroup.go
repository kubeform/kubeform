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

// AzurermContainerGroupsGetter has a method to return a AzurermContainerGroupInterface.
// A group's client should implement this interface.
type AzurermContainerGroupsGetter interface {
	AzurermContainerGroups() AzurermContainerGroupInterface
}

// AzurermContainerGroupInterface has methods to work with AzurermContainerGroup resources.
type AzurermContainerGroupInterface interface {
	Create(*v1alpha1.AzurermContainerGroup) (*v1alpha1.AzurermContainerGroup, error)
	Update(*v1alpha1.AzurermContainerGroup) (*v1alpha1.AzurermContainerGroup, error)
	UpdateStatus(*v1alpha1.AzurermContainerGroup) (*v1alpha1.AzurermContainerGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermContainerGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermContainerGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermContainerGroup, err error)
	AzurermContainerGroupExpansion
}

// azurermContainerGroups implements AzurermContainerGroupInterface
type azurermContainerGroups struct {
	client rest.Interface
}

// newAzurermContainerGroups returns a AzurermContainerGroups
func newAzurermContainerGroups(c *AzurermV1alpha1Client) *azurermContainerGroups {
	return &azurermContainerGroups{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermContainerGroup, and returns the corresponding azurermContainerGroup object, and an error if there is any.
func (c *azurermContainerGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermContainerGroup, err error) {
	result = &v1alpha1.AzurermContainerGroup{}
	err = c.client.Get().
		Resource("azurermcontainergroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermContainerGroups that match those selectors.
func (c *azurermContainerGroups) List(opts v1.ListOptions) (result *v1alpha1.AzurermContainerGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermContainerGroupList{}
	err = c.client.Get().
		Resource("azurermcontainergroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermContainerGroups.
func (c *azurermContainerGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermcontainergroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermContainerGroup and creates it.  Returns the server's representation of the azurermContainerGroup, and an error, if there is any.
func (c *azurermContainerGroups) Create(azurermContainerGroup *v1alpha1.AzurermContainerGroup) (result *v1alpha1.AzurermContainerGroup, err error) {
	result = &v1alpha1.AzurermContainerGroup{}
	err = c.client.Post().
		Resource("azurermcontainergroups").
		Body(azurermContainerGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermContainerGroup and updates it. Returns the server's representation of the azurermContainerGroup, and an error, if there is any.
func (c *azurermContainerGroups) Update(azurermContainerGroup *v1alpha1.AzurermContainerGroup) (result *v1alpha1.AzurermContainerGroup, err error) {
	result = &v1alpha1.AzurermContainerGroup{}
	err = c.client.Put().
		Resource("azurermcontainergroups").
		Name(azurermContainerGroup.Name).
		Body(azurermContainerGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermContainerGroups) UpdateStatus(azurermContainerGroup *v1alpha1.AzurermContainerGroup) (result *v1alpha1.AzurermContainerGroup, err error) {
	result = &v1alpha1.AzurermContainerGroup{}
	err = c.client.Put().
		Resource("azurermcontainergroups").
		Name(azurermContainerGroup.Name).
		SubResource("status").
		Body(azurermContainerGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermContainerGroup and deletes it. Returns an error if one occurs.
func (c *azurermContainerGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermcontainergroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermContainerGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermcontainergroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermContainerGroup.
func (c *azurermContainerGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermContainerGroup, err error) {
	result = &v1alpha1.AzurermContainerGroup{}
	err = c.client.Patch(pt).
		Resource("azurermcontainergroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
