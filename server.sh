#!/bin/bash
start() {
	./http_redirector &
	echo $! > server.pid
}
stop() {
	kill -KILL `cat server.pid`
}
case "$1" in
	start)
		start
	;;
	stop)
		stop
	;;
	restart)
		stop
		start
	;;
	*)
		echo start, stop or restart
	;;
esac