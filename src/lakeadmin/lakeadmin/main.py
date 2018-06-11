import bottle
from datetime import datetime
from bottle import request, response
import logging
import os
import json
from lakeadmin import services
from lakeadmin import LakeAdmin
from bottle.ext.websocket import GeventWebSocketServer
from bottle.ext.websocket import websocket

logging_level = logging.INFO
logging_format = '%(asctime)s %(filename)s:%(lineno)d %(levelname)s %(message)s'

app = application = bottle.app()
logging.basicConfig(level=logging_level, format=logging_format)

Admin = LakeAdmin.LakeAdmin()

def exception_to_response(e, response):
    resp = {
            'message': str(e)
            }
    response.status = 404
    return json.dumps(resp, indent=4)

@app.get('/v1/users')
def users_get():
    try:
        users = Admin.get_users()
        print("====>", users)
        return json.dumps(users, indent=4)
    except Exception as e:
        logging.exception(e)
        return exception_to_response(e, response)

@app.get('/v1/users/<uid>')
def user_get(uid):
    try:
        user = Admin.get_user(uid)
        return json.dumps(user, indent=4)
    except Exception as e:
        logging.exception(e)
        return exception_to_response(e, response)

@app.put('/v1/users')
def user_add():
    try:
        data = request.body.read()
        print("=====>", data)
        user = json.loads(data.decode('utf-8'))

        user =  Admin.create_user(user)

        return json.dumps(user, indent=4)
    except Exception as e:
        logging.exception(e)
        return exception_to_response(e, response)


@app.delete('/v1/users/<uid>')
def user_delete(uid):
    try:
        Admin.remove_user(uid)
    except Exception as e:
        logging.exception(e)
        return exception_to_response(e, response)

@app.get('/v1/credentials/<ak>')
def credential_get(ak):
    try:
        cred = Admin.get_credential(ak)
        return json.dumps(cred, indent=4)
    except Exception as e:
        logging.exception(e)
        return exception_to_response(e, response)


@app.put('/v1/das')
def create_da():
    try:
        data = request.body.read()
        da = json.loads(data.decode('utf-8'))
        result = Admin.create_da(da)

        return json.dumps(result, indent=4)
    except Exception as e:
        logging.exception(e)
        return exception_to_response(e, response)

@app.get('/v1/das')
def da_list():
    try:
        das = Admin.get_das()
        return json.dumps(das, indent=4)
    except Exception as e:
        logging.exception(e)
        return exception_to_response(e, response)

@app.get('/v1/das/<daid>')
def da_get(daid):
    try:
        da = Admin.get_da(daid)
        return json.dumps(da, indent=4)
    except Exception as e:
        logging.exception(e)
        return exception_to_response(e, response)

# archive job
@app.put('/v1/archive_jobs')
def archive_job_create():
    try:
        data = request.body.read()
        job = json.loads(data.decode('utf-8'))
        result = Admin.create_archive_job(job)

        return json.dumps(result, indent=4)
    except Exception as e:
        logging.exception(e)
        return exception_to_response(e, response)

@app.get('/v1/archive_jobs')
def archive_jobs_list():
    try:
        jobs = Admin.get_archive_jobs()
        return json.dumps(result, indent=4)
    except Exception as e:
        logging.exception(e)
        return exception_to_response(e, response)

@app.get('/v1/archive_jobs/<jobid>')
def archive_jobs_list(jobid):
    try:
        job = Admin.get_archive_job(jobid)
        return json.dumps(result, indent=4)
    except Exception as e:
        logging.exception(e)
        return exception_to_response(e, response)


@app.get('/v1/notification', apply=[websocket])
def notification():
    pass


def main():
    import logging
    import sys
    import requestlogger

    loggedapp = requestlogger.WSGILogger(
        app,
        [logging.StreamHandler(stream=sys.stdout)],
        requestlogger.ApacheFormatter())

    # Start all services
    try:
        #services.start_s3_services()
        pass
    except Exception as e:
        print("Got exception %s" % e)
        raise

    # Start the API server
    bottle.run(loggedapp, server='cherrypy', host='0.0.0.0', port=8080)
