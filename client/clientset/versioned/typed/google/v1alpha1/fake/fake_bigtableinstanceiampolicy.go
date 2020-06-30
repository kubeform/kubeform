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

	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBigtableInstanceIamPolicies implements BigtableInstanceIamPolicyInterface
type FakeBigtableInstanceIamPolicies struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var bigtableinstanceiampoliciesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "bigtableinstanceiampolicies"}

var bigtableinstanceiampoliciesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "BigtableInstanceIamPolicy"}

// Get takes name of the bigtableInstanceIamPolicy, and returns the corresponding bigtableInstanceIamPolicy object, and an error if there is any.
func (c *FakeBigtableInstanceIamPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BigtableInstanceIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(bigtableinstanceiampoliciesResource, c.ns, name), &v1alpha1.BigtableInstanceIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BigtableInstanceIamPolicy), err
}

// List takes label and field selectors, and returns the list of BigtableInstanceIamPolicies that match those selectors.
func (c *FakeBigtableInstanceIamPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BigtableInstanceIamPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(bigtableinstanceiampoliciesResource, bigtableinstanceiampoliciesKind, c.ns, opts), &v1alpha1.BigtableInstanceIamPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BigtableInstanceIamPolicyList{ListMeta: obj.(*v1alpha1.BigtableInstanceIamPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.BigtableInstanceIamPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested bigtableInstanceIamPolicies.
func (c *FakeBigtableInstanceIamPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(bigtableinstanceiampoliciesResource, c.ns, opts))

}

// Create takes the representation of a bigtableInstanceIamPolicy and creates it.  Returns the server's representation of the bigtableInstanceIamPolicy, and an error, if there is any.
func (c *FakeBigtableInstanceIamPolicies) Create(ctx context.Context, bigtableInstanceIamPolicy *v1alpha1.BigtableInstanceIamPolicy, opts v1.CreateOptions) (result *v1alpha1.BigtableInstanceIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(bigtableinstanceiampoliciesResource, c.ns, bigtableInstanceIamPolicy), &v1alpha1.BigtableInstanceIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BigtableInstanceIamPolicy), err
}

// Update takes the representation of a bigtableInstanceIamPolicy and updates it. Returns the server's representation of the bigtableInstanceIamPolicy, and an error, if there is any.
func (c *FakeBigtableInstanceIamPolicies) Update(ctx context.Context, bigtableInstanceIamPolicy *v1alpha1.BigtableInstanceIamPolicy, opts v1.UpdateOptions) (result *v1alpha1.BigtableInstanceIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(bigtableinstanceiampoliciesResource, c.ns, bigtableInstanceIamPolicy), &v1alpha1.BigtableInstanceIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BigtableInstanceIamPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBigtableInstanceIamPolicies) UpdateStatus(ctx context.Context, bigtableInstanceIamPolicy *v1alpha1.BigtableInstanceIamPolicy, opts v1.UpdateOptions) (*v1alpha1.BigtableInstanceIamPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(bigtableinstanceiampoliciesResource, "status", c.ns, bigtableInstanceIamPolicy), &v1alpha1.BigtableInstanceIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BigtableInstanceIamPolicy), err
}

// Delete takes name of the bigtableInstanceIamPolicy and deletes it. Returns an error if one occurs.
func (c *FakeBigtableInstanceIamPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(bigtableinstanceiampoliciesResource, c.ns, name), &v1alpha1.BigtableInstanceIamPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBigtableInstanceIamPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(bigtableinstanceiampoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BigtableInstanceIamPolicyList{})
	return err
}

// Patch applies the patch and returns the patched bigtableInstanceIamPolicy.
func (c *FakeBigtableInstanceIamPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BigtableInstanceIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(bigtableinstanceiampoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.BigtableInstanceIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BigtableInstanceIamPolicy), err
}