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

public class BACnetPriorityValueTime extends BACnetPriorityValue implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetApplicationTagTime timeValue;

  // Arguments.
  protected final BACnetObjectType objectTypeArgument;

  public BACnetPriorityValueTime(
      BACnetTagHeader peekedTagHeader,
      BACnetApplicationTagTime timeValue,
      BACnetObjectType objectTypeArgument) {
    super(peekedTagHeader, objectTypeArgument);
    this.timeValue = timeValue;
    this.objectTypeArgument = objectTypeArgument;
  }

  public BACnetApplicationTagTime getTimeValue() {
    return timeValue;
  }

  @Override
  protected void serializeBACnetPriorityValueChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetPriorityValueTime");

    // Simple Field (timeValue)
    writeSimpleField("timeValue", timeValue, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetPriorityValueTime");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetPriorityValueTime _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (timeValue)
    lengthInBits += timeValue.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetPriorityValueBuilder staticParseBACnetPriorityValueBuilder(
      ReadBuffer readBuffer, BACnetObjectType objectTypeArgument) throws ParseException {
    readBuffer.pullContext("BACnetPriorityValueTime");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetApplicationTagTime timeValue =
        readSimpleField(
            "timeValue",
            new DataReaderComplexDefault<>(
                () -> (BACnetApplicationTagTime) BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    readBuffer.closeContext("BACnetPriorityValueTime");
    // Create the instance
    return new BACnetPriorityValueTimeBuilderImpl(timeValue, objectTypeArgument);
  }

  public static class BACnetPriorityValueTimeBuilderImpl
      implements BACnetPriorityValue.BACnetPriorityValueBuilder {
    private final BACnetApplicationTagTime timeValue;
    private final BACnetObjectType objectTypeArgument;

    public BACnetPriorityValueTimeBuilderImpl(
        BACnetApplicationTagTime timeValue, BACnetObjectType objectTypeArgument) {
      this.timeValue = timeValue;
      this.objectTypeArgument = objectTypeArgument;
    }

    public BACnetPriorityValueTime build(
        BACnetTagHeader peekedTagHeader, BACnetObjectType objectTypeArgument) {
      BACnetPriorityValueTime bACnetPriorityValueTime =
          new BACnetPriorityValueTime(peekedTagHeader, timeValue, objectTypeArgument);
      return bACnetPriorityValueTime;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetPriorityValueTime)) {
      return false;
    }
    BACnetPriorityValueTime that = (BACnetPriorityValueTime) o;
    return (getTimeValue() == that.getTimeValue()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getTimeValue());
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
