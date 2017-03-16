terr(){
	echo "Build Error: " $1
	exit
}

go test -cover -bench 1000 ./src || terr "tests"
go build -v -o ./bin/compute ./src || terr "build"

echo "Build passed."
cd ./bin

err(){
	echo "E2E Error: " $1
	rm ./compute
	exit
}

./compute -grid=10 -circ=8 -pts=50 -its=1 | jq ".pi" &> /dev/null || err "pi field missing"
./compute -grid=10 -circ=8 -pts=50 -its=1 | jq ".png" &> /dev/null || err "png field missing"
! ./compute -grid=10 -circ=8 -pts=50 -its=0 > /dev/null || err "Fail on bad int"
! ./compute -grid=10 -circ=8 -pts=50 -its=-1 > /dev/null || err "Fail on bad int"
! ./compute -grid=10 -circ=8 -pts=50 -its=A > /dev/null || err "Fail on bad int" # '{"error":"Argument error."}'

echo "Tests pass!"

