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

// FakeAwsOpsworksRdsDbInstances implements AwsOpsworksRdsDbInstanceInterface
type FakeAwsOpsworksRdsDbInstances struct {
	Fake *FakeAwsV1alpha1
}

var awsopsworksrdsdbinstancesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsopsworksrdsdbinstances"}

var awsopsworksrdsdbinstancesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsOpsworksRdsDbInstance"}

// Get takes name of the awsOpsworksRdsDbInstance, and returns the corresponding awsOpsworksRdsDbInstance object, and an error if there is any.
func (c *FakeAwsOpsworksRdsDbInstances) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsOpsworksRdsDbInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsopsworksrdsdbinstancesResource, name), &v1alpha1.AwsOpsworksRdsDbInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOpsworksRdsDbInstance), err
}

// List takes label and field selectors, and returns the list of AwsOpsworksRdsDbInstances that match those selectors.
func (c *FakeAwsOpsworksRdsDbInstances) List(opts v1.ListOptions) (result *v1alpha1.AwsOpsworksRdsDbInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsopsworksrdsdbinstancesResource, awsopsworksrdsdbinstancesKind, opts), &v1alpha1.AwsOpsworksRdsDbInstanceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsOpsworksRdsDbInstanceList{ListMeta: obj.(*v1alpha1.AwsOpsworksRdsDbInstanceList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsOpsworksRdsDbInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsOpsworksRdsDbInstances.
func (c *FakeAwsOpsworksRdsDbInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsopsworksrdsdbinstancesResource, opts))
}

// Create takes the representation of a awsOpsworksRdsDbInstance and creates it.  Returns the server's representation of the awsOpsworksRdsDbInstance, and an error, if there is any.
func (c *FakeAwsOpsworksRdsDbInstances) Create(awsOpsworksRdsDbInstance *v1alpha1.AwsOpsworksRdsDbInstance) (result *v1alpha1.AwsOpsworksRdsDbInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsopsworksrdsdbinstancesResource, awsOpsworksRdsDbInstance), &v1alpha1.AwsOpsworksRdsDbInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOpsworksRdsDbInstance), err
}

// Update takes the representation of a awsOpsworksRdsDbInstance and updates it. Returns the server's representation of the awsOpsworksRdsDbInstance, and an error, if there is any.
func (c *FakeAwsOpsworksRdsDbInstances) Update(awsOpsworksRdsDbInstance *v1alpha1.AwsOpsworksRdsDbInstance) (result *v1alpha1.AwsOpsworksRdsDbInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsopsworksrdsdbinstancesResource, awsOpsworksRdsDbInstance), &v1alpha1.AwsOpsworksRdsDbInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOpsworksRdsDbInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsOpsworksRdsDbInstances) UpdateStatus(awsOpsworksRdsDbInstance *v1alpha1.AwsOpsworksRdsDbInstance) (*v1alpha1.AwsOpsworksRdsDbInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsopsworksrdsdbinstancesResource, "status", awsOpsworksRdsDbInstance), &v1alpha1.AwsOpsworksRdsDbInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOpsworksRdsDbInstance), err
}

// Delete takes name of the awsOpsworksRdsDbInstance and deletes it. Returns an error if one occurs.
func (c *FakeAwsOpsworksRdsDbInstances) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsopsworksrdsdbinstancesResource, name), &v1alpha1.AwsOpsworksRdsDbInstance{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsOpsworksRdsDbInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsopsworksrdsdbinstancesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsOpsworksRdsDbInstanceList{})
	return err
}

// Patch applies the patch and returns the patched awsOpsworksRdsDbInstance.
func (c *FakeAwsOpsworksRdsDbInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsOpsworksRdsDbInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsopsworksrdsdbinstancesResource, name, pt, data, subresources...), &v1alpha1.AwsOpsworksRdsDbInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOpsworksRdsDbInstance), err
}
