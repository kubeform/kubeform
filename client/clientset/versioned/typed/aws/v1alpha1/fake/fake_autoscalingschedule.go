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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAutoscalingSchedules implements AutoscalingScheduleInterface
type FakeAutoscalingSchedules struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var autoscalingschedulesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "autoscalingschedules"}

var autoscalingschedulesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AutoscalingSchedule"}

// Get takes name of the autoscalingSchedule, and returns the corresponding autoscalingSchedule object, and an error if there is any.
func (c *FakeAutoscalingSchedules) Get(name string, options v1.GetOptions) (result *v1alpha1.AutoscalingSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(autoscalingschedulesResource, c.ns, name), &v1alpha1.AutoscalingSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoscalingSchedule), err
}

// List takes label and field selectors, and returns the list of AutoscalingSchedules that match those selectors.
func (c *FakeAutoscalingSchedules) List(opts v1.ListOptions) (result *v1alpha1.AutoscalingScheduleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(autoscalingschedulesResource, autoscalingschedulesKind, c.ns, opts), &v1alpha1.AutoscalingScheduleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AutoscalingScheduleList{ListMeta: obj.(*v1alpha1.AutoscalingScheduleList).ListMeta}
	for _, item := range obj.(*v1alpha1.AutoscalingScheduleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested autoscalingSchedules.
func (c *FakeAutoscalingSchedules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(autoscalingschedulesResource, c.ns, opts))

}

// Create takes the representation of a autoscalingSchedule and creates it.  Returns the server's representation of the autoscalingSchedule, and an error, if there is any.
func (c *FakeAutoscalingSchedules) Create(autoscalingSchedule *v1alpha1.AutoscalingSchedule) (result *v1alpha1.AutoscalingSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(autoscalingschedulesResource, c.ns, autoscalingSchedule), &v1alpha1.AutoscalingSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoscalingSchedule), err
}

// Update takes the representation of a autoscalingSchedule and updates it. Returns the server's representation of the autoscalingSchedule, and an error, if there is any.
func (c *FakeAutoscalingSchedules) Update(autoscalingSchedule *v1alpha1.AutoscalingSchedule) (result *v1alpha1.AutoscalingSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(autoscalingschedulesResource, c.ns, autoscalingSchedule), &v1alpha1.AutoscalingSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoscalingSchedule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAutoscalingSchedules) UpdateStatus(autoscalingSchedule *v1alpha1.AutoscalingSchedule) (*v1alpha1.AutoscalingSchedule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(autoscalingschedulesResource, "status", c.ns, autoscalingSchedule), &v1alpha1.AutoscalingSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoscalingSchedule), err
}

// Delete takes name of the autoscalingSchedule and deletes it. Returns an error if one occurs.
func (c *FakeAutoscalingSchedules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(autoscalingschedulesResource, c.ns, name), &v1alpha1.AutoscalingSchedule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAutoscalingSchedules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(autoscalingschedulesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AutoscalingScheduleList{})
	return err
}

// Patch applies the patch and returns the patched autoscalingSchedule.
func (c *FakeAutoscalingSchedules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AutoscalingSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(autoscalingschedulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AutoscalingSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoscalingSchedule), err
}
