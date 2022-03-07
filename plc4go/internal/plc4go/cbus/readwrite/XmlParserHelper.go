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

package readwrite

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/cbus/readwrite/model"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

// Code generated by code-generation. DO NOT EDIT.

type CbusXmlParserHelper struct {
}

// Temporary imports to silent compiler warnings (TODO: migrate from static to emission based imports)
func init() {
	_ = strconv.ParseUint
	_ = strconv.ParseInt
	_ = strings.Join
	_ = utils.Dump
}

func (m CbusXmlParserHelper) Parse(typeName string, xmlString string, parserArguments ...string) (interface{}, error) {
	switch typeName {
	case "CALData":
		return model.CALDataParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "Checksum":
		return model.ChecksumParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "CALReply":
		return model.CALReplyParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ExclamationMark":
		return model.ExclamationMarkParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "NetworkRoute":
		return model.NetworkRouteParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "Application":
		return model.ApplicationParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "NetworkNumber":
		return model.NetworkNumberParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "StandardFormatStatusReply":
		return model.StandardFormatStatusReplyParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "CBusOptions":
		return model.CBusOptionsParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "SALData":
		return model.SALDataParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "CBusCommand":
		srchk := parserArguments[0] == "true"
		return model.CBusCommandParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), srchk)
	case "IdentifyReplyCommand":
		attribute := model.AttributeByName(parserArguments[0])
		return model.IdentifyReplyCommandParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), attribute)
	case "BridgeCount":
		return model.BridgeCountParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "PowerUp":
		return model.PowerUpParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "Reply":
		return model.ReplyParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "SerialInterfaceAddress":
		return model.SerialInterfaceAddressParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BridgeAddress":
		return model.BridgeAddressParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "MonitoredSAL":
		return model.MonitoredSALParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ParameterChange":
		return model.ParameterChangeParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "StatusByte":
		return model.StatusByteParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ReplyNetwork":
		return model.ReplyNetworkParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ExtendedStatusHeader":
		return model.ExtendedStatusHeaderParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "CommandHeader":
		return model.CommandHeaderParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "Confirmation":
		return model.ConfirmationParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "CBusPointToMultiPointCommand":
		srchk := parserArguments[0] == "true"
		return model.CBusPointToMultiPointCommandParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), srchk)
	case "StatusHeader":
		return model.StatusHeaderParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "StatusRequest":
		return model.StatusRequestParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "UnitAddress":
		return model.UnitAddressParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "NetworkProtocolControlInformation":
		return model.NetworkProtocolControlInformationParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ExtendedFormatStatusReply":
		return model.ExtendedFormatStatusReplyParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "CBusHeader":
		return model.CBusHeaderParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "CBusPointToPointCommand":
		srchk := parserArguments[0] == "true"
		return model.CBusPointToPointCommandParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), srchk)
	case "Alpha":
		return model.AlphaParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "CBusPointToPointToMultipointCommand":
		srchk := parserArguments[0] == "true"
		return model.CBusPointToPointToMultipointCommandParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), srchk)
	}
	return nil, errors.Errorf("Unsupported type %s", typeName)
}