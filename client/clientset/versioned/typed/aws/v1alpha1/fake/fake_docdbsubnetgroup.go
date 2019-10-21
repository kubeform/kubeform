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

// FakeDocdbSubnetGroups implements DocdbSubnetGroupInterface
type FakeDocdbSubnetGroups struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var docdbsubnetgroupsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "docdbsubnetgroups"}

var docdbsubnetgroupsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DocdbSubnetGroup"}

// Get takes name of the docdbSubnetGroup, and returns the corresponding docdbSubnetGroup object, and an error if there is any.
func (c *FakeDocdbSubnetGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.DocdbSubnetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(docdbsubnetgroupsResource, c.ns, name), &v1alpha1.DocdbSubnetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DocdbSubnetGroup), err
}

// List takes label and field selectors, and returns the list of DocdbSubnetGroups that match those selectors.
func (c *FakeDocdbSubnetGroups) List(opts v1.ListOptions) (result *v1alpha1.DocdbSubnetGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(docdbsubnetgroupsResource, docdbsubnetgroupsKind, c.ns, opts), &v1alpha1.DocdbSubnetGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DocdbSubnetGroupList{ListMeta: obj.(*v1alpha1.DocdbSubnetGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.DocdbSubnetGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested docdbSubnetGroups.
func (c *FakeDocdbSubnetGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(docdbsubnetgroupsResource, c.ns, opts))

}

// Create takes the representation of a docdbSubnetGroup and creates it.  Returns the server's representation of the docdbSubnetGroup, and an error, if there is any.
func (c *FakeDocdbSubnetGroups) Create(docdbSubnetGroup *v1alpha1.DocdbSubnetGroup) (result *v1alpha1.DocdbSubnetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(docdbsubnetgroupsResource, c.ns, docdbSubnetGroup), &v1alpha1.DocdbSubnetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DocdbSubnetGroup), err
}

// Update takes the representation of a docdbSubnetGroup and updates it. Returns the server's representation of the docdbSubnetGroup, and an error, if there is any.
func (c *FakeDocdbSubnetGroups) Update(docdbSubnetGroup *v1alpha1.DocdbSubnetGroup) (result *v1alpha1.DocdbSubnetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(docdbsubnetgroupsResource, c.ns, docdbSubnetGroup), &v1alpha1.DocdbSubnetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DocdbSubnetGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDocdbSubnetGroups) UpdateStatus(docdbSubnetGroup *v1alpha1.DocdbSubnetGroup) (*v1alpha1.DocdbSubnetGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(docdbsubnetgroupsResource, "status", c.ns, docdbSubnetGroup), &v1alpha1.DocdbSubnetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DocdbSubnetGroup), err
}

// Delete takes name of the docdbSubnetGroup and deletes it. Returns an error if one occurs.
func (c *FakeDocdbSubnetGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(docdbsubnetgroupsResource, c.ns, name), &v1alpha1.DocdbSubnetGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDocdbSubnetGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(docdbsubnetgroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DocdbSubnetGroupList{})
	return err
}

// Patch applies the patch and returns the patched docdbSubnetGroup.
func (c *FakeDocdbSubnetGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DocdbSubnetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(docdbsubnetgroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DocdbSubnetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DocdbSubnetGroup), err
}