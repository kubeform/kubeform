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

	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApiManagementProducts implements ApiManagementProductInterface
type FakeApiManagementProducts struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var apimanagementproductsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "apimanagementproducts"}

var apimanagementproductsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "ApiManagementProduct"}

// Get takes name of the apiManagementProduct, and returns the corresponding apiManagementProduct object, and an error if there is any.
func (c *FakeApiManagementProducts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ApiManagementProduct, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apimanagementproductsResource, c.ns, name), &v1alpha1.ApiManagementProduct{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementProduct), err
}

// List takes label and field selectors, and returns the list of ApiManagementProducts that match those selectors.
func (c *FakeApiManagementProducts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApiManagementProductList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apimanagementproductsResource, apimanagementproductsKind, c.ns, opts), &v1alpha1.ApiManagementProductList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApiManagementProductList{ListMeta: obj.(*v1alpha1.ApiManagementProductList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApiManagementProductList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apiManagementProducts.
func (c *FakeApiManagementProducts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apimanagementproductsResource, c.ns, opts))

}

// Create takes the representation of a apiManagementProduct and creates it.  Returns the server's representation of the apiManagementProduct, and an error, if there is any.
func (c *FakeApiManagementProducts) Create(ctx context.Context, apiManagementProduct *v1alpha1.ApiManagementProduct, opts v1.CreateOptions) (result *v1alpha1.ApiManagementProduct, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apimanagementproductsResource, c.ns, apiManagementProduct), &v1alpha1.ApiManagementProduct{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementProduct), err
}

// Update takes the representation of a apiManagementProduct and updates it. Returns the server's representation of the apiManagementProduct, and an error, if there is any.
func (c *FakeApiManagementProducts) Update(ctx context.Context, apiManagementProduct *v1alpha1.ApiManagementProduct, opts v1.UpdateOptions) (result *v1alpha1.ApiManagementProduct, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apimanagementproductsResource, c.ns, apiManagementProduct), &v1alpha1.ApiManagementProduct{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementProduct), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApiManagementProducts) UpdateStatus(ctx context.Context, apiManagementProduct *v1alpha1.ApiManagementProduct, opts v1.UpdateOptions) (*v1alpha1.ApiManagementProduct, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apimanagementproductsResource, "status", c.ns, apiManagementProduct), &v1alpha1.ApiManagementProduct{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementProduct), err
}

// Delete takes name of the apiManagementProduct and deletes it. Returns an error if one occurs.
func (c *FakeApiManagementProducts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(apimanagementproductsResource, c.ns, name), &v1alpha1.ApiManagementProduct{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApiManagementProducts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apimanagementproductsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApiManagementProductList{})
	return err
}

// Patch applies the patch and returns the patched apiManagementProduct.
func (c *FakeApiManagementProducts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApiManagementProduct, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apimanagementproductsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApiManagementProduct{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementProduct), err
}
