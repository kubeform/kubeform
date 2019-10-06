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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeSsmAssociations implements SsmAssociationInterface
type FakeSsmAssociations struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var ssmassociationsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "ssmassociations"}

var ssmassociationsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "SsmAssociation"}

// Get takes name of the ssmAssociation, and returns the corresponding ssmAssociation object, and an error if there is any.
func (c *FakeSsmAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.SsmAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ssmassociationsResource, c.ns, name), &v1alpha1.SsmAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsmAssociation), err
}

// List takes label and field selectors, and returns the list of SsmAssociations that match those selectors.
func (c *FakeSsmAssociations) List(opts v1.ListOptions) (result *v1alpha1.SsmAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ssmassociationsResource, ssmassociationsKind, c.ns, opts), &v1alpha1.SsmAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SsmAssociationList{ListMeta: obj.(*v1alpha1.SsmAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.SsmAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ssmAssociations.
func (c *FakeSsmAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ssmassociationsResource, c.ns, opts))

}

// Create takes the representation of a ssmAssociation and creates it.  Returns the server's representation of the ssmAssociation, and an error, if there is any.
func (c *FakeSsmAssociations) Create(ssmAssociation *v1alpha1.SsmAssociation) (result *v1alpha1.SsmAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ssmassociationsResource, c.ns, ssmAssociation), &v1alpha1.SsmAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsmAssociation), err
}

// Update takes the representation of a ssmAssociation and updates it. Returns the server's representation of the ssmAssociation, and an error, if there is any.
func (c *FakeSsmAssociations) Update(ssmAssociation *v1alpha1.SsmAssociation) (result *v1alpha1.SsmAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ssmassociationsResource, c.ns, ssmAssociation), &v1alpha1.SsmAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsmAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSsmAssociations) UpdateStatus(ssmAssociation *v1alpha1.SsmAssociation) (*v1alpha1.SsmAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ssmassociationsResource, "status", c.ns, ssmAssociation), &v1alpha1.SsmAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsmAssociation), err
}

// Delete takes name of the ssmAssociation and deletes it. Returns an error if one occurs.
func (c *FakeSsmAssociations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ssmassociationsResource, c.ns, name), &v1alpha1.SsmAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSsmAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ssmassociationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SsmAssociationList{})
	return err
}

// Patch applies the patch and returns the patched ssmAssociation.
func (c *FakeSsmAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SsmAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ssmassociationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SsmAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsmAssociation), err
}
