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

// FakeAzurermApiManagementLoggers implements AzurermApiManagementLoggerInterface
type FakeAzurermApiManagementLoggers struct {
	Fake *FakeAzurermV1alpha1
}

var azurermapimanagementloggersResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermapimanagementloggers"}

var azurermapimanagementloggersKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermApiManagementLogger"}

// Get takes name of the azurermApiManagementLogger, and returns the corresponding azurermApiManagementLogger object, and an error if there is any.
func (c *FakeAzurermApiManagementLoggers) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermApiManagementLogger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermapimanagementloggersResource, name), &v1alpha1.AzurermApiManagementLogger{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApiManagementLogger), err
}

// List takes label and field selectors, and returns the list of AzurermApiManagementLoggers that match those selectors.
func (c *FakeAzurermApiManagementLoggers) List(opts v1.ListOptions) (result *v1alpha1.AzurermApiManagementLoggerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermapimanagementloggersResource, azurermapimanagementloggersKind, opts), &v1alpha1.AzurermApiManagementLoggerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermApiManagementLoggerList{ListMeta: obj.(*v1alpha1.AzurermApiManagementLoggerList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermApiManagementLoggerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermApiManagementLoggers.
func (c *FakeAzurermApiManagementLoggers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermapimanagementloggersResource, opts))
}

// Create takes the representation of a azurermApiManagementLogger and creates it.  Returns the server's representation of the azurermApiManagementLogger, and an error, if there is any.
func (c *FakeAzurermApiManagementLoggers) Create(azurermApiManagementLogger *v1alpha1.AzurermApiManagementLogger) (result *v1alpha1.AzurermApiManagementLogger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermapimanagementloggersResource, azurermApiManagementLogger), &v1alpha1.AzurermApiManagementLogger{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApiManagementLogger), err
}

// Update takes the representation of a azurermApiManagementLogger and updates it. Returns the server's representation of the azurermApiManagementLogger, and an error, if there is any.
func (c *FakeAzurermApiManagementLoggers) Update(azurermApiManagementLogger *v1alpha1.AzurermApiManagementLogger) (result *v1alpha1.AzurermApiManagementLogger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermapimanagementloggersResource, azurermApiManagementLogger), &v1alpha1.AzurermApiManagementLogger{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApiManagementLogger), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermApiManagementLoggers) UpdateStatus(azurermApiManagementLogger *v1alpha1.AzurermApiManagementLogger) (*v1alpha1.AzurermApiManagementLogger, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermapimanagementloggersResource, "status", azurermApiManagementLogger), &v1alpha1.AzurermApiManagementLogger{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApiManagementLogger), err
}

// Delete takes name of the azurermApiManagementLogger and deletes it. Returns an error if one occurs.
func (c *FakeAzurermApiManagementLoggers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermapimanagementloggersResource, name), &v1alpha1.AzurermApiManagementLogger{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermApiManagementLoggers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermapimanagementloggersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermApiManagementLoggerList{})
	return err
}

// Patch applies the patch and returns the patched azurermApiManagementLogger.
func (c *FakeAzurermApiManagementLoggers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermApiManagementLogger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermapimanagementloggersResource, name, pt, data, subresources...), &v1alpha1.AzurermApiManagementLogger{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApiManagementLogger), err
}
