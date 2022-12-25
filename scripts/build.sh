#!/bin/sh
docker build -t codecatt/disco:base - < images/base.dockerfile
docker build -t codecatt/disco:js - < images/js.dockerfile
docker build -t codecatt/disco:py - < images/py.dockerfile
docker build -t codecatt/disco:php - < images/php.dockerfile
