// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/core/health_check.proto

package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf2 "github.com/golang/protobuf/ptypes/duration"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// [#not-implemented-hide:] Part of HDS.
type HealthStatus int32

const (
	// UNKNOWN should be treated by Envoy as HEALTHY.
	HealthStatus_UNKNOWN HealthStatus = 0
	// Healthy.
	HealthStatus_HEALTHY HealthStatus = 1
	// Unhealthy.
	HealthStatus_UNHEALTHY HealthStatus = 2
	// Connection draining in progress. E.g.,
	// https://aws.amazon.com/blogs/aws/elb-connection-draining-remove-instances-from-service-with-care/
	// or
	// https://cloud.google.com/compute/docs/load-balancing/enabling-connection-draining.
	HealthStatus_DRAINING HealthStatus = 3
	// This value is used by HDS Remote server. From Envoy’s perspective
	// TIMEOUT = UNHEALTHY in case EDS returns HealthStatus.
	HealthStatus_TIMEOUT HealthStatus = 4
)

var HealthStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "HEALTHY",
	2: "UNHEALTHY",
	3: "DRAINING",
	4: "TIMEOUT",
}
var HealthStatus_value = map[string]int32{
	"UNKNOWN":   0,
	"HEALTHY":   1,
	"UNHEALTHY": 2,
	"DRAINING":  3,
	"TIMEOUT":   4,
}

func (x HealthStatus) String() string {
	return proto.EnumName(HealthStatus_name, int32(x))
}
func (HealthStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

type HealthCheck struct {
	// The time to wait for a health check response. If the timeout is reached the
	// health check attempt will be considered a failure.
	Timeout *google_protobuf2.Duration `protobuf:"bytes,1,opt,name=timeout" json:"timeout,omitempty"`
	// The interval between health checks.
	Interval *google_protobuf2.Duration `protobuf:"bytes,2,opt,name=interval" json:"interval,omitempty"`
	// An optional jitter amount in millseconds. If specified, during every
	// internal Envoy will add 0 to interval_jitter to the wait time.
	IntervalJitter *google_protobuf2.Duration `protobuf:"bytes,3,opt,name=interval_jitter,json=intervalJitter" json:"interval_jitter,omitempty"`
	// The number of unhealthy health checks required before a host is marked
	// unhealthy. Note that for *http* health checking if a host responds with 503
	// this threshold is ignored and the host is considered unhealthy immediately.
	UnhealthyThreshold *google_protobuf.UInt32Value `protobuf:"bytes,4,opt,name=unhealthy_threshold,json=unhealthyThreshold" json:"unhealthy_threshold,omitempty"`
	// The number of healthy health checks required before a host is marked
	// healthy. Note that during startup, only a single successful health check is
	// required to mark a host healthy.
	HealthyThreshold *google_protobuf.UInt32Value `protobuf:"bytes,5,opt,name=healthy_threshold,json=healthyThreshold" json:"healthy_threshold,omitempty"`
	// [#not-implemented-hide:] Non-serving port for health checking.
	AltPort *google_protobuf.UInt32Value `protobuf:"bytes,6,opt,name=alt_port,json=altPort" json:"alt_port,omitempty"`
	// Reuse health check connection between health checks. Default is true.
	ReuseConnection *google_protobuf.BoolValue `protobuf:"bytes,7,opt,name=reuse_connection,json=reuseConnection" json:"reuse_connection,omitempty"`
	// Types that are valid to be assigned to HealthChecker:
	//	*HealthCheck_HttpHealthCheck_
	//	*HealthCheck_TcpHealthCheck_
	//	*HealthCheck_RedisHealthCheck_
	//	*HealthCheck_GrpcHealthCheck_
	HealthChecker isHealthCheck_HealthChecker `protobuf_oneof:"health_checker"`
}

func (m *HealthCheck) Reset()                    { *m = HealthCheck{} }
func (m *HealthCheck) String() string            { return proto.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()               {}
func (*HealthCheck) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

type isHealthCheck_HealthChecker interface {
	isHealthCheck_HealthChecker()
}

type HealthCheck_HttpHealthCheck_ struct {
	HttpHealthCheck *HealthCheck_HttpHealthCheck `protobuf:"bytes,8,opt,name=http_health_check,json=httpHealthCheck,oneof"`
}
type HealthCheck_TcpHealthCheck_ struct {
	TcpHealthCheck *HealthCheck_TcpHealthCheck `protobuf:"bytes,9,opt,name=tcp_health_check,json=tcpHealthCheck,oneof"`
}
type HealthCheck_RedisHealthCheck_ struct {
	RedisHealthCheck *HealthCheck_RedisHealthCheck `protobuf:"bytes,10,opt,name=redis_health_check,json=redisHealthCheck,oneof"`
}
type HealthCheck_GrpcHealthCheck_ struct {
	GrpcHealthCheck *HealthCheck_GrpcHealthCheck `protobuf:"bytes,11,opt,name=grpc_health_check,json=grpcHealthCheck,oneof"`
}

func (*HealthCheck_HttpHealthCheck_) isHealthCheck_HealthChecker()  {}
func (*HealthCheck_TcpHealthCheck_) isHealthCheck_HealthChecker()   {}
func (*HealthCheck_RedisHealthCheck_) isHealthCheck_HealthChecker() {}
func (*HealthCheck_GrpcHealthCheck_) isHealthCheck_HealthChecker()  {}

func (m *HealthCheck) GetHealthChecker() isHealthCheck_HealthChecker {
	if m != nil {
		return m.HealthChecker
	}
	return nil
}

func (m *HealthCheck) GetTimeout() *google_protobuf2.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *HealthCheck) GetInterval() *google_protobuf2.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *HealthCheck) GetIntervalJitter() *google_protobuf2.Duration {
	if m != nil {
		return m.IntervalJitter
	}
	return nil
}

func (m *HealthCheck) GetUnhealthyThreshold() *google_protobuf.UInt32Value {
	if m != nil {
		return m.UnhealthyThreshold
	}
	return nil
}

func (m *HealthCheck) GetHealthyThreshold() *google_protobuf.UInt32Value {
	if m != nil {
		return m.HealthyThreshold
	}
	return nil
}

func (m *HealthCheck) GetAltPort() *google_protobuf.UInt32Value {
	if m != nil {
		return m.AltPort
	}
	return nil
}

func (m *HealthCheck) GetReuseConnection() *google_protobuf.BoolValue {
	if m != nil {
		return m.ReuseConnection
	}
	return nil
}

func (m *HealthCheck) GetHttpHealthCheck() *HealthCheck_HttpHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_HttpHealthCheck_); ok {
		return x.HttpHealthCheck
	}
	return nil
}

func (m *HealthCheck) GetTcpHealthCheck() *HealthCheck_TcpHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_TcpHealthCheck_); ok {
		return x.TcpHealthCheck
	}
	return nil
}

func (m *HealthCheck) GetRedisHealthCheck() *HealthCheck_RedisHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_RedisHealthCheck_); ok {
		return x.RedisHealthCheck
	}
	return nil
}

func (m *HealthCheck) GetGrpcHealthCheck() *HealthCheck_GrpcHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_GrpcHealthCheck_); ok {
		return x.GrpcHealthCheck
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*HealthCheck) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _HealthCheck_OneofMarshaler, _HealthCheck_OneofUnmarshaler, _HealthCheck_OneofSizer, []interface{}{
		(*HealthCheck_HttpHealthCheck_)(nil),
		(*HealthCheck_TcpHealthCheck_)(nil),
		(*HealthCheck_RedisHealthCheck_)(nil),
		(*HealthCheck_GrpcHealthCheck_)(nil),
	}
}

func _HealthCheck_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*HealthCheck)
	// health_checker
	switch x := m.HealthChecker.(type) {
	case *HealthCheck_HttpHealthCheck_:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HttpHealthCheck); err != nil {
			return err
		}
	case *HealthCheck_TcpHealthCheck_:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TcpHealthCheck); err != nil {
			return err
		}
	case *HealthCheck_RedisHealthCheck_:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RedisHealthCheck); err != nil {
			return err
		}
	case *HealthCheck_GrpcHealthCheck_:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GrpcHealthCheck); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("HealthCheck.HealthChecker has unexpected type %T", x)
	}
	return nil
}

func _HealthCheck_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*HealthCheck)
	switch tag {
	case 8: // health_checker.http_health_check
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HealthCheck_HttpHealthCheck)
		err := b.DecodeMessage(msg)
		m.HealthChecker = &HealthCheck_HttpHealthCheck_{msg}
		return true, err
	case 9: // health_checker.tcp_health_check
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HealthCheck_TcpHealthCheck)
		err := b.DecodeMessage(msg)
		m.HealthChecker = &HealthCheck_TcpHealthCheck_{msg}
		return true, err
	case 10: // health_checker.redis_health_check
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HealthCheck_RedisHealthCheck)
		err := b.DecodeMessage(msg)
		m.HealthChecker = &HealthCheck_RedisHealthCheck_{msg}
		return true, err
	case 11: // health_checker.grpc_health_check
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HealthCheck_GrpcHealthCheck)
		err := b.DecodeMessage(msg)
		m.HealthChecker = &HealthCheck_GrpcHealthCheck_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _HealthCheck_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*HealthCheck)
	// health_checker
	switch x := m.HealthChecker.(type) {
	case *HealthCheck_HttpHealthCheck_:
		s := proto.Size(x.HttpHealthCheck)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *HealthCheck_TcpHealthCheck_:
		s := proto.Size(x.TcpHealthCheck)
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *HealthCheck_RedisHealthCheck_:
		s := proto.Size(x.RedisHealthCheck)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *HealthCheck_GrpcHealthCheck_:
		s := proto.Size(x.GrpcHealthCheck)
		n += proto.SizeVarint(11<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Describes the encoding of the payload bytes in the payload.
type HealthCheck_Payload struct {
	// Types that are valid to be assigned to Payload:
	//	*HealthCheck_Payload_Text
	//	*HealthCheck_Payload_Binary
	Payload isHealthCheck_Payload_Payload `protobuf_oneof:"payload"`
}

func (m *HealthCheck_Payload) Reset()                    { *m = HealthCheck_Payload{} }
func (m *HealthCheck_Payload) String() string            { return proto.CompactTextString(m) }
func (*HealthCheck_Payload) ProtoMessage()               {}
func (*HealthCheck_Payload) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0, 0} }

type isHealthCheck_Payload_Payload interface {
	isHealthCheck_Payload_Payload()
}

type HealthCheck_Payload_Text struct {
	Text string `protobuf:"bytes,1,opt,name=text,oneof"`
}
type HealthCheck_Payload_Binary struct {
	Binary []byte `protobuf:"bytes,2,opt,name=binary,proto3,oneof"`
}

func (*HealthCheck_Payload_Text) isHealthCheck_Payload_Payload()   {}
func (*HealthCheck_Payload_Binary) isHealthCheck_Payload_Payload() {}

func (m *HealthCheck_Payload) GetPayload() isHealthCheck_Payload_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *HealthCheck_Payload) GetText() string {
	if x, ok := m.GetPayload().(*HealthCheck_Payload_Text); ok {
		return x.Text
	}
	return ""
}

func (m *HealthCheck_Payload) GetBinary() []byte {
	if x, ok := m.GetPayload().(*HealthCheck_Payload_Binary); ok {
		return x.Binary
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*HealthCheck_Payload) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _HealthCheck_Payload_OneofMarshaler, _HealthCheck_Payload_OneofUnmarshaler, _HealthCheck_Payload_OneofSizer, []interface{}{
		(*HealthCheck_Payload_Text)(nil),
		(*HealthCheck_Payload_Binary)(nil),
	}
}

func _HealthCheck_Payload_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*HealthCheck_Payload)
	// payload
	switch x := m.Payload.(type) {
	case *HealthCheck_Payload_Text:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Text)
	case *HealthCheck_Payload_Binary:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.Binary)
	case nil:
	default:
		return fmt.Errorf("HealthCheck_Payload.Payload has unexpected type %T", x)
	}
	return nil
}

func _HealthCheck_Payload_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*HealthCheck_Payload)
	switch tag {
	case 1: // payload.text
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Payload = &HealthCheck_Payload_Text{x}
		return true, err
	case 2: // payload.binary
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Payload = &HealthCheck_Payload_Binary{x}
		return true, err
	default:
		return false, nil
	}
}

func _HealthCheck_Payload_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*HealthCheck_Payload)
	// payload
	switch x := m.Payload.(type) {
	case *HealthCheck_Payload_Text:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Text)))
		n += len(x.Text)
	case *HealthCheck_Payload_Binary:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Binary)))
		n += len(x.Binary)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type HealthCheck_HttpHealthCheck struct {
	// The value of the host header in the HTTP health check request. If
	// left empty (default value), the IP on behalf of which this health check is performed will be
	// used.
	Host string `protobuf:"bytes,1,opt,name=host" json:"host,omitempty"`
	// Specifies the HTTP path that will be requested during health checking. For example
	// */healthcheck*.
	Path string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	// [#not-implemented-hide:] HTTP specific payload.
	Send *HealthCheck_Payload `protobuf:"bytes,3,opt,name=send" json:"send,omitempty"`
	// [#not-implemented-hide:] HTTP specific response.
	Receive *HealthCheck_Payload `protobuf:"bytes,4,opt,name=receive" json:"receive,omitempty"`
	// An optional service name parameter which is used to validate the identity of
	// the health checked cluster. See the :ref:`architecture overview
	// <arch_overview_health_checking_identity>` for more information.
	ServiceName string `protobuf:"bytes,5,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
}

func (m *HealthCheck_HttpHealthCheck) Reset()                    { *m = HealthCheck_HttpHealthCheck{} }
func (m *HealthCheck_HttpHealthCheck) String() string            { return proto.CompactTextString(m) }
func (*HealthCheck_HttpHealthCheck) ProtoMessage()               {}
func (*HealthCheck_HttpHealthCheck) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0, 1} }

func (m *HealthCheck_HttpHealthCheck) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *HealthCheck_HttpHealthCheck) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *HealthCheck_HttpHealthCheck) GetSend() *HealthCheck_Payload {
	if m != nil {
		return m.Send
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetReceive() *HealthCheck_Payload {
	if m != nil {
		return m.Receive
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

type HealthCheck_TcpHealthCheck struct {
	// Empty payloads imply a connect-only health check.
	Send *HealthCheck_Payload `protobuf:"bytes,1,opt,name=send" json:"send,omitempty"`
	// When checking the response, “fuzzy” matching is performed such that each
	// binary block must be found, and in the order specified, but not
	// necessarily contiguous.
	Receive []*HealthCheck_Payload `protobuf:"bytes,2,rep,name=receive" json:"receive,omitempty"`
}

func (m *HealthCheck_TcpHealthCheck) Reset()                    { *m = HealthCheck_TcpHealthCheck{} }
func (m *HealthCheck_TcpHealthCheck) String() string            { return proto.CompactTextString(m) }
func (*HealthCheck_TcpHealthCheck) ProtoMessage()               {}
func (*HealthCheck_TcpHealthCheck) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0, 2} }

func (m *HealthCheck_TcpHealthCheck) GetSend() *HealthCheck_Payload {
	if m != nil {
		return m.Send
	}
	return nil
}

func (m *HealthCheck_TcpHealthCheck) GetReceive() []*HealthCheck_Payload {
	if m != nil {
		return m.Receive
	}
	return nil
}

type HealthCheck_RedisHealthCheck struct {
	// If set, optionally perform ``EXISTS <key>`` instead of ``PING``. A return value
	// from Redis of 0 (does not exist) is considered a passing healthcheck. A return value other
	// than 0 is considered a failure. This allows the user to mark a Redis instance for maintenance
	// by setting the specified key to any value and waiting for traffic to drain.
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *HealthCheck_RedisHealthCheck) Reset()                    { *m = HealthCheck_RedisHealthCheck{} }
func (m *HealthCheck_RedisHealthCheck) String() string            { return proto.CompactTextString(m) }
func (*HealthCheck_RedisHealthCheck) ProtoMessage()               {}
func (*HealthCheck_RedisHealthCheck) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0, 3} }

func (m *HealthCheck_RedisHealthCheck) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

// `grpc.health.v1.Health
// <https://github.com/grpc/grpc/blob/master/src/proto/grpc/health/v1/health.proto>`_-based
// healthcheck. See `gRPC doc <https://github.com/grpc/grpc/blob/master/doc/health-checking.md>`_
// for details.
type HealthCheck_GrpcHealthCheck struct {
	// An optional service name parameter which will be sent to gRPC service in
	// `grpc.health.v1.HealthCheckRequest
	// <https://github.com/grpc/grpc/blob/master/src/proto/grpc/health/v1/health.proto#L20>`_.
	// message. See `gRPC health-checking overview
	// <https://github.com/grpc/grpc/blob/master/doc/health-checking.md>`_ for more information.
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
}

func (m *HealthCheck_GrpcHealthCheck) Reset()                    { *m = HealthCheck_GrpcHealthCheck{} }
func (m *HealthCheck_GrpcHealthCheck) String() string            { return proto.CompactTextString(m) }
func (*HealthCheck_GrpcHealthCheck) ProtoMessage()               {}
func (*HealthCheck_GrpcHealthCheck) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0, 4} }

func (m *HealthCheck_GrpcHealthCheck) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func init() {
	proto.RegisterType((*HealthCheck)(nil), "envoy.api.v2.core.HealthCheck")
	proto.RegisterType((*HealthCheck_Payload)(nil), "envoy.api.v2.core.HealthCheck.Payload")
	proto.RegisterType((*HealthCheck_HttpHealthCheck)(nil), "envoy.api.v2.core.HealthCheck.HttpHealthCheck")
	proto.RegisterType((*HealthCheck_TcpHealthCheck)(nil), "envoy.api.v2.core.HealthCheck.TcpHealthCheck")
	proto.RegisterType((*HealthCheck_RedisHealthCheck)(nil), "envoy.api.v2.core.HealthCheck.RedisHealthCheck")
	proto.RegisterType((*HealthCheck_GrpcHealthCheck)(nil), "envoy.api.v2.core.HealthCheck.GrpcHealthCheck")
	proto.RegisterEnum("envoy.api.v2.core.HealthStatus", HealthStatus_name, HealthStatus_value)
}

func init() { proto.RegisterFile("envoy/api/v2/core/health_check.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 711 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0xc7, 0xe3, 0xc4, 0x6d, 0x92, 0x93, 0xdc, 0xc4, 0x9d, 0x7b, 0xa5, 0xfa, 0x46, 0x50, 0x0a,
	0xaa, 0x10, 0x42, 0xc2, 0x96, 0x52, 0x24, 0x24, 0x36, 0xd0, 0xb4, 0x55, 0x13, 0xa0, 0x69, 0x65,
	0x92, 0xa2, 0x4a, 0x48, 0xd6, 0xd4, 0x19, 0x62, 0x53, 0xd7, 0x63, 0x8d, 0x27, 0x81, 0xbc, 0x04,
	0xcf, 0x81, 0x90, 0xd8, 0x23, 0x56, 0x7d, 0x16, 0x24, 0x16, 0x7d, 0x0b, 0xe4, 0xb1, 0x1d, 0xc5,
	0xf6, 0x22, 0x01, 0x76, 0x67, 0xce, 0xcc, 0xff, 0x77, 0x3e, 0x7c, 0x8e, 0x61, 0x87, 0x78, 0x53,
	0x3a, 0xd3, 0xb1, 0xef, 0xe8, 0xd3, 0xb6, 0x6e, 0x51, 0x46, 0x74, 0x9b, 0x60, 0x97, 0xdb, 0xa6,
	0x65, 0x13, 0xeb, 0x52, 0xf3, 0x19, 0xe5, 0x14, 0x6d, 0x88, 0x57, 0x1a, 0xf6, 0x1d, 0x6d, 0xda,
	0xd6, 0xc2, 0x57, 0xad, 0xad, 0x31, 0xa5, 0x63, 0x97, 0xe8, 0xe2, 0xc1, 0xc5, 0xe4, 0x9d, 0x3e,
	0x9a, 0x30, 0xcc, 0x1d, 0xea, 0x45, 0x92, 0xfc, 0xfd, 0x07, 0x86, 0x7d, 0x9f, 0xb0, 0x20, 0xbe,
	0xdf, 0x9c, 0x62, 0xd7, 0x19, 0x61, 0x4e, 0xf4, 0xc4, 0x88, 0x2f, 0xfe, 0x1b, 0xd3, 0x31, 0x15,
	0xa6, 0x1e, 0x5a, 0x91, 0xf7, 0xde, 0xd7, 0x1a, 0xd4, 0xba, 0x22, 0xb1, 0xfd, 0x30, 0x2f, 0xf4,
	0x0c, 0xca, 0xdc, 0xb9, 0x22, 0x74, 0xc2, 0x55, 0x69, 0x5b, 0x7a, 0x50, 0x6b, 0xff, 0xaf, 0x45,
	0x01, 0xb5, 0x24, 0xa0, 0x76, 0x10, 0x27, 0xd4, 0x81, 0xef, 0x37, 0xd7, 0xa5, 0xb5, 0x2f, 0x52,
	0xb1, 0x22, 0x19, 0x89, 0x0a, 0xed, 0x41, 0xc5, 0xf1, 0x38, 0x61, 0x53, 0xec, 0xaa, 0xc5, 0xdf,
	0x21, 0xcc, 0x65, 0xa8, 0x03, 0xcd, 0xc4, 0x36, 0xdf, 0x3b, 0x9c, 0x13, 0xa6, 0x96, 0x96, 0x90,
	0x8c, 0x46, 0xa2, 0x78, 0x21, 0x04, 0xe8, 0x18, 0xfe, 0x9d, 0x78, 0x51, 0xc7, 0x67, 0x26, 0xb7,
	0x19, 0x09, 0x6c, 0xea, 0x8e, 0x54, 0x59, 0x70, 0x6e, 0xe5, 0x38, 0xc3, 0x9e, 0xc7, 0x77, 0xdb,
	0x67, 0xd8, 0x9d, 0x10, 0x03, 0xcd, 0x85, 0x83, 0x44, 0x87, 0x7a, 0xb0, 0x91, 0x87, 0xad, 0xad,
	0x00, 0x53, 0x72, 0xa8, 0x27, 0x50, 0xc1, 0x2e, 0x37, 0x7d, 0xca, 0xb8, 0xba, 0xbe, 0x02, 0xa1,
	0x8c, 0x5d, 0x7e, 0x4a, 0x19, 0x47, 0x87, 0xa0, 0x30, 0x32, 0x09, 0x88, 0x69, 0x51, 0xcf, 0x23,
	0x56, 0x58, 0xb6, 0x5a, 0x16, 0x80, 0x56, 0x0e, 0xd0, 0xa1, 0xd4, 0x8d, 0xe4, 0x4d, 0xa1, 0xd9,
	0x9f, 0x4b, 0xd0, 0x5b, 0xd8, 0xb0, 0x39, 0xf7, 0xcd, 0xc5, 0x71, 0x54, 0x2b, 0x82, 0xa3, 0x69,
	0xb9, 0x79, 0xd4, 0x16, 0x86, 0x43, 0xeb, 0x72, 0xee, 0x2f, 0x9c, 0xbb, 0x05, 0xa3, 0x69, 0xa7,
	0x5d, 0xe8, 0x1c, 0x14, 0x6e, 0x65, 0xe0, 0x55, 0x01, 0x7f, 0xb4, 0x04, 0x3e, 0xb0, 0x32, 0xec,
	0x06, 0x4f, 0x79, 0x90, 0x09, 0x88, 0x91, 0x91, 0x13, 0xa4, 0xe1, 0x20, 0xe0, 0xfa, 0x12, 0xb8,
	0x11, 0x0a, 0xd3, 0x78, 0x85, 0x65, 0x7c, 0x61, 0x67, 0xc6, 0xcc, 0xb7, 0xd2, 0xfc, 0xda, 0x4a,
	0x9d, 0x39, 0x62, 0xbe, 0x95, 0xe9, 0xcc, 0x38, 0xed, 0x6a, 0x9d, 0x41, 0xf9, 0x14, 0xcf, 0x5c,
	0x8a, 0x47, 0xe8, 0x0e, 0xc8, 0x9c, 0x7c, 0x8c, 0x36, 0xac, 0xda, 0xa9, 0x86, 0x4b, 0x20, 0xb3,
	0xe2, 0xb6, 0xd4, 0x2d, 0x18, 0xe2, 0x02, 0xa9, 0xb0, 0x7e, 0xe1, 0x78, 0x98, 0xcd, 0xc4, 0x0a,
	0xd5, 0xbb, 0x05, 0x23, 0x3e, 0x77, 0x14, 0x28, 0xfb, 0x31, 0x65, 0xed, 0xdb, 0xcd, 0x75, 0x49,
	0x6a, 0xfd, 0x94, 0xa0, 0x99, 0xf9, 0x30, 0x08, 0x81, 0x6c, 0xd3, 0x20, 0x0e, 0x60, 0x08, 0x1b,
	0xdd, 0x06, 0xd9, 0xc7, 0xdc, 0x16, 0xc4, 0xc5, 0xa0, 0x86, 0x70, 0xa3, 0xa7, 0x20, 0x07, 0xc4,
	0x1b, 0xc5, 0x9b, 0x76, 0x7f, 0x49, 0xbd, 0x71, 0x25, 0x86, 0xd0, 0xa0, 0xe7, 0x50, 0x66, 0xc4,
	0x22, 0xce, 0x94, 0xc4, 0x0b, 0xb6, 0xaa, 0x3c, 0x91, 0xa1, 0xbb, 0x50, 0x0f, 0x08, 0x9b, 0x3a,
	0x16, 0x31, 0x3d, 0x7c, 0x45, 0xc4, 0x6a, 0x55, 0x8d, 0x5a, 0xec, 0xeb, 0xe3, 0x2b, 0xd2, 0xfa,
	0x24, 0x41, 0x23, 0x3d, 0x23, 0xf3, 0x9c, 0xa5, 0xbf, 0xcb, 0xb9, 0xb8, 0x5d, 0xfa, 0x83, 0x9c,
	0x5b, 0x3b, 0xa0, 0x64, 0xc7, 0x0a, 0x29, 0x50, 0xba, 0x24, 0xb3, 0xb8, 0xef, 0xa1, 0xd9, 0x7a,
	0x0c, 0xcd, 0xcc, 0x70, 0xe4, 0x8a, 0x95, 0x72, 0xc5, 0x76, 0x36, 0xa1, 0xb1, 0x38, 0x85, 0x84,
	0xc5, 0x5f, 0xfb, 0xa1, 0x01, 0xf5, 0x08, 0xf5, 0x9a, 0x63, 0x3e, 0x09, 0x50, 0x0d, 0xca, 0xc3,
	0xfe, 0xcb, 0xfe, 0xc9, 0x9b, 0xbe, 0x52, 0x08, 0x0f, 0xdd, 0xc3, 0xbd, 0x57, 0x83, 0xee, 0xb9,
	0x22, 0xa1, 0x7f, 0xa0, 0x3a, 0xec, 0x27, 0xc7, 0x22, 0xaa, 0x43, 0xe5, 0xc0, 0xd8, 0xeb, 0xf5,
	0x7b, 0xfd, 0x23, 0xa5, 0x14, 0xbe, 0x1c, 0xf4, 0x8e, 0x0f, 0x4f, 0x86, 0x03, 0x45, 0xee, 0xc8,
	0x9f, 0x7f, 0x6c, 0x49, 0x17, 0xeb, 0xe2, 0xe7, 0xb1, 0xfb, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xa3,
	0xc7, 0x4f, 0x08, 0xba, 0x06, 0x00, 0x00,
}
