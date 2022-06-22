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

// COTPPacketDisconnectResponse is the corresponding interface of COTPPacketDisconnectResponse
type COTPPacketDisconnectResponse interface {
	utils.LengthAware
	utils.Serializable
	COTPPacket
	// GetDestinationReference returns DestinationReference (property field)
	GetDestinationReference() uint16
	// GetSourceReference returns SourceReference (property field)
	GetSourceReference() uint16
}

// COTPPacketDisconnectResponseExactly can be used when we want exactly this type and not a type which fulfills COTPPacketDisconnectResponse.
// This is useful for switch cases.
type COTPPacketDisconnectResponseExactly interface {
	COTPPacketDisconnectResponse
	isCOTPPacketDisconnectResponse() bool
}

// _COTPPacketDisconnectResponse is the data-structure of this message
type _COTPPacketDisconnectResponse struct {
	*_COTPPacket
	DestinationReference uint16
	SourceReference      uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_COTPPacketDisconnectResponse) GetTpduCode() uint8 {
	return 0xC0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_COTPPacketDisconnectResponse) InitializeParent(parent COTPPacket, parameters []COTPParameter, payload S7Message) {
	m.Parameters = parameters
	m.Payload = payload
}

func (m *_COTPPacketDisconnectResponse) GetParent() COTPPacket {
	return m._COTPPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_COTPPacketDisconnectResponse) GetDestinationReference() uint16 {
	return m.DestinationReference
}

func (m *_COTPPacketDisconnectResponse) GetSourceReference() uint16 {
	return m.SourceReference
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCOTPPacketDisconnectResponse factory function for _COTPPacketDisconnectResponse
func NewCOTPPacketDisconnectResponse(destinationReference uint16, sourceReference uint16, parameters []COTPParameter, payload S7Message, cotpLen uint16) *_COTPPacketDisconnectResponse {
	_result := &_COTPPacketDisconnectResponse{
		DestinationReference: destinationReference,
		SourceReference:      sourceReference,
		_COTPPacket:          NewCOTPPacket(parameters, payload, cotpLen),
	}
	_result._COTPPacket._COTPPacketChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCOTPPacketDisconnectResponse(structType interface{}) COTPPacketDisconnectResponse {
	if casted, ok := structType.(COTPPacketDisconnectResponse); ok {
		return casted
	}
	if casted, ok := structType.(*COTPPacketDisconnectResponse); ok {
		return *casted
	}
	return nil
}

func (m *_COTPPacketDisconnectResponse) GetTypeName() string {
	return "COTPPacketDisconnectResponse"
}

func (m *_COTPPacketDisconnectResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_COTPPacketDisconnectResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (destinationReference)
	lengthInBits += 16

	// Simple field (sourceReference)
	lengthInBits += 16

	return lengthInBits
}

func (m *_COTPPacketDisconnectResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func COTPPacketDisconnectResponseParse(readBuffer utils.ReadBuffer, cotpLen uint16) (COTPPacketDisconnectResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("COTPPacketDisconnectResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for COTPPacketDisconnectResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (destinationReference)
	_destinationReference, _destinationReferenceErr := readBuffer.ReadUint16("destinationReference", 16)
	if _destinationReferenceErr != nil {
		return nil, errors.Wrap(_destinationReferenceErr, "Error parsing 'destinationReference' field")
	}
	destinationReference := _destinationReference

	// Simple Field (sourceReference)
	_sourceReference, _sourceReferenceErr := readBuffer.ReadUint16("sourceReference", 16)
	if _sourceReferenceErr != nil {
		return nil, errors.Wrap(_sourceReferenceErr, "Error parsing 'sourceReference' field")
	}
	sourceReference := _sourceReference

	if closeErr := readBuffer.CloseContext("COTPPacketDisconnectResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for COTPPacketDisconnectResponse")
	}

	// Create a partially initialized instance
	_child := &_COTPPacketDisconnectResponse{
		DestinationReference: destinationReference,
		SourceReference:      sourceReference,
		_COTPPacket:          &_COTPPacket{},
	}
	_child._COTPPacket._COTPPacketChildRequirements = _child
	return _child, nil
}

func (m *_COTPPacketDisconnectResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("COTPPacketDisconnectResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for COTPPacketDisconnectResponse")
		}

		// Simple Field (destinationReference)
		destinationReference := uint16(m.GetDestinationReference())
		_destinationReferenceErr := writeBuffer.WriteUint16("destinationReference", 16, (destinationReference))
		if _destinationReferenceErr != nil {
			return errors.Wrap(_destinationReferenceErr, "Error serializing 'destinationReference' field")
		}

		// Simple Field (sourceReference)
		sourceReference := uint16(m.GetSourceReference())
		_sourceReferenceErr := writeBuffer.WriteUint16("sourceReference", 16, (sourceReference))
		if _sourceReferenceErr != nil {
			return errors.Wrap(_sourceReferenceErr, "Error serializing 'sourceReference' field")
		}

		if popErr := writeBuffer.PopContext("COTPPacketDisconnectResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for COTPPacketDisconnectResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_COTPPacketDisconnectResponse) isCOTPPacketDisconnectResponse() bool {
	return true
}

func (m *_COTPPacketDisconnectResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
