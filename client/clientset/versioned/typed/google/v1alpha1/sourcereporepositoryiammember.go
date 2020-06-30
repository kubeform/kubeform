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

// SourcerepoRepositoryIamMembersGetter has a method to return a SourcerepoRepositoryIamMemberInterface.
// A group's client should implement this interface.
type SourcerepoRepositoryIamMembersGetter interface {
	SourcerepoRepositoryIamMembers(namespace string) SourcerepoRepositoryIamMemberInterface
}

// SourcerepoRepositoryIamMemberInterface has methods to work with SourcerepoRepositoryIamMember resources.
type SourcerepoRepositoryIamMemberInterface interface {
	Create(*v1alpha1.SourcerepoRepositoryIamMember) (*v1alpha1.SourcerepoRepositoryIamMember, error)
	Update(*v1alpha1.SourcerepoRepositoryIamMember) (*v1alpha1.SourcerepoRepositoryIamMember, error)
	UpdateStatus(*v1alpha1.SourcerepoRepositoryIamMember) (*v1alpha1.SourcerepoRepositoryIamMember, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SourcerepoRepositoryIamMember, error)
	List(opts v1.ListOptions) (*v1alpha1.SourcerepoRepositoryIamMemberList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SourcerepoRepositoryIamMember, err error)
	SourcerepoRepositoryIamMemberExpansion
}

// sourcerepoRepositoryIamMembers implements SourcerepoRepositoryIamMemberInterface
type sourcerepoRepositoryIamMembers struct {
	client rest.Interface
	ns     string
}

// newSourcerepoRepositoryIamMembers returns a SourcerepoRepositoryIamMembers
func newSourcerepoRepositoryIamMembers(c *GoogleV1alpha1Client, namespace string) *sourcerepoRepositoryIamMembers {
	return &sourcerepoRepositoryIamMembers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sourcerepoRepositoryIamMember, and returns the corresponding sourcerepoRepositoryIamMember object, and an error if there is any.
func (c *sourcerepoRepositoryIamMembers) Get(name string, options v1.GetOptions) (result *v1alpha1.SourcerepoRepositoryIamMember, err error) {
	result = &v1alpha1.SourcerepoRepositoryIamMember{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sourcereporepositoryiammembers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SourcerepoRepositoryIamMembers that match those selectors.
func (c *sourcerepoRepositoryIamMembers) List(opts v1.ListOptions) (result *v1alpha1.SourcerepoRepositoryIamMemberList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SourcerepoRepositoryIamMemberList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sourcereporepositoryiammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sourcerepoRepositoryIamMembers.
func (c *sourcerepoRepositoryIamMembers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sourcereporepositoryiammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a sourcerepoRepositoryIamMember and creates it.  Returns the server's representation of the sourcerepoRepositoryIamMember, and an error, if there is any.
func (c *sourcerepoRepositoryIamMembers) Create(sourcerepoRepositoryIamMember *v1alpha1.SourcerepoRepositoryIamMember) (result *v1alpha1.SourcerepoRepositoryIamMember, err error) {
	result = &v1alpha1.SourcerepoRepositoryIamMember{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sourcereporepositoryiammembers").
		Body(sourcerepoRepositoryIamMember).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sourcerepoRepositoryIamMember and updates it. Returns the server's representation of the sourcerepoRepositoryIamMember, and an error, if there is any.
func (c *sourcerepoRepositoryIamMembers) Update(sourcerepoRepositoryIamMember *v1alpha1.SourcerepoRepositoryIamMember) (result *v1alpha1.SourcerepoRepositoryIamMember, err error) {
	result = &v1alpha1.SourcerepoRepositoryIamMember{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sourcereporepositoryiammembers").
		Name(sourcerepoRepositoryIamMember.Name).
		Body(sourcerepoRepositoryIamMember).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *sourcerepoRepositoryIamMembers) UpdateStatus(sourcerepoRepositoryIamMember *v1alpha1.SourcerepoRepositoryIamMember) (result *v1alpha1.SourcerepoRepositoryIamMember, err error) {
	result = &v1alpha1.SourcerepoRepositoryIamMember{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sourcereporepositoryiammembers").
		Name(sourcerepoRepositoryIamMember.Name).
		SubResource("status").
		Body(sourcerepoRepositoryIamMember).
		Do().
		Into(result)
	return
}

// Delete takes name of the sourcerepoRepositoryIamMember and deletes it. Returns an error if one occurs.
func (c *sourcerepoRepositoryIamMembers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sourcereporepositoryiammembers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sourcerepoRepositoryIamMembers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sourcereporepositoryiammembers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sourcerepoRepositoryIamMember.
func (c *sourcerepoRepositoryIamMembers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SourcerepoRepositoryIamMember, err error) {
	result = &v1alpha1.SourcerepoRepositoryIamMember{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sourcereporepositoryiammembers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
