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

// FakeLogAnalyticsLinkedServices implements LogAnalyticsLinkedServiceInterface
type FakeLogAnalyticsLinkedServices struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var loganalyticslinkedservicesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "loganalyticslinkedservices"}

var loganalyticslinkedservicesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "LogAnalyticsLinkedService"}

// Get takes name of the logAnalyticsLinkedService, and returns the corresponding logAnalyticsLinkedService object, and an error if there is any.
func (c *FakeLogAnalyticsLinkedServices) Get(name string, options v1.GetOptions) (result *v1alpha1.LogAnalyticsLinkedService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(loganalyticslinkedservicesResource, c.ns, name), &v1alpha1.LogAnalyticsLinkedService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogAnalyticsLinkedService), err
}

// List takes label and field selectors, and returns the list of LogAnalyticsLinkedServices that match those selectors.
func (c *FakeLogAnalyticsLinkedServices) List(opts v1.ListOptions) (result *v1alpha1.LogAnalyticsLinkedServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(loganalyticslinkedservicesResource, loganalyticslinkedservicesKind, c.ns, opts), &v1alpha1.LogAnalyticsLinkedServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LogAnalyticsLinkedServiceList{ListMeta: obj.(*v1alpha1.LogAnalyticsLinkedServiceList).ListMeta}
	for _, item := range obj.(*v1alpha1.LogAnalyticsLinkedServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested logAnalyticsLinkedServices.
func (c *FakeLogAnalyticsLinkedServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(loganalyticslinkedservicesResource, c.ns, opts))

}

// Create takes the representation of a logAnalyticsLinkedService and creates it.  Returns the server's representation of the logAnalyticsLinkedService, and an error, if there is any.
func (c *FakeLogAnalyticsLinkedServices) Create(logAnalyticsLinkedService *v1alpha1.LogAnalyticsLinkedService) (result *v1alpha1.LogAnalyticsLinkedService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(loganalyticslinkedservicesResource, c.ns, logAnalyticsLinkedService), &v1alpha1.LogAnalyticsLinkedService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogAnalyticsLinkedService), err
}

// Update takes the representation of a logAnalyticsLinkedService and updates it. Returns the server's representation of the logAnalyticsLinkedService, and an error, if there is any.
func (c *FakeLogAnalyticsLinkedServices) Update(logAnalyticsLinkedService *v1alpha1.LogAnalyticsLinkedService) (result *v1alpha1.LogAnalyticsLinkedService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(loganalyticslinkedservicesResource, c.ns, logAnalyticsLinkedService), &v1alpha1.LogAnalyticsLinkedService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogAnalyticsLinkedService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLogAnalyticsLinkedServices) UpdateStatus(logAnalyticsLinkedService *v1alpha1.LogAnalyticsLinkedService) (*v1alpha1.LogAnalyticsLinkedService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(loganalyticslinkedservicesResource, "status", c.ns, logAnalyticsLinkedService), &v1alpha1.LogAnalyticsLinkedService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogAnalyticsLinkedService), err
}

// Delete takes name of the logAnalyticsLinkedService and deletes it. Returns an error if one occurs.
func (c *FakeLogAnalyticsLinkedServices) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(loganalyticslinkedservicesResource, c.ns, name), &v1alpha1.LogAnalyticsLinkedService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLogAnalyticsLinkedServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(loganalyticslinkedservicesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.LogAnalyticsLinkedServiceList{})
	return err
}

// Patch applies the patch and returns the patched logAnalyticsLinkedService.
func (c *FakeLogAnalyticsLinkedServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LogAnalyticsLinkedService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(loganalyticslinkedservicesResource, c.ns, name, pt, data, subresources...), &v1alpha1.LogAnalyticsLinkedService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogAnalyticsLinkedService), err
}
