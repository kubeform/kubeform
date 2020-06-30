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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBackupPolicyFileShares implements BackupPolicyFileShareInterface
type FakeBackupPolicyFileShares struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var backuppolicyfilesharesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "backuppolicyfileshares"}

var backuppolicyfilesharesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "BackupPolicyFileShare"}

// Get takes name of the backupPolicyFileShare, and returns the corresponding backupPolicyFileShare object, and an error if there is any.
func (c *FakeBackupPolicyFileShares) Get(name string, options v1.GetOptions) (result *v1alpha1.BackupPolicyFileShare, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(backuppolicyfilesharesResource, c.ns, name), &v1alpha1.BackupPolicyFileShare{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupPolicyFileShare), err
}

// List takes label and field selectors, and returns the list of BackupPolicyFileShares that match those selectors.
func (c *FakeBackupPolicyFileShares) List(opts v1.ListOptions) (result *v1alpha1.BackupPolicyFileShareList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(backuppolicyfilesharesResource, backuppolicyfilesharesKind, c.ns, opts), &v1alpha1.BackupPolicyFileShareList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BackupPolicyFileShareList{ListMeta: obj.(*v1alpha1.BackupPolicyFileShareList).ListMeta}
	for _, item := range obj.(*v1alpha1.BackupPolicyFileShareList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested backupPolicyFileShares.
func (c *FakeBackupPolicyFileShares) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(backuppolicyfilesharesResource, c.ns, opts))

}

// Create takes the representation of a backupPolicyFileShare and creates it.  Returns the server's representation of the backupPolicyFileShare, and an error, if there is any.
func (c *FakeBackupPolicyFileShares) Create(backupPolicyFileShare *v1alpha1.BackupPolicyFileShare) (result *v1alpha1.BackupPolicyFileShare, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(backuppolicyfilesharesResource, c.ns, backupPolicyFileShare), &v1alpha1.BackupPolicyFileShare{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupPolicyFileShare), err
}

// Update takes the representation of a backupPolicyFileShare and updates it. Returns the server's representation of the backupPolicyFileShare, and an error, if there is any.
func (c *FakeBackupPolicyFileShares) Update(backupPolicyFileShare *v1alpha1.BackupPolicyFileShare) (result *v1alpha1.BackupPolicyFileShare, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(backuppolicyfilesharesResource, c.ns, backupPolicyFileShare), &v1alpha1.BackupPolicyFileShare{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupPolicyFileShare), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBackupPolicyFileShares) UpdateStatus(backupPolicyFileShare *v1alpha1.BackupPolicyFileShare) (*v1alpha1.BackupPolicyFileShare, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(backuppolicyfilesharesResource, "status", c.ns, backupPolicyFileShare), &v1alpha1.BackupPolicyFileShare{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupPolicyFileShare), err
}

// Delete takes name of the backupPolicyFileShare and deletes it. Returns an error if one occurs.
func (c *FakeBackupPolicyFileShares) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(backuppolicyfilesharesResource, c.ns, name), &v1alpha1.BackupPolicyFileShare{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBackupPolicyFileShares) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(backuppolicyfilesharesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.BackupPolicyFileShareList{})
	return err
}

// Patch applies the patch and returns the patched backupPolicyFileShare.
func (c *FakeBackupPolicyFileShares) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BackupPolicyFileShare, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(backuppolicyfilesharesResource, c.ns, name, pt, data, subresources...), &v1alpha1.BackupPolicyFileShare{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupPolicyFileShare), err
}
