/*
Copyright 2019 The KubeDB Authors.

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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeGoogleFolderOrganizationPolicies implements GoogleFolderOrganizationPolicyInterface
type FakeGoogleFolderOrganizationPolicies struct {
	Fake *FakeGoogleV1alpha1
}

var googlefolderorganizationpoliciesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googlefolderorganizationpolicies"}

var googlefolderorganizationpoliciesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GoogleFolderOrganizationPolicy"}

// Get takes name of the googleFolderOrganizationPolicy, and returns the corresponding googleFolderOrganizationPolicy object, and an error if there is any.
func (c *FakeGoogleFolderOrganizationPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleFolderOrganizationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googlefolderorganizationpoliciesResource, name), &v1alpha1.GoogleFolderOrganizationPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleFolderOrganizationPolicy), err
}

// List takes label and field selectors, and returns the list of GoogleFolderOrganizationPolicies that match those selectors.
func (c *FakeGoogleFolderOrganizationPolicies) List(opts v1.ListOptions) (result *v1alpha1.GoogleFolderOrganizationPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googlefolderorganizationpoliciesResource, googlefolderorganizationpoliciesKind, opts), &v1alpha1.GoogleFolderOrganizationPolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleFolderOrganizationPolicyList{ListMeta: obj.(*v1alpha1.GoogleFolderOrganizationPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleFolderOrganizationPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleFolderOrganizationPolicies.
func (c *FakeGoogleFolderOrganizationPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googlefolderorganizationpoliciesResource, opts))
}

// Create takes the representation of a googleFolderOrganizationPolicy and creates it.  Returns the server's representation of the googleFolderOrganizationPolicy, and an error, if there is any.
func (c *FakeGoogleFolderOrganizationPolicies) Create(googleFolderOrganizationPolicy *v1alpha1.GoogleFolderOrganizationPolicy) (result *v1alpha1.GoogleFolderOrganizationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googlefolderorganizationpoliciesResource, googleFolderOrganizationPolicy), &v1alpha1.GoogleFolderOrganizationPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleFolderOrganizationPolicy), err
}

// Update takes the representation of a googleFolderOrganizationPolicy and updates it. Returns the server's representation of the googleFolderOrganizationPolicy, and an error, if there is any.
func (c *FakeGoogleFolderOrganizationPolicies) Update(googleFolderOrganizationPolicy *v1alpha1.GoogleFolderOrganizationPolicy) (result *v1alpha1.GoogleFolderOrganizationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googlefolderorganizationpoliciesResource, googleFolderOrganizationPolicy), &v1alpha1.GoogleFolderOrganizationPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleFolderOrganizationPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleFolderOrganizationPolicies) UpdateStatus(googleFolderOrganizationPolicy *v1alpha1.GoogleFolderOrganizationPolicy) (*v1alpha1.GoogleFolderOrganizationPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googlefolderorganizationpoliciesResource, "status", googleFolderOrganizationPolicy), &v1alpha1.GoogleFolderOrganizationPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleFolderOrganizationPolicy), err
}

// Delete takes name of the googleFolderOrganizationPolicy and deletes it. Returns an error if one occurs.
func (c *FakeGoogleFolderOrganizationPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googlefolderorganizationpoliciesResource, name), &v1alpha1.GoogleFolderOrganizationPolicy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleFolderOrganizationPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googlefolderorganizationpoliciesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleFolderOrganizationPolicyList{})
	return err
}

// Patch applies the patch and returns the patched googleFolderOrganizationPolicy.
func (c *FakeGoogleFolderOrganizationPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleFolderOrganizationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googlefolderorganizationpoliciesResource, name, pt, data, subresources...), &v1alpha1.GoogleFolderOrganizationPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleFolderOrganizationPolicy), err
}
