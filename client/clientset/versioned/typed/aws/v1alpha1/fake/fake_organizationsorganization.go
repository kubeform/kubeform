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

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOrganizationsOrganizations implements OrganizationsOrganizationInterface
type FakeOrganizationsOrganizations struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var organizationsorganizationsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "organizationsorganizations"}

var organizationsorganizationsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "OrganizationsOrganization"}

// Get takes name of the organizationsOrganization, and returns the corresponding organizationsOrganization object, and an error if there is any.
func (c *FakeOrganizationsOrganizations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.OrganizationsOrganization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(organizationsorganizationsResource, c.ns, name), &v1alpha1.OrganizationsOrganization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationsOrganization), err
}

// List takes label and field selectors, and returns the list of OrganizationsOrganizations that match those selectors.
func (c *FakeOrganizationsOrganizations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.OrganizationsOrganizationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(organizationsorganizationsResource, organizationsorganizationsKind, c.ns, opts), &v1alpha1.OrganizationsOrganizationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OrganizationsOrganizationList{ListMeta: obj.(*v1alpha1.OrganizationsOrganizationList).ListMeta}
	for _, item := range obj.(*v1alpha1.OrganizationsOrganizationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested organizationsOrganizations.
func (c *FakeOrganizationsOrganizations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(organizationsorganizationsResource, c.ns, opts))

}

// Create takes the representation of a organizationsOrganization and creates it.  Returns the server's representation of the organizationsOrganization, and an error, if there is any.
func (c *FakeOrganizationsOrganizations) Create(ctx context.Context, organizationsOrganization *v1alpha1.OrganizationsOrganization, opts v1.CreateOptions) (result *v1alpha1.OrganizationsOrganization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(organizationsorganizationsResource, c.ns, organizationsOrganization), &v1alpha1.OrganizationsOrganization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationsOrganization), err
}

// Update takes the representation of a organizationsOrganization and updates it. Returns the server's representation of the organizationsOrganization, and an error, if there is any.
func (c *FakeOrganizationsOrganizations) Update(ctx context.Context, organizationsOrganization *v1alpha1.OrganizationsOrganization, opts v1.UpdateOptions) (result *v1alpha1.OrganizationsOrganization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(organizationsorganizationsResource, c.ns, organizationsOrganization), &v1alpha1.OrganizationsOrganization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationsOrganization), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOrganizationsOrganizations) UpdateStatus(ctx context.Context, organizationsOrganization *v1alpha1.OrganizationsOrganization, opts v1.UpdateOptions) (*v1alpha1.OrganizationsOrganization, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(organizationsorganizationsResource, "status", c.ns, organizationsOrganization), &v1alpha1.OrganizationsOrganization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationsOrganization), err
}

// Delete takes name of the organizationsOrganization and deletes it. Returns an error if one occurs.
func (c *FakeOrganizationsOrganizations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(organizationsorganizationsResource, c.ns, name), &v1alpha1.OrganizationsOrganization{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOrganizationsOrganizations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(organizationsorganizationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.OrganizationsOrganizationList{})
	return err
}

// Patch applies the patch and returns the patched organizationsOrganization.
func (c *FakeOrganizationsOrganizations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OrganizationsOrganization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(organizationsorganizationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.OrganizationsOrganization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationsOrganization), err
}
