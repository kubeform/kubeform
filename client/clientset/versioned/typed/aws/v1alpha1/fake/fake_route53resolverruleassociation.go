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

// FakeRoute53ResolverRuleAssociations implements Route53ResolverRuleAssociationInterface
type FakeRoute53ResolverRuleAssociations struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var route53resolverruleassociationsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "route53resolverruleassociations"}

var route53resolverruleassociationsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "Route53ResolverRuleAssociation"}

// Get takes name of the route53ResolverRuleAssociation, and returns the corresponding route53ResolverRuleAssociation object, and an error if there is any.
func (c *FakeRoute53ResolverRuleAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.Route53ResolverRuleAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(route53resolverruleassociationsResource, c.ns, name), &v1alpha1.Route53ResolverRuleAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Route53ResolverRuleAssociation), err
}

// List takes label and field selectors, and returns the list of Route53ResolverRuleAssociations that match those selectors.
func (c *FakeRoute53ResolverRuleAssociations) List(opts v1.ListOptions) (result *v1alpha1.Route53ResolverRuleAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(route53resolverruleassociationsResource, route53resolverruleassociationsKind, c.ns, opts), &v1alpha1.Route53ResolverRuleAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.Route53ResolverRuleAssociationList{ListMeta: obj.(*v1alpha1.Route53ResolverRuleAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.Route53ResolverRuleAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested route53ResolverRuleAssociations.
func (c *FakeRoute53ResolverRuleAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(route53resolverruleassociationsResource, c.ns, opts))

}

// Create takes the representation of a route53ResolverRuleAssociation and creates it.  Returns the server's representation of the route53ResolverRuleAssociation, and an error, if there is any.
func (c *FakeRoute53ResolverRuleAssociations) Create(route53ResolverRuleAssociation *v1alpha1.Route53ResolverRuleAssociation) (result *v1alpha1.Route53ResolverRuleAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(route53resolverruleassociationsResource, c.ns, route53ResolverRuleAssociation), &v1alpha1.Route53ResolverRuleAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Route53ResolverRuleAssociation), err
}

// Update takes the representation of a route53ResolverRuleAssociation and updates it. Returns the server's representation of the route53ResolverRuleAssociation, and an error, if there is any.
func (c *FakeRoute53ResolverRuleAssociations) Update(route53ResolverRuleAssociation *v1alpha1.Route53ResolverRuleAssociation) (result *v1alpha1.Route53ResolverRuleAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(route53resolverruleassociationsResource, c.ns, route53ResolverRuleAssociation), &v1alpha1.Route53ResolverRuleAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Route53ResolverRuleAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRoute53ResolverRuleAssociations) UpdateStatus(route53ResolverRuleAssociation *v1alpha1.Route53ResolverRuleAssociation) (*v1alpha1.Route53ResolverRuleAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(route53resolverruleassociationsResource, "status", c.ns, route53ResolverRuleAssociation), &v1alpha1.Route53ResolverRuleAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Route53ResolverRuleAssociation), err
}

// Delete takes name of the route53ResolverRuleAssociation and deletes it. Returns an error if one occurs.
func (c *FakeRoute53ResolverRuleAssociations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(route53resolverruleassociationsResource, c.ns, name), &v1alpha1.Route53ResolverRuleAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRoute53ResolverRuleAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(route53resolverruleassociationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.Route53ResolverRuleAssociationList{})
	return err
}

// Patch applies the patch and returns the patched route53ResolverRuleAssociation.
func (c *FakeRoute53ResolverRuleAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Route53ResolverRuleAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(route53resolverruleassociationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Route53ResolverRuleAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Route53ResolverRuleAssociation), err
}