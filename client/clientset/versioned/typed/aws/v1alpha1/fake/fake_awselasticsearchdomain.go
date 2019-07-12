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

// FakeAwsElasticsearchDomains implements AwsElasticsearchDomainInterface
type FakeAwsElasticsearchDomains struct {
	Fake *FakeAwsV1alpha1
}

var awselasticsearchdomainsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awselasticsearchdomains"}

var awselasticsearchdomainsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsElasticsearchDomain"}

// Get takes name of the awsElasticsearchDomain, and returns the corresponding awsElasticsearchDomain object, and an error if there is any.
func (c *FakeAwsElasticsearchDomains) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsElasticsearchDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awselasticsearchdomainsResource, name), &v1alpha1.AwsElasticsearchDomain{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElasticsearchDomain), err
}

// List takes label and field selectors, and returns the list of AwsElasticsearchDomains that match those selectors.
func (c *FakeAwsElasticsearchDomains) List(opts v1.ListOptions) (result *v1alpha1.AwsElasticsearchDomainList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awselasticsearchdomainsResource, awselasticsearchdomainsKind, opts), &v1alpha1.AwsElasticsearchDomainList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsElasticsearchDomainList{ListMeta: obj.(*v1alpha1.AwsElasticsearchDomainList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsElasticsearchDomainList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsElasticsearchDomains.
func (c *FakeAwsElasticsearchDomains) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awselasticsearchdomainsResource, opts))
}

// Create takes the representation of a awsElasticsearchDomain and creates it.  Returns the server's representation of the awsElasticsearchDomain, and an error, if there is any.
func (c *FakeAwsElasticsearchDomains) Create(awsElasticsearchDomain *v1alpha1.AwsElasticsearchDomain) (result *v1alpha1.AwsElasticsearchDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awselasticsearchdomainsResource, awsElasticsearchDomain), &v1alpha1.AwsElasticsearchDomain{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElasticsearchDomain), err
}

// Update takes the representation of a awsElasticsearchDomain and updates it. Returns the server's representation of the awsElasticsearchDomain, and an error, if there is any.
func (c *FakeAwsElasticsearchDomains) Update(awsElasticsearchDomain *v1alpha1.AwsElasticsearchDomain) (result *v1alpha1.AwsElasticsearchDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awselasticsearchdomainsResource, awsElasticsearchDomain), &v1alpha1.AwsElasticsearchDomain{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElasticsearchDomain), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsElasticsearchDomains) UpdateStatus(awsElasticsearchDomain *v1alpha1.AwsElasticsearchDomain) (*v1alpha1.AwsElasticsearchDomain, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awselasticsearchdomainsResource, "status", awsElasticsearchDomain), &v1alpha1.AwsElasticsearchDomain{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElasticsearchDomain), err
}

// Delete takes name of the awsElasticsearchDomain and deletes it. Returns an error if one occurs.
func (c *FakeAwsElasticsearchDomains) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awselasticsearchdomainsResource, name), &v1alpha1.AwsElasticsearchDomain{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsElasticsearchDomains) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awselasticsearchdomainsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsElasticsearchDomainList{})
	return err
}

// Patch applies the patch and returns the patched awsElasticsearchDomain.
func (c *FakeAwsElasticsearchDomains) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsElasticsearchDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awselasticsearchdomainsResource, name, pt, data, subresources...), &v1alpha1.AwsElasticsearchDomain{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElasticsearchDomain), err
}
