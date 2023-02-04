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

// BACnetServicesSupportedTagged is the corresponding interface of BACnetServicesSupportedTagged
type BACnetServicesSupportedTagged interface {
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadBitString
	// GetWriteGroup returns WriteGroup (virtual field)
	GetWriteGroup() bool
	// GetSubscribeCovPropertyMultiple returns SubscribeCovPropertyMultiple (virtual field)
	GetSubscribeCovPropertyMultiple() bool
	// GetConfirmedCovNotificationMultiple returns ConfirmedCovNotificationMultiple (virtual field)
	GetConfirmedCovNotificationMultiple() bool
	// GetUnconfirmedCovNotificationMultiple returns UnconfirmedCovNotificationMultiple (virtual field)
	GetUnconfirmedCovNotificationMultiple() bool
	// GetWhoIs returns WhoIs (virtual field)
	GetWhoIs() bool
	// GetReadRange returns ReadRange (virtual field)
	GetReadRange() bool
	// GetUtcTimeSynchronization returns UtcTimeSynchronization (virtual field)
	GetUtcTimeSynchronization() bool
	// GetLifeSafetyOperation returns LifeSafetyOperation (virtual field)
	GetLifeSafetyOperation() bool
	// GetSubscribeCovProperty returns SubscribeCovProperty (virtual field)
	GetSubscribeCovProperty() bool
	// GetGetEventInformation returns GetEventInformation (virtual field)
	GetGetEventInformation() bool
}

// BACnetServicesSupportedTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetServicesSupportedTagged.
// This is useful for switch cases.
type BACnetServicesSupportedTaggedExactly interface {
	BACnetServicesSupportedTagged
	isBACnetServicesSupportedTagged() bool
}

// _BACnetServicesSupportedTagged is the data-structure of this message
type _BACnetServicesSupportedTagged struct {
	Header  BACnetTagHeader
	Payload BACnetTagPayloadBitString

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServicesSupportedTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetServicesSupportedTagged) GetPayload() BACnetTagPayloadBitString {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetServicesSupportedTagged) GetWriteGroup() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (0))), func() interface{} { return bool(m.GetPayload().GetData()[0]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

func (m *_BACnetServicesSupportedTagged) GetSubscribeCovPropertyMultiple() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (1))), func() interface{} { return bool(m.GetPayload().GetData()[1]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

func (m *_BACnetServicesSupportedTagged) GetConfirmedCovNotificationMultiple() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (2))), func() interface{} { return bool(m.GetPayload().GetData()[2]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

func (m *_BACnetServicesSupportedTagged) GetUnconfirmedCovNotificationMultiple() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (3))), func() interface{} { return bool(m.GetPayload().GetData()[3]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

func (m *_BACnetServicesSupportedTagged) GetWhoIs() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (4))), func() interface{} { return bool(m.GetPayload().GetData()[4]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

func (m *_BACnetServicesSupportedTagged) GetReadRange() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (5))), func() interface{} { return bool(m.GetPayload().GetData()[5]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

func (m *_BACnetServicesSupportedTagged) GetUtcTimeSynchronization() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (6))), func() interface{} { return bool(m.GetPayload().GetData()[6]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

func (m *_BACnetServicesSupportedTagged) GetLifeSafetyOperation() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (7))), func() interface{} { return bool(m.GetPayload().GetData()[7]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

func (m *_BACnetServicesSupportedTagged) GetSubscribeCovProperty() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (8))), func() interface{} { return bool(m.GetPayload().GetData()[8]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

func (m *_BACnetServicesSupportedTagged) GetGetEventInformation() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (9))), func() interface{} { return bool(m.GetPayload().GetData()[9]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServicesSupportedTagged factory function for _BACnetServicesSupportedTagged
func NewBACnetServicesSupportedTagged(header BACnetTagHeader, payload BACnetTagPayloadBitString, tagNumber uint8, tagClass TagClass) *_BACnetServicesSupportedTagged {
	return &_BACnetServicesSupportedTagged{Header: header, Payload: payload, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetServicesSupportedTagged(structType interface{}) BACnetServicesSupportedTagged {
	if casted, ok := structType.(BACnetServicesSupportedTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServicesSupportedTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServicesSupportedTagged) GetTypeName() string {
	return "BACnetServicesSupportedTagged"
}

func (m *_BACnetServicesSupportedTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetServicesSupportedTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetServicesSupportedTaggedParse(theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetServicesSupportedTagged, error) {
	return BACnetServicesSupportedTaggedParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetServicesSupportedTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetServicesSupportedTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServicesSupportedTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServicesSupportedTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
	_header, _headerErr := BACnetTagHeaderParseWithBuffer(ctx, readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field of BACnetServicesSupportedTagged")
	}
	header := _header.(BACnetTagHeader)
	if closeErr := readBuffer.CloseContext("header"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for header")
	}

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{"tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{"tagnumber doesn't match"})
	}

	// Simple Field (payload)
	if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for payload")
	}
	_payload, _payloadErr := BACnetTagPayloadBitStringParseWithBuffer(ctx, readBuffer, uint32(header.GetActualLength()))
	if _payloadErr != nil {
		return nil, errors.Wrap(_payloadErr, "Error parsing 'payload' field of BACnetServicesSupportedTagged")
	}
	payload := _payload.(BACnetTagPayloadBitString)
	if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for payload")
	}

	// Virtual field
	_writeGroup := utils.InlineIf((bool((len(payload.GetData())) > (0))), func() interface{} { return bool(payload.GetData()[0]) }, func() interface{} { return bool(bool(false)) }).(bool)
	writeGroup := bool(_writeGroup)
	_ = writeGroup

	// Virtual field
	_subscribeCovPropertyMultiple := utils.InlineIf((bool((len(payload.GetData())) > (1))), func() interface{} { return bool(payload.GetData()[1]) }, func() interface{} { return bool(bool(false)) }).(bool)
	subscribeCovPropertyMultiple := bool(_subscribeCovPropertyMultiple)
	_ = subscribeCovPropertyMultiple

	// Virtual field
	_confirmedCovNotificationMultiple := utils.InlineIf((bool((len(payload.GetData())) > (2))), func() interface{} { return bool(payload.GetData()[2]) }, func() interface{} { return bool(bool(false)) }).(bool)
	confirmedCovNotificationMultiple := bool(_confirmedCovNotificationMultiple)
	_ = confirmedCovNotificationMultiple

	// Virtual field
	_unconfirmedCovNotificationMultiple := utils.InlineIf((bool((len(payload.GetData())) > (3))), func() interface{} { return bool(payload.GetData()[3]) }, func() interface{} { return bool(bool(false)) }).(bool)
	unconfirmedCovNotificationMultiple := bool(_unconfirmedCovNotificationMultiple)
	_ = unconfirmedCovNotificationMultiple

	// Virtual field
	_whoIs := utils.InlineIf((bool((len(payload.GetData())) > (4))), func() interface{} { return bool(payload.GetData()[4]) }, func() interface{} { return bool(bool(false)) }).(bool)
	whoIs := bool(_whoIs)
	_ = whoIs

	// Virtual field
	_readRange := utils.InlineIf((bool((len(payload.GetData())) > (5))), func() interface{} { return bool(payload.GetData()[5]) }, func() interface{} { return bool(bool(false)) }).(bool)
	readRange := bool(_readRange)
	_ = readRange

	// Virtual field
	_utcTimeSynchronization := utils.InlineIf((bool((len(payload.GetData())) > (6))), func() interface{} { return bool(payload.GetData()[6]) }, func() interface{} { return bool(bool(false)) }).(bool)
	utcTimeSynchronization := bool(_utcTimeSynchronization)
	_ = utcTimeSynchronization

	// Virtual field
	_lifeSafetyOperation := utils.InlineIf((bool((len(payload.GetData())) > (7))), func() interface{} { return bool(payload.GetData()[7]) }, func() interface{} { return bool(bool(false)) }).(bool)
	lifeSafetyOperation := bool(_lifeSafetyOperation)
	_ = lifeSafetyOperation

	// Virtual field
	_subscribeCovProperty := utils.InlineIf((bool((len(payload.GetData())) > (8))), func() interface{} { return bool(payload.GetData()[8]) }, func() interface{} { return bool(bool(false)) }).(bool)
	subscribeCovProperty := bool(_subscribeCovProperty)
	_ = subscribeCovProperty

	// Virtual field
	_getEventInformation := utils.InlineIf((bool((len(payload.GetData())) > (9))), func() interface{} { return bool(payload.GetData()[9]) }, func() interface{} { return bool(bool(false)) }).(bool)
	getEventInformation := bool(_getEventInformation)
	_ = getEventInformation

	if closeErr := readBuffer.CloseContext("BACnetServicesSupportedTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServicesSupportedTagged")
	}

	// Create the instance
	return &_BACnetServicesSupportedTagged{
		TagNumber: tagNumber,
		TagClass:  tagClass,
		Header:    header,
		Payload:   payload,
	}, nil
}

func (m *_BACnetServicesSupportedTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetServicesSupportedTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetServicesSupportedTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetServicesSupportedTagged")
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for header")
	}
	_headerErr := writeBuffer.WriteSerializable(ctx, m.GetHeader())
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for header")
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
	}

	// Simple Field (payload)
	if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for payload")
	}
	_payloadErr := writeBuffer.WriteSerializable(ctx, m.GetPayload())
	if popErr := writeBuffer.PopContext("payload"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for payload")
	}
	if _payloadErr != nil {
		return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
	}
	// Virtual field
	if _writeGroupErr := writeBuffer.WriteVirtual(ctx, "writeGroup", m.GetWriteGroup()); _writeGroupErr != nil {
		return errors.Wrap(_writeGroupErr, "Error serializing 'writeGroup' field")
	}
	// Virtual field
	if _subscribeCovPropertyMultipleErr := writeBuffer.WriteVirtual(ctx, "subscribeCovPropertyMultiple", m.GetSubscribeCovPropertyMultiple()); _subscribeCovPropertyMultipleErr != nil {
		return errors.Wrap(_subscribeCovPropertyMultipleErr, "Error serializing 'subscribeCovPropertyMultiple' field")
	}
	// Virtual field
	if _confirmedCovNotificationMultipleErr := writeBuffer.WriteVirtual(ctx, "confirmedCovNotificationMultiple", m.GetConfirmedCovNotificationMultiple()); _confirmedCovNotificationMultipleErr != nil {
		return errors.Wrap(_confirmedCovNotificationMultipleErr, "Error serializing 'confirmedCovNotificationMultiple' field")
	}
	// Virtual field
	if _unconfirmedCovNotificationMultipleErr := writeBuffer.WriteVirtual(ctx, "unconfirmedCovNotificationMultiple", m.GetUnconfirmedCovNotificationMultiple()); _unconfirmedCovNotificationMultipleErr != nil {
		return errors.Wrap(_unconfirmedCovNotificationMultipleErr, "Error serializing 'unconfirmedCovNotificationMultiple' field")
	}
	// Virtual field
	if _whoIsErr := writeBuffer.WriteVirtual(ctx, "whoIs", m.GetWhoIs()); _whoIsErr != nil {
		return errors.Wrap(_whoIsErr, "Error serializing 'whoIs' field")
	}
	// Virtual field
	if _readRangeErr := writeBuffer.WriteVirtual(ctx, "readRange", m.GetReadRange()); _readRangeErr != nil {
		return errors.Wrap(_readRangeErr, "Error serializing 'readRange' field")
	}
	// Virtual field
	if _utcTimeSynchronizationErr := writeBuffer.WriteVirtual(ctx, "utcTimeSynchronization", m.GetUtcTimeSynchronization()); _utcTimeSynchronizationErr != nil {
		return errors.Wrap(_utcTimeSynchronizationErr, "Error serializing 'utcTimeSynchronization' field")
	}
	// Virtual field
	if _lifeSafetyOperationErr := writeBuffer.WriteVirtual(ctx, "lifeSafetyOperation", m.GetLifeSafetyOperation()); _lifeSafetyOperationErr != nil {
		return errors.Wrap(_lifeSafetyOperationErr, "Error serializing 'lifeSafetyOperation' field")
	}
	// Virtual field
	if _subscribeCovPropertyErr := writeBuffer.WriteVirtual(ctx, "subscribeCovProperty", m.GetSubscribeCovProperty()); _subscribeCovPropertyErr != nil {
		return errors.Wrap(_subscribeCovPropertyErr, "Error serializing 'subscribeCovProperty' field")
	}
	// Virtual field
	if _getEventInformationErr := writeBuffer.WriteVirtual(ctx, "getEventInformation", m.GetGetEventInformation()); _getEventInformationErr != nil {
		return errors.Wrap(_getEventInformationErr, "Error serializing 'getEventInformation' field")
	}

	if popErr := writeBuffer.PopContext("BACnetServicesSupportedTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetServicesSupportedTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetServicesSupportedTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetServicesSupportedTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetServicesSupportedTagged) isBACnetServicesSupportedTagged() bool {
	return true
}

func (m *_BACnetServicesSupportedTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
