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

// IdentifyReplyCommandDelays is the data-structure of this message
type IdentifyReplyCommandDelays struct {
	*IdentifyReplyCommand
}

// IIdentifyReplyCommandDelays is the corresponding interface of IdentifyReplyCommandDelays
type IIdentifyReplyCommandDelays interface {
	IIdentifyReplyCommand
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

func (m *IdentifyReplyCommandDelays) GetAttribute() Attribute {
	return Attribute_Delays
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *IdentifyReplyCommandDelays) InitializeParent(parent *IdentifyReplyCommand) {}

func (m *IdentifyReplyCommandDelays) GetParent() *IdentifyReplyCommand {
	return m.IdentifyReplyCommand
}

// NewIdentifyReplyCommandDelays factory function for IdentifyReplyCommandDelays
func NewIdentifyReplyCommandDelays() *IdentifyReplyCommandDelays {
	_result := &IdentifyReplyCommandDelays{
		IdentifyReplyCommand: NewIdentifyReplyCommand(),
	}
	_result.Child = _result
	return _result
}

func CastIdentifyReplyCommandDelays(structType interface{}) *IdentifyReplyCommandDelays {
	if casted, ok := structType.(IdentifyReplyCommandDelays); ok {
		return &casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandDelays); ok {
		return casted
	}
	if casted, ok := structType.(IdentifyReplyCommand); ok {
		return CastIdentifyReplyCommandDelays(casted.Child)
	}
	if casted, ok := structType.(*IdentifyReplyCommand); ok {
		return CastIdentifyReplyCommandDelays(casted.Child)
	}
	return nil
}

func (m *IdentifyReplyCommandDelays) GetTypeName() string {
	return "IdentifyReplyCommandDelays"
}

func (m *IdentifyReplyCommandDelays) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *IdentifyReplyCommandDelays) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *IdentifyReplyCommandDelays) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func IdentifyReplyCommandDelaysParse(readBuffer utils.ReadBuffer, attribute Attribute) (*IdentifyReplyCommandDelays, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandDelays"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandDelays"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &IdentifyReplyCommandDelays{
		IdentifyReplyCommand: &IdentifyReplyCommand{},
	}
	_child.IdentifyReplyCommand.Child = _child
	return _child, nil
}

func (m *IdentifyReplyCommandDelays) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandDelays"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandDelays"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *IdentifyReplyCommandDelays) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}