CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    user_id INT UNIQUE,
    token_hash TEXT UNIQUE NOT NULL);


    insert into sessions (user_id, token_hash) 
    values ($1, $2)
    returning id;


    
	update sessions set token_hash = "abc123"
	where user_id = 1
    returning id;
    ;`, session.UserID, session.TokenHash

    "pPZeJJGFxXqO2xU7o2pL0xoY_vsGP8f36tGg8RApNMA="