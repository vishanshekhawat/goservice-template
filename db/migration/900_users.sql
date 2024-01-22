CREATE USER reader WITH PASSWORD 'reader_password';
GRANT pg_read_all_data TO reader;

CREATE USER writer WITH PASSWORD 'writer_password';
GRANT pg_read_all_data TO writer;
GRANT pg_write_all_data TO writer;
