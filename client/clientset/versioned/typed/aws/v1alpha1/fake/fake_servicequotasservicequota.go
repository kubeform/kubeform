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

// FakeServicequotasServiceQuotas implements ServicequotasServiceQuotaInterface
type FakeServicequotasServiceQuotas struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var servicequotasservicequotasResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "servicequotasservicequotas"}

var servicequotasservicequotasKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "ServicequotasServiceQuota"}

// Get takes name of the servicequotasServiceQuota, and returns the corresponding servicequotasServiceQuota object, and an error if there is any.
func (c *FakeServicequotasServiceQuotas) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ServicequotasServiceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicequotasservicequotasResource, c.ns, name), &v1alpha1.ServicequotasServiceQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicequotasServiceQuota), err
}

// List takes label and field selectors, and returns the list of ServicequotasServiceQuotas that match those selectors.
func (c *FakeServicequotasServiceQuotas) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServicequotasServiceQuotaList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicequotasservicequotasResource, servicequotasservicequotasKind, c.ns, opts), &v1alpha1.ServicequotasServiceQuotaList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ServicequotasServiceQuotaList{ListMeta: obj.(*v1alpha1.ServicequotasServiceQuotaList).ListMeta}
	for _, item := range obj.(*v1alpha1.ServicequotasServiceQuotaList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested servicequotasServiceQuotas.
func (c *FakeServicequotasServiceQuotas) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicequotasservicequotasResource, c.ns, opts))

}

// Create takes the representation of a servicequotasServiceQuota and creates it.  Returns the server's representation of the servicequotasServiceQuota, and an error, if there is any.
func (c *FakeServicequotasServiceQuotas) Create(ctx context.Context, servicequotasServiceQuota *v1alpha1.ServicequotasServiceQuota, opts v1.CreateOptions) (result *v1alpha1.ServicequotasServiceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicequotasservicequotasResource, c.ns, servicequotasServiceQuota), &v1alpha1.ServicequotasServiceQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicequotasServiceQuota), err
}

// Update takes the representation of a servicequotasServiceQuota and updates it. Returns the server's representation of the servicequotasServiceQuota, and an error, if there is any.
func (c *FakeServicequotasServiceQuotas) Update(ctx context.Context, servicequotasServiceQuota *v1alpha1.ServicequotasServiceQuota, opts v1.UpdateOptions) (result *v1alpha1.ServicequotasServiceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicequotasservicequotasResource, c.ns, servicequotasServiceQuota), &v1alpha1.ServicequotasServiceQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicequotasServiceQuota), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServicequotasServiceQuotas) UpdateStatus(ctx context.Context, servicequotasServiceQuota *v1alpha1.ServicequotasServiceQuota, opts v1.UpdateOptions) (*v1alpha1.ServicequotasServiceQuota, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(servicequotasservicequotasResource, "status", c.ns, servicequotasServiceQuota), &v1alpha1.ServicequotasServiceQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicequotasServiceQuota), err
}

// Delete takes name of the servicequotasServiceQuota and deletes it. Returns an error if one occurs.
func (c *FakeServicequotasServiceQuotas) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(servicequotasservicequotasResource, c.ns, name), &v1alpha1.ServicequotasServiceQuota{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServicequotasServiceQuotas) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicequotasservicequotasResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ServicequotasServiceQuotaList{})
	return err
}

// Patch applies the patch and returns the patched servicequotasServiceQuota.
func (c *FakeServicequotasServiceQuotas) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServicequotasServiceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicequotasservicequotasResource, c.ns, name, pt, data, subresources...), &v1alpha1.ServicequotasServiceQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicequotasServiceQuota), err
}