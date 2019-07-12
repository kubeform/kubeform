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

// FakeAzurermDataFactoryLinkedServiceMysqls implements AzurermDataFactoryLinkedServiceMysqlInterface
type FakeAzurermDataFactoryLinkedServiceMysqls struct {
	Fake *FakeAzurermV1alpha1
}

var azurermdatafactorylinkedservicemysqlsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermdatafactorylinkedservicemysqls"}

var azurermdatafactorylinkedservicemysqlsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermDataFactoryLinkedServiceMysql"}

// Get takes name of the azurermDataFactoryLinkedServiceMysql, and returns the corresponding azurermDataFactoryLinkedServiceMysql object, and an error if there is any.
func (c *FakeAzurermDataFactoryLinkedServiceMysqls) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermDataFactoryLinkedServiceMysql, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermdatafactorylinkedservicemysqlsResource, name), &v1alpha1.AzurermDataFactoryLinkedServiceMysql{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDataFactoryLinkedServiceMysql), err
}

// List takes label and field selectors, and returns the list of AzurermDataFactoryLinkedServiceMysqls that match those selectors.
func (c *FakeAzurermDataFactoryLinkedServiceMysqls) List(opts v1.ListOptions) (result *v1alpha1.AzurermDataFactoryLinkedServiceMysqlList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermdatafactorylinkedservicemysqlsResource, azurermdatafactorylinkedservicemysqlsKind, opts), &v1alpha1.AzurermDataFactoryLinkedServiceMysqlList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermDataFactoryLinkedServiceMysqlList{ListMeta: obj.(*v1alpha1.AzurermDataFactoryLinkedServiceMysqlList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermDataFactoryLinkedServiceMysqlList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermDataFactoryLinkedServiceMysqls.
func (c *FakeAzurermDataFactoryLinkedServiceMysqls) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermdatafactorylinkedservicemysqlsResource, opts))
}

// Create takes the representation of a azurermDataFactoryLinkedServiceMysql and creates it.  Returns the server's representation of the azurermDataFactoryLinkedServiceMysql, and an error, if there is any.
func (c *FakeAzurermDataFactoryLinkedServiceMysqls) Create(azurermDataFactoryLinkedServiceMysql *v1alpha1.AzurermDataFactoryLinkedServiceMysql) (result *v1alpha1.AzurermDataFactoryLinkedServiceMysql, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermdatafactorylinkedservicemysqlsResource, azurermDataFactoryLinkedServiceMysql), &v1alpha1.AzurermDataFactoryLinkedServiceMysql{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDataFactoryLinkedServiceMysql), err
}

// Update takes the representation of a azurermDataFactoryLinkedServiceMysql and updates it. Returns the server's representation of the azurermDataFactoryLinkedServiceMysql, and an error, if there is any.
func (c *FakeAzurermDataFactoryLinkedServiceMysqls) Update(azurermDataFactoryLinkedServiceMysql *v1alpha1.AzurermDataFactoryLinkedServiceMysql) (result *v1alpha1.AzurermDataFactoryLinkedServiceMysql, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermdatafactorylinkedservicemysqlsResource, azurermDataFactoryLinkedServiceMysql), &v1alpha1.AzurermDataFactoryLinkedServiceMysql{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDataFactoryLinkedServiceMysql), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermDataFactoryLinkedServiceMysqls) UpdateStatus(azurermDataFactoryLinkedServiceMysql *v1alpha1.AzurermDataFactoryLinkedServiceMysql) (*v1alpha1.AzurermDataFactoryLinkedServiceMysql, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermdatafactorylinkedservicemysqlsResource, "status", azurermDataFactoryLinkedServiceMysql), &v1alpha1.AzurermDataFactoryLinkedServiceMysql{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDataFactoryLinkedServiceMysql), err
}

// Delete takes name of the azurermDataFactoryLinkedServiceMysql and deletes it. Returns an error if one occurs.
func (c *FakeAzurermDataFactoryLinkedServiceMysqls) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermdatafactorylinkedservicemysqlsResource, name), &v1alpha1.AzurermDataFactoryLinkedServiceMysql{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermDataFactoryLinkedServiceMysqls) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermdatafactorylinkedservicemysqlsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermDataFactoryLinkedServiceMysqlList{})
	return err
}

// Patch applies the patch and returns the patched azurermDataFactoryLinkedServiceMysql.
func (c *FakeAzurermDataFactoryLinkedServiceMysqls) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermDataFactoryLinkedServiceMysql, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermdatafactorylinkedservicemysqlsResource, name, pt, data, subresources...), &v1alpha1.AzurermDataFactoryLinkedServiceMysql{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDataFactoryLinkedServiceMysql), err
}
