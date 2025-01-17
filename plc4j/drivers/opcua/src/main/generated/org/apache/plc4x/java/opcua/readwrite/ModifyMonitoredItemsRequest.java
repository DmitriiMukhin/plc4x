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
package org.apache.plc4x.java.opcua.readwrite;

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

public class ModifyMonitoredItemsRequest extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "763";
  }

  // Properties.
  protected final ExtensionObjectDefinition requestHeader;
  protected final long subscriptionId;
  protected final TimestampsToReturn timestampsToReturn;
  protected final int noOfItemsToModify;
  protected final List<ExtensionObjectDefinition> itemsToModify;

  public ModifyMonitoredItemsRequest(
      ExtensionObjectDefinition requestHeader,
      long subscriptionId,
      TimestampsToReturn timestampsToReturn,
      int noOfItemsToModify,
      List<ExtensionObjectDefinition> itemsToModify) {
    super();
    this.requestHeader = requestHeader;
    this.subscriptionId = subscriptionId;
    this.timestampsToReturn = timestampsToReturn;
    this.noOfItemsToModify = noOfItemsToModify;
    this.itemsToModify = itemsToModify;
  }

  public ExtensionObjectDefinition getRequestHeader() {
    return requestHeader;
  }

  public long getSubscriptionId() {
    return subscriptionId;
  }

  public TimestampsToReturn getTimestampsToReturn() {
    return timestampsToReturn;
  }

  public int getNoOfItemsToModify() {
    return noOfItemsToModify;
  }

  public List<ExtensionObjectDefinition> getItemsToModify() {
    return itemsToModify;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("ModifyMonitoredItemsRequest");

    // Simple Field (requestHeader)
    writeSimpleField("requestHeader", requestHeader, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (subscriptionId)
    writeSimpleField("subscriptionId", subscriptionId, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (timestampsToReturn)
    writeSimpleEnumField(
        "timestampsToReturn",
        "TimestampsToReturn",
        timestampsToReturn,
        new DataWriterEnumDefault<>(
            TimestampsToReturn::getValue,
            TimestampsToReturn::name,
            writeUnsignedLong(writeBuffer, 32)));

    // Simple Field (noOfItemsToModify)
    writeSimpleField("noOfItemsToModify", noOfItemsToModify, writeSignedInt(writeBuffer, 32));

    // Array Field (itemsToModify)
    writeComplexTypeArrayField("itemsToModify", itemsToModify, writeBuffer);

    writeBuffer.popContext("ModifyMonitoredItemsRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ModifyMonitoredItemsRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (requestHeader)
    lengthInBits += requestHeader.getLengthInBits();

    // Simple field (subscriptionId)
    lengthInBits += 32;

    // Simple field (timestampsToReturn)
    lengthInBits += 32;

    // Simple field (noOfItemsToModify)
    lengthInBits += 32;

    // Array field
    if (itemsToModify != null) {
      int i = 0;
      for (ExtensionObjectDefinition element : itemsToModify) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= itemsToModify.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("ModifyMonitoredItemsRequest");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ExtensionObjectDefinition requestHeader =
        readSimpleField(
            "requestHeader",
            new DataReaderComplexDefault<>(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("391")),
                readBuffer));

    long subscriptionId = readSimpleField("subscriptionId", readUnsignedLong(readBuffer, 32));

    TimestampsToReturn timestampsToReturn =
        readEnumField(
            "timestampsToReturn",
            "TimestampsToReturn",
            new DataReaderEnumDefault<>(
                TimestampsToReturn::enumForValue, readUnsignedLong(readBuffer, 32)));

    int noOfItemsToModify = readSimpleField("noOfItemsToModify", readSignedInt(readBuffer, 32));

    List<ExtensionObjectDefinition> itemsToModify =
        readCountArrayField(
            "itemsToModify",
            new DataReaderComplexDefault<>(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("757")),
                readBuffer),
            noOfItemsToModify);

    readBuffer.closeContext("ModifyMonitoredItemsRequest");
    // Create the instance
    return new ModifyMonitoredItemsRequestBuilderImpl(
        requestHeader, subscriptionId, timestampsToReturn, noOfItemsToModify, itemsToModify);
  }

  public static class ModifyMonitoredItemsRequestBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final ExtensionObjectDefinition requestHeader;
    private final long subscriptionId;
    private final TimestampsToReturn timestampsToReturn;
    private final int noOfItemsToModify;
    private final List<ExtensionObjectDefinition> itemsToModify;

    public ModifyMonitoredItemsRequestBuilderImpl(
        ExtensionObjectDefinition requestHeader,
        long subscriptionId,
        TimestampsToReturn timestampsToReturn,
        int noOfItemsToModify,
        List<ExtensionObjectDefinition> itemsToModify) {
      this.requestHeader = requestHeader;
      this.subscriptionId = subscriptionId;
      this.timestampsToReturn = timestampsToReturn;
      this.noOfItemsToModify = noOfItemsToModify;
      this.itemsToModify = itemsToModify;
    }

    public ModifyMonitoredItemsRequest build() {
      ModifyMonitoredItemsRequest modifyMonitoredItemsRequest =
          new ModifyMonitoredItemsRequest(
              requestHeader, subscriptionId, timestampsToReturn, noOfItemsToModify, itemsToModify);
      return modifyMonitoredItemsRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ModifyMonitoredItemsRequest)) {
      return false;
    }
    ModifyMonitoredItemsRequest that = (ModifyMonitoredItemsRequest) o;
    return (getRequestHeader() == that.getRequestHeader())
        && (getSubscriptionId() == that.getSubscriptionId())
        && (getTimestampsToReturn() == that.getTimestampsToReturn())
        && (getNoOfItemsToModify() == that.getNoOfItemsToModify())
        && (getItemsToModify() == that.getItemsToModify())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getRequestHeader(),
        getSubscriptionId(),
        getTimestampsToReturn(),
        getNoOfItemsToModify(),
        getItemsToModify());
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
