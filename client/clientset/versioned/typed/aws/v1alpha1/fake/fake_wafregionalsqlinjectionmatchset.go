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

// FakeWafregionalSqlInjectionMatchSets implements WafregionalSqlInjectionMatchSetInterface
type FakeWafregionalSqlInjectionMatchSets struct {
	Fake *FakeAwsV1alpha1
}

var wafregionalsqlinjectionmatchsetsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "wafregionalsqlinjectionmatchsets"}

var wafregionalsqlinjectionmatchsetsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "WafregionalSqlInjectionMatchSet"}

// Get takes name of the wafregionalSqlInjectionMatchSet, and returns the corresponding wafregionalSqlInjectionMatchSet object, and an error if there is any.
func (c *FakeWafregionalSqlInjectionMatchSets) Get(name string, options v1.GetOptions) (result *v1alpha1.WafregionalSqlInjectionMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(wafregionalsqlinjectionmatchsetsResource, name), &v1alpha1.WafregionalSqlInjectionMatchSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafregionalSqlInjectionMatchSet), err
}

// List takes label and field selectors, and returns the list of WafregionalSqlInjectionMatchSets that match those selectors.
func (c *FakeWafregionalSqlInjectionMatchSets) List(opts v1.ListOptions) (result *v1alpha1.WafregionalSqlInjectionMatchSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(wafregionalsqlinjectionmatchsetsResource, wafregionalsqlinjectionmatchsetsKind, opts), &v1alpha1.WafregionalSqlInjectionMatchSetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.WafregionalSqlInjectionMatchSetList{ListMeta: obj.(*v1alpha1.WafregionalSqlInjectionMatchSetList).ListMeta}
	for _, item := range obj.(*v1alpha1.WafregionalSqlInjectionMatchSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested wafregionalSqlInjectionMatchSets.
func (c *FakeWafregionalSqlInjectionMatchSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(wafregionalsqlinjectionmatchsetsResource, opts))
}

// Create takes the representation of a wafregionalSqlInjectionMatchSet and creates it.  Returns the server's representation of the wafregionalSqlInjectionMatchSet, and an error, if there is any.
func (c *FakeWafregionalSqlInjectionMatchSets) Create(wafregionalSqlInjectionMatchSet *v1alpha1.WafregionalSqlInjectionMatchSet) (result *v1alpha1.WafregionalSqlInjectionMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(wafregionalsqlinjectionmatchsetsResource, wafregionalSqlInjectionMatchSet), &v1alpha1.WafregionalSqlInjectionMatchSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafregionalSqlInjectionMatchSet), err
}

// Update takes the representation of a wafregionalSqlInjectionMatchSet and updates it. Returns the server's representation of the wafregionalSqlInjectionMatchSet, and an error, if there is any.
func (c *FakeWafregionalSqlInjectionMatchSets) Update(wafregionalSqlInjectionMatchSet *v1alpha1.WafregionalSqlInjectionMatchSet) (result *v1alpha1.WafregionalSqlInjectionMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(wafregionalsqlinjectionmatchsetsResource, wafregionalSqlInjectionMatchSet), &v1alpha1.WafregionalSqlInjectionMatchSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafregionalSqlInjectionMatchSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeWafregionalSqlInjectionMatchSets) UpdateStatus(wafregionalSqlInjectionMatchSet *v1alpha1.WafregionalSqlInjectionMatchSet) (*v1alpha1.WafregionalSqlInjectionMatchSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(wafregionalsqlinjectionmatchsetsResource, "status", wafregionalSqlInjectionMatchSet), &v1alpha1.WafregionalSqlInjectionMatchSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafregionalSqlInjectionMatchSet), err
}

// Delete takes name of the wafregionalSqlInjectionMatchSet and deletes it. Returns an error if one occurs.
func (c *FakeWafregionalSqlInjectionMatchSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(wafregionalsqlinjectionmatchsetsResource, name), &v1alpha1.WafregionalSqlInjectionMatchSet{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWafregionalSqlInjectionMatchSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(wafregionalsqlinjectionmatchsetsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.WafregionalSqlInjectionMatchSetList{})
	return err
}

// Patch applies the patch and returns the patched wafregionalSqlInjectionMatchSet.
func (c *FakeWafregionalSqlInjectionMatchSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WafregionalSqlInjectionMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(wafregionalsqlinjectionmatchsetsResource, name, pt, data, subresources...), &v1alpha1.WafregionalSqlInjectionMatchSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafregionalSqlInjectionMatchSet), err
}
