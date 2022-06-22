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

// BACnetTimerStateChangeValueOctetString is the corresponding interface of BACnetTimerStateChangeValueOctetString
type BACnetTimerStateChangeValueOctetString interface {
	utils.LengthAware
	utils.Serializable
	BACnetTimerStateChangeValue
	// GetOctetStringValue returns OctetStringValue (property field)
	GetOctetStringValue() BACnetApplicationTagOctetString
}

// BACnetTimerStateChangeValueOctetStringExactly can be used when we want exactly this type and not a type which fulfills BACnetTimerStateChangeValueOctetString.
// This is useful for switch cases.
type BACnetTimerStateChangeValueOctetStringExactly interface {
	BACnetTimerStateChangeValueOctetString
	isBACnetTimerStateChangeValueOctetString() bool
}

// _BACnetTimerStateChangeValueOctetString is the data-structure of this message
type _BACnetTimerStateChangeValueOctetString struct {
	*_BACnetTimerStateChangeValue
	OctetStringValue BACnetApplicationTagOctetString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetTimerStateChangeValueOctetString) InitializeParent(parent BACnetTimerStateChangeValue, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetTimerStateChangeValueOctetString) GetParent() BACnetTimerStateChangeValue {
	return m._BACnetTimerStateChangeValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimerStateChangeValueOctetString) GetOctetStringValue() BACnetApplicationTagOctetString {
	return m.OctetStringValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTimerStateChangeValueOctetString factory function for _BACnetTimerStateChangeValueOctetString
func NewBACnetTimerStateChangeValueOctetString(octetStringValue BACnetApplicationTagOctetString, peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType) *_BACnetTimerStateChangeValueOctetString {
	_result := &_BACnetTimerStateChangeValueOctetString{
		OctetStringValue:             octetStringValue,
		_BACnetTimerStateChangeValue: NewBACnetTimerStateChangeValue(peekedTagHeader, objectTypeArgument),
	}
	_result._BACnetTimerStateChangeValue._BACnetTimerStateChangeValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetTimerStateChangeValueOctetString(structType interface{}) BACnetTimerStateChangeValueOctetString {
	if casted, ok := structType.(BACnetTimerStateChangeValueOctetString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimerStateChangeValueOctetString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimerStateChangeValueOctetString) GetTypeName() string {
	return "BACnetTimerStateChangeValueOctetString"
}

func (m *_BACnetTimerStateChangeValueOctetString) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetTimerStateChangeValueOctetString) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (octetStringValue)
	lengthInBits += m.OctetStringValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetTimerStateChangeValueOctetString) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTimerStateChangeValueOctetStringParse(readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (BACnetTimerStateChangeValueOctetString, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimerStateChangeValueOctetString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimerStateChangeValueOctetString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (octetStringValue)
	if pullErr := readBuffer.PullContext("octetStringValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for octetStringValue")
	}
	_octetStringValue, _octetStringValueErr := BACnetApplicationTagParse(readBuffer)
	if _octetStringValueErr != nil {
		return nil, errors.Wrap(_octetStringValueErr, "Error parsing 'octetStringValue' field")
	}
	octetStringValue := _octetStringValue.(BACnetApplicationTagOctetString)
	if closeErr := readBuffer.CloseContext("octetStringValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for octetStringValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetTimerStateChangeValueOctetString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimerStateChangeValueOctetString")
	}

	// Create a partially initialized instance
	_child := &_BACnetTimerStateChangeValueOctetString{
		OctetStringValue:             octetStringValue,
		_BACnetTimerStateChangeValue: &_BACnetTimerStateChangeValue{},
	}
	_child._BACnetTimerStateChangeValue._BACnetTimerStateChangeValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetTimerStateChangeValueOctetString) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimerStateChangeValueOctetString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimerStateChangeValueOctetString")
		}

		// Simple Field (octetStringValue)
		if pushErr := writeBuffer.PushContext("octetStringValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for octetStringValue")
		}
		_octetStringValueErr := writeBuffer.WriteSerializable(m.GetOctetStringValue())
		if popErr := writeBuffer.PopContext("octetStringValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for octetStringValue")
		}
		if _octetStringValueErr != nil {
			return errors.Wrap(_octetStringValueErr, "Error serializing 'octetStringValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimerStateChangeValueOctetString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimerStateChangeValueOctetString")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetTimerStateChangeValueOctetString) isBACnetTimerStateChangeValueOctetString() bool {
	return true
}

func (m *_BACnetTimerStateChangeValueOctetString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
