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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// MonitorLogProfilesGetter has a method to return a MonitorLogProfileInterface.
// A group's client should implement this interface.
type MonitorLogProfilesGetter interface {
	MonitorLogProfiles(namespace string) MonitorLogProfileInterface
}

// MonitorLogProfileInterface has methods to work with MonitorLogProfile resources.
type MonitorLogProfileInterface interface {
	Create(*v1alpha1.MonitorLogProfile) (*v1alpha1.MonitorLogProfile, error)
	Update(*v1alpha1.MonitorLogProfile) (*v1alpha1.MonitorLogProfile, error)
	UpdateStatus(*v1alpha1.MonitorLogProfile) (*v1alpha1.MonitorLogProfile, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.MonitorLogProfile, error)
	List(opts v1.ListOptions) (*v1alpha1.MonitorLogProfileList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MonitorLogProfile, err error)
	MonitorLogProfileExpansion
}

// monitorLogProfiles implements MonitorLogProfileInterface
type monitorLogProfiles struct {
	client rest.Interface
	ns     string
}

// newMonitorLogProfiles returns a MonitorLogProfiles
func newMonitorLogProfiles(c *AzurermV1alpha1Client, namespace string) *monitorLogProfiles {
	return &monitorLogProfiles{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the monitorLogProfile, and returns the corresponding monitorLogProfile object, and an error if there is any.
func (c *monitorLogProfiles) Get(name string, options v1.GetOptions) (result *v1alpha1.MonitorLogProfile, err error) {
	result = &v1alpha1.MonitorLogProfile{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("monitorlogprofiles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MonitorLogProfiles that match those selectors.
func (c *monitorLogProfiles) List(opts v1.ListOptions) (result *v1alpha1.MonitorLogProfileList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MonitorLogProfileList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("monitorlogprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested monitorLogProfiles.
func (c *monitorLogProfiles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("monitorlogprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a monitorLogProfile and creates it.  Returns the server's representation of the monitorLogProfile, and an error, if there is any.
func (c *monitorLogProfiles) Create(monitorLogProfile *v1alpha1.MonitorLogProfile) (result *v1alpha1.MonitorLogProfile, err error) {
	result = &v1alpha1.MonitorLogProfile{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("monitorlogprofiles").
		Body(monitorLogProfile).
		Do().
		Into(result)
	return
}

// Update takes the representation of a monitorLogProfile and updates it. Returns the server's representation of the monitorLogProfile, and an error, if there is any.
func (c *monitorLogProfiles) Update(monitorLogProfile *v1alpha1.MonitorLogProfile) (result *v1alpha1.MonitorLogProfile, err error) {
	result = &v1alpha1.MonitorLogProfile{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("monitorlogprofiles").
		Name(monitorLogProfile.Name).
		Body(monitorLogProfile).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *monitorLogProfiles) UpdateStatus(monitorLogProfile *v1alpha1.MonitorLogProfile) (result *v1alpha1.MonitorLogProfile, err error) {
	result = &v1alpha1.MonitorLogProfile{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("monitorlogprofiles").
		Name(monitorLogProfile.Name).
		SubResource("status").
		Body(monitorLogProfile).
		Do().
		Into(result)
	return
}

// Delete takes name of the monitorLogProfile and deletes it. Returns an error if one occurs.
func (c *monitorLogProfiles) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("monitorlogprofiles").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *monitorLogProfiles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("monitorlogprofiles").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched monitorLogProfile.
func (c *monitorLogProfiles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MonitorLogProfile, err error) {
	result = &v1alpha1.MonitorLogProfile{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("monitorlogprofiles").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
