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

// FirmataMessageDigitalIO is the corresponding interface of FirmataMessageDigitalIO
type FirmataMessageDigitalIO interface {
	utils.LengthAware
	utils.Serializable
	FirmataMessage
	// GetPinBlock returns PinBlock (property field)
	GetPinBlock() uint8
	// GetData returns Data (property field)
	GetData() []int8
}

// FirmataMessageDigitalIOExactly can be used when we want exactly this type and not a type which fulfills FirmataMessageDigitalIO.
// This is useful for switch cases.
type FirmataMessageDigitalIOExactly interface {
	FirmataMessageDigitalIO
	isFirmataMessageDigitalIO() bool
}

// _FirmataMessageDigitalIO is the data-structure of this message
type _FirmataMessageDigitalIO struct {
	*_FirmataMessage
	PinBlock uint8
	Data     []int8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FirmataMessageDigitalIO) GetMessageType() uint8 {
	return 0x9
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FirmataMessageDigitalIO) InitializeParent(parent FirmataMessage) {}

func (m *_FirmataMessageDigitalIO) GetParent() FirmataMessage {
	return m._FirmataMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_FirmataMessageDigitalIO) GetPinBlock() uint8 {
	return m.PinBlock
}

func (m *_FirmataMessageDigitalIO) GetData() []int8 {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewFirmataMessageDigitalIO factory function for _FirmataMessageDigitalIO
func NewFirmataMessageDigitalIO(pinBlock uint8, data []int8, response bool) *_FirmataMessageDigitalIO {
	_result := &_FirmataMessageDigitalIO{
		PinBlock:        pinBlock,
		Data:            data,
		_FirmataMessage: NewFirmataMessage(response),
	}
	_result._FirmataMessage._FirmataMessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastFirmataMessageDigitalIO(structType interface{}) FirmataMessageDigitalIO {
	if casted, ok := structType.(FirmataMessageDigitalIO); ok {
		return casted
	}
	if casted, ok := structType.(*FirmataMessageDigitalIO); ok {
		return *casted
	}
	return nil
}

func (m *_FirmataMessageDigitalIO) GetTypeName() string {
	return "FirmataMessageDigitalIO"
}

func (m *_FirmataMessageDigitalIO) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_FirmataMessageDigitalIO) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (pinBlock)
	lengthInBits += 4

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_FirmataMessageDigitalIO) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func FirmataMessageDigitalIOParse(readBuffer utils.ReadBuffer, response bool) (FirmataMessageDigitalIO, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FirmataMessageDigitalIO"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FirmataMessageDigitalIO")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (pinBlock)
	_pinBlock, _pinBlockErr := readBuffer.ReadUint8("pinBlock", 4)
	if _pinBlockErr != nil {
		return nil, errors.Wrap(_pinBlockErr, "Error parsing 'pinBlock' field")
	}
	pinBlock := _pinBlock

	// Array field (data)
	if pullErr := readBuffer.PullContext("data", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for data")
	}
	// Count array
	data := make([]int8, uint16(2))
	// This happens when the size is set conditional to 0
	if len(data) == 0 {
		data = nil
	}
	{
		for curItem := uint16(0); curItem < uint16(uint16(2)); curItem++ {
			_item, _err := readBuffer.ReadInt8("", 8)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'data' field")
			}
			data[curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("data", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for data")
	}

	if closeErr := readBuffer.CloseContext("FirmataMessageDigitalIO"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FirmataMessageDigitalIO")
	}

	// Create a partially initialized instance
	_child := &_FirmataMessageDigitalIO{
		PinBlock:        pinBlock,
		Data:            data,
		_FirmataMessage: &_FirmataMessage{},
	}
	_child._FirmataMessage._FirmataMessageChildRequirements = _child
	return _child, nil
}

func (m *_FirmataMessageDigitalIO) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FirmataMessageDigitalIO"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FirmataMessageDigitalIO")
		}

		// Simple Field (pinBlock)
		pinBlock := uint8(m.GetPinBlock())
		_pinBlockErr := writeBuffer.WriteUint8("pinBlock", 4, (pinBlock))
		if _pinBlockErr != nil {
			return errors.Wrap(_pinBlockErr, "Error serializing 'pinBlock' field")
		}

		// Array Field (data)
		if pushErr := writeBuffer.PushContext("data", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for data")
		}
		for _, _element := range m.GetData() {
			_elementErr := writeBuffer.WriteInt8("", 8, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'data' field")
			}
		}
		if popErr := writeBuffer.PopContext("data", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for data")
		}

		if popErr := writeBuffer.PopContext("FirmataMessageDigitalIO"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FirmataMessageDigitalIO")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_FirmataMessageDigitalIO) isFirmataMessageDigitalIO() bool {
	return true
}

func (m *_FirmataMessageDigitalIO) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
