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

// FakeAwsPinpointApnsVoipChannels implements AwsPinpointApnsVoipChannelInterface
type FakeAwsPinpointApnsVoipChannels struct {
	Fake *FakeAwsV1alpha1
}

var awspinpointapnsvoipchannelsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awspinpointapnsvoipchannels"}

var awspinpointapnsvoipchannelsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsPinpointApnsVoipChannel"}

// Get takes name of the awsPinpointApnsVoipChannel, and returns the corresponding awsPinpointApnsVoipChannel object, and an error if there is any.
func (c *FakeAwsPinpointApnsVoipChannels) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsPinpointApnsVoipChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awspinpointapnsvoipchannelsResource, name), &v1alpha1.AwsPinpointApnsVoipChannel{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsPinpointApnsVoipChannel), err
}

// List takes label and field selectors, and returns the list of AwsPinpointApnsVoipChannels that match those selectors.
func (c *FakeAwsPinpointApnsVoipChannels) List(opts v1.ListOptions) (result *v1alpha1.AwsPinpointApnsVoipChannelList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awspinpointapnsvoipchannelsResource, awspinpointapnsvoipchannelsKind, opts), &v1alpha1.AwsPinpointApnsVoipChannelList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsPinpointApnsVoipChannelList{ListMeta: obj.(*v1alpha1.AwsPinpointApnsVoipChannelList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsPinpointApnsVoipChannelList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsPinpointApnsVoipChannels.
func (c *FakeAwsPinpointApnsVoipChannels) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awspinpointapnsvoipchannelsResource, opts))
}

// Create takes the representation of a awsPinpointApnsVoipChannel and creates it.  Returns the server's representation of the awsPinpointApnsVoipChannel, and an error, if there is any.
func (c *FakeAwsPinpointApnsVoipChannels) Create(awsPinpointApnsVoipChannel *v1alpha1.AwsPinpointApnsVoipChannel) (result *v1alpha1.AwsPinpointApnsVoipChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awspinpointapnsvoipchannelsResource, awsPinpointApnsVoipChannel), &v1alpha1.AwsPinpointApnsVoipChannel{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsPinpointApnsVoipChannel), err
}

// Update takes the representation of a awsPinpointApnsVoipChannel and updates it. Returns the server's representation of the awsPinpointApnsVoipChannel, and an error, if there is any.
func (c *FakeAwsPinpointApnsVoipChannels) Update(awsPinpointApnsVoipChannel *v1alpha1.AwsPinpointApnsVoipChannel) (result *v1alpha1.AwsPinpointApnsVoipChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awspinpointapnsvoipchannelsResource, awsPinpointApnsVoipChannel), &v1alpha1.AwsPinpointApnsVoipChannel{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsPinpointApnsVoipChannel), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsPinpointApnsVoipChannels) UpdateStatus(awsPinpointApnsVoipChannel *v1alpha1.AwsPinpointApnsVoipChannel) (*v1alpha1.AwsPinpointApnsVoipChannel, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awspinpointapnsvoipchannelsResource, "status", awsPinpointApnsVoipChannel), &v1alpha1.AwsPinpointApnsVoipChannel{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsPinpointApnsVoipChannel), err
}

// Delete takes name of the awsPinpointApnsVoipChannel and deletes it. Returns an error if one occurs.
func (c *FakeAwsPinpointApnsVoipChannels) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awspinpointapnsvoipchannelsResource, name), &v1alpha1.AwsPinpointApnsVoipChannel{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsPinpointApnsVoipChannels) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awspinpointapnsvoipchannelsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsPinpointApnsVoipChannelList{})
	return err
}

// Patch applies the patch and returns the patched awsPinpointApnsVoipChannel.
func (c *FakeAwsPinpointApnsVoipChannels) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsPinpointApnsVoipChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awspinpointapnsvoipchannelsResource, name, pt, data, subresources...), &v1alpha1.AwsPinpointApnsVoipChannel{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsPinpointApnsVoipChannel), err
}
