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

// GoogleStorageDefaultObjectAclsGetter has a method to return a GoogleStorageDefaultObjectAclInterface.
// A group's client should implement this interface.
type GoogleStorageDefaultObjectAclsGetter interface {
	GoogleStorageDefaultObjectAcls() GoogleStorageDefaultObjectAclInterface
}

// GoogleStorageDefaultObjectAclInterface has methods to work with GoogleStorageDefaultObjectAcl resources.
type GoogleStorageDefaultObjectAclInterface interface {
	Create(*v1alpha1.GoogleStorageDefaultObjectAcl) (*v1alpha1.GoogleStorageDefaultObjectAcl, error)
	Update(*v1alpha1.GoogleStorageDefaultObjectAcl) (*v1alpha1.GoogleStorageDefaultObjectAcl, error)
	UpdateStatus(*v1alpha1.GoogleStorageDefaultObjectAcl) (*v1alpha1.GoogleStorageDefaultObjectAcl, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GoogleStorageDefaultObjectAcl, error)
	List(opts v1.ListOptions) (*v1alpha1.GoogleStorageDefaultObjectAclList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleStorageDefaultObjectAcl, err error)
	GoogleStorageDefaultObjectAclExpansion
}

// googleStorageDefaultObjectAcls implements GoogleStorageDefaultObjectAclInterface
type googleStorageDefaultObjectAcls struct {
	client rest.Interface
}

// newGoogleStorageDefaultObjectAcls returns a GoogleStorageDefaultObjectAcls
func newGoogleStorageDefaultObjectAcls(c *GoogleV1alpha1Client) *googleStorageDefaultObjectAcls {
	return &googleStorageDefaultObjectAcls{
		client: c.RESTClient(),
	}
}

// Get takes name of the googleStorageDefaultObjectAcl, and returns the corresponding googleStorageDefaultObjectAcl object, and an error if there is any.
func (c *googleStorageDefaultObjectAcls) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleStorageDefaultObjectAcl, err error) {
	result = &v1alpha1.GoogleStorageDefaultObjectAcl{}
	err = c.client.Get().
		Resource("googlestoragedefaultobjectacls").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GoogleStorageDefaultObjectAcls that match those selectors.
func (c *googleStorageDefaultObjectAcls) List(opts v1.ListOptions) (result *v1alpha1.GoogleStorageDefaultObjectAclList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GoogleStorageDefaultObjectAclList{}
	err = c.client.Get().
		Resource("googlestoragedefaultobjectacls").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested googleStorageDefaultObjectAcls.
func (c *googleStorageDefaultObjectAcls) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("googlestoragedefaultobjectacls").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a googleStorageDefaultObjectAcl and creates it.  Returns the server's representation of the googleStorageDefaultObjectAcl, and an error, if there is any.
func (c *googleStorageDefaultObjectAcls) Create(googleStorageDefaultObjectAcl *v1alpha1.GoogleStorageDefaultObjectAcl) (result *v1alpha1.GoogleStorageDefaultObjectAcl, err error) {
	result = &v1alpha1.GoogleStorageDefaultObjectAcl{}
	err = c.client.Post().
		Resource("googlestoragedefaultobjectacls").
		Body(googleStorageDefaultObjectAcl).
		Do().
		Into(result)
	return
}

// Update takes the representation of a googleStorageDefaultObjectAcl and updates it. Returns the server's representation of the googleStorageDefaultObjectAcl, and an error, if there is any.
func (c *googleStorageDefaultObjectAcls) Update(googleStorageDefaultObjectAcl *v1alpha1.GoogleStorageDefaultObjectAcl) (result *v1alpha1.GoogleStorageDefaultObjectAcl, err error) {
	result = &v1alpha1.GoogleStorageDefaultObjectAcl{}
	err = c.client.Put().
		Resource("googlestoragedefaultobjectacls").
		Name(googleStorageDefaultObjectAcl.Name).
		Body(googleStorageDefaultObjectAcl).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *googleStorageDefaultObjectAcls) UpdateStatus(googleStorageDefaultObjectAcl *v1alpha1.GoogleStorageDefaultObjectAcl) (result *v1alpha1.GoogleStorageDefaultObjectAcl, err error) {
	result = &v1alpha1.GoogleStorageDefaultObjectAcl{}
	err = c.client.Put().
		Resource("googlestoragedefaultobjectacls").
		Name(googleStorageDefaultObjectAcl.Name).
		SubResource("status").
		Body(googleStorageDefaultObjectAcl).
		Do().
		Into(result)
	return
}

// Delete takes name of the googleStorageDefaultObjectAcl and deletes it. Returns an error if one occurs.
func (c *googleStorageDefaultObjectAcls) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("googlestoragedefaultobjectacls").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *googleStorageDefaultObjectAcls) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("googlestoragedefaultobjectacls").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched googleStorageDefaultObjectAcl.
func (c *googleStorageDefaultObjectAcls) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleStorageDefaultObjectAcl, err error) {
	result = &v1alpha1.GoogleStorageDefaultObjectAcl{}
	err = c.client.Patch(pt).
		Resource("googlestoragedefaultobjectacls").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
