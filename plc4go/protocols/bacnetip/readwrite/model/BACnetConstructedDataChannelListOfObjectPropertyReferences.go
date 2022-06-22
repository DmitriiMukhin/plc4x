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
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataChannelListOfObjectPropertyReferences is the corresponding interface of BACnetConstructedDataChannelListOfObjectPropertyReferences
type BACnetConstructedDataChannelListOfObjectPropertyReferences interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetReferences returns References (property field)
	GetReferences() []BACnetDeviceObjectPropertyReference
	// GetZero returns Zero (virtual field)
	GetZero() uint64
}

// BACnetConstructedDataChannelListOfObjectPropertyReferencesExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataChannelListOfObjectPropertyReferences.
// This is useful for switch cases.
type BACnetConstructedDataChannelListOfObjectPropertyReferencesExactly interface {
	BACnetConstructedDataChannelListOfObjectPropertyReferences
	isBACnetConstructedDataChannelListOfObjectPropertyReferences() bool
}

// _BACnetConstructedDataChannelListOfObjectPropertyReferences is the data-structure of this message
type _BACnetConstructedDataChannelListOfObjectPropertyReferences struct {
	*_BACnetConstructedData
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	References           []BACnetDeviceObjectPropertyReference
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataChannelListOfObjectPropertyReferences) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_CHANNEL
}

func (m *_BACnetConstructedDataChannelListOfObjectPropertyReferences) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LIST_OF_OBJECT_PROPERTY_REFERENCES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataChannelListOfObjectPropertyReferences) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataChannelListOfObjectPropertyReferences) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataChannelListOfObjectPropertyReferences) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataChannelListOfObjectPropertyReferences) GetReferences() []BACnetDeviceObjectPropertyReference {
	return m.References
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataChannelListOfObjectPropertyReferences) GetZero() uint64 {
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataChannelListOfObjectPropertyReferences factory function for _BACnetConstructedDataChannelListOfObjectPropertyReferences
func NewBACnetConstructedDataChannelListOfObjectPropertyReferences(numberOfDataElements BACnetApplicationTagUnsignedInteger, references []BACnetDeviceObjectPropertyReference, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataChannelListOfObjectPropertyReferences {
	_result := &_BACnetConstructedDataChannelListOfObjectPropertyReferences{
		NumberOfDataElements:   numberOfDataElements,
		References:             references,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataChannelListOfObjectPropertyReferences(structType interface{}) BACnetConstructedDataChannelListOfObjectPropertyReferences {
	if casted, ok := structType.(BACnetConstructedDataChannelListOfObjectPropertyReferences); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataChannelListOfObjectPropertyReferences); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataChannelListOfObjectPropertyReferences) GetTypeName() string {
	return "BACnetConstructedDataChannelListOfObjectPropertyReferences"
}

func (m *_BACnetConstructedDataChannelListOfObjectPropertyReferences) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataChannelListOfObjectPropertyReferences) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits()
	}

	// Array field
	if len(m.References) > 0 {
		for _, element := range m.References {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataChannelListOfObjectPropertyReferences) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataChannelListOfObjectPropertyReferencesParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataChannelListOfObjectPropertyReferences, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataChannelListOfObjectPropertyReferences"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataChannelListOfObjectPropertyReferences")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Virtual field
	_zero := uint64(0)
	zero := uint64(_zero)
	_ = zero

	// Optional Field (numberOfDataElements) (Can be skipped, if a given expression evaluates to false)
	var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
	if bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("numberOfDataElements"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for numberOfDataElements")
		}
		_val, _err := BACnetApplicationTagParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'numberOfDataElements' field")
		default:
			numberOfDataElements = _val.(BACnetApplicationTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("numberOfDataElements"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for numberOfDataElements")
			}
		}
	}

	// Array field (references)
	if pullErr := readBuffer.PullContext("references", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for references")
	}
	// Terminated array
	var references []BACnetDeviceObjectPropertyReference
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetDeviceObjectPropertyReferenceParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'references' field")
			}
			references = append(references, _item.(BACnetDeviceObjectPropertyReference))

		}
	}
	if closeErr := readBuffer.CloseContext("references", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for references")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataChannelListOfObjectPropertyReferences"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataChannelListOfObjectPropertyReferences")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataChannelListOfObjectPropertyReferences{
		NumberOfDataElements:   numberOfDataElements,
		References:             references,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataChannelListOfObjectPropertyReferences) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataChannelListOfObjectPropertyReferences"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataChannelListOfObjectPropertyReferences")
		}
		// Virtual field
		if _zeroErr := writeBuffer.WriteVirtual("zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		// Optional Field (numberOfDataElements) (Can be skipped, if the value is null)
		var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
		if m.GetNumberOfDataElements() != nil {
			if pushErr := writeBuffer.PushContext("numberOfDataElements"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for numberOfDataElements")
			}
			numberOfDataElements = m.GetNumberOfDataElements()
			_numberOfDataElementsErr := writeBuffer.WriteSerializable(numberOfDataElements)
			if popErr := writeBuffer.PopContext("numberOfDataElements"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for numberOfDataElements")
			}
			if _numberOfDataElementsErr != nil {
				return errors.Wrap(_numberOfDataElementsErr, "Error serializing 'numberOfDataElements' field")
			}
		}

		// Array Field (references)
		if pushErr := writeBuffer.PushContext("references", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for references")
		}
		for _, _element := range m.GetReferences() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'references' field")
			}
		}
		if popErr := writeBuffer.PopContext("references", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for references")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataChannelListOfObjectPropertyReferences"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataChannelListOfObjectPropertyReferences")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataChannelListOfObjectPropertyReferences) isBACnetConstructedDataChannelListOfObjectPropertyReferences() bool {
	return true
}

func (m *_BACnetConstructedDataChannelListOfObjectPropertyReferences) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
