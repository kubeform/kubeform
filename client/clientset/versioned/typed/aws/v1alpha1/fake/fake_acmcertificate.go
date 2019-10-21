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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeAcmCertificates implements AcmCertificateInterface
type FakeAcmCertificates struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var acmcertificatesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "acmcertificates"}

var acmcertificatesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AcmCertificate"}

// Get takes name of the acmCertificate, and returns the corresponding acmCertificate object, and an error if there is any.
func (c *FakeAcmCertificates) Get(name string, options v1.GetOptions) (result *v1alpha1.AcmCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(acmcertificatesResource, c.ns, name), &v1alpha1.AcmCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AcmCertificate), err
}

// List takes label and field selectors, and returns the list of AcmCertificates that match those selectors.
func (c *FakeAcmCertificates) List(opts v1.ListOptions) (result *v1alpha1.AcmCertificateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(acmcertificatesResource, acmcertificatesKind, c.ns, opts), &v1alpha1.AcmCertificateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AcmCertificateList{ListMeta: obj.(*v1alpha1.AcmCertificateList).ListMeta}
	for _, item := range obj.(*v1alpha1.AcmCertificateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested acmCertificates.
func (c *FakeAcmCertificates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(acmcertificatesResource, c.ns, opts))

}

// Create takes the representation of a acmCertificate and creates it.  Returns the server's representation of the acmCertificate, and an error, if there is any.
func (c *FakeAcmCertificates) Create(acmCertificate *v1alpha1.AcmCertificate) (result *v1alpha1.AcmCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(acmcertificatesResource, c.ns, acmCertificate), &v1alpha1.AcmCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AcmCertificate), err
}

// Update takes the representation of a acmCertificate and updates it. Returns the server's representation of the acmCertificate, and an error, if there is any.
func (c *FakeAcmCertificates) Update(acmCertificate *v1alpha1.AcmCertificate) (result *v1alpha1.AcmCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(acmcertificatesResource, c.ns, acmCertificate), &v1alpha1.AcmCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AcmCertificate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAcmCertificates) UpdateStatus(acmCertificate *v1alpha1.AcmCertificate) (*v1alpha1.AcmCertificate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(acmcertificatesResource, "status", c.ns, acmCertificate), &v1alpha1.AcmCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AcmCertificate), err
}

// Delete takes name of the acmCertificate and deletes it. Returns an error if one occurs.
func (c *FakeAcmCertificates) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(acmcertificatesResource, c.ns, name), &v1alpha1.AcmCertificate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAcmCertificates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(acmcertificatesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AcmCertificateList{})
	return err
}

// Patch applies the patch and returns the patched acmCertificate.
func (c *FakeAcmCertificates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AcmCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(acmcertificatesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AcmCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AcmCertificate), err
}