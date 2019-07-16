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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakePinpointApnsVoipSandboxChannels implements PinpointApnsVoipSandboxChannelInterface
type FakePinpointApnsVoipSandboxChannels struct {
	Fake *FakeAwsV1alpha1
}

var pinpointapnsvoipsandboxchannelsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "pinpointapnsvoipsandboxchannels"}

var pinpointapnsvoipsandboxchannelsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "PinpointApnsVoipSandboxChannel"}

// Get takes name of the pinpointApnsVoipSandboxChannel, and returns the corresponding pinpointApnsVoipSandboxChannel object, and an error if there is any.
func (c *FakePinpointApnsVoipSandboxChannels) Get(name string, options v1.GetOptions) (result *v1alpha1.PinpointApnsVoipSandboxChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(pinpointapnsvoipsandboxchannelsResource, name), &v1alpha1.PinpointApnsVoipSandboxChannel{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointApnsVoipSandboxChannel), err
}

// List takes label and field selectors, and returns the list of PinpointApnsVoipSandboxChannels that match those selectors.
func (c *FakePinpointApnsVoipSandboxChannels) List(opts v1.ListOptions) (result *v1alpha1.PinpointApnsVoipSandboxChannelList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(pinpointapnsvoipsandboxchannelsResource, pinpointapnsvoipsandboxchannelsKind, opts), &v1alpha1.PinpointApnsVoipSandboxChannelList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PinpointApnsVoipSandboxChannelList{ListMeta: obj.(*v1alpha1.PinpointApnsVoipSandboxChannelList).ListMeta}
	for _, item := range obj.(*v1alpha1.PinpointApnsVoipSandboxChannelList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pinpointApnsVoipSandboxChannels.
func (c *FakePinpointApnsVoipSandboxChannels) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(pinpointapnsvoipsandboxchannelsResource, opts))
}

// Create takes the representation of a pinpointApnsVoipSandboxChannel and creates it.  Returns the server's representation of the pinpointApnsVoipSandboxChannel, and an error, if there is any.
func (c *FakePinpointApnsVoipSandboxChannels) Create(pinpointApnsVoipSandboxChannel *v1alpha1.PinpointApnsVoipSandboxChannel) (result *v1alpha1.PinpointApnsVoipSandboxChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(pinpointapnsvoipsandboxchannelsResource, pinpointApnsVoipSandboxChannel), &v1alpha1.PinpointApnsVoipSandboxChannel{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointApnsVoipSandboxChannel), err
}

// Update takes the representation of a pinpointApnsVoipSandboxChannel and updates it. Returns the server's representation of the pinpointApnsVoipSandboxChannel, and an error, if there is any.
func (c *FakePinpointApnsVoipSandboxChannels) Update(pinpointApnsVoipSandboxChannel *v1alpha1.PinpointApnsVoipSandboxChannel) (result *v1alpha1.PinpointApnsVoipSandboxChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(pinpointapnsvoipsandboxchannelsResource, pinpointApnsVoipSandboxChannel), &v1alpha1.PinpointApnsVoipSandboxChannel{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointApnsVoipSandboxChannel), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePinpointApnsVoipSandboxChannels) UpdateStatus(pinpointApnsVoipSandboxChannel *v1alpha1.PinpointApnsVoipSandboxChannel) (*v1alpha1.PinpointApnsVoipSandboxChannel, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(pinpointapnsvoipsandboxchannelsResource, "status", pinpointApnsVoipSandboxChannel), &v1alpha1.PinpointApnsVoipSandboxChannel{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointApnsVoipSandboxChannel), err
}

// Delete takes name of the pinpointApnsVoipSandboxChannel and deletes it. Returns an error if one occurs.
func (c *FakePinpointApnsVoipSandboxChannels) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(pinpointapnsvoipsandboxchannelsResource, name), &v1alpha1.PinpointApnsVoipSandboxChannel{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePinpointApnsVoipSandboxChannels) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(pinpointapnsvoipsandboxchannelsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PinpointApnsVoipSandboxChannelList{})
	return err
}

// Patch applies the patch and returns the patched pinpointApnsVoipSandboxChannel.
func (c *FakePinpointApnsVoipSandboxChannels) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PinpointApnsVoipSandboxChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(pinpointapnsvoipsandboxchannelsResource, name, pt, data, subresources...), &v1alpha1.PinpointApnsVoipSandboxChannel{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointApnsVoipSandboxChannel), err
}
