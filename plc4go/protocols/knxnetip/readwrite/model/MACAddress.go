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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// MACAddress is the corresponding interface of MACAddress
type MACAddress interface {
	utils.LengthAware
	utils.Serializable
	// GetAddr returns Addr (property field)
	GetAddr() []byte
}

// MACAddressExactly can be used when we want exactly this type and not a type which fulfills MACAddress.
// This is useful for switch cases.
type MACAddressExactly interface {
	MACAddress
	isMACAddress() bool
}

// _MACAddress is the data-structure of this message
type _MACAddress struct {
	Addr []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MACAddress) GetAddr() []byte {
	return m.Addr
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMACAddress factory function for _MACAddress
func NewMACAddress(addr []byte) *_MACAddress {
	return &_MACAddress{Addr: addr}
}

// Deprecated: use the interface for direct cast
func CastMACAddress(structType interface{}) MACAddress {
	if casted, ok := structType.(MACAddress); ok {
		return casted
	}
	if casted, ok := structType.(*MACAddress); ok {
		return *casted
	}
	return nil
}

func (m *_MACAddress) GetTypeName() string {
	return "MACAddress"
}

func (m *_MACAddress) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Array field
	if len(m.Addr) > 0 {
		lengthInBits += 8 * uint16(len(m.Addr))
	}

	return lengthInBits
}

func (m *_MACAddress) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MACAddressParse(theBytes []byte) (MACAddress, error) {
	return MACAddressParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func MACAddressParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (MACAddress, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MACAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MACAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	// Byte Array field (addr)
	numberOfBytesaddr := int(uint16(6))
	addr, _readArrayErr := readBuffer.ReadByteArray("addr", numberOfBytesaddr)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'addr' field of MACAddress")
	}

	if closeErr := readBuffer.CloseContext("MACAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MACAddress")
	}

	// Create the instance
	return &_MACAddress{
		Addr: addr,
	}, nil
}

func (m *_MACAddress) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MACAddress) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("MACAddress"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for MACAddress")
	}

	// Array Field (addr)
	// Byte Array field (addr)
	if err := writeBuffer.WriteByteArray("addr", m.GetAddr()); err != nil {
		return errors.Wrap(err, "Error serializing 'addr' field")
	}

	if popErr := writeBuffer.PopContext("MACAddress"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for MACAddress")
	}
	return nil
}

func (m *_MACAddress) isMACAddress() bool {
	return true
}

func (m *_MACAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
