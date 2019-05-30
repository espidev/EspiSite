#!/bin/bash

cd generate/main
go build .
cd ../..
generate/main/generate
mv assets_vfsdata.go src/main
cd src/main
go build ./...
mv ./EspiSite ../../test

