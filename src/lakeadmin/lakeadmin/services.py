import os
import docker
from lakeadmin import LakeAdmin
import sys

Admin = LakeAdmin.LakeAdmin()

def start_s3_service(uid):
    user = Admin.get_user(uid)
    if not user:
        return

    ceph_rgw_url = os.environ.get('CEPH_RGW_URL')

    ak = ''
    sk = ''

    #docker_client = docker.DockerClient(base_url='unix://var/run/docker.sock')
    docker_client = docker.from_env()

    for key in user['keys']:
        ak = key.get('access_key')
        sk = key.get('secret_key')
        user = key.get('user')

        if not ak or not sk:
            continue

        image = 'ehualu.com/minio'
        #name = '%s.s3.datalake.com' % user
        name = 's3-%s' % user
        env = ['MINIO_ACCESS_KEY=%s' % ak, 'MINIO_SECRET_KEY=%s' % sk]
        args = ['gateway', 's3',  ceph_rgw_url]

        print(name)

        docker_client.services.create(
                image = image,
                name = name,
                env = env,
                networks = ['datalake'], 
                args = args)

def start_s3_services():
    for uid in Admin.get_users():
        # start up minio for this user
        try:
            print("Found user %s, trying to start s3 service" % uid)
            start_s3_service(uid)
        except Exception as e:
            print("Got exception %s while trying to start s3 service for %s" % (e, uid))

