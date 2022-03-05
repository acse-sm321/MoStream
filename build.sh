#! /bin/bash

# build web UI on local
cd ~/GolandProjects/MoStream/web
go install
cp ~/GolandProjects/bin/web ~/GolandProjects/bin/MoStream_ui/web
cp -R ~/GolandProjects/MoStream/templates ~/GolandProjects/bin/MoStream_ui