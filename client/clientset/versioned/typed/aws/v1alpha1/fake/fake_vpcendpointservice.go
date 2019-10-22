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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVpcEndpointServices implements VpcEndpointServiceInterface
type FakeVpcEndpointServices struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var vpcendpointservicesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "vpcendpointservices"}

var vpcendpointservicesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "VpcEndpointService"}

// Get takes name of the vpcEndpointService, and returns the corresponding vpcEndpointService object, and an error if there is any.
func (c *FakeVpcEndpointServices) Get(name string, options v1.GetOptions) (result *v1alpha1.VpcEndpointService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vpcendpointservicesResource, c.ns, name), &v1alpha1.VpcEndpointService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcEndpointService), err
}

// List takes label and field selectors, and returns the list of VpcEndpointServices that match those selectors.
func (c *FakeVpcEndpointServices) List(opts v1.ListOptions) (result *v1alpha1.VpcEndpointServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vpcendpointservicesResource, vpcendpointservicesKind, c.ns, opts), &v1alpha1.VpcEndpointServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VpcEndpointServiceList{ListMeta: obj.(*v1alpha1.VpcEndpointServiceList).ListMeta}
	for _, item := range obj.(*v1alpha1.VpcEndpointServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vpcEndpointServices.
func (c *FakeVpcEndpointServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vpcendpointservicesResource, c.ns, opts))

}

// Create takes the representation of a vpcEndpointService and creates it.  Returns the server's representation of the vpcEndpointService, and an error, if there is any.
func (c *FakeVpcEndpointServices) Create(vpcEndpointService *v1alpha1.VpcEndpointService) (result *v1alpha1.VpcEndpointService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vpcendpointservicesResource, c.ns, vpcEndpointService), &v1alpha1.VpcEndpointService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcEndpointService), err
}

// Update takes the representation of a vpcEndpointService and updates it. Returns the server's representation of the vpcEndpointService, and an error, if there is any.
func (c *FakeVpcEndpointServices) Update(vpcEndpointService *v1alpha1.VpcEndpointService) (result *v1alpha1.VpcEndpointService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vpcendpointservicesResource, c.ns, vpcEndpointService), &v1alpha1.VpcEndpointService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcEndpointService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVpcEndpointServices) UpdateStatus(vpcEndpointService *v1alpha1.VpcEndpointService) (*v1alpha1.VpcEndpointService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vpcendpointservicesResource, "status", c.ns, vpcEndpointService), &v1alpha1.VpcEndpointService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcEndpointService), err
}

// Delete takes name of the vpcEndpointService and deletes it. Returns an error if one occurs.
func (c *FakeVpcEndpointServices) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(vpcendpointservicesResource, c.ns, name), &v1alpha1.VpcEndpointService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVpcEndpointServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vpcendpointservicesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.VpcEndpointServiceList{})
	return err
}

// Patch applies the patch and returns the patched vpcEndpointService.
func (c *FakeVpcEndpointServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VpcEndpointService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vpcendpointservicesResource, c.ns, name, pt, data, subresources...), &v1alpha1.VpcEndpointService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcEndpointService), err
}
