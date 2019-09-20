#!/bin/bash

mysql -u root <<MYSQL_SCRIPT
  CREATE DATABASE IF NOT EXISTS library;
  USE library;

  CREATE TABLE IF NOT EXISTS books (
      isbn char(13) NOT NULL,
      title varchar(50) NOT NULL,
      author varchar(50) NOT NULL
  );
MYSQL_SCRIPT
