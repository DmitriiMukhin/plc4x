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

// ApduDataExtPropertyValueRead is the corresponding interface of ApduDataExtPropertyValueRead
type ApduDataExtPropertyValueRead interface {
	utils.LengthAware
	utils.Serializable
	ApduDataExt
	// GetObjectIndex returns ObjectIndex (property field)
	GetObjectIndex() uint8
	// GetPropertyId returns PropertyId (property field)
	GetPropertyId() uint8
	// GetCount returns Count (property field)
	GetCount() uint8
	// GetIndex returns Index (property field)
	GetIndex() uint16
}

// ApduDataExtPropertyValueReadExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtPropertyValueRead.
// This is useful for switch cases.
type ApduDataExtPropertyValueReadExactly interface {
	ApduDataExtPropertyValueRead
	isApduDataExtPropertyValueRead() bool
}

// _ApduDataExtPropertyValueRead is the data-structure of this message
type _ApduDataExtPropertyValueRead struct {
	*_ApduDataExt
	ObjectIndex uint8
	PropertyId  uint8
	Count       uint8
	Index       uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtPropertyValueRead) GetExtApciType() uint8 {
	return 0x15
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtPropertyValueRead) InitializeParent(parent ApduDataExt) {}

func (m *_ApduDataExtPropertyValueRead) GetParent() ApduDataExt {
	return m._ApduDataExt
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ApduDataExtPropertyValueRead) GetObjectIndex() uint8 {
	return m.ObjectIndex
}

func (m *_ApduDataExtPropertyValueRead) GetPropertyId() uint8 {
	return m.PropertyId
}

func (m *_ApduDataExtPropertyValueRead) GetCount() uint8 {
	return m.Count
}

func (m *_ApduDataExtPropertyValueRead) GetIndex() uint16 {
	return m.Index
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewApduDataExtPropertyValueRead factory function for _ApduDataExtPropertyValueRead
func NewApduDataExtPropertyValueRead(objectIndex uint8, propertyId uint8, count uint8, index uint16, length uint8) *_ApduDataExtPropertyValueRead {
	_result := &_ApduDataExtPropertyValueRead{
		ObjectIndex:  objectIndex,
		PropertyId:   propertyId,
		Count:        count,
		Index:        index,
		_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtPropertyValueRead(structType interface{}) ApduDataExtPropertyValueRead {
	if casted, ok := structType.(ApduDataExtPropertyValueRead); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtPropertyValueRead); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtPropertyValueRead) GetTypeName() string {
	return "ApduDataExtPropertyValueRead"
}

func (m *_ApduDataExtPropertyValueRead) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ApduDataExtPropertyValueRead) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (objectIndex)
	lengthInBits += 8

	// Simple field (propertyId)
	lengthInBits += 8

	// Simple field (count)
	lengthInBits += 4

	// Simple field (index)
	lengthInBits += 12

	return lengthInBits
}

func (m *_ApduDataExtPropertyValueRead) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtPropertyValueReadParse(readBuffer utils.ReadBuffer, length uint8) (ApduDataExtPropertyValueRead, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtPropertyValueRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtPropertyValueRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectIndex)
	_objectIndex, _objectIndexErr := readBuffer.ReadUint8("objectIndex", 8)
	if _objectIndexErr != nil {
		return nil, errors.Wrap(_objectIndexErr, "Error parsing 'objectIndex' field")
	}
	objectIndex := _objectIndex

	// Simple Field (propertyId)
	_propertyId, _propertyIdErr := readBuffer.ReadUint8("propertyId", 8)
	if _propertyIdErr != nil {
		return nil, errors.Wrap(_propertyIdErr, "Error parsing 'propertyId' field")
	}
	propertyId := _propertyId

	// Simple Field (count)
	_count, _countErr := readBuffer.ReadUint8("count", 4)
	if _countErr != nil {
		return nil, errors.Wrap(_countErr, "Error parsing 'count' field")
	}
	count := _count

	// Simple Field (index)
	_index, _indexErr := readBuffer.ReadUint16("index", 12)
	if _indexErr != nil {
		return nil, errors.Wrap(_indexErr, "Error parsing 'index' field")
	}
	index := _index

	if closeErr := readBuffer.CloseContext("ApduDataExtPropertyValueRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtPropertyValueRead")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtPropertyValueRead{
		ObjectIndex:  objectIndex,
		PropertyId:   propertyId,
		Count:        count,
		Index:        index,
		_ApduDataExt: &_ApduDataExt{},
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtPropertyValueRead) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtPropertyValueRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtPropertyValueRead")
		}

		// Simple Field (objectIndex)
		objectIndex := uint8(m.GetObjectIndex())
		_objectIndexErr := writeBuffer.WriteUint8("objectIndex", 8, (objectIndex))
		if _objectIndexErr != nil {
			return errors.Wrap(_objectIndexErr, "Error serializing 'objectIndex' field")
		}

		// Simple Field (propertyId)
		propertyId := uint8(m.GetPropertyId())
		_propertyIdErr := writeBuffer.WriteUint8("propertyId", 8, (propertyId))
		if _propertyIdErr != nil {
			return errors.Wrap(_propertyIdErr, "Error serializing 'propertyId' field")
		}

		// Simple Field (count)
		count := uint8(m.GetCount())
		_countErr := writeBuffer.WriteUint8("count", 4, (count))
		if _countErr != nil {
			return errors.Wrap(_countErr, "Error serializing 'count' field")
		}

		// Simple Field (index)
		index := uint16(m.GetIndex())
		_indexErr := writeBuffer.WriteUint16("index", 12, (index))
		if _indexErr != nil {
			return errors.Wrap(_indexErr, "Error serializing 'index' field")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtPropertyValueRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtPropertyValueRead")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ApduDataExtPropertyValueRead) isApduDataExtPropertyValueRead() bool {
	return true
}

func (m *_ApduDataExtPropertyValueRead) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
