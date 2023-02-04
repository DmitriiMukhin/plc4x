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

public class BACnetTimeStampSequence extends BACnetTimeStamp implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetContextTagUnsignedInteger sequenceNumber;

  public BACnetTimeStampSequence(
      BACnetTagHeader peekedTagHeader, BACnetContextTagUnsignedInteger sequenceNumber) {
    super(peekedTagHeader);
    this.sequenceNumber = sequenceNumber;
  }

  public BACnetContextTagUnsignedInteger getSequenceNumber() {
    return sequenceNumber;
  }

  @Override
  protected void serializeBACnetTimeStampChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetTimeStampSequence");

    // Simple Field (sequenceNumber)
    writeSimpleField("sequenceNumber", sequenceNumber, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetTimeStampSequence");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetTimeStampSequence _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (sequenceNumber)
    lengthInBits += sequenceNumber.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetTimeStampBuilder staticParseBACnetTimeStampBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("BACnetTimeStampSequence");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetContextTagUnsignedInteger sequenceNumber =
        readSimpleField(
            "sequenceNumber",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (1),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    readBuffer.closeContext("BACnetTimeStampSequence");
    // Create the instance
    return new BACnetTimeStampSequenceBuilderImpl(sequenceNumber);
  }

  public static class BACnetTimeStampSequenceBuilderImpl
      implements BACnetTimeStamp.BACnetTimeStampBuilder {
    private final BACnetContextTagUnsignedInteger sequenceNumber;

    public BACnetTimeStampSequenceBuilderImpl(BACnetContextTagUnsignedInteger sequenceNumber) {
      this.sequenceNumber = sequenceNumber;
    }

    public BACnetTimeStampSequence build(BACnetTagHeader peekedTagHeader) {
      BACnetTimeStampSequence bACnetTimeStampSequence =
          new BACnetTimeStampSequence(peekedTagHeader, sequenceNumber);
      return bACnetTimeStampSequence;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetTimeStampSequence)) {
      return false;
    }
    BACnetTimeStampSequence that = (BACnetTimeStampSequence) o;
    return (getSequenceNumber() == that.getSequenceNumber()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getSequenceNumber());
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
