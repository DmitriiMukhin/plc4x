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

// BACnetTimerStateChangeValueObjectidentifier is the corresponding interface of BACnetTimerStateChangeValueObjectidentifier
type BACnetTimerStateChangeValueObjectidentifier interface {
	utils.LengthAware
	utils.Serializable
	BACnetTimerStateChangeValue
	// GetObjectidentifierValue returns ObjectidentifierValue (property field)
	GetObjectidentifierValue() BACnetApplicationTagObjectIdentifier
}

// BACnetTimerStateChangeValueObjectidentifierExactly can be used when we want exactly this type and not a type which fulfills BACnetTimerStateChangeValueObjectidentifier.
// This is useful for switch cases.
type BACnetTimerStateChangeValueObjectidentifierExactly interface {
	BACnetTimerStateChangeValueObjectidentifier
	isBACnetTimerStateChangeValueObjectidentifier() bool
}

// _BACnetTimerStateChangeValueObjectidentifier is the data-structure of this message
type _BACnetTimerStateChangeValueObjectidentifier struct {
	*_BACnetTimerStateChangeValue
	ObjectidentifierValue BACnetApplicationTagObjectIdentifier
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetTimerStateChangeValueObjectidentifier) InitializeParent(parent BACnetTimerStateChangeValue, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetTimerStateChangeValueObjectidentifier) GetParent() BACnetTimerStateChangeValue {
	return m._BACnetTimerStateChangeValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimerStateChangeValueObjectidentifier) GetObjectidentifierValue() BACnetApplicationTagObjectIdentifier {
	return m.ObjectidentifierValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTimerStateChangeValueObjectidentifier factory function for _BACnetTimerStateChangeValueObjectidentifier
func NewBACnetTimerStateChangeValueObjectidentifier(objectidentifierValue BACnetApplicationTagObjectIdentifier, peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType) *_BACnetTimerStateChangeValueObjectidentifier {
	_result := &_BACnetTimerStateChangeValueObjectidentifier{
		ObjectidentifierValue:        objectidentifierValue,
		_BACnetTimerStateChangeValue: NewBACnetTimerStateChangeValue(peekedTagHeader, objectTypeArgument),
	}
	_result._BACnetTimerStateChangeValue._BACnetTimerStateChangeValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetTimerStateChangeValueObjectidentifier(structType interface{}) BACnetTimerStateChangeValueObjectidentifier {
	if casted, ok := structType.(BACnetTimerStateChangeValueObjectidentifier); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimerStateChangeValueObjectidentifier); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimerStateChangeValueObjectidentifier) GetTypeName() string {
	return "BACnetTimerStateChangeValueObjectidentifier"
}

func (m *_BACnetTimerStateChangeValueObjectidentifier) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetTimerStateChangeValueObjectidentifier) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (objectidentifierValue)
	lengthInBits += m.ObjectidentifierValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetTimerStateChangeValueObjectidentifier) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTimerStateChangeValueObjectidentifierParse(readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (BACnetTimerStateChangeValueObjectidentifier, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimerStateChangeValueObjectidentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimerStateChangeValueObjectidentifier")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectidentifierValue)
	if pullErr := readBuffer.PullContext("objectidentifierValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectidentifierValue")
	}
	_objectidentifierValue, _objectidentifierValueErr := BACnetApplicationTagParse(readBuffer)
	if _objectidentifierValueErr != nil {
		return nil, errors.Wrap(_objectidentifierValueErr, "Error parsing 'objectidentifierValue' field")
	}
	objectidentifierValue := _objectidentifierValue.(BACnetApplicationTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("objectidentifierValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for objectidentifierValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetTimerStateChangeValueObjectidentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimerStateChangeValueObjectidentifier")
	}

	// Create a partially initialized instance
	_child := &_BACnetTimerStateChangeValueObjectidentifier{
		ObjectidentifierValue:        objectidentifierValue,
		_BACnetTimerStateChangeValue: &_BACnetTimerStateChangeValue{},
	}
	_child._BACnetTimerStateChangeValue._BACnetTimerStateChangeValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetTimerStateChangeValueObjectidentifier) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimerStateChangeValueObjectidentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimerStateChangeValueObjectidentifier")
		}

		// Simple Field (objectidentifierValue)
		if pushErr := writeBuffer.PushContext("objectidentifierValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for objectidentifierValue")
		}
		_objectidentifierValueErr := writeBuffer.WriteSerializable(m.GetObjectidentifierValue())
		if popErr := writeBuffer.PopContext("objectidentifierValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for objectidentifierValue")
		}
		if _objectidentifierValueErr != nil {
			return errors.Wrap(_objectidentifierValueErr, "Error serializing 'objectidentifierValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimerStateChangeValueObjectidentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimerStateChangeValueObjectidentifier")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetTimerStateChangeValueObjectidentifier) isBACnetTimerStateChangeValueObjectidentifier() bool {
	return true
}

func (m *_BACnetTimerStateChangeValueObjectidentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
