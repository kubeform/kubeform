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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLbListenerCertificates implements LbListenerCertificateInterface
type FakeLbListenerCertificates struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var lblistenercertificatesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "lblistenercertificates"}

var lblistenercertificatesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "LbListenerCertificate"}

// Get takes name of the lbListenerCertificate, and returns the corresponding lbListenerCertificate object, and an error if there is any.
func (c *FakeLbListenerCertificates) Get(name string, options v1.GetOptions) (result *v1alpha1.LbListenerCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(lblistenercertificatesResource, c.ns, name), &v1alpha1.LbListenerCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbListenerCertificate), err
}

// List takes label and field selectors, and returns the list of LbListenerCertificates that match those selectors.
func (c *FakeLbListenerCertificates) List(opts v1.ListOptions) (result *v1alpha1.LbListenerCertificateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(lblistenercertificatesResource, lblistenercertificatesKind, c.ns, opts), &v1alpha1.LbListenerCertificateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LbListenerCertificateList{ListMeta: obj.(*v1alpha1.LbListenerCertificateList).ListMeta}
	for _, item := range obj.(*v1alpha1.LbListenerCertificateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested lbListenerCertificates.
func (c *FakeLbListenerCertificates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(lblistenercertificatesResource, c.ns, opts))

}

// Create takes the representation of a lbListenerCertificate and creates it.  Returns the server's representation of the lbListenerCertificate, and an error, if there is any.
func (c *FakeLbListenerCertificates) Create(lbListenerCertificate *v1alpha1.LbListenerCertificate) (result *v1alpha1.LbListenerCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(lblistenercertificatesResource, c.ns, lbListenerCertificate), &v1alpha1.LbListenerCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbListenerCertificate), err
}

// Update takes the representation of a lbListenerCertificate and updates it. Returns the server's representation of the lbListenerCertificate, and an error, if there is any.
func (c *FakeLbListenerCertificates) Update(lbListenerCertificate *v1alpha1.LbListenerCertificate) (result *v1alpha1.LbListenerCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(lblistenercertificatesResource, c.ns, lbListenerCertificate), &v1alpha1.LbListenerCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbListenerCertificate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLbListenerCertificates) UpdateStatus(lbListenerCertificate *v1alpha1.LbListenerCertificate) (*v1alpha1.LbListenerCertificate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(lblistenercertificatesResource, "status", c.ns, lbListenerCertificate), &v1alpha1.LbListenerCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbListenerCertificate), err
}

// Delete takes name of the lbListenerCertificate and deletes it. Returns an error if one occurs.
func (c *FakeLbListenerCertificates) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(lblistenercertificatesResource, c.ns, name), &v1alpha1.LbListenerCertificate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLbListenerCertificates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(lblistenercertificatesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.LbListenerCertificateList{})
	return err
}

// Patch applies the patch and returns the patched lbListenerCertificate.
func (c *FakeLbListenerCertificates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LbListenerCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(lblistenercertificatesResource, c.ns, name, pt, data, subresources...), &v1alpha1.LbListenerCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbListenerCertificate), err
}
