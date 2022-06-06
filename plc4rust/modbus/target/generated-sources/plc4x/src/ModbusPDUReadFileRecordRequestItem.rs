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
// package org.apache.plc4x.rust.modbus.readwrite;

// Code generated by code-generation. DO NOT EDIT.
use std::io::{Error, ErrorKind, Read, Write};
use plc4rust::{Message, NoOption};
use plc4rust::read_buffer::ReadBuffer;
use plc4rust::write_buffer::WriteBuffer;


#[derive(PartialEq, Debug, Clone)]
pub struct ModbusPDUReadFileRecordRequestItemOptions {
}
#[derive(PartialEq, Debug, Clone)]
pub struct ModbusPDUReadFileRecordRequestItem {
    pub referenceType: u8,
    pub fileNumber: u16,
    pub recordNumber: u16,
    pub recordLength: u16
}

impl ModbusPDUReadFileRecordRequestItem {
}

impl Message for ModbusPDUReadFileRecordRequestItem {
    type M = ModbusPDUReadFileRecordRequestItem;
    type P = ModbusPDUReadFileRecordRequestItemOptions;

    fn get_length_in_bits(&self) -> u32 {
        todo!()
    }

    fn serialize<T: Write>(&self, writer: &mut WriteBuffer<T>) -> Result<usize, Error> {
        writer.write_u8(self.referenceType)?;
        writer.write_u16(self.fileNumber)?;
        writer.write_u16(self.recordNumber)?;
        writer.write_u16(self.recordLength)?;
        Ok(0)
    }

    fn parse<T: Read>(reader: &mut ReadBuffer<T>, parameter: Option<Self::P>) -> Result<Self::M, Error> {
        // (Re-)define the options
        let parameter = parameter.unwrap();
        let referenceType = reader.read_u8()?;
        let fileNumber = reader.read_u16()?;
        let recordNumber = reader.read_u16()?;
        let recordLength = reader.read_u16()?;
        Ok(Self::M {
            referenceType,
            fileNumber,
            recordNumber,
            recordLength
        })
    }
}


