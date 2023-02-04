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

#ifndef PLC4C_S7_READ_WRITE_S7_VAR_REQUEST_PARAMETER_ITEM_H_
#define PLC4C_S7_READ_WRITE_S7_VAR_REQUEST_PARAMETER_ITEM_H_

#include <stdbool.h>
#include <stdint.h>
#include <plc4c/spi/context.h>
#include <plc4c/spi/read_buffer.h>
#include <plc4c/spi/write_buffer.h>
#include <plc4c/utils/list.h>
#include "s7_address.h"

// Code generated by code-generation. DO NOT EDIT.


// Structure used to contain the discriminator values for discriminated types using this as a parent
struct plc4c_s7_read_write_s7_var_request_parameter_item_discriminator {
  uint8_t itemType;
};
typedef struct plc4c_s7_read_write_s7_var_request_parameter_item_discriminator plc4c_s7_read_write_s7_var_request_parameter_item_discriminator;

// Enum assigning each subtype an individual id.
enum plc4c_s7_read_write_s7_var_request_parameter_item_type {
  plc4c_s7_read_write_s7_var_request_parameter_item_type_plc4c_s7_read_write_s7_var_request_parameter_item_address = 0};
typedef enum plc4c_s7_read_write_s7_var_request_parameter_item_type plc4c_s7_read_write_s7_var_request_parameter_item_type;

// Function to get the discriminator values for a given type.
plc4c_s7_read_write_s7_var_request_parameter_item_discriminator plc4c_s7_read_write_s7_var_request_parameter_item_get_discriminator(plc4c_s7_read_write_s7_var_request_parameter_item_type type);

struct plc4c_s7_read_write_s7_var_request_parameter_item {
  /* This is an abstract type so this property saves the type of this typed union */
  plc4c_s7_read_write_s7_var_request_parameter_item_type _type;
  /* Properties */
  union {
    struct { /* S7VarRequestParameterItemAddress */
      plc4c_s7_read_write_s7_address* s7_var_request_parameter_item_address_address;
    };
  };
};
typedef struct plc4c_s7_read_write_s7_var_request_parameter_item plc4c_s7_read_write_s7_var_request_parameter_item;

// Create an empty NULL-struct
plc4c_s7_read_write_s7_var_request_parameter_item plc4c_s7_read_write_s7_var_request_parameter_item_null();

plc4c_return_code plc4c_s7_read_write_s7_var_request_parameter_item_parse(plc4x_spi_context ctx, plc4c_spi_read_buffer* readBuffer, plc4c_s7_read_write_s7_var_request_parameter_item** message);

plc4c_return_code plc4c_s7_read_write_s7_var_request_parameter_item_serialize(plc4x_spi_context ctx, plc4c_spi_write_buffer* writeBuffer, plc4c_s7_read_write_s7_var_request_parameter_item* message);

uint16_t plc4c_s7_read_write_s7_var_request_parameter_item_length_in_bytes(plc4x_spi_context ctx, plc4c_s7_read_write_s7_var_request_parameter_item* message);

uint16_t plc4c_s7_read_write_s7_var_request_parameter_item_length_in_bits(plc4x_spi_context ctx, plc4c_s7_read_write_s7_var_request_parameter_item* message);

#endif  // PLC4C_S7_READ_WRITE_S7_VAR_REQUEST_PARAMETER_ITEM_H_
