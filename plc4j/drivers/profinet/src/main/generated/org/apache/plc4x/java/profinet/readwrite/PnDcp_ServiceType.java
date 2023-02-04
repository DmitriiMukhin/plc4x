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
package org.apache.plc4x.java.profinet.readwrite;

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

public class PnDcp_ServiceType implements Message {

  // Properties.
  protected final boolean notSupported;
  protected final boolean response;

  // Reserved Fields
  private Short reservedField0;
  private Byte reservedField1;

  public PnDcp_ServiceType(boolean notSupported, boolean response) {
    super();
    this.notSupported = notSupported;
    this.response = response;
  }

  public boolean getNotSupported() {
    return notSupported;
  }

  public boolean getResponse() {
    return response;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("PnDcp_ServiceType");

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField0 != null ? reservedField0 : (short) 0x00,
        writeUnsignedShort(writeBuffer, 5));

    // Simple Field (notSupported)
    writeSimpleField("notSupported", notSupported, writeBoolean(writeBuffer));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField1 != null ? reservedField1 : (byte) 0x00,
        writeUnsignedByte(writeBuffer, 1));

    // Simple Field (response)
    writeSimpleField("response", response, writeBoolean(writeBuffer));

    writeBuffer.popContext("PnDcp_ServiceType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    PnDcp_ServiceType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Reserved Field (reserved)
    lengthInBits += 5;

    // Simple field (notSupported)
    lengthInBits += 1;

    // Reserved Field (reserved)
    lengthInBits += 1;

    // Simple field (response)
    lengthInBits += 1;

    return lengthInBits;
  }

  public static PnDcp_ServiceType staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static PnDcp_ServiceType staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("PnDcp_ServiceType");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    Short reservedField0 =
        readReservedField("reserved", readUnsignedShort(readBuffer, 5), (short) 0x00);

    boolean notSupported = readSimpleField("notSupported", readBoolean(readBuffer));

    Byte reservedField1 =
        readReservedField("reserved", readUnsignedByte(readBuffer, 1), (byte) 0x00);

    boolean response = readSimpleField("response", readBoolean(readBuffer));

    readBuffer.closeContext("PnDcp_ServiceType");
    // Create the instance
    PnDcp_ServiceType _pnDcp_ServiceType;
    _pnDcp_ServiceType = new PnDcp_ServiceType(notSupported, response);
    _pnDcp_ServiceType.reservedField0 = reservedField0;
    _pnDcp_ServiceType.reservedField1 = reservedField1;
    return _pnDcp_ServiceType;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PnDcp_ServiceType)) {
      return false;
    }
    PnDcp_ServiceType that = (PnDcp_ServiceType) o;
    return (getNotSupported() == that.getNotSupported())
        && (getResponse() == that.getResponse())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getNotSupported(), getResponse());
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
