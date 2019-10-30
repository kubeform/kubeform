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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDataLakeAnalyticsAccounts implements DataLakeAnalyticsAccountInterface
type FakeDataLakeAnalyticsAccounts struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var datalakeanalyticsaccountsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "datalakeanalyticsaccounts"}

var datalakeanalyticsaccountsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "DataLakeAnalyticsAccount"}

// Get takes name of the dataLakeAnalyticsAccount, and returns the corresponding dataLakeAnalyticsAccount object, and an error if there is any.
func (c *FakeDataLakeAnalyticsAccounts) Get(name string, options v1.GetOptions) (result *v1alpha1.DataLakeAnalyticsAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(datalakeanalyticsaccountsResource, c.ns, name), &v1alpha1.DataLakeAnalyticsAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataLakeAnalyticsAccount), err
}

// List takes label and field selectors, and returns the list of DataLakeAnalyticsAccounts that match those selectors.
func (c *FakeDataLakeAnalyticsAccounts) List(opts v1.ListOptions) (result *v1alpha1.DataLakeAnalyticsAccountList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(datalakeanalyticsaccountsResource, datalakeanalyticsaccountsKind, c.ns, opts), &v1alpha1.DataLakeAnalyticsAccountList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DataLakeAnalyticsAccountList{ListMeta: obj.(*v1alpha1.DataLakeAnalyticsAccountList).ListMeta}
	for _, item := range obj.(*v1alpha1.DataLakeAnalyticsAccountList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dataLakeAnalyticsAccounts.
func (c *FakeDataLakeAnalyticsAccounts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(datalakeanalyticsaccountsResource, c.ns, opts))

}

// Create takes the representation of a dataLakeAnalyticsAccount and creates it.  Returns the server's representation of the dataLakeAnalyticsAccount, and an error, if there is any.
func (c *FakeDataLakeAnalyticsAccounts) Create(dataLakeAnalyticsAccount *v1alpha1.DataLakeAnalyticsAccount) (result *v1alpha1.DataLakeAnalyticsAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(datalakeanalyticsaccountsResource, c.ns, dataLakeAnalyticsAccount), &v1alpha1.DataLakeAnalyticsAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataLakeAnalyticsAccount), err
}

// Update takes the representation of a dataLakeAnalyticsAccount and updates it. Returns the server's representation of the dataLakeAnalyticsAccount, and an error, if there is any.
func (c *FakeDataLakeAnalyticsAccounts) Update(dataLakeAnalyticsAccount *v1alpha1.DataLakeAnalyticsAccount) (result *v1alpha1.DataLakeAnalyticsAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(datalakeanalyticsaccountsResource, c.ns, dataLakeAnalyticsAccount), &v1alpha1.DataLakeAnalyticsAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataLakeAnalyticsAccount), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDataLakeAnalyticsAccounts) UpdateStatus(dataLakeAnalyticsAccount *v1alpha1.DataLakeAnalyticsAccount) (*v1alpha1.DataLakeAnalyticsAccount, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(datalakeanalyticsaccountsResource, "status", c.ns, dataLakeAnalyticsAccount), &v1alpha1.DataLakeAnalyticsAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataLakeAnalyticsAccount), err
}

// Delete takes name of the dataLakeAnalyticsAccount and deletes it. Returns an error if one occurs.
func (c *FakeDataLakeAnalyticsAccounts) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(datalakeanalyticsaccountsResource, c.ns, name), &v1alpha1.DataLakeAnalyticsAccount{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDataLakeAnalyticsAccounts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(datalakeanalyticsaccountsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DataLakeAnalyticsAccountList{})
	return err
}

// Patch applies the patch and returns the patched dataLakeAnalyticsAccount.
func (c *FakeDataLakeAnalyticsAccounts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DataLakeAnalyticsAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(datalakeanalyticsaccountsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DataLakeAnalyticsAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataLakeAnalyticsAccount), err
}
