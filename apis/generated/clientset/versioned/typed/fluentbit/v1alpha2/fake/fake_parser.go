/*
Copyright 2022.

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

	v1alpha2 "github.com/fluent/fluent-operator/apis/fluentbit/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeParsers implements ParserInterface
type FakeParsers struct {
	Fake *FakeFluentbitV1alpha2
	ns   string
}

var parsersResource = schema.GroupVersionResource{Group: "fluentbit.fluent.io", Version: "v1alpha2", Resource: "parsers"}

var parsersKind = schema.GroupVersionKind{Group: "fluentbit.fluent.io", Version: "v1alpha2", Kind: "Parser"}

// Get takes name of the parser, and returns the corresponding parser object, and an error if there is any.
func (c *FakeParsers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.Parser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(parsersResource, c.ns, name), &v1alpha2.Parser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Parser), err
}

// List takes label and field selectors, and returns the list of Parsers that match those selectors.
func (c *FakeParsers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.ParserList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(parsersResource, parsersKind, c.ns, opts), &v1alpha2.ParserList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.ParserList{ListMeta: obj.(*v1alpha2.ParserList).ListMeta}
	for _, item := range obj.(*v1alpha2.ParserList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested parsers.
func (c *FakeParsers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(parsersResource, c.ns, opts))

}

// Create takes the representation of a parser and creates it.  Returns the server's representation of the parser, and an error, if there is any.
func (c *FakeParsers) Create(ctx context.Context, parser *v1alpha2.Parser, opts v1.CreateOptions) (result *v1alpha2.Parser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(parsersResource, c.ns, parser), &v1alpha2.Parser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Parser), err
}

// Update takes the representation of a parser and updates it. Returns the server's representation of the parser, and an error, if there is any.
func (c *FakeParsers) Update(ctx context.Context, parser *v1alpha2.Parser, opts v1.UpdateOptions) (result *v1alpha2.Parser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(parsersResource, c.ns, parser), &v1alpha2.Parser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Parser), err
}

// Delete takes name of the parser and deletes it. Returns an error if one occurs.
func (c *FakeParsers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(parsersResource, c.ns, name, opts), &v1alpha2.Parser{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeParsers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(parsersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.ParserList{})
	return err
}

// Patch applies the patch and returns the patched parser.
func (c *FakeParsers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.Parser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(parsersResource, c.ns, name, pt, data, subresources...), &v1alpha2.Parser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Parser), err
}
