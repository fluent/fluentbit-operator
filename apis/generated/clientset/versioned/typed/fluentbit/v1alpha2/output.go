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

package v1alpha2

import (
	"context"
	"time"

	v1alpha2 "github.com/fluent/fluent-operator/apis/fluentbit/v1alpha2"
	scheme "github.com/fluent/fluent-operator/apis/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// OutputsGetter has a method to return a OutputInterface.
// A group's client should implement this interface.
type OutputsGetter interface {
	Outputs(namespace string) OutputInterface
}

// OutputInterface has methods to work with Output resources.
type OutputInterface interface {
	Create(ctx context.Context, output *v1alpha2.Output, opts v1.CreateOptions) (*v1alpha2.Output, error)
	Update(ctx context.Context, output *v1alpha2.Output, opts v1.UpdateOptions) (*v1alpha2.Output, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha2.Output, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha2.OutputList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.Output, err error)
	OutputExpansion
}

// outputs implements OutputInterface
type outputs struct {
	client rest.Interface
	ns     string
}

// newOutputs returns a Outputs
func newOutputs(c *FluentbitV1alpha2Client, namespace string) *outputs {
	return &outputs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the output, and returns the corresponding output object, and an error if there is any.
func (c *outputs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.Output, err error) {
	result = &v1alpha2.Output{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("outputs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Outputs that match those selectors.
func (c *outputs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.OutputList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.OutputList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("outputs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested outputs.
func (c *outputs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("outputs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a output and creates it.  Returns the server's representation of the output, and an error, if there is any.
func (c *outputs) Create(ctx context.Context, output *v1alpha2.Output, opts v1.CreateOptions) (result *v1alpha2.Output, err error) {
	result = &v1alpha2.Output{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("outputs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(output).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a output and updates it. Returns the server's representation of the output, and an error, if there is any.
func (c *outputs) Update(ctx context.Context, output *v1alpha2.Output, opts v1.UpdateOptions) (result *v1alpha2.Output, err error) {
	result = &v1alpha2.Output{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("outputs").
		Name(output.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(output).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the output and deletes it. Returns an error if one occurs.
func (c *outputs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("outputs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *outputs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("outputs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched output.
func (c *outputs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.Output, err error) {
	result = &v1alpha2.Output{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("outputs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
