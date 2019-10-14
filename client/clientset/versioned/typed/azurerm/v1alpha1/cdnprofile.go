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

// CdnProfilesGetter has a method to return a CdnProfileInterface.
// A group's client should implement this interface.
type CdnProfilesGetter interface {
	CdnProfiles(namespace string) CdnProfileInterface
}

// CdnProfileInterface has methods to work with CdnProfile resources.
type CdnProfileInterface interface {
	Create(*v1alpha1.CdnProfile) (*v1alpha1.CdnProfile, error)
	Update(*v1alpha1.CdnProfile) (*v1alpha1.CdnProfile, error)
	UpdateStatus(*v1alpha1.CdnProfile) (*v1alpha1.CdnProfile, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CdnProfile, error)
	List(opts v1.ListOptions) (*v1alpha1.CdnProfileList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CdnProfile, err error)
	CdnProfileExpansion
}

// cdnProfiles implements CdnProfileInterface
type cdnProfiles struct {
	client rest.Interface
	ns     string
}

// newCdnProfiles returns a CdnProfiles
func newCdnProfiles(c *AzurermV1alpha1Client, namespace string) *cdnProfiles {
	return &cdnProfiles{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cdnProfile, and returns the corresponding cdnProfile object, and an error if there is any.
func (c *cdnProfiles) Get(name string, options v1.GetOptions) (result *v1alpha1.CdnProfile, err error) {
	result = &v1alpha1.CdnProfile{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cdnprofiles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CdnProfiles that match those selectors.
func (c *cdnProfiles) List(opts v1.ListOptions) (result *v1alpha1.CdnProfileList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CdnProfileList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cdnprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cdnProfiles.
func (c *cdnProfiles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cdnprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cdnProfile and creates it.  Returns the server's representation of the cdnProfile, and an error, if there is any.
func (c *cdnProfiles) Create(cdnProfile *v1alpha1.CdnProfile) (result *v1alpha1.CdnProfile, err error) {
	result = &v1alpha1.CdnProfile{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cdnprofiles").
		Body(cdnProfile).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cdnProfile and updates it. Returns the server's representation of the cdnProfile, and an error, if there is any.
func (c *cdnProfiles) Update(cdnProfile *v1alpha1.CdnProfile) (result *v1alpha1.CdnProfile, err error) {
	result = &v1alpha1.CdnProfile{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cdnprofiles").
		Name(cdnProfile.Name).
		Body(cdnProfile).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *cdnProfiles) UpdateStatus(cdnProfile *v1alpha1.CdnProfile) (result *v1alpha1.CdnProfile, err error) {
	result = &v1alpha1.CdnProfile{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cdnprofiles").
		Name(cdnProfile.Name).
		SubResource("status").
		Body(cdnProfile).
		Do().
		Into(result)
	return
}

// Delete takes name of the cdnProfile and deletes it. Returns an error if one occurs.
func (c *cdnProfiles) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cdnprofiles").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cdnProfiles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cdnprofiles").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cdnProfile.
func (c *cdnProfiles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CdnProfile, err error) {
	result = &v1alpha1.CdnProfile{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cdnprofiles").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
