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

public class ListOfCovNotificationsValue implements Message {

  // Properties.
  protected final BACnetPropertyIdentifierTagged propertyIdentifier;
  protected final BACnetContextTagUnsignedInteger arrayIndex;
  protected final BACnetConstructedData propertyValue;
  protected final BACnetContextTagTime timeOfChange;

  // Arguments.
  protected final BACnetObjectType objectTypeArgument;

  public ListOfCovNotificationsValue(
      BACnetPropertyIdentifierTagged propertyIdentifier,
      BACnetContextTagUnsignedInteger arrayIndex,
      BACnetConstructedData propertyValue,
      BACnetContextTagTime timeOfChange,
      BACnetObjectType objectTypeArgument) {
    super();
    this.propertyIdentifier = propertyIdentifier;
    this.arrayIndex = arrayIndex;
    this.propertyValue = propertyValue;
    this.timeOfChange = timeOfChange;
    this.objectTypeArgument = objectTypeArgument;
  }

  public BACnetPropertyIdentifierTagged getPropertyIdentifier() {
    return propertyIdentifier;
  }

  public BACnetContextTagUnsignedInteger getArrayIndex() {
    return arrayIndex;
  }

  public BACnetConstructedData getPropertyValue() {
    return propertyValue;
  }

  public BACnetContextTagTime getTimeOfChange() {
    return timeOfChange;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("ListOfCovNotificationsValue");

    // Simple Field (propertyIdentifier)
    writeSimpleField(
        "propertyIdentifier", propertyIdentifier, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (arrayIndex) (Can be skipped, if the value is null)
    writeOptionalField("arrayIndex", arrayIndex, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (propertyValue)
    writeSimpleField("propertyValue", propertyValue, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (timeOfChange) (Can be skipped, if the value is null)
    writeOptionalField("timeOfChange", timeOfChange, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("ListOfCovNotificationsValue");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    ListOfCovNotificationsValue _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (propertyIdentifier)
    lengthInBits += propertyIdentifier.getLengthInBits();

    // Optional Field (arrayIndex)
    if (arrayIndex != null) {
      lengthInBits += arrayIndex.getLengthInBits();
    }

    // Simple field (propertyValue)
    lengthInBits += propertyValue.getLengthInBits();

    // Optional Field (timeOfChange)
    if (timeOfChange != null) {
      lengthInBits += timeOfChange.getLengthInBits();
    }

    return lengthInBits;
  }

  public static ListOfCovNotificationsValue staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    if ((args == null) || (args.length != 1)) {
      throw new PlcRuntimeException(
          "Wrong number of arguments, expected 1, but got " + args.length);
    }
    BACnetObjectType objectTypeArgument;
    if (args[0] instanceof BACnetObjectType) {
      objectTypeArgument = (BACnetObjectType) args[0];
    } else if (args[0] instanceof String) {
      objectTypeArgument = BACnetObjectType.valueOf((String) args[0]);
    } else {
      throw new PlcRuntimeException(
          "Argument 0 expected to be of type BACnetObjectType or a string which is parseable but"
              + " was "
              + args[0].getClass().getName());
    }
    return staticParse(readBuffer, objectTypeArgument);
  }

  public static ListOfCovNotificationsValue staticParse(
      ReadBuffer readBuffer, BACnetObjectType objectTypeArgument) throws ParseException {
    readBuffer.pullContext("ListOfCovNotificationsValue");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetPropertyIdentifierTagged propertyIdentifier =
        readSimpleField(
            "propertyIdentifier",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetPropertyIdentifierTagged.staticParse(
                        readBuffer, (short) (0), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetContextTagUnsignedInteger arrayIndex =
        readOptionalField(
            "arrayIndex",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (1),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetConstructedData propertyValue =
        readSimpleField(
            "propertyValue",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetConstructedData.staticParse(
                        readBuffer,
                        (short) (2),
                        (BACnetObjectType) (objectTypeArgument),
                        (BACnetPropertyIdentifier) (propertyIdentifier.getValue()),
                        (BACnetTagPayloadUnsignedInteger)
                            (((((arrayIndex) != (null)) ? arrayIndex.getPayload() : null)))),
                readBuffer));

    BACnetContextTagTime timeOfChange =
        readOptionalField(
            "timeOfChange",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagTime)
                        BACnetContextTag.staticParse(
                            readBuffer, (short) (3), (BACnetDataType) (BACnetDataType.TIME)),
                readBuffer));

    readBuffer.closeContext("ListOfCovNotificationsValue");
    // Create the instance
    ListOfCovNotificationsValue _listOfCovNotificationsValue;
    _listOfCovNotificationsValue =
        new ListOfCovNotificationsValue(
            propertyIdentifier, arrayIndex, propertyValue, timeOfChange, objectTypeArgument);
    return _listOfCovNotificationsValue;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ListOfCovNotificationsValue)) {
      return false;
    }
    ListOfCovNotificationsValue that = (ListOfCovNotificationsValue) o;
    return (getPropertyIdentifier() == that.getPropertyIdentifier())
        && (getArrayIndex() == that.getArrayIndex())
        && (getPropertyValue() == that.getPropertyValue())
        && (getTimeOfChange() == that.getTimeOfChange())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getPropertyIdentifier(), getArrayIndex(), getPropertyValue(), getTimeOfChange());
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
