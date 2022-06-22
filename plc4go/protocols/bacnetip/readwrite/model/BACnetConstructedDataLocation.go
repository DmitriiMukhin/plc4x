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

// BACnetConstructedDataLocation is the corresponding interface of BACnetConstructedDataLocation
type BACnetConstructedDataLocation interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLocation returns Location (property field)
	GetLocation() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
}

// BACnetConstructedDataLocationExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLocation.
// This is useful for switch cases.
type BACnetConstructedDataLocationExactly interface {
	BACnetConstructedDataLocation
	isBACnetConstructedDataLocation() bool
}

// _BACnetConstructedDataLocation is the data-structure of this message
type _BACnetConstructedDataLocation struct {
	*_BACnetConstructedData
	Location BACnetApplicationTagCharacterString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLocation) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLocation) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOCATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLocation) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLocation) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLocation) GetLocation() BACnetApplicationTagCharacterString {
	return m.Location
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLocation) GetActualValue() BACnetApplicationTagCharacterString {
	return CastBACnetApplicationTagCharacterString(m.GetLocation())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLocation factory function for _BACnetConstructedDataLocation
func NewBACnetConstructedDataLocation(location BACnetApplicationTagCharacterString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLocation {
	_result := &_BACnetConstructedDataLocation{
		Location:               location,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLocation(structType interface{}) BACnetConstructedDataLocation {
	if casted, ok := structType.(BACnetConstructedDataLocation); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLocation); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLocation) GetTypeName() string {
	return "BACnetConstructedDataLocation"
}

func (m *_BACnetConstructedDataLocation) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataLocation) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (location)
	lengthInBits += m.Location.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLocation) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLocationParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLocation, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLocation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLocation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (location)
	if pullErr := readBuffer.PullContext("location"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for location")
	}
	_location, _locationErr := BACnetApplicationTagParse(readBuffer)
	if _locationErr != nil {
		return nil, errors.Wrap(_locationErr, "Error parsing 'location' field")
	}
	location := _location.(BACnetApplicationTagCharacterString)
	if closeErr := readBuffer.CloseContext("location"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for location")
	}

	// Virtual field
	_actualValue := location
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLocation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLocation")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLocation{
		Location:               location,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLocation) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLocation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLocation")
		}

		// Simple Field (location)
		if pushErr := writeBuffer.PushContext("location"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for location")
		}
		_locationErr := writeBuffer.WriteSerializable(m.GetLocation())
		if popErr := writeBuffer.PopContext("location"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for location")
		}
		if _locationErr != nil {
			return errors.Wrap(_locationErr, "Error serializing 'location' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLocation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLocation")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLocation) isBACnetConstructedDataLocation() bool {
	return true
}

func (m *_BACnetConstructedDataLocation) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
