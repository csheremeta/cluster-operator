// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	clusteroperator "github.com/openshift/cluster-operator/pkg/apis/clusteroperator"
	v1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_AWSClusterSpec_To_clusteroperator_AWSClusterSpec,
		Convert_clusteroperator_AWSClusterSpec_To_v1alpha1_AWSClusterSpec,
		Convert_v1alpha1_Cluster_To_clusteroperator_Cluster,
		Convert_clusteroperator_Cluster_To_v1alpha1_Cluster,
		Convert_v1alpha1_ClusterCondition_To_clusteroperator_ClusterCondition,
		Convert_clusteroperator_ClusterCondition_To_v1alpha1_ClusterCondition,
		Convert_v1alpha1_ClusterConfigSpec_To_clusteroperator_ClusterConfigSpec,
		Convert_clusteroperator_ClusterConfigSpec_To_v1alpha1_ClusterConfigSpec,
		Convert_v1alpha1_ClusterHardwareSpec_To_clusteroperator_ClusterHardwareSpec,
		Convert_clusteroperator_ClusterHardwareSpec_To_v1alpha1_ClusterHardwareSpec,
		Convert_v1alpha1_ClusterList_To_clusteroperator_ClusterList,
		Convert_clusteroperator_ClusterList_To_v1alpha1_ClusterList,
		Convert_v1alpha1_ClusterMachineSet_To_clusteroperator_ClusterMachineSet,
		Convert_clusteroperator_ClusterMachineSet_To_v1alpha1_ClusterMachineSet,
		Convert_v1alpha1_ClusterSpec_To_clusteroperator_ClusterSpec,
		Convert_clusteroperator_ClusterSpec_To_v1alpha1_ClusterSpec,
		Convert_v1alpha1_ClusterStatus_To_clusteroperator_ClusterStatus,
		Convert_clusteroperator_ClusterStatus_To_v1alpha1_ClusterStatus,
		Convert_v1alpha1_Machine_To_clusteroperator_Machine,
		Convert_clusteroperator_Machine_To_v1alpha1_Machine,
		Convert_v1alpha1_MachineList_To_clusteroperator_MachineList,
		Convert_clusteroperator_MachineList_To_v1alpha1_MachineList,
		Convert_v1alpha1_MachineSet_To_clusteroperator_MachineSet,
		Convert_clusteroperator_MachineSet_To_v1alpha1_MachineSet,
		Convert_v1alpha1_MachineSetAWSHardwareSpec_To_clusteroperator_MachineSetAWSHardwareSpec,
		Convert_clusteroperator_MachineSetAWSHardwareSpec_To_v1alpha1_MachineSetAWSHardwareSpec,
		Convert_v1alpha1_MachineSetCondition_To_clusteroperator_MachineSetCondition,
		Convert_clusteroperator_MachineSetCondition_To_v1alpha1_MachineSetCondition,
		Convert_v1alpha1_MachineSetConfig_To_clusteroperator_MachineSetConfig,
		Convert_clusteroperator_MachineSetConfig_To_v1alpha1_MachineSetConfig,
		Convert_v1alpha1_MachineSetHardwareSpec_To_clusteroperator_MachineSetHardwareSpec,
		Convert_clusteroperator_MachineSetHardwareSpec_To_v1alpha1_MachineSetHardwareSpec,
		Convert_v1alpha1_MachineSetList_To_clusteroperator_MachineSetList,
		Convert_clusteroperator_MachineSetList_To_v1alpha1_MachineSetList,
		Convert_v1alpha1_MachineSetSpec_To_clusteroperator_MachineSetSpec,
		Convert_clusteroperator_MachineSetSpec_To_v1alpha1_MachineSetSpec,
		Convert_v1alpha1_MachineSetStatus_To_clusteroperator_MachineSetStatus,
		Convert_clusteroperator_MachineSetStatus_To_v1alpha1_MachineSetStatus,
		Convert_v1alpha1_MachineSpec_To_clusteroperator_MachineSpec,
		Convert_clusteroperator_MachineSpec_To_v1alpha1_MachineSpec,
		Convert_v1alpha1_MachineStatus_To_clusteroperator_MachineStatus,
		Convert_clusteroperator_MachineStatus_To_v1alpha1_MachineStatus,
	)
}

func autoConvert_v1alpha1_AWSClusterSpec_To_clusteroperator_AWSClusterSpec(in *AWSClusterSpec, out *clusteroperator.AWSClusterSpec, s conversion.Scope) error {
	out.AccountSecret = in.AccountSecret
	out.Region = in.Region
	out.VPCName = in.VPCName
	out.VPCSubnet = in.VPCSubnet
	return nil
}

// Convert_v1alpha1_AWSClusterSpec_To_clusteroperator_AWSClusterSpec is an autogenerated conversion function.
func Convert_v1alpha1_AWSClusterSpec_To_clusteroperator_AWSClusterSpec(in *AWSClusterSpec, out *clusteroperator.AWSClusterSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_AWSClusterSpec_To_clusteroperator_AWSClusterSpec(in, out, s)
}

func autoConvert_clusteroperator_AWSClusterSpec_To_v1alpha1_AWSClusterSpec(in *clusteroperator.AWSClusterSpec, out *AWSClusterSpec, s conversion.Scope) error {
	out.AccountSecret = in.AccountSecret
	out.Region = in.Region
	out.VPCName = in.VPCName
	out.VPCSubnet = in.VPCSubnet
	return nil
}

// Convert_clusteroperator_AWSClusterSpec_To_v1alpha1_AWSClusterSpec is an autogenerated conversion function.
func Convert_clusteroperator_AWSClusterSpec_To_v1alpha1_AWSClusterSpec(in *clusteroperator.AWSClusterSpec, out *AWSClusterSpec, s conversion.Scope) error {
	return autoConvert_clusteroperator_AWSClusterSpec_To_v1alpha1_AWSClusterSpec(in, out, s)
}

func autoConvert_v1alpha1_Cluster_To_clusteroperator_Cluster(in *Cluster, out *clusteroperator.Cluster, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_ClusterSpec_To_clusteroperator_ClusterSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ClusterStatus_To_clusteroperator_ClusterStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Cluster_To_clusteroperator_Cluster is an autogenerated conversion function.
func Convert_v1alpha1_Cluster_To_clusteroperator_Cluster(in *Cluster, out *clusteroperator.Cluster, s conversion.Scope) error {
	return autoConvert_v1alpha1_Cluster_To_clusteroperator_Cluster(in, out, s)
}

func autoConvert_clusteroperator_Cluster_To_v1alpha1_Cluster(in *clusteroperator.Cluster, out *Cluster, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_clusteroperator_ClusterSpec_To_v1alpha1_ClusterSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_clusteroperator_ClusterStatus_To_v1alpha1_ClusterStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_clusteroperator_Cluster_To_v1alpha1_Cluster is an autogenerated conversion function.
func Convert_clusteroperator_Cluster_To_v1alpha1_Cluster(in *clusteroperator.Cluster, out *Cluster, s conversion.Scope) error {
	return autoConvert_clusteroperator_Cluster_To_v1alpha1_Cluster(in, out, s)
}

func autoConvert_v1alpha1_ClusterCondition_To_clusteroperator_ClusterCondition(in *ClusterCondition, out *clusteroperator.ClusterCondition, s conversion.Scope) error {
	out.Type = clusteroperator.ClusterConditionType(in.Type)
	out.Status = v1.ConditionStatus(in.Status)
	out.LastProbeTime = in.LastProbeTime
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_v1alpha1_ClusterCondition_To_clusteroperator_ClusterCondition is an autogenerated conversion function.
func Convert_v1alpha1_ClusterCondition_To_clusteroperator_ClusterCondition(in *ClusterCondition, out *clusteroperator.ClusterCondition, s conversion.Scope) error {
	return autoConvert_v1alpha1_ClusterCondition_To_clusteroperator_ClusterCondition(in, out, s)
}

func autoConvert_clusteroperator_ClusterCondition_To_v1alpha1_ClusterCondition(in *clusteroperator.ClusterCondition, out *ClusterCondition, s conversion.Scope) error {
	out.Type = ClusterConditionType(in.Type)
	out.Status = v1.ConditionStatus(in.Status)
	out.LastProbeTime = in.LastProbeTime
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_clusteroperator_ClusterCondition_To_v1alpha1_ClusterCondition is an autogenerated conversion function.
func Convert_clusteroperator_ClusterCondition_To_v1alpha1_ClusterCondition(in *clusteroperator.ClusterCondition, out *ClusterCondition, s conversion.Scope) error {
	return autoConvert_clusteroperator_ClusterCondition_To_v1alpha1_ClusterCondition(in, out, s)
}

func autoConvert_v1alpha1_ClusterConfigSpec_To_clusteroperator_ClusterConfigSpec(in *ClusterConfigSpec, out *clusteroperator.ClusterConfigSpec, s conversion.Scope) error {
	out.DeploymentType = clusteroperator.ClusterDeploymentType(in.DeploymentType)
	out.OpenshiftVersion = in.OpenshiftVersion
	out.SDNPluginName = in.SDNPluginName
	out.ServiceNetworkSubnet = in.ServiceNetworkSubnet
	out.PodNetworkSubnet = in.PodNetworkSubnet
	return nil
}

// Convert_v1alpha1_ClusterConfigSpec_To_clusteroperator_ClusterConfigSpec is an autogenerated conversion function.
func Convert_v1alpha1_ClusterConfigSpec_To_clusteroperator_ClusterConfigSpec(in *ClusterConfigSpec, out *clusteroperator.ClusterConfigSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_ClusterConfigSpec_To_clusteroperator_ClusterConfigSpec(in, out, s)
}

func autoConvert_clusteroperator_ClusterConfigSpec_To_v1alpha1_ClusterConfigSpec(in *clusteroperator.ClusterConfigSpec, out *ClusterConfigSpec, s conversion.Scope) error {
	out.DeploymentType = ClusterDeploymentType(in.DeploymentType)
	out.OpenshiftVersion = in.OpenshiftVersion
	out.SDNPluginName = in.SDNPluginName
	out.ServiceNetworkSubnet = in.ServiceNetworkSubnet
	out.PodNetworkSubnet = in.PodNetworkSubnet
	return nil
}

// Convert_clusteroperator_ClusterConfigSpec_To_v1alpha1_ClusterConfigSpec is an autogenerated conversion function.
func Convert_clusteroperator_ClusterConfigSpec_To_v1alpha1_ClusterConfigSpec(in *clusteroperator.ClusterConfigSpec, out *ClusterConfigSpec, s conversion.Scope) error {
	return autoConvert_clusteroperator_ClusterConfigSpec_To_v1alpha1_ClusterConfigSpec(in, out, s)
}

func autoConvert_v1alpha1_ClusterHardwareSpec_To_clusteroperator_ClusterHardwareSpec(in *ClusterHardwareSpec, out *clusteroperator.ClusterHardwareSpec, s conversion.Scope) error {
	out.AWS = (*clusteroperator.AWSClusterSpec)(unsafe.Pointer(in.AWS))
	return nil
}

// Convert_v1alpha1_ClusterHardwareSpec_To_clusteroperator_ClusterHardwareSpec is an autogenerated conversion function.
func Convert_v1alpha1_ClusterHardwareSpec_To_clusteroperator_ClusterHardwareSpec(in *ClusterHardwareSpec, out *clusteroperator.ClusterHardwareSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_ClusterHardwareSpec_To_clusteroperator_ClusterHardwareSpec(in, out, s)
}

func autoConvert_clusteroperator_ClusterHardwareSpec_To_v1alpha1_ClusterHardwareSpec(in *clusteroperator.ClusterHardwareSpec, out *ClusterHardwareSpec, s conversion.Scope) error {
	out.AWS = (*AWSClusterSpec)(unsafe.Pointer(in.AWS))
	return nil
}

// Convert_clusteroperator_ClusterHardwareSpec_To_v1alpha1_ClusterHardwareSpec is an autogenerated conversion function.
func Convert_clusteroperator_ClusterHardwareSpec_To_v1alpha1_ClusterHardwareSpec(in *clusteroperator.ClusterHardwareSpec, out *ClusterHardwareSpec, s conversion.Scope) error {
	return autoConvert_clusteroperator_ClusterHardwareSpec_To_v1alpha1_ClusterHardwareSpec(in, out, s)
}

func autoConvert_v1alpha1_ClusterList_To_clusteroperator_ClusterList(in *ClusterList, out *clusteroperator.ClusterList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]clusteroperator.Cluster)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_ClusterList_To_clusteroperator_ClusterList is an autogenerated conversion function.
func Convert_v1alpha1_ClusterList_To_clusteroperator_ClusterList(in *ClusterList, out *clusteroperator.ClusterList, s conversion.Scope) error {
	return autoConvert_v1alpha1_ClusterList_To_clusteroperator_ClusterList(in, out, s)
}

func autoConvert_clusteroperator_ClusterList_To_v1alpha1_ClusterList(in *clusteroperator.ClusterList, out *ClusterList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Cluster)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_clusteroperator_ClusterList_To_v1alpha1_ClusterList is an autogenerated conversion function.
func Convert_clusteroperator_ClusterList_To_v1alpha1_ClusterList(in *clusteroperator.ClusterList, out *ClusterList, s conversion.Scope) error {
	return autoConvert_clusteroperator_ClusterList_To_v1alpha1_ClusterList(in, out, s)
}

func autoConvert_v1alpha1_ClusterMachineSet_To_clusteroperator_ClusterMachineSet(in *ClusterMachineSet, out *clusteroperator.ClusterMachineSet, s conversion.Scope) error {
	out.Name = in.Name
	if err := Convert_v1alpha1_MachineSetConfig_To_clusteroperator_MachineSetConfig(&in.MachineSetConfig, &out.MachineSetConfig, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ClusterMachineSet_To_clusteroperator_ClusterMachineSet is an autogenerated conversion function.
func Convert_v1alpha1_ClusterMachineSet_To_clusteroperator_ClusterMachineSet(in *ClusterMachineSet, out *clusteroperator.ClusterMachineSet, s conversion.Scope) error {
	return autoConvert_v1alpha1_ClusterMachineSet_To_clusteroperator_ClusterMachineSet(in, out, s)
}

func autoConvert_clusteroperator_ClusterMachineSet_To_v1alpha1_ClusterMachineSet(in *clusteroperator.ClusterMachineSet, out *ClusterMachineSet, s conversion.Scope) error {
	out.Name = in.Name
	if err := Convert_clusteroperator_MachineSetConfig_To_v1alpha1_MachineSetConfig(&in.MachineSetConfig, &out.MachineSetConfig, s); err != nil {
		return err
	}
	return nil
}

// Convert_clusteroperator_ClusterMachineSet_To_v1alpha1_ClusterMachineSet is an autogenerated conversion function.
func Convert_clusteroperator_ClusterMachineSet_To_v1alpha1_ClusterMachineSet(in *clusteroperator.ClusterMachineSet, out *ClusterMachineSet, s conversion.Scope) error {
	return autoConvert_clusteroperator_ClusterMachineSet_To_v1alpha1_ClusterMachineSet(in, out, s)
}

func autoConvert_v1alpha1_ClusterSpec_To_clusteroperator_ClusterSpec(in *ClusterSpec, out *clusteroperator.ClusterSpec, s conversion.Scope) error {
	if err := Convert_v1alpha1_ClusterHardwareSpec_To_clusteroperator_ClusterHardwareSpec(&in.Hardware, &out.Hardware, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ClusterConfigSpec_To_clusteroperator_ClusterConfigSpec(&in.Config, &out.Config, s); err != nil {
		return err
	}
	out.DefaultHardwareSpec = (*clusteroperator.MachineSetHardwareSpec)(unsafe.Pointer(in.DefaultHardwareSpec))
	out.MachineSets = *(*[]clusteroperator.ClusterMachineSet)(unsafe.Pointer(&in.MachineSets))
	return nil
}

// Convert_v1alpha1_ClusterSpec_To_clusteroperator_ClusterSpec is an autogenerated conversion function.
func Convert_v1alpha1_ClusterSpec_To_clusteroperator_ClusterSpec(in *ClusterSpec, out *clusteroperator.ClusterSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_ClusterSpec_To_clusteroperator_ClusterSpec(in, out, s)
}

func autoConvert_clusteroperator_ClusterSpec_To_v1alpha1_ClusterSpec(in *clusteroperator.ClusterSpec, out *ClusterSpec, s conversion.Scope) error {
	if err := Convert_clusteroperator_ClusterHardwareSpec_To_v1alpha1_ClusterHardwareSpec(&in.Hardware, &out.Hardware, s); err != nil {
		return err
	}
	if err := Convert_clusteroperator_ClusterConfigSpec_To_v1alpha1_ClusterConfigSpec(&in.Config, &out.Config, s); err != nil {
		return err
	}
	out.DefaultHardwareSpec = (*MachineSetHardwareSpec)(unsafe.Pointer(in.DefaultHardwareSpec))
	out.MachineSets = *(*[]ClusterMachineSet)(unsafe.Pointer(&in.MachineSets))
	return nil
}

// Convert_clusteroperator_ClusterSpec_To_v1alpha1_ClusterSpec is an autogenerated conversion function.
func Convert_clusteroperator_ClusterSpec_To_v1alpha1_ClusterSpec(in *clusteroperator.ClusterSpec, out *ClusterSpec, s conversion.Scope) error {
	return autoConvert_clusteroperator_ClusterSpec_To_v1alpha1_ClusterSpec(in, out, s)
}

func autoConvert_v1alpha1_ClusterStatus_To_clusteroperator_ClusterStatus(in *ClusterStatus, out *clusteroperator.ClusterStatus, s conversion.Scope) error {
	out.MachineSetCount = in.MachineSetCount
	out.MasterMachineSetName = in.MasterMachineSetName
	out.InfraMachineSetName = in.InfraMachineSetName
	out.AdminKubeconfig = (*v1.LocalObjectReference)(unsafe.Pointer(in.AdminKubeconfig))
	out.Provisioned = in.Provisioned
	out.Running = in.Running
	out.Conditions = *(*[]clusteroperator.ClusterCondition)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_v1alpha1_ClusterStatus_To_clusteroperator_ClusterStatus is an autogenerated conversion function.
func Convert_v1alpha1_ClusterStatus_To_clusteroperator_ClusterStatus(in *ClusterStatus, out *clusteroperator.ClusterStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_ClusterStatus_To_clusteroperator_ClusterStatus(in, out, s)
}

func autoConvert_clusteroperator_ClusterStatus_To_v1alpha1_ClusterStatus(in *clusteroperator.ClusterStatus, out *ClusterStatus, s conversion.Scope) error {
	out.MachineSetCount = in.MachineSetCount
	out.MasterMachineSetName = in.MasterMachineSetName
	out.InfraMachineSetName = in.InfraMachineSetName
	out.AdminKubeconfig = (*v1.LocalObjectReference)(unsafe.Pointer(in.AdminKubeconfig))
	out.Provisioned = in.Provisioned
	out.Running = in.Running
	out.Conditions = *(*[]ClusterCondition)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_clusteroperator_ClusterStatus_To_v1alpha1_ClusterStatus is an autogenerated conversion function.
func Convert_clusteroperator_ClusterStatus_To_v1alpha1_ClusterStatus(in *clusteroperator.ClusterStatus, out *ClusterStatus, s conversion.Scope) error {
	return autoConvert_clusteroperator_ClusterStatus_To_v1alpha1_ClusterStatus(in, out, s)
}

func autoConvert_v1alpha1_Machine_To_clusteroperator_Machine(in *Machine, out *clusteroperator.Machine, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_MachineSpec_To_clusteroperator_MachineSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_MachineStatus_To_clusteroperator_MachineStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Machine_To_clusteroperator_Machine is an autogenerated conversion function.
func Convert_v1alpha1_Machine_To_clusteroperator_Machine(in *Machine, out *clusteroperator.Machine, s conversion.Scope) error {
	return autoConvert_v1alpha1_Machine_To_clusteroperator_Machine(in, out, s)
}

func autoConvert_clusteroperator_Machine_To_v1alpha1_Machine(in *clusteroperator.Machine, out *Machine, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_clusteroperator_MachineSpec_To_v1alpha1_MachineSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_clusteroperator_MachineStatus_To_v1alpha1_MachineStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_clusteroperator_Machine_To_v1alpha1_Machine is an autogenerated conversion function.
func Convert_clusteroperator_Machine_To_v1alpha1_Machine(in *clusteroperator.Machine, out *Machine, s conversion.Scope) error {
	return autoConvert_clusteroperator_Machine_To_v1alpha1_Machine(in, out, s)
}

func autoConvert_v1alpha1_MachineList_To_clusteroperator_MachineList(in *MachineList, out *clusteroperator.MachineList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]clusteroperator.Machine)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_MachineList_To_clusteroperator_MachineList is an autogenerated conversion function.
func Convert_v1alpha1_MachineList_To_clusteroperator_MachineList(in *MachineList, out *clusteroperator.MachineList, s conversion.Scope) error {
	return autoConvert_v1alpha1_MachineList_To_clusteroperator_MachineList(in, out, s)
}

func autoConvert_clusteroperator_MachineList_To_v1alpha1_MachineList(in *clusteroperator.MachineList, out *MachineList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Machine)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_clusteroperator_MachineList_To_v1alpha1_MachineList is an autogenerated conversion function.
func Convert_clusteroperator_MachineList_To_v1alpha1_MachineList(in *clusteroperator.MachineList, out *MachineList, s conversion.Scope) error {
	return autoConvert_clusteroperator_MachineList_To_v1alpha1_MachineList(in, out, s)
}

func autoConvert_v1alpha1_MachineSet_To_clusteroperator_MachineSet(in *MachineSet, out *clusteroperator.MachineSet, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_MachineSetSpec_To_clusteroperator_MachineSetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_MachineSetStatus_To_clusteroperator_MachineSetStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_MachineSet_To_clusteroperator_MachineSet is an autogenerated conversion function.
func Convert_v1alpha1_MachineSet_To_clusteroperator_MachineSet(in *MachineSet, out *clusteroperator.MachineSet, s conversion.Scope) error {
	return autoConvert_v1alpha1_MachineSet_To_clusteroperator_MachineSet(in, out, s)
}

func autoConvert_clusteroperator_MachineSet_To_v1alpha1_MachineSet(in *clusteroperator.MachineSet, out *MachineSet, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_clusteroperator_MachineSetSpec_To_v1alpha1_MachineSetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_clusteroperator_MachineSetStatus_To_v1alpha1_MachineSetStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_clusteroperator_MachineSet_To_v1alpha1_MachineSet is an autogenerated conversion function.
func Convert_clusteroperator_MachineSet_To_v1alpha1_MachineSet(in *clusteroperator.MachineSet, out *MachineSet, s conversion.Scope) error {
	return autoConvert_clusteroperator_MachineSet_To_v1alpha1_MachineSet(in, out, s)
}

func autoConvert_v1alpha1_MachineSetAWSHardwareSpec_To_clusteroperator_MachineSetAWSHardwareSpec(in *MachineSetAWSHardwareSpec, out *clusteroperator.MachineSetAWSHardwareSpec, s conversion.Scope) error {
	out.InstanceType = in.InstanceType
	out.AMIName = in.AMIName
	return nil
}

// Convert_v1alpha1_MachineSetAWSHardwareSpec_To_clusteroperator_MachineSetAWSHardwareSpec is an autogenerated conversion function.
func Convert_v1alpha1_MachineSetAWSHardwareSpec_To_clusteroperator_MachineSetAWSHardwareSpec(in *MachineSetAWSHardwareSpec, out *clusteroperator.MachineSetAWSHardwareSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_MachineSetAWSHardwareSpec_To_clusteroperator_MachineSetAWSHardwareSpec(in, out, s)
}

func autoConvert_clusteroperator_MachineSetAWSHardwareSpec_To_v1alpha1_MachineSetAWSHardwareSpec(in *clusteroperator.MachineSetAWSHardwareSpec, out *MachineSetAWSHardwareSpec, s conversion.Scope) error {
	out.InstanceType = in.InstanceType
	out.AMIName = in.AMIName
	return nil
}

// Convert_clusteroperator_MachineSetAWSHardwareSpec_To_v1alpha1_MachineSetAWSHardwareSpec is an autogenerated conversion function.
func Convert_clusteroperator_MachineSetAWSHardwareSpec_To_v1alpha1_MachineSetAWSHardwareSpec(in *clusteroperator.MachineSetAWSHardwareSpec, out *MachineSetAWSHardwareSpec, s conversion.Scope) error {
	return autoConvert_clusteroperator_MachineSetAWSHardwareSpec_To_v1alpha1_MachineSetAWSHardwareSpec(in, out, s)
}

func autoConvert_v1alpha1_MachineSetCondition_To_clusteroperator_MachineSetCondition(in *MachineSetCondition, out *clusteroperator.MachineSetCondition, s conversion.Scope) error {
	out.Type = clusteroperator.MachineSetConditionType(in.Type)
	out.Status = v1.ConditionStatus(in.Status)
	out.LastProbeTime = in.LastProbeTime
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_v1alpha1_MachineSetCondition_To_clusteroperator_MachineSetCondition is an autogenerated conversion function.
func Convert_v1alpha1_MachineSetCondition_To_clusteroperator_MachineSetCondition(in *MachineSetCondition, out *clusteroperator.MachineSetCondition, s conversion.Scope) error {
	return autoConvert_v1alpha1_MachineSetCondition_To_clusteroperator_MachineSetCondition(in, out, s)
}

func autoConvert_clusteroperator_MachineSetCondition_To_v1alpha1_MachineSetCondition(in *clusteroperator.MachineSetCondition, out *MachineSetCondition, s conversion.Scope) error {
	out.Type = MachineSetConditionType(in.Type)
	out.Status = v1.ConditionStatus(in.Status)
	out.LastProbeTime = in.LastProbeTime
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_clusteroperator_MachineSetCondition_To_v1alpha1_MachineSetCondition is an autogenerated conversion function.
func Convert_clusteroperator_MachineSetCondition_To_v1alpha1_MachineSetCondition(in *clusteroperator.MachineSetCondition, out *MachineSetCondition, s conversion.Scope) error {
	return autoConvert_clusteroperator_MachineSetCondition_To_v1alpha1_MachineSetCondition(in, out, s)
}

func autoConvert_v1alpha1_MachineSetConfig_To_clusteroperator_MachineSetConfig(in *MachineSetConfig, out *clusteroperator.MachineSetConfig, s conversion.Scope) error {
	out.NodeType = clusteroperator.NodeType(in.NodeType)
	out.Infra = in.Infra
	out.Size = in.Size
	if err := Convert_v1alpha1_MachineSetHardwareSpec_To_clusteroperator_MachineSetHardwareSpec(&in.Hardware, &out.Hardware, s); err != nil {
		return err
	}
	out.NodeLabels = *(*map[string]string)(unsafe.Pointer(&in.NodeLabels))
	return nil
}

// Convert_v1alpha1_MachineSetConfig_To_clusteroperator_MachineSetConfig is an autogenerated conversion function.
func Convert_v1alpha1_MachineSetConfig_To_clusteroperator_MachineSetConfig(in *MachineSetConfig, out *clusteroperator.MachineSetConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_MachineSetConfig_To_clusteroperator_MachineSetConfig(in, out, s)
}

func autoConvert_clusteroperator_MachineSetConfig_To_v1alpha1_MachineSetConfig(in *clusteroperator.MachineSetConfig, out *MachineSetConfig, s conversion.Scope) error {
	out.NodeType = NodeType(in.NodeType)
	out.Infra = in.Infra
	out.Size = in.Size
	if err := Convert_clusteroperator_MachineSetHardwareSpec_To_v1alpha1_MachineSetHardwareSpec(&in.Hardware, &out.Hardware, s); err != nil {
		return err
	}
	out.NodeLabels = *(*map[string]string)(unsafe.Pointer(&in.NodeLabels))
	return nil
}

// Convert_clusteroperator_MachineSetConfig_To_v1alpha1_MachineSetConfig is an autogenerated conversion function.
func Convert_clusteroperator_MachineSetConfig_To_v1alpha1_MachineSetConfig(in *clusteroperator.MachineSetConfig, out *MachineSetConfig, s conversion.Scope) error {
	return autoConvert_clusteroperator_MachineSetConfig_To_v1alpha1_MachineSetConfig(in, out, s)
}

func autoConvert_v1alpha1_MachineSetHardwareSpec_To_clusteroperator_MachineSetHardwareSpec(in *MachineSetHardwareSpec, out *clusteroperator.MachineSetHardwareSpec, s conversion.Scope) error {
	out.AWS = (*clusteroperator.MachineSetAWSHardwareSpec)(unsafe.Pointer(in.AWS))
	return nil
}

// Convert_v1alpha1_MachineSetHardwareSpec_To_clusteroperator_MachineSetHardwareSpec is an autogenerated conversion function.
func Convert_v1alpha1_MachineSetHardwareSpec_To_clusteroperator_MachineSetHardwareSpec(in *MachineSetHardwareSpec, out *clusteroperator.MachineSetHardwareSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_MachineSetHardwareSpec_To_clusteroperator_MachineSetHardwareSpec(in, out, s)
}

func autoConvert_clusteroperator_MachineSetHardwareSpec_To_v1alpha1_MachineSetHardwareSpec(in *clusteroperator.MachineSetHardwareSpec, out *MachineSetHardwareSpec, s conversion.Scope) error {
	out.AWS = (*MachineSetAWSHardwareSpec)(unsafe.Pointer(in.AWS))
	return nil
}

// Convert_clusteroperator_MachineSetHardwareSpec_To_v1alpha1_MachineSetHardwareSpec is an autogenerated conversion function.
func Convert_clusteroperator_MachineSetHardwareSpec_To_v1alpha1_MachineSetHardwareSpec(in *clusteroperator.MachineSetHardwareSpec, out *MachineSetHardwareSpec, s conversion.Scope) error {
	return autoConvert_clusteroperator_MachineSetHardwareSpec_To_v1alpha1_MachineSetHardwareSpec(in, out, s)
}

func autoConvert_v1alpha1_MachineSetList_To_clusteroperator_MachineSetList(in *MachineSetList, out *clusteroperator.MachineSetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]clusteroperator.MachineSet)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_MachineSetList_To_clusteroperator_MachineSetList is an autogenerated conversion function.
func Convert_v1alpha1_MachineSetList_To_clusteroperator_MachineSetList(in *MachineSetList, out *clusteroperator.MachineSetList, s conversion.Scope) error {
	return autoConvert_v1alpha1_MachineSetList_To_clusteroperator_MachineSetList(in, out, s)
}

func autoConvert_clusteroperator_MachineSetList_To_v1alpha1_MachineSetList(in *clusteroperator.MachineSetList, out *MachineSetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]MachineSet)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_clusteroperator_MachineSetList_To_v1alpha1_MachineSetList is an autogenerated conversion function.
func Convert_clusteroperator_MachineSetList_To_v1alpha1_MachineSetList(in *clusteroperator.MachineSetList, out *MachineSetList, s conversion.Scope) error {
	return autoConvert_clusteroperator_MachineSetList_To_v1alpha1_MachineSetList(in, out, s)
}

func autoConvert_v1alpha1_MachineSetSpec_To_clusteroperator_MachineSetSpec(in *MachineSetSpec, out *clusteroperator.MachineSetSpec, s conversion.Scope) error {
	if err := Convert_v1alpha1_MachineSetConfig_To_clusteroperator_MachineSetConfig(&in.MachineSetConfig, &out.MachineSetConfig, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_MachineSetSpec_To_clusteroperator_MachineSetSpec is an autogenerated conversion function.
func Convert_v1alpha1_MachineSetSpec_To_clusteroperator_MachineSetSpec(in *MachineSetSpec, out *clusteroperator.MachineSetSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_MachineSetSpec_To_clusteroperator_MachineSetSpec(in, out, s)
}

func autoConvert_clusteroperator_MachineSetSpec_To_v1alpha1_MachineSetSpec(in *clusteroperator.MachineSetSpec, out *MachineSetSpec, s conversion.Scope) error {
	if err := Convert_clusteroperator_MachineSetConfig_To_v1alpha1_MachineSetConfig(&in.MachineSetConfig, &out.MachineSetConfig, s); err != nil {
		return err
	}
	return nil
}

// Convert_clusteroperator_MachineSetSpec_To_v1alpha1_MachineSetSpec is an autogenerated conversion function.
func Convert_clusteroperator_MachineSetSpec_To_v1alpha1_MachineSetSpec(in *clusteroperator.MachineSetSpec, out *MachineSetSpec, s conversion.Scope) error {
	return autoConvert_clusteroperator_MachineSetSpec_To_v1alpha1_MachineSetSpec(in, out, s)
}

func autoConvert_v1alpha1_MachineSetStatus_To_clusteroperator_MachineSetStatus(in *MachineSetStatus, out *clusteroperator.MachineSetStatus, s conversion.Scope) error {
	out.MachinesProvisioned = in.MachinesProvisioned
	out.MachinesReady = in.MachinesReady
	out.Conditions = *(*[]clusteroperator.MachineSetCondition)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_v1alpha1_MachineSetStatus_To_clusteroperator_MachineSetStatus is an autogenerated conversion function.
func Convert_v1alpha1_MachineSetStatus_To_clusteroperator_MachineSetStatus(in *MachineSetStatus, out *clusteroperator.MachineSetStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_MachineSetStatus_To_clusteroperator_MachineSetStatus(in, out, s)
}

func autoConvert_clusteroperator_MachineSetStatus_To_v1alpha1_MachineSetStatus(in *clusteroperator.MachineSetStatus, out *MachineSetStatus, s conversion.Scope) error {
	out.MachinesProvisioned = in.MachinesProvisioned
	out.MachinesReady = in.MachinesReady
	out.Conditions = *(*[]MachineSetCondition)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_clusteroperator_MachineSetStatus_To_v1alpha1_MachineSetStatus is an autogenerated conversion function.
func Convert_clusteroperator_MachineSetStatus_To_v1alpha1_MachineSetStatus(in *clusteroperator.MachineSetStatus, out *MachineSetStatus, s conversion.Scope) error {
	return autoConvert_clusteroperator_MachineSetStatus_To_v1alpha1_MachineSetStatus(in, out, s)
}

func autoConvert_v1alpha1_MachineSpec_To_clusteroperator_MachineSpec(in *MachineSpec, out *clusteroperator.MachineSpec, s conversion.Scope) error {
	out.NodeType = clusteroperator.NodeType(in.NodeType)
	return nil
}

// Convert_v1alpha1_MachineSpec_To_clusteroperator_MachineSpec is an autogenerated conversion function.
func Convert_v1alpha1_MachineSpec_To_clusteroperator_MachineSpec(in *MachineSpec, out *clusteroperator.MachineSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_MachineSpec_To_clusteroperator_MachineSpec(in, out, s)
}

func autoConvert_clusteroperator_MachineSpec_To_v1alpha1_MachineSpec(in *clusteroperator.MachineSpec, out *MachineSpec, s conversion.Scope) error {
	out.NodeType = NodeType(in.NodeType)
	return nil
}

// Convert_clusteroperator_MachineSpec_To_v1alpha1_MachineSpec is an autogenerated conversion function.
func Convert_clusteroperator_MachineSpec_To_v1alpha1_MachineSpec(in *clusteroperator.MachineSpec, out *MachineSpec, s conversion.Scope) error {
	return autoConvert_clusteroperator_MachineSpec_To_v1alpha1_MachineSpec(in, out, s)
}

func autoConvert_v1alpha1_MachineStatus_To_clusteroperator_MachineStatus(in *MachineStatus, out *clusteroperator.MachineStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_MachineStatus_To_clusteroperator_MachineStatus is an autogenerated conversion function.
func Convert_v1alpha1_MachineStatus_To_clusteroperator_MachineStatus(in *MachineStatus, out *clusteroperator.MachineStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_MachineStatus_To_clusteroperator_MachineStatus(in, out, s)
}

func autoConvert_clusteroperator_MachineStatus_To_v1alpha1_MachineStatus(in *clusteroperator.MachineStatus, out *MachineStatus, s conversion.Scope) error {
	return nil
}

// Convert_clusteroperator_MachineStatus_To_v1alpha1_MachineStatus is an autogenerated conversion function.
func Convert_clusteroperator_MachineStatus_To_v1alpha1_MachineStatus(in *clusteroperator.MachineStatus, out *MachineStatus, s conversion.Scope) error {
	return autoConvert_clusteroperator_MachineStatus_To_v1alpha1_MachineStatus(in, out, s)
}
