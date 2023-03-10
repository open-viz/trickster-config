//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	reloadoptions "github.com/trickstercache/trickster/v2/cmd/trickster/config/reload/options"
	"github.com/trickstercache/trickster/v2/pkg/cache/negative"
	"github.com/trickstercache/trickster/v2/pkg/observability/logging/options"
	metricsoptions "github.com/trickstercache/trickster/v2/pkg/observability/metrics/options"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Trickster) DeepCopyInto(out *Trickster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Trickster.
func (in *Trickster) DeepCopy() *Trickster {
	if in == nil {
		return nil
	}
	out := new(Trickster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Trickster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TricksterBackend) DeepCopyInto(out *TricksterBackend) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TricksterBackend.
func (in *TricksterBackend) DeepCopy() *TricksterBackend {
	if in == nil {
		return nil
	}
	out := new(TricksterBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TricksterBackend) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TricksterBackendList) DeepCopyInto(out *TricksterBackendList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TricksterBackend, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TricksterBackendList.
func (in *TricksterBackendList) DeepCopy() *TricksterBackendList {
	if in == nil {
		return nil
	}
	out := new(TricksterBackendList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TricksterBackendList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TricksterBackendSpec) DeepCopyInto(out *TricksterBackendSpec) {
	*out = *in
	in.Options.DeepCopyInto(&out.Options)
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(v1.SecretProjection)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TricksterBackendSpec.
func (in *TricksterBackendSpec) DeepCopy() *TricksterBackendSpec {
	if in == nil {
		return nil
	}
	out := new(TricksterBackendSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TricksterBackendStatus) DeepCopyInto(out *TricksterBackendStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TricksterBackendStatus.
func (in *TricksterBackendStatus) DeepCopy() *TricksterBackendStatus {
	if in == nil {
		return nil
	}
	out := new(TricksterBackendStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TricksterCache) DeepCopyInto(out *TricksterCache) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TricksterCache.
func (in *TricksterCache) DeepCopy() *TricksterCache {
	if in == nil {
		return nil
	}
	out := new(TricksterCache)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TricksterCache) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TricksterCacheList) DeepCopyInto(out *TricksterCacheList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TricksterCache, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TricksterCacheList.
func (in *TricksterCacheList) DeepCopy() *TricksterCacheList {
	if in == nil {
		return nil
	}
	out := new(TricksterCacheList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TricksterCacheList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TricksterCacheSpec) DeepCopyInto(out *TricksterCacheSpec) {
	*out = *in
	in.Options.DeepCopyInto(&out.Options)
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(v1.SecretProjection)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TricksterCacheSpec.
func (in *TricksterCacheSpec) DeepCopy() *TricksterCacheSpec {
	if in == nil {
		return nil
	}
	out := new(TricksterCacheSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TricksterCacheStatus) DeepCopyInto(out *TricksterCacheStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TricksterCacheStatus.
func (in *TricksterCacheStatus) DeepCopy() *TricksterCacheStatus {
	if in == nil {
		return nil
	}
	out := new(TricksterCacheStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TricksterList) DeepCopyInto(out *TricksterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Trickster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TricksterList.
func (in *TricksterList) DeepCopy() *TricksterList {
	if in == nil {
		return nil
	}
	out := new(TricksterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TricksterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TricksterRequestRewriter) DeepCopyInto(out *TricksterRequestRewriter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TricksterRequestRewriter.
func (in *TricksterRequestRewriter) DeepCopy() *TricksterRequestRewriter {
	if in == nil {
		return nil
	}
	out := new(TricksterRequestRewriter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TricksterRequestRewriter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TricksterRequestRewriterList) DeepCopyInto(out *TricksterRequestRewriterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TricksterRequestRewriter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TricksterRequestRewriterList.
func (in *TricksterRequestRewriterList) DeepCopy() *TricksterRequestRewriterList {
	if in == nil {
		return nil
	}
	out := new(TricksterRequestRewriterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TricksterRequestRewriterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TricksterRequestRewriterSpec) DeepCopyInto(out *TricksterRequestRewriterSpec) {
	*out = *in
	in.Options.DeepCopyInto(&out.Options)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TricksterRequestRewriterSpec.
func (in *TricksterRequestRewriterSpec) DeepCopy() *TricksterRequestRewriterSpec {
	if in == nil {
		return nil
	}
	out := new(TricksterRequestRewriterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TricksterRequestRewriterStatus) DeepCopyInto(out *TricksterRequestRewriterStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TricksterRequestRewriterStatus.
func (in *TricksterRequestRewriterStatus) DeepCopy() *TricksterRequestRewriterStatus {
	if in == nil {
		return nil
	}
	out := new(TricksterRequestRewriterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TricksterRule) DeepCopyInto(out *TricksterRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TricksterRule.
func (in *TricksterRule) DeepCopy() *TricksterRule {
	if in == nil {
		return nil
	}
	out := new(TricksterRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TricksterRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TricksterRuleList) DeepCopyInto(out *TricksterRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TricksterRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TricksterRuleList.
func (in *TricksterRuleList) DeepCopy() *TricksterRuleList {
	if in == nil {
		return nil
	}
	out := new(TricksterRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TricksterRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TricksterRuleSpec) DeepCopyInto(out *TricksterRuleSpec) {
	*out = *in
	in.Options.DeepCopyInto(&out.Options)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TricksterRuleSpec.
func (in *TricksterRuleSpec) DeepCopy() *TricksterRuleSpec {
	if in == nil {
		return nil
	}
	out := new(TricksterRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TricksterRuleStatus) DeepCopyInto(out *TricksterRuleStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TricksterRuleStatus.
func (in *TricksterRuleStatus) DeepCopy() *TricksterRuleStatus {
	if in == nil {
		return nil
	}
	out := new(TricksterRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TricksterSpec) DeepCopyInto(out *TricksterSpec) {
	*out = *in
	if in.Main != nil {
		in, out := &in.Main, &out.Main
		*out = (*in).DeepCopy()
	}
	if in.Nats != nil {
		in, out := &in.Nats, &out.Nats
		*out = (*in).DeepCopy()
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(v1.SecretProjection)
		(*in).DeepCopyInto(*out)
	}
	if in.BackendSelector != nil {
		in, out := &in.BackendSelector, &out.BackendSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.CacheSelector != nil {
		in, out := &in.CacheSelector, &out.CacheSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Frontend != nil {
		in, out := &in.Frontend, &out.Frontend
		*out = (*in).DeepCopy()
	}
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(options.Options)
		**out = **in
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = new(metricsoptions.Options)
		**out = **in
	}
	if in.TracingConfigSelector != nil {
		in, out := &in.TracingConfigSelector, &out.TracingConfigSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NegativeCacheConfigs != nil {
		in, out := &in.NegativeCacheConfigs, &out.NegativeCacheConfigs
		*out = make(map[string]negative.Config, len(*in))
		for key, val := range *in {
			var outVal map[string]int
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(negative.Config, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.RuleSelector != nil {
		in, out := &in.RuleSelector, &out.RuleSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.RequestRewriterSelector != nil {
		in, out := &in.RequestRewriterSelector, &out.RequestRewriterSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ReloadConfig != nil {
		in, out := &in.ReloadConfig, &out.ReloadConfig
		*out = new(reloadoptions.Options)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TricksterSpec.
func (in *TricksterSpec) DeepCopy() *TricksterSpec {
	if in == nil {
		return nil
	}
	out := new(TricksterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TricksterStatus) DeepCopyInto(out *TricksterStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TricksterStatus.
func (in *TricksterStatus) DeepCopy() *TricksterStatus {
	if in == nil {
		return nil
	}
	out := new(TricksterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TricksterTracingConfig) DeepCopyInto(out *TricksterTracingConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TricksterTracingConfig.
func (in *TricksterTracingConfig) DeepCopy() *TricksterTracingConfig {
	if in == nil {
		return nil
	}
	out := new(TricksterTracingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TricksterTracingConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TricksterTracingConfigList) DeepCopyInto(out *TricksterTracingConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TricksterTracingConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TricksterTracingConfigList.
func (in *TricksterTracingConfigList) DeepCopy() *TricksterTracingConfigList {
	if in == nil {
		return nil
	}
	out := new(TricksterTracingConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TricksterTracingConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TricksterTracingConfigSpec) DeepCopyInto(out *TricksterTracingConfigSpec) {
	*out = *in
	in.Options.DeepCopyInto(&out.Options)
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(v1.SecretProjection)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TricksterTracingConfigSpec.
func (in *TricksterTracingConfigSpec) DeepCopy() *TricksterTracingConfigSpec {
	if in == nil {
		return nil
	}
	out := new(TricksterTracingConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TricksterTracingConfigStatus) DeepCopyInto(out *TricksterTracingConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TricksterTracingConfigStatus.
func (in *TricksterTracingConfigStatus) DeepCopy() *TricksterTracingConfigStatus {
	if in == nil {
		return nil
	}
	out := new(TricksterTracingConfigStatus)
	in.DeepCopyInto(out)
	return out
}
