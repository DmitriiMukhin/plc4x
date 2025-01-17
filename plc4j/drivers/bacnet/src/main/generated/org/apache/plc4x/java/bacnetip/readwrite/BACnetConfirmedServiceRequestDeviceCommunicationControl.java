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

public class BACnetConfirmedServiceRequestDeviceCommunicationControl
    extends BACnetConfirmedServiceRequest implements Message {

  // Accessors for discriminator values.
  public BACnetConfirmedServiceChoice getServiceChoice() {
    return BACnetConfirmedServiceChoice.DEVICE_COMMUNICATION_CONTROL;
  }

  // Properties.
  protected final BACnetContextTagUnsignedInteger timeDuration;
  protected final BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged
      enableDisable;
  protected final BACnetContextTagCharacterString password;

  // Arguments.
  protected final Long serviceRequestLength;

  public BACnetConfirmedServiceRequestDeviceCommunicationControl(
      BACnetContextTagUnsignedInteger timeDuration,
      BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged enableDisable,
      BACnetContextTagCharacterString password,
      Long serviceRequestLength) {
    super(serviceRequestLength);
    this.timeDuration = timeDuration;
    this.enableDisable = enableDisable;
    this.password = password;
    this.serviceRequestLength = serviceRequestLength;
  }

  public BACnetContextTagUnsignedInteger getTimeDuration() {
    return timeDuration;
  }

  public BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged
      getEnableDisable() {
    return enableDisable;
  }

  public BACnetContextTagCharacterString getPassword() {
    return password;
  }

  @Override
  protected void serializeBACnetConfirmedServiceRequestChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetConfirmedServiceRequestDeviceCommunicationControl");

    // Optional Field (timeDuration) (Can be skipped, if the value is null)
    writeOptionalField("timeDuration", timeDuration, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (enableDisable)
    writeSimpleField("enableDisable", enableDisable, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (password) (Can be skipped, if the value is null)
    writeOptionalField("password", password, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetConfirmedServiceRequestDeviceCommunicationControl");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetConfirmedServiceRequestDeviceCommunicationControl _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Optional Field (timeDuration)
    if (timeDuration != null) {
      lengthInBits += timeDuration.getLengthInBits();
    }

    // Simple field (enableDisable)
    lengthInBits += enableDisable.getLengthInBits();

    // Optional Field (password)
    if (password != null) {
      lengthInBits += password.getLengthInBits();
    }

    return lengthInBits;
  }

  public static BACnetConfirmedServiceRequestBuilder
      staticParseBACnetConfirmedServiceRequestBuilder(
          ReadBuffer readBuffer, Long serviceRequestLength) throws ParseException {
    readBuffer.pullContext("BACnetConfirmedServiceRequestDeviceCommunicationControl");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetContextTagUnsignedInteger timeDuration =
        readOptionalField(
            "timeDuration",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (0),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged enableDisable =
        readSimpleField(
            "enableDisable",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged
                        .staticParse(
                            readBuffer, (short) (1), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetContextTagCharacterString password =
        readOptionalField(
            "password",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagCharacterString)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (2),
                            (BACnetDataType) (BACnetDataType.CHARACTER_STRING)),
                readBuffer));

    readBuffer.closeContext("BACnetConfirmedServiceRequestDeviceCommunicationControl");
    // Create the instance
    return new BACnetConfirmedServiceRequestDeviceCommunicationControlBuilderImpl(
        timeDuration, enableDisable, password, serviceRequestLength);
  }

  public static class BACnetConfirmedServiceRequestDeviceCommunicationControlBuilderImpl
      implements BACnetConfirmedServiceRequest.BACnetConfirmedServiceRequestBuilder {
    private final BACnetContextTagUnsignedInteger timeDuration;
    private final BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged
        enableDisable;
    private final BACnetContextTagCharacterString password;
    private final Long serviceRequestLength;

    public BACnetConfirmedServiceRequestDeviceCommunicationControlBuilderImpl(
        BACnetContextTagUnsignedInteger timeDuration,
        BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged enableDisable,
        BACnetContextTagCharacterString password,
        Long serviceRequestLength) {
      this.timeDuration = timeDuration;
      this.enableDisable = enableDisable;
      this.password = password;
      this.serviceRequestLength = serviceRequestLength;
    }

    public BACnetConfirmedServiceRequestDeviceCommunicationControl build(
        Long serviceRequestLength) {
      BACnetConfirmedServiceRequestDeviceCommunicationControl
          bACnetConfirmedServiceRequestDeviceCommunicationControl =
              new BACnetConfirmedServiceRequestDeviceCommunicationControl(
                  timeDuration, enableDisable, password, serviceRequestLength);
      return bACnetConfirmedServiceRequestDeviceCommunicationControl;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetConfirmedServiceRequestDeviceCommunicationControl)) {
      return false;
    }
    BACnetConfirmedServiceRequestDeviceCommunicationControl that =
        (BACnetConfirmedServiceRequestDeviceCommunicationControl) o;
    return (getTimeDuration() == that.getTimeDuration())
        && (getEnableDisable() == that.getEnableDisable())
        && (getPassword() == that.getPassword())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getTimeDuration(), getEnableDisable(), getPassword());
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
