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

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DbSubnetGroupsGetter has a method to return a DbSubnetGroupInterface.
// A group's client should implement this interface.
type DbSubnetGroupsGetter interface {
	DbSubnetGroups(namespace string) DbSubnetGroupInterface
}

// DbSubnetGroupInterface has methods to work with DbSubnetGroup resources.
type DbSubnetGroupInterface interface {
	Create(*v1alpha1.DbSubnetGroup) (*v1alpha1.DbSubnetGroup, error)
	Update(*v1alpha1.DbSubnetGroup) (*v1alpha1.DbSubnetGroup, error)
	UpdateStatus(*v1alpha1.DbSubnetGroup) (*v1alpha1.DbSubnetGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DbSubnetGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.DbSubnetGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DbSubnetGroup, err error)
	DbSubnetGroupExpansion
}

// dbSubnetGroups implements DbSubnetGroupInterface
type dbSubnetGroups struct {
	client rest.Interface
	ns     string
}

// newDbSubnetGroups returns a DbSubnetGroups
func newDbSubnetGroups(c *AwsV1alpha1Client, namespace string) *dbSubnetGroups {
	return &dbSubnetGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dbSubnetGroup, and returns the corresponding dbSubnetGroup object, and an error if there is any.
func (c *dbSubnetGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.DbSubnetGroup, err error) {
	result = &v1alpha1.DbSubnetGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dbsubnetgroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DbSubnetGroups that match those selectors.
func (c *dbSubnetGroups) List(opts v1.ListOptions) (result *v1alpha1.DbSubnetGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DbSubnetGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dbsubnetgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dbSubnetGroups.
func (c *dbSubnetGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dbsubnetgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a dbSubnetGroup and creates it.  Returns the server's representation of the dbSubnetGroup, and an error, if there is any.
func (c *dbSubnetGroups) Create(dbSubnetGroup *v1alpha1.DbSubnetGroup) (result *v1alpha1.DbSubnetGroup, err error) {
	result = &v1alpha1.DbSubnetGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dbsubnetgroups").
		Body(dbSubnetGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dbSubnetGroup and updates it. Returns the server's representation of the dbSubnetGroup, and an error, if there is any.
func (c *dbSubnetGroups) Update(dbSubnetGroup *v1alpha1.DbSubnetGroup) (result *v1alpha1.DbSubnetGroup, err error) {
	result = &v1alpha1.DbSubnetGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dbsubnetgroups").
		Name(dbSubnetGroup.Name).
		Body(dbSubnetGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *dbSubnetGroups) UpdateStatus(dbSubnetGroup *v1alpha1.DbSubnetGroup) (result *v1alpha1.DbSubnetGroup, err error) {
	result = &v1alpha1.DbSubnetGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dbsubnetgroups").
		Name(dbSubnetGroup.Name).
		SubResource("status").
		Body(dbSubnetGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the dbSubnetGroup and deletes it. Returns an error if one occurs.
func (c *dbSubnetGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dbsubnetgroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dbSubnetGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dbsubnetgroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dbSubnetGroup.
func (c *dbSubnetGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DbSubnetGroup, err error) {
	result = &v1alpha1.DbSubnetGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dbsubnetgroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
