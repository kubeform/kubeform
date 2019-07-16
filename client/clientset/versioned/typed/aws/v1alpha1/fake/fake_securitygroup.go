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

// FakeSecurityGroups implements SecurityGroupInterface
type FakeSecurityGroups struct {
	Fake *FakeAwsV1alpha1
}

var securitygroupsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "securitygroups"}

var securitygroupsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "SecurityGroup"}

// Get takes name of the securityGroup, and returns the corresponding securityGroup object, and an error if there is any.
func (c *FakeSecurityGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.SecurityGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(securitygroupsResource, name), &v1alpha1.SecurityGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecurityGroup), err
}

// List takes label and field selectors, and returns the list of SecurityGroups that match those selectors.
func (c *FakeSecurityGroups) List(opts v1.ListOptions) (result *v1alpha1.SecurityGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(securitygroupsResource, securitygroupsKind, opts), &v1alpha1.SecurityGroupList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SecurityGroupList{ListMeta: obj.(*v1alpha1.SecurityGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.SecurityGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested securityGroups.
func (c *FakeSecurityGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(securitygroupsResource, opts))
}

// Create takes the representation of a securityGroup and creates it.  Returns the server's representation of the securityGroup, and an error, if there is any.
func (c *FakeSecurityGroups) Create(securityGroup *v1alpha1.SecurityGroup) (result *v1alpha1.SecurityGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(securitygroupsResource, securityGroup), &v1alpha1.SecurityGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecurityGroup), err
}

// Update takes the representation of a securityGroup and updates it. Returns the server's representation of the securityGroup, and an error, if there is any.
func (c *FakeSecurityGroups) Update(securityGroup *v1alpha1.SecurityGroup) (result *v1alpha1.SecurityGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(securitygroupsResource, securityGroup), &v1alpha1.SecurityGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecurityGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSecurityGroups) UpdateStatus(securityGroup *v1alpha1.SecurityGroup) (*v1alpha1.SecurityGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(securitygroupsResource, "status", securityGroup), &v1alpha1.SecurityGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecurityGroup), err
}

// Delete takes name of the securityGroup and deletes it. Returns an error if one occurs.
func (c *FakeSecurityGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(securitygroupsResource, name), &v1alpha1.SecurityGroup{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSecurityGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(securitygroupsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SecurityGroupList{})
	return err
}

// Patch applies the patch and returns the patched securityGroup.
func (c *FakeSecurityGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SecurityGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(securitygroupsResource, name, pt, data, subresources...), &v1alpha1.SecurityGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecurityGroup), err
}
