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

// FakeSimpledbDomains implements SimpledbDomainInterface
type FakeSimpledbDomains struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var simpledbdomainsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "simpledbdomains"}

var simpledbdomainsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "SimpledbDomain"}

// Get takes name of the simpledbDomain, and returns the corresponding simpledbDomain object, and an error if there is any.
func (c *FakeSimpledbDomains) Get(name string, options v1.GetOptions) (result *v1alpha1.SimpledbDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(simpledbdomainsResource, c.ns, name), &v1alpha1.SimpledbDomain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SimpledbDomain), err
}

// List takes label and field selectors, and returns the list of SimpledbDomains that match those selectors.
func (c *FakeSimpledbDomains) List(opts v1.ListOptions) (result *v1alpha1.SimpledbDomainList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(simpledbdomainsResource, simpledbdomainsKind, c.ns, opts), &v1alpha1.SimpledbDomainList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SimpledbDomainList{ListMeta: obj.(*v1alpha1.SimpledbDomainList).ListMeta}
	for _, item := range obj.(*v1alpha1.SimpledbDomainList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested simpledbDomains.
func (c *FakeSimpledbDomains) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(simpledbdomainsResource, c.ns, opts))

}

// Create takes the representation of a simpledbDomain and creates it.  Returns the server's representation of the simpledbDomain, and an error, if there is any.
func (c *FakeSimpledbDomains) Create(simpledbDomain *v1alpha1.SimpledbDomain) (result *v1alpha1.SimpledbDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(simpledbdomainsResource, c.ns, simpledbDomain), &v1alpha1.SimpledbDomain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SimpledbDomain), err
}

// Update takes the representation of a simpledbDomain and updates it. Returns the server's representation of the simpledbDomain, and an error, if there is any.
func (c *FakeSimpledbDomains) Update(simpledbDomain *v1alpha1.SimpledbDomain) (result *v1alpha1.SimpledbDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(simpledbdomainsResource, c.ns, simpledbDomain), &v1alpha1.SimpledbDomain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SimpledbDomain), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSimpledbDomains) UpdateStatus(simpledbDomain *v1alpha1.SimpledbDomain) (*v1alpha1.SimpledbDomain, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(simpledbdomainsResource, "status", c.ns, simpledbDomain), &v1alpha1.SimpledbDomain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SimpledbDomain), err
}

// Delete takes name of the simpledbDomain and deletes it. Returns an error if one occurs.
func (c *FakeSimpledbDomains) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(simpledbdomainsResource, c.ns, name), &v1alpha1.SimpledbDomain{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSimpledbDomains) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(simpledbdomainsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SimpledbDomainList{})
	return err
}

// Patch applies the patch and returns the patched simpledbDomain.
func (c *FakeSimpledbDomains) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SimpledbDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(simpledbdomainsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SimpledbDomain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SimpledbDomain), err
}
