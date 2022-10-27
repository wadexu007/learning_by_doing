## ProFTPD

[ProFTPD](http://www.proftpd.org/) is a high-performance, extremely configurable, and most of all a secure FTP server, featuring Apache-like configuration and blazing performance.

This page contains how to setup ProFTPD via authentication by pwd/key which stored in mysql.

### Prerequisites
* ProFTPD 1.3.6
* Mysql 5.7
* Linux Server CentOS 7


### Pre-Install
* make sure openssl-devel, gcc, mysql-devel and git are installed
```
yum install openssl-devel gcc mysql-devel git -y
```

* get proftpd source code ready
```
git clone -b 1.3.6  https://github.com/proftpd/proftpd/
# move to opt folder
mv proftpd/ /opt/
```

* Create mysql database
```
CREATE DATABASE IF NOT EXISTS sftp DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
GRANT ALL PRIVILEGES ON sftp.* TO sftpuser@'%' IDENTIFIED BY '<PASSWD>'; 
```

* Generate tables in database sftp
```
DROP TABLE IF EXISTS sftpgroup;CREATE TABLE sftpgroup (groupname VARCHAR(30) NOT NULL,gid INTEGER NOT NULL,members VARCHAR(255));CREATE INDEX groups_gid_idx ON sftpgroup (gid);
CREATE TABLE sftpuser (userid VARCHAR(30) NOT NULL UNIQUE,passwd VARCHAR(80) NOT NULL,uid INTEGER UNIQUE,gid INTEGER,homedir VARCHAR(255),shell VARCHAR(255));CREATE INDEX users_userid_idx ON sftpuser (userid);
CREATE TABLE sftpuserkeys (user VARCHAR(256) NOT NULL,user_key VARCHAR(8192) NOT NULL);CREATE INDEX sftpuserkeys_idx ON sftpuserkeys (user);
CREATE TABLE sftphostkeys (host VARCHAR(256) NOT NULL,host_key VARCHAR(8192) NOT NULL);CREATE INDEX sftphostkeys_idx ON sftphostkeys (host);
```

### Compile source code
Navigate to source code directory `cd /opt/proftpd/`

#### Option 1 without proxy_protocol
```
./configure --enable-dso --enable-ctrls --enable-openssl \
--with-modules=mod_sql:mod_sql_mysql:mod_sftp:mod_sftp_sql:mod_ban \
--with-shared=mod_ctrls_admin \
--with-includes=/usr/include/mysql \
--with-libraries=/usr/lib64/mysql
```

#### Option 2 with proxy_protocol
Download [mod_proxy_protocol](https://htmlpreview.github.io/?https://github.com/Castaglia/proftpd-mod_proxy_protocol/blob/master/mod_proxy_protocol.html#Installation) and copy mod_proxy_protocol.c to proftpd-dir/contrib/
```
./configure --enable-dso --enable-ctrls --enable-openssl \
--with-modules=mod_sql:mod_sql_mysql:mod_sftp:mod_sftp_sql:mod_ban:mod_proxy_protocol \
--with-shared=mod_ctrls_admin \
--with-includes=/usr/include/mysql \
--with-libraries=/usr/lib64/mysql
```

**Notes**: To get client’s source IP so that whitelist our SFTP, some Cloud TCP LB need enable proxy protocol, then our sftp server also need enable proxy protocol to receive client connection information passed through proxy servers and load balancers.


### Install
`make`

`make install` if no error occurs


### Configurations
create a new group for all sftp users
`groupadd -g 51 sftp`

change permission for ssh_host_rsa_key
`chmod 600 /etc/ssh/ssh_host_rsa_key`

 
```
mkdir /opt/proftpd.d/logs -p
touch /opt/proftpd.d/ban.tab
vi /opt/proftpd.d/mod_sql.conf
```

Quick way to start the service in backend:

`/opt/proftpd/proftpd -c /opt/proftpd.d/mod_sql.conf`

if you want to use proftpd directly,  you can add the command to PATH, vim ~/.bashrc in root user
```
% which proftpd
/usr/local/sbin/proftpd

# add to bottom
export PATH=$PATH:/usr/local/sbin/
```

also, add service start command to `/etc/rc.d/rc.local` for auto start when vm reboot.


For stop: `ps -ef | grep proftpd` to kill process id

### Storage
* Setup NFS
* AWS use s3fs mount bucket

### High Availability
Add another sftp server to LB backend services for HA.


### Configure SFTP User
**Manually**

Configure users manually which support both public key and password

Notice: in `mod_sql.conf` configuration `SQLMinUserUID 2000` so can't use id less than 2000

cat `/etc/passwd` to see what is current uid and never use a duplicate id.

```
useradd -u 2001 -g 51 -d /u01/sftp01 -s /sbin/nologin sftp01

# generate key pair
ssh-keygen -f /opt/sftp_genuser/keypairs/sftp01 -t rsa -C "sftp01" -P "" -N ""

# This public key format is OpenSSH, but SFTP server (proftpd) only support SSH2 public key,
# Output SSH2 format which for insert into mysql DB
ssh-keygen -e -f /opt/sftp_genuser/keypairs/sftp01.pub   

# connect to sftp db schema
mysql -h 10.0.5.4 -u sftpuser -p sftp

insert into sftpuser values ('sftp01','encrypt('pwd')','2001','51','/u01/sftp01','/sbin/nologin');

insert into sftpgroup values ('sftp','51','sftp01');

insert into sftpuserkeys values ('sftp01','<Public Key>');

#example:
insert into sftpuserkeys values ('sftp03','---- BEGIN SSH2 PUBLIC KEY ----
Comment: "2048-bit RSA, converted by root@fcplsftp01 from OpenSSH"
AAAAB3NzaC1yc2EAAAADXXXXXXXXXXXC2wmxAVb/xxxxxxxxx/xxxxxx/xxxxx/X
xxxxxxxx
xxxxxxxx
---- END SSH2 PUBLIC KEY ----');

```

Delete user

```
userdel -fr $user

# connect to sftp db schema
mysql -h 10.0.5.4 -u sftpuser -p sftp

delete from sftpgroup where members='xxx';

delete from sftpuser where uid='xxx';

delete from sftpuserkeys where user='xxx';"
```
Notes: need execute `useradd/userdel` in two vm instances if there is.

**Automatically**

Scripts to create/delete sftp user - **TBD**



### Test
Connect via key
```
sftp -P8888 -i <private_key> sftp01@<sftp server>
```

Connect via pwd
```
sftp -P8888 sftp01@<sftp server>
```

### FAQ

**Ban**
Because below settings MaxLoginAttempts 3

If you are locked by SFTP server. You can unlock in server via

```
#ssh to SFTP server

# check which ip banned
tail -f /opt/proftpd.d/logs/ban.log

# permit this ip
ftpdctl permit host 35.240.2.xx

```
Notes: if the IP is not client IP, but LB’s IP, then you need to fix this issue, otherwise all clients will be prevented to login sftp server once there is lock.


in Some Cloud LB, there is a proxy protocol, if enabled it, you will get Client’s IP, but will met errors when access SFTP.

```
Bad packet length 1349676916.
ssh_dispatch_run_fatal: Connection to 114.67.104.218 port 8888: message authentication code incorrect
Connection closed
```
In server side you will find errors like
```
Bad protocol version 'PROXY TCP4 58.241.167.212 10.0.4.3 47249 8888' from 10.0.4.7
```

To fix problem, need install mod_proxy_protocol following [Option 2](https://github.com/wadexu007/learning_by_doing/tree/main/ProFTPD#option-2-with-proxy_protocol)

Now if you want to access sftp instance IP from internal VPC, you will met errors, because SFTP server now running in proxy protocol, only can accept passthrough traffic from upstream LB which enable proxy protocol.

```
ssh_exchange_identification: read: Connection reset by peer
Couldn't read packet: Connection reset by peer
```

### Debug
enable trace logs

```
# Enable trace logs
# Update config then restart SFTP

Trace DEFAULT:10 sql:20
```


### List proftpd modules 
List all proftpd modules for reference.
```
% proftpd -l
Compiled-in modules:
  mod_core.c
  mod_xfer.c
  mod_rlimit.c
  mod_auth_unix.c
  mod_auth_file.c
  mod_auth.c
  mod_ls.c
  mod_log.c
  mod_site.c
  mod_delay.c
  mod_facts.c
  mod_dso.c
  mod_ident.c
  mod_sql.c
  mod_sql_mysql.c
  mod_sftp.c
  mod_sftp_sql.c
  mod_ban.c
  mod_cap.c
  mod_ctrls.c

# view settings of proftpd via -V 
% proftpd -V

```