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

// FakeDefaultVpcs implements DefaultVpcInterface
type FakeDefaultVpcs struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var defaultvpcsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "defaultvpcs"}

var defaultvpcsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DefaultVpc"}

// Get takes name of the defaultVpc, and returns the corresponding defaultVpc object, and an error if there is any.
func (c *FakeDefaultVpcs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DefaultVpc, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(defaultvpcsResource, c.ns, name), &v1alpha1.DefaultVpc{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DefaultVpc), err
}

// List takes label and field selectors, and returns the list of DefaultVpcs that match those selectors.
func (c *FakeDefaultVpcs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DefaultVpcList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(defaultvpcsResource, defaultvpcsKind, c.ns, opts), &v1alpha1.DefaultVpcList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DefaultVpcList{ListMeta: obj.(*v1alpha1.DefaultVpcList).ListMeta}
	for _, item := range obj.(*v1alpha1.DefaultVpcList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested defaultVpcs.
func (c *FakeDefaultVpcs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(defaultvpcsResource, c.ns, opts))

}

// Create takes the representation of a defaultVpc and creates it.  Returns the server's representation of the defaultVpc, and an error, if there is any.
func (c *FakeDefaultVpcs) Create(ctx context.Context, defaultVpc *v1alpha1.DefaultVpc, opts v1.CreateOptions) (result *v1alpha1.DefaultVpc, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(defaultvpcsResource, c.ns, defaultVpc), &v1alpha1.DefaultVpc{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DefaultVpc), err
}

// Update takes the representation of a defaultVpc and updates it. Returns the server's representation of the defaultVpc, and an error, if there is any.
func (c *FakeDefaultVpcs) Update(ctx context.Context, defaultVpc *v1alpha1.DefaultVpc, opts v1.UpdateOptions) (result *v1alpha1.DefaultVpc, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(defaultvpcsResource, c.ns, defaultVpc), &v1alpha1.DefaultVpc{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DefaultVpc), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDefaultVpcs) UpdateStatus(ctx context.Context, defaultVpc *v1alpha1.DefaultVpc, opts v1.UpdateOptions) (*v1alpha1.DefaultVpc, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(defaultvpcsResource, "status", c.ns, defaultVpc), &v1alpha1.DefaultVpc{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DefaultVpc), err
}

// Delete takes name of the defaultVpc and deletes it. Returns an error if one occurs.
func (c *FakeDefaultVpcs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(defaultvpcsResource, c.ns, name), &v1alpha1.DefaultVpc{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDefaultVpcs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(defaultvpcsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DefaultVpcList{})
	return err
}

// Patch applies the patch and returns the patched defaultVpc.
func (c *FakeDefaultVpcs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DefaultVpc, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(defaultvpcsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DefaultVpc{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DefaultVpc), err
}
