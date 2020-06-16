
function node_procs_defunct(PROC_COUNT) {
  line="{\"name\": \"node_procs_defunct\", \"label_name\": [],\"label_value\": [], \"value\": "PROC_COUNT"}"
  return line
}
BEGIN { i = 0; freesize = "" }
# main routine
{
    print node_procs_defunct($1)
}
END { 
}
