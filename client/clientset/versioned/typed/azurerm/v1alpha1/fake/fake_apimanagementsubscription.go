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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeApiManagementSubscriptions implements ApiManagementSubscriptionInterface
type FakeApiManagementSubscriptions struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var apimanagementsubscriptionsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "apimanagementsubscriptions"}

var apimanagementsubscriptionsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "ApiManagementSubscription"}

// Get takes name of the apiManagementSubscription, and returns the corresponding apiManagementSubscription object, and an error if there is any.
func (c *FakeApiManagementSubscriptions) Get(name string, options v1.GetOptions) (result *v1alpha1.ApiManagementSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apimanagementsubscriptionsResource, c.ns, name), &v1alpha1.ApiManagementSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementSubscription), err
}

// List takes label and field selectors, and returns the list of ApiManagementSubscriptions that match those selectors.
func (c *FakeApiManagementSubscriptions) List(opts v1.ListOptions) (result *v1alpha1.ApiManagementSubscriptionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apimanagementsubscriptionsResource, apimanagementsubscriptionsKind, c.ns, opts), &v1alpha1.ApiManagementSubscriptionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApiManagementSubscriptionList{ListMeta: obj.(*v1alpha1.ApiManagementSubscriptionList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApiManagementSubscriptionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apiManagementSubscriptions.
func (c *FakeApiManagementSubscriptions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apimanagementsubscriptionsResource, c.ns, opts))

}

// Create takes the representation of a apiManagementSubscription and creates it.  Returns the server's representation of the apiManagementSubscription, and an error, if there is any.
func (c *FakeApiManagementSubscriptions) Create(apiManagementSubscription *v1alpha1.ApiManagementSubscription) (result *v1alpha1.ApiManagementSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apimanagementsubscriptionsResource, c.ns, apiManagementSubscription), &v1alpha1.ApiManagementSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementSubscription), err
}

// Update takes the representation of a apiManagementSubscription and updates it. Returns the server's representation of the apiManagementSubscription, and an error, if there is any.
func (c *FakeApiManagementSubscriptions) Update(apiManagementSubscription *v1alpha1.ApiManagementSubscription) (result *v1alpha1.ApiManagementSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apimanagementsubscriptionsResource, c.ns, apiManagementSubscription), &v1alpha1.ApiManagementSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementSubscription), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApiManagementSubscriptions) UpdateStatus(apiManagementSubscription *v1alpha1.ApiManagementSubscription) (*v1alpha1.ApiManagementSubscription, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apimanagementsubscriptionsResource, "status", c.ns, apiManagementSubscription), &v1alpha1.ApiManagementSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementSubscription), err
}

// Delete takes name of the apiManagementSubscription and deletes it. Returns an error if one occurs.
func (c *FakeApiManagementSubscriptions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(apimanagementsubscriptionsResource, c.ns, name), &v1alpha1.ApiManagementSubscription{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApiManagementSubscriptions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apimanagementsubscriptionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApiManagementSubscriptionList{})
	return err
}

// Patch applies the patch and returns the patched apiManagementSubscription.
func (c *FakeApiManagementSubscriptions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiManagementSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apimanagementsubscriptionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApiManagementSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementSubscription), err
}
