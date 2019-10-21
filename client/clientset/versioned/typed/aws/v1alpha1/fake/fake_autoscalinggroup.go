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

// FakeAutoscalingGroups implements AutoscalingGroupInterface
type FakeAutoscalingGroups struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var autoscalinggroupsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "autoscalinggroups"}

var autoscalinggroupsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AutoscalingGroup"}

// Get takes name of the autoscalingGroup, and returns the corresponding autoscalingGroup object, and an error if there is any.
func (c *FakeAutoscalingGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AutoscalingGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(autoscalinggroupsResource, c.ns, name), &v1alpha1.AutoscalingGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoscalingGroup), err
}

// List takes label and field selectors, and returns the list of AutoscalingGroups that match those selectors.
func (c *FakeAutoscalingGroups) List(opts v1.ListOptions) (result *v1alpha1.AutoscalingGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(autoscalinggroupsResource, autoscalinggroupsKind, c.ns, opts), &v1alpha1.AutoscalingGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AutoscalingGroupList{ListMeta: obj.(*v1alpha1.AutoscalingGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.AutoscalingGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested autoscalingGroups.
func (c *FakeAutoscalingGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(autoscalinggroupsResource, c.ns, opts))

}

// Create takes the representation of a autoscalingGroup and creates it.  Returns the server's representation of the autoscalingGroup, and an error, if there is any.
func (c *FakeAutoscalingGroups) Create(autoscalingGroup *v1alpha1.AutoscalingGroup) (result *v1alpha1.AutoscalingGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(autoscalinggroupsResource, c.ns, autoscalingGroup), &v1alpha1.AutoscalingGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoscalingGroup), err
}

// Update takes the representation of a autoscalingGroup and updates it. Returns the server's representation of the autoscalingGroup, and an error, if there is any.
func (c *FakeAutoscalingGroups) Update(autoscalingGroup *v1alpha1.AutoscalingGroup) (result *v1alpha1.AutoscalingGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(autoscalinggroupsResource, c.ns, autoscalingGroup), &v1alpha1.AutoscalingGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoscalingGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAutoscalingGroups) UpdateStatus(autoscalingGroup *v1alpha1.AutoscalingGroup) (*v1alpha1.AutoscalingGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(autoscalinggroupsResource, "status", c.ns, autoscalingGroup), &v1alpha1.AutoscalingGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoscalingGroup), err
}

// Delete takes name of the autoscalingGroup and deletes it. Returns an error if one occurs.
func (c *FakeAutoscalingGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(autoscalinggroupsResource, c.ns, name), &v1alpha1.AutoscalingGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAutoscalingGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(autoscalinggroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AutoscalingGroupList{})
	return err
}

// Patch applies the patch and returns the patched autoscalingGroup.
func (c *FakeAutoscalingGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AutoscalingGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(autoscalinggroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AutoscalingGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoscalingGroup), err
}