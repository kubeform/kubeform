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

// FakeAwsVpcEndpointConnectionNotifications implements AwsVpcEndpointConnectionNotificationInterface
type FakeAwsVpcEndpointConnectionNotifications struct {
	Fake *FakeAwsV1alpha1
}

var awsvpcendpointconnectionnotificationsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsvpcendpointconnectionnotifications"}

var awsvpcendpointconnectionnotificationsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsVpcEndpointConnectionNotification"}

// Get takes name of the awsVpcEndpointConnectionNotification, and returns the corresponding awsVpcEndpointConnectionNotification object, and an error if there is any.
func (c *FakeAwsVpcEndpointConnectionNotifications) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsVpcEndpointConnectionNotification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsvpcendpointconnectionnotificationsResource, name), &v1alpha1.AwsVpcEndpointConnectionNotification{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVpcEndpointConnectionNotification), err
}

// List takes label and field selectors, and returns the list of AwsVpcEndpointConnectionNotifications that match those selectors.
func (c *FakeAwsVpcEndpointConnectionNotifications) List(opts v1.ListOptions) (result *v1alpha1.AwsVpcEndpointConnectionNotificationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsvpcendpointconnectionnotificationsResource, awsvpcendpointconnectionnotificationsKind, opts), &v1alpha1.AwsVpcEndpointConnectionNotificationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsVpcEndpointConnectionNotificationList{ListMeta: obj.(*v1alpha1.AwsVpcEndpointConnectionNotificationList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsVpcEndpointConnectionNotificationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsVpcEndpointConnectionNotifications.
func (c *FakeAwsVpcEndpointConnectionNotifications) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsvpcendpointconnectionnotificationsResource, opts))
}

// Create takes the representation of a awsVpcEndpointConnectionNotification and creates it.  Returns the server's representation of the awsVpcEndpointConnectionNotification, and an error, if there is any.
func (c *FakeAwsVpcEndpointConnectionNotifications) Create(awsVpcEndpointConnectionNotification *v1alpha1.AwsVpcEndpointConnectionNotification) (result *v1alpha1.AwsVpcEndpointConnectionNotification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsvpcendpointconnectionnotificationsResource, awsVpcEndpointConnectionNotification), &v1alpha1.AwsVpcEndpointConnectionNotification{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVpcEndpointConnectionNotification), err
}

// Update takes the representation of a awsVpcEndpointConnectionNotification and updates it. Returns the server's representation of the awsVpcEndpointConnectionNotification, and an error, if there is any.
func (c *FakeAwsVpcEndpointConnectionNotifications) Update(awsVpcEndpointConnectionNotification *v1alpha1.AwsVpcEndpointConnectionNotification) (result *v1alpha1.AwsVpcEndpointConnectionNotification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsvpcendpointconnectionnotificationsResource, awsVpcEndpointConnectionNotification), &v1alpha1.AwsVpcEndpointConnectionNotification{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVpcEndpointConnectionNotification), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsVpcEndpointConnectionNotifications) UpdateStatus(awsVpcEndpointConnectionNotification *v1alpha1.AwsVpcEndpointConnectionNotification) (*v1alpha1.AwsVpcEndpointConnectionNotification, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsvpcendpointconnectionnotificationsResource, "status", awsVpcEndpointConnectionNotification), &v1alpha1.AwsVpcEndpointConnectionNotification{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVpcEndpointConnectionNotification), err
}

// Delete takes name of the awsVpcEndpointConnectionNotification and deletes it. Returns an error if one occurs.
func (c *FakeAwsVpcEndpointConnectionNotifications) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsvpcendpointconnectionnotificationsResource, name), &v1alpha1.AwsVpcEndpointConnectionNotification{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsVpcEndpointConnectionNotifications) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsvpcendpointconnectionnotificationsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsVpcEndpointConnectionNotificationList{})
	return err
}

// Patch applies the patch and returns the patched awsVpcEndpointConnectionNotification.
func (c *FakeAwsVpcEndpointConnectionNotifications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsVpcEndpointConnectionNotification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsvpcendpointconnectionnotificationsResource, name, pt, data, subresources...), &v1alpha1.AwsVpcEndpointConnectionNotification{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVpcEndpointConnectionNotification), err
}
