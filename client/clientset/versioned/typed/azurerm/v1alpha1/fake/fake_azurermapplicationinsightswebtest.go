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

// FakeAzurermApplicationInsightsWebTests implements AzurermApplicationInsightsWebTestInterface
type FakeAzurermApplicationInsightsWebTests struct {
	Fake *FakeAzurermV1alpha1
}

var azurermapplicationinsightswebtestsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermapplicationinsightswebtests"}

var azurermapplicationinsightswebtestsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermApplicationInsightsWebTest"}

// Get takes name of the azurermApplicationInsightsWebTest, and returns the corresponding azurermApplicationInsightsWebTest object, and an error if there is any.
func (c *FakeAzurermApplicationInsightsWebTests) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermApplicationInsightsWebTest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermapplicationinsightswebtestsResource, name), &v1alpha1.AzurermApplicationInsightsWebTest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApplicationInsightsWebTest), err
}

// List takes label and field selectors, and returns the list of AzurermApplicationInsightsWebTests that match those selectors.
func (c *FakeAzurermApplicationInsightsWebTests) List(opts v1.ListOptions) (result *v1alpha1.AzurermApplicationInsightsWebTestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermapplicationinsightswebtestsResource, azurermapplicationinsightswebtestsKind, opts), &v1alpha1.AzurermApplicationInsightsWebTestList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermApplicationInsightsWebTestList{ListMeta: obj.(*v1alpha1.AzurermApplicationInsightsWebTestList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermApplicationInsightsWebTestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermApplicationInsightsWebTests.
func (c *FakeAzurermApplicationInsightsWebTests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermapplicationinsightswebtestsResource, opts))
}

// Create takes the representation of a azurermApplicationInsightsWebTest and creates it.  Returns the server's representation of the azurermApplicationInsightsWebTest, and an error, if there is any.
func (c *FakeAzurermApplicationInsightsWebTests) Create(azurermApplicationInsightsWebTest *v1alpha1.AzurermApplicationInsightsWebTest) (result *v1alpha1.AzurermApplicationInsightsWebTest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermapplicationinsightswebtestsResource, azurermApplicationInsightsWebTest), &v1alpha1.AzurermApplicationInsightsWebTest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApplicationInsightsWebTest), err
}

// Update takes the representation of a azurermApplicationInsightsWebTest and updates it. Returns the server's representation of the azurermApplicationInsightsWebTest, and an error, if there is any.
func (c *FakeAzurermApplicationInsightsWebTests) Update(azurermApplicationInsightsWebTest *v1alpha1.AzurermApplicationInsightsWebTest) (result *v1alpha1.AzurermApplicationInsightsWebTest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermapplicationinsightswebtestsResource, azurermApplicationInsightsWebTest), &v1alpha1.AzurermApplicationInsightsWebTest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApplicationInsightsWebTest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermApplicationInsightsWebTests) UpdateStatus(azurermApplicationInsightsWebTest *v1alpha1.AzurermApplicationInsightsWebTest) (*v1alpha1.AzurermApplicationInsightsWebTest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermapplicationinsightswebtestsResource, "status", azurermApplicationInsightsWebTest), &v1alpha1.AzurermApplicationInsightsWebTest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApplicationInsightsWebTest), err
}

// Delete takes name of the azurermApplicationInsightsWebTest and deletes it. Returns an error if one occurs.
func (c *FakeAzurermApplicationInsightsWebTests) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermapplicationinsightswebtestsResource, name), &v1alpha1.AzurermApplicationInsightsWebTest{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermApplicationInsightsWebTests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermapplicationinsightswebtestsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermApplicationInsightsWebTestList{})
	return err
}

// Patch applies the patch and returns the patched azurermApplicationInsightsWebTest.
func (c *FakeAzurermApplicationInsightsWebTests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermApplicationInsightsWebTest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermapplicationinsightswebtestsResource, name, pt, data, subresources...), &v1alpha1.AzurermApplicationInsightsWebTest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApplicationInsightsWebTest), err
}
