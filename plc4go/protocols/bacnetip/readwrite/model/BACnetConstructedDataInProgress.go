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

// BACnetConstructedDataInProgress is the corresponding interface of BACnetConstructedDataInProgress
type BACnetConstructedDataInProgress interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetInProgress returns InProgress (property field)
	GetInProgress() BACnetLightingInProgressTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetLightingInProgressTagged
}

// BACnetConstructedDataInProgressExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataInProgress.
// This is useful for switch cases.
type BACnetConstructedDataInProgressExactly interface {
	BACnetConstructedDataInProgress
	isBACnetConstructedDataInProgress() bool
}

// _BACnetConstructedDataInProgress is the data-structure of this message
type _BACnetConstructedDataInProgress struct {
	*_BACnetConstructedData
	InProgress BACnetLightingInProgressTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataInProgress) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataInProgress) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IN_PROGRESS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataInProgress) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataInProgress) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataInProgress) GetInProgress() BACnetLightingInProgressTagged {
	return m.InProgress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataInProgress) GetActualValue() BACnetLightingInProgressTagged {
	return CastBACnetLightingInProgressTagged(m.GetInProgress())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataInProgress factory function for _BACnetConstructedDataInProgress
func NewBACnetConstructedDataInProgress(inProgress BACnetLightingInProgressTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataInProgress {
	_result := &_BACnetConstructedDataInProgress{
		InProgress:             inProgress,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataInProgress(structType interface{}) BACnetConstructedDataInProgress {
	if casted, ok := structType.(BACnetConstructedDataInProgress); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataInProgress); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataInProgress) GetTypeName() string {
	return "BACnetConstructedDataInProgress"
}

func (m *_BACnetConstructedDataInProgress) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataInProgress) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (inProgress)
	lengthInBits += m.InProgress.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataInProgress) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataInProgressParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataInProgress, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataInProgress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataInProgress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (inProgress)
	if pullErr := readBuffer.PullContext("inProgress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for inProgress")
	}
	_inProgress, _inProgressErr := BACnetLightingInProgressTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _inProgressErr != nil {
		return nil, errors.Wrap(_inProgressErr, "Error parsing 'inProgress' field")
	}
	inProgress := _inProgress.(BACnetLightingInProgressTagged)
	if closeErr := readBuffer.CloseContext("inProgress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for inProgress")
	}

	// Virtual field
	_actualValue := inProgress
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataInProgress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataInProgress")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataInProgress{
		InProgress:             inProgress,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataInProgress) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataInProgress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataInProgress")
		}

		// Simple Field (inProgress)
		if pushErr := writeBuffer.PushContext("inProgress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for inProgress")
		}
		_inProgressErr := writeBuffer.WriteSerializable(m.GetInProgress())
		if popErr := writeBuffer.PopContext("inProgress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for inProgress")
		}
		if _inProgressErr != nil {
			return errors.Wrap(_inProgressErr, "Error serializing 'inProgress' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataInProgress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataInProgress")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataInProgress) isBACnetConstructedDataInProgress() bool {
	return true
}

func (m *_BACnetConstructedDataInProgress) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
