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

// AzurermApiManagementGroupUsersGetter has a method to return a AzurermApiManagementGroupUserInterface.
// A group's client should implement this interface.
type AzurermApiManagementGroupUsersGetter interface {
	AzurermApiManagementGroupUsers() AzurermApiManagementGroupUserInterface
}

// AzurermApiManagementGroupUserInterface has methods to work with AzurermApiManagementGroupUser resources.
type AzurermApiManagementGroupUserInterface interface {
	Create(*v1alpha1.AzurermApiManagementGroupUser) (*v1alpha1.AzurermApiManagementGroupUser, error)
	Update(*v1alpha1.AzurermApiManagementGroupUser) (*v1alpha1.AzurermApiManagementGroupUser, error)
	UpdateStatus(*v1alpha1.AzurermApiManagementGroupUser) (*v1alpha1.AzurermApiManagementGroupUser, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermApiManagementGroupUser, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermApiManagementGroupUserList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermApiManagementGroupUser, err error)
	AzurermApiManagementGroupUserExpansion
}

// azurermApiManagementGroupUsers implements AzurermApiManagementGroupUserInterface
type azurermApiManagementGroupUsers struct {
	client rest.Interface
}

// newAzurermApiManagementGroupUsers returns a AzurermApiManagementGroupUsers
func newAzurermApiManagementGroupUsers(c *AzurermV1alpha1Client) *azurermApiManagementGroupUsers {
	return &azurermApiManagementGroupUsers{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermApiManagementGroupUser, and returns the corresponding azurermApiManagementGroupUser object, and an error if there is any.
func (c *azurermApiManagementGroupUsers) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermApiManagementGroupUser, err error) {
	result = &v1alpha1.AzurermApiManagementGroupUser{}
	err = c.client.Get().
		Resource("azurermapimanagementgroupusers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermApiManagementGroupUsers that match those selectors.
func (c *azurermApiManagementGroupUsers) List(opts v1.ListOptions) (result *v1alpha1.AzurermApiManagementGroupUserList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermApiManagementGroupUserList{}
	err = c.client.Get().
		Resource("azurermapimanagementgroupusers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermApiManagementGroupUsers.
func (c *azurermApiManagementGroupUsers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermapimanagementgroupusers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermApiManagementGroupUser and creates it.  Returns the server's representation of the azurermApiManagementGroupUser, and an error, if there is any.
func (c *azurermApiManagementGroupUsers) Create(azurermApiManagementGroupUser *v1alpha1.AzurermApiManagementGroupUser) (result *v1alpha1.AzurermApiManagementGroupUser, err error) {
	result = &v1alpha1.AzurermApiManagementGroupUser{}
	err = c.client.Post().
		Resource("azurermapimanagementgroupusers").
		Body(azurermApiManagementGroupUser).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermApiManagementGroupUser and updates it. Returns the server's representation of the azurermApiManagementGroupUser, and an error, if there is any.
func (c *azurermApiManagementGroupUsers) Update(azurermApiManagementGroupUser *v1alpha1.AzurermApiManagementGroupUser) (result *v1alpha1.AzurermApiManagementGroupUser, err error) {
	result = &v1alpha1.AzurermApiManagementGroupUser{}
	err = c.client.Put().
		Resource("azurermapimanagementgroupusers").
		Name(azurermApiManagementGroupUser.Name).
		Body(azurermApiManagementGroupUser).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermApiManagementGroupUsers) UpdateStatus(azurermApiManagementGroupUser *v1alpha1.AzurermApiManagementGroupUser) (result *v1alpha1.AzurermApiManagementGroupUser, err error) {
	result = &v1alpha1.AzurermApiManagementGroupUser{}
	err = c.client.Put().
		Resource("azurermapimanagementgroupusers").
		Name(azurermApiManagementGroupUser.Name).
		SubResource("status").
		Body(azurermApiManagementGroupUser).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermApiManagementGroupUser and deletes it. Returns an error if one occurs.
func (c *azurermApiManagementGroupUsers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermapimanagementgroupusers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermApiManagementGroupUsers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermapimanagementgroupusers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermApiManagementGroupUser.
func (c *azurermApiManagementGroupUsers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermApiManagementGroupUser, err error) {
	result = &v1alpha1.AzurermApiManagementGroupUser{}
	err = c.client.Patch(pt).
		Resource("azurermapimanagementgroupusers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
