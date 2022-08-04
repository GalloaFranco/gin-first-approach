#!bin/bash
# =============================================================================
#  Author: GalloaFranco
#  Filename: setting-up.sh
#  Description: This script would create the directories, if are need of it, and run the compose file to setting up postgres database
#
# =============================================================================

BLUE='\033[1;34m'
RED='\033[1;31m'
GREEN='\033[1;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

location=$(pwd)
docker=$(command -v docker)
docker_compose=$(command -v docker-compose)

if [ -z $docker ]
then
 echo -e "${YELLOW}[Warning]: You don't have docker installed, is mandatory to have this install and configure in your machine ${NC}"
 exit
fi

if [ -z $docker_compose ]
then
 echo -e "${YELLOW}[Warning]: You don't have docker-compose installed, is mandatory to have this install and configure in your machine ${NC}"
 exit
fi

cd

if [ -d ${HOME}/docker/volumes/data-pg ]
then
  echo -e "${BLUE}[Info] Volume directory done! ${NC}"
else
  echo -e "${BLUE}[Info] Creating docker volumes... ${NC}"
  mkdir -p ~/docker/volumes/data-pg
  echo -e "${GREEN}[Ok] Volume creation done! ${NC}"
fi

cd $location

docker-compose up -d
