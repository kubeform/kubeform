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

// FakeAwsDocdbClusterParameterGroups implements AwsDocdbClusterParameterGroupInterface
type FakeAwsDocdbClusterParameterGroups struct {
	Fake *FakeAwsV1alpha1
}

var awsdocdbclusterparametergroupsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsdocdbclusterparametergroups"}

var awsdocdbclusterparametergroupsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsDocdbClusterParameterGroup"}

// Get takes name of the awsDocdbClusterParameterGroup, and returns the corresponding awsDocdbClusterParameterGroup object, and an error if there is any.
func (c *FakeAwsDocdbClusterParameterGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDocdbClusterParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsdocdbclusterparametergroupsResource, name), &v1alpha1.AwsDocdbClusterParameterGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDocdbClusterParameterGroup), err
}

// List takes label and field selectors, and returns the list of AwsDocdbClusterParameterGroups that match those selectors.
func (c *FakeAwsDocdbClusterParameterGroups) List(opts v1.ListOptions) (result *v1alpha1.AwsDocdbClusterParameterGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsdocdbclusterparametergroupsResource, awsdocdbclusterparametergroupsKind, opts), &v1alpha1.AwsDocdbClusterParameterGroupList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsDocdbClusterParameterGroupList{ListMeta: obj.(*v1alpha1.AwsDocdbClusterParameterGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsDocdbClusterParameterGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsDocdbClusterParameterGroups.
func (c *FakeAwsDocdbClusterParameterGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsdocdbclusterparametergroupsResource, opts))
}

// Create takes the representation of a awsDocdbClusterParameterGroup and creates it.  Returns the server's representation of the awsDocdbClusterParameterGroup, and an error, if there is any.
func (c *FakeAwsDocdbClusterParameterGroups) Create(awsDocdbClusterParameterGroup *v1alpha1.AwsDocdbClusterParameterGroup) (result *v1alpha1.AwsDocdbClusterParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsdocdbclusterparametergroupsResource, awsDocdbClusterParameterGroup), &v1alpha1.AwsDocdbClusterParameterGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDocdbClusterParameterGroup), err
}

// Update takes the representation of a awsDocdbClusterParameterGroup and updates it. Returns the server's representation of the awsDocdbClusterParameterGroup, and an error, if there is any.
func (c *FakeAwsDocdbClusterParameterGroups) Update(awsDocdbClusterParameterGroup *v1alpha1.AwsDocdbClusterParameterGroup) (result *v1alpha1.AwsDocdbClusterParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsdocdbclusterparametergroupsResource, awsDocdbClusterParameterGroup), &v1alpha1.AwsDocdbClusterParameterGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDocdbClusterParameterGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsDocdbClusterParameterGroups) UpdateStatus(awsDocdbClusterParameterGroup *v1alpha1.AwsDocdbClusterParameterGroup) (*v1alpha1.AwsDocdbClusterParameterGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsdocdbclusterparametergroupsResource, "status", awsDocdbClusterParameterGroup), &v1alpha1.AwsDocdbClusterParameterGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDocdbClusterParameterGroup), err
}

// Delete takes name of the awsDocdbClusterParameterGroup and deletes it. Returns an error if one occurs.
func (c *FakeAwsDocdbClusterParameterGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsdocdbclusterparametergroupsResource, name), &v1alpha1.AwsDocdbClusterParameterGroup{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsDocdbClusterParameterGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsdocdbclusterparametergroupsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsDocdbClusterParameterGroupList{})
	return err
}

// Patch applies the patch and returns the patched awsDocdbClusterParameterGroup.
func (c *FakeAwsDocdbClusterParameterGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDocdbClusterParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsdocdbclusterparametergroupsResource, name, pt, data, subresources...), &v1alpha1.AwsDocdbClusterParameterGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDocdbClusterParameterGroup), err
}
