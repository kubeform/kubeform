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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// AwsS3BucketObjectsGetter has a method to return a AwsS3BucketObjectInterface.
// A group's client should implement this interface.
type AwsS3BucketObjectsGetter interface {
	AwsS3BucketObjects() AwsS3BucketObjectInterface
}

// AwsS3BucketObjectInterface has methods to work with AwsS3BucketObject resources.
type AwsS3BucketObjectInterface interface {
	Create(*v1alpha1.AwsS3BucketObject) (*v1alpha1.AwsS3BucketObject, error)
	Update(*v1alpha1.AwsS3BucketObject) (*v1alpha1.AwsS3BucketObject, error)
	UpdateStatus(*v1alpha1.AwsS3BucketObject) (*v1alpha1.AwsS3BucketObject, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsS3BucketObject, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsS3BucketObjectList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsS3BucketObject, err error)
	AwsS3BucketObjectExpansion
}

// awsS3BucketObjects implements AwsS3BucketObjectInterface
type awsS3BucketObjects struct {
	client rest.Interface
}

// newAwsS3BucketObjects returns a AwsS3BucketObjects
func newAwsS3BucketObjects(c *AwsV1alpha1Client) *awsS3BucketObjects {
	return &awsS3BucketObjects{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsS3BucketObject, and returns the corresponding awsS3BucketObject object, and an error if there is any.
func (c *awsS3BucketObjects) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsS3BucketObject, err error) {
	result = &v1alpha1.AwsS3BucketObject{}
	err = c.client.Get().
		Resource("awss3bucketobjects").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsS3BucketObjects that match those selectors.
func (c *awsS3BucketObjects) List(opts v1.ListOptions) (result *v1alpha1.AwsS3BucketObjectList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsS3BucketObjectList{}
	err = c.client.Get().
		Resource("awss3bucketobjects").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsS3BucketObjects.
func (c *awsS3BucketObjects) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awss3bucketobjects").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsS3BucketObject and creates it.  Returns the server's representation of the awsS3BucketObject, and an error, if there is any.
func (c *awsS3BucketObjects) Create(awsS3BucketObject *v1alpha1.AwsS3BucketObject) (result *v1alpha1.AwsS3BucketObject, err error) {
	result = &v1alpha1.AwsS3BucketObject{}
	err = c.client.Post().
		Resource("awss3bucketobjects").
		Body(awsS3BucketObject).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsS3BucketObject and updates it. Returns the server's representation of the awsS3BucketObject, and an error, if there is any.
func (c *awsS3BucketObjects) Update(awsS3BucketObject *v1alpha1.AwsS3BucketObject) (result *v1alpha1.AwsS3BucketObject, err error) {
	result = &v1alpha1.AwsS3BucketObject{}
	err = c.client.Put().
		Resource("awss3bucketobjects").
		Name(awsS3BucketObject.Name).
		Body(awsS3BucketObject).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsS3BucketObjects) UpdateStatus(awsS3BucketObject *v1alpha1.AwsS3BucketObject) (result *v1alpha1.AwsS3BucketObject, err error) {
	result = &v1alpha1.AwsS3BucketObject{}
	err = c.client.Put().
		Resource("awss3bucketobjects").
		Name(awsS3BucketObject.Name).
		SubResource("status").
		Body(awsS3BucketObject).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsS3BucketObject and deletes it. Returns an error if one occurs.
func (c *awsS3BucketObjects) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awss3bucketobjects").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsS3BucketObjects) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awss3bucketobjects").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsS3BucketObject.
func (c *awsS3BucketObjects) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsS3BucketObject, err error) {
	result = &v1alpha1.AwsS3BucketObject{}
	err = c.client.Patch(pt).
		Resource("awss3bucketobjects").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
