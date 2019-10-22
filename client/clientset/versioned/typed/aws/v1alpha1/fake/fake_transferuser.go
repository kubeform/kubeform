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

// FakeTransferUsers implements TransferUserInterface
type FakeTransferUsers struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var transferusersResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "transferusers"}

var transferusersKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "TransferUser"}

// Get takes name of the transferUser, and returns the corresponding transferUser object, and an error if there is any.
func (c *FakeTransferUsers) Get(name string, options v1.GetOptions) (result *v1alpha1.TransferUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(transferusersResource, c.ns, name), &v1alpha1.TransferUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransferUser), err
}

// List takes label and field selectors, and returns the list of TransferUsers that match those selectors.
func (c *FakeTransferUsers) List(opts v1.ListOptions) (result *v1alpha1.TransferUserList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(transferusersResource, transferusersKind, c.ns, opts), &v1alpha1.TransferUserList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TransferUserList{ListMeta: obj.(*v1alpha1.TransferUserList).ListMeta}
	for _, item := range obj.(*v1alpha1.TransferUserList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested transferUsers.
func (c *FakeTransferUsers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(transferusersResource, c.ns, opts))

}

// Create takes the representation of a transferUser and creates it.  Returns the server's representation of the transferUser, and an error, if there is any.
func (c *FakeTransferUsers) Create(transferUser *v1alpha1.TransferUser) (result *v1alpha1.TransferUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(transferusersResource, c.ns, transferUser), &v1alpha1.TransferUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransferUser), err
}

// Update takes the representation of a transferUser and updates it. Returns the server's representation of the transferUser, and an error, if there is any.
func (c *FakeTransferUsers) Update(transferUser *v1alpha1.TransferUser) (result *v1alpha1.TransferUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(transferusersResource, c.ns, transferUser), &v1alpha1.TransferUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransferUser), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTransferUsers) UpdateStatus(transferUser *v1alpha1.TransferUser) (*v1alpha1.TransferUser, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(transferusersResource, "status", c.ns, transferUser), &v1alpha1.TransferUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransferUser), err
}

// Delete takes name of the transferUser and deletes it. Returns an error if one occurs.
func (c *FakeTransferUsers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(transferusersResource, c.ns, name), &v1alpha1.TransferUser{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTransferUsers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(transferusersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.TransferUserList{})
	return err
}

// Patch applies the patch and returns the patched transferUser.
func (c *FakeTransferUsers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TransferUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(transferusersResource, c.ns, name, pt, data, subresources...), &v1alpha1.TransferUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransferUser), err
}
