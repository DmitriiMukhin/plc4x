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

// ApduDataExtWriteRouterStatusRequest is the corresponding interface of ApduDataExtWriteRouterStatusRequest
type ApduDataExtWriteRouterStatusRequest interface {
	utils.LengthAware
	utils.Serializable
	ApduDataExt
}

// ApduDataExtWriteRouterStatusRequestExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtWriteRouterStatusRequest.
// This is useful for switch cases.
type ApduDataExtWriteRouterStatusRequestExactly interface {
	ApduDataExtWriteRouterStatusRequest
	isApduDataExtWriteRouterStatusRequest() bool
}

// _ApduDataExtWriteRouterStatusRequest is the data-structure of this message
type _ApduDataExtWriteRouterStatusRequest struct {
	*_ApduDataExt
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtWriteRouterStatusRequest) GetExtApciType() uint8 {
	return 0x0F
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtWriteRouterStatusRequest) InitializeParent(parent ApduDataExt) {}

func (m *_ApduDataExtWriteRouterStatusRequest) GetParent() ApduDataExt {
	return m._ApduDataExt
}

// NewApduDataExtWriteRouterStatusRequest factory function for _ApduDataExtWriteRouterStatusRequest
func NewApduDataExtWriteRouterStatusRequest(length uint8) *_ApduDataExtWriteRouterStatusRequest {
	_result := &_ApduDataExtWriteRouterStatusRequest{
		_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtWriteRouterStatusRequest(structType interface{}) ApduDataExtWriteRouterStatusRequest {
	if casted, ok := structType.(ApduDataExtWriteRouterStatusRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtWriteRouterStatusRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtWriteRouterStatusRequest) GetTypeName() string {
	return "ApduDataExtWriteRouterStatusRequest"
}

func (m *_ApduDataExtWriteRouterStatusRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ApduDataExtWriteRouterStatusRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_ApduDataExtWriteRouterStatusRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtWriteRouterStatusRequestParse(readBuffer utils.ReadBuffer, length uint8) (ApduDataExtWriteRouterStatusRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtWriteRouterStatusRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtWriteRouterStatusRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtWriteRouterStatusRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtWriteRouterStatusRequest")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtWriteRouterStatusRequest{
		_ApduDataExt: &_ApduDataExt{},
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtWriteRouterStatusRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtWriteRouterStatusRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtWriteRouterStatusRequest")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtWriteRouterStatusRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtWriteRouterStatusRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ApduDataExtWriteRouterStatusRequest) isApduDataExtWriteRouterStatusRequest() bool {
	return true
}

func (m *_ApduDataExtWriteRouterStatusRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
