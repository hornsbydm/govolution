RS-485 is a half duplex, multi-master bus. Nothing prevents two devices from sending data at the same time. Collisions happen, and the framing protocol detects and responds to them.

Frames consist of an 8 byte header, variable length data, and finally a CRC16 checksum over the header and data.

All devices receive all data and ignore frames not addressed to them. When a device receives a frame with a valid checksum and its own address as destination, it responds with a positive acknowledgement (`ACK02` or `ACK06`) or a negative acknowledgement (`NAK`). The sending device resends the frame if it does not receive an acknowledgement, as a collision or other data corruption may have occurred.

The thermostat (address 2001) is usually the master, with other devices responding to its requests. The SAM (address 9201, if present) is sometimes the master, usually sending frames to the thermostat. The small number of masters means there are relatively few collisions.

<table>
  <tr>
    <th colspan="12">Frame</th>
  </tr>
  <tr>
    <th colspan="8">Header</th>
    <th rowspan="2">Data</th>
    <th rowspan="2">Checksum</th>
  </tr>
  <tr>
    <th>1 byte</th>
    <th>1 byte</th>
    <th>1 byte</th>
    <th>1 byte</th>
    <th>1 byte</th>
    <th>1 byte</th>
    <th>1 byte</th>
    <th>1 byte</th>
  </tr>
  <tr>
    <td>Destination Address</td>
    <td>Destination Bus</td>
    <td>Source Address</td>
    <td>Source Bus</td>
    <td>Length</td>
    <td>PID</td>
    <td>EXT</td>
    <td>Function</td>
    <th>0-255 bytes</th>
    <th>2 bytes</th>
  </tr>
</table>

In practice for home use, the bus byte is always 0x01 and PID/EXT bytes are always 0x00.

* See [[Infinity - Known Devices]] for possible addresses; 0xf1f1 is the broadcast destination address
* The source address of a READ or WRITE frame is where the acknowledgement will be sent
* The length is the length in bytes of the data, excluding header and checksum
* The function is READ, WRITE, ACK06, etc. -- see below
* The checksum is CRC16 over header and data

Typical function codes:

    0x02 ACK02 -> "ACK02"
    0x06 ACK06 -> "ACK06"
         if length = 1 and data = 0x00 -– acknowledges a WRITE frame
         if length > 3 bytes -– a response to a READ frame; the first three data bytes are the register number and the remaining bytes are the contents of the register
    0x0B READ_TABLE_BLOCK -> "READ"
         length = 3 -- the register number to read
    0x0C WRITE_TABLE_BLOCK -> "WRITE"
         length > 3 bytes -- the first three data bytes are the register number and the remaining bytes are the new contents of the register
    0x15 NACK -> "NACK"
         length = 1 and data = 0x0A -- negative acknowledgement: no such table, table not writable, invalid data, function not supported

Other function codes:

    0x10 CHANGE_TABLE_NAME -> "CHGTBN"
    0x1E ALARM_PACKET -> "ALARM"
    0x22 READ_OBJECT_DATA -> "OBJRD"
    0x62 READ_VARIABLE -> "RDVAR"
    0x63 WRITE_VARIABLE -> "FORCE"
    0x64 AUTO_VARIABLE -> "AUTO"
    0x75 READ_LIST -> "LIST"

In observation, the thermostat (device 2001 or 1F01) always initiates a READ(0x0B) or WRITE(0x0C), and the response is always ACK06(0x06) or NAK(0x15). There is an instance of Device 0x4001 sending a READ(0x0B) frame to 0x4101, but it is unknown what this is for, as there is no response.

The thermostat regularly broadcasts (destination address F1F1) the date and time. A touch thermostat with source address 1F01 was reported to make additional broadcasts at startup time.