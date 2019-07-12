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

// FakeAwsWafByteMatchSets implements AwsWafByteMatchSetInterface
type FakeAwsWafByteMatchSets struct {
	Fake *FakeAwsV1alpha1
}

var awswafbytematchsetsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awswafbytematchsets"}

var awswafbytematchsetsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsWafByteMatchSet"}

// Get takes name of the awsWafByteMatchSet, and returns the corresponding awsWafByteMatchSet object, and an error if there is any.
func (c *FakeAwsWafByteMatchSets) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsWafByteMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awswafbytematchsetsResource, name), &v1alpha1.AwsWafByteMatchSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafByteMatchSet), err
}

// List takes label and field selectors, and returns the list of AwsWafByteMatchSets that match those selectors.
func (c *FakeAwsWafByteMatchSets) List(opts v1.ListOptions) (result *v1alpha1.AwsWafByteMatchSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awswafbytematchsetsResource, awswafbytematchsetsKind, opts), &v1alpha1.AwsWafByteMatchSetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsWafByteMatchSetList{ListMeta: obj.(*v1alpha1.AwsWafByteMatchSetList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsWafByteMatchSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsWafByteMatchSets.
func (c *FakeAwsWafByteMatchSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awswafbytematchsetsResource, opts))
}

// Create takes the representation of a awsWafByteMatchSet and creates it.  Returns the server's representation of the awsWafByteMatchSet, and an error, if there is any.
func (c *FakeAwsWafByteMatchSets) Create(awsWafByteMatchSet *v1alpha1.AwsWafByteMatchSet) (result *v1alpha1.AwsWafByteMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awswafbytematchsetsResource, awsWafByteMatchSet), &v1alpha1.AwsWafByteMatchSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafByteMatchSet), err
}

// Update takes the representation of a awsWafByteMatchSet and updates it. Returns the server's representation of the awsWafByteMatchSet, and an error, if there is any.
func (c *FakeAwsWafByteMatchSets) Update(awsWafByteMatchSet *v1alpha1.AwsWafByteMatchSet) (result *v1alpha1.AwsWafByteMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awswafbytematchsetsResource, awsWafByteMatchSet), &v1alpha1.AwsWafByteMatchSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafByteMatchSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsWafByteMatchSets) UpdateStatus(awsWafByteMatchSet *v1alpha1.AwsWafByteMatchSet) (*v1alpha1.AwsWafByteMatchSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awswafbytematchsetsResource, "status", awsWafByteMatchSet), &v1alpha1.AwsWafByteMatchSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafByteMatchSet), err
}

// Delete takes name of the awsWafByteMatchSet and deletes it. Returns an error if one occurs.
func (c *FakeAwsWafByteMatchSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awswafbytematchsetsResource, name), &v1alpha1.AwsWafByteMatchSet{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsWafByteMatchSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awswafbytematchsetsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsWafByteMatchSetList{})
	return err
}

// Patch applies the patch and returns the patched awsWafByteMatchSet.
func (c *FakeAwsWafByteMatchSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafByteMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awswafbytematchsetsResource, name, pt, data, subresources...), &v1alpha1.AwsWafByteMatchSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafByteMatchSet), err
}
