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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// GoogleStorageBucketIamBindingsGetter has a method to return a GoogleStorageBucketIamBindingInterface.
// A group's client should implement this interface.
type GoogleStorageBucketIamBindingsGetter interface {
	GoogleStorageBucketIamBindings() GoogleStorageBucketIamBindingInterface
}

// GoogleStorageBucketIamBindingInterface has methods to work with GoogleStorageBucketIamBinding resources.
type GoogleStorageBucketIamBindingInterface interface {
	Create(*v1alpha1.GoogleStorageBucketIamBinding) (*v1alpha1.GoogleStorageBucketIamBinding, error)
	Update(*v1alpha1.GoogleStorageBucketIamBinding) (*v1alpha1.GoogleStorageBucketIamBinding, error)
	UpdateStatus(*v1alpha1.GoogleStorageBucketIamBinding) (*v1alpha1.GoogleStorageBucketIamBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GoogleStorageBucketIamBinding, error)
	List(opts v1.ListOptions) (*v1alpha1.GoogleStorageBucketIamBindingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleStorageBucketIamBinding, err error)
	GoogleStorageBucketIamBindingExpansion
}

// googleStorageBucketIamBindings implements GoogleStorageBucketIamBindingInterface
type googleStorageBucketIamBindings struct {
	client rest.Interface
}

// newGoogleStorageBucketIamBindings returns a GoogleStorageBucketIamBindings
func newGoogleStorageBucketIamBindings(c *GoogleV1alpha1Client) *googleStorageBucketIamBindings {
	return &googleStorageBucketIamBindings{
		client: c.RESTClient(),
	}
}

// Get takes name of the googleStorageBucketIamBinding, and returns the corresponding googleStorageBucketIamBinding object, and an error if there is any.
func (c *googleStorageBucketIamBindings) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleStorageBucketIamBinding, err error) {
	result = &v1alpha1.GoogleStorageBucketIamBinding{}
	err = c.client.Get().
		Resource("googlestoragebucketiambindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GoogleStorageBucketIamBindings that match those selectors.
func (c *googleStorageBucketIamBindings) List(opts v1.ListOptions) (result *v1alpha1.GoogleStorageBucketIamBindingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GoogleStorageBucketIamBindingList{}
	err = c.client.Get().
		Resource("googlestoragebucketiambindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested googleStorageBucketIamBindings.
func (c *googleStorageBucketIamBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("googlestoragebucketiambindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a googleStorageBucketIamBinding and creates it.  Returns the server's representation of the googleStorageBucketIamBinding, and an error, if there is any.
func (c *googleStorageBucketIamBindings) Create(googleStorageBucketIamBinding *v1alpha1.GoogleStorageBucketIamBinding) (result *v1alpha1.GoogleStorageBucketIamBinding, err error) {
	result = &v1alpha1.GoogleStorageBucketIamBinding{}
	err = c.client.Post().
		Resource("googlestoragebucketiambindings").
		Body(googleStorageBucketIamBinding).
		Do().
		Into(result)
	return
}

// Update takes the representation of a googleStorageBucketIamBinding and updates it. Returns the server's representation of the googleStorageBucketIamBinding, and an error, if there is any.
func (c *googleStorageBucketIamBindings) Update(googleStorageBucketIamBinding *v1alpha1.GoogleStorageBucketIamBinding) (result *v1alpha1.GoogleStorageBucketIamBinding, err error) {
	result = &v1alpha1.GoogleStorageBucketIamBinding{}
	err = c.client.Put().
		Resource("googlestoragebucketiambindings").
		Name(googleStorageBucketIamBinding.Name).
		Body(googleStorageBucketIamBinding).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *googleStorageBucketIamBindings) UpdateStatus(googleStorageBucketIamBinding *v1alpha1.GoogleStorageBucketIamBinding) (result *v1alpha1.GoogleStorageBucketIamBinding, err error) {
	result = &v1alpha1.GoogleStorageBucketIamBinding{}
	err = c.client.Put().
		Resource("googlestoragebucketiambindings").
		Name(googleStorageBucketIamBinding.Name).
		SubResource("status").
		Body(googleStorageBucketIamBinding).
		Do().
		Into(result)
	return
}

// Delete takes name of the googleStorageBucketIamBinding and deletes it. Returns an error if one occurs.
func (c *googleStorageBucketIamBindings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("googlestoragebucketiambindings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *googleStorageBucketIamBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("googlestoragebucketiambindings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched googleStorageBucketIamBinding.
func (c *googleStorageBucketIamBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleStorageBucketIamBinding, err error) {
	result = &v1alpha1.GoogleStorageBucketIamBinding{}
	err = c.client.Patch(pt).
		Resource("googlestoragebucketiambindings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
