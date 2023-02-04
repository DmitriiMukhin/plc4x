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
package org.apache.plc4x.java.cbus.readwrite;

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

public class AirConditioningDataSetHumidityUpperGuardLimit extends AirConditioningData
    implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final byte zoneGroup;
  protected final HVACZoneList zoneList;
  protected final HVACHumidity limit;
  protected final HVACHumidityModeAndFlags hvacModeAndFlags;

  public AirConditioningDataSetHumidityUpperGuardLimit(
      AirConditioningCommandTypeContainer commandTypeContainer,
      byte zoneGroup,
      HVACZoneList zoneList,
      HVACHumidity limit,
      HVACHumidityModeAndFlags hvacModeAndFlags) {
    super(commandTypeContainer);
    this.zoneGroup = zoneGroup;
    this.zoneList = zoneList;
    this.limit = limit;
    this.hvacModeAndFlags = hvacModeAndFlags;
  }

  public byte getZoneGroup() {
    return zoneGroup;
  }

  public HVACZoneList getZoneList() {
    return zoneList;
  }

  public HVACHumidity getLimit() {
    return limit;
  }

  public HVACHumidityModeAndFlags getHvacModeAndFlags() {
    return hvacModeAndFlags;
  }

  @Override
  protected void serializeAirConditioningDataChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("AirConditioningDataSetHumidityUpperGuardLimit");

    // Simple Field (zoneGroup)
    writeSimpleField("zoneGroup", zoneGroup, writeByte(writeBuffer, 8));

    // Simple Field (zoneList)
    writeSimpleField("zoneList", zoneList, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (limit)
    writeSimpleField("limit", limit, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (hvacModeAndFlags)
    writeSimpleField(
        "hvacModeAndFlags", hvacModeAndFlags, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("AirConditioningDataSetHumidityUpperGuardLimit");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    AirConditioningDataSetHumidityUpperGuardLimit _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (zoneGroup)
    lengthInBits += 8;

    // Simple field (zoneList)
    lengthInBits += zoneList.getLengthInBits();

    // Simple field (limit)
    lengthInBits += limit.getLengthInBits();

    // Simple field (hvacModeAndFlags)
    lengthInBits += hvacModeAndFlags.getLengthInBits();

    return lengthInBits;
  }

  public static AirConditioningDataBuilder staticParseAirConditioningDataBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("AirConditioningDataSetHumidityUpperGuardLimit");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte zoneGroup = readSimpleField("zoneGroup", readByte(readBuffer, 8));

    HVACZoneList zoneList =
        readSimpleField(
            "zoneList",
            new DataReaderComplexDefault<>(() -> HVACZoneList.staticParse(readBuffer), readBuffer));

    HVACHumidity limit =
        readSimpleField(
            "limit",
            new DataReaderComplexDefault<>(() -> HVACHumidity.staticParse(readBuffer), readBuffer));

    HVACHumidityModeAndFlags hvacModeAndFlags =
        readSimpleField(
            "hvacModeAndFlags",
            new DataReaderComplexDefault<>(
                () -> HVACHumidityModeAndFlags.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("AirConditioningDataSetHumidityUpperGuardLimit");
    // Create the instance
    return new AirConditioningDataSetHumidityUpperGuardLimitBuilderImpl(
        zoneGroup, zoneList, limit, hvacModeAndFlags);
  }

  public static class AirConditioningDataSetHumidityUpperGuardLimitBuilderImpl
      implements AirConditioningData.AirConditioningDataBuilder {
    private final byte zoneGroup;
    private final HVACZoneList zoneList;
    private final HVACHumidity limit;
    private final HVACHumidityModeAndFlags hvacModeAndFlags;

    public AirConditioningDataSetHumidityUpperGuardLimitBuilderImpl(
        byte zoneGroup,
        HVACZoneList zoneList,
        HVACHumidity limit,
        HVACHumidityModeAndFlags hvacModeAndFlags) {
      this.zoneGroup = zoneGroup;
      this.zoneList = zoneList;
      this.limit = limit;
      this.hvacModeAndFlags = hvacModeAndFlags;
    }

    public AirConditioningDataSetHumidityUpperGuardLimit build(
        AirConditioningCommandTypeContainer commandTypeContainer) {
      AirConditioningDataSetHumidityUpperGuardLimit airConditioningDataSetHumidityUpperGuardLimit =
          new AirConditioningDataSetHumidityUpperGuardLimit(
              commandTypeContainer, zoneGroup, zoneList, limit, hvacModeAndFlags);
      return airConditioningDataSetHumidityUpperGuardLimit;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof AirConditioningDataSetHumidityUpperGuardLimit)) {
      return false;
    }
    AirConditioningDataSetHumidityUpperGuardLimit that =
        (AirConditioningDataSetHumidityUpperGuardLimit) o;
    return (getZoneGroup() == that.getZoneGroup())
        && (getZoneList() == that.getZoneList())
        && (getLimit() == that.getLimit())
        && (getHvacModeAndFlags() == that.getHvacModeAndFlags())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getZoneGroup(), getZoneList(), getLimit(), getHvacModeAndFlags());
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
