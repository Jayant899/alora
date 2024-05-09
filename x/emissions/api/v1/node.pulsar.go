// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package emissionsv1

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_OffchainNode               protoreflect.MessageDescriptor
	fd_OffchainNode_lib_p2p_key   protoreflect.FieldDescriptor
	fd_OffchainNode_multi_address protoreflect.FieldDescriptor
	fd_OffchainNode_owner         protoreflect.FieldDescriptor
	fd_OffchainNode_node_address  protoreflect.FieldDescriptor
	fd_OffchainNode_node_id       protoreflect.FieldDescriptor
)

func init() {
	file_emissions_v1_node_proto_init()
	md_OffchainNode = File_emissions_v1_node_proto.Messages().ByName("OffchainNode")
	fd_OffchainNode_lib_p2p_key = md_OffchainNode.Fields().ByName("lib_p2p_key")
	fd_OffchainNode_multi_address = md_OffchainNode.Fields().ByName("multi_address")
	fd_OffchainNode_owner = md_OffchainNode.Fields().ByName("owner")
	fd_OffchainNode_node_address = md_OffchainNode.Fields().ByName("node_address")
	fd_OffchainNode_node_id = md_OffchainNode.Fields().ByName("node_id")
}

var _ protoreflect.Message = (*fastReflection_OffchainNode)(nil)

type fastReflection_OffchainNode OffchainNode

func (x *OffchainNode) ProtoReflect() protoreflect.Message {
	return (*fastReflection_OffchainNode)(x)
}

func (x *OffchainNode) slowProtoReflect() protoreflect.Message {
	mi := &file_emissions_v1_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_OffchainNode_messageType fastReflection_OffchainNode_messageType
var _ protoreflect.MessageType = fastReflection_OffchainNode_messageType{}

type fastReflection_OffchainNode_messageType struct{}

func (x fastReflection_OffchainNode_messageType) Zero() protoreflect.Message {
	return (*fastReflection_OffchainNode)(nil)
}
func (x fastReflection_OffchainNode_messageType) New() protoreflect.Message {
	return new(fastReflection_OffchainNode)
}
func (x fastReflection_OffchainNode_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_OffchainNode
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_OffchainNode) Descriptor() protoreflect.MessageDescriptor {
	return md_OffchainNode
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_OffchainNode) Type() protoreflect.MessageType {
	return _fastReflection_OffchainNode_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_OffchainNode) New() protoreflect.Message {
	return new(fastReflection_OffchainNode)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_OffchainNode) Interface() protoreflect.ProtoMessage {
	return (*OffchainNode)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_OffchainNode) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.LibP2PKey != "" {
		value := protoreflect.ValueOfString(x.LibP2PKey)
		if !f(fd_OffchainNode_lib_p2p_key, value) {
			return
		}
	}
	if x.MultiAddress != "" {
		value := protoreflect.ValueOfString(x.MultiAddress)
		if !f(fd_OffchainNode_multi_address, value) {
			return
		}
	}
	if x.Owner != "" {
		value := protoreflect.ValueOfString(x.Owner)
		if !f(fd_OffchainNode_owner, value) {
			return
		}
	}
	if x.NodeAddress != "" {
		value := protoreflect.ValueOfString(x.NodeAddress)
		if !f(fd_OffchainNode_node_address, value) {
			return
		}
	}
	if x.NodeId != "" {
		value := protoreflect.ValueOfString(x.NodeId)
		if !f(fd_OffchainNode_node_id, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_OffchainNode) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "emissions.v1.OffchainNode.lib_p2p_key":
		return x.LibP2PKey != ""
	case "emissions.v1.OffchainNode.multi_address":
		return x.MultiAddress != ""
	case "emissions.v1.OffchainNode.owner":
		return x.Owner != ""
	case "emissions.v1.OffchainNode.node_address":
		return x.NodeAddress != ""
	case "emissions.v1.OffchainNode.node_id":
		return x.NodeId != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: emissions.v1.OffchainNode"))
		}
		panic(fmt.Errorf("message emissions.v1.OffchainNode does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_OffchainNode) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "emissions.v1.OffchainNode.lib_p2p_key":
		x.LibP2PKey = ""
	case "emissions.v1.OffchainNode.multi_address":
		x.MultiAddress = ""
	case "emissions.v1.OffchainNode.owner":
		x.Owner = ""
	case "emissions.v1.OffchainNode.node_address":
		x.NodeAddress = ""
	case "emissions.v1.OffchainNode.node_id":
		x.NodeId = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: emissions.v1.OffchainNode"))
		}
		panic(fmt.Errorf("message emissions.v1.OffchainNode does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_OffchainNode) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "emissions.v1.OffchainNode.lib_p2p_key":
		value := x.LibP2PKey
		return protoreflect.ValueOfString(value)
	case "emissions.v1.OffchainNode.multi_address":
		value := x.MultiAddress
		return protoreflect.ValueOfString(value)
	case "emissions.v1.OffchainNode.owner":
		value := x.Owner
		return protoreflect.ValueOfString(value)
	case "emissions.v1.OffchainNode.node_address":
		value := x.NodeAddress
		return protoreflect.ValueOfString(value)
	case "emissions.v1.OffchainNode.node_id":
		value := x.NodeId
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: emissions.v1.OffchainNode"))
		}
		panic(fmt.Errorf("message emissions.v1.OffchainNode does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_OffchainNode) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "emissions.v1.OffchainNode.lib_p2p_key":
		x.LibP2PKey = value.Interface().(string)
	case "emissions.v1.OffchainNode.multi_address":
		x.MultiAddress = value.Interface().(string)
	case "emissions.v1.OffchainNode.owner":
		x.Owner = value.Interface().(string)
	case "emissions.v1.OffchainNode.node_address":
		x.NodeAddress = value.Interface().(string)
	case "emissions.v1.OffchainNode.node_id":
		x.NodeId = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: emissions.v1.OffchainNode"))
		}
		panic(fmt.Errorf("message emissions.v1.OffchainNode does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_OffchainNode) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "emissions.v1.OffchainNode.lib_p2p_key":
		panic(fmt.Errorf("field lib_p2p_key of message emissions.v1.OffchainNode is not mutable"))
	case "emissions.v1.OffchainNode.multi_address":
		panic(fmt.Errorf("field multi_address of message emissions.v1.OffchainNode is not mutable"))
	case "emissions.v1.OffchainNode.owner":
		panic(fmt.Errorf("field owner of message emissions.v1.OffchainNode is not mutable"))
	case "emissions.v1.OffchainNode.node_address":
		panic(fmt.Errorf("field node_address of message emissions.v1.OffchainNode is not mutable"))
	case "emissions.v1.OffchainNode.node_id":
		panic(fmt.Errorf("field node_id of message emissions.v1.OffchainNode is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: emissions.v1.OffchainNode"))
		}
		panic(fmt.Errorf("message emissions.v1.OffchainNode does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_OffchainNode) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "emissions.v1.OffchainNode.lib_p2p_key":
		return protoreflect.ValueOfString("")
	case "emissions.v1.OffchainNode.multi_address":
		return protoreflect.ValueOfString("")
	case "emissions.v1.OffchainNode.owner":
		return protoreflect.ValueOfString("")
	case "emissions.v1.OffchainNode.node_address":
		return protoreflect.ValueOfString("")
	case "emissions.v1.OffchainNode.node_id":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: emissions.v1.OffchainNode"))
		}
		panic(fmt.Errorf("message emissions.v1.OffchainNode does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_OffchainNode) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in emissions.v1.OffchainNode", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_OffchainNode) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_OffchainNode) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_OffchainNode) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_OffchainNode) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*OffchainNode)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.LibP2PKey)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.MultiAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Owner)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.NodeAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.NodeId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*OffchainNode)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.NodeId) > 0 {
			i -= len(x.NodeId)
			copy(dAtA[i:], x.NodeId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.NodeId)))
			i--
			dAtA[i] = 0x2a
		}
		if len(x.NodeAddress) > 0 {
			i -= len(x.NodeAddress)
			copy(dAtA[i:], x.NodeAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.NodeAddress)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Owner) > 0 {
			i -= len(x.Owner)
			copy(dAtA[i:], x.Owner)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Owner)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.MultiAddress) > 0 {
			i -= len(x.MultiAddress)
			copy(dAtA[i:], x.MultiAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.MultiAddress)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.LibP2PKey) > 0 {
			i -= len(x.LibP2PKey)
			copy(dAtA[i:], x.LibP2PKey)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.LibP2PKey)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*OffchainNode)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: OffchainNode: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: OffchainNode: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field LibP2PKey", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.LibP2PKey = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MultiAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.MultiAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Owner = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field NodeAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.NodeAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field NodeId", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.NodeId = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: emissions/v1/node.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OffchainNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LibP2PKey    string `protobuf:"bytes,1,opt,name=lib_p2p_key,json=libP2pKey,proto3" json:"lib_p2p_key,omitempty"`        // LibP2P key of the node
	MultiAddress string `protobuf:"bytes,2,opt,name=multi_address,json=multiAddress,proto3" json:"multi_address,omitempty"` // Network address for accessing the node
	Owner        string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	NodeAddress  string `protobuf:"bytes,4,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty"`
	NodeId       string `protobuf:"bytes,5,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
}

func (x *OffchainNode) Reset() {
	*x = OffchainNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emissions_v1_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OffchainNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OffchainNode) ProtoMessage() {}

// Deprecated: Use OffchainNode.ProtoReflect.Descriptor instead.
func (*OffchainNode) Descriptor() ([]byte, []int) {
	return file_emissions_v1_node_proto_rawDescGZIP(), []int{0}
}

func (x *OffchainNode) GetLibP2PKey() string {
	if x != nil {
		return x.LibP2PKey
	}
	return ""
}

func (x *OffchainNode) GetMultiAddress() string {
	if x != nil {
		return x.MultiAddress
	}
	return ""
}

func (x *OffchainNode) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *OffchainNode) GetNodeAddress() string {
	if x != nil {
		return x.NodeAddress
	}
	return ""
}

func (x *OffchainNode) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

var File_emissions_v1_node_proto protoreflect.FileDescriptor

var file_emissions_v1_node_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x22, 0xa5, 0x01, 0x0a, 0x0c, 0x4f, 0x66, 0x66, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x6c, 0x69, 0x62, 0x5f,
	0x70, 0x32, 0x70, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c,
	0x69, 0x62, 0x50, 0x32, 0x70, 0x4b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x42,
	0xbf, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c,
	0x6c, 0x6f, 0x72, 0x61, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x61, 0x6c, 0x6c,
	0x6f, 0x72, 0x61, 0x2d, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x78, 0x2f, 0x65, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x45, 0x58, 0x58, 0xaa, 0x02, 0x0c, 0x45, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c, 0x45, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x45, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x0d, 0x45, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_emissions_v1_node_proto_rawDescOnce sync.Once
	file_emissions_v1_node_proto_rawDescData = file_emissions_v1_node_proto_rawDesc
)

func file_emissions_v1_node_proto_rawDescGZIP() []byte {
	file_emissions_v1_node_proto_rawDescOnce.Do(func() {
		file_emissions_v1_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_emissions_v1_node_proto_rawDescData)
	})
	return file_emissions_v1_node_proto_rawDescData
}

var file_emissions_v1_node_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_emissions_v1_node_proto_goTypes = []interface{}{
	(*OffchainNode)(nil), // 0: emissions.v1.OffchainNode
}
var file_emissions_v1_node_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_emissions_v1_node_proto_init() }
func file_emissions_v1_node_proto_init() {
	if File_emissions_v1_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_emissions_v1_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OffchainNode); i {
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
			RawDescriptor: file_emissions_v1_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_emissions_v1_node_proto_goTypes,
		DependencyIndexes: file_emissions_v1_node_proto_depIdxs,
		MessageInfos:      file_emissions_v1_node_proto_msgTypes,
	}.Build()
	File_emissions_v1_node_proto = out.File
	file_emissions_v1_node_proto_rawDesc = nil
	file_emissions_v1_node_proto_goTypes = nil
	file_emissions_v1_node_proto_depIdxs = nil
}