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
package org.apache.plc4x.java.bacnetip.readwrite;

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

public class BACnetConfirmedServiceRequestUnknown extends BACnetConfirmedServiceRequest
    implements Message {

  // Accessors for discriminator values.
  public BACnetConfirmedServiceChoice getServiceChoice() {
    return null;
  }

  // Properties.
  protected final byte[] unknownBytes;

  // Arguments.
  protected final Long serviceRequestPayloadLength;
  protected final Long serviceRequestLength;

  public BACnetConfirmedServiceRequestUnknown(
      byte[] unknownBytes, Long serviceRequestPayloadLength, Long serviceRequestLength) {
    super(serviceRequestLength);
    this.unknownBytes = unknownBytes;
    this.serviceRequestPayloadLength = serviceRequestPayloadLength;
    this.serviceRequestLength = serviceRequestLength;
  }

  public byte[] getUnknownBytes() {
    return unknownBytes;
  }

  @Override
  protected void serializeBACnetConfirmedServiceRequestChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetConfirmedServiceRequestUnknown");

    // Array Field (unknownBytes)
    writeByteArrayField("unknownBytes", unknownBytes, writeByteArray(writeBuffer, 8));

    writeBuffer.popContext("BACnetConfirmedServiceRequestUnknown");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetConfirmedServiceRequestUnknown _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Array field
    if (unknownBytes != null) {
      lengthInBits += 8 * unknownBytes.length;
    }

    return lengthInBits;
  }

  public static BACnetConfirmedServiceRequestBuilder
      staticParseBACnetConfirmedServiceRequestBuilder(
          ReadBuffer readBuffer, Long serviceRequestPayloadLength, Long serviceRequestLength)
          throws ParseException {
    readBuffer.pullContext("BACnetConfirmedServiceRequestUnknown");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte[] unknownBytes =
        readBuffer.readByteArray("unknownBytes", Math.toIntExact(serviceRequestPayloadLength));

    readBuffer.closeContext("BACnetConfirmedServiceRequestUnknown");
    // Create the instance
    return new BACnetConfirmedServiceRequestUnknownBuilderImpl(
        unknownBytes, serviceRequestPayloadLength, serviceRequestLength);
  }

  public static class BACnetConfirmedServiceRequestUnknownBuilderImpl
      implements BACnetConfirmedServiceRequest.BACnetConfirmedServiceRequestBuilder {
    private final byte[] unknownBytes;
    private final Long serviceRequestPayloadLength;
    private final Long serviceRequestLength;

    public BACnetConfirmedServiceRequestUnknownBuilderImpl(
        byte[] unknownBytes, Long serviceRequestPayloadLength, Long serviceRequestLength) {
      this.unknownBytes = unknownBytes;
      this.serviceRequestPayloadLength = serviceRequestPayloadLength;
      this.serviceRequestLength = serviceRequestLength;
    }

    public BACnetConfirmedServiceRequestUnknown build(Long serviceRequestLength) {

      BACnetConfirmedServiceRequestUnknown bACnetConfirmedServiceRequestUnknown =
          new BACnetConfirmedServiceRequestUnknown(
              unknownBytes, serviceRequestPayloadLength, serviceRequestLength);
      return bACnetConfirmedServiceRequestUnknown;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetConfirmedServiceRequestUnknown)) {
      return false;
    }
    BACnetConfirmedServiceRequestUnknown that = (BACnetConfirmedServiceRequestUnknown) o;
    return (getUnknownBytes() == that.getUnknownBytes()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getUnknownBytes());
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
