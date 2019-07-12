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

// FakeAwsCloudhsmV2Hsms implements AwsCloudhsmV2HsmInterface
type FakeAwsCloudhsmV2Hsms struct {
	Fake *FakeAwsV1alpha1
}

var awscloudhsmv2hsmsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awscloudhsmv2hsms"}

var awscloudhsmv2hsmsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsCloudhsmV2Hsm"}

// Get takes name of the awsCloudhsmV2Hsm, and returns the corresponding awsCloudhsmV2Hsm object, and an error if there is any.
func (c *FakeAwsCloudhsmV2Hsms) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsCloudhsmV2Hsm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awscloudhsmv2hsmsResource, name), &v1alpha1.AwsCloudhsmV2Hsm{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCloudhsmV2Hsm), err
}

// List takes label and field selectors, and returns the list of AwsCloudhsmV2Hsms that match those selectors.
func (c *FakeAwsCloudhsmV2Hsms) List(opts v1.ListOptions) (result *v1alpha1.AwsCloudhsmV2HsmList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awscloudhsmv2hsmsResource, awscloudhsmv2hsmsKind, opts), &v1alpha1.AwsCloudhsmV2HsmList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsCloudhsmV2HsmList{ListMeta: obj.(*v1alpha1.AwsCloudhsmV2HsmList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsCloudhsmV2HsmList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsCloudhsmV2Hsms.
func (c *FakeAwsCloudhsmV2Hsms) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awscloudhsmv2hsmsResource, opts))
}

// Create takes the representation of a awsCloudhsmV2Hsm and creates it.  Returns the server's representation of the awsCloudhsmV2Hsm, and an error, if there is any.
func (c *FakeAwsCloudhsmV2Hsms) Create(awsCloudhsmV2Hsm *v1alpha1.AwsCloudhsmV2Hsm) (result *v1alpha1.AwsCloudhsmV2Hsm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awscloudhsmv2hsmsResource, awsCloudhsmV2Hsm), &v1alpha1.AwsCloudhsmV2Hsm{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCloudhsmV2Hsm), err
}

// Update takes the representation of a awsCloudhsmV2Hsm and updates it. Returns the server's representation of the awsCloudhsmV2Hsm, and an error, if there is any.
func (c *FakeAwsCloudhsmV2Hsms) Update(awsCloudhsmV2Hsm *v1alpha1.AwsCloudhsmV2Hsm) (result *v1alpha1.AwsCloudhsmV2Hsm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awscloudhsmv2hsmsResource, awsCloudhsmV2Hsm), &v1alpha1.AwsCloudhsmV2Hsm{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCloudhsmV2Hsm), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsCloudhsmV2Hsms) UpdateStatus(awsCloudhsmV2Hsm *v1alpha1.AwsCloudhsmV2Hsm) (*v1alpha1.AwsCloudhsmV2Hsm, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awscloudhsmv2hsmsResource, "status", awsCloudhsmV2Hsm), &v1alpha1.AwsCloudhsmV2Hsm{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCloudhsmV2Hsm), err
}

// Delete takes name of the awsCloudhsmV2Hsm and deletes it. Returns an error if one occurs.
func (c *FakeAwsCloudhsmV2Hsms) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awscloudhsmv2hsmsResource, name), &v1alpha1.AwsCloudhsmV2Hsm{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsCloudhsmV2Hsms) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awscloudhsmv2hsmsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsCloudhsmV2HsmList{})
	return err
}

// Patch applies the patch and returns the patched awsCloudhsmV2Hsm.
func (c *FakeAwsCloudhsmV2Hsms) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCloudhsmV2Hsm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awscloudhsmv2hsmsResource, name, pt, data, subresources...), &v1alpha1.AwsCloudhsmV2Hsm{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCloudhsmV2Hsm), err
}
