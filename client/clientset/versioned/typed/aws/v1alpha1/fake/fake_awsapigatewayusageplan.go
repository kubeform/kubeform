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

// FakeAwsApiGatewayUsagePlans implements AwsApiGatewayUsagePlanInterface
type FakeAwsApiGatewayUsagePlans struct {
	Fake *FakeAwsV1alpha1
}

var awsapigatewayusageplansResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsapigatewayusageplans"}

var awsapigatewayusageplansKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsApiGatewayUsagePlan"}

// Get takes name of the awsApiGatewayUsagePlan, and returns the corresponding awsApiGatewayUsagePlan object, and an error if there is any.
func (c *FakeAwsApiGatewayUsagePlans) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsApiGatewayUsagePlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsapigatewayusageplansResource, name), &v1alpha1.AwsApiGatewayUsagePlan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayUsagePlan), err
}

// List takes label and field selectors, and returns the list of AwsApiGatewayUsagePlans that match those selectors.
func (c *FakeAwsApiGatewayUsagePlans) List(opts v1.ListOptions) (result *v1alpha1.AwsApiGatewayUsagePlanList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsapigatewayusageplansResource, awsapigatewayusageplansKind, opts), &v1alpha1.AwsApiGatewayUsagePlanList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsApiGatewayUsagePlanList{ListMeta: obj.(*v1alpha1.AwsApiGatewayUsagePlanList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsApiGatewayUsagePlanList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsApiGatewayUsagePlans.
func (c *FakeAwsApiGatewayUsagePlans) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsapigatewayusageplansResource, opts))
}

// Create takes the representation of a awsApiGatewayUsagePlan and creates it.  Returns the server's representation of the awsApiGatewayUsagePlan, and an error, if there is any.
func (c *FakeAwsApiGatewayUsagePlans) Create(awsApiGatewayUsagePlan *v1alpha1.AwsApiGatewayUsagePlan) (result *v1alpha1.AwsApiGatewayUsagePlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsapigatewayusageplansResource, awsApiGatewayUsagePlan), &v1alpha1.AwsApiGatewayUsagePlan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayUsagePlan), err
}

// Update takes the representation of a awsApiGatewayUsagePlan and updates it. Returns the server's representation of the awsApiGatewayUsagePlan, and an error, if there is any.
func (c *FakeAwsApiGatewayUsagePlans) Update(awsApiGatewayUsagePlan *v1alpha1.AwsApiGatewayUsagePlan) (result *v1alpha1.AwsApiGatewayUsagePlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsapigatewayusageplansResource, awsApiGatewayUsagePlan), &v1alpha1.AwsApiGatewayUsagePlan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayUsagePlan), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsApiGatewayUsagePlans) UpdateStatus(awsApiGatewayUsagePlan *v1alpha1.AwsApiGatewayUsagePlan) (*v1alpha1.AwsApiGatewayUsagePlan, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsapigatewayusageplansResource, "status", awsApiGatewayUsagePlan), &v1alpha1.AwsApiGatewayUsagePlan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayUsagePlan), err
}

// Delete takes name of the awsApiGatewayUsagePlan and deletes it. Returns an error if one occurs.
func (c *FakeAwsApiGatewayUsagePlans) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsapigatewayusageplansResource, name), &v1alpha1.AwsApiGatewayUsagePlan{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsApiGatewayUsagePlans) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsapigatewayusageplansResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsApiGatewayUsagePlanList{})
	return err
}

// Patch applies the patch and returns the patched awsApiGatewayUsagePlan.
func (c *FakeAwsApiGatewayUsagePlans) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayUsagePlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsapigatewayusageplansResource, name, pt, data, subresources...), &v1alpha1.AwsApiGatewayUsagePlan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayUsagePlan), err
}
