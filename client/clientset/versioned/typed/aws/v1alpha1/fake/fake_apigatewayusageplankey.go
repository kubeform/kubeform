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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApiGatewayUsagePlanKeys implements ApiGatewayUsagePlanKeyInterface
type FakeApiGatewayUsagePlanKeys struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var apigatewayusageplankeysResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "apigatewayusageplankeys"}

var apigatewayusageplankeysKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "ApiGatewayUsagePlanKey"}

// Get takes name of the apiGatewayUsagePlanKey, and returns the corresponding apiGatewayUsagePlanKey object, and an error if there is any.
func (c *FakeApiGatewayUsagePlanKeys) Get(name string, options v1.GetOptions) (result *v1alpha1.ApiGatewayUsagePlanKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apigatewayusageplankeysResource, c.ns, name), &v1alpha1.ApiGatewayUsagePlanKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayUsagePlanKey), err
}

// List takes label and field selectors, and returns the list of ApiGatewayUsagePlanKeys that match those selectors.
func (c *FakeApiGatewayUsagePlanKeys) List(opts v1.ListOptions) (result *v1alpha1.ApiGatewayUsagePlanKeyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apigatewayusageplankeysResource, apigatewayusageplankeysKind, c.ns, opts), &v1alpha1.ApiGatewayUsagePlanKeyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApiGatewayUsagePlanKeyList{ListMeta: obj.(*v1alpha1.ApiGatewayUsagePlanKeyList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApiGatewayUsagePlanKeyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apiGatewayUsagePlanKeys.
func (c *FakeApiGatewayUsagePlanKeys) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apigatewayusageplankeysResource, c.ns, opts))

}

// Create takes the representation of a apiGatewayUsagePlanKey and creates it.  Returns the server's representation of the apiGatewayUsagePlanKey, and an error, if there is any.
func (c *FakeApiGatewayUsagePlanKeys) Create(apiGatewayUsagePlanKey *v1alpha1.ApiGatewayUsagePlanKey) (result *v1alpha1.ApiGatewayUsagePlanKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apigatewayusageplankeysResource, c.ns, apiGatewayUsagePlanKey), &v1alpha1.ApiGatewayUsagePlanKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayUsagePlanKey), err
}

// Update takes the representation of a apiGatewayUsagePlanKey and updates it. Returns the server's representation of the apiGatewayUsagePlanKey, and an error, if there is any.
func (c *FakeApiGatewayUsagePlanKeys) Update(apiGatewayUsagePlanKey *v1alpha1.ApiGatewayUsagePlanKey) (result *v1alpha1.ApiGatewayUsagePlanKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apigatewayusageplankeysResource, c.ns, apiGatewayUsagePlanKey), &v1alpha1.ApiGatewayUsagePlanKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayUsagePlanKey), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApiGatewayUsagePlanKeys) UpdateStatus(apiGatewayUsagePlanKey *v1alpha1.ApiGatewayUsagePlanKey) (*v1alpha1.ApiGatewayUsagePlanKey, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apigatewayusageplankeysResource, "status", c.ns, apiGatewayUsagePlanKey), &v1alpha1.ApiGatewayUsagePlanKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayUsagePlanKey), err
}

// Delete takes name of the apiGatewayUsagePlanKey and deletes it. Returns an error if one occurs.
func (c *FakeApiGatewayUsagePlanKeys) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(apigatewayusageplankeysResource, c.ns, name), &v1alpha1.ApiGatewayUsagePlanKey{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApiGatewayUsagePlanKeys) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apigatewayusageplankeysResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApiGatewayUsagePlanKeyList{})
	return err
}

// Patch applies the patch and returns the patched apiGatewayUsagePlanKey.
func (c *FakeApiGatewayUsagePlanKeys) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiGatewayUsagePlanKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apigatewayusageplankeysResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApiGatewayUsagePlanKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayUsagePlanKey), err
}
