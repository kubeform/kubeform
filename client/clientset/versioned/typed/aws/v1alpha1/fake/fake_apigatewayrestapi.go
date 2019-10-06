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

// FakeApiGatewayRestAPIs implements ApiGatewayRestAPIInterface
type FakeApiGatewayRestAPIs struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var apigatewayrestapisResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "apigatewayrestapis"}

var apigatewayrestapisKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "ApiGatewayRestAPI"}

// Get takes name of the apiGatewayRestAPI, and returns the corresponding apiGatewayRestAPI object, and an error if there is any.
func (c *FakeApiGatewayRestAPIs) Get(name string, options v1.GetOptions) (result *v1alpha1.ApiGatewayRestAPI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apigatewayrestapisResource, c.ns, name), &v1alpha1.ApiGatewayRestAPI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayRestAPI), err
}

// List takes label and field selectors, and returns the list of ApiGatewayRestAPIs that match those selectors.
func (c *FakeApiGatewayRestAPIs) List(opts v1.ListOptions) (result *v1alpha1.ApiGatewayRestAPIList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apigatewayrestapisResource, apigatewayrestapisKind, c.ns, opts), &v1alpha1.ApiGatewayRestAPIList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApiGatewayRestAPIList{ListMeta: obj.(*v1alpha1.ApiGatewayRestAPIList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApiGatewayRestAPIList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apiGatewayRestAPIs.
func (c *FakeApiGatewayRestAPIs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apigatewayrestapisResource, c.ns, opts))

}

// Create takes the representation of a apiGatewayRestAPI and creates it.  Returns the server's representation of the apiGatewayRestAPI, and an error, if there is any.
func (c *FakeApiGatewayRestAPIs) Create(apiGatewayRestAPI *v1alpha1.ApiGatewayRestAPI) (result *v1alpha1.ApiGatewayRestAPI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apigatewayrestapisResource, c.ns, apiGatewayRestAPI), &v1alpha1.ApiGatewayRestAPI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayRestAPI), err
}

// Update takes the representation of a apiGatewayRestAPI and updates it. Returns the server's representation of the apiGatewayRestAPI, and an error, if there is any.
func (c *FakeApiGatewayRestAPIs) Update(apiGatewayRestAPI *v1alpha1.ApiGatewayRestAPI) (result *v1alpha1.ApiGatewayRestAPI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apigatewayrestapisResource, c.ns, apiGatewayRestAPI), &v1alpha1.ApiGatewayRestAPI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayRestAPI), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApiGatewayRestAPIs) UpdateStatus(apiGatewayRestAPI *v1alpha1.ApiGatewayRestAPI) (*v1alpha1.ApiGatewayRestAPI, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apigatewayrestapisResource, "status", c.ns, apiGatewayRestAPI), &v1alpha1.ApiGatewayRestAPI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayRestAPI), err
}

// Delete takes name of the apiGatewayRestAPI and deletes it. Returns an error if one occurs.
func (c *FakeApiGatewayRestAPIs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(apigatewayrestapisResource, c.ns, name), &v1alpha1.ApiGatewayRestAPI{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApiGatewayRestAPIs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apigatewayrestapisResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApiGatewayRestAPIList{})
	return err
}

// Patch applies the patch and returns the patched apiGatewayRestAPI.
func (c *FakeApiGatewayRestAPIs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiGatewayRestAPI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apigatewayrestapisResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApiGatewayRestAPI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayRestAPI), err
}
