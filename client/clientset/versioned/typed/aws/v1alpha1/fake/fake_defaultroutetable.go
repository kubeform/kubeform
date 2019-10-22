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

// FakeDefaultRouteTables implements DefaultRouteTableInterface
type FakeDefaultRouteTables struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var defaultroutetablesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "defaultroutetables"}

var defaultroutetablesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DefaultRouteTable"}

// Get takes name of the defaultRouteTable, and returns the corresponding defaultRouteTable object, and an error if there is any.
func (c *FakeDefaultRouteTables) Get(name string, options v1.GetOptions) (result *v1alpha1.DefaultRouteTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(defaultroutetablesResource, c.ns, name), &v1alpha1.DefaultRouteTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DefaultRouteTable), err
}

// List takes label and field selectors, and returns the list of DefaultRouteTables that match those selectors.
func (c *FakeDefaultRouteTables) List(opts v1.ListOptions) (result *v1alpha1.DefaultRouteTableList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(defaultroutetablesResource, defaultroutetablesKind, c.ns, opts), &v1alpha1.DefaultRouteTableList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DefaultRouteTableList{ListMeta: obj.(*v1alpha1.DefaultRouteTableList).ListMeta}
	for _, item := range obj.(*v1alpha1.DefaultRouteTableList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested defaultRouteTables.
func (c *FakeDefaultRouteTables) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(defaultroutetablesResource, c.ns, opts))

}

// Create takes the representation of a defaultRouteTable and creates it.  Returns the server's representation of the defaultRouteTable, and an error, if there is any.
func (c *FakeDefaultRouteTables) Create(defaultRouteTable *v1alpha1.DefaultRouteTable) (result *v1alpha1.DefaultRouteTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(defaultroutetablesResource, c.ns, defaultRouteTable), &v1alpha1.DefaultRouteTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DefaultRouteTable), err
}

// Update takes the representation of a defaultRouteTable and updates it. Returns the server's representation of the defaultRouteTable, and an error, if there is any.
func (c *FakeDefaultRouteTables) Update(defaultRouteTable *v1alpha1.DefaultRouteTable) (result *v1alpha1.DefaultRouteTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(defaultroutetablesResource, c.ns, defaultRouteTable), &v1alpha1.DefaultRouteTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DefaultRouteTable), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDefaultRouteTables) UpdateStatus(defaultRouteTable *v1alpha1.DefaultRouteTable) (*v1alpha1.DefaultRouteTable, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(defaultroutetablesResource, "status", c.ns, defaultRouteTable), &v1alpha1.DefaultRouteTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DefaultRouteTable), err
}

// Delete takes name of the defaultRouteTable and deletes it. Returns an error if one occurs.
func (c *FakeDefaultRouteTables) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(defaultroutetablesResource, c.ns, name), &v1alpha1.DefaultRouteTable{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDefaultRouteTables) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(defaultroutetablesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DefaultRouteTableList{})
	return err
}

// Patch applies the patch and returns the patched defaultRouteTable.
func (c *FakeDefaultRouteTables) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DefaultRouteTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(defaultroutetablesResource, c.ns, name, pt, data, subresources...), &v1alpha1.DefaultRouteTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DefaultRouteTable), err
}
