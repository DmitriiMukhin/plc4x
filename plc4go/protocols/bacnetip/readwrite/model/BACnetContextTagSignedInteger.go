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

// BACnetContextTagSignedInteger is the corresponding interface of BACnetContextTagSignedInteger
type BACnetContextTagSignedInteger interface {
	utils.LengthAware
	utils.Serializable
	BACnetContextTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadSignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() uint64
}

// BACnetContextTagSignedIntegerExactly can be used when we want exactly this type and not a type which fulfills BACnetContextTagSignedInteger.
// This is useful for switch cases.
type BACnetContextTagSignedIntegerExactly interface {
	BACnetContextTagSignedInteger
	isBACnetContextTagSignedInteger() bool
}

// _BACnetContextTagSignedInteger is the data-structure of this message
type _BACnetContextTagSignedInteger struct {
	*_BACnetContextTag
	Payload BACnetTagPayloadSignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetContextTagSignedInteger) GetDataType() BACnetDataType {
	return BACnetDataType_SIGNED_INTEGER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetContextTagSignedInteger) InitializeParent(parent BACnetContextTag, header BACnetTagHeader) {
	m.Header = header
}

func (m *_BACnetContextTagSignedInteger) GetParent() BACnetContextTag {
	return m._BACnetContextTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetContextTagSignedInteger) GetPayload() BACnetTagPayloadSignedInteger {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetContextTagSignedInteger) GetActualValue() uint64 {
	return uint64(m.GetPayload().GetActualValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetContextTagSignedInteger factory function for _BACnetContextTagSignedInteger
func NewBACnetContextTagSignedInteger(payload BACnetTagPayloadSignedInteger, header BACnetTagHeader, tagNumberArgument uint8) *_BACnetContextTagSignedInteger {
	_result := &_BACnetContextTagSignedInteger{
		Payload:           payload,
		_BACnetContextTag: NewBACnetContextTag(header, tagNumberArgument),
	}
	_result._BACnetContextTag._BACnetContextTagChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetContextTagSignedInteger(structType interface{}) BACnetContextTagSignedInteger {
	if casted, ok := structType.(BACnetContextTagSignedInteger); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetContextTagSignedInteger); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetContextTagSignedInteger) GetTypeName() string {
	return "BACnetContextTagSignedInteger"
}

func (m *_BACnetContextTagSignedInteger) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetContextTagSignedInteger) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetContextTagSignedInteger) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetContextTagSignedIntegerParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType, header BACnetTagHeader) (BACnetContextTagSignedInteger, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetContextTagSignedInteger"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTagSignedInteger")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (payload)
	if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for payload")
	}
	_payload, _payloadErr := BACnetTagPayloadSignedIntegerParse(readBuffer, uint32(header.GetActualLength()))
	if _payloadErr != nil {
		return nil, errors.Wrap(_payloadErr, "Error parsing 'payload' field")
	}
	payload := _payload.(BACnetTagPayloadSignedInteger)
	if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for payload")
	}

	// Virtual field
	_actualValue := payload.GetActualValue()
	actualValue := uint64(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetContextTagSignedInteger"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTagSignedInteger")
	}

	// Create a partially initialized instance
	_child := &_BACnetContextTagSignedInteger{
		Payload:           payload,
		_BACnetContextTag: &_BACnetContextTag{},
	}
	_child._BACnetContextTag._BACnetContextTagChildRequirements = _child
	return _child, nil
}

func (m *_BACnetContextTagSignedInteger) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagSignedInteger"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetContextTagSignedInteger")
		}

		// Simple Field (payload)
		if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for payload")
		}
		_payloadErr := writeBuffer.WriteSerializable(m.GetPayload())
		if popErr := writeBuffer.PopContext("payload"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for payload")
		}
		if _payloadErr != nil {
			return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagSignedInteger"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetContextTagSignedInteger")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetContextTagSignedInteger) isBACnetContextTagSignedInteger() bool {
	return true
}

func (m *_BACnetContextTagSignedInteger) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
