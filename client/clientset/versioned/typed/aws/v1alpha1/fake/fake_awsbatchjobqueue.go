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

// FakeAwsBatchJobQueues implements AwsBatchJobQueueInterface
type FakeAwsBatchJobQueues struct {
	Fake *FakeAwsV1alpha1
}

var awsbatchjobqueuesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsbatchjobqueues"}

var awsbatchjobqueuesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsBatchJobQueue"}

// Get takes name of the awsBatchJobQueue, and returns the corresponding awsBatchJobQueue object, and an error if there is any.
func (c *FakeAwsBatchJobQueues) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsBatchJobQueue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsbatchjobqueuesResource, name), &v1alpha1.AwsBatchJobQueue{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsBatchJobQueue), err
}

// List takes label and field selectors, and returns the list of AwsBatchJobQueues that match those selectors.
func (c *FakeAwsBatchJobQueues) List(opts v1.ListOptions) (result *v1alpha1.AwsBatchJobQueueList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsbatchjobqueuesResource, awsbatchjobqueuesKind, opts), &v1alpha1.AwsBatchJobQueueList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsBatchJobQueueList{ListMeta: obj.(*v1alpha1.AwsBatchJobQueueList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsBatchJobQueueList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsBatchJobQueues.
func (c *FakeAwsBatchJobQueues) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsbatchjobqueuesResource, opts))
}

// Create takes the representation of a awsBatchJobQueue and creates it.  Returns the server's representation of the awsBatchJobQueue, and an error, if there is any.
func (c *FakeAwsBatchJobQueues) Create(awsBatchJobQueue *v1alpha1.AwsBatchJobQueue) (result *v1alpha1.AwsBatchJobQueue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsbatchjobqueuesResource, awsBatchJobQueue), &v1alpha1.AwsBatchJobQueue{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsBatchJobQueue), err
}

// Update takes the representation of a awsBatchJobQueue and updates it. Returns the server's representation of the awsBatchJobQueue, and an error, if there is any.
func (c *FakeAwsBatchJobQueues) Update(awsBatchJobQueue *v1alpha1.AwsBatchJobQueue) (result *v1alpha1.AwsBatchJobQueue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsbatchjobqueuesResource, awsBatchJobQueue), &v1alpha1.AwsBatchJobQueue{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsBatchJobQueue), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsBatchJobQueues) UpdateStatus(awsBatchJobQueue *v1alpha1.AwsBatchJobQueue) (*v1alpha1.AwsBatchJobQueue, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsbatchjobqueuesResource, "status", awsBatchJobQueue), &v1alpha1.AwsBatchJobQueue{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsBatchJobQueue), err
}

// Delete takes name of the awsBatchJobQueue and deletes it. Returns an error if one occurs.
func (c *FakeAwsBatchJobQueues) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsbatchjobqueuesResource, name), &v1alpha1.AwsBatchJobQueue{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsBatchJobQueues) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsbatchjobqueuesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsBatchJobQueueList{})
	return err
}

// Patch applies the patch and returns the patched awsBatchJobQueue.
func (c *FakeAwsBatchJobQueues) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsBatchJobQueue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsbatchjobqueuesResource, name, pt, data, subresources...), &v1alpha1.AwsBatchJobQueue{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsBatchJobQueue), err
}
