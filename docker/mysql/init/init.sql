CREATE DATABASE test_db;

USE test_db;

CREATE TABLE companies (
  id bigint(20),
  name varchar(255),
  category varchar(255),
  address varchar(255),
  url text,
  employees_num int,
  sns_fb varchar(255),
  sns_tw varchar(255),
  sns_yb varchar(255),
  established_at datetime,
  created_at datetime,
  updated_at datetime
)
