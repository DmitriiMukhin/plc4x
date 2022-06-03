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

// BACnetConstructedDataWeeklySchedule is the data-structure of this message
type BACnetConstructedDataWeeklySchedule struct {
	*BACnetConstructedData
	WeeklySchedule []*BACnetDailySchedule

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataWeeklySchedule is the corresponding interface of BACnetConstructedDataWeeklySchedule
type IBACnetConstructedDataWeeklySchedule interface {
	IBACnetConstructedData
	// GetWeeklySchedule returns WeeklySchedule (property field)
	GetWeeklySchedule() []*BACnetDailySchedule
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

func (m *BACnetConstructedDataWeeklySchedule) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataWeeklySchedule) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_WEEKLY_SCHEDULE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataWeeklySchedule) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataWeeklySchedule) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataWeeklySchedule) GetWeeklySchedule() []*BACnetDailySchedule {
	return m.WeeklySchedule
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataWeeklySchedule factory function for BACnetConstructedDataWeeklySchedule
func NewBACnetConstructedDataWeeklySchedule(weeklySchedule []*BACnetDailySchedule, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataWeeklySchedule {
	_result := &BACnetConstructedDataWeeklySchedule{
		WeeklySchedule:        weeklySchedule,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataWeeklySchedule(structType interface{}) *BACnetConstructedDataWeeklySchedule {
	if casted, ok := structType.(BACnetConstructedDataWeeklySchedule); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataWeeklySchedule); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataWeeklySchedule(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataWeeklySchedule(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataWeeklySchedule) GetTypeName() string {
	return "BACnetConstructedDataWeeklySchedule"
}

func (m *BACnetConstructedDataWeeklySchedule) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataWeeklySchedule) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.WeeklySchedule) > 0 {
		for _, element := range m.WeeklySchedule {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConstructedDataWeeklySchedule) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataWeeklyScheduleParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataWeeklySchedule, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataWeeklySchedule"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (weeklySchedule)
	if pullErr := readBuffer.PullContext("weeklySchedule", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Terminated array
	weeklySchedule := make([]*BACnetDailySchedule, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetDailyScheduleParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'weeklySchedule' field")
			}
			weeklySchedule = append(weeklySchedule, CastBACnetDailySchedule(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("weeklySchedule", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	// Validation
	if !(bool((len(weeklySchedule)) == (7))) {
		return nil, utils.ParseValidationError{"weeklySchedule should have exactly 7 values"}
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataWeeklySchedule"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataWeeklySchedule{
		WeeklySchedule:        weeklySchedule,
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataWeeklySchedule) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataWeeklySchedule"); pushErr != nil {
			return pushErr
		}

		// Array Field (weeklySchedule)
		if m.WeeklySchedule != nil {
			if pushErr := writeBuffer.PushContext("weeklySchedule", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.WeeklySchedule {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'weeklySchedule' field")
				}
			}
			if popErr := writeBuffer.PopContext("weeklySchedule", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataWeeklySchedule"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataWeeklySchedule) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}