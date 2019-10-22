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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePinpointApnsVoipChannels implements PinpointApnsVoipChannelInterface
type FakePinpointApnsVoipChannels struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var pinpointapnsvoipchannelsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "pinpointapnsvoipchannels"}

var pinpointapnsvoipchannelsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "PinpointApnsVoipChannel"}

// Get takes name of the pinpointApnsVoipChannel, and returns the corresponding pinpointApnsVoipChannel object, and an error if there is any.
func (c *FakePinpointApnsVoipChannels) Get(name string, options v1.GetOptions) (result *v1alpha1.PinpointApnsVoipChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pinpointapnsvoipchannelsResource, c.ns, name), &v1alpha1.PinpointApnsVoipChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointApnsVoipChannel), err
}

// List takes label and field selectors, and returns the list of PinpointApnsVoipChannels that match those selectors.
func (c *FakePinpointApnsVoipChannels) List(opts v1.ListOptions) (result *v1alpha1.PinpointApnsVoipChannelList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pinpointapnsvoipchannelsResource, pinpointapnsvoipchannelsKind, c.ns, opts), &v1alpha1.PinpointApnsVoipChannelList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PinpointApnsVoipChannelList{ListMeta: obj.(*v1alpha1.PinpointApnsVoipChannelList).ListMeta}
	for _, item := range obj.(*v1alpha1.PinpointApnsVoipChannelList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pinpointApnsVoipChannels.
func (c *FakePinpointApnsVoipChannels) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pinpointapnsvoipchannelsResource, c.ns, opts))

}

// Create takes the representation of a pinpointApnsVoipChannel and creates it.  Returns the server's representation of the pinpointApnsVoipChannel, and an error, if there is any.
func (c *FakePinpointApnsVoipChannels) Create(pinpointApnsVoipChannel *v1alpha1.PinpointApnsVoipChannel) (result *v1alpha1.PinpointApnsVoipChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pinpointapnsvoipchannelsResource, c.ns, pinpointApnsVoipChannel), &v1alpha1.PinpointApnsVoipChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointApnsVoipChannel), err
}

// Update takes the representation of a pinpointApnsVoipChannel and updates it. Returns the server's representation of the pinpointApnsVoipChannel, and an error, if there is any.
func (c *FakePinpointApnsVoipChannels) Update(pinpointApnsVoipChannel *v1alpha1.PinpointApnsVoipChannel) (result *v1alpha1.PinpointApnsVoipChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pinpointapnsvoipchannelsResource, c.ns, pinpointApnsVoipChannel), &v1alpha1.PinpointApnsVoipChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointApnsVoipChannel), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePinpointApnsVoipChannels) UpdateStatus(pinpointApnsVoipChannel *v1alpha1.PinpointApnsVoipChannel) (*v1alpha1.PinpointApnsVoipChannel, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(pinpointapnsvoipchannelsResource, "status", c.ns, pinpointApnsVoipChannel), &v1alpha1.PinpointApnsVoipChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointApnsVoipChannel), err
}

// Delete takes name of the pinpointApnsVoipChannel and deletes it. Returns an error if one occurs.
func (c *FakePinpointApnsVoipChannels) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(pinpointapnsvoipchannelsResource, c.ns, name), &v1alpha1.PinpointApnsVoipChannel{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePinpointApnsVoipChannels) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pinpointapnsvoipchannelsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PinpointApnsVoipChannelList{})
	return err
}

// Patch applies the patch and returns the patched pinpointApnsVoipChannel.
func (c *FakePinpointApnsVoipChannels) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PinpointApnsVoipChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pinpointapnsvoipchannelsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PinpointApnsVoipChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointApnsVoipChannel), err
}
