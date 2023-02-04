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

public class BACnetContextTagBoolean extends BACnetContextTag implements Message {

  // Accessors for discriminator values.
  public BACnetDataType getDataType() {
    return BACnetDataType.BOOLEAN;
  }

  // Properties.
  protected final short value;
  protected final BACnetTagPayloadBoolean payload;

  // Arguments.
  protected final Short tagNumberArgument;

  public BACnetContextTagBoolean(
      BACnetTagHeader header,
      short value,
      BACnetTagPayloadBoolean payload,
      Short tagNumberArgument) {
    super(header, tagNumberArgument);
    this.value = value;
    this.payload = payload;
    this.tagNumberArgument = tagNumberArgument;
  }

  public short getValue() {
    return value;
  }

  public BACnetTagPayloadBoolean getPayload() {
    return payload;
  }

  public boolean getActualValue() {
    return (boolean) (getPayload().getValue());
  }

  @Override
  protected void serializeBACnetContextTagChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetContextTagBoolean");

    // Simple Field (value)
    writeSimpleField("value", value, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (payload)
    writeSimpleField("payload", payload, new DataWriterComplexDefault<>(writeBuffer));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean actualValue = getActualValue();
    writeBuffer.writeVirtual("actualValue", actualValue);

    writeBuffer.popContext("BACnetContextTagBoolean");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetContextTagBoolean _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (value)
    lengthInBits += 8;

    // Simple field (payload)
    lengthInBits += payload.getLengthInBits();

    // A virtual field doesn't have any in- or output.

    return lengthInBits;
  }

  public static BACnetContextTagBuilder staticParseBACnetContextTagBuilder(
      ReadBuffer readBuffer,
      BACnetTagHeader header,
      Short tagNumberArgument,
      BACnetDataType dataType)
      throws ParseException {
    readBuffer.pullContext("BACnetContextTagBoolean");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    // Validation
    if (!((header.getActualLength()) == (1))) {
      throw new ParseValidationException("length field should be 1");
    }

    short value = readSimpleField("value", readUnsignedShort(readBuffer, 8));

    BACnetTagPayloadBoolean payload =
        readSimpleField(
            "payload",
            new DataReaderComplexDefault<>(
                () -> BACnetTagPayloadBoolean.staticParse(readBuffer, (long) (value)), readBuffer));
    boolean actualValue = readVirtualField("actualValue", boolean.class, payload.getValue());

    readBuffer.closeContext("BACnetContextTagBoolean");
    // Create the instance
    return new BACnetContextTagBooleanBuilderImpl(value, payload, tagNumberArgument);
  }

  public static class BACnetContextTagBooleanBuilderImpl
      implements BACnetContextTag.BACnetContextTagBuilder {
    private final short value;
    private final BACnetTagPayloadBoolean payload;
    private final Short tagNumberArgument;

    public BACnetContextTagBooleanBuilderImpl(
        short value, BACnetTagPayloadBoolean payload, Short tagNumberArgument) {
      this.value = value;
      this.payload = payload;
      this.tagNumberArgument = tagNumberArgument;
    }

    public BACnetContextTagBoolean build(BACnetTagHeader header, Short tagNumberArgument) {
      BACnetContextTagBoolean bACnetContextTagBoolean =
          new BACnetContextTagBoolean(header, value, payload, tagNumberArgument);
      return bACnetContextTagBoolean;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetContextTagBoolean)) {
      return false;
    }
    BACnetContextTagBoolean that = (BACnetContextTagBoolean) o;
    return (getValue() == that.getValue())
        && (getPayload() == that.getPayload())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getValue(), getPayload());
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
