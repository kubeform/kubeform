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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeGoogleMonitoringNotificationChannels implements GoogleMonitoringNotificationChannelInterface
type FakeGoogleMonitoringNotificationChannels struct {
	Fake *FakeGoogleV1alpha1
}

var googlemonitoringnotificationchannelsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googlemonitoringnotificationchannels"}

var googlemonitoringnotificationchannelsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GoogleMonitoringNotificationChannel"}

// Get takes name of the googleMonitoringNotificationChannel, and returns the corresponding googleMonitoringNotificationChannel object, and an error if there is any.
func (c *FakeGoogleMonitoringNotificationChannels) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleMonitoringNotificationChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googlemonitoringnotificationchannelsResource, name), &v1alpha1.GoogleMonitoringNotificationChannel{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleMonitoringNotificationChannel), err
}

// List takes label and field selectors, and returns the list of GoogleMonitoringNotificationChannels that match those selectors.
func (c *FakeGoogleMonitoringNotificationChannels) List(opts v1.ListOptions) (result *v1alpha1.GoogleMonitoringNotificationChannelList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googlemonitoringnotificationchannelsResource, googlemonitoringnotificationchannelsKind, opts), &v1alpha1.GoogleMonitoringNotificationChannelList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleMonitoringNotificationChannelList{ListMeta: obj.(*v1alpha1.GoogleMonitoringNotificationChannelList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleMonitoringNotificationChannelList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleMonitoringNotificationChannels.
func (c *FakeGoogleMonitoringNotificationChannels) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googlemonitoringnotificationchannelsResource, opts))
}

// Create takes the representation of a googleMonitoringNotificationChannel and creates it.  Returns the server's representation of the googleMonitoringNotificationChannel, and an error, if there is any.
func (c *FakeGoogleMonitoringNotificationChannels) Create(googleMonitoringNotificationChannel *v1alpha1.GoogleMonitoringNotificationChannel) (result *v1alpha1.GoogleMonitoringNotificationChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googlemonitoringnotificationchannelsResource, googleMonitoringNotificationChannel), &v1alpha1.GoogleMonitoringNotificationChannel{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleMonitoringNotificationChannel), err
}

// Update takes the representation of a googleMonitoringNotificationChannel and updates it. Returns the server's representation of the googleMonitoringNotificationChannel, and an error, if there is any.
func (c *FakeGoogleMonitoringNotificationChannels) Update(googleMonitoringNotificationChannel *v1alpha1.GoogleMonitoringNotificationChannel) (result *v1alpha1.GoogleMonitoringNotificationChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googlemonitoringnotificationchannelsResource, googleMonitoringNotificationChannel), &v1alpha1.GoogleMonitoringNotificationChannel{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleMonitoringNotificationChannel), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleMonitoringNotificationChannels) UpdateStatus(googleMonitoringNotificationChannel *v1alpha1.GoogleMonitoringNotificationChannel) (*v1alpha1.GoogleMonitoringNotificationChannel, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googlemonitoringnotificationchannelsResource, "status", googleMonitoringNotificationChannel), &v1alpha1.GoogleMonitoringNotificationChannel{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleMonitoringNotificationChannel), err
}

// Delete takes name of the googleMonitoringNotificationChannel and deletes it. Returns an error if one occurs.
func (c *FakeGoogleMonitoringNotificationChannels) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googlemonitoringnotificationchannelsResource, name), &v1alpha1.GoogleMonitoringNotificationChannel{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleMonitoringNotificationChannels) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googlemonitoringnotificationchannelsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleMonitoringNotificationChannelList{})
	return err
}

// Patch applies the patch and returns the patched googleMonitoringNotificationChannel.
func (c *FakeGoogleMonitoringNotificationChannels) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleMonitoringNotificationChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googlemonitoringnotificationchannelsResource, name, pt, data, subresources...), &v1alpha1.GoogleMonitoringNotificationChannel{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleMonitoringNotificationChannel), err
}
