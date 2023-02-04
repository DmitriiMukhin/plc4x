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

public class BACnetFaultParameterFaultExtendedParametersEntryUnsigned
    extends BACnetFaultParameterFaultExtendedParametersEntry implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetApplicationTagUnsignedInteger unsignedValue;

  public BACnetFaultParameterFaultExtendedParametersEntryUnsigned(
      BACnetTagHeader peekedTagHeader, BACnetApplicationTagUnsignedInteger unsignedValue) {
    super(peekedTagHeader);
    this.unsignedValue = unsignedValue;
  }

  public BACnetApplicationTagUnsignedInteger getUnsignedValue() {
    return unsignedValue;
  }

  @Override
  protected void serializeBACnetFaultParameterFaultExtendedParametersEntryChild(
      WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetFaultParameterFaultExtendedParametersEntryUnsigned");

    // Simple Field (unsignedValue)
    writeSimpleField("unsignedValue", unsignedValue, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetFaultParameterFaultExtendedParametersEntryUnsigned");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetFaultParameterFaultExtendedParametersEntryUnsigned _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (unsignedValue)
    lengthInBits += unsignedValue.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetFaultParameterFaultExtendedParametersEntryBuilder
      staticParseBACnetFaultParameterFaultExtendedParametersEntryBuilder(ReadBuffer readBuffer)
          throws ParseException {
    readBuffer.pullContext("BACnetFaultParameterFaultExtendedParametersEntryUnsigned");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetApplicationTagUnsignedInteger unsignedValue =
        readSimpleField(
            "unsignedValue",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetApplicationTagUnsignedInteger)
                        BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    readBuffer.closeContext("BACnetFaultParameterFaultExtendedParametersEntryUnsigned");
    // Create the instance
    return new BACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilderImpl(unsignedValue);
  }

  public static class BACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilderImpl
      implements BACnetFaultParameterFaultExtendedParametersEntry
          .BACnetFaultParameterFaultExtendedParametersEntryBuilder {
    private final BACnetApplicationTagUnsignedInteger unsignedValue;

    public BACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilderImpl(
        BACnetApplicationTagUnsignedInteger unsignedValue) {
      this.unsignedValue = unsignedValue;
    }

    public BACnetFaultParameterFaultExtendedParametersEntryUnsigned build(
        BACnetTagHeader peekedTagHeader) {
      BACnetFaultParameterFaultExtendedParametersEntryUnsigned
          bACnetFaultParameterFaultExtendedParametersEntryUnsigned =
              new BACnetFaultParameterFaultExtendedParametersEntryUnsigned(
                  peekedTagHeader, unsignedValue);
      return bACnetFaultParameterFaultExtendedParametersEntryUnsigned;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetFaultParameterFaultExtendedParametersEntryUnsigned)) {
      return false;
    }
    BACnetFaultParameterFaultExtendedParametersEntryUnsigned that =
        (BACnetFaultParameterFaultExtendedParametersEntryUnsigned) o;
    return (getUnsignedValue() == that.getUnsignedValue()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getUnsignedValue());
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
