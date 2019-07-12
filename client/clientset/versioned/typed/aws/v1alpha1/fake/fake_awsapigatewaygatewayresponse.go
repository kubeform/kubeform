/*
Copyright 2019 The KubeDB Authors.

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

// FakeAwsApiGatewayGatewayResponses implements AwsApiGatewayGatewayResponseInterface
type FakeAwsApiGatewayGatewayResponses struct {
	Fake *FakeAwsV1alpha1
}

var awsapigatewaygatewayresponsesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsapigatewaygatewayresponses"}

var awsapigatewaygatewayresponsesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsApiGatewayGatewayResponse"}

// Get takes name of the awsApiGatewayGatewayResponse, and returns the corresponding awsApiGatewayGatewayResponse object, and an error if there is any.
func (c *FakeAwsApiGatewayGatewayResponses) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsApiGatewayGatewayResponse, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsapigatewaygatewayresponsesResource, name), &v1alpha1.AwsApiGatewayGatewayResponse{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayGatewayResponse), err
}

// List takes label and field selectors, and returns the list of AwsApiGatewayGatewayResponses that match those selectors.
func (c *FakeAwsApiGatewayGatewayResponses) List(opts v1.ListOptions) (result *v1alpha1.AwsApiGatewayGatewayResponseList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsapigatewaygatewayresponsesResource, awsapigatewaygatewayresponsesKind, opts), &v1alpha1.AwsApiGatewayGatewayResponseList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsApiGatewayGatewayResponseList{ListMeta: obj.(*v1alpha1.AwsApiGatewayGatewayResponseList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsApiGatewayGatewayResponseList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsApiGatewayGatewayResponses.
func (c *FakeAwsApiGatewayGatewayResponses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsapigatewaygatewayresponsesResource, opts))
}

// Create takes the representation of a awsApiGatewayGatewayResponse and creates it.  Returns the server's representation of the awsApiGatewayGatewayResponse, and an error, if there is any.
func (c *FakeAwsApiGatewayGatewayResponses) Create(awsApiGatewayGatewayResponse *v1alpha1.AwsApiGatewayGatewayResponse) (result *v1alpha1.AwsApiGatewayGatewayResponse, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsapigatewaygatewayresponsesResource, awsApiGatewayGatewayResponse), &v1alpha1.AwsApiGatewayGatewayResponse{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayGatewayResponse), err
}

// Update takes the representation of a awsApiGatewayGatewayResponse and updates it. Returns the server's representation of the awsApiGatewayGatewayResponse, and an error, if there is any.
func (c *FakeAwsApiGatewayGatewayResponses) Update(awsApiGatewayGatewayResponse *v1alpha1.AwsApiGatewayGatewayResponse) (result *v1alpha1.AwsApiGatewayGatewayResponse, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsapigatewaygatewayresponsesResource, awsApiGatewayGatewayResponse), &v1alpha1.AwsApiGatewayGatewayResponse{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayGatewayResponse), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsApiGatewayGatewayResponses) UpdateStatus(awsApiGatewayGatewayResponse *v1alpha1.AwsApiGatewayGatewayResponse) (*v1alpha1.AwsApiGatewayGatewayResponse, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsapigatewaygatewayresponsesResource, "status", awsApiGatewayGatewayResponse), &v1alpha1.AwsApiGatewayGatewayResponse{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayGatewayResponse), err
}

// Delete takes name of the awsApiGatewayGatewayResponse and deletes it. Returns an error if one occurs.
func (c *FakeAwsApiGatewayGatewayResponses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsapigatewaygatewayresponsesResource, name), &v1alpha1.AwsApiGatewayGatewayResponse{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsApiGatewayGatewayResponses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsapigatewaygatewayresponsesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsApiGatewayGatewayResponseList{})
	return err
}

// Patch applies the patch and returns the patched awsApiGatewayGatewayResponse.
func (c *FakeAwsApiGatewayGatewayResponses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayGatewayResponse, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsapigatewaygatewayresponsesResource, name, pt, data, subresources...), &v1alpha1.AwsApiGatewayGatewayResponse{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayGatewayResponse), err
}
