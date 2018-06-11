
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

echo set lifecycle
aws --endpoint-url http://localhost:9000 s3api put-bucket-lifecycle --bucket firstbucket --lifecycle-configuration file://lifecycle-rules.json

echo get lifecycle
aws --endpoint-url http://localhost:9000 s3api get-bucket-lifecycle --bucket firstbucket

echo delete lifecycle
aws --endpoint-url http://localhost:9000 s3api delete-bucket-lifecycle --bucket firstbucket

echo delete bucket
aws --endpoint-url http://localhost:9000 s3 rb s3://firstbucket 

echo delete user
curl -X DELETE $LAKEADMIN/v1/users/test

mv ~/.aws/credentials.smoke ~/.aws/credentials
