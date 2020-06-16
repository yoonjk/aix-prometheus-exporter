# start shell 
echo "["

# diskspace df -m
df -m | awk -f ./scripts/diskfree.awk -
# network packet
netstat -in | awk -f ./scripts/network_packet.awk -

# uptime
uptime | awk '{print $3" "substr($4, 0, length($4) -1) " " substr($5, 0, 2)"  hours "substr($5, 4,2) " minutes"}'|awk -f ./scripts/convert.awk -

# process metrics
ps -ef | wc -l | awk -f ./scripts/procs_all.awk -
# process namedgroup
ps -ef | grep -c java | awk -v proc_name='java' -f ./scripts/procs_namegroup.awk -
ps -ef | grep -c '/usr/sbin' | awk -v proc_name='/usr/sbin' -f ./scripts/procs_namegroup.awk -
ps -ef | grep -c '/opt/rsct/bin/rmcd -a' | awk -v proc_name='/opt/rsct/bin/rmcd -a' -f ./scripts/procs_namegroup.awk - 

# The last script don't have a comma in result
ps -ef| grep defunct | wc -l | awk -f ./scripts/procs_defunct.awk -
# end shell
echo "]"
