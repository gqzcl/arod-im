// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: conf/conf.proto

package conf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Bootstrap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server *Server `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Data   *Data   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Client *Client `protobuf:"bytes,3,opt,name=client,proto3" json:"client,omitempty"`
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap.ProtoReflect.Descriptor instead.
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Bootstrap) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Bootstrap) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Bootstrap) GetClient() *Client {
	if x != nil {
		return x.Client
	}
	return nil
}

type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Http     *Server_HTTP     `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	Grpc     *Server_GRPC     `protobuf:"bytes,2,opt,name=grpc,proto3" json:"grpc,omitempty"`
	Register *Server_Register `protobuf:"bytes,3,opt,name=register,proto3" json:"register,omitempty"`
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{1}
}

func (x *Server) GetHttp() *Server_HTTP {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *Server) GetGrpc() *Server_GRPC {
	if x != nil {
		return x.Grpc
	}
	return nil
}

func (x *Server) GetRegister() *Server_Register {
	if x != nil {
		return x.Register
	}
	return nil
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kafka *Data_Kafka `protobuf:"bytes,1,opt,name=kafka,proto3" json:"kafka,omitempty"`
	Redis *Data_Redis `protobuf:"bytes,2,opt,name=redis,proto3" json:"redis,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{2}
}

func (x *Data) GetKafka() *Data_Kafka {
	if x != nil {
		return x.Kafka
	}
	return nil
}

func (x *Data) GetRedis() *Data_Redis {
	if x != nil {
		return x.Redis
	}
	return nil
}

type Client struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RpcClient *Client_GRPC `protobuf:"bytes,1,opt,name=rpc_client,json=rpcClient,proto3" json:"rpc_client,omitempty"`
}

func (x *Client) Reset() {
	*x = Client{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Client) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client) ProtoMessage() {}

func (x *Client) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client.ProtoReflect.Descriptor instead.
func (*Client) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{3}
}

func (x *Client) GetRpcClient() *Client_GRPC {
	if x != nil {
		return x.RpcClient
	}
	return nil
}

type Server_HTTP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr    string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *Server_HTTP) Reset() {
	*x = Server_HTTP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_HTTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_HTTP) ProtoMessage() {}

func (x *Server_HTTP) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_HTTP.ProtoReflect.Descriptor instead.
func (*Server_HTTP) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Server_HTTP) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_HTTP) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_HTTP) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type Server_GRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network           string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr              string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout           *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	IdleTimeout       *durationpb.Duration `protobuf:"bytes,4,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
	MaxLifeTime       *durationpb.Duration `protobuf:"bytes,5,opt,name=max_life_time,json=maxLifeTime,proto3" json:"max_life_time,omitempty"`
	KeepAliveInterval *durationpb.Duration `protobuf:"bytes,6,opt,name=keep_alive_interval,json=keepAliveInterval,proto3" json:"keep_alive_interval,omitempty"`
	KeepAliveTimeout  *durationpb.Duration `protobuf:"bytes,7,opt,name=keep_alive_timeout,json=keepAliveTimeout,proto3" json:"keep_alive_timeout,omitempty"`
}

func (x *Server_GRPC) Reset() {
	*x = Server_GRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_GRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_GRPC) ProtoMessage() {}

func (x *Server_GRPC) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_GRPC.ProtoReflect.Descriptor instead.
func (*Server_GRPC) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Server_GRPC) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_GRPC) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_GRPC) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *Server_GRPC) GetIdleTimeout() *durationpb.Duration {
	if x != nil {
		return x.IdleTimeout
	}
	return nil
}

func (x *Server_GRPC) GetMaxLifeTime() *durationpb.Duration {
	if x != nil {
		return x.MaxLifeTime
	}
	return nil
}

func (x *Server_GRPC) GetKeepAliveInterval() *durationpb.Duration {
	if x != nil {
		return x.KeepAliveInterval
	}
	return nil
}

func (x *Server_GRPC) GetKeepAliveTimeout() *durationpb.Duration {
	if x != nil {
		return x.KeepAliveTimeout
	}
	return nil
}

type Server_Register struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address     string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Port        uint64 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	NamespaceId string `protobuf:"bytes,3,opt,name=namespaceId,proto3" json:"namespaceId,omitempty"`
	AccessKey   string `protobuf:"bytes,4,opt,name=accessKey,proto3" json:"accessKey,omitempty"`
	SecretKey   string `protobuf:"bytes,5,opt,name=secretKey,proto3" json:"secretKey,omitempty"`
	TimeoutMs   uint64 `protobuf:"varint,6,opt,name=timeoutMs,proto3" json:"timeoutMs,omitempty"`
	LogLevel    string `protobuf:"bytes,7,opt,name=logLevel,proto3" json:"logLevel,omitempty"`
}

func (x *Server_Register) Reset() {
	*x = Server_Register{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_Register) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_Register) ProtoMessage() {}

func (x *Server_Register) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_Register.ProtoReflect.Descriptor instead.
func (*Server_Register) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{1, 2}
}

func (x *Server_Register) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Server_Register) GetPort() uint64 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Server_Register) GetNamespaceId() string {
	if x != nil {
		return x.NamespaceId
	}
	return ""
}

func (x *Server_Register) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *Server_Register) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *Server_Register) GetTimeoutMs() uint64 {
	if x != nil {
		return x.TimeoutMs
	}
	return 0
}

func (x *Server_Register) GetLogLevel() string {
	if x != nil {
		return x.LogLevel
	}
	return ""
}

type Data_Kafka struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic   string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Brokers []string `protobuf:"bytes,2,rep,name=brokers,proto3" json:"brokers,omitempty"`
}

func (x *Data_Kafka) Reset() {
	*x = Data_Kafka{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_Kafka) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Kafka) ProtoMessage() {}

func (x *Data_Kafka) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Kafka.ProtoReflect.Descriptor instead.
func (*Data_Kafka) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Data_Kafka) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Data_Kafka) GetBrokers() []string {
	if x != nil {
		return x.Brokers
	}
	return nil
}

type Data_Redis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network      string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr         string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Auth         string               `protobuf:"bytes,3,opt,name=auth,proto3" json:"auth,omitempty"`
	Active       int64                `protobuf:"varint,4,opt,name=active,proto3" json:"active,omitempty"`
	Idle         int64                `protobuf:"varint,5,opt,name=idle,proto3" json:"idle,omitempty"`
	DailTimeout  *durationpb.Duration `protobuf:"bytes,6,opt,name=dail_timeout,json=dailTimeout,proto3" json:"dail_timeout,omitempty"`
	ReadTimeout  *durationpb.Duration `protobuf:"bytes,7,opt,name=read_timeout,json=readTimeout,proto3" json:"read_timeout,omitempty"`
	WriteTimeout *durationpb.Duration `protobuf:"bytes,8,opt,name=write_timeout,json=writeTimeout,proto3" json:"write_timeout,omitempty"`
	IdleTimeout  *durationpb.Duration `protobuf:"bytes,9,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
	Expire       *durationpb.Duration `protobuf:"bytes,10,opt,name=expire,proto3" json:"expire,omitempty"`
}

func (x *Data_Redis) Reset() {
	*x = Data_Redis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_Redis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Redis) ProtoMessage() {}

func (x *Data_Redis) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Redis.ProtoReflect.Descriptor instead.
func (*Data_Redis) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{2, 1}
}

func (x *Data_Redis) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Data_Redis) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Data_Redis) GetAuth() string {
	if x != nil {
		return x.Auth
	}
	return ""
}

func (x *Data_Redis) GetActive() int64 {
	if x != nil {
		return x.Active
	}
	return 0
}

func (x *Data_Redis) GetIdle() int64 {
	if x != nil {
		return x.Idle
	}
	return 0
}

func (x *Data_Redis) GetDailTimeout() *durationpb.Duration {
	if x != nil {
		return x.DailTimeout
	}
	return nil
}

func (x *Data_Redis) GetReadTimeout() *durationpb.Duration {
	if x != nil {
		return x.ReadTimeout
	}
	return nil
}

func (x *Data_Redis) GetWriteTimeout() *durationpb.Duration {
	if x != nil {
		return x.WriteTimeout
	}
	return nil
}

func (x *Data_Redis) GetIdleTimeout() *durationpb.Duration {
	if x != nil {
		return x.IdleTimeout
	}
	return nil
}

func (x *Data_Redis) GetExpire() *durationpb.Duration {
	if x != nil {
		return x.Expire
	}
	return nil
}

type Client_GRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dail    *durationpb.Duration `protobuf:"bytes,1,opt,name=dail,proto3" json:"dail,omitempty"`
	Timeout *durationpb.Duration `protobuf:"bytes,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *Client_GRPC) Reset() {
	*x = Client_GRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Client_GRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client_GRPC) ProtoMessage() {}

func (x *Client_GRPC) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client_GRPC.ProtoReflect.Descriptor instead.
func (*Client_GRPC) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{3, 0}
}

func (x *Client_GRPC) GetDail() *durationpb.Duration {
	if x != nil {
		return x.Dail
	}
	return nil
}

func (x *Client_GRPC) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

var File_conf_conf_proto protoreflect.FileDescriptor

var file_conf_conf_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01,
	0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x12, 0x2a, 0x0a, 0x06, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x72,
	0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52,
	0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a,
	0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0xd6, 0x06, 0x0a, 0x06, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x04, 0x68, 0x74, 0x74,
	0x70, 0x12, 0x2b, 0x0a, 0x04, 0x67, 0x72, 0x70, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x04, 0x67, 0x72, 0x70, 0x63, 0x12, 0x37,
	0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x08, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x69, 0x0a, 0x04, 0x48, 0x54, 0x54, 0x50, 0x12,
	0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x33, 0x0a,
	0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x1a, 0xfa, 0x02, 0x0a, 0x04, 0x47, 0x52, 0x50, 0x43, 0x12, 0x18, 0x0a, 0x07, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3c,
	0x0a, 0x0c, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x69, 0x64, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3d, 0x0a, 0x0d,
	0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b,
	0x6d, 0x61, 0x78, 0x4c, 0x69, 0x66, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x49, 0x0a, 0x13, 0x6b,
	0x65, 0x65, 0x70, 0x5f, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x11, 0x6b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x47, 0x0a, 0x12, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x61,
	0x6c, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x6b,
	0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x1a,
	0xd0, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x4d, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x4d, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x22, 0xc0, 0x04, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x05, 0x6b,
	0x61, 0x66, 0x6b, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x4b, 0x61, 0x66,
	0x6b, 0x61, 0x52, 0x05, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x12, 0x2c, 0x0a, 0x05, 0x72, 0x65, 0x64,
	0x69, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73,
	0x52, 0x05, 0x72, 0x65, 0x64, 0x69, 0x73, 0x1a, 0x37, 0x0a, 0x05, 0x4b, 0x61, 0x66, 0x6b, 0x61,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x73,
	0x1a, 0xa2, 0x03, 0x0a, 0x05, 0x52, 0x65, 0x64, 0x69, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x64, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x69, 0x64, 0x6c, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x64, 0x61, 0x69, 0x6c,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x61, 0x69, 0x6c, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3c, 0x0a, 0x0c, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x12, 0x3e, 0x0a, 0x0d, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x77, 0x72, 0x69, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x12, 0x3c, 0x0a, 0x0c, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x69, 0x64, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x22, 0xac, 0x01, 0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x36, 0x0a, 0x0a, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x09, 0x72,
	0x70, 0x63, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x6a, 0x0a, 0x04, 0x47, 0x52, 0x50, 0x43,
	0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x69, 0x6c, 0x12,
	0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x42, 0x1c, 0x5a, 0x1a, 0x61, 0x72, 0x6f, 0x64, 0x2d, 0x69, 0x6d, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f,
	0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conf_conf_proto_rawDescOnce sync.Once
	file_conf_conf_proto_rawDescData = file_conf_conf_proto_rawDesc
)

func file_conf_conf_proto_rawDescGZIP() []byte {
	file_conf_conf_proto_rawDescOnce.Do(func() {
		file_conf_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_conf_conf_proto_rawDescData)
	})
	return file_conf_conf_proto_rawDescData
}

var file_conf_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_conf_conf_proto_goTypes = []interface{}{
	(*Bootstrap)(nil),           // 0: kratos.api.Bootstrap
	(*Server)(nil),              // 1: kratos.api.Server
	(*Data)(nil),                // 2: kratos.api.Data
	(*Client)(nil),              // 3: kratos.api.Client
	(*Server_HTTP)(nil),         // 4: kratos.api.Server.HTTP
	(*Server_GRPC)(nil),         // 5: kratos.api.Server.GRPC
	(*Server_Register)(nil),     // 6: kratos.api.Server.Register
	(*Data_Kafka)(nil),          // 7: kratos.api.Data.Kafka
	(*Data_Redis)(nil),          // 8: kratos.api.Data.Redis
	(*Client_GRPC)(nil),         // 9: kratos.api.Client.GRPC
	(*durationpb.Duration)(nil), // 10: google.protobuf.Duration
}
var file_conf_conf_proto_depIdxs = []int32{
	1,  // 0: kratos.api.Bootstrap.server:type_name -> kratos.api.Server
	2,  // 1: kratos.api.Bootstrap.data:type_name -> kratos.api.Data
	3,  // 2: kratos.api.Bootstrap.client:type_name -> kratos.api.Client
	4,  // 3: kratos.api.Server.http:type_name -> kratos.api.Server.HTTP
	5,  // 4: kratos.api.Server.grpc:type_name -> kratos.api.Server.GRPC
	6,  // 5: kratos.api.Server.register:type_name -> kratos.api.Server.Register
	7,  // 6: kratos.api.Data.kafka:type_name -> kratos.api.Data.Kafka
	8,  // 7: kratos.api.Data.redis:type_name -> kratos.api.Data.Redis
	9,  // 8: kratos.api.Client.rpc_client:type_name -> kratos.api.Client.GRPC
	10, // 9: kratos.api.Server.HTTP.timeout:type_name -> google.protobuf.Duration
	10, // 10: kratos.api.Server.GRPC.timeout:type_name -> google.protobuf.Duration
	10, // 11: kratos.api.Server.GRPC.idle_timeout:type_name -> google.protobuf.Duration
	10, // 12: kratos.api.Server.GRPC.max_life_time:type_name -> google.protobuf.Duration
	10, // 13: kratos.api.Server.GRPC.keep_alive_interval:type_name -> google.protobuf.Duration
	10, // 14: kratos.api.Server.GRPC.keep_alive_timeout:type_name -> google.protobuf.Duration
	10, // 15: kratos.api.Data.Redis.dail_timeout:type_name -> google.protobuf.Duration
	10, // 16: kratos.api.Data.Redis.read_timeout:type_name -> google.protobuf.Duration
	10, // 17: kratos.api.Data.Redis.write_timeout:type_name -> google.protobuf.Duration
	10, // 18: kratos.api.Data.Redis.idle_timeout:type_name -> google.protobuf.Duration
	10, // 19: kratos.api.Data.Redis.expire:type_name -> google.protobuf.Duration
	10, // 20: kratos.api.Client.GRPC.dail:type_name -> google.protobuf.Duration
	10, // 21: kratos.api.Client.GRPC.timeout:type_name -> google.protobuf.Duration
	22, // [22:22] is the sub-list for method output_type
	22, // [22:22] is the sub-list for method input_type
	22, // [22:22] is the sub-list for extension type_name
	22, // [22:22] is the sub-list for extension extendee
	0,  // [0:22] is the sub-list for field type_name
}

func init() { file_conf_conf_proto_init() }
func file_conf_conf_proto_init() {
	if File_conf_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_conf_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bootstrap); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_conf_conf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_conf_conf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_conf_conf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Client); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_conf_conf_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_HTTP); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_conf_conf_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_GRPC); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_conf_conf_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_Register); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_conf_conf_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_Kafka); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_conf_conf_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_Redis); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_conf_conf_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Client_GRPC); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_conf_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_conf_proto_goTypes,
		DependencyIndexes: file_conf_conf_proto_depIdxs,
		MessageInfos:      file_conf_conf_proto_msgTypes,
	}.Build()
	File_conf_conf_proto = out.File
	file_conf_conf_proto_rawDesc = nil
	file_conf_conf_proto_goTypes = nil
	file_conf_conf_proto_depIdxs = nil
}
