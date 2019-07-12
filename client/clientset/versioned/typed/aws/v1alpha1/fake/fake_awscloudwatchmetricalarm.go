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

// FakeAwsCloudwatchMetricAlarms implements AwsCloudwatchMetricAlarmInterface
type FakeAwsCloudwatchMetricAlarms struct {
	Fake *FakeAwsV1alpha1
}

var awscloudwatchmetricalarmsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awscloudwatchmetricalarms"}

var awscloudwatchmetricalarmsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsCloudwatchMetricAlarm"}

// Get takes name of the awsCloudwatchMetricAlarm, and returns the corresponding awsCloudwatchMetricAlarm object, and an error if there is any.
func (c *FakeAwsCloudwatchMetricAlarms) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsCloudwatchMetricAlarm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awscloudwatchmetricalarmsResource, name), &v1alpha1.AwsCloudwatchMetricAlarm{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCloudwatchMetricAlarm), err
}

// List takes label and field selectors, and returns the list of AwsCloudwatchMetricAlarms that match those selectors.
func (c *FakeAwsCloudwatchMetricAlarms) List(opts v1.ListOptions) (result *v1alpha1.AwsCloudwatchMetricAlarmList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awscloudwatchmetricalarmsResource, awscloudwatchmetricalarmsKind, opts), &v1alpha1.AwsCloudwatchMetricAlarmList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsCloudwatchMetricAlarmList{ListMeta: obj.(*v1alpha1.AwsCloudwatchMetricAlarmList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsCloudwatchMetricAlarmList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsCloudwatchMetricAlarms.
func (c *FakeAwsCloudwatchMetricAlarms) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awscloudwatchmetricalarmsResource, opts))
}

// Create takes the representation of a awsCloudwatchMetricAlarm and creates it.  Returns the server's representation of the awsCloudwatchMetricAlarm, and an error, if there is any.
func (c *FakeAwsCloudwatchMetricAlarms) Create(awsCloudwatchMetricAlarm *v1alpha1.AwsCloudwatchMetricAlarm) (result *v1alpha1.AwsCloudwatchMetricAlarm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awscloudwatchmetricalarmsResource, awsCloudwatchMetricAlarm), &v1alpha1.AwsCloudwatchMetricAlarm{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCloudwatchMetricAlarm), err
}

// Update takes the representation of a awsCloudwatchMetricAlarm and updates it. Returns the server's representation of the awsCloudwatchMetricAlarm, and an error, if there is any.
func (c *FakeAwsCloudwatchMetricAlarms) Update(awsCloudwatchMetricAlarm *v1alpha1.AwsCloudwatchMetricAlarm) (result *v1alpha1.AwsCloudwatchMetricAlarm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awscloudwatchmetricalarmsResource, awsCloudwatchMetricAlarm), &v1alpha1.AwsCloudwatchMetricAlarm{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCloudwatchMetricAlarm), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsCloudwatchMetricAlarms) UpdateStatus(awsCloudwatchMetricAlarm *v1alpha1.AwsCloudwatchMetricAlarm) (*v1alpha1.AwsCloudwatchMetricAlarm, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awscloudwatchmetricalarmsResource, "status", awsCloudwatchMetricAlarm), &v1alpha1.AwsCloudwatchMetricAlarm{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCloudwatchMetricAlarm), err
}

// Delete takes name of the awsCloudwatchMetricAlarm and deletes it. Returns an error if one occurs.
func (c *FakeAwsCloudwatchMetricAlarms) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awscloudwatchmetricalarmsResource, name), &v1alpha1.AwsCloudwatchMetricAlarm{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsCloudwatchMetricAlarms) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awscloudwatchmetricalarmsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsCloudwatchMetricAlarmList{})
	return err
}

// Patch applies the patch and returns the patched awsCloudwatchMetricAlarm.
func (c *FakeAwsCloudwatchMetricAlarms) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCloudwatchMetricAlarm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awscloudwatchmetricalarmsResource, name, pt, data, subresources...), &v1alpha1.AwsCloudwatchMetricAlarm{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCloudwatchMetricAlarm), err
}
