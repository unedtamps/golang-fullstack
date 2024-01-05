-- name: CreateUser :execrows

INSERT INTO
    "users"(
        id,
        first_name,
        last_name,
        email,
        password,
        image_url,
        description
    )
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: CreateMessage :one

INSERT INTO
    "messages" (
        id,
        connection_id,
        message_content,
        user_id
    )
VALUES ($1, $2, $3, $4) RETURNING id;

-- name: CreateUserConnectionWith :one

INSERT INTO
    "user_connections" (user_1, user_2, id)
VALUES ($1, $2, $3) RETURNING id,
    status;

-- name: CreateAttachment :one

INSERT INTO
    "attachment" (id, connection_id, id_messages)
VALUES ($1, $2, $3) RETURNING id,
    attachment_url;

-- name: CreateGroupParticiPants :execrows

INSERT INTO "group_participants" VALUES($1,$2);

-- name: CreateGroup :one

INSERT INTO
    "groups" (id, name, description)
VALUES ($1, $2, $3) RETURNING id;

-- name: CreateGroupMessages :one

INSERT INTO
    "group_messages" (
        id,
        user_id,
        group_id,
        message_content
    )
VALUES ($1, $2, $3, $4) RETURNING id;

-- name: CreateGroupAttachement :one

INSERT INTO
    "group_attachment"(
        id,
        attachment_url,
        group_id,
        id_message
    )
VALUES ($1, $2, $3, $4) RETURNING id,
    attachment_url;
