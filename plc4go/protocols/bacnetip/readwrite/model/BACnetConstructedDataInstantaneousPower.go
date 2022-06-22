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

// BACnetConstructedDataInstantaneousPower is the corresponding interface of BACnetConstructedDataInstantaneousPower
type BACnetConstructedDataInstantaneousPower interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetInstantaneousPower returns InstantaneousPower (property field)
	GetInstantaneousPower() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
}

// BACnetConstructedDataInstantaneousPowerExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataInstantaneousPower.
// This is useful for switch cases.
type BACnetConstructedDataInstantaneousPowerExactly interface {
	BACnetConstructedDataInstantaneousPower
	isBACnetConstructedDataInstantaneousPower() bool
}

// _BACnetConstructedDataInstantaneousPower is the data-structure of this message
type _BACnetConstructedDataInstantaneousPower struct {
	*_BACnetConstructedData
	InstantaneousPower BACnetApplicationTagReal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataInstantaneousPower) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataInstantaneousPower) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_INSTANTANEOUS_POWER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataInstantaneousPower) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataInstantaneousPower) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataInstantaneousPower) GetInstantaneousPower() BACnetApplicationTagReal {
	return m.InstantaneousPower
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataInstantaneousPower) GetActualValue() BACnetApplicationTagReal {
	return CastBACnetApplicationTagReal(m.GetInstantaneousPower())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataInstantaneousPower factory function for _BACnetConstructedDataInstantaneousPower
func NewBACnetConstructedDataInstantaneousPower(instantaneousPower BACnetApplicationTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataInstantaneousPower {
	_result := &_BACnetConstructedDataInstantaneousPower{
		InstantaneousPower:     instantaneousPower,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataInstantaneousPower(structType interface{}) BACnetConstructedDataInstantaneousPower {
	if casted, ok := structType.(BACnetConstructedDataInstantaneousPower); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataInstantaneousPower); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataInstantaneousPower) GetTypeName() string {
	return "BACnetConstructedDataInstantaneousPower"
}

func (m *_BACnetConstructedDataInstantaneousPower) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataInstantaneousPower) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (instantaneousPower)
	lengthInBits += m.InstantaneousPower.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataInstantaneousPower) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataInstantaneousPowerParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataInstantaneousPower, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataInstantaneousPower"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataInstantaneousPower")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (instantaneousPower)
	if pullErr := readBuffer.PullContext("instantaneousPower"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for instantaneousPower")
	}
	_instantaneousPower, _instantaneousPowerErr := BACnetApplicationTagParse(readBuffer)
	if _instantaneousPowerErr != nil {
		return nil, errors.Wrap(_instantaneousPowerErr, "Error parsing 'instantaneousPower' field")
	}
	instantaneousPower := _instantaneousPower.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("instantaneousPower"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for instantaneousPower")
	}

	// Virtual field
	_actualValue := instantaneousPower
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataInstantaneousPower"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataInstantaneousPower")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataInstantaneousPower{
		InstantaneousPower:     instantaneousPower,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataInstantaneousPower) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataInstantaneousPower"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataInstantaneousPower")
		}

		// Simple Field (instantaneousPower)
		if pushErr := writeBuffer.PushContext("instantaneousPower"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for instantaneousPower")
		}
		_instantaneousPowerErr := writeBuffer.WriteSerializable(m.GetInstantaneousPower())
		if popErr := writeBuffer.PopContext("instantaneousPower"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for instantaneousPower")
		}
		if _instantaneousPowerErr != nil {
			return errors.Wrap(_instantaneousPowerErr, "Error serializing 'instantaneousPower' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataInstantaneousPower"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataInstantaneousPower")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataInstantaneousPower) isBACnetConstructedDataInstantaneousPower() bool {
	return true
}

func (m *_BACnetConstructedDataInstantaneousPower) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
