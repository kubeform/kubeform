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

// FakeDirectoryServiceLogSubscriptions implements DirectoryServiceLogSubscriptionInterface
type FakeDirectoryServiceLogSubscriptions struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var directoryservicelogsubscriptionsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "directoryservicelogsubscriptions"}

var directoryservicelogsubscriptionsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DirectoryServiceLogSubscription"}

// Get takes name of the directoryServiceLogSubscription, and returns the corresponding directoryServiceLogSubscription object, and an error if there is any.
func (c *FakeDirectoryServiceLogSubscriptions) Get(name string, options v1.GetOptions) (result *v1alpha1.DirectoryServiceLogSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(directoryservicelogsubscriptionsResource, c.ns, name), &v1alpha1.DirectoryServiceLogSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DirectoryServiceLogSubscription), err
}

// List takes label and field selectors, and returns the list of DirectoryServiceLogSubscriptions that match those selectors.
func (c *FakeDirectoryServiceLogSubscriptions) List(opts v1.ListOptions) (result *v1alpha1.DirectoryServiceLogSubscriptionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(directoryservicelogsubscriptionsResource, directoryservicelogsubscriptionsKind, c.ns, opts), &v1alpha1.DirectoryServiceLogSubscriptionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DirectoryServiceLogSubscriptionList{ListMeta: obj.(*v1alpha1.DirectoryServiceLogSubscriptionList).ListMeta}
	for _, item := range obj.(*v1alpha1.DirectoryServiceLogSubscriptionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested directoryServiceLogSubscriptions.
func (c *FakeDirectoryServiceLogSubscriptions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(directoryservicelogsubscriptionsResource, c.ns, opts))

}

// Create takes the representation of a directoryServiceLogSubscription and creates it.  Returns the server's representation of the directoryServiceLogSubscription, and an error, if there is any.
func (c *FakeDirectoryServiceLogSubscriptions) Create(directoryServiceLogSubscription *v1alpha1.DirectoryServiceLogSubscription) (result *v1alpha1.DirectoryServiceLogSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(directoryservicelogsubscriptionsResource, c.ns, directoryServiceLogSubscription), &v1alpha1.DirectoryServiceLogSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DirectoryServiceLogSubscription), err
}

// Update takes the representation of a directoryServiceLogSubscription and updates it. Returns the server's representation of the directoryServiceLogSubscription, and an error, if there is any.
func (c *FakeDirectoryServiceLogSubscriptions) Update(directoryServiceLogSubscription *v1alpha1.DirectoryServiceLogSubscription) (result *v1alpha1.DirectoryServiceLogSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(directoryservicelogsubscriptionsResource, c.ns, directoryServiceLogSubscription), &v1alpha1.DirectoryServiceLogSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DirectoryServiceLogSubscription), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDirectoryServiceLogSubscriptions) UpdateStatus(directoryServiceLogSubscription *v1alpha1.DirectoryServiceLogSubscription) (*v1alpha1.DirectoryServiceLogSubscription, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(directoryservicelogsubscriptionsResource, "status", c.ns, directoryServiceLogSubscription), &v1alpha1.DirectoryServiceLogSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DirectoryServiceLogSubscription), err
}

// Delete takes name of the directoryServiceLogSubscription and deletes it. Returns an error if one occurs.
func (c *FakeDirectoryServiceLogSubscriptions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(directoryservicelogsubscriptionsResource, c.ns, name), &v1alpha1.DirectoryServiceLogSubscription{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDirectoryServiceLogSubscriptions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(directoryservicelogsubscriptionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DirectoryServiceLogSubscriptionList{})
	return err
}

// Patch applies the patch and returns the patched directoryServiceLogSubscription.
func (c *FakeDirectoryServiceLogSubscriptions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DirectoryServiceLogSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(directoryservicelogsubscriptionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DirectoryServiceLogSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DirectoryServiceLogSubscription), err
}