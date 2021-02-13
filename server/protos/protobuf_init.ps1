echo "building to go"
protoc --go_out=./ *.proto
echo "building to cpp"
protoc --cpp_out=../../client/UserClient/ShockChat_Qt/protos *.proto
echo "finished"