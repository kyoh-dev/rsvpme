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
-- Name: event_log_updated_at(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.event_log_updated_at() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
    BEGIN
        NEW.updated_at = NOW();
        RETURN NEW;
    END;
$$;


--
-- Name: invitee_log_updated_at(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.invitee_log_updated_at() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
    BEGIN
        NEW.updated_at = NOW();
        RETURN NEW;
    END;
$$;


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: event; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.event (
    id integer NOT NULL,
    uuid uuid DEFAULT gen_random_uuid() NOT NULL,
    title text NOT NULL,
    description text NOT NULL,
    start_datetime timestamp without time zone NOT NULL,
    finish_datetime timestamp without time zone,
    address text NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


--
-- Name: invitee; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.invitee (
    id integer NOT NULL,
    first_name text NOT NULL,
    last_name text NOT NULL,
    email text,
    rsvp boolean NOT NULL,
    event_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


--
-- Name: event_detail; Type: VIEW; Schema: public; Owner: -
--

CREATE VIEW public.event_detail AS
 SELECT event.uuid,
    event.title,
    event.description,
    event.start_datetime,
    event.finish_datetime,
    event.address,
    ( SELECT json_agg(json_build_object('first_name', invitee.first_name, 'last_name', invitee.last_name, 'email', invitee.email, 'rsvp', invitee.rsvp)) AS invitees_array
           FROM public.invitee
          WHERE (event.id = invitee.event_id)) AS invitees
   FROM public.event;


--
-- Name: event_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.event ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.event_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: invitee_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.invitee ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.invitee_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version character varying(128) NOT NULL
);


--
-- Name: event event_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event
    ADD CONSTRAINT event_pkey PRIMARY KEY (id);


--
-- Name: invitee invitee_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.invitee
    ADD CONSTRAINT invitee_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: event_uuid_unique_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX event_uuid_unique_idx ON public.event USING btree (uuid);


--
-- Name: event event_log_updated_at; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER event_log_updated_at BEFORE UPDATE ON public.event FOR EACH ROW EXECUTE FUNCTION public.event_log_updated_at();


--
-- Name: invitee invitee_log_updated_at; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER invitee_log_updated_at BEFORE UPDATE ON public.invitee FOR EACH ROW EXECUTE FUNCTION public.invitee_log_updated_at();


--
-- Name: invitee invitee_event_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.invitee
    ADD CONSTRAINT invitee_event_id_fkey FOREIGN KEY (event_id) REFERENCES public.event(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO public.schema_migrations (version) VALUES
    ('20240202061936'),
    ('20240202063939'),
    ('20240207015424');
