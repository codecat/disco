#!/bin/sh
docker build -t ghcr.io/codecat/disco:base - < images/base.dockerfile
docker build -t ghcr.io/codecat/disco:js - < images/js.dockerfile
docker build -t ghcr.io/codecat/disco:py - < images/py.dockerfile
