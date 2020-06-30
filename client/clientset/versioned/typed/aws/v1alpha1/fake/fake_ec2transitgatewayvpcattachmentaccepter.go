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

// FakeEc2TransitGatewayVpcAttachmentAccepters implements Ec2TransitGatewayVpcAttachmentAccepterInterface
type FakeEc2TransitGatewayVpcAttachmentAccepters struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var ec2transitgatewayvpcattachmentacceptersResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "ec2transitgatewayvpcattachmentaccepters"}

var ec2transitgatewayvpcattachmentacceptersKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "Ec2TransitGatewayVpcAttachmentAccepter"}

// Get takes name of the ec2TransitGatewayVpcAttachmentAccepter, and returns the corresponding ec2TransitGatewayVpcAttachmentAccepter object, and an error if there is any.
func (c *FakeEc2TransitGatewayVpcAttachmentAccepters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Ec2TransitGatewayVpcAttachmentAccepter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ec2transitgatewayvpcattachmentacceptersResource, c.ns, name), &v1alpha1.Ec2TransitGatewayVpcAttachmentAccepter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ec2TransitGatewayVpcAttachmentAccepter), err
}

// List takes label and field selectors, and returns the list of Ec2TransitGatewayVpcAttachmentAccepters that match those selectors.
func (c *FakeEc2TransitGatewayVpcAttachmentAccepters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.Ec2TransitGatewayVpcAttachmentAccepterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ec2transitgatewayvpcattachmentacceptersResource, ec2transitgatewayvpcattachmentacceptersKind, c.ns, opts), &v1alpha1.Ec2TransitGatewayVpcAttachmentAccepterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.Ec2TransitGatewayVpcAttachmentAccepterList{ListMeta: obj.(*v1alpha1.Ec2TransitGatewayVpcAttachmentAccepterList).ListMeta}
	for _, item := range obj.(*v1alpha1.Ec2TransitGatewayVpcAttachmentAccepterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ec2TransitGatewayVpcAttachmentAccepters.
func (c *FakeEc2TransitGatewayVpcAttachmentAccepters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ec2transitgatewayvpcattachmentacceptersResource, c.ns, opts))

}

// Create takes the representation of a ec2TransitGatewayVpcAttachmentAccepter and creates it.  Returns the server's representation of the ec2TransitGatewayVpcAttachmentAccepter, and an error, if there is any.
func (c *FakeEc2TransitGatewayVpcAttachmentAccepters) Create(ctx context.Context, ec2TransitGatewayVpcAttachmentAccepter *v1alpha1.Ec2TransitGatewayVpcAttachmentAccepter, opts v1.CreateOptions) (result *v1alpha1.Ec2TransitGatewayVpcAttachmentAccepter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ec2transitgatewayvpcattachmentacceptersResource, c.ns, ec2TransitGatewayVpcAttachmentAccepter), &v1alpha1.Ec2TransitGatewayVpcAttachmentAccepter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ec2TransitGatewayVpcAttachmentAccepter), err
}

// Update takes the representation of a ec2TransitGatewayVpcAttachmentAccepter and updates it. Returns the server's representation of the ec2TransitGatewayVpcAttachmentAccepter, and an error, if there is any.
func (c *FakeEc2TransitGatewayVpcAttachmentAccepters) Update(ctx context.Context, ec2TransitGatewayVpcAttachmentAccepter *v1alpha1.Ec2TransitGatewayVpcAttachmentAccepter, opts v1.UpdateOptions) (result *v1alpha1.Ec2TransitGatewayVpcAttachmentAccepter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ec2transitgatewayvpcattachmentacceptersResource, c.ns, ec2TransitGatewayVpcAttachmentAccepter), &v1alpha1.Ec2TransitGatewayVpcAttachmentAccepter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ec2TransitGatewayVpcAttachmentAccepter), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEc2TransitGatewayVpcAttachmentAccepters) UpdateStatus(ctx context.Context, ec2TransitGatewayVpcAttachmentAccepter *v1alpha1.Ec2TransitGatewayVpcAttachmentAccepter, opts v1.UpdateOptions) (*v1alpha1.Ec2TransitGatewayVpcAttachmentAccepter, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ec2transitgatewayvpcattachmentacceptersResource, "status", c.ns, ec2TransitGatewayVpcAttachmentAccepter), &v1alpha1.Ec2TransitGatewayVpcAttachmentAccepter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ec2TransitGatewayVpcAttachmentAccepter), err
}

// Delete takes name of the ec2TransitGatewayVpcAttachmentAccepter and deletes it. Returns an error if one occurs.
func (c *FakeEc2TransitGatewayVpcAttachmentAccepters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ec2transitgatewayvpcattachmentacceptersResource, c.ns, name), &v1alpha1.Ec2TransitGatewayVpcAttachmentAccepter{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEc2TransitGatewayVpcAttachmentAccepters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ec2transitgatewayvpcattachmentacceptersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.Ec2TransitGatewayVpcAttachmentAccepterList{})
	return err
}

// Patch applies the patch and returns the patched ec2TransitGatewayVpcAttachmentAccepter.
func (c *FakeEc2TransitGatewayVpcAttachmentAccepters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Ec2TransitGatewayVpcAttachmentAccepter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ec2transitgatewayvpcattachmentacceptersResource, c.ns, name, pt, data, subresources...), &v1alpha1.Ec2TransitGatewayVpcAttachmentAccepter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ec2TransitGatewayVpcAttachmentAccepter), err
}
