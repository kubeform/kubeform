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

// PinpointApnsVoipSandboxChannelsGetter has a method to return a PinpointApnsVoipSandboxChannelInterface.
// A group's client should implement this interface.
type PinpointApnsVoipSandboxChannelsGetter interface {
	PinpointApnsVoipSandboxChannels() PinpointApnsVoipSandboxChannelInterface
}

// PinpointApnsVoipSandboxChannelInterface has methods to work with PinpointApnsVoipSandboxChannel resources.
type PinpointApnsVoipSandboxChannelInterface interface {
	Create(*v1alpha1.PinpointApnsVoipSandboxChannel) (*v1alpha1.PinpointApnsVoipSandboxChannel, error)
	Update(*v1alpha1.PinpointApnsVoipSandboxChannel) (*v1alpha1.PinpointApnsVoipSandboxChannel, error)
	UpdateStatus(*v1alpha1.PinpointApnsVoipSandboxChannel) (*v1alpha1.PinpointApnsVoipSandboxChannel, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.PinpointApnsVoipSandboxChannel, error)
	List(opts v1.ListOptions) (*v1alpha1.PinpointApnsVoipSandboxChannelList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PinpointApnsVoipSandboxChannel, err error)
	PinpointApnsVoipSandboxChannelExpansion
}

// pinpointApnsVoipSandboxChannels implements PinpointApnsVoipSandboxChannelInterface
type pinpointApnsVoipSandboxChannels struct {
	client rest.Interface
}

// newPinpointApnsVoipSandboxChannels returns a PinpointApnsVoipSandboxChannels
func newPinpointApnsVoipSandboxChannels(c *AwsV1alpha1Client) *pinpointApnsVoipSandboxChannels {
	return &pinpointApnsVoipSandboxChannels{
		client: c.RESTClient(),
	}
}

// Get takes name of the pinpointApnsVoipSandboxChannel, and returns the corresponding pinpointApnsVoipSandboxChannel object, and an error if there is any.
func (c *pinpointApnsVoipSandboxChannels) Get(name string, options v1.GetOptions) (result *v1alpha1.PinpointApnsVoipSandboxChannel, err error) {
	result = &v1alpha1.PinpointApnsVoipSandboxChannel{}
	err = c.client.Get().
		Resource("pinpointapnsvoipsandboxchannels").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PinpointApnsVoipSandboxChannels that match those selectors.
func (c *pinpointApnsVoipSandboxChannels) List(opts v1.ListOptions) (result *v1alpha1.PinpointApnsVoipSandboxChannelList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PinpointApnsVoipSandboxChannelList{}
	err = c.client.Get().
		Resource("pinpointapnsvoipsandboxchannels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested pinpointApnsVoipSandboxChannels.
func (c *pinpointApnsVoipSandboxChannels) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("pinpointapnsvoipsandboxchannels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a pinpointApnsVoipSandboxChannel and creates it.  Returns the server's representation of the pinpointApnsVoipSandboxChannel, and an error, if there is any.
func (c *pinpointApnsVoipSandboxChannels) Create(pinpointApnsVoipSandboxChannel *v1alpha1.PinpointApnsVoipSandboxChannel) (result *v1alpha1.PinpointApnsVoipSandboxChannel, err error) {
	result = &v1alpha1.PinpointApnsVoipSandboxChannel{}
	err = c.client.Post().
		Resource("pinpointapnsvoipsandboxchannels").
		Body(pinpointApnsVoipSandboxChannel).
		Do().
		Into(result)
	return
}

// Update takes the representation of a pinpointApnsVoipSandboxChannel and updates it. Returns the server's representation of the pinpointApnsVoipSandboxChannel, and an error, if there is any.
func (c *pinpointApnsVoipSandboxChannels) Update(pinpointApnsVoipSandboxChannel *v1alpha1.PinpointApnsVoipSandboxChannel) (result *v1alpha1.PinpointApnsVoipSandboxChannel, err error) {
	result = &v1alpha1.PinpointApnsVoipSandboxChannel{}
	err = c.client.Put().
		Resource("pinpointapnsvoipsandboxchannels").
		Name(pinpointApnsVoipSandboxChannel.Name).
		Body(pinpointApnsVoipSandboxChannel).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *pinpointApnsVoipSandboxChannels) UpdateStatus(pinpointApnsVoipSandboxChannel *v1alpha1.PinpointApnsVoipSandboxChannel) (result *v1alpha1.PinpointApnsVoipSandboxChannel, err error) {
	result = &v1alpha1.PinpointApnsVoipSandboxChannel{}
	err = c.client.Put().
		Resource("pinpointapnsvoipsandboxchannels").
		Name(pinpointApnsVoipSandboxChannel.Name).
		SubResource("status").
		Body(pinpointApnsVoipSandboxChannel).
		Do().
		Into(result)
	return
}

// Delete takes name of the pinpointApnsVoipSandboxChannel and deletes it. Returns an error if one occurs.
func (c *pinpointApnsVoipSandboxChannels) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("pinpointapnsvoipsandboxchannels").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *pinpointApnsVoipSandboxChannels) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("pinpointapnsvoipsandboxchannels").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched pinpointApnsVoipSandboxChannel.
func (c *pinpointApnsVoipSandboxChannels) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PinpointApnsVoipSandboxChannel, err error) {
	result = &v1alpha1.PinpointApnsVoipSandboxChannel{}
	err = c.client.Patch(pt).
		Resource("pinpointapnsvoipsandboxchannels").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
