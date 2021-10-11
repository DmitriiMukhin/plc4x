/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.java.spi.codegen.fields;

import jdk.internal.org.objectweb.asm.TypeReference;
import org.apache.plc4x.java.spi.codegen.io.DataReader;
import org.apache.plc4x.java.spi.generation.ParseException;
import org.apache.plc4x.java.spi.generation.WithReaderArgs;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class FieldReaderArray<T> implements FieldReader<T> {

    private static final Logger LOGGER = LoggerFactory.getLogger(FieldReaderArray.class);

    @Override
    public T readField(String logicalName, DataReader<T> dataReader, WithReaderArgs... readerArgs) throws ParseException {
        return switchByteOrderIfNecessary(() -> dataReader.read(logicalName, readerArgs), dataReader, extractByteOder(readerArgs).orElse(null));
    }

    public T[] readFieldCount(String logicalName, DataReader<T> dataReader, int count, WithReaderArgs... readerArgs) throws ParseException {
        return null;
    }

    public T[] readFieldLength(String logicalName, DataReader<T> dataReader, int length, WithReaderArgs... readerArgs) throws ParseException {
        return null;
    }

    public T[] readFieldTerminated(String logicalName, DataReader<T> dataReader, T termination, WithReaderArgs... readerArgs) throws ParseException {
        return null;
    }

}
