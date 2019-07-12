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

// FakeAwsSagemakerNotebookInstances implements AwsSagemakerNotebookInstanceInterface
type FakeAwsSagemakerNotebookInstances struct {
	Fake *FakeAwsV1alpha1
}

var awssagemakernotebookinstancesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awssagemakernotebookinstances"}

var awssagemakernotebookinstancesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsSagemakerNotebookInstance"}

// Get takes name of the awsSagemakerNotebookInstance, and returns the corresponding awsSagemakerNotebookInstance object, and an error if there is any.
func (c *FakeAwsSagemakerNotebookInstances) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSagemakerNotebookInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awssagemakernotebookinstancesResource, name), &v1alpha1.AwsSagemakerNotebookInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSagemakerNotebookInstance), err
}

// List takes label and field selectors, and returns the list of AwsSagemakerNotebookInstances that match those selectors.
func (c *FakeAwsSagemakerNotebookInstances) List(opts v1.ListOptions) (result *v1alpha1.AwsSagemakerNotebookInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awssagemakernotebookinstancesResource, awssagemakernotebookinstancesKind, opts), &v1alpha1.AwsSagemakerNotebookInstanceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsSagemakerNotebookInstanceList{ListMeta: obj.(*v1alpha1.AwsSagemakerNotebookInstanceList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsSagemakerNotebookInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsSagemakerNotebookInstances.
func (c *FakeAwsSagemakerNotebookInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awssagemakernotebookinstancesResource, opts))
}

// Create takes the representation of a awsSagemakerNotebookInstance and creates it.  Returns the server's representation of the awsSagemakerNotebookInstance, and an error, if there is any.
func (c *FakeAwsSagemakerNotebookInstances) Create(awsSagemakerNotebookInstance *v1alpha1.AwsSagemakerNotebookInstance) (result *v1alpha1.AwsSagemakerNotebookInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awssagemakernotebookinstancesResource, awsSagemakerNotebookInstance), &v1alpha1.AwsSagemakerNotebookInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSagemakerNotebookInstance), err
}

// Update takes the representation of a awsSagemakerNotebookInstance and updates it. Returns the server's representation of the awsSagemakerNotebookInstance, and an error, if there is any.
func (c *FakeAwsSagemakerNotebookInstances) Update(awsSagemakerNotebookInstance *v1alpha1.AwsSagemakerNotebookInstance) (result *v1alpha1.AwsSagemakerNotebookInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awssagemakernotebookinstancesResource, awsSagemakerNotebookInstance), &v1alpha1.AwsSagemakerNotebookInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSagemakerNotebookInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsSagemakerNotebookInstances) UpdateStatus(awsSagemakerNotebookInstance *v1alpha1.AwsSagemakerNotebookInstance) (*v1alpha1.AwsSagemakerNotebookInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awssagemakernotebookinstancesResource, "status", awsSagemakerNotebookInstance), &v1alpha1.AwsSagemakerNotebookInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSagemakerNotebookInstance), err
}

// Delete takes name of the awsSagemakerNotebookInstance and deletes it. Returns an error if one occurs.
func (c *FakeAwsSagemakerNotebookInstances) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awssagemakernotebookinstancesResource, name), &v1alpha1.AwsSagemakerNotebookInstance{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsSagemakerNotebookInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awssagemakernotebookinstancesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsSagemakerNotebookInstanceList{})
	return err
}

// Patch applies the patch and returns the patched awsSagemakerNotebookInstance.
func (c *FakeAwsSagemakerNotebookInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSagemakerNotebookInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awssagemakernotebookinstancesResource, name, pt, data, subresources...), &v1alpha1.AwsSagemakerNotebookInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSagemakerNotebookInstance), err
}
