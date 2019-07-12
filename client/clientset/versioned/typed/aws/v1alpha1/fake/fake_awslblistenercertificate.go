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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeAwsLbListenerCertificates implements AwsLbListenerCertificateInterface
type FakeAwsLbListenerCertificates struct {
	Fake *FakeAwsV1alpha1
}

var awslblistenercertificatesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awslblistenercertificates"}

var awslblistenercertificatesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsLbListenerCertificate"}

// Get takes name of the awsLbListenerCertificate, and returns the corresponding awsLbListenerCertificate object, and an error if there is any.
func (c *FakeAwsLbListenerCertificates) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsLbListenerCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awslblistenercertificatesResource, name), &v1alpha1.AwsLbListenerCertificate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLbListenerCertificate), err
}

// List takes label and field selectors, and returns the list of AwsLbListenerCertificates that match those selectors.
func (c *FakeAwsLbListenerCertificates) List(opts v1.ListOptions) (result *v1alpha1.AwsLbListenerCertificateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awslblistenercertificatesResource, awslblistenercertificatesKind, opts), &v1alpha1.AwsLbListenerCertificateList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsLbListenerCertificateList{ListMeta: obj.(*v1alpha1.AwsLbListenerCertificateList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsLbListenerCertificateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsLbListenerCertificates.
func (c *FakeAwsLbListenerCertificates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awslblistenercertificatesResource, opts))
}

// Create takes the representation of a awsLbListenerCertificate and creates it.  Returns the server's representation of the awsLbListenerCertificate, and an error, if there is any.
func (c *FakeAwsLbListenerCertificates) Create(awsLbListenerCertificate *v1alpha1.AwsLbListenerCertificate) (result *v1alpha1.AwsLbListenerCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awslblistenercertificatesResource, awsLbListenerCertificate), &v1alpha1.AwsLbListenerCertificate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLbListenerCertificate), err
}

// Update takes the representation of a awsLbListenerCertificate and updates it. Returns the server's representation of the awsLbListenerCertificate, and an error, if there is any.
func (c *FakeAwsLbListenerCertificates) Update(awsLbListenerCertificate *v1alpha1.AwsLbListenerCertificate) (result *v1alpha1.AwsLbListenerCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awslblistenercertificatesResource, awsLbListenerCertificate), &v1alpha1.AwsLbListenerCertificate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLbListenerCertificate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsLbListenerCertificates) UpdateStatus(awsLbListenerCertificate *v1alpha1.AwsLbListenerCertificate) (*v1alpha1.AwsLbListenerCertificate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awslblistenercertificatesResource, "status", awsLbListenerCertificate), &v1alpha1.AwsLbListenerCertificate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLbListenerCertificate), err
}

// Delete takes name of the awsLbListenerCertificate and deletes it. Returns an error if one occurs.
func (c *FakeAwsLbListenerCertificates) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awslblistenercertificatesResource, name), &v1alpha1.AwsLbListenerCertificate{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsLbListenerCertificates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awslblistenercertificatesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsLbListenerCertificateList{})
	return err
}

// Patch applies the patch and returns the patched awsLbListenerCertificate.
func (c *FakeAwsLbListenerCertificates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsLbListenerCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awslblistenercertificatesResource, name, pt, data, subresources...), &v1alpha1.AwsLbListenerCertificate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLbListenerCertificate), err
}
