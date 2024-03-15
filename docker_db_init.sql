
--CREATE DATABASE postgres;

--CREATE SCHEMA public;

\connect postgres;

CREATE TABLE public.trucker (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    age INTEGER NOT NULL,
    nationality VARCHAR NOT NULL
);


CREATE TABLE public.vehicle (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    brand VARCHAR NOT NULL,
    color VARCHAR NOT NULL,
    year INTEGER NOT NULL,
    plate VARCHAR NOT NULL
);

CREATE TABLE  public.assign (
    ID SERIAL PRIMARY KEY,
    trucker_id INTEGER,
    vehicle_id INTEGER,
    CONSTRAINT fk_trucker FOREIGN KEY (trucker_id) REFERENCES public.trucker (ID),
    CONSTRAINT fk_vehicle FOREIGN KEY (vehicle_id) REFERENCES public.vehicle (ID)
);