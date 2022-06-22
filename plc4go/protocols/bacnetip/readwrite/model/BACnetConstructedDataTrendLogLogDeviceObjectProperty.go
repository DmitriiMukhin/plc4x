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

// BACnetConstructedDataTrendLogLogDeviceObjectProperty is the corresponding interface of BACnetConstructedDataTrendLogLogDeviceObjectProperty
type BACnetConstructedDataTrendLogLogDeviceObjectProperty interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLogDeviceObjectProperty returns LogDeviceObjectProperty (property field)
	GetLogDeviceObjectProperty() BACnetDeviceObjectPropertyReference
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDeviceObjectPropertyReference
}

// BACnetConstructedDataTrendLogLogDeviceObjectPropertyExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataTrendLogLogDeviceObjectProperty.
// This is useful for switch cases.
type BACnetConstructedDataTrendLogLogDeviceObjectPropertyExactly interface {
	BACnetConstructedDataTrendLogLogDeviceObjectProperty
	isBACnetConstructedDataTrendLogLogDeviceObjectProperty() bool
}

// _BACnetConstructedDataTrendLogLogDeviceObjectProperty is the data-structure of this message
type _BACnetConstructedDataTrendLogLogDeviceObjectProperty struct {
	*_BACnetConstructedData
	LogDeviceObjectProperty BACnetDeviceObjectPropertyReference
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_TREND_LOG
}

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOG_DEVICE_OBJECT_PROPERTY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetLogDeviceObjectProperty() BACnetDeviceObjectPropertyReference {
	return m.LogDeviceObjectProperty
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetActualValue() BACnetDeviceObjectPropertyReference {
	return CastBACnetDeviceObjectPropertyReference(m.GetLogDeviceObjectProperty())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTrendLogLogDeviceObjectProperty factory function for _BACnetConstructedDataTrendLogLogDeviceObjectProperty
func NewBACnetConstructedDataTrendLogLogDeviceObjectProperty(logDeviceObjectProperty BACnetDeviceObjectPropertyReference, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTrendLogLogDeviceObjectProperty {
	_result := &_BACnetConstructedDataTrendLogLogDeviceObjectProperty{
		LogDeviceObjectProperty: logDeviceObjectProperty,
		_BACnetConstructedData:  NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTrendLogLogDeviceObjectProperty(structType interface{}) BACnetConstructedDataTrendLogLogDeviceObjectProperty {
	if casted, ok := structType.(BACnetConstructedDataTrendLogLogDeviceObjectProperty); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTrendLogLogDeviceObjectProperty); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetTypeName() string {
	return "BACnetConstructedDataTrendLogLogDeviceObjectProperty"
}

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (logDeviceObjectProperty)
	lengthInBits += m.LogDeviceObjectProperty.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataTrendLogLogDeviceObjectPropertyParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTrendLogLogDeviceObjectProperty, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTrendLogLogDeviceObjectProperty"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTrendLogLogDeviceObjectProperty")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (logDeviceObjectProperty)
	if pullErr := readBuffer.PullContext("logDeviceObjectProperty"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for logDeviceObjectProperty")
	}
	_logDeviceObjectProperty, _logDeviceObjectPropertyErr := BACnetDeviceObjectPropertyReferenceParse(readBuffer)
	if _logDeviceObjectPropertyErr != nil {
		return nil, errors.Wrap(_logDeviceObjectPropertyErr, "Error parsing 'logDeviceObjectProperty' field")
	}
	logDeviceObjectProperty := _logDeviceObjectProperty.(BACnetDeviceObjectPropertyReference)
	if closeErr := readBuffer.CloseContext("logDeviceObjectProperty"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for logDeviceObjectProperty")
	}

	// Virtual field
	_actualValue := logDeviceObjectProperty
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTrendLogLogDeviceObjectProperty"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTrendLogLogDeviceObjectProperty")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataTrendLogLogDeviceObjectProperty{
		LogDeviceObjectProperty: logDeviceObjectProperty,
		_BACnetConstructedData:  &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTrendLogLogDeviceObjectProperty"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTrendLogLogDeviceObjectProperty")
		}

		// Simple Field (logDeviceObjectProperty)
		if pushErr := writeBuffer.PushContext("logDeviceObjectProperty"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for logDeviceObjectProperty")
		}
		_logDeviceObjectPropertyErr := writeBuffer.WriteSerializable(m.GetLogDeviceObjectProperty())
		if popErr := writeBuffer.PopContext("logDeviceObjectProperty"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for logDeviceObjectProperty")
		}
		if _logDeviceObjectPropertyErr != nil {
			return errors.Wrap(_logDeviceObjectPropertyErr, "Error serializing 'logDeviceObjectProperty' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTrendLogLogDeviceObjectProperty"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTrendLogLogDeviceObjectProperty")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) isBACnetConstructedDataTrendLogLogDeviceObjectProperty() bool {
	return true
}

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
