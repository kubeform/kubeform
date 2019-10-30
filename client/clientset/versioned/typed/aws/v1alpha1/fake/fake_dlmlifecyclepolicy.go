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

// FakeDlmLifecyclePolicies implements DlmLifecyclePolicyInterface
type FakeDlmLifecyclePolicies struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var dlmlifecyclepoliciesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "dlmlifecyclepolicies"}

var dlmlifecyclepoliciesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DlmLifecyclePolicy"}

// Get takes name of the dlmLifecyclePolicy, and returns the corresponding dlmLifecyclePolicy object, and an error if there is any.
func (c *FakeDlmLifecyclePolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.DlmLifecyclePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dlmlifecyclepoliciesResource, c.ns, name), &v1alpha1.DlmLifecyclePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DlmLifecyclePolicy), err
}

// List takes label and field selectors, and returns the list of DlmLifecyclePolicies that match those selectors.
func (c *FakeDlmLifecyclePolicies) List(opts v1.ListOptions) (result *v1alpha1.DlmLifecyclePolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dlmlifecyclepoliciesResource, dlmlifecyclepoliciesKind, c.ns, opts), &v1alpha1.DlmLifecyclePolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DlmLifecyclePolicyList{ListMeta: obj.(*v1alpha1.DlmLifecyclePolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.DlmLifecyclePolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dlmLifecyclePolicies.
func (c *FakeDlmLifecyclePolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dlmlifecyclepoliciesResource, c.ns, opts))

}

// Create takes the representation of a dlmLifecyclePolicy and creates it.  Returns the server's representation of the dlmLifecyclePolicy, and an error, if there is any.
func (c *FakeDlmLifecyclePolicies) Create(dlmLifecyclePolicy *v1alpha1.DlmLifecyclePolicy) (result *v1alpha1.DlmLifecyclePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dlmlifecyclepoliciesResource, c.ns, dlmLifecyclePolicy), &v1alpha1.DlmLifecyclePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DlmLifecyclePolicy), err
}

// Update takes the representation of a dlmLifecyclePolicy and updates it. Returns the server's representation of the dlmLifecyclePolicy, and an error, if there is any.
func (c *FakeDlmLifecyclePolicies) Update(dlmLifecyclePolicy *v1alpha1.DlmLifecyclePolicy) (result *v1alpha1.DlmLifecyclePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dlmlifecyclepoliciesResource, c.ns, dlmLifecyclePolicy), &v1alpha1.DlmLifecyclePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DlmLifecyclePolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDlmLifecyclePolicies) UpdateStatus(dlmLifecyclePolicy *v1alpha1.DlmLifecyclePolicy) (*v1alpha1.DlmLifecyclePolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dlmlifecyclepoliciesResource, "status", c.ns, dlmLifecyclePolicy), &v1alpha1.DlmLifecyclePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DlmLifecyclePolicy), err
}

// Delete takes name of the dlmLifecyclePolicy and deletes it. Returns an error if one occurs.
func (c *FakeDlmLifecyclePolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dlmlifecyclepoliciesResource, c.ns, name), &v1alpha1.DlmLifecyclePolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDlmLifecyclePolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dlmlifecyclepoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DlmLifecyclePolicyList{})
	return err
}

// Patch applies the patch and returns the patched dlmLifecyclePolicy.
func (c *FakeDlmLifecyclePolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DlmLifecyclePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dlmlifecyclepoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.DlmLifecyclePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DlmLifecyclePolicy), err
}
