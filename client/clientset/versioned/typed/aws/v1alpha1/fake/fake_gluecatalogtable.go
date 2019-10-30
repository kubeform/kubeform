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

// FakeGlueCatalogTables implements GlueCatalogTableInterface
type FakeGlueCatalogTables struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var gluecatalogtablesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "gluecatalogtables"}

var gluecatalogtablesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "GlueCatalogTable"}

// Get takes name of the glueCatalogTable, and returns the corresponding glueCatalogTable object, and an error if there is any.
func (c *FakeGlueCatalogTables) Get(name string, options v1.GetOptions) (result *v1alpha1.GlueCatalogTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(gluecatalogtablesResource, c.ns, name), &v1alpha1.GlueCatalogTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlueCatalogTable), err
}

// List takes label and field selectors, and returns the list of GlueCatalogTables that match those selectors.
func (c *FakeGlueCatalogTables) List(opts v1.ListOptions) (result *v1alpha1.GlueCatalogTableList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(gluecatalogtablesResource, gluecatalogtablesKind, c.ns, opts), &v1alpha1.GlueCatalogTableList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GlueCatalogTableList{ListMeta: obj.(*v1alpha1.GlueCatalogTableList).ListMeta}
	for _, item := range obj.(*v1alpha1.GlueCatalogTableList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested glueCatalogTables.
func (c *FakeGlueCatalogTables) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(gluecatalogtablesResource, c.ns, opts))

}

// Create takes the representation of a glueCatalogTable and creates it.  Returns the server's representation of the glueCatalogTable, and an error, if there is any.
func (c *FakeGlueCatalogTables) Create(glueCatalogTable *v1alpha1.GlueCatalogTable) (result *v1alpha1.GlueCatalogTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(gluecatalogtablesResource, c.ns, glueCatalogTable), &v1alpha1.GlueCatalogTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlueCatalogTable), err
}

// Update takes the representation of a glueCatalogTable and updates it. Returns the server's representation of the glueCatalogTable, and an error, if there is any.
func (c *FakeGlueCatalogTables) Update(glueCatalogTable *v1alpha1.GlueCatalogTable) (result *v1alpha1.GlueCatalogTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(gluecatalogtablesResource, c.ns, glueCatalogTable), &v1alpha1.GlueCatalogTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlueCatalogTable), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGlueCatalogTables) UpdateStatus(glueCatalogTable *v1alpha1.GlueCatalogTable) (*v1alpha1.GlueCatalogTable, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(gluecatalogtablesResource, "status", c.ns, glueCatalogTable), &v1alpha1.GlueCatalogTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlueCatalogTable), err
}

// Delete takes name of the glueCatalogTable and deletes it. Returns an error if one occurs.
func (c *FakeGlueCatalogTables) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(gluecatalogtablesResource, c.ns, name), &v1alpha1.GlueCatalogTable{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGlueCatalogTables) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(gluecatalogtablesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GlueCatalogTableList{})
	return err
}

// Patch applies the patch and returns the patched glueCatalogTable.
func (c *FakeGlueCatalogTables) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GlueCatalogTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(gluecatalogtablesResource, c.ns, name, pt, data, subresources...), &v1alpha1.GlueCatalogTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlueCatalogTable), err
}
