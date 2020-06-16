df -m |awk '{print "disk_mount_volumn{Filesystem=\""$1"\",block=\"MB\"} "$2}'
dk -m |awk '{print "disk_mount_volumn_all_block{Filesystem=\""$1"\",mount=\""$7"\"} "$2}'

