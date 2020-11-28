#!/bin/bash

PRONAME="go-admin"

BIN=/usr/local/go-admin/api/$PRONAME
STDLOG=/usr/local/go-admin/api/output.log

chmod u+x $BIN

ID=$(/usr/sbin/pidof "$BIN")
if [ "$ID" ] ; then
        echo "kill -SIGINT $ID"
        kill -2 $ID
fi

count=0
while :
do
        ID=$(/usr/sbin/pidof "$BIN")
        if [ "$ID" ] ; then
                let count+=1;
                if [ "$count" -gt 50 ] ; then
                    echo "$PRONAME长时间未结束，启动失败，请重试"
                fi
                echo "$PRONAME still running...wait"
                sleep 0.1
        else
                echo "$PRONAME service was not started"
                echo "Starting service..."

                NOW=$(date +%Y-%m-%d,%H:%m:%s)
                #nohup $BIN > $STDLOG 2>&1 &
                nohup $BIN server -c config/settings-test.yml > "/data/logs/$PRONAME/output.$NOW" 2>&1 &
                break
        fi
done

ps -aux | grep $BIN