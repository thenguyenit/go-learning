#!/bin/csh
while true ; do

for env in ent-bptu4vkg5jfiw-production-vohbr3y@ssh.us.magentosite.cloud
do
        for file in /var/log/platform/bptu4vkg5jfiw/cron.log /app/bptu4vkg5jfiw/var/log/logic_broker/pull_returns.log
        do
            printf "$file\n"
            rsync -cavz $env:$file ./log/production/
        done
done

printf 'Monitoring \n'
go run main.go

printf 'Sleeping \n'
sleep 60
done
