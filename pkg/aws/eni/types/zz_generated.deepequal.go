//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: Apache-2.0
// Copyright 2017-2021 Authors of Cilium

// Code generated by main. DO NOT EDIT.

package types

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *AwsSubnet) DeepEqual(other *AwsSubnet) bool {
	if other == nil {
		return false
	}

	if in.ID != other.ID {
		return false
	}
	if in.CIDR != other.CIDR {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *AwsVPC) DeepEqual(other *AwsVPC) bool {
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
func (in *ENI) DeepEqual(other *ENI) bool {
	if other == nil {
		return false
	}

	if in.ID != other.ID {
		return false
	}
	if in.IP != other.IP {
		return false
	}
	if in.MAC != other.MAC {
		return false
	}
	if in.AvailabilityZone != other.AvailabilityZone {
		return false
	}
	if in.Description != other.Description {
		return false
	}
	if in.Number != other.Number {
		return false
	}
	if in.Subnet != other.Subnet {
		return false
	}

	if !in.VPC.DeepEqual(&other.VPC) {
		return false
	}

	if ((in.Addresses != nil) && (other.Addresses != nil)) || ((in.Addresses == nil) != (other.Addresses == nil)) {
		in, other := &in.Addresses, &other.Addresses
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

	if ((in.Prefixes != nil) && (other.Prefixes != nil)) || ((in.Prefixes == nil) != (other.Prefixes == nil)) {
		in, other := &in.Prefixes, &other.Prefixes
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

	if ((in.SecurityGroups != nil) && (other.SecurityGroups != nil)) || ((in.SecurityGroups == nil) != (other.SecurityGroups == nil)) {
		in, other := &in.SecurityGroups, &other.SecurityGroups
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
func (in *ENISpec) DeepEqual(other *ENISpec) bool {
	if other == nil {
		return false
	}

	if in.InstanceID != other.InstanceID {
		return false
	}
	if in.InstanceType != other.InstanceType {
		return false
	}
	if in.MinAllocate != other.MinAllocate {
		return false
	}
	if in.PreAllocate != other.PreAllocate {
		return false
	}
	if in.MaxAboveWatermark != other.MaxAboveWatermark {
		return false
	}
	if (in.FirstInterfaceIndex == nil) != (other.FirstInterfaceIndex == nil) {
		return false
	} else if in.FirstInterfaceIndex != nil {
		if *in.FirstInterfaceIndex != *other.FirstInterfaceIndex {
			return false
		}
	}

	if ((in.SecurityGroups != nil) && (other.SecurityGroups != nil)) || ((in.SecurityGroups == nil) != (other.SecurityGroups == nil)) {
		in, other := &in.SecurityGroups, &other.SecurityGroups
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

	if ((in.SecurityGroupTags != nil) && (other.SecurityGroupTags != nil)) || ((in.SecurityGroupTags == nil) != (other.SecurityGroupTags == nil)) {
		in, other := &in.SecurityGroupTags, &other.SecurityGroupTags
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

	if ((in.SubnetIDs != nil) && (other.SubnetIDs != nil)) || ((in.SubnetIDs == nil) != (other.SubnetIDs == nil)) {
		in, other := &in.SubnetIDs, &other.SubnetIDs
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

	if ((in.SubnetTags != nil) && (other.SubnetTags != nil)) || ((in.SubnetTags == nil) != (other.SubnetTags == nil)) {
		in, other := &in.SubnetTags, &other.SubnetTags
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

	if in.VpcID != other.VpcID {
		return false
	}
	if in.AvailabilityZone != other.AvailabilityZone {
		return false
	}
	if (in.DeleteOnTermination == nil) != (other.DeleteOnTermination == nil) {
		return false
	} else if in.DeleteOnTermination != nil {
		if *in.DeleteOnTermination != *other.DeleteOnTermination {
			return false
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *ENIStatus) DeepEqual(other *ENIStatus) bool {
	if other == nil {
		return false
	}

	if ((in.ENIs != nil) && (other.ENIs != nil)) || ((in.ENIs == nil) != (other.ENIs == nil)) {
		in, other := &in.ENIs, &other.ENIs
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
	}

	return true
}
