/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetContextTagDouble struct {
	*BACnetContextTag
	Value float64
}

// The corresponding interface
type IBACnetContextTagDouble interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetContextTagDouble) DataType() BACnetDataType {
	return BACnetDataType_DOUBLE
}

func (m *BACnetContextTagDouble) InitializeParent(parent *BACnetContextTag, header *BACnetTagHeader, tagNumber uint8, actualLength uint32) {
	m.Header = header
}

func NewBACnetContextTagDouble(value float64, header *BACnetTagHeader, tagNumber uint8, actualLength uint32) *BACnetContextTag {
	child := &BACnetContextTagDouble{
		Value:            value,
		BACnetContextTag: NewBACnetContextTag(header, tagNumber, actualLength),
	}
	child.Child = child
	return child.BACnetContextTag
}

func CastBACnetContextTagDouble(structType interface{}) *BACnetContextTagDouble {
	castFunc := func(typ interface{}) *BACnetContextTagDouble {
		if casted, ok := typ.(BACnetContextTagDouble); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetContextTagDouble); ok {
			return casted
		}
		if casted, ok := typ.(BACnetContextTag); ok {
			return CastBACnetContextTagDouble(casted.Child)
		}
		if casted, ok := typ.(*BACnetContextTag); ok {
			return CastBACnetContextTagDouble(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetContextTagDouble) GetTypeName() string {
	return "BACnetContextTagDouble"
}

func (m *BACnetContextTagDouble) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetContextTagDouble) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (value)
	lengthInBits += 64

	return lengthInBits
}

func (m *BACnetContextTagDouble) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetContextTagDoubleParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType, actualLength uint32) (*BACnetContextTag, error) {
	if pullErr := readBuffer.PullContext("BACnetContextTagDouble"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (value)
	_value, _valueErr := readBuffer.ReadFloat64("value", 64)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := _value

	if closeErr := readBuffer.CloseContext("BACnetContextTagDouble"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetContextTagDouble{
		Value:            value,
		BACnetContextTag: &BACnetContextTag{},
	}
	_child.BACnetContextTag.Child = _child
	return _child.BACnetContextTag, nil
}

func (m *BACnetContextTagDouble) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagDouble"); pushErr != nil {
			return pushErr
		}

		// Simple Field (value)
		value := float64(m.Value)
		_valueErr := writeBuffer.WriteFloat64("value", 64, (value))
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagDouble"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetContextTagDouble) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}