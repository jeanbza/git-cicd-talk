# git-cicd-talk

[![Build Status](https://travis-ci.org/jadekler/git-cicd-talk.svg?branch=master)](https://travis-ci.org/jadekler/git-cicd-talk)

A talk on CI/CD.

THIS README IS IN PROGRESS

## Running

1. Install [golang](https://golang.org/dl/)
1. `go run main.go`

## Testing

1. `go run main.go`
1. `go test`

## Installation

1. Set up an EC2 VM (this README uses an amazon image (as opposed to rhel, windows, etc))
1. Update the security group for this VM (default will be something like `launch-wizard-1`) by opening up TCP port 8080
1. Download the `.pem` file. We'll assume it's called `concourse_server.pem` for this README
1. `chmod 400 concourse_server.pem`
1. Navigate to your EC2 page and grab the public DNS. It should look something like this: `ec2-12-34-56-78.us-west-2.compute.amazonaws.com`. The rest of this README will assume your DNS address is this - change accordingly
1. `ssh -i "concourse_server.pem" ec2-12-34-56-78.us-west-2.compute.amazonaws.com` (replace appropriately)
1. On the box
    1. Install and set up postgresql
    
        ```
        sudo yum install postgresql94-server.x86_64 # use yum list postgres*server to find an appropriate version
        sudo su - # log in as root from ec2-user
        su - postgres # switch to postgres user from root
        mkdir -p /var/lib/pgsql94/9.4/data
        initdb -D /var/lib/pgsql94/9.4/data/
        pg_ctl -D /var/lib/pgsql94/9.4/data/ -l logfile start
        
        psql # into the postgres shell
        
        CREATE USER "ec2-user";
        CREATE DATABASE atc;
        GRANT ALL ON DATABASE atc TO "ec2-user";
        
        \q # quit psql
        
        exit # back to root from postgres user
        exit # back to ec2-user from root  
        ```

    1. Install and set up concourse
    
        ```
        ssh-keygen -t rsa -f host_key -N ''
        ssh-keygen -t rsa -f worker_key -N ''
        ssh-keygen -t rsa -f session_signing_key -N ''
        cp worker_key.pub authorized_worker_keys
        
        wget https://github.com/concourse/concourse/releases/download/v1.2.0/concourse_linux_amd64
        chmod a+x concourse_linux_amd64
        sudo mv concourse_linux_amd64 /usr/local/bin/concourse
        concourse web \
          --basic-auth-username myuser \
          --basic-auth-password mypass \
          --session-signing-key session_signing_key \
          --tsa-host-key host_key \
          --tsa-authorized-keys authorized_worker_keys \
          --bind-port 8080 \
          --external-url http://ec2-12-34-56-78.us-west-2.compute.amazonaws.com:8080
        ```
1. Navigate to `http://ec2-12-34-56-78.us-west-2.compute.amazonaws.com:8080` in your browser
