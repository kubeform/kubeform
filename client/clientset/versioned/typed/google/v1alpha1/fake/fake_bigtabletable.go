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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeBigtableTables implements BigtableTableInterface
type FakeBigtableTables struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var bigtabletablesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "bigtabletables"}

var bigtabletablesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "BigtableTable"}

// Get takes name of the bigtableTable, and returns the corresponding bigtableTable object, and an error if there is any.
func (c *FakeBigtableTables) Get(name string, options v1.GetOptions) (result *v1alpha1.BigtableTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(bigtabletablesResource, c.ns, name), &v1alpha1.BigtableTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BigtableTable), err
}

// List takes label and field selectors, and returns the list of BigtableTables that match those selectors.
func (c *FakeBigtableTables) List(opts v1.ListOptions) (result *v1alpha1.BigtableTableList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(bigtabletablesResource, bigtabletablesKind, c.ns, opts), &v1alpha1.BigtableTableList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BigtableTableList{ListMeta: obj.(*v1alpha1.BigtableTableList).ListMeta}
	for _, item := range obj.(*v1alpha1.BigtableTableList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested bigtableTables.
func (c *FakeBigtableTables) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(bigtabletablesResource, c.ns, opts))

}

// Create takes the representation of a bigtableTable and creates it.  Returns the server's representation of the bigtableTable, and an error, if there is any.
func (c *FakeBigtableTables) Create(bigtableTable *v1alpha1.BigtableTable) (result *v1alpha1.BigtableTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(bigtabletablesResource, c.ns, bigtableTable), &v1alpha1.BigtableTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BigtableTable), err
}

// Update takes the representation of a bigtableTable and updates it. Returns the server's representation of the bigtableTable, and an error, if there is any.
func (c *FakeBigtableTables) Update(bigtableTable *v1alpha1.BigtableTable) (result *v1alpha1.BigtableTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(bigtabletablesResource, c.ns, bigtableTable), &v1alpha1.BigtableTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BigtableTable), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBigtableTables) UpdateStatus(bigtableTable *v1alpha1.BigtableTable) (*v1alpha1.BigtableTable, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(bigtabletablesResource, "status", c.ns, bigtableTable), &v1alpha1.BigtableTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BigtableTable), err
}

// Delete takes name of the bigtableTable and deletes it. Returns an error if one occurs.
func (c *FakeBigtableTables) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(bigtabletablesResource, c.ns, name), &v1alpha1.BigtableTable{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBigtableTables) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(bigtabletablesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.BigtableTableList{})
	return err
}

// Patch applies the patch and returns the patched bigtableTable.
func (c *FakeBigtableTables) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BigtableTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(bigtabletablesResource, c.ns, name, pt, data, subresources...), &v1alpha1.BigtableTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BigtableTable), err
}
