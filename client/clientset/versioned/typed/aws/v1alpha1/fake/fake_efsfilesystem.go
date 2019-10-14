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

// FakeEfsFileSystems implements EfsFileSystemInterface
type FakeEfsFileSystems struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var efsfilesystemsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "efsfilesystems"}

var efsfilesystemsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "EfsFileSystem"}

// Get takes name of the efsFileSystem, and returns the corresponding efsFileSystem object, and an error if there is any.
func (c *FakeEfsFileSystems) Get(name string, options v1.GetOptions) (result *v1alpha1.EfsFileSystem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(efsfilesystemsResource, c.ns, name), &v1alpha1.EfsFileSystem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EfsFileSystem), err
}

// List takes label and field selectors, and returns the list of EfsFileSystems that match those selectors.
func (c *FakeEfsFileSystems) List(opts v1.ListOptions) (result *v1alpha1.EfsFileSystemList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(efsfilesystemsResource, efsfilesystemsKind, c.ns, opts), &v1alpha1.EfsFileSystemList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EfsFileSystemList{ListMeta: obj.(*v1alpha1.EfsFileSystemList).ListMeta}
	for _, item := range obj.(*v1alpha1.EfsFileSystemList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested efsFileSystems.
func (c *FakeEfsFileSystems) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(efsfilesystemsResource, c.ns, opts))

}

// Create takes the representation of a efsFileSystem and creates it.  Returns the server's representation of the efsFileSystem, and an error, if there is any.
func (c *FakeEfsFileSystems) Create(efsFileSystem *v1alpha1.EfsFileSystem) (result *v1alpha1.EfsFileSystem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(efsfilesystemsResource, c.ns, efsFileSystem), &v1alpha1.EfsFileSystem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EfsFileSystem), err
}

// Update takes the representation of a efsFileSystem and updates it. Returns the server's representation of the efsFileSystem, and an error, if there is any.
func (c *FakeEfsFileSystems) Update(efsFileSystem *v1alpha1.EfsFileSystem) (result *v1alpha1.EfsFileSystem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(efsfilesystemsResource, c.ns, efsFileSystem), &v1alpha1.EfsFileSystem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EfsFileSystem), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEfsFileSystems) UpdateStatus(efsFileSystem *v1alpha1.EfsFileSystem) (*v1alpha1.EfsFileSystem, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(efsfilesystemsResource, "status", c.ns, efsFileSystem), &v1alpha1.EfsFileSystem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EfsFileSystem), err
}

// Delete takes name of the efsFileSystem and deletes it. Returns an error if one occurs.
func (c *FakeEfsFileSystems) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(efsfilesystemsResource, c.ns, name), &v1alpha1.EfsFileSystem{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEfsFileSystems) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(efsfilesystemsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.EfsFileSystemList{})
	return err
}

// Patch applies the patch and returns the patched efsFileSystem.
func (c *FakeEfsFileSystems) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EfsFileSystem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(efsfilesystemsResource, c.ns, name, pt, data, subresources...), &v1alpha1.EfsFileSystem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EfsFileSystem), err
}
