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

// FakeRedshiftParameterGroups implements RedshiftParameterGroupInterface
type FakeRedshiftParameterGroups struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var redshiftparametergroupsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "redshiftparametergroups"}

var redshiftparametergroupsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "RedshiftParameterGroup"}

// Get takes name of the redshiftParameterGroup, and returns the corresponding redshiftParameterGroup object, and an error if there is any.
func (c *FakeRedshiftParameterGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.RedshiftParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(redshiftparametergroupsResource, c.ns, name), &v1alpha1.RedshiftParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedshiftParameterGroup), err
}

// List takes label and field selectors, and returns the list of RedshiftParameterGroups that match those selectors.
func (c *FakeRedshiftParameterGroups) List(opts v1.ListOptions) (result *v1alpha1.RedshiftParameterGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(redshiftparametergroupsResource, redshiftparametergroupsKind, c.ns, opts), &v1alpha1.RedshiftParameterGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RedshiftParameterGroupList{ListMeta: obj.(*v1alpha1.RedshiftParameterGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.RedshiftParameterGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested redshiftParameterGroups.
func (c *FakeRedshiftParameterGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(redshiftparametergroupsResource, c.ns, opts))

}

// Create takes the representation of a redshiftParameterGroup and creates it.  Returns the server's representation of the redshiftParameterGroup, and an error, if there is any.
func (c *FakeRedshiftParameterGroups) Create(redshiftParameterGroup *v1alpha1.RedshiftParameterGroup) (result *v1alpha1.RedshiftParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(redshiftparametergroupsResource, c.ns, redshiftParameterGroup), &v1alpha1.RedshiftParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedshiftParameterGroup), err
}

// Update takes the representation of a redshiftParameterGroup and updates it. Returns the server's representation of the redshiftParameterGroup, and an error, if there is any.
func (c *FakeRedshiftParameterGroups) Update(redshiftParameterGroup *v1alpha1.RedshiftParameterGroup) (result *v1alpha1.RedshiftParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(redshiftparametergroupsResource, c.ns, redshiftParameterGroup), &v1alpha1.RedshiftParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedshiftParameterGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRedshiftParameterGroups) UpdateStatus(redshiftParameterGroup *v1alpha1.RedshiftParameterGroup) (*v1alpha1.RedshiftParameterGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(redshiftparametergroupsResource, "status", c.ns, redshiftParameterGroup), &v1alpha1.RedshiftParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedshiftParameterGroup), err
}

// Delete takes name of the redshiftParameterGroup and deletes it. Returns an error if one occurs.
func (c *FakeRedshiftParameterGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(redshiftparametergroupsResource, c.ns, name), &v1alpha1.RedshiftParameterGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRedshiftParameterGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(redshiftparametergroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.RedshiftParameterGroupList{})
	return err
}

// Patch applies the patch and returns the patched redshiftParameterGroup.
func (c *FakeRedshiftParameterGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RedshiftParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(redshiftparametergroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.RedshiftParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedshiftParameterGroup), err
}