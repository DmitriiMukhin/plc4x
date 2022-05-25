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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse is the data-structure of this message
type S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse struct {
	*S7PayloadUserDataItem
}

// IS7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse is the corresponding interface of S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse
type IS7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse interface {
	IS7PayloadUserDataItem
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) GetCpuFunctionType() uint8 {
	return 0x08
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) GetCpuSubfunction() uint8 {
	return 0x02
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) GetDataLength() uint16 {
	return 0x00
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) InitializeParent(parent *S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize) {
	m.S7PayloadUserDataItem.ReturnCode = returnCode
	m.S7PayloadUserDataItem.TransportSize = transportSize
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) GetParent() *S7PayloadUserDataItem {
	return m.S7PayloadUserDataItem
}

// NewS7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse factory function for S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse
func NewS7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse(returnCode DataTransportErrorCode, transportSize DataTransportSize) *S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse {
	_result := &S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse{
		S7PayloadUserDataItem: NewS7PayloadUserDataItem(returnCode, transportSize),
	}
	_result.Child = _result
	return _result
}

func CastS7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse(structType interface{}) *S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse {
	if casted, ok := structType.(S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse); ok {
		return casted
	}
	if casted, ok := structType.(S7PayloadUserDataItem); ok {
		return CastS7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse(casted.Child)
	}
	if casted, ok := structType.(*S7PayloadUserDataItem); ok {
		return CastS7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse(casted.Child)
	}
	return nil
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) GetTypeName() string {
	return "S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse"
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseParse(readBuffer utils.ReadBuffer, cpuFunctionType uint8, cpuSubfunction uint8) (*S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse{
		S7PayloadUserDataItem: &S7PayloadUserDataItem{},
	}
	_child.S7PayloadUserDataItem.Child = _child
	return _child, nil
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}