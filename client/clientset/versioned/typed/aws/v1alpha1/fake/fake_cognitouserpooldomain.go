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

// FakeCognitoUserPoolDomains implements CognitoUserPoolDomainInterface
type FakeCognitoUserPoolDomains struct {
	Fake *FakeAwsV1alpha1
}

var cognitouserpooldomainsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "cognitouserpooldomains"}

var cognitouserpooldomainsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "CognitoUserPoolDomain"}

// Get takes name of the cognitoUserPoolDomain, and returns the corresponding cognitoUserPoolDomain object, and an error if there is any.
func (c *FakeCognitoUserPoolDomains) Get(name string, options v1.GetOptions) (result *v1alpha1.CognitoUserPoolDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(cognitouserpooldomainsResource, name), &v1alpha1.CognitoUserPoolDomain{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CognitoUserPoolDomain), err
}

// List takes label and field selectors, and returns the list of CognitoUserPoolDomains that match those selectors.
func (c *FakeCognitoUserPoolDomains) List(opts v1.ListOptions) (result *v1alpha1.CognitoUserPoolDomainList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(cognitouserpooldomainsResource, cognitouserpooldomainsKind, opts), &v1alpha1.CognitoUserPoolDomainList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CognitoUserPoolDomainList{ListMeta: obj.(*v1alpha1.CognitoUserPoolDomainList).ListMeta}
	for _, item := range obj.(*v1alpha1.CognitoUserPoolDomainList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cognitoUserPoolDomains.
func (c *FakeCognitoUserPoolDomains) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(cognitouserpooldomainsResource, opts))
}

// Create takes the representation of a cognitoUserPoolDomain and creates it.  Returns the server's representation of the cognitoUserPoolDomain, and an error, if there is any.
func (c *FakeCognitoUserPoolDomains) Create(cognitoUserPoolDomain *v1alpha1.CognitoUserPoolDomain) (result *v1alpha1.CognitoUserPoolDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(cognitouserpooldomainsResource, cognitoUserPoolDomain), &v1alpha1.CognitoUserPoolDomain{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CognitoUserPoolDomain), err
}

// Update takes the representation of a cognitoUserPoolDomain and updates it. Returns the server's representation of the cognitoUserPoolDomain, and an error, if there is any.
func (c *FakeCognitoUserPoolDomains) Update(cognitoUserPoolDomain *v1alpha1.CognitoUserPoolDomain) (result *v1alpha1.CognitoUserPoolDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(cognitouserpooldomainsResource, cognitoUserPoolDomain), &v1alpha1.CognitoUserPoolDomain{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CognitoUserPoolDomain), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCognitoUserPoolDomains) UpdateStatus(cognitoUserPoolDomain *v1alpha1.CognitoUserPoolDomain) (*v1alpha1.CognitoUserPoolDomain, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(cognitouserpooldomainsResource, "status", cognitoUserPoolDomain), &v1alpha1.CognitoUserPoolDomain{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CognitoUserPoolDomain), err
}

// Delete takes name of the cognitoUserPoolDomain and deletes it. Returns an error if one occurs.
func (c *FakeCognitoUserPoolDomains) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(cognitouserpooldomainsResource, name), &v1alpha1.CognitoUserPoolDomain{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCognitoUserPoolDomains) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(cognitouserpooldomainsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CognitoUserPoolDomainList{})
	return err
}

// Patch applies the patch and returns the patched cognitoUserPoolDomain.
func (c *FakeCognitoUserPoolDomains) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CognitoUserPoolDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(cognitouserpooldomainsResource, name, pt, data, subresources...), &v1alpha1.CognitoUserPoolDomain{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CognitoUserPoolDomain), err
}
