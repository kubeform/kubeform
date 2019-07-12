/*
Copyright 2019 The KubeDB Authors.

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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeAzurermCosmosdbSqlDatabases implements AzurermCosmosdbSqlDatabaseInterface
type FakeAzurermCosmosdbSqlDatabases struct {
	Fake *FakeAzurermV1alpha1
}

var azurermcosmosdbsqldatabasesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermcosmosdbsqldatabases"}

var azurermcosmosdbsqldatabasesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermCosmosdbSqlDatabase"}

// Get takes name of the azurermCosmosdbSqlDatabase, and returns the corresponding azurermCosmosdbSqlDatabase object, and an error if there is any.
func (c *FakeAzurermCosmosdbSqlDatabases) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermCosmosdbSqlDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermcosmosdbsqldatabasesResource, name), &v1alpha1.AzurermCosmosdbSqlDatabase{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermCosmosdbSqlDatabase), err
}

// List takes label and field selectors, and returns the list of AzurermCosmosdbSqlDatabases that match those selectors.
func (c *FakeAzurermCosmosdbSqlDatabases) List(opts v1.ListOptions) (result *v1alpha1.AzurermCosmosdbSqlDatabaseList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermcosmosdbsqldatabasesResource, azurermcosmosdbsqldatabasesKind, opts), &v1alpha1.AzurermCosmosdbSqlDatabaseList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermCosmosdbSqlDatabaseList{ListMeta: obj.(*v1alpha1.AzurermCosmosdbSqlDatabaseList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermCosmosdbSqlDatabaseList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermCosmosdbSqlDatabases.
func (c *FakeAzurermCosmosdbSqlDatabases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermcosmosdbsqldatabasesResource, opts))
}

// Create takes the representation of a azurermCosmosdbSqlDatabase and creates it.  Returns the server's representation of the azurermCosmosdbSqlDatabase, and an error, if there is any.
func (c *FakeAzurermCosmosdbSqlDatabases) Create(azurermCosmosdbSqlDatabase *v1alpha1.AzurermCosmosdbSqlDatabase) (result *v1alpha1.AzurermCosmosdbSqlDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermcosmosdbsqldatabasesResource, azurermCosmosdbSqlDatabase), &v1alpha1.AzurermCosmosdbSqlDatabase{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermCosmosdbSqlDatabase), err
}

// Update takes the representation of a azurermCosmosdbSqlDatabase and updates it. Returns the server's representation of the azurermCosmosdbSqlDatabase, and an error, if there is any.
func (c *FakeAzurermCosmosdbSqlDatabases) Update(azurermCosmosdbSqlDatabase *v1alpha1.AzurermCosmosdbSqlDatabase) (result *v1alpha1.AzurermCosmosdbSqlDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermcosmosdbsqldatabasesResource, azurermCosmosdbSqlDatabase), &v1alpha1.AzurermCosmosdbSqlDatabase{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermCosmosdbSqlDatabase), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermCosmosdbSqlDatabases) UpdateStatus(azurermCosmosdbSqlDatabase *v1alpha1.AzurermCosmosdbSqlDatabase) (*v1alpha1.AzurermCosmosdbSqlDatabase, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermcosmosdbsqldatabasesResource, "status", azurermCosmosdbSqlDatabase), &v1alpha1.AzurermCosmosdbSqlDatabase{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermCosmosdbSqlDatabase), err
}

// Delete takes name of the azurermCosmosdbSqlDatabase and deletes it. Returns an error if one occurs.
func (c *FakeAzurermCosmosdbSqlDatabases) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermcosmosdbsqldatabasesResource, name), &v1alpha1.AzurermCosmosdbSqlDatabase{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermCosmosdbSqlDatabases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermcosmosdbsqldatabasesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermCosmosdbSqlDatabaseList{})
	return err
}

// Patch applies the patch and returns the patched azurermCosmosdbSqlDatabase.
func (c *FakeAzurermCosmosdbSqlDatabases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermCosmosdbSqlDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermcosmosdbsqldatabasesResource, name, pt, data, subresources...), &v1alpha1.AzurermCosmosdbSqlDatabase{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermCosmosdbSqlDatabase), err
}
