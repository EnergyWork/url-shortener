--
-- PostgreSQL database dump
--

-- Dumped from database version 14.2 (Debian 14.2-1.pgdg110+1)
-- Dumped by pg_dump version 14.2

-- Started on 2022-04-03 03:54:12

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
-- TOC entry 210 (class 1259 OID 16395)
-- Name: url_representation; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.url_representation (
    id integer NOT NULL,
    url_long character varying NOT NULL,
    url_short character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    CONSTRAINT url_short_null CHECK (((url_short)::text <> ''::text))
);


ALTER TABLE public.url_representation OWNER TO postgres;

--
-- TOC entry 209 (class 1259 OID 16394)
-- Name: url_representation_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.url_representation_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.url_representation_id_seq OWNER TO postgres;

--
-- TOC entry 3318 (class 0 OID 0)
-- Dependencies: 209
-- Name: url_representation_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.url_representation_id_seq OWNED BY public.url_representation.id;


--
-- TOC entry 3167 (class 2604 OID 16398)
-- Name: url_representation id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.url_representation ALTER COLUMN id SET DEFAULT nextval('public.url_representation_id_seq'::regclass);


--
-- TOC entry 3312 (class 0 OID 16395)
-- Dependencies: 210
-- Data for Name: url_representation; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.url_representation (id, url_long, url_short, created_at) FROM stdin;
\.


--
-- TOC entry 3319 (class 0 OID 0)
-- Dependencies: 209
-- Name: url_representation_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.url_representation_id_seq', 12, true);


--
-- TOC entry 3171 (class 2606 OID 16403)
-- Name: url_representation url_representation_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.url_representation
    ADD CONSTRAINT url_representation_pkey PRIMARY KEY (id);


-- Completed on 2022-04-03 03:54:13

--
-- PostgreSQL database dump complete
--

