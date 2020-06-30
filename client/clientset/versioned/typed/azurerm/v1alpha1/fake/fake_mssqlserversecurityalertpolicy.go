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

// FakeMssqlServerSecurityAlertPolicies implements MssqlServerSecurityAlertPolicyInterface
type FakeMssqlServerSecurityAlertPolicies struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var mssqlserversecurityalertpoliciesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "mssqlserversecurityalertpolicies"}

var mssqlserversecurityalertpoliciesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "MssqlServerSecurityAlertPolicy"}

// Get takes name of the mssqlServerSecurityAlertPolicy, and returns the corresponding mssqlServerSecurityAlertPolicy object, and an error if there is any.
func (c *FakeMssqlServerSecurityAlertPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MssqlServerSecurityAlertPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mssqlserversecurityalertpoliciesResource, c.ns, name), &v1alpha1.MssqlServerSecurityAlertPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MssqlServerSecurityAlertPolicy), err
}

// List takes label and field selectors, and returns the list of MssqlServerSecurityAlertPolicies that match those selectors.
func (c *FakeMssqlServerSecurityAlertPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MssqlServerSecurityAlertPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mssqlserversecurityalertpoliciesResource, mssqlserversecurityalertpoliciesKind, c.ns, opts), &v1alpha1.MssqlServerSecurityAlertPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MssqlServerSecurityAlertPolicyList{ListMeta: obj.(*v1alpha1.MssqlServerSecurityAlertPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.MssqlServerSecurityAlertPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mssqlServerSecurityAlertPolicies.
func (c *FakeMssqlServerSecurityAlertPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mssqlserversecurityalertpoliciesResource, c.ns, opts))

}

// Create takes the representation of a mssqlServerSecurityAlertPolicy and creates it.  Returns the server's representation of the mssqlServerSecurityAlertPolicy, and an error, if there is any.
func (c *FakeMssqlServerSecurityAlertPolicies) Create(ctx context.Context, mssqlServerSecurityAlertPolicy *v1alpha1.MssqlServerSecurityAlertPolicy, opts v1.CreateOptions) (result *v1alpha1.MssqlServerSecurityAlertPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mssqlserversecurityalertpoliciesResource, c.ns, mssqlServerSecurityAlertPolicy), &v1alpha1.MssqlServerSecurityAlertPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MssqlServerSecurityAlertPolicy), err
}

// Update takes the representation of a mssqlServerSecurityAlertPolicy and updates it. Returns the server's representation of the mssqlServerSecurityAlertPolicy, and an error, if there is any.
func (c *FakeMssqlServerSecurityAlertPolicies) Update(ctx context.Context, mssqlServerSecurityAlertPolicy *v1alpha1.MssqlServerSecurityAlertPolicy, opts v1.UpdateOptions) (result *v1alpha1.MssqlServerSecurityAlertPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mssqlserversecurityalertpoliciesResource, c.ns, mssqlServerSecurityAlertPolicy), &v1alpha1.MssqlServerSecurityAlertPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MssqlServerSecurityAlertPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMssqlServerSecurityAlertPolicies) UpdateStatus(ctx context.Context, mssqlServerSecurityAlertPolicy *v1alpha1.MssqlServerSecurityAlertPolicy, opts v1.UpdateOptions) (*v1alpha1.MssqlServerSecurityAlertPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(mssqlserversecurityalertpoliciesResource, "status", c.ns, mssqlServerSecurityAlertPolicy), &v1alpha1.MssqlServerSecurityAlertPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MssqlServerSecurityAlertPolicy), err
}

// Delete takes name of the mssqlServerSecurityAlertPolicy and deletes it. Returns an error if one occurs.
func (c *FakeMssqlServerSecurityAlertPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(mssqlserversecurityalertpoliciesResource, c.ns, name), &v1alpha1.MssqlServerSecurityAlertPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMssqlServerSecurityAlertPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mssqlserversecurityalertpoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MssqlServerSecurityAlertPolicyList{})
	return err
}

// Patch applies the patch and returns the patched mssqlServerSecurityAlertPolicy.
func (c *FakeMssqlServerSecurityAlertPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MssqlServerSecurityAlertPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mssqlserversecurityalertpoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.MssqlServerSecurityAlertPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MssqlServerSecurityAlertPolicy), err
}
