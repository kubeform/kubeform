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

// FakeAwsInspectorAssessmentTargets implements AwsInspectorAssessmentTargetInterface
type FakeAwsInspectorAssessmentTargets struct {
	Fake *FakeAwsV1alpha1
}

var awsinspectorassessmenttargetsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsinspectorassessmenttargets"}

var awsinspectorassessmenttargetsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsInspectorAssessmentTarget"}

// Get takes name of the awsInspectorAssessmentTarget, and returns the corresponding awsInspectorAssessmentTarget object, and an error if there is any.
func (c *FakeAwsInspectorAssessmentTargets) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsInspectorAssessmentTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsinspectorassessmenttargetsResource, name), &v1alpha1.AwsInspectorAssessmentTarget{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsInspectorAssessmentTarget), err
}

// List takes label and field selectors, and returns the list of AwsInspectorAssessmentTargets that match those selectors.
func (c *FakeAwsInspectorAssessmentTargets) List(opts v1.ListOptions) (result *v1alpha1.AwsInspectorAssessmentTargetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsinspectorassessmenttargetsResource, awsinspectorassessmenttargetsKind, opts), &v1alpha1.AwsInspectorAssessmentTargetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsInspectorAssessmentTargetList{ListMeta: obj.(*v1alpha1.AwsInspectorAssessmentTargetList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsInspectorAssessmentTargetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsInspectorAssessmentTargets.
func (c *FakeAwsInspectorAssessmentTargets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsinspectorassessmenttargetsResource, opts))
}

// Create takes the representation of a awsInspectorAssessmentTarget and creates it.  Returns the server's representation of the awsInspectorAssessmentTarget, and an error, if there is any.
func (c *FakeAwsInspectorAssessmentTargets) Create(awsInspectorAssessmentTarget *v1alpha1.AwsInspectorAssessmentTarget) (result *v1alpha1.AwsInspectorAssessmentTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsinspectorassessmenttargetsResource, awsInspectorAssessmentTarget), &v1alpha1.AwsInspectorAssessmentTarget{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsInspectorAssessmentTarget), err
}

// Update takes the representation of a awsInspectorAssessmentTarget and updates it. Returns the server's representation of the awsInspectorAssessmentTarget, and an error, if there is any.
func (c *FakeAwsInspectorAssessmentTargets) Update(awsInspectorAssessmentTarget *v1alpha1.AwsInspectorAssessmentTarget) (result *v1alpha1.AwsInspectorAssessmentTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsinspectorassessmenttargetsResource, awsInspectorAssessmentTarget), &v1alpha1.AwsInspectorAssessmentTarget{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsInspectorAssessmentTarget), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsInspectorAssessmentTargets) UpdateStatus(awsInspectorAssessmentTarget *v1alpha1.AwsInspectorAssessmentTarget) (*v1alpha1.AwsInspectorAssessmentTarget, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsinspectorassessmenttargetsResource, "status", awsInspectorAssessmentTarget), &v1alpha1.AwsInspectorAssessmentTarget{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsInspectorAssessmentTarget), err
}

// Delete takes name of the awsInspectorAssessmentTarget and deletes it. Returns an error if one occurs.
func (c *FakeAwsInspectorAssessmentTargets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsinspectorassessmenttargetsResource, name), &v1alpha1.AwsInspectorAssessmentTarget{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsInspectorAssessmentTargets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsinspectorassessmenttargetsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsInspectorAssessmentTargetList{})
	return err
}

// Patch applies the patch and returns the patched awsInspectorAssessmentTarget.
func (c *FakeAwsInspectorAssessmentTargets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsInspectorAssessmentTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsinspectorassessmenttargetsResource, name, pt, data, subresources...), &v1alpha1.AwsInspectorAssessmentTarget{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsInspectorAssessmentTarget), err
}
