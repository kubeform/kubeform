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

// GoogleServiceAccountIamMembersGetter has a method to return a GoogleServiceAccountIamMemberInterface.
// A group's client should implement this interface.
type GoogleServiceAccountIamMembersGetter interface {
	GoogleServiceAccountIamMembers() GoogleServiceAccountIamMemberInterface
}

// GoogleServiceAccountIamMemberInterface has methods to work with GoogleServiceAccountIamMember resources.
type GoogleServiceAccountIamMemberInterface interface {
	Create(*v1alpha1.GoogleServiceAccountIamMember) (*v1alpha1.GoogleServiceAccountIamMember, error)
	Update(*v1alpha1.GoogleServiceAccountIamMember) (*v1alpha1.GoogleServiceAccountIamMember, error)
	UpdateStatus(*v1alpha1.GoogleServiceAccountIamMember) (*v1alpha1.GoogleServiceAccountIamMember, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GoogleServiceAccountIamMember, error)
	List(opts v1.ListOptions) (*v1alpha1.GoogleServiceAccountIamMemberList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleServiceAccountIamMember, err error)
	GoogleServiceAccountIamMemberExpansion
}

// googleServiceAccountIamMembers implements GoogleServiceAccountIamMemberInterface
type googleServiceAccountIamMembers struct {
	client rest.Interface
}

// newGoogleServiceAccountIamMembers returns a GoogleServiceAccountIamMembers
func newGoogleServiceAccountIamMembers(c *GoogleV1alpha1Client) *googleServiceAccountIamMembers {
	return &googleServiceAccountIamMembers{
		client: c.RESTClient(),
	}
}

// Get takes name of the googleServiceAccountIamMember, and returns the corresponding googleServiceAccountIamMember object, and an error if there is any.
func (c *googleServiceAccountIamMembers) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleServiceAccountIamMember, err error) {
	result = &v1alpha1.GoogleServiceAccountIamMember{}
	err = c.client.Get().
		Resource("googleserviceaccountiammembers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GoogleServiceAccountIamMembers that match those selectors.
func (c *googleServiceAccountIamMembers) List(opts v1.ListOptions) (result *v1alpha1.GoogleServiceAccountIamMemberList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GoogleServiceAccountIamMemberList{}
	err = c.client.Get().
		Resource("googleserviceaccountiammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested googleServiceAccountIamMembers.
func (c *googleServiceAccountIamMembers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("googleserviceaccountiammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a googleServiceAccountIamMember and creates it.  Returns the server's representation of the googleServiceAccountIamMember, and an error, if there is any.
func (c *googleServiceAccountIamMembers) Create(googleServiceAccountIamMember *v1alpha1.GoogleServiceAccountIamMember) (result *v1alpha1.GoogleServiceAccountIamMember, err error) {
	result = &v1alpha1.GoogleServiceAccountIamMember{}
	err = c.client.Post().
		Resource("googleserviceaccountiammembers").
		Body(googleServiceAccountIamMember).
		Do().
		Into(result)
	return
}

// Update takes the representation of a googleServiceAccountIamMember and updates it. Returns the server's representation of the googleServiceAccountIamMember, and an error, if there is any.
func (c *googleServiceAccountIamMembers) Update(googleServiceAccountIamMember *v1alpha1.GoogleServiceAccountIamMember) (result *v1alpha1.GoogleServiceAccountIamMember, err error) {
	result = &v1alpha1.GoogleServiceAccountIamMember{}
	err = c.client.Put().
		Resource("googleserviceaccountiammembers").
		Name(googleServiceAccountIamMember.Name).
		Body(googleServiceAccountIamMember).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *googleServiceAccountIamMembers) UpdateStatus(googleServiceAccountIamMember *v1alpha1.GoogleServiceAccountIamMember) (result *v1alpha1.GoogleServiceAccountIamMember, err error) {
	result = &v1alpha1.GoogleServiceAccountIamMember{}
	err = c.client.Put().
		Resource("googleserviceaccountiammembers").
		Name(googleServiceAccountIamMember.Name).
		SubResource("status").
		Body(googleServiceAccountIamMember).
		Do().
		Into(result)
	return
}

// Delete takes name of the googleServiceAccountIamMember and deletes it. Returns an error if one occurs.
func (c *googleServiceAccountIamMembers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("googleserviceaccountiammembers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *googleServiceAccountIamMembers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("googleserviceaccountiammembers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched googleServiceAccountIamMember.
func (c *googleServiceAccountIamMembers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleServiceAccountIamMember, err error) {
	result = &v1alpha1.GoogleServiceAccountIamMember{}
	err = c.client.Patch(pt).
		Resource("googleserviceaccountiammembers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
