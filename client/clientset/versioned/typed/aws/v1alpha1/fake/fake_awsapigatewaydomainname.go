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

// FakeAwsApiGatewayDomainNames implements AwsApiGatewayDomainNameInterface
type FakeAwsApiGatewayDomainNames struct {
	Fake *FakeAwsV1alpha1
}

var awsapigatewaydomainnamesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsapigatewaydomainnames"}

var awsapigatewaydomainnamesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsApiGatewayDomainName"}

// Get takes name of the awsApiGatewayDomainName, and returns the corresponding awsApiGatewayDomainName object, and an error if there is any.
func (c *FakeAwsApiGatewayDomainNames) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsApiGatewayDomainName, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsapigatewaydomainnamesResource, name), &v1alpha1.AwsApiGatewayDomainName{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayDomainName), err
}

// List takes label and field selectors, and returns the list of AwsApiGatewayDomainNames that match those selectors.
func (c *FakeAwsApiGatewayDomainNames) List(opts v1.ListOptions) (result *v1alpha1.AwsApiGatewayDomainNameList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsapigatewaydomainnamesResource, awsapigatewaydomainnamesKind, opts), &v1alpha1.AwsApiGatewayDomainNameList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsApiGatewayDomainNameList{ListMeta: obj.(*v1alpha1.AwsApiGatewayDomainNameList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsApiGatewayDomainNameList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsApiGatewayDomainNames.
func (c *FakeAwsApiGatewayDomainNames) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsapigatewaydomainnamesResource, opts))
}

// Create takes the representation of a awsApiGatewayDomainName and creates it.  Returns the server's representation of the awsApiGatewayDomainName, and an error, if there is any.
func (c *FakeAwsApiGatewayDomainNames) Create(awsApiGatewayDomainName *v1alpha1.AwsApiGatewayDomainName) (result *v1alpha1.AwsApiGatewayDomainName, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsapigatewaydomainnamesResource, awsApiGatewayDomainName), &v1alpha1.AwsApiGatewayDomainName{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayDomainName), err
}

// Update takes the representation of a awsApiGatewayDomainName and updates it. Returns the server's representation of the awsApiGatewayDomainName, and an error, if there is any.
func (c *FakeAwsApiGatewayDomainNames) Update(awsApiGatewayDomainName *v1alpha1.AwsApiGatewayDomainName) (result *v1alpha1.AwsApiGatewayDomainName, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsapigatewaydomainnamesResource, awsApiGatewayDomainName), &v1alpha1.AwsApiGatewayDomainName{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayDomainName), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsApiGatewayDomainNames) UpdateStatus(awsApiGatewayDomainName *v1alpha1.AwsApiGatewayDomainName) (*v1alpha1.AwsApiGatewayDomainName, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsapigatewaydomainnamesResource, "status", awsApiGatewayDomainName), &v1alpha1.AwsApiGatewayDomainName{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayDomainName), err
}

// Delete takes name of the awsApiGatewayDomainName and deletes it. Returns an error if one occurs.
func (c *FakeAwsApiGatewayDomainNames) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsapigatewaydomainnamesResource, name), &v1alpha1.AwsApiGatewayDomainName{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsApiGatewayDomainNames) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsapigatewaydomainnamesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsApiGatewayDomainNameList{})
	return err
}

// Patch applies the patch and returns the patched awsApiGatewayDomainName.
func (c *FakeAwsApiGatewayDomainNames) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayDomainName, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsapigatewaydomainnamesResource, name, pt, data, subresources...), &v1alpha1.AwsApiGatewayDomainName{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayDomainName), err
}
