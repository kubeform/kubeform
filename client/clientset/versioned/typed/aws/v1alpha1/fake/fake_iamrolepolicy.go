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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIamRolePolicies implements IamRolePolicyInterface
type FakeIamRolePolicies struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var iamrolepoliciesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "iamrolepolicies"}

var iamrolepoliciesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "IamRolePolicy"}

// Get takes name of the iamRolePolicy, and returns the corresponding iamRolePolicy object, and an error if there is any.
func (c *FakeIamRolePolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.IamRolePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iamrolepoliciesResource, c.ns, name), &v1alpha1.IamRolePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamRolePolicy), err
}

// List takes label and field selectors, and returns the list of IamRolePolicies that match those selectors.
func (c *FakeIamRolePolicies) List(opts v1.ListOptions) (result *v1alpha1.IamRolePolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iamrolepoliciesResource, iamrolepoliciesKind, c.ns, opts), &v1alpha1.IamRolePolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IamRolePolicyList{ListMeta: obj.(*v1alpha1.IamRolePolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.IamRolePolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iamRolePolicies.
func (c *FakeIamRolePolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iamrolepoliciesResource, c.ns, opts))

}

// Create takes the representation of a iamRolePolicy and creates it.  Returns the server's representation of the iamRolePolicy, and an error, if there is any.
func (c *FakeIamRolePolicies) Create(iamRolePolicy *v1alpha1.IamRolePolicy) (result *v1alpha1.IamRolePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iamrolepoliciesResource, c.ns, iamRolePolicy), &v1alpha1.IamRolePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamRolePolicy), err
}

// Update takes the representation of a iamRolePolicy and updates it. Returns the server's representation of the iamRolePolicy, and an error, if there is any.
func (c *FakeIamRolePolicies) Update(iamRolePolicy *v1alpha1.IamRolePolicy) (result *v1alpha1.IamRolePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iamrolepoliciesResource, c.ns, iamRolePolicy), &v1alpha1.IamRolePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamRolePolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIamRolePolicies) UpdateStatus(iamRolePolicy *v1alpha1.IamRolePolicy) (*v1alpha1.IamRolePolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(iamrolepoliciesResource, "status", c.ns, iamRolePolicy), &v1alpha1.IamRolePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamRolePolicy), err
}

// Delete takes name of the iamRolePolicy and deletes it. Returns an error if one occurs.
func (c *FakeIamRolePolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(iamrolepoliciesResource, c.ns, name), &v1alpha1.IamRolePolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIamRolePolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iamrolepoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.IamRolePolicyList{})
	return err
}

// Patch applies the patch and returns the patched iamRolePolicy.
func (c *FakeIamRolePolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IamRolePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iamrolepoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.IamRolePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamRolePolicy), err
}
