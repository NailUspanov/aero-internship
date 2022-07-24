--
-- PostgreSQL database dump
--

-- Dumped from database version 14.3 (Debian 14.3-1.pgdg110+1)
-- Dumped by pg_dump version 14.1

-- Started on 2022-07-24 22:20:24

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 209 (class 1259 OID 16385)
-- Name: files; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.files (
    id integer NOT NULL,
    user_id integer NOT NULL,
    name character varying(200) NOT NULL,
    ext character varying(15) NOT NULL,
    base64 character varying(172) NOT NULL,
    datecreate timestamp with time zone NOT NULL
);


ALTER TABLE public.files OWNER TO postgres;

--
-- TOC entry 210 (class 1259 OID 16388)
-- Name: files_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.files ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.files_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 211 (class 1259 OID 16389)
-- Name: news; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.news (
    id integer NOT NULL,
    title character varying(200) NOT NULL,
    author_id integer NOT NULL,
    active boolean NOT NULL,
    activefrom timestamp with time zone NOT NULL,
    text text NOT NULL,
    textjson text NOT NULL,
    tags integer[] NOT NULL,
    files integer[] NOT NULL,
    isimportant boolean NOT NULL
);


ALTER TABLE public.news OWNER TO postgres;

--
-- TOC entry 212 (class 1259 OID 16394)
-- Name: news_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.news ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.news_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 218 (class 1259 OID 16428)
-- Name: refreshsessions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.refreshsessions (
    id integer NOT NULL,
    userid integer,
    refreshtoken text NOT NULL,
    expiresin integer NOT NULL,
    createdat integer NOT NULL
);


ALTER TABLE public.refreshsessions OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 16427)
-- Name: refreshsessions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.refreshsessions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.refreshsessions_id_seq OWNER TO postgres;

--
-- TOC entry 3359 (class 0 OID 0)
-- Dependencies: 217
-- Name: refreshsessions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.refreshsessions_id_seq OWNED BY public.refreshsessions.id;


--
-- TOC entry 213 (class 1259 OID 16395)
-- Name: tags; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tags (
    id integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE public.tags OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 16400)
-- Name: tags_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.tags ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.tags_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 215 (class 1259 OID 16401)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    name text NOT NULL,
    isadmin boolean NOT NULL,
    passwordhash text NOT NULL,
    registeredfrom integer NOT NULL,
    email text NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 216 (class 1259 OID 16406)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.users ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 3187 (class 2604 OID 16431)
-- Name: refreshsessions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.refreshsessions ALTER COLUMN id SET DEFAULT nextval('public.refreshsessions_id_seq'::regclass);


--
-- TOC entry 3344 (class 0 OID 16385)
-- Dependencies: 209
-- Data for Name: files; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.files (id, user_id, name, ext, base64, datecreate) FROM stdin;
\.


--
-- TOC entry 3346 (class 0 OID 16389)
-- Dependencies: 211
-- Data for Name: news; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.news (id, title, author_id, active, activefrom, text, textjson, tags, files, isimportant) FROM stdin;
\.


--
-- TOC entry 3353 (class 0 OID 16428)
-- Dependencies: 218
-- Data for Name: refreshsessions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.refreshsessions (id, userid, refreshtoken, expiresin, createdat) FROM stdin;
1	1	bea28334-f203-4ef7-aff8-99d754a69e1a	1661268568	1658676568
2	1	ef1632c8-f9c5-40ea-b5d6-e3296836cfd8	1661272629	1658680629
3	1	e92dc938-9fe2-473a-bce2-2475001fd71c	1661274190	1658682190
4	1	341c55a1-dab0-4f01-87f4-5954ec1a2d57	1661275289	1658683289
5	1	6f0a0ef8-e053-4646-bd9c-5a27dec72f15	1661278102	1658686102
\.


--
-- TOC entry 3348 (class 0 OID 16395)
-- Dependencies: 213
-- Data for Name: tags; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tags (id, name) FROM stdin;
\.


--
-- TOC entry 3350 (class 0 OID 16401)
-- Dependencies: 215
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, name, isadmin, passwordhash, registeredfrom, email) FROM stdin;
1	Danil	f	$2a$04$oYUK3Wyiwb5Mwbw0cDspJOIzSWhecI/gj5vfhKx13hKtaPIGN6Cte	1658676410	email2@mail.com
2	Danil	f	$2a$04$RqBVvkS3GS.M6gLHJdUyq.TLUHTUMODaVTi1JJeSqEkkgmccVcHkm	1658676568	email2@mail.com
\.


--
-- TOC entry 3360 (class 0 OID 0)
-- Dependencies: 210
-- Name: files_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.files_id_seq', 1, false);


--
-- TOC entry 3361 (class 0 OID 0)
-- Dependencies: 212
-- Name: news_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.news_id_seq', 1, false);


--
-- TOC entry 3362 (class 0 OID 0)
-- Dependencies: 217
-- Name: refreshsessions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.refreshsessions_id_seq', 5, true);


--
-- TOC entry 3363 (class 0 OID 0)
-- Dependencies: 214
-- Name: tags_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.tags_id_seq', 1, false);


--
-- TOC entry 3364 (class 0 OID 0)
-- Dependencies: 216
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 2, true);


--
-- TOC entry 3189 (class 2606 OID 16408)
-- Name: files files_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.files
    ADD CONSTRAINT files_pkey PRIMARY KEY (id);


--
-- TOC entry 3191 (class 2606 OID 16410)
-- Name: news news_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.news
    ADD CONSTRAINT news_pkey PRIMARY KEY (id);


--
-- TOC entry 3199 (class 2606 OID 16435)
-- Name: refreshsessions refreshsessions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.refreshsessions
    ADD CONSTRAINT refreshsessions_pkey PRIMARY KEY (id);


--
-- TOC entry 3201 (class 2606 OID 16437)
-- Name: refreshsessions refreshsessions_refreshtoken_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.refreshsessions
    ADD CONSTRAINT refreshsessions_refreshtoken_key UNIQUE (refreshtoken);


--
-- TOC entry 3193 (class 2606 OID 16412)
-- Name: tags tags_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_name_key UNIQUE (name);


--
-- TOC entry 3195 (class 2606 OID 16414)
-- Name: tags tags_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_pkey PRIMARY KEY (id);


--
-- TOC entry 3197 (class 2606 OID 16416)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 3202 (class 2606 OID 16417)
-- Name: files files_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.files
    ADD CONSTRAINT files_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3203 (class 2606 OID 16422)
-- Name: news news_author_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.news
    ADD CONSTRAINT news_author_id_fkey FOREIGN KEY (author_id) REFERENCES public.users(id);


--
-- TOC entry 3204 (class 2606 OID 16438)
-- Name: refreshsessions refreshsessions_userid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.refreshsessions
    ADD CONSTRAINT refreshsessions_userid_fkey FOREIGN KEY (userid) REFERENCES public.users(id);


-- Completed on 2022-07-24 22:20:25

--
-- PostgreSQL database dump complete
--

