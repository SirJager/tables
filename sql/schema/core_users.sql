CREATE TABLE "core_users" (
  "id" BIGSERIAL PRIMARY KEY NOT NULL,
  "email" VARCHAR(255) UNIQUE NOT NULL,
  "username" VARCHAR(60) UNIQUE NOT NULL,
  "password" VARCHAR NOT NULL,
  "fullname" VARCHAR(255) NOT NULL,
  "public" BOOLEAN NOT NULL DEFAULT FALSE,
  "verified" BOOLEAN NOT NULL DEFAULT FALSE,
  "blocked" BOOLEAN NOT NULL DEFAULT FALSE,
  "updated" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "created" TIMESTAMPTZ NOT NULL DEFAULT (now())
);


COMMENT ON COLUMN "core_users"."id" IS 'server-side auto generated id';
COMMENT ON COLUMN "core_users"."email" IS 'required,max=255';
COMMENT ON COLUMN "core_users"."username" IS 'required,alphanumeric,min=3,max=60';
COMMENT ON COLUMN "core_users"."password" IS 'required,hashed';
COMMENT ON COLUMN "core_users"."fullname" IS 'optional,max=255';
COMMENT ON COLUMN "public"."core_users"."public" IS 'optinal,default=false,desc=profile visible by others or not';
COMMENT ON COLUMN "public"."core_users"."blocked" IS 'optinal,default=false,desc=profile is accessible or not';
COMMENT ON COLUMN "public"."core_users"."verified" IS 'optinal,default=false,desc=email is verified or not';
COMMENT ON COLUMN "core_users"."updated" IS 'server-side auto generated timestamp with time zone';
COMMENT ON COLUMN "core_users"."created" IS 'server-side auto generated timestamp with time zone';

