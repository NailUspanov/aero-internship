--
-- PostgreSQL database dump
--

-- Dumped from database version 14.3 (Debian 14.3-1.pgdg110+1)
-- Dumped by pg_dump version 14.1

-- Started on 2022-07-28 18:10:50

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 2 (class 3079 OID 16385)
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- TOC entry 3365 (class 0 OID 0)
-- Dependencies: 2
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 210 (class 1259 OID 16400)
-- Name: news; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.news (
    title character varying(200) NOT NULL,
    active boolean NOT NULL,
    activefrom timestamp with time zone NOT NULL,
    text text NOT NULL,
    textjson text NOT NULL,
    isimportant boolean NOT NULL,
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    author_id uuid NOT NULL
);


ALTER TABLE public.news OWNER TO postgres;

--
-- TOC entry 211 (class 1259 OID 16410)
-- Name: news_tags; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.news_tags (
    news_id uuid NOT NULL,
    tag_id uuid NOT NULL,
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL
);


ALTER TABLE public.news_tags OWNER TO postgres;

--
-- TOC entry 212 (class 1259 OID 16414)
-- Name: refreshsessions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.refreshsessions (
    refreshtoken text NOT NULL,
    expiresin integer NOT NULL,
    createdat integer NOT NULL,
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id uuid NOT NULL
);


ALTER TABLE public.refreshsessions OWNER TO postgres;

--
-- TOC entry 213 (class 1259 OID 16420)
-- Name: tags; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tags (
    name text NOT NULL,
    uid uuid DEFAULT public.uuid_generate_v4() NOT NULL
);


ALTER TABLE public.tags OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 16426)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    name text NOT NULL,
    isadmin boolean NOT NULL,
    passwordhash text NOT NULL,
    registeredfrom integer NOT NULL,
    email text NOT NULL,
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 3355 (class 0 OID 16400)
-- Dependencies: 210
-- Data for Name: news; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.news (title, active, activefrom, text, textjson, isimportant, id, author_id) FROM stdin;
\.


--
-- TOC entry 3356 (class 0 OID 16410)
-- Dependencies: 211
-- Data for Name: news_tags; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.news_tags (news_id, tag_id, id) FROM stdin;
\.


--
-- TOC entry 3357 (class 0 OID 16414)
-- Dependencies: 212
-- Data for Name: refreshsessions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.refreshsessions (refreshtoken, expiresin, createdat, id, user_id) FROM stdin;
2f397e55-dbc5-4370-a513-e9ab297e6a03	1661607245	1659015245	cfe696e1-ce86-4fe0-b692-7623e42e4d11	e21e1087-fc7c-4283-8230-8a4a15f946c7
b811fa90-f558-4d74-8997-0c755add41e1	1661609377	1659017377	b7dc013f-4b92-4960-9733-e75b3801ccb0	e21e1087-fc7c-4283-8230-8a4a15f946c7
\.


--
-- TOC entry 3358 (class 0 OID 16420)
-- Dependencies: 213
-- Data for Name: tags; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tags (name, uid) FROM stdin;
\.


--
-- TOC entry 3359 (class 0 OID 16426)
-- Dependencies: 214
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (name, isadmin, passwordhash, registeredfrom, email, id) FROM stdin;
Danil2	f	$2a$04$uCYle5ZkhCqCtfvznnxgEu7Xp16QuJLGU3QxQ3pPgcAmSqOBLiW26	1659014971	email3@mail.com	e21e1087-fc7c-4283-8230-8a4a15f946c7
Danil2	f	$2a$04$t/YyjYfjbXkh8bvgpyofR.M6aTEKuAJjbVaXaNFFjH1/9hQSyeUWi	1659015094	email3@mail.com	3490c62e-175b-430f-9ee1-d90299b8a03d
Danil5	f	$2a$04$R/uTba1sRFqalPnEAYXsxuO9WlFZ4FBGR5xvNHHoSBvymH.an9pgC	1659015175	email3@mail.com	ca2e502a-7050-40ff-9dde-a35117a8d466
Danil6	f	$2a$04$zlSfyCca3zeqWhFA78CeM.LacWswGX/OVjRiXpZkwsyEn64yTPmAG	1659015245	email3@mail.com	6f601a8d-a1f6-40dc-b693-7ceb019b2996
\.


--
-- TOC entry 3199 (class 2606 OID 16437)
-- Name: news news_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.news
    ADD CONSTRAINT news_pkey PRIMARY KEY (id);


--
-- TOC entry 3201 (class 2606 OID 16439)
-- Name: news_tags news_tags_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.news_tags
    ADD CONSTRAINT news_tags_pkey PRIMARY KEY (id);


--
-- TOC entry 3203 (class 2606 OID 16441)
-- Name: refreshsessions refreshsessions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.refreshsessions
    ADD CONSTRAINT refreshsessions_pkey PRIMARY KEY (id);


--
-- TOC entry 3205 (class 2606 OID 16443)
-- Name: refreshsessions refreshsessions_refreshtoken_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.refreshsessions
    ADD CONSTRAINT refreshsessions_refreshtoken_key UNIQUE (refreshtoken);


--
-- TOC entry 3207 (class 2606 OID 16445)
-- Name: tags tags_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_name_key UNIQUE (name);


--
-- TOC entry 3209 (class 2606 OID 16447)
-- Name: tags tags_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_pkey PRIMARY KEY (uid);


--
-- TOC entry 3211 (class 2606 OID 16449)
-- Name: users uuidpk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT uuidpk PRIMARY KEY (id);


--
-- TOC entry 3212 (class 2606 OID 16450)
-- Name: news news_author_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.news
    ADD CONSTRAINT news_author_id_fkey FOREIGN KEY (author_id) REFERENCES public.users(id) NOT VALID;


--
-- TOC entry 3213 (class 2606 OID 16465)
-- Name: news_tags news_tags_news_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.news_tags
    ADD CONSTRAINT news_tags_news_id_fkey FOREIGN KEY (news_id) REFERENCES public.news(id);


--
-- TOC entry 3214 (class 2606 OID 16470)
-- Name: news_tags news_tags_tag_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.news_tags
    ADD CONSTRAINT news_tags_tag_id_fkey FOREIGN KEY (tag_id) REFERENCES public.tags(uid);


--
-- TOC entry 3215 (class 2606 OID 16475)
-- Name: refreshsessions refreshsessions_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.refreshsessions
    ADD CONSTRAINT refreshsessions_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) NOT VALID;


-- Completed on 2022-07-28 18:10:51

--
-- PostgreSQL database dump complete
--

