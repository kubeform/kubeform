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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeGoogleCloudbuildTriggers implements GoogleCloudbuildTriggerInterface
type FakeGoogleCloudbuildTriggers struct {
	Fake *FakeGoogleV1alpha1
}

var googlecloudbuildtriggersResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googlecloudbuildtriggers"}

var googlecloudbuildtriggersKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GoogleCloudbuildTrigger"}

// Get takes name of the googleCloudbuildTrigger, and returns the corresponding googleCloudbuildTrigger object, and an error if there is any.
func (c *FakeGoogleCloudbuildTriggers) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleCloudbuildTrigger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googlecloudbuildtriggersResource, name), &v1alpha1.GoogleCloudbuildTrigger{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleCloudbuildTrigger), err
}

// List takes label and field selectors, and returns the list of GoogleCloudbuildTriggers that match those selectors.
func (c *FakeGoogleCloudbuildTriggers) List(opts v1.ListOptions) (result *v1alpha1.GoogleCloudbuildTriggerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googlecloudbuildtriggersResource, googlecloudbuildtriggersKind, opts), &v1alpha1.GoogleCloudbuildTriggerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleCloudbuildTriggerList{ListMeta: obj.(*v1alpha1.GoogleCloudbuildTriggerList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleCloudbuildTriggerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleCloudbuildTriggers.
func (c *FakeGoogleCloudbuildTriggers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googlecloudbuildtriggersResource, opts))
}

// Create takes the representation of a googleCloudbuildTrigger and creates it.  Returns the server's representation of the googleCloudbuildTrigger, and an error, if there is any.
func (c *FakeGoogleCloudbuildTriggers) Create(googleCloudbuildTrigger *v1alpha1.GoogleCloudbuildTrigger) (result *v1alpha1.GoogleCloudbuildTrigger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googlecloudbuildtriggersResource, googleCloudbuildTrigger), &v1alpha1.GoogleCloudbuildTrigger{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleCloudbuildTrigger), err
}

// Update takes the representation of a googleCloudbuildTrigger and updates it. Returns the server's representation of the googleCloudbuildTrigger, and an error, if there is any.
func (c *FakeGoogleCloudbuildTriggers) Update(googleCloudbuildTrigger *v1alpha1.GoogleCloudbuildTrigger) (result *v1alpha1.GoogleCloudbuildTrigger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googlecloudbuildtriggersResource, googleCloudbuildTrigger), &v1alpha1.GoogleCloudbuildTrigger{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleCloudbuildTrigger), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleCloudbuildTriggers) UpdateStatus(googleCloudbuildTrigger *v1alpha1.GoogleCloudbuildTrigger) (*v1alpha1.GoogleCloudbuildTrigger, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googlecloudbuildtriggersResource, "status", googleCloudbuildTrigger), &v1alpha1.GoogleCloudbuildTrigger{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleCloudbuildTrigger), err
}

// Delete takes name of the googleCloudbuildTrigger and deletes it. Returns an error if one occurs.
func (c *FakeGoogleCloudbuildTriggers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googlecloudbuildtriggersResource, name), &v1alpha1.GoogleCloudbuildTrigger{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleCloudbuildTriggers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googlecloudbuildtriggersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleCloudbuildTriggerList{})
	return err
}

// Patch applies the patch and returns the patched googleCloudbuildTrigger.
func (c *FakeGoogleCloudbuildTriggers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleCloudbuildTrigger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googlecloudbuildtriggersResource, name, pt, data, subresources...), &v1alpha1.GoogleCloudbuildTrigger{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleCloudbuildTrigger), err
}
