## Installing concourse

We will need to set up a concourse web VM and a concourse worker VM

##### Set up the concourse web

1. Set up an EC2 VM (you MUST use Ubuntu) (concourse 1.3+ recommended)
1. Update the security group for this VM (default will be something like `launch-wizard-1`) by opening up TCP port 8080
1. Download the `.pem` file. We'll assume it's called `concourse_server.pem` for this README
1. `chmod 400 concourse_server.pem`
1. Navigate to your EC2 page and grab the public DNS. It should look something like this: `ec2-12-34-56-78.us-west-2.compute.amazonaws.com`. The rest of this README will assume your DNS address is this - change accordingly
1. `ssh -i "concourse_server.pem" ubuntu@ec2-12-34-56-78.us-west-2.compute.amazonaws.com` (replace appropriately)
1. On the box
    1. Install and set up postgresql
    
        ```
        sudo apt-get update
        sudo apt-get install postgresql-9.3
        sudo apt-get install postgresql-client-9.3
        sudo su - # switch to root
        su - postgres # switch to postgres user from root
        mkdir -p /var/lib/postgresql/9.3/data
        
        psql # into the postgres shell
        
        CREATE USER ubuntu;
        ALTER USER ubuntu PASSWORD '';
        CREATE DATABASE atc;
        GRANT ALL ON DATABASE atc TO ubuntu;
        
        \q # quit psql
        
        exit # back to root from postgres user
        exit # back to ubuntu from root user
        ```

    1. Install and set up concourse
    
        ```
        ssh-keygen -t rsa -f host_key -N ''
        ssh-keygen -t rsa -f worker_key -N ''
        ssh-keygen -t rsa -f session_signing_key -N ''
        cp worker_key.pub authorized_worker_keys
        
        wget https://github.com/concourse/concourse/releases/download/v1.3.1/concourse_linux_amd64
        chmod a+x concourse_linux_amd64
        mv concourse_linux_amd64 /usr/local/bin/concourse
        concourse web \
          --basic-auth-username myuser \
          --basic-auth-password mypass \
          --session-signing-key session_signing_key \
          --tsa-host-key host_key \
          --tsa-authorized-keys authorized_worker_keys \
          --bind-port 8080 \
          --external-url http://ec2-12-34-56-78.us-west-2.compute.amazonaws.com:8080 &
        ```
1. Navigate to `http://ec2-12-34-56-78.us-west-2.compute.amazonaws.com:8080` in your browser

##### Set up the concourse worker

1. SSH into your EC2 instance again: `ssh -i "concourse_server.pem" ubuntu@ec2-12-34-56-78.us-west-2.compute.amazonaws.com` (replace appropriately)
1. On the box
    1. Start concourse worker
    
        ```
        sudo mkdir -p /opt/concourse/worker
        sudo chmod 777 /opt/concourse/worker
        
        sudo concourse worker \
          --work-dir /opt/concourse/worker \
          --tsa-host ec2-12-34-56-78.us-west-2.compute.amazonaws.com \
          --tsa-public-key host_key.pub \
          --tsa-worker-private-key worker_key &
        ```
1. Navigate to `http://ec2-12-34-56-78.us-west-2.compute.amazonaws.com:8080` in your browser

##### Check your installation

1. Once you've logged in at `http://ec2-12-34-56-78.us-west-2.compute.amazonaws.com:8080`, you should be able to download the fly CLI
1. With the fly CLI,
    1. `fly login -t concourse -c http://ec2-12-34-56-78.us-west-2.compute.amazonaws.com:8080`
    1. `fly -t concourse workers`
    1. You should see your one worker!
    
## Set up a pipeline

(this will not require you to SSH into your concourse box)

1. `fly login -t concourse -c http://ec2-12-34-56-78.us-west-2.compute.amazonaws.com:8080`
1. `fly -t concourse set-pipeline -p hello -c hello.yml`
1. `fly -t concourse unpause-pipeline -p hello`
1. Navigate to http://ec2-12-34-56-78.us-west-2.compute.amazonaws.com:8080/pipelines/hello to see your pipeline!