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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeComputeBackendBuckets implements ComputeBackendBucketInterface
type FakeComputeBackendBuckets struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var computebackendbucketsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "computebackendbuckets"}

var computebackendbucketsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "ComputeBackendBucket"}

// Get takes name of the computeBackendBucket, and returns the corresponding computeBackendBucket object, and an error if there is any.
func (c *FakeComputeBackendBuckets) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeBackendBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computebackendbucketsResource, c.ns, name), &v1alpha1.ComputeBackendBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeBackendBucket), err
}

// List takes label and field selectors, and returns the list of ComputeBackendBuckets that match those selectors.
func (c *FakeComputeBackendBuckets) List(opts v1.ListOptions) (result *v1alpha1.ComputeBackendBucketList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computebackendbucketsResource, computebackendbucketsKind, c.ns, opts), &v1alpha1.ComputeBackendBucketList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComputeBackendBucketList{ListMeta: obj.(*v1alpha1.ComputeBackendBucketList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComputeBackendBucketList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeBackendBuckets.
func (c *FakeComputeBackendBuckets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computebackendbucketsResource, c.ns, opts))

}

// Create takes the representation of a computeBackendBucket and creates it.  Returns the server's representation of the computeBackendBucket, and an error, if there is any.
func (c *FakeComputeBackendBuckets) Create(computeBackendBucket *v1alpha1.ComputeBackendBucket) (result *v1alpha1.ComputeBackendBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computebackendbucketsResource, c.ns, computeBackendBucket), &v1alpha1.ComputeBackendBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeBackendBucket), err
}

// Update takes the representation of a computeBackendBucket and updates it. Returns the server's representation of the computeBackendBucket, and an error, if there is any.
func (c *FakeComputeBackendBuckets) Update(computeBackendBucket *v1alpha1.ComputeBackendBucket) (result *v1alpha1.ComputeBackendBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computebackendbucketsResource, c.ns, computeBackendBucket), &v1alpha1.ComputeBackendBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeBackendBucket), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeBackendBuckets) UpdateStatus(computeBackendBucket *v1alpha1.ComputeBackendBucket) (*v1alpha1.ComputeBackendBucket, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computebackendbucketsResource, "status", c.ns, computeBackendBucket), &v1alpha1.ComputeBackendBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeBackendBucket), err
}

// Delete takes name of the computeBackendBucket and deletes it. Returns an error if one occurs.
func (c *FakeComputeBackendBuckets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(computebackendbucketsResource, c.ns, name), &v1alpha1.ComputeBackendBucket{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeBackendBuckets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computebackendbucketsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComputeBackendBucketList{})
	return err
}

// Patch applies the patch and returns the patched computeBackendBucket.
func (c *FakeComputeBackendBuckets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeBackendBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computebackendbucketsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ComputeBackendBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeBackendBucket), err
}
