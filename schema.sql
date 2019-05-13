--
-- PostgreSQL database dump
--

-- Dumped from database version 10.4 (Debian 10.4-1.pgdg90+1)
-- Dumped by pg_dump version 11.2

-- Started on 2019-05-13 01:00:22 UTC

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

SET default_with_oids = false;

--
-- TOC entry 197 (class 1259 OID 33945)
-- Name: categories; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.categories (
    id integer NOT NULL,
    name character varying(64) NOT NULL
);


--
-- TOC entry 196 (class 1259 OID 33943)
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.categories_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 2894 (class 0 OID 0)
-- Dependencies: 196
-- Name: categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;


--
-- TOC entry 199 (class 1259 OID 33953)
-- Name: hotspots; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.hotspots (
    id integer NOT NULL,
    category_id integer NOT NULL,
    name character varying(64) NOT NULL,
    description character varying(512),
    latitude double precision NOT NULL,
    longitude double precision NOT NULL,
    type character varying(64) NOT NULL
);


--
-- TOC entry 198 (class 1259 OID 33951)
-- Name: hotspots_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.hotspots_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 2895 (class 0 OID 0)
-- Dependencies: 198
-- Name: hotspots_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.hotspots_id_seq OWNED BY public.hotspots.id;


--
-- TOC entry 204 (class 1259 OID 33992)
-- Name: route_points; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.route_points (
    id integer NOT NULL,
    route_id integer NOT NULL,
    latitude double precision NOT NULL,
    longitude double precision NOT NULL
);


--
-- TOC entry 203 (class 1259 OID 33990)
-- Name: route_points_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.route_points_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 2896 (class 0 OID 0)
-- Dependencies: 203
-- Name: route_points_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.route_points_id_seq OWNED BY public.route_points.id;


--
-- TOC entry 202 (class 1259 OID 33976)
-- Name: routes; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.routes (
    id integer NOT NULL,
    category_id integer NOT NULL,
    name character varying(64) NOT NULL,
    description character varying(512)
);


--
-- TOC entry 201 (class 1259 OID 33974)
-- Name: routes_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.routes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 2897 (class 0 OID 0)
-- Dependencies: 201
-- Name: routes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.routes_id_seq OWNED BY public.routes.id;


--
-- TOC entry 200 (class 1259 OID 33967)
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    username character varying(64),
    id uuid NOT NULL,
    email character varying(64) NOT NULL,
    password character varying(128) NOT NULL
);


--
-- TOC entry 2749 (class 2604 OID 33948)
-- Name: categories id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.categories ALTER COLUMN id SET DEFAULT nextval('public.categories_id_seq'::regclass);


--
-- TOC entry 2750 (class 2604 OID 33956)
-- Name: hotspots id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hotspots ALTER COLUMN id SET DEFAULT nextval('public.hotspots_id_seq'::regclass);


--
-- TOC entry 2752 (class 2604 OID 33995)
-- Name: route_points id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.route_points ALTER COLUMN id SET DEFAULT nextval('public.route_points_id_seq'::regclass);


--
-- TOC entry 2751 (class 2604 OID 33979)
-- Name: routes id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.routes ALTER COLUMN id SET DEFAULT nextval('public.routes_id_seq'::regclass);


--
-- TOC entry 2758 (class 2606 OID 33973)
-- Name: users UNIQUE_email; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT "UNIQUE_email" UNIQUE (email);


--
-- TOC entry 2754 (class 2606 OID 33950)
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- TOC entry 2756 (class 2606 OID 33958)
-- Name: hotspots hotspots_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hotspots
    ADD CONSTRAINT hotspots_pkey PRIMARY KEY (id);


--
-- TOC entry 2764 (class 2606 OID 33997)
-- Name: route_points route_points_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.route_points
    ADD CONSTRAINT route_points_pkey PRIMARY KEY (id, route_id);


--
-- TOC entry 2762 (class 2606 OID 33984)
-- Name: routes routes_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.routes
    ADD CONSTRAINT routes_pkey PRIMARY KEY (id);


--
-- TOC entry 2760 (class 2606 OID 33971)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 2765 (class 2606 OID 33959)
-- Name: hotspots FK_category_id; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hotspots
    ADD CONSTRAINT "FK_category_id" FOREIGN KEY (category_id) REFERENCES public.categories(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 2766 (class 2606 OID 33985)
-- Name: routes FK_category_id; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.routes
    ADD CONSTRAINT "FK_category_id" FOREIGN KEY (category_id) REFERENCES public.categories(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 2767 (class 2606 OID 33998)
-- Name: route_points FK_route_id; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.route_points
    ADD CONSTRAINT "FK_route_id" FOREIGN KEY (route_id) REFERENCES public.routes(id) ON UPDATE CASCADE ON DELETE CASCADE;


-- Completed on 2019-05-13 01:00:22 UTC

--
-- PostgreSQL database dump complete
--

