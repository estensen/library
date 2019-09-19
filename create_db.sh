#!/bin/bash

mysql -u root <<MYSQL_SCRIPT
  CREATE DATABASE IF NOT EXISTS library;
  USE library;

  CREATE TABLE IF NOT EXISTS books (
      isbn char(13),
      title varchar(50),
      author varchar(50)
  );
MYSQL_SCRIPT
