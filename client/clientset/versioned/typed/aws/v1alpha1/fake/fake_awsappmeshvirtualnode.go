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

// FakeAwsAppmeshVirtualNodes implements AwsAppmeshVirtualNodeInterface
type FakeAwsAppmeshVirtualNodes struct {
	Fake *FakeAwsV1alpha1
}

var awsappmeshvirtualnodesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsappmeshvirtualnodes"}

var awsappmeshvirtualnodesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsAppmeshVirtualNode"}

// Get takes name of the awsAppmeshVirtualNode, and returns the corresponding awsAppmeshVirtualNode object, and an error if there is any.
func (c *FakeAwsAppmeshVirtualNodes) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsAppmeshVirtualNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsappmeshvirtualnodesResource, name), &v1alpha1.AwsAppmeshVirtualNode{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAppmeshVirtualNode), err
}

// List takes label and field selectors, and returns the list of AwsAppmeshVirtualNodes that match those selectors.
func (c *FakeAwsAppmeshVirtualNodes) List(opts v1.ListOptions) (result *v1alpha1.AwsAppmeshVirtualNodeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsappmeshvirtualnodesResource, awsappmeshvirtualnodesKind, opts), &v1alpha1.AwsAppmeshVirtualNodeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsAppmeshVirtualNodeList{ListMeta: obj.(*v1alpha1.AwsAppmeshVirtualNodeList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsAppmeshVirtualNodeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsAppmeshVirtualNodes.
func (c *FakeAwsAppmeshVirtualNodes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsappmeshvirtualnodesResource, opts))
}

// Create takes the representation of a awsAppmeshVirtualNode and creates it.  Returns the server's representation of the awsAppmeshVirtualNode, and an error, if there is any.
func (c *FakeAwsAppmeshVirtualNodes) Create(awsAppmeshVirtualNode *v1alpha1.AwsAppmeshVirtualNode) (result *v1alpha1.AwsAppmeshVirtualNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsappmeshvirtualnodesResource, awsAppmeshVirtualNode), &v1alpha1.AwsAppmeshVirtualNode{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAppmeshVirtualNode), err
}

// Update takes the representation of a awsAppmeshVirtualNode and updates it. Returns the server's representation of the awsAppmeshVirtualNode, and an error, if there is any.
func (c *FakeAwsAppmeshVirtualNodes) Update(awsAppmeshVirtualNode *v1alpha1.AwsAppmeshVirtualNode) (result *v1alpha1.AwsAppmeshVirtualNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsappmeshvirtualnodesResource, awsAppmeshVirtualNode), &v1alpha1.AwsAppmeshVirtualNode{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAppmeshVirtualNode), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsAppmeshVirtualNodes) UpdateStatus(awsAppmeshVirtualNode *v1alpha1.AwsAppmeshVirtualNode) (*v1alpha1.AwsAppmeshVirtualNode, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsappmeshvirtualnodesResource, "status", awsAppmeshVirtualNode), &v1alpha1.AwsAppmeshVirtualNode{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAppmeshVirtualNode), err
}

// Delete takes name of the awsAppmeshVirtualNode and deletes it. Returns an error if one occurs.
func (c *FakeAwsAppmeshVirtualNodes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsappmeshvirtualnodesResource, name), &v1alpha1.AwsAppmeshVirtualNode{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsAppmeshVirtualNodes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsappmeshvirtualnodesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsAppmeshVirtualNodeList{})
	return err
}

// Patch applies the patch and returns the patched awsAppmeshVirtualNode.
func (c *FakeAwsAppmeshVirtualNodes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAppmeshVirtualNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsappmeshvirtualnodesResource, name, pt, data, subresources...), &v1alpha1.AwsAppmeshVirtualNode{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAppmeshVirtualNode), err
}
