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

// FakeDaxParameterGroups implements DaxParameterGroupInterface
type FakeDaxParameterGroups struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var daxparametergroupsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "daxparametergroups"}

var daxparametergroupsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DaxParameterGroup"}

// Get takes name of the daxParameterGroup, and returns the corresponding daxParameterGroup object, and an error if there is any.
func (c *FakeDaxParameterGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.DaxParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(daxparametergroupsResource, c.ns, name), &v1alpha1.DaxParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DaxParameterGroup), err
}

// List takes label and field selectors, and returns the list of DaxParameterGroups that match those selectors.
func (c *FakeDaxParameterGroups) List(opts v1.ListOptions) (result *v1alpha1.DaxParameterGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(daxparametergroupsResource, daxparametergroupsKind, c.ns, opts), &v1alpha1.DaxParameterGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DaxParameterGroupList{ListMeta: obj.(*v1alpha1.DaxParameterGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.DaxParameterGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested daxParameterGroups.
func (c *FakeDaxParameterGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(daxparametergroupsResource, c.ns, opts))

}

// Create takes the representation of a daxParameterGroup and creates it.  Returns the server's representation of the daxParameterGroup, and an error, if there is any.
func (c *FakeDaxParameterGroups) Create(daxParameterGroup *v1alpha1.DaxParameterGroup) (result *v1alpha1.DaxParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(daxparametergroupsResource, c.ns, daxParameterGroup), &v1alpha1.DaxParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DaxParameterGroup), err
}

// Update takes the representation of a daxParameterGroup and updates it. Returns the server's representation of the daxParameterGroup, and an error, if there is any.
func (c *FakeDaxParameterGroups) Update(daxParameterGroup *v1alpha1.DaxParameterGroup) (result *v1alpha1.DaxParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(daxparametergroupsResource, c.ns, daxParameterGroup), &v1alpha1.DaxParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DaxParameterGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDaxParameterGroups) UpdateStatus(daxParameterGroup *v1alpha1.DaxParameterGroup) (*v1alpha1.DaxParameterGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(daxparametergroupsResource, "status", c.ns, daxParameterGroup), &v1alpha1.DaxParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DaxParameterGroup), err
}

// Delete takes name of the daxParameterGroup and deletes it. Returns an error if one occurs.
func (c *FakeDaxParameterGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(daxparametergroupsResource, c.ns, name), &v1alpha1.DaxParameterGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDaxParameterGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(daxparametergroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DaxParameterGroupList{})
	return err
}

// Patch applies the patch and returns the patched daxParameterGroup.
func (c *FakeDaxParameterGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DaxParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(daxparametergroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DaxParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DaxParameterGroup), err
}