/*
Copyright The Kubeform Authors.

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
	"context"

	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCloudSchedulerJobs implements CloudSchedulerJobInterface
type FakeCloudSchedulerJobs struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var cloudschedulerjobsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "cloudschedulerjobs"}

var cloudschedulerjobsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "CloudSchedulerJob"}

// Get takes name of the cloudSchedulerJob, and returns the corresponding cloudSchedulerJob object, and an error if there is any.
func (c *FakeCloudSchedulerJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CloudSchedulerJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cloudschedulerjobsResource, c.ns, name), &v1alpha1.CloudSchedulerJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudSchedulerJob), err
}

// List takes label and field selectors, and returns the list of CloudSchedulerJobs that match those selectors.
func (c *FakeCloudSchedulerJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CloudSchedulerJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cloudschedulerjobsResource, cloudschedulerjobsKind, c.ns, opts), &v1alpha1.CloudSchedulerJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CloudSchedulerJobList{ListMeta: obj.(*v1alpha1.CloudSchedulerJobList).ListMeta}
	for _, item := range obj.(*v1alpha1.CloudSchedulerJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudSchedulerJobs.
func (c *FakeCloudSchedulerJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cloudschedulerjobsResource, c.ns, opts))

}

// Create takes the representation of a cloudSchedulerJob and creates it.  Returns the server's representation of the cloudSchedulerJob, and an error, if there is any.
func (c *FakeCloudSchedulerJobs) Create(ctx context.Context, cloudSchedulerJob *v1alpha1.CloudSchedulerJob, opts v1.CreateOptions) (result *v1alpha1.CloudSchedulerJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cloudschedulerjobsResource, c.ns, cloudSchedulerJob), &v1alpha1.CloudSchedulerJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudSchedulerJob), err
}

// Update takes the representation of a cloudSchedulerJob and updates it. Returns the server's representation of the cloudSchedulerJob, and an error, if there is any.
func (c *FakeCloudSchedulerJobs) Update(ctx context.Context, cloudSchedulerJob *v1alpha1.CloudSchedulerJob, opts v1.UpdateOptions) (result *v1alpha1.CloudSchedulerJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cloudschedulerjobsResource, c.ns, cloudSchedulerJob), &v1alpha1.CloudSchedulerJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudSchedulerJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloudSchedulerJobs) UpdateStatus(ctx context.Context, cloudSchedulerJob *v1alpha1.CloudSchedulerJob, opts v1.UpdateOptions) (*v1alpha1.CloudSchedulerJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cloudschedulerjobsResource, "status", c.ns, cloudSchedulerJob), &v1alpha1.CloudSchedulerJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudSchedulerJob), err
}

// Delete takes name of the cloudSchedulerJob and deletes it. Returns an error if one occurs.
func (c *FakeCloudSchedulerJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cloudschedulerjobsResource, c.ns, name), &v1alpha1.CloudSchedulerJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudSchedulerJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cloudschedulerjobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CloudSchedulerJobList{})
	return err
}

// Patch applies the patch and returns the patched cloudSchedulerJob.
func (c *FakeCloudSchedulerJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CloudSchedulerJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cloudschedulerjobsResource, c.ns, name, pt, data, subresources...), &v1alpha1.CloudSchedulerJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudSchedulerJob), err
}