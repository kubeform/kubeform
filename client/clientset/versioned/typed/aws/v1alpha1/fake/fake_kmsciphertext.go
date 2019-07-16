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

// FakeKmsCiphertexts implements KmsCiphertextInterface
type FakeKmsCiphertexts struct {
	Fake *FakeAwsV1alpha1
}

var kmsciphertextsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "kmsciphertexts"}

var kmsciphertextsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "KmsCiphertext"}

// Get takes name of the kmsCiphertext, and returns the corresponding kmsCiphertext object, and an error if there is any.
func (c *FakeKmsCiphertexts) Get(name string, options v1.GetOptions) (result *v1alpha1.KmsCiphertext, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(kmsciphertextsResource, name), &v1alpha1.KmsCiphertext{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KmsCiphertext), err
}

// List takes label and field selectors, and returns the list of KmsCiphertexts that match those selectors.
func (c *FakeKmsCiphertexts) List(opts v1.ListOptions) (result *v1alpha1.KmsCiphertextList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(kmsciphertextsResource, kmsciphertextsKind, opts), &v1alpha1.KmsCiphertextList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KmsCiphertextList{ListMeta: obj.(*v1alpha1.KmsCiphertextList).ListMeta}
	for _, item := range obj.(*v1alpha1.KmsCiphertextList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kmsCiphertexts.
func (c *FakeKmsCiphertexts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(kmsciphertextsResource, opts))
}

// Create takes the representation of a kmsCiphertext and creates it.  Returns the server's representation of the kmsCiphertext, and an error, if there is any.
func (c *FakeKmsCiphertexts) Create(kmsCiphertext *v1alpha1.KmsCiphertext) (result *v1alpha1.KmsCiphertext, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(kmsciphertextsResource, kmsCiphertext), &v1alpha1.KmsCiphertext{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KmsCiphertext), err
}

// Update takes the representation of a kmsCiphertext and updates it. Returns the server's representation of the kmsCiphertext, and an error, if there is any.
func (c *FakeKmsCiphertexts) Update(kmsCiphertext *v1alpha1.KmsCiphertext) (result *v1alpha1.KmsCiphertext, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(kmsciphertextsResource, kmsCiphertext), &v1alpha1.KmsCiphertext{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KmsCiphertext), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKmsCiphertexts) UpdateStatus(kmsCiphertext *v1alpha1.KmsCiphertext) (*v1alpha1.KmsCiphertext, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(kmsciphertextsResource, "status", kmsCiphertext), &v1alpha1.KmsCiphertext{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KmsCiphertext), err
}

// Delete takes name of the kmsCiphertext and deletes it. Returns an error if one occurs.
func (c *FakeKmsCiphertexts) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(kmsciphertextsResource, name), &v1alpha1.KmsCiphertext{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKmsCiphertexts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(kmsciphertextsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.KmsCiphertextList{})
	return err
}

// Patch applies the patch and returns the patched kmsCiphertext.
func (c *FakeKmsCiphertexts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KmsCiphertext, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(kmsciphertextsResource, name, pt, data, subresources...), &v1alpha1.KmsCiphertext{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KmsCiphertext), err
}