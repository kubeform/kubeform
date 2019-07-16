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

// FakeGuarddutyThreatintelsets implements GuarddutyThreatintelsetInterface
type FakeGuarddutyThreatintelsets struct {
	Fake *FakeAwsV1alpha1
}

var guarddutythreatintelsetsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "guarddutythreatintelsets"}

var guarddutythreatintelsetsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "GuarddutyThreatintelset"}

// Get takes name of the guarddutyThreatintelset, and returns the corresponding guarddutyThreatintelset object, and an error if there is any.
func (c *FakeGuarddutyThreatintelsets) Get(name string, options v1.GetOptions) (result *v1alpha1.GuarddutyThreatintelset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(guarddutythreatintelsetsResource, name), &v1alpha1.GuarddutyThreatintelset{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GuarddutyThreatintelset), err
}

// List takes label and field selectors, and returns the list of GuarddutyThreatintelsets that match those selectors.
func (c *FakeGuarddutyThreatintelsets) List(opts v1.ListOptions) (result *v1alpha1.GuarddutyThreatintelsetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(guarddutythreatintelsetsResource, guarddutythreatintelsetsKind, opts), &v1alpha1.GuarddutyThreatintelsetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GuarddutyThreatintelsetList{ListMeta: obj.(*v1alpha1.GuarddutyThreatintelsetList).ListMeta}
	for _, item := range obj.(*v1alpha1.GuarddutyThreatintelsetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested guarddutyThreatintelsets.
func (c *FakeGuarddutyThreatintelsets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(guarddutythreatintelsetsResource, opts))
}

// Create takes the representation of a guarddutyThreatintelset and creates it.  Returns the server's representation of the guarddutyThreatintelset, and an error, if there is any.
func (c *FakeGuarddutyThreatintelsets) Create(guarddutyThreatintelset *v1alpha1.GuarddutyThreatintelset) (result *v1alpha1.GuarddutyThreatintelset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(guarddutythreatintelsetsResource, guarddutyThreatintelset), &v1alpha1.GuarddutyThreatintelset{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GuarddutyThreatintelset), err
}

// Update takes the representation of a guarddutyThreatintelset and updates it. Returns the server's representation of the guarddutyThreatintelset, and an error, if there is any.
func (c *FakeGuarddutyThreatintelsets) Update(guarddutyThreatintelset *v1alpha1.GuarddutyThreatintelset) (result *v1alpha1.GuarddutyThreatintelset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(guarddutythreatintelsetsResource, guarddutyThreatintelset), &v1alpha1.GuarddutyThreatintelset{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GuarddutyThreatintelset), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGuarddutyThreatintelsets) UpdateStatus(guarddutyThreatintelset *v1alpha1.GuarddutyThreatintelset) (*v1alpha1.GuarddutyThreatintelset, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(guarddutythreatintelsetsResource, "status", guarddutyThreatintelset), &v1alpha1.GuarddutyThreatintelset{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GuarddutyThreatintelset), err
}

// Delete takes name of the guarddutyThreatintelset and deletes it. Returns an error if one occurs.
func (c *FakeGuarddutyThreatintelsets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(guarddutythreatintelsetsResource, name), &v1alpha1.GuarddutyThreatintelset{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGuarddutyThreatintelsets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(guarddutythreatintelsetsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GuarddutyThreatintelsetList{})
	return err
}

// Patch applies the patch and returns the patched guarddutyThreatintelset.
func (c *FakeGuarddutyThreatintelsets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GuarddutyThreatintelset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(guarddutythreatintelsetsResource, name, pt, data, subresources...), &v1alpha1.GuarddutyThreatintelset{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GuarddutyThreatintelset), err
}
