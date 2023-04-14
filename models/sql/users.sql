CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL
);


select user.id,
users.email,
users.password_hash,
from sessions
join users on users.id = sessions.user_id
where sessions.token_hash = $1;
