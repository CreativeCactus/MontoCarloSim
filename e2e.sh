terr(){
       echo "Build Error: " $1
       exit
}

option=${1:-"all"}

if [[ "testonly all" == *"$option"* ]]; then 
    go test -cover -bench 1000 ./src || terr "tests"
    echo "Tests pass!"
fi

if [[ "buildonly all" == *"$option"* ]]; then 
    go build -v -o ./bin/compute ./src || terr "build"
    echo "Built! Find binary in bin/compute"
fi

if [[ "testonly buildonly" == *"$option"* ]]; then 
    echo "Done!"
    exit
fi

cd bin

err(){
       echo "E2E Error: " $1
       rm ./compute
       exit
}

jq --version &> /dev/null || err "JQ missing"
./compute -grid=10 -circ=8 -pts=50 -its=1  -j | jq ".pi"  &> /dev/null || err "pi field missing"
./compute -grid=10 -circ=8 -pts=50 -its=1  -j | jq ".png" &> /dev/null || err "png field missing"
./compute -grid=10 -circ=8 -pts=50 -its=0  -j 2> /dev/null | grep "^{\"error\":\"Argument error.\"}$" &> /dev/null && err "Allowed zero arg"
./compute -grid=10 -circ=8 -pts=50 -its=-1 -j 2> /dev/null | grep "^{\"error\":\"Argument error.\"}$" &> /dev/null && err "Allowed negative arg"
./compute -grid=10 -circ=8 -pts=50 -its=A     &> /dev/null && err "Allowed invalid arg"
echo "End-to-ends pass!"

echo "Starting server on IPs: $(hostname -I)"
cd ..
node web/server.js
