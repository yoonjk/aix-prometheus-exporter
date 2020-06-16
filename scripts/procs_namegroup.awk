
function node_procs_named(PROC_NAME, PROC_COUNT) {
  line="{\"name\": \"process_namegroup_num_procs\", \"label_name\": [\"groupname\"],\"label_value\": [\""PROC_NAME"\"], \"value\": "PROC_COUNT"},"
  return line
}
BEGIN { }
# main routine
{
    print node_procs_named(proc_name, $1)
}
END { 
}
