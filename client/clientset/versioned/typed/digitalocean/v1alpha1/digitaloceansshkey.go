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
	v1alpha1 "kubeform.dev/kubeform/apis/digitalocean/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// DigitaloceanSshKeysGetter has a method to return a DigitaloceanSshKeyInterface.
// A group's client should implement this interface.
type DigitaloceanSshKeysGetter interface {
	DigitaloceanSshKeys() DigitaloceanSshKeyInterface
}

// DigitaloceanSshKeyInterface has methods to work with DigitaloceanSshKey resources.
type DigitaloceanSshKeyInterface interface {
	Create(*v1alpha1.DigitaloceanSshKey) (*v1alpha1.DigitaloceanSshKey, error)
	Update(*v1alpha1.DigitaloceanSshKey) (*v1alpha1.DigitaloceanSshKey, error)
	UpdateStatus(*v1alpha1.DigitaloceanSshKey) (*v1alpha1.DigitaloceanSshKey, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DigitaloceanSshKey, error)
	List(opts v1.ListOptions) (*v1alpha1.DigitaloceanSshKeyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DigitaloceanSshKey, err error)
	DigitaloceanSshKeyExpansion
}

// digitaloceanSshKeys implements DigitaloceanSshKeyInterface
type digitaloceanSshKeys struct {
	client rest.Interface
}

// newDigitaloceanSshKeys returns a DigitaloceanSshKeys
func newDigitaloceanSshKeys(c *DigitaloceanV1alpha1Client) *digitaloceanSshKeys {
	return &digitaloceanSshKeys{
		client: c.RESTClient(),
	}
}

// Get takes name of the digitaloceanSshKey, and returns the corresponding digitaloceanSshKey object, and an error if there is any.
func (c *digitaloceanSshKeys) Get(name string, options v1.GetOptions) (result *v1alpha1.DigitaloceanSshKey, err error) {
	result = &v1alpha1.DigitaloceanSshKey{}
	err = c.client.Get().
		Resource("digitaloceansshkeys").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DigitaloceanSshKeys that match those selectors.
func (c *digitaloceanSshKeys) List(opts v1.ListOptions) (result *v1alpha1.DigitaloceanSshKeyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DigitaloceanSshKeyList{}
	err = c.client.Get().
		Resource("digitaloceansshkeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested digitaloceanSshKeys.
func (c *digitaloceanSshKeys) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("digitaloceansshkeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a digitaloceanSshKey and creates it.  Returns the server's representation of the digitaloceanSshKey, and an error, if there is any.
func (c *digitaloceanSshKeys) Create(digitaloceanSshKey *v1alpha1.DigitaloceanSshKey) (result *v1alpha1.DigitaloceanSshKey, err error) {
	result = &v1alpha1.DigitaloceanSshKey{}
	err = c.client.Post().
		Resource("digitaloceansshkeys").
		Body(digitaloceanSshKey).
		Do().
		Into(result)
	return
}

// Update takes the representation of a digitaloceanSshKey and updates it. Returns the server's representation of the digitaloceanSshKey, and an error, if there is any.
func (c *digitaloceanSshKeys) Update(digitaloceanSshKey *v1alpha1.DigitaloceanSshKey) (result *v1alpha1.DigitaloceanSshKey, err error) {
	result = &v1alpha1.DigitaloceanSshKey{}
	err = c.client.Put().
		Resource("digitaloceansshkeys").
		Name(digitaloceanSshKey.Name).
		Body(digitaloceanSshKey).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *digitaloceanSshKeys) UpdateStatus(digitaloceanSshKey *v1alpha1.DigitaloceanSshKey) (result *v1alpha1.DigitaloceanSshKey, err error) {
	result = &v1alpha1.DigitaloceanSshKey{}
	err = c.client.Put().
		Resource("digitaloceansshkeys").
		Name(digitaloceanSshKey.Name).
		SubResource("status").
		Body(digitaloceanSshKey).
		Do().
		Into(result)
	return
}

// Delete takes name of the digitaloceanSshKey and deletes it. Returns an error if one occurs.
func (c *digitaloceanSshKeys) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("digitaloceansshkeys").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *digitaloceanSshKeys) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("digitaloceansshkeys").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched digitaloceanSshKey.
func (c *digitaloceanSshKeys) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DigitaloceanSshKey, err error) {
	result = &v1alpha1.DigitaloceanSshKey{}
	err = c.client.Patch(pt).
		Resource("digitaloceansshkeys").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
