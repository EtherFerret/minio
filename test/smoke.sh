LAKEADMIN=http://localhost:8080

curl $LAKEADMIN/v1/users
curl $LAKEADMIN/v1/users/admin
curl -X PUT $LAKEADMIN/v1/users -d @user.json

access_key=P258ESU80MFG7M9A2NDY
secret_key=IbBXixBUveywDD2jypsxACXt3NPbX3aE32HM3YHw

mv ~/.aws/credentials ~/.aws/credentials.smoke
cat > ~/.aws/credentials << EOF
[default]
aws_access_key_id = $access_key
aws_secret_access_key = $secret_key
EOF

echo
echo list buckets
aws --endpoint-url http://localhost:9000 s3 ls
echo make bucket
aws --endpoint-url http://localhost:9000 s3 mb s3://firstbucket 
echo list buckets again
aws --endpoint-url http://localhost:9000 s3 ls
echo upload file
aws --endpoint-url http://localhost:9000 s3 cp user.json s3://firstbucket
echo download file
aws --endpoint-url http://localhost:9000 s3 cp s3://firstbucket/user.json /tmp/smoke.user.json
echo compare file
diff user.json /tmp/smoke.user.json
echo delete file
aws --endpoint-url http://localhost:9000 s3 rm s3://firstbucket/user.json
echo delete bucket
aws --endpoint-url http://localhost:9000 s3 rb s3://firstbucket 

echo delete user
curl -X DELETE $LAKEADMIN/v1/users/test

mv ~/.aws/credentials.smoke ~/.aws/credentials
