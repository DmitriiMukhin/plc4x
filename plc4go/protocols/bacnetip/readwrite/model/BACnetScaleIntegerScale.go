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

// BACnetScaleIntegerScale is the corresponding interface of BACnetScaleIntegerScale
type BACnetScaleIntegerScale interface {
	utils.LengthAware
	utils.Serializable
	BACnetScale
	// GetIntegerScale returns IntegerScale (property field)
	GetIntegerScale() BACnetContextTagSignedInteger
}

// BACnetScaleIntegerScaleExactly can be used when we want exactly this type and not a type which fulfills BACnetScaleIntegerScale.
// This is useful for switch cases.
type BACnetScaleIntegerScaleExactly interface {
	BACnetScaleIntegerScale
	isBACnetScaleIntegerScale() bool
}

// _BACnetScaleIntegerScale is the data-structure of this message
type _BACnetScaleIntegerScale struct {
	*_BACnetScale
	IntegerScale BACnetContextTagSignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetScaleIntegerScale) InitializeParent(parent BACnetScale, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetScaleIntegerScale) GetParent() BACnetScale {
	return m._BACnetScale
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetScaleIntegerScale) GetIntegerScale() BACnetContextTagSignedInteger {
	return m.IntegerScale
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetScaleIntegerScale factory function for _BACnetScaleIntegerScale
func NewBACnetScaleIntegerScale(integerScale BACnetContextTagSignedInteger, peekedTagHeader BACnetTagHeader) *_BACnetScaleIntegerScale {
	_result := &_BACnetScaleIntegerScale{
		IntegerScale: integerScale,
		_BACnetScale: NewBACnetScale(peekedTagHeader),
	}
	_result._BACnetScale._BACnetScaleChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetScaleIntegerScale(structType interface{}) BACnetScaleIntegerScale {
	if casted, ok := structType.(BACnetScaleIntegerScale); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetScaleIntegerScale); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetScaleIntegerScale) GetTypeName() string {
	return "BACnetScaleIntegerScale"
}

func (m *_BACnetScaleIntegerScale) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (integerScale)
	lengthInBits += m.IntegerScale.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetScaleIntegerScale) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetScaleIntegerScaleParse(theBytes []byte) (BACnetScaleIntegerScale, error) {
	return BACnetScaleIntegerScaleParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetScaleIntegerScaleParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetScaleIntegerScale, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetScaleIntegerScale"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetScaleIntegerScale")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (integerScale)
	if pullErr := readBuffer.PullContext("integerScale"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for integerScale")
	}
	_integerScale, _integerScaleErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_SIGNED_INTEGER))
	if _integerScaleErr != nil {
		return nil, errors.Wrap(_integerScaleErr, "Error parsing 'integerScale' field of BACnetScaleIntegerScale")
	}
	integerScale := _integerScale.(BACnetContextTagSignedInteger)
	if closeErr := readBuffer.CloseContext("integerScale"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for integerScale")
	}

	if closeErr := readBuffer.CloseContext("BACnetScaleIntegerScale"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetScaleIntegerScale")
	}

	// Create a partially initialized instance
	_child := &_BACnetScaleIntegerScale{
		_BACnetScale: &_BACnetScale{},
		IntegerScale: integerScale,
	}
	_child._BACnetScale._BACnetScaleChildRequirements = _child
	return _child, nil
}

func (m *_BACnetScaleIntegerScale) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetScaleIntegerScale) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetScaleIntegerScale"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetScaleIntegerScale")
		}

		// Simple Field (integerScale)
		if pushErr := writeBuffer.PushContext("integerScale"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for integerScale")
		}
		_integerScaleErr := writeBuffer.WriteSerializable(ctx, m.GetIntegerScale())
		if popErr := writeBuffer.PopContext("integerScale"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for integerScale")
		}
		if _integerScaleErr != nil {
			return errors.Wrap(_integerScaleErr, "Error serializing 'integerScale' field")
		}

		if popErr := writeBuffer.PopContext("BACnetScaleIntegerScale"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetScaleIntegerScale")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetScaleIntegerScale) isBACnetScaleIntegerScale() bool {
	return true
}

func (m *_BACnetScaleIntegerScale) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
