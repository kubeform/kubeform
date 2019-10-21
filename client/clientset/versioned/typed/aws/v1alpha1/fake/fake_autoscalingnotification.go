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

// FakeAutoscalingNotifications implements AutoscalingNotificationInterface
type FakeAutoscalingNotifications struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var autoscalingnotificationsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "autoscalingnotifications"}

var autoscalingnotificationsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AutoscalingNotification"}

// Get takes name of the autoscalingNotification, and returns the corresponding autoscalingNotification object, and an error if there is any.
func (c *FakeAutoscalingNotifications) Get(name string, options v1.GetOptions) (result *v1alpha1.AutoscalingNotification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(autoscalingnotificationsResource, c.ns, name), &v1alpha1.AutoscalingNotification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoscalingNotification), err
}

// List takes label and field selectors, and returns the list of AutoscalingNotifications that match those selectors.
func (c *FakeAutoscalingNotifications) List(opts v1.ListOptions) (result *v1alpha1.AutoscalingNotificationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(autoscalingnotificationsResource, autoscalingnotificationsKind, c.ns, opts), &v1alpha1.AutoscalingNotificationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AutoscalingNotificationList{ListMeta: obj.(*v1alpha1.AutoscalingNotificationList).ListMeta}
	for _, item := range obj.(*v1alpha1.AutoscalingNotificationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested autoscalingNotifications.
func (c *FakeAutoscalingNotifications) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(autoscalingnotificationsResource, c.ns, opts))

}

// Create takes the representation of a autoscalingNotification and creates it.  Returns the server's representation of the autoscalingNotification, and an error, if there is any.
func (c *FakeAutoscalingNotifications) Create(autoscalingNotification *v1alpha1.AutoscalingNotification) (result *v1alpha1.AutoscalingNotification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(autoscalingnotificationsResource, c.ns, autoscalingNotification), &v1alpha1.AutoscalingNotification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoscalingNotification), err
}

// Update takes the representation of a autoscalingNotification and updates it. Returns the server's representation of the autoscalingNotification, and an error, if there is any.
func (c *FakeAutoscalingNotifications) Update(autoscalingNotification *v1alpha1.AutoscalingNotification) (result *v1alpha1.AutoscalingNotification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(autoscalingnotificationsResource, c.ns, autoscalingNotification), &v1alpha1.AutoscalingNotification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoscalingNotification), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAutoscalingNotifications) UpdateStatus(autoscalingNotification *v1alpha1.AutoscalingNotification) (*v1alpha1.AutoscalingNotification, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(autoscalingnotificationsResource, "status", c.ns, autoscalingNotification), &v1alpha1.AutoscalingNotification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoscalingNotification), err
}

// Delete takes name of the autoscalingNotification and deletes it. Returns an error if one occurs.
func (c *FakeAutoscalingNotifications) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(autoscalingnotificationsResource, c.ns, name), &v1alpha1.AutoscalingNotification{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAutoscalingNotifications) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(autoscalingnotificationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AutoscalingNotificationList{})
	return err
}

// Patch applies the patch and returns the patched autoscalingNotification.
func (c *FakeAutoscalingNotifications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AutoscalingNotification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(autoscalingnotificationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AutoscalingNotification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoscalingNotification), err
}