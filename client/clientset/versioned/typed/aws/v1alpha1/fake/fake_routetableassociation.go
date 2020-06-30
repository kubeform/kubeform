/*
Copyright The Kubeform Authors.

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
	"context"

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRouteTableAssociations implements RouteTableAssociationInterface
type FakeRouteTableAssociations struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var routetableassociationsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "routetableassociations"}

var routetableassociationsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "RouteTableAssociation"}

// Get takes name of the routeTableAssociation, and returns the corresponding routeTableAssociation object, and an error if there is any.
func (c *FakeRouteTableAssociations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RouteTableAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(routetableassociationsResource, c.ns, name), &v1alpha1.RouteTableAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RouteTableAssociation), err
}

// List takes label and field selectors, and returns the list of RouteTableAssociations that match those selectors.
func (c *FakeRouteTableAssociations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RouteTableAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(routetableassociationsResource, routetableassociationsKind, c.ns, opts), &v1alpha1.RouteTableAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RouteTableAssociationList{ListMeta: obj.(*v1alpha1.RouteTableAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.RouteTableAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested routeTableAssociations.
func (c *FakeRouteTableAssociations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(routetableassociationsResource, c.ns, opts))

}

// Create takes the representation of a routeTableAssociation and creates it.  Returns the server's representation of the routeTableAssociation, and an error, if there is any.
func (c *FakeRouteTableAssociations) Create(ctx context.Context, routeTableAssociation *v1alpha1.RouteTableAssociation, opts v1.CreateOptions) (result *v1alpha1.RouteTableAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(routetableassociationsResource, c.ns, routeTableAssociation), &v1alpha1.RouteTableAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RouteTableAssociation), err
}

// Update takes the representation of a routeTableAssociation and updates it. Returns the server's representation of the routeTableAssociation, and an error, if there is any.
func (c *FakeRouteTableAssociations) Update(ctx context.Context, routeTableAssociation *v1alpha1.RouteTableAssociation, opts v1.UpdateOptions) (result *v1alpha1.RouteTableAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(routetableassociationsResource, c.ns, routeTableAssociation), &v1alpha1.RouteTableAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RouteTableAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRouteTableAssociations) UpdateStatus(ctx context.Context, routeTableAssociation *v1alpha1.RouteTableAssociation, opts v1.UpdateOptions) (*v1alpha1.RouteTableAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(routetableassociationsResource, "status", c.ns, routeTableAssociation), &v1alpha1.RouteTableAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RouteTableAssociation), err
}

// Delete takes name of the routeTableAssociation and deletes it. Returns an error if one occurs.
func (c *FakeRouteTableAssociations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(routetableassociationsResource, c.ns, name), &v1alpha1.RouteTableAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRouteTableAssociations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(routetableassociationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RouteTableAssociationList{})
	return err
}

// Patch applies the patch and returns the patched routeTableAssociation.
func (c *FakeRouteTableAssociations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RouteTableAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(routetableassociationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.RouteTableAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RouteTableAssociation), err
}
