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

// FakeSecurityhubStandardsSubscriptions implements SecurityhubStandardsSubscriptionInterface
type FakeSecurityhubStandardsSubscriptions struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var securityhubstandardssubscriptionsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "securityhubstandardssubscriptions"}

var securityhubstandardssubscriptionsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "SecurityhubStandardsSubscription"}

// Get takes name of the securityhubStandardsSubscription, and returns the corresponding securityhubStandardsSubscription object, and an error if there is any.
func (c *FakeSecurityhubStandardsSubscriptions) Get(name string, options v1.GetOptions) (result *v1alpha1.SecurityhubStandardsSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(securityhubstandardssubscriptionsResource, c.ns, name), &v1alpha1.SecurityhubStandardsSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecurityhubStandardsSubscription), err
}

// List takes label and field selectors, and returns the list of SecurityhubStandardsSubscriptions that match those selectors.
func (c *FakeSecurityhubStandardsSubscriptions) List(opts v1.ListOptions) (result *v1alpha1.SecurityhubStandardsSubscriptionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(securityhubstandardssubscriptionsResource, securityhubstandardssubscriptionsKind, c.ns, opts), &v1alpha1.SecurityhubStandardsSubscriptionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SecurityhubStandardsSubscriptionList{ListMeta: obj.(*v1alpha1.SecurityhubStandardsSubscriptionList).ListMeta}
	for _, item := range obj.(*v1alpha1.SecurityhubStandardsSubscriptionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested securityhubStandardsSubscriptions.
func (c *FakeSecurityhubStandardsSubscriptions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(securityhubstandardssubscriptionsResource, c.ns, opts))

}

// Create takes the representation of a securityhubStandardsSubscription and creates it.  Returns the server's representation of the securityhubStandardsSubscription, and an error, if there is any.
func (c *FakeSecurityhubStandardsSubscriptions) Create(securityhubStandardsSubscription *v1alpha1.SecurityhubStandardsSubscription) (result *v1alpha1.SecurityhubStandardsSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(securityhubstandardssubscriptionsResource, c.ns, securityhubStandardsSubscription), &v1alpha1.SecurityhubStandardsSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecurityhubStandardsSubscription), err
}

// Update takes the representation of a securityhubStandardsSubscription and updates it. Returns the server's representation of the securityhubStandardsSubscription, and an error, if there is any.
func (c *FakeSecurityhubStandardsSubscriptions) Update(securityhubStandardsSubscription *v1alpha1.SecurityhubStandardsSubscription) (result *v1alpha1.SecurityhubStandardsSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(securityhubstandardssubscriptionsResource, c.ns, securityhubStandardsSubscription), &v1alpha1.SecurityhubStandardsSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecurityhubStandardsSubscription), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSecurityhubStandardsSubscriptions) UpdateStatus(securityhubStandardsSubscription *v1alpha1.SecurityhubStandardsSubscription) (*v1alpha1.SecurityhubStandardsSubscription, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(securityhubstandardssubscriptionsResource, "status", c.ns, securityhubStandardsSubscription), &v1alpha1.SecurityhubStandardsSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecurityhubStandardsSubscription), err
}

// Delete takes name of the securityhubStandardsSubscription and deletes it. Returns an error if one occurs.
func (c *FakeSecurityhubStandardsSubscriptions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(securityhubstandardssubscriptionsResource, c.ns, name), &v1alpha1.SecurityhubStandardsSubscription{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSecurityhubStandardsSubscriptions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(securityhubstandardssubscriptionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SecurityhubStandardsSubscriptionList{})
	return err
}

// Patch applies the patch and returns the patched securityhubStandardsSubscription.
func (c *FakeSecurityhubStandardsSubscriptions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SecurityhubStandardsSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(securityhubstandardssubscriptionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SecurityhubStandardsSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecurityhubStandardsSubscription), err
}
