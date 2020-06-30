/*
Copyright The Kubeform Authors.

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

// FakeDataFactoryIntegrationRuntimeManageds implements DataFactoryIntegrationRuntimeManagedInterface
type FakeDataFactoryIntegrationRuntimeManageds struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var datafactoryintegrationruntimemanagedsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "datafactoryintegrationruntimemanageds"}

var datafactoryintegrationruntimemanagedsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "DataFactoryIntegrationRuntimeManaged"}

// Get takes name of the dataFactoryIntegrationRuntimeManaged, and returns the corresponding dataFactoryIntegrationRuntimeManaged object, and an error if there is any.
func (c *FakeDataFactoryIntegrationRuntimeManageds) Get(name string, options v1.GetOptions) (result *v1alpha1.DataFactoryIntegrationRuntimeManaged, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(datafactoryintegrationruntimemanagedsResource, c.ns, name), &v1alpha1.DataFactoryIntegrationRuntimeManaged{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataFactoryIntegrationRuntimeManaged), err
}

// List takes label and field selectors, and returns the list of DataFactoryIntegrationRuntimeManageds that match those selectors.
func (c *FakeDataFactoryIntegrationRuntimeManageds) List(opts v1.ListOptions) (result *v1alpha1.DataFactoryIntegrationRuntimeManagedList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(datafactoryintegrationruntimemanagedsResource, datafactoryintegrationruntimemanagedsKind, c.ns, opts), &v1alpha1.DataFactoryIntegrationRuntimeManagedList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DataFactoryIntegrationRuntimeManagedList{ListMeta: obj.(*v1alpha1.DataFactoryIntegrationRuntimeManagedList).ListMeta}
	for _, item := range obj.(*v1alpha1.DataFactoryIntegrationRuntimeManagedList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dataFactoryIntegrationRuntimeManageds.
func (c *FakeDataFactoryIntegrationRuntimeManageds) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(datafactoryintegrationruntimemanagedsResource, c.ns, opts))

}

// Create takes the representation of a dataFactoryIntegrationRuntimeManaged and creates it.  Returns the server's representation of the dataFactoryIntegrationRuntimeManaged, and an error, if there is any.
func (c *FakeDataFactoryIntegrationRuntimeManageds) Create(dataFactoryIntegrationRuntimeManaged *v1alpha1.DataFactoryIntegrationRuntimeManaged) (result *v1alpha1.DataFactoryIntegrationRuntimeManaged, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(datafactoryintegrationruntimemanagedsResource, c.ns, dataFactoryIntegrationRuntimeManaged), &v1alpha1.DataFactoryIntegrationRuntimeManaged{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataFactoryIntegrationRuntimeManaged), err
}

// Update takes the representation of a dataFactoryIntegrationRuntimeManaged and updates it. Returns the server's representation of the dataFactoryIntegrationRuntimeManaged, and an error, if there is any.
func (c *FakeDataFactoryIntegrationRuntimeManageds) Update(dataFactoryIntegrationRuntimeManaged *v1alpha1.DataFactoryIntegrationRuntimeManaged) (result *v1alpha1.DataFactoryIntegrationRuntimeManaged, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(datafactoryintegrationruntimemanagedsResource, c.ns, dataFactoryIntegrationRuntimeManaged), &v1alpha1.DataFactoryIntegrationRuntimeManaged{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataFactoryIntegrationRuntimeManaged), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDataFactoryIntegrationRuntimeManageds) UpdateStatus(dataFactoryIntegrationRuntimeManaged *v1alpha1.DataFactoryIntegrationRuntimeManaged) (*v1alpha1.DataFactoryIntegrationRuntimeManaged, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(datafactoryintegrationruntimemanagedsResource, "status", c.ns, dataFactoryIntegrationRuntimeManaged), &v1alpha1.DataFactoryIntegrationRuntimeManaged{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataFactoryIntegrationRuntimeManaged), err
}

// Delete takes name of the dataFactoryIntegrationRuntimeManaged and deletes it. Returns an error if one occurs.
func (c *FakeDataFactoryIntegrationRuntimeManageds) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(datafactoryintegrationruntimemanagedsResource, c.ns, name), &v1alpha1.DataFactoryIntegrationRuntimeManaged{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDataFactoryIntegrationRuntimeManageds) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(datafactoryintegrationruntimemanagedsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DataFactoryIntegrationRuntimeManagedList{})
	return err
}

// Patch applies the patch and returns the patched dataFactoryIntegrationRuntimeManaged.
func (c *FakeDataFactoryIntegrationRuntimeManageds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DataFactoryIntegrationRuntimeManaged, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(datafactoryintegrationruntimemanagedsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DataFactoryIntegrationRuntimeManaged{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataFactoryIntegrationRuntimeManaged), err
}
