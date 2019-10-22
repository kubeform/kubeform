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

// FakeIamRoles implements IamRoleInterface
type FakeIamRoles struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var iamrolesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "iamroles"}

var iamrolesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "IamRole"}

// Get takes name of the iamRole, and returns the corresponding iamRole object, and an error if there is any.
func (c *FakeIamRoles) Get(name string, options v1.GetOptions) (result *v1alpha1.IamRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iamrolesResource, c.ns, name), &v1alpha1.IamRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamRole), err
}

// List takes label and field selectors, and returns the list of IamRoles that match those selectors.
func (c *FakeIamRoles) List(opts v1.ListOptions) (result *v1alpha1.IamRoleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iamrolesResource, iamrolesKind, c.ns, opts), &v1alpha1.IamRoleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IamRoleList{ListMeta: obj.(*v1alpha1.IamRoleList).ListMeta}
	for _, item := range obj.(*v1alpha1.IamRoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iamRoles.
func (c *FakeIamRoles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iamrolesResource, c.ns, opts))

}

// Create takes the representation of a iamRole and creates it.  Returns the server's representation of the iamRole, and an error, if there is any.
func (c *FakeIamRoles) Create(iamRole *v1alpha1.IamRole) (result *v1alpha1.IamRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iamrolesResource, c.ns, iamRole), &v1alpha1.IamRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamRole), err
}

// Update takes the representation of a iamRole and updates it. Returns the server's representation of the iamRole, and an error, if there is any.
func (c *FakeIamRoles) Update(iamRole *v1alpha1.IamRole) (result *v1alpha1.IamRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iamrolesResource, c.ns, iamRole), &v1alpha1.IamRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamRole), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIamRoles) UpdateStatus(iamRole *v1alpha1.IamRole) (*v1alpha1.IamRole, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(iamrolesResource, "status", c.ns, iamRole), &v1alpha1.IamRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamRole), err
}

// Delete takes name of the iamRole and deletes it. Returns an error if one occurs.
func (c *FakeIamRoles) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(iamrolesResource, c.ns, name), &v1alpha1.IamRole{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIamRoles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iamrolesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.IamRoleList{})
	return err
}

// Patch applies the patch and returns the patched iamRole.
func (c *FakeIamRoles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IamRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iamrolesResource, c.ns, name, pt, data, subresources...), &v1alpha1.IamRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamRole), err
}
