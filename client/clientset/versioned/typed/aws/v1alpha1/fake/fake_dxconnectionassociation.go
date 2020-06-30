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

// FakeDxConnectionAssociations implements DxConnectionAssociationInterface
type FakeDxConnectionAssociations struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var dxconnectionassociationsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "dxconnectionassociations"}

var dxconnectionassociationsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DxConnectionAssociation"}

// Get takes name of the dxConnectionAssociation, and returns the corresponding dxConnectionAssociation object, and an error if there is any.
func (c *FakeDxConnectionAssociations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DxConnectionAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dxconnectionassociationsResource, c.ns, name), &v1alpha1.DxConnectionAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxConnectionAssociation), err
}

// List takes label and field selectors, and returns the list of DxConnectionAssociations that match those selectors.
func (c *FakeDxConnectionAssociations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DxConnectionAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dxconnectionassociationsResource, dxconnectionassociationsKind, c.ns, opts), &v1alpha1.DxConnectionAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DxConnectionAssociationList{ListMeta: obj.(*v1alpha1.DxConnectionAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.DxConnectionAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dxConnectionAssociations.
func (c *FakeDxConnectionAssociations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dxconnectionassociationsResource, c.ns, opts))

}

// Create takes the representation of a dxConnectionAssociation and creates it.  Returns the server's representation of the dxConnectionAssociation, and an error, if there is any.
func (c *FakeDxConnectionAssociations) Create(ctx context.Context, dxConnectionAssociation *v1alpha1.DxConnectionAssociation, opts v1.CreateOptions) (result *v1alpha1.DxConnectionAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dxconnectionassociationsResource, c.ns, dxConnectionAssociation), &v1alpha1.DxConnectionAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxConnectionAssociation), err
}

// Update takes the representation of a dxConnectionAssociation and updates it. Returns the server's representation of the dxConnectionAssociation, and an error, if there is any.
func (c *FakeDxConnectionAssociations) Update(ctx context.Context, dxConnectionAssociation *v1alpha1.DxConnectionAssociation, opts v1.UpdateOptions) (result *v1alpha1.DxConnectionAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dxconnectionassociationsResource, c.ns, dxConnectionAssociation), &v1alpha1.DxConnectionAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxConnectionAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDxConnectionAssociations) UpdateStatus(ctx context.Context, dxConnectionAssociation *v1alpha1.DxConnectionAssociation, opts v1.UpdateOptions) (*v1alpha1.DxConnectionAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dxconnectionassociationsResource, "status", c.ns, dxConnectionAssociation), &v1alpha1.DxConnectionAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxConnectionAssociation), err
}

// Delete takes name of the dxConnectionAssociation and deletes it. Returns an error if one occurs.
func (c *FakeDxConnectionAssociations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dxconnectionassociationsResource, c.ns, name), &v1alpha1.DxConnectionAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDxConnectionAssociations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dxconnectionassociationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DxConnectionAssociationList{})
	return err
}

// Patch applies the patch and returns the patched dxConnectionAssociation.
func (c *FakeDxConnectionAssociations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DxConnectionAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dxconnectionassociationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DxConnectionAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxConnectionAssociation), err
}
