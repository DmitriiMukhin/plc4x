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

// BACnetConstructedDataDatetimepatternValueAll is the corresponding interface of BACnetConstructedDataDatetimepatternValueAll
type BACnetConstructedDataDatetimepatternValueAll interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
}

// BACnetConstructedDataDatetimepatternValueAllExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDatetimepatternValueAll.
// This is useful for switch cases.
type BACnetConstructedDataDatetimepatternValueAllExactly interface {
	BACnetConstructedDataDatetimepatternValueAll
	isBACnetConstructedDataDatetimepatternValueAll() bool
}

// _BACnetConstructedDataDatetimepatternValueAll is the data-structure of this message
type _BACnetConstructedDataDatetimepatternValueAll struct {
	*_BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDatetimepatternValueAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_DATETIMEPATTERN_VALUE
}

func (m *_BACnetConstructedDataDatetimepatternValueAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDatetimepatternValueAll) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataDatetimepatternValueAll) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

// NewBACnetConstructedDataDatetimepatternValueAll factory function for _BACnetConstructedDataDatetimepatternValueAll
func NewBACnetConstructedDataDatetimepatternValueAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDatetimepatternValueAll {
	_result := &_BACnetConstructedDataDatetimepatternValueAll{
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDatetimepatternValueAll(structType interface{}) BACnetConstructedDataDatetimepatternValueAll {
	if casted, ok := structType.(BACnetConstructedDataDatetimepatternValueAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDatetimepatternValueAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDatetimepatternValueAll) GetTypeName() string {
	return "BACnetConstructedDataDatetimepatternValueAll"
}

func (m *_BACnetConstructedDataDatetimepatternValueAll) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataDatetimepatternValueAll) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_BACnetConstructedDataDatetimepatternValueAll) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDatetimepatternValueAllParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDatetimepatternValueAll, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDatetimepatternValueAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDatetimepatternValueAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{"All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDatetimepatternValueAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDatetimepatternValueAll")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataDatetimepatternValueAll{
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataDatetimepatternValueAll) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDatetimepatternValueAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDatetimepatternValueAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDatetimepatternValueAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDatetimepatternValueAll")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDatetimepatternValueAll) isBACnetConstructedDataDatetimepatternValueAll() bool {
	return true
}

func (m *_BACnetConstructedDataDatetimepatternValueAll) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
