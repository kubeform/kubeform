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

// FakeAwsWafregionalRegexMatchSets implements AwsWafregionalRegexMatchSetInterface
type FakeAwsWafregionalRegexMatchSets struct {
	Fake *FakeAwsV1alpha1
}

var awswafregionalregexmatchsetsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awswafregionalregexmatchsets"}

var awswafregionalregexmatchsetsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsWafregionalRegexMatchSet"}

// Get takes name of the awsWafregionalRegexMatchSet, and returns the corresponding awsWafregionalRegexMatchSet object, and an error if there is any.
func (c *FakeAwsWafregionalRegexMatchSets) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsWafregionalRegexMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awswafregionalregexmatchsetsResource, name), &v1alpha1.AwsWafregionalRegexMatchSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafregionalRegexMatchSet), err
}

// List takes label and field selectors, and returns the list of AwsWafregionalRegexMatchSets that match those selectors.
func (c *FakeAwsWafregionalRegexMatchSets) List(opts v1.ListOptions) (result *v1alpha1.AwsWafregionalRegexMatchSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awswafregionalregexmatchsetsResource, awswafregionalregexmatchsetsKind, opts), &v1alpha1.AwsWafregionalRegexMatchSetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsWafregionalRegexMatchSetList{ListMeta: obj.(*v1alpha1.AwsWafregionalRegexMatchSetList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsWafregionalRegexMatchSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsWafregionalRegexMatchSets.
func (c *FakeAwsWafregionalRegexMatchSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awswafregionalregexmatchsetsResource, opts))
}

// Create takes the representation of a awsWafregionalRegexMatchSet and creates it.  Returns the server's representation of the awsWafregionalRegexMatchSet, and an error, if there is any.
func (c *FakeAwsWafregionalRegexMatchSets) Create(awsWafregionalRegexMatchSet *v1alpha1.AwsWafregionalRegexMatchSet) (result *v1alpha1.AwsWafregionalRegexMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awswafregionalregexmatchsetsResource, awsWafregionalRegexMatchSet), &v1alpha1.AwsWafregionalRegexMatchSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafregionalRegexMatchSet), err
}

// Update takes the representation of a awsWafregionalRegexMatchSet and updates it. Returns the server's representation of the awsWafregionalRegexMatchSet, and an error, if there is any.
func (c *FakeAwsWafregionalRegexMatchSets) Update(awsWafregionalRegexMatchSet *v1alpha1.AwsWafregionalRegexMatchSet) (result *v1alpha1.AwsWafregionalRegexMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awswafregionalregexmatchsetsResource, awsWafregionalRegexMatchSet), &v1alpha1.AwsWafregionalRegexMatchSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafregionalRegexMatchSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsWafregionalRegexMatchSets) UpdateStatus(awsWafregionalRegexMatchSet *v1alpha1.AwsWafregionalRegexMatchSet) (*v1alpha1.AwsWafregionalRegexMatchSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awswafregionalregexmatchsetsResource, "status", awsWafregionalRegexMatchSet), &v1alpha1.AwsWafregionalRegexMatchSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafregionalRegexMatchSet), err
}

// Delete takes name of the awsWafregionalRegexMatchSet and deletes it. Returns an error if one occurs.
func (c *FakeAwsWafregionalRegexMatchSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awswafregionalregexmatchsetsResource, name), &v1alpha1.AwsWafregionalRegexMatchSet{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsWafregionalRegexMatchSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awswafregionalregexmatchsetsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsWafregionalRegexMatchSetList{})
	return err
}

// Patch applies the patch and returns the patched awsWafregionalRegexMatchSet.
func (c *FakeAwsWafregionalRegexMatchSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafregionalRegexMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awswafregionalregexmatchsetsResource, name, pt, data, subresources...), &v1alpha1.AwsWafregionalRegexMatchSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafregionalRegexMatchSet), err
}
