#!/bin/bash
tar -xvf backup_veterinaria_citiaps.tar.gz -C /backup/
mongorestore /backup/
