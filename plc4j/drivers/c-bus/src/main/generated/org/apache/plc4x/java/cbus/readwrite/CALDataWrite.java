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

public class CALDataWrite extends CALData implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final Parameter paramNo;
  protected final byte code;
  protected final ParameterValue parameterValue;

  // Arguments.
  protected final RequestContext requestContext;

  public CALDataWrite(
      CALCommandTypeContainer commandTypeContainer,
      CALData additionalData,
      Parameter paramNo,
      byte code,
      ParameterValue parameterValue,
      RequestContext requestContext) {
    super(commandTypeContainer, additionalData, requestContext);
    this.paramNo = paramNo;
    this.code = code;
    this.parameterValue = parameterValue;
    this.requestContext = requestContext;
  }

  public Parameter getParamNo() {
    return paramNo;
  }

  public byte getCode() {
    return code;
  }

  public ParameterValue getParameterValue() {
    return parameterValue;
  }

  @Override
  protected void serializeCALDataChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("CALDataWrite");

    // Simple Field (paramNo)
    writeSimpleEnumField(
        "paramNo",
        "Parameter",
        paramNo,
        new DataWriterEnumDefault<>(
            Parameter::getValue, Parameter::name, writeUnsignedShort(writeBuffer, 8)));

    // Simple Field (code)
    writeSimpleField("code", code, writeByte(writeBuffer, 8));

    // Simple Field (parameterValue)
    writeSimpleField("parameterValue", parameterValue, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("CALDataWrite");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    CALDataWrite _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (paramNo)
    lengthInBits += 8;

    // Simple field (code)
    lengthInBits += 8;

    // Simple field (parameterValue)
    lengthInBits += parameterValue.getLengthInBits();

    return lengthInBits;
  }

  public static CALDataBuilder staticParseCALDataBuilder(
      ReadBuffer readBuffer,
      CALCommandTypeContainer commandTypeContainer,
      RequestContext requestContext)
      throws ParseException {
    readBuffer.pullContext("CALDataWrite");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    Parameter paramNo =
        readEnumField(
            "paramNo",
            "Parameter",
            new DataReaderEnumDefault<>(Parameter::enumForValue, readUnsignedShort(readBuffer, 8)));

    byte code = readSimpleField("code", readByte(readBuffer, 8));

    ParameterValue parameterValue =
        readSimpleField(
            "parameterValue",
            new DataReaderComplexDefault<>(
                () ->
                    ParameterValue.staticParse(
                        readBuffer,
                        (ParameterType) (paramNo.getParameterType()),
                        (short) ((commandTypeContainer.getNumBytes()) - (2))),
                readBuffer));

    readBuffer.closeContext("CALDataWrite");
    // Create the instance
    return new CALDataWriteBuilderImpl(paramNo, code, parameterValue, requestContext);
  }

  public static class CALDataWriteBuilderImpl implements CALData.CALDataBuilder {
    private final Parameter paramNo;
    private final byte code;
    private final ParameterValue parameterValue;
    private final RequestContext requestContext;

    public CALDataWriteBuilderImpl(
        Parameter paramNo,
        byte code,
        ParameterValue parameterValue,
        RequestContext requestContext) {
      this.paramNo = paramNo;
      this.code = code;
      this.parameterValue = parameterValue;
      this.requestContext = requestContext;
    }

    public CALDataWrite build(
        CALCommandTypeContainer commandTypeContainer,
        CALData additionalData,
        RequestContext requestContext) {
      CALDataWrite cALDataWrite =
          new CALDataWrite(
              commandTypeContainer, additionalData, paramNo, code, parameterValue, requestContext);
      return cALDataWrite;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CALDataWrite)) {
      return false;
    }
    CALDataWrite that = (CALDataWrite) o;
    return (getParamNo() == that.getParamNo())
        && (getCode() == that.getCode())
        && (getParameterValue() == that.getParameterValue())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getParamNo(), getCode(), getParameterValue());
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
