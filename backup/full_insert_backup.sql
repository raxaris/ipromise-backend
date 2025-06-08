--
-- PostgreSQL database dump
--

-- Dumped from database version 16.9 (Debian 16.9-1.pgdg120+1)
-- Dumped by pg_dump version 16.9 (Debian 16.9-1.pgdg120+1)

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
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: attachments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.attachments (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    post_id uuid,
    user_id uuid,
    attachment_type character varying(20),
    file_url text NOT NULL,
    file_type character varying(50) NOT NULL,
    created_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.attachments OWNER TO postgres;

--
-- Name: badges; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.badges (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    code character varying(50) NOT NULL,
    title character varying(100) NOT NULL,
    description character varying(255),
    icon_url character varying(255),
    created_at timestamp with time zone
);


ALTER TABLE public.badges OWNER TO postgres;

--
-- Name: followers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.followers (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    follower_id uuid NOT NULL,
    following_id uuid NOT NULL,
    status character varying(20) NOT NULL,
    created_at timestamp with time zone
);


ALTER TABLE public.followers OWNER TO postgres;

--
-- Name: likes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.likes (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id uuid NOT NULL,
    post_id uuid NOT NULL,
    created_at timestamp with time zone
);


ALTER TABLE public.likes OWNER TO postgres;

--
-- Name: microtasks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.microtasks (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    promise_id uuid NOT NULL,
    title character varying(255) NOT NULL,
    steps_planned bigint DEFAULT 1,
    status character varying(20) NOT NULL,
    microtask_order bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.microtasks OWNER TO postgres;

--
-- Name: notifications; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.notifications (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id uuid NOT NULL,
    type text NOT NULL,
    message text NOT NULL,
    is_read boolean DEFAULT false,
    related_id uuid,
    created_at timestamp with time zone
);


ALTER TABLE public.notifications OWNER TO postgres;

--
-- Name: posts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.posts (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id uuid NOT NULL,
    promise_id uuid NOT NULL,
    microtask_id uuid NOT NULL,
    parent_id uuid,
    content text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.posts OWNER TO postgres;

--
-- Name: predictions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.predictions (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    promise_id uuid NOT NULL,
    success_rate numeric NOT NULL,
    advice text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.predictions OWNER TO postgres;

--
-- Name: promises; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.promises (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id uuid NOT NULL,
    title character varying(255) NOT NULL,
    description text,
    deadline timestamp with time zone NOT NULL,
    is_private boolean DEFAULT false NOT NULL,
    status character varying(20) NOT NULL,
    category character varying(30),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.promises OWNER TO postgres;

--
-- Name: refresh_tokens; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.refresh_tokens (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id uuid NOT NULL,
    token text NOT NULL,
    expires_at timestamp with time zone NOT NULL,
    created_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.refresh_tokens OWNER TO postgres;

--
-- Name: user_badges; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_badges (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id uuid NOT NULL,
    badge_id uuid NOT NULL,
    awarded_at timestamp with time zone
);


ALTER TABLE public.user_badges OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    username character varying(50) NOT NULL,
    email character varying(100) NOT NULL,
    password text NOT NULL,
    role character varying(15) DEFAULT 'user'::character varying NOT NULL,
    avatar_url text,
    bio character varying(160),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Data for Name: attachments; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.attachments VALUES ('e23914f0-75d8-4519-8ebb-d1114b6c3254', NULL, '7e4a119f-4121-424e-af44-a9e97cbae049', 'avatar', 'http://localhost:9000/ipromise/4478dce2-4b15-413b-8813-639e8170fb5f.png', 'image/png', '2025-06-06 12:47:36.940883+00', NULL);
INSERT INTO public.attachments VALUES ('4d77930b-b331-4f6d-b138-0d597f15d3b0', '23ad69f9-8c1f-45e1-bdb6-86ae97f0a2d2', NULL, 'post_image', 'http://localhost:9000/ipromise/490e1f2c-3323-414b-b0e1-2cbb80ed173d.jpeg', 'image/jpeg', '2025-06-07 16:15:30.660681+00', NULL);
INSERT INTO public.attachments VALUES ('e430a0b7-03ac-4aa9-8139-0928a31e4dd4', '741d9f45-74da-4da2-b40e-06b325c5342f', NULL, 'post_image', 'http://localhost:9000/ipromise/151b108a-054a-43ca-bbdd-f2df0566ea86.jpeg', 'image/jpeg', '2025-06-07 16:20:17.666089+00', NULL);
INSERT INTO public.attachments VALUES ('d12fa1b9-29ba-43c2-bbe5-4b6990a1f0fd', '60333cb2-3eac-4c88-b1ab-187788f3a753', NULL, 'post_image', 'http://localhost:9000/ipromise/b3f04097-83e5-4fd8-873c-8d9c48c76338.jpeg', 'image/jpeg', '2025-06-07 16:23:59.439292+00', NULL);
INSERT INTO public.attachments VALUES ('e31356b1-00d7-4fe9-81e7-21e1ea39a141', '31f82a0d-869c-4ea9-8c72-152d229893cf', NULL, 'post_image', 'http://localhost:9000/ipromise/c16494e9-842f-4c57-b270-257a05282b87.png', 'image/png', '2025-06-07 16:49:02.689124+00', NULL);
INSERT INTO public.attachments VALUES ('ea0b6ca3-7b06-44ef-889e-97bb687eb6f3', '223da9a0-82b1-4745-9a4a-5b657d47299e', NULL, 'post_image', 'http://localhost:9000/ipromise/a1af8183-b41a-4069-9c4a-f47bc02ec65d.jpg', 'image/jpeg', '2025-06-07 18:31:21.425609+00', NULL);
INSERT INTO public.attachments VALUES ('67da62a4-f159-4e0d-94b5-05ddf10d7b72', 'a3f8de46-d113-48a1-b569-a20561d328f5', NULL, 'post_image', 'http://localhost:9000/ipromise/2e04b9e8-805f-419e-8825-091950d86313.jpeg', 'image/jpeg', '2025-06-07 18:33:11.6177+00', NULL);
INSERT INTO public.attachments VALUES ('e3cb3985-9f10-44fe-ae61-9f70deea6d8a', 'b3621746-9032-46c1-9e9c-072c55b709a7', NULL, 'post_image', 'http://localhost:9000/ipromise/c10f2317-0330-40ee-ba63-25c9240416bd.png', 'image/png', '2025-06-07 18:42:48.861851+00', NULL);
INSERT INTO public.attachments VALUES ('2fe0f1b5-6ed0-4b65-98c8-4bc305f231e1', '5931007c-9050-44ae-90af-0c9732651ff3', NULL, 'post_image', 'http://localhost:9000/ipromise/6c27818a-f165-4c87-91ce-4972b81b14dc.jpg', 'image/jpeg', '2025-06-07 19:11:55.902955+00', NULL);
INSERT INTO public.attachments VALUES ('199d8280-ba24-46a6-8e52-a6c37e821697', '660c8c08-468b-4a0d-a37d-fe58d7d0b075', NULL, 'post_image', 'http://localhost:9000/ipromise/6ff6a754-e644-4230-9f69-dbb63111636b.jpg', 'image/jpeg', '2025-06-07 19:20:10.473915+00', NULL);
INSERT INTO public.attachments VALUES ('656e329d-421f-45fc-addd-3a0ae4091cd0', 'c4e4435b-e210-481c-9603-880796a0b7df', NULL, 'post_image', 'http://localhost:9000/ipromise/6b328193-df3f-4e0f-b70e-881df4f5d348.png', 'image/png', '2025-06-07 19:22:43.05936+00', NULL);
INSERT INTO public.attachments VALUES ('01cedf4a-459f-4f5f-ac1a-d95e3ee1a565', '6a82001d-b416-4515-913e-5b6e3f57fbdf', NULL, 'post_image', 'http://localhost:9000/ipromise/afc95188-62eb-4ffa-b314-88aa5e601815.jpg', 'image/jpeg', '2025-06-07 19:23:48.4721+00', NULL);
INSERT INTO public.attachments VALUES ('ff3bc8d9-16fa-486f-aa70-8268d3e62d11', NULL, 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'avatar', 'http://localhost:9000/ipromise/e2d9aaaf-0998-4ab5-a4f5-e3e826a7b90b.jpg', 'image/jpeg', '2025-06-08 16:32:20.701731+00', NULL);
INSERT INTO public.attachments VALUES ('3c429f8f-ba81-42dc-8e54-961876cac126', NULL, '3b0ebec9-3805-490b-987f-4c6e7286319e', 'avatar', 'http://localhost:9000/ipromise/8f0e65a2-569d-4660-876c-e34f820b1d22.jpg', 'image/jpeg', '2025-06-08 16:40:27.81271+00', NULL);
INSERT INTO public.attachments VALUES ('87607e12-3fa4-4e00-a7b2-de320d7d6108', NULL, '56e89866-3a90-4f97-8461-34b24f842da4', 'avatar', 'http://localhost:9000/ipromise/e6e344b9-3bce-4ffd-a17d-b5cc14b39c74.jpg', 'image/jpeg', '2025-06-08 16:40:57.004486+00', NULL);
INSERT INTO public.attachments VALUES ('5b15e21f-eb9c-44d1-88a7-91fd3bf48263', NULL, '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'avatar', 'http://localhost:9000/ipromise/8f250274-a5cf-400b-95d8-a31e4c38955f.jpg', 'image/jpeg', '2025-06-08 16:42:38.038403+00', NULL);
INSERT INTO public.attachments VALUES ('4713bc5c-c4a4-43f1-bc7b-4308c6bc60e4', NULL, '6e137217-71a9-439b-a579-4bcd1a729af9', 'avatar', 'http://localhost:9000/ipromise/6ea78a78-8ac9-491d-8cfe-ed485e5f4c01.jpg', 'image/jpeg', '2025-06-08 16:43:02.026153+00', NULL);
INSERT INTO public.attachments VALUES ('e9d90b7a-c572-41b4-9827-4c60ac5889d9', NULL, '79758f94-70d8-4a11-b05f-793fcb57e093', 'avatar', 'http://localhost:9000/ipromise/888a18c4-ac96-43a7-98a9-1c6b45d520ff.jpg', 'image/jpeg', '2025-06-08 16:43:28.029193+00', NULL);
INSERT INTO public.attachments VALUES ('2a326447-603e-4749-bd6b-096c95f2a378', NULL, 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'avatar', 'http://localhost:9000/ipromise/a772e9c4-6b63-47d8-8d8d-49b95a630e21.jpg', 'image/jpeg', '2025-06-08 16:43:51.685717+00', NULL);
INSERT INTO public.attachments VALUES ('f0a29130-b8d6-4e73-b356-8c5092d3fd79', NULL, 'a203c89a-3ab9-42b2-ae92-0c979d8425de', 'avatar', 'http://localhost:9000/ipromise/af40e94d-8961-45b6-95cc-13fdff717704.jpg', 'image/jpeg', '2025-06-08 16:44:12.164318+00', NULL);
INSERT INTO public.attachments VALUES ('47c6e858-e769-410b-9931-822aec863290', NULL, 'ae884ed2-075a-4ed2-9dde-736921d80e9f', 'avatar', 'http://localhost:9000/ipromise/019511ab-4a38-4a10-9e0d-b6bdb8c6befd.jpg', 'image/jpeg', '2025-06-08 16:44:44.01106+00', NULL);
INSERT INTO public.attachments VALUES ('6764451f-47c1-471c-a6cd-b0f07b72f005', NULL, 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', 'avatar', 'http://localhost:9000/ipromise/0a5f1083-c1bc-43da-bab3-7aeb2ffee375.jpg', 'image/jpeg', '2025-06-08 16:45:10.231864+00', NULL);
INSERT INTO public.attachments VALUES ('6bc218e1-38fa-4a77-8546-af274c48d7b7', NULL, 'd943512c-a462-4f9e-8881-13dc2b9b59e1', 'avatar', 'http://localhost:9000/ipromise/7b21eafa-b1bf-4745-be9e-59b2e8e838c0.jpg', 'image/jpeg', '2025-06-08 16:45:43.18496+00', NULL);
INSERT INTO public.attachments VALUES ('78403fda-97c0-4c54-a97a-6e8f725e69d6', NULL, 'f399a996-1567-4fc0-a86a-92774494ca44', 'avatar', 'http://localhost:9000/ipromise/80ed7c51-695d-4f6a-ad7b-ecb3f79a44ba.jpg', 'image/jpeg', '2025-06-08 16:46:10.473935+00', NULL);
INSERT INTO public.attachments VALUES ('5e6053f4-6050-4440-a1a3-b3cd9d32aa2e', NULL, '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'avatar', 'http://localhost:9000/ipromise/44ae11d5-10f6-43b1-9ef8-9bc12283faa7.jpg', 'image/jpeg', '2025-06-08 16:47:03.423807+00', NULL);


--
-- Data for Name: badges; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.badges VALUES ('4769b089-f1c3-4cb6-88af-7f53614d1e6d', 'post_1', 'First post', 'You created your first post!', 'http://localhost:9000/ipromise/1stlevel.png', '2025-05-18 15:08:17.169147+00');
INSERT INTO public.badges VALUES ('4fe90924-2c53-4882-b404-e27d066ef682', 'post_50', '50 posts', '50 posts! You''re a legend!', 'http://localhost:9000/ipromise/4rthlevel.png', '2025-05-18 15:08:17.169147+00');
INSERT INTO public.badges VALUES ('5db810f1-1a7f-4022-9ca1-85b3b727db56', 'friends_5', '5 friends', '5 friends! You''re not alone anymore!', 'http://localhost:9000/ipromise/2ndlevel.png', '2025-05-18 15:08:17.169147+00');
INSERT INTO public.badges VALUES ('6eb08fda-7fae-42fc-a7e9-8c2aa4df7d35', 'friends_15', '15 friends', '15 friends! You''re the life of the party', 'http://localhost:9000/ipromise/4rthlevel.png', '2025-05-18 15:08:17.169147+00');
INSERT INTO public.badges VALUES ('712e8ca2-37e2-46eb-98e6-1615f66d0fd7', 'friends_1', 'First friend', 'You made your first friend!', 'http://localhost:9000/ipromise/1stlevel.png', '2025-05-18 15:08:17.169147+00');
INSERT INTO public.badges VALUES ('a725e609-3804-43d6-ada5-f17e92458f6c', 'promise_5', '5 promises', 'You''ve already made 5 promises!', 'http://localhost:9000/ipromise/2ndlevel.png', '2025-05-18 15:08:17.169147+00');
INSERT INTO public.badges VALUES ('d67d0bd3-06ad-4f4c-bf43-09ebcb3240a0', 'post_10', '10 posts', 'You published 10 posts!', 'http://localhost:9000/ipromise/2ndlevel.png', '2025-05-18 15:08:17.169147+00');
INSERT INTO public.badges VALUES ('f0963252-7299-43e2-8849-6cce3b3ea52a', 'promise_1', 'First promise', 'Your first promise is created!', 'http://localhost:9000/ipromise/1stlevel.png', '2025-05-18 15:08:17.169147+00');
INSERT INTO public.badges VALUES ('f926bd1b-370b-40b0-854a-51a036c346e2', 'promise_15', '15 promises', '15 promises! You''re a motivator!', 'http://localhost:9000/ipromise/4rthlevel.png', '2025-05-18 15:08:17.169147+00');


--
-- Data for Name: followers; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.followers VALUES ('6d3bffeb-e1ec-4387-aae4-fcd5ecc8c6ef', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '6e137217-71a9-439b-a579-4bcd1a729af9', 'accepted', '2025-06-08 16:46:53.481291+00');
INSERT INTO public.followers VALUES ('0c2d3b79-42d5-4561-af0c-7c597eaf1a9a', '6e137217-71a9-439b-a579-4bcd1a729af9', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'accepted', '2025-06-08 16:54:38.481291+00');
INSERT INTO public.followers VALUES ('7232dc07-17e0-45d9-aa39-eb652143e96c', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'f399a996-1567-4fc0-a86a-92774494ca44', 'accepted', '2025-06-08 17:29:07.481291+00');
INSERT INTO public.followers VALUES ('d002ebc6-884f-4565-95e1-dadcf3c7373b', 'f399a996-1567-4fc0-a86a-92774494ca44', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'accepted', '2025-06-08 17:32:06.481291+00');
INSERT INTO public.followers VALUES ('db028601-4fb3-4a34-a346-525d96def010', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', 'accepted', '2025-06-08 17:40:19.481291+00');
INSERT INTO public.followers VALUES ('6eb625a5-b4e8-4a3d-af73-1680f08f4012', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'accepted', '2025-06-08 17:46:28.481291+00');
INSERT INTO public.followers VALUES ('38c69390-c34d-41ea-af69-d310d75324bf', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'accepted', '2025-06-08 15:19:49.481291+00');
INSERT INTO public.followers VALUES ('98ff9cbd-f1fc-4b03-bb59-f7e7fad92c9c', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'accepted', '2025-06-08 15:27:39.481291+00');
INSERT INTO public.followers VALUES ('8a856465-9f8b-4b46-a952-03337198ffb8', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', 'pending', '2025-06-08 20:33:53.481291+00');
INSERT INTO public.followers VALUES ('9901861d-2157-40d9-a742-9a72b5ae5415', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '56e89866-3a90-4f97-8461-34b24f842da4', 'pending', '2025-06-08 19:13:00.481291+00');
INSERT INTO public.followers VALUES ('ef662b34-fddb-4f96-aedb-7bea173adc57', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', 'accepted', '2025-06-08 15:32:06.481291+00');
INSERT INTO public.followers VALUES ('ffa81b92-e9ec-4553-b8bd-0cf16b1dc2d0', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'accepted', '2025-06-08 15:48:19.481291+00');
INSERT INTO public.followers VALUES ('95f99121-025c-4cc2-b8dd-ceda1c4a5769', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', 'accepted', '2025-06-08 17:58:34.481291+00');
INSERT INTO public.followers VALUES ('997ba12b-82b3-4842-ba59-d8473401c698', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'accepted', '2025-06-08 18:12:39.481291+00');
INSERT INTO public.followers VALUES ('1d1627e4-f516-4357-abe5-32420cd1867d', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'accepted', '2025-06-08 15:13:23.481291+00');
INSERT INTO public.followers VALUES ('20b97815-a33b-4158-9b82-dabfdbb410ec', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'accepted', '2025-06-08 15:14:38.481291+00');
INSERT INTO public.followers VALUES ('db6010ca-465b-4d0b-a668-c7c93e9684d4', '3b0ebec9-3805-490b-987f-4c6e7286319e', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'accepted', '2025-06-08 16:55:06.481291+00');
INSERT INTO public.followers VALUES ('1bfc223a-d990-460b-a866-8d1e91bffdc2', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'accepted', '2025-06-08 16:55:11.481291+00');
INSERT INTO public.followers VALUES ('9e78e136-fa91-452a-b1f7-8991f6ed8d0a', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', 'pending', '2025-06-08 18:35:55.481291+00');
INSERT INTO public.followers VALUES ('01fae9bc-29f7-472d-b0dd-ce3f2dcd81ca', '3b0ebec9-3805-490b-987f-4c6e7286319e', '79758f94-70d8-4a11-b05f-793fcb57e093', 'pending', '2025-06-08 18:08:22.481291+00');
INSERT INTO public.followers VALUES ('7381fa68-fb7b-4214-9f58-214aa19b0aa8', '56e89866-3a90-4f97-8461-34b24f842da4', '79758f94-70d8-4a11-b05f-793fcb57e093', 'accepted', '2025-06-08 17:47:43.481291+00');
INSERT INTO public.followers VALUES ('a0b574ef-824d-4bd5-8dca-14f5d9fe1ea8', '79758f94-70d8-4a11-b05f-793fcb57e093', '56e89866-3a90-4f97-8461-34b24f842da4', 'accepted', '2025-06-08 18:02:59.481291+00');
INSERT INTO public.followers VALUES ('bc8cc456-1ff0-4c6e-b988-7f268461c922', '56e89866-3a90-4f97-8461-34b24f842da4', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'accepted', '2025-06-08 17:08:18.481291+00');
INSERT INTO public.followers VALUES ('afe2b316-302b-465e-82d6-643685646133', '56e89866-3a90-4f97-8461-34b24f842da4', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', 'accepted', '2025-06-08 17:30:31.481291+00');
INSERT INTO public.followers VALUES ('18261ff6-e313-4f30-9a94-34c568069322', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '56e89866-3a90-4f97-8461-34b24f842da4', 'accepted', '2025-06-08 17:37:15.481291+00');
INSERT INTO public.followers VALUES ('91a83f38-8abd-4890-882f-c3147be995fe', '56e89866-3a90-4f97-8461-34b24f842da4', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', 'pending', '2025-06-08 20:14:10.481291+00');
INSERT INTO public.followers VALUES ('b9752e6a-707a-45c9-bf68-8fa61c6ab95f', '56e89866-3a90-4f97-8461-34b24f842da4', 'f399a996-1567-4fc0-a86a-92774494ca44', 'pending', '2025-06-08 17:59:09.481291+00');
INSERT INTO public.followers VALUES ('0542249c-7ffa-49d5-98c4-413af9fec41d', '6e137217-71a9-439b-a579-4bcd1a729af9', '79758f94-70d8-4a11-b05f-793fcb57e093', 'accepted', '2025-06-08 16:04:32.481291+00');
INSERT INTO public.followers VALUES ('e5c037c1-f016-4122-b2c9-6568f103f656', '79758f94-70d8-4a11-b05f-793fcb57e093', '6e137217-71a9-439b-a579-4bcd1a729af9', 'accepted', '2025-06-08 16:04:44.481291+00');
INSERT INTO public.followers VALUES ('e947a942-b74d-42ee-9193-008969c7a76b', '6e137217-71a9-439b-a579-4bcd1a729af9', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', 'accepted', '2025-06-08 16:03:10.481291+00');
INSERT INTO public.followers VALUES ('5a725d04-e673-4631-b416-8e518daab9f8', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', '6e137217-71a9-439b-a579-4bcd1a729af9', 'accepted', '2025-06-08 16:08:40.481291+00');
INSERT INTO public.followers VALUES ('c3313a6b-6f75-4a29-b093-ec54b2eb3f4c', '6e137217-71a9-439b-a579-4bcd1a729af9', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'accepted', '2025-06-08 17:27:00.481291+00');
INSERT INTO public.followers VALUES ('b3ece98b-c567-4af9-89cd-1cce69b1e56a', '3b0ebec9-3805-490b-987f-4c6e7286319e', '6e137217-71a9-439b-a579-4bcd1a729af9', 'accepted', '2025-06-08 17:40:21.481291+00');
INSERT INTO public.followers VALUES ('e5685c74-ee06-40e5-a864-0aedd030a991', '6e137217-71a9-439b-a579-4bcd1a729af9', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', 'pending', '2025-06-08 19:29:09.481291+00');
INSERT INTO public.followers VALUES ('e199df5f-95a8-4e13-a6f1-210d5ea2351b', '6e137217-71a9-439b-a579-4bcd1a729af9', '56e89866-3a90-4f97-8461-34b24f842da4', 'pending', '2025-06-08 19:26:52.481291+00');
INSERT INTO public.followers VALUES ('b79323d4-fcfe-4074-9ea8-e416ab1201a1', '79758f94-70d8-4a11-b05f-793fcb57e093', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', 'accepted', '2025-06-08 17:35:01.481291+00');
INSERT INTO public.followers VALUES ('36fd031d-94c7-4d41-8519-59182caa3c4a', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '79758f94-70d8-4a11-b05f-793fcb57e093', 'accepted', '2025-06-08 17:37:41.481291+00');
INSERT INTO public.followers VALUES ('2a463547-4022-4db6-9ea6-cbfcb0445b52', '79758f94-70d8-4a11-b05f-793fcb57e093', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'accepted', '2025-06-08 17:12:35.481291+00');
INSERT INTO public.followers VALUES ('3af993cb-a7c0-4dd4-992e-5739ebaa79eb', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '79758f94-70d8-4a11-b05f-793fcb57e093', 'accepted', '2025-06-08 17:19:18.481291+00');
INSERT INTO public.followers VALUES ('6a0f4302-9cb6-43c6-8781-5673858e5dc4', '79758f94-70d8-4a11-b05f-793fcb57e093', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'accepted', '2025-06-08 16:03:17.481291+00');
INSERT INTO public.followers VALUES ('a8f3ce10-9011-412a-92b9-463e87d0fbe1', '79758f94-70d8-4a11-b05f-793fcb57e093', 'f399a996-1567-4fc0-a86a-92774494ca44', 'pending', '2025-06-08 20:36:34.481291+00');
INSERT INTO public.followers VALUES ('494215d0-d8f2-4ae9-bc3a-6a9dfb6a941c', '79758f94-70d8-4a11-b05f-793fcb57e093', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', 'pending', '2025-06-08 20:28:04.481291+00');
INSERT INTO public.followers VALUES ('47ea85db-1a5b-4031-8323-6125afd2c1ab', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '6e137217-71a9-439b-a579-4bcd1a729af9', 'accepted', '2025-06-08 15:25:28.481291+00');
INSERT INTO public.followers VALUES ('1d5d49f4-e6f1-48ae-a718-e04d4a85f004', '6e137217-71a9-439b-a579-4bcd1a729af9', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'accepted', '2025-06-08 15:41:56.481291+00');
INSERT INTO public.followers VALUES ('f995084b-d73d-41ff-bc4e-fb1d625a2571', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', 'pending', '2025-06-08 19:40:36.481291+00');
INSERT INTO public.followers VALUES ('c6ef739d-d2df-4073-a493-66da2f0c734a', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', 'pending', '2025-06-08 20:38:40.481291+00');
INSERT INTO public.followers VALUES ('5322de2e-3917-4110-94f4-1ac24f90c65b', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', 'accepted', '2025-06-08 15:31:01.481291+00');
INSERT INTO public.followers VALUES ('fcc10f9d-8641-4fea-b64c-2c1d14217292', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', 'accepted', '2025-06-08 15:39:23.481291+00');
INSERT INTO public.followers VALUES ('bc473a32-3549-4508-99c2-269c94341c96', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '56e89866-3a90-4f97-8461-34b24f842da4', 'accepted', '2025-06-08 17:54:05.481291+00');
INSERT INTO public.followers VALUES ('90a3f0ab-c590-492b-9ae9-dc544b4d3f73', '56e89866-3a90-4f97-8461-34b24f842da4', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', 'accepted', '2025-06-08 18:00:12.481291+00');
INSERT INTO public.followers VALUES ('15f07b40-1c1c-419c-b89a-5e7937273812', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', 'f399a996-1567-4fc0-a86a-92774494ca44', 'accepted', '2025-06-08 16:05:45.481291+00');
INSERT INTO public.followers VALUES ('d7898ac2-1854-4c96-9d0b-1c110bf536ca', 'f399a996-1567-4fc0-a86a-92774494ca44', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', 'accepted', '2025-06-08 16:13:03.481291+00');
INSERT INTO public.followers VALUES ('907f6d1e-6828-4e07-aa79-c8fd6dfd35e3', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', '79758f94-70d8-4a11-b05f-793fcb57e093', 'accepted', '2025-06-08 17:21:11.481291+00');
INSERT INTO public.followers VALUES ('619e4146-09e6-4c5d-b4ed-1c222cd4ec70', '79758f94-70d8-4a11-b05f-793fcb57e093', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', 'accepted', '2025-06-08 17:26:59.481291+00');
INSERT INTO public.followers VALUES ('6559c051-dbc4-4c19-9823-3aa1da2284b2', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', '56e89866-3a90-4f97-8461-34b24f842da4', 'accepted', '2025-06-08 17:54:02.481291+00');
INSERT INTO public.followers VALUES ('a743efc9-3130-442b-a48f-a6d3fece60e7', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', 'pending', '2025-06-08 18:38:28.481291+00');
INSERT INTO public.followers VALUES ('ad78f149-9558-438f-82c6-4fa55072e469', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', 'f399a996-1567-4fc0-a86a-92774494ca44', 'pending', '2025-06-08 18:23:12.481291+00');
INSERT INTO public.followers VALUES ('b2a80acb-64d2-4932-a1bd-25598e0cb729', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '6e137217-71a9-439b-a579-4bcd1a729af9', 'accepted', '2025-06-08 15:19:08.481291+00');
INSERT INTO public.followers VALUES ('c5bf83d8-b023-4ce6-bab2-050a2f95dced', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'accepted', '2025-06-08 17:36:00.481291+00');
INSERT INTO public.followers VALUES ('a58c1147-2a25-47c2-a5c9-646e3e4ffd6d', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', 'accepted', '2025-06-08 17:41:12.481291+00');
INSERT INTO public.followers VALUES ('7fc6d4d7-698b-41ce-9097-a50497fc0790', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'accepted', '2025-06-08 15:58:51.481291+00');
INSERT INTO public.followers VALUES ('0d03bf53-9a65-4e23-addb-52918ad56fd5', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '79758f94-70d8-4a11-b05f-793fcb57e093', 'accepted', '2025-06-08 17:18:43.481291+00');
INSERT INTO public.followers VALUES ('9352094c-f9fd-43a6-a23e-047745a3f26c', '79758f94-70d8-4a11-b05f-793fcb57e093', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', 'accepted', '2025-06-08 17:33:45.481291+00');
INSERT INTO public.followers VALUES ('47a807ea-6e1c-419b-94c4-823d012e3d38', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', 'pending', '2025-06-08 20:42:55.481291+00');
INSERT INTO public.followers VALUES ('8fc317ae-bc0f-4fbd-afa9-555f66fd698d', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', 'accepted', '2025-06-08 16:56:22.481291+00');
INSERT INTO public.followers VALUES ('51dae433-b8a8-40ea-8a2b-80a562b2cf32', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', 'accepted', '2025-06-08 17:00:20.481291+00');
INSERT INTO public.followers VALUES ('52c2cb73-a728-48f1-b390-b2cad34a3672', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'accepted', '2025-06-08 16:25:08.481291+00');
INSERT INTO public.followers VALUES ('4a3d1ac2-79ca-4353-94c9-ec7f3c039db7', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', 'f399a996-1567-4fc0-a86a-92774494ca44', 'accepted', '2025-06-08 16:10:17.481291+00');
INSERT INTO public.followers VALUES ('4e9280a3-8208-42da-8ae9-b9666ee90727', 'f399a996-1567-4fc0-a86a-92774494ca44', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', 'accepted', '2025-06-08 16:17:22.481291+00');
INSERT INTO public.followers VALUES ('442e589e-5e53-4a6f-9512-07c7bdf0e16e', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', 'pending', '2025-06-08 18:09:17.481291+00');
INSERT INTO public.followers VALUES ('0fa2dd4f-a250-4e4a-8cde-8b456cd8bdbf', 'f399a996-1567-4fc0-a86a-92774494ca44', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', 'accepted', '2025-06-08 15:23:27.481291+00');
INSERT INTO public.followers VALUES ('cbea1a56-d471-4c37-aab8-699a76d923b4', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', 'f399a996-1567-4fc0-a86a-92774494ca44', 'accepted', '2025-06-08 15:32:48.481291+00');
INSERT INTO public.followers VALUES ('44ed9e56-1fc0-4d85-be2f-ddd963617545', 'f399a996-1567-4fc0-a86a-92774494ca44', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'accepted', '2025-06-08 17:37:00.481291+00');
INSERT INTO public.followers VALUES ('6fbb8099-5d25-4c59-bb11-5393532a52f1', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'f399a996-1567-4fc0-a86a-92774494ca44', 'accepted', '2025-06-08 17:53:29.481291+00');
INSERT INTO public.followers VALUES ('c39f30d8-0ae8-46f0-8772-899ed0933bd9', 'f399a996-1567-4fc0-a86a-92774494ca44', '56e89866-3a90-4f97-8461-34b24f842da4', 'accepted', '2025-06-08 16:44:37.481291+00');
INSERT INTO public.followers VALUES ('c212ab2c-8660-4f89-b01a-199d3dc468ca', 'f399a996-1567-4fc0-a86a-92774494ca44', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'pending', '2025-06-08 20:10:03.481291+00');


--
-- Data for Name: likes; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.likes VALUES ('425f19e4-f605-440d-83a8-828d77e6975a', '6e137217-71a9-439b-a579-4bcd1a729af9', '02dd7349-fe64-4ff1-974b-b1673d7ac009', '2025-06-03 21:08:36.101041+00');
INSERT INTO public.likes VALUES ('3620c233-41e2-4921-a8e9-3b61eaa1fb75', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '02dd7349-fe64-4ff1-974b-b1673d7ac009', '2025-06-02 16:40:36.101127+00');
INSERT INTO public.likes VALUES ('1c44c3bd-7ab6-4521-998e-2ffc0b69bca4', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '02dd7349-fe64-4ff1-974b-b1673d7ac009', '2025-06-04 18:55:36.101252+00');
INSERT INTO public.likes VALUES ('744d33c8-1020-45d0-9b8a-5047fef02101', '3b0ebec9-3805-490b-987f-4c6e7286319e', '02dd7349-fe64-4ff1-974b-b1673d7ac009', '2025-06-05 00:20:36.101407+00');
INSERT INTO public.likes VALUES ('38e1c3c1-04c6-4f59-a63b-9a38d6a0f5dd', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '02dd7349-fe64-4ff1-974b-b1673d7ac009', '2025-06-06 09:35:36.101759+00');
INSERT INTO public.likes VALUES ('89b7cbfe-311f-4bb4-9d94-34ead1a6a822', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '02dd7349-fe64-4ff1-974b-b1673d7ac009', '2025-06-03 13:12:36.102386+00');
INSERT INTO public.likes VALUES ('51d6e6ce-dc25-43f5-9ae1-19f5d8207af2', '56e89866-3a90-4f97-8461-34b24f842da4', '02dd7349-fe64-4ff1-974b-b1673d7ac009', '2025-06-06 18:20:36.102922+00');
INSERT INTO public.likes VALUES ('e3bca902-4cec-42ad-b6e4-75fb526ba147', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '02dd7349-fe64-4ff1-974b-b1673d7ac009', '2025-06-07 00:55:36.103085+00');
INSERT INTO public.likes VALUES ('3a21ed0c-5712-4086-b008-5ec5c31c5e5d', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '03ebcf42-bfcf-4ee0-abae-752a1ded3096', '2025-06-05 00:19:36.103225+00');
INSERT INTO public.likes VALUES ('f58e9f60-bfff-4b5d-85ad-1dbba1b1dc07', '3b0ebec9-3805-490b-987f-4c6e7286319e', '03ebcf42-bfcf-4ee0-abae-752a1ded3096', '2025-06-06 14:37:36.103282+00');
INSERT INTO public.likes VALUES ('9eef3b92-ed32-451a-9fb6-00407cbaf07c', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '03ebcf42-bfcf-4ee0-abae-752a1ded3096', '2025-06-06 11:22:36.103347+00');
INSERT INTO public.likes VALUES ('a78a4844-6374-4776-ae83-4947b6dde4ee', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '03ebcf42-bfcf-4ee0-abae-752a1ded3096', '2025-06-07 01:07:36.103367+00');
INSERT INTO public.likes VALUES ('00627cd1-5ad2-402c-b538-75a3f7dd9c34', 'f399a996-1567-4fc0-a86a-92774494ca44', '03ebcf42-bfcf-4ee0-abae-752a1ded3096', '2025-06-05 18:58:36.103408+00');
INSERT INTO public.likes VALUES ('169ec0ae-0f72-4bbe-bb46-7a0789eff49d', '3b0ebec9-3805-490b-987f-4c6e7286319e', '0526699f-47c4-46b2-8846-a298a864ed24', '2025-06-06 18:20:36.103746+00');
INSERT INTO public.likes VALUES ('3af96425-4ddd-4dc0-bd84-db372d084b56', '56e89866-3a90-4f97-8461-34b24f842da4', '0526699f-47c4-46b2-8846-a298a864ed24', '2025-06-07 16:13:36.103791+00');
INSERT INTO public.likes VALUES ('7218b737-6a2b-482e-9bcb-9681c17f5ae4', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', '0526699f-47c4-46b2-8846-a298a864ed24', '2025-06-07 08:48:36.103805+00');
INSERT INTO public.likes VALUES ('ac0253aa-c826-46c4-abab-0c5dc7ecb496', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '0526699f-47c4-46b2-8846-a298a864ed24', '2025-06-07 00:18:36.103815+00');
INSERT INTO public.likes VALUES ('bcda6a40-8335-4b1d-9ac7-ae676fecf591', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '0526699f-47c4-46b2-8846-a298a864ed24', '2025-06-08 02:52:36.103825+00');
INSERT INTO public.likes VALUES ('dac988d0-6feb-46a7-a72a-a38c5811439c', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '0526699f-47c4-46b2-8846-a298a864ed24', '2025-06-03 07:46:36.103836+00');
INSERT INTO public.likes VALUES ('5590568a-08d1-4e66-a1fd-1b0fb3007760', '79758f94-70d8-4a11-b05f-793fcb57e093', '0526699f-47c4-46b2-8846-a298a864ed24', '2025-06-05 09:36:36.103845+00');
INSERT INTO public.likes VALUES ('9592b58a-8c14-432e-a5d0-1993277d174a', '6e137217-71a9-439b-a579-4bcd1a729af9', '0526699f-47c4-46b2-8846-a298a864ed24', '2025-06-03 07:52:36.103859+00');
INSERT INTO public.likes VALUES ('40b1346e-6f22-4e39-8560-1f21d49551e9', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '0526699f-47c4-46b2-8846-a298a864ed24', '2025-06-08 00:02:36.103894+00');
INSERT INTO public.likes VALUES ('771a5e97-174e-4a56-810b-0e47bc825f23', 'f399a996-1567-4fc0-a86a-92774494ca44', '0526699f-47c4-46b2-8846-a298a864ed24', '2025-06-07 09:36:36.10393+00');
INSERT INTO public.likes VALUES ('2475f6b1-8a27-4765-b85d-f0b7861c385c', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '0526699f-47c4-46b2-8846-a298a864ed24', '2025-06-01 14:11:36.103951+00');
INSERT INTO public.likes VALUES ('5933c16d-519a-4db7-9ba6-83a9d076268f', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '0abf63f3-6d91-4737-add9-edb192813b76', '2025-06-02 20:20:36.104014+00');
INSERT INTO public.likes VALUES ('41b86ca4-6c1f-49eb-ba04-a9cdf9ba6a9e', '6e137217-71a9-439b-a579-4bcd1a729af9', '0abf63f3-6d91-4737-add9-edb192813b76', '2025-06-02 07:45:36.104028+00');
INSERT INTO public.likes VALUES ('fd3e4608-1809-47a8-af39-2672978c03fe', '3b0ebec9-3805-490b-987f-4c6e7286319e', '0abf63f3-6d91-4737-add9-edb192813b76', '2025-06-02 13:55:36.104037+00');
INSERT INTO public.likes VALUES ('5bc02c82-8462-4171-875e-5a1f2ce9a4c1', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '0abf63f3-6d91-4737-add9-edb192813b76', '2025-06-05 02:06:36.104047+00');
INSERT INTO public.likes VALUES ('0cb789d0-e881-40b4-bf40-b376024081db', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '0abf63f3-6d91-4737-add9-edb192813b76', '2025-06-03 06:14:36.104057+00');
INSERT INTO public.likes VALUES ('c732a00f-0f9f-4c14-a9a2-2d45ce87d4e5', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '0abf63f3-6d91-4737-add9-edb192813b76', '2025-06-03 12:52:36.104066+00');
INSERT INTO public.likes VALUES ('678b7f9c-a298-4075-a37f-f39543774607', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', '0abf63f3-6d91-4737-add9-edb192813b76', '2025-06-08 00:06:36.104076+00');
INSERT INTO public.likes VALUES ('9ca85c39-f7f6-450f-8476-2f4eef46b7b9', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '0abf63f3-6d91-4737-add9-edb192813b76', '2025-06-03 18:16:36.104086+00');
INSERT INTO public.likes VALUES ('042197b4-8ea3-4010-bde5-a0970f39f4dc', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', '1402f5b0-c845-4665-998c-e5597f5f5521', '2025-06-07 02:28:36.104111+00');
INSERT INTO public.likes VALUES ('24e7c98b-76c1-4dbe-b2b8-0dba78950711', '3b0ebec9-3805-490b-987f-4c6e7286319e', '1402f5b0-c845-4665-998c-e5597f5f5521', '2025-06-03 20:24:36.104122+00');
INSERT INTO public.likes VALUES ('621ba366-16b2-41c7-a4ec-0ccf275f2a28', '79758f94-70d8-4a11-b05f-793fcb57e093', '1402f5b0-c845-4665-998c-e5597f5f5521', '2025-06-06 19:49:36.104131+00');
INSERT INTO public.likes VALUES ('746c28e4-fdee-4b3b-bf44-42be349ad479', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '1402f5b0-c845-4665-998c-e5597f5f5521', '2025-06-05 04:08:36.104141+00');
INSERT INTO public.likes VALUES ('abb5e142-e6b4-4546-8848-ed03e8af622f', '56e89866-3a90-4f97-8461-34b24f842da4', '1402f5b0-c845-4665-998c-e5597f5f5521', '2025-06-03 16:21:36.10415+00');
INSERT INTO public.likes VALUES ('982e3b5d-c860-4e05-9de2-a322659cf28f', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '1402f5b0-c845-4665-998c-e5597f5f5521', '2025-06-07 05:57:36.104159+00');
INSERT INTO public.likes VALUES ('32c24d60-3449-4dfa-b875-8742dd27a35a', '56e89866-3a90-4f97-8461-34b24f842da4', '19e2449f-0228-4c1a-b7c7-c7c8c0d7a4a4', '2025-06-01 21:12:36.10418+00');
INSERT INTO public.likes VALUES ('2fd5918c-45e0-4ca1-9e57-387fac5e364d', '79758f94-70d8-4a11-b05f-793fcb57e093', '19e2449f-0228-4c1a-b7c7-c7c8c0d7a4a4', '2025-06-07 05:13:36.104196+00');
INSERT INTO public.likes VALUES ('43c34d93-3ad8-4af4-9c1a-88da55499d93', '6e137217-71a9-439b-a579-4bcd1a729af9', '19e2449f-0228-4c1a-b7c7-c7c8c0d7a4a4', '2025-06-07 15:05:36.104212+00');
INSERT INTO public.likes VALUES ('d5e69037-a292-4bcb-bd0d-2406c8901331', 'f399a996-1567-4fc0-a86a-92774494ca44', '19e2449f-0228-4c1a-b7c7-c7c8c0d7a4a4', '2025-06-02 05:54:36.10423+00');
INSERT INTO public.likes VALUES ('2862a6da-5cc2-4868-a6c4-d8db93fa379b', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '19e2449f-0228-4c1a-b7c7-c7c8c0d7a4a4', '2025-06-02 12:45:36.104248+00');
INSERT INTO public.likes VALUES ('11157ce9-ca06-4570-9ffb-fe1101460aa2', 'f399a996-1567-4fc0-a86a-92774494ca44', '1e976362-73b0-400a-a74c-3fa951afdbeb', '2025-06-02 21:25:36.104302+00');
INSERT INTO public.likes VALUES ('73e58b62-d217-4cb0-9370-34af1b972bc9', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '1e976362-73b0-400a-a74c-3fa951afdbeb', '2025-06-03 07:34:36.104329+00');
INSERT INTO public.likes VALUES ('9d886846-6232-449b-b004-cdb413e86c5f', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '1e976362-73b0-400a-a74c-3fa951afdbeb', '2025-06-01 15:56:36.10434+00');
INSERT INTO public.likes VALUES ('0af8cc03-124a-4add-8d5e-a607171d2219', '6e137217-71a9-439b-a579-4bcd1a729af9', '1e976362-73b0-400a-a74c-3fa951afdbeb', '2025-06-01 16:21:36.104362+00');
INSERT INTO public.likes VALUES ('d7b130ef-c95d-4d1b-b00d-ec640d40aebe', '79758f94-70d8-4a11-b05f-793fcb57e093', '1e976362-73b0-400a-a74c-3fa951afdbeb', '2025-06-04 12:13:36.104372+00');
INSERT INTO public.likes VALUES ('33fe6295-a3b6-49bc-b2c9-c1531ee27aee', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '1e976362-73b0-400a-a74c-3fa951afdbeb', '2025-06-02 17:48:36.104385+00');
INSERT INTO public.likes VALUES ('e33be987-382b-47b4-ab7a-15c11394474d', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '23a24eb3-9a76-4807-8ff2-24ebdd6f11bb', '2025-06-07 14:10:36.104409+00');
INSERT INTO public.likes VALUES ('2bd91bee-cf84-4c6f-a406-e7702f354018', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', '23a24eb3-9a76-4807-8ff2-24ebdd6f11bb', '2025-06-05 03:05:36.10442+00');
INSERT INTO public.likes VALUES ('ecd71b2c-2cb3-433c-a299-1dda80415c45', '56e89866-3a90-4f97-8461-34b24f842da4', '23a24eb3-9a76-4807-8ff2-24ebdd6f11bb', '2025-06-02 01:58:36.104439+00');
INSERT INTO public.likes VALUES ('e7479bd8-55b7-45ba-8a9f-8dd693751cdd', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '23a24eb3-9a76-4807-8ff2-24ebdd6f11bb', '2025-06-05 03:26:36.104448+00');
INSERT INTO public.likes VALUES ('84b1d5b4-e223-4bb3-a6f4-c6c02f97226b', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '23a24eb3-9a76-4807-8ff2-24ebdd6f11bb', '2025-06-02 12:40:36.104458+00');
INSERT INTO public.likes VALUES ('916a5b28-2629-4012-8a81-1d6f694f1e6e', '6e137217-71a9-439b-a579-4bcd1a729af9', '23a24eb3-9a76-4807-8ff2-24ebdd6f11bb', '2025-06-06 14:59:36.104469+00');
INSERT INTO public.likes VALUES ('3f049a3c-d0f9-49d3-b127-aa4c59349ab6', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '23a24eb3-9a76-4807-8ff2-24ebdd6f11bb', '2025-06-03 06:48:36.104478+00');
INSERT INTO public.likes VALUES ('70d675f2-3687-41ac-a4cb-6cef83509fb8', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '23a24eb3-9a76-4807-8ff2-24ebdd6f11bb', '2025-06-02 20:16:36.1045+00');
INSERT INTO public.likes VALUES ('50e763c9-5704-48f2-b6c8-339e05493e2f', '79758f94-70d8-4a11-b05f-793fcb57e093', '23a24eb3-9a76-4807-8ff2-24ebdd6f11bb', '2025-06-04 09:32:36.104509+00');
INSERT INTO public.likes VALUES ('5a72fe2a-25ef-4d42-b35c-561f551c7e1a', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '23a24eb3-9a76-4807-8ff2-24ebdd6f11bb', '2025-06-07 04:59:36.104518+00');
INSERT INTO public.likes VALUES ('e481e7a6-a41c-40a4-bd07-cc77d9ec5f9d', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '23ad69f9-8c1f-45e1-bdb6-86ae97f0a2d2', '2025-06-03 01:00:36.10454+00');
INSERT INTO public.likes VALUES ('ed3b785e-753e-4f8b-90bf-e16611213f28', '79758f94-70d8-4a11-b05f-793fcb57e093', '23ad69f9-8c1f-45e1-bdb6-86ae97f0a2d2', '2025-06-07 03:12:36.1046+00');
INSERT INTO public.likes VALUES ('1a2b5861-55b0-401f-b3b4-c1064939ec38', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '23ad69f9-8c1f-45e1-bdb6-86ae97f0a2d2', '2025-06-02 22:24:36.104628+00');
INSERT INTO public.likes VALUES ('19fad761-ef72-46b1-ad94-41a3dfdf8d7f', 'f399a996-1567-4fc0-a86a-92774494ca44', '23ad69f9-8c1f-45e1-bdb6-86ae97f0a2d2', '2025-06-05 11:42:36.104639+00');
INSERT INTO public.likes VALUES ('e8a11fd0-a253-42e8-8191-2464f9a39d82', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '23ad69f9-8c1f-45e1-bdb6-86ae97f0a2d2', '2025-06-02 14:17:36.104648+00');
INSERT INTO public.likes VALUES ('6f47f212-38ce-4cd8-a0ef-a197841ce1d5', '3b0ebec9-3805-490b-987f-4c6e7286319e', '23ad69f9-8c1f-45e1-bdb6-86ae97f0a2d2', '2025-06-05 00:38:36.104657+00');
INSERT INTO public.likes VALUES ('574ab563-8ca8-4eea-809d-ded171500132', '6e137217-71a9-439b-a579-4bcd1a729af9', '23ad69f9-8c1f-45e1-bdb6-86ae97f0a2d2', '2025-06-05 04:19:36.104671+00');
INSERT INTO public.likes VALUES ('955e80be-4211-4853-9f46-d27cf6cd1867', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', '23ad69f9-8c1f-45e1-bdb6-86ae97f0a2d2', '2025-06-04 21:22:36.10468+00');
INSERT INTO public.likes VALUES ('258b5b6d-4dbb-46e3-8630-1b845d9dba1c', '56e89866-3a90-4f97-8461-34b24f842da4', '23ad69f9-8c1f-45e1-bdb6-86ae97f0a2d2', '2025-06-07 05:49:36.104692+00');
INSERT INTO public.likes VALUES ('06637418-b9d9-4618-bd89-e05fc47b6d7d', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '23ad69f9-8c1f-45e1-bdb6-86ae97f0a2d2', '2025-06-06 17:26:36.104707+00');
INSERT INTO public.likes VALUES ('5a21c37b-9de6-4296-aa05-9267885c2c14', 'f399a996-1567-4fc0-a86a-92774494ca44', '26e50906-4a53-4537-89c5-9971e86bb0ad', '2025-06-05 05:37:36.104768+00');
INSERT INTO public.likes VALUES ('c33db237-3524-40e0-b067-ec4de268e646', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', '26e50906-4a53-4537-89c5-9971e86bb0ad', '2025-06-02 14:12:36.104781+00');
INSERT INTO public.likes VALUES ('15c7f36e-d236-497f-bab1-2c734d3563a1', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '26e50906-4a53-4537-89c5-9971e86bb0ad', '2025-06-06 04:21:36.104791+00');
INSERT INTO public.likes VALUES ('406154e9-f449-4fbf-9387-f3a4bf884717', '3b0ebec9-3805-490b-987f-4c6e7286319e', '26e50906-4a53-4537-89c5-9971e86bb0ad', '2025-06-05 14:44:36.1048+00');
INSERT INTO public.likes VALUES ('5127f3f7-fcff-47ed-95fe-9240b0093303', '56e89866-3a90-4f97-8461-34b24f842da4', '26e50906-4a53-4537-89c5-9971e86bb0ad', '2025-06-03 12:23:36.104809+00');
INSERT INTO public.likes VALUES ('10cd0b2d-b2c9-4a44-bd4e-6842115e844e', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '26e50906-4a53-4537-89c5-9971e86bb0ad', '2025-06-08 01:04:36.104817+00');
INSERT INTO public.likes VALUES ('4b826b88-3aab-4e6d-acdc-9bb2f7f0ada7', '79758f94-70d8-4a11-b05f-793fcb57e093', '26e50906-4a53-4537-89c5-9971e86bb0ad', '2025-06-02 22:20:36.104825+00');
INSERT INTO public.likes VALUES ('d4913cc6-8707-410e-87f6-19a45dfd54b1', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '26e50906-4a53-4537-89c5-9971e86bb0ad', '2025-06-01 18:52:36.104834+00');
INSERT INTO public.likes VALUES ('b7482459-bb0a-4124-83f4-65819828b8aa', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '26e50906-4a53-4537-89c5-9971e86bb0ad', '2025-06-02 00:07:36.104844+00');
INSERT INTO public.likes VALUES ('2759c48b-cd4d-485d-9795-cea7ab9efedb', '6e137217-71a9-439b-a579-4bcd1a729af9', '26e50906-4a53-4537-89c5-9971e86bb0ad', '2025-06-05 10:25:36.104854+00');
INSERT INTO public.likes VALUES ('915bc11e-3e83-4508-84c9-9d942e06ff06', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '26e50906-4a53-4537-89c5-9971e86bb0ad', '2025-06-04 12:27:36.104864+00');
INSERT INTO public.likes VALUES ('d8e4d58e-52a7-4467-9911-b33ed3b6882e', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', '2d2a93c2-c55a-4fcc-8fb2-72b38f9e96a4', '2025-06-05 23:46:36.104883+00');
INSERT INTO public.likes VALUES ('ac8ba6a3-6da0-4449-a46b-7bf9fb8b0e7d', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '2d2a93c2-c55a-4fcc-8fb2-72b38f9e96a4', '2025-06-02 16:12:36.104892+00');
INSERT INTO public.likes VALUES ('9c04ca28-33b9-499e-b61e-9b185398d77b', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '2d2a93c2-c55a-4fcc-8fb2-72b38f9e96a4', '2025-06-03 15:22:36.104901+00');
INSERT INTO public.likes VALUES ('096b9ff6-0dc7-4d25-988c-ea0c1b696949', '56e89866-3a90-4f97-8461-34b24f842da4', '2d2a93c2-c55a-4fcc-8fb2-72b38f9e96a4', '2025-06-07 01:33:36.10491+00');
INSERT INTO public.likes VALUES ('d3823ce3-aeae-44e0-b6d3-efe15e3e3ae1', 'f399a996-1567-4fc0-a86a-92774494ca44', '2d2a93c2-c55a-4fcc-8fb2-72b38f9e96a4', '2025-06-07 07:51:36.104919+00');
INSERT INTO public.likes VALUES ('91b4adfb-b64b-44f7-a942-740a8a7b0e54', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '31f82a0d-869c-4ea9-8c72-152d229893cf', '2025-06-02 16:51:36.104948+00');
INSERT INTO public.likes VALUES ('69ecd14f-718b-4037-bcd3-d1faffa476dd', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '31f82a0d-869c-4ea9-8c72-152d229893cf', '2025-06-07 03:09:36.104968+00');
INSERT INTO public.likes VALUES ('c6fbc127-6ba1-4a51-acbd-91480e7512b2', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '31f82a0d-869c-4ea9-8c72-152d229893cf', '2025-06-06 04:04:36.104984+00');
INSERT INTO public.likes VALUES ('1636b08e-f6ce-4861-81c5-8c856b7188a5', '79758f94-70d8-4a11-b05f-793fcb57e093', '31f82a0d-869c-4ea9-8c72-152d229893cf', '2025-06-02 07:07:36.104997+00');
INSERT INTO public.likes VALUES ('8bc4a1be-8a7e-40b9-a549-078a3e508670', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '31f82a0d-869c-4ea9-8c72-152d229893cf', '2025-06-08 11:01:36.105014+00');
INSERT INTO public.likes VALUES ('591be700-538b-4a67-9126-102aab2232d8', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '31f82a0d-869c-4ea9-8c72-152d229893cf', '2025-06-03 07:49:36.105031+00');
INSERT INTO public.likes VALUES ('6448b8de-86d1-452a-b5ed-cf74d1d88133', 'f399a996-1567-4fc0-a86a-92774494ca44', '31f82a0d-869c-4ea9-8c72-152d229893cf', '2025-06-06 03:10:36.105043+00');
INSERT INTO public.likes VALUES ('836f03d9-54b0-4639-963f-f15bac877cd8', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '31f82a0d-869c-4ea9-8c72-152d229893cf', '2025-06-07 19:12:36.105051+00');
INSERT INTO public.likes VALUES ('3e2a61fc-1754-469a-b569-a9bc216407a1', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', '31f82a0d-869c-4ea9-8c72-152d229893cf', '2025-06-08 02:02:36.105062+00');
INSERT INTO public.likes VALUES ('3dee6f12-9dd7-4730-a87b-28a14f16a199', '56e89866-3a90-4f97-8461-34b24f842da4', '31f82a0d-869c-4ea9-8c72-152d229893cf', '2025-06-04 03:13:36.105071+00');
INSERT INTO public.likes VALUES ('fc1acb7b-5297-4ec7-ab07-f179515b82e5', '3b0ebec9-3805-490b-987f-4c6e7286319e', '33f2ef0e-d58d-4b7a-a5fe-89b87390accd', '2025-06-04 21:27:36.105091+00');
INSERT INTO public.likes VALUES ('c90e0cf2-b7cb-490c-a27e-48dae781df39', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '33f2ef0e-d58d-4b7a-a5fe-89b87390accd', '2025-06-03 18:38:36.1051+00');
INSERT INTO public.likes VALUES ('cbe916e7-e75a-46a0-ae9e-947b8180d418', '6e137217-71a9-439b-a579-4bcd1a729af9', '33f2ef0e-d58d-4b7a-a5fe-89b87390accd', '2025-06-06 14:14:36.105111+00');
INSERT INTO public.likes VALUES ('d072a6eb-3eea-49c4-a8c1-85599d5e6d2c', '79758f94-70d8-4a11-b05f-793fcb57e093', '33f2ef0e-d58d-4b7a-a5fe-89b87390accd', '2025-06-03 01:29:36.10512+00');
INSERT INTO public.likes VALUES ('cf568724-f322-4588-bf73-2a183703ae43', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '33f2ef0e-d58d-4b7a-a5fe-89b87390accd', '2025-06-04 03:29:36.105129+00');
INSERT INTO public.likes VALUES ('6bae2aa8-c418-49b3-a1b5-e0d3991c31cb', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '3413acc6-1419-4111-9328-0485f0ed1ef8', '2025-06-02 22:31:36.105149+00');
INSERT INTO public.likes VALUES ('646e1b53-8806-4fed-a94f-923a6b1b8ffc', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '3413acc6-1419-4111-9328-0485f0ed1ef8', '2025-06-02 20:54:36.10516+00');
INSERT INTO public.likes VALUES ('d0591b17-7bfe-4965-9e89-8a5181163a06', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', '3413acc6-1419-4111-9328-0485f0ed1ef8', '2025-06-07 23:09:36.105169+00');
INSERT INTO public.likes VALUES ('f747ae5a-6d6a-457a-b075-6b5096beed97', 'f399a996-1567-4fc0-a86a-92774494ca44', '3413acc6-1419-4111-9328-0485f0ed1ef8', '2025-06-08 09:04:36.10518+00');
INSERT INTO public.likes VALUES ('209e8365-6008-4ba7-87a3-38ef6e9b5040', '3b0ebec9-3805-490b-987f-4c6e7286319e', '3413acc6-1419-4111-9328-0485f0ed1ef8', '2025-06-04 01:19:36.105191+00');
INSERT INTO public.likes VALUES ('57cc22c5-8b6a-4a7e-80e9-2b2b3e3f41f8', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '3413acc6-1419-4111-9328-0485f0ed1ef8', '2025-06-02 13:57:36.105201+00');
INSERT INTO public.likes VALUES ('4b7c08a4-15a4-40e8-8f08-66ef5445ebaa', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '3487baed-69d9-4a07-ba09-ea0e9c2699a4', '2025-06-05 17:31:36.105219+00');
INSERT INTO public.likes VALUES ('cd6aeb35-d3f2-47a4-afd0-dad73418f955', '3b0ebec9-3805-490b-987f-4c6e7286319e', '3487baed-69d9-4a07-ba09-ea0e9c2699a4', '2025-06-08 00:24:36.105228+00');
INSERT INTO public.likes VALUES ('7e598d62-6e0a-4df1-a521-b757dc963324', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '3487baed-69d9-4a07-ba09-ea0e9c2699a4', '2025-06-02 17:32:36.105237+00');
INSERT INTO public.likes VALUES ('fd8bedd1-50d7-442e-9e15-c70bc8b54dc6', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '3487baed-69d9-4a07-ba09-ea0e9c2699a4', '2025-06-04 17:53:36.105246+00');
INSERT INTO public.likes VALUES ('a05eba02-689b-41a0-a718-198458bb9b82', '56e89866-3a90-4f97-8461-34b24f842da4', '3487baed-69d9-4a07-ba09-ea0e9c2699a4', '2025-06-04 16:19:36.105255+00');
INSERT INTO public.likes VALUES ('af6860a6-ddcf-49a3-b037-2cd1ce130fb9', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '3487baed-69d9-4a07-ba09-ea0e9c2699a4', '2025-06-06 08:34:36.105263+00');
INSERT INTO public.likes VALUES ('8575b8c0-b6e5-44df-9365-8b842acd52b4', '6e137217-71a9-439b-a579-4bcd1a729af9', '3487baed-69d9-4a07-ba09-ea0e9c2699a4', '2025-06-02 23:47:36.105273+00');
INSERT INTO public.likes VALUES ('48476354-51cc-4eb3-9a1f-1aa31685fdbe', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '3487baed-69d9-4a07-ba09-ea0e9c2699a4', '2025-06-06 12:14:36.105288+00');
INSERT INTO public.likes VALUES ('b35634cc-fe79-491d-b6cd-25991dc0922d', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '40123287-2503-4a02-bd89-93f13eea21be', '2025-06-06 20:53:36.105364+00');
INSERT INTO public.likes VALUES ('cb0a94a3-6e91-42d2-83e9-0eacc2339e01', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', '40123287-2503-4a02-bd89-93f13eea21be', '2025-06-05 22:25:36.105397+00');
INSERT INTO public.likes VALUES ('15d344c3-855f-49c0-b8ce-0948dff91d0e', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '40123287-2503-4a02-bd89-93f13eea21be', '2025-06-02 23:17:36.105408+00');
INSERT INTO public.likes VALUES ('5684b0fa-451b-4c26-a17e-b8ac401c1864', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '40123287-2503-4a02-bd89-93f13eea21be', '2025-06-05 19:04:36.105422+00');
INSERT INTO public.likes VALUES ('82a788f9-3e74-40e6-a522-bac4ff13ac34', '6e137217-71a9-439b-a579-4bcd1a729af9', '40123287-2503-4a02-bd89-93f13eea21be', '2025-06-01 21:44:36.105431+00');
INSERT INTO public.likes VALUES ('546d81b8-9863-452f-8431-f392b3da4e47', '56e89866-3a90-4f97-8461-34b24f842da4', '40123287-2503-4a02-bd89-93f13eea21be', '2025-06-01 14:16:36.10544+00');
INSERT INTO public.likes VALUES ('5e570374-c009-4496-a7e4-0b6136002f0f', 'f399a996-1567-4fc0-a86a-92774494ca44', '40123287-2503-4a02-bd89-93f13eea21be', '2025-06-04 09:45:36.105448+00');
INSERT INTO public.likes VALUES ('314c79c7-3ca1-4c0f-bef3-4e5bca59f7a9', '3b0ebec9-3805-490b-987f-4c6e7286319e', '40123287-2503-4a02-bd89-93f13eea21be', '2025-06-02 10:52:36.105458+00');
INSERT INTO public.likes VALUES ('aaf336ac-4378-4e80-82b6-889aa067ecfe', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '40123287-2503-4a02-bd89-93f13eea21be', '2025-06-07 12:11:36.105467+00');
INSERT INTO public.likes VALUES ('70047e27-fa3b-4f8a-8e7a-5e0f1ffdf857', '3b0ebec9-3805-490b-987f-4c6e7286319e', '47d7751c-7bc4-4dd3-8c10-a25918732741', '2025-06-06 21:15:36.105492+00');
INSERT INTO public.likes VALUES ('9ed14cee-d069-495c-a8c0-e2365ecfab26', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '47d7751c-7bc4-4dd3-8c10-a25918732741', '2025-06-05 08:49:36.105502+00');
INSERT INTO public.likes VALUES ('bc2ac65a-0962-4f3c-b880-abba87d2bb5f', '6e137217-71a9-439b-a579-4bcd1a729af9', '47d7751c-7bc4-4dd3-8c10-a25918732741', '2025-06-06 06:47:36.105512+00');
INSERT INTO public.likes VALUES ('6738a6af-7e08-4cd0-a58a-fcff2a901747', 'f399a996-1567-4fc0-a86a-92774494ca44', '47d7751c-7bc4-4dd3-8c10-a25918732741', '2025-06-04 06:47:36.105522+00');
INSERT INTO public.likes VALUES ('910abacc-4c24-4208-ada2-9a0775a00a2a', '79758f94-70d8-4a11-b05f-793fcb57e093', '47d7751c-7bc4-4dd3-8c10-a25918732741', '2025-06-02 17:01:36.105532+00');
INSERT INTO public.likes VALUES ('b258e25b-ea98-49aa-a106-59d6973c458f', 'f399a996-1567-4fc0-a86a-92774494ca44', '4836de17-5868-4800-bdc7-57144e5c3807', '2025-06-03 16:57:36.105549+00');
INSERT INTO public.likes VALUES ('ad1815a0-6aa6-42db-b422-089fb235501c', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', '4836de17-5868-4800-bdc7-57144e5c3807', '2025-06-02 18:25:36.105561+00');
INSERT INTO public.likes VALUES ('7245f937-7a29-4349-ad0e-5b5aba5455dd', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '4836de17-5868-4800-bdc7-57144e5c3807', '2025-06-05 16:46:36.105571+00');
INSERT INTO public.likes VALUES ('898a74da-6657-4b45-9efb-b5d0dad013ff', '6e137217-71a9-439b-a579-4bcd1a729af9', '4836de17-5868-4800-bdc7-57144e5c3807', '2025-06-04 21:08:36.10558+00');
INSERT INTO public.likes VALUES ('32c23bb1-530a-4dba-a6be-f3d5357b8c7b', '79758f94-70d8-4a11-b05f-793fcb57e093', '4836de17-5868-4800-bdc7-57144e5c3807', '2025-06-04 21:35:36.105588+00');
INSERT INTO public.likes VALUES ('6e7a411d-0fc7-4a1c-b5b0-d5c2896e445e', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '4fcdd3cc-c626-4dd4-904b-426f0839c71c', '2025-06-02 20:07:36.105608+00');
INSERT INTO public.likes VALUES ('470e880e-2aae-4587-9ecb-7b93a3966500', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '4fcdd3cc-c626-4dd4-904b-426f0839c71c', '2025-06-03 10:12:36.105617+00');
INSERT INTO public.likes VALUES ('c90b5524-a603-4a65-8787-a1fefba4f518', '56e89866-3a90-4f97-8461-34b24f842da4', '4fcdd3cc-c626-4dd4-904b-426f0839c71c', '2025-06-08 08:21:36.105626+00');
INSERT INTO public.likes VALUES ('ece559a8-6530-4842-b760-c9b0adf737c3', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '4fcdd3cc-c626-4dd4-904b-426f0839c71c', '2025-06-04 02:58:36.105637+00');
INSERT INTO public.likes VALUES ('d9ea6fdf-0300-4f59-b369-3e72ec62d9a3', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '4fcdd3cc-c626-4dd4-904b-426f0839c71c', '2025-06-03 21:41:36.105653+00');
INSERT INTO public.likes VALUES ('b041759b-61c4-430e-8f77-4a330e8423a4', '3b0ebec9-3805-490b-987f-4c6e7286319e', '4fcdd3cc-c626-4dd4-904b-426f0839c71c', '2025-06-03 14:10:36.105671+00');
INSERT INTO public.likes VALUES ('45030c5d-63bf-475e-95d7-f18ce4508713', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '4fcdd3cc-c626-4dd4-904b-426f0839c71c', '2025-06-02 21:22:36.105692+00');
INSERT INTO public.likes VALUES ('ee733ccb-4c7c-49fc-bd02-21d5a5731fc7', '3b0ebec9-3805-490b-987f-4c6e7286319e', '53626074-a976-43cb-83b1-b812a80f3244', '2025-06-06 17:09:36.105742+00');
INSERT INTO public.likes VALUES ('c79ca2a1-291c-4590-8aba-1e799a5ca18f', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '53626074-a976-43cb-83b1-b812a80f3244', '2025-06-06 12:49:36.105752+00');
INSERT INTO public.likes VALUES ('afdbfbd4-0b0a-4941-891f-bae718c34bb8', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '53626074-a976-43cb-83b1-b812a80f3244', '2025-06-02 17:51:36.105761+00');
INSERT INTO public.likes VALUES ('6d6cdba7-932a-4ed3-a3b3-a2f8ed78a8c0', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', '53626074-a976-43cb-83b1-b812a80f3244', '2025-06-05 18:44:36.105769+00');
INSERT INTO public.likes VALUES ('b840f55d-6c18-4a34-84a7-f2bd836df868', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '53626074-a976-43cb-83b1-b812a80f3244', '2025-06-02 14:53:36.105778+00');
INSERT INTO public.likes VALUES ('5ddc8624-35e1-40e5-bd58-d5df4b45ad93', '79758f94-70d8-4a11-b05f-793fcb57e093', '53626074-a976-43cb-83b1-b812a80f3244', '2025-06-04 00:36:36.105788+00');
INSERT INTO public.likes VALUES ('55122b33-db0d-4702-9fcd-d6a3c9cb98e0', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '53626074-a976-43cb-83b1-b812a80f3244', '2025-06-02 16:22:36.1058+00');
INSERT INTO public.likes VALUES ('985e1201-d2b3-4bef-ac79-633b81593270', '6e137217-71a9-439b-a579-4bcd1a729af9', '53626074-a976-43cb-83b1-b812a80f3244', '2025-06-03 05:54:36.10583+00');
INSERT INTO public.likes VALUES ('fe768267-f3a2-476f-af08-73a15e46eb79', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '53626074-a976-43cb-83b1-b812a80f3244', '2025-06-01 21:13:36.105848+00');
INSERT INTO public.likes VALUES ('20623379-db22-4789-a822-ff6e497ed505', '56e89866-3a90-4f97-8461-34b24f842da4', '53626074-a976-43cb-83b1-b812a80f3244', '2025-06-01 16:44:36.105862+00');
INSERT INTO public.likes VALUES ('8e2b81d6-f775-4532-93c2-d510b8f0c9d3', 'f399a996-1567-4fc0-a86a-92774494ca44', '53626074-a976-43cb-83b1-b812a80f3244', '2025-06-03 17:53:36.105872+00');
INSERT INTO public.likes VALUES ('cda60214-f66d-4548-881d-c1740d8cb8d3', '56e89866-3a90-4f97-8461-34b24f842da4', '5931007c-9050-44ae-90af-0c9732651ff3', '2025-06-05 11:56:36.105899+00');
INSERT INTO public.likes VALUES ('5b38eb18-f00b-497c-bdb9-ea70038a3c09', '3b0ebec9-3805-490b-987f-4c6e7286319e', '5931007c-9050-44ae-90af-0c9732651ff3', '2025-06-08 06:14:36.105925+00');
INSERT INTO public.likes VALUES ('5a0ea8a8-210a-412f-87cd-af9c36c51db7', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '5931007c-9050-44ae-90af-0c9732651ff3', '2025-06-05 17:08:36.105934+00');
INSERT INTO public.likes VALUES ('5b5c0b74-136a-46dc-ace5-8fdb8745ab2f', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '5931007c-9050-44ae-90af-0c9732651ff3', '2025-06-03 17:14:36.105943+00');
INSERT INTO public.likes VALUES ('b4cd3747-a0a3-4507-a7cc-2ead92aec348', '6e137217-71a9-439b-a579-4bcd1a729af9', '5931007c-9050-44ae-90af-0c9732651ff3', '2025-06-05 05:50:36.105951+00');
INSERT INTO public.likes VALUES ('f70623ad-3c03-43db-8952-95b53914a079', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '5931007c-9050-44ae-90af-0c9732651ff3', '2025-06-06 23:02:36.10596+00');
INSERT INTO public.likes VALUES ('6844380c-3568-4326-9089-c5c497ff8f6a', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '5931007c-9050-44ae-90af-0c9732651ff3', '2025-06-02 08:19:36.105968+00');
INSERT INTO public.likes VALUES ('a5b04c4f-024e-4dea-b61b-0fdd70fdec67', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '5931007c-9050-44ae-90af-0c9732651ff3', '2025-06-06 16:25:36.105976+00');
INSERT INTO public.likes VALUES ('920075ab-186b-4d6e-9656-e5403c653d73', '56e89866-3a90-4f97-8461-34b24f842da4', '596dd13f-30c4-402d-9eb4-5e041370d12d', '2025-06-02 21:21:36.10601+00');
INSERT INTO public.likes VALUES ('b9d9325b-27ae-4f14-b31c-cbdb5aa6eb56', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '596dd13f-30c4-402d-9eb4-5e041370d12d', '2025-06-07 02:34:36.106038+00');
INSERT INTO public.likes VALUES ('06d86be9-6010-4505-bce8-7dab63ed22bd', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '596dd13f-30c4-402d-9eb4-5e041370d12d', '2025-06-02 00:59:36.106055+00');
INSERT INTO public.likes VALUES ('7bcaadfe-aa94-4db5-ae20-befe03e353cb', '3b0ebec9-3805-490b-987f-4c6e7286319e', '596dd13f-30c4-402d-9eb4-5e041370d12d', '2025-06-03 11:17:36.106065+00');
INSERT INTO public.likes VALUES ('71d62dcc-9e43-4985-830e-faed5a48de12', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '596dd13f-30c4-402d-9eb4-5e041370d12d', '2025-06-07 01:00:36.106075+00');
INSERT INTO public.likes VALUES ('3909dea4-c64d-4a5e-be17-3721ab9b9ec1', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '596dd13f-30c4-402d-9eb4-5e041370d12d', '2025-06-06 12:05:36.106083+00');
INSERT INTO public.likes VALUES ('a390702b-a5c5-4110-99e7-90ee3a759e87', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', '596dd13f-30c4-402d-9eb4-5e041370d12d', '2025-06-05 15:34:36.106092+00');
INSERT INTO public.likes VALUES ('91c2c6ec-cfd0-4084-bc43-6b055c8ea2e4', 'f399a996-1567-4fc0-a86a-92774494ca44', '596dd13f-30c4-402d-9eb4-5e041370d12d', '2025-06-08 07:14:36.106102+00');
INSERT INTO public.likes VALUES ('369743b8-2c80-49f8-9279-934cf5ecada1', '56e89866-3a90-4f97-8461-34b24f842da4', '60333cb2-3eac-4c88-b1ab-187788f3a753', '2025-06-02 23:47:36.106125+00');
INSERT INTO public.likes VALUES ('80f25e84-3682-4fad-b49f-cfd8384f35c0', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '60333cb2-3eac-4c88-b1ab-187788f3a753', '2025-06-07 12:40:36.106136+00');
INSERT INTO public.likes VALUES ('6ce3d446-ffde-431d-b410-dd58fa3f21fd', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '60333cb2-3eac-4c88-b1ab-187788f3a753', '2025-06-03 00:45:36.106145+00');
INSERT INTO public.likes VALUES ('9b9dcec4-76ce-4fce-974f-f83b63625b45', '79758f94-70d8-4a11-b05f-793fcb57e093', '60333cb2-3eac-4c88-b1ab-187788f3a753', '2025-06-04 00:17:36.106153+00');
INSERT INTO public.likes VALUES ('6272c230-23e8-48d4-9455-47bd0ca37712', 'f399a996-1567-4fc0-a86a-92774494ca44', '60333cb2-3eac-4c88-b1ab-187788f3a753', '2025-06-04 21:36:36.106162+00');
INSERT INTO public.likes VALUES ('8b129281-ffe3-4eea-ad4c-e0c2d046b94d', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', '60333cb2-3eac-4c88-b1ab-187788f3a753', '2025-06-06 21:16:36.10617+00');
INSERT INTO public.likes VALUES ('60409185-cb7f-4258-99da-8b3fccd7da12', '6e137217-71a9-439b-a579-4bcd1a729af9', '60333cb2-3eac-4c88-b1ab-187788f3a753', '2025-06-03 17:11:36.106182+00');
INSERT INTO public.likes VALUES ('fcacfb18-712f-4170-a081-b85d1485d26a', '3b0ebec9-3805-490b-987f-4c6e7286319e', '60333cb2-3eac-4c88-b1ab-187788f3a753', '2025-06-02 00:17:36.106191+00');
INSERT INTO public.likes VALUES ('f30dfec8-a8c5-4e50-9ac3-6a703449b6ba', '6e137217-71a9-439b-a579-4bcd1a729af9', '678fdbee-e90c-445c-949e-e3962b0503ce', '2025-06-05 22:43:36.10623+00');
INSERT INTO public.likes VALUES ('39894a62-1a73-436d-bcc8-1c8ce89b5e19', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '678fdbee-e90c-445c-949e-e3962b0503ce', '2025-06-06 14:04:36.106239+00');
INSERT INTO public.likes VALUES ('9f62f9b3-679e-4fc4-b241-873779df4ff6', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '678fdbee-e90c-445c-949e-e3962b0503ce', '2025-06-03 22:00:36.106248+00');
INSERT INTO public.likes VALUES ('eff4d226-adc7-45c6-a9e0-b9622ff18520', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', '678fdbee-e90c-445c-949e-e3962b0503ce', '2025-06-06 17:49:36.106256+00');
INSERT INTO public.likes VALUES ('c20986a6-3d0e-4d66-bd90-60ef250b8ebf', '56e89866-3a90-4f97-8461-34b24f842da4', '678fdbee-e90c-445c-949e-e3962b0503ce', '2025-06-02 09:30:36.106265+00');
INSERT INTO public.likes VALUES ('5d58ee55-85b2-4de0-ab7d-c7f60bfb2f44', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '678fdbee-e90c-445c-949e-e3962b0503ce', '2025-06-06 17:57:36.106274+00');
INSERT INTO public.likes VALUES ('fff7fb4e-0911-4f49-87cd-e68d3669a10e', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '678fdbee-e90c-445c-949e-e3962b0503ce', '2025-06-05 00:45:36.106282+00');
INSERT INTO public.likes VALUES ('bf63cff0-d851-4810-97e4-0ef4a1fff993', '79758f94-70d8-4a11-b05f-793fcb57e093', '678fdbee-e90c-445c-949e-e3962b0503ce', '2025-06-05 02:36:36.10629+00');
INSERT INTO public.likes VALUES ('87e25e5c-de6d-455c-baeb-f426e27189e2', 'f399a996-1567-4fc0-a86a-92774494ca44', '678fdbee-e90c-445c-949e-e3962b0503ce', '2025-06-07 05:37:36.106299+00');
INSERT INTO public.likes VALUES ('6b29f1ae-68d0-41e1-9ef0-617d7dd24c8e', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '678fdbee-e90c-445c-949e-e3962b0503ce', '2025-06-04 04:48:36.106307+00');
INSERT INTO public.likes VALUES ('95d09b20-f45e-40e0-b8a4-03978a8f73f4', 'f399a996-1567-4fc0-a86a-92774494ca44', '741d9f45-74da-4da2-b40e-06b325c5342f', '2025-06-04 09:38:36.106326+00');
INSERT INTO public.likes VALUES ('b8eaa369-5bd2-4d52-8666-985ffaabbeb1', '56e89866-3a90-4f97-8461-34b24f842da4', '741d9f45-74da-4da2-b40e-06b325c5342f', '2025-06-05 05:53:36.106336+00');
INSERT INTO public.likes VALUES ('0a67e82b-b822-4e05-b6f5-ce517718024c', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '741d9f45-74da-4da2-b40e-06b325c5342f', '2025-06-03 07:50:36.106346+00');
INSERT INTO public.likes VALUES ('3ead9135-a749-4bf5-83e6-57c609391355', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '741d9f45-74da-4da2-b40e-06b325c5342f', '2025-06-05 05:46:36.106355+00');
INSERT INTO public.likes VALUES ('627524bb-6c08-47a3-bbff-7a7429809c50', '79758f94-70d8-4a11-b05f-793fcb57e093', '741d9f45-74da-4da2-b40e-06b325c5342f', '2025-06-02 15:38:36.106368+00');
INSERT INTO public.likes VALUES ('36b2e65c-a4c4-442c-a039-d445f090733a', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '741d9f45-74da-4da2-b40e-06b325c5342f', '2025-06-05 19:34:36.106383+00');
INSERT INTO public.likes VALUES ('15085b58-e01d-4860-a130-2fc91b0fbba4', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', '741d9f45-74da-4da2-b40e-06b325c5342f', '2025-06-03 20:08:36.106405+00');
INSERT INTO public.likes VALUES ('298ac925-f655-4b07-82f6-44cb6cf634cd', '3b0ebec9-3805-490b-987f-4c6e7286319e', '789062db-6670-4972-97f9-21f61cdbd723', '2025-06-05 10:38:36.10647+00');
INSERT INTO public.likes VALUES ('d9545ebe-e1e3-4543-99e1-e00bd9dcd1e1', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '789062db-6670-4972-97f9-21f61cdbd723', '2025-06-05 16:52:36.10651+00');
INSERT INTO public.likes VALUES ('f9834215-94f9-4197-8292-34ed7ff03288', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '789062db-6670-4972-97f9-21f61cdbd723', '2025-06-03 14:20:36.106521+00');
INSERT INTO public.likes VALUES ('4e5ced96-8a95-4556-aa11-f6defe698364', 'f399a996-1567-4fc0-a86a-92774494ca44', '789062db-6670-4972-97f9-21f61cdbd723', '2025-06-08 03:58:36.106531+00');
INSERT INTO public.likes VALUES ('277b5440-dcd2-4c26-a9cd-e5b2cbf9bac3', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '789062db-6670-4972-97f9-21f61cdbd723', '2025-06-06 03:43:36.10654+00');
INSERT INTO public.likes VALUES ('3f3a7362-4924-44ba-b9ea-fd505eaa3eb9', '56e89866-3a90-4f97-8461-34b24f842da4', '789062db-6670-4972-97f9-21f61cdbd723', '2025-06-02 23:05:36.106553+00');
INSERT INTO public.likes VALUES ('28c7af60-f526-4aa2-a3f2-202b6b55c0aa', '79758f94-70d8-4a11-b05f-793fcb57e093', '789062db-6670-4972-97f9-21f61cdbd723', '2025-06-08 00:37:36.106563+00');
INSERT INTO public.likes VALUES ('d36106d7-99fd-471a-b2d9-b07d800d08c1', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '789062db-6670-4972-97f9-21f61cdbd723', '2025-06-03 16:55:36.106573+00');
INSERT INTO public.likes VALUES ('f89a0db0-9225-496c-aa5b-f91704d27854', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', '789062db-6670-4972-97f9-21f61cdbd723', '2025-06-05 12:14:36.106584+00');
INSERT INTO public.likes VALUES ('344183eb-3376-4950-a4fc-c2a9e58cd82e', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', '866d6d95-29d3-41e2-adb7-a01b09fd5616', '2025-06-07 06:10:36.106647+00');
INSERT INTO public.likes VALUES ('320650a5-c7d2-4f79-b6fa-ae722bad4fce', '56e89866-3a90-4f97-8461-34b24f842da4', '866d6d95-29d3-41e2-adb7-a01b09fd5616', '2025-06-07 12:51:36.106659+00');
INSERT INTO public.likes VALUES ('f7613a7d-103d-4972-8d5c-6da9dcb09917', '3b0ebec9-3805-490b-987f-4c6e7286319e', '866d6d95-29d3-41e2-adb7-a01b09fd5616', '2025-06-08 07:07:36.106669+00');
INSERT INTO public.likes VALUES ('e0cc2391-9d40-4841-bb9b-d6ad110b3699', '6e137217-71a9-439b-a579-4bcd1a729af9', '866d6d95-29d3-41e2-adb7-a01b09fd5616', '2025-06-06 01:04:36.106689+00');
INSERT INTO public.likes VALUES ('e3f6bb02-ebb5-4706-9a1a-b81ca0a77e49', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '866d6d95-29d3-41e2-adb7-a01b09fd5616', '2025-06-01 14:55:36.106701+00');
INSERT INTO public.likes VALUES ('97633637-984b-4120-bb0b-eff3ce071b5f', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '866d6d95-29d3-41e2-adb7-a01b09fd5616', '2025-06-07 12:48:36.106713+00');
INSERT INTO public.likes VALUES ('528b78c4-e0b7-4c02-9d1d-ea9b829ed07e', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '866d6d95-29d3-41e2-adb7-a01b09fd5616', '2025-06-05 07:04:36.106729+00');
INSERT INTO public.likes VALUES ('1163075f-2a26-4c60-9c58-6b9f6e7adbff', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '866d6d95-29d3-41e2-adb7-a01b09fd5616', '2025-06-06 06:44:36.106743+00');
INSERT INTO public.likes VALUES ('4adc0801-7078-495b-97e3-f00c184700ee', 'f399a996-1567-4fc0-a86a-92774494ca44', '866d6d95-29d3-41e2-adb7-a01b09fd5616', '2025-06-04 12:37:36.106762+00');
INSERT INTO public.likes VALUES ('09e65d77-0ca3-448b-89ae-ffe9ad31875c', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '866d6d95-29d3-41e2-adb7-a01b09fd5616', '2025-06-03 22:25:36.106782+00');
INSERT INTO public.likes VALUES ('b0f60dfd-2297-4868-a81d-c1d1c6378aaf', '79758f94-70d8-4a11-b05f-793fcb57e093', '866d6d95-29d3-41e2-adb7-a01b09fd5616', '2025-06-04 13:36:36.106796+00');
INSERT INTO public.likes VALUES ('ab0e42ab-0d05-45cf-ae87-4df55095fc39', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '8b91fce6-8708-47e5-b7c5-9eee8c55b696', '2025-06-04 14:08:36.106818+00');
INSERT INTO public.likes VALUES ('ba7f92f8-2db2-41ad-a992-6b19aa97b64c', '79758f94-70d8-4a11-b05f-793fcb57e093', '8b91fce6-8708-47e5-b7c5-9eee8c55b696', '2025-06-07 16:06:36.106828+00');
INSERT INTO public.likes VALUES ('f98cbcc5-bdb2-416f-898b-91c4762ff78f', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', '8b91fce6-8708-47e5-b7c5-9eee8c55b696', '2025-06-02 07:52:36.106836+00');
INSERT INTO public.likes VALUES ('f727c41c-ebd4-4d66-8b7e-2ea5fcedad1a', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '8b91fce6-8708-47e5-b7c5-9eee8c55b696', '2025-06-02 10:29:36.106845+00');
INSERT INTO public.likes VALUES ('60a36908-cc63-490e-852b-2b54a8ae5143', '56e89866-3a90-4f97-8461-34b24f842da4', '8b91fce6-8708-47e5-b7c5-9eee8c55b696', '2025-06-08 07:54:36.106853+00');
INSERT INTO public.likes VALUES ('7a609c2d-a16c-4501-a7da-11fce0c2826e', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', '91d813fb-0bb2-4a9f-98cd-40c4c86de628', '2025-06-01 13:42:36.106873+00');
INSERT INTO public.likes VALUES ('79a775a7-41cb-490f-9071-b5c7463081d6', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '91d813fb-0bb2-4a9f-98cd-40c4c86de628', '2025-06-04 18:53:36.106884+00');
INSERT INTO public.likes VALUES ('71fdb617-97e8-4791-b3dd-c21a6a8d8437', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '91d813fb-0bb2-4a9f-98cd-40c4c86de628', '2025-06-07 13:39:36.1069+00');
INSERT INTO public.likes VALUES ('244b7037-8763-4a5a-a8fe-26d521886257', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '91d813fb-0bb2-4a9f-98cd-40c4c86de628', '2025-06-05 20:48:36.106916+00');
INSERT INTO public.likes VALUES ('2bf94509-b38c-4d40-92de-45d3361756fe', '56e89866-3a90-4f97-8461-34b24f842da4', '91d813fb-0bb2-4a9f-98cd-40c4c86de628', '2025-06-07 17:31:36.106952+00');
INSERT INTO public.likes VALUES ('6acf4312-eddc-4033-b497-903ad06274ea', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '91d813fb-0bb2-4a9f-98cd-40c4c86de628', '2025-06-06 10:15:36.107117+00');
INSERT INTO public.likes VALUES ('83d359de-4216-4d3b-9054-a45ef55f169c', '79758f94-70d8-4a11-b05f-793fcb57e093', '91d813fb-0bb2-4a9f-98cd-40c4c86de628', '2025-06-03 08:30:36.107179+00');
INSERT INTO public.likes VALUES ('13a15936-045b-47ed-bd3e-9e330c85d188', 'f399a996-1567-4fc0-a86a-92774494ca44', '91d813fb-0bb2-4a9f-98cd-40c4c86de628', '2025-06-05 00:14:36.107192+00');
INSERT INTO public.likes VALUES ('24320051-69ae-4fa8-a1cb-a9248037e9e1', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '91d813fb-0bb2-4a9f-98cd-40c4c86de628', '2025-06-05 18:56:36.107201+00');
INSERT INTO public.likes VALUES ('26b4a88a-a15d-4417-ac03-a83f63eb1f37', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '91d813fb-0bb2-4a9f-98cd-40c4c86de628', '2025-06-06 17:18:36.107211+00');
INSERT INTO public.likes VALUES ('b0550428-af37-46d1-ba6b-edacf0777a57', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '97000960-eabc-450b-b074-609480dc17e8', '2025-06-06 02:15:36.107248+00');
INSERT INTO public.likes VALUES ('73d83fca-3523-4397-b450-79c9288329a0', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '97000960-eabc-450b-b074-609480dc17e8', '2025-06-03 18:14:36.107259+00');
INSERT INTO public.likes VALUES ('dd0cb43a-1ca4-4233-93f5-867467fc011b', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '97000960-eabc-450b-b074-609480dc17e8', '2025-06-04 09:21:36.107267+00');
INSERT INTO public.likes VALUES ('084b0617-f245-4841-9e2e-4de6c874704c', '6e137217-71a9-439b-a579-4bcd1a729af9', '97000960-eabc-450b-b074-609480dc17e8', '2025-06-05 23:35:36.107276+00');
INSERT INTO public.likes VALUES ('374a53ca-44b5-4dce-b52f-6ffdea4a765e', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '97000960-eabc-450b-b074-609480dc17e8', '2025-06-04 19:56:36.107285+00');
INSERT INTO public.likes VALUES ('23444e9f-9d46-42c6-8507-992bff71f711', 'f399a996-1567-4fc0-a86a-92774494ca44', '97000960-eabc-450b-b074-609480dc17e8', '2025-06-02 08:01:36.107294+00');
INSERT INTO public.likes VALUES ('88a4325e-d1bf-4881-93c7-39282cdc39fd', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '970217d3-a37f-4bea-a894-b26a21d7845f', '2025-06-08 09:58:36.107352+00');
INSERT INTO public.likes VALUES ('9dfd40e7-87d1-47d8-b4e5-109e8c7224e4', '3b0ebec9-3805-490b-987f-4c6e7286319e', '970217d3-a37f-4bea-a894-b26a21d7845f', '2025-06-07 21:05:36.107363+00');
INSERT INTO public.likes VALUES ('b0b8ea9b-aa45-4755-b851-24f9a30661ff', '56e89866-3a90-4f97-8461-34b24f842da4', '970217d3-a37f-4bea-a894-b26a21d7845f', '2025-06-07 17:11:36.107371+00');
INSERT INTO public.likes VALUES ('fc9ee984-860b-4378-a260-ebf72ddf96bc', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', '970217d3-a37f-4bea-a894-b26a21d7845f', '2025-06-03 00:51:36.10738+00');
INSERT INTO public.likes VALUES ('eb42fe53-6001-406b-a13a-70ef674d4fed', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '970217d3-a37f-4bea-a894-b26a21d7845f', '2025-06-06 16:16:36.107389+00');
INSERT INTO public.likes VALUES ('ec7d6861-34f2-46a9-bb00-bff7ed5fbede', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '970217d3-a37f-4bea-a894-b26a21d7845f', '2025-06-02 06:19:36.107397+00');
INSERT INTO public.likes VALUES ('a1ad5518-1e5d-4477-b6ff-21f549e1fc9f', '79758f94-70d8-4a11-b05f-793fcb57e093', '970217d3-a37f-4bea-a894-b26a21d7845f', '2025-06-07 05:40:36.107406+00');
INSERT INTO public.likes VALUES ('af33df52-576e-45e4-8703-0357db7f457d', '6e137217-71a9-439b-a579-4bcd1a729af9', '970217d3-a37f-4bea-a894-b26a21d7845f', '2025-06-01 17:36:36.107415+00');
INSERT INTO public.likes VALUES ('06021efb-d518-46cd-91c4-b57f342f9ac6', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '970217d3-a37f-4bea-a894-b26a21d7845f', '2025-06-03 15:40:36.107424+00');
INSERT INTO public.likes VALUES ('78d332c1-04fd-4795-afae-79c32f65f4b5', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '970217d3-a37f-4bea-a894-b26a21d7845f', '2025-06-07 12:37:36.107438+00');
INSERT INTO public.likes VALUES ('f89977bc-b82e-4d8a-9e94-2753e5a9f942', 'f399a996-1567-4fc0-a86a-92774494ca44', '970217d3-a37f-4bea-a894-b26a21d7845f', '2025-06-05 15:23:36.107447+00');
INSERT INTO public.likes VALUES ('8f5b8fd2-12df-4b47-8acc-89d52225e63a', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'd3aef78d-bbf8-4543-bda4-1946e61bcc4b', '2025-06-06 08:02:36.108641+00');
INSERT INTO public.likes VALUES ('50069c73-662d-466f-9032-7d96cec4d149', '56e89866-3a90-4f97-8461-34b24f842da4', 'd3aef78d-bbf8-4543-bda4-1946e61bcc4b', '2025-06-08 10:54:36.108803+00');
INSERT INTO public.likes VALUES ('096252c2-21e7-4325-9450-e59d41e2c4fc', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', 'd3aef78d-bbf8-4543-bda4-1946e61bcc4b', '2025-06-02 17:31:36.108908+00');
INSERT INTO public.likes VALUES ('2a5aa79a-b14d-4f89-a5c3-4b851c1f0156', '6e137217-71a9-439b-a579-4bcd1a729af9', 'd3aef78d-bbf8-4543-bda4-1946e61bcc4b', '2025-06-05 05:09:36.108966+00');
INSERT INTO public.likes VALUES ('88fc8c39-c1fe-47d5-b5eb-2dd911befdc8', 'f399a996-1567-4fc0-a86a-92774494ca44', 'd3aef78d-bbf8-4543-bda4-1946e61bcc4b', '2025-06-04 16:56:36.109+00');
INSERT INTO public.likes VALUES ('94e3b474-d793-475f-916c-0a25135691ad', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', 'd3aef78d-bbf8-4543-bda4-1946e61bcc4b', '2025-06-07 13:13:36.109016+00');
INSERT INTO public.likes VALUES ('dcfdc3cf-8858-49cd-87b3-f9ffc8c8979c', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'd3aef78d-bbf8-4543-bda4-1946e61bcc4b', '2025-06-05 14:35:36.10903+00');
INSERT INTO public.likes VALUES ('58c7de19-cf49-4234-ab3d-5580c27e5e78', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', 'd3aef78d-bbf8-4543-bda4-1946e61bcc4b', '2025-06-02 13:06:36.109044+00');
INSERT INTO public.likes VALUES ('d24e2448-7b85-4605-868a-cd182c092013', '79758f94-70d8-4a11-b05f-793fcb57e093', 'd3aef78d-bbf8-4543-bda4-1946e61bcc4b', '2025-06-03 17:49:36.109059+00');
INSERT INTO public.likes VALUES ('3161c9e7-4d46-4fae-a1d9-9286836e1ce7', '6e137217-71a9-439b-a579-4bcd1a729af9', 'a15c1a09-2a18-4cc2-9915-6c3aa7007455', NULL);
INSERT INTO public.likes VALUES ('e3d5fd3a-f8b9-4aa2-a947-4d539458573c', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', 'a15c1a09-2a18-4cc2-9915-6c3aa7007455', NULL);
INSERT INTO public.likes VALUES ('800c0772-af22-4ee0-8323-633c17b592c1', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', 'a15c1a09-2a18-4cc2-9915-6c3aa7007455', NULL);
INSERT INTO public.likes VALUES ('df2db509-6c5b-4be4-bd64-2b9ea672367a', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'a15c1a09-2a18-4cc2-9915-6c3aa7007455', NULL);
INSERT INTO public.likes VALUES ('134e8143-e975-43a7-9a72-30638ed12cbd', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', 'a15c1a09-2a18-4cc2-9915-6c3aa7007455', NULL);
INSERT INTO public.likes VALUES ('0e8b496f-396e-4196-8d63-35db9603e138', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', 'a15c1a09-2a18-4cc2-9915-6c3aa7007455', NULL);
INSERT INTO public.likes VALUES ('2572ce1d-6c0c-4867-8eb9-60cc950bfec6', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'a15c1a09-2a18-4cc2-9915-6c3aa7007455', NULL);
INSERT INTO public.likes VALUES ('a116a100-a695-4ce0-836d-26c5191953a3', 'f399a996-1567-4fc0-a86a-92774494ca44', 'a15c1a09-2a18-4cc2-9915-6c3aa7007455', NULL);
INSERT INTO public.likes VALUES ('817cf807-8864-4363-9ac9-9fe1a105f556', '79758f94-70d8-4a11-b05f-793fcb57e093', '6a82001d-b416-4515-913e-5b6e3f57fbdf', NULL);
INSERT INTO public.likes VALUES ('e44b8bc8-869a-4c7f-ad7c-02740f7a710e', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '6a82001d-b416-4515-913e-5b6e3f57fbdf', NULL);
INSERT INTO public.likes VALUES ('96175716-aa0e-42ca-a9fd-36233b7014d7', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '6a82001d-b416-4515-913e-5b6e3f57fbdf', NULL);
INSERT INTO public.likes VALUES ('49ef84f6-183b-48db-928d-7400ef356650', '6e137217-71a9-439b-a579-4bcd1a729af9', '6a82001d-b416-4515-913e-5b6e3f57fbdf', NULL);
INSERT INTO public.likes VALUES ('49cffb70-b7ff-44f5-9df9-5cd05a1b686d', '56e89866-3a90-4f97-8461-34b24f842da4', '6a82001d-b416-4515-913e-5b6e3f57fbdf', NULL);
INSERT INTO public.likes VALUES ('d3ac646f-519a-49ce-aeec-3949fa750b64', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '6a82001d-b416-4515-913e-5b6e3f57fbdf', NULL);
INSERT INTO public.likes VALUES ('91356577-4ab1-4acc-a816-198c083cb1dc', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '6a82001d-b416-4515-913e-5b6e3f57fbdf', NULL);
INSERT INTO public.likes VALUES ('d9f20255-e22a-459b-95b1-023636761791', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '6a82001d-b416-4515-913e-5b6e3f57fbdf', NULL);
INSERT INTO public.likes VALUES ('ac303033-fb56-4959-8bc2-ba5aecb8519f', '3b0ebec9-3805-490b-987f-4c6e7286319e', '6a82001d-b416-4515-913e-5b6e3f57fbdf', NULL);
INSERT INTO public.likes VALUES ('ad3469a2-5ea3-46ba-b9e7-7dfa6b827eca', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '6a82001d-b416-4515-913e-5b6e3f57fbdf', NULL);
INSERT INTO public.likes VALUES ('6b8e887e-e193-4385-a8aa-b14da70f47c6', '6e137217-71a9-439b-a579-4bcd1a729af9', 'b3621746-9032-46c1-9e9c-072c55b709a7', NULL);
INSERT INTO public.likes VALUES ('cfba3a85-1b25-449a-85d5-e7fd21343e9f', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'b3621746-9032-46c1-9e9c-072c55b709a7', NULL);
INSERT INTO public.likes VALUES ('1a5a7e3d-02fa-4192-b562-6341ea131868', '79758f94-70d8-4a11-b05f-793fcb57e093', 'b3621746-9032-46c1-9e9c-072c55b709a7', NULL);
INSERT INTO public.likes VALUES ('c7f3b8d9-503d-40dd-88bc-002392208ef6', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', 'b3621746-9032-46c1-9e9c-072c55b709a7', NULL);
INSERT INTO public.likes VALUES ('f109ff56-79f9-4a84-aee2-1f204d7149e8', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'b3621746-9032-46c1-9e9c-072c55b709a7', NULL);
INSERT INTO public.likes VALUES ('54f09ae5-74bc-47c5-b865-ec5fb37ec97f', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'b3621746-9032-46c1-9e9c-072c55b709a7', NULL);
INSERT INTO public.likes VALUES ('e56a1514-2c00-4634-bda6-bbefdcfb29e5', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', 'b3621746-9032-46c1-9e9c-072c55b709a7', NULL);
INSERT INTO public.likes VALUES ('3b0bb46c-aa93-46f6-9c4f-12d3740a5947', '56e89866-3a90-4f97-8461-34b24f842da4', 'b3621746-9032-46c1-9e9c-072c55b709a7', NULL);
INSERT INTO public.likes VALUES ('48510d39-b9bb-4644-9219-42155723f2e5', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '223da9a0-82b1-4745-9a4a-5b657d47299e', NULL);
INSERT INTO public.likes VALUES ('e03f3a9f-da90-43ec-9c84-0da214d10820', '3b0ebec9-3805-490b-987f-4c6e7286319e', '223da9a0-82b1-4745-9a4a-5b657d47299e', NULL);
INSERT INTO public.likes VALUES ('efefaf51-b047-4f62-a0bc-cbfb143ed28b', 'f399a996-1567-4fc0-a86a-92774494ca44', '223da9a0-82b1-4745-9a4a-5b657d47299e', NULL);
INSERT INTO public.likes VALUES ('31bdd026-6b34-4855-9b33-5f01b8680b56', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '223da9a0-82b1-4745-9a4a-5b657d47299e', NULL);
INSERT INTO public.likes VALUES ('b2cd24ed-c95d-4cd8-a00d-bbe1ae975b48', '6e137217-71a9-439b-a579-4bcd1a729af9', '223da9a0-82b1-4745-9a4a-5b657d47299e', NULL);
INSERT INTO public.likes VALUES ('5d78eb2c-ee40-4794-a102-1bfd8f54775c', 'f399a996-1567-4fc0-a86a-92774494ca44', '807dc60c-27a5-44e8-a20c-0ee8ce41ede8', NULL);
INSERT INTO public.likes VALUES ('012e5274-f174-4c42-9d5b-6d50a308e159', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', '807dc60c-27a5-44e8-a20c-0ee8ce41ede8', NULL);
INSERT INTO public.likes VALUES ('2828ea1f-0cfa-4504-bbc2-77a4a0f49d74', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '807dc60c-27a5-44e8-a20c-0ee8ce41ede8', NULL);
INSERT INTO public.likes VALUES ('849cc274-4035-48a8-b2e2-8eb2c5ad1a40', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '807dc60c-27a5-44e8-a20c-0ee8ce41ede8', NULL);
INSERT INTO public.likes VALUES ('04a67002-c760-4ac2-9838-9ce3c3744561', '3b0ebec9-3805-490b-987f-4c6e7286319e', '807dc60c-27a5-44e8-a20c-0ee8ce41ede8', NULL);
INSERT INTO public.likes VALUES ('6605cd00-527d-427e-8344-becc29bad1d0', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '807dc60c-27a5-44e8-a20c-0ee8ce41ede8', NULL);
INSERT INTO public.likes VALUES ('23bece5f-faab-48fd-a819-a2c26a6588d7', '6e137217-71a9-439b-a579-4bcd1a729af9', '807dc60c-27a5-44e8-a20c-0ee8ce41ede8', NULL);
INSERT INTO public.likes VALUES ('12dd8978-ec3a-42ef-86a7-259ef542cca7', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '219facb1-badb-49bc-9395-e2c980236804', NULL);
INSERT INTO public.likes VALUES ('05011793-0ce5-418c-82e2-6f2c5f6879c6', '56e89866-3a90-4f97-8461-34b24f842da4', '219facb1-badb-49bc-9395-e2c980236804', NULL);
INSERT INTO public.likes VALUES ('622d2eb6-4010-47d4-8019-c852fe0cb656', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '219facb1-badb-49bc-9395-e2c980236804', NULL);
INSERT INTO public.likes VALUES ('5807a339-7b14-4f89-8b1d-f7d14d140516', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', '219facb1-badb-49bc-9395-e2c980236804', NULL);
INSERT INTO public.likes VALUES ('03deb71e-99c9-4d6e-9950-6c3c0108453d', '3b0ebec9-3805-490b-987f-4c6e7286319e', '219facb1-badb-49bc-9395-e2c980236804', NULL);
INSERT INTO public.likes VALUES ('2dd8ec78-0989-422c-b027-78bad9fc3ed5', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', 'b8ea0106-08f2-45fc-92ea-77a61af34a1d', NULL);
INSERT INTO public.likes VALUES ('5bba2156-4582-4c04-93b3-1731b0164a32', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', 'b8ea0106-08f2-45fc-92ea-77a61af34a1d', NULL);
INSERT INTO public.likes VALUES ('046bd5f0-6c45-4a7a-a33d-3201bb282a17', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'b8ea0106-08f2-45fc-92ea-77a61af34a1d', NULL);
INSERT INTO public.likes VALUES ('5f1370a6-2601-447a-b685-14597050716e', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'b8ea0106-08f2-45fc-92ea-77a61af34a1d', NULL);
INSERT INTO public.likes VALUES ('dba8d1fe-7aba-47b9-9891-977ae6fe86a3', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'b8ea0106-08f2-45fc-92ea-77a61af34a1d', NULL);
INSERT INTO public.likes VALUES ('b8e92a2e-84fe-4010-8512-8181ba675639', '56e89866-3a90-4f97-8461-34b24f842da4', 'b8ea0106-08f2-45fc-92ea-77a61af34a1d', NULL);
INSERT INTO public.likes VALUES ('51a7fd6c-86aa-4aa5-81a7-01362bdb802c', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'dee234bb-394b-4471-a913-7e6153be0bb7', NULL);
INSERT INTO public.likes VALUES ('78f056e9-16a2-4fc6-be5b-5e6524a10c8c', 'f399a996-1567-4fc0-a86a-92774494ca44', 'dee234bb-394b-4471-a913-7e6153be0bb7', NULL);
INSERT INTO public.likes VALUES ('befed642-d74e-4fc7-b19c-7763d91f75df', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'dee234bb-394b-4471-a913-7e6153be0bb7', NULL);
INSERT INTO public.likes VALUES ('aec7dc2b-4422-434e-aa87-886019b97995', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', 'dee234bb-394b-4471-a913-7e6153be0bb7', NULL);
INSERT INTO public.likes VALUES ('1012c204-9429-48a5-a222-8a5411f9edc9', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', 'dee234bb-394b-4471-a913-7e6153be0bb7', NULL);
INSERT INTO public.likes VALUES ('05aac842-ba6f-4218-a72f-48f32041558d', '56e89866-3a90-4f97-8461-34b24f842da4', '14cd020e-3ee9-4114-829d-e3b66173c5f0', NULL);
INSERT INTO public.likes VALUES ('87bbc734-824c-45c1-a108-04c5b33701bc', 'f399a996-1567-4fc0-a86a-92774494ca44', '14cd020e-3ee9-4114-829d-e3b66173c5f0', NULL);
INSERT INTO public.likes VALUES ('4857ee42-e6a3-4f86-a37e-aaa50e426740', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', '14cd020e-3ee9-4114-829d-e3b66173c5f0', NULL);
INSERT INTO public.likes VALUES ('710db9ae-8a2f-4a64-a5b9-a330ef66853d', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '14cd020e-3ee9-4114-829d-e3b66173c5f0', NULL);
INSERT INTO public.likes VALUES ('bb58140e-c656-4acf-aae7-713f72c6d797', '6e137217-71a9-439b-a579-4bcd1a729af9', '14cd020e-3ee9-4114-829d-e3b66173c5f0', NULL);
INSERT INTO public.likes VALUES ('457d302e-3e7b-4540-a45c-b96af561facd', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '14cd020e-3ee9-4114-829d-e3b66173c5f0', NULL);
INSERT INTO public.likes VALUES ('2c68c92a-a9c1-4304-af0c-edf1b3692eee', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'f4d4efdf-82d7-491f-b9fc-4d10d40e6a1b', NULL);
INSERT INTO public.likes VALUES ('26dcc0c8-40c2-447b-81e3-67887090e61d', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'f4d4efdf-82d7-491f-b9fc-4d10d40e6a1b', NULL);
INSERT INTO public.likes VALUES ('b20f65a3-73c2-47f4-9749-1281346aa747', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'f4d4efdf-82d7-491f-b9fc-4d10d40e6a1b', NULL);
INSERT INTO public.likes VALUES ('b22375d7-1b37-4aa0-9289-b943364a4501', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', 'f4d4efdf-82d7-491f-b9fc-4d10d40e6a1b', NULL);
INSERT INTO public.likes VALUES ('c11f4117-b518-41a6-8e76-aef83ea8a29b', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', 'f4d4efdf-82d7-491f-b9fc-4d10d40e6a1b', NULL);
INSERT INTO public.likes VALUES ('eded7c78-0c2e-424d-bfa4-299f1f15b09f', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '71a82c6e-6511-4cf5-8a80-099a7ee1f829', NULL);
INSERT INTO public.likes VALUES ('91cb2459-e2ce-4119-a8a0-2326981b886a', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', '71a82c6e-6511-4cf5-8a80-099a7ee1f829', NULL);
INSERT INTO public.likes VALUES ('4e3606a3-acf4-46ce-b192-9b6938b8023e', '3b0ebec9-3805-490b-987f-4c6e7286319e', '71a82c6e-6511-4cf5-8a80-099a7ee1f829', NULL);
INSERT INTO public.likes VALUES ('b8ea38be-b4b2-46b4-a26a-3763049a5771', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '71a82c6e-6511-4cf5-8a80-099a7ee1f829', NULL);
INSERT INTO public.likes VALUES ('935ed2d6-8000-47c8-8fcb-0d6d00aa6ccf', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '71a82c6e-6511-4cf5-8a80-099a7ee1f829', NULL);
INSERT INTO public.likes VALUES ('0e02bcda-394e-4993-a1a1-d6ad2edc255e', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '71a82c6e-6511-4cf5-8a80-099a7ee1f829', NULL);
INSERT INTO public.likes VALUES ('ff126039-c827-40f7-b023-f2e757feb08f', '79758f94-70d8-4a11-b05f-793fcb57e093', '71a82c6e-6511-4cf5-8a80-099a7ee1f829', NULL);
INSERT INTO public.likes VALUES ('bf096e90-6095-4908-b0fe-cfed476a2f8f', 'f399a996-1567-4fc0-a86a-92774494ca44', '824b208f-844a-4b1c-b9e2-ce7e8148f06e', NULL);
INSERT INTO public.likes VALUES ('91385742-a113-4be1-b951-89988c0a5a50', '6e137217-71a9-439b-a579-4bcd1a729af9', '824b208f-844a-4b1c-b9e2-ce7e8148f06e', NULL);
INSERT INTO public.likes VALUES ('1bf6df31-cacc-43bf-a221-c21ac8486292', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', '824b208f-844a-4b1c-b9e2-ce7e8148f06e', NULL);
INSERT INTO public.likes VALUES ('8a268a32-bfc8-4940-870b-2ae5f8938e22', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '824b208f-844a-4b1c-b9e2-ce7e8148f06e', NULL);
INSERT INTO public.likes VALUES ('ef98da23-e537-4768-af94-5d026854d731', '79758f94-70d8-4a11-b05f-793fcb57e093', '824b208f-844a-4b1c-b9e2-ce7e8148f06e', NULL);
INSERT INTO public.likes VALUES ('2d8a9d94-035e-4145-acc0-f051c46e134f', '3b0ebec9-3805-490b-987f-4c6e7286319e', '824b208f-844a-4b1c-b9e2-ce7e8148f06e', NULL);
INSERT INTO public.likes VALUES ('ad531a02-f284-4ad9-9252-0e9fd7cc3d81', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '824b208f-844a-4b1c-b9e2-ce7e8148f06e', NULL);
INSERT INTO public.likes VALUES ('eb65f5c7-ac96-4839-8929-a13a7d18b7b6', '6e137217-71a9-439b-a579-4bcd1a729af9', '9a0f1426-9829-4c07-89b4-288a6e2d2a33', NULL);
INSERT INTO public.likes VALUES ('67da13f4-d25e-4ede-b539-a596208317b5', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '9a0f1426-9829-4c07-89b4-288a6e2d2a33', NULL);
INSERT INTO public.likes VALUES ('7bb603d5-c43f-487b-80e6-3ee6efaec4a5', '79758f94-70d8-4a11-b05f-793fcb57e093', '9a0f1426-9829-4c07-89b4-288a6e2d2a33', NULL);
INSERT INTO public.likes VALUES ('5446877f-75ec-466a-b56d-d6b9b088fd16', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '9a0f1426-9829-4c07-89b4-288a6e2d2a33', NULL);
INSERT INTO public.likes VALUES ('4c04f459-b5d1-4b06-81ba-ec1606389489', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', 'acc9d130-e2ea-426e-9553-d3ee35589a7f', NULL);
INSERT INTO public.likes VALUES ('82a8c707-46f9-49de-a554-764ce3bec7af', '79758f94-70d8-4a11-b05f-793fcb57e093', 'acc9d130-e2ea-426e-9553-d3ee35589a7f', NULL);
INSERT INTO public.likes VALUES ('fceea6df-c8db-4b01-adb9-a8b86c1b2914', '6e137217-71a9-439b-a579-4bcd1a729af9', 'acc9d130-e2ea-426e-9553-d3ee35589a7f', NULL);
INSERT INTO public.likes VALUES ('ce91a302-59dd-4559-83bc-4cab934aebb1', 'f399a996-1567-4fc0-a86a-92774494ca44', 'acc9d130-e2ea-426e-9553-d3ee35589a7f', NULL);
INSERT INTO public.likes VALUES ('1b06843e-399d-4665-b505-a96884c84524', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', 'acc9d130-e2ea-426e-9553-d3ee35589a7f', NULL);
INSERT INTO public.likes VALUES ('edf6df40-60e2-4443-b28b-1a96dd43734b', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', 'acc9d130-e2ea-426e-9553-d3ee35589a7f', NULL);
INSERT INTO public.likes VALUES ('2672a39d-20ed-4e34-bd83-7c9cec21b031', '56e89866-3a90-4f97-8461-34b24f842da4', 'acc9d130-e2ea-426e-9553-d3ee35589a7f', NULL);
INSERT INTO public.likes VALUES ('025a2faf-b09f-4ef3-b66d-3d42dee69627', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'acc9d130-e2ea-426e-9553-d3ee35589a7f', NULL);
INSERT INTO public.likes VALUES ('a94c5196-21f7-46b7-9f04-17ee871b15ce', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'acc9d130-e2ea-426e-9553-d3ee35589a7f', NULL);
INSERT INTO public.likes VALUES ('98cda13e-4473-4726-bc73-9c36e1125fcc', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'c2901639-580e-46f1-a0b9-12e6a970e7d4', NULL);
INSERT INTO public.likes VALUES ('22c81e5a-0030-4d48-a14a-d780717ae69e', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'c2901639-580e-46f1-a0b9-12e6a970e7d4', NULL);
INSERT INTO public.likes VALUES ('96f74bfd-eb78-45a3-8b14-b3fc6b710a70', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'c2901639-580e-46f1-a0b9-12e6a970e7d4', NULL);
INSERT INTO public.likes VALUES ('d545df61-4eda-444d-a54c-c82a9f47b056', '79758f94-70d8-4a11-b05f-793fcb57e093', 'c2901639-580e-46f1-a0b9-12e6a970e7d4', NULL);
INSERT INTO public.likes VALUES ('006fc1ce-29f9-4038-8026-a6b7eb19c8f4', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', 'c2901639-580e-46f1-a0b9-12e6a970e7d4', NULL);
INSERT INTO public.likes VALUES ('b32e6359-f5fa-4698-991c-ea8b29ac777c', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', 'c2901639-580e-46f1-a0b9-12e6a970e7d4', NULL);
INSERT INTO public.likes VALUES ('36b76310-afcc-4832-af58-070479f16dd6', '6e137217-71a9-439b-a579-4bcd1a729af9', 'c2901639-580e-46f1-a0b9-12e6a970e7d4', NULL);
INSERT INTO public.likes VALUES ('d73d20d2-5cd6-45cf-b44a-ecb41f623df1', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', 'c2901639-580e-46f1-a0b9-12e6a970e7d4', NULL);
INSERT INTO public.likes VALUES ('ad295cfd-7cc0-43da-8f53-6de96726b16e', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'c56780c9-07da-467d-8d5b-5236215603af', NULL);
INSERT INTO public.likes VALUES ('d04f187b-bde2-461b-adf0-d7f747bdad7c', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'c56780c9-07da-467d-8d5b-5236215603af', NULL);
INSERT INTO public.likes VALUES ('276a89c3-6ce6-46fb-a0dd-7c0131deb559', 'f399a996-1567-4fc0-a86a-92774494ca44', 'c56780c9-07da-467d-8d5b-5236215603af', NULL);
INSERT INTO public.likes VALUES ('b447b12e-bedf-4d29-a00b-0dbcb9ca7bab', 'f399a996-1567-4fc0-a86a-92774494ca44', '9508badc-12e8-4390-aa8f-7763d4dbde96', NULL);
INSERT INTO public.likes VALUES ('62711efe-9312-4bf0-92ea-c3c70e1acbee', '6e137217-71a9-439b-a579-4bcd1a729af9', '9508badc-12e8-4390-aa8f-7763d4dbde96', NULL);
INSERT INTO public.likes VALUES ('e6830cdd-52a3-434f-9eb3-4252ada10661', '3b0ebec9-3805-490b-987f-4c6e7286319e', '9508badc-12e8-4390-aa8f-7763d4dbde96', NULL);
INSERT INTO public.likes VALUES ('71e0e00f-2ed1-417a-b7d5-c1a5d1c73958', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '9508badc-12e8-4390-aa8f-7763d4dbde96', NULL);
INSERT INTO public.likes VALUES ('b150126e-e3bd-438a-ab24-834b97f2d04a', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '9508badc-12e8-4390-aa8f-7763d4dbde96', NULL);
INSERT INTO public.likes VALUES ('3f0a5ab9-f620-4008-9550-cb97d47f1ea9', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', '9508badc-12e8-4390-aa8f-7763d4dbde96', NULL);
INSERT INTO public.likes VALUES ('27b0a3df-df2a-4940-bdde-0a875cf15fe0', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '9508badc-12e8-4390-aa8f-7763d4dbde96', NULL);
INSERT INTO public.likes VALUES ('9d184ef6-f17d-4567-8424-1809633461fb', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '9508badc-12e8-4390-aa8f-7763d4dbde96', NULL);
INSERT INTO public.likes VALUES ('dbe8fd8d-1bac-4b15-a9f3-fe79bd8af129', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'e5ac024b-a6ee-469f-834f-704f37620b25', NULL);
INSERT INTO public.likes VALUES ('5448c73b-2e9f-4ca6-a65b-385067204ebc', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', 'e5ac024b-a6ee-469f-834f-704f37620b25', NULL);
INSERT INTO public.likes VALUES ('0381fba2-ef2d-438a-8b05-af5106873e07', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'e5ac024b-a6ee-469f-834f-704f37620b25', NULL);
INSERT INTO public.likes VALUES ('c36850e9-4550-4017-9921-8a44704c544b', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'e5ac024b-a6ee-469f-834f-704f37620b25', NULL);
INSERT INTO public.likes VALUES ('8187f449-5521-4773-8845-beed95eebae7', '56e89866-3a90-4f97-8461-34b24f842da4', 'e5ac024b-a6ee-469f-834f-704f37620b25', NULL);
INSERT INTO public.likes VALUES ('0890eebe-ae1c-4864-a47e-ddcbea0354b0', '6e137217-71a9-439b-a579-4bcd1a729af9', 'e5ac024b-a6ee-469f-834f-704f37620b25', NULL);
INSERT INTO public.likes VALUES ('75788b80-741e-40cd-80b0-25b69730aaec', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', 'e5ac024b-a6ee-469f-834f-704f37620b25', NULL);
INSERT INTO public.likes VALUES ('10a37038-655d-4b1a-acc4-f3d3777facc2', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', 'e5ac024b-a6ee-469f-834f-704f37620b25', NULL);
INSERT INTO public.likes VALUES ('47ee4efa-6307-4e71-b07e-9b543bf16deb', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'c4e4435b-e210-481c-9603-880796a0b7df', NULL);
INSERT INTO public.likes VALUES ('e931e4c3-6e71-4e2e-94ce-f88c8a4d712e', '6e137217-71a9-439b-a579-4bcd1a729af9', 'c4e4435b-e210-481c-9603-880796a0b7df', NULL);
INSERT INTO public.likes VALUES ('8e1ecc17-15ad-414c-990d-5bc1f4b3b3c7', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', 'c4e4435b-e210-481c-9603-880796a0b7df', NULL);
INSERT INTO public.likes VALUES ('fc3c4085-e560-4464-b408-8960ec27d6c7', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'c4e4435b-e210-481c-9603-880796a0b7df', NULL);
INSERT INTO public.likes VALUES ('f4e41ffa-6c51-4c5f-bb49-ea01161b11c4', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'c4e4435b-e210-481c-9603-880796a0b7df', NULL);
INSERT INTO public.likes VALUES ('ef37e597-29d1-4a0b-94c1-190297a85d70', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', 'c4e4435b-e210-481c-9603-880796a0b7df', NULL);
INSERT INTO public.likes VALUES ('2e9e34c8-97ff-494c-ad5b-a18c06d4a1bc', '79758f94-70d8-4a11-b05f-793fcb57e093', 'c4e4435b-e210-481c-9603-880796a0b7df', NULL);
INSERT INTO public.likes VALUES ('9cab7a30-69ba-409a-951f-aeafea76b8ec', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', 'c4e4435b-e210-481c-9603-880796a0b7df', NULL);
INSERT INTO public.likes VALUES ('24af1791-d766-4cdf-9cb9-510c9bfd6a39', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', 'c4e4435b-e210-481c-9603-880796a0b7df', NULL);
INSERT INTO public.likes VALUES ('f1323d5d-8a98-490d-ac0c-4f40fe76c8f5', 'f399a996-1567-4fc0-a86a-92774494ca44', 'c4e4435b-e210-481c-9603-880796a0b7df', NULL);
INSERT INTO public.likes VALUES ('f8399d7b-e503-4f31-8c10-4dbd8953deaa', '56e89866-3a90-4f97-8461-34b24f842da4', '90782b89-1c9c-4c2f-a514-7719331b9c72', NULL);
INSERT INTO public.likes VALUES ('c0b64227-e578-4ef8-aa18-8a6320b61b3e', '79758f94-70d8-4a11-b05f-793fcb57e093', '90782b89-1c9c-4c2f-a514-7719331b9c72', NULL);
INSERT INTO public.likes VALUES ('6f86296c-bd4e-4104-b4cb-7aa5f6d4467e', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', '90782b89-1c9c-4c2f-a514-7719331b9c72', NULL);
INSERT INTO public.likes VALUES ('ff510c8a-5c13-4f31-89b2-c993d5187f62', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '90782b89-1c9c-4c2f-a514-7719331b9c72', NULL);
INSERT INTO public.likes VALUES ('cc86d489-269c-47a4-adff-9ccbad6b286c', '6e137217-71a9-439b-a579-4bcd1a729af9', '90782b89-1c9c-4c2f-a514-7719331b9c72', NULL);
INSERT INTO public.likes VALUES ('84144e0a-35c3-44ac-a648-437ee1152163', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '90782b89-1c9c-4c2f-a514-7719331b9c72', NULL);
INSERT INTO public.likes VALUES ('4863b6b3-4d91-4be0-8002-f2145f8feefd', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '90782b89-1c9c-4c2f-a514-7719331b9c72', NULL);
INSERT INTO public.likes VALUES ('b8b774da-46e5-4bf4-a136-5ff887e3a89a', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '90782b89-1c9c-4c2f-a514-7719331b9c72', NULL);
INSERT INTO public.likes VALUES ('55428b23-0be1-4b6d-ae6d-e321bd0fc44b', 'f399a996-1567-4fc0-a86a-92774494ca44', '90782b89-1c9c-4c2f-a514-7719331b9c72', NULL);
INSERT INTO public.likes VALUES ('11993024-3ea0-495e-af16-85d11fc064e5', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', 'cdcaa739-1c4d-422c-abea-b234c55cc559', NULL);
INSERT INTO public.likes VALUES ('4b79c952-9821-4fe7-b547-777b779e0b1a', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'cdcaa739-1c4d-422c-abea-b234c55cc559', NULL);
INSERT INTO public.likes VALUES ('4f1b6c8f-4fba-413f-9d18-d24f811599bc', 'f399a996-1567-4fc0-a86a-92774494ca44', 'cdcaa739-1c4d-422c-abea-b234c55cc559', NULL);
INSERT INTO public.likes VALUES ('2983f5dc-c7ca-4bc5-9fd6-32137d648ce6', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', 'cdcaa739-1c4d-422c-abea-b234c55cc559', NULL);
INSERT INTO public.likes VALUES ('de19bafc-c80d-47c4-8742-07067080f70c', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', 'cdcaa739-1c4d-422c-abea-b234c55cc559', NULL);
INSERT INTO public.likes VALUES ('0fcef47b-f66c-4f4d-b140-7f01fc0de1fc', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'cdcaa739-1c4d-422c-abea-b234c55cc559', NULL);
INSERT INTO public.likes VALUES ('77a61158-0122-4824-8e79-bc44242e4886', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'cdcaa739-1c4d-422c-abea-b234c55cc559', NULL);
INSERT INTO public.likes VALUES ('9389741f-b57b-417f-9f32-8a5564186d64', '56e89866-3a90-4f97-8461-34b24f842da4', 'cdcaa739-1c4d-422c-abea-b234c55cc559', NULL);
INSERT INTO public.likes VALUES ('61c1cd07-81a0-49e2-b64b-da7f6b500c00', '79758f94-70d8-4a11-b05f-793fcb57e093', 'cdcaa739-1c4d-422c-abea-b234c55cc559', NULL);
INSERT INTO public.likes VALUES ('9d18a763-3a6b-4bc6-ab72-5404156e4f8c', '6e137217-71a9-439b-a579-4bcd1a729af9', 'e003776a-dd4e-4c34-bfe2-503f25a993c3', NULL);
INSERT INTO public.likes VALUES ('b4e9b688-b4c6-47ed-9f8a-6e84392aea45', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', 'e003776a-dd4e-4c34-bfe2-503f25a993c3', NULL);
INSERT INTO public.likes VALUES ('33b2f65f-9c0c-48e1-8a38-964dc4b7e7f2', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'e003776a-dd4e-4c34-bfe2-503f25a993c3', NULL);
INSERT INTO public.likes VALUES ('e0972853-384b-474b-a165-93b7130d07ec', '79758f94-70d8-4a11-b05f-793fcb57e093', 'e003776a-dd4e-4c34-bfe2-503f25a993c3', NULL);
INSERT INTO public.likes VALUES ('a9f0d9f1-d464-4ace-9e2e-6ba3440e4458', 'f399a996-1567-4fc0-a86a-92774494ca44', 'e003776a-dd4e-4c34-bfe2-503f25a993c3', NULL);
INSERT INTO public.likes VALUES ('2ff04b68-1dad-4a91-95cf-5d3c147fc64b', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', 'e003776a-dd4e-4c34-bfe2-503f25a993c3', NULL);
INSERT INTO public.likes VALUES ('149a7b06-ff86-4352-bdfc-32d215423af3', '79758f94-70d8-4a11-b05f-793fcb57e093', '4ddee261-c8d1-4980-9a3e-0baef66bbf7c', NULL);
INSERT INTO public.likes VALUES ('56ae3069-cff9-421f-8ed8-68a9cc7bc01d', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '4ddee261-c8d1-4980-9a3e-0baef66bbf7c', NULL);
INSERT INTO public.likes VALUES ('c67eef67-3927-4a8d-b204-e032304e70ce', '6e137217-71a9-439b-a579-4bcd1a729af9', '4ddee261-c8d1-4980-9a3e-0baef66bbf7c', NULL);
INSERT INTO public.likes VALUES ('9779d95e-a8c1-4d14-b1ff-dfb7dd62db14', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '4ddee261-c8d1-4980-9a3e-0baef66bbf7c', NULL);
INSERT INTO public.likes VALUES ('0e67bf91-1bcd-4340-b712-69cbd325e243', 'f399a996-1567-4fc0-a86a-92774494ca44', '4ddee261-c8d1-4980-9a3e-0baef66bbf7c', NULL);
INSERT INTO public.likes VALUES ('07c7bfb2-0a06-420f-a3a7-f284cf950a9d', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '4ddee261-c8d1-4980-9a3e-0baef66bbf7c', NULL);
INSERT INTO public.likes VALUES ('89a0003e-c6be-4b67-926a-73d0770a4a6d', '3b0ebec9-3805-490b-987f-4c6e7286319e', '4ddee261-c8d1-4980-9a3e-0baef66bbf7c', NULL);
INSERT INTO public.likes VALUES ('9d2324e2-54ce-4a41-9bc5-6378a4de9537', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '4ddee261-c8d1-4980-9a3e-0baef66bbf7c', NULL);
INSERT INTO public.likes VALUES ('deca6f8b-7a42-4dbf-a3a5-33a10ec7b773', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '4ddee261-c8d1-4980-9a3e-0baef66bbf7c', NULL);
INSERT INTO public.likes VALUES ('afd938ab-e674-4076-be74-73298d672c94', '6e137217-71a9-439b-a579-4bcd1a729af9', 'fbbcab19-51fd-4eb4-aa4b-750782900d5d', NULL);
INSERT INTO public.likes VALUES ('4fd6c946-6455-41a3-95f1-4b76cbd51c35', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'fbbcab19-51fd-4eb4-aa4b-750782900d5d', NULL);
INSERT INTO public.likes VALUES ('af101c52-6015-4cb4-b503-3f95268ace7b', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', 'fbbcab19-51fd-4eb4-aa4b-750782900d5d', NULL);
INSERT INTO public.likes VALUES ('1df53647-b3da-43e8-940d-d9cacf6410be', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', 'fbbcab19-51fd-4eb4-aa4b-750782900d5d', NULL);
INSERT INTO public.likes VALUES ('ad56657a-6b31-4597-b5ec-5ebdab145bf6', 'f399a996-1567-4fc0-a86a-92774494ca44', 'fbbcab19-51fd-4eb4-aa4b-750782900d5d', NULL);
INSERT INTO public.likes VALUES ('56fdafa0-f950-4b76-b7c1-836300a89384', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '6a54c0a1-ec0f-45b9-8620-c5d20ba82a03', NULL);
INSERT INTO public.likes VALUES ('3c2d0372-6b64-448b-b49c-480cd3068b19', '79758f94-70d8-4a11-b05f-793fcb57e093', '6a54c0a1-ec0f-45b9-8620-c5d20ba82a03', NULL);
INSERT INTO public.likes VALUES ('5bbdc2dd-ad8a-445e-a5b1-9a8b8467fb4d', '56e89866-3a90-4f97-8461-34b24f842da4', '6a54c0a1-ec0f-45b9-8620-c5d20ba82a03', NULL);
INSERT INTO public.likes VALUES ('4a908a55-04d6-4904-838e-362808be16bf', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '6a54c0a1-ec0f-45b9-8620-c5d20ba82a03', NULL);
INSERT INTO public.likes VALUES ('31cdae9f-aff4-46ee-b41e-e242b12da817', '3b0ebec9-3805-490b-987f-4c6e7286319e', '6a54c0a1-ec0f-45b9-8620-c5d20ba82a03', NULL);
INSERT INTO public.likes VALUES ('a519b3fe-60cd-49c5-82c7-fc8bb12cc0d0', '3b0ebec9-3805-490b-987f-4c6e7286319e', '85a48eda-14df-462a-a173-2c1b653afa4e', NULL);
INSERT INTO public.likes VALUES ('273af40e-713e-4adc-9902-1018282a4fb4', 'f399a996-1567-4fc0-a86a-92774494ca44', '85a48eda-14df-462a-a173-2c1b653afa4e', NULL);
INSERT INTO public.likes VALUES ('10d8e4ef-4a89-43e0-936c-a3d4efd5106b', '79758f94-70d8-4a11-b05f-793fcb57e093', '85a48eda-14df-462a-a173-2c1b653afa4e', NULL);
INSERT INTO public.likes VALUES ('63fa26bd-015e-4bf9-bcae-7f6ce00db19e', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '85a48eda-14df-462a-a173-2c1b653afa4e', NULL);
INSERT INTO public.likes VALUES ('33a55a2b-7ca8-405d-b5c3-1345a878440a', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '85a48eda-14df-462a-a173-2c1b653afa4e', NULL);
INSERT INTO public.likes VALUES ('36d0862a-2e33-4edf-8b0b-fc174f453297', '6e137217-71a9-439b-a579-4bcd1a729af9', '85a48eda-14df-462a-a173-2c1b653afa4e', NULL);
INSERT INTO public.likes VALUES ('812515ce-91a3-4d65-9791-0da991628998', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', 'a21898ff-48aa-4eb6-aeab-050acb249a73', NULL);
INSERT INTO public.likes VALUES ('16bc044c-18a1-437e-b455-c57ed5092297', '79758f94-70d8-4a11-b05f-793fcb57e093', 'a21898ff-48aa-4eb6-aeab-050acb249a73', NULL);
INSERT INTO public.likes VALUES ('f945980d-8e5c-41db-938a-a53c3551d704', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'a21898ff-48aa-4eb6-aeab-050acb249a73', NULL);
INSERT INTO public.likes VALUES ('deb6b805-ca0a-4a8c-bb03-1835ff3a4039', 'f399a996-1567-4fc0-a86a-92774494ca44', 'a21898ff-48aa-4eb6-aeab-050acb249a73', NULL);
INSERT INTO public.likes VALUES ('93eb212b-0df3-4b97-8303-a9fd0ee2802b', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'a21898ff-48aa-4eb6-aeab-050acb249a73', NULL);
INSERT INTO public.likes VALUES ('edd8b968-a22c-48c2-bc9f-098a769531eb', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', 'a21898ff-48aa-4eb6-aeab-050acb249a73', NULL);
INSERT INTO public.likes VALUES ('37b02b36-745d-4739-a641-e6c5faca6320', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'a21898ff-48aa-4eb6-aeab-050acb249a73', NULL);
INSERT INTO public.likes VALUES ('cf13f693-18ec-4b63-a54d-4eecea5e1518', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', 'a21898ff-48aa-4eb6-aeab-050acb249a73', NULL);
INSERT INTO public.likes VALUES ('e64ac0e2-ee2e-40ed-89e2-28c0bf6a5d18', '3b0ebec9-3805-490b-987f-4c6e7286319e', '7f9fcdb1-7602-4b70-8626-204ff26dae55', NULL);
INSERT INTO public.likes VALUES ('617aa502-a8d7-407b-beb8-18d6b222e0c6', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '7f9fcdb1-7602-4b70-8626-204ff26dae55', NULL);
INSERT INTO public.likes VALUES ('6c580cd8-e892-417f-97fc-a3c7bba0356d', '56e89866-3a90-4f97-8461-34b24f842da4', '7f9fcdb1-7602-4b70-8626-204ff26dae55', NULL);
INSERT INTO public.likes VALUES ('48adcea3-fc30-4189-b737-14863154d17f', '6e137217-71a9-439b-a579-4bcd1a729af9', '7f9fcdb1-7602-4b70-8626-204ff26dae55', NULL);
INSERT INTO public.likes VALUES ('7f54bfb0-e89d-4d30-9f9e-ce5b2938d132', '79758f94-70d8-4a11-b05f-793fcb57e093', '93ae9cbf-d56a-4041-b094-ec070ef69045', NULL);
INSERT INTO public.likes VALUES ('aa3c9765-9333-4ea5-9fc1-9d358e78a370', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '93ae9cbf-d56a-4041-b094-ec070ef69045', NULL);
INSERT INTO public.likes VALUES ('e2c07e27-ab1c-4088-b07a-f7fa8aa4d31b', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '93ae9cbf-d56a-4041-b094-ec070ef69045', NULL);
INSERT INTO public.likes VALUES ('9e230789-d3b2-4c49-a7b4-562cbecd628f', '56e89866-3a90-4f97-8461-34b24f842da4', '93ae9cbf-d56a-4041-b094-ec070ef69045', NULL);
INSERT INTO public.likes VALUES ('7d284411-8f33-4ce3-9d38-097c7cc292ab', 'f399a996-1567-4fc0-a86a-92774494ca44', '93ae9cbf-d56a-4041-b094-ec070ef69045', NULL);
INSERT INTO public.likes VALUES ('d25e10c5-3ad0-48ff-9b77-5e3061524676', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '93ae9cbf-d56a-4041-b094-ec070ef69045', NULL);
INSERT INTO public.likes VALUES ('94355f18-0f62-4f9a-8c24-caf88573e77b', '6e137217-71a9-439b-a579-4bcd1a729af9', '93ae9cbf-d56a-4041-b094-ec070ef69045', NULL);
INSERT INTO public.likes VALUES ('fd515b09-a639-4923-a81e-8e90d2ec2232', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '93ae9cbf-d56a-4041-b094-ec070ef69045', NULL);
INSERT INTO public.likes VALUES ('a37df7e6-38fd-4866-a165-0ae788bd2577', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '93ae9cbf-d56a-4041-b094-ec070ef69045', NULL);
INSERT INTO public.likes VALUES ('3bd2a021-1b49-4faf-be10-efd359007f50', '79758f94-70d8-4a11-b05f-793fcb57e093', 'a3f8de46-d113-48a1-b569-a20561d328f5', NULL);
INSERT INTO public.likes VALUES ('9fc829c4-2062-47b9-8586-43606dd221bc', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'a3f8de46-d113-48a1-b569-a20561d328f5', NULL);
INSERT INTO public.likes VALUES ('52d114ab-f706-4c50-a3f5-ba1fb7d7765c', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'a3f8de46-d113-48a1-b569-a20561d328f5', NULL);
INSERT INTO public.likes VALUES ('457241f6-623f-47c4-bc23-2be2b4e0930f', '56e89866-3a90-4f97-8461-34b24f842da4', '8249f900-b093-4d7a-b8f8-67be7fb47448', NULL);
INSERT INTO public.likes VALUES ('f5d8630d-3565-44ce-a3bc-823721771f12', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '8249f900-b093-4d7a-b8f8-67be7fb47448', NULL);
INSERT INTO public.likes VALUES ('53b50a20-8cd1-4232-b8ed-6728deff7aab', '6e137217-71a9-439b-a579-4bcd1a729af9', '8249f900-b093-4d7a-b8f8-67be7fb47448', NULL);
INSERT INTO public.likes VALUES ('4ca21268-8370-4d6a-9ed5-24288547bf1b', '3b0ebec9-3805-490b-987f-4c6e7286319e', '8249f900-b093-4d7a-b8f8-67be7fb47448', NULL);
INSERT INTO public.likes VALUES ('4a5c1039-f535-441a-acea-0dbe1a3853ae', '6e137217-71a9-439b-a579-4bcd1a729af9', 'cf9676ef-e717-4e1e-8ffa-80d6f42d14f5', NULL);
INSERT INTO public.likes VALUES ('03bbfb03-424d-47ff-af33-830e544d148f', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'cf9676ef-e717-4e1e-8ffa-80d6f42d14f5', NULL);
INSERT INTO public.likes VALUES ('6e46aecf-41be-4da5-a321-48e1078c4afe', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', 'cf9676ef-e717-4e1e-8ffa-80d6f42d14f5', NULL);
INSERT INTO public.likes VALUES ('aee9c997-4d64-4db6-819a-05dd64175ad8', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'cf9676ef-e717-4e1e-8ffa-80d6f42d14f5', NULL);
INSERT INTO public.likes VALUES ('41df0533-1d26-43d5-bcfc-00881ee16ce6', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '660c8c08-468b-4a0d-a37d-fe58d7d0b075', NULL);
INSERT INTO public.likes VALUES ('49c99154-5011-4f4c-a0c4-7a86855c6e13', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '660c8c08-468b-4a0d-a37d-fe58d7d0b075', NULL);
INSERT INTO public.likes VALUES ('d3aa3a07-fcfa-4ada-9a9c-d945ef9cef0c', '3b0ebec9-3805-490b-987f-4c6e7286319e', '660c8c08-468b-4a0d-a37d-fe58d7d0b075', NULL);
INSERT INTO public.likes VALUES ('fce53ffb-7520-40d8-8c31-8ed9db258df0', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '660c8c08-468b-4a0d-a37d-fe58d7d0b075', NULL);
INSERT INTO public.likes VALUES ('3b8c3f07-be3e-42bb-b297-4ef0ba9243aa', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '660c8c08-468b-4a0d-a37d-fe58d7d0b075', NULL);
INSERT INTO public.likes VALUES ('f8226ff3-05a7-4e62-a587-d60590e9e1f3', '6e137217-71a9-439b-a579-4bcd1a729af9', '660c8c08-468b-4a0d-a37d-fe58d7d0b075', NULL);
INSERT INTO public.likes VALUES ('08731ad2-32fe-480b-a965-016df9881531', '56e89866-3a90-4f97-8461-34b24f842da4', '660c8c08-468b-4a0d-a37d-fe58d7d0b075', NULL);
INSERT INTO public.likes VALUES ('6a18d303-57d8-4353-9208-94d897d6c972', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '660c8c08-468b-4a0d-a37d-fe58d7d0b075', NULL);
INSERT INTO public.likes VALUES ('ea2bb48d-6153-4e1e-bfc1-e139baab0ae1', '79758f94-70d8-4a11-b05f-793fcb57e093', '660c8c08-468b-4a0d-a37d-fe58d7d0b075', NULL);
INSERT INTO public.likes VALUES ('15127c31-ed44-4324-bb37-bed268b0b73b', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'e2d16c26-ba68-4a8d-ab33-8a68890f1683', NULL);
INSERT INTO public.likes VALUES ('d433d01d-fb2a-4108-bdfd-19812d70c0e3', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'e2d16c26-ba68-4a8d-ab33-8a68890f1683', NULL);
INSERT INTO public.likes VALUES ('8cbcd5b7-d180-4395-8a84-b5dd2f78d406', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', 'e2d16c26-ba68-4a8d-ab33-8a68890f1683', NULL);
INSERT INTO public.likes VALUES ('b4efc41f-90d1-4a87-8990-a36aba7cad51', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', 'a3624d47-236a-4b55-be16-d717bfcd8384', NULL);
INSERT INTO public.likes VALUES ('5928c7b7-f8eb-40cf-85af-830406832d28', '79758f94-70d8-4a11-b05f-793fcb57e093', 'a3624d47-236a-4b55-be16-d717bfcd8384', NULL);
INSERT INTO public.likes VALUES ('2cc8b085-df1a-401d-ba79-0af5985f7abc', '6e137217-71a9-439b-a579-4bcd1a729af9', 'a3624d47-236a-4b55-be16-d717bfcd8384', NULL);
INSERT INTO public.likes VALUES ('de3f677c-aed4-4a4e-89fb-8d1a7f1ac655', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'a3624d47-236a-4b55-be16-d717bfcd8384', NULL);
INSERT INTO public.likes VALUES ('11253c2d-f58c-4be2-a411-1dd80790b544', 'f399a996-1567-4fc0-a86a-92774494ca44', '384be8c2-daf3-4bb6-8e8a-69d83285c0df', NULL);
INSERT INTO public.likes VALUES ('405c4db9-4a1e-40d4-96fc-9f9987f48f9b', '3b0ebec9-3805-490b-987f-4c6e7286319e', '384be8c2-daf3-4bb6-8e8a-69d83285c0df', NULL);
INSERT INTO public.likes VALUES ('942f9679-97de-4390-997e-23559b1366ae', '56e89866-3a90-4f97-8461-34b24f842da4', '384be8c2-daf3-4bb6-8e8a-69d83285c0df', NULL);
INSERT INTO public.likes VALUES ('eba38d3f-e677-420d-aa5c-a764882d6e4e', '79758f94-70d8-4a11-b05f-793fcb57e093', '384be8c2-daf3-4bb6-8e8a-69d83285c0df', NULL);
INSERT INTO public.likes VALUES ('52e760a2-bb95-4522-891c-d10bb127a72a', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '384be8c2-daf3-4bb6-8e8a-69d83285c0df', NULL);
INSERT INTO public.likes VALUES ('1432eb42-00d3-4de3-9d5a-d039ae67161c', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '384be8c2-daf3-4bb6-8e8a-69d83285c0df', NULL);
INSERT INTO public.likes VALUES ('7df9ef1f-628f-4cb5-967b-e8ae56568ec4', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', 'd85d141a-e806-45e6-83bd-3b48cef86d95', NULL);
INSERT INTO public.likes VALUES ('8fbe344a-b8fa-4b73-b3c4-3836353cf966', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'd85d141a-e806-45e6-83bd-3b48cef86d95', NULL);
INSERT INTO public.likes VALUES ('bfcfa091-19ff-4099-ab9a-d91a523bd234', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', 'd85d141a-e806-45e6-83bd-3b48cef86d95', NULL);
INSERT INTO public.likes VALUES ('5d67a21c-454a-4267-94a5-cbe99ed4bc2c', 'f399a996-1567-4fc0-a86a-92774494ca44', 'd85d141a-e806-45e6-83bd-3b48cef86d95', NULL);
INSERT INTO public.likes VALUES ('776b8551-5360-4bd2-85ac-83dd94925f5d', '79758f94-70d8-4a11-b05f-793fcb57e093', 'd85d141a-e806-45e6-83bd-3b48cef86d95', NULL);
INSERT INTO public.likes VALUES ('ec4ab314-67e6-45a3-b57c-902b13ed2dd4', '6e137217-71a9-439b-a579-4bcd1a729af9', 'd85d141a-e806-45e6-83bd-3b48cef86d95', NULL);
INSERT INTO public.likes VALUES ('bbc8fdf1-7f68-4c80-a23e-fa3c5a8de457', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', 'd85d141a-e806-45e6-83bd-3b48cef86d95', NULL);
INSERT INTO public.likes VALUES ('85fd5c06-215b-483a-87ed-a4baa148a31c', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'ab773301-3e88-4c8b-9b75-34d5654e2a2a', NULL);
INSERT INTO public.likes VALUES ('509ccfdf-a7e6-4011-a9cd-f44439616b73', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', 'ab773301-3e88-4c8b-9b75-34d5654e2a2a', NULL);
INSERT INTO public.likes VALUES ('592aaa32-6b59-45c6-a5fb-d03266ff6d27', '6e137217-71a9-439b-a579-4bcd1a729af9', 'ab773301-3e88-4c8b-9b75-34d5654e2a2a', NULL);
INSERT INTO public.likes VALUES ('63fb2c3f-d87d-4aa5-b3a6-dfc6e5559712', 'f399a996-1567-4fc0-a86a-92774494ca44', 'ab773301-3e88-4c8b-9b75-34d5654e2a2a', NULL);
INSERT INTO public.likes VALUES ('0f58ac94-22a6-4344-be4a-c81c4f5e8e7b', '79758f94-70d8-4a11-b05f-793fcb57e093', 'ab773301-3e88-4c8b-9b75-34d5654e2a2a', NULL);
INSERT INTO public.likes VALUES ('ebcedf5e-85ec-43c2-8948-e6ec9e8897b3', '56e89866-3a90-4f97-8461-34b24f842da4', 'ab773301-3e88-4c8b-9b75-34d5654e2a2a', NULL);
INSERT INTO public.likes VALUES ('167695f4-9760-4a07-a70b-8231534b4e2a', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', 'ab773301-3e88-4c8b-9b75-34d5654e2a2a', NULL);
INSERT INTO public.likes VALUES ('152b1c95-9aa8-46c6-9c40-b9d2dfbf53f4', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'ab773301-3e88-4c8b-9b75-34d5654e2a2a', NULL);
INSERT INTO public.likes VALUES ('9f2e76a8-e6f0-4ba7-a9e9-b931ec1bb286', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', 'aa09c190-fd10-4554-966e-835aefbafcc5', NULL);
INSERT INTO public.likes VALUES ('381c4c43-0d18-4e24-b086-b59461f08153', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', 'aa09c190-fd10-4554-966e-835aefbafcc5', NULL);
INSERT INTO public.likes VALUES ('ec3873ae-f723-4e44-8bc4-d7218e3c3174', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'aa09c190-fd10-4554-966e-835aefbafcc5', NULL);
INSERT INTO public.likes VALUES ('6556b4b1-d3a7-4935-bff3-516bde78d6ff', '79758f94-70d8-4a11-b05f-793fcb57e093', 'aa09c190-fd10-4554-966e-835aefbafcc5', NULL);
INSERT INTO public.likes VALUES ('3f37174f-a827-4a8e-b85c-a881b58067d3', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', 'aa09c190-fd10-4554-966e-835aefbafcc5', NULL);
INSERT INTO public.likes VALUES ('d5b024c0-b0f4-4fa2-9fb8-6a96b671b7ad', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'aa09c190-fd10-4554-966e-835aefbafcc5', NULL);
INSERT INTO public.likes VALUES ('fdbd1c9b-e5b3-4a1d-97fd-04eb73b7d12b', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'aa09c190-fd10-4554-966e-835aefbafcc5', NULL);
INSERT INTO public.likes VALUES ('c9dca898-2de3-40be-aa84-d1f596af37a2', '6e137217-71a9-439b-a579-4bcd1a729af9', 'aa09c190-fd10-4554-966e-835aefbafcc5', NULL);
INSERT INTO public.likes VALUES ('93240e1b-27d5-48a8-9091-9d663bb80c74', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', 'cfd2ccd0-49d2-4a73-8f13-14348a779f6f', NULL);
INSERT INTO public.likes VALUES ('63a7802c-2a5a-4926-a107-9049976742e6', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'cfd2ccd0-49d2-4a73-8f13-14348a779f6f', NULL);
INSERT INTO public.likes VALUES ('8228c602-6ae0-47d5-addf-7b87c557e693', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'cfd2ccd0-49d2-4a73-8f13-14348a779f6f', NULL);
INSERT INTO public.likes VALUES ('d28f2b43-b672-43cc-b193-eba1dcf9dc01', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', 'cfd2ccd0-49d2-4a73-8f13-14348a779f6f', NULL);
INSERT INTO public.likes VALUES ('fa11ac76-42cb-4cad-9e23-6991f49a3eaa', 'f399a996-1567-4fc0-a86a-92774494ca44', 'cfd2ccd0-49d2-4a73-8f13-14348a779f6f', NULL);
INSERT INTO public.likes VALUES ('bd58bd4a-cf9d-4fff-a13d-eaf7cfefbc36', '56e89866-3a90-4f97-8461-34b24f842da4', 'cfd2ccd0-49d2-4a73-8f13-14348a779f6f', NULL);
INSERT INTO public.likes VALUES ('eacd776b-4a78-4d6f-baa0-0e60fcd90c7c', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'cfd2ccd0-49d2-4a73-8f13-14348a779f6f', NULL);
INSERT INTO public.likes VALUES ('1339c498-36a2-4229-9f80-4bb44c92fd11', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'cfd2ccd0-49d2-4a73-8f13-14348a779f6f', NULL);


--
-- Data for Name: microtasks; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.microtasks VALUES ('be99dca9-40d3-475d-8ca1-a0f599df7795', '9dd80410-bd88-455a-b8fe-a5c2e39edcf0', 'Watch video tutorial', 1, 'in progress', 0, '2025-06-07 15:52:06.391299+00', '2025-06-07 15:52:06.391299+00', NULL);
INSERT INTO public.microtasks VALUES ('62fd5b6a-1833-4551-b403-82eb0dab6724', '9dd80410-bd88-455a-b8fe-a5c2e39edcf0', 'Buy ingredients', 2, 'in progress', 1, '2025-06-07 15:52:06.391301+00', '2025-06-07 15:52:06.391301+00', NULL);
INSERT INTO public.microtasks VALUES ('9b3da3d5-8874-4398-a1ac-07e835ce813b', '9dd80410-bd88-455a-b8fe-a5c2e39edcf0', 'Cook samsa', 1, 'in progress', 2, '2025-06-07 15:52:06.391302+00', '2025-06-07 15:52:06.391302+00', NULL);
INSERT INTO public.microtasks VALUES ('b377c3ec-dfcf-48b0-9ece-f3898d31ae9c', '9edaf28d-e47c-4bec-aa12-f4d6b78c6edb', 'Read Atomic Habits', 3, 'in progress', 0, '2025-06-07 15:52:21.291749+00', '2025-06-07 15:52:21.291749+00', NULL);
INSERT INTO public.microtasks VALUES ('78d799bd-4a56-4538-bd57-60deda16ded5', '9edaf28d-e47c-4bec-aa12-f4d6b78c6edb', 'Read Deep Work', 2, 'in progress', 1, '2025-06-07 15:52:21.29175+00', '2025-06-07 15:52:21.29175+00', NULL);
INSERT INTO public.microtasks VALUES ('dfc33a5b-552c-4d12-b43f-abad5ee9a14d', '9edaf28d-e47c-4bec-aa12-f4d6b78c6edb', 'Summarize and apply', 2, 'in progress', 2, '2025-06-07 15:52:21.291752+00', '2025-06-07 15:52:21.291752+00', NULL);
INSERT INTO public.microtasks VALUES ('239bd055-307a-4344-97a7-211b820620f0', '609eba42-64bc-460a-9468-bc011dae591c', 'Set up meditation app', 1, 'in progress', 0, '2025-06-07 15:57:22.966273+00', '2025-06-07 15:57:22.966273+00', NULL);
INSERT INTO public.microtasks VALUES ('941639f8-f2f9-4b34-b71c-374442355a50', '609eba42-64bc-460a-9468-bc011dae591c', 'Morning meditation routine', 5, 'in progress', 1, '2025-06-07 15:57:22.966274+00', '2025-06-07 15:57:22.966274+00', NULL);
INSERT INTO public.microtasks VALUES ('9b6fcfac-2b07-47ec-be05-6948ebe66541', '8ecb2652-80cf-4fe2-bab8-f3841d663a2d', 'Wake up at 6 AM', 2, 'in progress', 0, '2025-06-07 16:44:05.365203+00', '2025-06-07 16:44:05.365203+00', NULL);
INSERT INTO public.microtasks VALUES ('36590d41-6f25-4b2d-9df5-b0def4bef879', '8ecb2652-80cf-4fe2-bab8-f3841d663a2d', 'Do 10-minute meditation', 1, 'in progress', 1, '2025-06-07 16:44:05.365205+00', '2025-06-07 16:44:05.365205+00', NULL);
INSERT INTO public.microtasks VALUES ('1eabc31c-8f77-4779-804a-0df1b535b12b', '8ecb2652-80cf-4fe2-bab8-f3841d663a2d', 'Stretch or short walk', 1, 'in progress', 2, '2025-06-07 16:44:05.365207+00', '2025-06-07 16:44:05.365207+00', NULL);
INSERT INTO public.microtasks VALUES ('d89edf48-a9de-456a-aa44-9e669a382fd5', '66b19273-144f-4dd0-971e-830c18c26b38', 'Finish 3 Easy problems', 3, 'in progress', 0, '2025-06-07 16:44:14.301359+00', '2025-06-07 16:44:14.301359+00', NULL);
INSERT INTO public.microtasks VALUES ('adb2d568-a751-4303-a5cd-67362c864091', '66b19273-144f-4dd0-971e-830c18c26b38', 'Solve 5 Medium problems', 5, 'in progress', 1, '2025-06-07 16:44:14.30136+00', '2025-06-07 16:44:14.301361+00', NULL);
INSERT INTO public.microtasks VALUES ('4d8c1e09-2684-49bf-904a-c8bfd973feae', '66b19273-144f-4dd0-971e-830c18c26b38', 'Attempt 2 Hard problems', 2, 'in progress', 2, '2025-06-07 16:44:14.301362+00', '2025-06-07 16:44:14.301362+00', NULL);
INSERT INTO public.microtasks VALUES ('2c31cf98-f564-4537-8a12-a71a705103ea', 'e5bf481a-040b-4b1b-96c7-787909876292', 'Read 1 Juz every day', 30, 'in progress', 0, '2025-06-07 16:44:38.050593+00', '2025-06-07 16:44:38.050593+00', NULL);
INSERT INTO public.microtasks VALUES ('ce0aa660-fdc1-4342-af7d-c55f54b1904e', 'e5bf481a-040b-4b1b-96c7-787909876292', 'Take notes of key reflections', 5, 'in progress', 1, '2025-06-07 16:44:38.050595+00', '2025-06-07 16:44:38.050595+00', NULL);
INSERT INTO public.microtasks VALUES ('ff2ec17a-020e-44c4-b3bc-3dc64e88ed45', '912c9992-debd-42b8-9df7-28a9c4b7a319', 'Buy a journal or app', 1, 'in progress', 0, '2025-06-07 16:55:21.47355+00', '2025-06-07 16:55:21.47355+00', NULL);
INSERT INTO public.microtasks VALUES ('d57937da-618e-4c6b-bc5b-e757af98582e', '912c9992-debd-42b8-9df7-28a9c4b7a319', 'Write entries for 7 consecutive days', 7, 'in progress', 1, '2025-06-07 16:55:21.473551+00', '2025-06-07 16:55:21.473551+00', NULL);
INSERT INTO public.microtasks VALUES ('c72f0f00-94be-41ae-ade0-954267d26c9d', '6b7011e6-fb89-4542-8eba-31b219970185', 'Run 1km three times', 3, 'in progress', 0, '2025-06-07 16:55:33.701382+00', '2025-06-07 16:55:33.701382+00', NULL);
INSERT INTO public.microtasks VALUES ('5501f6a8-a3be-46fe-a00e-46ef96f0dedd', '6b7011e6-fb89-4542-8eba-31b219970185', 'Run 3km two times', 2, 'in progress', 1, '2025-06-07 16:55:33.701383+00', '2025-06-07 16:55:33.701383+00', NULL);
INSERT INTO public.microtasks VALUES ('7214781e-ee0e-444d-9a4c-a6ec49f8b9be', '6b7011e6-fb89-4542-8eba-31b219970185', 'Run full 5km once', 1, 'in progress', 2, '2025-06-07 16:55:33.701385+00', '2025-06-07 16:55:33.701385+00', NULL);
INSERT INTO public.microtasks VALUES ('429eac06-eca0-49fe-9cc0-6c731b297206', 'c84f3282-901c-4433-82ee-42fc3f4ddf3a', 'Learn greetings and introductions', 2, 'in progress', 0, '2025-06-07 16:55:44.029716+00', '2025-06-07 16:55:44.029716+00', NULL);
INSERT INTO public.microtasks VALUES ('20e7d6e0-aaad-4946-84a9-7be9471fb444', 'c84f3282-901c-4433-82ee-42fc3f4ddf3a', 'Practice common travel phrases', 3, 'in progress', 1, '2025-06-07 16:55:44.029717+00', '2025-06-07 16:55:44.029717+00', NULL);
INSERT INTO public.microtasks VALUES ('0e96f9b2-7005-4c25-9d0a-14bd722bebcb', 'c84f3282-901c-4433-82ee-42fc3f4ddf3a', 'Watch German videos with subtitles', 2, 'in progress', 2, '2025-06-07 16:55:44.029719+00', '2025-06-07 16:55:44.029719+00', NULL);
INSERT INTO public.microtasks VALUES ('4db717b6-a224-4748-accf-17c23d4a7015', '47b0bc29-2ef0-4b57-a3b9-849968031526', 'Buy a domain', 1, 'in progress', 0, '2025-06-07 18:36:06.325129+00', '2025-06-07 18:36:06.325129+00', NULL);
INSERT INTO public.microtasks VALUES ('cab6f8f2-e842-45ee-8b47-75ee8c6a8c73', '47b0bc29-2ef0-4b57-a3b9-849968031526', 'Design homepage and about me', 2, 'in progress', 1, '2025-06-07 18:36:06.32513+00', '2025-06-07 18:36:06.32513+00', NULL);
INSERT INTO public.microtasks VALUES ('ae132a3a-a947-45a1-949b-cf40ff3b4560', '47b0bc29-2ef0-4b57-a3b9-849968031526', 'Deploy site to hosting', 1, 'in progress', 2, '2025-06-07 18:36:06.325132+00', '2025-06-07 18:36:06.325132+00', NULL);
INSERT INTO public.microtasks VALUES ('7e1d9562-3ffd-40be-830d-0fd0e005beb7', 'c3f3d0c5-29e5-4167-87ab-8178d06ac7c2', 'Do cardio 3 times a week', 3, 'in progress', 0, '2025-06-07 18:36:23.143519+00', '2025-06-07 18:36:23.143519+00', NULL);
INSERT INTO public.microtasks VALUES ('47bf77f8-3f6d-48f0-bdae-27574c0b902e', 'c3f3d0c5-29e5-4167-87ab-8178d06ac7c2', 'Strength training twice a week', 2, 'in progress', 1, '2025-06-07 18:36:23.143521+00', '2025-06-07 18:36:23.143521+00', NULL);
INSERT INTO public.microtasks VALUES ('da1ef535-f977-4327-a352-fb097383d231', '111e36b5-3a31-4079-a5c8-8ead7454723b', 'Track time spent on tasks', 2, 'in progress', 0, '2025-06-07 18:36:43.706858+00', '2025-06-07 18:36:43.706858+00', NULL);
INSERT INTO public.microtasks VALUES ('0ea22b31-fb48-49de-9299-446fdf46623d', '111e36b5-3a31-4079-a5c8-8ead7454723b', 'Set weekly goals and reflect', 2, 'in progress', 1, '2025-06-07 18:36:43.70686+00', '2025-06-07 18:36:43.70686+00', NULL);
INSERT INTO public.microtasks VALUES ('d36ac6f3-57a2-4d11-ba34-9674e46c7bec', 'c4c10327-612e-4a10-a744-853e80c41a49', 'Install drawing software', 1, 'in progress', 0, '2025-06-07 19:15:45.857299+00', '2025-06-07 19:15:45.857299+00', NULL);
INSERT INTO public.microtasks VALUES ('106936c7-f18f-48c6-81ab-6f931815cced', 'c4c10327-612e-4a10-a744-853e80c41a49', 'Follow a basic tutorial', 2, 'in progress', 1, '2025-06-07 19:15:45.857301+00', '2025-06-07 19:15:45.857301+00', NULL);
INSERT INTO public.microtasks VALUES ('c6dd0c27-420c-4123-81c0-0744d6e3a7a6', 'c4c10327-612e-4a10-a744-853e80c41a49', 'Draw 3 simple characters', 3, 'in progress', 2, '2025-06-07 19:15:45.857303+00', '2025-06-07 19:15:45.857303+00', NULL);
INSERT INTO public.microtasks VALUES ('6eea2e14-0992-4f13-a5e8-d927886af1e1', 'ac557757-04c4-428d-a616-655ecda8b969', 'Group words by topic', 1, 'in progress', 0, '2025-06-07 19:15:57.014698+00', '2025-06-07 19:15:57.014698+00', NULL);
INSERT INTO public.microtasks VALUES ('adbd6df3-b447-4970-8d81-bb54b15310b5', 'ac557757-04c4-428d-a616-655ecda8b969', 'Use flashcards daily', 5, 'in progress', 1, '2025-06-07 19:15:57.0147+00', '2025-06-07 19:15:57.0147+00', NULL);
INSERT INTO public.microtasks VALUES ('b03ade4c-2749-4068-b191-d2d31129ec5c', 'ac557757-04c4-428d-a616-655ecda8b969', 'Write sentences with new words', 2, 'in progress', 2, '2025-06-07 19:15:57.014701+00', '2025-06-07 19:15:57.014701+00', NULL);
INSERT INTO public.microtasks VALUES ('771b2a91-478a-46de-afc3-b87f15a693fd', '30a43be9-a3fe-43c9-b536-25cb0ee96c1b', 'Throw away unused items', 2, 'in progress', 0, '2025-06-07 19:16:09.648618+00', '2025-06-07 19:16:09.648618+00', NULL);
INSERT INTO public.microtasks VALUES ('fe218910-6c0b-423c-bbf6-468e5b86844c', '30a43be9-a3fe-43c9-b536-25cb0ee96c1b', 'Rearrange desk layout', 1, 'in progress', 1, '2025-06-07 19:16:09.648619+00', '2025-06-07 19:16:09.648619+00', NULL);
INSERT INTO public.microtasks VALUES ('ba72e44e-efb7-4331-a79f-a36933bc4f35', '70679b66-1611-48c0-9bcb-0558b6d28965', 'Finish A2 grammar book', 5, 'in progress', 0, '2025-06-07 19:31:01.900587+00', '2025-06-07 19:31:01.900587+00', NULL);
INSERT INTO public.microtasks VALUES ('7dbec861-f7dd-47dc-9695-c3fdd5f60791', '70679b66-1611-48c0-9bcb-0558b6d28965', 'Take 3 mock exams', 3, 'in progress', 1, '2025-06-07 19:31:01.90059+00', '2025-06-07 19:31:01.90059+00', NULL);
INSERT INTO public.microtasks VALUES ('12cccfe8-320f-4992-9d2f-b59ef5ea4f6e', '70679b66-1611-48c0-9bcb-0558b6d28965', 'Practice speaking weekly', 4, 'in progress', 2, '2025-06-07 19:31:01.900593+00', '2025-06-07 19:31:01.900593+00', NULL);
INSERT INTO public.microtasks VALUES ('4d943fed-81fb-4660-a2f5-c64779b63a7e', '7e8e87bc-1b4c-44e9-832c-6e29617642d8', 'Do posture stretches every morning', 5, 'in progress', 0, '2025-06-07 19:31:13.329994+00', '2025-06-07 19:31:13.329994+00', NULL);
INSERT INTO public.microtasks VALUES ('4369d19b-7570-47b9-9f06-f72b2eee78f2', '7e8e87bc-1b4c-44e9-832c-6e29617642d8', 'Sit with lumbar support', 1, 'in progress', 1, '2025-06-07 19:31:13.329995+00', '2025-06-07 19:31:13.329995+00', NULL);
INSERT INTO public.microtasks VALUES ('75a93f4b-253b-414d-a112-85ada92226c2', '7e8e87bc-1b4c-44e9-832c-6e29617642d8', 'Take walking breaks during work', 3, 'in progress', 2, '2025-06-07 19:31:13.329997+00', '2025-06-07 19:31:13.329997+00', NULL);
INSERT INTO public.microtasks VALUES ('fe5cb029-80a1-4c37-8fd0-62af0afc0976', 'a9528f6b-3895-4cab-af6c-62a93ef1b83e', 'Design the layout in Figma', 1, 'in progress', 0, '2025-06-07 19:31:54.123191+00', '2025-06-07 19:31:54.123191+00', NULL);
INSERT INTO public.microtasks VALUES ('d81f001e-91a4-4a34-aeb8-90544b4c328e', 'a9528f6b-3895-4cab-af6c-62a93ef1b83e', 'Build with React and Tailwind', 3, 'in progress', 1, '2025-06-07 19:31:54.123192+00', '2025-06-07 19:31:54.123192+00', NULL);
INSERT INTO public.microtasks VALUES ('0fb658ac-876a-448a-ada4-e96dd4c3e95b', 'a9528f6b-3895-4cab-af6c-62a93ef1b83e', 'Deploy to GitHub Pages', 1, 'in progress', 2, '2025-06-07 19:31:54.123194+00', '2025-06-07 19:31:54.123194+00', NULL);
INSERT INTO public.microtasks VALUES ('9aead199-68d1-4dc2-b42e-af82e4932380', 'b438fcc1-9526-4e3c-8bd2-039bc7e71e7b', 'Outline blog structure', 1, 'in progress', 0, '2025-06-07 19:54:56.347179+00', '2025-06-07 19:54:56.347179+00', NULL);
INSERT INTO public.microtasks VALUES ('a40da453-3b25-4f00-97c3-04c8ac10f93b', 'b438fcc1-9526-4e3c-8bd2-039bc7e71e7b', 'Write first draft', 2, 'in progress', 1, '2025-06-07 19:54:56.347181+00', '2025-06-07 19:54:56.347181+00', NULL);
INSERT INTO public.microtasks VALUES ('556ed912-4d18-49d3-80c4-521abefbeced', 'b438fcc1-9526-4e3c-8bd2-039bc7e71e7b', 'Revise and publish', 1, 'in progress', 2, '2025-06-07 19:54:56.347182+00', '2025-06-07 19:54:56.347182+00', NULL);
INSERT INTO public.microtasks VALUES ('97ff5a74-3b7d-40ce-ab6b-7153f65a446a', '513f748c-0dc9-4b04-a8f0-f33ea5892151', 'Design database schema', 1, 'in progress', 0, '2025-06-07 19:55:01.280566+00', '2025-06-07 19:55:01.280566+00', NULL);
INSERT INTO public.microtasks VALUES ('4ef71a7d-5c7a-47b4-bead-b825ef583b04', '513f748c-0dc9-4b04-a8f0-f33ea5892151', 'Implement CRUD APIs', 3, 'in progress', 1, '2025-06-07 19:55:01.280567+00', '2025-06-07 19:55:01.280567+00', NULL);
INSERT INTO public.microtasks VALUES ('1729e316-eb3d-485e-bd1d-0ad114b0f79c', '513f748c-0dc9-4b04-a8f0-f33ea5892151', 'Deploy to Render or Fly.io', 1, 'in progress', 2, '2025-06-07 19:55:01.280568+00', '2025-06-07 19:55:01.280568+00', NULL);
INSERT INTO public.microtasks VALUES ('950b125d-11c7-4703-877c-a87c1cb027b0', 'f3984ee6-5f7b-47b0-a88c-c684e10eae43', 'Wake up by 6:30 AM', 7, 'in progress', 0, '2025-06-07 19:55:08.848614+00', '2025-06-07 19:55:08.848614+00', NULL);
INSERT INTO public.microtasks VALUES ('141ea941-bf50-45ec-8387-deefb7e716ed', 'f3984ee6-5f7b-47b0-a88c-c684e10eae43', 'Do 10 minutes of stretching', 5, 'in progress', 1, '2025-06-07 19:55:08.848616+00', '2025-06-07 19:55:08.848616+00', NULL);
INSERT INTO public.microtasks VALUES ('15f6b3ad-4408-457a-879a-cc080aeb3746', 'f3984ee6-5f7b-47b0-a88c-c684e10eae43', 'No phone for 1 hour after waking', 7, 'in progress', 2, '2025-06-07 19:55:08.848617+00', '2025-06-07 19:55:08.848617+00', NULL);


--
-- Data for Name: notifications; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: posts; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.posts VALUES ('866d6d95-29d3-41e2-adb7-a01b09fd5616', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '9dd80410-bd88-455a-b8fe-a5c2e39edcf0', 'be99dca9-40d3-475d-8ca1-a0f599df7795', NULL, 'Just watched a detailed Uzbek chef explain how to make traditional samsa with lamb. Looks delicious and easier than I thought!', '2025-06-07 18:40:29.136373+00', '2025-06-07 18:40:29.136373+00', NULL);
INSERT INTO public.posts VALUES ('23ad69f9-8c1f-45e1-bdb6-86ae97f0a2d2', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '9dd80410-bd88-455a-b8fe-a5c2e39edcf0', '62fd5b6a-1833-4551-b403-82eb0dab6724', NULL, 'Went to the market and grabbed everything I need: flour, onions, meat, cumin. Almost forgot the sesame seeds!', '2025-06-07 19:10:03.559862+00', '2025-06-07 19:10:03.559862+00', NULL);
INSERT INTO public.posts VALUES ('4fcdd3cc-c626-4dd4-904b-426f0839c71c', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '9dd80410-bd88-455a-b8fe-a5c2e39edcf0', '62fd5b6a-1833-4551-b403-82eb0dab6724', NULL, 'Shopping was chaotic, but I found all the ingredients in the end. Lamb was surprisingly hard to find!', '2025-06-07 19:38:19.693846+00', '2025-06-07 19:38:19.693846+00', NULL);
INSERT INTO public.posts VALUES ('26e50906-4a53-4537-89c5-9971e86bb0ad', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '9dd80410-bd88-455a-b8fe-a5c2e39edcf0', '9b3da3d5-8874-4398-a1ac-07e835ce813b', NULL, 'Made samsa today for the first time. The crust is crispy, filling is juicy тАФ family loved it!', '2025-06-07 16:49:02.639102+00', '2025-06-07 16:49:02.639102+00', NULL);
INSERT INTO public.posts VALUES ('19e2449f-0228-4c1a-b7c7-c7c8c0d7a4a4', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '609eba42-64bc-460a-9468-bc011dae591c', '239bd055-307a-4344-97a7-211b820620f0', NULL, 'Installed Headspace and set my first daily reminder. Time to finally take meditation seriously.', '2025-06-07 19:20:10.456679+00', '2025-06-07 19:20:10.456679+00', NULL);
INSERT INTO public.posts VALUES ('741d9f45-74da-4da2-b40e-06b325c5342f', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '609eba42-64bc-460a-9468-bc011dae591c', '941639f8-f2f9-4b34-b71c-374442355a50', NULL, 'Started my morning with 10 minutes of silence. My thoughts finally slowed down. Surprised how hard it is to focus!', '2025-06-07 19:37:38.589553+00', '2025-06-07 19:37:38.589553+00', NULL);
INSERT INTO public.posts VALUES ('678fdbee-e90c-445c-949e-e3962b0503ce', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '609eba42-64bc-460a-9468-bc011dae591c', '941639f8-f2f9-4b34-b71c-374442355a50', NULL, 'Second day of meditation. Added soft ambient music тАФ helped me stay grounded.', '2025-06-07 18:32:00.525467+00', '2025-06-07 18:32:00.525467+00', NULL);
INSERT INTO public.posts VALUES ('8249f900-b093-4d7a-b8f8-67be7fb47448', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '609eba42-64bc-460a-9468-bc011dae591c', '941639f8-f2f9-4b34-b71c-374442355a50', NULL, 'Missed the alarm, but still managed to meditate for 5 minutes. Small wins matter.', '2025-06-07 19:59:21.610526+00', '2025-06-07 19:59:21.610526+00', NULL);
INSERT INTO public.posts VALUES ('53626074-a976-43cb-83b1-b812a80f3244', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '9edaf28d-e47c-4bec-aa12-f4d6b78c6edb', 'b377c3ec-dfcf-48b0-9ece-f3898d31ae9c', NULL, 'Started reading *Atomic Habits*. The first chapter already made me rethink how small habits shape identity.', '2025-06-07 19:11:55.864884+00', '2025-06-07 19:11:55.864884+00', NULL);
INSERT INTO public.posts VALUES ('1402f5b0-c845-4665-998c-e5597f5f5521', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '9edaf28d-e47c-4bec-aa12-f4d6b78c6edb', 'b377c3ec-dfcf-48b0-9ece-f3898d31ae9c', NULL, 'Noted key ideas about habit stacking. Planning to implement a morning reading cue.', '2025-06-07 19:23:48.43565+00', '2025-06-07 19:23:48.43565+00', NULL);
INSERT INTO public.posts VALUES ('60333cb2-3eac-4c88-b1ab-187788f3a753', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '9edaf28d-e47c-4bec-aa12-f4d6b78c6edb', '78d799bd-4a56-4538-bd57-60deda16ded5', NULL, 'Just started Deep Work by Cal Newport. Realized how much shallow distractions eat my day.', '2025-06-07 16:15:59.127193+00', '2025-06-07 16:15:59.127193+00', NULL);
INSERT INTO public.posts VALUES ('cdcaa739-1c4d-422c-abea-b234c55cc559', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '9edaf28d-e47c-4bec-aa12-f4d6b78c6edb', 'dfc33a5b-552c-4d12-b43f-abad5ee9a14d', NULL, 'Wrote summaries for both books. Planning how to integrate these ideas into my workflow.', '2025-06-07 19:21:23.645203+00', '2025-06-07 19:21:23.645203+00', NULL);
INSERT INTO public.posts VALUES ('23a24eb3-9a76-4807-8ff2-24ebdd6f11bb', '3b0ebec9-3805-490b-987f-4c6e7286319e', '66b19273-144f-4dd0-971e-830c18c26b38', 'd89edf48-a9de-456a-aa44-9e669a382fd5', NULL, 'Started solving my first easy LeetCode problem today. Felt good!', '2025-06-07 19:58:56.195766+00', '2025-06-07 19:58:56.195766+00', NULL);
INSERT INTO public.posts VALUES ('b8ea0106-08f2-45fc-92ea-77a61af34a1d', '3b0ebec9-3805-490b-987f-4c6e7286319e', '66b19273-144f-4dd0-971e-830c18c26b38', 'd89edf48-a9de-456a-aa44-9e669a382fd5', NULL, 'Completed two more easy problems. Progress is visible!', '2025-06-07 19:39:18.582408+00', '2025-06-07 19:39:18.582408+00', NULL);
INSERT INTO public.posts VALUES ('31f82a0d-869c-4ea9-8c72-152d229893cf', '3b0ebec9-3805-490b-987f-4c6e7286319e', '66b19273-144f-4dd0-971e-830c18c26b38', 'adb2d568-a751-4303-a5cd-67362c864091', NULL, 'Medium problems are a bit tricky, but I''m learning a lot.', '2025-06-07 19:59:11.392577+00', '2025-06-07 19:59:11.392577+00', NULL);
INSERT INTO public.posts VALUES ('91d813fb-0bb2-4a9f-98cd-40c4c86de628', '3b0ebec9-3805-490b-987f-4c6e7286319e', '66b19273-144f-4dd0-971e-830c18c26b38', '4d8c1e09-2684-49bf-904a-c8bfd973feae', NULL, 'Attempted my first hard problem today. Got stuck but learned heaps.', '2025-06-07 19:58:51.57268+00', '2025-06-07 19:58:51.57268+00', NULL);
INSERT INTO public.posts VALUES ('e5ac024b-a6ee-469f-834f-704f37620b25', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'e5bf481a-040b-4b1b-96c7-787909876292', '2c31cf98-f564-4537-8a12-a71a705103ea', NULL, 'Day 1: Finished Juz 1. Beautiful reminders throughout.', '2025-06-07 19:22:43.030841+00', '2025-06-07 19:22:43.030841+00', NULL);
INSERT INTO public.posts VALUES ('aa09c190-fd10-4554-966e-835aefbafcc5', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'e5bf481a-040b-4b1b-96c7-787909876292', '2c31cf98-f564-4537-8a12-a71a705103ea', NULL, 'Day 2: Read Juz 2 after Fajr. Feeling refreshed spiritually.', '2025-06-07 19:19:00.389621+00', '2025-06-07 19:19:00.389621+00', NULL);
INSERT INTO public.posts VALUES ('4ddee261-c8d1-4980-9a3e-0baef66bbf7c', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'e5bf481a-040b-4b1b-96c7-787909876292', 'ce0aa660-fdc1-4342-af7d-c55f54b1904e', NULL, 'Wrote reflections on Surah Baqarah today. Deep meanings.', '2025-06-07 18:31:21.341003+00', '2025-06-07 18:31:21.341003+00', NULL);
INSERT INTO public.posts VALUES ('0abf63f3-6d91-4737-add9-edb192813b76', '56e89866-3a90-4f97-8461-34b24f842da4', '6b7011e6-fb89-4542-8eba-31b219970185', 'c72f0f00-94be-41ae-ade0-954267d26c9d', NULL, 'Started training today by running 1km. Felt harder than expected, but I made it.', '2025-06-07 16:49:23.923606+00', '2025-06-07 16:49:23.923606+00', NULL);
INSERT INTO public.posts VALUES ('97000960-eabc-450b-b074-609480dc17e8', '56e89866-3a90-4f97-8461-34b24f842da4', '6b7011e6-fb89-4542-8eba-31b219970185', 'c72f0f00-94be-41ae-ade0-954267d26c9d', NULL, 'Third 1km run done! My pace is improving each day.', '2025-06-07 19:37:10.005337+00', '2025-06-07 19:37:10.005337+00', NULL);
INSERT INTO public.posts VALUES ('d85d141a-e806-45e6-83bd-3b48cef86d95', '56e89866-3a90-4f97-8461-34b24f842da4', '6b7011e6-fb89-4542-8eba-31b219970185', '5501f6a8-a3be-46fe-a00e-46ef96f0dedd', NULL, 'Completed 3km today! Legs sore but happy with progress.', '2025-06-07 19:22:58.948263+00', '2025-06-07 19:22:58.948263+00', NULL);
INSERT INTO public.posts VALUES ('223da9a0-82b1-4745-9a4a-5b657d47299e', '56e89866-3a90-4f97-8461-34b24f842da4', '6b7011e6-fb89-4542-8eba-31b219970185', '7214781e-ee0e-444d-9a4c-a6ec49f8b9be', NULL, 'Ran full 5km for the first time. Incredible feeling!', '2025-06-07 16:21:50.550104+00', '2025-06-07 16:21:50.550104+00', NULL);
INSERT INTO public.posts VALUES ('1e976362-73b0-400a-a74c-3fa951afdbeb', '56e89866-3a90-4f97-8461-34b24f842da4', '912c9992-debd-42b8-9df7-28a9c4b7a319', 'ff2ec17a-020e-44c4-b3bc-3dc64e88ed45', NULL, 'Bought a beautiful leather journal. Time to fill it with thoughts.', '2025-06-07 19:39:04.766141+00', '2025-06-07 19:39:04.766141+00', NULL);
INSERT INTO public.posts VALUES ('47d7751c-7bc4-4dd3-8c10-a25918732741', '56e89866-3a90-4f97-8461-34b24f842da4', '912c9992-debd-42b8-9df7-28a9c4b7a319', 'd57937da-618e-4c6b-bc5b-e757af98582e', NULL, 'Wrote consistently for 4 days. Writing helps me organize my thoughts.', '2025-06-07 19:39:29.731122+00', '2025-06-07 19:39:29.731122+00', NULL);
INSERT INTO public.posts VALUES ('3413acc6-1419-4111-9328-0485f0ed1ef8', '56e89866-3a90-4f97-8461-34b24f842da4', 'c84f3282-901c-4433-82ee-42fc3f4ddf3a', '429eac06-eca0-49fe-9cc0-6c731b297206', NULL, 'Practiced German greetings. Saying ''Guten Tag'' to everyone now!', '2025-06-07 19:21:15.233012+00', '2025-06-07 19:21:15.233012+00', NULL);
INSERT INTO public.posts VALUES ('71a82c6e-6511-4cf5-8a80-099a7ee1f829', '56e89866-3a90-4f97-8461-34b24f842da4', 'c84f3282-901c-4433-82ee-42fc3f4ddf3a', '20e7d6e0-aaad-4946-84a9-7be9471fb444', NULL, 'Practiced travel phrases like ''Wo ist der Bahnhof?''. Feeling more confident.', '2025-06-07 16:49:42.213778+00', '2025-06-07 16:49:42.213778+00', NULL);
INSERT INTO public.posts VALUES ('a3f8de46-d113-48a1-b569-a20561d328f5', '56e89866-3a90-4f97-8461-34b24f842da4', 'c84f3282-901c-4433-82ee-42fc3f4ddf3a', '0e96f9b2-7005-4c25-9d0a-14bd722bebcb', NULL, 'Watched a German video with subtitles today. Caught a few new words!', '2025-06-07 16:21:38.416333+00', '2025-06-07 16:21:38.416333+00', NULL);
INSERT INTO public.posts VALUES ('14cd020e-3ee9-4114-829d-e3b66173c5f0', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '111e36b5-3a31-4079-a5c8-8ead7454723b', 'da1ef535-f977-4327-a352-fb097383d231', NULL, 'Started tracking my time with Toggl. Realized I''m spending too much time on YouTube.', '2025-06-07 19:12:13.847255+00', '2025-06-07 19:12:13.847255+00', NULL);
INSERT INTO public.posts VALUES ('03ebcf42-bfcf-4ee0-abae-752a1ded3096', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '111e36b5-3a31-4079-a5c8-8ead7454723b', 'da1ef535-f977-4327-a352-fb097383d231', NULL, 'Refined categories for better insights. Will try weekly review next.', '2025-06-07 18:32:16.865329+00', '2025-06-07 18:32:16.865329+00', NULL);
INSERT INTO public.posts VALUES ('f4d4efdf-82d7-491f-b9fc-4d10d40e6a1b', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '111e36b5-3a31-4079-a5c8-8ead7454723b', '0ea22b31-fb48-49de-9299-446fdf46623d', NULL, 'Set 3 weekly goals using Notion. Feeling more organized already.', '2025-06-07 18:29:50.533708+00', '2025-06-07 18:29:50.533708+00', NULL);
INSERT INTO public.posts VALUES ('93ae9cbf-d56a-4041-b094-ec070ef69045', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '111e36b5-3a31-4079-a5c8-8ead7454723b', '0ea22b31-fb48-49de-9299-446fdf46623d', NULL, 'Reflected on this week: hit 2 of 3 goals. Need to improve consistency.', '2025-06-07 18:41:04.121645+00', '2025-06-07 18:41:04.121645+00', NULL);
INSERT INTO public.posts VALUES ('e003776a-dd4e-4c34-bfe2-503f25a993c3', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '47b0bc29-2ef0-4b57-a3b9-849968031526', '4db717b6-a224-4748-accf-17c23d4a7015', NULL, 'Bought the domain myportfolio.dev! Time to build.', '2025-06-07 16:48:14.923984+00', '2025-06-07 16:48:14.923984+00', NULL);
INSERT INTO public.posts VALUES ('b3621746-9032-46c1-9e9c-072c55b709a7', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '47b0bc29-2ef0-4b57-a3b9-849968031526', 'cab6f8f2-e842-45ee-8b47-75ee8c6a8c73', NULL, 'Designed homepage wireframe in Figma. Excited for the result.', '2025-06-07 19:20:57.420885+00', '2025-06-07 19:20:57.420885+00', NULL);
INSERT INTO public.posts VALUES ('9a0f1426-9829-4c07-89b4-288a6e2d2a33', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '47b0bc29-2ef0-4b57-a3b9-849968031526', 'cab6f8f2-e842-45ee-8b47-75ee8c6a8c73', NULL, 'Finalized the ''About Me'' section тАФ now it actually reflects me.', '2025-06-07 18:29:36.719493+00', '2025-06-07 18:29:36.719493+00', NULL);
INSERT INTO public.posts VALUES ('a21898ff-48aa-4eb6-aeab-050acb249a73', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '47b0bc29-2ef0-4b57-a3b9-849968031526', 'ae132a3a-a947-45a1-949b-cf40ff3b4560', NULL, 'Deployed my site using Vercel. ItтАЩs live!', '2025-06-07 16:22:13.988036+00', '2025-06-07 16:22:13.988036+00', NULL);
INSERT INTO public.posts VALUES ('d3aef78d-bbf8-4543-bda4-1946e61bcc4b', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '47b0bc29-2ef0-4b57-a3b9-849968031526', '7e1d9562-3ffd-40be-830d-0fd0e005beb7', NULL, 'Completed first cardio session. Felt amazing!', '2025-06-07 19:59:45.085868+00', '2025-06-07 19:59:45.085868+00', NULL);
INSERT INTO public.posts VALUES ('4836de17-5868-4800-bdc7-57144e5c3807', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '47b0bc29-2ef0-4b57-a3b9-849968031526', '7e1d9562-3ffd-40be-830d-0fd0e005beb7', NULL, 'Day 2 done тАФ increased pace slightly.', '2025-06-07 16:24:18.984369+00', '2025-06-07 16:24:18.984369+00', NULL);
INSERT INTO public.posts VALUES ('5931007c-9050-44ae-90af-0c9732651ff3', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '47b0bc29-2ef0-4b57-a3b9-849968031526', '7e1d9562-3ffd-40be-830d-0fd0e005beb7', NULL, 'Final session for the week. Feeling consistent!', '2025-06-07 16:19:31.707751+00', '2025-06-07 16:19:31.707751+00', NULL);
INSERT INTO public.posts VALUES ('0526699f-47c4-46b2-8846-a298a864ed24', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '47b0bc29-2ef0-4b57-a3b9-849968031526', '47bf77f8-3f6d-48f0-bdae-27574c0b902e', NULL, 'Did my first strength workout. Focused on upper body.', '2025-06-07 19:20:45.664415+00', '2025-06-07 19:20:45.664415+00', NULL);
INSERT INTO public.posts VALUES ('90782b89-1c9c-4c2f-a514-7719331b9c72', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '47b0bc29-2ef0-4b57-a3b9-849968031526', '47bf77f8-3f6d-48f0-bdae-27574c0b902e', NULL, 'Second strength day тАФ leg day!', '2025-06-07 18:31:42.232831+00', '2025-06-07 18:31:42.232831+00', NULL);
INSERT INTO public.posts VALUES ('85a48eda-14df-462a-a173-2c1b653afa4e', '6e137217-71a9-439b-a579-4bcd1a729af9', '30a43be9-a3fe-43c9-b536-25cb0ee96c1b', '771b2a91-478a-46de-afc3-b87f15a693fd', NULL, 'Cleared out a ton of old papers and cables today. Already feels more spacious!', '2025-06-07 19:38:34.013715+00', '2025-06-07 19:38:34.013715+00', NULL);
INSERT INTO public.posts VALUES ('384be8c2-daf3-4bb6-8e8a-69d83285c0df', '6e137217-71a9-439b-a579-4bcd1a729af9', '30a43be9-a3fe-43c9-b536-25cb0ee96c1b', '771b2a91-478a-46de-afc3-b87f15a693fd', NULL, 'Donated unused gadgets to a local center. Decluttering feels so freeing.', '2025-06-07 19:10:12.771858+00', '2025-06-07 19:10:12.771858+00', NULL);
INSERT INTO public.posts VALUES ('660c8c08-468b-4a0d-a37d-fe58d7d0b075', '6e137217-71a9-439b-a579-4bcd1a729af9', '30a43be9-a3fe-43c9-b536-25cb0ee96c1b', 'fe218910-6c0b-423c-bbf6-468e5b86844c', NULL, 'Moved my desk closer to the window. Natural light makes a big difference!', '2025-06-07 16:15:30.512253+00', '2025-06-07 16:15:30.512253+00', NULL);
INSERT INTO public.posts VALUES ('2d2a93c2-c55a-4fcc-8fb2-72b38f9e96a4', '6e137217-71a9-439b-a579-4bcd1a729af9', 'ac557757-04c4-428d-a616-655ecda8b969', '6eea2e14-0992-4f13-a5e8-d927886af1e1', NULL, 'Grouped 50 new English words into food, travel, and tech categories. Easier to memorize!', '2025-06-07 16:20:17.654467+00', '2025-06-07 16:20:17.654467+00', NULL);
INSERT INTO public.posts VALUES ('824b208f-844a-4b1c-b9e2-ce7e8148f06e', '6e137217-71a9-439b-a579-4bcd1a729af9', 'ac557757-04c4-428d-a616-655ecda8b969', 'adbd6df3-b447-4970-8d81-bb54b15310b5', NULL, 'Used flashcards during breakfast and on the bus. Repetition really works!', '2025-06-07 19:59:54.725475+00', '2025-06-07 19:59:54.725475+00', NULL);
INSERT INTO public.posts VALUES ('596dd13f-30c4-402d-9eb4-5e041370d12d', '6e137217-71a9-439b-a579-4bcd1a729af9', 'ac557757-04c4-428d-a616-655ecda8b969', 'adbd6df3-b447-4970-8d81-bb54b15310b5', NULL, 'Day 3: Reviewed verbs and adjectives. I''m starting to recall them faster.', '2025-06-07 19:58:41.513687+00', '2025-06-07 19:58:41.513687+00', NULL);
INSERT INTO public.posts VALUES ('a3624d47-236a-4b55-be16-d717bfcd8384', '6e137217-71a9-439b-a579-4bcd1a729af9', 'ac557757-04c4-428d-a616-655ecda8b969', 'b03ade4c-2749-4068-b191-d2d31129ec5c', NULL, 'Wrote a short story using 10 new words. Feeling more confident now.', '2025-06-07 16:49:49.831369+00', '2025-06-07 16:49:49.831369+00', NULL);
INSERT INTO public.posts VALUES ('7f9fcdb1-7602-4b70-8626-204ff26dae55', '6e137217-71a9-439b-a579-4bcd1a729af9', 'ac557757-04c4-428d-a616-655ecda8b969', 'b03ade4c-2749-4068-b191-d2d31129ec5c', NULL, 'Practiced sentence construction with complex verbs. Not easy, but fun!', '2025-06-07 19:09:35.519597+00', '2025-06-07 19:09:35.519597+00', NULL);
INSERT INTO public.posts VALUES ('c4e4435b-e210-481c-9603-880796a0b7df', '6e137217-71a9-439b-a579-4bcd1a729af9', 'c4c10327-612e-4a10-a744-853e80c41a49', 'd36ac6f3-57a2-4d11-ba34-9674e46c7bec', NULL, 'Installed Krita on my tablet. Interface looks neat!', '2025-06-07 16:16:16.75267+00', '2025-06-07 16:16:16.75267+00', NULL);
INSERT INTO public.posts VALUES ('dee234bb-394b-4471-a913-7e6153be0bb7', '6e137217-71a9-439b-a579-4bcd1a729af9', 'c4c10327-612e-4a10-a744-853e80c41a49', '106936c7-f18f-48c6-81ab-6f931815cced', NULL, 'Followed a tutorial on drawing cartoon eyes. Simple but satisfying!', '2025-06-07 18:42:48.757811+00', '2025-06-07 18:42:48.757811+00', NULL);
INSERT INTO public.posts VALUES ('9508badc-12e8-4390-aa8f-7763d4dbde96', '6e137217-71a9-439b-a579-4bcd1a729af9', 'c4c10327-612e-4a10-a744-853e80c41a49', '106936c7-f18f-48c6-81ab-6f931815cced', NULL, 'Practiced shading techniques using layers. Digital drawing is awesome.', '2025-06-07 16:48:07.003084+00', '2025-06-07 16:48:07.003084+00', NULL);
INSERT INTO public.posts VALUES ('6a82001d-b416-4515-913e-5b6e3f57fbdf', '6e137217-71a9-439b-a579-4bcd1a729af9', 'c4c10327-612e-4a10-a744-853e80c41a49', 'c6dd0c27-420c-4123-81c0-0744d6e3a7a6', NULL, 'Drew 3 simple characters using different brushes. Felt proud of the progress!', '2025-06-07 18:40:43.109651+00', '2025-06-07 18:40:43.109651+00', NULL);
INSERT INTO public.posts VALUES ('3487baed-69d9-4a07-ba09-ea0e9c2699a4', '79758f94-70d8-4a11-b05f-793fcb57e093', '70679b66-1611-48c0-9bcb-0558b6d28965', 'ba72e44e-efb7-4331-a79f-a36933bc4f35', NULL, 'Completed Chapter 1 of the grammar book. A bit tricky, but IтАЩm getting there!', '2025-06-07 16:22:21.42636+00', '2025-06-07 16:22:21.42636+00', NULL);
INSERT INTO public.posts VALUES ('ab773301-3e88-4c8b-9b75-34d5654e2a2a', '79758f94-70d8-4a11-b05f-793fcb57e093', '70679b66-1611-48c0-9bcb-0558b6d28965', 'ba72e44e-efb7-4331-a79f-a36933bc4f35', NULL, 'Reviewed dative and accusative cases. ItтАЩs starting to make sense now!', '2025-06-07 19:09:50.73246+00', '2025-06-07 19:09:50.73246+00', NULL);
INSERT INTO public.posts VALUES ('807dc60c-27a5-44e8-a20c-0ee8ce41ede8', '79758f94-70d8-4a11-b05f-793fcb57e093', '70679b66-1611-48c0-9bcb-0558b6d28965', '7dbec861-f7dd-47dc-9695-c3fdd5f60791', NULL, 'Took my first mock A2 test. Got 75%, need to improve listening section.', '2025-06-07 18:30:03.693167+00', '2025-06-07 18:30:03.693167+00', NULL);
INSERT INTO public.posts VALUES ('c2901639-580e-46f1-a0b9-12e6a970e7d4', '79758f94-70d8-4a11-b05f-793fcb57e093', '70679b66-1611-48c0-9bcb-0558b6d28965', '12cccfe8-320f-4992-9d2f-b59ef5ea4f6e', NULL, 'Had a 30-minute Zoom call practicing small talk with a German buddy.', '2025-06-07 19:38:45.225947+00', '2025-06-07 19:38:45.225947+00', NULL);
INSERT INTO public.posts VALUES ('c56780c9-07da-467d-8d5b-5236215603af', '79758f94-70d8-4a11-b05f-793fcb57e093', '7e8e87bc-1b4c-44e9-832c-6e29617642d8', '4d943fed-81fb-4660-a2f5-c64779b63a7e', NULL, 'Did morning stretches for 4 days straight. My back feels much looser!', '2025-06-07 16:14:35.740502+00', '2025-06-07 16:14:35.740502+00', NULL);
INSERT INTO public.posts VALUES ('40123287-2503-4a02-bd89-93f13eea21be', '79758f94-70d8-4a11-b05f-793fcb57e093', '7e8e87bc-1b4c-44e9-832c-6e29617642d8', '4369d19b-7570-47b9-9f06-f72b2eee78f2', NULL, 'Bought a lumbar support cushion. Already noticing better posture at my desk.', '2025-06-07 19:19:09.181576+00', '2025-06-07 19:19:09.181576+00', NULL);
INSERT INTO public.posts VALUES ('fbbcab19-51fd-4eb4-aa4b-750782900d5d', '79758f94-70d8-4a11-b05f-793fcb57e093', 'a9528f6b-3895-4cab-af6c-62a93ef1b83e', 'fe5cb029-80a1-4c37-8fd0-62af0afc0976', NULL, 'Finished the homepage wireframe in Figma. Happy with the minimal style!', '2025-06-07 18:32:29.925977+00', '2025-06-07 18:32:29.925977+00', NULL);
INSERT INTO public.posts VALUES ('cfd2ccd0-49d2-4a73-8f13-14348a779f6f', '79758f94-70d8-4a11-b05f-793fcb57e093', 'a9528f6b-3895-4cab-af6c-62a93ef1b83e', 'd81f001e-91a4-4a34-aeb8-90544b4c328e', NULL, 'Implemented responsive navbar and hero section with TailwindCSS.', '2025-06-07 18:33:11.591281+00', '2025-06-07 18:33:11.591281+00', NULL);
INSERT INTO public.posts VALUES ('970217d3-a37f-4bea-a894-b26a21d7845f', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '513f748c-0dc9-4b04-a8f0-f33ea5892151', '97ff5a74-3b7d-40ce-ab6b-7153f65a446a', NULL, 'Sketched out the initial DB schema using dbdiagram.io тАФ looks clean and minimal.', '2025-06-07 20:00:00.609925+00', '2025-06-07 20:00:00.609925+00', NULL);
INSERT INTO public.posts VALUES ('8b91fce6-8708-47e5-b7c5-9eee8c55b696', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '513f748c-0dc9-4b04-a8f0-f33ea5892151', '4ef71a7d-5c7a-47b4-bead-b825ef583b04', NULL, 'Implemented Create and Read endpoints using Gin framework. Things are going smooth!', '2025-06-07 19:23:09.668842+00', '2025-06-07 19:23:09.668842+00', NULL);
INSERT INTO public.posts VALUES ('acc9d130-e2ea-426e-9553-d3ee35589a7f', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '513f748c-0dc9-4b04-a8f0-f33ea5892151', '4ef71a7d-5c7a-47b4-bead-b825ef583b04', NULL, 'Added error handling and validations for the API inputs.', '2025-06-07 18:40:55.774318+00', '2025-06-07 18:40:55.774318+00', NULL);
INSERT INTO public.posts VALUES ('33f2ef0e-d58d-4b7a-a5fe-89b87390accd', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'b438fcc1-9526-4e3c-8bd2-039bc7e71e7b', '9aead199-68d1-4dc2-b42e-af82e4932380', NULL, 'Outlined key sections: Introduction, Applications, Challenges, and Future Outlook.', '2025-06-07 16:23:59.428892+00', '2025-06-07 16:23:59.428892+00', NULL);
INSERT INTO public.posts VALUES ('a15c1a09-2a18-4cc2-9915-6c3aa7007455', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'b438fcc1-9526-4e3c-8bd2-039bc7e71e7b', 'a40da453-3b25-4f00-97c3-04c8ac10f93b', NULL, 'Completed the introduction and added examples of ChatGPT in classrooms.', '2025-06-07 16:50:01.949929+00', '2025-06-07 16:50:01.949929+00', NULL);
INSERT INTO public.posts VALUES ('789062db-6670-4972-97f9-21f61cdbd723', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'b438fcc1-9526-4e3c-8bd2-039bc7e71e7b', 'a40da453-3b25-4f00-97c3-04c8ac10f93b', NULL, 'Explained AI benefits for adaptive learning in section two.', '2025-06-07 19:37:25.747028+00', '2025-06-07 19:37:25.747028+00', NULL);
INSERT INTO public.posts VALUES ('e2d16c26-ba68-4a8d-ab33-8a68890f1683', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'f3984ee6-5f7b-47b0-a88c-c684e10eae43', '950b125d-11c7-4703-877c-a87c1cb027b0', NULL, 'Managed to wake up at 6:25 AM for three days straight.', '2025-06-07 19:59:26.408471+00', '2025-06-07 19:59:26.408471+00', NULL);
INSERT INTO public.posts VALUES ('6a54c0a1-ec0f-45b9-8620-c5d20ba82a03', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'f3984ee6-5f7b-47b0-a88c-c684e10eae43', '141ea941-bf50-45ec-8387-deefb7e716ed', NULL, 'Did 10 min yoga flow for neck and spine today.', '2025-06-07 18:41:33.756394+00', '2025-06-07 18:41:33.756394+00', NULL);
INSERT INTO public.posts VALUES ('219facb1-badb-49bc-9395-e2c980236804', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'f3984ee6-5f7b-47b0-a88c-c684e10eae43', '141ea941-bf50-45ec-8387-deefb7e716ed', NULL, 'Stretched hamstrings and shoulders post-walk. Felt great!', '2025-06-07 19:21:03.688414+00', '2025-06-07 19:21:03.688414+00', NULL);
INSERT INTO public.posts VALUES ('02dd7349-fe64-4ff1-974b-b1673d7ac009', '79758f94-70d8-4a11-b05f-793fcb57e093', 'a9528f6b-3895-4cab-af6c-62a93ef1b83e', 'd81f001e-91a4-4a34-aeb8-90544b4c328e', NULL, 'Deployed first version of the portfolio to GitHub Pages! ', '2025-06-07 19:12:21.201689+00', '2025-06-07 19:12:21.201689+00', NULL);
INSERT INTO public.posts VALUES ('1c13eebf-531d-4c1e-9560-69c914c649e6', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'a9528f6b-3895-4cab-af6c-62a93ef1b83e', 'd81f001e-91a4-4a34-aeb8-90544b4c328e', '02dd7349-fe64-4ff1-974b-b1673d7ac009', 'Motivating stuff!', '2025-06-07 21:12:21.201689+00', '2025-06-07 21:12:21.201689+00', NULL);
INSERT INTO public.posts VALUES ('f8f59a6d-6a50-405b-9b62-9561d552d736', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'a9528f6b-3895-4cab-af6c-62a93ef1b83e', 'd81f001e-91a4-4a34-aeb8-90544b4c328e', '02dd7349-fe64-4ff1-974b-b1673d7ac009', 'I had the same thought.', '2025-06-07 21:38:21.201689+00', '2025-06-07 21:38:21.201689+00', NULL);
INSERT INTO public.posts VALUES ('637f254f-af75-4690-a93f-c6652addffcc', '56e89866-3a90-4f97-8461-34b24f842da4', 'a9528f6b-3895-4cab-af6c-62a93ef1b83e', 'd81f001e-91a4-4a34-aeb8-90544b4c328e', '02dd7349-fe64-4ff1-974b-b1673d7ac009', 'Love this!', '2025-06-07 19:31:21.201689+00', '2025-06-07 19:31:21.201689+00', NULL);
INSERT INTO public.posts VALUES ('50e48281-cb2b-4ca3-a595-330b8d2fef53', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'a9528f6b-3895-4cab-af6c-62a93ef1b83e', 'd81f001e-91a4-4a34-aeb8-90544b4c328e', '02dd7349-fe64-4ff1-974b-b1673d7ac009', 'Nice progress!', '2025-06-07 19:40:21.201689+00', '2025-06-07 19:40:21.201689+00', NULL);
INSERT INTO public.posts VALUES ('cadea142-016d-4fa9-b144-221a5632c242', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'a9528f6b-3895-4cab-af6c-62a93ef1b83e', 'd81f001e-91a4-4a34-aeb8-90544b4c328e', '02dd7349-fe64-4ff1-974b-b1673d7ac009', 'Keep going!', '2025-06-07 20:26:21.201689+00', '2025-06-07 20:26:21.201689+00', NULL);
INSERT INTO public.posts VALUES ('b116de78-41ca-498d-98ad-899c376c80b1', '6e137217-71a9-439b-a579-4bcd1a729af9', 'a9528f6b-3895-4cab-af6c-62a93ef1b83e', 'd81f001e-91a4-4a34-aeb8-90544b4c328e', '02dd7349-fe64-4ff1-974b-b1673d7ac009', 'Following your example.', '2025-06-07 20:41:21.201689+00', '2025-06-07 20:41:21.201689+00', NULL);
INSERT INTO public.posts VALUES ('0206b5cc-367f-4b99-bf67-2f77c54f5913', '6e137217-71a9-439b-a579-4bcd1a729af9', '111e36b5-3a31-4079-a5c8-8ead7454723b', 'da1ef535-f977-4327-a352-fb097383d231', '03ebcf42-bfcf-4ee0-abae-752a1ded3096', 'Love this!', '2025-06-07 20:29:16.865329+00', '2025-06-07 20:29:16.865329+00', NULL);
INSERT INTO public.posts VALUES ('5ce00cfd-afea-4785-b9d3-677eb2338dee', '56e89866-3a90-4f97-8461-34b24f842da4', '111e36b5-3a31-4079-a5c8-8ead7454723b', 'da1ef535-f977-4327-a352-fb097383d231', '03ebcf42-bfcf-4ee0-abae-752a1ded3096', 'Nice progress!', '2025-06-07 19:12:16.865329+00', '2025-06-07 19:12:16.865329+00', NULL);
INSERT INTO public.posts VALUES ('a9f0f95b-0fbb-4feb-bd14-31c5e4e89cd0', '79758f94-70d8-4a11-b05f-793fcb57e093', '111e36b5-3a31-4079-a5c8-8ead7454723b', 'da1ef535-f977-4327-a352-fb097383d231', '03ebcf42-bfcf-4ee0-abae-752a1ded3096', 'I had the same thought.', '2025-06-07 20:30:16.865329+00', '2025-06-07 20:30:16.865329+00', NULL);
INSERT INTO public.posts VALUES ('1ea7a694-3858-4fc9-8234-2f1c97266c12', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '111e36b5-3a31-4079-a5c8-8ead7454723b', 'da1ef535-f977-4327-a352-fb097383d231', '03ebcf42-bfcf-4ee0-abae-752a1ded3096', 'Following your example.', '2025-06-07 18:36:16.865329+00', '2025-06-07 18:36:16.865329+00', NULL);
INSERT INTO public.posts VALUES ('7482f89f-91f0-42ae-8582-36c6e0e8ab80', '3b0ebec9-3805-490b-987f-4c6e7286319e', '111e36b5-3a31-4079-a5c8-8ead7454723b', 'da1ef535-f977-4327-a352-fb097383d231', '03ebcf42-bfcf-4ee0-abae-752a1ded3096', 'Interesting!', '2025-06-07 19:21:16.865329+00', '2025-06-07 19:21:16.865329+00', NULL);
INSERT INTO public.posts VALUES ('843f7468-f256-4241-b1f0-7a5220338502', '56e89866-3a90-4f97-8461-34b24f842da4', '47b0bc29-2ef0-4b57-a3b9-849968031526', '47bf77f8-3f6d-48f0-bdae-27574c0b902e', '0526699f-47c4-46b2-8846-a298a864ed24', 'Motivating stuff!', '2025-06-07 21:24:45.664415+00', '2025-06-07 21:24:45.664415+00', NULL);
INSERT INTO public.posts VALUES ('8bd157e5-35cc-4991-9ca6-fd6bccd9a275', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '47b0bc29-2ef0-4b57-a3b9-849968031526', '47bf77f8-3f6d-48f0-bdae-27574c0b902e', '0526699f-47c4-46b2-8846-a298a864ed24', 'Following your example.', '2025-06-07 20:23:45.664415+00', '2025-06-07 20:23:45.664415+00', NULL);
INSERT INTO public.posts VALUES ('9da69496-7eac-4e47-a21b-8f02cfd61ae9', '3b0ebec9-3805-490b-987f-4c6e7286319e', '47b0bc29-2ef0-4b57-a3b9-849968031526', '47bf77f8-3f6d-48f0-bdae-27574c0b902e', '0526699f-47c4-46b2-8846-a298a864ed24', 'I had the same thought.', '2025-06-07 20:11:45.664415+00', '2025-06-07 20:11:45.664415+00', NULL);
INSERT INTO public.posts VALUES ('5f8be4cb-d4ab-4968-871a-0c98b3e1c4ad', '6e137217-71a9-439b-a579-4bcd1a729af9', '47b0bc29-2ef0-4b57-a3b9-849968031526', '47bf77f8-3f6d-48f0-bdae-27574c0b902e', '0526699f-47c4-46b2-8846-a298a864ed24', 'Very cool!', '2025-06-07 21:56:45.664415+00', '2025-06-07 21:56:45.664415+00', NULL);
INSERT INTO public.posts VALUES ('e38e8216-cf08-4577-9bbb-32e3f544c02d', '6e137217-71a9-439b-a579-4bcd1a729af9', '6b7011e6-fb89-4542-8eba-31b219970185', 'c72f0f00-94be-41ae-ade0-954267d26c9d', '0abf63f3-6d91-4737-add9-edb192813b76', 'I had the same thought.', '2025-06-07 17:04:23.923606+00', '2025-06-07 17:04:23.923606+00', NULL);
INSERT INTO public.posts VALUES ('778ea4a7-c1a3-43d8-9ca3-f017584a7894', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '6b7011e6-fb89-4542-8eba-31b219970185', 'c72f0f00-94be-41ae-ade0-954267d26c9d', '0abf63f3-6d91-4737-add9-edb192813b76', 'Love this!', '2025-06-07 17:24:23.923606+00', '2025-06-07 17:24:23.923606+00', NULL);
INSERT INTO public.posts VALUES ('833f48f2-0d57-423a-8f9f-d8790684be6b', '3b0ebec9-3805-490b-987f-4c6e7286319e', '6b7011e6-fb89-4542-8eba-31b219970185', 'c72f0f00-94be-41ae-ade0-954267d26c9d', '0abf63f3-6d91-4737-add9-edb192813b76', 'Interesting!', '2025-06-07 19:01:23.923606+00', '2025-06-07 19:01:23.923606+00', NULL);
INSERT INTO public.posts VALUES ('4271c429-9631-4285-b85d-fc53dcd1b61d', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '6b7011e6-fb89-4542-8eba-31b219970185', 'c72f0f00-94be-41ae-ade0-954267d26c9d', '0abf63f3-6d91-4737-add9-edb192813b76', 'Keep going!', '2025-06-07 19:41:23.923606+00', '2025-06-07 19:41:23.923606+00', NULL);
INSERT INTO public.posts VALUES ('5e81ffb4-1692-45bf-87b2-523d4153ad82', '79758f94-70d8-4a11-b05f-793fcb57e093', '6b7011e6-fb89-4542-8eba-31b219970185', 'c72f0f00-94be-41ae-ade0-954267d26c9d', '0abf63f3-6d91-4737-add9-edb192813b76', 'That’s awesome!', '2025-06-07 18:36:23.923606+00', '2025-06-07 18:36:23.923606+00', NULL);
INSERT INTO public.posts VALUES ('b60f407b-3dcb-4287-9802-b341c16c83b2', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '6b7011e6-fb89-4542-8eba-31b219970185', 'c72f0f00-94be-41ae-ade0-954267d26c9d', '0abf63f3-6d91-4737-add9-edb192813b76', 'Motivating stuff!', '2025-06-07 19:05:23.923606+00', '2025-06-07 19:05:23.923606+00', NULL);
INSERT INTO public.posts VALUES ('2bba9765-112a-470b-831a-1ed94b8ef6e8', '6e137217-71a9-439b-a579-4bcd1a729af9', '9edaf28d-e47c-4bec-aa12-f4d6b78c6edb', 'b377c3ec-dfcf-48b0-9ece-f3898d31ae9c', '1402f5b0-c845-4665-998c-e5597f5f5521', 'Nice progress!', '2025-06-07 21:44:48.43565+00', '2025-06-07 21:44:48.43565+00', NULL);
INSERT INTO public.posts VALUES ('c53eca81-cb47-46cf-a597-81d35cd71c41', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '9edaf28d-e47c-4bec-aa12-f4d6b78c6edb', 'b377c3ec-dfcf-48b0-9ece-f3898d31ae9c', '1402f5b0-c845-4665-998c-e5597f5f5521', 'Keep going!', '2025-06-07 20:53:48.43565+00', '2025-06-07 20:53:48.43565+00', NULL);
INSERT INTO public.posts VALUES ('cb86771d-43ff-4bfb-b5c4-6304a38fbe8f', '79758f94-70d8-4a11-b05f-793fcb57e093', '9edaf28d-e47c-4bec-aa12-f4d6b78c6edb', 'b377c3ec-dfcf-48b0-9ece-f3898d31ae9c', '1402f5b0-c845-4665-998c-e5597f5f5521', 'Interesting!', '2025-06-07 19:29:48.43565+00', '2025-06-07 19:29:48.43565+00', NULL);
INSERT INTO public.posts VALUES ('02e29416-17f6-4086-8609-ca3113030ac9', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '9edaf28d-e47c-4bec-aa12-f4d6b78c6edb', 'b377c3ec-dfcf-48b0-9ece-f3898d31ae9c', '1402f5b0-c845-4665-998c-e5597f5f5521', 'Motivating stuff!', '2025-06-07 21:00:48.43565+00', '2025-06-07 21:00:48.43565+00', NULL);
INSERT INTO public.posts VALUES ('c0bdc85b-fcae-4132-8aee-3039dd805359', '6e137217-71a9-439b-a579-4bcd1a729af9', '111e36b5-3a31-4079-a5c8-8ead7454723b', 'da1ef535-f977-4327-a352-fb097383d231', '14cd020e-3ee9-4114-829d-e3b66173c5f0', 'Keep going!', '2025-06-07 21:26:13.847255+00', '2025-06-07 21:26:13.847255+00', NULL);
INSERT INTO public.posts VALUES ('9cf3d6aa-cd8c-4769-9870-95273bea0706', '3b0ebec9-3805-490b-987f-4c6e7286319e', '111e36b5-3a31-4079-a5c8-8ead7454723b', 'da1ef535-f977-4327-a352-fb097383d231', '14cd020e-3ee9-4114-829d-e3b66173c5f0', 'That’s awesome!', '2025-06-07 20:06:13.847255+00', '2025-06-07 20:06:13.847255+00', NULL);
INSERT INTO public.posts VALUES ('0b6299e8-cf60-4890-be87-f56cf32df4fa', '56e89866-3a90-4f97-8461-34b24f842da4', '111e36b5-3a31-4079-a5c8-8ead7454723b', 'da1ef535-f977-4327-a352-fb097383d231', '14cd020e-3ee9-4114-829d-e3b66173c5f0', 'Love this!', '2025-06-07 19:57:13.847255+00', '2025-06-07 19:57:13.847255+00', NULL);
INSERT INTO public.posts VALUES ('d1c1777b-aed3-4dfd-895e-e0efabc39863', '79758f94-70d8-4a11-b05f-793fcb57e093', '111e36b5-3a31-4079-a5c8-8ead7454723b', 'da1ef535-f977-4327-a352-fb097383d231', '14cd020e-3ee9-4114-829d-e3b66173c5f0', 'Following your example.', '2025-06-07 21:22:13.847255+00', '2025-06-07 21:22:13.847255+00', NULL);
INSERT INTO public.posts VALUES ('acfb86a4-d5d4-40af-afcb-3023fb08a119', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '111e36b5-3a31-4079-a5c8-8ead7454723b', 'da1ef535-f977-4327-a352-fb097383d231', '14cd020e-3ee9-4114-829d-e3b66173c5f0', 'Very cool!', '2025-06-07 22:12:13.847255+00', '2025-06-07 22:12:13.847255+00', NULL);
INSERT INTO public.posts VALUES ('e80a2495-961a-4721-9702-07c4fe64687c', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '111e36b5-3a31-4079-a5c8-8ead7454723b', 'da1ef535-f977-4327-a352-fb097383d231', '14cd020e-3ee9-4114-829d-e3b66173c5f0', 'Nice progress!', '2025-06-07 19:14:13.847255+00', '2025-06-07 19:14:13.847255+00', NULL);
INSERT INTO public.posts VALUES ('c7a75cdc-3c78-4a64-9917-a6020a26f972', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '609eba42-64bc-460a-9468-bc011dae591c', '239bd055-307a-4344-97a7-211b820620f0', '19e2449f-0228-4c1a-b7c7-c7c8c0d7a4a4', 'Keep going!', '2025-06-07 20:32:10.456679+00', '2025-06-07 20:32:10.456679+00', NULL);
INSERT INTO public.posts VALUES ('4c166df3-7213-457a-9ec0-d1055b4267b8', '56e89866-3a90-4f97-8461-34b24f842da4', '609eba42-64bc-460a-9468-bc011dae591c', '239bd055-307a-4344-97a7-211b820620f0', '19e2449f-0228-4c1a-b7c7-c7c8c0d7a4a4', 'This inspired me.', '2025-06-07 21:24:10.456679+00', '2025-06-07 21:24:10.456679+00', NULL);
INSERT INTO public.posts VALUES ('5cc67d7a-a28e-4b4b-9d05-9fa71e2c3516', '3b0ebec9-3805-490b-987f-4c6e7286319e', '609eba42-64bc-460a-9468-bc011dae591c', '239bd055-307a-4344-97a7-211b820620f0', '19e2449f-0228-4c1a-b7c7-c7c8c0d7a4a4', 'Following your example.', '2025-06-07 21:52:10.456679+00', '2025-06-07 21:52:10.456679+00', NULL);
INSERT INTO public.posts VALUES ('b0e6f6f3-600c-47e8-9925-dfd365663ce7', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '609eba42-64bc-460a-9468-bc011dae591c', '239bd055-307a-4344-97a7-211b820620f0', '19e2449f-0228-4c1a-b7c7-c7c8c0d7a4a4', 'That’s awesome!', '2025-06-07 21:52:10.456679+00', '2025-06-07 21:52:10.456679+00', NULL);
INSERT INTO public.posts VALUES ('bd1f9b54-3796-4e68-95d0-1ea0ade7a28d', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '912c9992-debd-42b8-9df7-28a9c4b7a319', 'ff2ec17a-020e-44c4-b3bc-3dc64e88ed45', '1e976362-73b0-400a-a74c-3fa951afdbeb', 'I had the same thought.', '2025-06-07 22:14:04.766141+00', '2025-06-07 22:14:04.766141+00', NULL);
INSERT INTO public.posts VALUES ('f3decec2-f963-46ca-b772-0ac93db3633e', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '912c9992-debd-42b8-9df7-28a9c4b7a319', 'ff2ec17a-020e-44c4-b3bc-3dc64e88ed45', '1e976362-73b0-400a-a74c-3fa951afdbeb', 'Nice progress!', '2025-06-07 21:11:04.766141+00', '2025-06-07 21:11:04.766141+00', NULL);
INSERT INTO public.posts VALUES ('90767895-7c2b-4771-a644-1e95b5103feb', '79758f94-70d8-4a11-b05f-793fcb57e093', '912c9992-debd-42b8-9df7-28a9c4b7a319', 'ff2ec17a-020e-44c4-b3bc-3dc64e88ed45', '1e976362-73b0-400a-a74c-3fa951afdbeb', 'Very cool!', '2025-06-07 20:54:04.766141+00', '2025-06-07 20:54:04.766141+00', NULL);
INSERT INTO public.posts VALUES ('e7e7a1f9-0976-4222-a220-3896f5aba5c2', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '912c9992-debd-42b8-9df7-28a9c4b7a319', 'ff2ec17a-020e-44c4-b3bc-3dc64e88ed45', '1e976362-73b0-400a-a74c-3fa951afdbeb', 'Following your example.', '2025-06-07 21:02:04.766141+00', '2025-06-07 21:02:04.766141+00', NULL);
INSERT INTO public.posts VALUES ('47b93a87-9475-45da-ab9d-2e733f1e0e35', '6e137217-71a9-439b-a579-4bcd1a729af9', 'f3984ee6-5f7b-47b0-a88c-c684e10eae43', '141ea941-bf50-45ec-8387-deefb7e716ed', '219facb1-badb-49bc-9395-e2c980236804', 'Love this!', '2025-06-07 22:18:03.688414+00', '2025-06-07 22:18:03.688414+00', NULL);
INSERT INTO public.posts VALUES ('e36ac40f-db03-4497-8687-792ae0fd3d52', '56e89866-3a90-4f97-8461-34b24f842da4', 'f3984ee6-5f7b-47b0-a88c-c684e10eae43', '141ea941-bf50-45ec-8387-deefb7e716ed', '219facb1-badb-49bc-9395-e2c980236804', 'Following your example.', '2025-06-07 21:28:03.688414+00', '2025-06-07 21:28:03.688414+00', NULL);
INSERT INTO public.posts VALUES ('bc494d68-0bdd-45fe-a120-0ab08384eae5', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'f3984ee6-5f7b-47b0-a88c-c684e10eae43', '141ea941-bf50-45ec-8387-deefb7e716ed', '219facb1-badb-49bc-9395-e2c980236804', 'Nice progress!', '2025-06-07 20:35:03.688414+00', '2025-06-07 20:35:03.688414+00', NULL);
INSERT INTO public.posts VALUES ('220c8be5-6bd5-4a86-a3a0-b43d0804a83a', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'f3984ee6-5f7b-47b0-a88c-c684e10eae43', '141ea941-bf50-45ec-8387-deefb7e716ed', '219facb1-badb-49bc-9395-e2c980236804', 'I had the same thought.', '2025-06-07 19:24:03.688414+00', '2025-06-07 19:24:03.688414+00', NULL);
INSERT INTO public.posts VALUES ('200bbfa6-1006-4844-8c0a-b7cd3aceb5b0', '79758f94-70d8-4a11-b05f-793fcb57e093', 'f3984ee6-5f7b-47b0-a88c-c684e10eae43', '141ea941-bf50-45ec-8387-deefb7e716ed', '219facb1-badb-49bc-9395-e2c980236804', 'Interesting!', '2025-06-07 22:06:03.688414+00', '2025-06-07 22:06:03.688414+00', NULL);
INSERT INTO public.posts VALUES ('74a5911d-9439-4da5-bc7c-c170294788f3', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'f3984ee6-5f7b-47b0-a88c-c684e10eae43', '141ea941-bf50-45ec-8387-deefb7e716ed', '219facb1-badb-49bc-9395-e2c980236804', 'Motivating stuff!', '2025-06-07 19:47:03.688414+00', '2025-06-07 19:47:03.688414+00', NULL);
INSERT INTO public.posts VALUES ('beb6e3f1-ad34-408c-86e9-264266ad6c29', '6e137217-71a9-439b-a579-4bcd1a729af9', '6b7011e6-fb89-4542-8eba-31b219970185', '7214781e-ee0e-444d-9a4c-a6ec49f8b9be', '223da9a0-82b1-4745-9a4a-5b657d47299e', 'Nice progress!', '2025-06-07 18:05:50.550104+00', '2025-06-07 18:05:50.550104+00', NULL);
INSERT INTO public.posts VALUES ('c9badfc6-2a80-4558-9606-012ea4069f80', '3b0ebec9-3805-490b-987f-4c6e7286319e', '6b7011e6-fb89-4542-8eba-31b219970185', '7214781e-ee0e-444d-9a4c-a6ec49f8b9be', '223da9a0-82b1-4745-9a4a-5b657d47299e', 'That’s awesome!', '2025-06-07 19:05:50.550104+00', '2025-06-07 19:05:50.550104+00', NULL);
INSERT INTO public.posts VALUES ('e28c01a0-e4bd-4cb4-b3c5-b4f2edadd426', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '6b7011e6-fb89-4542-8eba-31b219970185', '7214781e-ee0e-444d-9a4c-a6ec49f8b9be', '223da9a0-82b1-4745-9a4a-5b657d47299e', 'I had the same thought.', '2025-06-07 17:13:50.550104+00', '2025-06-07 17:13:50.550104+00', NULL);
INSERT INTO public.posts VALUES ('c7963d2b-926f-430e-8f37-a7ee167de346', '79758f94-70d8-4a11-b05f-793fcb57e093', '6b7011e6-fb89-4542-8eba-31b219970185', '7214781e-ee0e-444d-9a4c-a6ec49f8b9be', '223da9a0-82b1-4745-9a4a-5b657d47299e', 'Motivating stuff!', '2025-06-07 18:20:50.550104+00', '2025-06-07 18:20:50.550104+00', NULL);
INSERT INTO public.posts VALUES ('5cb35c5c-3a2e-4cfc-8d69-ae78e89b5709', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '6b7011e6-fb89-4542-8eba-31b219970185', '7214781e-ee0e-444d-9a4c-a6ec49f8b9be', '223da9a0-82b1-4745-9a4a-5b657d47299e', 'This inspired me.', '2025-06-07 17:31:50.550104+00', '2025-06-07 17:31:50.550104+00', NULL);
INSERT INTO public.posts VALUES ('cbf428db-3ecd-4ef2-8ad4-2b2f9359aa3b', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '6b7011e6-fb89-4542-8eba-31b219970185', '7214781e-ee0e-444d-9a4c-a6ec49f8b9be', '223da9a0-82b1-4745-9a4a-5b657d47299e', 'Love this!', '2025-06-07 16:25:50.550104+00', '2025-06-07 16:25:50.550104+00', NULL);
INSERT INTO public.posts VALUES ('901dfd62-3db9-4546-bc1f-b3054db54b86', '6e137217-71a9-439b-a579-4bcd1a729af9', '66b19273-144f-4dd0-971e-830c18c26b38', 'd89edf48-a9de-456a-aa44-9e669a382fd5', '23a24eb3-9a76-4807-8ff2-24ebdd6f11bb', 'I had the same thought.', '2025-06-07 21:29:56.195766+00', '2025-06-07 21:29:56.195766+00', NULL);
INSERT INTO public.posts VALUES ('da34bec5-8a5a-40f1-b6ed-48b69d045734', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '66b19273-144f-4dd0-971e-830c18c26b38', 'd89edf48-a9de-456a-aa44-9e669a382fd5', '23a24eb3-9a76-4807-8ff2-24ebdd6f11bb', 'Keep going!', '2025-06-07 20:31:56.195766+00', '2025-06-07 20:31:56.195766+00', NULL);
INSERT INTO public.posts VALUES ('5b70c4b0-9881-4e85-b6c7-92ba29c573a4', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '66b19273-144f-4dd0-971e-830c18c26b38', 'd89edf48-a9de-456a-aa44-9e669a382fd5', '23a24eb3-9a76-4807-8ff2-24ebdd6f11bb', 'That’s awesome!', '2025-06-07 21:24:56.195766+00', '2025-06-07 21:24:56.195766+00', NULL);
INSERT INTO public.posts VALUES ('ec6f60c1-4d84-44c8-81df-5bd83340011f', '79758f94-70d8-4a11-b05f-793fcb57e093', '66b19273-144f-4dd0-971e-830c18c26b38', 'd89edf48-a9de-456a-aa44-9e669a382fd5', '23a24eb3-9a76-4807-8ff2-24ebdd6f11bb', 'Interesting!', '2025-06-07 21:36:56.195766+00', '2025-06-07 21:36:56.195766+00', NULL);
INSERT INTO public.posts VALUES ('edae5488-c61f-4358-a1c8-93d0d7ab7b89', '56e89866-3a90-4f97-8461-34b24f842da4', '66b19273-144f-4dd0-971e-830c18c26b38', 'd89edf48-a9de-456a-aa44-9e669a382fd5', '23a24eb3-9a76-4807-8ff2-24ebdd6f11bb', 'Love this!', '2025-06-07 21:07:56.195766+00', '2025-06-07 21:07:56.195766+00', NULL);
INSERT INTO public.posts VALUES ('9a424d27-23e9-4713-aad1-e5810e002406', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '66b19273-144f-4dd0-971e-830c18c26b38', 'd89edf48-a9de-456a-aa44-9e669a382fd5', '23a24eb3-9a76-4807-8ff2-24ebdd6f11bb', 'Following your example.', '2025-06-07 22:17:56.195766+00', '2025-06-07 22:17:56.195766+00', NULL);
INSERT INTO public.posts VALUES ('4d27b072-2283-4fca-89bc-8fc1d41bdc47', '6e137217-71a9-439b-a579-4bcd1a729af9', '9dd80410-bd88-455a-b8fe-a5c2e39edcf0', '62fd5b6a-1833-4551-b403-82eb0dab6724', '23ad69f9-8c1f-45e1-bdb6-86ae97f0a2d2', 'This inspired me.', '2025-06-07 22:06:03.559862+00', '2025-06-07 22:06:03.559862+00', NULL);
INSERT INTO public.posts VALUES ('b628425a-e1b4-4c61-947d-8851a524a4e2', '3b0ebec9-3805-490b-987f-4c6e7286319e', '9dd80410-bd88-455a-b8fe-a5c2e39edcf0', '62fd5b6a-1833-4551-b403-82eb0dab6724', '23ad69f9-8c1f-45e1-bdb6-86ae97f0a2d2', 'Nice progress!', '2025-06-07 21:23:03.559862+00', '2025-06-07 21:23:03.559862+00', NULL);
INSERT INTO public.posts VALUES ('f53990ad-8e4b-42c1-9373-b5f58b633b00', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '9dd80410-bd88-455a-b8fe-a5c2e39edcf0', '62fd5b6a-1833-4551-b403-82eb0dab6724', '23ad69f9-8c1f-45e1-bdb6-86ae97f0a2d2', 'Keep going!', '2025-06-07 19:18:03.559862+00', '2025-06-07 19:18:03.559862+00', NULL);
INSERT INTO public.posts VALUES ('47f1acd6-ab7b-48cb-ac0e-79c7c17c62ff', '6e137217-71a9-439b-a579-4bcd1a729af9', '9dd80410-bd88-455a-b8fe-a5c2e39edcf0', '9b3da3d5-8874-4398-a1ac-07e835ce813b', '26e50906-4a53-4537-89c5-9971e86bb0ad', 'Interesting!', '2025-06-07 18:05:02.639102+00', '2025-06-07 18:05:02.639102+00', NULL);
INSERT INTO public.posts VALUES ('3bcbd51f-abda-4329-bc7a-2e8059a434f1', '56e89866-3a90-4f97-8461-34b24f842da4', '9dd80410-bd88-455a-b8fe-a5c2e39edcf0', '9b3da3d5-8874-4398-a1ac-07e835ce813b', '26e50906-4a53-4537-89c5-9971e86bb0ad', 'Love this!', '2025-06-07 18:53:02.639102+00', '2025-06-07 18:53:02.639102+00', NULL);
INSERT INTO public.posts VALUES ('e1b0467f-f445-4bb6-a90e-1036e3537255', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '9dd80410-bd88-455a-b8fe-a5c2e39edcf0', '9b3da3d5-8874-4398-a1ac-07e835ce813b', '26e50906-4a53-4537-89c5-9971e86bb0ad', 'Very cool!', '2025-06-07 18:05:02.639102+00', '2025-06-07 18:05:02.639102+00', NULL);
INSERT INTO public.posts VALUES ('5e592e68-e6dc-491b-9f99-e3720507d215', '3b0ebec9-3805-490b-987f-4c6e7286319e', '9dd80410-bd88-455a-b8fe-a5c2e39edcf0', '9b3da3d5-8874-4398-a1ac-07e835ce813b', '26e50906-4a53-4537-89c5-9971e86bb0ad', 'I had the same thought.', '2025-06-07 18:10:02.639102+00', '2025-06-07 18:10:02.639102+00', NULL);
INSERT INTO public.posts VALUES ('4a734e46-ae36-405c-b8e7-53804088bca0', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '9dd80410-bd88-455a-b8fe-a5c2e39edcf0', '9b3da3d5-8874-4398-a1ac-07e835ce813b', '26e50906-4a53-4537-89c5-9971e86bb0ad', 'Motivating stuff!', '2025-06-07 19:31:02.639102+00', '2025-06-07 19:31:02.639102+00', NULL);
INSERT INTO public.posts VALUES ('d34818aa-5d3d-45a6-82bc-070c48258597', '79758f94-70d8-4a11-b05f-793fcb57e093', '9dd80410-bd88-455a-b8fe-a5c2e39edcf0', '9b3da3d5-8874-4398-a1ac-07e835ce813b', '26e50906-4a53-4537-89c5-9971e86bb0ad', 'Nice progress!', '2025-06-07 17:38:02.639102+00', '2025-06-07 17:38:02.639102+00', NULL);
INSERT INTO public.posts VALUES ('e129c8a4-cc56-4557-a2c5-0ad0cea9e539', '79758f94-70d8-4a11-b05f-793fcb57e093', 'ac557757-04c4-428d-a616-655ecda8b969', '6eea2e14-0992-4f13-a5e8-d927886af1e1', '2d2a93c2-c55a-4fcc-8fb2-72b38f9e96a4', 'I had the same thought.', '2025-06-07 18:12:17.654467+00', '2025-06-07 18:12:17.654467+00', NULL);
INSERT INTO public.posts VALUES ('3f683356-475b-4c5d-ac9b-d05299d1908d', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'ac557757-04c4-428d-a616-655ecda8b969', '6eea2e14-0992-4f13-a5e8-d927886af1e1', '2d2a93c2-c55a-4fcc-8fb2-72b38f9e96a4', 'Love this!', '2025-06-07 18:48:17.654467+00', '2025-06-07 18:48:17.654467+00', NULL);
INSERT INTO public.posts VALUES ('72b93c25-602a-4299-bdab-7268d5f6e2f8', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'ac557757-04c4-428d-a616-655ecda8b969', '6eea2e14-0992-4f13-a5e8-d927886af1e1', '2d2a93c2-c55a-4fcc-8fb2-72b38f9e96a4', 'Following your example.', '2025-06-07 18:12:17.654467+00', '2025-06-07 18:12:17.654467+00', NULL);
INSERT INTO public.posts VALUES ('95196fc2-6f65-4b5d-9941-4894d1b6303e', '56e89866-3a90-4f97-8461-34b24f842da4', 'ac557757-04c4-428d-a616-655ecda8b969', '6eea2e14-0992-4f13-a5e8-d927886af1e1', '2d2a93c2-c55a-4fcc-8fb2-72b38f9e96a4', 'Keep going!', '2025-06-07 18:42:17.654467+00', '2025-06-07 18:42:17.654467+00', NULL);
INSERT INTO public.posts VALUES ('2f38cfba-f9e9-46df-a55c-74aec94b1732', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'ac557757-04c4-428d-a616-655ecda8b969', '6eea2e14-0992-4f13-a5e8-d927886af1e1', '2d2a93c2-c55a-4fcc-8fb2-72b38f9e96a4', 'Nice progress!', '2025-06-07 18:12:17.654467+00', '2025-06-07 18:12:17.654467+00', NULL);
INSERT INTO public.posts VALUES ('d0a5551b-9372-4712-9a7b-5b644df52011', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'ac557757-04c4-428d-a616-655ecda8b969', '6eea2e14-0992-4f13-a5e8-d927886af1e1', '2d2a93c2-c55a-4fcc-8fb2-72b38f9e96a4', 'Very cool!', '2025-06-07 18:11:17.654467+00', '2025-06-07 18:11:17.654467+00', NULL);
INSERT INTO public.posts VALUES ('0b7b73f7-a030-41d0-96a4-625ff53c5532', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '66b19273-144f-4dd0-971e-830c18c26b38', 'adb2d568-a751-4303-a5cd-67362c864091', '31f82a0d-869c-4ea9-8c72-152d229893cf', 'This inspired me.', '2025-06-07 22:38:11.392577+00', '2025-06-07 22:38:11.392577+00', NULL);
INSERT INTO public.posts VALUES ('035d8fa6-fe68-4575-bea8-3e49a4ab8389', '56e89866-3a90-4f97-8461-34b24f842da4', '66b19273-144f-4dd0-971e-830c18c26b38', 'adb2d568-a751-4303-a5cd-67362c864091', '31f82a0d-869c-4ea9-8c72-152d229893cf', 'That’s awesome!', '2025-06-07 22:19:11.392577+00', '2025-06-07 22:19:11.392577+00', NULL);
INSERT INTO public.posts VALUES ('c1c8464c-e5b5-4b2a-be04-ecfe6d366667', '79758f94-70d8-4a11-b05f-793fcb57e093', '66b19273-144f-4dd0-971e-830c18c26b38', 'adb2d568-a751-4303-a5cd-67362c864091', '31f82a0d-869c-4ea9-8c72-152d229893cf', 'Following your example.', '2025-06-07 20:52:11.392577+00', '2025-06-07 20:52:11.392577+00', NULL);
INSERT INTO public.posts VALUES ('cf9ae5e6-de60-422f-9401-233484e7dbb9', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '66b19273-144f-4dd0-971e-830c18c26b38', 'adb2d568-a751-4303-a5cd-67362c864091', '31f82a0d-869c-4ea9-8c72-152d229893cf', 'Love this!', '2025-06-07 20:40:11.392577+00', '2025-06-07 20:40:11.392577+00', NULL);
INSERT INTO public.posts VALUES ('a990a344-2ba6-4d9c-b186-bc4b5804c8e0', '6e137217-71a9-439b-a579-4bcd1a729af9', '66b19273-144f-4dd0-971e-830c18c26b38', 'adb2d568-a751-4303-a5cd-67362c864091', '31f82a0d-869c-4ea9-8c72-152d229893cf', 'I had the same thought.', '2025-06-07 20:42:11.392577+00', '2025-06-07 20:42:11.392577+00', NULL);
INSERT INTO public.posts VALUES ('f470b873-853f-474e-8a05-784d09601684', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '66b19273-144f-4dd0-971e-830c18c26b38', 'adb2d568-a751-4303-a5cd-67362c864091', '31f82a0d-869c-4ea9-8c72-152d229893cf', 'Interesting!', '2025-06-07 21:52:11.392577+00', '2025-06-07 21:52:11.392577+00', NULL);
INSERT INTO public.posts VALUES ('f42662bd-7fa2-470f-a0cc-64a9f23c8128', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'b438fcc1-9526-4e3c-8bd2-039bc7e71e7b', '9aead199-68d1-4dc2-b42e-af82e4932380', '33f2ef0e-d58d-4b7a-a5fe-89b87390accd', 'That’s awesome!', '2025-06-07 17:02:59.428892+00', '2025-06-07 17:02:59.428892+00', NULL);
INSERT INTO public.posts VALUES ('366c30a0-c870-4d9c-97a0-4ffd597e0e0f', '56e89866-3a90-4f97-8461-34b24f842da4', 'b438fcc1-9526-4e3c-8bd2-039bc7e71e7b', '9aead199-68d1-4dc2-b42e-af82e4932380', '33f2ef0e-d58d-4b7a-a5fe-89b87390accd', 'Following your example.', '2025-06-07 18:46:59.428892+00', '2025-06-07 18:46:59.428892+00', NULL);
INSERT INTO public.posts VALUES ('8b7c4ed1-e3cf-45a2-8451-7d515f21375f', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'b438fcc1-9526-4e3c-8bd2-039bc7e71e7b', '9aead199-68d1-4dc2-b42e-af82e4932380', '33f2ef0e-d58d-4b7a-a5fe-89b87390accd', 'Very cool!', '2025-06-07 16:32:59.428892+00', '2025-06-07 16:32:59.428892+00', NULL);
INSERT INTO public.posts VALUES ('c0854578-4332-46d4-b09c-5639a0b000e5', '79758f94-70d8-4a11-b05f-793fcb57e093', 'b438fcc1-9526-4e3c-8bd2-039bc7e71e7b', '9aead199-68d1-4dc2-b42e-af82e4932380', '33f2ef0e-d58d-4b7a-a5fe-89b87390accd', 'Interesting!', '2025-06-07 18:58:59.428892+00', '2025-06-07 18:58:59.428892+00', NULL);
INSERT INTO public.posts VALUES ('9e050e56-935c-49f3-bf24-b8e538a38b82', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'b438fcc1-9526-4e3c-8bd2-039bc7e71e7b', '9aead199-68d1-4dc2-b42e-af82e4932380', '33f2ef0e-d58d-4b7a-a5fe-89b87390accd', 'This inspired me.', '2025-06-07 19:18:59.428892+00', '2025-06-07 19:18:59.428892+00', NULL);
INSERT INTO public.posts VALUES ('5a6d4b65-2d68-4114-9fbb-15b90fb044f9', '6e137217-71a9-439b-a579-4bcd1a729af9', 'b438fcc1-9526-4e3c-8bd2-039bc7e71e7b', '9aead199-68d1-4dc2-b42e-af82e4932380', '33f2ef0e-d58d-4b7a-a5fe-89b87390accd', 'Love this!', '2025-06-07 17:50:59.428892+00', '2025-06-07 17:50:59.428892+00', NULL);
INSERT INTO public.posts VALUES ('2afc5e20-e109-447b-88a0-d7d34b42054e', '79758f94-70d8-4a11-b05f-793fcb57e093', 'c84f3282-901c-4433-82ee-42fc3f4ddf3a', '429eac06-eca0-49fe-9cc0-6c731b297206', '3413acc6-1419-4111-9328-0485f0ed1ef8', 'Interesting!', '2025-06-07 20:11:15.233012+00', '2025-06-07 20:11:15.233012+00', NULL);
INSERT INTO public.posts VALUES ('e5f41cd8-9f8a-4aa9-8500-447c2839d9f0', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'c84f3282-901c-4433-82ee-42fc3f4ddf3a', '429eac06-eca0-49fe-9cc0-6c731b297206', '3413acc6-1419-4111-9328-0485f0ed1ef8', 'Very cool!', '2025-06-07 21:03:15.233012+00', '2025-06-07 21:03:15.233012+00', NULL);
INSERT INTO public.posts VALUES ('5ab71ccf-eacc-49f6-9702-4cf9842e7030', '6e137217-71a9-439b-a579-4bcd1a729af9', 'c84f3282-901c-4433-82ee-42fc3f4ddf3a', '429eac06-eca0-49fe-9cc0-6c731b297206', '3413acc6-1419-4111-9328-0485f0ed1ef8', 'Following your example.', '2025-06-07 20:04:15.233012+00', '2025-06-07 20:04:15.233012+00', NULL);
INSERT INTO public.posts VALUES ('4e2d9dc5-fe87-4fd8-ba99-9e2a272dbaab', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'c84f3282-901c-4433-82ee-42fc3f4ddf3a', '429eac06-eca0-49fe-9cc0-6c731b297206', '3413acc6-1419-4111-9328-0485f0ed1ef8', 'Keep going!', '2025-06-07 20:05:15.233012+00', '2025-06-07 20:05:15.233012+00', NULL);
INSERT INTO public.posts VALUES ('5eefa106-2377-4960-8968-9908e9eae43b', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'c84f3282-901c-4433-82ee-42fc3f4ddf3a', '429eac06-eca0-49fe-9cc0-6c731b297206', '3413acc6-1419-4111-9328-0485f0ed1ef8', 'This inspired me.', '2025-06-07 21:41:15.233012+00', '2025-06-07 21:41:15.233012+00', NULL);
INSERT INTO public.posts VALUES ('2ca85ba4-bc1f-4e45-9f67-c2dcf45675b3', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'c84f3282-901c-4433-82ee-42fc3f4ddf3a', '429eac06-eca0-49fe-9cc0-6c731b297206', '3413acc6-1419-4111-9328-0485f0ed1ef8', 'Love this!', '2025-06-07 20:00:15.233012+00', '2025-06-07 20:00:15.233012+00', NULL);
INSERT INTO public.posts VALUES ('c6f33d3f-d649-404f-b43c-8bb0aa6ba56c', '6e137217-71a9-439b-a579-4bcd1a729af9', '70679b66-1611-48c0-9bcb-0558b6d28965', 'ba72e44e-efb7-4331-a79f-a36933bc4f35', '3487baed-69d9-4a07-ba09-ea0e9c2699a4', 'This inspired me.', '2025-06-07 17:33:21.42636+00', '2025-06-07 17:33:21.42636+00', NULL);
INSERT INTO public.posts VALUES ('2d5aad28-f2f8-4ae6-9fd5-77ed71d970af', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '70679b66-1611-48c0-9bcb-0558b6d28965', 'ba72e44e-efb7-4331-a79f-a36933bc4f35', '3487baed-69d9-4a07-ba09-ea0e9c2699a4', 'Motivating stuff!', '2025-06-07 16:56:21.42636+00', '2025-06-07 16:56:21.42636+00', NULL);
INSERT INTO public.posts VALUES ('3578dc9f-ff29-47d0-979f-6111180c199f', '3b0ebec9-3805-490b-987f-4c6e7286319e', '70679b66-1611-48c0-9bcb-0558b6d28965', 'ba72e44e-efb7-4331-a79f-a36933bc4f35', '3487baed-69d9-4a07-ba09-ea0e9c2699a4', 'That’s awesome!', '2025-06-07 17:14:21.42636+00', '2025-06-07 17:14:21.42636+00', NULL);
INSERT INTO public.posts VALUES ('f7e27a88-1f46-444a-8652-d17333a3308b', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '70679b66-1611-48c0-9bcb-0558b6d28965', 'ba72e44e-efb7-4331-a79f-a36933bc4f35', '3487baed-69d9-4a07-ba09-ea0e9c2699a4', 'Keep going!', '2025-06-07 18:52:21.42636+00', '2025-06-07 18:52:21.42636+00', NULL);
INSERT INTO public.posts VALUES ('b634ccdc-df63-409d-b80c-85cc55dab872', '56e89866-3a90-4f97-8461-34b24f842da4', '70679b66-1611-48c0-9bcb-0558b6d28965', 'ba72e44e-efb7-4331-a79f-a36933bc4f35', '3487baed-69d9-4a07-ba09-ea0e9c2699a4', 'Very cool!', '2025-06-07 17:02:21.42636+00', '2025-06-07 17:02:21.42636+00', NULL);
INSERT INTO public.posts VALUES ('912ee6b5-66d3-41de-a033-6986f9298c05', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '70679b66-1611-48c0-9bcb-0558b6d28965', 'ba72e44e-efb7-4331-a79f-a36933bc4f35', '3487baed-69d9-4a07-ba09-ea0e9c2699a4', 'Following your example.', '2025-06-07 18:52:21.42636+00', '2025-06-07 18:52:21.42636+00', NULL);
INSERT INTO public.posts VALUES ('54972a89-d4e1-4e74-bc36-f93e2dc31399', '56e89866-3a90-4f97-8461-34b24f842da4', '30a43be9-a3fe-43c9-b536-25cb0ee96c1b', '771b2a91-478a-46de-afc3-b87f15a693fd', '384be8c2-daf3-4bb6-8e8a-69d83285c0df', 'Motivating stuff!', '2025-06-07 20:04:12.771858+00', '2025-06-07 20:04:12.771858+00', NULL);
INSERT INTO public.posts VALUES ('6d1871e5-5fc4-4d20-b04c-7a72d87d67fa', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '30a43be9-a3fe-43c9-b536-25cb0ee96c1b', '771b2a91-478a-46de-afc3-b87f15a693fd', '384be8c2-daf3-4bb6-8e8a-69d83285c0df', 'Interesting!', '2025-06-07 20:33:12.771858+00', '2025-06-07 20:33:12.771858+00', NULL);
INSERT INTO public.posts VALUES ('1be44b59-9d36-4642-a79a-80c7b8f547f6', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '30a43be9-a3fe-43c9-b536-25cb0ee96c1b', '771b2a91-478a-46de-afc3-b87f15a693fd', '384be8c2-daf3-4bb6-8e8a-69d83285c0df', 'Nice progress!', '2025-06-07 20:45:12.771858+00', '2025-06-07 20:45:12.771858+00', NULL);
INSERT INTO public.posts VALUES ('1c03a974-f9a6-469e-9fb5-e40d24c409da', '3b0ebec9-3805-490b-987f-4c6e7286319e', '30a43be9-a3fe-43c9-b536-25cb0ee96c1b', '771b2a91-478a-46de-afc3-b87f15a693fd', '384be8c2-daf3-4bb6-8e8a-69d83285c0df', 'This inspired me.', '2025-06-07 19:55:12.771858+00', '2025-06-07 19:55:12.771858+00', NULL);
INSERT INTO public.posts VALUES ('7e82ebe1-d3b1-4cee-b0b3-aa81450021bc', '79758f94-70d8-4a11-b05f-793fcb57e093', '30a43be9-a3fe-43c9-b536-25cb0ee96c1b', '771b2a91-478a-46de-afc3-b87f15a693fd', '384be8c2-daf3-4bb6-8e8a-69d83285c0df', 'Keep going!', '2025-06-07 22:06:12.771858+00', '2025-06-07 22:06:12.771858+00', NULL);
INSERT INTO public.posts VALUES ('df7c1c7a-e603-4e1c-b3f1-f01d5daedc0c', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '30a43be9-a3fe-43c9-b536-25cb0ee96c1b', '771b2a91-478a-46de-afc3-b87f15a693fd', '384be8c2-daf3-4bb6-8e8a-69d83285c0df', 'Following your example.', '2025-06-07 21:42:12.771858+00', '2025-06-07 21:42:12.771858+00', NULL);
INSERT INTO public.posts VALUES ('d8095bd3-c981-4d71-9c0b-ccdae83eb4d4', '56e89866-3a90-4f97-8461-34b24f842da4', '7e8e87bc-1b4c-44e9-832c-6e29617642d8', '4369d19b-7570-47b9-9f06-f72b2eee78f2', '40123287-2503-4a02-bd89-93f13eea21be', 'I had the same thought.', '2025-06-07 19:33:09.181576+00', '2025-06-07 19:33:09.181576+00', NULL);
INSERT INTO public.posts VALUES ('01ca61c9-d947-4a6a-8a7a-fed592577f42', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '7e8e87bc-1b4c-44e9-832c-6e29617642d8', '4369d19b-7570-47b9-9f06-f72b2eee78f2', '40123287-2503-4a02-bd89-93f13eea21be', 'Keep going!', '2025-06-07 21:18:09.181576+00', '2025-06-07 21:18:09.181576+00', NULL);
INSERT INTO public.posts VALUES ('41ca627b-6b17-41b4-9c5e-58696fe92e87', '3b0ebec9-3805-490b-987f-4c6e7286319e', '7e8e87bc-1b4c-44e9-832c-6e29617642d8', '4369d19b-7570-47b9-9f06-f72b2eee78f2', '40123287-2503-4a02-bd89-93f13eea21be', 'Nice progress!', '2025-06-07 21:59:09.181576+00', '2025-06-07 21:59:09.181576+00', NULL);
INSERT INTO public.posts VALUES ('fa6fb49a-8fa8-44bf-a771-89725bfac79e', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '7e8e87bc-1b4c-44e9-832c-6e29617642d8', '4369d19b-7570-47b9-9f06-f72b2eee78f2', '40123287-2503-4a02-bd89-93f13eea21be', 'Interesting!', '2025-06-07 20:11:09.181576+00', '2025-06-07 20:11:09.181576+00', NULL);
INSERT INTO public.posts VALUES ('2fcb405f-d82f-45ce-9b47-0c3de6255931', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '912c9992-debd-42b8-9df7-28a9c4b7a319', 'd57937da-618e-4c6b-bc5b-e757af98582e', '47d7751c-7bc4-4dd3-8c10-a25918732741', 'Nice progress!', '2025-06-07 20:22:29.731122+00', '2025-06-07 20:22:29.731122+00', NULL);
INSERT INTO public.posts VALUES ('e5a3fc99-59c5-40f1-975c-5c386322fd0a', '3b0ebec9-3805-490b-987f-4c6e7286319e', '912c9992-debd-42b8-9df7-28a9c4b7a319', 'd57937da-618e-4c6b-bc5b-e757af98582e', '47d7751c-7bc4-4dd3-8c10-a25918732741', 'Following your example.', '2025-06-07 21:42:29.731122+00', '2025-06-07 21:42:29.731122+00', NULL);
INSERT INTO public.posts VALUES ('155fe2e5-9135-407e-b06b-2be1773b1f0b', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '912c9992-debd-42b8-9df7-28a9c4b7a319', 'd57937da-618e-4c6b-bc5b-e757af98582e', '47d7751c-7bc4-4dd3-8c10-a25918732741', 'Motivating stuff!', '2025-06-07 20:29:29.731122+00', '2025-06-07 20:29:29.731122+00', NULL);
INSERT INTO public.posts VALUES ('64543caf-ec34-418e-9dea-e2e0cfa6529d', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '912c9992-debd-42b8-9df7-28a9c4b7a319', 'd57937da-618e-4c6b-bc5b-e757af98582e', '47d7751c-7bc4-4dd3-8c10-a25918732741', 'I had the same thought.', '2025-06-07 22:10:29.731122+00', '2025-06-07 22:10:29.731122+00', NULL);
INSERT INTO public.posts VALUES ('5f7ae2ec-6cc9-4c08-8e2b-ad15d187da62', '6e137217-71a9-439b-a579-4bcd1a729af9', '912c9992-debd-42b8-9df7-28a9c4b7a319', 'd57937da-618e-4c6b-bc5b-e757af98582e', '47d7751c-7bc4-4dd3-8c10-a25918732741', 'Love this!', '2025-06-07 21:57:29.731122+00', '2025-06-07 21:57:29.731122+00', NULL);
INSERT INTO public.posts VALUES ('a9910bfe-1f6e-4046-ab16-a6eaa43f47b2', '79758f94-70d8-4a11-b05f-793fcb57e093', '912c9992-debd-42b8-9df7-28a9c4b7a319', 'd57937da-618e-4c6b-bc5b-e757af98582e', '47d7751c-7bc4-4dd3-8c10-a25918732741', 'Keep going!', '2025-06-07 19:59:29.731122+00', '2025-06-07 19:59:29.731122+00', NULL);
INSERT INTO public.posts VALUES ('47eb7be7-e1f9-4c7e-8b1c-25476db1826d', '6e137217-71a9-439b-a579-4bcd1a729af9', '47b0bc29-2ef0-4b57-a3b9-849968031526', '7e1d9562-3ffd-40be-830d-0fd0e005beb7', '4836de17-5868-4800-bdc7-57144e5c3807', 'Very cool!', '2025-06-07 18:32:18.984369+00', '2025-06-07 18:32:18.984369+00', NULL);
INSERT INTO public.posts VALUES ('e5669143-22ab-4685-a830-fe3e2dffc6cd', '3b0ebec9-3805-490b-987f-4c6e7286319e', '47b0bc29-2ef0-4b57-a3b9-849968031526', '7e1d9562-3ffd-40be-830d-0fd0e005beb7', '4836de17-5868-4800-bdc7-57144e5c3807', 'Nice progress!', '2025-06-07 17:47:18.984369+00', '2025-06-07 17:47:18.984369+00', NULL);
INSERT INTO public.posts VALUES ('8d67f139-165c-4c79-85c1-de77035e412f', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '47b0bc29-2ef0-4b57-a3b9-849968031526', '7e1d9562-3ffd-40be-830d-0fd0e005beb7', '4836de17-5868-4800-bdc7-57144e5c3807', 'That’s awesome!', '2025-06-07 18:13:18.984369+00', '2025-06-07 18:13:18.984369+00', NULL);
INSERT INTO public.posts VALUES ('713b8343-6394-4484-a395-6edf16d4aee0', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '47b0bc29-2ef0-4b57-a3b9-849968031526', '7e1d9562-3ffd-40be-830d-0fd0e005beb7', '4836de17-5868-4800-bdc7-57144e5c3807', 'I had the same thought.', '2025-06-07 19:09:18.984369+00', '2025-06-07 19:09:18.984369+00', NULL);
INSERT INTO public.posts VALUES ('1e2ecbd6-2ca9-4e13-be21-13459839f472', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'e5bf481a-040b-4b1b-96c7-787909876292', 'ce0aa660-fdc1-4342-af7d-c55f54b1904e', '4ddee261-c8d1-4980-9a3e-0baef66bbf7c', 'Very cool!', '2025-06-07 21:20:21.341003+00', '2025-06-07 21:20:21.341003+00', NULL);
INSERT INTO public.posts VALUES ('c11ce564-6b53-4aa9-a2c1-2ada621ace85', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'e5bf481a-040b-4b1b-96c7-787909876292', 'ce0aa660-fdc1-4342-af7d-c55f54b1904e', '4ddee261-c8d1-4980-9a3e-0baef66bbf7c', 'This inspired me.', '2025-06-07 18:35:21.341003+00', '2025-06-07 18:35:21.341003+00', NULL);
INSERT INTO public.posts VALUES ('c8b3815c-0fd5-4e50-bdcb-696dc8379931', '56e89866-3a90-4f97-8461-34b24f842da4', 'e5bf481a-040b-4b1b-96c7-787909876292', 'ce0aa660-fdc1-4342-af7d-c55f54b1904e', '4ddee261-c8d1-4980-9a3e-0baef66bbf7c', 'Following your example.', '2025-06-07 19:53:21.341003+00', '2025-06-07 19:53:21.341003+00', NULL);
INSERT INTO public.posts VALUES ('c0c04522-05c7-4cce-a347-60b060ff62d5', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'e5bf481a-040b-4b1b-96c7-787909876292', 'ce0aa660-fdc1-4342-af7d-c55f54b1904e', '4ddee261-c8d1-4980-9a3e-0baef66bbf7c', 'Interesting!', '2025-06-07 20:32:21.341003+00', '2025-06-07 20:32:21.341003+00', NULL);
INSERT INTO public.posts VALUES ('ead07ed9-a362-48e4-a068-36d2f1afe5a8', '79758f94-70d8-4a11-b05f-793fcb57e093', 'e5bf481a-040b-4b1b-96c7-787909876292', 'ce0aa660-fdc1-4342-af7d-c55f54b1904e', '4ddee261-c8d1-4980-9a3e-0baef66bbf7c', 'Nice progress!', '2025-06-07 20:25:21.341003+00', '2025-06-07 20:25:21.341003+00', NULL);
INSERT INTO public.posts VALUES ('f202f8fe-c249-4df0-9d7d-33e08262cc0f', '6e137217-71a9-439b-a579-4bcd1a729af9', 'e5bf481a-040b-4b1b-96c7-787909876292', 'ce0aa660-fdc1-4342-af7d-c55f54b1904e', '4ddee261-c8d1-4980-9a3e-0baef66bbf7c', 'Love this!', '2025-06-07 19:36:21.341003+00', '2025-06-07 19:36:21.341003+00', NULL);
INSERT INTO public.posts VALUES ('0696896e-012f-42c1-a14c-86af5d76d186', '79758f94-70d8-4a11-b05f-793fcb57e093', '9dd80410-bd88-455a-b8fe-a5c2e39edcf0', '62fd5b6a-1833-4551-b403-82eb0dab6724', '4fcdd3cc-c626-4dd4-904b-426f0839c71c', 'Following your example.', '2025-06-07 22:05:19.693846+00', '2025-06-07 22:05:19.693846+00', NULL);
INSERT INTO public.posts VALUES ('907efae7-ddc2-4385-8b05-1b0ea33e4562', '6e137217-71a9-439b-a579-4bcd1a729af9', '9dd80410-bd88-455a-b8fe-a5c2e39edcf0', '62fd5b6a-1833-4551-b403-82eb0dab6724', '4fcdd3cc-c626-4dd4-904b-426f0839c71c', 'Nice progress!', '2025-06-07 21:21:19.693846+00', '2025-06-07 21:21:19.693846+00', NULL);
INSERT INTO public.posts VALUES ('8eabc043-a56e-423d-976d-42bab37ca1f1', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '9dd80410-bd88-455a-b8fe-a5c2e39edcf0', '62fd5b6a-1833-4551-b403-82eb0dab6724', '4fcdd3cc-c626-4dd4-904b-426f0839c71c', 'Very cool!', '2025-06-07 21:56:19.693846+00', '2025-06-07 21:56:19.693846+00', NULL);
INSERT INTO public.posts VALUES ('5258c4ff-d350-43a6-93cd-f990f7b39868', '56e89866-3a90-4f97-8461-34b24f842da4', '9dd80410-bd88-455a-b8fe-a5c2e39edcf0', '62fd5b6a-1833-4551-b403-82eb0dab6724', '4fcdd3cc-c626-4dd4-904b-426f0839c71c', 'That’s awesome!', '2025-06-07 21:01:19.693846+00', '2025-06-07 21:01:19.693846+00', NULL);
INSERT INTO public.posts VALUES ('b2e45e53-ee94-4ecf-86ba-eda7fc7d007a', '3b0ebec9-3805-490b-987f-4c6e7286319e', '9dd80410-bd88-455a-b8fe-a5c2e39edcf0', '62fd5b6a-1833-4551-b403-82eb0dab6724', '4fcdd3cc-c626-4dd4-904b-426f0839c71c', 'This inspired me.', '2025-06-07 20:00:19.693846+00', '2025-06-07 20:00:19.693846+00', NULL);
INSERT INTO public.posts VALUES ('26de32e8-f69e-4636-836d-a4c8744b3f7f', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '9dd80410-bd88-455a-b8fe-a5c2e39edcf0', '62fd5b6a-1833-4551-b403-82eb0dab6724', '4fcdd3cc-c626-4dd4-904b-426f0839c71c', 'Keep going!', '2025-06-07 21:55:19.693846+00', '2025-06-07 21:55:19.693846+00', NULL);
INSERT INTO public.posts VALUES ('b75b498e-f280-48a1-871b-73d9d848caa1', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '9edaf28d-e47c-4bec-aa12-f4d6b78c6edb', 'b377c3ec-dfcf-48b0-9ece-f3898d31ae9c', '53626074-a976-43cb-83b1-b812a80f3244', 'Motivating stuff!', '2025-06-07 22:00:55.864884+00', '2025-06-07 22:00:55.864884+00', NULL);
INSERT INTO public.posts VALUES ('17085ae1-30e9-45f4-8437-144ef8270e62', '6e137217-71a9-439b-a579-4bcd1a729af9', '9edaf28d-e47c-4bec-aa12-f4d6b78c6edb', 'b377c3ec-dfcf-48b0-9ece-f3898d31ae9c', '53626074-a976-43cb-83b1-b812a80f3244', 'I had the same thought.', '2025-06-07 20:24:55.864884+00', '2025-06-07 20:24:55.864884+00', NULL);
INSERT INTO public.posts VALUES ('d9c7d7df-f9ed-4a84-b529-3126449bacb2', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '9edaf28d-e47c-4bec-aa12-f4d6b78c6edb', 'b377c3ec-dfcf-48b0-9ece-f3898d31ae9c', '53626074-a976-43cb-83b1-b812a80f3244', 'Love this!', '2025-06-07 20:05:55.864884+00', '2025-06-07 20:05:55.864884+00', NULL);
INSERT INTO public.posts VALUES ('344ae4b8-ae5b-4e75-ab8b-35ef1492ebb2', '79758f94-70d8-4a11-b05f-793fcb57e093', '9edaf28d-e47c-4bec-aa12-f4d6b78c6edb', 'b377c3ec-dfcf-48b0-9ece-f3898d31ae9c', '53626074-a976-43cb-83b1-b812a80f3244', 'This inspired me.', '2025-06-07 19:23:55.864884+00', '2025-06-07 19:23:55.864884+00', NULL);
INSERT INTO public.posts VALUES ('c20d6c8f-5142-4cd2-8d19-1a38f3c88aab', '3b0ebec9-3805-490b-987f-4c6e7286319e', '9edaf28d-e47c-4bec-aa12-f4d6b78c6edb', 'b377c3ec-dfcf-48b0-9ece-f3898d31ae9c', '53626074-a976-43cb-83b1-b812a80f3244', 'Interesting!', '2025-06-07 20:57:55.864884+00', '2025-06-07 20:57:55.864884+00', NULL);
INSERT INTO public.posts VALUES ('381932f3-5411-42c2-bbf5-a7958bb7c85c', '56e89866-3a90-4f97-8461-34b24f842da4', '9edaf28d-e47c-4bec-aa12-f4d6b78c6edb', 'b377c3ec-dfcf-48b0-9ece-f3898d31ae9c', '53626074-a976-43cb-83b1-b812a80f3244', 'Nice progress!', '2025-06-07 19:55:55.864884+00', '2025-06-07 19:55:55.864884+00', NULL);
INSERT INTO public.posts VALUES ('441ae36e-7b7d-4972-a92e-d24c77994c68', '3b0ebec9-3805-490b-987f-4c6e7286319e', '47b0bc29-2ef0-4b57-a3b9-849968031526', '7e1d9562-3ffd-40be-830d-0fd0e005beb7', '5931007c-9050-44ae-90af-0c9732651ff3', 'Following your example.', '2025-06-07 19:10:31.707751+00', '2025-06-07 19:10:31.707751+00', NULL);
INSERT INTO public.posts VALUES ('06861b72-662d-4e5b-ad9b-78c0bb3ff99e', '6e137217-71a9-439b-a579-4bcd1a729af9', '47b0bc29-2ef0-4b57-a3b9-849968031526', '7e1d9562-3ffd-40be-830d-0fd0e005beb7', '5931007c-9050-44ae-90af-0c9732651ff3', 'I had the same thought.', '2025-06-07 16:36:31.707751+00', '2025-06-07 16:36:31.707751+00', NULL);
INSERT INTO public.posts VALUES ('1c083b5d-c23e-47ec-b1aa-d35771ead679', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '47b0bc29-2ef0-4b57-a3b9-849968031526', '7e1d9562-3ffd-40be-830d-0fd0e005beb7', '5931007c-9050-44ae-90af-0c9732651ff3', 'Motivating stuff!', '2025-06-07 18:40:31.707751+00', '2025-06-07 18:40:31.707751+00', NULL);
INSERT INTO public.posts VALUES ('8c366194-5265-40f3-9d7b-67ca1c9dad27', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '47b0bc29-2ef0-4b57-a3b9-849968031526', '7e1d9562-3ffd-40be-830d-0fd0e005beb7', '5931007c-9050-44ae-90af-0c9732651ff3', 'Keep going!', '2025-06-07 16:56:31.707751+00', '2025-06-07 16:56:31.707751+00', NULL);
INSERT INTO public.posts VALUES ('31d6d01d-721b-4707-8ad7-bb720d4686e7', '79758f94-70d8-4a11-b05f-793fcb57e093', '47b0bc29-2ef0-4b57-a3b9-849968031526', '7e1d9562-3ffd-40be-830d-0fd0e005beb7', '5931007c-9050-44ae-90af-0c9732651ff3', 'Very cool!', '2025-06-07 19:13:31.707751+00', '2025-06-07 19:13:31.707751+00', NULL);
INSERT INTO public.posts VALUES ('b3f14d80-5b01-4b35-9607-4410855116bd', '56e89866-3a90-4f97-8461-34b24f842da4', '47b0bc29-2ef0-4b57-a3b9-849968031526', '7e1d9562-3ffd-40be-830d-0fd0e005beb7', '5931007c-9050-44ae-90af-0c9732651ff3', 'Love this!', '2025-06-07 17:28:31.707751+00', '2025-06-07 17:28:31.707751+00', NULL);
INSERT INTO public.posts VALUES ('c1271cb4-b7b7-4266-8988-ed98804803ac', '79758f94-70d8-4a11-b05f-793fcb57e093', 'ac557757-04c4-428d-a616-655ecda8b969', 'adbd6df3-b447-4970-8d81-bb54b15310b5', '596dd13f-30c4-402d-9eb4-5e041370d12d', 'Love this!', '2025-06-07 22:18:41.513687+00', '2025-06-07 22:18:41.513687+00', NULL);
INSERT INTO public.posts VALUES ('e3fe0090-acd9-4cf1-8bc2-1f5355776d24', '56e89866-3a90-4f97-8461-34b24f842da4', 'ac557757-04c4-428d-a616-655ecda8b969', 'adbd6df3-b447-4970-8d81-bb54b15310b5', '596dd13f-30c4-402d-9eb4-5e041370d12d', 'I had the same thought.', '2025-06-07 21:39:41.513687+00', '2025-06-07 21:39:41.513687+00', NULL);
INSERT INTO public.posts VALUES ('7fd8646c-087d-4a95-8845-d0ef7d16d05e', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'ac557757-04c4-428d-a616-655ecda8b969', 'adbd6df3-b447-4970-8d81-bb54b15310b5', '596dd13f-30c4-402d-9eb4-5e041370d12d', 'Interesting!', '2025-06-07 21:06:41.513687+00', '2025-06-07 21:06:41.513687+00', NULL);
INSERT INTO public.posts VALUES ('6a049793-7d93-440d-9aaa-5f6b393d1af3', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'ac557757-04c4-428d-a616-655ecda8b969', 'adbd6df3-b447-4970-8d81-bb54b15310b5', '596dd13f-30c4-402d-9eb4-5e041370d12d', 'That’s awesome!', '2025-06-07 20:19:41.513687+00', '2025-06-07 20:19:41.513687+00', NULL);
INSERT INTO public.posts VALUES ('7fb3653a-74bf-41ae-9668-a311ed0e8727', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '9edaf28d-e47c-4bec-aa12-f4d6b78c6edb', '78d799bd-4a56-4538-bd57-60deda16ded5', '60333cb2-3eac-4c88-b1ab-187788f3a753', 'This inspired me.', '2025-06-07 16:23:59.127193+00', '2025-06-07 16:23:59.127193+00', NULL);
INSERT INTO public.posts VALUES ('0c2dc9dc-0c82-40ee-a31b-5a6fbd2241c5', '6e137217-71a9-439b-a579-4bcd1a729af9', '9edaf28d-e47c-4bec-aa12-f4d6b78c6edb', '78d799bd-4a56-4538-bd57-60deda16ded5', '60333cb2-3eac-4c88-b1ab-187788f3a753', 'Love this!', '2025-06-07 18:09:59.127193+00', '2025-06-07 18:09:59.127193+00', NULL);
INSERT INTO public.posts VALUES ('e20ccb7a-0234-4986-b162-3c63400bac4a', '79758f94-70d8-4a11-b05f-793fcb57e093', '9edaf28d-e47c-4bec-aa12-f4d6b78c6edb', '78d799bd-4a56-4538-bd57-60deda16ded5', '60333cb2-3eac-4c88-b1ab-187788f3a753', 'Keep going!', '2025-06-07 19:11:59.127193+00', '2025-06-07 19:11:59.127193+00', NULL);
INSERT INTO public.posts VALUES ('0db7da75-9124-4adc-8ec5-38df2ccfebf1', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '9edaf28d-e47c-4bec-aa12-f4d6b78c6edb', '78d799bd-4a56-4538-bd57-60deda16ded5', '60333cb2-3eac-4c88-b1ab-187788f3a753', 'That’s awesome!', '2025-06-07 17:47:59.127193+00', '2025-06-07 17:47:59.127193+00', NULL);
INSERT INTO public.posts VALUES ('4bd8691d-da55-4997-963c-bc6a323b63f7', '3b0ebec9-3805-490b-987f-4c6e7286319e', '9edaf28d-e47c-4bec-aa12-f4d6b78c6edb', '78d799bd-4a56-4538-bd57-60deda16ded5', '60333cb2-3eac-4c88-b1ab-187788f3a753', 'Following your example.', '2025-06-07 18:48:59.127193+00', '2025-06-07 18:48:59.127193+00', NULL);
INSERT INTO public.posts VALUES ('34a41990-5c63-4fe6-b899-e5d6c8540022', '56e89866-3a90-4f97-8461-34b24f842da4', '9edaf28d-e47c-4bec-aa12-f4d6b78c6edb', '78d799bd-4a56-4538-bd57-60deda16ded5', '60333cb2-3eac-4c88-b1ab-187788f3a753', 'Nice progress!', '2025-06-07 18:44:59.127193+00', '2025-06-07 18:44:59.127193+00', NULL);
INSERT INTO public.posts VALUES ('9e0a9d97-e28e-4dc7-9940-a90220d166b3', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '30a43be9-a3fe-43c9-b536-25cb0ee96c1b', 'fe218910-6c0b-423c-bbf6-468e5b86844c', '660c8c08-468b-4a0d-a37d-fe58d7d0b075', 'That’s awesome!', '2025-06-07 17:37:30.512253+00', '2025-06-07 17:37:30.512253+00', NULL);
INSERT INTO public.posts VALUES ('4ff54ab1-4aad-41ae-9100-3575dcb38cbc', '3b0ebec9-3805-490b-987f-4c6e7286319e', '30a43be9-a3fe-43c9-b536-25cb0ee96c1b', 'fe218910-6c0b-423c-bbf6-468e5b86844c', '660c8c08-468b-4a0d-a37d-fe58d7d0b075', 'Following your example.', '2025-06-07 17:10:30.512253+00', '2025-06-07 17:10:30.512253+00', NULL);
INSERT INTO public.posts VALUES ('5fb8b9a8-10a2-4009-9115-c1a58cd867f9', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '30a43be9-a3fe-43c9-b536-25cb0ee96c1b', 'fe218910-6c0b-423c-bbf6-468e5b86844c', '660c8c08-468b-4a0d-a37d-fe58d7d0b075', 'This inspired me.', '2025-06-07 18:17:30.512253+00', '2025-06-07 18:17:30.512253+00', NULL);
INSERT INTO public.posts VALUES ('035b2b29-1f6a-4cc7-9dd2-4acf84fd5ec7', '79758f94-70d8-4a11-b05f-793fcb57e093', '30a43be9-a3fe-43c9-b536-25cb0ee96c1b', 'fe218910-6c0b-423c-bbf6-468e5b86844c', '660c8c08-468b-4a0d-a37d-fe58d7d0b075', 'Very cool!', '2025-06-07 17:05:30.512253+00', '2025-06-07 17:05:30.512253+00', NULL);
INSERT INTO public.posts VALUES ('479f0862-1a8b-4237-b300-9579e7c419dc', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '30a43be9-a3fe-43c9-b536-25cb0ee96c1b', 'fe218910-6c0b-423c-bbf6-468e5b86844c', '660c8c08-468b-4a0d-a37d-fe58d7d0b075', 'Nice progress!', '2025-06-07 18:59:30.512253+00', '2025-06-07 18:59:30.512253+00', NULL);
INSERT INTO public.posts VALUES ('a6337109-af59-49ce-94e3-2d33ac45a239', '56e89866-3a90-4f97-8461-34b24f842da4', '30a43be9-a3fe-43c9-b536-25cb0ee96c1b', 'fe218910-6c0b-423c-bbf6-468e5b86844c', '660c8c08-468b-4a0d-a37d-fe58d7d0b075', 'Keep going!', '2025-06-07 18:12:30.512253+00', '2025-06-07 18:12:30.512253+00', NULL);
INSERT INTO public.posts VALUES ('60678403-fd1a-4293-9717-2ae74a84663d', '56e89866-3a90-4f97-8461-34b24f842da4', '609eba42-64bc-460a-9468-bc011dae591c', '941639f8-f2f9-4b34-b71c-374442355a50', '678fdbee-e90c-445c-949e-e3962b0503ce', 'Following your example.', '2025-06-07 18:34:00.525467+00', '2025-06-07 18:34:00.525467+00', NULL);
INSERT INTO public.posts VALUES ('b4c17e4f-38c2-4cb4-9aa9-69017a3f540d', '3b0ebec9-3805-490b-987f-4c6e7286319e', '609eba42-64bc-460a-9468-bc011dae591c', '941639f8-f2f9-4b34-b71c-374442355a50', '678fdbee-e90c-445c-949e-e3962b0503ce', 'Keep going!', '2025-06-07 20:35:00.525467+00', '2025-06-07 20:35:00.525467+00', NULL);
INSERT INTO public.posts VALUES ('84309536-9568-43c2-b64b-eff9cf570e31', '79758f94-70d8-4a11-b05f-793fcb57e093', '609eba42-64bc-460a-9468-bc011dae591c', '941639f8-f2f9-4b34-b71c-374442355a50', '678fdbee-e90c-445c-949e-e3962b0503ce', 'Very cool!', '2025-06-07 21:06:00.525467+00', '2025-06-07 21:06:00.525467+00', NULL);
INSERT INTO public.posts VALUES ('9c2072ac-0de0-4754-8a17-d08563755407', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '609eba42-64bc-460a-9468-bc011dae591c', '941639f8-f2f9-4b34-b71c-374442355a50', '678fdbee-e90c-445c-949e-e3962b0503ce', 'Love this!', '2025-06-07 20:16:00.525467+00', '2025-06-07 20:16:00.525467+00', NULL);
INSERT INTO public.posts VALUES ('74919479-08e8-416b-a82d-a301e6e2fd70', '6e137217-71a9-439b-a579-4bcd1a729af9', '609eba42-64bc-460a-9468-bc011dae591c', '941639f8-f2f9-4b34-b71c-374442355a50', '678fdbee-e90c-445c-949e-e3962b0503ce', 'This inspired me.', '2025-06-07 19:55:00.525467+00', '2025-06-07 19:55:00.525467+00', NULL);
INSERT INTO public.posts VALUES ('784bc747-bca0-4f0d-b836-3aaf8dabe87f', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '609eba42-64bc-460a-9468-bc011dae591c', '941639f8-f2f9-4b34-b71c-374442355a50', '678fdbee-e90c-445c-949e-e3962b0503ce', 'I had the same thought.', '2025-06-07 21:26:00.525467+00', '2025-06-07 21:26:00.525467+00', NULL);
INSERT INTO public.posts VALUES ('7085241e-d881-4e3d-9f10-a47b911913a4', '56e89866-3a90-4f97-8461-34b24f842da4', 'f3984ee6-5f7b-47b0-a88c-c684e10eae43', '141ea941-bf50-45ec-8387-deefb7e716ed', '6a54c0a1-ec0f-45b9-8620-c5d20ba82a03', 'Nice progress!', '2025-06-07 19:12:33.756394+00', '2025-06-07 19:12:33.756394+00', NULL);
INSERT INTO public.posts VALUES ('28852612-3701-47a1-911b-8f043b090390', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'f3984ee6-5f7b-47b0-a88c-c684e10eae43', '141ea941-bf50-45ec-8387-deefb7e716ed', '6a54c0a1-ec0f-45b9-8620-c5d20ba82a03', 'Love this!', '2025-06-07 21:41:33.756394+00', '2025-06-07 21:41:33.756394+00', NULL);
INSERT INTO public.posts VALUES ('567efc28-d569-41e5-81dc-7e81e2575d9f', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'f3984ee6-5f7b-47b0-a88c-c684e10eae43', '141ea941-bf50-45ec-8387-deefb7e716ed', '6a54c0a1-ec0f-45b9-8620-c5d20ba82a03', 'Following your example.', '2025-06-07 19:43:33.756394+00', '2025-06-07 19:43:33.756394+00', NULL);
INSERT INTO public.posts VALUES ('5eb75915-d81c-4325-a3c9-cb45a4a22b6d', '79758f94-70d8-4a11-b05f-793fcb57e093', 'f3984ee6-5f7b-47b0-a88c-c684e10eae43', '141ea941-bf50-45ec-8387-deefb7e716ed', '6a54c0a1-ec0f-45b9-8620-c5d20ba82a03', 'This inspired me.', '2025-06-07 19:56:33.756394+00', '2025-06-07 19:56:33.756394+00', NULL);
INSERT INTO public.posts VALUES ('8cbec237-f647-4186-a86d-069200988f4e', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'f3984ee6-5f7b-47b0-a88c-c684e10eae43', '141ea941-bf50-45ec-8387-deefb7e716ed', '6a54c0a1-ec0f-45b9-8620-c5d20ba82a03', 'Interesting!', '2025-06-07 19:04:33.756394+00', '2025-06-07 19:04:33.756394+00', NULL);
INSERT INTO public.posts VALUES ('02e38a1e-a1fe-40d4-8085-5bda9f90ba9d', '6e137217-71a9-439b-a579-4bcd1a729af9', 'f3984ee6-5f7b-47b0-a88c-c684e10eae43', '141ea941-bf50-45ec-8387-deefb7e716ed', '6a54c0a1-ec0f-45b9-8620-c5d20ba82a03', 'Very cool!', '2025-06-07 20:54:33.756394+00', '2025-06-07 20:54:33.756394+00', NULL);
INSERT INTO public.posts VALUES ('3cdeb7d8-01ca-43b8-82f3-a871409cb70e', '79758f94-70d8-4a11-b05f-793fcb57e093', 'c4c10327-612e-4a10-a744-853e80c41a49', 'c6dd0c27-420c-4123-81c0-0744d6e3a7a6', '6a82001d-b416-4515-913e-5b6e3f57fbdf', 'Love this!', '2025-06-07 19:22:43.109651+00', '2025-06-07 19:22:43.109651+00', NULL);
INSERT INTO public.posts VALUES ('729851d0-a291-4abb-a053-d4e3f4355486', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'c4c10327-612e-4a10-a744-853e80c41a49', 'c6dd0c27-420c-4123-81c0-0744d6e3a7a6', '6a82001d-b416-4515-913e-5b6e3f57fbdf', 'Nice progress!', '2025-06-07 19:04:43.109651+00', '2025-06-07 19:04:43.109651+00', NULL);
INSERT INTO public.posts VALUES ('77d6b7e6-9bbd-4997-b4e2-aa37cab1acad', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'c4c10327-612e-4a10-a744-853e80c41a49', 'c6dd0c27-420c-4123-81c0-0744d6e3a7a6', '6a82001d-b416-4515-913e-5b6e3f57fbdf', 'That’s awesome!', '2025-06-07 20:44:43.109651+00', '2025-06-07 20:44:43.109651+00', NULL);
INSERT INTO public.posts VALUES ('f9a99865-ec07-41e5-8938-7c4d3d3fd6e3', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'c4c10327-612e-4a10-a744-853e80c41a49', 'c6dd0c27-420c-4123-81c0-0744d6e3a7a6', '6a82001d-b416-4515-913e-5b6e3f57fbdf', 'Very cool!', '2025-06-07 18:55:43.109651+00', '2025-06-07 18:55:43.109651+00', NULL);
INSERT INTO public.posts VALUES ('464f5d1b-2f43-4c4a-9768-01cf5f435232', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'c4c10327-612e-4a10-a744-853e80c41a49', 'c6dd0c27-420c-4123-81c0-0744d6e3a7a6', '6a82001d-b416-4515-913e-5b6e3f57fbdf', 'Following your example.', '2025-06-07 19:51:43.109651+00', '2025-06-07 19:51:43.109651+00', NULL);
INSERT INTO public.posts VALUES ('c594f373-0534-4b12-91de-37701f7d4b35', '56e89866-3a90-4f97-8461-34b24f842da4', 'c4c10327-612e-4a10-a744-853e80c41a49', 'c6dd0c27-420c-4123-81c0-0744d6e3a7a6', '6a82001d-b416-4515-913e-5b6e3f57fbdf', 'Keep going!', '2025-06-07 19:16:43.109651+00', '2025-06-07 19:16:43.109651+00', NULL);
INSERT INTO public.posts VALUES ('9e1a0703-f116-40a7-be3e-71bcf28dc96e', '79758f94-70d8-4a11-b05f-793fcb57e093', 'c84f3282-901c-4433-82ee-42fc3f4ddf3a', '20e7d6e0-aaad-4946-84a9-7be9471fb444', '71a82c6e-6511-4cf5-8a80-099a7ee1f829', 'I had the same thought.', '2025-06-07 19:10:42.213778+00', '2025-06-07 19:10:42.213778+00', NULL);
INSERT INTO public.posts VALUES ('2e102ddf-9ee7-45b3-a1ac-0925ab875e0b', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'c84f3282-901c-4433-82ee-42fc3f4ddf3a', '20e7d6e0-aaad-4946-84a9-7be9471fb444', '71a82c6e-6511-4cf5-8a80-099a7ee1f829', 'This inspired me.', '2025-06-07 17:26:42.213778+00', '2025-06-07 17:26:42.213778+00', NULL);
INSERT INTO public.posts VALUES ('28f63f9c-460d-4575-9938-ad44a0f107ce', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'c84f3282-901c-4433-82ee-42fc3f4ddf3a', '20e7d6e0-aaad-4946-84a9-7be9471fb444', '71a82c6e-6511-4cf5-8a80-099a7ee1f829', 'Interesting!', '2025-06-07 17:02:42.213778+00', '2025-06-07 17:02:42.213778+00', NULL);
INSERT INTO public.posts VALUES ('de3f7caf-6eda-4d58-b94f-edcc35a499e1', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'c84f3282-901c-4433-82ee-42fc3f4ddf3a', '20e7d6e0-aaad-4946-84a9-7be9471fb444', '71a82c6e-6511-4cf5-8a80-099a7ee1f829', 'Keep going!', '2025-06-07 19:34:42.213778+00', '2025-06-07 19:34:42.213778+00', NULL);
INSERT INTO public.posts VALUES ('64488a6f-46e9-497a-b666-2d0692c6efd5', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'c84f3282-901c-4433-82ee-42fc3f4ddf3a', '20e7d6e0-aaad-4946-84a9-7be9471fb444', '71a82c6e-6511-4cf5-8a80-099a7ee1f829', 'Love this!', '2025-06-07 18:46:42.213778+00', '2025-06-07 18:46:42.213778+00', NULL);
INSERT INTO public.posts VALUES ('140aa7e0-6817-442d-945c-4c85eb3a8756', '6e137217-71a9-439b-a579-4bcd1a729af9', 'c84f3282-901c-4433-82ee-42fc3f4ddf3a', '20e7d6e0-aaad-4946-84a9-7be9471fb444', '71a82c6e-6511-4cf5-8a80-099a7ee1f829', 'Motivating stuff!', '2025-06-07 17:09:42.213778+00', '2025-06-07 17:09:42.213778+00', NULL);
INSERT INTO public.posts VALUES ('6b6107d6-0120-4aa9-95b1-41736b740497', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '609eba42-64bc-460a-9468-bc011dae591c', '941639f8-f2f9-4b34-b71c-374442355a50', '741d9f45-74da-4da2-b40e-06b325c5342f', 'Motivating stuff!', '2025-06-07 19:41:38.589553+00', '2025-06-07 19:41:38.589553+00', NULL);
INSERT INTO public.posts VALUES ('001789cc-7d41-483a-a89e-f3e2417b129d', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '609eba42-64bc-460a-9468-bc011dae591c', '941639f8-f2f9-4b34-b71c-374442355a50', '741d9f45-74da-4da2-b40e-06b325c5342f', 'Nice progress!', '2025-06-07 22:22:38.589553+00', '2025-06-07 22:22:38.589553+00', NULL);
INSERT INTO public.posts VALUES ('e4bf8168-0372-4368-8d67-2614792d505b', '3b0ebec9-3805-490b-987f-4c6e7286319e', '609eba42-64bc-460a-9468-bc011dae591c', '941639f8-f2f9-4b34-b71c-374442355a50', '741d9f45-74da-4da2-b40e-06b325c5342f', 'Interesting!', '2025-06-07 22:04:38.589553+00', '2025-06-07 22:04:38.589553+00', NULL);
INSERT INTO public.posts VALUES ('619cbf01-bfdb-49c4-aff5-4129fa8e027c', '56e89866-3a90-4f97-8461-34b24f842da4', '609eba42-64bc-460a-9468-bc011dae591c', '941639f8-f2f9-4b34-b71c-374442355a50', '741d9f45-74da-4da2-b40e-06b325c5342f', 'Following your example.', '2025-06-07 20:12:38.589553+00', '2025-06-07 20:12:38.589553+00', NULL);
INSERT INTO public.posts VALUES ('673d98d4-4ed9-46a2-95b2-4d20c65d21da', '6e137217-71a9-439b-a579-4bcd1a729af9', '609eba42-64bc-460a-9468-bc011dae591c', '941639f8-f2f9-4b34-b71c-374442355a50', '741d9f45-74da-4da2-b40e-06b325c5342f', 'Very cool!', '2025-06-07 21:12:38.589553+00', '2025-06-07 21:12:38.589553+00', NULL);
INSERT INTO public.posts VALUES ('76890084-caab-467d-984e-0914254a4076', '79758f94-70d8-4a11-b05f-793fcb57e093', '609eba42-64bc-460a-9468-bc011dae591c', '941639f8-f2f9-4b34-b71c-374442355a50', '741d9f45-74da-4da2-b40e-06b325c5342f', 'Keep going!', '2025-06-07 20:34:38.589553+00', '2025-06-07 20:34:38.589553+00', NULL);
INSERT INTO public.posts VALUES ('e18b27fe-ee85-485f-91b3-02338caf64c0', '79758f94-70d8-4a11-b05f-793fcb57e093', 'b438fcc1-9526-4e3c-8bd2-039bc7e71e7b', 'a40da453-3b25-4f00-97c3-04c8ac10f93b', '789062db-6670-4972-97f9-21f61cdbd723', 'This inspired me.', '2025-06-07 21:13:25.747028+00', '2025-06-07 21:13:25.747028+00', NULL);
INSERT INTO public.posts VALUES ('2e389f91-ee03-41ae-aecf-ab2980823eef', '6e137217-71a9-439b-a579-4bcd1a729af9', 'b438fcc1-9526-4e3c-8bd2-039bc7e71e7b', 'a40da453-3b25-4f00-97c3-04c8ac10f93b', '789062db-6670-4972-97f9-21f61cdbd723', 'I had the same thought.', '2025-06-07 19:45:25.747028+00', '2025-06-07 19:45:25.747028+00', NULL);
INSERT INTO public.posts VALUES ('dfc4b8f9-62db-40fa-aac7-d499b5b0df19', '56e89866-3a90-4f97-8461-34b24f842da4', 'b438fcc1-9526-4e3c-8bd2-039bc7e71e7b', 'a40da453-3b25-4f00-97c3-04c8ac10f93b', '789062db-6670-4972-97f9-21f61cdbd723', 'Very cool!', '2025-06-07 22:13:25.747028+00', '2025-06-07 22:13:25.747028+00', NULL);
INSERT INTO public.posts VALUES ('58658d10-9553-4f3d-bf88-1ee825c8c90e', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'b438fcc1-9526-4e3c-8bd2-039bc7e71e7b', 'a40da453-3b25-4f00-97c3-04c8ac10f93b', '789062db-6670-4972-97f9-21f61cdbd723', 'Nice progress!', '2025-06-07 22:34:25.747028+00', '2025-06-07 22:34:25.747028+00', NULL);
INSERT INTO public.posts VALUES ('2dc062d5-f259-4251-9a93-53efedbc027b', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'b438fcc1-9526-4e3c-8bd2-039bc7e71e7b', 'a40da453-3b25-4f00-97c3-04c8ac10f93b', '789062db-6670-4972-97f9-21f61cdbd723', 'Following your example.', '2025-06-07 21:38:25.747028+00', '2025-06-07 21:38:25.747028+00', NULL);
INSERT INTO public.posts VALUES ('3aff7ebf-0eea-4fa4-a1f5-ff5ee5799110', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'b438fcc1-9526-4e3c-8bd2-039bc7e71e7b', 'a40da453-3b25-4f00-97c3-04c8ac10f93b', '789062db-6670-4972-97f9-21f61cdbd723', 'That’s awesome!', '2025-06-07 20:25:25.747028+00', '2025-06-07 20:25:25.747028+00', NULL);
INSERT INTO public.posts VALUES ('da5f5e19-7e8e-403a-a8bd-0e3d0f59ac8f', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'ac557757-04c4-428d-a616-655ecda8b969', 'b03ade4c-2749-4068-b191-d2d31129ec5c', '7f9fcdb1-7602-4b70-8626-204ff26dae55', 'I had the same thought.', '2025-06-07 20:35:35.519597+00', '2025-06-07 20:35:35.519597+00', NULL);
INSERT INTO public.posts VALUES ('b64333c9-4982-4650-9790-dae48d6bf9a1', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'ac557757-04c4-428d-a616-655ecda8b969', 'b03ade4c-2749-4068-b191-d2d31129ec5c', '7f9fcdb1-7602-4b70-8626-204ff26dae55', 'This inspired me.', '2025-06-07 21:46:35.519597+00', '2025-06-07 21:46:35.519597+00', NULL);
INSERT INTO public.posts VALUES ('f10081ca-4b25-4cb0-90c6-7628038177d6', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'ac557757-04c4-428d-a616-655ecda8b969', 'b03ade4c-2749-4068-b191-d2d31129ec5c', '7f9fcdb1-7602-4b70-8626-204ff26dae55', 'Love this!', '2025-06-07 21:43:35.519597+00', '2025-06-07 21:43:35.519597+00', NULL);
INSERT INTO public.posts VALUES ('5ffe0723-fa7f-4202-a999-3f6dc9253c6a', '79758f94-70d8-4a11-b05f-793fcb57e093', 'ac557757-04c4-428d-a616-655ecda8b969', 'b03ade4c-2749-4068-b191-d2d31129ec5c', '7f9fcdb1-7602-4b70-8626-204ff26dae55', 'Very cool!', '2025-06-07 20:40:35.519597+00', '2025-06-07 20:40:35.519597+00', NULL);
INSERT INTO public.posts VALUES ('f2d9c789-c808-4b9a-88b0-2b56a02712f1', '56e89866-3a90-4f97-8461-34b24f842da4', 'ac557757-04c4-428d-a616-655ecda8b969', 'b03ade4c-2749-4068-b191-d2d31129ec5c', '7f9fcdb1-7602-4b70-8626-204ff26dae55', 'Interesting!', '2025-06-07 21:16:35.519597+00', '2025-06-07 21:16:35.519597+00', NULL);
INSERT INTO public.posts VALUES ('2289e8ba-9852-49df-baff-3bd5a4f273e6', '3b0ebec9-3805-490b-987f-4c6e7286319e', '70679b66-1611-48c0-9bcb-0558b6d28965', '7dbec861-f7dd-47dc-9695-c3fdd5f60791', '807dc60c-27a5-44e8-a20c-0ee8ce41ede8', 'Keep going!', '2025-06-07 19:13:03.693167+00', '2025-06-07 19:13:03.693167+00', NULL);
INSERT INTO public.posts VALUES ('bf28ca7e-9db6-469c-9a71-0f7a6632ac3e', '56e89866-3a90-4f97-8461-34b24f842da4', '70679b66-1611-48c0-9bcb-0558b6d28965', '7dbec861-f7dd-47dc-9695-c3fdd5f60791', '807dc60c-27a5-44e8-a20c-0ee8ce41ede8', 'I had the same thought.', '2025-06-07 18:56:03.693167+00', '2025-06-07 18:56:03.693167+00', NULL);
INSERT INTO public.posts VALUES ('5f1de2a7-377d-4209-bb2a-013fcd01babf', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '70679b66-1611-48c0-9bcb-0558b6d28965', '7dbec861-f7dd-47dc-9695-c3fdd5f60791', '807dc60c-27a5-44e8-a20c-0ee8ce41ede8', 'That’s awesome!', '2025-06-07 19:11:03.693167+00', '2025-06-07 19:11:03.693167+00', NULL);
INSERT INTO public.posts VALUES ('8b37efb7-a4bc-441f-8a2c-08e9a7e6be4d', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '70679b66-1611-48c0-9bcb-0558b6d28965', '7dbec861-f7dd-47dc-9695-c3fdd5f60791', '807dc60c-27a5-44e8-a20c-0ee8ce41ede8', 'Very cool!', '2025-06-07 20:59:03.693167+00', '2025-06-07 20:59:03.693167+00', NULL);
INSERT INTO public.posts VALUES ('e6c86f4c-7f4f-459e-b43f-d055f3cd3080', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '70679b66-1611-48c0-9bcb-0558b6d28965', '7dbec861-f7dd-47dc-9695-c3fdd5f60791', '807dc60c-27a5-44e8-a20c-0ee8ce41ede8', 'Motivating stuff!', '2025-06-07 18:57:03.693167+00', '2025-06-07 18:57:03.693167+00', NULL);
INSERT INTO public.posts VALUES ('a10989d1-0c81-4269-a226-93276c27da86', '6e137217-71a9-439b-a579-4bcd1a729af9', '70679b66-1611-48c0-9bcb-0558b6d28965', '7dbec861-f7dd-47dc-9695-c3fdd5f60791', '807dc60c-27a5-44e8-a20c-0ee8ce41ede8', 'Following your example.', '2025-06-07 19:04:03.693167+00', '2025-06-07 19:04:03.693167+00', NULL);
INSERT INTO public.posts VALUES ('396c260d-55ee-43a6-95b1-8c04ddf512e9', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '609eba42-64bc-460a-9468-bc011dae591c', '941639f8-f2f9-4b34-b71c-374442355a50', '8249f900-b093-4d7a-b8f8-67be7fb47448', 'Motivating stuff!', '2025-06-07 22:34:21.610526+00', '2025-06-07 22:34:21.610526+00', NULL);
INSERT INTO public.posts VALUES ('30215eac-8efd-4488-b502-7505379ac9c9', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '609eba42-64bc-460a-9468-bc011dae591c', '941639f8-f2f9-4b34-b71c-374442355a50', '8249f900-b093-4d7a-b8f8-67be7fb47448', 'That’s awesome!', '2025-06-07 20:16:21.610526+00', '2025-06-07 20:16:21.610526+00', NULL);
INSERT INTO public.posts VALUES ('9ba932e7-499f-4ef4-8311-42bb2cf48353', '6e137217-71a9-439b-a579-4bcd1a729af9', '609eba42-64bc-460a-9468-bc011dae591c', '941639f8-f2f9-4b34-b71c-374442355a50', '8249f900-b093-4d7a-b8f8-67be7fb47448', 'Keep going!', '2025-06-07 20:32:21.610526+00', '2025-06-07 20:32:21.610526+00', NULL);
INSERT INTO public.posts VALUES ('f176daff-d23e-428b-b044-a73540d1d0da', '56e89866-3a90-4f97-8461-34b24f842da4', 'ac557757-04c4-428d-a616-655ecda8b969', 'adbd6df3-b447-4970-8d81-bb54b15310b5', '824b208f-844a-4b1c-b9e2-ce7e8148f06e', 'Love this!', '2025-06-07 21:32:54.725475+00', '2025-06-07 21:32:54.725475+00', NULL);
INSERT INTO public.posts VALUES ('11014fdd-b696-4c04-94a1-007e4d9e9eac', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'ac557757-04c4-428d-a616-655ecda8b969', 'adbd6df3-b447-4970-8d81-bb54b15310b5', '824b208f-844a-4b1c-b9e2-ce7e8148f06e', 'Motivating stuff!', '2025-06-07 20:24:54.725475+00', '2025-06-07 20:24:54.725475+00', NULL);
INSERT INTO public.posts VALUES ('ec729ed5-81b1-43ba-8a3c-796457cc011d', '79758f94-70d8-4a11-b05f-793fcb57e093', 'ac557757-04c4-428d-a616-655ecda8b969', 'adbd6df3-b447-4970-8d81-bb54b15310b5', '824b208f-844a-4b1c-b9e2-ce7e8148f06e', 'I had the same thought.', '2025-06-07 21:09:54.725475+00', '2025-06-07 21:09:54.725475+00', NULL);
INSERT INTO public.posts VALUES ('7f6b07fe-e4e9-468f-ab18-cfd39cbd4e7a', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'ac557757-04c4-428d-a616-655ecda8b969', 'adbd6df3-b447-4970-8d81-bb54b15310b5', '824b208f-844a-4b1c-b9e2-ce7e8148f06e', 'Keep going!', '2025-06-07 22:59:54.725475+00', '2025-06-07 22:59:54.725475+00', NULL);
INSERT INTO public.posts VALUES ('b3534311-29d8-4fc7-be52-c9aa8b5bf77f', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'ac557757-04c4-428d-a616-655ecda8b969', 'adbd6df3-b447-4970-8d81-bb54b15310b5', '824b208f-844a-4b1c-b9e2-ce7e8148f06e', 'Nice progress!', '2025-06-07 21:33:54.725475+00', '2025-06-07 21:33:54.725475+00', NULL);
INSERT INTO public.posts VALUES ('91381b81-132a-4798-a3f5-4f520de3821b', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'ac557757-04c4-428d-a616-655ecda8b969', 'adbd6df3-b447-4970-8d81-bb54b15310b5', '824b208f-844a-4b1c-b9e2-ce7e8148f06e', 'Interesting!', '2025-06-07 22:13:54.725475+00', '2025-06-07 22:13:54.725475+00', NULL);
INSERT INTO public.posts VALUES ('7203873a-06c2-49b8-8c80-0fdfbd39aed3', '56e89866-3a90-4f97-8461-34b24f842da4', '30a43be9-a3fe-43c9-b536-25cb0ee96c1b', '771b2a91-478a-46de-afc3-b87f15a693fd', '85a48eda-14df-462a-a173-2c1b653afa4e', 'Keep going!', '2025-06-07 21:45:34.013715+00', '2025-06-07 21:45:34.013715+00', NULL);
INSERT INTO public.posts VALUES ('c97fe57e-606e-46c9-964e-649c915df5ef', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '30a43be9-a3fe-43c9-b536-25cb0ee96c1b', '771b2a91-478a-46de-afc3-b87f15a693fd', '85a48eda-14df-462a-a173-2c1b653afa4e', 'This inspired me.', '2025-06-07 19:52:34.013715+00', '2025-06-07 19:52:34.013715+00', NULL);
INSERT INTO public.posts VALUES ('66e928fe-5645-418f-96d6-3b875bea535b', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '30a43be9-a3fe-43c9-b536-25cb0ee96c1b', '771b2a91-478a-46de-afc3-b87f15a693fd', '85a48eda-14df-462a-a173-2c1b653afa4e', 'Following your example.', '2025-06-07 22:19:34.013715+00', '2025-06-07 22:19:34.013715+00', NULL);
INSERT INTO public.posts VALUES ('1edf0c7f-a1c7-40b4-83db-bc122b45a4c5', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '30a43be9-a3fe-43c9-b536-25cb0ee96c1b', '771b2a91-478a-46de-afc3-b87f15a693fd', '85a48eda-14df-462a-a173-2c1b653afa4e', 'Love this!', '2025-06-07 20:24:34.013715+00', '2025-06-07 20:24:34.013715+00', NULL);
INSERT INTO public.posts VALUES ('8741a70a-51ae-4458-8721-bc19e9da804d', '79758f94-70d8-4a11-b05f-793fcb57e093', '30a43be9-a3fe-43c9-b536-25cb0ee96c1b', '771b2a91-478a-46de-afc3-b87f15a693fd', '85a48eda-14df-462a-a173-2c1b653afa4e', 'Motivating stuff!', '2025-06-07 20:52:34.013715+00', '2025-06-07 20:52:34.013715+00', NULL);
INSERT INTO public.posts VALUES ('3ef42ed3-c898-4610-baea-b0541bd29dce', '3b0ebec9-3805-490b-987f-4c6e7286319e', '30a43be9-a3fe-43c9-b536-25cb0ee96c1b', '771b2a91-478a-46de-afc3-b87f15a693fd', '85a48eda-14df-462a-a173-2c1b653afa4e', 'Nice progress!', '2025-06-07 20:13:34.013715+00', '2025-06-07 20:13:34.013715+00', NULL);
INSERT INTO public.posts VALUES ('dd7f6fb9-1175-41ea-b2f0-98daf3ca1e71', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '9dd80410-bd88-455a-b8fe-a5c2e39edcf0', 'be99dca9-40d3-475d-8ca1-a0f599df7795', '866d6d95-29d3-41e2-adb7-a01b09fd5616', 'Interesting!', '2025-06-07 20:48:29.136373+00', '2025-06-07 20:48:29.136373+00', NULL);
INSERT INTO public.posts VALUES ('06c32191-b799-4e87-b42a-ef8f3c91fea3', '6e137217-71a9-439b-a579-4bcd1a729af9', '9dd80410-bd88-455a-b8fe-a5c2e39edcf0', 'be99dca9-40d3-475d-8ca1-a0f599df7795', '866d6d95-29d3-41e2-adb7-a01b09fd5616', 'Nice progress!', '2025-06-07 19:25:29.136373+00', '2025-06-07 19:25:29.136373+00', NULL);
INSERT INTO public.posts VALUES ('9081c528-ae7b-4e77-b636-b9b289aa7045', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '9dd80410-bd88-455a-b8fe-a5c2e39edcf0', 'be99dca9-40d3-475d-8ca1-a0f599df7795', '866d6d95-29d3-41e2-adb7-a01b09fd5616', 'Following your example.', '2025-06-07 21:39:29.136373+00', '2025-06-07 21:39:29.136373+00', NULL);
INSERT INTO public.posts VALUES ('8cf1a3cc-8274-4604-a0ab-7972822c664d', '3b0ebec9-3805-490b-987f-4c6e7286319e', '9dd80410-bd88-455a-b8fe-a5c2e39edcf0', 'be99dca9-40d3-475d-8ca1-a0f599df7795', '866d6d95-29d3-41e2-adb7-a01b09fd5616', 'Love this!', '2025-06-07 19:49:29.136373+00', '2025-06-07 19:49:29.136373+00', NULL);
INSERT INTO public.posts VALUES ('10bf733b-c904-4062-8fc5-244db7ca484b', '56e89866-3a90-4f97-8461-34b24f842da4', '9dd80410-bd88-455a-b8fe-a5c2e39edcf0', 'be99dca9-40d3-475d-8ca1-a0f599df7795', '866d6d95-29d3-41e2-adb7-a01b09fd5616', 'I had the same thought.', '2025-06-07 20:35:29.136373+00', '2025-06-07 20:35:29.136373+00', NULL);
INSERT INTO public.posts VALUES ('b23e5e0b-0754-4f6d-8a57-ade3c7ffda18', '79758f94-70d8-4a11-b05f-793fcb57e093', '9dd80410-bd88-455a-b8fe-a5c2e39edcf0', 'be99dca9-40d3-475d-8ca1-a0f599df7795', '866d6d95-29d3-41e2-adb7-a01b09fd5616', 'Keep going!', '2025-06-07 19:34:29.136373+00', '2025-06-07 19:34:29.136373+00', NULL);
INSERT INTO public.posts VALUES ('526a064b-865f-4edf-9882-0faf914525d1', '79758f94-70d8-4a11-b05f-793fcb57e093', '513f748c-0dc9-4b04-a8f0-f33ea5892151', '4ef71a7d-5c7a-47b4-bead-b825ef583b04', '8b91fce6-8708-47e5-b7c5-9eee8c55b696', 'I had the same thought.', '2025-06-07 20:39:09.668842+00', '2025-06-07 20:39:09.668842+00', NULL);
INSERT INTO public.posts VALUES ('aae6c796-2462-4780-b62b-41ef9ea901a2', '3b0ebec9-3805-490b-987f-4c6e7286319e', '513f748c-0dc9-4b04-a8f0-f33ea5892151', '4ef71a7d-5c7a-47b4-bead-b825ef583b04', '8b91fce6-8708-47e5-b7c5-9eee8c55b696', 'Very cool!', '2025-06-07 21:19:09.668842+00', '2025-06-07 21:19:09.668842+00', NULL);
INSERT INTO public.posts VALUES ('fabedf04-e966-4dcd-8122-0e884fc7e115', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '513f748c-0dc9-4b04-a8f0-f33ea5892151', '4ef71a7d-5c7a-47b4-bead-b825ef583b04', '8b91fce6-8708-47e5-b7c5-9eee8c55b696', 'Love this!', '2025-06-07 21:46:09.668842+00', '2025-06-07 21:46:09.668842+00', NULL);
INSERT INTO public.posts VALUES ('9287a6ab-0067-454d-bfc8-19abe57cd804', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '513f748c-0dc9-4b04-a8f0-f33ea5892151', '4ef71a7d-5c7a-47b4-bead-b825ef583b04', '8b91fce6-8708-47e5-b7c5-9eee8c55b696', 'That’s awesome!', '2025-06-07 20:28:09.668842+00', '2025-06-07 20:28:09.668842+00', NULL);
INSERT INTO public.posts VALUES ('f141bd54-8d09-49b4-a8cb-a576a0b6812e', '6e137217-71a9-439b-a579-4bcd1a729af9', '513f748c-0dc9-4b04-a8f0-f33ea5892151', '4ef71a7d-5c7a-47b4-bead-b825ef583b04', '8b91fce6-8708-47e5-b7c5-9eee8c55b696', 'Keep going!', '2025-06-07 20:58:09.668842+00', '2025-06-07 20:58:09.668842+00', NULL);
INSERT INTO public.posts VALUES ('a57f056d-d47f-45b8-add6-20e374c34692', '56e89866-3a90-4f97-8461-34b24f842da4', '513f748c-0dc9-4b04-a8f0-f33ea5892151', '4ef71a7d-5c7a-47b4-bead-b825ef583b04', '8b91fce6-8708-47e5-b7c5-9eee8c55b696', 'Nice progress!', '2025-06-07 21:09:09.668842+00', '2025-06-07 21:09:09.668842+00', NULL);
INSERT INTO public.posts VALUES ('6d999417-0af2-4761-84f1-6560932bca5e', '56e89866-3a90-4f97-8461-34b24f842da4', '47b0bc29-2ef0-4b57-a3b9-849968031526', '47bf77f8-3f6d-48f0-bdae-27574c0b902e', '90782b89-1c9c-4c2f-a514-7719331b9c72', 'Nice progress!', '2025-06-07 20:24:42.232831+00', '2025-06-07 20:24:42.232831+00', NULL);
INSERT INTO public.posts VALUES ('96bf5dbf-0732-4678-bb06-12543ea50df8', '3b0ebec9-3805-490b-987f-4c6e7286319e', '47b0bc29-2ef0-4b57-a3b9-849968031526', '47bf77f8-3f6d-48f0-bdae-27574c0b902e', '90782b89-1c9c-4c2f-a514-7719331b9c72', 'Following your example.', '2025-06-07 20:55:42.232831+00', '2025-06-07 20:55:42.232831+00', NULL);
INSERT INTO public.posts VALUES ('62ae29f9-0ce0-4c86-bd1f-4caa23213e9b', '79758f94-70d8-4a11-b05f-793fcb57e093', '47b0bc29-2ef0-4b57-a3b9-849968031526', '47bf77f8-3f6d-48f0-bdae-27574c0b902e', '90782b89-1c9c-4c2f-a514-7719331b9c72', 'Keep going!', '2025-06-07 19:36:42.232831+00', '2025-06-07 19:36:42.232831+00', NULL);
INSERT INTO public.posts VALUES ('2cc2db9b-0a97-412a-9276-71e068e40df7', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '47b0bc29-2ef0-4b57-a3b9-849968031526', '47bf77f8-3f6d-48f0-bdae-27574c0b902e', '90782b89-1c9c-4c2f-a514-7719331b9c72', 'Motivating stuff!', '2025-06-07 20:07:42.232831+00', '2025-06-07 20:07:42.232831+00', NULL);
INSERT INTO public.posts VALUES ('a970daa3-2558-43ea-b493-df1996be7eca', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '47b0bc29-2ef0-4b57-a3b9-849968031526', '47bf77f8-3f6d-48f0-bdae-27574c0b902e', '90782b89-1c9c-4c2f-a514-7719331b9c72', 'Love this!', '2025-06-07 19:53:42.232831+00', '2025-06-07 19:53:42.232831+00', NULL);
INSERT INTO public.posts VALUES ('aac53972-f51a-47b2-8453-0964dd3cdcdf', '6e137217-71a9-439b-a579-4bcd1a729af9', '47b0bc29-2ef0-4b57-a3b9-849968031526', '47bf77f8-3f6d-48f0-bdae-27574c0b902e', '90782b89-1c9c-4c2f-a514-7719331b9c72', 'That’s awesome!', '2025-06-07 20:59:42.232831+00', '2025-06-07 20:59:42.232831+00', NULL);
INSERT INTO public.posts VALUES ('869a5a1e-d001-41df-950e-14f081a3ab5e', '79758f94-70d8-4a11-b05f-793fcb57e093', '66b19273-144f-4dd0-971e-830c18c26b38', '4d8c1e09-2684-49bf-904a-c8bfd973feae', '91d813fb-0bb2-4a9f-98cd-40c4c86de628', 'Love this!', '2025-06-07 21:26:51.57268+00', '2025-06-07 21:26:51.57268+00', NULL);
INSERT INTO public.posts VALUES ('37665743-f593-4269-bf66-02e3ddad9e02', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '66b19273-144f-4dd0-971e-830c18c26b38', '4d8c1e09-2684-49bf-904a-c8bfd973feae', '91d813fb-0bb2-4a9f-98cd-40c4c86de628', 'Interesting!', '2025-06-07 21:01:51.57268+00', '2025-06-07 21:01:51.57268+00', NULL);
INSERT INTO public.posts VALUES ('db5dfd5a-ee03-47df-b14f-f5afa98103c2', '6e137217-71a9-439b-a579-4bcd1a729af9', '66b19273-144f-4dd0-971e-830c18c26b38', '4d8c1e09-2684-49bf-904a-c8bfd973feae', '91d813fb-0bb2-4a9f-98cd-40c4c86de628', 'Keep going!', '2025-06-07 22:24:51.57268+00', '2025-06-07 22:24:51.57268+00', NULL);
INSERT INTO public.posts VALUES ('de69803d-b8d6-432e-a3b4-35d9e356d86c', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '66b19273-144f-4dd0-971e-830c18c26b38', '4d8c1e09-2684-49bf-904a-c8bfd973feae', '91d813fb-0bb2-4a9f-98cd-40c4c86de628', 'Motivating stuff!', '2025-06-07 21:01:51.57268+00', '2025-06-07 21:01:51.57268+00', NULL);
INSERT INTO public.posts VALUES ('df3db3e0-8db0-4dc3-9378-9a5208694742', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '66b19273-144f-4dd0-971e-830c18c26b38', '4d8c1e09-2684-49bf-904a-c8bfd973feae', '91d813fb-0bb2-4a9f-98cd-40c4c86de628', 'This inspired me.', '2025-06-07 22:34:51.57268+00', '2025-06-07 22:34:51.57268+00', NULL);
INSERT INTO public.posts VALUES ('bc137f08-b09f-4069-a065-063d87452b9f', '56e89866-3a90-4f97-8461-34b24f842da4', '66b19273-144f-4dd0-971e-830c18c26b38', '4d8c1e09-2684-49bf-904a-c8bfd973feae', '91d813fb-0bb2-4a9f-98cd-40c4c86de628', 'Nice progress!', '2025-06-07 21:19:51.57268+00', '2025-06-07 21:19:51.57268+00', NULL);
INSERT INTO public.posts VALUES ('c9eb79e1-3a8b-4f50-ba6d-c11a7a1e4105', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '111e36b5-3a31-4079-a5c8-8ead7454723b', '0ea22b31-fb48-49de-9299-446fdf46623d', '93ae9cbf-d56a-4041-b094-ec070ef69045', 'Nice progress!', '2025-06-07 18:58:04.121645+00', '2025-06-07 18:58:04.121645+00', NULL);
INSERT INTO public.posts VALUES ('37f6f051-aa01-4782-a2c6-0b645c4c4795', '6e137217-71a9-439b-a579-4bcd1a729af9', '111e36b5-3a31-4079-a5c8-8ead7454723b', '0ea22b31-fb48-49de-9299-446fdf46623d', '93ae9cbf-d56a-4041-b094-ec070ef69045', 'Love this!', '2025-06-07 21:07:04.121645+00', '2025-06-07 21:07:04.121645+00', NULL);
INSERT INTO public.posts VALUES ('ceb07012-f1e8-4f4a-bcd9-4f0cc2158364', '79758f94-70d8-4a11-b05f-793fcb57e093', '111e36b5-3a31-4079-a5c8-8ead7454723b', '0ea22b31-fb48-49de-9299-446fdf46623d', '93ae9cbf-d56a-4041-b094-ec070ef69045', 'This inspired me.', '2025-06-07 20:30:04.121645+00', '2025-06-07 20:30:04.121645+00', NULL);
INSERT INTO public.posts VALUES ('1ecd3791-435f-4cb1-867b-a6a948eeade9', '56e89866-3a90-4f97-8461-34b24f842da4', '111e36b5-3a31-4079-a5c8-8ead7454723b', '0ea22b31-fb48-49de-9299-446fdf46623d', '93ae9cbf-d56a-4041-b094-ec070ef69045', 'Very cool!', '2025-06-07 21:34:04.121645+00', '2025-06-07 21:34:04.121645+00', NULL);
INSERT INTO public.posts VALUES ('4c396115-5aab-440e-8596-2cd08fe930fb', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '111e36b5-3a31-4079-a5c8-8ead7454723b', '0ea22b31-fb48-49de-9299-446fdf46623d', '93ae9cbf-d56a-4041-b094-ec070ef69045', 'Interesting!', '2025-06-07 20:30:04.121645+00', '2025-06-07 20:30:04.121645+00', NULL);
INSERT INTO public.posts VALUES ('a3857fdb-5579-48e4-835b-623801be4c63', '3b0ebec9-3805-490b-987f-4c6e7286319e', '111e36b5-3a31-4079-a5c8-8ead7454723b', '0ea22b31-fb48-49de-9299-446fdf46623d', '93ae9cbf-d56a-4041-b094-ec070ef69045', 'Keep going!', '2025-06-07 19:20:04.121645+00', '2025-06-07 19:20:04.121645+00', NULL);
INSERT INTO public.posts VALUES ('2d399963-c95f-4b82-84bd-e1fc2163db2f', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'c4c10327-612e-4a10-a744-853e80c41a49', '106936c7-f18f-48c6-81ab-6f931815cced', '9508badc-12e8-4390-aa8f-7763d4dbde96', 'Keep going!', '2025-06-07 17:04:07.003084+00', '2025-06-07 17:04:07.003084+00', NULL);
INSERT INTO public.posts VALUES ('b6f645f8-39e7-433e-8564-90ffa50c5d24', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'c4c10327-612e-4a10-a744-853e80c41a49', '106936c7-f18f-48c6-81ab-6f931815cced', '9508badc-12e8-4390-aa8f-7763d4dbde96', 'Interesting!', '2025-06-07 18:19:07.003084+00', '2025-06-07 18:19:07.003084+00', NULL);
INSERT INTO public.posts VALUES ('8fcd7d61-2048-4444-b863-9176c17e0ce7', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'c4c10327-612e-4a10-a744-853e80c41a49', '106936c7-f18f-48c6-81ab-6f931815cced', '9508badc-12e8-4390-aa8f-7763d4dbde96', 'Nice progress!', '2025-06-07 16:53:07.003084+00', '2025-06-07 16:53:07.003084+00', NULL);
INSERT INTO public.posts VALUES ('fb441b28-ffa0-4223-ae2c-ec5133696d4e', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'c4c10327-612e-4a10-a744-853e80c41a49', '106936c7-f18f-48c6-81ab-6f931815cced', '9508badc-12e8-4390-aa8f-7763d4dbde96', 'Motivating stuff!', '2025-06-07 17:30:07.003084+00', '2025-06-07 17:30:07.003084+00', NULL);
INSERT INTO public.posts VALUES ('0d32eb05-7e86-4ca1-9532-21cfad8bb31a', '79758f94-70d8-4a11-b05f-793fcb57e093', 'c4c10327-612e-4a10-a744-853e80c41a49', '106936c7-f18f-48c6-81ab-6f931815cced', '9508badc-12e8-4390-aa8f-7763d4dbde96', 'Very cool!', '2025-06-07 18:02:07.003084+00', '2025-06-07 18:02:07.003084+00', NULL);
INSERT INTO public.posts VALUES ('5ce35945-569c-44ee-bddc-d42deac28a3f', '56e89866-3a90-4f97-8461-34b24f842da4', 'c4c10327-612e-4a10-a744-853e80c41a49', '106936c7-f18f-48c6-81ab-6f931815cced', '9508badc-12e8-4390-aa8f-7763d4dbde96', 'I had the same thought.', '2025-06-07 17:38:07.003084+00', '2025-06-07 17:38:07.003084+00', NULL);
INSERT INTO public.posts VALUES ('dc1d6ac9-ab2d-4ab0-a7ab-1c910f232897', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '6b7011e6-fb89-4542-8eba-31b219970185', 'c72f0f00-94be-41ae-ade0-954267d26c9d', '97000960-eabc-450b-b074-609480dc17e8', 'This inspired me.', '2025-06-07 20:50:10.005337+00', '2025-06-07 20:50:10.005337+00', NULL);
INSERT INTO public.posts VALUES ('139ac944-44a0-4be6-acd6-9195a477a271', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '6b7011e6-fb89-4542-8eba-31b219970185', 'c72f0f00-94be-41ae-ade0-954267d26c9d', '97000960-eabc-450b-b074-609480dc17e8', 'That’s awesome!', '2025-06-07 21:13:10.005337+00', '2025-06-07 21:13:10.005337+00', NULL);
INSERT INTO public.posts VALUES ('02ffcdbe-b507-4364-a3cc-f2ec47efd02f', '3b0ebec9-3805-490b-987f-4c6e7286319e', '6b7011e6-fb89-4542-8eba-31b219970185', 'c72f0f00-94be-41ae-ade0-954267d26c9d', '97000960-eabc-450b-b074-609480dc17e8', 'Keep going!', '2025-06-07 22:24:10.005337+00', '2025-06-07 22:24:10.005337+00', NULL);
INSERT INTO public.posts VALUES ('2cf85137-bf73-4698-b215-f2ed66f8a82d', '79758f94-70d8-4a11-b05f-793fcb57e093', '513f748c-0dc9-4b04-a8f0-f33ea5892151', '97ff5a74-3b7d-40ce-ab6b-7153f65a446a', '970217d3-a37f-4bea-a894-b26a21d7845f', 'Keep going!', '2025-06-07 20:17:00.609925+00', '2025-06-07 20:17:00.609925+00', NULL);
INSERT INTO public.posts VALUES ('4b6e72ff-fac5-48fd-b857-81ce1210461f', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '513f748c-0dc9-4b04-a8f0-f33ea5892151', '97ff5a74-3b7d-40ce-ab6b-7153f65a446a', '970217d3-a37f-4bea-a894-b26a21d7845f', 'I had the same thought.', '2025-06-07 20:24:00.609925+00', '2025-06-07 20:24:00.609925+00', NULL);
INSERT INTO public.posts VALUES ('c054d7ca-5f0e-407f-bcf2-9dbc88bc87ee', '6e137217-71a9-439b-a579-4bcd1a729af9', '513f748c-0dc9-4b04-a8f0-f33ea5892151', '97ff5a74-3b7d-40ce-ab6b-7153f65a446a', '970217d3-a37f-4bea-a894-b26a21d7845f', 'Nice progress!', '2025-06-07 20:04:00.609925+00', '2025-06-07 20:04:00.609925+00', NULL);
INSERT INTO public.posts VALUES ('b273066d-085f-45fc-a173-7d8039a7b125', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '513f748c-0dc9-4b04-a8f0-f33ea5892151', '97ff5a74-3b7d-40ce-ab6b-7153f65a446a', '970217d3-a37f-4bea-a894-b26a21d7845f', 'This inspired me.', '2025-06-07 20:23:00.609925+00', '2025-06-07 20:23:00.609925+00', NULL);
INSERT INTO public.posts VALUES ('7080a94c-55d0-482d-8b3e-24a79690add6', '3b0ebec9-3805-490b-987f-4c6e7286319e', '513f748c-0dc9-4b04-a8f0-f33ea5892151', '97ff5a74-3b7d-40ce-ab6b-7153f65a446a', '970217d3-a37f-4bea-a894-b26a21d7845f', 'Love this!', '2025-06-07 20:33:00.609925+00', '2025-06-07 20:33:00.609925+00', NULL);
INSERT INTO public.posts VALUES ('8bab410d-d621-4a78-b204-d049246e3f94', '56e89866-3a90-4f97-8461-34b24f842da4', '513f748c-0dc9-4b04-a8f0-f33ea5892151', '97ff5a74-3b7d-40ce-ab6b-7153f65a446a', '970217d3-a37f-4bea-a894-b26a21d7845f', 'Interesting!', '2025-06-07 20:59:00.609925+00', '2025-06-07 20:59:00.609925+00', NULL);
INSERT INTO public.posts VALUES ('33441010-3154-4d51-aa24-f9280cc6f762', '3b0ebec9-3805-490b-987f-4c6e7286319e', '47b0bc29-2ef0-4b57-a3b9-849968031526', 'cab6f8f2-e842-45ee-8b47-75ee8c6a8c73', '9a0f1426-9829-4c07-89b4-288a6e2d2a33', 'Love this!', '2025-06-07 20:52:36.719493+00', '2025-06-07 20:52:36.719493+00', NULL);
INSERT INTO public.posts VALUES ('71ac4a96-8b6c-4a24-8363-93c0a463dfee', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '47b0bc29-2ef0-4b57-a3b9-849968031526', 'cab6f8f2-e842-45ee-8b47-75ee8c6a8c73', '9a0f1426-9829-4c07-89b4-288a6e2d2a33', 'I had the same thought.', '2025-06-07 18:35:36.719493+00', '2025-06-07 18:35:36.719493+00', NULL);
INSERT INTO public.posts VALUES ('405558b4-2e46-4698-b5a7-da3c912ddc97', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '47b0bc29-2ef0-4b57-a3b9-849968031526', 'cab6f8f2-e842-45ee-8b47-75ee8c6a8c73', '9a0f1426-9829-4c07-89b4-288a6e2d2a33', 'Nice progress!', '2025-06-07 19:27:36.719493+00', '2025-06-07 19:27:36.719493+00', NULL);
INSERT INTO public.posts VALUES ('d29885fa-bf10-42ac-bccc-342dfcc89b3f', '79758f94-70d8-4a11-b05f-793fcb57e093', '47b0bc29-2ef0-4b57-a3b9-849968031526', 'cab6f8f2-e842-45ee-8b47-75ee8c6a8c73', '9a0f1426-9829-4c07-89b4-288a6e2d2a33', 'Motivating stuff!', '2025-06-07 21:23:36.719493+00', '2025-06-07 21:23:36.719493+00', NULL);
INSERT INTO public.posts VALUES ('414ee4f0-ac80-47fd-991e-7ef47bc1b7c9', '6e137217-71a9-439b-a579-4bcd1a729af9', '47b0bc29-2ef0-4b57-a3b9-849968031526', 'cab6f8f2-e842-45ee-8b47-75ee8c6a8c73', '9a0f1426-9829-4c07-89b4-288a6e2d2a33', 'Interesting!', '2025-06-07 18:51:36.719493+00', '2025-06-07 18:51:36.719493+00', NULL);
INSERT INTO public.posts VALUES ('3addff12-5ebd-434c-bf01-5256c18475bb', '56e89866-3a90-4f97-8461-34b24f842da4', '47b0bc29-2ef0-4b57-a3b9-849968031526', 'cab6f8f2-e842-45ee-8b47-75ee8c6a8c73', '9a0f1426-9829-4c07-89b4-288a6e2d2a33', 'Following your example.', '2025-06-07 19:35:36.719493+00', '2025-06-07 19:35:36.719493+00', NULL);
INSERT INTO public.posts VALUES ('37647def-b51c-445b-89d3-e4b02c1dfe82', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'b438fcc1-9526-4e3c-8bd2-039bc7e71e7b', 'a40da453-3b25-4f00-97c3-04c8ac10f93b', 'a15c1a09-2a18-4cc2-9915-6c3aa7007455', 'This inspired me.', '2025-06-07 18:54:01.949929+00', '2025-06-07 18:54:01.949929+00', NULL);
INSERT INTO public.posts VALUES ('760dfc9b-b2e2-42b0-8f16-a3c53ed699b7', '6e137217-71a9-439b-a579-4bcd1a729af9', 'b438fcc1-9526-4e3c-8bd2-039bc7e71e7b', 'a40da453-3b25-4f00-97c3-04c8ac10f93b', 'a15c1a09-2a18-4cc2-9915-6c3aa7007455', 'Love this!', '2025-06-07 18:22:01.949929+00', '2025-06-07 18:22:01.949929+00', NULL);
INSERT INTO public.posts VALUES ('e13f421e-263a-42a1-913a-e99561a4962c', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'b438fcc1-9526-4e3c-8bd2-039bc7e71e7b', 'a40da453-3b25-4f00-97c3-04c8ac10f93b', 'a15c1a09-2a18-4cc2-9915-6c3aa7007455', 'I had the same thought.', '2025-06-07 17:50:01.949929+00', '2025-06-07 17:50:01.949929+00', NULL);
INSERT INTO public.posts VALUES ('af2ece0b-ceb2-475f-a6b7-3c638f4afbd2', '56e89866-3a90-4f97-8461-34b24f842da4', 'b438fcc1-9526-4e3c-8bd2-039bc7e71e7b', 'a40da453-3b25-4f00-97c3-04c8ac10f93b', 'a15c1a09-2a18-4cc2-9915-6c3aa7007455', 'Nice progress!', '2025-06-07 18:22:01.949929+00', '2025-06-07 18:22:01.949929+00', NULL);
INSERT INTO public.posts VALUES ('16491e43-d592-407a-b155-a52c4dc60bdb', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'b438fcc1-9526-4e3c-8bd2-039bc7e71e7b', 'a40da453-3b25-4f00-97c3-04c8ac10f93b', 'a15c1a09-2a18-4cc2-9915-6c3aa7007455', 'Motivating stuff!', '2025-06-07 19:37:01.949929+00', '2025-06-07 19:37:01.949929+00', NULL);
INSERT INTO public.posts VALUES ('3cfb760d-2e9c-4637-ae51-8da924867fbd', '79758f94-70d8-4a11-b05f-793fcb57e093', 'b438fcc1-9526-4e3c-8bd2-039bc7e71e7b', 'a40da453-3b25-4f00-97c3-04c8ac10f93b', 'a15c1a09-2a18-4cc2-9915-6c3aa7007455', 'Interesting!', '2025-06-07 18:03:01.949929+00', '2025-06-07 18:03:01.949929+00', NULL);
INSERT INTO public.posts VALUES ('837d9a3a-27d1-4136-9057-26619adaa867', '3b0ebec9-3805-490b-987f-4c6e7286319e', '47b0bc29-2ef0-4b57-a3b9-849968031526', 'ae132a3a-a947-45a1-949b-cf40ff3b4560', 'a21898ff-48aa-4eb6-aeab-050acb249a73', 'Very cool!', '2025-06-07 18:19:13.988036+00', '2025-06-07 18:19:13.988036+00', NULL);
INSERT INTO public.posts VALUES ('ce6aab61-09a4-4751-a1b9-f0b53be912ab', '56e89866-3a90-4f97-8461-34b24f842da4', '47b0bc29-2ef0-4b57-a3b9-849968031526', 'ae132a3a-a947-45a1-949b-cf40ff3b4560', 'a21898ff-48aa-4eb6-aeab-050acb249a73', 'Following your example.', '2025-06-07 16:54:13.988036+00', '2025-06-07 16:54:13.988036+00', NULL);
INSERT INTO public.posts VALUES ('3cfb95fe-36e9-474b-a46c-d657dc951f0b', '6e137217-71a9-439b-a579-4bcd1a729af9', '47b0bc29-2ef0-4b57-a3b9-849968031526', 'ae132a3a-a947-45a1-949b-cf40ff3b4560', 'a21898ff-48aa-4eb6-aeab-050acb249a73', 'That’s awesome!', '2025-06-07 18:00:13.988036+00', '2025-06-07 18:00:13.988036+00', NULL);
INSERT INTO public.posts VALUES ('6c1d47e6-a277-4a11-be2a-6589f5919ad1', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '47b0bc29-2ef0-4b57-a3b9-849968031526', 'ae132a3a-a947-45a1-949b-cf40ff3b4560', 'a21898ff-48aa-4eb6-aeab-050acb249a73', 'Motivating stuff!', '2025-06-07 18:57:13.988036+00', '2025-06-07 18:57:13.988036+00', NULL);
INSERT INTO public.posts VALUES ('7ab41c01-f5ea-4d8a-8d12-b4c1b565391c', '79758f94-70d8-4a11-b05f-793fcb57e093', '47b0bc29-2ef0-4b57-a3b9-849968031526', 'ae132a3a-a947-45a1-949b-cf40ff3b4560', 'a21898ff-48aa-4eb6-aeab-050acb249a73', 'Nice progress!', '2025-06-07 18:11:13.988036+00', '2025-06-07 18:11:13.988036+00', NULL);
INSERT INTO public.posts VALUES ('5248e6db-3ae4-420e-81f8-7ffd535ff8f7', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '47b0bc29-2ef0-4b57-a3b9-849968031526', 'ae132a3a-a947-45a1-949b-cf40ff3b4560', 'a21898ff-48aa-4eb6-aeab-050acb249a73', 'Love this!', '2025-06-07 18:55:13.988036+00', '2025-06-07 18:55:13.988036+00', NULL);
INSERT INTO public.posts VALUES ('33ce3ea5-c21f-4243-af13-1283062aa9a6', '56e89866-3a90-4f97-8461-34b24f842da4', 'ac557757-04c4-428d-a616-655ecda8b969', 'b03ade4c-2749-4068-b191-d2d31129ec5c', 'a3624d47-236a-4b55-be16-d717bfcd8384', 'Very cool!', '2025-06-07 16:56:49.831369+00', '2025-06-07 16:56:49.831369+00', NULL);
INSERT INTO public.posts VALUES ('4c235b97-4966-447d-852c-2c8871c78cef', '79758f94-70d8-4a11-b05f-793fcb57e093', 'ac557757-04c4-428d-a616-655ecda8b969', 'b03ade4c-2749-4068-b191-d2d31129ec5c', 'a3624d47-236a-4b55-be16-d717bfcd8384', 'Love this!', '2025-06-07 17:50:49.831369+00', '2025-06-07 17:50:49.831369+00', NULL);
INSERT INTO public.posts VALUES ('43a33873-7148-4843-b5e0-77bfbbbf4e99', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'ac557757-04c4-428d-a616-655ecda8b969', 'b03ade4c-2749-4068-b191-d2d31129ec5c', 'a3624d47-236a-4b55-be16-d717bfcd8384', 'Interesting!', '2025-06-07 18:15:49.831369+00', '2025-06-07 18:15:49.831369+00', NULL);
INSERT INTO public.posts VALUES ('c5e46978-06c4-4a5f-851b-317b59132801', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'ac557757-04c4-428d-a616-655ecda8b969', 'b03ade4c-2749-4068-b191-d2d31129ec5c', 'a3624d47-236a-4b55-be16-d717bfcd8384', 'Motivating stuff!', '2025-06-07 17:21:49.831369+00', '2025-06-07 17:21:49.831369+00', NULL);
INSERT INTO public.posts VALUES ('f42e2934-5f80-4682-84b7-280ad298c9c8', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'ac557757-04c4-428d-a616-655ecda8b969', 'b03ade4c-2749-4068-b191-d2d31129ec5c', 'a3624d47-236a-4b55-be16-d717bfcd8384', 'Nice progress!', '2025-06-07 19:47:49.831369+00', '2025-06-07 19:47:49.831369+00', NULL);
INSERT INTO public.posts VALUES ('a883030a-47a4-416b-9620-6cba3c7ec93f', '6e137217-71a9-439b-a579-4bcd1a729af9', 'c84f3282-901c-4433-82ee-42fc3f4ddf3a', '0e96f9b2-7005-4c25-9d0a-14bd722bebcb', 'a3f8de46-d113-48a1-b569-a20561d328f5', 'Nice progress!', '2025-06-07 18:40:38.416333+00', '2025-06-07 18:40:38.416333+00', NULL);
INSERT INTO public.posts VALUES ('97c8229d-5296-4946-8b3a-c6f908fcec33', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'c84f3282-901c-4433-82ee-42fc3f4ddf3a', '0e96f9b2-7005-4c25-9d0a-14bd722bebcb', 'a3f8de46-d113-48a1-b569-a20561d328f5', 'Very cool!', '2025-06-07 17:29:38.416333+00', '2025-06-07 17:29:38.416333+00', NULL);
INSERT INTO public.posts VALUES ('468920c6-0d82-4d57-926a-e9531f5dd197', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'c84f3282-901c-4433-82ee-42fc3f4ddf3a', '0e96f9b2-7005-4c25-9d0a-14bd722bebcb', 'a3f8de46-d113-48a1-b569-a20561d328f5', 'This inspired me.', '2025-06-07 16:33:38.416333+00', '2025-06-07 16:33:38.416333+00', NULL);
INSERT INTO public.posts VALUES ('48430cfb-17b8-4103-b024-7c945b6a30e2', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'c84f3282-901c-4433-82ee-42fc3f4ddf3a', '0e96f9b2-7005-4c25-9d0a-14bd722bebcb', 'a3f8de46-d113-48a1-b569-a20561d328f5', 'I had the same thought.', '2025-06-07 17:48:38.416333+00', '2025-06-07 17:48:38.416333+00', NULL);
INSERT INTO public.posts VALUES ('d62ffb49-e0e2-4fff-a437-b5b50b55f679', '79758f94-70d8-4a11-b05f-793fcb57e093', 'c84f3282-901c-4433-82ee-42fc3f4ddf3a', '0e96f9b2-7005-4c25-9d0a-14bd722bebcb', 'a3f8de46-d113-48a1-b569-a20561d328f5', 'Keep going!', '2025-06-07 19:15:38.416333+00', '2025-06-07 19:15:38.416333+00', NULL);
INSERT INTO public.posts VALUES ('c77c13df-0acd-47f7-8b8c-0a720e0c57cb', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'c84f3282-901c-4433-82ee-42fc3f4ddf3a', '0e96f9b2-7005-4c25-9d0a-14bd722bebcb', 'a3f8de46-d113-48a1-b569-a20561d328f5', 'Following your example.', '2025-06-07 17:32:38.416333+00', '2025-06-07 17:32:38.416333+00', NULL);
INSERT INTO public.posts VALUES ('a5e1fb18-2ada-4565-bb35-e36c8df4d709', '56e89866-3a90-4f97-8461-34b24f842da4', 'e5bf481a-040b-4b1b-96c7-787909876292', '2c31cf98-f564-4537-8a12-a71a705103ea', 'aa09c190-fd10-4554-966e-835aefbafcc5', 'This inspired me.', '2025-06-07 21:41:00.389621+00', '2025-06-07 21:41:00.389621+00', NULL);
INSERT INTO public.posts VALUES ('ad41fca1-6771-4a36-a5cf-259fcfd94f77', '79758f94-70d8-4a11-b05f-793fcb57e093', 'e5bf481a-040b-4b1b-96c7-787909876292', '2c31cf98-f564-4537-8a12-a71a705103ea', 'aa09c190-fd10-4554-966e-835aefbafcc5', 'Love this!', '2025-06-07 21:07:00.389621+00', '2025-06-07 21:07:00.389621+00', NULL);
INSERT INTO public.posts VALUES ('5c359637-fbf9-4137-baa5-acacaa4534bb', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'e5bf481a-040b-4b1b-96c7-787909876292', '2c31cf98-f564-4537-8a12-a71a705103ea', 'aa09c190-fd10-4554-966e-835aefbafcc5', 'Interesting!', '2025-06-07 20:48:00.389621+00', '2025-06-07 20:48:00.389621+00', NULL);
INSERT INTO public.posts VALUES ('973fe308-6360-4976-9022-303229057f4a', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'e5bf481a-040b-4b1b-96c7-787909876292', '2c31cf98-f564-4537-8a12-a71a705103ea', 'aa09c190-fd10-4554-966e-835aefbafcc5', 'Keep going!', '2025-06-07 22:16:00.389621+00', '2025-06-07 22:16:00.389621+00', NULL);
INSERT INTO public.posts VALUES ('b47b90fc-e48a-42e6-b623-ddb8d71c9110', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'e5bf481a-040b-4b1b-96c7-787909876292', '2c31cf98-f564-4537-8a12-a71a705103ea', 'aa09c190-fd10-4554-966e-835aefbafcc5', 'I had the same thought.', '2025-06-07 21:48:00.389621+00', '2025-06-07 21:48:00.389621+00', NULL);
INSERT INTO public.posts VALUES ('7ff384ee-755f-404d-b780-96429e20cad8', '6e137217-71a9-439b-a579-4bcd1a729af9', 'e5bf481a-040b-4b1b-96c7-787909876292', '2c31cf98-f564-4537-8a12-a71a705103ea', 'aa09c190-fd10-4554-966e-835aefbafcc5', 'Nice progress!', '2025-06-07 21:54:00.389621+00', '2025-06-07 21:54:00.389621+00', NULL);
INSERT INTO public.posts VALUES ('d3f760ae-142d-445e-bb0d-c0521e3f012d', '6e137217-71a9-439b-a579-4bcd1a729af9', '70679b66-1611-48c0-9bcb-0558b6d28965', 'ba72e44e-efb7-4331-a79f-a36933bc4f35', 'ab773301-3e88-4c8b-9b75-34d5654e2a2a', 'Love this!', '2025-06-07 19:30:50.73246+00', '2025-06-07 19:30:50.73246+00', NULL);
INSERT INTO public.posts VALUES ('821a5a12-b588-4869-a3e2-159a370b0db4', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '70679b66-1611-48c0-9bcb-0558b6d28965', 'ba72e44e-efb7-4331-a79f-a36933bc4f35', 'ab773301-3e88-4c8b-9b75-34d5654e2a2a', 'That’s awesome!', '2025-06-07 20:16:50.73246+00', '2025-06-07 20:16:50.73246+00', NULL);
INSERT INTO public.posts VALUES ('90feb60a-0aa6-4c6e-8a78-479e9b48a513', '56e89866-3a90-4f97-8461-34b24f842da4', '70679b66-1611-48c0-9bcb-0558b6d28965', 'ba72e44e-efb7-4331-a79f-a36933bc4f35', 'ab773301-3e88-4c8b-9b75-34d5654e2a2a', 'This inspired me.', '2025-06-07 20:41:50.73246+00', '2025-06-07 20:41:50.73246+00', NULL);
INSERT INTO public.posts VALUES ('1fe63ee6-fee4-41aa-ae84-6fae14bda85b', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '70679b66-1611-48c0-9bcb-0558b6d28965', 'ba72e44e-efb7-4331-a79f-a36933bc4f35', 'ab773301-3e88-4c8b-9b75-34d5654e2a2a', 'Interesting!', '2025-06-07 19:36:50.73246+00', '2025-06-07 19:36:50.73246+00', NULL);
INSERT INTO public.posts VALUES ('6275d821-021d-4c1d-940e-f5742994532b', '3b0ebec9-3805-490b-987f-4c6e7286319e', '70679b66-1611-48c0-9bcb-0558b6d28965', 'ba72e44e-efb7-4331-a79f-a36933bc4f35', 'ab773301-3e88-4c8b-9b75-34d5654e2a2a', 'Following your example.', '2025-06-07 20:38:50.73246+00', '2025-06-07 20:38:50.73246+00', NULL);
INSERT INTO public.posts VALUES ('97e2caef-f4e6-4a95-840e-abc090e67515', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '70679b66-1611-48c0-9bcb-0558b6d28965', 'ba72e44e-efb7-4331-a79f-a36933bc4f35', 'ab773301-3e88-4c8b-9b75-34d5654e2a2a', 'Very cool!', '2025-06-07 21:52:50.73246+00', '2025-06-07 21:52:50.73246+00', NULL);
INSERT INTO public.posts VALUES ('7eba5100-8254-4b93-b24a-64d22832f26d', '79758f94-70d8-4a11-b05f-793fcb57e093', '513f748c-0dc9-4b04-a8f0-f33ea5892151', '4ef71a7d-5c7a-47b4-bead-b825ef583b04', 'acc9d130-e2ea-426e-9553-d3ee35589a7f', 'Following your example.', '2025-06-07 20:24:55.774318+00', '2025-06-07 20:24:55.774318+00', NULL);
INSERT INTO public.posts VALUES ('f9399428-eb18-4100-a6ad-e2a971149e5c', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '513f748c-0dc9-4b04-a8f0-f33ea5892151', '4ef71a7d-5c7a-47b4-bead-b825ef583b04', 'acc9d130-e2ea-426e-9553-d3ee35589a7f', 'Keep going!', '2025-06-07 19:53:55.774318+00', '2025-06-07 19:53:55.774318+00', NULL);
INSERT INTO public.posts VALUES ('5f5b3fd7-1360-48a7-be69-6c401fff6871', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '513f748c-0dc9-4b04-a8f0-f33ea5892151', '4ef71a7d-5c7a-47b4-bead-b825ef583b04', 'acc9d130-e2ea-426e-9553-d3ee35589a7f', 'Nice progress!', '2025-06-07 19:27:55.774318+00', '2025-06-07 19:27:55.774318+00', NULL);
INSERT INTO public.posts VALUES ('79176376-3a35-4d5f-be1a-70ff80cccae2', '56e89866-3a90-4f97-8461-34b24f842da4', '513f748c-0dc9-4b04-a8f0-f33ea5892151', '4ef71a7d-5c7a-47b4-bead-b825ef583b04', 'acc9d130-e2ea-426e-9553-d3ee35589a7f', 'Motivating stuff!', '2025-06-07 20:28:55.774318+00', '2025-06-07 20:28:55.774318+00', NULL);
INSERT INTO public.posts VALUES ('6441c744-b8d6-44cf-a49c-7886ddb370e5', '3b0ebec9-3805-490b-987f-4c6e7286319e', '513f748c-0dc9-4b04-a8f0-f33ea5892151', '4ef71a7d-5c7a-47b4-bead-b825ef583b04', 'acc9d130-e2ea-426e-9553-d3ee35589a7f', 'I had the same thought.', '2025-06-07 19:08:55.774318+00', '2025-06-07 19:08:55.774318+00', NULL);
INSERT INTO public.posts VALUES ('2cd8e961-a39b-4b6c-8406-95a091b95193', '6e137217-71a9-439b-a579-4bcd1a729af9', '513f748c-0dc9-4b04-a8f0-f33ea5892151', '4ef71a7d-5c7a-47b4-bead-b825ef583b04', 'acc9d130-e2ea-426e-9553-d3ee35589a7f', 'Love this!', '2025-06-07 20:21:55.774318+00', '2025-06-07 20:21:55.774318+00', NULL);
INSERT INTO public.posts VALUES ('03693cf7-09a9-4e2d-a7b4-2bb061780b33', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '47b0bc29-2ef0-4b57-a3b9-849968031526', 'cab6f8f2-e842-45ee-8b47-75ee8c6a8c73', 'b3621746-9032-46c1-9e9c-072c55b709a7', 'Motivating stuff!', '2025-06-07 19:55:57.420885+00', '2025-06-07 19:55:57.420885+00', NULL);
INSERT INTO public.posts VALUES ('9a983f94-9ffa-4cbb-9cc8-e02e5caf916e', '3b0ebec9-3805-490b-987f-4c6e7286319e', '47b0bc29-2ef0-4b57-a3b9-849968031526', 'cab6f8f2-e842-45ee-8b47-75ee8c6a8c73', 'b3621746-9032-46c1-9e9c-072c55b709a7', 'I had the same thought.', '2025-06-07 19:26:57.420885+00', '2025-06-07 19:26:57.420885+00', NULL);
INSERT INTO public.posts VALUES ('5cfe8503-eece-445f-9e23-2ceede4c4a5d', '79758f94-70d8-4a11-b05f-793fcb57e093', '47b0bc29-2ef0-4b57-a3b9-849968031526', 'cab6f8f2-e842-45ee-8b47-75ee8c6a8c73', 'b3621746-9032-46c1-9e9c-072c55b709a7', 'Following your example.', '2025-06-07 20:58:57.420885+00', '2025-06-07 20:58:57.420885+00', NULL);
INSERT INTO public.posts VALUES ('3817238f-26c2-430c-9443-90b9cc3d0d4c', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '47b0bc29-2ef0-4b57-a3b9-849968031526', 'cab6f8f2-e842-45ee-8b47-75ee8c6a8c73', 'b3621746-9032-46c1-9e9c-072c55b709a7', 'Interesting!', '2025-06-07 22:08:57.420885+00', '2025-06-07 22:08:57.420885+00', NULL);
INSERT INTO public.posts VALUES ('e9da1728-7743-465e-9c2b-3cc400be00c8', '56e89866-3a90-4f97-8461-34b24f842da4', '47b0bc29-2ef0-4b57-a3b9-849968031526', 'cab6f8f2-e842-45ee-8b47-75ee8c6a8c73', 'b3621746-9032-46c1-9e9c-072c55b709a7', 'That’s awesome!', '2025-06-07 21:00:57.420885+00', '2025-06-07 21:00:57.420885+00', NULL);
INSERT INTO public.posts VALUES ('f651d8ba-63d7-4651-99b9-e78650533dca', '6e137217-71a9-439b-a579-4bcd1a729af9', '47b0bc29-2ef0-4b57-a3b9-849968031526', 'cab6f8f2-e842-45ee-8b47-75ee8c6a8c73', 'b3621746-9032-46c1-9e9c-072c55b709a7', 'Nice progress!', '2025-06-07 21:28:57.420885+00', '2025-06-07 21:28:57.420885+00', NULL);
INSERT INTO public.posts VALUES ('5e3529d8-a27d-4351-ad80-228ee4ba8dea', '79758f94-70d8-4a11-b05f-793fcb57e093', '66b19273-144f-4dd0-971e-830c18c26b38', 'd89edf48-a9de-456a-aa44-9e669a382fd5', 'b8ea0106-08f2-45fc-92ea-77a61af34a1d', 'Keep going!', '2025-06-07 21:19:18.582408+00', '2025-06-07 21:19:18.582408+00', NULL);
INSERT INTO public.posts VALUES ('9b4479fb-6ad4-4d0c-b42d-5d9ba962ecc9', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '66b19273-144f-4dd0-971e-830c18c26b38', 'd89edf48-a9de-456a-aa44-9e669a382fd5', 'b8ea0106-08f2-45fc-92ea-77a61af34a1d', 'Following your example.', '2025-06-07 21:17:18.582408+00', '2025-06-07 21:17:18.582408+00', NULL);
INSERT INTO public.posts VALUES ('fd1f52a0-085c-409f-a2e0-a7c0f234f8a9', '6e137217-71a9-439b-a579-4bcd1a729af9', '66b19273-144f-4dd0-971e-830c18c26b38', 'd89edf48-a9de-456a-aa44-9e669a382fd5', 'b8ea0106-08f2-45fc-92ea-77a61af34a1d', 'Love this!', '2025-06-07 19:44:18.582408+00', '2025-06-07 19:44:18.582408+00', NULL);
INSERT INTO public.posts VALUES ('88ad4c04-4d5f-4303-92a8-c6c9906f2c08', '56e89866-3a90-4f97-8461-34b24f842da4', '66b19273-144f-4dd0-971e-830c18c26b38', 'd89edf48-a9de-456a-aa44-9e669a382fd5', 'b8ea0106-08f2-45fc-92ea-77a61af34a1d', 'I had the same thought.', '2025-06-07 21:22:18.582408+00', '2025-06-07 21:22:18.582408+00', NULL);
INSERT INTO public.posts VALUES ('cf6f3dfd-add2-400b-b21f-da360336b27d', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '66b19273-144f-4dd0-971e-830c18c26b38', 'd89edf48-a9de-456a-aa44-9e669a382fd5', 'b8ea0106-08f2-45fc-92ea-77a61af34a1d', 'This inspired me.', '2025-06-07 20:31:18.582408+00', '2025-06-07 20:31:18.582408+00', NULL);
INSERT INTO public.posts VALUES ('e1cdd912-0b2c-4b6e-b800-0a3bf5c6c1a0', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '66b19273-144f-4dd0-971e-830c18c26b38', 'd89edf48-a9de-456a-aa44-9e669a382fd5', 'b8ea0106-08f2-45fc-92ea-77a61af34a1d', 'Motivating stuff!', '2025-06-07 22:04:18.582408+00', '2025-06-07 22:04:18.582408+00', NULL);
INSERT INTO public.posts VALUES ('d3335898-cdfc-4f98-850d-9dcfc67ba74f', '56e89866-3a90-4f97-8461-34b24f842da4', '70679b66-1611-48c0-9bcb-0558b6d28965', '12cccfe8-320f-4992-9d2f-b59ef5ea4f6e', 'c2901639-580e-46f1-a0b9-12e6a970e7d4', 'Love this!', '2025-06-07 21:28:45.225947+00', '2025-06-07 21:28:45.225947+00', NULL);
INSERT INTO public.posts VALUES ('d0d9112f-8261-4a60-b0ae-b3744ca8a532', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '70679b66-1611-48c0-9bcb-0558b6d28965', '12cccfe8-320f-4992-9d2f-b59ef5ea4f6e', 'c2901639-580e-46f1-a0b9-12e6a970e7d4', 'Motivating stuff!', '2025-06-07 21:00:45.225947+00', '2025-06-07 21:00:45.225947+00', NULL);
INSERT INTO public.posts VALUES ('a3a97fb7-1cee-4d9e-a9ae-98b12669ac36', '3b0ebec9-3805-490b-987f-4c6e7286319e', '70679b66-1611-48c0-9bcb-0558b6d28965', '12cccfe8-320f-4992-9d2f-b59ef5ea4f6e', 'c2901639-580e-46f1-a0b9-12e6a970e7d4', 'Nice progress!', '2025-06-07 20:18:45.225947+00', '2025-06-07 20:18:45.225947+00', NULL);
INSERT INTO public.posts VALUES ('c705f2e5-0275-4f4d-b6cb-4f8135b4d9c6', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '70679b66-1611-48c0-9bcb-0558b6d28965', '12cccfe8-320f-4992-9d2f-b59ef5ea4f6e', 'c2901639-580e-46f1-a0b9-12e6a970e7d4', 'That’s awesome!', '2025-06-07 21:28:45.225947+00', '2025-06-07 21:28:45.225947+00', NULL);
INSERT INTO public.posts VALUES ('97196d01-3074-44cb-b802-0aafa97b248c', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '70679b66-1611-48c0-9bcb-0558b6d28965', '12cccfe8-320f-4992-9d2f-b59ef5ea4f6e', 'c2901639-580e-46f1-a0b9-12e6a970e7d4', 'I had the same thought.', '2025-06-07 21:29:45.225947+00', '2025-06-07 21:29:45.225947+00', NULL);
INSERT INTO public.posts VALUES ('62718577-2d36-4940-8ebb-e9a16d02dcf6', '6e137217-71a9-439b-a579-4bcd1a729af9', '70679b66-1611-48c0-9bcb-0558b6d28965', '12cccfe8-320f-4992-9d2f-b59ef5ea4f6e', 'c2901639-580e-46f1-a0b9-12e6a970e7d4', 'Keep going!', '2025-06-07 20:34:45.225947+00', '2025-06-07 20:34:45.225947+00', NULL);
INSERT INTO public.posts VALUES ('83f2ecad-a326-4781-b941-b80da4812d5a', '56e89866-3a90-4f97-8461-34b24f842da4', 'c4c10327-612e-4a10-a744-853e80c41a49', 'd36ac6f3-57a2-4d11-ba34-9674e46c7bec', 'c4e4435b-e210-481c-9603-880796a0b7df', 'Love this!', '2025-06-07 18:42:16.75267+00', '2025-06-07 18:42:16.75267+00', NULL);
INSERT INTO public.posts VALUES ('43684694-5d2c-4aa8-894f-e64394d2afc4', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'c4c10327-612e-4a10-a744-853e80c41a49', 'd36ac6f3-57a2-4d11-ba34-9674e46c7bec', 'c4e4435b-e210-481c-9603-880796a0b7df', 'This inspired me.', '2025-06-07 18:11:16.75267+00', '2025-06-07 18:11:16.75267+00', NULL);
INSERT INTO public.posts VALUES ('2a1d55e1-37d4-4ded-b815-56ed354aea35', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'c4c10327-612e-4a10-a744-853e80c41a49', 'd36ac6f3-57a2-4d11-ba34-9674e46c7bec', 'c4e4435b-e210-481c-9603-880796a0b7df', 'Keep going!', '2025-06-07 18:00:16.75267+00', '2025-06-07 18:00:16.75267+00', NULL);
INSERT INTO public.posts VALUES ('a9803e2b-be79-4719-9e77-6c4c337c77f0', '79758f94-70d8-4a11-b05f-793fcb57e093', 'c4c10327-612e-4a10-a744-853e80c41a49', 'd36ac6f3-57a2-4d11-ba34-9674e46c7bec', 'c4e4435b-e210-481c-9603-880796a0b7df', 'Nice progress!', '2025-06-07 17:31:16.75267+00', '2025-06-07 17:31:16.75267+00', NULL);
INSERT INTO public.posts VALUES ('24c3d169-24eb-4576-82ef-d0c0d98d4c0a', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'c4c10327-612e-4a10-a744-853e80c41a49', 'd36ac6f3-57a2-4d11-ba34-9674e46c7bec', 'c4e4435b-e210-481c-9603-880796a0b7df', 'I had the same thought.', '2025-06-07 19:03:16.75267+00', '2025-06-07 19:03:16.75267+00', NULL);
INSERT INTO public.posts VALUES ('439e902d-724f-4d87-a27a-abdac09cd274', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'c4c10327-612e-4a10-a744-853e80c41a49', 'd36ac6f3-57a2-4d11-ba34-9674e46c7bec', 'c4e4435b-e210-481c-9603-880796a0b7df', 'Motivating stuff!', '2025-06-07 16:43:16.75267+00', '2025-06-07 16:43:16.75267+00', NULL);
INSERT INTO public.posts VALUES ('d2c3e269-9188-481b-af35-148150302af8', '56e89866-3a90-4f97-8461-34b24f842da4', '7e8e87bc-1b4c-44e9-832c-6e29617642d8', '4d943fed-81fb-4660-a2f5-c64779b63a7e', 'c56780c9-07da-467d-8d5b-5236215603af', 'Motivating stuff!', '2025-06-07 17:46:35.740502+00', '2025-06-07 17:46:35.740502+00', NULL);
INSERT INTO public.posts VALUES ('6774310e-3c89-493b-86ec-e26c6d2d509e', '3b0ebec9-3805-490b-987f-4c6e7286319e', '7e8e87bc-1b4c-44e9-832c-6e29617642d8', '4d943fed-81fb-4660-a2f5-c64779b63a7e', 'c56780c9-07da-467d-8d5b-5236215603af', 'Nice progress!', '2025-06-07 17:57:35.740502+00', '2025-06-07 17:57:35.740502+00', NULL);
INSERT INTO public.posts VALUES ('dadff8ef-b0ca-4131-b152-18351716921e', '6e137217-71a9-439b-a579-4bcd1a729af9', '7e8e87bc-1b4c-44e9-832c-6e29617642d8', '4d943fed-81fb-4660-a2f5-c64779b63a7e', 'c56780c9-07da-467d-8d5b-5236215603af', 'Following your example.', '2025-06-07 18:50:35.740502+00', '2025-06-07 18:50:35.740502+00', NULL);
INSERT INTO public.posts VALUES ('f32e6e3e-33ba-4e4e-aa39-61d4c11744ca', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '7e8e87bc-1b4c-44e9-832c-6e29617642d8', '4d943fed-81fb-4660-a2f5-c64779b63a7e', 'c56780c9-07da-467d-8d5b-5236215603af', 'This inspired me.', '2025-06-07 17:01:35.740502+00', '2025-06-07 17:01:35.740502+00', NULL);
INSERT INTO public.posts VALUES ('3729218b-13ea-4106-a24a-55fcb986b4e0', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '7e8e87bc-1b4c-44e9-832c-6e29617642d8', '4d943fed-81fb-4660-a2f5-c64779b63a7e', 'c56780c9-07da-467d-8d5b-5236215603af', 'Love this!', '2025-06-07 18:10:35.740502+00', '2025-06-07 18:10:35.740502+00', NULL);
INSERT INTO public.posts VALUES ('15238ef6-8b8a-43ad-af02-0601cb39aaac', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '7e8e87bc-1b4c-44e9-832c-6e29617642d8', '4d943fed-81fb-4660-a2f5-c64779b63a7e', 'c56780c9-07da-467d-8d5b-5236215603af', 'Very cool!', '2025-06-07 17:23:35.740502+00', '2025-06-07 17:23:35.740502+00', NULL);
INSERT INTO public.posts VALUES ('8919084f-5c89-46b5-835e-d86180215c89', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '9edaf28d-e47c-4bec-aa12-f4d6b78c6edb', 'dfc33a5b-552c-4d12-b43f-abad5ee9a14d', 'cdcaa739-1c4d-422c-abea-b234c55cc559', 'Nice progress!', '2025-06-07 21:51:23.645203+00', '2025-06-07 21:51:23.645203+00', NULL);
INSERT INTO public.posts VALUES ('1ce878b7-d1ed-4e04-80c5-c3f1d7716f05', '56e89866-3a90-4f97-8461-34b24f842da4', '9edaf28d-e47c-4bec-aa12-f4d6b78c6edb', 'dfc33a5b-552c-4d12-b43f-abad5ee9a14d', 'cdcaa739-1c4d-422c-abea-b234c55cc559', 'Love this!', '2025-06-07 22:06:23.645203+00', '2025-06-07 22:06:23.645203+00', NULL);
INSERT INTO public.posts VALUES ('f864a9c9-6d24-49f3-b997-41c62e4ad601', '6e137217-71a9-439b-a579-4bcd1a729af9', '9edaf28d-e47c-4bec-aa12-f4d6b78c6edb', 'dfc33a5b-552c-4d12-b43f-abad5ee9a14d', 'cdcaa739-1c4d-422c-abea-b234c55cc559', 'Following your example.', '2025-06-07 19:43:23.645203+00', '2025-06-07 19:43:23.645203+00', NULL);
INSERT INTO public.posts VALUES ('558991ca-e608-447a-b441-e9618f061716', '79758f94-70d8-4a11-b05f-793fcb57e093', '9edaf28d-e47c-4bec-aa12-f4d6b78c6edb', 'dfc33a5b-552c-4d12-b43f-abad5ee9a14d', 'cdcaa739-1c4d-422c-abea-b234c55cc559', 'Interesting!', '2025-06-07 20:21:23.645203+00', '2025-06-07 20:21:23.645203+00', NULL);
INSERT INTO public.posts VALUES ('31508dfe-6efd-41e7-841c-a340627021a0', '3b0ebec9-3805-490b-987f-4c6e7286319e', '9edaf28d-e47c-4bec-aa12-f4d6b78c6edb', 'dfc33a5b-552c-4d12-b43f-abad5ee9a14d', 'cdcaa739-1c4d-422c-abea-b234c55cc559', 'This inspired me.', '2025-06-07 21:48:23.645203+00', '2025-06-07 21:48:23.645203+00', NULL);
INSERT INTO public.posts VALUES ('afb7fe88-1b3d-49e9-838c-1db13039a425', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '9edaf28d-e47c-4bec-aa12-f4d6b78c6edb', 'dfc33a5b-552c-4d12-b43f-abad5ee9a14d', 'cdcaa739-1c4d-422c-abea-b234c55cc559', 'Keep going!', '2025-06-07 21:45:23.645203+00', '2025-06-07 21:45:23.645203+00', NULL);
INSERT INTO public.posts VALUES ('1e103266-d7d5-432d-b00d-15646b7edb35', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'a9528f6b-3895-4cab-af6c-62a93ef1b83e', 'd81f001e-91a4-4a34-aeb8-90544b4c328e', 'cfd2ccd0-49d2-4a73-8f13-14348a779f6f', 'This inspired me.', '2025-06-07 18:41:11.591281+00', '2025-06-07 18:41:11.591281+00', NULL);
INSERT INTO public.posts VALUES ('32ef9dbd-be4c-4fb5-9462-29d1e6ca2c8b', '6e137217-71a9-439b-a579-4bcd1a729af9', 'a9528f6b-3895-4cab-af6c-62a93ef1b83e', 'd81f001e-91a4-4a34-aeb8-90544b4c328e', 'cfd2ccd0-49d2-4a73-8f13-14348a779f6f', 'I had the same thought.', '2025-06-07 19:13:11.591281+00', '2025-06-07 19:13:11.591281+00', NULL);
INSERT INTO public.posts VALUES ('c9285fea-771e-4a93-a4b3-46437524eeab', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'a9528f6b-3895-4cab-af6c-62a93ef1b83e', 'd81f001e-91a4-4a34-aeb8-90544b4c328e', 'cfd2ccd0-49d2-4a73-8f13-14348a779f6f', 'Interesting!', '2025-06-07 18:45:11.591281+00', '2025-06-07 18:45:11.591281+00', NULL);
INSERT INTO public.posts VALUES ('ffd656d7-5339-4fd7-a078-60836e7d6fb5', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '47b0bc29-2ef0-4b57-a3b9-849968031526', '7e1d9562-3ffd-40be-830d-0fd0e005beb7', 'd3aef78d-bbf8-4543-bda4-1946e61bcc4b', 'Motivating stuff!', '2025-06-07 20:22:45.085868+00', '2025-06-07 20:22:45.085868+00', NULL);
INSERT INTO public.posts VALUES ('df97dc63-46ab-467a-bf6c-d863db027eb6', '6e137217-71a9-439b-a579-4bcd1a729af9', '47b0bc29-2ef0-4b57-a3b9-849968031526', '7e1d9562-3ffd-40be-830d-0fd0e005beb7', 'd3aef78d-bbf8-4543-bda4-1946e61bcc4b', 'This inspired me.', '2025-06-07 20:04:45.085868+00', '2025-06-07 20:04:45.085868+00', NULL);
INSERT INTO public.posts VALUES ('5726ea2e-7562-4e91-a847-a1a1ce44ef7d', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '47b0bc29-2ef0-4b57-a3b9-849968031526', '7e1d9562-3ffd-40be-830d-0fd0e005beb7', 'd3aef78d-bbf8-4543-bda4-1946e61bcc4b', 'I had the same thought.', '2025-06-07 20:18:45.085868+00', '2025-06-07 20:18:45.085868+00', NULL);
INSERT INTO public.posts VALUES ('295f865a-6a84-4641-8986-1c065cff9952', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '6b7011e6-fb89-4542-8eba-31b219970185', '5501f6a8-a3be-46fe-a00e-46ef96f0dedd', 'd85d141a-e806-45e6-83bd-3b48cef86d95', 'Very cool!', '2025-06-07 20:32:58.948263+00', '2025-06-07 20:32:58.948263+00', NULL);
INSERT INTO public.posts VALUES ('c1544597-2835-448b-8a57-67dba3f6b591', '6e137217-71a9-439b-a579-4bcd1a729af9', '6b7011e6-fb89-4542-8eba-31b219970185', '5501f6a8-a3be-46fe-a00e-46ef96f0dedd', 'd85d141a-e806-45e6-83bd-3b48cef86d95', 'Motivating stuff!', '2025-06-07 19:53:58.948263+00', '2025-06-07 19:53:58.948263+00', NULL);
INSERT INTO public.posts VALUES ('8344d706-6723-4474-91f0-f1d102eb2b80', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '6b7011e6-fb89-4542-8eba-31b219970185', '5501f6a8-a3be-46fe-a00e-46ef96f0dedd', 'd85d141a-e806-45e6-83bd-3b48cef86d95', 'Nice progress!', '2025-06-07 19:54:58.948263+00', '2025-06-07 19:54:58.948263+00', NULL);
INSERT INTO public.posts VALUES ('f12ae689-bb1a-49c4-9dd3-9698e4502310', '3b0ebec9-3805-490b-987f-4c6e7286319e', '6b7011e6-fb89-4542-8eba-31b219970185', '5501f6a8-a3be-46fe-a00e-46ef96f0dedd', 'd85d141a-e806-45e6-83bd-3b48cef86d95', 'Love this!', '2025-06-07 22:19:58.948263+00', '2025-06-07 22:19:58.948263+00', NULL);
INSERT INTO public.posts VALUES ('9f44d18d-2c82-42d2-8a3a-0241a2bad009', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'c4c10327-612e-4a10-a744-853e80c41a49', '106936c7-f18f-48c6-81ab-6f931815cced', 'dee234bb-394b-4471-a913-7e6153be0bb7', 'Keep going!', '2025-06-07 21:23:48.757811+00', '2025-06-07 21:23:48.757811+00', NULL);
INSERT INTO public.posts VALUES ('f9fa180b-133b-47c0-9bae-bae5fd2692d5', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'c4c10327-612e-4a10-a744-853e80c41a49', '106936c7-f18f-48c6-81ab-6f931815cced', 'dee234bb-394b-4471-a913-7e6153be0bb7', 'That’s awesome!', '2025-06-07 20:26:48.757811+00', '2025-06-07 20:26:48.757811+00', NULL);
INSERT INTO public.posts VALUES ('ac61f633-2754-4fa5-860d-5a1c50398e07', '56e89866-3a90-4f97-8461-34b24f842da4', 'c4c10327-612e-4a10-a744-853e80c41a49', '106936c7-f18f-48c6-81ab-6f931815cced', 'dee234bb-394b-4471-a913-7e6153be0bb7', 'Nice progress!', '2025-06-07 19:31:48.757811+00', '2025-06-07 19:31:48.757811+00', NULL);
INSERT INTO public.posts VALUES ('ed69b0a3-fb72-4665-99b0-326b9b607100', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'c4c10327-612e-4a10-a744-853e80c41a49', '106936c7-f18f-48c6-81ab-6f931815cced', 'dee234bb-394b-4471-a913-7e6153be0bb7', 'Interesting!', '2025-06-07 21:23:48.757811+00', '2025-06-07 21:23:48.757811+00', NULL);
INSERT INTO public.posts VALUES ('6738d34f-3fa3-4df6-8b40-24bd8046b489', '56e89866-3a90-4f97-8461-34b24f842da4', '47b0bc29-2ef0-4b57-a3b9-849968031526', '4db717b6-a224-4748-accf-17c23d4a7015', 'e003776a-dd4e-4c34-bfe2-503f25a993c3', 'Very cool!', '2025-06-07 19:38:14.923984+00', '2025-06-07 19:38:14.923984+00', NULL);
INSERT INTO public.posts VALUES ('6e93ee80-fa53-4c9c-a7c3-01c1eedf9cf1', '79758f94-70d8-4a11-b05f-793fcb57e093', '47b0bc29-2ef0-4b57-a3b9-849968031526', '4db717b6-a224-4748-accf-17c23d4a7015', 'e003776a-dd4e-4c34-bfe2-503f25a993c3', 'This inspired me.', '2025-06-07 17:16:14.923984+00', '2025-06-07 17:16:14.923984+00', NULL);
INSERT INTO public.posts VALUES ('be79505d-558c-45bb-a28b-52dd6a8635c1', '3b0ebec9-3805-490b-987f-4c6e7286319e', '47b0bc29-2ef0-4b57-a3b9-849968031526', '4db717b6-a224-4748-accf-17c23d4a7015', 'e003776a-dd4e-4c34-bfe2-503f25a993c3', 'Love this!', '2025-06-07 19:23:14.923984+00', '2025-06-07 19:23:14.923984+00', NULL);
INSERT INTO public.posts VALUES ('c04adb23-1f1a-4911-b4a8-ddc09efb104c', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'f3984ee6-5f7b-47b0-a88c-c684e10eae43', '950b125d-11c7-4703-877c-a87c1cb027b0', 'e2d16c26-ba68-4a8d-ab33-8a68890f1683', 'Following your example.', '2025-06-07 21:54:26.408471+00', '2025-06-07 21:54:26.408471+00', NULL);
INSERT INTO public.posts VALUES ('102da257-388f-4e53-9ae4-5d918d124b64', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'f3984ee6-5f7b-47b0-a88c-c684e10eae43', '950b125d-11c7-4703-877c-a87c1cb027b0', 'e2d16c26-ba68-4a8d-ab33-8a68890f1683', 'This inspired me.', '2025-06-07 22:46:26.408471+00', '2025-06-07 22:46:26.408471+00', NULL);
INSERT INTO public.posts VALUES ('cb3ef559-3bbe-490f-821d-321f2971b687', '79758f94-70d8-4a11-b05f-793fcb57e093', 'f3984ee6-5f7b-47b0-a88c-c684e10eae43', '950b125d-11c7-4703-877c-a87c1cb027b0', 'e2d16c26-ba68-4a8d-ab33-8a68890f1683', 'Interesting!', '2025-06-07 20:14:26.408471+00', '2025-06-07 20:14:26.408471+00', NULL);
INSERT INTO public.posts VALUES ('86ed5060-9a13-4005-92fa-56306287cdc1', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'f3984ee6-5f7b-47b0-a88c-c684e10eae43', '950b125d-11c7-4703-877c-a87c1cb027b0', 'e2d16c26-ba68-4a8d-ab33-8a68890f1683', 'Motivating stuff!', '2025-06-07 21:36:26.408471+00', '2025-06-07 21:36:26.408471+00', NULL);
INSERT INTO public.posts VALUES ('90dca6db-4b8c-480b-9981-2434e22c2584', '6e137217-71a9-439b-a579-4bcd1a729af9', 'f3984ee6-5f7b-47b0-a88c-c684e10eae43', '950b125d-11c7-4703-877c-a87c1cb027b0', 'e2d16c26-ba68-4a8d-ab33-8a68890f1683', 'Very cool!', '2025-06-07 21:10:26.408471+00', '2025-06-07 21:10:26.408471+00', NULL);
INSERT INTO public.posts VALUES ('19b7321a-eaf3-4c12-b5e4-ad99bae338d9', '56e89866-3a90-4f97-8461-34b24f842da4', 'f3984ee6-5f7b-47b0-a88c-c684e10eae43', '950b125d-11c7-4703-877c-a87c1cb027b0', 'e2d16c26-ba68-4a8d-ab33-8a68890f1683', 'Keep going!', '2025-06-07 22:41:26.408471+00', '2025-06-07 22:41:26.408471+00', NULL);
INSERT INTO public.posts VALUES ('3214bfee-5f62-4e4b-99f6-5ecc44813168', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'e5bf481a-040b-4b1b-96c7-787909876292', '2c31cf98-f564-4537-8a12-a71a705103ea', 'e5ac024b-a6ee-469f-834f-704f37620b25', 'This inspired me.', '2025-06-07 19:40:43.030841+00', '2025-06-07 19:40:43.030841+00', NULL);
INSERT INTO public.posts VALUES ('8b50b491-d101-49dd-8b91-7dd1a3a7c80f', '56e89866-3a90-4f97-8461-34b24f842da4', 'e5bf481a-040b-4b1b-96c7-787909876292', '2c31cf98-f564-4537-8a12-a71a705103ea', 'e5ac024b-a6ee-469f-834f-704f37620b25', 'Interesting!', '2025-06-07 20:19:43.030841+00', '2025-06-07 20:19:43.030841+00', NULL);
INSERT INTO public.posts VALUES ('c3384fc8-b882-41ea-af5b-4987052834ec', '79758f94-70d8-4a11-b05f-793fcb57e093', 'e5bf481a-040b-4b1b-96c7-787909876292', '2c31cf98-f564-4537-8a12-a71a705103ea', 'e5ac024b-a6ee-469f-834f-704f37620b25', 'Motivating stuff!', '2025-06-07 21:28:43.030841+00', '2025-06-07 21:28:43.030841+00', NULL);
INSERT INTO public.posts VALUES ('315d379e-756c-4c2d-a361-c48b25b88e57', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'e5bf481a-040b-4b1b-96c7-787909876292', '2c31cf98-f564-4537-8a12-a71a705103ea', 'e5ac024b-a6ee-469f-834f-704f37620b25', 'Nice progress!', '2025-06-07 20:55:43.030841+00', '2025-06-07 20:55:43.030841+00', NULL);
INSERT INTO public.posts VALUES ('053c1e83-5165-4e41-80e3-c5e478b16e3a', '6e137217-71a9-439b-a579-4bcd1a729af9', 'e5bf481a-040b-4b1b-96c7-787909876292', '2c31cf98-f564-4537-8a12-a71a705103ea', 'e5ac024b-a6ee-469f-834f-704f37620b25', 'Very cool!', '2025-06-07 20:51:43.030841+00', '2025-06-07 20:51:43.030841+00', NULL);
INSERT INTO public.posts VALUES ('a9339692-0332-49e4-b9fd-ca379661d4ff', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'e5bf481a-040b-4b1b-96c7-787909876292', '2c31cf98-f564-4537-8a12-a71a705103ea', 'e5ac024b-a6ee-469f-834f-704f37620b25', 'That’s awesome!', '2025-06-07 22:11:43.030841+00', '2025-06-07 22:11:43.030841+00', NULL);
INSERT INTO public.posts VALUES ('c469e127-ead1-45c2-9286-758a8346ab0a', '56e89866-3a90-4f97-8461-34b24f842da4', '111e36b5-3a31-4079-a5c8-8ead7454723b', '0ea22b31-fb48-49de-9299-446fdf46623d', 'f4d4efdf-82d7-491f-b9fc-4d10d40e6a1b', 'Love this!', '2025-06-07 21:28:50.533708+00', '2025-06-07 21:28:50.533708+00', NULL);
INSERT INTO public.posts VALUES ('477353f1-f8e6-4793-a895-e33331d283dd', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '111e36b5-3a31-4079-a5c8-8ead7454723b', '0ea22b31-fb48-49de-9299-446fdf46623d', 'f4d4efdf-82d7-491f-b9fc-4d10d40e6a1b', 'Motivating stuff!', '2025-06-07 19:17:50.533708+00', '2025-06-07 19:17:50.533708+00', NULL);
INSERT INTO public.posts VALUES ('7f8ceb0b-bf4e-409f-b35b-1384bddbf46f', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '111e36b5-3a31-4079-a5c8-8ead7454723b', '0ea22b31-fb48-49de-9299-446fdf46623d', 'f4d4efdf-82d7-491f-b9fc-4d10d40e6a1b', 'That’s awesome!', '2025-06-07 19:45:50.533708+00', '2025-06-07 19:45:50.533708+00', NULL);
INSERT INTO public.posts VALUES ('642d3409-c4e9-4c0b-b457-75e6fd3c0669', '6e137217-71a9-439b-a579-4bcd1a729af9', '111e36b5-3a31-4079-a5c8-8ead7454723b', '0ea22b31-fb48-49de-9299-446fdf46623d', 'f4d4efdf-82d7-491f-b9fc-4d10d40e6a1b', 'Following your example.', '2025-06-07 19:27:50.533708+00', '2025-06-07 19:27:50.533708+00', NULL);
INSERT INTO public.posts VALUES ('b45b8344-088d-4bd2-99e8-b1ea4b2f6154', '3b0ebec9-3805-490b-987f-4c6e7286319e', '111e36b5-3a31-4079-a5c8-8ead7454723b', '0ea22b31-fb48-49de-9299-446fdf46623d', 'f4d4efdf-82d7-491f-b9fc-4d10d40e6a1b', 'This inspired me.', '2025-06-07 19:18:50.533708+00', '2025-06-07 19:18:50.533708+00', NULL);
INSERT INTO public.posts VALUES ('d31be3a5-bc2b-416c-a1bc-385ae6c36d69', '79758f94-70d8-4a11-b05f-793fcb57e093', '111e36b5-3a31-4079-a5c8-8ead7454723b', '0ea22b31-fb48-49de-9299-446fdf46623d', 'f4d4efdf-82d7-491f-b9fc-4d10d40e6a1b', 'Nice progress!', '2025-06-07 19:58:50.533708+00', '2025-06-07 19:58:50.533708+00', NULL);
INSERT INTO public.posts VALUES ('d78db93b-e801-4bf1-8209-c7647cd228b4', '56e89866-3a90-4f97-8461-34b24f842da4', 'a9528f6b-3895-4cab-af6c-62a93ef1b83e', 'fe5cb029-80a1-4c37-8fd0-62af0afc0976', 'fbbcab19-51fd-4eb4-aa4b-750782900d5d', 'Nice progress!', '2025-06-07 21:11:29.925977+00', '2025-06-07 21:11:29.925977+00', NULL);
INSERT INTO public.posts VALUES ('777ef179-62f9-4c2e-b802-fde3003ebe4d', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'a9528f6b-3895-4cab-af6c-62a93ef1b83e', 'fe5cb029-80a1-4c37-8fd0-62af0afc0976', 'fbbcab19-51fd-4eb4-aa4b-750782900d5d', 'Very cool!', '2025-06-07 19:09:29.925977+00', '2025-06-07 19:09:29.925977+00', NULL);
INSERT INTO public.posts VALUES ('621c5d70-414c-4c1d-88fc-da3afb8d5514', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'a9528f6b-3895-4cab-af6c-62a93ef1b83e', 'fe5cb029-80a1-4c37-8fd0-62af0afc0976', 'fbbcab19-51fd-4eb4-aa4b-750782900d5d', 'Following your example.', '2025-06-07 18:54:29.925977+00', '2025-06-07 18:54:29.925977+00', NULL);
INSERT INTO public.posts VALUES ('4965441a-5a07-4a50-8254-f02de72b5712', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'a9528f6b-3895-4cab-af6c-62a93ef1b83e', 'fe5cb029-80a1-4c37-8fd0-62af0afc0976', 'fbbcab19-51fd-4eb4-aa4b-750782900d5d', 'Motivating stuff!', '2025-06-07 21:06:29.925977+00', '2025-06-07 21:06:29.925977+00', NULL);
INSERT INTO public.posts VALUES ('13e04411-896b-4bf0-92f0-47f896648ae7', '6e137217-71a9-439b-a579-4bcd1a729af9', 'a9528f6b-3895-4cab-af6c-62a93ef1b83e', 'fe5cb029-80a1-4c37-8fd0-62af0afc0976', 'fbbcab19-51fd-4eb4-aa4b-750782900d5d', 'That’s awesome!', '2025-06-07 19:08:29.925977+00', '2025-06-07 19:08:29.925977+00', NULL);
INSERT INTO public.posts VALUES ('0ed9ebd1-3993-4485-ae06-a99cf706aac8', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'a9528f6b-3895-4cab-af6c-62a93ef1b83e', 'fe5cb029-80a1-4c37-8fd0-62af0afc0976', 'fbbcab19-51fd-4eb4-aa4b-750782900d5d', 'Keep going!', '2025-06-07 19:51:29.925977+00', '2025-06-07 19:51:29.925977+00', NULL);


--
-- Data for Name: predictions; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.predictions VALUES ('69549c2c-cfcb-490e-be05-c4a228988cf6', '9dd80410-bd88-455a-b8fe-a5c2e39edcf0', 85, 'You''ve taken the first step towards mastering a new skill, which is great! Cooking is a skill that requires practice, so start small and be consistent. Adjust recipes to your liking - it''s part of the process. Embrace mistakes, they''re only opportunities to learn. You''ve got this!', '2025-06-07 15:52:06.388359+00', '2025-06-07 15:52:06.388359+00');
INSERT INTO public.predictions VALUES ('9ad163d9-5520-432b-8471-d1af8eafa59c', '9edaf28d-e47c-4bec-aa12-f4d6b78c6edb', 85, 'Your goal aligns with enhancing your self-growth, and it''s achievable within the timeframe. Break it into small tasks: choose the books, set a reading schedule, and apply what you learn. Remember, productivity is not about being busy, but being efficient. Keep up the pace!', '2025-06-07 15:52:21.291314+00', '2025-06-07 15:52:21.291314+00');
INSERT INTO public.predictions VALUES ('d2846b13-1ad4-4382-90bd-df42fa5d72bf', '609eba42-64bc-460a-9468-bc011dae591c', 85, 'Your decision to cultivate a daily meditation habit is a powerful step towards personal wellness. Start with short sessions and gradually increase them. Do it first thing in the morning and award yourself for consistency. Be patient with yourself and remember, all progress is good progress. You''ve got this!', '2025-06-07 15:57:22.96576+00', '2025-06-07 15:57:22.96576+00');
INSERT INTO public.predictions VALUES ('6b555959-f6b2-470e-bf67-d93d13bba0ed', '8ecb2652-80cf-4fe2-bab8-f3841d663a2d', 75, 'Establishing a healthy morning routine is an excellent way to start your day. Remember, consistency is key. Each day try to do a little more than the previous one. Don''t rush it - progress over perfection. Believe in yourself and keep going. Your wellness is worth it!', '2025-06-07 16:44:05.361271+00', '2025-06-07 16:44:05.361271+00');
INSERT INTO public.predictions VALUES ('0900eb38-4b86-4648-9f5b-d5403cbd03ee', '66b19273-144f-4dd0-971e-830c18c26b38', 90, 'Great goal focusing on tactical skill enhancement! Consistency is key here. Take one problem at a time, practice regularly, and you''ll swiftly conquer all 10 problems. Remember, every problem solved is a step closer to your career advancement. Keep pushing, success is sure!', '2025-06-07 16:44:14.300485+00', '2025-06-07 16:44:14.300485+00');
INSERT INTO public.predictions VALUES ('991e30e8-9c7e-4628-b883-028347839099', 'e5bf481a-040b-4b1b-96c7-787909876292', 80, 'Achieving spiritual growth through regular Quran reading is a profound journey. It''s essential not to rush. Allow yourself to understand and internalize each verse, which could elevate your success probability. Remember, it''s about making consistent progress and meaningful reflection, not simply finishing within the set met time frame. Good luck!', '2025-06-07 16:44:38.049773+00', '2025-06-07 16:44:38.049773+00');
INSERT INTO public.predictions VALUES ('3379e208-050b-40ff-b092-fdcab8e2ec8e', '912c9992-debd-42b8-9df7-28a9c4b7a319', 80, 'Journaling can be empowering and therapeutic. Start gradually, perhaps a few minutes per day, to make it enjoyable rather than a task. Remember, this journey of self-growth doesn''t demand perfection, be as raw and genuine as possible.', '2025-06-07 16:55:21.472665+00', '2025-06-07 16:55:21.472665+00');
INSERT INTO public.predictions VALUES ('585eb144-27ef-4bc2-8782-5b1936324ff0', '6b7011e6-fb89-4542-8eba-31b219970185', 90, 'Your determination for better health and stamina is impressive. Consistency is key in reaching your goal. Start by running short distances and gradually increase it. Remember, it''s about steady progress, not instant perfection. You''ve got this!', '2025-06-07 16:55:33.69937+00', '2025-06-07 16:55:33.69937+00');
INSERT INTO public.predictions VALUES ('2d911abd-2c88-4a3c-82b7-d4c0ca4c1e20', 'c84f3282-901c-4433-82ee-42fc3f4ddf3a', 90, 'Your ample preparation time is a great asset. Immersion is key. Try labeling items around your house with their translations and practice at least 15 minutes per day. Viel Erfolg!', '2025-06-07 16:55:44.02934+00', '2025-06-07 16:55:44.02934+00');
INSERT INTO public.predictions VALUES ('4c3331d9-c029-4c76-9fa4-66e20f6cb876', '47b0bc29-2ef0-4b57-a3b9-849968031526', 85, 'Building a personal website is a fantastic way to showcase your abilities. Remember, consistency is key. Allocate time every day to learning, planning and designing your website. Step by step, you will be amazed how much progress you can make. Don''t let roadblocks deter you, they are opportunities to learn and grow. Best of luck!', '2025-06-07 18:36:06.323029+00', '2025-06-07 18:36:06.323029+00');
INSERT INTO public.predictions VALUES ('0c4156ab-a9e7-42c8-993a-cf33a27f28cb', 'c3f3d0c5-29e5-4167-87ab-8178d06ac7c2', 75, 'Great goal! Remember, consistency is key when starting a workout routine. Start small and gradually increase the intensity to avoid burnout. Fuel your body with nutritious foods and keep hydrated. You''re stronger than you think, believe in yourself!', '2025-06-07 18:36:23.143144+00', '2025-06-07 18:36:23.143144+00');
INSERT INTO public.predictions VALUES ('aaea9672-7bbb-46a3-82c3-33bc97feb95f', '111e36b5-3a31-4079-a5c8-8ead7454723b', 70, 'Persistence, determination, and daily practice are keys for improving time management skills. Prioritizing tasks, breaking them up, and scheduling dedicated focus times can be beneficial. Remember, every minute counts; utilize them wisely. You''re on a wonderful journey of self-growth. Keep pushing!', '2025-06-07 18:36:43.706219+00', '2025-06-07 18:36:43.706219+00');
INSERT INTO public.predictions VALUES ('b0924872-1201-48ae-8dcd-95ecf2a8498c', 'c4c10327-612e-4a10-a744-853e80c41a49', 85, 'Embrace the learning process and have fun. Don''t worry about creating a masterpiece right away, focus on consistent practice and gradual improvement. Remember that every artist once started from zero, don''t be discouraged, your passion will drive your progress.', '2025-06-07 19:15:45.854921+00', '2025-06-07 19:15:45.854921+00');
INSERT INTO public.predictions VALUES ('349caff1-9016-41a8-a4c0-5dc58ef6edf8', 'ac557757-04c4-428d-a616-655ecda8b969', 95, 'Steady and consistent progress is key. Set aside specific daily time to learn and review the words. You''ve got plenty of time to reach your goal. Remember, the beauty of language is not just in knowing but using it. Practice incorporating your new words into daily conversations. Keep going, you''ve got this!', '2025-06-07 19:15:57.014213+00', '2025-06-07 19:15:57.014213+00');
INSERT INTO public.predictions VALUES ('1c141c9f-9a28-4a4b-a40f-b5ade81fcabb', '30a43be9-a3fe-43c9-b536-25cb0ee96c1b', 85, 'A clutter-free workspace can lead to a focused mind! Organizing things can seem overwhelming, but spread the task over a few days and tackle one area at a time. Make it a fun endeavor by adding music or podcasts into the mix. Start today and enjoy a more productive, clutter-free environment sooner!', '2025-06-07 19:16:09.647857+00', '2025-06-07 19:16:09.647857+00');
INSERT INTO public.predictions VALUES ('7c915f20-cac1-4ef3-8773-a5d9d14e992c', '70679b66-1611-48c0-9bcb-0558b6d28965', 85, 'Success is right within reach. Consistent study and immersion can lead to rapid language growth. Try speaking German in your daily life or practicing with native speakers to expand your skills faster. You have the aim, now work hard for it, und viel Glück!', '2025-06-07 19:31:01.898856+00', '2025-06-07 19:31:01.898856+00');
INSERT INTO public.predictions VALUES ('9ff10245-bfd3-49d8-9136-159d50e0c62a', '7e8e87bc-1b4c-44e9-832c-6e29617642d8', 70, 'Your body considers mindfulness very rewarding. Integrate stretching routines into your daily schedule, and remember to correct your posture regularly. With constant attention and effort, your back pain will reduce significantly. You''re on a journey of wellness. Make every step count.', '2025-06-07 19:31:13.329452+00', '2025-06-07 19:31:13.329452+00');
INSERT INTO public.predictions VALUES ('10138381-ad74-4e31-bb05-1bf65c2c1e06', 'a9528f6b-3895-4cab-af6c-62a93ef1b83e', 85, 'You''ve already taken the first step in your career advancement by deciding to create a personal portfolio online. Now it''s time to persevere and consistently work on it. Remember, your portfolio will be a reflection of your skills and dedication. Break it into smaller tasks to manage it better and avoid getting overwhelmed. Keep your focus, success is right around the corner!', '2025-06-07 19:31:54.122748+00', '2025-06-07 19:31:54.122748+00');
INSERT INTO public.predictions VALUES ('d2d23a8c-338b-463c-ace8-ddd1fbde9640', 'b438fcc1-9526-4e3c-8bd2-039bc7e71e7b', 90, 'Your goal is challenging yet achievable, with a clear focus. Stay proactive, outline your ideas, and start drafting immediately. Remember that the first draft doesn''t need to be perfect; optimize as you go. Your unique experiences with ChatGPT in learning would surely add value in the AI in education space. Good luck!', '2025-06-07 19:54:56.346002+00', '2025-06-07 19:54:56.346002+00');
INSERT INTO public.predictions VALUES ('2b153f5c-8ae8-4547-9ddb-0b219cb91372', '513f748c-0dc9-4b04-a8f0-f33ea5892151', 85, 'Embrace the adventure of learning Go. Consistency is key in backend development, so schedule regular coding hours. Small, daily progress adds up to big victories. Remember, every expert was once a beginner. You''ve got this!', '2025-06-07 19:55:01.280127+00', '2025-06-07 19:55:01.280127+00');
INSERT INTO public.predictions VALUES ('41e6e1c0-d1f3-4dd3-ab97-490951676e89', 'f3984ee6-5f7b-47b0-a88c-c684e10eae43', 70, 'Consistency and discipline will be your keys to success. Start with small changes; they are easier to maintain and accumulate over time. Remember, each day is an opportunity to become a healthier version of yourself. Stay committed, and you''ll become your own motivation. Keep going!', '2025-06-07 19:55:08.848015+00', '2025-06-07 19:55:08.848015+00');


--
-- Data for Name: promises; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.promises VALUES ('9dd80410-bd88-455a-b8fe-a5c2e39edcf0', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'Learn how to cook samsa', 'I want to learn how to cook samsa to replace raw chicken with rice.', '2025-07-30 00:00:00+00', false, 'in_progress', 'personal', '2025-06-07 15:52:02.460083+00', '2025-06-07 15:52:02.460083+00', NULL);
INSERT INTO public.promises VALUES ('9edaf28d-e47c-4bec-aa12-f4d6b78c6edb', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'Read 3 books on productivity', 'I want to become more efficient in my daily life by reading useful books.', '2025-08-15 00:00:00+00', false, 'in_progress', 'self-growth', '2025-06-07 15:52:18.108362+00', '2025-06-07 15:52:18.108362+00', NULL);
INSERT INTO public.promises VALUES ('609eba42-64bc-460a-9468-bc011dae591c', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'Practice daily meditation', 'Build a habit of meditating every morning.', '2025-07-20 00:00:00+00', false, 'in_progress', 'wellness', '2025-06-07 15:57:16.024924+00', '2025-06-07 15:57:16.024925+00', NULL);
INSERT INTO public.promises VALUES ('8ecb2652-80cf-4fe2-bab8-f3841d663a2d', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'Start a morning routine', 'Build a healthy start to my day with consistent habits.', '2025-08-10 00:00:00+00', false, 'in_progress', 'wellness', '2025-06-07 16:44:01.500466+00', '2025-06-07 16:44:01.500466+00', NULL);
INSERT INTO public.promises VALUES ('66b19273-144f-4dd0-971e-830c18c26b38', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'Complete 10 LeetCode problems', 'To prepare for interviews and sharpen problem-solving skills.', '2025-08-20 00:00:00+00', false, 'in_progress', 'career', '2025-06-07 16:44:11.615415+00', '2025-06-07 16:44:11.615415+00', NULL);
INSERT INTO public.promises VALUES ('e5bf481a-040b-4b1b-96c7-787909876292', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'Read the Quran in a month', 'Spiritual goal: read and reflect on the Quran regularly.', '2025-08-31 00:00:00+00', true, 'in_progress', 'spiritual', '2025-06-07 16:44:35.20675+00', '2025-06-07 16:44:35.20675+00', NULL);
INSERT INTO public.promises VALUES ('912c9992-debd-42b8-9df7-28a9c4b7a319', '56e89866-3a90-4f97-8461-34b24f842da4', 'Start a daily journaling habit', 'I want to reflect on my days and improve self-awareness', '2025-08-10 00:00:00+00', false, 'in_progress', 'self-growth', '2025-06-07 16:55:17.668905+00', '2025-06-07 16:55:17.668905+00', NULL);
INSERT INTO public.promises VALUES ('6b7011e6-fb89-4542-8eba-31b219970185', '56e89866-3a90-4f97-8461-34b24f842da4', 'Run 5km without stopping', 'Improve stamina and health by training regularly', '2025-09-01 00:00:00+00', false, 'in_progress', 'fitness', '2025-06-07 16:55:30.313291+00', '2025-06-07 16:55:30.313291+00', NULL);
INSERT INTO public.promises VALUES ('c84f3282-901c-4433-82ee-42fc3f4ddf3a', '56e89866-3a90-4f97-8461-34b24f842da4', 'Learn basic German phrases', 'Prepare for trip to Germany by learning key phrases', '2025-08-25 00:00:00+00', false, 'in_progress', 'language', '2025-06-07 16:55:41.745948+00', '2025-06-07 16:55:41.745948+00', NULL);
INSERT INTO public.promises VALUES ('47b0bc29-2ef0-4b57-a3b9-849968031526', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'Build a personal website', 'I want to showcase my projects and skills online.', '2025-09-30 00:00:00+00', false, 'in_progress', 'career', '2025-06-07 18:36:02.99803+00', '2025-06-07 18:36:02.99803+00', NULL);
INSERT INTO public.promises VALUES ('c3f3d0c5-29e5-4167-87ab-8178d06ac7c2', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'Start a workout routine', 'I want to get in better shape and feel more energetic.', '2025-08-30 00:00:00+00', false, 'in_progress', 'fitness', '2025-06-07 18:36:20.43694+00', '2025-06-07 18:36:20.43694+00', NULL);
INSERT INTO public.promises VALUES ('111e36b5-3a31-4079-a5c8-8ead7454723b', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'Improve time management', 'I want to be more productive and focused.', '2025-08-20 00:00:00+00', false, 'in_progress', 'self-growth', '2025-06-07 18:36:40.687889+00', '2025-06-07 18:36:40.687889+00', NULL);
INSERT INTO public.promises VALUES ('c4c10327-612e-4a10-a744-853e80c41a49', '6e137217-71a9-439b-a579-4bcd1a729af9', 'Learn to draw digitally', 'Practice digital drawing to improve my art skills.', '2025-09-15 00:00:00+00', false, 'in_progress', 'hobby', '2025-06-07 19:15:43.098632+00', '2025-06-07 19:15:43.098632+00', NULL);
INSERT INTO public.promises VALUES ('ac557757-04c4-428d-a616-655ecda8b969', '6e137217-71a9-439b-a579-4bcd1a729af9', 'Memorize 50 new English words', 'Improve my vocabulary by learning 50 words.', '2025-08-28 00:00:00+00', false, 'in_progress', 'language', '2025-06-07 19:15:54.292575+00', '2025-06-07 19:15:54.292576+00', NULL);
INSERT INTO public.promises VALUES ('30a43be9-a3fe-43c9-b536-25cb0ee96c1b', '6e137217-71a9-439b-a579-4bcd1a729af9', 'Organize my workspace', 'Declutter and optimize my desk for better focus.', '2025-08-15 00:00:00+00', true, 'in_progress', 'personal', '2025-06-07 19:16:06.481473+00', '2025-06-07 19:16:06.481473+00', NULL);
INSERT INTO public.promises VALUES ('70679b66-1611-48c0-9bcb-0558b6d28965', '79758f94-70d8-4a11-b05f-793fcb57e093', 'Pass German A2 Exam by October', 'I need to pass the A2 level to apply for residency.', '2025-10-01 00:00:00+00', false, 'in_progress', 'language', '2025-06-07 19:30:59.472305+00', '2025-06-07 19:30:59.472306+00', NULL);
INSERT INTO public.promises VALUES ('7e8e87bc-1b4c-44e9-832c-6e29617642d8', '79758f94-70d8-4a11-b05f-793fcb57e093', 'Fix posture and reduce back pain', 'Improve daily posture and reduce pain from long work hours.', '2025-09-01 00:00:00+00', false, 'in_progress', 'health', '2025-06-07 19:31:11.050676+00', '2025-06-07 19:31:11.050676+00', NULL);
INSERT INTO public.promises VALUES ('a9528f6b-3895-4cab-af6c-62a93ef1b83e', '79758f94-70d8-4a11-b05f-793fcb57e093', 'Make a personal online portfolio', 'To showcase my projects and resume to employers.', '2025-08-25 00:00:00+00', false, 'in_progress', 'career', '2025-06-07 19:31:50.507042+00', '2025-06-07 19:31:50.507042+00', NULL);
INSERT INTO public.promises VALUES ('b438fcc1-9526-4e3c-8bd2-039bc7e71e7b', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'Write and publish a blog post about AI in education', 'Share my insights and experience with using ChatGPT in learning.', '2025-08-12 00:00:00+00', false, 'in_progress', 'career', '2025-06-07 19:54:52.5664+00', '2025-06-07 19:54:52.5664+00', NULL);
INSERT INTO public.promises VALUES ('513f748c-0dc9-4b04-a8f0-f33ea5892151', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'Build a basic to-do app with Go', 'Get comfortable with backend development in Go by making a simple web app.', '2025-08-20 00:00:00+00', false, 'in_progress', 'career', '2025-06-07 19:54:58.449098+00', '2025-06-07 19:54:58.449098+00', NULL);
INSERT INTO public.promises VALUES ('f3984ee6-5f7b-47b0-a88c-c684e10eae43', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'Switch to a healthier morning routine', 'Establish better habits for energy and focus.', '2025-08-05 00:00:00+00', true, 'in_progress', 'wellness', '2025-06-07 19:55:06.226955+00', '2025-06-07 19:55:06.226955+00', NULL);


--
-- Data for Name: refresh_tokens; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.refresh_tokens VALUES ('8822bd60-0fd9-4e2a-8be0-0ac1b9395156', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDk5ODQ0OTEsInJvbGUiOiJ1c2VyIiwidXNlcl9pZCI6ImEwM2IzMzAwLTBkMTQtNDU0ZS04MTEyLTdiM2VhZGE5MGFlOSJ9.nmhpbUVwjdd23eU01dx0LOmieS2P8h73zqptjmvZVkY', '2025-06-15 10:48:11.449591+00', '2025-06-08 10:48:11.450048+00', NULL);
INSERT INTO public.refresh_tokens VALUES ('a5241f73-7e5a-48a4-96b1-3b0c025367c2', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTAwMDQ5MzMsInJvbGUiOiJ1c2VyIiwidXNlcl9pZCI6IjA5MmI0NGQ5LWU5ZDctNGE0MC1hNTlmLWQ4NjhkMTkzZGQwZCJ9.t33GuyaLKUmtQV0da9OF1bCo6irCGWefD2dHuZU7W1s', '2025-06-15 16:28:53.010632+00', '2025-06-08 16:28:53.011018+00', NULL);
INSERT INTO public.refresh_tokens VALUES ('b37c7d1e-058e-41bd-ad42-5b4b5bd3e71d', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTAwMDU2MDMsInJvbGUiOiJ1c2VyIiwidXNlcl9pZCI6IjNiMGViZWM5LTM4MDUtNDkwYi05ODdmLTRjNmU3Mjg2MzE5ZSJ9.-kvyUOSBSEw0CBaaLdH3Zvy9CBOKFvNB3djHNrlnQ2o', '2025-06-15 16:40:03.298776+00', '2025-06-08 16:40:03.299133+00', NULL);
INSERT INTO public.refresh_tokens VALUES ('430be8ad-aa15-48a1-80d1-b3694bc81c9e', '56e89866-3a90-4f97-8461-34b24f842da4', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTAwMDU2MzgsInJvbGUiOiJ1c2VyIiwidXNlcl9pZCI6IjU2ZTg5ODY2LTNhOTAtNGY5Ny04NDYxLTM0YjI0Zjg0MmRhNCJ9.xBmbB-2_TODqyKA1okr2V7rOxc3ZUVqJVOukm4O7RPw', '2025-06-15 16:40:38.134017+00', '2025-06-08 16:40:38.134385+00', NULL);
INSERT INTO public.refresh_tokens VALUES ('f05372f1-c338-404d-9c23-49f95e84fd67', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTAwMDU3NDUsInJvbGUiOiJ1c2VyIiwidXNlcl9pZCI6IjY5MjlhYWE5LWYzYTgtNGUyOC1iOTU5LWY2NTg1ZjY1YjZiMSJ9.IAPVl-w-ecx-oGVzuGjI88KOIiFJloKTYnWLcqvkZSU', '2025-06-15 16:42:25.089589+00', '2025-06-08 16:42:25.089935+00', NULL);
INSERT INTO public.refresh_tokens VALUES ('9dbf324b-3dfe-479c-ab70-59cd24c6abf9', '6e137217-71a9-439b-a579-4bcd1a729af9', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTAwMDU3NzMsInJvbGUiOiJ1c2VyIiwidXNlcl9pZCI6IjZlMTM3MjE3LTcxYTktNDM5Yi1hNTc5LTRiY2QxYTcyOWFmOSJ9.VNP2GW8MStrGV3r7U204C59tF9bed6FgOfHm1jWmVB8', '2025-06-15 16:42:53.853144+00', '2025-06-08 16:42:53.853526+00', NULL);
INSERT INTO public.refresh_tokens VALUES ('fff01343-4132-4f0d-a0f6-ccd661361f58', '79758f94-70d8-4a11-b05f-793fcb57e093', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTAwMDU3OTEsInJvbGUiOiJ1c2VyIiwidXNlcl9pZCI6Ijc5NzU4Zjk0LTcwZDgtNGExMS1iMDVmLTc5M2ZjYjU3ZTA5MyJ9.9TA7Am4MiC0qmA63xSz6b4T_KmZpQMnpdk7BQx8bTGU', '2025-06-15 16:43:11.306128+00', '2025-06-08 16:43:11.306538+00', NULL);
INSERT INTO public.refresh_tokens VALUES ('f68566b0-a265-40aa-a806-3cfd02718cc8', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTAwMDU4MjEsInJvbGUiOiJ1c2VyIiwidXNlcl9pZCI6ImEwM2IzMzAwLTBkMTQtNDU0ZS04MTEyLTdiM2VhZGE5MGFlOSJ9.g2RzBk6MBwJUY7IRAUmt_lhnwzxLtW1UPYmmPIDjgIg', '2025-06-15 16:43:41.265524+00', '2025-06-08 16:43:41.265865+00', NULL);
INSERT INTO public.refresh_tokens VALUES ('8d2ab5fd-2705-4f7a-ba47-ea06ff3a8d37', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTAwMDU4MzksInJvbGUiOiJ1c2VyIiwidXNlcl9pZCI6ImEyMDNjODlhLTNhYjktNDJiMi1hZTkyLTBjOTc5ZDg0MjVkZSJ9.-WTWv_E4Q9WP_JWSy30hnCGBxvMg4PP5hIIG_obwmXk', '2025-06-15 16:43:59.772413+00', '2025-06-08 16:43:59.772769+00', NULL);
INSERT INTO public.refresh_tokens VALUES ('3f3312ca-262c-4e3f-a121-ad719fc3fe4d', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTAwMDU4NzEsInJvbGUiOiJ1c2VyIiwidXNlcl9pZCI6ImFlODg0ZWQyLTA3NWEtNGVkMi05ZGRlLTczNjkyMWQ4MGU5ZiJ9.Pt89W8oZsAeWeNGbPOxLYTKby6RDmTgc0HkvVj8AS68', '2025-06-15 16:44:31.01195+00', '2025-06-08 16:44:31.012351+00', NULL);
INSERT INTO public.refresh_tokens VALUES ('8b9a2432-0671-47e7-a48b-679097e180de', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTAwMDU4OTQsInJvbGUiOiJ1c2VyIiwidXNlcl9pZCI6ImQ0NDBmZjczLTgxYWUtNGE3Yy1iN2NiLTMzMzZlM2IyNGE0YyJ9.YKJuugTsP3nRFd-Y6NpDhVKnforWm85_lMRpoYAEP_M', '2025-06-15 16:44:54.390702+00', '2025-06-08 16:44:54.391095+00', NULL);
INSERT INTO public.refresh_tokens VALUES ('5c813180-ad6d-4f9a-b750-5f3ac2bad683', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTAwMDU5MjksInJvbGUiOiJ1c2VyIiwidXNlcl9pZCI6ImQ5NDM1MTJjLWE0NjItNGY5ZS04ODgxLTEzZGMyYjliNTllMSJ9.D0DOTLENIVE5ABo7nyUOSxteUy9GFkHUWXcCtWyqNds', '2025-06-15 16:45:29.098496+00', '2025-06-08 16:45:29.098869+00', NULL);
INSERT INTO public.refresh_tokens VALUES ('fb6d86a2-497d-48ce-9734-de615fc024d5', 'f399a996-1567-4fc0-a86a-92774494ca44', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTAwMDU5NTcsInJvbGUiOiJ1c2VyIiwidXNlcl9pZCI6ImYzOTlhOTk2LTE1NjctNGZjMC1hODZhLTkyNzc0NDk0Y2E0NCJ9.Fe3EuNh30pngdmSWTJTqxsL3sSUsy3T8qKeDjECio9Y', '2025-06-15 16:45:57.242201+00', '2025-06-08 16:45:57.242537+00', NULL);
INSERT INTO public.refresh_tokens VALUES ('27ff2dc1-d063-479c-bf19-15a7fea14b6f', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTAwMDYwMDAsInJvbGUiOiJ1c2VyIiwidXNlcl9pZCI6IjA5MmI0NGQ5LWU5ZDctNGE0MC1hNTlmLWQ4NjhkMTkzZGQwZCJ9.YqP1-MiCGyjgzgtn4F-V3HaVoId2XR388Nsltunk5Lo', '2025-06-15 16:46:40.777304+00', '2025-06-08 16:46:40.777645+00', NULL);


--
-- Data for Name: user_badges; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.user_badges VALUES ('6b895dfe-df3e-45c2-b4a0-439047c150a7', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'f0963252-7299-43e2-8849-6cce3b3ea52a', '2025-06-08 00:09:22.3649+00');
INSERT INTO public.user_badges VALUES ('9ce0144d-157c-4fa3-809f-ece1c283338b', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '4769b089-f1c3-4cb6-88af-7f53614d1e6d', '2025-06-08 00:09:22.372037+00');
INSERT INTO public.user_badges VALUES ('38078625-c5a9-4c1c-9b02-f2ddaa8c0681', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'f0963252-7299-43e2-8849-6cce3b3ea52a', '2025-06-08 00:09:22.376947+00');
INSERT INTO public.user_badges VALUES ('add43dbd-fc72-4a87-aee3-e04ecfebf602', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '4769b089-f1c3-4cb6-88af-7f53614d1e6d', '2025-06-08 00:09:22.38352+00');
INSERT INTO public.user_badges VALUES ('12b5956b-9322-466e-87fe-1f307bae4a05', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'd67d0bd3-06ad-4f4c-bf43-09ebcb3240a0', '2025-06-08 00:09:22.387813+00');
INSERT INTO public.user_badges VALUES ('507ab96e-d6b8-4e2b-96e3-e6948c99fb89', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'f0963252-7299-43e2-8849-6cce3b3ea52a', '2025-06-08 00:09:22.392321+00');
INSERT INTO public.user_badges VALUES ('071d8ee2-de78-485f-8800-83b5965da208', '3b0ebec9-3805-490b-987f-4c6e7286319e', '4769b089-f1c3-4cb6-88af-7f53614d1e6d', '2025-06-08 00:09:22.398682+00');
INSERT INTO public.user_badges VALUES ('48dae91b-59bc-4959-811a-3b6aa531cdf0', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'f0963252-7299-43e2-8849-6cce3b3ea52a', '2025-06-08 00:09:22.403283+00');
INSERT INTO public.user_badges VALUES ('995feda3-8229-4c54-b1ea-25c443766cdf', '6e137217-71a9-439b-a579-4bcd1a729af9', 'f0963252-7299-43e2-8849-6cce3b3ea52a', '2025-06-08 00:09:22.410088+00');
INSERT INTO public.user_badges VALUES ('878e148e-b48b-40b4-9328-0edb7264cc55', '79758f94-70d8-4a11-b05f-793fcb57e093', 'f0963252-7299-43e2-8849-6cce3b3ea52a', '2025-06-08 00:09:22.415496+00');
INSERT INTO public.user_badges VALUES ('8f6199d6-6ac3-4c65-b4e7-0ebd0399401a', '56e89866-3a90-4f97-8461-34b24f842da4', '4769b089-f1c3-4cb6-88af-7f53614d1e6d', '2025-06-08 00:09:22.420916+00');
INSERT INTO public.user_badges VALUES ('8e1d7f24-a85f-4afb-8275-ec8401dc2579', '56e89866-3a90-4f97-8461-34b24f842da4', 'f0963252-7299-43e2-8849-6cce3b3ea52a', '2025-06-08 00:09:22.425801+00');
INSERT INTO public.user_badges VALUES ('b9c1db49-3493-47a4-81f9-3dd4fd415c69', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '4769b089-f1c3-4cb6-88af-7f53614d1e6d', '2025-06-08 10:47:20.324201+00');
INSERT INTO public.user_badges VALUES ('5790e99b-fa2b-46cd-9f06-10f60ec219c5', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'd67d0bd3-06ad-4f4c-bf43-09ebcb3240a0', '2025-06-08 10:47:20.330171+00');
INSERT INTO public.user_badges VALUES ('49ea490d-2a5e-496b-8f21-c295bbc7867f', '6e137217-71a9-439b-a579-4bcd1a729af9', '4769b089-f1c3-4cb6-88af-7f53614d1e6d', '2025-06-08 10:47:20.338701+00');
INSERT INTO public.user_badges VALUES ('db3e7619-3705-4eb8-b665-a601e6bbaf17', '6e137217-71a9-439b-a579-4bcd1a729af9', 'd67d0bd3-06ad-4f4c-bf43-09ebcb3240a0', '2025-06-08 10:47:20.341486+00');
INSERT INTO public.user_badges VALUES ('7edc7d04-99d2-4809-9cba-52f6651ba7e4', '79758f94-70d8-4a11-b05f-793fcb57e093', '4769b089-f1c3-4cb6-88af-7f53614d1e6d', '2025-06-08 10:47:20.345514+00');
INSERT INTO public.user_badges VALUES ('a7a7ec38-18af-43fb-9955-4ef36d149687', '79758f94-70d8-4a11-b05f-793fcb57e093', 'd67d0bd3-06ad-4f4c-bf43-09ebcb3240a0', '2025-06-08 11:01:46.835483+00');
INSERT INTO public.user_badges VALUES ('3bf2ee02-5964-4666-9065-6d2397683bbb', 'f399a996-1567-4fc0-a86a-92774494ca44', '712e8ca2-37e2-46eb-98e6-1615f66d0fd7', '2025-06-08 15:46:13.173896+00');
INSERT INTO public.user_badges VALUES ('24cba622-3b91-439d-888a-f28f225f150b', 'f399a996-1567-4fc0-a86a-92774494ca44', '5db810f1-1a7f-4022-9ca1-85b3b727db56', '2025-06-08 15:46:13.178574+00');
INSERT INTO public.user_badges VALUES ('d0536961-99c7-4055-96f5-8e8cfdb80c0e', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '712e8ca2-37e2-46eb-98e6-1615f66d0fd7', '2025-06-08 15:46:13.18236+00');
INSERT INTO public.user_badges VALUES ('055a25bc-4f09-4258-ad92-59ccb7f98a23', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '5db810f1-1a7f-4022-9ca1-85b3b727db56', '2025-06-08 15:46:13.184845+00');
INSERT INTO public.user_badges VALUES ('00706ea6-ed01-49fb-bf37-184e729bad92', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '712e8ca2-37e2-46eb-98e6-1615f66d0fd7', '2025-06-08 15:46:13.189417+00');
INSERT INTO public.user_badges VALUES ('4d011725-d9ea-496b-9d22-25bd6bbaea5f', 'd440ff73-81ae-4a7c-b7cb-3336e3b24a4c', '5db810f1-1a7f-4022-9ca1-85b3b727db56', '2025-06-08 15:46:13.191874+00');
INSERT INTO public.user_badges VALUES ('fe545b47-3b32-47a1-8a20-6421f71c9a80', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '712e8ca2-37e2-46eb-98e6-1615f66d0fd7', '2025-06-08 15:46:13.195738+00');
INSERT INTO public.user_badges VALUES ('12c71fb3-f49e-439c-8000-96da5d6225db', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '5db810f1-1a7f-4022-9ca1-85b3b727db56', '2025-06-08 15:46:13.198074+00');
INSERT INTO public.user_badges VALUES ('e5451999-88de-4584-abee-33237e261f0d', '3b0ebec9-3805-490b-987f-4c6e7286319e', '712e8ca2-37e2-46eb-98e6-1615f66d0fd7', '2025-06-08 15:46:13.201419+00');
INSERT INTO public.user_badges VALUES ('9bcbeb99-07ab-4a72-b707-7f79279e1552', '3b0ebec9-3805-490b-987f-4c6e7286319e', '5db810f1-1a7f-4022-9ca1-85b3b727db56', '2025-06-08 15:46:13.203582+00');
INSERT INTO public.user_badges VALUES ('1b61aa8e-1596-4d00-a5ba-199caafca4a2', 'ae884ed2-075a-4ed2-9dde-736921d80e9f', '712e8ca2-37e2-46eb-98e6-1615f66d0fd7', '2025-06-08 15:46:13.206594+00');
INSERT INTO public.user_badges VALUES ('578b76c0-50e9-4f53-badf-2b4b2b391ee4', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '712e8ca2-37e2-46eb-98e6-1615f66d0fd7', '2025-06-08 15:46:13.209542+00');
INSERT INTO public.user_badges VALUES ('96c6a045-83ae-4119-86d7-e4037ee1137c', 'a203c89a-3ab9-42b2-ae92-0c979d8425de', '5db810f1-1a7f-4022-9ca1-85b3b727db56', '2025-06-08 15:46:13.21329+00');
INSERT INTO public.user_badges VALUES ('6d4b904a-ecd4-42ef-af83-a8c79f1a47dd', '6e137217-71a9-439b-a579-4bcd1a729af9', '712e8ca2-37e2-46eb-98e6-1615f66d0fd7', '2025-06-08 15:46:13.217053+00');
INSERT INTO public.user_badges VALUES ('3ccdd0bd-a27f-4a42-bc61-0c9a15b37125', '6e137217-71a9-439b-a579-4bcd1a729af9', '5db810f1-1a7f-4022-9ca1-85b3b727db56', '2025-06-08 15:46:13.219315+00');
INSERT INTO public.user_badges VALUES ('ff0be8c7-fe58-4d73-bcf3-67ae2b51356f', '79758f94-70d8-4a11-b05f-793fcb57e093', '712e8ca2-37e2-46eb-98e6-1615f66d0fd7', '2025-06-08 15:46:13.222819+00');
INSERT INTO public.user_badges VALUES ('c71371e0-0f5e-4b19-9519-9bf74daf94d7', '79758f94-70d8-4a11-b05f-793fcb57e093', '5db810f1-1a7f-4022-9ca1-85b3b727db56', '2025-06-08 15:46:13.225032+00');
INSERT INTO public.user_badges VALUES ('9389d4b8-efb6-4757-96fd-0d7524010f19', '56e89866-3a90-4f97-8461-34b24f842da4', '712e8ca2-37e2-46eb-98e6-1615f66d0fd7', '2025-06-08 15:46:13.228452+00');
INSERT INTO public.user_badges VALUES ('923f0858-7119-488c-ae7a-c98446dbf760', 'd943512c-a462-4f9e-8881-13dc2b9b59e1', '712e8ca2-37e2-46eb-98e6-1615f66d0fd7', '2025-06-08 15:46:13.231362+00');
INSERT INTO public.user_badges VALUES ('1a1f943f-206e-404c-ac32-481c5b621b67', 'a03b3300-0d14-454e-8112-7b3eada90ae9', 'd67d0bd3-06ad-4f4c-bf43-09ebcb3240a0', '2025-06-08 16:14:08.060812+00');
INSERT INTO public.user_badges VALUES ('da2fdf46-f9e6-4aa7-8725-84b123f1dfa9', 'a03b3300-0d14-454e-8112-7b3eada90ae9', '4fe90924-2c53-4882-b404-e27d066ef682', '2025-06-08 16:14:08.063959+00');
INSERT INTO public.user_badges VALUES ('f9501491-ca83-45d2-b548-5c9ea50d7551', '6929aaa9-f3a8-4e28-b959-f6585f65b6b1', '4fe90924-2c53-4882-b404-e27d066ef682', '2025-06-08 16:14:08.068683+00');
INSERT INTO public.user_badges VALUES ('4973cff6-75d1-4b3e-bc57-b97a685c8dee', '092b44d9-e9d7-4a40-a59f-d868d193dd0d', '4fe90924-2c53-4882-b404-e27d066ef682', '2025-06-08 16:14:08.074228+00');
INSERT INTO public.user_badges VALUES ('11407bfa-3846-44ef-a3a6-67b494e3b907', '3b0ebec9-3805-490b-987f-4c6e7286319e', 'd67d0bd3-06ad-4f4c-bf43-09ebcb3240a0', '2025-06-08 16:14:08.078025+00');
INSERT INTO public.user_badges VALUES ('02c34ec3-38f3-4cc1-9c93-806883cedd0f', '3b0ebec9-3805-490b-987f-4c6e7286319e', '4fe90924-2c53-4882-b404-e27d066ef682', '2025-06-08 16:14:08.080191+00');
INSERT INTO public.user_badges VALUES ('b1175a0c-9e74-4d8b-a205-2cb42528200d', '6e137217-71a9-439b-a579-4bcd1a729af9', '4fe90924-2c53-4882-b404-e27d066ef682', '2025-06-08 16:14:08.086398+00');
INSERT INTO public.user_badges VALUES ('0328ba87-fad3-41ff-9881-1a630bc0ec1f', '79758f94-70d8-4a11-b05f-793fcb57e093', '4fe90924-2c53-4882-b404-e27d066ef682', '2025-06-08 16:14:08.090491+00');
INSERT INTO public.user_badges VALUES ('a9a8903d-6a30-4dec-8141-6680ffca5c12', '56e89866-3a90-4f97-8461-34b24f842da4', 'd67d0bd3-06ad-4f4c-bf43-09ebcb3240a0', '2025-06-08 16:14:08.094236+00');
INSERT INTO public.user_badges VALUES ('f9622581-229d-4377-ac6e-b205581eafd2', '56e89866-3a90-4f97-8461-34b24f842da4', '4fe90924-2c53-4882-b404-e27d066ef682', '2025-06-08 16:14:08.096447+00');


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users VALUES ('f399a996-1567-4fc0-a86a-92774494ca44', 'johndoe', 'johndoe@example.com', '$2a$10$5Ai8T0Vueev0ll6cKobWhOhkWG/Hca/tYuiLFzoz0tj44ZXScvMPa', 'user', 'http://localhost:9000/ipromise/80ed7c51-695d-4f6a-ad7b-ecb3f79a44ba.jpg', 'Music, motivation, and making it happen.', '2025-05-25 20:28:15.427525+00', '2025-06-08 16:46:10.475315+00', NULL);
INSERT INTO public.users VALUES ('092b44d9-e9d7-4a40-a59f-d868d193dd0d', 'michaelbrown', 'michaelbrown@example.com', '$2a$10$RnPMerRGCDnb9E7Wq6hRmuIqSjVmjLf7ZISrsKK2J9k62WOhmitFq', 'user', 'http://localhost:9000/ipromise/44ae11d5-10f6-43b1-9ef8-9bc12283faa7.jpg', 'Productive mornings, strong coffee, and weekend hikes.', '2025-05-25 20:28:56.632514+00', '2025-06-08 16:47:03.425325+00', NULL);
INSERT INTO public.users VALUES ('3b0ebec9-3805-490b-987f-4c6e7286319e', 'sarahlee', 'sarahlee@example.com', '$2a$10$a6TS0FN4IIc1i35ZjgX6WOvuPlI.E6dTyuRzEDOQ8/lojTAbtd3hq', 'user', 'http://localhost:9000/ipromise/8f0e65a2-569d-4660-876c-e34f820b1d22.jpg', 'Designer by day, plant mom by night.', '2025-05-25 20:29:02.859366+00', '2025-06-08 16:40:27.814247+00', NULL);
INSERT INTO public.users VALUES ('56e89866-3a90-4f97-8461-34b24f842da4', 'emanuelneuer', 'emanuelneuer@example.com', '$2a$10$KyEdttw7eTelHj7p3bwCzuW40AJbpkM4eCS1M0gPQsmcNMO0ztwGy', 'user', 'http://localhost:9000/ipromise/e6e344b9-3bce-4ffd-a17d-b5cc14b39c74.jpg', 'Keeper of goals, lover of clean sheets.', '2025-06-04 19:20:34.356125+00', '2025-06-08 16:40:57.005873+00', NULL);
INSERT INTO public.users VALUES ('6929aaa9-f3a8-4e28-b959-f6585f65b6b1', 'alexsmith', 'alexsmith@example.com', '$2a$10$2uPDDCCPZ6nNqV.k7yqxZO3bZRncpWd0y/Cp1F0birRwibcfCoNqK', 'user', 'http://localhost:9000/ipromise/8f250274-a5cf-400b-95d8-a31e4c38955f.jpg', 'Fitness. Code. Repeat.', '2025-05-25 20:28:47.413388+00', '2025-06-08 16:42:38.040308+00', NULL);
INSERT INTO public.users VALUES ('6e137217-71a9-439b-a579-4bcd1a729af9', 'danielmartin', 'danielmartin@example.com', '$2a$10$.4yRj.ICdkXJHNehph5fRu4A5MeNtBULMWpEvEZ/DaEZKphc9vOrK', 'user', 'http://localhost:9000/ipromise/6ea78a78-8ac9-491d-8cfe-ed485e5f4c01.jpg', 'Books, bikes & backend APIs.', '2025-05-25 20:29:16.99026+00', '2025-06-08 16:43:02.027876+00', NULL);
INSERT INTO public.users VALUES ('79758f94-70d8-4a11-b05f-793fcb57e093', 'lauragarcia', 'lauragarcia@example.com', '$2a$10$iQdfkkzgZOOvZRSUEYDuze1m/7xvsr6jli3JJHxfvBKOYVp8OsLQy', 'user', 'http://localhost:9000/ipromise/888a18c4-ac96-43a7-98a9-1c6b45d520ff.jpg', 'Capturing moments one photo at a time.', '2025-05-25 20:29:20.474495+00', '2025-06-08 16:43:28.030584+00', NULL);
INSERT INTO public.users VALUES ('a03b3300-0d14-454e-8112-7b3eada90ae9', 'janedoe', 'janedoe@example.com', '$2a$10$POkvoB2up.f0tFVShVuSyOIgSteS5A3/YKbw3yWU6Hm37w1bO7ztS', 'user', 'http://localhost:9000/ipromise/a772e9c4-6b63-47d8-8d8d-49b95a630e21.jpg', 'Doodles, coffee & deep thoughts.', '2025-05-25 20:28:41.77735+00', '2025-06-08 16:43:51.687264+00', NULL);
INSERT INTO public.users VALUES ('a203c89a-3ab9-42b2-ae92-0c979d8425de', 'nataliewilson', 'nataliewilson@example.com', '$2a$10$B2PzfU7hK0oETWeZ2aTXSeSO4ClexCTT9HJNUDjmF/xkzkp6WuAtC', 'user', 'http://localhost:9000/ipromise/af40e94d-8961-45b6-95cc-13fdff717704.jpg', 'Always planning the next adventure.', '2025-05-25 20:29:11.983565+00', '2025-06-08 16:44:12.165928+00', NULL);
INSERT INTO public.users VALUES ('ae884ed2-075a-4ed2-9dde-736921d80e9f', 'davidclark', 'davidclark@example.com', '$2a$10$ETRwUmBOQxkYgukM6X0DA.I1F5UpGGyT1C0DgwsB8DpRjnzcdwmI6', 'user', 'http://localhost:9000/ipromise/019511ab-4a38-4a10-9e0d-b6bdb8c6befd.jpg', 'Writing code & chasing bugs.', '2025-05-25 20:29:07.36962+00', '2025-06-08 16:44:44.012578+00', NULL);
INSERT INTO public.users VALUES ('d440ff73-81ae-4a7c-b7cb-3336e3b24a4c', 'emilyjones', 'emilyjones@example.com', '$2a$10$KTUBOvC1hurqcGSJJ8IqNuKf/BAhK2HMZjwWoEzrawIz81l1P3G1K', 'user', 'http://localhost:9000/ipromise/0a5f1083-c1bc-43da-bab3-7aeb2ffee375.jpg', 'Dog lover & weekend baker.', '2025-05-25 20:28:51.592047+00', '2025-06-08 16:45:10.233377+00', NULL);
INSERT INTO public.users VALUES ('d943512c-a462-4f9e-8881-13dc2b9b59e1', 'raxarbek', 'raxarbek@example.com', '$2a$10$2gPsgmsSPVyix/xroKyBousthYebDTd.7WXhunCh4sDx0gCXelPZu', 'user', 'http://localhost:9000/ipromise/7b21eafa-b1bf-4745-be9e-59b2e8e838c0.jpg', 'Learning every day. Sharing what I know.', '2025-06-04 21:33:46.271911+00', '2025-06-08 16:45:43.186382+00', NULL);


--
-- Name: attachments attachments_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.attachments
    ADD CONSTRAINT attachments_pkey PRIMARY KEY (id);


--
-- Name: badges badges_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.badges
    ADD CONSTRAINT badges_pkey PRIMARY KEY (id);


--
-- Name: followers followers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.followers
    ADD CONSTRAINT followers_pkey PRIMARY KEY (id);


--
-- Name: likes likes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.likes
    ADD CONSTRAINT likes_pkey PRIMARY KEY (id);


--
-- Name: microtasks microtasks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.microtasks
    ADD CONSTRAINT microtasks_pkey PRIMARY KEY (id);


--
-- Name: notifications notifications_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT notifications_pkey PRIMARY KEY (id);


--
-- Name: posts posts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_pkey PRIMARY KEY (id);


--
-- Name: predictions predictions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.predictions
    ADD CONSTRAINT predictions_pkey PRIMARY KEY (id);


--
-- Name: promises promises_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.promises
    ADD CONSTRAINT promises_pkey PRIMARY KEY (id);


--
-- Name: refresh_tokens refresh_tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.refresh_tokens
    ADD CONSTRAINT refresh_tokens_pkey PRIMARY KEY (id);


--
-- Name: badges uni_badges_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.badges
    ADD CONSTRAINT uni_badges_code UNIQUE (code);


--
-- Name: refresh_tokens uni_refresh_tokens_token; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.refresh_tokens
    ADD CONSTRAINT uni_refresh_tokens_token UNIQUE (token);


--
-- Name: users uni_users_email; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT uni_users_email UNIQUE (email);


--
-- Name: users uni_users_username; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT uni_users_username UNIQUE (username);


--
-- Name: user_badges user_badges_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_badges
    ADD CONSTRAINT user_badges_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_attachments_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_attachments_deleted_at ON public.attachments USING btree (deleted_at);


--
-- Name: idx_attachments_post_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_attachments_post_id ON public.attachments USING btree (post_id);


--
-- Name: idx_attachments_user_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_attachments_user_id ON public.attachments USING btree (user_id);


--
-- Name: idx_follower_following; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_follower_following ON public.followers USING btree (follower_id, following_id);


--
-- Name: idx_likes_post_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_likes_post_id ON public.likes USING btree (post_id);


--
-- Name: idx_microtasks_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_microtasks_deleted_at ON public.microtasks USING btree (deleted_at);


--
-- Name: idx_microtasks_promise_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_microtasks_promise_id ON public.microtasks USING btree (promise_id);


--
-- Name: idx_notifications_user_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_notifications_user_id ON public.notifications USING btree (user_id);


--
-- Name: idx_posts_created_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_posts_created_at ON public.posts USING btree (created_at);


--
-- Name: idx_posts_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_posts_deleted_at ON public.posts USING btree (deleted_at);


--
-- Name: idx_posts_microtask_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_posts_microtask_id ON public.posts USING btree (microtask_id);


--
-- Name: idx_posts_parent_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_posts_parent_id ON public.posts USING btree (parent_id);


--
-- Name: idx_posts_promise_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_posts_promise_id ON public.posts USING btree (promise_id);


--
-- Name: idx_posts_user_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_posts_user_id ON public.posts USING btree (user_id);


--
-- Name: idx_predictions_promise_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_predictions_promise_id ON public.predictions USING btree (promise_id);


--
-- Name: idx_promises_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_promises_deleted_at ON public.promises USING btree (deleted_at);


--
-- Name: idx_promises_user_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_promises_user_id ON public.promises USING btree (user_id);


--
-- Name: idx_refresh_tokens_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_refresh_tokens_deleted_at ON public.refresh_tokens USING btree (deleted_at);


--
-- Name: idx_refresh_tokens_user_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_refresh_tokens_user_id ON public.refresh_tokens USING btree (user_id);


--
-- Name: idx_user_badge; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_user_badge ON public.user_badges USING btree (user_id, badge_id);


--
-- Name: idx_user_post; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_user_post ON public.likes USING btree (user_id, post_id);


--
-- Name: idx_users_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


--
-- PostgreSQL database dump complete
--

