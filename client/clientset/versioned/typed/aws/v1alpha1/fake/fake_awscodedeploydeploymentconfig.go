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

// FakeAwsCodedeployDeploymentConfigs implements AwsCodedeployDeploymentConfigInterface
type FakeAwsCodedeployDeploymentConfigs struct {
	Fake *FakeAwsV1alpha1
}

var awscodedeploydeploymentconfigsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awscodedeploydeploymentconfigs"}

var awscodedeploydeploymentconfigsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsCodedeployDeploymentConfig"}

// Get takes name of the awsCodedeployDeploymentConfig, and returns the corresponding awsCodedeployDeploymentConfig object, and an error if there is any.
func (c *FakeAwsCodedeployDeploymentConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsCodedeployDeploymentConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awscodedeploydeploymentconfigsResource, name), &v1alpha1.AwsCodedeployDeploymentConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCodedeployDeploymentConfig), err
}

// List takes label and field selectors, and returns the list of AwsCodedeployDeploymentConfigs that match those selectors.
func (c *FakeAwsCodedeployDeploymentConfigs) List(opts v1.ListOptions) (result *v1alpha1.AwsCodedeployDeploymentConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awscodedeploydeploymentconfigsResource, awscodedeploydeploymentconfigsKind, opts), &v1alpha1.AwsCodedeployDeploymentConfigList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsCodedeployDeploymentConfigList{ListMeta: obj.(*v1alpha1.AwsCodedeployDeploymentConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsCodedeployDeploymentConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsCodedeployDeploymentConfigs.
func (c *FakeAwsCodedeployDeploymentConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awscodedeploydeploymentconfigsResource, opts))
}

// Create takes the representation of a awsCodedeployDeploymentConfig and creates it.  Returns the server's representation of the awsCodedeployDeploymentConfig, and an error, if there is any.
func (c *FakeAwsCodedeployDeploymentConfigs) Create(awsCodedeployDeploymentConfig *v1alpha1.AwsCodedeployDeploymentConfig) (result *v1alpha1.AwsCodedeployDeploymentConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awscodedeploydeploymentconfigsResource, awsCodedeployDeploymentConfig), &v1alpha1.AwsCodedeployDeploymentConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCodedeployDeploymentConfig), err
}

// Update takes the representation of a awsCodedeployDeploymentConfig and updates it. Returns the server's representation of the awsCodedeployDeploymentConfig, and an error, if there is any.
func (c *FakeAwsCodedeployDeploymentConfigs) Update(awsCodedeployDeploymentConfig *v1alpha1.AwsCodedeployDeploymentConfig) (result *v1alpha1.AwsCodedeployDeploymentConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awscodedeploydeploymentconfigsResource, awsCodedeployDeploymentConfig), &v1alpha1.AwsCodedeployDeploymentConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCodedeployDeploymentConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsCodedeployDeploymentConfigs) UpdateStatus(awsCodedeployDeploymentConfig *v1alpha1.AwsCodedeployDeploymentConfig) (*v1alpha1.AwsCodedeployDeploymentConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awscodedeploydeploymentconfigsResource, "status", awsCodedeployDeploymentConfig), &v1alpha1.AwsCodedeployDeploymentConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCodedeployDeploymentConfig), err
}

// Delete takes name of the awsCodedeployDeploymentConfig and deletes it. Returns an error if one occurs.
func (c *FakeAwsCodedeployDeploymentConfigs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awscodedeploydeploymentconfigsResource, name), &v1alpha1.AwsCodedeployDeploymentConfig{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsCodedeployDeploymentConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awscodedeploydeploymentconfigsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsCodedeployDeploymentConfigList{})
	return err
}

// Patch applies the patch and returns the patched awsCodedeployDeploymentConfig.
func (c *FakeAwsCodedeployDeploymentConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCodedeployDeploymentConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awscodedeploydeploymentconfigsResource, name, pt, data, subresources...), &v1alpha1.AwsCodedeployDeploymentConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCodedeployDeploymentConfig), err
}
