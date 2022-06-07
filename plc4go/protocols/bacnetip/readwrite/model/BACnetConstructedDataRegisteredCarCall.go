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
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataRegisteredCarCall is the data-structure of this message
type BACnetConstructedDataRegisteredCarCall struct {
	*BACnetConstructedData
	RegisteredCarCall []*BACnetLiftCarCallList

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataRegisteredCarCall is the corresponding interface of BACnetConstructedDataRegisteredCarCall
type IBACnetConstructedDataRegisteredCarCall interface {
	IBACnetConstructedData
	// GetRegisteredCarCall returns RegisteredCarCall (property field)
	GetRegisteredCarCall() []*BACnetLiftCarCallList
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

func (m *BACnetConstructedDataRegisteredCarCall) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataRegisteredCarCall) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_REGISTERED_CAR_CALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataRegisteredCarCall) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataRegisteredCarCall) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataRegisteredCarCall) GetRegisteredCarCall() []*BACnetLiftCarCallList {
	return m.RegisteredCarCall
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataRegisteredCarCall factory function for BACnetConstructedDataRegisteredCarCall
func NewBACnetConstructedDataRegisteredCarCall(registeredCarCall []*BACnetLiftCarCallList, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataRegisteredCarCall {
	_result := &BACnetConstructedDataRegisteredCarCall{
		RegisteredCarCall:     registeredCarCall,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataRegisteredCarCall(structType interface{}) *BACnetConstructedDataRegisteredCarCall {
	if casted, ok := structType.(BACnetConstructedDataRegisteredCarCall); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataRegisteredCarCall); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataRegisteredCarCall(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataRegisteredCarCall(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataRegisteredCarCall) GetTypeName() string {
	return "BACnetConstructedDataRegisteredCarCall"
}

func (m *BACnetConstructedDataRegisteredCarCall) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataRegisteredCarCall) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.RegisteredCarCall) > 0 {
		for _, element := range m.RegisteredCarCall {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConstructedDataRegisteredCarCall) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataRegisteredCarCallParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataRegisteredCarCall, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataRegisteredCarCall"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (registeredCarCall)
	if pullErr := readBuffer.PullContext("registeredCarCall", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Terminated array
	registeredCarCall := make([]*BACnetLiftCarCallList, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetLiftCarCallListParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'registeredCarCall' field")
			}
			registeredCarCall = append(registeredCarCall, CastBACnetLiftCarCallList(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("registeredCarCall", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataRegisteredCarCall"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataRegisteredCarCall{
		RegisteredCarCall:     registeredCarCall,
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataRegisteredCarCall) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataRegisteredCarCall"); pushErr != nil {
			return pushErr
		}

		// Array Field (registeredCarCall)
		if m.RegisteredCarCall != nil {
			if pushErr := writeBuffer.PushContext("registeredCarCall", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.RegisteredCarCall {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'registeredCarCall' field")
				}
			}
			if popErr := writeBuffer.PopContext("registeredCarCall", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataRegisteredCarCall"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataRegisteredCarCall) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}