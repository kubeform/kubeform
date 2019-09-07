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

// FakePinpointGcmChannels implements PinpointGcmChannelInterface
type FakePinpointGcmChannels struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var pinpointgcmchannelsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "pinpointgcmchannels"}

var pinpointgcmchannelsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "PinpointGcmChannel"}

// Get takes name of the pinpointGcmChannel, and returns the corresponding pinpointGcmChannel object, and an error if there is any.
func (c *FakePinpointGcmChannels) Get(name string, options v1.GetOptions) (result *v1alpha1.PinpointGcmChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pinpointgcmchannelsResource, c.ns, name), &v1alpha1.PinpointGcmChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointGcmChannel), err
}

// List takes label and field selectors, and returns the list of PinpointGcmChannels that match those selectors.
func (c *FakePinpointGcmChannels) List(opts v1.ListOptions) (result *v1alpha1.PinpointGcmChannelList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pinpointgcmchannelsResource, pinpointgcmchannelsKind, c.ns, opts), &v1alpha1.PinpointGcmChannelList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PinpointGcmChannelList{ListMeta: obj.(*v1alpha1.PinpointGcmChannelList).ListMeta}
	for _, item := range obj.(*v1alpha1.PinpointGcmChannelList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pinpointGcmChannels.
func (c *FakePinpointGcmChannels) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pinpointgcmchannelsResource, c.ns, opts))

}

// Create takes the representation of a pinpointGcmChannel and creates it.  Returns the server's representation of the pinpointGcmChannel, and an error, if there is any.
func (c *FakePinpointGcmChannels) Create(pinpointGcmChannel *v1alpha1.PinpointGcmChannel) (result *v1alpha1.PinpointGcmChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pinpointgcmchannelsResource, c.ns, pinpointGcmChannel), &v1alpha1.PinpointGcmChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointGcmChannel), err
}

// Update takes the representation of a pinpointGcmChannel and updates it. Returns the server's representation of the pinpointGcmChannel, and an error, if there is any.
func (c *FakePinpointGcmChannels) Update(pinpointGcmChannel *v1alpha1.PinpointGcmChannel) (result *v1alpha1.PinpointGcmChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pinpointgcmchannelsResource, c.ns, pinpointGcmChannel), &v1alpha1.PinpointGcmChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointGcmChannel), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePinpointGcmChannels) UpdateStatus(pinpointGcmChannel *v1alpha1.PinpointGcmChannel) (*v1alpha1.PinpointGcmChannel, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(pinpointgcmchannelsResource, "status", c.ns, pinpointGcmChannel), &v1alpha1.PinpointGcmChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointGcmChannel), err
}

// Delete takes name of the pinpointGcmChannel and deletes it. Returns an error if one occurs.
func (c *FakePinpointGcmChannels) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(pinpointgcmchannelsResource, c.ns, name), &v1alpha1.PinpointGcmChannel{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePinpointGcmChannels) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pinpointgcmchannelsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PinpointGcmChannelList{})
	return err
}

// Patch applies the patch and returns the patched pinpointGcmChannel.
func (c *FakePinpointGcmChannels) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PinpointGcmChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pinpointgcmchannelsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PinpointGcmChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointGcmChannel), err
}