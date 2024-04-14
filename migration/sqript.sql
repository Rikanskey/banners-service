CREATE DATABASE banners;

create table if not exists tag
(
    id serial
    primary key
);

create table if not exists feature
(
    id serial
    primary key
);

create table if not exists banner
(
    id         serial
    primary key,
    feature_id integer
    references feature
    on update cascade on delete restrict,
    content    json,
    created    timestamp default now() not null,
    updated    timestamp default now() not null,
    is_active  boolean
    );

comment on column banner.created is 'created';

comment on column banner.is_active is 'active';

create table if not exists banner_tag
(
    id        serial
    primary key,
    banner_id integer
    references banner
    on update cascade on delete restrict,
    tag_id    integer
    references tag
    on update cascade on delete restrict
);

INSERT INTO tag (id) VALUES (1), (2), (3), (4), (5), (6), (7), (8), (9), (10);

INSERT INTO feature (id) VALUES (1), (2), (3), (4), (5);

INSERT INTO banner (feature_id, content, created, updated, is_active) VALUES (1, '{"a": "b"}', now(), now(), true),
                                                                             (1, '{"bf": "f"}', now(), now(), true),
                                                                             (2, '{"d": "b"}', now(), now(), true),
                                                                             (3, '{"a": "b"}', now(), now(), true);

INSERT INTO banner_tag (banner_id, tag_id) VALUES (1, 1), (1, 2), (1, 3), (2, 3), (3, 1);