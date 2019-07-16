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

// IamUserLoginProfilesGetter has a method to return a IamUserLoginProfileInterface.
// A group's client should implement this interface.
type IamUserLoginProfilesGetter interface {
	IamUserLoginProfiles() IamUserLoginProfileInterface
}

// IamUserLoginProfileInterface has methods to work with IamUserLoginProfile resources.
type IamUserLoginProfileInterface interface {
	Create(*v1alpha1.IamUserLoginProfile) (*v1alpha1.IamUserLoginProfile, error)
	Update(*v1alpha1.IamUserLoginProfile) (*v1alpha1.IamUserLoginProfile, error)
	UpdateStatus(*v1alpha1.IamUserLoginProfile) (*v1alpha1.IamUserLoginProfile, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.IamUserLoginProfile, error)
	List(opts v1.ListOptions) (*v1alpha1.IamUserLoginProfileList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IamUserLoginProfile, err error)
	IamUserLoginProfileExpansion
}

// iamUserLoginProfiles implements IamUserLoginProfileInterface
type iamUserLoginProfiles struct {
	client rest.Interface
}

// newIamUserLoginProfiles returns a IamUserLoginProfiles
func newIamUserLoginProfiles(c *AwsV1alpha1Client) *iamUserLoginProfiles {
	return &iamUserLoginProfiles{
		client: c.RESTClient(),
	}
}

// Get takes name of the iamUserLoginProfile, and returns the corresponding iamUserLoginProfile object, and an error if there is any.
func (c *iamUserLoginProfiles) Get(name string, options v1.GetOptions) (result *v1alpha1.IamUserLoginProfile, err error) {
	result = &v1alpha1.IamUserLoginProfile{}
	err = c.client.Get().
		Resource("iamuserloginprofiles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IamUserLoginProfiles that match those selectors.
func (c *iamUserLoginProfiles) List(opts v1.ListOptions) (result *v1alpha1.IamUserLoginProfileList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.IamUserLoginProfileList{}
	err = c.client.Get().
		Resource("iamuserloginprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iamUserLoginProfiles.
func (c *iamUserLoginProfiles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("iamuserloginprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a iamUserLoginProfile and creates it.  Returns the server's representation of the iamUserLoginProfile, and an error, if there is any.
func (c *iamUserLoginProfiles) Create(iamUserLoginProfile *v1alpha1.IamUserLoginProfile) (result *v1alpha1.IamUserLoginProfile, err error) {
	result = &v1alpha1.IamUserLoginProfile{}
	err = c.client.Post().
		Resource("iamuserloginprofiles").
		Body(iamUserLoginProfile).
		Do().
		Into(result)
	return
}

// Update takes the representation of a iamUserLoginProfile and updates it. Returns the server's representation of the iamUserLoginProfile, and an error, if there is any.
func (c *iamUserLoginProfiles) Update(iamUserLoginProfile *v1alpha1.IamUserLoginProfile) (result *v1alpha1.IamUserLoginProfile, err error) {
	result = &v1alpha1.IamUserLoginProfile{}
	err = c.client.Put().
		Resource("iamuserloginprofiles").
		Name(iamUserLoginProfile.Name).
		Body(iamUserLoginProfile).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *iamUserLoginProfiles) UpdateStatus(iamUserLoginProfile *v1alpha1.IamUserLoginProfile) (result *v1alpha1.IamUserLoginProfile, err error) {
	result = &v1alpha1.IamUserLoginProfile{}
	err = c.client.Put().
		Resource("iamuserloginprofiles").
		Name(iamUserLoginProfile.Name).
		SubResource("status").
		Body(iamUserLoginProfile).
		Do().
		Into(result)
	return
}

// Delete takes name of the iamUserLoginProfile and deletes it. Returns an error if one occurs.
func (c *iamUserLoginProfiles) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("iamuserloginprofiles").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *iamUserLoginProfiles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("iamuserloginprofiles").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched iamUserLoginProfile.
func (c *iamUserLoginProfiles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IamUserLoginProfile, err error) {
	result = &v1alpha1.IamUserLoginProfile{}
	err = c.client.Patch(pt).
		Resource("iamuserloginprofiles").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
