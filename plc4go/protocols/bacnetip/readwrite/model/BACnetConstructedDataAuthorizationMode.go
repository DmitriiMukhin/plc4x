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

// BACnetConstructedDataAuthorizationMode is the corresponding interface of BACnetConstructedDataAuthorizationMode
type BACnetConstructedDataAuthorizationMode interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAuthorizationMode returns AuthorizationMode (property field)
	GetAuthorizationMode() BACnetAuthorizationModeTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetAuthorizationModeTagged
}

// BACnetConstructedDataAuthorizationModeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAuthorizationMode.
// This is useful for switch cases.
type BACnetConstructedDataAuthorizationModeExactly interface {
	BACnetConstructedDataAuthorizationMode
	isBACnetConstructedDataAuthorizationMode() bool
}

// _BACnetConstructedDataAuthorizationMode is the data-structure of this message
type _BACnetConstructedDataAuthorizationMode struct {
	*_BACnetConstructedData
	AuthorizationMode BACnetAuthorizationModeTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAuthorizationMode) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAuthorizationMode) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_AUTHORIZATION_MODE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAuthorizationMode) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAuthorizationMode) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAuthorizationMode) GetAuthorizationMode() BACnetAuthorizationModeTagged {
	return m.AuthorizationMode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAuthorizationMode) GetActualValue() BACnetAuthorizationModeTagged {
	return CastBACnetAuthorizationModeTagged(m.GetAuthorizationMode())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAuthorizationMode factory function for _BACnetConstructedDataAuthorizationMode
func NewBACnetConstructedDataAuthorizationMode(authorizationMode BACnetAuthorizationModeTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAuthorizationMode {
	_result := &_BACnetConstructedDataAuthorizationMode{
		AuthorizationMode:      authorizationMode,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAuthorizationMode(structType interface{}) BACnetConstructedDataAuthorizationMode {
	if casted, ok := structType.(BACnetConstructedDataAuthorizationMode); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAuthorizationMode); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAuthorizationMode) GetTypeName() string {
	return "BACnetConstructedDataAuthorizationMode"
}

func (m *_BACnetConstructedDataAuthorizationMode) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataAuthorizationMode) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (authorizationMode)
	lengthInBits += m.AuthorizationMode.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAuthorizationMode) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAuthorizationModeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAuthorizationMode, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAuthorizationMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAuthorizationMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (authorizationMode)
	if pullErr := readBuffer.PullContext("authorizationMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for authorizationMode")
	}
	_authorizationMode, _authorizationModeErr := BACnetAuthorizationModeTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _authorizationModeErr != nil {
		return nil, errors.Wrap(_authorizationModeErr, "Error parsing 'authorizationMode' field")
	}
	authorizationMode := _authorizationMode.(BACnetAuthorizationModeTagged)
	if closeErr := readBuffer.CloseContext("authorizationMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for authorizationMode")
	}

	// Virtual field
	_actualValue := authorizationMode
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAuthorizationMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAuthorizationMode")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAuthorizationMode{
		AuthorizationMode:      authorizationMode,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAuthorizationMode) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAuthorizationMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAuthorizationMode")
		}

		// Simple Field (authorizationMode)
		if pushErr := writeBuffer.PushContext("authorizationMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for authorizationMode")
		}
		_authorizationModeErr := writeBuffer.WriteSerializable(m.GetAuthorizationMode())
		if popErr := writeBuffer.PopContext("authorizationMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for authorizationMode")
		}
		if _authorizationModeErr != nil {
			return errors.Wrap(_authorizationModeErr, "Error serializing 'authorizationMode' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAuthorizationMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAuthorizationMode")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAuthorizationMode) isBACnetConstructedDataAuthorizationMode() bool {
	return true
}

func (m *_BACnetConstructedDataAuthorizationMode) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
