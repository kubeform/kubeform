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

// FakeGoogleSpannerDatabaseIamPolicies implements GoogleSpannerDatabaseIamPolicyInterface
type FakeGoogleSpannerDatabaseIamPolicies struct {
	Fake *FakeGoogleV1alpha1
}

var googlespannerdatabaseiampoliciesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googlespannerdatabaseiampolicies"}

var googlespannerdatabaseiampoliciesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GoogleSpannerDatabaseIamPolicy"}

// Get takes name of the googleSpannerDatabaseIamPolicy, and returns the corresponding googleSpannerDatabaseIamPolicy object, and an error if there is any.
func (c *FakeGoogleSpannerDatabaseIamPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleSpannerDatabaseIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googlespannerdatabaseiampoliciesResource, name), &v1alpha1.GoogleSpannerDatabaseIamPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleSpannerDatabaseIamPolicy), err
}

// List takes label and field selectors, and returns the list of GoogleSpannerDatabaseIamPolicies that match those selectors.
func (c *FakeGoogleSpannerDatabaseIamPolicies) List(opts v1.ListOptions) (result *v1alpha1.GoogleSpannerDatabaseIamPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googlespannerdatabaseiampoliciesResource, googlespannerdatabaseiampoliciesKind, opts), &v1alpha1.GoogleSpannerDatabaseIamPolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleSpannerDatabaseIamPolicyList{ListMeta: obj.(*v1alpha1.GoogleSpannerDatabaseIamPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleSpannerDatabaseIamPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleSpannerDatabaseIamPolicies.
func (c *FakeGoogleSpannerDatabaseIamPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googlespannerdatabaseiampoliciesResource, opts))
}

// Create takes the representation of a googleSpannerDatabaseIamPolicy and creates it.  Returns the server's representation of the googleSpannerDatabaseIamPolicy, and an error, if there is any.
func (c *FakeGoogleSpannerDatabaseIamPolicies) Create(googleSpannerDatabaseIamPolicy *v1alpha1.GoogleSpannerDatabaseIamPolicy) (result *v1alpha1.GoogleSpannerDatabaseIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googlespannerdatabaseiampoliciesResource, googleSpannerDatabaseIamPolicy), &v1alpha1.GoogleSpannerDatabaseIamPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleSpannerDatabaseIamPolicy), err
}

// Update takes the representation of a googleSpannerDatabaseIamPolicy and updates it. Returns the server's representation of the googleSpannerDatabaseIamPolicy, and an error, if there is any.
func (c *FakeGoogleSpannerDatabaseIamPolicies) Update(googleSpannerDatabaseIamPolicy *v1alpha1.GoogleSpannerDatabaseIamPolicy) (result *v1alpha1.GoogleSpannerDatabaseIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googlespannerdatabaseiampoliciesResource, googleSpannerDatabaseIamPolicy), &v1alpha1.GoogleSpannerDatabaseIamPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleSpannerDatabaseIamPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleSpannerDatabaseIamPolicies) UpdateStatus(googleSpannerDatabaseIamPolicy *v1alpha1.GoogleSpannerDatabaseIamPolicy) (*v1alpha1.GoogleSpannerDatabaseIamPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googlespannerdatabaseiampoliciesResource, "status", googleSpannerDatabaseIamPolicy), &v1alpha1.GoogleSpannerDatabaseIamPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleSpannerDatabaseIamPolicy), err
}

// Delete takes name of the googleSpannerDatabaseIamPolicy and deletes it. Returns an error if one occurs.
func (c *FakeGoogleSpannerDatabaseIamPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googlespannerdatabaseiampoliciesResource, name), &v1alpha1.GoogleSpannerDatabaseIamPolicy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleSpannerDatabaseIamPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googlespannerdatabaseiampoliciesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleSpannerDatabaseIamPolicyList{})
	return err
}

// Patch applies the patch and returns the patched googleSpannerDatabaseIamPolicy.
func (c *FakeGoogleSpannerDatabaseIamPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleSpannerDatabaseIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googlespannerdatabaseiampoliciesResource, name, pt, data, subresources...), &v1alpha1.GoogleSpannerDatabaseIamPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleSpannerDatabaseIamPolicy), err
}
