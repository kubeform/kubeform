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

// FakeEbsDefaultKmsKeys implements EbsDefaultKmsKeyInterface
type FakeEbsDefaultKmsKeys struct {
	Fake *FakeAwsV1alpha1
}

var ebsdefaultkmskeysResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "ebsdefaultkmskeys"}

var ebsdefaultkmskeysKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "EbsDefaultKmsKey"}

// Get takes name of the ebsDefaultKmsKey, and returns the corresponding ebsDefaultKmsKey object, and an error if there is any.
func (c *FakeEbsDefaultKmsKeys) Get(name string, options v1.GetOptions) (result *v1alpha1.EbsDefaultKmsKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(ebsdefaultkmskeysResource, name), &v1alpha1.EbsDefaultKmsKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EbsDefaultKmsKey), err
}

// List takes label and field selectors, and returns the list of EbsDefaultKmsKeys that match those selectors.
func (c *FakeEbsDefaultKmsKeys) List(opts v1.ListOptions) (result *v1alpha1.EbsDefaultKmsKeyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(ebsdefaultkmskeysResource, ebsdefaultkmskeysKind, opts), &v1alpha1.EbsDefaultKmsKeyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EbsDefaultKmsKeyList{ListMeta: obj.(*v1alpha1.EbsDefaultKmsKeyList).ListMeta}
	for _, item := range obj.(*v1alpha1.EbsDefaultKmsKeyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ebsDefaultKmsKeys.
func (c *FakeEbsDefaultKmsKeys) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(ebsdefaultkmskeysResource, opts))
}

// Create takes the representation of a ebsDefaultKmsKey and creates it.  Returns the server's representation of the ebsDefaultKmsKey, and an error, if there is any.
func (c *FakeEbsDefaultKmsKeys) Create(ebsDefaultKmsKey *v1alpha1.EbsDefaultKmsKey) (result *v1alpha1.EbsDefaultKmsKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(ebsdefaultkmskeysResource, ebsDefaultKmsKey), &v1alpha1.EbsDefaultKmsKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EbsDefaultKmsKey), err
}

// Update takes the representation of a ebsDefaultKmsKey and updates it. Returns the server's representation of the ebsDefaultKmsKey, and an error, if there is any.
func (c *FakeEbsDefaultKmsKeys) Update(ebsDefaultKmsKey *v1alpha1.EbsDefaultKmsKey) (result *v1alpha1.EbsDefaultKmsKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(ebsdefaultkmskeysResource, ebsDefaultKmsKey), &v1alpha1.EbsDefaultKmsKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EbsDefaultKmsKey), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEbsDefaultKmsKeys) UpdateStatus(ebsDefaultKmsKey *v1alpha1.EbsDefaultKmsKey) (*v1alpha1.EbsDefaultKmsKey, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(ebsdefaultkmskeysResource, "status", ebsDefaultKmsKey), &v1alpha1.EbsDefaultKmsKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EbsDefaultKmsKey), err
}

// Delete takes name of the ebsDefaultKmsKey and deletes it. Returns an error if one occurs.
func (c *FakeEbsDefaultKmsKeys) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(ebsdefaultkmskeysResource, name), &v1alpha1.EbsDefaultKmsKey{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEbsDefaultKmsKeys) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(ebsdefaultkmskeysResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.EbsDefaultKmsKeyList{})
	return err
}

// Patch applies the patch and returns the patched ebsDefaultKmsKey.
func (c *FakeEbsDefaultKmsKeys) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EbsDefaultKmsKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(ebsdefaultkmskeysResource, name, pt, data, subresources...), &v1alpha1.EbsDefaultKmsKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EbsDefaultKmsKey), err
}
