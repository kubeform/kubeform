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

// FakeVpcDHCPOptionses implements VpcDHCPOptionsInterface
type FakeVpcDHCPOptionses struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var vpcdhcpoptionsesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "vpcdhcpoptionses"}

var vpcdhcpoptionsesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "VpcDHCPOptions"}

// Get takes name of the vpcDHCPOptions, and returns the corresponding vpcDHCPOptions object, and an error if there is any.
func (c *FakeVpcDHCPOptionses) Get(name string, options v1.GetOptions) (result *v1alpha1.VpcDHCPOptions, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vpcdhcpoptionsesResource, c.ns, name), &v1alpha1.VpcDHCPOptions{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcDHCPOptions), err
}

// List takes label and field selectors, and returns the list of VpcDHCPOptionses that match those selectors.
func (c *FakeVpcDHCPOptionses) List(opts v1.ListOptions) (result *v1alpha1.VpcDHCPOptionsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vpcdhcpoptionsesResource, vpcdhcpoptionsesKind, c.ns, opts), &v1alpha1.VpcDHCPOptionsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VpcDHCPOptionsList{ListMeta: obj.(*v1alpha1.VpcDHCPOptionsList).ListMeta}
	for _, item := range obj.(*v1alpha1.VpcDHCPOptionsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vpcDHCPOptionses.
func (c *FakeVpcDHCPOptionses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vpcdhcpoptionsesResource, c.ns, opts))

}

// Create takes the representation of a vpcDHCPOptions and creates it.  Returns the server's representation of the vpcDHCPOptions, and an error, if there is any.
func (c *FakeVpcDHCPOptionses) Create(vpcDHCPOptions *v1alpha1.VpcDHCPOptions) (result *v1alpha1.VpcDHCPOptions, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vpcdhcpoptionsesResource, c.ns, vpcDHCPOptions), &v1alpha1.VpcDHCPOptions{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcDHCPOptions), err
}

// Update takes the representation of a vpcDHCPOptions and updates it. Returns the server's representation of the vpcDHCPOptions, and an error, if there is any.
func (c *FakeVpcDHCPOptionses) Update(vpcDHCPOptions *v1alpha1.VpcDHCPOptions) (result *v1alpha1.VpcDHCPOptions, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vpcdhcpoptionsesResource, c.ns, vpcDHCPOptions), &v1alpha1.VpcDHCPOptions{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcDHCPOptions), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVpcDHCPOptionses) UpdateStatus(vpcDHCPOptions *v1alpha1.VpcDHCPOptions) (*v1alpha1.VpcDHCPOptions, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vpcdhcpoptionsesResource, "status", c.ns, vpcDHCPOptions), &v1alpha1.VpcDHCPOptions{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcDHCPOptions), err
}

// Delete takes name of the vpcDHCPOptions and deletes it. Returns an error if one occurs.
func (c *FakeVpcDHCPOptionses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(vpcdhcpoptionsesResource, c.ns, name), &v1alpha1.VpcDHCPOptions{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVpcDHCPOptionses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vpcdhcpoptionsesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.VpcDHCPOptionsList{})
	return err
}

// Patch applies the patch and returns the patched vpcDHCPOptions.
func (c *FakeVpcDHCPOptionses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VpcDHCPOptions, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vpcdhcpoptionsesResource, c.ns, name, pt, data, subresources...), &v1alpha1.VpcDHCPOptions{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcDHCPOptions), err
}