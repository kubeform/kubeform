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

// S3BucketNotificationsGetter has a method to return a S3BucketNotificationInterface.
// A group's client should implement this interface.
type S3BucketNotificationsGetter interface {
	S3BucketNotifications(namespace string) S3BucketNotificationInterface
}

// S3BucketNotificationInterface has methods to work with S3BucketNotification resources.
type S3BucketNotificationInterface interface {
	Create(*v1alpha1.S3BucketNotification) (*v1alpha1.S3BucketNotification, error)
	Update(*v1alpha1.S3BucketNotification) (*v1alpha1.S3BucketNotification, error)
	UpdateStatus(*v1alpha1.S3BucketNotification) (*v1alpha1.S3BucketNotification, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.S3BucketNotification, error)
	List(opts v1.ListOptions) (*v1alpha1.S3BucketNotificationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.S3BucketNotification, err error)
	S3BucketNotificationExpansion
}

// s3BucketNotifications implements S3BucketNotificationInterface
type s3BucketNotifications struct {
	client rest.Interface
	ns     string
}

// newS3BucketNotifications returns a S3BucketNotifications
func newS3BucketNotifications(c *AwsV1alpha1Client, namespace string) *s3BucketNotifications {
	return &s3BucketNotifications{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the s3BucketNotification, and returns the corresponding s3BucketNotification object, and an error if there is any.
func (c *s3BucketNotifications) Get(name string, options v1.GetOptions) (result *v1alpha1.S3BucketNotification, err error) {
	result = &v1alpha1.S3BucketNotification{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("s3bucketnotifications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of S3BucketNotifications that match those selectors.
func (c *s3BucketNotifications) List(opts v1.ListOptions) (result *v1alpha1.S3BucketNotificationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.S3BucketNotificationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("s3bucketnotifications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested s3BucketNotifications.
func (c *s3BucketNotifications) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("s3bucketnotifications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a s3BucketNotification and creates it.  Returns the server's representation of the s3BucketNotification, and an error, if there is any.
func (c *s3BucketNotifications) Create(s3BucketNotification *v1alpha1.S3BucketNotification) (result *v1alpha1.S3BucketNotification, err error) {
	result = &v1alpha1.S3BucketNotification{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("s3bucketnotifications").
		Body(s3BucketNotification).
		Do().
		Into(result)
	return
}

// Update takes the representation of a s3BucketNotification and updates it. Returns the server's representation of the s3BucketNotification, and an error, if there is any.
func (c *s3BucketNotifications) Update(s3BucketNotification *v1alpha1.S3BucketNotification) (result *v1alpha1.S3BucketNotification, err error) {
	result = &v1alpha1.S3BucketNotification{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("s3bucketnotifications").
		Name(s3BucketNotification.Name).
		Body(s3BucketNotification).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *s3BucketNotifications) UpdateStatus(s3BucketNotification *v1alpha1.S3BucketNotification) (result *v1alpha1.S3BucketNotification, err error) {
	result = &v1alpha1.S3BucketNotification{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("s3bucketnotifications").
		Name(s3BucketNotification.Name).
		SubResource("status").
		Body(s3BucketNotification).
		Do().
		Into(result)
	return
}

// Delete takes name of the s3BucketNotification and deletes it. Returns an error if one occurs.
func (c *s3BucketNotifications) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("s3bucketnotifications").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *s3BucketNotifications) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("s3bucketnotifications").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched s3BucketNotification.
func (c *s3BucketNotifications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.S3BucketNotification, err error) {
	result = &v1alpha1.S3BucketNotification{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("s3bucketnotifications").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
