
CREATE TABLE
    "group_messages"(
        "id" CHAR(26),
        "user_id" CHAR(26) NOT NULL,
        "group_id" CHAR(26) NOT NULL,
        "message_content" TEXT NOT NULL,
        "created_at" TIMESTAMP
        WITH
            TIME zone DEFAULT CURRENT_TIMESTAMP,
            "updated_at" TIMESTAMP
        WITH
            TIME zone DEFAULT CURRENT_TIMESTAMP
    );

ALTER TABLE "group_messages" ADD PRIMARY KEY("id");

CREATE TABLE
    "messages"(
        "id" CHAR(26),
        "connection_id" CHAR(26) NOT NULL,
        "created_at" TIMESTAMP
        WITH
            TIME zone DEFAULT CURRENT_TIMESTAMP,
            "updated_at" TIMESTAMP
        WITH
            TIME zone DEFAULT CURRENT_TIMESTAMP,
            "message_content" TEXT,
            "user_id" CHAR(26) NOT NULL
    );

ALTER TABLE "messages" ADD PRIMARY KEY("id");

CREATE TYPE status as ENUM ('active', 'block', 'mute');

CREATE TABLE
    "user_connections"(
        "user_1" CHAR(26) NOT NULL,
        "user_2" CHAR(26) NOT NULL,
        "id" CHAR(26) NOT NULL,
        "created_at" TIMESTAMP
        WITH
            TIME zone DEFAULT CURRENT_TIMESTAMP,
            "updated_at" TIMESTAMP
        WITH
            TIME zone DEFAULT CURRENT_TIMESTAMP,
            "status" status DEFAULT 'active'
    );

CREATE INDEX
    "user_connections_user_1_index" ON "user_connections"("user_1");

CREATE INDEX
    "user_connections_user_2_index" ON "user_connections"("user_2");

ALTER TABLE "user_connections" ADD PRIMARY KEY("id");

CREATE TABLE
    "attachment"(
        "created_at" TIMESTAMP
        WITH
            TIME zone DEFAULT CURRENT_TIMESTAMP,
            "updated_at" TIMESTAMP
        WITH
            TIME zone DEFAULT CURRENT_TIMESTAMP,
            "id_messages" CHAR(26) NOT NULL,
            "attachment_url" VARCHAR(512) NOT NULL,
            "id" CHAR(26) NOT NULL,
            "connection_id" CHAR(26) NOT NULL
    );

ALTER TABLE "attachment" ADD PRIMARY KEY("id");

CREATE TABLE
    "groups"(
        "id" CHAR(26) NOT NULL,
        "name" VARCHAR(255) NOT NULL,
        "description" TEXT,
        "created_at" TIMESTAMP
        WITH
            TIME zone DEFAULT CURRENT_TIMESTAMP,
            "updated_at" TIMESTAMP
        WITH
            TIME zone DEFAULT CURRENT_TIMESTAMP 
    );

ALTER TABLE "groups" ADD PRIMARY KEY("id");

CREATE TABLE
    "group_participants"(
        "group_id" CHAR(26) NOT NULL,
        "user_id" CHAR(26) NOT NULL,
        PRIMARY KEY ("group_id", "user_id")
    );

CREATE TABLE
    "group_attachment"(
        "id" CHAR(26) NOT NULL,
        "attachment_url" VARCHAR(512) NOT NULL,
        "created_at" TIMESTAMP
        WITH
            TIME zone DEFAULT CURRENT_TIMESTAMP,
            "updated_at" TIMESTAMP
        WITH
            TIME zone DEFAULT CURRENT_TIMESTAMP, 
            "group_id" CHAR(26) NOT NULL,
            "id_message" CHAR(26) NOT NULL
    );

ALTER TABLE "group_attachment" ADD PRIMARY KEY("id");

CREATE TABLE
    "users"(
        "id" CHAR(26) NOT NULL,
        "first_name" VARCHAR(255) NOT NULL,
        "last_name" VARCHAR(255) NOT NULL,
        "email" VARCHAR(255) NOT NULL,
        "password" VARCHAR(255) NOT NULL,
        "created_at" TIMESTAMP
        WITH
            TIME zone DEFAULT CURRENT_TIMESTAMP,
            "updated_at" TIMESTAMP
        WITH
            TIME zone DEFAULT CURRENT_TIMESTAMP,
            "image_url" VARCHAR(512),
            "description" TEXT
    );

ALTER TABLE "users" ADD PRIMARY KEY("id");

ALTER TABLE "group_attachment"
ADD
    CONSTRAINT "group_attachment_group_id_foreign" FOREIGN KEY("group_id") REFERENCES "groups"("id") ON DELETE CASCADE;

ALTER TABLE "attachment"
ADD
    CONSTRAINT "attachment_id_messages_foreign" FOREIGN KEY("id_messages") REFERENCES "messages"("id") ON DELETE CASCADE;

ALTER TABLE "user_connections"
ADD
    CONSTRAINT "user_connections_user_2_foreign" FOREIGN KEY("user_2") REFERENCES "users"("id") ON DELETE CASCADE;

ALTER TABLE "attachment"
ADD
    CONSTRAINT "attachment_connection_id_foreign" FOREIGN KEY("connection_id") REFERENCES "user_connections"("id") ON DELETE CASCADE;

ALTER TABLE "group_attachment"
ADD
    CONSTRAINT "group_attachment_id_message_foreign" FOREIGN KEY("id_message") REFERENCES "group_messages"("id") ON DELETE CASCADE;

ALTER TABLE
    "group_participants"
ADD
    CONSTRAINT "group_participants_group_id_foreign" FOREIGN KEY("group_id") REFERENCES "groups"("id") ON DELETE CASCADE;

ALTER TABLE "group_messages"
ADD
    CONSTRAINT "group_messages_group_id_foreign" FOREIGN KEY("group_id") REFERENCES "groups"("id") ON DELETE CASCADE;

ALTER TABLE "messages"
ADD
    CONSTRAINT "messages_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id") ON DELETE CASCADE;

ALTER TABLE
    "group_participants"
ADD
    CONSTRAINT "group_participants_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id") ON DELETE CASCADE;

ALTER TABLE "user_connections"
ADD
    CONSTRAINT "user_connections_user_1_foreign" FOREIGN KEY("user_1") REFERENCES "users"("id") ON DELETE CASCADE;

ALTER TABLE "messages"
ADD
    CONSTRAINT "messages_connection_id_foreign" FOREIGN KEY("connection_id") REFERENCES "user_connections"("id") ON DELETE CASCADE;