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

// BACnetConstructedDataValidSamples is the corresponding interface of BACnetConstructedDataValidSamples
type BACnetConstructedDataValidSamples interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetValidSamples returns ValidSamples (property field)
	GetValidSamples() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataValidSamplesExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataValidSamples.
// This is useful for switch cases.
type BACnetConstructedDataValidSamplesExactly interface {
	BACnetConstructedDataValidSamples
	isBACnetConstructedDataValidSamples() bool
}

// _BACnetConstructedDataValidSamples is the data-structure of this message
type _BACnetConstructedDataValidSamples struct {
	*_BACnetConstructedData
	ValidSamples BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataValidSamples) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataValidSamples) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_VALID_SAMPLES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataValidSamples) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataValidSamples) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataValidSamples) GetValidSamples() BACnetApplicationTagUnsignedInteger {
	return m.ValidSamples
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataValidSamples) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetValidSamples())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataValidSamples factory function for _BACnetConstructedDataValidSamples
func NewBACnetConstructedDataValidSamples(validSamples BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataValidSamples {
	_result := &_BACnetConstructedDataValidSamples{
		ValidSamples:           validSamples,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataValidSamples(structType interface{}) BACnetConstructedDataValidSamples {
	if casted, ok := structType.(BACnetConstructedDataValidSamples); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataValidSamples); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataValidSamples) GetTypeName() string {
	return "BACnetConstructedDataValidSamples"
}

func (m *_BACnetConstructedDataValidSamples) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataValidSamples) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (validSamples)
	lengthInBits += m.ValidSamples.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataValidSamples) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataValidSamplesParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataValidSamples, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataValidSamples"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataValidSamples")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (validSamples)
	if pullErr := readBuffer.PullContext("validSamples"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for validSamples")
	}
	_validSamples, _validSamplesErr := BACnetApplicationTagParse(readBuffer)
	if _validSamplesErr != nil {
		return nil, errors.Wrap(_validSamplesErr, "Error parsing 'validSamples' field")
	}
	validSamples := _validSamples.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("validSamples"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for validSamples")
	}

	// Virtual field
	_actualValue := validSamples
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataValidSamples"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataValidSamples")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataValidSamples{
		ValidSamples:           validSamples,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataValidSamples) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataValidSamples"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataValidSamples")
		}

		// Simple Field (validSamples)
		if pushErr := writeBuffer.PushContext("validSamples"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for validSamples")
		}
		_validSamplesErr := writeBuffer.WriteSerializable(m.GetValidSamples())
		if popErr := writeBuffer.PopContext("validSamples"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for validSamples")
		}
		if _validSamplesErr != nil {
			return errors.Wrap(_validSamplesErr, "Error serializing 'validSamples' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataValidSamples"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataValidSamples")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataValidSamples) isBACnetConstructedDataValidSamples() bool {
	return true
}

func (m *_BACnetConstructedDataValidSamples) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
