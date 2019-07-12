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

// FakeAwsEbsVolumes implements AwsEbsVolumeInterface
type FakeAwsEbsVolumes struct {
	Fake *FakeAwsV1alpha1
}

var awsebsvolumesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsebsvolumes"}

var awsebsvolumesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsEbsVolume"}

// Get takes name of the awsEbsVolume, and returns the corresponding awsEbsVolume object, and an error if there is any.
func (c *FakeAwsEbsVolumes) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsEbsVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsebsvolumesResource, name), &v1alpha1.AwsEbsVolume{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEbsVolume), err
}

// List takes label and field selectors, and returns the list of AwsEbsVolumes that match those selectors.
func (c *FakeAwsEbsVolumes) List(opts v1.ListOptions) (result *v1alpha1.AwsEbsVolumeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsebsvolumesResource, awsebsvolumesKind, opts), &v1alpha1.AwsEbsVolumeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsEbsVolumeList{ListMeta: obj.(*v1alpha1.AwsEbsVolumeList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsEbsVolumeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsEbsVolumes.
func (c *FakeAwsEbsVolumes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsebsvolumesResource, opts))
}

// Create takes the representation of a awsEbsVolume and creates it.  Returns the server's representation of the awsEbsVolume, and an error, if there is any.
func (c *FakeAwsEbsVolumes) Create(awsEbsVolume *v1alpha1.AwsEbsVolume) (result *v1alpha1.AwsEbsVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsebsvolumesResource, awsEbsVolume), &v1alpha1.AwsEbsVolume{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEbsVolume), err
}

// Update takes the representation of a awsEbsVolume and updates it. Returns the server's representation of the awsEbsVolume, and an error, if there is any.
func (c *FakeAwsEbsVolumes) Update(awsEbsVolume *v1alpha1.AwsEbsVolume) (result *v1alpha1.AwsEbsVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsebsvolumesResource, awsEbsVolume), &v1alpha1.AwsEbsVolume{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEbsVolume), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsEbsVolumes) UpdateStatus(awsEbsVolume *v1alpha1.AwsEbsVolume) (*v1alpha1.AwsEbsVolume, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsebsvolumesResource, "status", awsEbsVolume), &v1alpha1.AwsEbsVolume{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEbsVolume), err
}

// Delete takes name of the awsEbsVolume and deletes it. Returns an error if one occurs.
func (c *FakeAwsEbsVolumes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsebsvolumesResource, name), &v1alpha1.AwsEbsVolume{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsEbsVolumes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsebsvolumesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsEbsVolumeList{})
	return err
}

// Patch applies the patch and returns the patched awsEbsVolume.
func (c *FakeAwsEbsVolumes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsEbsVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsebsvolumesResource, name, pt, data, subresources...), &v1alpha1.AwsEbsVolume{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEbsVolume), err
}
