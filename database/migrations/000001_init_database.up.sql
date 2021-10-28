-- Install Extensions --------------------------------------------
------------------------------------------------------------------

CREATE EXTENSION IF NOT EXISTS timescaledb;
CREATE EXTENSION IF NOT EXISTS pgcrypto;

ALTER DATABASE "analytics" SET timescaledb.telemetry_level = 'off';

CREATE SCHEMA IF NOT EXISTS data;

-- CREATE updated_at trigger -------------------------------------
------------------------------------------------------------------

CREATE OR REPLACE FUNCTION data.trigger_set_timestamp()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;