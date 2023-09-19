#!/bin/bash
GO_VERSION="1.21.0"
VERSION=$(grep -Eo 'version\s*=\s*"[^"]+"' main.go | grep -Eo '"[^"]+"' | tr -d '"')
NAME=$(grep 'Name' suitekit.rpm.spec | awk '{print $2}')
GO_LINK="https://go.dev/dl/go$GO_VERSION.linux-amd64.tar.gz"

yum install -y rpm-build wget 
# tar current project
tar -czf ~/source.tgz .
mkdir -p ~/rpmbuild/SOURCES/
# prepare source 
mv ~/source.tgz ~/rpmbuild/SOURCES/
wget -O ~/rpmbuild/SOURCES/go.tgz $GO_LINK
# build rpm
rpmbuild -ba suitekit.rpm.spec --define "version $VERSION"
# copy rpm pkg to current dir
cp ~/rpmbuild/RPMS/x86_64/$NAME-$VERSION-1.x86_64.rpm .
