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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePolicyDefinitions implements PolicyDefinitionInterface
type FakePolicyDefinitions struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var policydefinitionsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "policydefinitions"}

var policydefinitionsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "PolicyDefinition"}

// Get takes name of the policyDefinition, and returns the corresponding policyDefinition object, and an error if there is any.
func (c *FakePolicyDefinitions) Get(name string, options v1.GetOptions) (result *v1alpha1.PolicyDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(policydefinitionsResource, c.ns, name), &v1alpha1.PolicyDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyDefinition), err
}

// List takes label and field selectors, and returns the list of PolicyDefinitions that match those selectors.
func (c *FakePolicyDefinitions) List(opts v1.ListOptions) (result *v1alpha1.PolicyDefinitionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(policydefinitionsResource, policydefinitionsKind, c.ns, opts), &v1alpha1.PolicyDefinitionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PolicyDefinitionList{ListMeta: obj.(*v1alpha1.PolicyDefinitionList).ListMeta}
	for _, item := range obj.(*v1alpha1.PolicyDefinitionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested policyDefinitions.
func (c *FakePolicyDefinitions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(policydefinitionsResource, c.ns, opts))

}

// Create takes the representation of a policyDefinition and creates it.  Returns the server's representation of the policyDefinition, and an error, if there is any.
func (c *FakePolicyDefinitions) Create(policyDefinition *v1alpha1.PolicyDefinition) (result *v1alpha1.PolicyDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(policydefinitionsResource, c.ns, policyDefinition), &v1alpha1.PolicyDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyDefinition), err
}

// Update takes the representation of a policyDefinition and updates it. Returns the server's representation of the policyDefinition, and an error, if there is any.
func (c *FakePolicyDefinitions) Update(policyDefinition *v1alpha1.PolicyDefinition) (result *v1alpha1.PolicyDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(policydefinitionsResource, c.ns, policyDefinition), &v1alpha1.PolicyDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyDefinition), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePolicyDefinitions) UpdateStatus(policyDefinition *v1alpha1.PolicyDefinition) (*v1alpha1.PolicyDefinition, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(policydefinitionsResource, "status", c.ns, policyDefinition), &v1alpha1.PolicyDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyDefinition), err
}

// Delete takes name of the policyDefinition and deletes it. Returns an error if one occurs.
func (c *FakePolicyDefinitions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(policydefinitionsResource, c.ns, name), &v1alpha1.PolicyDefinition{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePolicyDefinitions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(policydefinitionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PolicyDefinitionList{})
	return err
}

// Patch applies the patch and returns the patched policyDefinition.
func (c *FakePolicyDefinitions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PolicyDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(policydefinitionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PolicyDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyDefinition), err
}
