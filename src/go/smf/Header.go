// automatically generated by the FlatBuffers compiler, do not modify

package smf

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// \brief: header parsed by rpc engine
/// must be sizeof()'able
/// that is, must be a struct in fbs language
///
/// layout
/// [ 8bits(compression) + 8bits(bitflags) + 16bits(session) + 32bits(size) + 32bits(checksum) + 32bits(meta) ]
/// total = 128bits == 16bytes
///
type Header struct {
	_tab flatbuffers.Struct
}

func (rcv *Header) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Header) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *Header) Compression() int8 {
	return rcv._tab.GetInt8(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
func (rcv *Header) MutateCompression(n int8) bool {
	return rcv._tab.MutateInt8(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

func (rcv *Header) Bitflags() int8 {
	return rcv._tab.GetInt8(rcv._tab.Pos + flatbuffers.UOffsetT(1))
}
func (rcv *Header) MutateBitflags(n int8) bool {
	return rcv._tab.MutateInt8(rcv._tab.Pos+flatbuffers.UOffsetT(1), n)
}

/// 16 bits for storing the actual session id.
/// used for streaming client and slot allocation
func (rcv *Header) Session() uint16 {
	return rcv._tab.GetUint16(rcv._tab.Pos + flatbuffers.UOffsetT(2))
}

/// 16 bits for storing the actual session id.
/// used for streaming client and slot allocation
func (rcv *Header) MutateSession(n uint16) bool {
	return rcv._tab.MutateUint16(rcv._tab.Pos+flatbuffers.UOffsetT(2), n)
}

/// size of the next payload
func (rcv *Header) Size() uint32 {
	return rcv._tab.GetUint32(rcv._tab.Pos + flatbuffers.UOffsetT(4))
}

/// size of the next payload
func (rcv *Header) MutateSize(n uint32) bool {
	return rcv._tab.MutateUint32(rcv._tab.Pos+flatbuffers.UOffsetT(4), n)
}

/// currently we use xxhash32
func (rcv *Header) Checksum() uint32 {
	return rcv._tab.GetUint32(rcv._tab.Pos + flatbuffers.UOffsetT(8))
}

/// currently we use xxhash32
func (rcv *Header) MutateChecksum(n uint32) bool {
	return rcv._tab.MutateUint32(rcv._tab.Pos+flatbuffers.UOffsetT(8), n)
}

/// \brief used for sending and receiving, read carefully.
///
/// Receiving:
///
/// Uses the same as HTTP status - on the receiving end
/// We don't want to pay the cost of parsing a header
/// On every response as does HTTP. std::to_string and std::stol()
/// are needlesly expensive
///
/// Sending:
///
/// Used with the xor hash of Service::ID() ^ Service::Method::ID()
/// This is how the server multiplexer figures out what function pointer
/// to call
///
func (rcv *Header) Meta() uint32 {
	return rcv._tab.GetUint32(rcv._tab.Pos + flatbuffers.UOffsetT(12))
}

/// \brief used for sending and receiving, read carefully.
///
/// Receiving:
///
/// Uses the same as HTTP status - on the receiving end
/// We don't want to pay the cost of parsing a header
/// On every response as does HTTP. std::to_string and std::stol()
/// are needlesly expensive
///
/// Sending:
///
/// Used with the xor hash of Service::ID() ^ Service::Method::ID()
/// This is how the server multiplexer figures out what function pointer
/// to call
///
func (rcv *Header) MutateMeta(n uint32) bool {
	return rcv._tab.MutateUint32(rcv._tab.Pos+flatbuffers.UOffsetT(12), n)
}

func CreateHeader(builder *flatbuffers.Builder, compression int8, bitflags int8, session uint16, size uint32, checksum uint32, meta uint32) flatbuffers.UOffsetT {
	builder.Prep(4, 16)
	builder.PrependUint32(meta)
	builder.PrependUint32(checksum)
	builder.PrependUint32(size)
	builder.PrependUint16(session)
	builder.PrependInt8(bitflags)
	builder.PrependInt8(compression)
	return builder.Offset()
}
