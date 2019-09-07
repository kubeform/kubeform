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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// ComposerEnvironmentsGetter has a method to return a ComposerEnvironmentInterface.
// A group's client should implement this interface.
type ComposerEnvironmentsGetter interface {
	ComposerEnvironments(namespace string) ComposerEnvironmentInterface
}

// ComposerEnvironmentInterface has methods to work with ComposerEnvironment resources.
type ComposerEnvironmentInterface interface {
	Create(*v1alpha1.ComposerEnvironment) (*v1alpha1.ComposerEnvironment, error)
	Update(*v1alpha1.ComposerEnvironment) (*v1alpha1.ComposerEnvironment, error)
	UpdateStatus(*v1alpha1.ComposerEnvironment) (*v1alpha1.ComposerEnvironment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ComposerEnvironment, error)
	List(opts v1.ListOptions) (*v1alpha1.ComposerEnvironmentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComposerEnvironment, err error)
	ComposerEnvironmentExpansion
}

// composerEnvironments implements ComposerEnvironmentInterface
type composerEnvironments struct {
	client rest.Interface
	ns     string
}

// newComposerEnvironments returns a ComposerEnvironments
func newComposerEnvironments(c *GoogleV1alpha1Client, namespace string) *composerEnvironments {
	return &composerEnvironments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the composerEnvironment, and returns the corresponding composerEnvironment object, and an error if there is any.
func (c *composerEnvironments) Get(name string, options v1.GetOptions) (result *v1alpha1.ComposerEnvironment, err error) {
	result = &v1alpha1.ComposerEnvironment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("composerenvironments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComposerEnvironments that match those selectors.
func (c *composerEnvironments) List(opts v1.ListOptions) (result *v1alpha1.ComposerEnvironmentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ComposerEnvironmentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("composerenvironments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested composerEnvironments.
func (c *composerEnvironments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("composerenvironments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a composerEnvironment and creates it.  Returns the server's representation of the composerEnvironment, and an error, if there is any.
func (c *composerEnvironments) Create(composerEnvironment *v1alpha1.ComposerEnvironment) (result *v1alpha1.ComposerEnvironment, err error) {
	result = &v1alpha1.ComposerEnvironment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("composerenvironments").
		Body(composerEnvironment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a composerEnvironment and updates it. Returns the server's representation of the composerEnvironment, and an error, if there is any.
func (c *composerEnvironments) Update(composerEnvironment *v1alpha1.ComposerEnvironment) (result *v1alpha1.ComposerEnvironment, err error) {
	result = &v1alpha1.ComposerEnvironment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("composerenvironments").
		Name(composerEnvironment.Name).
		Body(composerEnvironment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *composerEnvironments) UpdateStatus(composerEnvironment *v1alpha1.ComposerEnvironment) (result *v1alpha1.ComposerEnvironment, err error) {
	result = &v1alpha1.ComposerEnvironment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("composerenvironments").
		Name(composerEnvironment.Name).
		SubResource("status").
		Body(composerEnvironment).
		Do().
		Into(result)
	return
}

// Delete takes name of the composerEnvironment and deletes it. Returns an error if one occurs.
func (c *composerEnvironments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("composerenvironments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *composerEnvironments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("composerenvironments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched composerEnvironment.
func (c *composerEnvironments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComposerEnvironment, err error) {
	result = &v1alpha1.ComposerEnvironment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("composerenvironments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}