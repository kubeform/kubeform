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

// FakeStorageTransferJobs implements StorageTransferJobInterface
type FakeStorageTransferJobs struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var storagetransferjobsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "storagetransferjobs"}

var storagetransferjobsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "StorageTransferJob"}

// Get takes name of the storageTransferJob, and returns the corresponding storageTransferJob object, and an error if there is any.
func (c *FakeStorageTransferJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.StorageTransferJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(storagetransferjobsResource, c.ns, name), &v1alpha1.StorageTransferJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageTransferJob), err
}

// List takes label and field selectors, and returns the list of StorageTransferJobs that match those selectors.
func (c *FakeStorageTransferJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.StorageTransferJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(storagetransferjobsResource, storagetransferjobsKind, c.ns, opts), &v1alpha1.StorageTransferJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StorageTransferJobList{ListMeta: obj.(*v1alpha1.StorageTransferJobList).ListMeta}
	for _, item := range obj.(*v1alpha1.StorageTransferJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storageTransferJobs.
func (c *FakeStorageTransferJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(storagetransferjobsResource, c.ns, opts))

}

// Create takes the representation of a storageTransferJob and creates it.  Returns the server's representation of the storageTransferJob, and an error, if there is any.
func (c *FakeStorageTransferJobs) Create(ctx context.Context, storageTransferJob *v1alpha1.StorageTransferJob, opts v1.CreateOptions) (result *v1alpha1.StorageTransferJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(storagetransferjobsResource, c.ns, storageTransferJob), &v1alpha1.StorageTransferJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageTransferJob), err
}

// Update takes the representation of a storageTransferJob and updates it. Returns the server's representation of the storageTransferJob, and an error, if there is any.
func (c *FakeStorageTransferJobs) Update(ctx context.Context, storageTransferJob *v1alpha1.StorageTransferJob, opts v1.UpdateOptions) (result *v1alpha1.StorageTransferJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(storagetransferjobsResource, c.ns, storageTransferJob), &v1alpha1.StorageTransferJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageTransferJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStorageTransferJobs) UpdateStatus(ctx context.Context, storageTransferJob *v1alpha1.StorageTransferJob, opts v1.UpdateOptions) (*v1alpha1.StorageTransferJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(storagetransferjobsResource, "status", c.ns, storageTransferJob), &v1alpha1.StorageTransferJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageTransferJob), err
}

// Delete takes name of the storageTransferJob and deletes it. Returns an error if one occurs.
func (c *FakeStorageTransferJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(storagetransferjobsResource, c.ns, name), &v1alpha1.StorageTransferJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStorageTransferJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(storagetransferjobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.StorageTransferJobList{})
	return err
}

// Patch applies the patch and returns the patched storageTransferJob.
func (c *FakeStorageTransferJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.StorageTransferJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(storagetransferjobsResource, c.ns, name, pt, data, subresources...), &v1alpha1.StorageTransferJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageTransferJob), err
}