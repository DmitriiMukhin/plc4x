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
package org.apache.plc4x.java.s7.readwrite;

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

public class S7PayloadAlarmS extends S7PayloadUserDataItem implements Message {

  // Accessors for discriminator values.
  public Byte getCpuFunctionType() {
    return (byte) 0x00;
  }

  public Short getCpuSubfunction() {
    return (short) 0x12;
  }

  public Integer getDataLength() {
    return 0;
  }

  // Properties.
  protected final AlarmMessagePushType alarmMessage;

  public S7PayloadAlarmS(
      DataTransportErrorCode returnCode,
      DataTransportSize transportSize,
      AlarmMessagePushType alarmMessage) {
    super(returnCode, transportSize);
    this.alarmMessage = alarmMessage;
  }

  public AlarmMessagePushType getAlarmMessage() {
    return alarmMessage;
  }

  @Override
  protected void serializeS7PayloadUserDataItemChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("S7PayloadAlarmS");

    // Simple Field (alarmMessage)
    writeSimpleField("alarmMessage", alarmMessage, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("S7PayloadAlarmS");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    S7PayloadAlarmS _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (alarmMessage)
    lengthInBits += alarmMessage.getLengthInBits();

    return lengthInBits;
  }

  public static S7PayloadUserDataItemBuilder staticParseS7PayloadUserDataItemBuilder(
      ReadBuffer readBuffer, Byte cpuFunctionType, Short cpuSubfunction) throws ParseException {
    readBuffer.pullContext("S7PayloadAlarmS");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    AlarmMessagePushType alarmMessage =
        readSimpleField(
            "alarmMessage",
            new DataReaderComplexDefault<>(
                () -> AlarmMessagePushType.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("S7PayloadAlarmS");
    // Create the instance
    return new S7PayloadAlarmSBuilderImpl(alarmMessage);
  }

  public static class S7PayloadAlarmSBuilderImpl
      implements S7PayloadUserDataItem.S7PayloadUserDataItemBuilder {
    private final AlarmMessagePushType alarmMessage;

    public S7PayloadAlarmSBuilderImpl(AlarmMessagePushType alarmMessage) {
      this.alarmMessage = alarmMessage;
    }

    public S7PayloadAlarmS build(
        DataTransportErrorCode returnCode, DataTransportSize transportSize) {
      S7PayloadAlarmS s7PayloadAlarmS =
          new S7PayloadAlarmS(returnCode, transportSize, alarmMessage);
      return s7PayloadAlarmS;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof S7PayloadAlarmS)) {
      return false;
    }
    S7PayloadAlarmS that = (S7PayloadAlarmS) o;
    return (getAlarmMessage() == that.getAlarmMessage()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getAlarmMessage());
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
