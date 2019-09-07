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

// FakeDbSnapshots implements DbSnapshotInterface
type FakeDbSnapshots struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var dbsnapshotsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "dbsnapshots"}

var dbsnapshotsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DbSnapshot"}

// Get takes name of the dbSnapshot, and returns the corresponding dbSnapshot object, and an error if there is any.
func (c *FakeDbSnapshots) Get(name string, options v1.GetOptions) (result *v1alpha1.DbSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dbsnapshotsResource, c.ns, name), &v1alpha1.DbSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DbSnapshot), err
}

// List takes label and field selectors, and returns the list of DbSnapshots that match those selectors.
func (c *FakeDbSnapshots) List(opts v1.ListOptions) (result *v1alpha1.DbSnapshotList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dbsnapshotsResource, dbsnapshotsKind, c.ns, opts), &v1alpha1.DbSnapshotList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DbSnapshotList{ListMeta: obj.(*v1alpha1.DbSnapshotList).ListMeta}
	for _, item := range obj.(*v1alpha1.DbSnapshotList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dbSnapshots.
func (c *FakeDbSnapshots) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dbsnapshotsResource, c.ns, opts))

}

// Create takes the representation of a dbSnapshot and creates it.  Returns the server's representation of the dbSnapshot, and an error, if there is any.
func (c *FakeDbSnapshots) Create(dbSnapshot *v1alpha1.DbSnapshot) (result *v1alpha1.DbSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dbsnapshotsResource, c.ns, dbSnapshot), &v1alpha1.DbSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DbSnapshot), err
}

// Update takes the representation of a dbSnapshot and updates it. Returns the server's representation of the dbSnapshot, and an error, if there is any.
func (c *FakeDbSnapshots) Update(dbSnapshot *v1alpha1.DbSnapshot) (result *v1alpha1.DbSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dbsnapshotsResource, c.ns, dbSnapshot), &v1alpha1.DbSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DbSnapshot), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDbSnapshots) UpdateStatus(dbSnapshot *v1alpha1.DbSnapshot) (*v1alpha1.DbSnapshot, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dbsnapshotsResource, "status", c.ns, dbSnapshot), &v1alpha1.DbSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DbSnapshot), err
}

// Delete takes name of the dbSnapshot and deletes it. Returns an error if one occurs.
func (c *FakeDbSnapshots) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dbsnapshotsResource, c.ns, name), &v1alpha1.DbSnapshot{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDbSnapshots) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dbsnapshotsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DbSnapshotList{})
	return err
}

// Patch applies the patch and returns the patched dbSnapshot.
func (c *FakeDbSnapshots) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DbSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dbsnapshotsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DbSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DbSnapshot), err
}