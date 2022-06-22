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

// BACnetConstructedDataAccessEventAuthenticationFactor is the corresponding interface of BACnetConstructedDataAccessEventAuthenticationFactor
type BACnetConstructedDataAccessEventAuthenticationFactor interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAccessEventAuthenticationFactor returns AccessEventAuthenticationFactor (property field)
	GetAccessEventAuthenticationFactor() BACnetAuthenticationFactor
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetAuthenticationFactor
}

// BACnetConstructedDataAccessEventAuthenticationFactorExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAccessEventAuthenticationFactor.
// This is useful for switch cases.
type BACnetConstructedDataAccessEventAuthenticationFactorExactly interface {
	BACnetConstructedDataAccessEventAuthenticationFactor
	isBACnetConstructedDataAccessEventAuthenticationFactor() bool
}

// _BACnetConstructedDataAccessEventAuthenticationFactor is the data-structure of this message
type _BACnetConstructedDataAccessEventAuthenticationFactor struct {
	*_BACnetConstructedData
	AccessEventAuthenticationFactor BACnetAuthenticationFactor
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACCESS_EVENT_AUTHENTICATION_FACTOR
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) GetAccessEventAuthenticationFactor() BACnetAuthenticationFactor {
	return m.AccessEventAuthenticationFactor
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) GetActualValue() BACnetAuthenticationFactor {
	return CastBACnetAuthenticationFactor(m.GetAccessEventAuthenticationFactor())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAccessEventAuthenticationFactor factory function for _BACnetConstructedDataAccessEventAuthenticationFactor
func NewBACnetConstructedDataAccessEventAuthenticationFactor(accessEventAuthenticationFactor BACnetAuthenticationFactor, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccessEventAuthenticationFactor {
	_result := &_BACnetConstructedDataAccessEventAuthenticationFactor{
		AccessEventAuthenticationFactor: accessEventAuthenticationFactor,
		_BACnetConstructedData:          NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccessEventAuthenticationFactor(structType interface{}) BACnetConstructedDataAccessEventAuthenticationFactor {
	if casted, ok := structType.(BACnetConstructedDataAccessEventAuthenticationFactor); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccessEventAuthenticationFactor); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) GetTypeName() string {
	return "BACnetConstructedDataAccessEventAuthenticationFactor"
}

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (accessEventAuthenticationFactor)
	lengthInBits += m.AccessEventAuthenticationFactor.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAccessEventAuthenticationFactorParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAccessEventAuthenticationFactor, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccessEventAuthenticationFactor"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccessEventAuthenticationFactor")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (accessEventAuthenticationFactor)
	if pullErr := readBuffer.PullContext("accessEventAuthenticationFactor"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for accessEventAuthenticationFactor")
	}
	_accessEventAuthenticationFactor, _accessEventAuthenticationFactorErr := BACnetAuthenticationFactorParse(readBuffer)
	if _accessEventAuthenticationFactorErr != nil {
		return nil, errors.Wrap(_accessEventAuthenticationFactorErr, "Error parsing 'accessEventAuthenticationFactor' field")
	}
	accessEventAuthenticationFactor := _accessEventAuthenticationFactor.(BACnetAuthenticationFactor)
	if closeErr := readBuffer.CloseContext("accessEventAuthenticationFactor"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for accessEventAuthenticationFactor")
	}

	// Virtual field
	_actualValue := accessEventAuthenticationFactor
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccessEventAuthenticationFactor"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccessEventAuthenticationFactor")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAccessEventAuthenticationFactor{
		AccessEventAuthenticationFactor: accessEventAuthenticationFactor,
		_BACnetConstructedData:          &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccessEventAuthenticationFactor"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccessEventAuthenticationFactor")
		}

		// Simple Field (accessEventAuthenticationFactor)
		if pushErr := writeBuffer.PushContext("accessEventAuthenticationFactor"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for accessEventAuthenticationFactor")
		}
		_accessEventAuthenticationFactorErr := writeBuffer.WriteSerializable(m.GetAccessEventAuthenticationFactor())
		if popErr := writeBuffer.PopContext("accessEventAuthenticationFactor"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for accessEventAuthenticationFactor")
		}
		if _accessEventAuthenticationFactorErr != nil {
			return errors.Wrap(_accessEventAuthenticationFactorErr, "Error serializing 'accessEventAuthenticationFactor' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccessEventAuthenticationFactor"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccessEventAuthenticationFactor")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) isBACnetConstructedDataAccessEventAuthenticationFactor() bool {
	return true
}

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
