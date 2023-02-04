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

public class S7PayloadWriteVarResponse extends S7Payload implements Message {

  // Accessors for discriminator values.
  public Short getParameterParameterType() {
    return (short) 0x05;
  }

  public Short getMessageType() {
    return (short) 0x03;
  }

  // Properties.
  protected final List<S7VarPayloadStatusItem> items;

  // Arguments.
  protected final S7Parameter parameter;

  public S7PayloadWriteVarResponse(List<S7VarPayloadStatusItem> items, S7Parameter parameter) {
    super(parameter);
    this.items = items;
    this.parameter = parameter;
  }

  public List<S7VarPayloadStatusItem> getItems() {
    return items;
  }

  @Override
  protected void serializeS7PayloadChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("S7PayloadWriteVarResponse");

    // Array Field (items)
    writeComplexTypeArrayField("items", items, writeBuffer);

    writeBuffer.popContext("S7PayloadWriteVarResponse");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    S7PayloadWriteVarResponse _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Array field
    if (items != null) {
      int i = 0;
      for (S7VarPayloadStatusItem element : items) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= items.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static S7PayloadBuilder staticParseS7PayloadBuilder(
      ReadBuffer readBuffer, Short messageType, S7Parameter parameter) throws ParseException {
    readBuffer.pullContext("S7PayloadWriteVarResponse");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    List<S7VarPayloadStatusItem> items =
        readCountArrayField(
            "items",
            new DataReaderComplexDefault<>(
                () -> S7VarPayloadStatusItem.staticParse(readBuffer), readBuffer),
            CAST(parameter, S7ParameterWriteVarResponse.class).getNumItems());

    readBuffer.closeContext("S7PayloadWriteVarResponse");
    // Create the instance
    return new S7PayloadWriteVarResponseBuilderImpl(items, parameter);
  }

  public static class S7PayloadWriteVarResponseBuilderImpl implements S7Payload.S7PayloadBuilder {
    private final List<S7VarPayloadStatusItem> items;
    private final S7Parameter parameter;

    public S7PayloadWriteVarResponseBuilderImpl(
        List<S7VarPayloadStatusItem> items, S7Parameter parameter) {
      this.items = items;
      this.parameter = parameter;
    }

    public S7PayloadWriteVarResponse build(S7Parameter parameter) {

      S7PayloadWriteVarResponse s7PayloadWriteVarResponse =
          new S7PayloadWriteVarResponse(items, parameter);
      return s7PayloadWriteVarResponse;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof S7PayloadWriteVarResponse)) {
      return false;
    }
    S7PayloadWriteVarResponse that = (S7PayloadWriteVarResponse) o;
    return (getItems() == that.getItems()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getItems());
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
