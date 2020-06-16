function disk_volume_size(FS_NAME, MNT_POINT, BLOCK_SIZE)
{
  line="{\"name\": \"disk_volume_size\", \"label_name\": [\"filesystem\", \"mount\", \"unit\"],\"label_value\": [\""FS_NAME"\",\""MNT_POINT"\", \"MB blocks\"], \"value\": "BLOCK_SIZE"},"
  return line
}

function disk_volume_freesize(FS_NAME, MNT_POINT, BLOCK_SIZE)
{
  line="{\"name\": \"disk_mount_volumn_freesize\", \"label_name\": [ \"filesystem\",\"mount\", \"unit\"],\"label_value\":[\""FS_NAME"\",\""MNT_POINT"\",\"MB blocks\"], \"value\":"BLOCK_SIZE"},"
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
    print disk_volume_size($1, $7, format($2))
    freesize = freesize "" disk_volume_freesize($1, $7, format($3)) "\n"
  }
}
END { 
  print substr(freesize, 0, length(freesize) - 1)
}
