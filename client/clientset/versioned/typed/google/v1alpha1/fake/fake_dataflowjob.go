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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeDataflowJobs implements DataflowJobInterface
type FakeDataflowJobs struct {
	Fake *FakeGoogleV1alpha1
}

var dataflowjobsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "dataflowjobs"}

var dataflowjobsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "DataflowJob"}

// Get takes name of the dataflowJob, and returns the corresponding dataflowJob object, and an error if there is any.
func (c *FakeDataflowJobs) Get(name string, options v1.GetOptions) (result *v1alpha1.DataflowJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(dataflowjobsResource, name), &v1alpha1.DataflowJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataflowJob), err
}

// List takes label and field selectors, and returns the list of DataflowJobs that match those selectors.
func (c *FakeDataflowJobs) List(opts v1.ListOptions) (result *v1alpha1.DataflowJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(dataflowjobsResource, dataflowjobsKind, opts), &v1alpha1.DataflowJobList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DataflowJobList{ListMeta: obj.(*v1alpha1.DataflowJobList).ListMeta}
	for _, item := range obj.(*v1alpha1.DataflowJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dataflowJobs.
func (c *FakeDataflowJobs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(dataflowjobsResource, opts))
}

// Create takes the representation of a dataflowJob and creates it.  Returns the server's representation of the dataflowJob, and an error, if there is any.
func (c *FakeDataflowJobs) Create(dataflowJob *v1alpha1.DataflowJob) (result *v1alpha1.DataflowJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(dataflowjobsResource, dataflowJob), &v1alpha1.DataflowJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataflowJob), err
}

// Update takes the representation of a dataflowJob and updates it. Returns the server's representation of the dataflowJob, and an error, if there is any.
func (c *FakeDataflowJobs) Update(dataflowJob *v1alpha1.DataflowJob) (result *v1alpha1.DataflowJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(dataflowjobsResource, dataflowJob), &v1alpha1.DataflowJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataflowJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDataflowJobs) UpdateStatus(dataflowJob *v1alpha1.DataflowJob) (*v1alpha1.DataflowJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(dataflowjobsResource, "status", dataflowJob), &v1alpha1.DataflowJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataflowJob), err
}

// Delete takes name of the dataflowJob and deletes it. Returns an error if one occurs.
func (c *FakeDataflowJobs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(dataflowjobsResource, name), &v1alpha1.DataflowJob{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDataflowJobs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(dataflowjobsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DataflowJobList{})
	return err
}

// Patch applies the patch and returns the patched dataflowJob.
func (c *FakeDataflowJobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DataflowJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(dataflowjobsResource, name, pt, data, subresources...), &v1alpha1.DataflowJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataflowJob), err
}
