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

// FakeDxConnections implements DxConnectionInterface
type FakeDxConnections struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var dxconnectionsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "dxconnections"}

var dxconnectionsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DxConnection"}

// Get takes name of the dxConnection, and returns the corresponding dxConnection object, and an error if there is any.
func (c *FakeDxConnections) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DxConnection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dxconnectionsResource, c.ns, name), &v1alpha1.DxConnection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxConnection), err
}

// List takes label and field selectors, and returns the list of DxConnections that match those selectors.
func (c *FakeDxConnections) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DxConnectionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dxconnectionsResource, dxconnectionsKind, c.ns, opts), &v1alpha1.DxConnectionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DxConnectionList{ListMeta: obj.(*v1alpha1.DxConnectionList).ListMeta}
	for _, item := range obj.(*v1alpha1.DxConnectionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dxConnections.
func (c *FakeDxConnections) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dxconnectionsResource, c.ns, opts))

}

// Create takes the representation of a dxConnection and creates it.  Returns the server's representation of the dxConnection, and an error, if there is any.
func (c *FakeDxConnections) Create(ctx context.Context, dxConnection *v1alpha1.DxConnection, opts v1.CreateOptions) (result *v1alpha1.DxConnection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dxconnectionsResource, c.ns, dxConnection), &v1alpha1.DxConnection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxConnection), err
}

// Update takes the representation of a dxConnection and updates it. Returns the server's representation of the dxConnection, and an error, if there is any.
func (c *FakeDxConnections) Update(ctx context.Context, dxConnection *v1alpha1.DxConnection, opts v1.UpdateOptions) (result *v1alpha1.DxConnection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dxconnectionsResource, c.ns, dxConnection), &v1alpha1.DxConnection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxConnection), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDxConnections) UpdateStatus(ctx context.Context, dxConnection *v1alpha1.DxConnection, opts v1.UpdateOptions) (*v1alpha1.DxConnection, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dxconnectionsResource, "status", c.ns, dxConnection), &v1alpha1.DxConnection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxConnection), err
}

// Delete takes name of the dxConnection and deletes it. Returns an error if one occurs.
func (c *FakeDxConnections) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dxconnectionsResource, c.ns, name), &v1alpha1.DxConnection{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDxConnections) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dxconnectionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DxConnectionList{})
	return err
}

// Patch applies the patch and returns the patched dxConnection.
func (c *FakeDxConnections) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DxConnection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dxconnectionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DxConnection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxConnection), err
}
