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

// BACnetTimerStateChangeValueDateTime is the corresponding interface of BACnetTimerStateChangeValueDateTime
type BACnetTimerStateChangeValueDateTime interface {
	utils.LengthAware
	utils.Serializable
	BACnetTimerStateChangeValue
	// GetDateTimeValue returns DateTimeValue (property field)
	GetDateTimeValue() BACnetDateTimeEnclosed
}

// BACnetTimerStateChangeValueDateTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetTimerStateChangeValueDateTime.
// This is useful for switch cases.
type BACnetTimerStateChangeValueDateTimeExactly interface {
	BACnetTimerStateChangeValueDateTime
	isBACnetTimerStateChangeValueDateTime() bool
}

// _BACnetTimerStateChangeValueDateTime is the data-structure of this message
type _BACnetTimerStateChangeValueDateTime struct {
	*_BACnetTimerStateChangeValue
	DateTimeValue BACnetDateTimeEnclosed
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetTimerStateChangeValueDateTime) InitializeParent(parent BACnetTimerStateChangeValue, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetTimerStateChangeValueDateTime) GetParent() BACnetTimerStateChangeValue {
	return m._BACnetTimerStateChangeValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimerStateChangeValueDateTime) GetDateTimeValue() BACnetDateTimeEnclosed {
	return m.DateTimeValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTimerStateChangeValueDateTime factory function for _BACnetTimerStateChangeValueDateTime
func NewBACnetTimerStateChangeValueDateTime(dateTimeValue BACnetDateTimeEnclosed, peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType) *_BACnetTimerStateChangeValueDateTime {
	_result := &_BACnetTimerStateChangeValueDateTime{
		DateTimeValue:                dateTimeValue,
		_BACnetTimerStateChangeValue: NewBACnetTimerStateChangeValue(peekedTagHeader, objectTypeArgument),
	}
	_result._BACnetTimerStateChangeValue._BACnetTimerStateChangeValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetTimerStateChangeValueDateTime(structType interface{}) BACnetTimerStateChangeValueDateTime {
	if casted, ok := structType.(BACnetTimerStateChangeValueDateTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimerStateChangeValueDateTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimerStateChangeValueDateTime) GetTypeName() string {
	return "BACnetTimerStateChangeValueDateTime"
}

func (m *_BACnetTimerStateChangeValueDateTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetTimerStateChangeValueDateTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (dateTimeValue)
	lengthInBits += m.DateTimeValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetTimerStateChangeValueDateTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTimerStateChangeValueDateTimeParse(readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (BACnetTimerStateChangeValueDateTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimerStateChangeValueDateTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimerStateChangeValueDateTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (dateTimeValue)
	if pullErr := readBuffer.PullContext("dateTimeValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dateTimeValue")
	}
	_dateTimeValue, _dateTimeValueErr := BACnetDateTimeEnclosedParse(readBuffer, uint8(uint8(2)))
	if _dateTimeValueErr != nil {
		return nil, errors.Wrap(_dateTimeValueErr, "Error parsing 'dateTimeValue' field")
	}
	dateTimeValue := _dateTimeValue.(BACnetDateTimeEnclosed)
	if closeErr := readBuffer.CloseContext("dateTimeValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dateTimeValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetTimerStateChangeValueDateTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimerStateChangeValueDateTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetTimerStateChangeValueDateTime{
		DateTimeValue:                dateTimeValue,
		_BACnetTimerStateChangeValue: &_BACnetTimerStateChangeValue{},
	}
	_child._BACnetTimerStateChangeValue._BACnetTimerStateChangeValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetTimerStateChangeValueDateTime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimerStateChangeValueDateTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimerStateChangeValueDateTime")
		}

		// Simple Field (dateTimeValue)
		if pushErr := writeBuffer.PushContext("dateTimeValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for dateTimeValue")
		}
		_dateTimeValueErr := writeBuffer.WriteSerializable(m.GetDateTimeValue())
		if popErr := writeBuffer.PopContext("dateTimeValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for dateTimeValue")
		}
		if _dateTimeValueErr != nil {
			return errors.Wrap(_dateTimeValueErr, "Error serializing 'dateTimeValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimerStateChangeValueDateTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimerStateChangeValueDateTime")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetTimerStateChangeValueDateTime) isBACnetTimerStateChangeValueDateTime() bool {
	return true
}

func (m *_BACnetTimerStateChangeValueDateTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
