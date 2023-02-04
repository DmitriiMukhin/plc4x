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
package org.apache.plc4x.java.openprotocol.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class OpenProtocolMessageApplicationGenericDataRequestRev1
    extends OpenProtocolMessageApplicationGenericDataRequest implements Message {

  // Accessors for discriminator values.
  public Long getRevision() {
    return (long) 1;
  }

  // Properties.
  protected final Mid requestMid;
  protected final long wantedRevision;
  protected final byte[] extraData;

  public OpenProtocolMessageApplicationGenericDataRequestRev1(
      Long midRevision,
      Short noAckFlag,
      Integer targetStationId,
      Integer targetSpindleId,
      Integer sequenceNumber,
      Short numberOfMessageParts,
      Short messagePartNumber,
      Mid requestMid,
      long wantedRevision,
      byte[] extraData) {
    super(
        midRevision,
        noAckFlag,
        targetStationId,
        targetSpindleId,
        sequenceNumber,
        numberOfMessageParts,
        messagePartNumber);
    this.requestMid = requestMid;
    this.wantedRevision = wantedRevision;
    this.extraData = extraData;
  }

  public Mid getRequestMid() {
    return requestMid;
  }

  public long getWantedRevision() {
    return wantedRevision;
  }

  public byte[] getExtraData() {
    return extraData;
  }

  @Override
  protected void serializeOpenProtocolMessageApplicationGenericDataRequestChild(
      WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("OpenProtocolMessageApplicationGenericDataRequestRev1");

    // Simple Field (requestMid)
    writeSimpleEnumField(
        "requestMid",
        "Mid",
        requestMid,
        new DataWriterEnumDefault<>(Mid::getValue, Mid::name, writeUnsignedLong(writeBuffer, 32)),
        WithOption.WithEncoding("ASCII"));

    // Simple Field (wantedRevision)
    writeSimpleField(
        "wantedRevision",
        wantedRevision,
        writeUnsignedLong(writeBuffer, 24),
        WithOption.WithEncoding("ASCII"));

    // Implicit Field (extraDataLength) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int extraDataLength = (int) (COUNT(getExtraData()));
    writeImplicitField(
        "extraDataLength",
        extraDataLength,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithEncoding("ASCII"));

    // Array Field (extraData)
    writeByteArrayField("extraData", extraData, writeByteArray(writeBuffer, 8));

    writeBuffer.popContext("OpenProtocolMessageApplicationGenericDataRequestRev1");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    OpenProtocolMessageApplicationGenericDataRequestRev1 _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (requestMid)
    lengthInBits += 32;

    // Simple field (wantedRevision)
    lengthInBits += 24;

    // Implicit Field (extraDataLength)
    lengthInBits += 16;

    // Array field
    if (extraData != null) {
      lengthInBits += 8 * extraData.length;
    }

    return lengthInBits;
  }

  public static OpenProtocolMessageApplicationGenericDataRequestBuilder
      staticParseOpenProtocolMessageApplicationGenericDataRequestBuilder(
          ReadBuffer readBuffer, Long revision) throws ParseException {
    readBuffer.pullContext("OpenProtocolMessageApplicationGenericDataRequestRev1");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    Mid requestMid =
        readEnumField(
            "requestMid",
            "Mid",
            new DataReaderEnumDefault<>(Mid::enumForValue, readUnsignedLong(readBuffer, 32)),
            WithOption.WithEncoding("ASCII"));

    long wantedRevision =
        readSimpleField(
            "wantedRevision", readUnsignedLong(readBuffer, 24), WithOption.WithEncoding("ASCII"));

    int extraDataLength =
        readImplicitField(
            "extraDataLength", readUnsignedInt(readBuffer, 16), WithOption.WithEncoding("ASCII"));

    byte[] extraData = readBuffer.readByteArray("extraData", Math.toIntExact(extraDataLength));

    readBuffer.closeContext("OpenProtocolMessageApplicationGenericDataRequestRev1");
    // Create the instance
    return new OpenProtocolMessageApplicationGenericDataRequestRev1BuilderImpl(
        requestMid, wantedRevision, extraData);
  }

  public static class OpenProtocolMessageApplicationGenericDataRequestRev1BuilderImpl
      implements OpenProtocolMessageApplicationGenericDataRequest
          .OpenProtocolMessageApplicationGenericDataRequestBuilder {
    private final Mid requestMid;
    private final long wantedRevision;
    private final byte[] extraData;

    public OpenProtocolMessageApplicationGenericDataRequestRev1BuilderImpl(
        Mid requestMid, long wantedRevision, byte[] extraData) {
      this.requestMid = requestMid;
      this.wantedRevision = wantedRevision;
      this.extraData = extraData;
    }

    public OpenProtocolMessageApplicationGenericDataRequestRev1 build(
        Long midRevision,
        Short noAckFlag,
        Integer targetStationId,
        Integer targetSpindleId,
        Integer sequenceNumber,
        Short numberOfMessageParts,
        Short messagePartNumber) {
      OpenProtocolMessageApplicationGenericDataRequestRev1
          openProtocolMessageApplicationGenericDataRequestRev1 =
              new OpenProtocolMessageApplicationGenericDataRequestRev1(
                  midRevision,
                  noAckFlag,
                  targetStationId,
                  targetSpindleId,
                  sequenceNumber,
                  numberOfMessageParts,
                  messagePartNumber,
                  requestMid,
                  wantedRevision,
                  extraData);
      return openProtocolMessageApplicationGenericDataRequestRev1;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof OpenProtocolMessageApplicationGenericDataRequestRev1)) {
      return false;
    }
    OpenProtocolMessageApplicationGenericDataRequestRev1 that =
        (OpenProtocolMessageApplicationGenericDataRequestRev1) o;
    return (getRequestMid() == that.getRequestMid())
        && (getWantedRevision() == that.getWantedRevision())
        && (getExtraData() == that.getExtraData())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getRequestMid(), getWantedRevision(), getExtraData());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}