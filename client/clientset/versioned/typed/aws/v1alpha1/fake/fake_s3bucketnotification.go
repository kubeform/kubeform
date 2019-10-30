/*
Copyright The Kubeform Authors.

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

// FakeS3BucketNotifications implements S3BucketNotificationInterface
type FakeS3BucketNotifications struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var s3bucketnotificationsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "s3bucketnotifications"}

var s3bucketnotificationsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "S3BucketNotification"}

// Get takes name of the s3BucketNotification, and returns the corresponding s3BucketNotification object, and an error if there is any.
func (c *FakeS3BucketNotifications) Get(name string, options v1.GetOptions) (result *v1alpha1.S3BucketNotification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(s3bucketnotificationsResource, c.ns, name), &v1alpha1.S3BucketNotification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.S3BucketNotification), err
}

// List takes label and field selectors, and returns the list of S3BucketNotifications that match those selectors.
func (c *FakeS3BucketNotifications) List(opts v1.ListOptions) (result *v1alpha1.S3BucketNotificationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(s3bucketnotificationsResource, s3bucketnotificationsKind, c.ns, opts), &v1alpha1.S3BucketNotificationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.S3BucketNotificationList{ListMeta: obj.(*v1alpha1.S3BucketNotificationList).ListMeta}
	for _, item := range obj.(*v1alpha1.S3BucketNotificationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested s3BucketNotifications.
func (c *FakeS3BucketNotifications) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(s3bucketnotificationsResource, c.ns, opts))

}

// Create takes the representation of a s3BucketNotification and creates it.  Returns the server's representation of the s3BucketNotification, and an error, if there is any.
func (c *FakeS3BucketNotifications) Create(s3BucketNotification *v1alpha1.S3BucketNotification) (result *v1alpha1.S3BucketNotification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(s3bucketnotificationsResource, c.ns, s3BucketNotification), &v1alpha1.S3BucketNotification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.S3BucketNotification), err
}

// Update takes the representation of a s3BucketNotification and updates it. Returns the server's representation of the s3BucketNotification, and an error, if there is any.
func (c *FakeS3BucketNotifications) Update(s3BucketNotification *v1alpha1.S3BucketNotification) (result *v1alpha1.S3BucketNotification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(s3bucketnotificationsResource, c.ns, s3BucketNotification), &v1alpha1.S3BucketNotification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.S3BucketNotification), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeS3BucketNotifications) UpdateStatus(s3BucketNotification *v1alpha1.S3BucketNotification) (*v1alpha1.S3BucketNotification, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(s3bucketnotificationsResource, "status", c.ns, s3BucketNotification), &v1alpha1.S3BucketNotification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.S3BucketNotification), err
}

// Delete takes name of the s3BucketNotification and deletes it. Returns an error if one occurs.
func (c *FakeS3BucketNotifications) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(s3bucketnotificationsResource, c.ns, name), &v1alpha1.S3BucketNotification{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeS3BucketNotifications) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(s3bucketnotificationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.S3BucketNotificationList{})
	return err
}

// Patch applies the patch and returns the patched s3BucketNotification.
func (c *FakeS3BucketNotifications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.S3BucketNotification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(s3bucketnotificationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.S3BucketNotification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.S3BucketNotification), err
}
