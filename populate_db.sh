#!/bin/bash

mysql -u root <<MYSQL_SCRIPT
  USE library;

  INSERT INTO books (isbn, title, author) VALUES
    ('9781503261969', 'Emma', 'Jayne Austen'),
    ('9781505255607', 'The Time Machine', 'H. G. Wells'),
    ('9781503379640', 'The Prince', 'Niccolò Machiavelli');
MYSQL_SCRIPT
