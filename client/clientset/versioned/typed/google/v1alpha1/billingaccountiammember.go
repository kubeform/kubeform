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

// BillingAccountIamMembersGetter has a method to return a BillingAccountIamMemberInterface.
// A group's client should implement this interface.
type BillingAccountIamMembersGetter interface {
	BillingAccountIamMembers(namespace string) BillingAccountIamMemberInterface
}

// BillingAccountIamMemberInterface has methods to work with BillingAccountIamMember resources.
type BillingAccountIamMemberInterface interface {
	Create(*v1alpha1.BillingAccountIamMember) (*v1alpha1.BillingAccountIamMember, error)
	Update(*v1alpha1.BillingAccountIamMember) (*v1alpha1.BillingAccountIamMember, error)
	UpdateStatus(*v1alpha1.BillingAccountIamMember) (*v1alpha1.BillingAccountIamMember, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.BillingAccountIamMember, error)
	List(opts v1.ListOptions) (*v1alpha1.BillingAccountIamMemberList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BillingAccountIamMember, err error)
	BillingAccountIamMemberExpansion
}

// billingAccountIamMembers implements BillingAccountIamMemberInterface
type billingAccountIamMembers struct {
	client rest.Interface
	ns     string
}

// newBillingAccountIamMembers returns a BillingAccountIamMembers
func newBillingAccountIamMembers(c *GoogleV1alpha1Client, namespace string) *billingAccountIamMembers {
	return &billingAccountIamMembers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the billingAccountIamMember, and returns the corresponding billingAccountIamMember object, and an error if there is any.
func (c *billingAccountIamMembers) Get(name string, options v1.GetOptions) (result *v1alpha1.BillingAccountIamMember, err error) {
	result = &v1alpha1.BillingAccountIamMember{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("billingaccountiammembers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BillingAccountIamMembers that match those selectors.
func (c *billingAccountIamMembers) List(opts v1.ListOptions) (result *v1alpha1.BillingAccountIamMemberList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BillingAccountIamMemberList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("billingaccountiammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested billingAccountIamMembers.
func (c *billingAccountIamMembers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("billingaccountiammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a billingAccountIamMember and creates it.  Returns the server's representation of the billingAccountIamMember, and an error, if there is any.
func (c *billingAccountIamMembers) Create(billingAccountIamMember *v1alpha1.BillingAccountIamMember) (result *v1alpha1.BillingAccountIamMember, err error) {
	result = &v1alpha1.BillingAccountIamMember{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("billingaccountiammembers").
		Body(billingAccountIamMember).
		Do().
		Into(result)
	return
}

// Update takes the representation of a billingAccountIamMember and updates it. Returns the server's representation of the billingAccountIamMember, and an error, if there is any.
func (c *billingAccountIamMembers) Update(billingAccountIamMember *v1alpha1.BillingAccountIamMember) (result *v1alpha1.BillingAccountIamMember, err error) {
	result = &v1alpha1.BillingAccountIamMember{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("billingaccountiammembers").
		Name(billingAccountIamMember.Name).
		Body(billingAccountIamMember).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *billingAccountIamMembers) UpdateStatus(billingAccountIamMember *v1alpha1.BillingAccountIamMember) (result *v1alpha1.BillingAccountIamMember, err error) {
	result = &v1alpha1.BillingAccountIamMember{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("billingaccountiammembers").
		Name(billingAccountIamMember.Name).
		SubResource("status").
		Body(billingAccountIamMember).
		Do().
		Into(result)
	return
}

// Delete takes name of the billingAccountIamMember and deletes it. Returns an error if one occurs.
func (c *billingAccountIamMembers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("billingaccountiammembers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *billingAccountIamMembers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("billingaccountiammembers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched billingAccountIamMember.
func (c *billingAccountIamMembers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BillingAccountIamMember, err error) {
	result = &v1alpha1.BillingAccountIamMember{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("billingaccountiammembers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
