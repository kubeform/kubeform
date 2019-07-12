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

// AzurermEventgridTopicsGetter has a method to return a AzurermEventgridTopicInterface.
// A group's client should implement this interface.
type AzurermEventgridTopicsGetter interface {
	AzurermEventgridTopics() AzurermEventgridTopicInterface
}

// AzurermEventgridTopicInterface has methods to work with AzurermEventgridTopic resources.
type AzurermEventgridTopicInterface interface {
	Create(*v1alpha1.AzurermEventgridTopic) (*v1alpha1.AzurermEventgridTopic, error)
	Update(*v1alpha1.AzurermEventgridTopic) (*v1alpha1.AzurermEventgridTopic, error)
	UpdateStatus(*v1alpha1.AzurermEventgridTopic) (*v1alpha1.AzurermEventgridTopic, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermEventgridTopic, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermEventgridTopicList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermEventgridTopic, err error)
	AzurermEventgridTopicExpansion
}

// azurermEventgridTopics implements AzurermEventgridTopicInterface
type azurermEventgridTopics struct {
	client rest.Interface
}

// newAzurermEventgridTopics returns a AzurermEventgridTopics
func newAzurermEventgridTopics(c *AzurermV1alpha1Client) *azurermEventgridTopics {
	return &azurermEventgridTopics{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermEventgridTopic, and returns the corresponding azurermEventgridTopic object, and an error if there is any.
func (c *azurermEventgridTopics) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermEventgridTopic, err error) {
	result = &v1alpha1.AzurermEventgridTopic{}
	err = c.client.Get().
		Resource("azurermeventgridtopics").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermEventgridTopics that match those selectors.
func (c *azurermEventgridTopics) List(opts v1.ListOptions) (result *v1alpha1.AzurermEventgridTopicList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermEventgridTopicList{}
	err = c.client.Get().
		Resource("azurermeventgridtopics").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermEventgridTopics.
func (c *azurermEventgridTopics) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermeventgridtopics").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermEventgridTopic and creates it.  Returns the server's representation of the azurermEventgridTopic, and an error, if there is any.
func (c *azurermEventgridTopics) Create(azurermEventgridTopic *v1alpha1.AzurermEventgridTopic) (result *v1alpha1.AzurermEventgridTopic, err error) {
	result = &v1alpha1.AzurermEventgridTopic{}
	err = c.client.Post().
		Resource("azurermeventgridtopics").
		Body(azurermEventgridTopic).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermEventgridTopic and updates it. Returns the server's representation of the azurermEventgridTopic, and an error, if there is any.
func (c *azurermEventgridTopics) Update(azurermEventgridTopic *v1alpha1.AzurermEventgridTopic) (result *v1alpha1.AzurermEventgridTopic, err error) {
	result = &v1alpha1.AzurermEventgridTopic{}
	err = c.client.Put().
		Resource("azurermeventgridtopics").
		Name(azurermEventgridTopic.Name).
		Body(azurermEventgridTopic).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermEventgridTopics) UpdateStatus(azurermEventgridTopic *v1alpha1.AzurermEventgridTopic) (result *v1alpha1.AzurermEventgridTopic, err error) {
	result = &v1alpha1.AzurermEventgridTopic{}
	err = c.client.Put().
		Resource("azurermeventgridtopics").
		Name(azurermEventgridTopic.Name).
		SubResource("status").
		Body(azurermEventgridTopic).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermEventgridTopic and deletes it. Returns an error if one occurs.
func (c *azurermEventgridTopics) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermeventgridtopics").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermEventgridTopics) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermeventgridtopics").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermEventgridTopic.
func (c *azurermEventgridTopics) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermEventgridTopic, err error) {
	result = &v1alpha1.AzurermEventgridTopic{}
	err = c.client.Patch(pt).
		Resource("azurermeventgridtopics").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
