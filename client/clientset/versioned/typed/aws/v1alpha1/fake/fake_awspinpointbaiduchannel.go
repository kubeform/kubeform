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

// FakeAwsPinpointBaiduChannels implements AwsPinpointBaiduChannelInterface
type FakeAwsPinpointBaiduChannels struct {
	Fake *FakeAwsV1alpha1
}

var awspinpointbaiduchannelsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awspinpointbaiduchannels"}

var awspinpointbaiduchannelsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsPinpointBaiduChannel"}

// Get takes name of the awsPinpointBaiduChannel, and returns the corresponding awsPinpointBaiduChannel object, and an error if there is any.
func (c *FakeAwsPinpointBaiduChannels) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsPinpointBaiduChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awspinpointbaiduchannelsResource, name), &v1alpha1.AwsPinpointBaiduChannel{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsPinpointBaiduChannel), err
}

// List takes label and field selectors, and returns the list of AwsPinpointBaiduChannels that match those selectors.
func (c *FakeAwsPinpointBaiduChannels) List(opts v1.ListOptions) (result *v1alpha1.AwsPinpointBaiduChannelList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awspinpointbaiduchannelsResource, awspinpointbaiduchannelsKind, opts), &v1alpha1.AwsPinpointBaiduChannelList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsPinpointBaiduChannelList{ListMeta: obj.(*v1alpha1.AwsPinpointBaiduChannelList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsPinpointBaiduChannelList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsPinpointBaiduChannels.
func (c *FakeAwsPinpointBaiduChannels) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awspinpointbaiduchannelsResource, opts))
}

// Create takes the representation of a awsPinpointBaiduChannel and creates it.  Returns the server's representation of the awsPinpointBaiduChannel, and an error, if there is any.
func (c *FakeAwsPinpointBaiduChannels) Create(awsPinpointBaiduChannel *v1alpha1.AwsPinpointBaiduChannel) (result *v1alpha1.AwsPinpointBaiduChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awspinpointbaiduchannelsResource, awsPinpointBaiduChannel), &v1alpha1.AwsPinpointBaiduChannel{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsPinpointBaiduChannel), err
}

// Update takes the representation of a awsPinpointBaiduChannel and updates it. Returns the server's representation of the awsPinpointBaiduChannel, and an error, if there is any.
func (c *FakeAwsPinpointBaiduChannels) Update(awsPinpointBaiduChannel *v1alpha1.AwsPinpointBaiduChannel) (result *v1alpha1.AwsPinpointBaiduChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awspinpointbaiduchannelsResource, awsPinpointBaiduChannel), &v1alpha1.AwsPinpointBaiduChannel{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsPinpointBaiduChannel), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsPinpointBaiduChannels) UpdateStatus(awsPinpointBaiduChannel *v1alpha1.AwsPinpointBaiduChannel) (*v1alpha1.AwsPinpointBaiduChannel, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awspinpointbaiduchannelsResource, "status", awsPinpointBaiduChannel), &v1alpha1.AwsPinpointBaiduChannel{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsPinpointBaiduChannel), err
}

// Delete takes name of the awsPinpointBaiduChannel and deletes it. Returns an error if one occurs.
func (c *FakeAwsPinpointBaiduChannels) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awspinpointbaiduchannelsResource, name), &v1alpha1.AwsPinpointBaiduChannel{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsPinpointBaiduChannels) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awspinpointbaiduchannelsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsPinpointBaiduChannelList{})
	return err
}

// Patch applies the patch and returns the patched awsPinpointBaiduChannel.
func (c *FakeAwsPinpointBaiduChannels) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsPinpointBaiduChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awspinpointbaiduchannelsResource, name, pt, data, subresources...), &v1alpha1.AwsPinpointBaiduChannel{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsPinpointBaiduChannel), err
}
