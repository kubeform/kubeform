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

// FakeAzurermMssqlElasticpools implements AzurermMssqlElasticpoolInterface
type FakeAzurermMssqlElasticpools struct {
	Fake *FakeAzurermV1alpha1
}

var azurermmssqlelasticpoolsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermmssqlelasticpools"}

var azurermmssqlelasticpoolsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermMssqlElasticpool"}

// Get takes name of the azurermMssqlElasticpool, and returns the corresponding azurermMssqlElasticpool object, and an error if there is any.
func (c *FakeAzurermMssqlElasticpools) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermMssqlElasticpool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermmssqlelasticpoolsResource, name), &v1alpha1.AzurermMssqlElasticpool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMssqlElasticpool), err
}

// List takes label and field selectors, and returns the list of AzurermMssqlElasticpools that match those selectors.
func (c *FakeAzurermMssqlElasticpools) List(opts v1.ListOptions) (result *v1alpha1.AzurermMssqlElasticpoolList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermmssqlelasticpoolsResource, azurermmssqlelasticpoolsKind, opts), &v1alpha1.AzurermMssqlElasticpoolList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermMssqlElasticpoolList{ListMeta: obj.(*v1alpha1.AzurermMssqlElasticpoolList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermMssqlElasticpoolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermMssqlElasticpools.
func (c *FakeAzurermMssqlElasticpools) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermmssqlelasticpoolsResource, opts))
}

// Create takes the representation of a azurermMssqlElasticpool and creates it.  Returns the server's representation of the azurermMssqlElasticpool, and an error, if there is any.
func (c *FakeAzurermMssqlElasticpools) Create(azurermMssqlElasticpool *v1alpha1.AzurermMssqlElasticpool) (result *v1alpha1.AzurermMssqlElasticpool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermmssqlelasticpoolsResource, azurermMssqlElasticpool), &v1alpha1.AzurermMssqlElasticpool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMssqlElasticpool), err
}

// Update takes the representation of a azurermMssqlElasticpool and updates it. Returns the server's representation of the azurermMssqlElasticpool, and an error, if there is any.
func (c *FakeAzurermMssqlElasticpools) Update(azurermMssqlElasticpool *v1alpha1.AzurermMssqlElasticpool) (result *v1alpha1.AzurermMssqlElasticpool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermmssqlelasticpoolsResource, azurermMssqlElasticpool), &v1alpha1.AzurermMssqlElasticpool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMssqlElasticpool), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermMssqlElasticpools) UpdateStatus(azurermMssqlElasticpool *v1alpha1.AzurermMssqlElasticpool) (*v1alpha1.AzurermMssqlElasticpool, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermmssqlelasticpoolsResource, "status", azurermMssqlElasticpool), &v1alpha1.AzurermMssqlElasticpool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMssqlElasticpool), err
}

// Delete takes name of the azurermMssqlElasticpool and deletes it. Returns an error if one occurs.
func (c *FakeAzurermMssqlElasticpools) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermmssqlelasticpoolsResource, name), &v1alpha1.AzurermMssqlElasticpool{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermMssqlElasticpools) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermmssqlelasticpoolsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermMssqlElasticpoolList{})
	return err
}

// Patch applies the patch and returns the patched azurermMssqlElasticpool.
func (c *FakeAzurermMssqlElasticpools) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermMssqlElasticpool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermmssqlelasticpoolsResource, name, pt, data, subresources...), &v1alpha1.AzurermMssqlElasticpool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMssqlElasticpool), err
}
