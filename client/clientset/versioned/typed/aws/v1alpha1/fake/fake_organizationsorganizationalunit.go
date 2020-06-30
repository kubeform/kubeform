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

// FakeOrganizationsOrganizationalUnits implements OrganizationsOrganizationalUnitInterface
type FakeOrganizationsOrganizationalUnits struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var organizationsorganizationalunitsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "organizationsorganizationalunits"}

var organizationsorganizationalunitsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "OrganizationsOrganizationalUnit"}

// Get takes name of the organizationsOrganizationalUnit, and returns the corresponding organizationsOrganizationalUnit object, and an error if there is any.
func (c *FakeOrganizationsOrganizationalUnits) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.OrganizationsOrganizationalUnit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(organizationsorganizationalunitsResource, c.ns, name), &v1alpha1.OrganizationsOrganizationalUnit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationsOrganizationalUnit), err
}

// List takes label and field selectors, and returns the list of OrganizationsOrganizationalUnits that match those selectors.
func (c *FakeOrganizationsOrganizationalUnits) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.OrganizationsOrganizationalUnitList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(organizationsorganizationalunitsResource, organizationsorganizationalunitsKind, c.ns, opts), &v1alpha1.OrganizationsOrganizationalUnitList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OrganizationsOrganizationalUnitList{ListMeta: obj.(*v1alpha1.OrganizationsOrganizationalUnitList).ListMeta}
	for _, item := range obj.(*v1alpha1.OrganizationsOrganizationalUnitList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested organizationsOrganizationalUnits.
func (c *FakeOrganizationsOrganizationalUnits) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(organizationsorganizationalunitsResource, c.ns, opts))

}

// Create takes the representation of a organizationsOrganizationalUnit and creates it.  Returns the server's representation of the organizationsOrganizationalUnit, and an error, if there is any.
func (c *FakeOrganizationsOrganizationalUnits) Create(ctx context.Context, organizationsOrganizationalUnit *v1alpha1.OrganizationsOrganizationalUnit, opts v1.CreateOptions) (result *v1alpha1.OrganizationsOrganizationalUnit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(organizationsorganizationalunitsResource, c.ns, organizationsOrganizationalUnit), &v1alpha1.OrganizationsOrganizationalUnit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationsOrganizationalUnit), err
}

// Update takes the representation of a organizationsOrganizationalUnit and updates it. Returns the server's representation of the organizationsOrganizationalUnit, and an error, if there is any.
func (c *FakeOrganizationsOrganizationalUnits) Update(ctx context.Context, organizationsOrganizationalUnit *v1alpha1.OrganizationsOrganizationalUnit, opts v1.UpdateOptions) (result *v1alpha1.OrganizationsOrganizationalUnit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(organizationsorganizationalunitsResource, c.ns, organizationsOrganizationalUnit), &v1alpha1.OrganizationsOrganizationalUnit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationsOrganizationalUnit), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOrganizationsOrganizationalUnits) UpdateStatus(ctx context.Context, organizationsOrganizationalUnit *v1alpha1.OrganizationsOrganizationalUnit, opts v1.UpdateOptions) (*v1alpha1.OrganizationsOrganizationalUnit, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(organizationsorganizationalunitsResource, "status", c.ns, organizationsOrganizationalUnit), &v1alpha1.OrganizationsOrganizationalUnit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationsOrganizationalUnit), err
}

// Delete takes name of the organizationsOrganizationalUnit and deletes it. Returns an error if one occurs.
func (c *FakeOrganizationsOrganizationalUnits) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(organizationsorganizationalunitsResource, c.ns, name), &v1alpha1.OrganizationsOrganizationalUnit{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOrganizationsOrganizationalUnits) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(organizationsorganizationalunitsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.OrganizationsOrganizationalUnitList{})
	return err
}

// Patch applies the patch and returns the patched organizationsOrganizationalUnit.
func (c *FakeOrganizationsOrganizationalUnits) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OrganizationsOrganizationalUnit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(organizationsorganizationalunitsResource, c.ns, name, pt, data, subresources...), &v1alpha1.OrganizationsOrganizationalUnit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationsOrganizationalUnit), err
}
