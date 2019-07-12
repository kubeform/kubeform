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

// FakeAzurermBatchCertificates implements AzurermBatchCertificateInterface
type FakeAzurermBatchCertificates struct {
	Fake *FakeAzurermV1alpha1
}

var azurermbatchcertificatesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermbatchcertificates"}

var azurermbatchcertificatesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermBatchCertificate"}

// Get takes name of the azurermBatchCertificate, and returns the corresponding azurermBatchCertificate object, and an error if there is any.
func (c *FakeAzurermBatchCertificates) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermBatchCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermbatchcertificatesResource, name), &v1alpha1.AzurermBatchCertificate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermBatchCertificate), err
}

// List takes label and field selectors, and returns the list of AzurermBatchCertificates that match those selectors.
func (c *FakeAzurermBatchCertificates) List(opts v1.ListOptions) (result *v1alpha1.AzurermBatchCertificateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermbatchcertificatesResource, azurermbatchcertificatesKind, opts), &v1alpha1.AzurermBatchCertificateList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermBatchCertificateList{ListMeta: obj.(*v1alpha1.AzurermBatchCertificateList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermBatchCertificateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermBatchCertificates.
func (c *FakeAzurermBatchCertificates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermbatchcertificatesResource, opts))
}

// Create takes the representation of a azurermBatchCertificate and creates it.  Returns the server's representation of the azurermBatchCertificate, and an error, if there is any.
func (c *FakeAzurermBatchCertificates) Create(azurermBatchCertificate *v1alpha1.AzurermBatchCertificate) (result *v1alpha1.AzurermBatchCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermbatchcertificatesResource, azurermBatchCertificate), &v1alpha1.AzurermBatchCertificate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermBatchCertificate), err
}

// Update takes the representation of a azurermBatchCertificate and updates it. Returns the server's representation of the azurermBatchCertificate, and an error, if there is any.
func (c *FakeAzurermBatchCertificates) Update(azurermBatchCertificate *v1alpha1.AzurermBatchCertificate) (result *v1alpha1.AzurermBatchCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermbatchcertificatesResource, azurermBatchCertificate), &v1alpha1.AzurermBatchCertificate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermBatchCertificate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermBatchCertificates) UpdateStatus(azurermBatchCertificate *v1alpha1.AzurermBatchCertificate) (*v1alpha1.AzurermBatchCertificate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermbatchcertificatesResource, "status", azurermBatchCertificate), &v1alpha1.AzurermBatchCertificate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermBatchCertificate), err
}

// Delete takes name of the azurermBatchCertificate and deletes it. Returns an error if one occurs.
func (c *FakeAzurermBatchCertificates) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermbatchcertificatesResource, name), &v1alpha1.AzurermBatchCertificate{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermBatchCertificates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermbatchcertificatesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermBatchCertificateList{})
	return err
}

// Patch applies the patch and returns the patched azurermBatchCertificate.
func (c *FakeAzurermBatchCertificates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermBatchCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermbatchcertificatesResource, name, pt, data, subresources...), &v1alpha1.AzurermBatchCertificate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermBatchCertificate), err
}
