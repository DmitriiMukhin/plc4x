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
	"context"
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// ModbusRtuADU is the corresponding interface of ModbusRtuADU
type ModbusRtuADU interface {
	utils.LengthAware
	utils.Serializable
	ModbusADU
	// GetAddress returns Address (property field)
	GetAddress() uint8
	// GetPdu returns Pdu (property field)
	GetPdu() ModbusPDU
}

// ModbusRtuADUExactly can be used when we want exactly this type and not a type which fulfills ModbusRtuADU.
// This is useful for switch cases.
type ModbusRtuADUExactly interface {
	ModbusRtuADU
	isModbusRtuADU() bool
}

// _ModbusRtuADU is the data-structure of this message
type _ModbusRtuADU struct {
	*_ModbusADU
	Address uint8
	Pdu     ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusRtuADU) GetDriverType() DriverType {
	return DriverType_MODBUS_RTU
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusRtuADU) InitializeParent(parent ModbusADU) {}

func (m *_ModbusRtuADU) GetParent() ModbusADU {
	return m._ModbusADU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusRtuADU) GetAddress() uint8 {
	return m.Address
}

func (m *_ModbusRtuADU) GetPdu() ModbusPDU {
	return m.Pdu
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusRtuADU factory function for _ModbusRtuADU
func NewModbusRtuADU(address uint8, pdu ModbusPDU, response bool) *_ModbusRtuADU {
	_result := &_ModbusRtuADU{
		Address:    address,
		Pdu:        pdu,
		_ModbusADU: NewModbusADU(response),
	}
	_result._ModbusADU._ModbusADUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusRtuADU(structType interface{}) ModbusRtuADU {
	if casted, ok := structType.(ModbusRtuADU); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusRtuADU); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusRtuADU) GetTypeName() string {
	return "ModbusRtuADU"
}

func (m *_ModbusRtuADU) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (address)
	lengthInBits += 8

	// Simple field (pdu)
	lengthInBits += m.Pdu.GetLengthInBits(ctx)

	// Checksum Field (checksum)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ModbusRtuADU) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusRtuADUParse(theBytes []byte, driverType DriverType, response bool) (ModbusRtuADU, error) {
	return ModbusRtuADUParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), driverType, response)
}

func ModbusRtuADUParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, driverType DriverType, response bool) (ModbusRtuADU, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusRtuADU"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusRtuADU")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (address)
	_address, _addressErr := readBuffer.ReadUint8("address", 8)
	if _addressErr != nil {
		return nil, errors.Wrap(_addressErr, "Error parsing 'address' field of ModbusRtuADU")
	}
	address := _address

	// Simple Field (pdu)
	if pullErr := readBuffer.PullContext("pdu"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for pdu")
	}
	_pdu, _pduErr := ModbusPDUParseWithBuffer(ctx, readBuffer, bool(response))
	if _pduErr != nil {
		return nil, errors.Wrap(_pduErr, "Error parsing 'pdu' field of ModbusRtuADU")
	}
	pdu := _pdu.(ModbusPDU)
	if closeErr := readBuffer.CloseContext("pdu"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for pdu")
	}

	// Checksum Field (checksum)
	{
		checksumRef, _checksumRefErr := readBuffer.ReadUint16("checksum", 16)
		if _checksumRefErr != nil {
			return nil, errors.Wrap(_checksumRefErr, "Error parsing 'checksum' field of ModbusRtuADU")
		}
		checksum, _checksumErr := RtuCrcCheck(address, pdu)
		if _checksumErr != nil {
			return nil, errors.Wrap(_checksumErr, "Checksum verification failed")
		}
		if checksum != checksumRef {
			return nil, errors.Errorf("Checksum verification failed. Expected %x but got %x", checksumRef, checksum)
		}
	}

	if closeErr := readBuffer.CloseContext("ModbusRtuADU"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusRtuADU")
	}

	// Create a partially initialized instance
	_child := &_ModbusRtuADU{
		_ModbusADU: &_ModbusADU{
			Response: response,
		},
		Address: address,
		Pdu:     pdu,
	}
	_child._ModbusADU._ModbusADUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusRtuADU) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusRtuADU) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusRtuADU"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusRtuADU")
		}

		// Simple Field (address)
		address := uint8(m.GetAddress())
		_addressErr := writeBuffer.WriteUint8("address", 8, (address))
		if _addressErr != nil {
			return errors.Wrap(_addressErr, "Error serializing 'address' field")
		}

		// Simple Field (pdu)
		if pushErr := writeBuffer.PushContext("pdu"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for pdu")
		}
		_pduErr := writeBuffer.WriteSerializable(ctx, m.GetPdu())
		if popErr := writeBuffer.PopContext("pdu"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for pdu")
		}
		if _pduErr != nil {
			return errors.Wrap(_pduErr, "Error serializing 'pdu' field")
		}

		// Checksum Field (checksum) (Calculated)
		{
			_checksum, _checksumErr := RtuCrcCheck(m.GetAddress(), m.GetPdu())
			if _checksumErr != nil {
				return errors.Wrap(_checksumErr, "Checksum calculation failed")
			}
			_checksumWriteErr := writeBuffer.WriteUint16("checksum", 16, (_checksum))
			if _checksumWriteErr != nil {
				return errors.Wrap(_checksumWriteErr, "Error serializing 'checksum' field")
			}
		}

		if popErr := writeBuffer.PopContext("ModbusRtuADU"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusRtuADU")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusRtuADU) isModbusRtuADU() bool {
	return true
}

func (m *_ModbusRtuADU) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
