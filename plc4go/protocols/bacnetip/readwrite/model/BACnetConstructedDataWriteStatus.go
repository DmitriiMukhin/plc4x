/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataWriteStatus is the data-structure of this message
type BACnetConstructedDataWriteStatus struct {
	*BACnetConstructedData
	WriteStatus *BACnetWriteStatusTagged

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataWriteStatus is the corresponding interface of BACnetConstructedDataWriteStatus
type IBACnetConstructedDataWriteStatus interface {
	IBACnetConstructedData
	// GetWriteStatus returns WriteStatus (property field)
	GetWriteStatus() *BACnetWriteStatusTagged
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetConstructedDataWriteStatus) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataWriteStatus) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_WRITE_STATUS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataWriteStatus) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataWriteStatus) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataWriteStatus) GetWriteStatus() *BACnetWriteStatusTagged {
	return m.WriteStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataWriteStatus factory function for BACnetConstructedDataWriteStatus
func NewBACnetConstructedDataWriteStatus(writeStatus *BACnetWriteStatusTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataWriteStatus {
	_result := &BACnetConstructedDataWriteStatus{
		WriteStatus:           writeStatus,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataWriteStatus(structType interface{}) *BACnetConstructedDataWriteStatus {
	if casted, ok := structType.(BACnetConstructedDataWriteStatus); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataWriteStatus); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataWriteStatus(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataWriteStatus(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataWriteStatus) GetTypeName() string {
	return "BACnetConstructedDataWriteStatus"
}

func (m *BACnetConstructedDataWriteStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataWriteStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (writeStatus)
	lengthInBits += m.WriteStatus.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataWriteStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataWriteStatusParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataWriteStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataWriteStatus"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (writeStatus)
	if pullErr := readBuffer.PullContext("writeStatus"); pullErr != nil {
		return nil, pullErr
	}
	_writeStatus, _writeStatusErr := BACnetWriteStatusTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _writeStatusErr != nil {
		return nil, errors.Wrap(_writeStatusErr, "Error parsing 'writeStatus' field")
	}
	writeStatus := CastBACnetWriteStatusTagged(_writeStatus)
	if closeErr := readBuffer.CloseContext("writeStatus"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataWriteStatus"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataWriteStatus{
		WriteStatus:           CastBACnetWriteStatusTagged(writeStatus),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataWriteStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataWriteStatus"); pushErr != nil {
			return pushErr
		}

		// Simple Field (writeStatus)
		if pushErr := writeBuffer.PushContext("writeStatus"); pushErr != nil {
			return pushErr
		}
		_writeStatusErr := m.WriteStatus.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("writeStatus"); popErr != nil {
			return popErr
		}
		if _writeStatusErr != nil {
			return errors.Wrap(_writeStatusErr, "Error serializing 'writeStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataWriteStatus"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataWriteStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}