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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeDataFactoryLinkedServiceMysqls implements DataFactoryLinkedServiceMysqlInterface
type FakeDataFactoryLinkedServiceMysqls struct {
	Fake *FakeAzurermV1alpha1
}

var datafactorylinkedservicemysqlsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "datafactorylinkedservicemysqls"}

var datafactorylinkedservicemysqlsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "DataFactoryLinkedServiceMysql"}

// Get takes name of the dataFactoryLinkedServiceMysql, and returns the corresponding dataFactoryLinkedServiceMysql object, and an error if there is any.
func (c *FakeDataFactoryLinkedServiceMysqls) Get(name string, options v1.GetOptions) (result *v1alpha1.DataFactoryLinkedServiceMysql, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(datafactorylinkedservicemysqlsResource, name), &v1alpha1.DataFactoryLinkedServiceMysql{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataFactoryLinkedServiceMysql), err
}

// List takes label and field selectors, and returns the list of DataFactoryLinkedServiceMysqls that match those selectors.
func (c *FakeDataFactoryLinkedServiceMysqls) List(opts v1.ListOptions) (result *v1alpha1.DataFactoryLinkedServiceMysqlList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(datafactorylinkedservicemysqlsResource, datafactorylinkedservicemysqlsKind, opts), &v1alpha1.DataFactoryLinkedServiceMysqlList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DataFactoryLinkedServiceMysqlList{ListMeta: obj.(*v1alpha1.DataFactoryLinkedServiceMysqlList).ListMeta}
	for _, item := range obj.(*v1alpha1.DataFactoryLinkedServiceMysqlList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dataFactoryLinkedServiceMysqls.
func (c *FakeDataFactoryLinkedServiceMysqls) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(datafactorylinkedservicemysqlsResource, opts))
}

// Create takes the representation of a dataFactoryLinkedServiceMysql and creates it.  Returns the server's representation of the dataFactoryLinkedServiceMysql, and an error, if there is any.
func (c *FakeDataFactoryLinkedServiceMysqls) Create(dataFactoryLinkedServiceMysql *v1alpha1.DataFactoryLinkedServiceMysql) (result *v1alpha1.DataFactoryLinkedServiceMysql, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(datafactorylinkedservicemysqlsResource, dataFactoryLinkedServiceMysql), &v1alpha1.DataFactoryLinkedServiceMysql{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataFactoryLinkedServiceMysql), err
}

// Update takes the representation of a dataFactoryLinkedServiceMysql and updates it. Returns the server's representation of the dataFactoryLinkedServiceMysql, and an error, if there is any.
func (c *FakeDataFactoryLinkedServiceMysqls) Update(dataFactoryLinkedServiceMysql *v1alpha1.DataFactoryLinkedServiceMysql) (result *v1alpha1.DataFactoryLinkedServiceMysql, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(datafactorylinkedservicemysqlsResource, dataFactoryLinkedServiceMysql), &v1alpha1.DataFactoryLinkedServiceMysql{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataFactoryLinkedServiceMysql), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDataFactoryLinkedServiceMysqls) UpdateStatus(dataFactoryLinkedServiceMysql *v1alpha1.DataFactoryLinkedServiceMysql) (*v1alpha1.DataFactoryLinkedServiceMysql, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(datafactorylinkedservicemysqlsResource, "status", dataFactoryLinkedServiceMysql), &v1alpha1.DataFactoryLinkedServiceMysql{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataFactoryLinkedServiceMysql), err
}

// Delete takes name of the dataFactoryLinkedServiceMysql and deletes it. Returns an error if one occurs.
func (c *FakeDataFactoryLinkedServiceMysqls) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(datafactorylinkedservicemysqlsResource, name), &v1alpha1.DataFactoryLinkedServiceMysql{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDataFactoryLinkedServiceMysqls) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(datafactorylinkedservicemysqlsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DataFactoryLinkedServiceMysqlList{})
	return err
}

// Patch applies the patch and returns the patched dataFactoryLinkedServiceMysql.
func (c *FakeDataFactoryLinkedServiceMysqls) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DataFactoryLinkedServiceMysql, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(datafactorylinkedservicemysqlsResource, name, pt, data, subresources...), &v1alpha1.DataFactoryLinkedServiceMysql{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataFactoryLinkedServiceMysql), err
}
