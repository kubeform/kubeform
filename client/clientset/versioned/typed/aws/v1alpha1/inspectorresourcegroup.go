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

// InspectorResourceGroupsGetter has a method to return a InspectorResourceGroupInterface.
// A group's client should implement this interface.
type InspectorResourceGroupsGetter interface {
	InspectorResourceGroups(namespace string) InspectorResourceGroupInterface
}

// InspectorResourceGroupInterface has methods to work with InspectorResourceGroup resources.
type InspectorResourceGroupInterface interface {
	Create(*v1alpha1.InspectorResourceGroup) (*v1alpha1.InspectorResourceGroup, error)
	Update(*v1alpha1.InspectorResourceGroup) (*v1alpha1.InspectorResourceGroup, error)
	UpdateStatus(*v1alpha1.InspectorResourceGroup) (*v1alpha1.InspectorResourceGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.InspectorResourceGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.InspectorResourceGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.InspectorResourceGroup, err error)
	InspectorResourceGroupExpansion
}

// inspectorResourceGroups implements InspectorResourceGroupInterface
type inspectorResourceGroups struct {
	client rest.Interface
	ns     string
}

// newInspectorResourceGroups returns a InspectorResourceGroups
func newInspectorResourceGroups(c *AwsV1alpha1Client, namespace string) *inspectorResourceGroups {
	return &inspectorResourceGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the inspectorResourceGroup, and returns the corresponding inspectorResourceGroup object, and an error if there is any.
func (c *inspectorResourceGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.InspectorResourceGroup, err error) {
	result = &v1alpha1.InspectorResourceGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("inspectorresourcegroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of InspectorResourceGroups that match those selectors.
func (c *inspectorResourceGroups) List(opts v1.ListOptions) (result *v1alpha1.InspectorResourceGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.InspectorResourceGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("inspectorresourcegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested inspectorResourceGroups.
func (c *inspectorResourceGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("inspectorresourcegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a inspectorResourceGroup and creates it.  Returns the server's representation of the inspectorResourceGroup, and an error, if there is any.
func (c *inspectorResourceGroups) Create(inspectorResourceGroup *v1alpha1.InspectorResourceGroup) (result *v1alpha1.InspectorResourceGroup, err error) {
	result = &v1alpha1.InspectorResourceGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("inspectorresourcegroups").
		Body(inspectorResourceGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a inspectorResourceGroup and updates it. Returns the server's representation of the inspectorResourceGroup, and an error, if there is any.
func (c *inspectorResourceGroups) Update(inspectorResourceGroup *v1alpha1.InspectorResourceGroup) (result *v1alpha1.InspectorResourceGroup, err error) {
	result = &v1alpha1.InspectorResourceGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("inspectorresourcegroups").
		Name(inspectorResourceGroup.Name).
		Body(inspectorResourceGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *inspectorResourceGroups) UpdateStatus(inspectorResourceGroup *v1alpha1.InspectorResourceGroup) (result *v1alpha1.InspectorResourceGroup, err error) {
	result = &v1alpha1.InspectorResourceGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("inspectorresourcegroups").
		Name(inspectorResourceGroup.Name).
		SubResource("status").
		Body(inspectorResourceGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the inspectorResourceGroup and deletes it. Returns an error if one occurs.
func (c *inspectorResourceGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("inspectorresourcegroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *inspectorResourceGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("inspectorresourcegroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched inspectorResourceGroup.
func (c *inspectorResourceGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.InspectorResourceGroup, err error) {
	result = &v1alpha1.InspectorResourceGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("inspectorresourcegroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
