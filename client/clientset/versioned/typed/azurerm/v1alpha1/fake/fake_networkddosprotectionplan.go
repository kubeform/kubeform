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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeNetworkDdosProtectionPlans implements NetworkDdosProtectionPlanInterface
type FakeNetworkDdosProtectionPlans struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var networkddosprotectionplansResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "networkddosprotectionplans"}

var networkddosprotectionplansKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "NetworkDdosProtectionPlan"}

// Get takes name of the networkDdosProtectionPlan, and returns the corresponding networkDdosProtectionPlan object, and an error if there is any.
func (c *FakeNetworkDdosProtectionPlans) Get(name string, options v1.GetOptions) (result *v1alpha1.NetworkDdosProtectionPlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkddosprotectionplansResource, c.ns, name), &v1alpha1.NetworkDdosProtectionPlan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkDdosProtectionPlan), err
}

// List takes label and field selectors, and returns the list of NetworkDdosProtectionPlans that match those selectors.
func (c *FakeNetworkDdosProtectionPlans) List(opts v1.ListOptions) (result *v1alpha1.NetworkDdosProtectionPlanList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkddosprotectionplansResource, networkddosprotectionplansKind, c.ns, opts), &v1alpha1.NetworkDdosProtectionPlanList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NetworkDdosProtectionPlanList{ListMeta: obj.(*v1alpha1.NetworkDdosProtectionPlanList).ListMeta}
	for _, item := range obj.(*v1alpha1.NetworkDdosProtectionPlanList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkDdosProtectionPlans.
func (c *FakeNetworkDdosProtectionPlans) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkddosprotectionplansResource, c.ns, opts))

}

// Create takes the representation of a networkDdosProtectionPlan and creates it.  Returns the server's representation of the networkDdosProtectionPlan, and an error, if there is any.
func (c *FakeNetworkDdosProtectionPlans) Create(networkDdosProtectionPlan *v1alpha1.NetworkDdosProtectionPlan) (result *v1alpha1.NetworkDdosProtectionPlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkddosprotectionplansResource, c.ns, networkDdosProtectionPlan), &v1alpha1.NetworkDdosProtectionPlan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkDdosProtectionPlan), err
}

// Update takes the representation of a networkDdosProtectionPlan and updates it. Returns the server's representation of the networkDdosProtectionPlan, and an error, if there is any.
func (c *FakeNetworkDdosProtectionPlans) Update(networkDdosProtectionPlan *v1alpha1.NetworkDdosProtectionPlan) (result *v1alpha1.NetworkDdosProtectionPlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkddosprotectionplansResource, c.ns, networkDdosProtectionPlan), &v1alpha1.NetworkDdosProtectionPlan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkDdosProtectionPlan), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkDdosProtectionPlans) UpdateStatus(networkDdosProtectionPlan *v1alpha1.NetworkDdosProtectionPlan) (*v1alpha1.NetworkDdosProtectionPlan, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networkddosprotectionplansResource, "status", c.ns, networkDdosProtectionPlan), &v1alpha1.NetworkDdosProtectionPlan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkDdosProtectionPlan), err
}

// Delete takes name of the networkDdosProtectionPlan and deletes it. Returns an error if one occurs.
func (c *FakeNetworkDdosProtectionPlans) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(networkddosprotectionplansResource, c.ns, name), &v1alpha1.NetworkDdosProtectionPlan{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkDdosProtectionPlans) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkddosprotectionplansResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.NetworkDdosProtectionPlanList{})
	return err
}

// Patch applies the patch and returns the patched networkDdosProtectionPlan.
func (c *FakeNetworkDdosProtectionPlans) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetworkDdosProtectionPlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkddosprotectionplansResource, c.ns, name, pt, data, subresources...), &v1alpha1.NetworkDdosProtectionPlan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkDdosProtectionPlan), err
}
