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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStoragegatewayNfsFileShares implements StoragegatewayNfsFileShareInterface
type FakeStoragegatewayNfsFileShares struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var storagegatewaynfsfilesharesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "storagegatewaynfsfileshares"}

var storagegatewaynfsfilesharesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "StoragegatewayNfsFileShare"}

// Get takes name of the storagegatewayNfsFileShare, and returns the corresponding storagegatewayNfsFileShare object, and an error if there is any.
func (c *FakeStoragegatewayNfsFileShares) Get(name string, options v1.GetOptions) (result *v1alpha1.StoragegatewayNfsFileShare, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(storagegatewaynfsfilesharesResource, c.ns, name), &v1alpha1.StoragegatewayNfsFileShare{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StoragegatewayNfsFileShare), err
}

// List takes label and field selectors, and returns the list of StoragegatewayNfsFileShares that match those selectors.
func (c *FakeStoragegatewayNfsFileShares) List(opts v1.ListOptions) (result *v1alpha1.StoragegatewayNfsFileShareList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(storagegatewaynfsfilesharesResource, storagegatewaynfsfilesharesKind, c.ns, opts), &v1alpha1.StoragegatewayNfsFileShareList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StoragegatewayNfsFileShareList{ListMeta: obj.(*v1alpha1.StoragegatewayNfsFileShareList).ListMeta}
	for _, item := range obj.(*v1alpha1.StoragegatewayNfsFileShareList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storagegatewayNfsFileShares.
func (c *FakeStoragegatewayNfsFileShares) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(storagegatewaynfsfilesharesResource, c.ns, opts))

}

// Create takes the representation of a storagegatewayNfsFileShare and creates it.  Returns the server's representation of the storagegatewayNfsFileShare, and an error, if there is any.
func (c *FakeStoragegatewayNfsFileShares) Create(storagegatewayNfsFileShare *v1alpha1.StoragegatewayNfsFileShare) (result *v1alpha1.StoragegatewayNfsFileShare, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(storagegatewaynfsfilesharesResource, c.ns, storagegatewayNfsFileShare), &v1alpha1.StoragegatewayNfsFileShare{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StoragegatewayNfsFileShare), err
}

// Update takes the representation of a storagegatewayNfsFileShare and updates it. Returns the server's representation of the storagegatewayNfsFileShare, and an error, if there is any.
func (c *FakeStoragegatewayNfsFileShares) Update(storagegatewayNfsFileShare *v1alpha1.StoragegatewayNfsFileShare) (result *v1alpha1.StoragegatewayNfsFileShare, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(storagegatewaynfsfilesharesResource, c.ns, storagegatewayNfsFileShare), &v1alpha1.StoragegatewayNfsFileShare{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StoragegatewayNfsFileShare), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStoragegatewayNfsFileShares) UpdateStatus(storagegatewayNfsFileShare *v1alpha1.StoragegatewayNfsFileShare) (*v1alpha1.StoragegatewayNfsFileShare, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(storagegatewaynfsfilesharesResource, "status", c.ns, storagegatewayNfsFileShare), &v1alpha1.StoragegatewayNfsFileShare{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StoragegatewayNfsFileShare), err
}

// Delete takes name of the storagegatewayNfsFileShare and deletes it. Returns an error if one occurs.
func (c *FakeStoragegatewayNfsFileShares) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(storagegatewaynfsfilesharesResource, c.ns, name), &v1alpha1.StoragegatewayNfsFileShare{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStoragegatewayNfsFileShares) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(storagegatewaynfsfilesharesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.StoragegatewayNfsFileShareList{})
	return err
}

// Patch applies the patch and returns the patched storagegatewayNfsFileShare.
func (c *FakeStoragegatewayNfsFileShares) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StoragegatewayNfsFileShare, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(storagegatewaynfsfilesharesResource, c.ns, name, pt, data, subresources...), &v1alpha1.StoragegatewayNfsFileShare{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StoragegatewayNfsFileShare), err
}
