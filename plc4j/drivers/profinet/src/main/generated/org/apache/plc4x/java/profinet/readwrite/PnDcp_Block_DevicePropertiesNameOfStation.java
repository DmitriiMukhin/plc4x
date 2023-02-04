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

public class PnDcp_Block_DevicePropertiesNameOfStation extends PnDcp_Block implements Message {

  // Accessors for discriminator values.
  public PnDcp_BlockOptions getOption() {
    return PnDcp_BlockOptions.DEVICE_PROPERTIES_OPTION;
  }

  public Short getSuboption() {
    return (short) 2;
  }

  // Properties.
  protected final byte[] nameOfStation;

  // Arguments.
  protected final Integer blockLength;
  // Reserved Fields
  private Integer reservedField0;

  public PnDcp_Block_DevicePropertiesNameOfStation(byte[] nameOfStation, Integer blockLength) {
    super();
    this.nameOfStation = nameOfStation;
    this.blockLength = blockLength;
  }

  public byte[] getNameOfStation() {
    return nameOfStation;
  }

  @Override
  protected void serializePnDcp_BlockChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("PnDcp_Block_DevicePropertiesNameOfStation");

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField0 != null ? reservedField0 : (int) 0x0000,
        writeUnsignedInt(writeBuffer, 16));

    // Array Field (nameOfStation)
    writeByteArrayField("nameOfStation", nameOfStation, writeByteArray(writeBuffer, 8));

    // Padding Field (padding)
    writePaddingField(
        "padding",
        (int)
            ((org.apache.plc4x.java.profinet.readwrite.utils.StaticHelper.arrayLength(
                    nameOfStation))
                % (2)),
        (short) 0x00,
        writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("PnDcp_Block_DevicePropertiesNameOfStation");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    PnDcp_Block_DevicePropertiesNameOfStation _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Reserved Field (reserved)
    lengthInBits += 16;

    // Array field
    if (nameOfStation != null) {
      lengthInBits += 8 * nameOfStation.length;
    }

    // Padding Field (padding)
    int _timesPadding =
        (int)
            ((org.apache.plc4x.java.profinet.readwrite.utils.StaticHelper.arrayLength(
                    nameOfStation))
                % (2));
    while (_timesPadding-- > 0) {
      lengthInBits += 8;
    }

    return lengthInBits;
  }

  public static PnDcp_BlockBuilder staticParsePnDcp_BlockBuilder(
      ReadBuffer readBuffer, Integer blockLength) throws ParseException {
    readBuffer.pullContext("PnDcp_Block_DevicePropertiesNameOfStation");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    Integer reservedField0 =
        readReservedField("reserved", readUnsignedInt(readBuffer, 16), (int) 0x0000);

    byte[] nameOfStation =
        readBuffer.readByteArray("nameOfStation", Math.toIntExact((blockLength) - (2)));

    readPaddingField(
        readUnsignedShort(readBuffer, 8),
        (int)
            ((org.apache.plc4x.java.profinet.readwrite.utils.StaticHelper.arrayLength(
                    nameOfStation))
                % (2)));

    readBuffer.closeContext("PnDcp_Block_DevicePropertiesNameOfStation");
    // Create the instance
    return new PnDcp_Block_DevicePropertiesNameOfStationBuilderImpl(
        nameOfStation, blockLength, reservedField0);
  }

  public static class PnDcp_Block_DevicePropertiesNameOfStationBuilderImpl
      implements PnDcp_Block.PnDcp_BlockBuilder {
    private final byte[] nameOfStation;
    private final Integer blockLength;
    private final Integer reservedField0;

    public PnDcp_Block_DevicePropertiesNameOfStationBuilderImpl(
        byte[] nameOfStation, Integer blockLength, Integer reservedField0) {
      this.nameOfStation = nameOfStation;
      this.blockLength = blockLength;
      this.reservedField0 = reservedField0;
    }

    public PnDcp_Block_DevicePropertiesNameOfStation build() {
      PnDcp_Block_DevicePropertiesNameOfStation pnDcp_Block_DevicePropertiesNameOfStation =
          new PnDcp_Block_DevicePropertiesNameOfStation(nameOfStation, blockLength);
      pnDcp_Block_DevicePropertiesNameOfStation.reservedField0 = reservedField0;
      return pnDcp_Block_DevicePropertiesNameOfStation;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PnDcp_Block_DevicePropertiesNameOfStation)) {
      return false;
    }
    PnDcp_Block_DevicePropertiesNameOfStation that = (PnDcp_Block_DevicePropertiesNameOfStation) o;
    return (getNameOfStation() == that.getNameOfStation()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getNameOfStation());
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
