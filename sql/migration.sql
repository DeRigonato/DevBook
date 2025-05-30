CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(100) not null,
    nick varchar(50) not null unique,
    email varchar(100) not null unique,
    senha varchar(100) not null unique, 
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;