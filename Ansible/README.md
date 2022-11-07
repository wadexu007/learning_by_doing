## What Is Ansible
[Ansible](https://www.ansible.com/) is a suite of software tools that enables infrastructure as code. It is open-source and the suite includes software provisioning, configuration management, and application deployment functionality. 

### Installation
Just execute `brew install ansible` in MacOS

```
ansible version [core 2.13.5]
python version = 3.10.8
```

### Quick Start 
Use **Command** line

touch a hosts file with all your host IPs
```
# stop nginx in all hosts
ansible "*" -i hosts -m shell -a "service nginx stop" -u wade_xu -b

# remove docker container in all hosts
ansible "*" -i hosts -m shell -a "docker ps | awk '/8080-/''{print }' | cut -c -12 | xargs -I {} sudo  docker rm -f {}" -u wade_xu -b

# use chage command update xxx user in all hosts with pwd never expired
ansible "*" -i hosts -m shell -a "chage -I -1 -m 0 -M 99999 -E -1 xxx" -u wade_xu -b
```

### Playbook Usage
Demo 3 features about ansible
* use shell command, like `chage`
* create a file `config.yaml` and set owner/group
* restart a service, like nginx

Replace hosts IP in `demo/hosts/eng-dev.yaml`
```
cd demo

ansible-playbook demo1.yaml \
    -i hosts/eng-dev.yaml \
    --extra-vars "user_id=wade_xu key=123456 project_name=test-eng-dev region=apac"
```

### Test result
A config.yaml generated in /opt folder with correct owner and group.
```
$ ls -l /opt
total 3
-rwxr-xr-x  1 wade_xu root 149 Nov  7 18:14 config.yaml
```
```
$ cat /opt/config.yaml 
#########################
## Basic Configuration ##
#########################

env: test-eng-dev
product: product-demo
test_key: 123456
region: apac
```

Nginx restart and running just now at 2022-11-07 18:15:03
```
$ service nginx status
Redirecting to /bin/systemctl status nginx.service
● nginx.service - The nginx HTTP and reverse proxy server
   Loaded: loaded (/usr/lib/systemd/system/nginx.service; enabled; vendor preset: disabled)
   Active: active (running) since Mon 2022-11-07 18:15:03 AEDT; 45s ago
```

<br>
