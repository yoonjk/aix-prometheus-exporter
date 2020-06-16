BEGIN { a["seconds"]=1; a["minutes"]=60
        a["hours"]=3600; a["days"]=86400;
        a["weeks"]=7*86400 }
{ tmp=tolower($0) }
{ while (match(tmp,/[0-9]+[^a-z]*[a-z]+/)) {
    u=substr(tmp,RSTART,RLENGTH); q=u+0; gsub(/[^a-z]/,"",u)
    t+=q*a[u]
    tmp=substr(tmp,RSTART+RLENGTH)
  }
  print "{\"name\": \"node_boot_time_seconds\", \"label_name\":[], \"label_value\":[], \"value\": "t"\},"
}
