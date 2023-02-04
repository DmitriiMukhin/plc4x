/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"context"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetVendorIdTagged is the corresponding interface of BACnetVendorIdTagged
type BACnetVendorIdTagged interface {
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetVendorId
	// GetUnknownId returns UnknownId (property field)
	GetUnknownId() uint32
	// GetIsUnknownId returns IsUnknownId (virtual field)
	GetIsUnknownId() bool
}

// BACnetVendorIdTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetVendorIdTagged.
// This is useful for switch cases.
type BACnetVendorIdTaggedExactly interface {
	BACnetVendorIdTagged
	isBACnetVendorIdTagged() bool
}

// _BACnetVendorIdTagged is the data-structure of this message
type _BACnetVendorIdTagged struct {
	Header    BACnetTagHeader
	Value     BACnetVendorId
	UnknownId uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetVendorIdTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetVendorIdTagged) GetValue() BACnetVendorId {
	return m.Value
}

func (m *_BACnetVendorIdTagged) GetUnknownId() uint32 {
	return m.UnknownId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetVendorIdTagged) GetIsUnknownId() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetVendorId_UNKNOWN_VENDOR)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetVendorIdTagged factory function for _BACnetVendorIdTagged
func NewBACnetVendorIdTagged(header BACnetTagHeader, value BACnetVendorId, unknownId uint32, tagNumber uint8, tagClass TagClass) *_BACnetVendorIdTagged {
	return &_BACnetVendorIdTagged{Header: header, Value: value, UnknownId: unknownId, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetVendorIdTagged(structType interface{}) BACnetVendorIdTagged {
	if casted, ok := structType.(BACnetVendorIdTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetVendorIdTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetVendorIdTagged) GetTypeName() string {
	return "BACnetVendorIdTagged"
}

func (m *_BACnetVendorIdTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(utils.InlineIf(m.GetIsUnknownId(), func() interface{} { return int32(int32(0)) }, func() interface{} { return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8)))) }).(int32))

	// A virtual field doesn't have any in- or output.

	// Manual Field (unknownId)
	lengthInBits += uint16(utils.InlineIf(m.GetIsUnknownId(), func() interface{} { return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8)))) }, func() interface{} { return int32(int32(0)) }).(int32))

	return lengthInBits
}

func (m *_BACnetVendorIdTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetVendorIdTaggedParse(theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetVendorIdTagged, error) {
	return BACnetVendorIdTaggedParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetVendorIdTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetVendorIdTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetVendorIdTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetVendorIdTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
	_header, _headerErr := BACnetTagHeaderParseWithBuffer(ctx, readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field of BACnetVendorIdTagged")
	}
	header := _header.(BACnetTagHeader)
	if closeErr := readBuffer.CloseContext("header"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for header")
	}

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{"tag class doesn't match"})
	}

	// Validation
	if !(bool((bool(bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS))) && bool(bool((header.GetActualTagNumber()) == (2))))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{"tagnumber doesn't match"})
	}

	// Manual Field (value)
	_value, _valueErr := ReadEnumGeneric(readBuffer, header.GetActualLength(), BACnetVendorId_UNKNOWN_VENDOR)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of BACnetVendorIdTagged")
	}
	var value BACnetVendorId
	if _value != nil {
		value = _value.(BACnetVendorId)
	}

	// Virtual field
	_isUnknownId := bool((value) == (BACnetVendorId_UNKNOWN_VENDOR))
	isUnknownId := bool(_isUnknownId)
	_ = isUnknownId

	// Manual Field (unknownId)
	_unknownId, _unknownIdErr := ReadProprietaryEnumGeneric(readBuffer, header.GetActualLength(), isUnknownId)
	if _unknownIdErr != nil {
		return nil, errors.Wrap(_unknownIdErr, "Error parsing 'unknownId' field of BACnetVendorIdTagged")
	}
	var unknownId uint32
	if _unknownId != nil {
		unknownId = _unknownId.(uint32)
	}

	if closeErr := readBuffer.CloseContext("BACnetVendorIdTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetVendorIdTagged")
	}

	// Create the instance
	return &_BACnetVendorIdTagged{
		TagNumber: tagNumber,
		TagClass:  tagClass,
		Header:    header,
		Value:     value,
		UnknownId: unknownId,
	}, nil
}

func (m *_BACnetVendorIdTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetVendorIdTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetVendorIdTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetVendorIdTagged")
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for header")
	}
	_headerErr := writeBuffer.WriteSerializable(ctx, m.GetHeader())
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for header")
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
	}

	// Manual Field (value)
	_valueErr := WriteEnumGeneric(writeBuffer, m.GetValue())
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}
	// Virtual field
	if _isUnknownIdErr := writeBuffer.WriteVirtual(ctx, "isUnknownId", m.GetIsUnknownId()); _isUnknownIdErr != nil {
		return errors.Wrap(_isUnknownIdErr, "Error serializing 'isUnknownId' field")
	}

	// Manual Field (unknownId)
	_unknownIdErr := WriteProprietaryEnumGeneric(writeBuffer, m.GetUnknownId(), m.GetIsUnknownId())
	if _unknownIdErr != nil {
		return errors.Wrap(_unknownIdErr, "Error serializing 'unknownId' field")
	}

	if popErr := writeBuffer.PopContext("BACnetVendorIdTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetVendorIdTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetVendorIdTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetVendorIdTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetVendorIdTagged) isBACnetVendorIdTagged() bool {
	return true
}

func (m *_BACnetVendorIdTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
