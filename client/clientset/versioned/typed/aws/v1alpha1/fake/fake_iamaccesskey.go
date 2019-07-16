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

// FakeIamAccessKeys implements IamAccessKeyInterface
type FakeIamAccessKeys struct {
	Fake *FakeAwsV1alpha1
}

var iamaccesskeysResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "iamaccesskeys"}

var iamaccesskeysKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "IamAccessKey"}

// Get takes name of the iamAccessKey, and returns the corresponding iamAccessKey object, and an error if there is any.
func (c *FakeIamAccessKeys) Get(name string, options v1.GetOptions) (result *v1alpha1.IamAccessKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(iamaccesskeysResource, name), &v1alpha1.IamAccessKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamAccessKey), err
}

// List takes label and field selectors, and returns the list of IamAccessKeys that match those selectors.
func (c *FakeIamAccessKeys) List(opts v1.ListOptions) (result *v1alpha1.IamAccessKeyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(iamaccesskeysResource, iamaccesskeysKind, opts), &v1alpha1.IamAccessKeyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IamAccessKeyList{ListMeta: obj.(*v1alpha1.IamAccessKeyList).ListMeta}
	for _, item := range obj.(*v1alpha1.IamAccessKeyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iamAccessKeys.
func (c *FakeIamAccessKeys) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(iamaccesskeysResource, opts))
}

// Create takes the representation of a iamAccessKey and creates it.  Returns the server's representation of the iamAccessKey, and an error, if there is any.
func (c *FakeIamAccessKeys) Create(iamAccessKey *v1alpha1.IamAccessKey) (result *v1alpha1.IamAccessKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(iamaccesskeysResource, iamAccessKey), &v1alpha1.IamAccessKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamAccessKey), err
}

// Update takes the representation of a iamAccessKey and updates it. Returns the server's representation of the iamAccessKey, and an error, if there is any.
func (c *FakeIamAccessKeys) Update(iamAccessKey *v1alpha1.IamAccessKey) (result *v1alpha1.IamAccessKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(iamaccesskeysResource, iamAccessKey), &v1alpha1.IamAccessKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamAccessKey), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIamAccessKeys) UpdateStatus(iamAccessKey *v1alpha1.IamAccessKey) (*v1alpha1.IamAccessKey, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(iamaccesskeysResource, "status", iamAccessKey), &v1alpha1.IamAccessKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamAccessKey), err
}

// Delete takes name of the iamAccessKey and deletes it. Returns an error if one occurs.
func (c *FakeIamAccessKeys) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(iamaccesskeysResource, name), &v1alpha1.IamAccessKey{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIamAccessKeys) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(iamaccesskeysResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.IamAccessKeyList{})
	return err
}

// Patch applies the patch and returns the patched iamAccessKey.
func (c *FakeIamAccessKeys) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IamAccessKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(iamaccesskeysResource, name, pt, data, subresources...), &v1alpha1.IamAccessKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamAccessKey), err
}
