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

// DmsReplicationSubnetGroupsGetter has a method to return a DmsReplicationSubnetGroupInterface.
// A group's client should implement this interface.
type DmsReplicationSubnetGroupsGetter interface {
	DmsReplicationSubnetGroups(namespace string) DmsReplicationSubnetGroupInterface
}

// DmsReplicationSubnetGroupInterface has methods to work with DmsReplicationSubnetGroup resources.
type DmsReplicationSubnetGroupInterface interface {
	Create(*v1alpha1.DmsReplicationSubnetGroup) (*v1alpha1.DmsReplicationSubnetGroup, error)
	Update(*v1alpha1.DmsReplicationSubnetGroup) (*v1alpha1.DmsReplicationSubnetGroup, error)
	UpdateStatus(*v1alpha1.DmsReplicationSubnetGroup) (*v1alpha1.DmsReplicationSubnetGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DmsReplicationSubnetGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.DmsReplicationSubnetGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DmsReplicationSubnetGroup, err error)
	DmsReplicationSubnetGroupExpansion
}

// dmsReplicationSubnetGroups implements DmsReplicationSubnetGroupInterface
type dmsReplicationSubnetGroups struct {
	client rest.Interface
	ns     string
}

// newDmsReplicationSubnetGroups returns a DmsReplicationSubnetGroups
func newDmsReplicationSubnetGroups(c *AwsV1alpha1Client, namespace string) *dmsReplicationSubnetGroups {
	return &dmsReplicationSubnetGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dmsReplicationSubnetGroup, and returns the corresponding dmsReplicationSubnetGroup object, and an error if there is any.
func (c *dmsReplicationSubnetGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.DmsReplicationSubnetGroup, err error) {
	result = &v1alpha1.DmsReplicationSubnetGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dmsreplicationsubnetgroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DmsReplicationSubnetGroups that match those selectors.
func (c *dmsReplicationSubnetGroups) List(opts v1.ListOptions) (result *v1alpha1.DmsReplicationSubnetGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DmsReplicationSubnetGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dmsreplicationsubnetgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dmsReplicationSubnetGroups.
func (c *dmsReplicationSubnetGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dmsreplicationsubnetgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a dmsReplicationSubnetGroup and creates it.  Returns the server's representation of the dmsReplicationSubnetGroup, and an error, if there is any.
func (c *dmsReplicationSubnetGroups) Create(dmsReplicationSubnetGroup *v1alpha1.DmsReplicationSubnetGroup) (result *v1alpha1.DmsReplicationSubnetGroup, err error) {
	result = &v1alpha1.DmsReplicationSubnetGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dmsreplicationsubnetgroups").
		Body(dmsReplicationSubnetGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dmsReplicationSubnetGroup and updates it. Returns the server's representation of the dmsReplicationSubnetGroup, and an error, if there is any.
func (c *dmsReplicationSubnetGroups) Update(dmsReplicationSubnetGroup *v1alpha1.DmsReplicationSubnetGroup) (result *v1alpha1.DmsReplicationSubnetGroup, err error) {
	result = &v1alpha1.DmsReplicationSubnetGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dmsreplicationsubnetgroups").
		Name(dmsReplicationSubnetGroup.Name).
		Body(dmsReplicationSubnetGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *dmsReplicationSubnetGroups) UpdateStatus(dmsReplicationSubnetGroup *v1alpha1.DmsReplicationSubnetGroup) (result *v1alpha1.DmsReplicationSubnetGroup, err error) {
	result = &v1alpha1.DmsReplicationSubnetGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dmsreplicationsubnetgroups").
		Name(dmsReplicationSubnetGroup.Name).
		SubResource("status").
		Body(dmsReplicationSubnetGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the dmsReplicationSubnetGroup and deletes it. Returns an error if one occurs.
func (c *dmsReplicationSubnetGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dmsreplicationsubnetgroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dmsReplicationSubnetGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dmsreplicationsubnetgroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dmsReplicationSubnetGroup.
func (c *dmsReplicationSubnetGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DmsReplicationSubnetGroup, err error) {
	result = &v1alpha1.DmsReplicationSubnetGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dmsreplicationsubnetgroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
