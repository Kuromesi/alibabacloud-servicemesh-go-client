// Copyright 2025 Alibaba Cloud Service Mesh
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.
// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	scheme "istio.io/client-go/asm/pkg/clientset/scheme"
	"context"
	"time"

	v1 "istio.io/api/alibabacloudservicemesh/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ASMDecompressorsGetter has a method to return a ASMDecompressorInterface.
// A group's client should implement this interface.
type ASMDecompressorsGetter interface {
	ASMDecompressors(namespace string) ASMDecompressorInterface
}

// ASMDecompressorInterface has methods to work with ASMDecompressor resources.
type ASMDecompressorInterface interface {
	Create(ctx context.Context, aSMDecompressor *v1.ASMDecompressor, opts metav1.CreateOptions) (*v1.ASMDecompressor, error)
	Update(ctx context.Context, aSMDecompressor *v1.ASMDecompressor, opts metav1.UpdateOptions) (*v1.ASMDecompressor, error)
	UpdateStatus(ctx context.Context, aSMDecompressor *v1.ASMDecompressor, opts metav1.UpdateOptions) (*v1.ASMDecompressor, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ASMDecompressor, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ASMDecompressorList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ASMDecompressor, err error)
	ASMDecompressorExpansion
}

// aSMDecompressors implements ASMDecompressorInterface
type aSMDecompressors struct {
	client rest.Interface
	ns     string
}

// newASMDecompressors returns a ASMDecompressors
func newASMDecompressors(c *IstioV1Client, namespace string) *aSMDecompressors {
	return &aSMDecompressors{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the aSMDecompressor, and returns the corresponding aSMDecompressor object, and an error if there is any.
func (c *aSMDecompressors) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ASMDecompressor, err error) {
	result = &v1.ASMDecompressor{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("asmdecompressors").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ASMDecompressors that match those selectors.
func (c *aSMDecompressors) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ASMDecompressorList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ASMDecompressorList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("asmdecompressors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested aSMDecompressors.
func (c *aSMDecompressors) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("asmdecompressors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a aSMDecompressor and creates it.  Returns the server's representation of the aSMDecompressor, and an error, if there is any.
func (c *aSMDecompressors) Create(ctx context.Context, aSMDecompressor *v1.ASMDecompressor, opts metav1.CreateOptions) (result *v1.ASMDecompressor, err error) {
	result = &v1.ASMDecompressor{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("asmdecompressors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aSMDecompressor).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a aSMDecompressor and updates it. Returns the server's representation of the aSMDecompressor, and an error, if there is any.
func (c *aSMDecompressors) Update(ctx context.Context, aSMDecompressor *v1.ASMDecompressor, opts metav1.UpdateOptions) (result *v1.ASMDecompressor, err error) {
	result = &v1.ASMDecompressor{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("asmdecompressors").
		Name(aSMDecompressor.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aSMDecompressor).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *aSMDecompressors) UpdateStatus(ctx context.Context, aSMDecompressor *v1.ASMDecompressor, opts metav1.UpdateOptions) (result *v1.ASMDecompressor, err error) {
	result = &v1.ASMDecompressor{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("asmdecompressors").
		Name(aSMDecompressor.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aSMDecompressor).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the aSMDecompressor and deletes it. Returns an error if one occurs.
func (c *aSMDecompressors) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("asmdecompressors").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *aSMDecompressors) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("asmdecompressors").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched aSMDecompressor.
func (c *aSMDecompressors) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ASMDecompressor, err error) {
	result = &v1.ASMDecompressor{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("asmdecompressors").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
