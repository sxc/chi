CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    user_id INT UNIQUE REFERENCES users(id) on delete cascade,
    token_hash TEXT UNIQUE NOT NULL
    );


ALTER TABLE sessions
 ADD CONSTRAINT sessions_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id)
  ON DELETE CASCADE;


    insert into sessions (user_id, token_hash) 
    values ($1, $2)
    returning id;


    
	update sessions set token_hash = "abc123"
	where user_id = 1
    returning id;
    ;`, session.UserID, session.TokenHash

    "pPZeJJGFxXqO2xU7o2pL0xoY_vsGP8f36tGg8RApNMA="



    DELETE FROM sessions
    WHERE token_hash = $1