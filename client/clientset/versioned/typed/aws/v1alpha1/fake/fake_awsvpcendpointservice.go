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

// FakeAwsVpcEndpointServices implements AwsVpcEndpointServiceInterface
type FakeAwsVpcEndpointServices struct {
	Fake *FakeAwsV1alpha1
}

var awsvpcendpointservicesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsvpcendpointservices"}

var awsvpcendpointservicesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsVpcEndpointService"}

// Get takes name of the awsVpcEndpointService, and returns the corresponding awsVpcEndpointService object, and an error if there is any.
func (c *FakeAwsVpcEndpointServices) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsVpcEndpointService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsvpcendpointservicesResource, name), &v1alpha1.AwsVpcEndpointService{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVpcEndpointService), err
}

// List takes label and field selectors, and returns the list of AwsVpcEndpointServices that match those selectors.
func (c *FakeAwsVpcEndpointServices) List(opts v1.ListOptions) (result *v1alpha1.AwsVpcEndpointServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsvpcendpointservicesResource, awsvpcendpointservicesKind, opts), &v1alpha1.AwsVpcEndpointServiceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsVpcEndpointServiceList{ListMeta: obj.(*v1alpha1.AwsVpcEndpointServiceList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsVpcEndpointServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsVpcEndpointServices.
func (c *FakeAwsVpcEndpointServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsvpcendpointservicesResource, opts))
}

// Create takes the representation of a awsVpcEndpointService and creates it.  Returns the server's representation of the awsVpcEndpointService, and an error, if there is any.
func (c *FakeAwsVpcEndpointServices) Create(awsVpcEndpointService *v1alpha1.AwsVpcEndpointService) (result *v1alpha1.AwsVpcEndpointService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsvpcendpointservicesResource, awsVpcEndpointService), &v1alpha1.AwsVpcEndpointService{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVpcEndpointService), err
}

// Update takes the representation of a awsVpcEndpointService and updates it. Returns the server's representation of the awsVpcEndpointService, and an error, if there is any.
func (c *FakeAwsVpcEndpointServices) Update(awsVpcEndpointService *v1alpha1.AwsVpcEndpointService) (result *v1alpha1.AwsVpcEndpointService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsvpcendpointservicesResource, awsVpcEndpointService), &v1alpha1.AwsVpcEndpointService{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVpcEndpointService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsVpcEndpointServices) UpdateStatus(awsVpcEndpointService *v1alpha1.AwsVpcEndpointService) (*v1alpha1.AwsVpcEndpointService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsvpcendpointservicesResource, "status", awsVpcEndpointService), &v1alpha1.AwsVpcEndpointService{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVpcEndpointService), err
}

// Delete takes name of the awsVpcEndpointService and deletes it. Returns an error if one occurs.
func (c *FakeAwsVpcEndpointServices) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsvpcendpointservicesResource, name), &v1alpha1.AwsVpcEndpointService{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsVpcEndpointServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsvpcendpointservicesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsVpcEndpointServiceList{})
	return err
}

// Patch applies the patch and returns the patched awsVpcEndpointService.
func (c *FakeAwsVpcEndpointServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsVpcEndpointService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsvpcendpointservicesResource, name, pt, data, subresources...), &v1alpha1.AwsVpcEndpointService{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVpcEndpointService), err
}
