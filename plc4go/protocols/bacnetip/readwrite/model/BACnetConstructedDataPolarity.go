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

// BACnetConstructedDataPolarity is the corresponding interface of BACnetConstructedDataPolarity
type BACnetConstructedDataPolarity interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetPolarity returns Polarity (property field)
	GetPolarity() BACnetPolarityTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetPolarityTagged
}

// BACnetConstructedDataPolarityExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataPolarity.
// This is useful for switch cases.
type BACnetConstructedDataPolarityExactly interface {
	BACnetConstructedDataPolarity
	isBACnetConstructedDataPolarity() bool
}

// _BACnetConstructedDataPolarity is the data-structure of this message
type _BACnetConstructedDataPolarity struct {
	*_BACnetConstructedData
	Polarity BACnetPolarityTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataPolarity) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataPolarity) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_POLARITY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataPolarity) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataPolarity) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataPolarity) GetPolarity() BACnetPolarityTagged {
	return m.Polarity
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataPolarity) GetActualValue() BACnetPolarityTagged {
	return CastBACnetPolarityTagged(m.GetPolarity())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataPolarity factory function for _BACnetConstructedDataPolarity
func NewBACnetConstructedDataPolarity(polarity BACnetPolarityTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataPolarity {
	_result := &_BACnetConstructedDataPolarity{
		Polarity:               polarity,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataPolarity(structType interface{}) BACnetConstructedDataPolarity {
	if casted, ok := structType.(BACnetConstructedDataPolarity); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPolarity); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataPolarity) GetTypeName() string {
	return "BACnetConstructedDataPolarity"
}

func (m *_BACnetConstructedDataPolarity) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataPolarity) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (polarity)
	lengthInBits += m.Polarity.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataPolarity) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataPolarityParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataPolarity, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPolarity"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPolarity")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (polarity)
	if pullErr := readBuffer.PullContext("polarity"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for polarity")
	}
	_polarity, _polarityErr := BACnetPolarityTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _polarityErr != nil {
		return nil, errors.Wrap(_polarityErr, "Error parsing 'polarity' field")
	}
	polarity := _polarity.(BACnetPolarityTagged)
	if closeErr := readBuffer.CloseContext("polarity"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for polarity")
	}

	// Virtual field
	_actualValue := polarity
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPolarity"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPolarity")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataPolarity{
		Polarity:               polarity,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataPolarity) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPolarity"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPolarity")
		}

		// Simple Field (polarity)
		if pushErr := writeBuffer.PushContext("polarity"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for polarity")
		}
		_polarityErr := writeBuffer.WriteSerializable(m.GetPolarity())
		if popErr := writeBuffer.PopContext("polarity"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for polarity")
		}
		if _polarityErr != nil {
			return errors.Wrap(_polarityErr, "Error serializing 'polarity' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPolarity"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPolarity")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataPolarity) isBACnetConstructedDataPolarity() bool {
	return true
}

func (m *_BACnetConstructedDataPolarity) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
