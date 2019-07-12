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

// FakeAwsSqsQueues implements AwsSqsQueueInterface
type FakeAwsSqsQueues struct {
	Fake *FakeAwsV1alpha1
}

var awssqsqueuesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awssqsqueues"}

var awssqsqueuesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsSqsQueue"}

// Get takes name of the awsSqsQueue, and returns the corresponding awsSqsQueue object, and an error if there is any.
func (c *FakeAwsSqsQueues) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSqsQueue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awssqsqueuesResource, name), &v1alpha1.AwsSqsQueue{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSqsQueue), err
}

// List takes label and field selectors, and returns the list of AwsSqsQueues that match those selectors.
func (c *FakeAwsSqsQueues) List(opts v1.ListOptions) (result *v1alpha1.AwsSqsQueueList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awssqsqueuesResource, awssqsqueuesKind, opts), &v1alpha1.AwsSqsQueueList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsSqsQueueList{ListMeta: obj.(*v1alpha1.AwsSqsQueueList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsSqsQueueList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsSqsQueues.
func (c *FakeAwsSqsQueues) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awssqsqueuesResource, opts))
}

// Create takes the representation of a awsSqsQueue and creates it.  Returns the server's representation of the awsSqsQueue, and an error, if there is any.
func (c *FakeAwsSqsQueues) Create(awsSqsQueue *v1alpha1.AwsSqsQueue) (result *v1alpha1.AwsSqsQueue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awssqsqueuesResource, awsSqsQueue), &v1alpha1.AwsSqsQueue{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSqsQueue), err
}

// Update takes the representation of a awsSqsQueue and updates it. Returns the server's representation of the awsSqsQueue, and an error, if there is any.
func (c *FakeAwsSqsQueues) Update(awsSqsQueue *v1alpha1.AwsSqsQueue) (result *v1alpha1.AwsSqsQueue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awssqsqueuesResource, awsSqsQueue), &v1alpha1.AwsSqsQueue{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSqsQueue), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsSqsQueues) UpdateStatus(awsSqsQueue *v1alpha1.AwsSqsQueue) (*v1alpha1.AwsSqsQueue, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awssqsqueuesResource, "status", awsSqsQueue), &v1alpha1.AwsSqsQueue{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSqsQueue), err
}

// Delete takes name of the awsSqsQueue and deletes it. Returns an error if one occurs.
func (c *FakeAwsSqsQueues) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awssqsqueuesResource, name), &v1alpha1.AwsSqsQueue{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsSqsQueues) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awssqsqueuesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsSqsQueueList{})
	return err
}

// Patch applies the patch and returns the patched awsSqsQueue.
func (c *FakeAwsSqsQueues) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSqsQueue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awssqsqueuesResource, name, pt, data, subresources...), &v1alpha1.AwsSqsQueue{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSqsQueue), err
}
