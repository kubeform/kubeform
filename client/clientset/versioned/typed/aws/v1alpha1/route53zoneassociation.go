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

// Route53ZoneAssociationsGetter has a method to return a Route53ZoneAssociationInterface.
// A group's client should implement this interface.
type Route53ZoneAssociationsGetter interface {
	Route53ZoneAssociations(namespace string) Route53ZoneAssociationInterface
}

// Route53ZoneAssociationInterface has methods to work with Route53ZoneAssociation resources.
type Route53ZoneAssociationInterface interface {
	Create(*v1alpha1.Route53ZoneAssociation) (*v1alpha1.Route53ZoneAssociation, error)
	Update(*v1alpha1.Route53ZoneAssociation) (*v1alpha1.Route53ZoneAssociation, error)
	UpdateStatus(*v1alpha1.Route53ZoneAssociation) (*v1alpha1.Route53ZoneAssociation, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Route53ZoneAssociation, error)
	List(opts v1.ListOptions) (*v1alpha1.Route53ZoneAssociationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Route53ZoneAssociation, err error)
	Route53ZoneAssociationExpansion
}

// route53ZoneAssociations implements Route53ZoneAssociationInterface
type route53ZoneAssociations struct {
	client rest.Interface
	ns     string
}

// newRoute53ZoneAssociations returns a Route53ZoneAssociations
func newRoute53ZoneAssociations(c *AwsV1alpha1Client, namespace string) *route53ZoneAssociations {
	return &route53ZoneAssociations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the route53ZoneAssociation, and returns the corresponding route53ZoneAssociation object, and an error if there is any.
func (c *route53ZoneAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.Route53ZoneAssociation, err error) {
	result = &v1alpha1.Route53ZoneAssociation{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("route53zoneassociations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Route53ZoneAssociations that match those selectors.
func (c *route53ZoneAssociations) List(opts v1.ListOptions) (result *v1alpha1.Route53ZoneAssociationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.Route53ZoneAssociationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("route53zoneassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested route53ZoneAssociations.
func (c *route53ZoneAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("route53zoneassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a route53ZoneAssociation and creates it.  Returns the server's representation of the route53ZoneAssociation, and an error, if there is any.
func (c *route53ZoneAssociations) Create(route53ZoneAssociation *v1alpha1.Route53ZoneAssociation) (result *v1alpha1.Route53ZoneAssociation, err error) {
	result = &v1alpha1.Route53ZoneAssociation{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("route53zoneassociations").
		Body(route53ZoneAssociation).
		Do().
		Into(result)
	return
}

// Update takes the representation of a route53ZoneAssociation and updates it. Returns the server's representation of the route53ZoneAssociation, and an error, if there is any.
func (c *route53ZoneAssociations) Update(route53ZoneAssociation *v1alpha1.Route53ZoneAssociation) (result *v1alpha1.Route53ZoneAssociation, err error) {
	result = &v1alpha1.Route53ZoneAssociation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("route53zoneassociations").
		Name(route53ZoneAssociation.Name).
		Body(route53ZoneAssociation).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *route53ZoneAssociations) UpdateStatus(route53ZoneAssociation *v1alpha1.Route53ZoneAssociation) (result *v1alpha1.Route53ZoneAssociation, err error) {
	result = &v1alpha1.Route53ZoneAssociation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("route53zoneassociations").
		Name(route53ZoneAssociation.Name).
		SubResource("status").
		Body(route53ZoneAssociation).
		Do().
		Into(result)
	return
}

// Delete takes name of the route53ZoneAssociation and deletes it. Returns an error if one occurs.
func (c *route53ZoneAssociations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("route53zoneassociations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *route53ZoneAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("route53zoneassociations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched route53ZoneAssociation.
func (c *route53ZoneAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Route53ZoneAssociation, err error) {
	result = &v1alpha1.Route53ZoneAssociation{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("route53zoneassociations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}