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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeAwsWafGeoMatchSets implements AwsWafGeoMatchSetInterface
type FakeAwsWafGeoMatchSets struct {
	Fake *FakeAwsV1alpha1
}

var awswafgeomatchsetsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awswafgeomatchsets"}

var awswafgeomatchsetsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsWafGeoMatchSet"}

// Get takes name of the awsWafGeoMatchSet, and returns the corresponding awsWafGeoMatchSet object, and an error if there is any.
func (c *FakeAwsWafGeoMatchSets) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsWafGeoMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awswafgeomatchsetsResource, name), &v1alpha1.AwsWafGeoMatchSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafGeoMatchSet), err
}

// List takes label and field selectors, and returns the list of AwsWafGeoMatchSets that match those selectors.
func (c *FakeAwsWafGeoMatchSets) List(opts v1.ListOptions) (result *v1alpha1.AwsWafGeoMatchSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awswafgeomatchsetsResource, awswafgeomatchsetsKind, opts), &v1alpha1.AwsWafGeoMatchSetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsWafGeoMatchSetList{ListMeta: obj.(*v1alpha1.AwsWafGeoMatchSetList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsWafGeoMatchSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsWafGeoMatchSets.
func (c *FakeAwsWafGeoMatchSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awswafgeomatchsetsResource, opts))
}

// Create takes the representation of a awsWafGeoMatchSet and creates it.  Returns the server's representation of the awsWafGeoMatchSet, and an error, if there is any.
func (c *FakeAwsWafGeoMatchSets) Create(awsWafGeoMatchSet *v1alpha1.AwsWafGeoMatchSet) (result *v1alpha1.AwsWafGeoMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awswafgeomatchsetsResource, awsWafGeoMatchSet), &v1alpha1.AwsWafGeoMatchSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafGeoMatchSet), err
}

// Update takes the representation of a awsWafGeoMatchSet and updates it. Returns the server's representation of the awsWafGeoMatchSet, and an error, if there is any.
func (c *FakeAwsWafGeoMatchSets) Update(awsWafGeoMatchSet *v1alpha1.AwsWafGeoMatchSet) (result *v1alpha1.AwsWafGeoMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awswafgeomatchsetsResource, awsWafGeoMatchSet), &v1alpha1.AwsWafGeoMatchSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafGeoMatchSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsWafGeoMatchSets) UpdateStatus(awsWafGeoMatchSet *v1alpha1.AwsWafGeoMatchSet) (*v1alpha1.AwsWafGeoMatchSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awswafgeomatchsetsResource, "status", awsWafGeoMatchSet), &v1alpha1.AwsWafGeoMatchSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafGeoMatchSet), err
}

// Delete takes name of the awsWafGeoMatchSet and deletes it. Returns an error if one occurs.
func (c *FakeAwsWafGeoMatchSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awswafgeomatchsetsResource, name), &v1alpha1.AwsWafGeoMatchSet{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsWafGeoMatchSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awswafgeomatchsetsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsWafGeoMatchSetList{})
	return err
}

// Patch applies the patch and returns the patched awsWafGeoMatchSet.
func (c *FakeAwsWafGeoMatchSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafGeoMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awswafgeomatchsetsResource, name, pt, data, subresources...), &v1alpha1.AwsWafGeoMatchSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafGeoMatchSet), err
}
