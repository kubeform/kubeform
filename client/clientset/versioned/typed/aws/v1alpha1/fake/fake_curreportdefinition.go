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

// FakeCurReportDefinitions implements CurReportDefinitionInterface
type FakeCurReportDefinitions struct {
	Fake *FakeAwsV1alpha1
}

var curreportdefinitionsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "curreportdefinitions"}

var curreportdefinitionsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "CurReportDefinition"}

// Get takes name of the curReportDefinition, and returns the corresponding curReportDefinition object, and an error if there is any.
func (c *FakeCurReportDefinitions) Get(name string, options v1.GetOptions) (result *v1alpha1.CurReportDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(curreportdefinitionsResource, name), &v1alpha1.CurReportDefinition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CurReportDefinition), err
}

// List takes label and field selectors, and returns the list of CurReportDefinitions that match those selectors.
func (c *FakeCurReportDefinitions) List(opts v1.ListOptions) (result *v1alpha1.CurReportDefinitionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(curreportdefinitionsResource, curreportdefinitionsKind, opts), &v1alpha1.CurReportDefinitionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CurReportDefinitionList{ListMeta: obj.(*v1alpha1.CurReportDefinitionList).ListMeta}
	for _, item := range obj.(*v1alpha1.CurReportDefinitionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested curReportDefinitions.
func (c *FakeCurReportDefinitions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(curreportdefinitionsResource, opts))
}

// Create takes the representation of a curReportDefinition and creates it.  Returns the server's representation of the curReportDefinition, and an error, if there is any.
func (c *FakeCurReportDefinitions) Create(curReportDefinition *v1alpha1.CurReportDefinition) (result *v1alpha1.CurReportDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(curreportdefinitionsResource, curReportDefinition), &v1alpha1.CurReportDefinition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CurReportDefinition), err
}

// Update takes the representation of a curReportDefinition and updates it. Returns the server's representation of the curReportDefinition, and an error, if there is any.
func (c *FakeCurReportDefinitions) Update(curReportDefinition *v1alpha1.CurReportDefinition) (result *v1alpha1.CurReportDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(curreportdefinitionsResource, curReportDefinition), &v1alpha1.CurReportDefinition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CurReportDefinition), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCurReportDefinitions) UpdateStatus(curReportDefinition *v1alpha1.CurReportDefinition) (*v1alpha1.CurReportDefinition, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(curreportdefinitionsResource, "status", curReportDefinition), &v1alpha1.CurReportDefinition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CurReportDefinition), err
}

// Delete takes name of the curReportDefinition and deletes it. Returns an error if one occurs.
func (c *FakeCurReportDefinitions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(curreportdefinitionsResource, name), &v1alpha1.CurReportDefinition{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCurReportDefinitions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(curreportdefinitionsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CurReportDefinitionList{})
	return err
}

// Patch applies the patch and returns the patched curReportDefinition.
func (c *FakeCurReportDefinitions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CurReportDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(curreportdefinitionsResource, name, pt, data, subresources...), &v1alpha1.CurReportDefinition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CurReportDefinition), err
}
