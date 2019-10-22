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

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DbParameterGroupsGetter has a method to return a DbParameterGroupInterface.
// A group's client should implement this interface.
type DbParameterGroupsGetter interface {
	DbParameterGroups(namespace string) DbParameterGroupInterface
}

// DbParameterGroupInterface has methods to work with DbParameterGroup resources.
type DbParameterGroupInterface interface {
	Create(*v1alpha1.DbParameterGroup) (*v1alpha1.DbParameterGroup, error)
	Update(*v1alpha1.DbParameterGroup) (*v1alpha1.DbParameterGroup, error)
	UpdateStatus(*v1alpha1.DbParameterGroup) (*v1alpha1.DbParameterGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DbParameterGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.DbParameterGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DbParameterGroup, err error)
	DbParameterGroupExpansion
}

// dbParameterGroups implements DbParameterGroupInterface
type dbParameterGroups struct {
	client rest.Interface
	ns     string
}

// newDbParameterGroups returns a DbParameterGroups
func newDbParameterGroups(c *AwsV1alpha1Client, namespace string) *dbParameterGroups {
	return &dbParameterGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dbParameterGroup, and returns the corresponding dbParameterGroup object, and an error if there is any.
func (c *dbParameterGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.DbParameterGroup, err error) {
	result = &v1alpha1.DbParameterGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dbparametergroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DbParameterGroups that match those selectors.
func (c *dbParameterGroups) List(opts v1.ListOptions) (result *v1alpha1.DbParameterGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DbParameterGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dbparametergroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dbParameterGroups.
func (c *dbParameterGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dbparametergroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a dbParameterGroup and creates it.  Returns the server's representation of the dbParameterGroup, and an error, if there is any.
func (c *dbParameterGroups) Create(dbParameterGroup *v1alpha1.DbParameterGroup) (result *v1alpha1.DbParameterGroup, err error) {
	result = &v1alpha1.DbParameterGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dbparametergroups").
		Body(dbParameterGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dbParameterGroup and updates it. Returns the server's representation of the dbParameterGroup, and an error, if there is any.
func (c *dbParameterGroups) Update(dbParameterGroup *v1alpha1.DbParameterGroup) (result *v1alpha1.DbParameterGroup, err error) {
	result = &v1alpha1.DbParameterGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dbparametergroups").
		Name(dbParameterGroup.Name).
		Body(dbParameterGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *dbParameterGroups) UpdateStatus(dbParameterGroup *v1alpha1.DbParameterGroup) (result *v1alpha1.DbParameterGroup, err error) {
	result = &v1alpha1.DbParameterGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dbparametergroups").
		Name(dbParameterGroup.Name).
		SubResource("status").
		Body(dbParameterGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the dbParameterGroup and deletes it. Returns an error if one occurs.
func (c *dbParameterGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dbparametergroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dbParameterGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dbparametergroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dbParameterGroup.
func (c *dbParameterGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DbParameterGroup, err error) {
	result = &v1alpha1.DbParameterGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dbparametergroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
