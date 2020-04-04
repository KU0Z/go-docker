-- Database: users_database

-- DROP DATABASE users_database;

CREATE DATABASE users_database
    WITH OWNER = demo
       ENCODING = 'UTF8'
       CONNECTION LIMIT = -1;

\connect users_database 
-- Table: users

-- DROP TABLE users;

CREATE TABLE users
(
 id serial NOT NULL,
 name character varying NOT NULL,
 lastname character varying ,
 faculty character varying,
 carer character varying,
 carne integer,
 publication_date date,
 CONSTRAINT pk_users PRIMARY KEY (id)
)
WITH (
 OIDS=FALSE
);