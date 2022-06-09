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

// BACnetConstructedDataCredentialDataInputUpdateTime is the data-structure of this message
type BACnetConstructedDataCredentialDataInputUpdateTime struct {
	*BACnetConstructedData
	UpdateTime *BACnetTimeStamp

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataCredentialDataInputUpdateTime is the corresponding interface of BACnetConstructedDataCredentialDataInputUpdateTime
type IBACnetConstructedDataCredentialDataInputUpdateTime interface {
	IBACnetConstructedData
	// GetUpdateTime returns UpdateTime (property field)
	GetUpdateTime() *BACnetTimeStamp
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

func (m *BACnetConstructedDataCredentialDataInputUpdateTime) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_CREDENTIAL_DATA_INPUT
}

func (m *BACnetConstructedDataCredentialDataInputUpdateTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_UPDATE_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataCredentialDataInputUpdateTime) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataCredentialDataInputUpdateTime) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataCredentialDataInputUpdateTime) GetUpdateTime() *BACnetTimeStamp {
	return m.UpdateTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCredentialDataInputUpdateTime factory function for BACnetConstructedDataCredentialDataInputUpdateTime
func NewBACnetConstructedDataCredentialDataInputUpdateTime(updateTime *BACnetTimeStamp, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataCredentialDataInputUpdateTime {
	_result := &BACnetConstructedDataCredentialDataInputUpdateTime{
		UpdateTime:            updateTime,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataCredentialDataInputUpdateTime(structType interface{}) *BACnetConstructedDataCredentialDataInputUpdateTime {
	if casted, ok := structType.(BACnetConstructedDataCredentialDataInputUpdateTime); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCredentialDataInputUpdateTime); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataCredentialDataInputUpdateTime(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataCredentialDataInputUpdateTime(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataCredentialDataInputUpdateTime) GetTypeName() string {
	return "BACnetConstructedDataCredentialDataInputUpdateTime"
}

func (m *BACnetConstructedDataCredentialDataInputUpdateTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataCredentialDataInputUpdateTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (updateTime)
	lengthInBits += m.UpdateTime.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataCredentialDataInputUpdateTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataCredentialDataInputUpdateTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataCredentialDataInputUpdateTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCredentialDataInputUpdateTime"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (updateTime)
	if pullErr := readBuffer.PullContext("updateTime"); pullErr != nil {
		return nil, pullErr
	}
	_updateTime, _updateTimeErr := BACnetTimeStampParse(readBuffer)
	if _updateTimeErr != nil {
		return nil, errors.Wrap(_updateTimeErr, "Error parsing 'updateTime' field")
	}
	updateTime := CastBACnetTimeStamp(_updateTime)
	if closeErr := readBuffer.CloseContext("updateTime"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCredentialDataInputUpdateTime"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataCredentialDataInputUpdateTime{
		UpdateTime:            CastBACnetTimeStamp(updateTime),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataCredentialDataInputUpdateTime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCredentialDataInputUpdateTime"); pushErr != nil {
			return pushErr
		}

		// Simple Field (updateTime)
		if pushErr := writeBuffer.PushContext("updateTime"); pushErr != nil {
			return pushErr
		}
		_updateTimeErr := m.UpdateTime.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("updateTime"); popErr != nil {
			return popErr
		}
		if _updateTimeErr != nil {
			return errors.Wrap(_updateTimeErr, "Error serializing 'updateTime' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCredentialDataInputUpdateTime"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataCredentialDataInputUpdateTime) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}