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

// FakeCognitoIdentityProviders implements CognitoIdentityProviderInterface
type FakeCognitoIdentityProviders struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var cognitoidentityprovidersResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "cognitoidentityproviders"}

var cognitoidentityprovidersKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "CognitoIdentityProvider"}

// Get takes name of the cognitoIdentityProvider, and returns the corresponding cognitoIdentityProvider object, and an error if there is any.
func (c *FakeCognitoIdentityProviders) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CognitoIdentityProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cognitoidentityprovidersResource, c.ns, name), &v1alpha1.CognitoIdentityProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CognitoIdentityProvider), err
}

// List takes label and field selectors, and returns the list of CognitoIdentityProviders that match those selectors.
func (c *FakeCognitoIdentityProviders) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CognitoIdentityProviderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cognitoidentityprovidersResource, cognitoidentityprovidersKind, c.ns, opts), &v1alpha1.CognitoIdentityProviderList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CognitoIdentityProviderList{ListMeta: obj.(*v1alpha1.CognitoIdentityProviderList).ListMeta}
	for _, item := range obj.(*v1alpha1.CognitoIdentityProviderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cognitoIdentityProviders.
func (c *FakeCognitoIdentityProviders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cognitoidentityprovidersResource, c.ns, opts))

}

// Create takes the representation of a cognitoIdentityProvider and creates it.  Returns the server's representation of the cognitoIdentityProvider, and an error, if there is any.
func (c *FakeCognitoIdentityProviders) Create(ctx context.Context, cognitoIdentityProvider *v1alpha1.CognitoIdentityProvider, opts v1.CreateOptions) (result *v1alpha1.CognitoIdentityProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cognitoidentityprovidersResource, c.ns, cognitoIdentityProvider), &v1alpha1.CognitoIdentityProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CognitoIdentityProvider), err
}

// Update takes the representation of a cognitoIdentityProvider and updates it. Returns the server's representation of the cognitoIdentityProvider, and an error, if there is any.
func (c *FakeCognitoIdentityProviders) Update(ctx context.Context, cognitoIdentityProvider *v1alpha1.CognitoIdentityProvider, opts v1.UpdateOptions) (result *v1alpha1.CognitoIdentityProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cognitoidentityprovidersResource, c.ns, cognitoIdentityProvider), &v1alpha1.CognitoIdentityProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CognitoIdentityProvider), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCognitoIdentityProviders) UpdateStatus(ctx context.Context, cognitoIdentityProvider *v1alpha1.CognitoIdentityProvider, opts v1.UpdateOptions) (*v1alpha1.CognitoIdentityProvider, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cognitoidentityprovidersResource, "status", c.ns, cognitoIdentityProvider), &v1alpha1.CognitoIdentityProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CognitoIdentityProvider), err
}

// Delete takes name of the cognitoIdentityProvider and deletes it. Returns an error if one occurs.
func (c *FakeCognitoIdentityProviders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cognitoidentityprovidersResource, c.ns, name), &v1alpha1.CognitoIdentityProvider{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCognitoIdentityProviders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cognitoidentityprovidersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CognitoIdentityProviderList{})
	return err
}

// Patch applies the patch and returns the patched cognitoIdentityProvider.
func (c *FakeCognitoIdentityProviders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CognitoIdentityProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cognitoidentityprovidersResource, c.ns, name, pt, data, subresources...), &v1alpha1.CognitoIdentityProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CognitoIdentityProvider), err
}
