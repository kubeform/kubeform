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

// GoogleProjectIamCustomRolesGetter has a method to return a GoogleProjectIamCustomRoleInterface.
// A group's client should implement this interface.
type GoogleProjectIamCustomRolesGetter interface {
	GoogleProjectIamCustomRoles() GoogleProjectIamCustomRoleInterface
}

// GoogleProjectIamCustomRoleInterface has methods to work with GoogleProjectIamCustomRole resources.
type GoogleProjectIamCustomRoleInterface interface {
	Create(*v1alpha1.GoogleProjectIamCustomRole) (*v1alpha1.GoogleProjectIamCustomRole, error)
	Update(*v1alpha1.GoogleProjectIamCustomRole) (*v1alpha1.GoogleProjectIamCustomRole, error)
	UpdateStatus(*v1alpha1.GoogleProjectIamCustomRole) (*v1alpha1.GoogleProjectIamCustomRole, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GoogleProjectIamCustomRole, error)
	List(opts v1.ListOptions) (*v1alpha1.GoogleProjectIamCustomRoleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleProjectIamCustomRole, err error)
	GoogleProjectIamCustomRoleExpansion
}

// googleProjectIamCustomRoles implements GoogleProjectIamCustomRoleInterface
type googleProjectIamCustomRoles struct {
	client rest.Interface
}

// newGoogleProjectIamCustomRoles returns a GoogleProjectIamCustomRoles
func newGoogleProjectIamCustomRoles(c *GoogleV1alpha1Client) *googleProjectIamCustomRoles {
	return &googleProjectIamCustomRoles{
		client: c.RESTClient(),
	}
}

// Get takes name of the googleProjectIamCustomRole, and returns the corresponding googleProjectIamCustomRole object, and an error if there is any.
func (c *googleProjectIamCustomRoles) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleProjectIamCustomRole, err error) {
	result = &v1alpha1.GoogleProjectIamCustomRole{}
	err = c.client.Get().
		Resource("googleprojectiamcustomroles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GoogleProjectIamCustomRoles that match those selectors.
func (c *googleProjectIamCustomRoles) List(opts v1.ListOptions) (result *v1alpha1.GoogleProjectIamCustomRoleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GoogleProjectIamCustomRoleList{}
	err = c.client.Get().
		Resource("googleprojectiamcustomroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested googleProjectIamCustomRoles.
func (c *googleProjectIamCustomRoles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("googleprojectiamcustomroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a googleProjectIamCustomRole and creates it.  Returns the server's representation of the googleProjectIamCustomRole, and an error, if there is any.
func (c *googleProjectIamCustomRoles) Create(googleProjectIamCustomRole *v1alpha1.GoogleProjectIamCustomRole) (result *v1alpha1.GoogleProjectIamCustomRole, err error) {
	result = &v1alpha1.GoogleProjectIamCustomRole{}
	err = c.client.Post().
		Resource("googleprojectiamcustomroles").
		Body(googleProjectIamCustomRole).
		Do().
		Into(result)
	return
}

// Update takes the representation of a googleProjectIamCustomRole and updates it. Returns the server's representation of the googleProjectIamCustomRole, and an error, if there is any.
func (c *googleProjectIamCustomRoles) Update(googleProjectIamCustomRole *v1alpha1.GoogleProjectIamCustomRole) (result *v1alpha1.GoogleProjectIamCustomRole, err error) {
	result = &v1alpha1.GoogleProjectIamCustomRole{}
	err = c.client.Put().
		Resource("googleprojectiamcustomroles").
		Name(googleProjectIamCustomRole.Name).
		Body(googleProjectIamCustomRole).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *googleProjectIamCustomRoles) UpdateStatus(googleProjectIamCustomRole *v1alpha1.GoogleProjectIamCustomRole) (result *v1alpha1.GoogleProjectIamCustomRole, err error) {
	result = &v1alpha1.GoogleProjectIamCustomRole{}
	err = c.client.Put().
		Resource("googleprojectiamcustomroles").
		Name(googleProjectIamCustomRole.Name).
		SubResource("status").
		Body(googleProjectIamCustomRole).
		Do().
		Into(result)
	return
}

// Delete takes name of the googleProjectIamCustomRole and deletes it. Returns an error if one occurs.
func (c *googleProjectIamCustomRoles) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("googleprojectiamcustomroles").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *googleProjectIamCustomRoles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("googleprojectiamcustomroles").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched googleProjectIamCustomRole.
func (c *googleProjectIamCustomRoles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleProjectIamCustomRole, err error) {
	result = &v1alpha1.GoogleProjectIamCustomRole{}
	err = c.client.Patch(pt).
		Resource("googleprojectiamcustomroles").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
