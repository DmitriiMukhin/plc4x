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

// UnitStatus is an enum
type UnitStatus uint8

type IUnitStatus interface {
	utils.Serializable
}

const (
	UnitStatus_OK          UnitStatus = 0
	UnitStatus_NACK        UnitStatus = 1
	UnitStatus_NO_RESPONSE UnitStatus = 2
)

var UnitStatusValues []UnitStatus

func init() {
	_ = errors.New
	UnitStatusValues = []UnitStatus{
		UnitStatus_OK,
		UnitStatus_NACK,
		UnitStatus_NO_RESPONSE,
	}
}

func UnitStatusByValue(value uint8) (enum UnitStatus, ok bool) {
	switch value {
	case 0:
		return UnitStatus_OK, true
	case 1:
		return UnitStatus_NACK, true
	case 2:
		return UnitStatus_NO_RESPONSE, true
	}
	return 0, false
}

func UnitStatusByName(value string) (enum UnitStatus, ok bool) {
	switch value {
	case "OK":
		return UnitStatus_OK, true
	case "NACK":
		return UnitStatus_NACK, true
	case "NO_RESPONSE":
		return UnitStatus_NO_RESPONSE, true
	}
	return 0, false
}

func UnitStatusKnows(value uint8) bool {
	for _, typeValue := range UnitStatusValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastUnitStatus(structType interface{}) UnitStatus {
	castFunc := func(typ interface{}) UnitStatus {
		if sUnitStatus, ok := typ.(UnitStatus); ok {
			return sUnitStatus
		}
		return 0
	}
	return castFunc(structType)
}

func (m UnitStatus) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m UnitStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func UnitStatusParse(ctx context.Context, theBytes []byte) (UnitStatus, error) {
	return UnitStatusParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func UnitStatusParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (UnitStatus, error) {
	val, err := readBuffer.ReadUint8("UnitStatus", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading UnitStatus")
	}
	if enum, ok := UnitStatusByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return UnitStatus(val), nil
	} else {
		return enum, nil
	}
}

func (e UnitStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e UnitStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("UnitStatus", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e UnitStatus) PLC4XEnumName() string {
	switch e {
	case UnitStatus_OK:
		return "OK"
	case UnitStatus_NACK:
		return "NACK"
	case UnitStatus_NO_RESPONSE:
		return "NO_RESPONSE"
	}
	return ""
}

func (e UnitStatus) String() string {
	return e.PLC4XEnumName()
}
