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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeAwsGlueCatalogDatabases implements AwsGlueCatalogDatabaseInterface
type FakeAwsGlueCatalogDatabases struct {
	Fake *FakeAwsV1alpha1
}

var awsgluecatalogdatabasesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsgluecatalogdatabases"}

var awsgluecatalogdatabasesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsGlueCatalogDatabase"}

// Get takes name of the awsGlueCatalogDatabase, and returns the corresponding awsGlueCatalogDatabase object, and an error if there is any.
func (c *FakeAwsGlueCatalogDatabases) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsGlueCatalogDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsgluecatalogdatabasesResource, name), &v1alpha1.AwsGlueCatalogDatabase{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGlueCatalogDatabase), err
}

// List takes label and field selectors, and returns the list of AwsGlueCatalogDatabases that match those selectors.
func (c *FakeAwsGlueCatalogDatabases) List(opts v1.ListOptions) (result *v1alpha1.AwsGlueCatalogDatabaseList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsgluecatalogdatabasesResource, awsgluecatalogdatabasesKind, opts), &v1alpha1.AwsGlueCatalogDatabaseList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsGlueCatalogDatabaseList{ListMeta: obj.(*v1alpha1.AwsGlueCatalogDatabaseList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsGlueCatalogDatabaseList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsGlueCatalogDatabases.
func (c *FakeAwsGlueCatalogDatabases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsgluecatalogdatabasesResource, opts))
}

// Create takes the representation of a awsGlueCatalogDatabase and creates it.  Returns the server's representation of the awsGlueCatalogDatabase, and an error, if there is any.
func (c *FakeAwsGlueCatalogDatabases) Create(awsGlueCatalogDatabase *v1alpha1.AwsGlueCatalogDatabase) (result *v1alpha1.AwsGlueCatalogDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsgluecatalogdatabasesResource, awsGlueCatalogDatabase), &v1alpha1.AwsGlueCatalogDatabase{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGlueCatalogDatabase), err
}

// Update takes the representation of a awsGlueCatalogDatabase and updates it. Returns the server's representation of the awsGlueCatalogDatabase, and an error, if there is any.
func (c *FakeAwsGlueCatalogDatabases) Update(awsGlueCatalogDatabase *v1alpha1.AwsGlueCatalogDatabase) (result *v1alpha1.AwsGlueCatalogDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsgluecatalogdatabasesResource, awsGlueCatalogDatabase), &v1alpha1.AwsGlueCatalogDatabase{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGlueCatalogDatabase), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsGlueCatalogDatabases) UpdateStatus(awsGlueCatalogDatabase *v1alpha1.AwsGlueCatalogDatabase) (*v1alpha1.AwsGlueCatalogDatabase, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsgluecatalogdatabasesResource, "status", awsGlueCatalogDatabase), &v1alpha1.AwsGlueCatalogDatabase{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGlueCatalogDatabase), err
}

// Delete takes name of the awsGlueCatalogDatabase and deletes it. Returns an error if one occurs.
func (c *FakeAwsGlueCatalogDatabases) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsgluecatalogdatabasesResource, name), &v1alpha1.AwsGlueCatalogDatabase{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsGlueCatalogDatabases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsgluecatalogdatabasesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsGlueCatalogDatabaseList{})
	return err
}

// Patch applies the patch and returns the patched awsGlueCatalogDatabase.
func (c *FakeAwsGlueCatalogDatabases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsGlueCatalogDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsgluecatalogdatabasesResource, name, pt, data, subresources...), &v1alpha1.AwsGlueCatalogDatabase{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGlueCatalogDatabase), err
}
