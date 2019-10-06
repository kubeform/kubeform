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

// FakeCodedeployDeploymentGroups implements CodedeployDeploymentGroupInterface
type FakeCodedeployDeploymentGroups struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var codedeploydeploymentgroupsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "codedeploydeploymentgroups"}

var codedeploydeploymentgroupsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "CodedeployDeploymentGroup"}

// Get takes name of the codedeployDeploymentGroup, and returns the corresponding codedeployDeploymentGroup object, and an error if there is any.
func (c *FakeCodedeployDeploymentGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.CodedeployDeploymentGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(codedeploydeploymentgroupsResource, c.ns, name), &v1alpha1.CodedeployDeploymentGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CodedeployDeploymentGroup), err
}

// List takes label and field selectors, and returns the list of CodedeployDeploymentGroups that match those selectors.
func (c *FakeCodedeployDeploymentGroups) List(opts v1.ListOptions) (result *v1alpha1.CodedeployDeploymentGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(codedeploydeploymentgroupsResource, codedeploydeploymentgroupsKind, c.ns, opts), &v1alpha1.CodedeployDeploymentGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CodedeployDeploymentGroupList{ListMeta: obj.(*v1alpha1.CodedeployDeploymentGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.CodedeployDeploymentGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested codedeployDeploymentGroups.
func (c *FakeCodedeployDeploymentGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(codedeploydeploymentgroupsResource, c.ns, opts))

}

// Create takes the representation of a codedeployDeploymentGroup and creates it.  Returns the server's representation of the codedeployDeploymentGroup, and an error, if there is any.
func (c *FakeCodedeployDeploymentGroups) Create(codedeployDeploymentGroup *v1alpha1.CodedeployDeploymentGroup) (result *v1alpha1.CodedeployDeploymentGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(codedeploydeploymentgroupsResource, c.ns, codedeployDeploymentGroup), &v1alpha1.CodedeployDeploymentGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CodedeployDeploymentGroup), err
}

// Update takes the representation of a codedeployDeploymentGroup and updates it. Returns the server's representation of the codedeployDeploymentGroup, and an error, if there is any.
func (c *FakeCodedeployDeploymentGroups) Update(codedeployDeploymentGroup *v1alpha1.CodedeployDeploymentGroup) (result *v1alpha1.CodedeployDeploymentGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(codedeploydeploymentgroupsResource, c.ns, codedeployDeploymentGroup), &v1alpha1.CodedeployDeploymentGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CodedeployDeploymentGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCodedeployDeploymentGroups) UpdateStatus(codedeployDeploymentGroup *v1alpha1.CodedeployDeploymentGroup) (*v1alpha1.CodedeployDeploymentGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(codedeploydeploymentgroupsResource, "status", c.ns, codedeployDeploymentGroup), &v1alpha1.CodedeployDeploymentGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CodedeployDeploymentGroup), err
}

// Delete takes name of the codedeployDeploymentGroup and deletes it. Returns an error if one occurs.
func (c *FakeCodedeployDeploymentGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(codedeploydeploymentgroupsResource, c.ns, name), &v1alpha1.CodedeployDeploymentGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCodedeployDeploymentGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(codedeploydeploymentgroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CodedeployDeploymentGroupList{})
	return err
}

// Patch applies the patch and returns the patched codedeployDeploymentGroup.
func (c *FakeCodedeployDeploymentGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CodedeployDeploymentGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(codedeploydeploymentgroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.CodedeployDeploymentGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CodedeployDeploymentGroup), err
}
