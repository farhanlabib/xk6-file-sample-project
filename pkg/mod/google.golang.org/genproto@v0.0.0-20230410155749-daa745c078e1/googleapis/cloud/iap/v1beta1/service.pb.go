// Copyright 2021 Google LLC
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/cloud/iap/v1beta1/service.proto

package iap

import (
	context "context"
	reflect "reflect"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	v1 "google.golang.org/genproto/googleapis/iam/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_google_cloud_iap_v1beta1_service_proto protoreflect.FileDescriptor

var file_google_cloud_iap_v1beta1_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x69,
	0x61, 0x70, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x80, 0x04, 0x0a, 0x1e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x41, 0x77, 0x61, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31, 0x12, 0x79, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x49,
	0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x49, 0x61, 0x6d, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x22, 0x23, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x7b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3d,
	0x2a, 0x2a, 0x7d, 0x3a, 0x73, 0x65, 0x74, 0x49, 0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x3a, 0x01, 0x2a, 0x12, 0x79, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x49, 0x61, 0x6d, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x12, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x2e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x22, 0x23, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x7b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3d, 0x2a, 0x2a, 0x7d, 0x3a, 0x67,
	0x65, 0x74, 0x49, 0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x9f,
	0x01, 0x0a, 0x12, 0x54, 0x65, 0x73, 0x74, 0x49, 0x61, 0x6d, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x49, 0x61, 0x6d, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x49, 0x61, 0x6d, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x2e, 0x22, 0x29, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x7b, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3d, 0x2a, 0x2a, 0x7d, 0x3a, 0x74, 0x65, 0x73, 0x74, 0x49,
	0x61, 0x6d, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x01, 0x2a,
	0x1a, 0x46, 0xca, 0x41, 0x12, 0x69, 0x61, 0x70, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x5d, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x61, 0x70,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x69, 0x61, 0x70, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x3b, 0x69, 0x61, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_google_cloud_iap_v1beta1_service_proto_goTypes = []interface{}{
	(*v1.SetIamPolicyRequest)(nil),        // 0: google.iam.v1.SetIamPolicyRequest
	(*v1.GetIamPolicyRequest)(nil),        // 1: google.iam.v1.GetIamPolicyRequest
	(*v1.TestIamPermissionsRequest)(nil),  // 2: google.iam.v1.TestIamPermissionsRequest
	(*v1.Policy)(nil),                     // 3: google.iam.v1.Policy
	(*v1.TestIamPermissionsResponse)(nil), // 4: google.iam.v1.TestIamPermissionsResponse
}
var file_google_cloud_iap_v1beta1_service_proto_depIdxs = []int32{
	0, // 0: google.cloud.iap.v1beta1.IdentityAwareProxyAdminV1Beta1.SetIamPolicy:input_type -> google.iam.v1.SetIamPolicyRequest
	1, // 1: google.cloud.iap.v1beta1.IdentityAwareProxyAdminV1Beta1.GetIamPolicy:input_type -> google.iam.v1.GetIamPolicyRequest
	2, // 2: google.cloud.iap.v1beta1.IdentityAwareProxyAdminV1Beta1.TestIamPermissions:input_type -> google.iam.v1.TestIamPermissionsRequest
	3, // 3: google.cloud.iap.v1beta1.IdentityAwareProxyAdminV1Beta1.SetIamPolicy:output_type -> google.iam.v1.Policy
	3, // 4: google.cloud.iap.v1beta1.IdentityAwareProxyAdminV1Beta1.GetIamPolicy:output_type -> google.iam.v1.Policy
	4, // 5: google.cloud.iap.v1beta1.IdentityAwareProxyAdminV1Beta1.TestIamPermissions:output_type -> google.iam.v1.TestIamPermissionsResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_iap_v1beta1_service_proto_init() }
func file_google_cloud_iap_v1beta1_service_proto_init() {
	if File_google_cloud_iap_v1beta1_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_iap_v1beta1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_iap_v1beta1_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_iap_v1beta1_service_proto_depIdxs,
	}.Build()
	File_google_cloud_iap_v1beta1_service_proto = out.File
	file_google_cloud_iap_v1beta1_service_proto_rawDesc = nil
	file_google_cloud_iap_v1beta1_service_proto_goTypes = nil
	file_google_cloud_iap_v1beta1_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// IdentityAwareProxyAdminV1Beta1Client is the client API for IdentityAwareProxyAdminV1Beta1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IdentityAwareProxyAdminV1Beta1Client interface {
	// Sets the access control policy for an Identity-Aware Proxy protected
	// resource. Replaces any existing policy.
	// More information about managing access via IAP can be found at:
	// https://cloud.google.com/iap/docs/managing-access#managing_access_via_the_api
	SetIamPolicy(ctx context.Context, in *v1.SetIamPolicyRequest, opts ...grpc.CallOption) (*v1.Policy, error)
	// Gets the access control policy for an Identity-Aware Proxy protected
	// resource.
	// More information about managing access via IAP can be found at:
	// https://cloud.google.com/iap/docs/managing-access#managing_access_via_the_api
	GetIamPolicy(ctx context.Context, in *v1.GetIamPolicyRequest, opts ...grpc.CallOption) (*v1.Policy, error)
	// Returns permissions that a caller has on the Identity-Aware Proxy protected
	// resource. If the resource does not exist or the caller does not have
	// Identity-Aware Proxy permissions a [google.rpc.Code.PERMISSION_DENIED]
	// will be returned.
	// More information about managing access via IAP can be found at:
	// https://cloud.google.com/iap/docs/managing-access#managing_access_via_the_api
	TestIamPermissions(ctx context.Context, in *v1.TestIamPermissionsRequest, opts ...grpc.CallOption) (*v1.TestIamPermissionsResponse, error)
}

type identityAwareProxyAdminV1Beta1Client struct {
	cc grpc.ClientConnInterface
}

func NewIdentityAwareProxyAdminV1Beta1Client(cc grpc.ClientConnInterface) IdentityAwareProxyAdminV1Beta1Client {
	return &identityAwareProxyAdminV1Beta1Client{cc}
}

func (c *identityAwareProxyAdminV1Beta1Client) SetIamPolicy(ctx context.Context, in *v1.SetIamPolicyRequest, opts ...grpc.CallOption) (*v1.Policy, error) {
	out := new(v1.Policy)
	err := c.cc.Invoke(ctx, "/google.cloud.iap.v1beta1.IdentityAwareProxyAdminV1Beta1/SetIamPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityAwareProxyAdminV1Beta1Client) GetIamPolicy(ctx context.Context, in *v1.GetIamPolicyRequest, opts ...grpc.CallOption) (*v1.Policy, error) {
	out := new(v1.Policy)
	err := c.cc.Invoke(ctx, "/google.cloud.iap.v1beta1.IdentityAwareProxyAdminV1Beta1/GetIamPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityAwareProxyAdminV1Beta1Client) TestIamPermissions(ctx context.Context, in *v1.TestIamPermissionsRequest, opts ...grpc.CallOption) (*v1.TestIamPermissionsResponse, error) {
	out := new(v1.TestIamPermissionsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.iap.v1beta1.IdentityAwareProxyAdminV1Beta1/TestIamPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityAwareProxyAdminV1Beta1Server is the server API for IdentityAwareProxyAdminV1Beta1 service.
type IdentityAwareProxyAdminV1Beta1Server interface {
	// Sets the access control policy for an Identity-Aware Proxy protected
	// resource. Replaces any existing policy.
	// More information about managing access via IAP can be found at:
	// https://cloud.google.com/iap/docs/managing-access#managing_access_via_the_api
	SetIamPolicy(context.Context, *v1.SetIamPolicyRequest) (*v1.Policy, error)
	// Gets the access control policy for an Identity-Aware Proxy protected
	// resource.
	// More information about managing access via IAP can be found at:
	// https://cloud.google.com/iap/docs/managing-access#managing_access_via_the_api
	GetIamPolicy(context.Context, *v1.GetIamPolicyRequest) (*v1.Policy, error)
	// Returns permissions that a caller has on the Identity-Aware Proxy protected
	// resource. If the resource does not exist or the caller does not have
	// Identity-Aware Proxy permissions a [google.rpc.Code.PERMISSION_DENIED]
	// will be returned.
	// More information about managing access via IAP can be found at:
	// https://cloud.google.com/iap/docs/managing-access#managing_access_via_the_api
	TestIamPermissions(context.Context, *v1.TestIamPermissionsRequest) (*v1.TestIamPermissionsResponse, error)
}

// UnimplementedIdentityAwareProxyAdminV1Beta1Server can be embedded to have forward compatible implementations.
type UnimplementedIdentityAwareProxyAdminV1Beta1Server struct {
}

func (*UnimplementedIdentityAwareProxyAdminV1Beta1Server) SetIamPolicy(context.Context, *v1.SetIamPolicyRequest) (*v1.Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetIamPolicy not implemented")
}
func (*UnimplementedIdentityAwareProxyAdminV1Beta1Server) GetIamPolicy(context.Context, *v1.GetIamPolicyRequest) (*v1.Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIamPolicy not implemented")
}
func (*UnimplementedIdentityAwareProxyAdminV1Beta1Server) TestIamPermissions(context.Context, *v1.TestIamPermissionsRequest) (*v1.TestIamPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestIamPermissions not implemented")
}

func RegisterIdentityAwareProxyAdminV1Beta1Server(s *grpc.Server, srv IdentityAwareProxyAdminV1Beta1Server) {
	s.RegisterService(&_IdentityAwareProxyAdminV1Beta1_serviceDesc, srv)
}

func _IdentityAwareProxyAdminV1Beta1_SetIamPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.SetIamPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityAwareProxyAdminV1Beta1Server).SetIamPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.iap.v1beta1.IdentityAwareProxyAdminV1Beta1/SetIamPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityAwareProxyAdminV1Beta1Server).SetIamPolicy(ctx, req.(*v1.SetIamPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityAwareProxyAdminV1Beta1_GetIamPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetIamPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityAwareProxyAdminV1Beta1Server).GetIamPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.iap.v1beta1.IdentityAwareProxyAdminV1Beta1/GetIamPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityAwareProxyAdminV1Beta1Server).GetIamPolicy(ctx, req.(*v1.GetIamPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityAwareProxyAdminV1Beta1_TestIamPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.TestIamPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityAwareProxyAdminV1Beta1Server).TestIamPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.iap.v1beta1.IdentityAwareProxyAdminV1Beta1/TestIamPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityAwareProxyAdminV1Beta1Server).TestIamPermissions(ctx, req.(*v1.TestIamPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IdentityAwareProxyAdminV1Beta1_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.iap.v1beta1.IdentityAwareProxyAdminV1Beta1",
	HandlerType: (*IdentityAwareProxyAdminV1Beta1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetIamPolicy",
			Handler:    _IdentityAwareProxyAdminV1Beta1_SetIamPolicy_Handler,
		},
		{
			MethodName: "GetIamPolicy",
			Handler:    _IdentityAwareProxyAdminV1Beta1_GetIamPolicy_Handler,
		},
		{
			MethodName: "TestIamPermissions",
			Handler:    _IdentityAwareProxyAdminV1Beta1_TestIamPermissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/iap/v1beta1/service.proto",
}