// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server/worker/datum/datum.proto

package datum

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	pps "github.com/pachyderm/pachyderm/src/client/pps"
	common "github.com/pachyderm/pachyderm/src/server/worker/common"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type State int32

const (
	State_PROCESSED State = 0
	State_FAILED    State = 1
	State_RECOVERED State = 2
)

var State_name = map[int32]string{
	0: "PROCESSED",
	1: "FAILED",
	2: "RECOVERED",
}

var State_value = map[string]int32{
	"PROCESSED": 0,
	"FAILED":    1,
	"RECOVERED": 2,
}

func (x State) String() string {
	return proto.EnumName(State_name, int32(x))
}

func (State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_96ec7427544ac634, []int{0}
}

type Meta struct {
	JobID                string            `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Inputs               []*common.InputV2 `protobuf:"bytes,2,rep,name=inputs,proto3" json:"inputs,omitempty"`
	Hash                 string            `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	State                State             `protobuf:"varint,4,opt,name=state,proto3,enum=datum.State" json:"state,omitempty"`
	Reason               string            `protobuf:"bytes,5,opt,name=reason,proto3" json:"reason,omitempty"`
	Stats                *pps.ProcessStats `protobuf:"bytes,6,opt,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Meta) Reset()         { *m = Meta{} }
func (m *Meta) String() string { return proto.CompactTextString(m) }
func (*Meta) ProtoMessage()    {}
func (*Meta) Descriptor() ([]byte, []int) {
	return fileDescriptor_96ec7427544ac634, []int{0}
}
func (m *Meta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Meta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Meta.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Meta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Meta.Merge(m, src)
}
func (m *Meta) XXX_Size() int {
	return m.Size()
}
func (m *Meta) XXX_DiscardUnknown() {
	xxx_messageInfo_Meta.DiscardUnknown(m)
}

var xxx_messageInfo_Meta proto.InternalMessageInfo

func (m *Meta) GetJobID() string {
	if m != nil {
		return m.JobID
	}
	return ""
}

func (m *Meta) GetInputs() []*common.InputV2 {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *Meta) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Meta) GetState() State {
	if m != nil {
		return m.State
	}
	return State_PROCESSED
}

func (m *Meta) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *Meta) GetStats() *pps.ProcessStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

type Stats struct {
	ProcessStats         *pps.ProcessStats `protobuf:"bytes,1,opt,name=process_stats,json=processStats,proto3" json:"process_stats,omitempty"`
	Processed            int64             `protobuf:"varint,2,opt,name=processed,proto3" json:"processed,omitempty"`
	Skipped              int64             `protobuf:"varint,3,opt,name=skipped,proto3" json:"skipped,omitempty"`
	Failed               int64             `protobuf:"varint,4,opt,name=failed,proto3" json:"failed,omitempty"`
	Recovered            int64             `protobuf:"varint,5,opt,name=recovered,proto3" json:"recovered,omitempty"`
	FailedID             string            `protobuf:"bytes,6,opt,name=failed_id,json=failedId,proto3" json:"failed_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Stats) Reset()         { *m = Stats{} }
func (m *Stats) String() string { return proto.CompactTextString(m) }
func (*Stats) ProtoMessage()    {}
func (*Stats) Descriptor() ([]byte, []int) {
	return fileDescriptor_96ec7427544ac634, []int{1}
}
func (m *Stats) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Stats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Stats.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Stats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stats.Merge(m, src)
}
func (m *Stats) XXX_Size() int {
	return m.Size()
}
func (m *Stats) XXX_DiscardUnknown() {
	xxx_messageInfo_Stats.DiscardUnknown(m)
}

var xxx_messageInfo_Stats proto.InternalMessageInfo

func (m *Stats) GetProcessStats() *pps.ProcessStats {
	if m != nil {
		return m.ProcessStats
	}
	return nil
}

func (m *Stats) GetProcessed() int64 {
	if m != nil {
		return m.Processed
	}
	return 0
}

func (m *Stats) GetSkipped() int64 {
	if m != nil {
		return m.Skipped
	}
	return 0
}

func (m *Stats) GetFailed() int64 {
	if m != nil {
		return m.Failed
	}
	return 0
}

func (m *Stats) GetRecovered() int64 {
	if m != nil {
		return m.Recovered
	}
	return 0
}

func (m *Stats) GetFailedID() string {
	if m != nil {
		return m.FailedID
	}
	return ""
}

func init() {
	proto.RegisterEnum("datum.State", State_name, State_value)
	proto.RegisterType((*Meta)(nil), "datum.Meta")
	proto.RegisterType((*Stats)(nil), "datum.Stats")
}

func init() { proto.RegisterFile("server/worker/datum/datum.proto", fileDescriptor_96ec7427544ac634) }

var fileDescriptor_96ec7427544ac634 = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0x4d, 0x13, 0x16, 0xaf, 0x83, 0x62, 0x4d, 0x28, 0x9a, 0x50, 0x1b, 0x7a, 0x59, 0xe0,
	0xd0, 0x88, 0x22, 0xed, 0xce, 0x96, 0x4c, 0x0a, 0x02, 0x6d, 0x72, 0xa5, 0x1d, 0xb8, 0x4c, 0x49,
	0xfc, 0x68, 0xb3, 0xad, 0xb1, 0x65, 0xbb, 0x43, 0xfc, 0x43, 0xc4, 0x89, 0x23, 0xa7, 0x09, 0xe5,
	0x97, 0x20, 0xdb, 0x41, 0x1b, 0x12, 0x1c, 0x92, 0xbc, 0xef, 0xfb, 0xde, 0xfb, 0x94, 0xcf, 0x7e,
	0x78, 0xaa, 0x40, 0xde, 0x82, 0x4c, 0xbf, 0x70, 0x79, 0x0d, 0x32, 0x65, 0xa5, 0xde, 0x6e, 0xdc,
	0x7b, 0x2e, 0x24, 0xd7, 0x9c, 0xf8, 0x16, 0x1c, 0xec, 0xaf, 0xf8, 0x8a, 0x5b, 0x26, 0x35, 0x95,
	0x13, 0x0f, 0xf6, 0xeb, 0x9b, 0x06, 0x5a, 0x9d, 0x0a, 0xa1, 0xcc, 0xd3, 0xb3, 0x2f, 0xff, 0xf6,
	0xac, 0xf9, 0x66, 0xc3, 0xdb, 0xfe, 0xe3, 0x5a, 0x66, 0xdf, 0x11, 0x1e, 0x7e, 0x04, 0x5d, 0x92,
	0x18, 0x07, 0x57, 0xbc, 0xba, 0x6c, 0x58, 0x84, 0x62, 0x94, 0x84, 0xc7, 0x61, 0x77, 0x37, 0xf5,
	0xdf, 0xf3, 0xaa, 0xc8, 0xa8, 0x7f, 0xc5, 0xab, 0x82, 0x91, 0x43, 0x1c, 0x34, 0xad, 0xd8, 0x6a,
	0x15, 0x0d, 0x62, 0x2f, 0xd9, 0x5d, 0x3c, 0x9d, 0xf7, 0x4e, 0x85, 0x61, 0x2f, 0x16, 0xb4, 0x97,
	0x09, 0xc1, 0xc3, 0x75, 0xa9, 0xd6, 0x91, 0x67, 0x8c, 0xa8, 0xad, 0xc9, 0x0c, 0xfb, 0x4a, 0x97,
	0x1a, 0xa2, 0x61, 0x8c, 0x92, 0x27, 0x8b, 0xd1, 0xdc, 0x45, 0x5b, 0x1a, 0x8e, 0x3a, 0x89, 0x3c,
	0xc7, 0x81, 0x84, 0x52, 0xf1, 0x36, 0xf2, 0xed, 0x64, 0x8f, 0xc8, 0xa1, 0x9b, 0x55, 0x51, 0x10,
	0xa3, 0x64, 0x77, 0xf1, 0x6c, 0x6e, 0x12, 0x9e, 0x4b, 0x5e, 0x83, 0x52, 0xc6, 0x40, 0x39, 0x03,
	0x35, 0xfb, 0x89, 0xb0, 0x6f, 0x09, 0x72, 0x84, 0xf7, 0x84, 0x6b, 0xb8, 0x74, 0xa3, 0xe8, 0x7f,
	0xa3, 0x23, 0xf1, 0x00, 0x91, 0x17, 0x38, 0xec, 0x31, 0xb0, 0x68, 0x10, 0xa3, 0xc4, 0xa3, 0xf7,
	0x04, 0x89, 0xf0, 0x63, 0x75, 0xdd, 0x08, 0x01, 0xcc, 0x66, 0xf3, 0xe8, 0x1f, 0x68, 0x7e, 0xfd,
	0x73, 0xd9, 0xdc, 0x00, 0xb3, 0xf9, 0x3c, 0xda, 0x23, 0xe3, 0x27, 0xa1, 0xe6, 0xb7, 0x20, 0x81,
	0xd9, 0x54, 0x1e, 0xbd, 0x27, 0xc8, 0x2b, 0x1c, 0xba, 0x3e, 0x73, 0xec, 0x81, 0x3d, 0xf6, 0x51,
	0x77, 0x37, 0xdd, 0x39, 0xb5, 0x64, 0x91, 0xd1, 0x1d, 0x27, 0x17, 0xec, 0xf5, 0x1b, 0x97, 0x0c,
	0xc8, 0x1e, 0x0e, 0xcf, 0xe9, 0xd9, 0x49, 0xbe, 0x5c, 0xe6, 0xd9, 0xf8, 0x11, 0xc1, 0x38, 0x38,
	0x7d, 0x57, 0x7c, 0xc8, 0xb3, 0x31, 0x32, 0x12, 0xcd, 0x4f, 0xce, 0x2e, 0x72, 0x9a, 0x67, 0xe3,
	0xc1, 0x71, 0xf6, 0xad, 0x9b, 0xa0, 0x1f, 0xdd, 0x04, 0xfd, 0xea, 0x26, 0xe8, 0xd3, 0xd1, 0xaa,
	0xd1, 0xeb, 0x6d, 0x65, 0xee, 0x2c, 0x15, 0x65, 0xbd, 0xfe, 0xca, 0x40, 0x3e, 0xac, 0x94, 0xac,
	0xd3, 0x7f, 0xac, 0x60, 0x15, 0xd8, 0x3d, 0x79, 0xfb, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x4d,
	0x4c, 0x4e, 0xa0, 0x02, 0x00, 0x00,
}

func (m *Meta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Meta) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Meta) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Stats != nil {
		{
			size, err := m.Stats.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDatum(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.Reason) > 0 {
		i -= len(m.Reason)
		copy(dAtA[i:], m.Reason)
		i = encodeVarintDatum(dAtA, i, uint64(len(m.Reason)))
		i--
		dAtA[i] = 0x2a
	}
	if m.State != 0 {
		i = encodeVarintDatum(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintDatum(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Inputs) > 0 {
		for iNdEx := len(m.Inputs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Inputs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDatum(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.JobID) > 0 {
		i -= len(m.JobID)
		copy(dAtA[i:], m.JobID)
		i = encodeVarintDatum(dAtA, i, uint64(len(m.JobID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Stats) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Stats) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Stats) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.FailedID) > 0 {
		i -= len(m.FailedID)
		copy(dAtA[i:], m.FailedID)
		i = encodeVarintDatum(dAtA, i, uint64(len(m.FailedID)))
		i--
		dAtA[i] = 0x32
	}
	if m.Recovered != 0 {
		i = encodeVarintDatum(dAtA, i, uint64(m.Recovered))
		i--
		dAtA[i] = 0x28
	}
	if m.Failed != 0 {
		i = encodeVarintDatum(dAtA, i, uint64(m.Failed))
		i--
		dAtA[i] = 0x20
	}
	if m.Skipped != 0 {
		i = encodeVarintDatum(dAtA, i, uint64(m.Skipped))
		i--
		dAtA[i] = 0x18
	}
	if m.Processed != 0 {
		i = encodeVarintDatum(dAtA, i, uint64(m.Processed))
		i--
		dAtA[i] = 0x10
	}
	if m.ProcessStats != nil {
		{
			size, err := m.ProcessStats.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDatum(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDatum(dAtA []byte, offset int, v uint64) int {
	offset -= sovDatum(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Meta) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.JobID)
	if l > 0 {
		n += 1 + l + sovDatum(uint64(l))
	}
	if len(m.Inputs) > 0 {
		for _, e := range m.Inputs {
			l = e.Size()
			n += 1 + l + sovDatum(uint64(l))
		}
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovDatum(uint64(l))
	}
	if m.State != 0 {
		n += 1 + sovDatum(uint64(m.State))
	}
	l = len(m.Reason)
	if l > 0 {
		n += 1 + l + sovDatum(uint64(l))
	}
	if m.Stats != nil {
		l = m.Stats.Size()
		n += 1 + l + sovDatum(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Stats) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProcessStats != nil {
		l = m.ProcessStats.Size()
		n += 1 + l + sovDatum(uint64(l))
	}
	if m.Processed != 0 {
		n += 1 + sovDatum(uint64(m.Processed))
	}
	if m.Skipped != 0 {
		n += 1 + sovDatum(uint64(m.Skipped))
	}
	if m.Failed != 0 {
		n += 1 + sovDatum(uint64(m.Failed))
	}
	if m.Recovered != 0 {
		n += 1 + sovDatum(uint64(m.Recovered))
	}
	l = len(m.FailedID)
	if l > 0 {
		n += 1 + l + sovDatum(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDatum(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDatum(x uint64) (n int) {
	return sovDatum(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Meta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDatum
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Meta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Meta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatum
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDatum
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDatum
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JobID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inputs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatum
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDatum
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDatum
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Inputs = append(m.Inputs, &common.InputV2{})
			if err := m.Inputs[len(m.Inputs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatum
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDatum
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDatum
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatum
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= State(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatum
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDatum
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDatum
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatum
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDatum
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDatum
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Stats == nil {
				m.Stats = &pps.ProcessStats{}
			}
			if err := m.Stats.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDatum(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDatum
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDatum
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Stats) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDatum
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Stats: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Stats: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProcessStats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatum
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDatum
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDatum
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ProcessStats == nil {
				m.ProcessStats = &pps.ProcessStats{}
			}
			if err := m.ProcessStats.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Processed", wireType)
			}
			m.Processed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatum
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Processed |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Skipped", wireType)
			}
			m.Skipped = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatum
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Skipped |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Failed", wireType)
			}
			m.Failed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatum
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Failed |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recovered", wireType)
			}
			m.Recovered = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatum
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Recovered |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FailedID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatum
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDatum
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDatum
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FailedID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDatum(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDatum
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDatum
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDatum(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDatum
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDatum
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDatum
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthDatum
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDatum
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDatum
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDatum        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDatum          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDatum = fmt.Errorf("proto: unexpected end of group")
)
