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

// FakeAwsGuarddutyThreatintelsets implements AwsGuarddutyThreatintelsetInterface
type FakeAwsGuarddutyThreatintelsets struct {
	Fake *FakeAwsV1alpha1
}

var awsguarddutythreatintelsetsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsguarddutythreatintelsets"}

var awsguarddutythreatintelsetsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsGuarddutyThreatintelset"}

// Get takes name of the awsGuarddutyThreatintelset, and returns the corresponding awsGuarddutyThreatintelset object, and an error if there is any.
func (c *FakeAwsGuarddutyThreatintelsets) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsGuarddutyThreatintelset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsguarddutythreatintelsetsResource, name), &v1alpha1.AwsGuarddutyThreatintelset{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGuarddutyThreatintelset), err
}

// List takes label and field selectors, and returns the list of AwsGuarddutyThreatintelsets that match those selectors.
func (c *FakeAwsGuarddutyThreatintelsets) List(opts v1.ListOptions) (result *v1alpha1.AwsGuarddutyThreatintelsetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsguarddutythreatintelsetsResource, awsguarddutythreatintelsetsKind, opts), &v1alpha1.AwsGuarddutyThreatintelsetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsGuarddutyThreatintelsetList{ListMeta: obj.(*v1alpha1.AwsGuarddutyThreatintelsetList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsGuarddutyThreatintelsetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsGuarddutyThreatintelsets.
func (c *FakeAwsGuarddutyThreatintelsets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsguarddutythreatintelsetsResource, opts))
}

// Create takes the representation of a awsGuarddutyThreatintelset and creates it.  Returns the server's representation of the awsGuarddutyThreatintelset, and an error, if there is any.
func (c *FakeAwsGuarddutyThreatintelsets) Create(awsGuarddutyThreatintelset *v1alpha1.AwsGuarddutyThreatintelset) (result *v1alpha1.AwsGuarddutyThreatintelset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsguarddutythreatintelsetsResource, awsGuarddutyThreatintelset), &v1alpha1.AwsGuarddutyThreatintelset{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGuarddutyThreatintelset), err
}

// Update takes the representation of a awsGuarddutyThreatintelset and updates it. Returns the server's representation of the awsGuarddutyThreatintelset, and an error, if there is any.
func (c *FakeAwsGuarddutyThreatintelsets) Update(awsGuarddutyThreatintelset *v1alpha1.AwsGuarddutyThreatintelset) (result *v1alpha1.AwsGuarddutyThreatintelset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsguarddutythreatintelsetsResource, awsGuarddutyThreatintelset), &v1alpha1.AwsGuarddutyThreatintelset{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGuarddutyThreatintelset), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsGuarddutyThreatintelsets) UpdateStatus(awsGuarddutyThreatintelset *v1alpha1.AwsGuarddutyThreatintelset) (*v1alpha1.AwsGuarddutyThreatintelset, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsguarddutythreatintelsetsResource, "status", awsGuarddutyThreatintelset), &v1alpha1.AwsGuarddutyThreatintelset{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGuarddutyThreatintelset), err
}

// Delete takes name of the awsGuarddutyThreatintelset and deletes it. Returns an error if one occurs.
func (c *FakeAwsGuarddutyThreatintelsets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsguarddutythreatintelsetsResource, name), &v1alpha1.AwsGuarddutyThreatintelset{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsGuarddutyThreatintelsets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsguarddutythreatintelsetsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsGuarddutyThreatintelsetList{})
	return err
}

// Patch applies the patch and returns the patched awsGuarddutyThreatintelset.
func (c *FakeAwsGuarddutyThreatintelsets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsGuarddutyThreatintelset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsguarddutythreatintelsetsResource, name, pt, data, subresources...), &v1alpha1.AwsGuarddutyThreatintelset{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGuarddutyThreatintelset), err
}
