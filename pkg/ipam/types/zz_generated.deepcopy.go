// +build !ignore_autogenerated

// Copyright 2017-2021 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by deepcopy-gen. DO NOT EDIT.

package types

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocationIP) DeepCopyInto(out *AllocationIP) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocationIP.
func (in *AllocationIP) DeepCopy() *AllocationIP {
	if in == nil {
		return nil
	}
	out := new(AllocationIP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in AllocationMap) DeepCopyInto(out *AllocationMap) {
	{
		in := &in
		*out = make(AllocationMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocationMap.
func (in AllocationMap) DeepCopy() AllocationMap {
	if in == nil {
		return nil
	}
	out := new(AllocationMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAMSpec) DeepCopyInto(out *IPAMSpec) {
	*out = *in
	if in.Pool != nil {
		in, out := &in.Pool, &out.Pool
		*out = make(AllocationMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodCIDRs != nil {
		in, out := &in.PodCIDRs, &out.PodCIDRs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAMSpec.
func (in *IPAMSpec) DeepCopy() *IPAMSpec {
	if in == nil {
		return nil
	}
	out := new(IPAMSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAMStatus) DeepCopyInto(out *IPAMStatus) {
	*out = *in
	if in.Used != nil {
		in, out := &in.Used, &out.Used
		*out = make(AllocationMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.OperatorStatus = in.OperatorStatus
	if in.ReleaseIPs != nil {
		in, out := &in.ReleaseIPs, &out.ReleaseIPs
		*out = make(map[string]IPReleaseStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAMStatus.
func (in *IPAMStatus) DeepCopy() *IPAMStatus {
	if in == nil {
		return nil
	}
	out := new(IPAMStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Limits) DeepCopyInto(out *Limits) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Limits.
func (in *Limits) DeepCopy() *Limits {
	if in == nil {
		return nil
	}
	out := new(Limits)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorStatus) DeepCopyInto(out *OperatorStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorStatus.
func (in *OperatorStatus) DeepCopy() *OperatorStatus {
	if in == nil {
		return nil
	}
	out := new(OperatorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolQuota) DeepCopyInto(out *PoolQuota) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolQuota.
func (in *PoolQuota) DeepCopy() *PoolQuota {
	if in == nil {
		return nil
	}
	out := new(PoolQuota)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in PoolQuotaMap) DeepCopyInto(out *PoolQuotaMap) {
	{
		in := &in
		*out = make(PoolQuotaMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolQuotaMap.
func (in PoolQuotaMap) DeepCopy() PoolQuotaMap {
	if in == nil {
		return nil
	}
	out := new(PoolQuotaMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subnet) DeepCopyInto(out *Subnet) {
	*out = *in
	if in.CIDR != nil {
		in, out := &in.CIDR, &out.CIDR
		*out = (*in).DeepCopy()
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(Tags, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subnet.
func (in *Subnet) DeepCopy() *Subnet {
	if in == nil {
		return nil
	}
	out := new(Subnet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in SubnetMap) DeepCopyInto(out *SubnetMap) {
	{
		in := &in
		*out = make(SubnetMap, len(*in))
		for key, val := range *in {
			var outVal *Subnet
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(Subnet)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetMap.
func (in SubnetMap) DeepCopy() SubnetMap {
	if in == nil {
		return nil
	}
	out := new(SubnetMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Tags) DeepCopyInto(out *Tags) {
	{
		in := &in
		*out = make(Tags, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tags.
func (in Tags) DeepCopy() Tags {
	if in == nil {
		return nil
	}
	out := new(Tags)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNetwork) DeepCopyInto(out *VirtualNetwork) {
	*out = *in
	if in.CIDRs != nil {
		in, out := &in.CIDRs, &out.CIDRs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNetwork.
func (in *VirtualNetwork) DeepCopy() *VirtualNetwork {
	if in == nil {
		return nil
	}
	out := new(VirtualNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in VirtualNetworkMap) DeepCopyInto(out *VirtualNetworkMap) {
	{
		in := &in
		*out = make(VirtualNetworkMap, len(*in))
		for key, val := range *in {
			var outVal *VirtualNetwork
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(VirtualNetwork)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNetworkMap.
func (in VirtualNetworkMap) DeepCopy() VirtualNetworkMap {
	if in == nil {
		return nil
	}
	out := new(VirtualNetworkMap)
	in.DeepCopyInto(out)
	return *out
}
