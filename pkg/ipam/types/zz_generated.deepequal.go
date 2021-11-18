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

// Code generated by main. DO NOT EDIT.

package types

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *AllocationIP) DeepEqual(other *AllocationIP) bool {
	if other == nil {
		return false
	}

	if in.Owner != other.Owner {
		return false
	}
	if in.Resource != other.Resource {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *AllocationMap) DeepEqual(other *AllocationMap) bool {
	if other == nil {
		return false
	}

	if len(*in) != len(*other) {
		return false
	} else {
		for key, inValue := range *in {
			if otherValue, present := (*other)[key]; !present {
				return false
			} else {
				if !inValue.DeepEqual(&otherValue) {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *IPAMSpec) DeepEqual(other *IPAMSpec) bool {
	if other == nil {
		return false
	}

	if ((in.Pool != nil) && (other.Pool != nil)) || ((in.Pool == nil) != (other.Pool == nil)) {
		in, other := &in.Pool, &other.Pool
		if other == nil || !in.DeepEqual(other) {
			return false
		}
	}

	if ((in.PodCIDRs != nil) && (other.PodCIDRs != nil)) || ((in.PodCIDRs == nil) != (other.PodCIDRs == nil)) {
		in, other := &in.PodCIDRs, &other.PodCIDRs
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	if in.MinAllocate != other.MinAllocate {
		return false
	}
	if in.MaxAllocate != other.MaxAllocate {
		return false
	}
	if in.PreAllocate != other.PreAllocate {
		return false
	}
	if in.MaxAboveWatermark != other.MaxAboveWatermark {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *IPAMStatus) DeepEqual(other *IPAMStatus) bool {
	if other == nil {
		return false
	}

	if ((in.Used != nil) && (other.Used != nil)) || ((in.Used == nil) != (other.Used == nil)) {
		in, other := &in.Used, &other.Used
		if other == nil || !in.DeepEqual(other) {
			return false
		}
	}

	if in.OperatorStatus != other.OperatorStatus {
		return false
	}

	if ((in.ReleaseIPs != nil) && (other.ReleaseIPs != nil)) || ((in.ReleaseIPs == nil) != (other.ReleaseIPs == nil)) {
		in, other := &in.ReleaseIPs, &other.ReleaseIPs
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if inValue != otherValue {
						return false
					}
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *Limits) DeepEqual(other *Limits) bool {
	if other == nil {
		return false
	}

	if in.Adapters != other.Adapters {
		return false
	}
	if in.IPv4 != other.IPv4 {
		return false
	}
	if in.IPv6 != other.IPv6 {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *OperatorStatus) DeepEqual(other *OperatorStatus) bool {
	if other == nil {
		return false
	}

	if in.Error != other.Error {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *PoolQuota) DeepEqual(other *PoolQuota) bool {
	if other == nil {
		return false
	}

	if in.AvailabilityZone != other.AvailabilityZone {
		return false
	}
	if in.AvailableIPs != other.AvailableIPs {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *PoolQuotaMap) DeepEqual(other *PoolQuotaMap) bool {
	if other == nil {
		return false
	}

	if len(*in) != len(*other) {
		return false
	} else {
		for key, inValue := range *in {
			if otherValue, present := (*other)[key]; !present {
				return false
			} else {
				if !inValue.DeepEqual(&otherValue) {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *Subnet) DeepEqual(other *Subnet) bool {
	if other == nil {
		return false
	}

	if in.ID != other.ID {
		return false
	}
	if in.Name != other.Name {
		return false
	}
	if (in.CIDR == nil) != (other.CIDR == nil) {
		return false
	} else if in.CIDR != nil {
		if !in.CIDR.DeepEqual(other.CIDR) {
			return false
		}
	}

	if in.AvailabilityZone != other.AvailabilityZone {
		return false
	}
	if in.VirtualNetworkID != other.VirtualNetworkID {
		return false
	}
	if in.AvailableAddresses != other.AvailableAddresses {
		return false
	}
	if ((in.Tags != nil) && (other.Tags != nil)) || ((in.Tags == nil) != (other.Tags == nil)) {
		in, other := &in.Tags, &other.Tags
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if inValue != otherValue {
						return false
					}
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *SubnetMap) DeepEqual(other *SubnetMap) bool {
	if other == nil {
		return false
	}

	if len(*in) != len(*other) {
		return false
	} else {
		for key, inValue := range *in {
			if otherValue, present := (*other)[key]; !present {
				return false
			} else {
				if !inValue.DeepEqual(otherValue) {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *Tags) DeepEqual(other *Tags) bool {
	if other == nil {
		return false
	}

	if len(*in) != len(*other) {
		return false
	} else {
		for key, inValue := range *in {
			if otherValue, present := (*other)[key]; !present {
				return false
			} else {
				if inValue != otherValue {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *VirtualNetwork) DeepEqual(other *VirtualNetwork) bool {
	if other == nil {
		return false
	}

	if in.ID != other.ID {
		return false
	}
	if in.PrimaryCIDR != other.PrimaryCIDR {
		return false
	}
	if ((in.CIDRs != nil) && (other.CIDRs != nil)) || ((in.CIDRs == nil) != (other.CIDRs == nil)) {
		in, other := &in.CIDRs, &other.CIDRs
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *VirtualNetworkMap) DeepEqual(other *VirtualNetworkMap) bool {
	if other == nil {
		return false
	}

	if len(*in) != len(*other) {
		return false
	} else {
		for key, inValue := range *in {
			if otherValue, present := (*other)[key]; !present {
				return false
			} else {
				if !inValue.DeepEqual(otherValue) {
					return false
				}
			}
		}
	}

	return true
}
