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

// ApduDataExtDomainAddressRead is the corresponding interface of ApduDataExtDomainAddressRead
type ApduDataExtDomainAddressRead interface {
	utils.LengthAware
	utils.Serializable
	ApduDataExt
}

// ApduDataExtDomainAddressReadExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtDomainAddressRead.
// This is useful for switch cases.
type ApduDataExtDomainAddressReadExactly interface {
	ApduDataExtDomainAddressRead
	isApduDataExtDomainAddressRead() bool
}

// _ApduDataExtDomainAddressRead is the data-structure of this message
type _ApduDataExtDomainAddressRead struct {
	*_ApduDataExt
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtDomainAddressRead) GetExtApciType() uint8 {
	return 0x21
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtDomainAddressRead) InitializeParent(parent ApduDataExt) {}

func (m *_ApduDataExtDomainAddressRead) GetParent() ApduDataExt {
	return m._ApduDataExt
}

// NewApduDataExtDomainAddressRead factory function for _ApduDataExtDomainAddressRead
func NewApduDataExtDomainAddressRead(length uint8) *_ApduDataExtDomainAddressRead {
	_result := &_ApduDataExtDomainAddressRead{
		_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtDomainAddressRead(structType interface{}) ApduDataExtDomainAddressRead {
	if casted, ok := structType.(ApduDataExtDomainAddressRead); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtDomainAddressRead); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtDomainAddressRead) GetTypeName() string {
	return "ApduDataExtDomainAddressRead"
}

func (m *_ApduDataExtDomainAddressRead) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ApduDataExtDomainAddressRead) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_ApduDataExtDomainAddressRead) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtDomainAddressReadParse(readBuffer utils.ReadBuffer, length uint8) (ApduDataExtDomainAddressRead, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtDomainAddressRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtDomainAddressRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtDomainAddressRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtDomainAddressRead")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtDomainAddressRead{
		_ApduDataExt: &_ApduDataExt{},
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtDomainAddressRead) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtDomainAddressRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtDomainAddressRead")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtDomainAddressRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtDomainAddressRead")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ApduDataExtDomainAddressRead) isApduDataExtDomainAddressRead() bool {
	return true
}

func (m *_ApduDataExtDomainAddressRead) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
