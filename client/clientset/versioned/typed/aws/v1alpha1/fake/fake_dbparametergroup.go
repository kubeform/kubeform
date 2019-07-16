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

// FakeDbParameterGroups implements DbParameterGroupInterface
type FakeDbParameterGroups struct {
	Fake *FakeAwsV1alpha1
}

var dbparametergroupsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "dbparametergroups"}

var dbparametergroupsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DbParameterGroup"}

// Get takes name of the dbParameterGroup, and returns the corresponding dbParameterGroup object, and an error if there is any.
func (c *FakeDbParameterGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.DbParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(dbparametergroupsResource, name), &v1alpha1.DbParameterGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DbParameterGroup), err
}

// List takes label and field selectors, and returns the list of DbParameterGroups that match those selectors.
func (c *FakeDbParameterGroups) List(opts v1.ListOptions) (result *v1alpha1.DbParameterGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(dbparametergroupsResource, dbparametergroupsKind, opts), &v1alpha1.DbParameterGroupList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DbParameterGroupList{ListMeta: obj.(*v1alpha1.DbParameterGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.DbParameterGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dbParameterGroups.
func (c *FakeDbParameterGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(dbparametergroupsResource, opts))
}

// Create takes the representation of a dbParameterGroup and creates it.  Returns the server's representation of the dbParameterGroup, and an error, if there is any.
func (c *FakeDbParameterGroups) Create(dbParameterGroup *v1alpha1.DbParameterGroup) (result *v1alpha1.DbParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(dbparametergroupsResource, dbParameterGroup), &v1alpha1.DbParameterGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DbParameterGroup), err
}

// Update takes the representation of a dbParameterGroup and updates it. Returns the server's representation of the dbParameterGroup, and an error, if there is any.
func (c *FakeDbParameterGroups) Update(dbParameterGroup *v1alpha1.DbParameterGroup) (result *v1alpha1.DbParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(dbparametergroupsResource, dbParameterGroup), &v1alpha1.DbParameterGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DbParameterGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDbParameterGroups) UpdateStatus(dbParameterGroup *v1alpha1.DbParameterGroup) (*v1alpha1.DbParameterGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(dbparametergroupsResource, "status", dbParameterGroup), &v1alpha1.DbParameterGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DbParameterGroup), err
}

// Delete takes name of the dbParameterGroup and deletes it. Returns an error if one occurs.
func (c *FakeDbParameterGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(dbparametergroupsResource, name), &v1alpha1.DbParameterGroup{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDbParameterGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(dbparametergroupsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DbParameterGroupList{})
	return err
}

// Patch applies the patch and returns the patched dbParameterGroup.
func (c *FakeDbParameterGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DbParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(dbparametergroupsResource, name, pt, data, subresources...), &v1alpha1.DbParameterGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DbParameterGroup), err
}
