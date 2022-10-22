--
-- PostgreSQL database dump
--

-- Dumped from database version 14.5 (Ubuntu 14.5-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.5 (Ubuntu 14.5-0ubuntu0.22.04.1)

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
-- Name: menu; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.menu (
    id character varying(60) NOT NULL,
    name character varying(100) NOT NULL,
    price integer,
    stock integer,
    image text
);


ALTER TABLE public.menu OWNER TO postgres;

--
-- Name: transaction; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transaction (
    id character varying(60) NOT NULL,
    total_price integer,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone
);


ALTER TABLE public.transaction OWNER TO postgres;

--
-- Name: transaction_detail; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transaction_detail (
    transaction_id character varying(60),
    menu_id character varying(60),
    qty integer,
    subtotal integer,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone
);


ALTER TABLE public.transaction_detail OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    username character varying(32) NOT NULL,
    password text NOT NULL,
    image text
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Data for Name: menu; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.menu (id, name, price, stock, image) FROM stdin;
8503e898-ab0a-4691-80af-4eb04f7065dd	ketoprak	12500	51	8503e898-ab0a-4691-80af-4eb04f7065dd.jpg
aceabd1a-beab-4884-bb7c-3665271a3b10	mie ayam	10000	50	aceabd1a-beab-4884-bb7c-3665271a3b10.jpg
\.


--
-- Data for Name: transaction; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.transaction (id, total_price, created_at, updated_at) FROM stdin;
d6a4738b-7b34-4cdf-93f1-22b7c3779eb3	0	2022-10-16 21:18:59.774226+07	\N
5a959740-0c25-4fee-a2a2-c229bc903427	0	2022-10-16 21:27:00.861721+07	\N
a3f7818f-5072-4655-8067-52fb2b6716bc	0	2022-10-16 21:28:53.038139+07	\N
def29083-d42a-4d80-8012-12be5de8c653	0	2022-10-16 21:29:50.444184+07	\N
a101ac55-07b0-4b3b-af9a-1d61ae084f4e	0	2022-10-16 21:49:38.203464+07	\N
0fe225c2-3e4e-42d3-80bb-8d554e83c63e	0	2022-10-16 21:50:29.970281+07	\N
03342830-1207-42c6-8cba-4a01baadd27e	0	2022-10-16 21:55:48.986808+07	\N
80329264-7a31-40b1-8d28-20b8f0ad62c2	300000	2022-10-16 22:05:25.512189+07	\N
11d624a8-c2c8-4d0b-9c43-17cb9bc63951	45000	2022-10-16 22:07:25.535363+07	\N
89f85e10-4ff1-4090-a81c-e740ca007b2a	45000	2022-10-16 22:18:43.261553+07	\N
82af03c9-309b-4d37-91c8-85ac9f3071d7	0	2022-10-19 11:34:09.417618+07	\N
45c541bd-8b29-4f95-b900-5a871f8dd839	32500	2022-10-19 11:42:19.488093+07	\N
\.


--
-- Data for Name: transaction_detail; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.transaction_detail (transaction_id, menu_id, qty, subtotal, created_at, updated_at) FROM stdin;
5a959740-0c25-4fee-a2a2-c229bc903427	93a23afd-d231-42f9-a8af-7d3d08e32682	10	0	2022-10-16 21:27:00.861721+07	\N
5a959740-0c25-4fee-a2a2-c229bc903427	b53e8ea9-f180-4515-91a4-d1f3c67d3913	10	0	2022-10-16 21:27:00.861721+07	\N
a3f7818f-5072-4655-8067-52fb2b6716bc	93a23afd-d231-42f9-a8af-7d3d08e32682	10	0	2022-10-16 21:28:53.038139+07	\N
a3f7818f-5072-4655-8067-52fb2b6716bc	b53e8ea9-f180-4515-91a4-d1f3c67d3913	10	0	2022-10-16 21:28:53.038139+07	\N
def29083-d42a-4d80-8012-12be5de8c653	93a23afd-d231-42f9-a8af-7d3d08e32682	10	0	2022-10-16 21:29:50.444184+07	\N
def29083-d42a-4d80-8012-12be5de8c653	b53e8ea9-f180-4515-91a4-d1f3c67d3913	10	0	2022-10-16 21:29:50.444184+07	\N
c3dd2d15-f0f6-43a4-9c95-b5be765b27af	93a23afd-d231-42f9-a8af-7d3d08e32682	10	150000	2022-10-16 21:49:38.203464+07	\N
c3dd2d15-f0f6-43a4-9c95-b5be765b27af	b53e8ea9-f180-4515-91a4-d1f3c67d3913	10	150000	2022-10-16 21:49:38.203464+07	\N
0fe225c2-3e4e-42d3-80bb-8d554e83c63e	93a23afd-d231-42f9-a8af-7d3d08e32682	10	150000	2022-10-16 21:50:29.970281+07	\N
0fe225c2-3e4e-42d3-80bb-8d554e83c63e	b53e8ea9-f180-4515-91a4-d1f3c67d3913	10	150000	2022-10-16 21:50:29.970281+07	\N
03342830-1207-42c6-8cba-4a01baadd27e	93a23afd-d231-42f9-a8af-7d3d08e32682	10	150000	2022-10-16 21:55:48.986808+07	\N
03342830-1207-42c6-8cba-4a01baadd27e	b53e8ea9-f180-4515-91a4-d1f3c67d3913	10	150000	2022-10-16 21:55:48.986808+07	\N
80329264-7a31-40b1-8d28-20b8f0ad62c2	93a23afd-d231-42f9-a8af-7d3d08e32682	10	150000	2022-10-16 22:05:25.512189+07	\N
80329264-7a31-40b1-8d28-20b8f0ad62c2	b53e8ea9-f180-4515-91a4-d1f3c67d3913	10	150000	2022-10-16 22:05:25.512189+07	\N
11d624a8-c2c8-4d0b-9c43-17cb9bc63951	93a23afd-d231-42f9-a8af-7d3d08e32682	1	15000	2022-10-16 22:07:25.535363+07	\N
11d624a8-c2c8-4d0b-9c43-17cb9bc63951	b53e8ea9-f180-4515-91a4-d1f3c67d3913	2	30000	2022-10-16 22:07:25.535363+07	\N
89f85e10-4ff1-4090-a81c-e740ca007b2a	93a23afd-d231-42f9-a8af-7d3d08e32682	1	15000	2022-10-16 22:18:43.261553+07	\N
89f85e10-4ff1-4090-a81c-e740ca007b2a	b53e8ea9-f180-4515-91a4-d1f3c67d3913	2	30000	2022-10-16 22:18:43.261553+07	\N
	93a23afd-d231-42f9-a8af-7d3d08e32682	1	0	2022-10-19 11:34:09.417618+07	\N
	b53e8ea9-f180-4515-91a4-d1f3c67d3913	2	0	2022-10-19 11:34:09.417618+07	\N
45c541bd-8b29-4f95-b900-5a871f8dd839	8503e898-ab0a-4691-80af-4eb04f7065dd	1	12500	2022-10-19 11:42:19.488093+07	\N
45c541bd-8b29-4f95-b900-5a871f8dd839	aceabd1a-beab-4884-bb7c-3665271a3b10	2	20000	2022-10-19 11:42:19.488093+07	\N
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, name, username, password, image) FROM stdin;
8ded252a-f510-4a67-b711-73a31cc208be	vuero eruko	vueko	vueko123	8ded252a-f510-4a67-b711-73a31cc208be.jpg
449360c8-914d-4c03-8c2a-d511dcfd4b82	admin	admin	admin123	449360c8-914d-4c03-8c2a-d511dcfd4b82.jpg
4c180e5a-cab4-40f6-855b-af45ab884b05	ueno	ueno	ueno123	4c180e5a-cab4-40f6-855b-af45ab884b05.jpg
\.


--
-- Name: transaction transaction_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transaction
    ADD CONSTRAINT transaction_pkey PRIMARY KEY (id);


--
-- Name: users unique_username; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT unique_username UNIQUE (username);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

