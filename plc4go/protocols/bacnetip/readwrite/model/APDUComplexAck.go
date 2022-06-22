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
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// APDUComplexAck is the corresponding interface of APDUComplexAck
type APDUComplexAck interface {
	utils.LengthAware
	utils.Serializable
	APDU
	// GetSegmentedMessage returns SegmentedMessage (property field)
	GetSegmentedMessage() bool
	// GetMoreFollows returns MoreFollows (property field)
	GetMoreFollows() bool
	// GetOriginalInvokeId returns OriginalInvokeId (property field)
	GetOriginalInvokeId() uint8
	// GetSequenceNumber returns SequenceNumber (property field)
	GetSequenceNumber() *uint8
	// GetProposedWindowSize returns ProposedWindowSize (property field)
	GetProposedWindowSize() *uint8
	// GetServiceAck returns ServiceAck (property field)
	GetServiceAck() BACnetServiceAck
	// GetSegmentServiceChoice returns SegmentServiceChoice (property field)
	GetSegmentServiceChoice() *uint8
	// GetSegment returns Segment (property field)
	GetSegment() []byte
	// GetApduHeaderReduction returns ApduHeaderReduction (virtual field)
	GetApduHeaderReduction() uint16
	// GetSegmentReduction returns SegmentReduction (virtual field)
	GetSegmentReduction() uint16
}

// APDUComplexAckExactly can be used when we want exactly this type and not a type which fulfills APDUComplexAck.
// This is useful for switch cases.
type APDUComplexAckExactly interface {
	APDUComplexAck
	isAPDUComplexAck() bool
}

// _APDUComplexAck is the data-structure of this message
type _APDUComplexAck struct {
	*_APDU
	SegmentedMessage     bool
	MoreFollows          bool
	OriginalInvokeId     uint8
	SequenceNumber       *uint8
	ProposedWindowSize   *uint8
	ServiceAck           BACnetServiceAck
	SegmentServiceChoice *uint8
	Segment              []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_APDUComplexAck) GetApduType() ApduType {
	return ApduType_COMPLEX_ACK_PDU
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_APDUComplexAck) InitializeParent(parent APDU) {}

func (m *_APDUComplexAck) GetParent() APDU {
	return m._APDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_APDUComplexAck) GetSegmentedMessage() bool {
	return m.SegmentedMessage
}

func (m *_APDUComplexAck) GetMoreFollows() bool {
	return m.MoreFollows
}

func (m *_APDUComplexAck) GetOriginalInvokeId() uint8 {
	return m.OriginalInvokeId
}

func (m *_APDUComplexAck) GetSequenceNumber() *uint8 {
	return m.SequenceNumber
}

func (m *_APDUComplexAck) GetProposedWindowSize() *uint8 {
	return m.ProposedWindowSize
}

func (m *_APDUComplexAck) GetServiceAck() BACnetServiceAck {
	return m.ServiceAck
}

func (m *_APDUComplexAck) GetSegmentServiceChoice() *uint8 {
	return m.SegmentServiceChoice
}

func (m *_APDUComplexAck) GetSegment() []byte {
	return m.Segment
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_APDUComplexAck) GetApduHeaderReduction() uint16 {
	sequenceNumber := m.SequenceNumber
	_ = sequenceNumber
	proposedWindowSize := m.ProposedWindowSize
	_ = proposedWindowSize
	serviceAck := m.ServiceAck
	_ = serviceAck
	segmentServiceChoice := m.SegmentServiceChoice
	_ = segmentServiceChoice
	return uint16(uint16(uint16(2)) + uint16(uint16(utils.InlineIf(m.GetSegmentedMessage(), func() interface{} { return uint16(uint16(2)) }, func() interface{} { return uint16(uint16(0)) }).(uint16))))
}

func (m *_APDUComplexAck) GetSegmentReduction() uint16 {
	sequenceNumber := m.SequenceNumber
	_ = sequenceNumber
	proposedWindowSize := m.ProposedWindowSize
	_ = proposedWindowSize
	serviceAck := m.ServiceAck
	_ = serviceAck
	segmentServiceChoice := m.SegmentServiceChoice
	_ = segmentServiceChoice
	return uint16(utils.InlineIf(bool(bool((m.GetSegmentServiceChoice()) != (nil))), func() interface{} { return uint16(uint16(uint16(m.GetApduHeaderReduction()) + uint16(uint16(1)))) }, func() interface{} { return uint16(m.GetApduHeaderReduction()) }).(uint16))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAPDUComplexAck factory function for _APDUComplexAck
func NewAPDUComplexAck(segmentedMessage bool, moreFollows bool, originalInvokeId uint8, sequenceNumber *uint8, proposedWindowSize *uint8, serviceAck BACnetServiceAck, segmentServiceChoice *uint8, segment []byte, apduLength uint16) *_APDUComplexAck {
	_result := &_APDUComplexAck{
		SegmentedMessage:     segmentedMessage,
		MoreFollows:          moreFollows,
		OriginalInvokeId:     originalInvokeId,
		SequenceNumber:       sequenceNumber,
		ProposedWindowSize:   proposedWindowSize,
		ServiceAck:           serviceAck,
		SegmentServiceChoice: segmentServiceChoice,
		Segment:              segment,
		_APDU:                NewAPDU(apduLength),
	}
	_result._APDU._APDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAPDUComplexAck(structType interface{}) APDUComplexAck {
	if casted, ok := structType.(APDUComplexAck); ok {
		return casted
	}
	if casted, ok := structType.(*APDUComplexAck); ok {
		return *casted
	}
	return nil
}

func (m *_APDUComplexAck) GetTypeName() string {
	return "APDUComplexAck"
}

func (m *_APDUComplexAck) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_APDUComplexAck) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (segmentedMessage)
	lengthInBits += 1

	// Simple field (moreFollows)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 2

	// Simple field (originalInvokeId)
	lengthInBits += 8

	// Optional Field (sequenceNumber)
	if m.SequenceNumber != nil {
		lengthInBits += 8
	}

	// Optional Field (proposedWindowSize)
	if m.ProposedWindowSize != nil {
		lengthInBits += 8
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (serviceAck)
	if m.ServiceAck != nil {
		lengthInBits += m.ServiceAck.GetLengthInBits()
	}

	// Optional Field (segmentServiceChoice)
	if m.SegmentServiceChoice != nil {
		lengthInBits += 8
	}

	// A virtual field doesn't have any in- or output.

	// Array field
	if len(m.Segment) > 0 {
		lengthInBits += 8 * uint16(len(m.Segment))
	}

	return lengthInBits
}

func (m *_APDUComplexAck) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func APDUComplexAckParse(readBuffer utils.ReadBuffer, apduLength uint16) (APDUComplexAck, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("APDUComplexAck"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for APDUComplexAck")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (segmentedMessage)
	_segmentedMessage, _segmentedMessageErr := readBuffer.ReadBit("segmentedMessage")
	if _segmentedMessageErr != nil {
		return nil, errors.Wrap(_segmentedMessageErr, "Error parsing 'segmentedMessage' field")
	}
	segmentedMessage := _segmentedMessage

	// Simple Field (moreFollows)
	_moreFollows, _moreFollowsErr := readBuffer.ReadBit("moreFollows")
	if _moreFollowsErr != nil {
		return nil, errors.Wrap(_moreFollowsErr, "Error parsing 'moreFollows' field")
	}
	moreFollows := _moreFollows

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 2)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (originalInvokeId)
	_originalInvokeId, _originalInvokeIdErr := readBuffer.ReadUint8("originalInvokeId", 8)
	if _originalInvokeIdErr != nil {
		return nil, errors.Wrap(_originalInvokeIdErr, "Error parsing 'originalInvokeId' field")
	}
	originalInvokeId := _originalInvokeId

	// Optional Field (sequenceNumber) (Can be skipped, if a given expression evaluates to false)
	var sequenceNumber *uint8 = nil
	if segmentedMessage {
		_val, _err := readBuffer.ReadUint8("sequenceNumber", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'sequenceNumber' field")
		}
		sequenceNumber = &_val
	}

	// Optional Field (proposedWindowSize) (Can be skipped, if a given expression evaluates to false)
	var proposedWindowSize *uint8 = nil
	if segmentedMessage {
		_val, _err := readBuffer.ReadUint8("proposedWindowSize", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'proposedWindowSize' field")
		}
		proposedWindowSize = &_val
	}

	// Virtual field
	_apduHeaderReduction := uint16(uint16(2)) + uint16(uint16(utils.InlineIf(segmentedMessage, func() interface{} { return uint16(uint16(2)) }, func() interface{} { return uint16(uint16(0)) }).(uint16)))
	apduHeaderReduction := uint16(_apduHeaderReduction)
	_ = apduHeaderReduction

	// Optional Field (serviceAck) (Can be skipped, if a given expression evaluates to false)
	var serviceAck BACnetServiceAck = nil
	if !(segmentedMessage) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("serviceAck"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for serviceAck")
		}
		_val, _err := BACnetServiceAckParse(readBuffer, uint16(apduLength)-uint16(apduHeaderReduction))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'serviceAck' field")
		default:
			serviceAck = _val.(BACnetServiceAck)
			if closeErr := readBuffer.CloseContext("serviceAck"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for serviceAck")
			}
		}
	}

	// Validation
	if !(bool(bool(bool(!(segmentedMessage)) && bool(bool((serviceAck) != (nil))))) || bool(segmentedMessage)) {
		return nil, errors.WithStack(utils.ParseValidationError{"service ack should be set"})
	}

	// Optional Field (segmentServiceChoice) (Can be skipped, if a given expression evaluates to false)
	var segmentServiceChoice *uint8 = nil
	if bool(segmentedMessage) && bool(bool((*sequenceNumber) != (0))) {
		_val, _err := readBuffer.ReadUint8("segmentServiceChoice", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'segmentServiceChoice' field")
		}
		segmentServiceChoice = &_val
	}

	// Virtual field
	_segmentReduction := utils.InlineIf(bool(bool((segmentServiceChoice) != (nil))), func() interface{} { return uint16(uint16(uint16(apduHeaderReduction) + uint16(uint16(1)))) }, func() interface{} { return uint16(apduHeaderReduction) }).(uint16)
	segmentReduction := uint16(_segmentReduction)
	_ = segmentReduction
	// Byte Array field (segment)
	numberOfBytessegment := int(utils.InlineIf(segmentedMessage, func() interface{} {
		return uint16(uint16(utils.InlineIf(bool(bool((apduLength) > (0))), func() interface{} { return uint16(uint16(uint16(apduLength) - uint16(segmentReduction))) }, func() interface{} { return uint16(uint16(0)) }).(uint16)))
	}, func() interface{} { return uint16(uint16(0)) }).(uint16))
	segment, _readArrayErr := readBuffer.ReadByteArray("segment", numberOfBytessegment)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'segment' field")
	}

	if closeErr := readBuffer.CloseContext("APDUComplexAck"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for APDUComplexAck")
	}

	// Create a partially initialized instance
	_child := &_APDUComplexAck{
		SegmentedMessage:     segmentedMessage,
		MoreFollows:          moreFollows,
		OriginalInvokeId:     originalInvokeId,
		SequenceNumber:       sequenceNumber,
		ProposedWindowSize:   proposedWindowSize,
		ServiceAck:           serviceAck,
		SegmentServiceChoice: segmentServiceChoice,
		Segment:              segment,
		_APDU:                &_APDU{},
	}
	_child._APDU._APDUChildRequirements = _child
	return _child, nil
}

func (m *_APDUComplexAck) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUComplexAck"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for APDUComplexAck")
		}

		// Simple Field (segmentedMessage)
		segmentedMessage := bool(m.GetSegmentedMessage())
		_segmentedMessageErr := writeBuffer.WriteBit("segmentedMessage", (segmentedMessage))
		if _segmentedMessageErr != nil {
			return errors.Wrap(_segmentedMessageErr, "Error serializing 'segmentedMessage' field")
		}

		// Simple Field (moreFollows)
		moreFollows := bool(m.GetMoreFollows())
		_moreFollowsErr := writeBuffer.WriteBit("moreFollows", (moreFollows))
		if _moreFollowsErr != nil {
			return errors.Wrap(_moreFollowsErr, "Error serializing 'moreFollows' field")
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 2, uint8(0))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (originalInvokeId)
		originalInvokeId := uint8(m.GetOriginalInvokeId())
		_originalInvokeIdErr := writeBuffer.WriteUint8("originalInvokeId", 8, (originalInvokeId))
		if _originalInvokeIdErr != nil {
			return errors.Wrap(_originalInvokeIdErr, "Error serializing 'originalInvokeId' field")
		}

		// Optional Field (sequenceNumber) (Can be skipped, if the value is null)
		var sequenceNumber *uint8 = nil
		if m.GetSequenceNumber() != nil {
			sequenceNumber = m.GetSequenceNumber()
			_sequenceNumberErr := writeBuffer.WriteUint8("sequenceNumber", 8, *(sequenceNumber))
			if _sequenceNumberErr != nil {
				return errors.Wrap(_sequenceNumberErr, "Error serializing 'sequenceNumber' field")
			}
		}

		// Optional Field (proposedWindowSize) (Can be skipped, if the value is null)
		var proposedWindowSize *uint8 = nil
		if m.GetProposedWindowSize() != nil {
			proposedWindowSize = m.GetProposedWindowSize()
			_proposedWindowSizeErr := writeBuffer.WriteUint8("proposedWindowSize", 8, *(proposedWindowSize))
			if _proposedWindowSizeErr != nil {
				return errors.Wrap(_proposedWindowSizeErr, "Error serializing 'proposedWindowSize' field")
			}
		}
		// Virtual field
		if _apduHeaderReductionErr := writeBuffer.WriteVirtual("apduHeaderReduction", m.GetApduHeaderReduction()); _apduHeaderReductionErr != nil {
			return errors.Wrap(_apduHeaderReductionErr, "Error serializing 'apduHeaderReduction' field")
		}

		// Optional Field (serviceAck) (Can be skipped, if the value is null)
		var serviceAck BACnetServiceAck = nil
		if m.GetServiceAck() != nil {
			if pushErr := writeBuffer.PushContext("serviceAck"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for serviceAck")
			}
			serviceAck = m.GetServiceAck()
			_serviceAckErr := writeBuffer.WriteSerializable(serviceAck)
			if popErr := writeBuffer.PopContext("serviceAck"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for serviceAck")
			}
			if _serviceAckErr != nil {
				return errors.Wrap(_serviceAckErr, "Error serializing 'serviceAck' field")
			}
		}

		// Optional Field (segmentServiceChoice) (Can be skipped, if the value is null)
		var segmentServiceChoice *uint8 = nil
		if m.GetSegmentServiceChoice() != nil {
			segmentServiceChoice = m.GetSegmentServiceChoice()
			_segmentServiceChoiceErr := writeBuffer.WriteUint8("segmentServiceChoice", 8, *(segmentServiceChoice))
			if _segmentServiceChoiceErr != nil {
				return errors.Wrap(_segmentServiceChoiceErr, "Error serializing 'segmentServiceChoice' field")
			}
		}
		// Virtual field
		if _segmentReductionErr := writeBuffer.WriteVirtual("segmentReduction", m.GetSegmentReduction()); _segmentReductionErr != nil {
			return errors.Wrap(_segmentReductionErr, "Error serializing 'segmentReduction' field")
		}

		// Array Field (segment)
		// Byte Array field (segment)
		if err := writeBuffer.WriteByteArray("segment", m.GetSegment()); err != nil {
			return errors.Wrap(err, "Error serializing 'segment' field")
		}

		if popErr := writeBuffer.PopContext("APDUComplexAck"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for APDUComplexAck")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_APDUComplexAck) isAPDUComplexAck() bool {
	return true
}

func (m *_APDUComplexAck) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
