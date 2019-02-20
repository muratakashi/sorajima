CREATE TABLE public.user
(
  id                   bigserial,
  email                character varying(255)      NOT NULL,
  encrypted_password   character varying(255)      NOT NULL,
  first_name           character varying(255)      NOT NULL,
  last_name            character varying(255)      NOT NULL,
  created_at           timestamp without time zone NOT NULL,
  updated_at           timestamp without time zone NOT NULL,
  CONSTRAINT user_pkey PRIMARY KEY (id),
  CONSTRAINT user_unique_key UNIQUE (email)
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
COMMENT ON COLUMN public.user.updated_at         IS '更新日時';

CREATE TABLE public.workhour
(
  id         bigserial,
  user_id    bigint                      NOT NULL,
  work_day   date                        NOT NULL,
  hello      timestamp without time zone NOT NULL,
  goodbye    timestamp without time zone NOT NULL,
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL,
  CONSTRAINT workhour_pkey PRIMARY KEY (id),
  CONSTRAINT workhour_unique_key UNIQUE (user_id, work_day)
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
COMMENT ON COLUMN public.workhour.updated_at IS '更新日時';

CREATE TABLE public.project
(
  id         bigserial
  code       character varying(50) COLLATE pg_catalog."default" NOT NULL,
  name       character varying(200) COLLATE pg_catalog."default" NOT NULL,
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL,
  alias_name character varying(100) COLLATE pg_catalog."default" NOT NULL,
  CONSTRAINT project_pkey PRIMARY KEY (id),
  CONSTRAINT project_unique_key UNIQUE (code)
)
WITH (
  OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.project OWNER to postgres;

COMMENT ON TABLE public.project             IS '受託開発、派遣業務をプロジェクト情報として管理';
COMMENT ON COLUMN public.project.id         IS 'ID';
COMMENT ON COLUMN public.project.code       IS 'プロジェクトに付与した番号';
COMMENT ON COLUMN public.project.name       IS 'プロジェクト名';
COMMENT ON COLUMN public.project.alias_name IS 'プロジェクト通称';
COMMENT ON COLUMN public.project.created_at IS '登録日時';
COMMENT ON COLUMN public.project.updated_at IS '更新日時';
