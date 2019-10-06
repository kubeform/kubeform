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

// RamPrincipalAssociationsGetter has a method to return a RamPrincipalAssociationInterface.
// A group's client should implement this interface.
type RamPrincipalAssociationsGetter interface {
	RamPrincipalAssociations(namespace string) RamPrincipalAssociationInterface
}

// RamPrincipalAssociationInterface has methods to work with RamPrincipalAssociation resources.
type RamPrincipalAssociationInterface interface {
	Create(*v1alpha1.RamPrincipalAssociation) (*v1alpha1.RamPrincipalAssociation, error)
	Update(*v1alpha1.RamPrincipalAssociation) (*v1alpha1.RamPrincipalAssociation, error)
	UpdateStatus(*v1alpha1.RamPrincipalAssociation) (*v1alpha1.RamPrincipalAssociation, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.RamPrincipalAssociation, error)
	List(opts v1.ListOptions) (*v1alpha1.RamPrincipalAssociationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RamPrincipalAssociation, err error)
	RamPrincipalAssociationExpansion
}

// ramPrincipalAssociations implements RamPrincipalAssociationInterface
type ramPrincipalAssociations struct {
	client rest.Interface
	ns     string
}

// newRamPrincipalAssociations returns a RamPrincipalAssociations
func newRamPrincipalAssociations(c *AwsV1alpha1Client, namespace string) *ramPrincipalAssociations {
	return &ramPrincipalAssociations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the ramPrincipalAssociation, and returns the corresponding ramPrincipalAssociation object, and an error if there is any.
func (c *ramPrincipalAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.RamPrincipalAssociation, err error) {
	result = &v1alpha1.RamPrincipalAssociation{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ramprincipalassociations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RamPrincipalAssociations that match those selectors.
func (c *ramPrincipalAssociations) List(opts v1.ListOptions) (result *v1alpha1.RamPrincipalAssociationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.RamPrincipalAssociationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ramprincipalassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ramPrincipalAssociations.
func (c *ramPrincipalAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("ramprincipalassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a ramPrincipalAssociation and creates it.  Returns the server's representation of the ramPrincipalAssociation, and an error, if there is any.
func (c *ramPrincipalAssociations) Create(ramPrincipalAssociation *v1alpha1.RamPrincipalAssociation) (result *v1alpha1.RamPrincipalAssociation, err error) {
	result = &v1alpha1.RamPrincipalAssociation{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("ramprincipalassociations").
		Body(ramPrincipalAssociation).
		Do().
		Into(result)
	return
}

// Update takes the representation of a ramPrincipalAssociation and updates it. Returns the server's representation of the ramPrincipalAssociation, and an error, if there is any.
func (c *ramPrincipalAssociations) Update(ramPrincipalAssociation *v1alpha1.RamPrincipalAssociation) (result *v1alpha1.RamPrincipalAssociation, err error) {
	result = &v1alpha1.RamPrincipalAssociation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ramprincipalassociations").
		Name(ramPrincipalAssociation.Name).
		Body(ramPrincipalAssociation).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *ramPrincipalAssociations) UpdateStatus(ramPrincipalAssociation *v1alpha1.RamPrincipalAssociation) (result *v1alpha1.RamPrincipalAssociation, err error) {
	result = &v1alpha1.RamPrincipalAssociation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ramprincipalassociations").
		Name(ramPrincipalAssociation.Name).
		SubResource("status").
		Body(ramPrincipalAssociation).
		Do().
		Into(result)
	return
}

// Delete takes name of the ramPrincipalAssociation and deletes it. Returns an error if one occurs.
func (c *ramPrincipalAssociations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ramprincipalassociations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ramPrincipalAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ramprincipalassociations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched ramPrincipalAssociation.
func (c *ramPrincipalAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RamPrincipalAssociation, err error) {
	result = &v1alpha1.RamPrincipalAssociation{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("ramprincipalassociations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
