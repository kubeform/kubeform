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

// FakePostgresqlDatabases implements PostgresqlDatabaseInterface
type FakePostgresqlDatabases struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var postgresqldatabasesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "postgresqldatabases"}

var postgresqldatabasesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "PostgresqlDatabase"}

// Get takes name of the postgresqlDatabase, and returns the corresponding postgresqlDatabase object, and an error if there is any.
func (c *FakePostgresqlDatabases) Get(name string, options v1.GetOptions) (result *v1alpha1.PostgresqlDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(postgresqldatabasesResource, c.ns, name), &v1alpha1.PostgresqlDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PostgresqlDatabase), err
}

// List takes label and field selectors, and returns the list of PostgresqlDatabases that match those selectors.
func (c *FakePostgresqlDatabases) List(opts v1.ListOptions) (result *v1alpha1.PostgresqlDatabaseList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(postgresqldatabasesResource, postgresqldatabasesKind, c.ns, opts), &v1alpha1.PostgresqlDatabaseList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PostgresqlDatabaseList{ListMeta: obj.(*v1alpha1.PostgresqlDatabaseList).ListMeta}
	for _, item := range obj.(*v1alpha1.PostgresqlDatabaseList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested postgresqlDatabases.
func (c *FakePostgresqlDatabases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(postgresqldatabasesResource, c.ns, opts))

}

// Create takes the representation of a postgresqlDatabase and creates it.  Returns the server's representation of the postgresqlDatabase, and an error, if there is any.
func (c *FakePostgresqlDatabases) Create(postgresqlDatabase *v1alpha1.PostgresqlDatabase) (result *v1alpha1.PostgresqlDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(postgresqldatabasesResource, c.ns, postgresqlDatabase), &v1alpha1.PostgresqlDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PostgresqlDatabase), err
}

// Update takes the representation of a postgresqlDatabase and updates it. Returns the server's representation of the postgresqlDatabase, and an error, if there is any.
func (c *FakePostgresqlDatabases) Update(postgresqlDatabase *v1alpha1.PostgresqlDatabase) (result *v1alpha1.PostgresqlDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(postgresqldatabasesResource, c.ns, postgresqlDatabase), &v1alpha1.PostgresqlDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PostgresqlDatabase), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePostgresqlDatabases) UpdateStatus(postgresqlDatabase *v1alpha1.PostgresqlDatabase) (*v1alpha1.PostgresqlDatabase, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(postgresqldatabasesResource, "status", c.ns, postgresqlDatabase), &v1alpha1.PostgresqlDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PostgresqlDatabase), err
}

// Delete takes name of the postgresqlDatabase and deletes it. Returns an error if one occurs.
func (c *FakePostgresqlDatabases) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(postgresqldatabasesResource, c.ns, name), &v1alpha1.PostgresqlDatabase{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePostgresqlDatabases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(postgresqldatabasesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PostgresqlDatabaseList{})
	return err
}

// Patch applies the patch and returns the patched postgresqlDatabase.
func (c *FakePostgresqlDatabases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PostgresqlDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(postgresqldatabasesResource, c.ns, name, pt, data, subresources...), &v1alpha1.PostgresqlDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PostgresqlDatabase), err
}
