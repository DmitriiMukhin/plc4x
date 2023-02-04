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
package org.apache.plc4x.java.modbus.readwrite;

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

public class ModbusPDUReadDeviceIdentificationRequest extends ModbusPDU implements Message {

  // Accessors for discriminator values.
  public Boolean getErrorFlag() {
    return (boolean) false;
  }

  public Short getFunctionFlag() {
    return (short) 0x2B;
  }

  public Boolean getResponse() {
    return (boolean) false;
  }

  // Constant values.
  public static final Short MEITYPE = 0x0E;

  // Properties.
  protected final ModbusDeviceInformationLevel level;
  protected final short objectId;

  public ModbusPDUReadDeviceIdentificationRequest(
      ModbusDeviceInformationLevel level, short objectId) {
    super();
    this.level = level;
    this.objectId = objectId;
  }

  public ModbusDeviceInformationLevel getLevel() {
    return level;
  }

  public short getObjectId() {
    return objectId;
  }

  public short getMeiType() {
    return MEITYPE;
  }

  @Override
  protected void serializeModbusPDUChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("ModbusPDUReadDeviceIdentificationRequest");

    // Const Field (meiType)
    writeConstField("meiType", MEITYPE, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (level)
    writeSimpleEnumField(
        "level",
        "ModbusDeviceInformationLevel",
        level,
        new DataWriterEnumDefault<>(
            ModbusDeviceInformationLevel::getValue,
            ModbusDeviceInformationLevel::name,
            writeUnsignedShort(writeBuffer, 8)));

    // Simple Field (objectId)
    writeSimpleField("objectId", objectId, writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("ModbusPDUReadDeviceIdentificationRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ModbusPDUReadDeviceIdentificationRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Const Field (meiType)
    lengthInBits += 8;

    // Simple field (level)
    lengthInBits += 8;

    // Simple field (objectId)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static ModbusPDUBuilder staticParseModbusPDUBuilder(
      ReadBuffer readBuffer, Boolean response) throws ParseException {
    readBuffer.pullContext("ModbusPDUReadDeviceIdentificationRequest");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short meiType =
        readConstField(
            "meiType",
            readUnsignedShort(readBuffer, 8),
            ModbusPDUReadDeviceIdentificationRequest.MEITYPE);

    ModbusDeviceInformationLevel level =
        readEnumField(
            "level",
            "ModbusDeviceInformationLevel",
            new DataReaderEnumDefault<>(
                ModbusDeviceInformationLevel::enumForValue, readUnsignedShort(readBuffer, 8)));

    short objectId = readSimpleField("objectId", readUnsignedShort(readBuffer, 8));

    readBuffer.closeContext("ModbusPDUReadDeviceIdentificationRequest");
    // Create the instance
    return new ModbusPDUReadDeviceIdentificationRequestBuilderImpl(level, objectId);
  }

  public static class ModbusPDUReadDeviceIdentificationRequestBuilderImpl
      implements ModbusPDU.ModbusPDUBuilder {
    private final ModbusDeviceInformationLevel level;
    private final short objectId;

    public ModbusPDUReadDeviceIdentificationRequestBuilderImpl(
        ModbusDeviceInformationLevel level, short objectId) {
      this.level = level;
      this.objectId = objectId;
    }

    public ModbusPDUReadDeviceIdentificationRequest build() {
      ModbusPDUReadDeviceIdentificationRequest modbusPDUReadDeviceIdentificationRequest =
          new ModbusPDUReadDeviceIdentificationRequest(level, objectId);
      return modbusPDUReadDeviceIdentificationRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ModbusPDUReadDeviceIdentificationRequest)) {
      return false;
    }
    ModbusPDUReadDeviceIdentificationRequest that = (ModbusPDUReadDeviceIdentificationRequest) o;
    return (getLevel() == that.getLevel())
        && (getObjectId() == that.getObjectId())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getLevel(), getObjectId());
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
