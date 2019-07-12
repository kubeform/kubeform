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

// AwsVpcEndpointConnectionNotificationsGetter has a method to return a AwsVpcEndpointConnectionNotificationInterface.
// A group's client should implement this interface.
type AwsVpcEndpointConnectionNotificationsGetter interface {
	AwsVpcEndpointConnectionNotifications() AwsVpcEndpointConnectionNotificationInterface
}

// AwsVpcEndpointConnectionNotificationInterface has methods to work with AwsVpcEndpointConnectionNotification resources.
type AwsVpcEndpointConnectionNotificationInterface interface {
	Create(*v1alpha1.AwsVpcEndpointConnectionNotification) (*v1alpha1.AwsVpcEndpointConnectionNotification, error)
	Update(*v1alpha1.AwsVpcEndpointConnectionNotification) (*v1alpha1.AwsVpcEndpointConnectionNotification, error)
	UpdateStatus(*v1alpha1.AwsVpcEndpointConnectionNotification) (*v1alpha1.AwsVpcEndpointConnectionNotification, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsVpcEndpointConnectionNotification, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsVpcEndpointConnectionNotificationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsVpcEndpointConnectionNotification, err error)
	AwsVpcEndpointConnectionNotificationExpansion
}

// awsVpcEndpointConnectionNotifications implements AwsVpcEndpointConnectionNotificationInterface
type awsVpcEndpointConnectionNotifications struct {
	client rest.Interface
}

// newAwsVpcEndpointConnectionNotifications returns a AwsVpcEndpointConnectionNotifications
func newAwsVpcEndpointConnectionNotifications(c *AwsV1alpha1Client) *awsVpcEndpointConnectionNotifications {
	return &awsVpcEndpointConnectionNotifications{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsVpcEndpointConnectionNotification, and returns the corresponding awsVpcEndpointConnectionNotification object, and an error if there is any.
func (c *awsVpcEndpointConnectionNotifications) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsVpcEndpointConnectionNotification, err error) {
	result = &v1alpha1.AwsVpcEndpointConnectionNotification{}
	err = c.client.Get().
		Resource("awsvpcendpointconnectionnotifications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsVpcEndpointConnectionNotifications that match those selectors.
func (c *awsVpcEndpointConnectionNotifications) List(opts v1.ListOptions) (result *v1alpha1.AwsVpcEndpointConnectionNotificationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsVpcEndpointConnectionNotificationList{}
	err = c.client.Get().
		Resource("awsvpcendpointconnectionnotifications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsVpcEndpointConnectionNotifications.
func (c *awsVpcEndpointConnectionNotifications) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsvpcendpointconnectionnotifications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsVpcEndpointConnectionNotification and creates it.  Returns the server's representation of the awsVpcEndpointConnectionNotification, and an error, if there is any.
func (c *awsVpcEndpointConnectionNotifications) Create(awsVpcEndpointConnectionNotification *v1alpha1.AwsVpcEndpointConnectionNotification) (result *v1alpha1.AwsVpcEndpointConnectionNotification, err error) {
	result = &v1alpha1.AwsVpcEndpointConnectionNotification{}
	err = c.client.Post().
		Resource("awsvpcendpointconnectionnotifications").
		Body(awsVpcEndpointConnectionNotification).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsVpcEndpointConnectionNotification and updates it. Returns the server's representation of the awsVpcEndpointConnectionNotification, and an error, if there is any.
func (c *awsVpcEndpointConnectionNotifications) Update(awsVpcEndpointConnectionNotification *v1alpha1.AwsVpcEndpointConnectionNotification) (result *v1alpha1.AwsVpcEndpointConnectionNotification, err error) {
	result = &v1alpha1.AwsVpcEndpointConnectionNotification{}
	err = c.client.Put().
		Resource("awsvpcendpointconnectionnotifications").
		Name(awsVpcEndpointConnectionNotification.Name).
		Body(awsVpcEndpointConnectionNotification).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsVpcEndpointConnectionNotifications) UpdateStatus(awsVpcEndpointConnectionNotification *v1alpha1.AwsVpcEndpointConnectionNotification) (result *v1alpha1.AwsVpcEndpointConnectionNotification, err error) {
	result = &v1alpha1.AwsVpcEndpointConnectionNotification{}
	err = c.client.Put().
		Resource("awsvpcendpointconnectionnotifications").
		Name(awsVpcEndpointConnectionNotification.Name).
		SubResource("status").
		Body(awsVpcEndpointConnectionNotification).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsVpcEndpointConnectionNotification and deletes it. Returns an error if one occurs.
func (c *awsVpcEndpointConnectionNotifications) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsvpcendpointconnectionnotifications").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsVpcEndpointConnectionNotifications) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsvpcendpointconnectionnotifications").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsVpcEndpointConnectionNotification.
func (c *awsVpcEndpointConnectionNotifications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsVpcEndpointConnectionNotification, err error) {
	result = &v1alpha1.AwsVpcEndpointConnectionNotification{}
	err = c.client.Patch(pt).
		Resource("awsvpcendpointconnectionnotifications").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
