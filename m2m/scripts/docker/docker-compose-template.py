import sys

localTemplate = \
'''
version: "2"
services:
  mxprotocol-server:
    build: 
      context: .
      dockerfile: Dockerfile-devel
    volumes:
      - ./configuration:/etc/mxprotocol-server
      - ./:/mxprotocol-server
    links:
      - postgresql
      - redis
      - mosquitto
    environment:
      - APPSERVER=http://localhost:8080
      - MXPROTOCOL_SERVER=http://localhost:4000
    tty: true
    ports:
      - 4000:4000
    security_opt:
      - seccomp:unconfined
    cap_add:
      - SYS_PTRACE
      
  postgresql:
    image: postgres:9.6-alpine
    volumes:
      - ./.docker-compose/postgresql/initdb:/docker-entrypoint-initdb.d

  redis:
    image: redis:5-alpine

  mosquitto:
    image: eclipse-mosquitto
'''

remoteTemplate = \
'''   
version: "2"
services:
  mxprotocol-server:
    build: 
      context: .
      dockerfile: Dockerfile-devel
    volumes:
      - ./configuration:/etc/mxprotocol-server
      - ./:/mxprotocol-server
    links:
      - redis
    environment:
      - APPSERVER=http://localhost:8080
      - MXPROTOCOL_SERVER=http://localhost:4000
      - REMOTE_SERVER_NAME={}
    tty: true
    ports:
      - 4000:4000
    security_opt:
      - seccomp:unconfined
    cap_add:
      - SYS_PTRACE

  redis:
    image: redis:5-alpine
'''

imageTemplate = \
'''
version: "2"
services:
  mxprotocol-server:
    build: 
      context: .
      dockerfile: Dockerfile
    volumes:
      - ./configuration:/etc/mxprotocol-server
    ports:
      - 4000:4000
    environment:
      - APPSERVER=http://localhost:8080
      - MXPROTOCOL_SERVER=http://localhost:4000
'''

if __name__ == "__main__":
    if len(sys.argv) == 2:
        if sys.argv[1] == "local":
            print(localTemplate)
            exit()

        if sys.argv[1] == "image":
            print(imageTemplate)
            exit()

    if (len(sys.argv) == 3) and (sys.argv[1] == "remote"):
        print(remoteTemplate.format(sys.argv[2]))
        exit()

    print("invalid argument")

