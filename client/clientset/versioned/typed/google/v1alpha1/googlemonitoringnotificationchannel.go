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

// GoogleMonitoringNotificationChannelsGetter has a method to return a GoogleMonitoringNotificationChannelInterface.
// A group's client should implement this interface.
type GoogleMonitoringNotificationChannelsGetter interface {
	GoogleMonitoringNotificationChannels() GoogleMonitoringNotificationChannelInterface
}

// GoogleMonitoringNotificationChannelInterface has methods to work with GoogleMonitoringNotificationChannel resources.
type GoogleMonitoringNotificationChannelInterface interface {
	Create(*v1alpha1.GoogleMonitoringNotificationChannel) (*v1alpha1.GoogleMonitoringNotificationChannel, error)
	Update(*v1alpha1.GoogleMonitoringNotificationChannel) (*v1alpha1.GoogleMonitoringNotificationChannel, error)
	UpdateStatus(*v1alpha1.GoogleMonitoringNotificationChannel) (*v1alpha1.GoogleMonitoringNotificationChannel, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GoogleMonitoringNotificationChannel, error)
	List(opts v1.ListOptions) (*v1alpha1.GoogleMonitoringNotificationChannelList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleMonitoringNotificationChannel, err error)
	GoogleMonitoringNotificationChannelExpansion
}

// googleMonitoringNotificationChannels implements GoogleMonitoringNotificationChannelInterface
type googleMonitoringNotificationChannels struct {
	client rest.Interface
}

// newGoogleMonitoringNotificationChannels returns a GoogleMonitoringNotificationChannels
func newGoogleMonitoringNotificationChannels(c *GoogleV1alpha1Client) *googleMonitoringNotificationChannels {
	return &googleMonitoringNotificationChannels{
		client: c.RESTClient(),
	}
}

// Get takes name of the googleMonitoringNotificationChannel, and returns the corresponding googleMonitoringNotificationChannel object, and an error if there is any.
func (c *googleMonitoringNotificationChannels) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleMonitoringNotificationChannel, err error) {
	result = &v1alpha1.GoogleMonitoringNotificationChannel{}
	err = c.client.Get().
		Resource("googlemonitoringnotificationchannels").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GoogleMonitoringNotificationChannels that match those selectors.
func (c *googleMonitoringNotificationChannels) List(opts v1.ListOptions) (result *v1alpha1.GoogleMonitoringNotificationChannelList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GoogleMonitoringNotificationChannelList{}
	err = c.client.Get().
		Resource("googlemonitoringnotificationchannels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested googleMonitoringNotificationChannels.
func (c *googleMonitoringNotificationChannels) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("googlemonitoringnotificationchannels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a googleMonitoringNotificationChannel and creates it.  Returns the server's representation of the googleMonitoringNotificationChannel, and an error, if there is any.
func (c *googleMonitoringNotificationChannels) Create(googleMonitoringNotificationChannel *v1alpha1.GoogleMonitoringNotificationChannel) (result *v1alpha1.GoogleMonitoringNotificationChannel, err error) {
	result = &v1alpha1.GoogleMonitoringNotificationChannel{}
	err = c.client.Post().
		Resource("googlemonitoringnotificationchannels").
		Body(googleMonitoringNotificationChannel).
		Do().
		Into(result)
	return
}

// Update takes the representation of a googleMonitoringNotificationChannel and updates it. Returns the server's representation of the googleMonitoringNotificationChannel, and an error, if there is any.
func (c *googleMonitoringNotificationChannels) Update(googleMonitoringNotificationChannel *v1alpha1.GoogleMonitoringNotificationChannel) (result *v1alpha1.GoogleMonitoringNotificationChannel, err error) {
	result = &v1alpha1.GoogleMonitoringNotificationChannel{}
	err = c.client.Put().
		Resource("googlemonitoringnotificationchannels").
		Name(googleMonitoringNotificationChannel.Name).
		Body(googleMonitoringNotificationChannel).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *googleMonitoringNotificationChannels) UpdateStatus(googleMonitoringNotificationChannel *v1alpha1.GoogleMonitoringNotificationChannel) (result *v1alpha1.GoogleMonitoringNotificationChannel, err error) {
	result = &v1alpha1.GoogleMonitoringNotificationChannel{}
	err = c.client.Put().
		Resource("googlemonitoringnotificationchannels").
		Name(googleMonitoringNotificationChannel.Name).
		SubResource("status").
		Body(googleMonitoringNotificationChannel).
		Do().
		Into(result)
	return
}

// Delete takes name of the googleMonitoringNotificationChannel and deletes it. Returns an error if one occurs.
func (c *googleMonitoringNotificationChannels) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("googlemonitoringnotificationchannels").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *googleMonitoringNotificationChannels) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("googlemonitoringnotificationchannels").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched googleMonitoringNotificationChannel.
func (c *googleMonitoringNotificationChannels) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleMonitoringNotificationChannel, err error) {
	result = &v1alpha1.GoogleMonitoringNotificationChannel{}
	err = c.client.Patch(pt).
		Resource("googlemonitoringnotificationchannels").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
