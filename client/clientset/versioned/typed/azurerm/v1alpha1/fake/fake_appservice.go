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

// FakeAppServices implements AppServiceInterface
type FakeAppServices struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var appservicesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "appservices"}

var appservicesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AppService"}

// Get takes name of the appService, and returns the corresponding appService object, and an error if there is any.
func (c *FakeAppServices) Get(name string, options v1.GetOptions) (result *v1alpha1.AppService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(appservicesResource, c.ns, name), &v1alpha1.AppService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppService), err
}

// List takes label and field selectors, and returns the list of AppServices that match those selectors.
func (c *FakeAppServices) List(opts v1.ListOptions) (result *v1alpha1.AppServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(appservicesResource, appservicesKind, c.ns, opts), &v1alpha1.AppServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AppServiceList{ListMeta: obj.(*v1alpha1.AppServiceList).ListMeta}
	for _, item := range obj.(*v1alpha1.AppServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested appServices.
func (c *FakeAppServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(appservicesResource, c.ns, opts))

}

// Create takes the representation of a appService and creates it.  Returns the server's representation of the appService, and an error, if there is any.
func (c *FakeAppServices) Create(appService *v1alpha1.AppService) (result *v1alpha1.AppService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(appservicesResource, c.ns, appService), &v1alpha1.AppService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppService), err
}

// Update takes the representation of a appService and updates it. Returns the server's representation of the appService, and an error, if there is any.
func (c *FakeAppServices) Update(appService *v1alpha1.AppService) (result *v1alpha1.AppService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(appservicesResource, c.ns, appService), &v1alpha1.AppService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAppServices) UpdateStatus(appService *v1alpha1.AppService) (*v1alpha1.AppService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(appservicesResource, "status", c.ns, appService), &v1alpha1.AppService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppService), err
}

// Delete takes name of the appService and deletes it. Returns an error if one occurs.
func (c *FakeAppServices) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(appservicesResource, c.ns, name), &v1alpha1.AppService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAppServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(appservicesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AppServiceList{})
	return err
}

// Patch applies the patch and returns the patched appService.
func (c *FakeAppServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AppService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(appservicesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AppService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppService), err
}
