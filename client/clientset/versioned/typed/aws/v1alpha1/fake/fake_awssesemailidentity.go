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

// FakeAwsSesEmailIdentities implements AwsSesEmailIdentityInterface
type FakeAwsSesEmailIdentities struct {
	Fake *FakeAwsV1alpha1
}

var awssesemailidentitiesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awssesemailidentities"}

var awssesemailidentitiesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsSesEmailIdentity"}

// Get takes name of the awsSesEmailIdentity, and returns the corresponding awsSesEmailIdentity object, and an error if there is any.
func (c *FakeAwsSesEmailIdentities) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSesEmailIdentity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awssesemailidentitiesResource, name), &v1alpha1.AwsSesEmailIdentity{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSesEmailIdentity), err
}

// List takes label and field selectors, and returns the list of AwsSesEmailIdentities that match those selectors.
func (c *FakeAwsSesEmailIdentities) List(opts v1.ListOptions) (result *v1alpha1.AwsSesEmailIdentityList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awssesemailidentitiesResource, awssesemailidentitiesKind, opts), &v1alpha1.AwsSesEmailIdentityList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsSesEmailIdentityList{ListMeta: obj.(*v1alpha1.AwsSesEmailIdentityList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsSesEmailIdentityList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsSesEmailIdentities.
func (c *FakeAwsSesEmailIdentities) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awssesemailidentitiesResource, opts))
}

// Create takes the representation of a awsSesEmailIdentity and creates it.  Returns the server's representation of the awsSesEmailIdentity, and an error, if there is any.
func (c *FakeAwsSesEmailIdentities) Create(awsSesEmailIdentity *v1alpha1.AwsSesEmailIdentity) (result *v1alpha1.AwsSesEmailIdentity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awssesemailidentitiesResource, awsSesEmailIdentity), &v1alpha1.AwsSesEmailIdentity{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSesEmailIdentity), err
}

// Update takes the representation of a awsSesEmailIdentity and updates it. Returns the server's representation of the awsSesEmailIdentity, and an error, if there is any.
func (c *FakeAwsSesEmailIdentities) Update(awsSesEmailIdentity *v1alpha1.AwsSesEmailIdentity) (result *v1alpha1.AwsSesEmailIdentity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awssesemailidentitiesResource, awsSesEmailIdentity), &v1alpha1.AwsSesEmailIdentity{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSesEmailIdentity), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsSesEmailIdentities) UpdateStatus(awsSesEmailIdentity *v1alpha1.AwsSesEmailIdentity) (*v1alpha1.AwsSesEmailIdentity, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awssesemailidentitiesResource, "status", awsSesEmailIdentity), &v1alpha1.AwsSesEmailIdentity{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSesEmailIdentity), err
}

// Delete takes name of the awsSesEmailIdentity and deletes it. Returns an error if one occurs.
func (c *FakeAwsSesEmailIdentities) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awssesemailidentitiesResource, name), &v1alpha1.AwsSesEmailIdentity{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsSesEmailIdentities) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awssesemailidentitiesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsSesEmailIdentityList{})
	return err
}

// Patch applies the patch and returns the patched awsSesEmailIdentity.
func (c *FakeAwsSesEmailIdentities) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSesEmailIdentity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awssesemailidentitiesResource, name, pt, data, subresources...), &v1alpha1.AwsSesEmailIdentity{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSesEmailIdentity), err
}
