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

// OrganizationsAccountsGetter has a method to return a OrganizationsAccountInterface.
// A group's client should implement this interface.
type OrganizationsAccountsGetter interface {
	OrganizationsAccounts(namespace string) OrganizationsAccountInterface
}

// OrganizationsAccountInterface has methods to work with OrganizationsAccount resources.
type OrganizationsAccountInterface interface {
	Create(*v1alpha1.OrganizationsAccount) (*v1alpha1.OrganizationsAccount, error)
	Update(*v1alpha1.OrganizationsAccount) (*v1alpha1.OrganizationsAccount, error)
	UpdateStatus(*v1alpha1.OrganizationsAccount) (*v1alpha1.OrganizationsAccount, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.OrganizationsAccount, error)
	List(opts v1.ListOptions) (*v1alpha1.OrganizationsAccountList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OrganizationsAccount, err error)
	OrganizationsAccountExpansion
}

// organizationsAccounts implements OrganizationsAccountInterface
type organizationsAccounts struct {
	client rest.Interface
	ns     string
}

// newOrganizationsAccounts returns a OrganizationsAccounts
func newOrganizationsAccounts(c *AwsV1alpha1Client, namespace string) *organizationsAccounts {
	return &organizationsAccounts{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the organizationsAccount, and returns the corresponding organizationsAccount object, and an error if there is any.
func (c *organizationsAccounts) Get(name string, options v1.GetOptions) (result *v1alpha1.OrganizationsAccount, err error) {
	result = &v1alpha1.OrganizationsAccount{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("organizationsaccounts").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of OrganizationsAccounts that match those selectors.
func (c *organizationsAccounts) List(opts v1.ListOptions) (result *v1alpha1.OrganizationsAccountList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.OrganizationsAccountList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("organizationsaccounts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested organizationsAccounts.
func (c *organizationsAccounts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("organizationsaccounts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a organizationsAccount and creates it.  Returns the server's representation of the organizationsAccount, and an error, if there is any.
func (c *organizationsAccounts) Create(organizationsAccount *v1alpha1.OrganizationsAccount) (result *v1alpha1.OrganizationsAccount, err error) {
	result = &v1alpha1.OrganizationsAccount{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("organizationsaccounts").
		Body(organizationsAccount).
		Do().
		Into(result)
	return
}

// Update takes the representation of a organizationsAccount and updates it. Returns the server's representation of the organizationsAccount, and an error, if there is any.
func (c *organizationsAccounts) Update(organizationsAccount *v1alpha1.OrganizationsAccount) (result *v1alpha1.OrganizationsAccount, err error) {
	result = &v1alpha1.OrganizationsAccount{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("organizationsaccounts").
		Name(organizationsAccount.Name).
		Body(organizationsAccount).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *organizationsAccounts) UpdateStatus(organizationsAccount *v1alpha1.OrganizationsAccount) (result *v1alpha1.OrganizationsAccount, err error) {
	result = &v1alpha1.OrganizationsAccount{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("organizationsaccounts").
		Name(organizationsAccount.Name).
		SubResource("status").
		Body(organizationsAccount).
		Do().
		Into(result)
	return
}

// Delete takes name of the organizationsAccount and deletes it. Returns an error if one occurs.
func (c *organizationsAccounts) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("organizationsaccounts").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *organizationsAccounts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("organizationsaccounts").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched organizationsAccount.
func (c *organizationsAccounts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OrganizationsAccount, err error) {
	result = &v1alpha1.OrganizationsAccount{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("organizationsaccounts").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}