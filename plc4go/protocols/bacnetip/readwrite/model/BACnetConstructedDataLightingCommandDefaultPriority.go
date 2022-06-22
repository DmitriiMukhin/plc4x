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

// BACnetConstructedDataLightingCommandDefaultPriority is the corresponding interface of BACnetConstructedDataLightingCommandDefaultPriority
type BACnetConstructedDataLightingCommandDefaultPriority interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLightingCommandDefaultPriority returns LightingCommandDefaultPriority (property field)
	GetLightingCommandDefaultPriority() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataLightingCommandDefaultPriorityExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLightingCommandDefaultPriority.
// This is useful for switch cases.
type BACnetConstructedDataLightingCommandDefaultPriorityExactly interface {
	BACnetConstructedDataLightingCommandDefaultPriority
	isBACnetConstructedDataLightingCommandDefaultPriority() bool
}

// _BACnetConstructedDataLightingCommandDefaultPriority is the data-structure of this message
type _BACnetConstructedDataLightingCommandDefaultPriority struct {
	*_BACnetConstructedData
	LightingCommandDefaultPriority BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LIGHTING_COMMAND_DEFAULT_PRIORITY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) GetLightingCommandDefaultPriority() BACnetApplicationTagUnsignedInteger {
	return m.LightingCommandDefaultPriority
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetLightingCommandDefaultPriority())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLightingCommandDefaultPriority factory function for _BACnetConstructedDataLightingCommandDefaultPriority
func NewBACnetConstructedDataLightingCommandDefaultPriority(lightingCommandDefaultPriority BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLightingCommandDefaultPriority {
	_result := &_BACnetConstructedDataLightingCommandDefaultPriority{
		LightingCommandDefaultPriority: lightingCommandDefaultPriority,
		_BACnetConstructedData:         NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLightingCommandDefaultPriority(structType interface{}) BACnetConstructedDataLightingCommandDefaultPriority {
	if casted, ok := structType.(BACnetConstructedDataLightingCommandDefaultPriority); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLightingCommandDefaultPriority); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) GetTypeName() string {
	return "BACnetConstructedDataLightingCommandDefaultPriority"
}

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (lightingCommandDefaultPriority)
	lengthInBits += m.LightingCommandDefaultPriority.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLightingCommandDefaultPriorityParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLightingCommandDefaultPriority, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLightingCommandDefaultPriority"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLightingCommandDefaultPriority")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lightingCommandDefaultPriority)
	if pullErr := readBuffer.PullContext("lightingCommandDefaultPriority"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lightingCommandDefaultPriority")
	}
	_lightingCommandDefaultPriority, _lightingCommandDefaultPriorityErr := BACnetApplicationTagParse(readBuffer)
	if _lightingCommandDefaultPriorityErr != nil {
		return nil, errors.Wrap(_lightingCommandDefaultPriorityErr, "Error parsing 'lightingCommandDefaultPriority' field")
	}
	lightingCommandDefaultPriority := _lightingCommandDefaultPriority.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("lightingCommandDefaultPriority"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lightingCommandDefaultPriority")
	}

	// Virtual field
	_actualValue := lightingCommandDefaultPriority
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLightingCommandDefaultPriority"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLightingCommandDefaultPriority")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLightingCommandDefaultPriority{
		LightingCommandDefaultPriority: lightingCommandDefaultPriority,
		_BACnetConstructedData:         &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLightingCommandDefaultPriority"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLightingCommandDefaultPriority")
		}

		// Simple Field (lightingCommandDefaultPriority)
		if pushErr := writeBuffer.PushContext("lightingCommandDefaultPriority"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lightingCommandDefaultPriority")
		}
		_lightingCommandDefaultPriorityErr := writeBuffer.WriteSerializable(m.GetLightingCommandDefaultPriority())
		if popErr := writeBuffer.PopContext("lightingCommandDefaultPriority"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lightingCommandDefaultPriority")
		}
		if _lightingCommandDefaultPriorityErr != nil {
			return errors.Wrap(_lightingCommandDefaultPriorityErr, "Error serializing 'lightingCommandDefaultPriority' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLightingCommandDefaultPriority"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLightingCommandDefaultPriority")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) isBACnetConstructedDataLightingCommandDefaultPriority() bool {
	return true
}

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
