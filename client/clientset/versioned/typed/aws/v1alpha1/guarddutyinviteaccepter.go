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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// GuarddutyInviteAcceptersGetter has a method to return a GuarddutyInviteAccepterInterface.
// A group's client should implement this interface.
type GuarddutyInviteAcceptersGetter interface {
	GuarddutyInviteAccepters() GuarddutyInviteAccepterInterface
}

// GuarddutyInviteAccepterInterface has methods to work with GuarddutyInviteAccepter resources.
type GuarddutyInviteAccepterInterface interface {
	Create(*v1alpha1.GuarddutyInviteAccepter) (*v1alpha1.GuarddutyInviteAccepter, error)
	Update(*v1alpha1.GuarddutyInviteAccepter) (*v1alpha1.GuarddutyInviteAccepter, error)
	UpdateStatus(*v1alpha1.GuarddutyInviteAccepter) (*v1alpha1.GuarddutyInviteAccepter, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GuarddutyInviteAccepter, error)
	List(opts v1.ListOptions) (*v1alpha1.GuarddutyInviteAccepterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GuarddutyInviteAccepter, err error)
	GuarddutyInviteAccepterExpansion
}

// guarddutyInviteAccepters implements GuarddutyInviteAccepterInterface
type guarddutyInviteAccepters struct {
	client rest.Interface
}

// newGuarddutyInviteAccepters returns a GuarddutyInviteAccepters
func newGuarddutyInviteAccepters(c *AwsV1alpha1Client) *guarddutyInviteAccepters {
	return &guarddutyInviteAccepters{
		client: c.RESTClient(),
	}
}

// Get takes name of the guarddutyInviteAccepter, and returns the corresponding guarddutyInviteAccepter object, and an error if there is any.
func (c *guarddutyInviteAccepters) Get(name string, options v1.GetOptions) (result *v1alpha1.GuarddutyInviteAccepter, err error) {
	result = &v1alpha1.GuarddutyInviteAccepter{}
	err = c.client.Get().
		Resource("guarddutyinviteaccepters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GuarddutyInviteAccepters that match those selectors.
func (c *guarddutyInviteAccepters) List(opts v1.ListOptions) (result *v1alpha1.GuarddutyInviteAccepterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GuarddutyInviteAccepterList{}
	err = c.client.Get().
		Resource("guarddutyinviteaccepters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested guarddutyInviteAccepters.
func (c *guarddutyInviteAccepters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("guarddutyinviteaccepters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a guarddutyInviteAccepter and creates it.  Returns the server's representation of the guarddutyInviteAccepter, and an error, if there is any.
func (c *guarddutyInviteAccepters) Create(guarddutyInviteAccepter *v1alpha1.GuarddutyInviteAccepter) (result *v1alpha1.GuarddutyInviteAccepter, err error) {
	result = &v1alpha1.GuarddutyInviteAccepter{}
	err = c.client.Post().
		Resource("guarddutyinviteaccepters").
		Body(guarddutyInviteAccepter).
		Do().
		Into(result)
	return
}

// Update takes the representation of a guarddutyInviteAccepter and updates it. Returns the server's representation of the guarddutyInviteAccepter, and an error, if there is any.
func (c *guarddutyInviteAccepters) Update(guarddutyInviteAccepter *v1alpha1.GuarddutyInviteAccepter) (result *v1alpha1.GuarddutyInviteAccepter, err error) {
	result = &v1alpha1.GuarddutyInviteAccepter{}
	err = c.client.Put().
		Resource("guarddutyinviteaccepters").
		Name(guarddutyInviteAccepter.Name).
		Body(guarddutyInviteAccepter).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *guarddutyInviteAccepters) UpdateStatus(guarddutyInviteAccepter *v1alpha1.GuarddutyInviteAccepter) (result *v1alpha1.GuarddutyInviteAccepter, err error) {
	result = &v1alpha1.GuarddutyInviteAccepter{}
	err = c.client.Put().
		Resource("guarddutyinviteaccepters").
		Name(guarddutyInviteAccepter.Name).
		SubResource("status").
		Body(guarddutyInviteAccepter).
		Do().
		Into(result)
	return
}

// Delete takes name of the guarddutyInviteAccepter and deletes it. Returns an error if one occurs.
func (c *guarddutyInviteAccepters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("guarddutyinviteaccepters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *guarddutyInviteAccepters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("guarddutyinviteaccepters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched guarddutyInviteAccepter.
func (c *guarddutyInviteAccepters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GuarddutyInviteAccepter, err error) {
	result = &v1alpha1.GuarddutyInviteAccepter{}
	err = c.client.Patch(pt).
		Resource("guarddutyinviteaccepters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
