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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCosmosdbSQLDatabases implements CosmosdbSQLDatabaseInterface
type FakeCosmosdbSQLDatabases struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var cosmosdbsqldatabasesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "cosmosdbsqldatabases"}

var cosmosdbsqldatabasesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "CosmosdbSQLDatabase"}

// Get takes name of the cosmosdbSQLDatabase, and returns the corresponding cosmosdbSQLDatabase object, and an error if there is any.
func (c *FakeCosmosdbSQLDatabases) Get(name string, options v1.GetOptions) (result *v1alpha1.CosmosdbSQLDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cosmosdbsqldatabasesResource, c.ns, name), &v1alpha1.CosmosdbSQLDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CosmosdbSQLDatabase), err
}

// List takes label and field selectors, and returns the list of CosmosdbSQLDatabases that match those selectors.
func (c *FakeCosmosdbSQLDatabases) List(opts v1.ListOptions) (result *v1alpha1.CosmosdbSQLDatabaseList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cosmosdbsqldatabasesResource, cosmosdbsqldatabasesKind, c.ns, opts), &v1alpha1.CosmosdbSQLDatabaseList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CosmosdbSQLDatabaseList{ListMeta: obj.(*v1alpha1.CosmosdbSQLDatabaseList).ListMeta}
	for _, item := range obj.(*v1alpha1.CosmosdbSQLDatabaseList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cosmosdbSQLDatabases.
func (c *FakeCosmosdbSQLDatabases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cosmosdbsqldatabasesResource, c.ns, opts))

}

// Create takes the representation of a cosmosdbSQLDatabase and creates it.  Returns the server's representation of the cosmosdbSQLDatabase, and an error, if there is any.
func (c *FakeCosmosdbSQLDatabases) Create(cosmosdbSQLDatabase *v1alpha1.CosmosdbSQLDatabase) (result *v1alpha1.CosmosdbSQLDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cosmosdbsqldatabasesResource, c.ns, cosmosdbSQLDatabase), &v1alpha1.CosmosdbSQLDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CosmosdbSQLDatabase), err
}

// Update takes the representation of a cosmosdbSQLDatabase and updates it. Returns the server's representation of the cosmosdbSQLDatabase, and an error, if there is any.
func (c *FakeCosmosdbSQLDatabases) Update(cosmosdbSQLDatabase *v1alpha1.CosmosdbSQLDatabase) (result *v1alpha1.CosmosdbSQLDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cosmosdbsqldatabasesResource, c.ns, cosmosdbSQLDatabase), &v1alpha1.CosmosdbSQLDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CosmosdbSQLDatabase), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCosmosdbSQLDatabases) UpdateStatus(cosmosdbSQLDatabase *v1alpha1.CosmosdbSQLDatabase) (*v1alpha1.CosmosdbSQLDatabase, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cosmosdbsqldatabasesResource, "status", c.ns, cosmosdbSQLDatabase), &v1alpha1.CosmosdbSQLDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CosmosdbSQLDatabase), err
}

// Delete takes name of the cosmosdbSQLDatabase and deletes it. Returns an error if one occurs.
func (c *FakeCosmosdbSQLDatabases) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cosmosdbsqldatabasesResource, c.ns, name), &v1alpha1.CosmosdbSQLDatabase{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCosmosdbSQLDatabases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cosmosdbsqldatabasesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CosmosdbSQLDatabaseList{})
	return err
}

// Patch applies the patch and returns the patched cosmosdbSQLDatabase.
func (c *FakeCosmosdbSQLDatabases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CosmosdbSQLDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cosmosdbsqldatabasesResource, c.ns, name, pt, data, subresources...), &v1alpha1.CosmosdbSQLDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CosmosdbSQLDatabase), err
}
