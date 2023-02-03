DROP DATABASE video_demo;
CREATE DATABASE video_demo;
USE video_demo;

CREATE TABLE IF NOT EXISTS t_user (
    id              BIGINT        NOT NULL AUTO_INCREMENT,
    username        CHAR(20)    NOT NULL,
    follow_count    BIGINT        ,
    follower_count  BIGINT,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS t_video (
    id              BIGINT      NOT NULL AUTO_INCREMENT,
    author_id       BIGINT      NOT NULL,
    favorite_count  BIGINT      NOT NULL,
    comment_count   BIGINT      NOT NULL,
    title           CHAR(50)    NOT NULL,
    publish_date    TIMESTAMP   NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY(author_id) REFERENCES t_user(id)
);