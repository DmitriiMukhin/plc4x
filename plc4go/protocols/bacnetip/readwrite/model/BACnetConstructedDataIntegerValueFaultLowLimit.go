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

// BACnetConstructedDataIntegerValueFaultLowLimit is the data-structure of this message
type BACnetConstructedDataIntegerValueFaultLowLimit struct {
	*BACnetConstructedData
	FaultLowLimit *BACnetApplicationTagSignedInteger

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataIntegerValueFaultLowLimit is the corresponding interface of BACnetConstructedDataIntegerValueFaultLowLimit
type IBACnetConstructedDataIntegerValueFaultLowLimit interface {
	IBACnetConstructedData
	// GetFaultLowLimit returns FaultLowLimit (property field)
	GetFaultLowLimit() *BACnetApplicationTagSignedInteger
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

func (m *BACnetConstructedDataIntegerValueFaultLowLimit) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_INTEGER_VALUE
}

func (m *BACnetConstructedDataIntegerValueFaultLowLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAULT_LOW_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataIntegerValueFaultLowLimit) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataIntegerValueFaultLowLimit) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataIntegerValueFaultLowLimit) GetFaultLowLimit() *BACnetApplicationTagSignedInteger {
	return m.FaultLowLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataIntegerValueFaultLowLimit factory function for BACnetConstructedDataIntegerValueFaultLowLimit
func NewBACnetConstructedDataIntegerValueFaultLowLimit(faultLowLimit *BACnetApplicationTagSignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataIntegerValueFaultLowLimit {
	_result := &BACnetConstructedDataIntegerValueFaultLowLimit{
		FaultLowLimit:         faultLowLimit,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataIntegerValueFaultLowLimit(structType interface{}) *BACnetConstructedDataIntegerValueFaultLowLimit {
	if casted, ok := structType.(BACnetConstructedDataIntegerValueFaultLowLimit); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIntegerValueFaultLowLimit); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataIntegerValueFaultLowLimit(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataIntegerValueFaultLowLimit(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataIntegerValueFaultLowLimit) GetTypeName() string {
	return "BACnetConstructedDataIntegerValueFaultLowLimit"
}

func (m *BACnetConstructedDataIntegerValueFaultLowLimit) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataIntegerValueFaultLowLimit) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (faultLowLimit)
	lengthInBits += m.FaultLowLimit.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataIntegerValueFaultLowLimit) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataIntegerValueFaultLowLimitParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataIntegerValueFaultLowLimit, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIntegerValueFaultLowLimit"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (faultLowLimit)
	if pullErr := readBuffer.PullContext("faultLowLimit"); pullErr != nil {
		return nil, pullErr
	}
	_faultLowLimit, _faultLowLimitErr := BACnetApplicationTagParse(readBuffer)
	if _faultLowLimitErr != nil {
		return nil, errors.Wrap(_faultLowLimitErr, "Error parsing 'faultLowLimit' field")
	}
	faultLowLimit := CastBACnetApplicationTagSignedInteger(_faultLowLimit)
	if closeErr := readBuffer.CloseContext("faultLowLimit"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIntegerValueFaultLowLimit"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataIntegerValueFaultLowLimit{
		FaultLowLimit:         CastBACnetApplicationTagSignedInteger(faultLowLimit),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataIntegerValueFaultLowLimit) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIntegerValueFaultLowLimit"); pushErr != nil {
			return pushErr
		}

		// Simple Field (faultLowLimit)
		if pushErr := writeBuffer.PushContext("faultLowLimit"); pushErr != nil {
			return pushErr
		}
		_faultLowLimitErr := m.FaultLowLimit.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("faultLowLimit"); popErr != nil {
			return popErr
		}
		if _faultLowLimitErr != nil {
			return errors.Wrap(_faultLowLimitErr, "Error serializing 'faultLowLimit' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIntegerValueFaultLowLimit"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataIntegerValueFaultLowLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}