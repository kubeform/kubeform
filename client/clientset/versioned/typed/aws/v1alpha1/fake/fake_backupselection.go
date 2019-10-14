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

// FakeBackupSelections implements BackupSelectionInterface
type FakeBackupSelections struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var backupselectionsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "backupselections"}

var backupselectionsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "BackupSelection"}

// Get takes name of the backupSelection, and returns the corresponding backupSelection object, and an error if there is any.
func (c *FakeBackupSelections) Get(name string, options v1.GetOptions) (result *v1alpha1.BackupSelection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(backupselectionsResource, c.ns, name), &v1alpha1.BackupSelection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupSelection), err
}

// List takes label and field selectors, and returns the list of BackupSelections that match those selectors.
func (c *FakeBackupSelections) List(opts v1.ListOptions) (result *v1alpha1.BackupSelectionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(backupselectionsResource, backupselectionsKind, c.ns, opts), &v1alpha1.BackupSelectionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BackupSelectionList{ListMeta: obj.(*v1alpha1.BackupSelectionList).ListMeta}
	for _, item := range obj.(*v1alpha1.BackupSelectionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested backupSelections.
func (c *FakeBackupSelections) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(backupselectionsResource, c.ns, opts))

}

// Create takes the representation of a backupSelection and creates it.  Returns the server's representation of the backupSelection, and an error, if there is any.
func (c *FakeBackupSelections) Create(backupSelection *v1alpha1.BackupSelection) (result *v1alpha1.BackupSelection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(backupselectionsResource, c.ns, backupSelection), &v1alpha1.BackupSelection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupSelection), err
}

// Update takes the representation of a backupSelection and updates it. Returns the server's representation of the backupSelection, and an error, if there is any.
func (c *FakeBackupSelections) Update(backupSelection *v1alpha1.BackupSelection) (result *v1alpha1.BackupSelection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(backupselectionsResource, c.ns, backupSelection), &v1alpha1.BackupSelection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupSelection), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBackupSelections) UpdateStatus(backupSelection *v1alpha1.BackupSelection) (*v1alpha1.BackupSelection, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(backupselectionsResource, "status", c.ns, backupSelection), &v1alpha1.BackupSelection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupSelection), err
}

// Delete takes name of the backupSelection and deletes it. Returns an error if one occurs.
func (c *FakeBackupSelections) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(backupselectionsResource, c.ns, name), &v1alpha1.BackupSelection{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBackupSelections) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(backupselectionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.BackupSelectionList{})
	return err
}

// Patch applies the patch and returns the patched backupSelection.
func (c *FakeBackupSelections) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BackupSelection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(backupselectionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.BackupSelection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupSelection), err
}
