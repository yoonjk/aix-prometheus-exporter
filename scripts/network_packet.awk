function network_transmit_packets(NET_NAME, PACKET_SIZE)
{
  line="{\"name\": \"network_transmit_packets_total\", \"label_name\": [\"device\"],\"label_value\": [\""NET_NAME"\"], \"value\": "PACKET_SIZE"},"
  return line
}

function network_receive_packets(NET_NAME, PACKET_SIZE)
{
  line="{\"name\": \"network_receive_packets_total\", \"label_name\": [ \"device\"],\"label_value\":[\""NET_NAME"\"], \"value\":"PACKET_SIZE"},"
  return line 
}

function format(BLOCK_SIZE) 
{
  if (BLOCK_SIZE >= 0.0) {
    return BLOCK_SIZE
  } else { 
    return "0.00"
  }
}
BEGIN { i = 0; freesize = "" }
# main routine
{
  if (i++ > 0) {
    if ($4 ~ /^[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+$/) {
      print network_transmit_packets($1, $5)
      freesize = freesize "" network_receive_packets($1, $7) "\n"
    }
  }
}
END { 
   print substr(freesize, 0, length(freesize) - 1) 
}
