// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/dam/v1/dam_secrets.proto

// Package dam provides protocol buffer versions of the DAM API, allowing
// end points to receive requests and returns responses using these messages.

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type DamSecrets struct {
	Version              string                          `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Revision             int64                           `protobuf:"varint,2,opt,name=revision,proto3" json:"revision,omitempty"`
	CommitTime           float64                         `protobuf:"fixed64,3,opt,name=commit_time,json=commitTime,proto3" json:"commit_time,omitempty"`
	ClientSecrets        map[string]string               `protobuf:"bytes,4,rep,name=client_secrets,json=clientSecrets,proto3" json:"client_secrets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	PublicTokenKeys      map[string]string               `protobuf:"bytes,5,rep,name=public_token_keys,json=publicTokenKeys,proto3" json:"public_token_keys,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	GatekeeperTokenKeys  *DamSecrets_GatekeeperTokenKeys `protobuf:"bytes,6,opt,name=gatekeeper_token_keys,json=gatekeeperTokenKeys,proto3" json:"gatekeeper_token_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *DamSecrets) Reset()         { *m = DamSecrets{} }
func (m *DamSecrets) String() string { return proto.CompactTextString(m) }
func (*DamSecrets) ProtoMessage()    {}
func (*DamSecrets) Descriptor() ([]byte, []int) {
	return fileDescriptor_e301ca973d5091dc, []int{0}
}

func (m *DamSecrets) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DamSecrets.Unmarshal(m, b)
}
func (m *DamSecrets) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DamSecrets.Marshal(b, m, deterministic)
}
func (m *DamSecrets) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DamSecrets.Merge(m, src)
}
func (m *DamSecrets) XXX_Size() int {
	return xxx_messageInfo_DamSecrets.Size(m)
}
func (m *DamSecrets) XXX_DiscardUnknown() {
	xxx_messageInfo_DamSecrets.DiscardUnknown(m)
}

var xxx_messageInfo_DamSecrets proto.InternalMessageInfo

func (m *DamSecrets) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *DamSecrets) GetRevision() int64 {
	if m != nil {
		return m.Revision
	}
	return 0
}

func (m *DamSecrets) GetCommitTime() float64 {
	if m != nil {
		return m.CommitTime
	}
	return 0
}

func (m *DamSecrets) GetClientSecrets() map[string]string {
	if m != nil {
		return m.ClientSecrets
	}
	return nil
}

func (m *DamSecrets) GetPublicTokenKeys() map[string]string {
	if m != nil {
		return m.PublicTokenKeys
	}
	return nil
}

func (m *DamSecrets) GetGatekeeperTokenKeys() *DamSecrets_GatekeeperTokenKeys {
	if m != nil {
		return m.GatekeeperTokenKeys
	}
	return nil
}

type DamSecrets_GatekeeperTokenKeys struct {
	PrivateKey           string   `protobuf:"bytes,1,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	PublicKey            string   `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DamSecrets_GatekeeperTokenKeys) Reset()         { *m = DamSecrets_GatekeeperTokenKeys{} }
func (m *DamSecrets_GatekeeperTokenKeys) String() string { return proto.CompactTextString(m) }
func (*DamSecrets_GatekeeperTokenKeys) ProtoMessage()    {}
func (*DamSecrets_GatekeeperTokenKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_e301ca973d5091dc, []int{0, 2}
}

func (m *DamSecrets_GatekeeperTokenKeys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DamSecrets_GatekeeperTokenKeys.Unmarshal(m, b)
}
func (m *DamSecrets_GatekeeperTokenKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DamSecrets_GatekeeperTokenKeys.Marshal(b, m, deterministic)
}
func (m *DamSecrets_GatekeeperTokenKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DamSecrets_GatekeeperTokenKeys.Merge(m, src)
}
func (m *DamSecrets_GatekeeperTokenKeys) XXX_Size() int {
	return xxx_messageInfo_DamSecrets_GatekeeperTokenKeys.Size(m)
}
func (m *DamSecrets_GatekeeperTokenKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_DamSecrets_GatekeeperTokenKeys.DiscardUnknown(m)
}

var xxx_messageInfo_DamSecrets_GatekeeperTokenKeys proto.InternalMessageInfo

func (m *DamSecrets_GatekeeperTokenKeys) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

func (m *DamSecrets_GatekeeperTokenKeys) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func init() {
	proto.RegisterType((*DamSecrets)(nil), "dam.v1.DamSecrets")
	proto.RegisterMapType((map[string]string)(nil), "dam.v1.DamSecrets.ClientSecretsEntry")
	proto.RegisterMapType((map[string]string)(nil), "dam.v1.DamSecrets.PublicTokenKeysEntry")
	proto.RegisterType((*DamSecrets_GatekeeperTokenKeys)(nil), "dam.v1.DamSecrets.GatekeeperTokenKeys")
}

func init() { proto.RegisterFile("proto/dam/v1/dam_secrets.proto", fileDescriptor_e301ca973d5091dc) }

var fileDescriptor_e301ca973d5091dc = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4b, 0x6f, 0x9b, 0x40,
	0x14, 0x85, 0x45, 0x48, 0xdc, 0xfa, 0x5a, 0x7d, 0x4d, 0x52, 0x09, 0x59, 0x6a, 0x8b, 0x2a, 0xb5,
	0x65, 0x63, 0x50, 0xd2, 0x4d, 0xd5, 0x55, 0x95, 0xb4, 0xca, 0xc2, 0x5d, 0x58, 0xd8, 0xdd, 0x78,
	0x83, 0xc6, 0xc3, 0x35, 0x1e, 0xc1, 0x30, 0x68, 0x66, 0x40, 0xe2, 0xd7, 0xf5, 0xaf, 0x55, 0x3c,
	0xfc, 0xa8, 0xcc, 0x26, 0x2b, 0x38, 0xe7, 0x1e, 0x7d, 0x3a, 0xf7, 0x02, 0xbc, 0x2f, 0x94, 0x34,
	0x32, 0x88, 0xa9, 0x08, 0xaa, 0xdb, 0xe6, 0x11, 0x69, 0x64, 0x0a, 0x8d, 0xf6, 0xdb, 0x01, 0x19,
	0xc5, 0x54, 0xf8, 0xd5, 0xed, 0xc7, 0xbf, 0x97, 0x00, 0x3f, 0xa9, 0x58, 0x76, 0x43, 0xe2, 0xc0,
	0xb3, 0x0a, 0x95, 0xe6, 0x32, 0x77, 0x2c, 0xd7, 0xf2, 0xc6, 0xe1, 0x5e, 0x92, 0x29, 0x3c, 0x57,
	0x58, 0xf1, 0x76, 0x74, 0xe1, 0x5a, 0x9e, 0x1d, 0x1e, 0x34, 0xf9, 0x00, 0x13, 0x26, 0x85, 0xe0,
	0x26, 0x32, 0x5c, 0xa0, 0x63, 0xbb, 0x96, 0x67, 0x85, 0xd0, 0x59, 0x2b, 0x2e, 0x90, 0xfc, 0x86,
	0x97, 0x2c, 0xe3, 0x98, 0x9b, 0x7d, 0x0b, 0xe7, 0xd2, 0xb5, 0xbd, 0xc9, 0xdd, 0x27, 0xbf, 0xab,
	0xe1, 0x1f, 0x2b, 0xf8, 0x0f, 0x6d, 0xb0, 0x57, 0xbf, 0x72, 0xa3, 0xea, 0xf0, 0x05, 0x3b, 0xf5,
	0xc8, 0x12, 0xde, 0x14, 0xe5, 0x26, 0xe3, 0x2c, 0x32, 0x32, 0xc5, 0x3c, 0x4a, 0xb1, 0xd6, 0xce,
	0x55, 0x0b, 0xfc, 0x32, 0x00, 0x5c, 0xb4, 0xd9, 0x55, 0x13, 0x9d, 0x63, 0xdd, 0x23, 0x5f, 0x15,
	0xff, 0xbb, 0x64, 0x0d, 0x6f, 0x13, 0x6a, 0x30, 0x45, 0x2c, 0x50, 0x9d, 0x82, 0x47, 0xae, 0xe5,
	0x4d, 0xee, 0x3e, 0x0f, 0x80, 0x1f, 0x0f, 0xf9, 0x03, 0x26, 0xbc, 0x4e, 0xce, 0xcd, 0xe9, 0x0f,
	0x20, 0xe7, 0x5b, 0x91, 0xd7, 0x60, 0xa7, 0x58, 0xf7, 0x77, 0x6e, 0x5e, 0xc9, 0x0d, 0x5c, 0x55,
	0x34, 0x2b, 0xb1, 0x3d, 0xf0, 0x38, 0xec, 0xc4, 0xf7, 0x8b, 0x6f, 0xd6, 0xf4, 0x1e, 0x6e, 0x86,
	0xd6, 0x78, 0x12, 0xe3, 0x0f, 0x5c, 0x0f, 0x34, 0x6e, 0x3e, 0x5e, 0xa1, 0x78, 0x45, 0x0d, 0x46,
	0x47, 0x14, 0xf4, 0xd6, 0x1c, 0x6b, 0xf2, 0x0e, 0xa0, 0x3f, 0x77, 0x33, 0xef, 0xb0, 0xe3, 0xce,
	0x99, 0x63, 0x7d, 0x1f, 0xae, 0x17, 0x09, 0x37, 0xbb, 0x72, 0xe3, 0x33, 0x29, 0x82, 0x47, 0x29,
	0x93, 0x0c, 0x1f, 0x32, 0x59, 0xc6, 0x8b, 0x8c, 0x9a, 0xad, 0x54, 0x22, 0xd8, 0x21, 0xcd, 0xcc,
	0x8e, 0x51, 0x85, 0xb3, 0x2d, 0xc6, 0xa8, 0xa8, 0xc1, 0x78, 0x46, 0x19, 0x43, 0xad, 0x67, 0x1a,
	0x55, 0xc5, 0x19, 0xea, 0xe0, 0xf4, 0x7f, 0xdd, 0x8c, 0x5a, 0xf5, 0xf5, 0x5f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x30, 0x5c, 0xff, 0xe0, 0xc6, 0x02, 0x00, 0x00,
}
