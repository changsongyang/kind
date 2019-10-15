/*
Copyright 2019 The Kubernetes Authors.

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

package config

import (
	v1alpha3 "sigs.k8s.io/kind/pkg/apis/config/v1alpha3"
)

func Convertv1alpha3(in *v1alpha3.Cluster) *Cluster {
	out := &Cluster{
		Nodes:                        make([]Node, len(in.Nodes)),
		KubeadmConfigPatches:         in.KubeadmConfigPatches,
		KubeadmConfigPatchesJSON6902: make([]PatchJSON6902, len(in.KubeadmConfigPatchesJSON6902)),
	}

	for i := range in.Nodes {
		convertv1alpha3Node(&in.Nodes[i], &out.Nodes[i])
	}

	convertv1alpha3Networking(&in.Networking, &out.Networking)

	for i := range in.KubeadmConfigPatchesJSON6902 {
		convertv1alphaPatchJSON6902(&in.KubeadmConfigPatchesJSON6902[i], &out.KubeadmConfigPatchesJSON6902[i])
	}

	return out
}

func convertv1alpha3Node(in *v1alpha3.Node, out *Node) {
	out.Role = NodeRole(in.Role)
	out.Image = in.Image
	out.ExtraMounts = in.ExtraMounts
	out.ExtraPortMappings = in.ExtraPortMappings
}

func convertv1alphaPatchJSON6902(in *v1alpha3.PatchJSON6902, out *PatchJSON6902) {
	out.Group = in.Group
	out.Version = in.Version
	out.Kind = in.Kind
	out.Name = in.Name
	out.Namespace = in.Namespace
	out.Patch = in.Patch
}

func convertv1alpha3Networking(in *v1alpha3.Networking, out *Networking) {
	out.IPFamily = ClusterIPFamily(in.IPFamily)
	out.APIServerPort = in.APIServerPort
	out.APIServerAddress = in.APIServerAddress
	out.PodSubnet = in.PodSubnet
	out.ServiceSubnet = in.ServiceSubnet
	out.DisableDefaultCNI = in.DisableDefaultCNI
}
