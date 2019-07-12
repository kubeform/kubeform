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

// GoogleComputeInstanceGroupManagersGetter has a method to return a GoogleComputeInstanceGroupManagerInterface.
// A group's client should implement this interface.
type GoogleComputeInstanceGroupManagersGetter interface {
	GoogleComputeInstanceGroupManagers() GoogleComputeInstanceGroupManagerInterface
}

// GoogleComputeInstanceGroupManagerInterface has methods to work with GoogleComputeInstanceGroupManager resources.
type GoogleComputeInstanceGroupManagerInterface interface {
	Create(*v1alpha1.GoogleComputeInstanceGroupManager) (*v1alpha1.GoogleComputeInstanceGroupManager, error)
	Update(*v1alpha1.GoogleComputeInstanceGroupManager) (*v1alpha1.GoogleComputeInstanceGroupManager, error)
	UpdateStatus(*v1alpha1.GoogleComputeInstanceGroupManager) (*v1alpha1.GoogleComputeInstanceGroupManager, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GoogleComputeInstanceGroupManager, error)
	List(opts v1.ListOptions) (*v1alpha1.GoogleComputeInstanceGroupManagerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleComputeInstanceGroupManager, err error)
	GoogleComputeInstanceGroupManagerExpansion
}

// googleComputeInstanceGroupManagers implements GoogleComputeInstanceGroupManagerInterface
type googleComputeInstanceGroupManagers struct {
	client rest.Interface
}

// newGoogleComputeInstanceGroupManagers returns a GoogleComputeInstanceGroupManagers
func newGoogleComputeInstanceGroupManagers(c *GoogleV1alpha1Client) *googleComputeInstanceGroupManagers {
	return &googleComputeInstanceGroupManagers{
		client: c.RESTClient(),
	}
}

// Get takes name of the googleComputeInstanceGroupManager, and returns the corresponding googleComputeInstanceGroupManager object, and an error if there is any.
func (c *googleComputeInstanceGroupManagers) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleComputeInstanceGroupManager, err error) {
	result = &v1alpha1.GoogleComputeInstanceGroupManager{}
	err = c.client.Get().
		Resource("googlecomputeinstancegroupmanagers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GoogleComputeInstanceGroupManagers that match those selectors.
func (c *googleComputeInstanceGroupManagers) List(opts v1.ListOptions) (result *v1alpha1.GoogleComputeInstanceGroupManagerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GoogleComputeInstanceGroupManagerList{}
	err = c.client.Get().
		Resource("googlecomputeinstancegroupmanagers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested googleComputeInstanceGroupManagers.
func (c *googleComputeInstanceGroupManagers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("googlecomputeinstancegroupmanagers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a googleComputeInstanceGroupManager and creates it.  Returns the server's representation of the googleComputeInstanceGroupManager, and an error, if there is any.
func (c *googleComputeInstanceGroupManagers) Create(googleComputeInstanceGroupManager *v1alpha1.GoogleComputeInstanceGroupManager) (result *v1alpha1.GoogleComputeInstanceGroupManager, err error) {
	result = &v1alpha1.GoogleComputeInstanceGroupManager{}
	err = c.client.Post().
		Resource("googlecomputeinstancegroupmanagers").
		Body(googleComputeInstanceGroupManager).
		Do().
		Into(result)
	return
}

// Update takes the representation of a googleComputeInstanceGroupManager and updates it. Returns the server's representation of the googleComputeInstanceGroupManager, and an error, if there is any.
func (c *googleComputeInstanceGroupManagers) Update(googleComputeInstanceGroupManager *v1alpha1.GoogleComputeInstanceGroupManager) (result *v1alpha1.GoogleComputeInstanceGroupManager, err error) {
	result = &v1alpha1.GoogleComputeInstanceGroupManager{}
	err = c.client.Put().
		Resource("googlecomputeinstancegroupmanagers").
		Name(googleComputeInstanceGroupManager.Name).
		Body(googleComputeInstanceGroupManager).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *googleComputeInstanceGroupManagers) UpdateStatus(googleComputeInstanceGroupManager *v1alpha1.GoogleComputeInstanceGroupManager) (result *v1alpha1.GoogleComputeInstanceGroupManager, err error) {
	result = &v1alpha1.GoogleComputeInstanceGroupManager{}
	err = c.client.Put().
		Resource("googlecomputeinstancegroupmanagers").
		Name(googleComputeInstanceGroupManager.Name).
		SubResource("status").
		Body(googleComputeInstanceGroupManager).
		Do().
		Into(result)
	return
}

// Delete takes name of the googleComputeInstanceGroupManager and deletes it. Returns an error if one occurs.
func (c *googleComputeInstanceGroupManagers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("googlecomputeinstancegroupmanagers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *googleComputeInstanceGroupManagers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("googlecomputeinstancegroupmanagers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched googleComputeInstanceGroupManager.
func (c *googleComputeInstanceGroupManagers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleComputeInstanceGroupManager, err error) {
	result = &v1alpha1.GoogleComputeInstanceGroupManager{}
	err = c.client.Patch(pt).
		Resource("googlecomputeinstancegroupmanagers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
