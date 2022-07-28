CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

CREATE TABLE public.users
(
    name           text                                   NOT NULL,
    isadmin        boolean                                NOT NULL,
    passwordhash   text                                   NOT NULL,
    registeredfrom integer                                NOT NULL,
    email          text                                   NOT NULL,
    id             uuid DEFAULT public.uuid_generate_v4() NOT NULL PRIMARY KEY
);

CREATE TABLE public.news
(
    title       character varying(200)                       NOT NULL,
    active      boolean                                      NOT NULL,
    activefrom  timestamp with time zone                     NOT NULL,
    text        text                                         NOT NULL,
    textjson    text                                         NOT NULL,
    isimportant boolean                                      NOT NULL,
    id          uuid DEFAULT public.uuid_generate_v4()       NOT NULL PRIMARY KEY,
    author_id   uuid references users (id) on delete cascade NOT NULL
);

CREATE TABLE public.tags
(
    name text                                   NOT NULL unique,
    id   uuid DEFAULT public.uuid_generate_v4() NOT NULL PRIMARY KEY
);

CREATE TABLE public.news_tags
(
    news_id uuid references news (id) on delete cascade NOT NULL,
    tag_id  uuid references tags (id) on delete cascade NOT NULL,
    id      uuid DEFAULT public.uuid_generate_v4()      NOT NULL PRIMARY KEY
);

CREATE TABLE public.refreshsessions
(
    refreshtoken text                                         NOT NULL,
    expiresin    integer                                      NOT NULL,
    createdat    integer                                      NOT NULL,
    id           uuid DEFAULT public.uuid_generate_v4()       NOT NULL PRIMARY KEY,
    user_id      uuid references users (id) on delete cascade NOT NULL
);