/*
Copyright 2019 The KubeDB Authors.

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

// GoogleMonitoringGroupsGetter has a method to return a GoogleMonitoringGroupInterface.
// A group's client should implement this interface.
type GoogleMonitoringGroupsGetter interface {
	GoogleMonitoringGroups() GoogleMonitoringGroupInterface
}

// GoogleMonitoringGroupInterface has methods to work with GoogleMonitoringGroup resources.
type GoogleMonitoringGroupInterface interface {
	Create(*v1alpha1.GoogleMonitoringGroup) (*v1alpha1.GoogleMonitoringGroup, error)
	Update(*v1alpha1.GoogleMonitoringGroup) (*v1alpha1.GoogleMonitoringGroup, error)
	UpdateStatus(*v1alpha1.GoogleMonitoringGroup) (*v1alpha1.GoogleMonitoringGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GoogleMonitoringGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.GoogleMonitoringGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleMonitoringGroup, err error)
	GoogleMonitoringGroupExpansion
}

// googleMonitoringGroups implements GoogleMonitoringGroupInterface
type googleMonitoringGroups struct {
	client rest.Interface
}

// newGoogleMonitoringGroups returns a GoogleMonitoringGroups
func newGoogleMonitoringGroups(c *GoogleV1alpha1Client) *googleMonitoringGroups {
	return &googleMonitoringGroups{
		client: c.RESTClient(),
	}
}

// Get takes name of the googleMonitoringGroup, and returns the corresponding googleMonitoringGroup object, and an error if there is any.
func (c *googleMonitoringGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleMonitoringGroup, err error) {
	result = &v1alpha1.GoogleMonitoringGroup{}
	err = c.client.Get().
		Resource("googlemonitoringgroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GoogleMonitoringGroups that match those selectors.
func (c *googleMonitoringGroups) List(opts v1.ListOptions) (result *v1alpha1.GoogleMonitoringGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GoogleMonitoringGroupList{}
	err = c.client.Get().
		Resource("googlemonitoringgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested googleMonitoringGroups.
func (c *googleMonitoringGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("googlemonitoringgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a googleMonitoringGroup and creates it.  Returns the server's representation of the googleMonitoringGroup, and an error, if there is any.
func (c *googleMonitoringGroups) Create(googleMonitoringGroup *v1alpha1.GoogleMonitoringGroup) (result *v1alpha1.GoogleMonitoringGroup, err error) {
	result = &v1alpha1.GoogleMonitoringGroup{}
	err = c.client.Post().
		Resource("googlemonitoringgroups").
		Body(googleMonitoringGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a googleMonitoringGroup and updates it. Returns the server's representation of the googleMonitoringGroup, and an error, if there is any.
func (c *googleMonitoringGroups) Update(googleMonitoringGroup *v1alpha1.GoogleMonitoringGroup) (result *v1alpha1.GoogleMonitoringGroup, err error) {
	result = &v1alpha1.GoogleMonitoringGroup{}
	err = c.client.Put().
		Resource("googlemonitoringgroups").
		Name(googleMonitoringGroup.Name).
		Body(googleMonitoringGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *googleMonitoringGroups) UpdateStatus(googleMonitoringGroup *v1alpha1.GoogleMonitoringGroup) (result *v1alpha1.GoogleMonitoringGroup, err error) {
	result = &v1alpha1.GoogleMonitoringGroup{}
	err = c.client.Put().
		Resource("googlemonitoringgroups").
		Name(googleMonitoringGroup.Name).
		SubResource("status").
		Body(googleMonitoringGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the googleMonitoringGroup and deletes it. Returns an error if one occurs.
func (c *googleMonitoringGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("googlemonitoringgroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *googleMonitoringGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("googlemonitoringgroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched googleMonitoringGroup.
func (c *googleMonitoringGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleMonitoringGroup, err error) {
	result = &v1alpha1.GoogleMonitoringGroup{}
	err = c.client.Patch(pt).
		Resource("googlemonitoringgroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
