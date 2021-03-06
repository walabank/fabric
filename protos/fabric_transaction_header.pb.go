// Code generated by protoc-gen-go.
// source: fabric_transaction_header.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Header_Type int32

const (
	Header_UNDEFINED Header_Type = 0
	Header_CHAINCODE Header_Type = 1
)

var Header_Type_name = map[int32]string{
	0: "UNDEFINED",
	1: "CHAINCODE",
}
var Header_Type_value = map[string]int32{
	"UNDEFINED": 0,
	"CHAINCODE": 1,
}

func (x Header_Type) String() string {
	return proto.EnumName(Header_Type_name, int32(x))
}
func (Header_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor13, []int{0, 0} }

// A Header contains fields that are common to all proposals and all
// transactions, no matter their type.  It can include also type-dependant
// fields by using the 'extensions' field.  This header is on purpose the same
// header for proposals (a request to do "something" on the ledger) and a
// transaction (the endorsed actions following from some request).
// Furthermore, a proposal, its endorsements and the resulting transaction are
// linked together by this message, as follows
// 1. a Proposal contains a Header
// 2. the hash of the Header of a proposal is included in the proposal response
//    generated by each endorser as a result of that proposal
// 3. a TransactionAction contains both i) the *same* Header (byte-by-byte) of
//    the corresponsing Proposal and ii) the hash of the Header in each of the
//    endorsed actions
type Header struct {
	// Version indicates message protocol version
	Version int32 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	// Timestamp is the local time when the message was created
	// by the sender
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	// Type of the transaction
	Type Header_Type `protobuf:"varint,3,opt,name=type,enum=protos.Header_Type" json:"type,omitempty"`
	// Creator of the header (and encapsulating message). This is usually a tcert
	// or ecert identifying the entity who submits the proposal/transaction.  The
	// creator identifies the signer of
	// 1. a proposal (if this is the header of a Proposal message)
	// 2. a transaction (if this is the header of a TransactionAction message)
	Creator []byte `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
	// Arbitrary number that may only be used once. This ensures the hash of
	// the proposal is unique and may be used in replay detection
	Nonce []byte `protobuf:"bytes,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// Identifier of the chain this header targets to
	ChainID []byte `protobuf:"bytes,6,opt,name=chainID,proto3" json:"chainID,omitempty"`
	// Extensions is used to include type-dependant fields
	Extensions []byte `protobuf:"bytes,7,opt,name=extensions,proto3" json:"extensions,omitempty"`
}

func (m *Header) Reset()                    { *m = Header{} }
func (m *Header) String() string            { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()               {}
func (*Header) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *Header) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func init() {
	proto.RegisterType((*Header)(nil), "protos.Header")
	proto.RegisterEnum("protos.Header_Type", Header_Type_name, Header_Type_value)
}

func init() { proto.RegisterFile("fabric_transaction_header.proto", fileDescriptor13) }

var fileDescriptor13 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x90, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xcd, 0x6c, 0x3b, 0xf6, 0xfc, 0xc1, 0x88, 0x1e, 0xc2, 0x0e, 0xae, 0x0c, 0xc1, 0x9e,
	0x32, 0x98, 0x17, 0xaf, 0xb2, 0x56, 0xd6, 0x4b, 0x85, 0x32, 0xcf, 0x23, 0xad, 0x6f, 0x5b, 0xc1,
	0x25, 0x25, 0x89, 0xe2, 0xfe, 0x10, 0xff, 0x5f, 0x69, 0x62, 0xd5, 0x53, 0xf8, 0xbc, 0x7c, 0xde,
	0xf7, 0x1b, 0x02, 0xd3, 0xad, 0xa8, 0x74, 0x53, 0x6f, 0xac, 0x16, 0xd2, 0x88, 0xda, 0x36, 0x4a,
	0x6e, 0xf6, 0x28, 0x5e, 0x51, 0xf3, 0x56, 0x2b, 0xab, 0x68, 0xe4, 0x0e, 0x33, 0x99, 0xee, 0x94,
	0xda, 0xbd, 0xe1, 0xdc, 0x61, 0xf5, 0xbe, 0x9d, 0xdb, 0xe6, 0x80, 0xc6, 0x8a, 0x43, 0xeb, 0xc5,
	0xd9, 0xd7, 0x00, 0xa2, 0x95, 0xdb, 0xa4, 0x0c, 0x86, 0x1f, 0xa8, 0x4d, 0xa3, 0x24, 0x23, 0x31,
	0x49, 0xc2, 0xb2, 0x47, 0xfa, 0x00, 0xa3, 0xdf, 0x3d, 0x36, 0x88, 0x49, 0x72, 0xb6, 0x98, 0x70,
	0x9f, 0xcc, 0xfb, 0x64, 0xbe, 0xee, 0x8d, 0xf2, 0x4f, 0xa6, 0x77, 0x10, 0xd8, 0x63, 0x8b, 0xec,
	0x34, 0x26, 0xc9, 0xe5, 0xe2, 0xca, 0xdb, 0x86, 0xfb, 0x46, 0xbe, 0x3e, 0xb6, 0x58, 0x3a, 0xa1,
	0x2b, 0xaf, 0x35, 0x0a, 0xab, 0x34, 0x0b, 0x62, 0x92, 0x9c, 0x97, 0x3d, 0xd2, 0x6b, 0x08, 0xa5,
	0x92, 0x35, 0xb2, 0xd0, 0xcd, 0x3d, 0x38, 0x7f, 0x2f, 0x1a, 0x99, 0xa7, 0x2c, 0xfa, 0xf1, 0x3d,
	0xd2, 0x1b, 0x00, 0xfc, 0xb4, 0x28, 0xbb, 0x97, 0x1b, 0x36, 0x74, 0x97, 0xff, 0x26, 0xb3, 0x5b,
	0x08, 0xba, 0x5e, 0x7a, 0x01, 0xa3, 0x97, 0x22, 0xcd, 0x9e, 0xf2, 0x22, 0x4b, 0xc7, 0x27, 0x1d,
	0x2e, 0x57, 0x8f, 0x79, 0xb1, 0x7c, 0x4e, 0xb3, 0x31, 0xa9, 0xfc, 0x07, 0xde, 0x7f, 0x07, 0x00,
	0x00, 0xff, 0xff, 0x36, 0x2e, 0x02, 0xb6, 0x6a, 0x01, 0x00, 0x00,
}
