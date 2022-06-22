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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataProcessIdentifierFilter is the corresponding interface of BACnetConstructedDataProcessIdentifierFilter
type BACnetConstructedDataProcessIdentifierFilter interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetProcessIdentifierFilter returns ProcessIdentifierFilter (property field)
	GetProcessIdentifierFilter() BACnetProcessIdSelection
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetProcessIdSelection
}

// BACnetConstructedDataProcessIdentifierFilterExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataProcessIdentifierFilter.
// This is useful for switch cases.
type BACnetConstructedDataProcessIdentifierFilterExactly interface {
	BACnetConstructedDataProcessIdentifierFilter
	isBACnetConstructedDataProcessIdentifierFilter() bool
}

// _BACnetConstructedDataProcessIdentifierFilter is the data-structure of this message
type _BACnetConstructedDataProcessIdentifierFilter struct {
	*_BACnetConstructedData
	ProcessIdentifierFilter BACnetProcessIdSelection
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataProcessIdentifierFilter) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataProcessIdentifierFilter) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PROCESS_IDENTIFIER_FILTER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataProcessIdentifierFilter) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataProcessIdentifierFilter) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataProcessIdentifierFilter) GetProcessIdentifierFilter() BACnetProcessIdSelection {
	return m.ProcessIdentifierFilter
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataProcessIdentifierFilter) GetActualValue() BACnetProcessIdSelection {
	return CastBACnetProcessIdSelection(m.GetProcessIdentifierFilter())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataProcessIdentifierFilter factory function for _BACnetConstructedDataProcessIdentifierFilter
func NewBACnetConstructedDataProcessIdentifierFilter(processIdentifierFilter BACnetProcessIdSelection, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataProcessIdentifierFilter {
	_result := &_BACnetConstructedDataProcessIdentifierFilter{
		ProcessIdentifierFilter: processIdentifierFilter,
		_BACnetConstructedData:  NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataProcessIdentifierFilter(structType interface{}) BACnetConstructedDataProcessIdentifierFilter {
	if casted, ok := structType.(BACnetConstructedDataProcessIdentifierFilter); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataProcessIdentifierFilter); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataProcessIdentifierFilter) GetTypeName() string {
	return "BACnetConstructedDataProcessIdentifierFilter"
}

func (m *_BACnetConstructedDataProcessIdentifierFilter) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataProcessIdentifierFilter) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (processIdentifierFilter)
	lengthInBits += m.ProcessIdentifierFilter.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataProcessIdentifierFilter) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataProcessIdentifierFilterParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataProcessIdentifierFilter, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataProcessIdentifierFilter"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataProcessIdentifierFilter")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (processIdentifierFilter)
	if pullErr := readBuffer.PullContext("processIdentifierFilter"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for processIdentifierFilter")
	}
	_processIdentifierFilter, _processIdentifierFilterErr := BACnetProcessIdSelectionParse(readBuffer)
	if _processIdentifierFilterErr != nil {
		return nil, errors.Wrap(_processIdentifierFilterErr, "Error parsing 'processIdentifierFilter' field")
	}
	processIdentifierFilter := _processIdentifierFilter.(BACnetProcessIdSelection)
	if closeErr := readBuffer.CloseContext("processIdentifierFilter"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for processIdentifierFilter")
	}

	// Virtual field
	_actualValue := processIdentifierFilter
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataProcessIdentifierFilter"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataProcessIdentifierFilter")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataProcessIdentifierFilter{
		ProcessIdentifierFilter: processIdentifierFilter,
		_BACnetConstructedData:  &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataProcessIdentifierFilter) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataProcessIdentifierFilter"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataProcessIdentifierFilter")
		}

		// Simple Field (processIdentifierFilter)
		if pushErr := writeBuffer.PushContext("processIdentifierFilter"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for processIdentifierFilter")
		}
		_processIdentifierFilterErr := writeBuffer.WriteSerializable(m.GetProcessIdentifierFilter())
		if popErr := writeBuffer.PopContext("processIdentifierFilter"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for processIdentifierFilter")
		}
		if _processIdentifierFilterErr != nil {
			return errors.Wrap(_processIdentifierFilterErr, "Error serializing 'processIdentifierFilter' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataProcessIdentifierFilter"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataProcessIdentifierFilter")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataProcessIdentifierFilter) isBACnetConstructedDataProcessIdentifierFilter() bool {
	return true
}

func (m *_BACnetConstructedDataProcessIdentifierFilter) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
