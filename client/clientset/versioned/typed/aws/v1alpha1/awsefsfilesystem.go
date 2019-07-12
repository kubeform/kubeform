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

// AwsEfsFileSystemsGetter has a method to return a AwsEfsFileSystemInterface.
// A group's client should implement this interface.
type AwsEfsFileSystemsGetter interface {
	AwsEfsFileSystems() AwsEfsFileSystemInterface
}

// AwsEfsFileSystemInterface has methods to work with AwsEfsFileSystem resources.
type AwsEfsFileSystemInterface interface {
	Create(*v1alpha1.AwsEfsFileSystem) (*v1alpha1.AwsEfsFileSystem, error)
	Update(*v1alpha1.AwsEfsFileSystem) (*v1alpha1.AwsEfsFileSystem, error)
	UpdateStatus(*v1alpha1.AwsEfsFileSystem) (*v1alpha1.AwsEfsFileSystem, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsEfsFileSystem, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsEfsFileSystemList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsEfsFileSystem, err error)
	AwsEfsFileSystemExpansion
}

// awsEfsFileSystems implements AwsEfsFileSystemInterface
type awsEfsFileSystems struct {
	client rest.Interface
}

// newAwsEfsFileSystems returns a AwsEfsFileSystems
func newAwsEfsFileSystems(c *AwsV1alpha1Client) *awsEfsFileSystems {
	return &awsEfsFileSystems{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsEfsFileSystem, and returns the corresponding awsEfsFileSystem object, and an error if there is any.
func (c *awsEfsFileSystems) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsEfsFileSystem, err error) {
	result = &v1alpha1.AwsEfsFileSystem{}
	err = c.client.Get().
		Resource("awsefsfilesystems").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsEfsFileSystems that match those selectors.
func (c *awsEfsFileSystems) List(opts v1.ListOptions) (result *v1alpha1.AwsEfsFileSystemList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsEfsFileSystemList{}
	err = c.client.Get().
		Resource("awsefsfilesystems").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsEfsFileSystems.
func (c *awsEfsFileSystems) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsefsfilesystems").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsEfsFileSystem and creates it.  Returns the server's representation of the awsEfsFileSystem, and an error, if there is any.
func (c *awsEfsFileSystems) Create(awsEfsFileSystem *v1alpha1.AwsEfsFileSystem) (result *v1alpha1.AwsEfsFileSystem, err error) {
	result = &v1alpha1.AwsEfsFileSystem{}
	err = c.client.Post().
		Resource("awsefsfilesystems").
		Body(awsEfsFileSystem).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsEfsFileSystem and updates it. Returns the server's representation of the awsEfsFileSystem, and an error, if there is any.
func (c *awsEfsFileSystems) Update(awsEfsFileSystem *v1alpha1.AwsEfsFileSystem) (result *v1alpha1.AwsEfsFileSystem, err error) {
	result = &v1alpha1.AwsEfsFileSystem{}
	err = c.client.Put().
		Resource("awsefsfilesystems").
		Name(awsEfsFileSystem.Name).
		Body(awsEfsFileSystem).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsEfsFileSystems) UpdateStatus(awsEfsFileSystem *v1alpha1.AwsEfsFileSystem) (result *v1alpha1.AwsEfsFileSystem, err error) {
	result = &v1alpha1.AwsEfsFileSystem{}
	err = c.client.Put().
		Resource("awsefsfilesystems").
		Name(awsEfsFileSystem.Name).
		SubResource("status").
		Body(awsEfsFileSystem).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsEfsFileSystem and deletes it. Returns an error if one occurs.
func (c *awsEfsFileSystems) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsefsfilesystems").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsEfsFileSystems) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsefsfilesystems").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsEfsFileSystem.
func (c *awsEfsFileSystems) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsEfsFileSystem, err error) {
	result = &v1alpha1.AwsEfsFileSystem{}
	err = c.client.Patch(pt).
		Resource("awsefsfilesystems").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
