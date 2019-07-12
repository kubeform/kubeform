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

// FakeAwsLambdaAliases implements AwsLambdaAliasInterface
type FakeAwsLambdaAliases struct {
	Fake *FakeAwsV1alpha1
}

var awslambdaaliasesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awslambdaaliases"}

var awslambdaaliasesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsLambdaAlias"}

// Get takes name of the awsLambdaAlias, and returns the corresponding awsLambdaAlias object, and an error if there is any.
func (c *FakeAwsLambdaAliases) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsLambdaAlias, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awslambdaaliasesResource, name), &v1alpha1.AwsLambdaAlias{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLambdaAlias), err
}

// List takes label and field selectors, and returns the list of AwsLambdaAliases that match those selectors.
func (c *FakeAwsLambdaAliases) List(opts v1.ListOptions) (result *v1alpha1.AwsLambdaAliasList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awslambdaaliasesResource, awslambdaaliasesKind, opts), &v1alpha1.AwsLambdaAliasList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsLambdaAliasList{ListMeta: obj.(*v1alpha1.AwsLambdaAliasList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsLambdaAliasList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsLambdaAliases.
func (c *FakeAwsLambdaAliases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awslambdaaliasesResource, opts))
}

// Create takes the representation of a awsLambdaAlias and creates it.  Returns the server's representation of the awsLambdaAlias, and an error, if there is any.
func (c *FakeAwsLambdaAliases) Create(awsLambdaAlias *v1alpha1.AwsLambdaAlias) (result *v1alpha1.AwsLambdaAlias, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awslambdaaliasesResource, awsLambdaAlias), &v1alpha1.AwsLambdaAlias{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLambdaAlias), err
}

// Update takes the representation of a awsLambdaAlias and updates it. Returns the server's representation of the awsLambdaAlias, and an error, if there is any.
func (c *FakeAwsLambdaAliases) Update(awsLambdaAlias *v1alpha1.AwsLambdaAlias) (result *v1alpha1.AwsLambdaAlias, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awslambdaaliasesResource, awsLambdaAlias), &v1alpha1.AwsLambdaAlias{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLambdaAlias), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsLambdaAliases) UpdateStatus(awsLambdaAlias *v1alpha1.AwsLambdaAlias) (*v1alpha1.AwsLambdaAlias, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awslambdaaliasesResource, "status", awsLambdaAlias), &v1alpha1.AwsLambdaAlias{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLambdaAlias), err
}

// Delete takes name of the awsLambdaAlias and deletes it. Returns an error if one occurs.
func (c *FakeAwsLambdaAliases) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awslambdaaliasesResource, name), &v1alpha1.AwsLambdaAlias{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsLambdaAliases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awslambdaaliasesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsLambdaAliasList{})
	return err
}

// Patch applies the patch and returns the patched awsLambdaAlias.
func (c *FakeAwsLambdaAliases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsLambdaAlias, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awslambdaaliasesResource, name, pt, data, subresources...), &v1alpha1.AwsLambdaAlias{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLambdaAlias), err
}
