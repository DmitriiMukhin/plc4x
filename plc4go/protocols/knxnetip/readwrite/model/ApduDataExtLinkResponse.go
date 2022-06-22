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

// ApduDataExtLinkResponse is the corresponding interface of ApduDataExtLinkResponse
type ApduDataExtLinkResponse interface {
	utils.LengthAware
	utils.Serializable
	ApduDataExt
}

// ApduDataExtLinkResponseExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtLinkResponse.
// This is useful for switch cases.
type ApduDataExtLinkResponseExactly interface {
	ApduDataExtLinkResponse
	isApduDataExtLinkResponse() bool
}

// _ApduDataExtLinkResponse is the data-structure of this message
type _ApduDataExtLinkResponse struct {
	*_ApduDataExt
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtLinkResponse) GetExtApciType() uint8 {
	return 0x26
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtLinkResponse) InitializeParent(parent ApduDataExt) {}

func (m *_ApduDataExtLinkResponse) GetParent() ApduDataExt {
	return m._ApduDataExt
}

// NewApduDataExtLinkResponse factory function for _ApduDataExtLinkResponse
func NewApduDataExtLinkResponse(length uint8) *_ApduDataExtLinkResponse {
	_result := &_ApduDataExtLinkResponse{
		_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtLinkResponse(structType interface{}) ApduDataExtLinkResponse {
	if casted, ok := structType.(ApduDataExtLinkResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtLinkResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtLinkResponse) GetTypeName() string {
	return "ApduDataExtLinkResponse"
}

func (m *_ApduDataExtLinkResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ApduDataExtLinkResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_ApduDataExtLinkResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtLinkResponseParse(readBuffer utils.ReadBuffer, length uint8) (ApduDataExtLinkResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtLinkResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtLinkResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtLinkResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtLinkResponse")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtLinkResponse{
		_ApduDataExt: &_ApduDataExt{},
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtLinkResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtLinkResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtLinkResponse")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtLinkResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtLinkResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ApduDataExtLinkResponse) isApduDataExtLinkResponse() bool {
	return true
}

func (m *_ApduDataExtLinkResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
