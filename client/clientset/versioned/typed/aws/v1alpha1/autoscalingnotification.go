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

// AutoscalingNotificationsGetter has a method to return a AutoscalingNotificationInterface.
// A group's client should implement this interface.
type AutoscalingNotificationsGetter interface {
	AutoscalingNotifications(namespace string) AutoscalingNotificationInterface
}

// AutoscalingNotificationInterface has methods to work with AutoscalingNotification resources.
type AutoscalingNotificationInterface interface {
	Create(*v1alpha1.AutoscalingNotification) (*v1alpha1.AutoscalingNotification, error)
	Update(*v1alpha1.AutoscalingNotification) (*v1alpha1.AutoscalingNotification, error)
	UpdateStatus(*v1alpha1.AutoscalingNotification) (*v1alpha1.AutoscalingNotification, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AutoscalingNotification, error)
	List(opts v1.ListOptions) (*v1alpha1.AutoscalingNotificationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AutoscalingNotification, err error)
	AutoscalingNotificationExpansion
}

// autoscalingNotifications implements AutoscalingNotificationInterface
type autoscalingNotifications struct {
	client rest.Interface
	ns     string
}

// newAutoscalingNotifications returns a AutoscalingNotifications
func newAutoscalingNotifications(c *AwsV1alpha1Client, namespace string) *autoscalingNotifications {
	return &autoscalingNotifications{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the autoscalingNotification, and returns the corresponding autoscalingNotification object, and an error if there is any.
func (c *autoscalingNotifications) Get(name string, options v1.GetOptions) (result *v1alpha1.AutoscalingNotification, err error) {
	result = &v1alpha1.AutoscalingNotification{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("autoscalingnotifications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AutoscalingNotifications that match those selectors.
func (c *autoscalingNotifications) List(opts v1.ListOptions) (result *v1alpha1.AutoscalingNotificationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AutoscalingNotificationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("autoscalingnotifications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested autoscalingNotifications.
func (c *autoscalingNotifications) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("autoscalingnotifications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a autoscalingNotification and creates it.  Returns the server's representation of the autoscalingNotification, and an error, if there is any.
func (c *autoscalingNotifications) Create(autoscalingNotification *v1alpha1.AutoscalingNotification) (result *v1alpha1.AutoscalingNotification, err error) {
	result = &v1alpha1.AutoscalingNotification{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("autoscalingnotifications").
		Body(autoscalingNotification).
		Do().
		Into(result)
	return
}

// Update takes the representation of a autoscalingNotification and updates it. Returns the server's representation of the autoscalingNotification, and an error, if there is any.
func (c *autoscalingNotifications) Update(autoscalingNotification *v1alpha1.AutoscalingNotification) (result *v1alpha1.AutoscalingNotification, err error) {
	result = &v1alpha1.AutoscalingNotification{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("autoscalingnotifications").
		Name(autoscalingNotification.Name).
		Body(autoscalingNotification).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *autoscalingNotifications) UpdateStatus(autoscalingNotification *v1alpha1.AutoscalingNotification) (result *v1alpha1.AutoscalingNotification, err error) {
	result = &v1alpha1.AutoscalingNotification{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("autoscalingnotifications").
		Name(autoscalingNotification.Name).
		SubResource("status").
		Body(autoscalingNotification).
		Do().
		Into(result)
	return
}

// Delete takes name of the autoscalingNotification and deletes it. Returns an error if one occurs.
func (c *autoscalingNotifications) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("autoscalingnotifications").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *autoscalingNotifications) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("autoscalingnotifications").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched autoscalingNotification.
func (c *autoscalingNotifications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AutoscalingNotification, err error) {
	result = &v1alpha1.AutoscalingNotification{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("autoscalingnotifications").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}