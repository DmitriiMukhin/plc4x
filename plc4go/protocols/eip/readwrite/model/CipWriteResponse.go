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
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// CipWriteResponse is the corresponding interface of CipWriteResponse
type CipWriteResponse interface {
	utils.LengthAware
	utils.Serializable
	CipService
	// GetStatus returns Status (property field)
	GetStatus() uint8
	// GetExtStatus returns ExtStatus (property field)
	GetExtStatus() uint8
}

// CipWriteResponseExactly can be used when we want exactly this type and not a type which fulfills CipWriteResponse.
// This is useful for switch cases.
type CipWriteResponseExactly interface {
	CipWriteResponse
	isCipWriteResponse() bool
}

// _CipWriteResponse is the data-structure of this message
type _CipWriteResponse struct {
	*_CipService
	Status    uint8
	ExtStatus uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CipWriteResponse) GetService() uint8 {
	return 0xCD
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CipWriteResponse) InitializeParent(parent CipService) {}

func (m *_CipWriteResponse) GetParent() CipService {
	return m._CipService
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CipWriteResponse) GetStatus() uint8 {
	return m.Status
}

func (m *_CipWriteResponse) GetExtStatus() uint8 {
	return m.ExtStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCipWriteResponse factory function for _CipWriteResponse
func NewCipWriteResponse(status uint8, extStatus uint8, serviceLen uint16) *_CipWriteResponse {
	_result := &_CipWriteResponse{
		Status:      status,
		ExtStatus:   extStatus,
		_CipService: NewCipService(serviceLen),
	}
	_result._CipService._CipServiceChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCipWriteResponse(structType interface{}) CipWriteResponse {
	if casted, ok := structType.(CipWriteResponse); ok {
		return casted
	}
	if casted, ok := structType.(*CipWriteResponse); ok {
		return *casted
	}
	return nil
}

func (m *_CipWriteResponse) GetTypeName() string {
	return "CipWriteResponse"
}

func (m *_CipWriteResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CipWriteResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	// Simple field (extStatus)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CipWriteResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CipWriteResponseParse(readBuffer utils.ReadBuffer, serviceLen uint16) (CipWriteResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CipWriteResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CipWriteResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (status)
	_status, _statusErr := readBuffer.ReadUint8("status", 8)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field")
	}
	status := _status

	// Simple Field (extStatus)
	_extStatus, _extStatusErr := readBuffer.ReadUint8("extStatus", 8)
	if _extStatusErr != nil {
		return nil, errors.Wrap(_extStatusErr, "Error parsing 'extStatus' field")
	}
	extStatus := _extStatus

	if closeErr := readBuffer.CloseContext("CipWriteResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CipWriteResponse")
	}

	// Create a partially initialized instance
	_child := &_CipWriteResponse{
		Status:      status,
		ExtStatus:   extStatus,
		_CipService: &_CipService{},
	}
	_child._CipService._CipServiceChildRequirements = _child
	return _child, nil
}

func (m *_CipWriteResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipWriteResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CipWriteResponse")
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 8, uint8(0x00))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (status)
		status := uint8(m.GetStatus())
		_statusErr := writeBuffer.WriteUint8("status", 8, (status))
		if _statusErr != nil {
			return errors.Wrap(_statusErr, "Error serializing 'status' field")
		}

		// Simple Field (extStatus)
		extStatus := uint8(m.GetExtStatus())
		_extStatusErr := writeBuffer.WriteUint8("extStatus", 8, (extStatus))
		if _extStatusErr != nil {
			return errors.Wrap(_extStatusErr, "Error serializing 'extStatus' field")
		}

		if popErr := writeBuffer.PopContext("CipWriteResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CipWriteResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_CipWriteResponse) isCipWriteResponse() bool {
	return true
}

func (m *_CipWriteResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
