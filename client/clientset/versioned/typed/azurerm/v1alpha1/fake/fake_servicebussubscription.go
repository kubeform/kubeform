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

// FakeServicebusSubscriptions implements ServicebusSubscriptionInterface
type FakeServicebusSubscriptions struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var servicebussubscriptionsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "servicebussubscriptions"}

var servicebussubscriptionsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "ServicebusSubscription"}

// Get takes name of the servicebusSubscription, and returns the corresponding servicebusSubscription object, and an error if there is any.
func (c *FakeServicebusSubscriptions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ServicebusSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicebussubscriptionsResource, c.ns, name), &v1alpha1.ServicebusSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicebusSubscription), err
}

// List takes label and field selectors, and returns the list of ServicebusSubscriptions that match those selectors.
func (c *FakeServicebusSubscriptions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServicebusSubscriptionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicebussubscriptionsResource, servicebussubscriptionsKind, c.ns, opts), &v1alpha1.ServicebusSubscriptionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ServicebusSubscriptionList{ListMeta: obj.(*v1alpha1.ServicebusSubscriptionList).ListMeta}
	for _, item := range obj.(*v1alpha1.ServicebusSubscriptionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested servicebusSubscriptions.
func (c *FakeServicebusSubscriptions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicebussubscriptionsResource, c.ns, opts))

}

// Create takes the representation of a servicebusSubscription and creates it.  Returns the server's representation of the servicebusSubscription, and an error, if there is any.
func (c *FakeServicebusSubscriptions) Create(ctx context.Context, servicebusSubscription *v1alpha1.ServicebusSubscription, opts v1.CreateOptions) (result *v1alpha1.ServicebusSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicebussubscriptionsResource, c.ns, servicebusSubscription), &v1alpha1.ServicebusSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicebusSubscription), err
}

// Update takes the representation of a servicebusSubscription and updates it. Returns the server's representation of the servicebusSubscription, and an error, if there is any.
func (c *FakeServicebusSubscriptions) Update(ctx context.Context, servicebusSubscription *v1alpha1.ServicebusSubscription, opts v1.UpdateOptions) (result *v1alpha1.ServicebusSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicebussubscriptionsResource, c.ns, servicebusSubscription), &v1alpha1.ServicebusSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicebusSubscription), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServicebusSubscriptions) UpdateStatus(ctx context.Context, servicebusSubscription *v1alpha1.ServicebusSubscription, opts v1.UpdateOptions) (*v1alpha1.ServicebusSubscription, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(servicebussubscriptionsResource, "status", c.ns, servicebusSubscription), &v1alpha1.ServicebusSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicebusSubscription), err
}

// Delete takes name of the servicebusSubscription and deletes it. Returns an error if one occurs.
func (c *FakeServicebusSubscriptions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(servicebussubscriptionsResource, c.ns, name), &v1alpha1.ServicebusSubscription{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServicebusSubscriptions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicebussubscriptionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ServicebusSubscriptionList{})
	return err
}

// Patch applies the patch and returns the patched servicebusSubscription.
func (c *FakeServicebusSubscriptions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServicebusSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicebussubscriptionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ServicebusSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicebusSubscription), err
}
