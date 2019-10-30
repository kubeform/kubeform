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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTransferSSHKeys implements TransferSSHKeyInterface
type FakeTransferSSHKeys struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var transfersshkeysResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "transfersshkeys"}

var transfersshkeysKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "TransferSSHKey"}

// Get takes name of the transferSSHKey, and returns the corresponding transferSSHKey object, and an error if there is any.
func (c *FakeTransferSSHKeys) Get(name string, options v1.GetOptions) (result *v1alpha1.TransferSSHKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(transfersshkeysResource, c.ns, name), &v1alpha1.TransferSSHKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransferSSHKey), err
}

// List takes label and field selectors, and returns the list of TransferSSHKeys that match those selectors.
func (c *FakeTransferSSHKeys) List(opts v1.ListOptions) (result *v1alpha1.TransferSSHKeyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(transfersshkeysResource, transfersshkeysKind, c.ns, opts), &v1alpha1.TransferSSHKeyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TransferSSHKeyList{ListMeta: obj.(*v1alpha1.TransferSSHKeyList).ListMeta}
	for _, item := range obj.(*v1alpha1.TransferSSHKeyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested transferSSHKeys.
func (c *FakeTransferSSHKeys) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(transfersshkeysResource, c.ns, opts))

}

// Create takes the representation of a transferSSHKey and creates it.  Returns the server's representation of the transferSSHKey, and an error, if there is any.
func (c *FakeTransferSSHKeys) Create(transferSSHKey *v1alpha1.TransferSSHKey) (result *v1alpha1.TransferSSHKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(transfersshkeysResource, c.ns, transferSSHKey), &v1alpha1.TransferSSHKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransferSSHKey), err
}

// Update takes the representation of a transferSSHKey and updates it. Returns the server's representation of the transferSSHKey, and an error, if there is any.
func (c *FakeTransferSSHKeys) Update(transferSSHKey *v1alpha1.TransferSSHKey) (result *v1alpha1.TransferSSHKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(transfersshkeysResource, c.ns, transferSSHKey), &v1alpha1.TransferSSHKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransferSSHKey), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTransferSSHKeys) UpdateStatus(transferSSHKey *v1alpha1.TransferSSHKey) (*v1alpha1.TransferSSHKey, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(transfersshkeysResource, "status", c.ns, transferSSHKey), &v1alpha1.TransferSSHKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransferSSHKey), err
}

// Delete takes name of the transferSSHKey and deletes it. Returns an error if one occurs.
func (c *FakeTransferSSHKeys) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(transfersshkeysResource, c.ns, name), &v1alpha1.TransferSSHKey{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTransferSSHKeys) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(transfersshkeysResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.TransferSSHKeyList{})
	return err
}

// Patch applies the patch and returns the patched transferSSHKey.
func (c *FakeTransferSSHKeys) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TransferSSHKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(transfersshkeysResource, c.ns, name, pt, data, subresources...), &v1alpha1.TransferSSHKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransferSSHKey), err
}
