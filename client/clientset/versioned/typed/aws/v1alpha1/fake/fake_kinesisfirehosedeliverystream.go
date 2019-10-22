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

// FakeKinesisFirehoseDeliveryStreams implements KinesisFirehoseDeliveryStreamInterface
type FakeKinesisFirehoseDeliveryStreams struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var kinesisfirehosedeliverystreamsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "kinesisfirehosedeliverystreams"}

var kinesisfirehosedeliverystreamsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "KinesisFirehoseDeliveryStream"}

// Get takes name of the kinesisFirehoseDeliveryStream, and returns the corresponding kinesisFirehoseDeliveryStream object, and an error if there is any.
func (c *FakeKinesisFirehoseDeliveryStreams) Get(name string, options v1.GetOptions) (result *v1alpha1.KinesisFirehoseDeliveryStream, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kinesisfirehosedeliverystreamsResource, c.ns, name), &v1alpha1.KinesisFirehoseDeliveryStream{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KinesisFirehoseDeliveryStream), err
}

// List takes label and field selectors, and returns the list of KinesisFirehoseDeliveryStreams that match those selectors.
func (c *FakeKinesisFirehoseDeliveryStreams) List(opts v1.ListOptions) (result *v1alpha1.KinesisFirehoseDeliveryStreamList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kinesisfirehosedeliverystreamsResource, kinesisfirehosedeliverystreamsKind, c.ns, opts), &v1alpha1.KinesisFirehoseDeliveryStreamList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KinesisFirehoseDeliveryStreamList{ListMeta: obj.(*v1alpha1.KinesisFirehoseDeliveryStreamList).ListMeta}
	for _, item := range obj.(*v1alpha1.KinesisFirehoseDeliveryStreamList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kinesisFirehoseDeliveryStreams.
func (c *FakeKinesisFirehoseDeliveryStreams) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kinesisfirehosedeliverystreamsResource, c.ns, opts))

}

// Create takes the representation of a kinesisFirehoseDeliveryStream and creates it.  Returns the server's representation of the kinesisFirehoseDeliveryStream, and an error, if there is any.
func (c *FakeKinesisFirehoseDeliveryStreams) Create(kinesisFirehoseDeliveryStream *v1alpha1.KinesisFirehoseDeliveryStream) (result *v1alpha1.KinesisFirehoseDeliveryStream, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kinesisfirehosedeliverystreamsResource, c.ns, kinesisFirehoseDeliveryStream), &v1alpha1.KinesisFirehoseDeliveryStream{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KinesisFirehoseDeliveryStream), err
}

// Update takes the representation of a kinesisFirehoseDeliveryStream and updates it. Returns the server's representation of the kinesisFirehoseDeliveryStream, and an error, if there is any.
func (c *FakeKinesisFirehoseDeliveryStreams) Update(kinesisFirehoseDeliveryStream *v1alpha1.KinesisFirehoseDeliveryStream) (result *v1alpha1.KinesisFirehoseDeliveryStream, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kinesisfirehosedeliverystreamsResource, c.ns, kinesisFirehoseDeliveryStream), &v1alpha1.KinesisFirehoseDeliveryStream{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KinesisFirehoseDeliveryStream), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKinesisFirehoseDeliveryStreams) UpdateStatus(kinesisFirehoseDeliveryStream *v1alpha1.KinesisFirehoseDeliveryStream) (*v1alpha1.KinesisFirehoseDeliveryStream, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kinesisfirehosedeliverystreamsResource, "status", c.ns, kinesisFirehoseDeliveryStream), &v1alpha1.KinesisFirehoseDeliveryStream{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KinesisFirehoseDeliveryStream), err
}

// Delete takes name of the kinesisFirehoseDeliveryStream and deletes it. Returns an error if one occurs.
func (c *FakeKinesisFirehoseDeliveryStreams) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kinesisfirehosedeliverystreamsResource, c.ns, name), &v1alpha1.KinesisFirehoseDeliveryStream{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKinesisFirehoseDeliveryStreams) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kinesisfirehosedeliverystreamsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.KinesisFirehoseDeliveryStreamList{})
	return err
}

// Patch applies the patch and returns the patched kinesisFirehoseDeliveryStream.
func (c *FakeKinesisFirehoseDeliveryStreams) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KinesisFirehoseDeliveryStream, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kinesisfirehosedeliverystreamsResource, c.ns, name, pt, data, subresources...), &v1alpha1.KinesisFirehoseDeliveryStream{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KinesisFirehoseDeliveryStream), err
}
