// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeApiserverConfig) DeepCopyInto(out *KubeApiserverConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeApiserverConfig.
func (in *KubeApiserverConfig) DeepCopy() *KubeApiserverConfig {
	if in == nil {
		return nil
	}
	out := new(KubeApiserverConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeApiserverConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeApiserverOperatorConfig) DeepCopyInto(out *KubeApiserverOperatorConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeApiserverOperatorConfig.
func (in *KubeApiserverOperatorConfig) DeepCopy() *KubeApiserverOperatorConfig {
	if in == nil {
		return nil
	}
	out := new(KubeApiserverOperatorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeApiserverOperatorConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeApiserverOperatorConfigList) DeepCopyInto(out *KubeApiserverOperatorConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubeApiserverOperatorConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeApiserverOperatorConfigList.
func (in *KubeApiserverOperatorConfigList) DeepCopy() *KubeApiserverOperatorConfigList {
	if in == nil {
		return nil
	}
	out := new(KubeApiserverOperatorConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeApiserverOperatorConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeApiserverOperatorConfigSpec) DeepCopyInto(out *KubeApiserverOperatorConfigSpec) {
	*out = *in
	out.OperatorSpec = in.OperatorSpec
	in.KubeApiserverConfig.DeepCopyInto(&out.KubeApiserverConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeApiserverOperatorConfigSpec.
func (in *KubeApiserverOperatorConfigSpec) DeepCopy() *KubeApiserverOperatorConfigSpec {
	if in == nil {
		return nil
	}
	out := new(KubeApiserverOperatorConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeApiserverOperatorConfigStatus) DeepCopyInto(out *KubeApiserverOperatorConfigStatus) {
	*out = *in
	in.OperatorStatus.DeepCopyInto(&out.OperatorStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeApiserverOperatorConfigStatus.
func (in *KubeApiserverOperatorConfigStatus) DeepCopy() *KubeApiserverOperatorConfigStatus {
	if in == nil {
		return nil
	}
	out := new(KubeApiserverOperatorConfigStatus)
	in.DeepCopyInto(out)
	return out
}
