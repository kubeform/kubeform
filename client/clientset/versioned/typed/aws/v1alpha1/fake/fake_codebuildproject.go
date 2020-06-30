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
	"context"

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCodebuildProjects implements CodebuildProjectInterface
type FakeCodebuildProjects struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var codebuildprojectsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "codebuildprojects"}

var codebuildprojectsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "CodebuildProject"}

// Get takes name of the codebuildProject, and returns the corresponding codebuildProject object, and an error if there is any.
func (c *FakeCodebuildProjects) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CodebuildProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(codebuildprojectsResource, c.ns, name), &v1alpha1.CodebuildProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CodebuildProject), err
}

// List takes label and field selectors, and returns the list of CodebuildProjects that match those selectors.
func (c *FakeCodebuildProjects) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CodebuildProjectList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(codebuildprojectsResource, codebuildprojectsKind, c.ns, opts), &v1alpha1.CodebuildProjectList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CodebuildProjectList{ListMeta: obj.(*v1alpha1.CodebuildProjectList).ListMeta}
	for _, item := range obj.(*v1alpha1.CodebuildProjectList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested codebuildProjects.
func (c *FakeCodebuildProjects) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(codebuildprojectsResource, c.ns, opts))

}

// Create takes the representation of a codebuildProject and creates it.  Returns the server's representation of the codebuildProject, and an error, if there is any.
func (c *FakeCodebuildProjects) Create(ctx context.Context, codebuildProject *v1alpha1.CodebuildProject, opts v1.CreateOptions) (result *v1alpha1.CodebuildProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(codebuildprojectsResource, c.ns, codebuildProject), &v1alpha1.CodebuildProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CodebuildProject), err
}

// Update takes the representation of a codebuildProject and updates it. Returns the server's representation of the codebuildProject, and an error, if there is any.
func (c *FakeCodebuildProjects) Update(ctx context.Context, codebuildProject *v1alpha1.CodebuildProject, opts v1.UpdateOptions) (result *v1alpha1.CodebuildProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(codebuildprojectsResource, c.ns, codebuildProject), &v1alpha1.CodebuildProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CodebuildProject), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCodebuildProjects) UpdateStatus(ctx context.Context, codebuildProject *v1alpha1.CodebuildProject, opts v1.UpdateOptions) (*v1alpha1.CodebuildProject, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(codebuildprojectsResource, "status", c.ns, codebuildProject), &v1alpha1.CodebuildProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CodebuildProject), err
}

// Delete takes name of the codebuildProject and deletes it. Returns an error if one occurs.
func (c *FakeCodebuildProjects) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(codebuildprojectsResource, c.ns, name), &v1alpha1.CodebuildProject{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCodebuildProjects) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(codebuildprojectsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CodebuildProjectList{})
	return err
}

// Patch applies the patch and returns the patched codebuildProject.
func (c *FakeCodebuildProjects) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CodebuildProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(codebuildprojectsResource, c.ns, name, pt, data, subresources...), &v1alpha1.CodebuildProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CodebuildProject), err
}
