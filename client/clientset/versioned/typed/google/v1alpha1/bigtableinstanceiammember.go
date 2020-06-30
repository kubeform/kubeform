/*
Copyright The Kubeform Authors.

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

	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BigtableInstanceIamMembersGetter has a method to return a BigtableInstanceIamMemberInterface.
// A group's client should implement this interface.
type BigtableInstanceIamMembersGetter interface {
	BigtableInstanceIamMembers(namespace string) BigtableInstanceIamMemberInterface
}

// BigtableInstanceIamMemberInterface has methods to work with BigtableInstanceIamMember resources.
type BigtableInstanceIamMemberInterface interface {
	Create(*v1alpha1.BigtableInstanceIamMember) (*v1alpha1.BigtableInstanceIamMember, error)
	Update(*v1alpha1.BigtableInstanceIamMember) (*v1alpha1.BigtableInstanceIamMember, error)
	UpdateStatus(*v1alpha1.BigtableInstanceIamMember) (*v1alpha1.BigtableInstanceIamMember, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.BigtableInstanceIamMember, error)
	List(opts v1.ListOptions) (*v1alpha1.BigtableInstanceIamMemberList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BigtableInstanceIamMember, err error)
	BigtableInstanceIamMemberExpansion
}

// bigtableInstanceIamMembers implements BigtableInstanceIamMemberInterface
type bigtableInstanceIamMembers struct {
	client rest.Interface
	ns     string
}

// newBigtableInstanceIamMembers returns a BigtableInstanceIamMembers
func newBigtableInstanceIamMembers(c *GoogleV1alpha1Client, namespace string) *bigtableInstanceIamMembers {
	return &bigtableInstanceIamMembers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the bigtableInstanceIamMember, and returns the corresponding bigtableInstanceIamMember object, and an error if there is any.
func (c *bigtableInstanceIamMembers) Get(name string, options v1.GetOptions) (result *v1alpha1.BigtableInstanceIamMember, err error) {
	result = &v1alpha1.BigtableInstanceIamMember{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bigtableinstanceiammembers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BigtableInstanceIamMembers that match those selectors.
func (c *bigtableInstanceIamMembers) List(opts v1.ListOptions) (result *v1alpha1.BigtableInstanceIamMemberList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BigtableInstanceIamMemberList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bigtableinstanceiammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested bigtableInstanceIamMembers.
func (c *bigtableInstanceIamMembers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("bigtableinstanceiammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a bigtableInstanceIamMember and creates it.  Returns the server's representation of the bigtableInstanceIamMember, and an error, if there is any.
func (c *bigtableInstanceIamMembers) Create(bigtableInstanceIamMember *v1alpha1.BigtableInstanceIamMember) (result *v1alpha1.BigtableInstanceIamMember, err error) {
	result = &v1alpha1.BigtableInstanceIamMember{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("bigtableinstanceiammembers").
		Body(bigtableInstanceIamMember).
		Do().
		Into(result)
	return
}

// Update takes the representation of a bigtableInstanceIamMember and updates it. Returns the server's representation of the bigtableInstanceIamMember, and an error, if there is any.
func (c *bigtableInstanceIamMembers) Update(bigtableInstanceIamMember *v1alpha1.BigtableInstanceIamMember) (result *v1alpha1.BigtableInstanceIamMember, err error) {
	result = &v1alpha1.BigtableInstanceIamMember{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bigtableinstanceiammembers").
		Name(bigtableInstanceIamMember.Name).
		Body(bigtableInstanceIamMember).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *bigtableInstanceIamMembers) UpdateStatus(bigtableInstanceIamMember *v1alpha1.BigtableInstanceIamMember) (result *v1alpha1.BigtableInstanceIamMember, err error) {
	result = &v1alpha1.BigtableInstanceIamMember{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bigtableinstanceiammembers").
		Name(bigtableInstanceIamMember.Name).
		SubResource("status").
		Body(bigtableInstanceIamMember).
		Do().
		Into(result)
	return
}

// Delete takes name of the bigtableInstanceIamMember and deletes it. Returns an error if one occurs.
func (c *bigtableInstanceIamMembers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bigtableinstanceiammembers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *bigtableInstanceIamMembers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bigtableinstanceiammembers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched bigtableInstanceIamMember.
func (c *bigtableInstanceIamMembers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BigtableInstanceIamMember, err error) {
	result = &v1alpha1.BigtableInstanceIamMember{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("bigtableinstanceiammembers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
