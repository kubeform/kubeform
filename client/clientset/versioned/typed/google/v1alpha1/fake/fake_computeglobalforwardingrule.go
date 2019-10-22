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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeComputeGlobalForwardingRules implements ComputeGlobalForwardingRuleInterface
type FakeComputeGlobalForwardingRules struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var computeglobalforwardingrulesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "computeglobalforwardingrules"}

var computeglobalforwardingrulesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "ComputeGlobalForwardingRule"}

// Get takes name of the computeGlobalForwardingRule, and returns the corresponding computeGlobalForwardingRule object, and an error if there is any.
func (c *FakeComputeGlobalForwardingRules) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeGlobalForwardingRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computeglobalforwardingrulesResource, c.ns, name), &v1alpha1.ComputeGlobalForwardingRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeGlobalForwardingRule), err
}

// List takes label and field selectors, and returns the list of ComputeGlobalForwardingRules that match those selectors.
func (c *FakeComputeGlobalForwardingRules) List(opts v1.ListOptions) (result *v1alpha1.ComputeGlobalForwardingRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computeglobalforwardingrulesResource, computeglobalforwardingrulesKind, c.ns, opts), &v1alpha1.ComputeGlobalForwardingRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComputeGlobalForwardingRuleList{ListMeta: obj.(*v1alpha1.ComputeGlobalForwardingRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComputeGlobalForwardingRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeGlobalForwardingRules.
func (c *FakeComputeGlobalForwardingRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computeglobalforwardingrulesResource, c.ns, opts))

}

// Create takes the representation of a computeGlobalForwardingRule and creates it.  Returns the server's representation of the computeGlobalForwardingRule, and an error, if there is any.
func (c *FakeComputeGlobalForwardingRules) Create(computeGlobalForwardingRule *v1alpha1.ComputeGlobalForwardingRule) (result *v1alpha1.ComputeGlobalForwardingRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computeglobalforwardingrulesResource, c.ns, computeGlobalForwardingRule), &v1alpha1.ComputeGlobalForwardingRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeGlobalForwardingRule), err
}

// Update takes the representation of a computeGlobalForwardingRule and updates it. Returns the server's representation of the computeGlobalForwardingRule, and an error, if there is any.
func (c *FakeComputeGlobalForwardingRules) Update(computeGlobalForwardingRule *v1alpha1.ComputeGlobalForwardingRule) (result *v1alpha1.ComputeGlobalForwardingRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computeglobalforwardingrulesResource, c.ns, computeGlobalForwardingRule), &v1alpha1.ComputeGlobalForwardingRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeGlobalForwardingRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeGlobalForwardingRules) UpdateStatus(computeGlobalForwardingRule *v1alpha1.ComputeGlobalForwardingRule) (*v1alpha1.ComputeGlobalForwardingRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computeglobalforwardingrulesResource, "status", c.ns, computeGlobalForwardingRule), &v1alpha1.ComputeGlobalForwardingRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeGlobalForwardingRule), err
}

// Delete takes name of the computeGlobalForwardingRule and deletes it. Returns an error if one occurs.
func (c *FakeComputeGlobalForwardingRules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(computeglobalforwardingrulesResource, c.ns, name), &v1alpha1.ComputeGlobalForwardingRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeGlobalForwardingRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computeglobalforwardingrulesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComputeGlobalForwardingRuleList{})
	return err
}

// Patch applies the patch and returns the patched computeGlobalForwardingRule.
func (c *FakeComputeGlobalForwardingRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeGlobalForwardingRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computeglobalforwardingrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ComputeGlobalForwardingRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeGlobalForwardingRule), err
}
