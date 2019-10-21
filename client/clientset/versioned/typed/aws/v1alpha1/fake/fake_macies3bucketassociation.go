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

// FakeMacieS3BucketAssociations implements MacieS3BucketAssociationInterface
type FakeMacieS3BucketAssociations struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var macies3bucketassociationsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "macies3bucketassociations"}

var macies3bucketassociationsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "MacieS3BucketAssociation"}

// Get takes name of the macieS3BucketAssociation, and returns the corresponding macieS3BucketAssociation object, and an error if there is any.
func (c *FakeMacieS3BucketAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.MacieS3BucketAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(macies3bucketassociationsResource, c.ns, name), &v1alpha1.MacieS3BucketAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MacieS3BucketAssociation), err
}

// List takes label and field selectors, and returns the list of MacieS3BucketAssociations that match those selectors.
func (c *FakeMacieS3BucketAssociations) List(opts v1.ListOptions) (result *v1alpha1.MacieS3BucketAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(macies3bucketassociationsResource, macies3bucketassociationsKind, c.ns, opts), &v1alpha1.MacieS3BucketAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MacieS3BucketAssociationList{ListMeta: obj.(*v1alpha1.MacieS3BucketAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.MacieS3BucketAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested macieS3BucketAssociations.
func (c *FakeMacieS3BucketAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(macies3bucketassociationsResource, c.ns, opts))

}

// Create takes the representation of a macieS3BucketAssociation and creates it.  Returns the server's representation of the macieS3BucketAssociation, and an error, if there is any.
func (c *FakeMacieS3BucketAssociations) Create(macieS3BucketAssociation *v1alpha1.MacieS3BucketAssociation) (result *v1alpha1.MacieS3BucketAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(macies3bucketassociationsResource, c.ns, macieS3BucketAssociation), &v1alpha1.MacieS3BucketAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MacieS3BucketAssociation), err
}

// Update takes the representation of a macieS3BucketAssociation and updates it. Returns the server's representation of the macieS3BucketAssociation, and an error, if there is any.
func (c *FakeMacieS3BucketAssociations) Update(macieS3BucketAssociation *v1alpha1.MacieS3BucketAssociation) (result *v1alpha1.MacieS3BucketAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(macies3bucketassociationsResource, c.ns, macieS3BucketAssociation), &v1alpha1.MacieS3BucketAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MacieS3BucketAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMacieS3BucketAssociations) UpdateStatus(macieS3BucketAssociation *v1alpha1.MacieS3BucketAssociation) (*v1alpha1.MacieS3BucketAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(macies3bucketassociationsResource, "status", c.ns, macieS3BucketAssociation), &v1alpha1.MacieS3BucketAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MacieS3BucketAssociation), err
}

// Delete takes name of the macieS3BucketAssociation and deletes it. Returns an error if one occurs.
func (c *FakeMacieS3BucketAssociations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(macies3bucketassociationsResource, c.ns, name), &v1alpha1.MacieS3BucketAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMacieS3BucketAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(macies3bucketassociationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.MacieS3BucketAssociationList{})
	return err
}

// Patch applies the patch and returns the patched macieS3BucketAssociation.
func (c *FakeMacieS3BucketAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MacieS3BucketAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(macies3bucketassociationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.MacieS3BucketAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MacieS3BucketAssociation), err
}