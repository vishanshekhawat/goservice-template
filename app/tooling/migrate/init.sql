CREATE TABLE users
(
    id              INT unsigned NOT NULL AUTO_INCREMENT, # Unique ID for the record
    name            VARCHAR(150) NOT NULL,                # Name of the cat
    city           VARCHAR(150) NOT NULL,                # City of the user
    email           VARCHAR(150) NOT NULL,                      # Email of the user
    PRIMARY KEY     (id)                                  # Make the id the primary key
);
