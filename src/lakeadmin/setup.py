#!/usr/bin/env python

from __future__ import with_statement
import ast
import re
try:
    from setuptools import setup
    extra = dict(test_suite="tests.test.suite", include_package_data=True)
except ImportError:
    from distutils.core import setup
    extra = {}

long_description = '''
'''

install_requires = [
    "requests",
    "gevent",
    "docker",
    "boto3",
    "rgwadmin",
    "kafka",
    "bottle==0.12.13",
    "cherrypy==8.9.1",
    "wsgi-request-logger",
    "prometheus_client",
    "retrying",
    "bottle-websocket"
]

_version_re = re.compile(r'__version__\s+=\s+(.*)')
with open('lakeadmin/__init__.py', 'rb') as f:
    version = str(ast.literal_eval(_version_re.search(
        f.read().decode('utf-8')).group(1)))

setup(
    name="lakeadmin",
    packages=["lakeadmin"],
    version=version,
    install_requires=install_requires,
    author="James Pan",
    author_email="panjm@ehualu.com",
    license="LGPL v2.1",
    description="Python Datalake Admin API",
    long_description=long_description,
    keywords=["datalake", "admin api"],

    scripts = [],
    entry_points = {
        'console_scripts': [
            'lakeadmind = lakeadmin.main:main'
            ]
        },

    classifiers=[
        'Development Status :: 5 - Production/Stable',
	'Programming Language :: Python :: 2',
    	'Programming Language :: Python :: 2.7',
    	'Programming Language :: Python :: 3',
    	'Programming Language :: Python :: 3.2',
    	'Programming Language :: Python :: 3.3',
    	'Programming Language :: Python :: 3.4',
    	'Programming Language :: Python :: 3.5',
    	'Programming Language :: Python :: 3.6'],
    **extra
)
