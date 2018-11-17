CREATE TABLE public.user
(
  id                   bigserial,
  email                character varying(255)      NOT NULL,
  encrypted_password   character varying(255)      NOT NULL,
  first_name           character varying(255)      NOT NULL,
  last_name            character varying(255)      NOT NULL,
  created_at           timestamp without time zone NOT NULL,
  created_by           character varying(255)      NOT NULL,
  updated_at           timestamp without time zone NOT NULL,
  updated_by           character varying(255)      NOT NULL,
  CONSTRAINT users_pkey PRIMARY KEY (id),
  CONSTRAINT users_unique_key UNIQUE (email)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.user OWNER to postgres;

COMMENT ON TABLE public.user                     IS 'ユーザ情報を管理';
COMMENT ON COLUMN public.user.id                 IS 'ID';
COMMENT ON COLUMN public.user.email              IS 'メールアドレス';
COMMENT ON COLUMN public.user.encrypted_password IS 'パスワード';
COMMENT ON COLUMN public.user.first_name         IS '姓';
COMMENT ON COLUMN public.user.last_name          IS '名';
COMMENT ON COLUMN public.user.created_at         IS '登録日時';
COMMENT ON COLUMN public.user.created_by         IS '登録者';
COMMENT ON COLUMN public.user.updated_at         IS '更新日時';
COMMENT ON COLUMN public.user.updated_by         IS '更新者';

CREATE TABLE public.workhour
(
  id         bigserial,
  user_id    bigint                      NOT NULL,
  work_day   date                        NOT NULL,
  hello      timestamp without time zone NOT NULL,
  goodbye    timestamp without time zone NOT NULL,
  created_at timestamp without time zone NOT NULL,
  created_by character varying(255)      NOT NULL,
  updated_at timestamp without time zone NOT NULL,
  updated_by character varying(255)      NOT NULL,
  CONSTRAINT hours_pkey PRIMARY KEY (id),
  CONSTRAINT hours_unique_key UNIQUE (user_id, work_day)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.workhour OWNER to postgres;

COMMENT ON TABLE public.workhour             IS '勤怠を管理';
COMMENT ON COLUMN public.workhour.id         IS 'ID';
COMMENT ON COLUMN public.workhour.user_id    IS 'ユーザーID';
COMMENT ON COLUMN public.workhour.work_day   IS '出社日';
COMMENT ON COLUMN public.workhour.hello      IS '出社時間';
COMMENT ON COLUMN public.workhour.goodbye    IS '退社時間';
COMMENT ON COLUMN public.workhour.created_at IS '登録日時';
COMMENT ON COLUMN public.workhour.created_by IS '登録者';
COMMENT ON COLUMN public.workhour.updated_at IS '更新日時';
COMMENT ON COLUMN public.workhour.updated_by IS '更新者';

