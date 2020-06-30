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

// FakeKeyVaultCertificates implements KeyVaultCertificateInterface
type FakeKeyVaultCertificates struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var keyvaultcertificatesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "keyvaultcertificates"}

var keyvaultcertificatesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "KeyVaultCertificate"}

// Get takes name of the keyVaultCertificate, and returns the corresponding keyVaultCertificate object, and an error if there is any.
func (c *FakeKeyVaultCertificates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.KeyVaultCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(keyvaultcertificatesResource, c.ns, name), &v1alpha1.KeyVaultCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeyVaultCertificate), err
}

// List takes label and field selectors, and returns the list of KeyVaultCertificates that match those selectors.
func (c *FakeKeyVaultCertificates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KeyVaultCertificateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(keyvaultcertificatesResource, keyvaultcertificatesKind, c.ns, opts), &v1alpha1.KeyVaultCertificateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KeyVaultCertificateList{ListMeta: obj.(*v1alpha1.KeyVaultCertificateList).ListMeta}
	for _, item := range obj.(*v1alpha1.KeyVaultCertificateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested keyVaultCertificates.
func (c *FakeKeyVaultCertificates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(keyvaultcertificatesResource, c.ns, opts))

}

// Create takes the representation of a keyVaultCertificate and creates it.  Returns the server's representation of the keyVaultCertificate, and an error, if there is any.
func (c *FakeKeyVaultCertificates) Create(ctx context.Context, keyVaultCertificate *v1alpha1.KeyVaultCertificate, opts v1.CreateOptions) (result *v1alpha1.KeyVaultCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(keyvaultcertificatesResource, c.ns, keyVaultCertificate), &v1alpha1.KeyVaultCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeyVaultCertificate), err
}

// Update takes the representation of a keyVaultCertificate and updates it. Returns the server's representation of the keyVaultCertificate, and an error, if there is any.
func (c *FakeKeyVaultCertificates) Update(ctx context.Context, keyVaultCertificate *v1alpha1.KeyVaultCertificate, opts v1.UpdateOptions) (result *v1alpha1.KeyVaultCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(keyvaultcertificatesResource, c.ns, keyVaultCertificate), &v1alpha1.KeyVaultCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeyVaultCertificate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKeyVaultCertificates) UpdateStatus(ctx context.Context, keyVaultCertificate *v1alpha1.KeyVaultCertificate, opts v1.UpdateOptions) (*v1alpha1.KeyVaultCertificate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(keyvaultcertificatesResource, "status", c.ns, keyVaultCertificate), &v1alpha1.KeyVaultCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeyVaultCertificate), err
}

// Delete takes name of the keyVaultCertificate and deletes it. Returns an error if one occurs.
func (c *FakeKeyVaultCertificates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(keyvaultcertificatesResource, c.ns, name), &v1alpha1.KeyVaultCertificate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKeyVaultCertificates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(keyvaultcertificatesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.KeyVaultCertificateList{})
	return err
}

// Patch applies the patch and returns the patched keyVaultCertificate.
func (c *FakeKeyVaultCertificates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KeyVaultCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(keyvaultcertificatesResource, c.ns, name, pt, data, subresources...), &v1alpha1.KeyVaultCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeyVaultCertificate), err
}
