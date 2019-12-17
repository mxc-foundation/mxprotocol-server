![m2m-wallet_ui_public_logo_mxc_logo.png](docs/pics/m2m-wallet_ui_public_logo_mxc_logo.png)

-----------------

-m2m-wallet
--------------
M2M wallet, as a part of MXProtocol is responsible for the Payments and accounting among machines (gateways, and devices) in MXC network.   
M2M wallet will be used for accounting allocated and obtained network resources, Smart Machine Bidding and Data Market Place of MXProtocol.  
Every organization (may have gateways or/and IoT devices) in MXC network has a corresponding wallet in M2M wallet.

At the first stage, MXC M2M wallet will be used in the MXProtocol MVP. MXProtocol MVP will be released in Q4 2019.  


__Note.__ This preliminary version of M2M wallet is under development and supposed to be improved. 
Additional features will be added to M2M wallet  based on the MXProtocol design.

# Setup

See MXC Developer Handbook for further information.

Note: UI part from m2m has been merged into lpwan-app-server, m2m no longer contains UI part.  
However part of the APIs get data from m2m service, you need to start m2m service for accessing all features correctly.

## Environment

#### Set up docker
- [Install Docker](https://docs.docker.com/install/linux/docker-ce/ubuntu/)  
Just follow __Install using the repository / SET UP THE REPOSITORY__, no need to install docker engine community

- [Install docker-compose](https://docs.docker.com/compose/install/)
Just follow __Install Compose on Linux systems__

- Add user to docker group
```bash
$ sudo usermod -aG docker $USER
```

## Clone the repo:

```bash
git clone git@gitlab.com:MXCFoundation/cloud/mxprotocol-server.git &&
cd mxprotocol-server
```

## Fetch latest develop branch:

```bash
git fetch origin develop:develop &&
git checkout develop &&
git pull --rebase origin develop
```

## Existing or new feature branch

* New feature branch required?

```
git checkout -b feature/MCL-XXX
```

* Existing feature branch?

> Example: If there is a "feature" branch that you are working on in Jira
(i.e. feature/MCL-117) and you are working on a task of that feature,
then create a branch from that feature that is prefixed with your name
(i.e. luke/MCL-118-page-network-servers)

```bash
git fetch origin feature/MCL-117:feature/MCL-117 &&
git checkout feature/MCL-117 &&
git pull --rebase origin feature/MCL-117
```

## Create task branch from feature branch:

```bash
git checkout -b luke/MCL-118-page-network-servers
```

## Build Docker container and start container shell session:

```bash
docker-compose up -d && docker-compose exec mxprotocol-server bash
```

## Start Mxprotocol Server:

```bash
make clean &&
make &&
./m2m/build/m2m
```


## Configuration

##### - redirect database
For sharing testing data during development, set postgresql service server wherever it is needed.
Change in configuration/lora-app-server.toml
```toml
[postgresql]
dsn="postgres://USERNAME:PASSWORD@SERVICE_SERVER_DOMAIN_NAME:5432/DATABASE_NAME?sslmode=disable"
```

After changing config file, simply restart the service in docker container again

```bash
$ ./m2m/build/m2m -c configuration/mxprotocol-server.toml
```

