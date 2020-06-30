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
	"context"

	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNotificationHubAuthorizationRules implements NotificationHubAuthorizationRuleInterface
type FakeNotificationHubAuthorizationRules struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var notificationhubauthorizationrulesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "notificationhubauthorizationrules"}

var notificationhubauthorizationrulesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "NotificationHubAuthorizationRule"}

// Get takes name of the notificationHubAuthorizationRule, and returns the corresponding notificationHubAuthorizationRule object, and an error if there is any.
func (c *FakeNotificationHubAuthorizationRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NotificationHubAuthorizationRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(notificationhubauthorizationrulesResource, c.ns, name), &v1alpha1.NotificationHubAuthorizationRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NotificationHubAuthorizationRule), err
}

// List takes label and field selectors, and returns the list of NotificationHubAuthorizationRules that match those selectors.
func (c *FakeNotificationHubAuthorizationRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NotificationHubAuthorizationRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(notificationhubauthorizationrulesResource, notificationhubauthorizationrulesKind, c.ns, opts), &v1alpha1.NotificationHubAuthorizationRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NotificationHubAuthorizationRuleList{ListMeta: obj.(*v1alpha1.NotificationHubAuthorizationRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.NotificationHubAuthorizationRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested notificationHubAuthorizationRules.
func (c *FakeNotificationHubAuthorizationRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(notificationhubauthorizationrulesResource, c.ns, opts))

}

// Create takes the representation of a notificationHubAuthorizationRule and creates it.  Returns the server's representation of the notificationHubAuthorizationRule, and an error, if there is any.
func (c *FakeNotificationHubAuthorizationRules) Create(ctx context.Context, notificationHubAuthorizationRule *v1alpha1.NotificationHubAuthorizationRule, opts v1.CreateOptions) (result *v1alpha1.NotificationHubAuthorizationRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(notificationhubauthorizationrulesResource, c.ns, notificationHubAuthorizationRule), &v1alpha1.NotificationHubAuthorizationRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NotificationHubAuthorizationRule), err
}

// Update takes the representation of a notificationHubAuthorizationRule and updates it. Returns the server's representation of the notificationHubAuthorizationRule, and an error, if there is any.
func (c *FakeNotificationHubAuthorizationRules) Update(ctx context.Context, notificationHubAuthorizationRule *v1alpha1.NotificationHubAuthorizationRule, opts v1.UpdateOptions) (result *v1alpha1.NotificationHubAuthorizationRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(notificationhubauthorizationrulesResource, c.ns, notificationHubAuthorizationRule), &v1alpha1.NotificationHubAuthorizationRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NotificationHubAuthorizationRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNotificationHubAuthorizationRules) UpdateStatus(ctx context.Context, notificationHubAuthorizationRule *v1alpha1.NotificationHubAuthorizationRule, opts v1.UpdateOptions) (*v1alpha1.NotificationHubAuthorizationRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(notificationhubauthorizationrulesResource, "status", c.ns, notificationHubAuthorizationRule), &v1alpha1.NotificationHubAuthorizationRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NotificationHubAuthorizationRule), err
}

// Delete takes name of the notificationHubAuthorizationRule and deletes it. Returns an error if one occurs.
func (c *FakeNotificationHubAuthorizationRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(notificationhubauthorizationrulesResource, c.ns, name), &v1alpha1.NotificationHubAuthorizationRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNotificationHubAuthorizationRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(notificationhubauthorizationrulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NotificationHubAuthorizationRuleList{})
	return err
}

// Patch applies the patch and returns the patched notificationHubAuthorizationRule.
func (c *FakeNotificationHubAuthorizationRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NotificationHubAuthorizationRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(notificationhubauthorizationrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.NotificationHubAuthorizationRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NotificationHubAuthorizationRule), err
}
