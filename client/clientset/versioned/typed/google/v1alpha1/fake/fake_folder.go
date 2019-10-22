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

// FakeFolders implements FolderInterface
type FakeFolders struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var foldersResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "folders"}

var foldersKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "Folder"}

// Get takes name of the folder, and returns the corresponding folder object, and an error if there is any.
func (c *FakeFolders) Get(name string, options v1.GetOptions) (result *v1alpha1.Folder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(foldersResource, c.ns, name), &v1alpha1.Folder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Folder), err
}

// List takes label and field selectors, and returns the list of Folders that match those selectors.
func (c *FakeFolders) List(opts v1.ListOptions) (result *v1alpha1.FolderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(foldersResource, foldersKind, c.ns, opts), &v1alpha1.FolderList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FolderList{ListMeta: obj.(*v1alpha1.FolderList).ListMeta}
	for _, item := range obj.(*v1alpha1.FolderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested folders.
func (c *FakeFolders) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(foldersResource, c.ns, opts))

}

// Create takes the representation of a folder and creates it.  Returns the server's representation of the folder, and an error, if there is any.
func (c *FakeFolders) Create(folder *v1alpha1.Folder) (result *v1alpha1.Folder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(foldersResource, c.ns, folder), &v1alpha1.Folder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Folder), err
}

// Update takes the representation of a folder and updates it. Returns the server's representation of the folder, and an error, if there is any.
func (c *FakeFolders) Update(folder *v1alpha1.Folder) (result *v1alpha1.Folder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(foldersResource, c.ns, folder), &v1alpha1.Folder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Folder), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFolders) UpdateStatus(folder *v1alpha1.Folder) (*v1alpha1.Folder, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(foldersResource, "status", c.ns, folder), &v1alpha1.Folder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Folder), err
}

// Delete takes name of the folder and deletes it. Returns an error if one occurs.
func (c *FakeFolders) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(foldersResource, c.ns, name), &v1alpha1.Folder{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFolders) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(foldersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.FolderList{})
	return err
}

// Patch applies the patch and returns the patched folder.
func (c *FakeFolders) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Folder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(foldersResource, c.ns, name, pt, data, subresources...), &v1alpha1.Folder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Folder), err
}
