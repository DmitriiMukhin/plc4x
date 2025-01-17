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
package org.apache.plc4x.java.test.readwrite;

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

public abstract class TypeSwitchInTypeSwitchParentTypeChild0Infant0
    extends TypeSwitchInTypeSwitchParentTypeChild0 implements Message {

  // Accessors for discriminator values.
  public Short getChildNumber() {
    return (short) 0x01;
  }
  // Abstract accessors for discriminator values.
  public abstract Short getInfantNumber();

  // Properties.
  protected final short infantSomeField00;

  public TypeSwitchInTypeSwitchParentTypeChild0Infant0(
      short parentFieldHurz, short childFieldwolf, short infantSomeField00) {
    super(parentFieldHurz, childFieldwolf);
    this.infantSomeField00 = infantSomeField00;
  }

  public short getInfantSomeField00() {
    return infantSomeField00;
  }

  protected abstract void serializeTypeSwitchInTypeSwitchParentTypeChild0Infant0Child(
      WriteBuffer writeBuffer) throws SerializationException;

  @Override
  protected void serializeTypeSwitchInTypeSwitchParentTypeChild0Child(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("TypeSwitchInTypeSwitchParentTypeChild0Infant0");

    // Discriminator Field (infantNumber) (Used as input to a switch field)
    writeDiscriminatorField("infantNumber", getInfantNumber(), writeUnsignedShort(writeBuffer, 8));

    // Simple Field (infantSomeField00)
    writeSimpleField("infantSomeField00", infantSomeField00, writeUnsignedShort(writeBuffer, 8));

    // Switch field (Serialize the sub-type)
    serializeTypeSwitchInTypeSwitchParentTypeChild0Infant0Child(writeBuffer);

    writeBuffer.popContext("TypeSwitchInTypeSwitchParentTypeChild0Infant0");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    TypeSwitchInTypeSwitchParentTypeChild0Infant0 _value = this;

    // Discriminator Field (infantNumber)
    lengthInBits += 8;

    // Simple field (infantSomeField00)
    lengthInBits += 8;

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static TypeSwitchInTypeSwitchParentTypeChild0Builder
      staticParseTypeSwitchInTypeSwitchParentTypeChild0Builder(ReadBuffer readBuffer)
          throws ParseException {
    readBuffer.pullContext("TypeSwitchInTypeSwitchParentTypeChild0Infant0");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    short infantNumber = readDiscriminatorField("infantNumber", readUnsignedShort(readBuffer, 8));

    short infantSomeField00 =
        readSimpleField("infantSomeField00", readUnsignedShort(readBuffer, 8));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    TypeSwitchInTypeSwitchParentTypeChild0Infant0Builder builder = null;
    if (EvaluationHelper.equals(infantNumber, (short) 0x98)) {
      builder =
          TypeSwitchInTypeSwitchParentTypeChild0Infant0InfantsChild0
              .staticParseTypeSwitchInTypeSwitchParentTypeChild0Infant0Builder(readBuffer);
    } else if (EvaluationHelper.equals(infantNumber, (short) 0x99)) {
      builder =
          TypeSwitchInTypeSwitchParentTypeChild0Infant0InfantsChild1
              .staticParseTypeSwitchInTypeSwitchParentTypeChild0Infant0Builder(readBuffer);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "infantNumber="
              + infantNumber
              + "]");
    }

    readBuffer.closeContext("TypeSwitchInTypeSwitchParentTypeChild0Infant0");
    // Create the instance
    return new TypeSwitchInTypeSwitchParentTypeChild0Infant0BuilderImpl(infantSomeField00, builder);
  }

  public interface TypeSwitchInTypeSwitchParentTypeChild0Infant0Builder {
    TypeSwitchInTypeSwitchParentTypeChild0Infant0 build(
        short parentFieldHurz, short childFieldwolf, short infantSomeField00);
  }

  public static class TypeSwitchInTypeSwitchParentTypeChild0Infant0BuilderImpl
      implements TypeSwitchInTypeSwitchParentTypeChild0
          .TypeSwitchInTypeSwitchParentTypeChild0Builder {
    private final short infantSomeField00;
    private final TypeSwitchInTypeSwitchParentTypeChild0Infant0Builder builder;

    public TypeSwitchInTypeSwitchParentTypeChild0Infant0BuilderImpl(
        short infantSomeField00, TypeSwitchInTypeSwitchParentTypeChild0Infant0Builder builder) {

      this.infantSomeField00 = infantSomeField00;
      this.builder = builder;
    }

    public TypeSwitchInTypeSwitchParentTypeChild0Infant0 build(
        short parentFieldHurz, short childFieldwolf) {
      return builder.build(parentFieldHurz, childFieldwolf, infantSomeField00);
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof TypeSwitchInTypeSwitchParentTypeChild0Infant0)) {
      return false;
    }
    TypeSwitchInTypeSwitchParentTypeChild0Infant0 that =
        (TypeSwitchInTypeSwitchParentTypeChild0Infant0) o;
    return (getInfantSomeField00() == that.getInfantSomeField00()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getInfantSomeField00());
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
