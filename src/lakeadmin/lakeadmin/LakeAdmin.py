import os
from rgwadmin import RGWAdmin
import boto3
import sys
if sys.version_info > (3,0):
    import urllib.parse as urlparse
else:
    import urlparse
import json
import logging

from kafka import KafkaProducer
from kafka import KafkaConsumer

class LakeAdmin:
    def __init__(self):
        # Env
        admin_ak = os.environ.get('ADMIN_AK')
        admin_sk = os.environ.get('ADMIN_SK')
        ceph_rgw_url = os.environ.get('CEPH_RGW_URL')
        kafka_ep = os.environ.get('KAFKA_EP')

        if not admin_ak or not admin_sk or not ceph_rgw_url:
            raise Exception("Please check the environment ADMIN_AK, ADMIN_SK, CEPH_RGW_URL!")

        print("AK %s SK %s URL %s" % (admin_ak, admin_sk, ceph_rgw_url))

        url = urlparse.urlparse(ceph_rgw_url)

        secure = False
        if url.scheme == 'https':
            secure = True

        self.rgw = RGWAdmin(access_key = admin_ak, secret_key = admin_sk, server = url.netloc, secure = secure)

        self.s3 = boto3.resource("s3",
                endpoint_url = ceph_rgw_url,
                aws_access_key_id = admin_ak,
                aws_secret_access_key = admin_sk)

        self.kafka_producer = KafkaProducer(bootstrap_servers = kafka_ep)
        self.kafka_consumer = KafkaConsumer(bootstrap_servers = kafka_ep)

        try:
            self.s3.Bucket('credentials').create()
            self.s3.Bucket('da_configuration').create()
            self.s3.Bucket('lifecycle_configuration').create()
            self.s3.Bucket('archive_jobs').create()
        except Exception as e:
            logging.exception(e)
            pass

    def get_user(self, uid):
        return self.rgw.get_user(uid)

    def remove_user(self, uid):
        try:
            user = self.rgw.get_user(uid)
            if not user:
                return

            for key in user['keys']:
                ak = key.get('access_key')
                sk = key.get('access_key')
                if ak:
                    self.s3.Object('credentials', ak).delete()
            self.rgw.remove_user(uid)
        except Exception as e:
            logging.exception(e)
            raise
 

    def get_credential(self, ak):
        obj = self.s3.Object('credentials', ak).get()
        data = obj['Body'].read()

        return json.loads(data.decode("utf-8"))

    def get_users(self):
        try:
            return self.rgw.get_users()
        except Exception as e:
            logging.exception(e)
            raise

    def create_user(self, user):
        uid = user['uid']

        display_name = uid
        email = ""
        max_buckets = 10000
        ak = None
        sk = None

        if 'display_name' in user:
            display_name = user.get('display_name')

        if 'email' in user:
            email = user.get('email')

        if 'access_key' in user:
            ak = user.get('access_key')
        
        if 'secret_key' in user:
            sk = user.get('secret_key')

        if 'max_buckets' in user:
            max_buckets = int(user.get('max_buckets'))

        self.rgw.create_user(
                uid = uid,
                display_name = display_name,
                email = email,
                max_buckets = max_buckets,
                access_key = ak,
                secret_key = sk
                )
        
        user = self.rgw.get_user(uid)
        for key in user['keys']:
            ak = key.get('access_key')
            sk = key.get('secret_key')
            if ak:
                data = json.dumps({'uid': uid, 'access_key': ak, 'secret_key': sk})
                print("!!writting data::", data)

                self.s3.Object('credentials', ak).put(Body=data)

        return user

    def create_da(da):
        daid = da['id']
        data = json.dumps(da)
        self.s3.Object('da_configuration', daid).put(Body=data)

    def get_das():
        das = []
        for obj in self.s3.Bucket('da_configuration'):
            das.append(obj.key)

        return das

    def get_da(daid):
        obj = self.s3.Object('da_configuration', daid).get()
        data = obj['Body'].read()
        return json.loads(data.decode("utf-8"))


    def create_archive_job(job):
        jobid = job['id']

        bucket = '*'
        if 'bucket' in job:
            bucket = job['bucket']

        data = json.dumps(job)
        self.s3.Object('archive_jobs', jobid).put(Body=data)
        event = {'sender': 'lakeadmin', 'type': 'create_archive_job', 'bucket': bucket}

        self.kafka_producer_send('lakeadmin-events', json.dumps())

    def get_archive_jobs():
        jobs = []
        for obj in self.s3.Bucket('archive_jobs'):
            jobs.append(obj.key)

        return jobs

    def get_archive_job(jobid):
        obj = self.s3.Object('archive_jobs', jobid).get()
        data = obj['Body'].read()
        return json.loads(data.decode("utf-8"))

