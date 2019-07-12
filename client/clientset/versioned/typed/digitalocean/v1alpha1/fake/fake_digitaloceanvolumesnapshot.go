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
	v1alpha1 "kubeform.dev/kubeform/apis/digitalocean/v1alpha1"
)

// FakeDigitaloceanVolumeSnapshots implements DigitaloceanVolumeSnapshotInterface
type FakeDigitaloceanVolumeSnapshots struct {
	Fake *FakeDigitaloceanV1alpha1
}

var digitaloceanvolumesnapshotsResource = schema.GroupVersionResource{Group: "digitalocean.kubeform.com", Version: "v1alpha1", Resource: "digitaloceanvolumesnapshots"}

var digitaloceanvolumesnapshotsKind = schema.GroupVersionKind{Group: "digitalocean.kubeform.com", Version: "v1alpha1", Kind: "DigitaloceanVolumeSnapshot"}

// Get takes name of the digitaloceanVolumeSnapshot, and returns the corresponding digitaloceanVolumeSnapshot object, and an error if there is any.
func (c *FakeDigitaloceanVolumeSnapshots) Get(name string, options v1.GetOptions) (result *v1alpha1.DigitaloceanVolumeSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(digitaloceanvolumesnapshotsResource, name), &v1alpha1.DigitaloceanVolumeSnapshot{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DigitaloceanVolumeSnapshot), err
}

// List takes label and field selectors, and returns the list of DigitaloceanVolumeSnapshots that match those selectors.
func (c *FakeDigitaloceanVolumeSnapshots) List(opts v1.ListOptions) (result *v1alpha1.DigitaloceanVolumeSnapshotList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(digitaloceanvolumesnapshotsResource, digitaloceanvolumesnapshotsKind, opts), &v1alpha1.DigitaloceanVolumeSnapshotList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DigitaloceanVolumeSnapshotList{ListMeta: obj.(*v1alpha1.DigitaloceanVolumeSnapshotList).ListMeta}
	for _, item := range obj.(*v1alpha1.DigitaloceanVolumeSnapshotList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested digitaloceanVolumeSnapshots.
func (c *FakeDigitaloceanVolumeSnapshots) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(digitaloceanvolumesnapshotsResource, opts))
}

// Create takes the representation of a digitaloceanVolumeSnapshot and creates it.  Returns the server's representation of the digitaloceanVolumeSnapshot, and an error, if there is any.
func (c *FakeDigitaloceanVolumeSnapshots) Create(digitaloceanVolumeSnapshot *v1alpha1.DigitaloceanVolumeSnapshot) (result *v1alpha1.DigitaloceanVolumeSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(digitaloceanvolumesnapshotsResource, digitaloceanVolumeSnapshot), &v1alpha1.DigitaloceanVolumeSnapshot{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DigitaloceanVolumeSnapshot), err
}

// Update takes the representation of a digitaloceanVolumeSnapshot and updates it. Returns the server's representation of the digitaloceanVolumeSnapshot, and an error, if there is any.
func (c *FakeDigitaloceanVolumeSnapshots) Update(digitaloceanVolumeSnapshot *v1alpha1.DigitaloceanVolumeSnapshot) (result *v1alpha1.DigitaloceanVolumeSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(digitaloceanvolumesnapshotsResource, digitaloceanVolumeSnapshot), &v1alpha1.DigitaloceanVolumeSnapshot{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DigitaloceanVolumeSnapshot), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDigitaloceanVolumeSnapshots) UpdateStatus(digitaloceanVolumeSnapshot *v1alpha1.DigitaloceanVolumeSnapshot) (*v1alpha1.DigitaloceanVolumeSnapshot, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(digitaloceanvolumesnapshotsResource, "status", digitaloceanVolumeSnapshot), &v1alpha1.DigitaloceanVolumeSnapshot{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DigitaloceanVolumeSnapshot), err
}

// Delete takes name of the digitaloceanVolumeSnapshot and deletes it. Returns an error if one occurs.
func (c *FakeDigitaloceanVolumeSnapshots) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(digitaloceanvolumesnapshotsResource, name), &v1alpha1.DigitaloceanVolumeSnapshot{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDigitaloceanVolumeSnapshots) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(digitaloceanvolumesnapshotsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DigitaloceanVolumeSnapshotList{})
	return err
}

// Patch applies the patch and returns the patched digitaloceanVolumeSnapshot.
func (c *FakeDigitaloceanVolumeSnapshots) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DigitaloceanVolumeSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(digitaloceanvolumesnapshotsResource, name, pt, data, subresources...), &v1alpha1.DigitaloceanVolumeSnapshot{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DigitaloceanVolumeSnapshot), err
}
