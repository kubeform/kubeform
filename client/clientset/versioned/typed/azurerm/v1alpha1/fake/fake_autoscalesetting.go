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

// FakeAutoscaleSettings implements AutoscaleSettingInterface
type FakeAutoscaleSettings struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var autoscalesettingsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "autoscalesettings"}

var autoscalesettingsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AutoscaleSetting"}

// Get takes name of the autoscaleSetting, and returns the corresponding autoscaleSetting object, and an error if there is any.
func (c *FakeAutoscaleSettings) Get(name string, options v1.GetOptions) (result *v1alpha1.AutoscaleSetting, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(autoscalesettingsResource, c.ns, name), &v1alpha1.AutoscaleSetting{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoscaleSetting), err
}

// List takes label and field selectors, and returns the list of AutoscaleSettings that match those selectors.
func (c *FakeAutoscaleSettings) List(opts v1.ListOptions) (result *v1alpha1.AutoscaleSettingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(autoscalesettingsResource, autoscalesettingsKind, c.ns, opts), &v1alpha1.AutoscaleSettingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AutoscaleSettingList{ListMeta: obj.(*v1alpha1.AutoscaleSettingList).ListMeta}
	for _, item := range obj.(*v1alpha1.AutoscaleSettingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested autoscaleSettings.
func (c *FakeAutoscaleSettings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(autoscalesettingsResource, c.ns, opts))

}

// Create takes the representation of a autoscaleSetting and creates it.  Returns the server's representation of the autoscaleSetting, and an error, if there is any.
func (c *FakeAutoscaleSettings) Create(autoscaleSetting *v1alpha1.AutoscaleSetting) (result *v1alpha1.AutoscaleSetting, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(autoscalesettingsResource, c.ns, autoscaleSetting), &v1alpha1.AutoscaleSetting{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoscaleSetting), err
}

// Update takes the representation of a autoscaleSetting and updates it. Returns the server's representation of the autoscaleSetting, and an error, if there is any.
func (c *FakeAutoscaleSettings) Update(autoscaleSetting *v1alpha1.AutoscaleSetting) (result *v1alpha1.AutoscaleSetting, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(autoscalesettingsResource, c.ns, autoscaleSetting), &v1alpha1.AutoscaleSetting{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoscaleSetting), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAutoscaleSettings) UpdateStatus(autoscaleSetting *v1alpha1.AutoscaleSetting) (*v1alpha1.AutoscaleSetting, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(autoscalesettingsResource, "status", c.ns, autoscaleSetting), &v1alpha1.AutoscaleSetting{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoscaleSetting), err
}

// Delete takes name of the autoscaleSetting and deletes it. Returns an error if one occurs.
func (c *FakeAutoscaleSettings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(autoscalesettingsResource, c.ns, name), &v1alpha1.AutoscaleSetting{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAutoscaleSettings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(autoscalesettingsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AutoscaleSettingList{})
	return err
}

// Patch applies the patch and returns the patched autoscaleSetting.
func (c *FakeAutoscaleSettings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AutoscaleSetting, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(autoscalesettingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AutoscaleSetting{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoscaleSetting), err
}
