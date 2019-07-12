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

// GoogleStorageObjectAccessControlsGetter has a method to return a GoogleStorageObjectAccessControlInterface.
// A group's client should implement this interface.
type GoogleStorageObjectAccessControlsGetter interface {
	GoogleStorageObjectAccessControls() GoogleStorageObjectAccessControlInterface
}

// GoogleStorageObjectAccessControlInterface has methods to work with GoogleStorageObjectAccessControl resources.
type GoogleStorageObjectAccessControlInterface interface {
	Create(*v1alpha1.GoogleStorageObjectAccessControl) (*v1alpha1.GoogleStorageObjectAccessControl, error)
	Update(*v1alpha1.GoogleStorageObjectAccessControl) (*v1alpha1.GoogleStorageObjectAccessControl, error)
	UpdateStatus(*v1alpha1.GoogleStorageObjectAccessControl) (*v1alpha1.GoogleStorageObjectAccessControl, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GoogleStorageObjectAccessControl, error)
	List(opts v1.ListOptions) (*v1alpha1.GoogleStorageObjectAccessControlList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleStorageObjectAccessControl, err error)
	GoogleStorageObjectAccessControlExpansion
}

// googleStorageObjectAccessControls implements GoogleStorageObjectAccessControlInterface
type googleStorageObjectAccessControls struct {
	client rest.Interface
}

// newGoogleStorageObjectAccessControls returns a GoogleStorageObjectAccessControls
func newGoogleStorageObjectAccessControls(c *GoogleV1alpha1Client) *googleStorageObjectAccessControls {
	return &googleStorageObjectAccessControls{
		client: c.RESTClient(),
	}
}

// Get takes name of the googleStorageObjectAccessControl, and returns the corresponding googleStorageObjectAccessControl object, and an error if there is any.
func (c *googleStorageObjectAccessControls) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleStorageObjectAccessControl, err error) {
	result = &v1alpha1.GoogleStorageObjectAccessControl{}
	err = c.client.Get().
		Resource("googlestorageobjectaccesscontrols").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GoogleStorageObjectAccessControls that match those selectors.
func (c *googleStorageObjectAccessControls) List(opts v1.ListOptions) (result *v1alpha1.GoogleStorageObjectAccessControlList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GoogleStorageObjectAccessControlList{}
	err = c.client.Get().
		Resource("googlestorageobjectaccesscontrols").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested googleStorageObjectAccessControls.
func (c *googleStorageObjectAccessControls) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("googlestorageobjectaccesscontrols").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a googleStorageObjectAccessControl and creates it.  Returns the server's representation of the googleStorageObjectAccessControl, and an error, if there is any.
func (c *googleStorageObjectAccessControls) Create(googleStorageObjectAccessControl *v1alpha1.GoogleStorageObjectAccessControl) (result *v1alpha1.GoogleStorageObjectAccessControl, err error) {
	result = &v1alpha1.GoogleStorageObjectAccessControl{}
	err = c.client.Post().
		Resource("googlestorageobjectaccesscontrols").
		Body(googleStorageObjectAccessControl).
		Do().
		Into(result)
	return
}

// Update takes the representation of a googleStorageObjectAccessControl and updates it. Returns the server's representation of the googleStorageObjectAccessControl, and an error, if there is any.
func (c *googleStorageObjectAccessControls) Update(googleStorageObjectAccessControl *v1alpha1.GoogleStorageObjectAccessControl) (result *v1alpha1.GoogleStorageObjectAccessControl, err error) {
	result = &v1alpha1.GoogleStorageObjectAccessControl{}
	err = c.client.Put().
		Resource("googlestorageobjectaccesscontrols").
		Name(googleStorageObjectAccessControl.Name).
		Body(googleStorageObjectAccessControl).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *googleStorageObjectAccessControls) UpdateStatus(googleStorageObjectAccessControl *v1alpha1.GoogleStorageObjectAccessControl) (result *v1alpha1.GoogleStorageObjectAccessControl, err error) {
	result = &v1alpha1.GoogleStorageObjectAccessControl{}
	err = c.client.Put().
		Resource("googlestorageobjectaccesscontrols").
		Name(googleStorageObjectAccessControl.Name).
		SubResource("status").
		Body(googleStorageObjectAccessControl).
		Do().
		Into(result)
	return
}

// Delete takes name of the googleStorageObjectAccessControl and deletes it. Returns an error if one occurs.
func (c *googleStorageObjectAccessControls) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("googlestorageobjectaccesscontrols").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *googleStorageObjectAccessControls) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("googlestorageobjectaccesscontrols").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched googleStorageObjectAccessControl.
func (c *googleStorageObjectAccessControls) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleStorageObjectAccessControl, err error) {
	result = &v1alpha1.GoogleStorageObjectAccessControl{}
	err = c.client.Patch(pt).
		Resource("googlestorageobjectaccesscontrols").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
