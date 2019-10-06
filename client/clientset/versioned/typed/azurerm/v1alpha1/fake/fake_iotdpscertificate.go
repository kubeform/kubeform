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

// FakeIotDpsCertificates implements IotDpsCertificateInterface
type FakeIotDpsCertificates struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var iotdpscertificatesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "iotdpscertificates"}

var iotdpscertificatesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "IotDpsCertificate"}

// Get takes name of the iotDpsCertificate, and returns the corresponding iotDpsCertificate object, and an error if there is any.
func (c *FakeIotDpsCertificates) Get(name string, options v1.GetOptions) (result *v1alpha1.IotDpsCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iotdpscertificatesResource, c.ns, name), &v1alpha1.IotDpsCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IotDpsCertificate), err
}

// List takes label and field selectors, and returns the list of IotDpsCertificates that match those selectors.
func (c *FakeIotDpsCertificates) List(opts v1.ListOptions) (result *v1alpha1.IotDpsCertificateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iotdpscertificatesResource, iotdpscertificatesKind, c.ns, opts), &v1alpha1.IotDpsCertificateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IotDpsCertificateList{ListMeta: obj.(*v1alpha1.IotDpsCertificateList).ListMeta}
	for _, item := range obj.(*v1alpha1.IotDpsCertificateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iotDpsCertificates.
func (c *FakeIotDpsCertificates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iotdpscertificatesResource, c.ns, opts))

}

// Create takes the representation of a iotDpsCertificate and creates it.  Returns the server's representation of the iotDpsCertificate, and an error, if there is any.
func (c *FakeIotDpsCertificates) Create(iotDpsCertificate *v1alpha1.IotDpsCertificate) (result *v1alpha1.IotDpsCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iotdpscertificatesResource, c.ns, iotDpsCertificate), &v1alpha1.IotDpsCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IotDpsCertificate), err
}

// Update takes the representation of a iotDpsCertificate and updates it. Returns the server's representation of the iotDpsCertificate, and an error, if there is any.
func (c *FakeIotDpsCertificates) Update(iotDpsCertificate *v1alpha1.IotDpsCertificate) (result *v1alpha1.IotDpsCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iotdpscertificatesResource, c.ns, iotDpsCertificate), &v1alpha1.IotDpsCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IotDpsCertificate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIotDpsCertificates) UpdateStatus(iotDpsCertificate *v1alpha1.IotDpsCertificate) (*v1alpha1.IotDpsCertificate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(iotdpscertificatesResource, "status", c.ns, iotDpsCertificate), &v1alpha1.IotDpsCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IotDpsCertificate), err
}

// Delete takes name of the iotDpsCertificate and deletes it. Returns an error if one occurs.
func (c *FakeIotDpsCertificates) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(iotdpscertificatesResource, c.ns, name), &v1alpha1.IotDpsCertificate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIotDpsCertificates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iotdpscertificatesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.IotDpsCertificateList{})
	return err
}

// Patch applies the patch and returns the patched iotDpsCertificate.
func (c *FakeIotDpsCertificates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IotDpsCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iotdpscertificatesResource, c.ns, name, pt, data, subresources...), &v1alpha1.IotDpsCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IotDpsCertificate), err
}
