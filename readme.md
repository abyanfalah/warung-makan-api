# warung-makan-api
## Installation
If you wish to run the app in localhost:
```bash
git clone https://github.com/abyanfalah/warung-makan-api.git

cd warung-makan-api

go build

./warung-makan-api
```


## Config
Change the value of constants in main.go to suit your need.
Port can be left blank, the app will automatically initiate
the API at port :8000, and will iterate if the current port
is used.


## Database
```sql

CREATE TABLE public.menu (
    id character varying(60) NOT NULL,
    name character varying(100) NOT NULL,
    price integer,
    stock integer,
    image text
);


CREATE TABLE public.transaction (
    id character varying(60) NOT NULL,
    total_price integer,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone
);


CREATE TABLE public.transaction_detail (
    transaction_id character varying(60),
    menu_id character varying(60),
    qty integer,
    subtotal integer,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone
);


CREATE TABLE public.users (
    id character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    username character varying(32) NOT NULL,
    password text NOT NULL,
    image text
);

ALTER TABLE ONLY public.transaction
    ADD CONSTRAINT transaction_pkey PRIMARY KEY (id);


ALTER TABLE ONLY public.users
    ADD CONSTRAINT unique_username UNIQUE (username);

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);

```
