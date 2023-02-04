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

// BACnetLandingCallStatusCommandDestination is the corresponding interface of BACnetLandingCallStatusCommandDestination
type BACnetLandingCallStatusCommandDestination interface {
	utils.LengthAware
	utils.Serializable
	BACnetLandingCallStatusCommand
	// GetDestination returns Destination (property field)
	GetDestination() BACnetContextTagUnsignedInteger
}

// BACnetLandingCallStatusCommandDestinationExactly can be used when we want exactly this type and not a type which fulfills BACnetLandingCallStatusCommandDestination.
// This is useful for switch cases.
type BACnetLandingCallStatusCommandDestinationExactly interface {
	BACnetLandingCallStatusCommandDestination
	isBACnetLandingCallStatusCommandDestination() bool
}

// _BACnetLandingCallStatusCommandDestination is the data-structure of this message
type _BACnetLandingCallStatusCommandDestination struct {
	*_BACnetLandingCallStatusCommand
	Destination BACnetContextTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetLandingCallStatusCommandDestination) InitializeParent(parent BACnetLandingCallStatusCommand, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetLandingCallStatusCommandDestination) GetParent() BACnetLandingCallStatusCommand {
	return m._BACnetLandingCallStatusCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLandingCallStatusCommandDestination) GetDestination() BACnetContextTagUnsignedInteger {
	return m.Destination
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLandingCallStatusCommandDestination factory function for _BACnetLandingCallStatusCommandDestination
func NewBACnetLandingCallStatusCommandDestination(destination BACnetContextTagUnsignedInteger, peekedTagHeader BACnetTagHeader) *_BACnetLandingCallStatusCommandDestination {
	_result := &_BACnetLandingCallStatusCommandDestination{
		Destination:                     destination,
		_BACnetLandingCallStatusCommand: NewBACnetLandingCallStatusCommand(peekedTagHeader),
	}
	_result._BACnetLandingCallStatusCommand._BACnetLandingCallStatusCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetLandingCallStatusCommandDestination(structType interface{}) BACnetLandingCallStatusCommandDestination {
	if casted, ok := structType.(BACnetLandingCallStatusCommandDestination); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLandingCallStatusCommandDestination); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLandingCallStatusCommandDestination) GetTypeName() string {
	return "BACnetLandingCallStatusCommandDestination"
}

func (m *_BACnetLandingCallStatusCommandDestination) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (destination)
	lengthInBits += m.Destination.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLandingCallStatusCommandDestination) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLandingCallStatusCommandDestinationParse(theBytes []byte) (BACnetLandingCallStatusCommandDestination, error) {
	return BACnetLandingCallStatusCommandDestinationParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetLandingCallStatusCommandDestinationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLandingCallStatusCommandDestination, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLandingCallStatusCommandDestination"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLandingCallStatusCommandDestination")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (destination)
	if pullErr := readBuffer.PullContext("destination"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for destination")
	}
	_destination, _destinationErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _destinationErr != nil {
		return nil, errors.Wrap(_destinationErr, "Error parsing 'destination' field of BACnetLandingCallStatusCommandDestination")
	}
	destination := _destination.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("destination"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for destination")
	}

	if closeErr := readBuffer.CloseContext("BACnetLandingCallStatusCommandDestination"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLandingCallStatusCommandDestination")
	}

	// Create a partially initialized instance
	_child := &_BACnetLandingCallStatusCommandDestination{
		_BACnetLandingCallStatusCommand: &_BACnetLandingCallStatusCommand{},
		Destination:                     destination,
	}
	_child._BACnetLandingCallStatusCommand._BACnetLandingCallStatusCommandChildRequirements = _child
	return _child, nil
}

func (m *_BACnetLandingCallStatusCommandDestination) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLandingCallStatusCommandDestination) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLandingCallStatusCommandDestination"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLandingCallStatusCommandDestination")
		}

		// Simple Field (destination)
		if pushErr := writeBuffer.PushContext("destination"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for destination")
		}
		_destinationErr := writeBuffer.WriteSerializable(ctx, m.GetDestination())
		if popErr := writeBuffer.PopContext("destination"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for destination")
		}
		if _destinationErr != nil {
			return errors.Wrap(_destinationErr, "Error serializing 'destination' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLandingCallStatusCommandDestination"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLandingCallStatusCommandDestination")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetLandingCallStatusCommandDestination) isBACnetLandingCallStatusCommandDestination() bool {
	return true
}

func (m *_BACnetLandingCallStatusCommandDestination) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
