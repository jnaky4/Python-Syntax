-- CREATE TABLE
DROP TABLE IF EXISTS Pokemon;
CREATE TABLE Pokemon (
Description,Category,Leveling_Speed,Base_Exp,Catch_Rate
    dexnum SMALLINT PRIMARY KEY,
    name VARCHAR NOT NULL,
    type1 VARCHAR NOT NULL,
    type1 VARCHAR,
    stage VARCHAR NOT NULL,
    evolve_lvl SMALLINT NOT NULL,
    gender_ratio VARCHAR NOT NULL,
    height REAL NOT NULL,
    weight REAL NOT NULL,
    description VARCHAR NOT NULL,
    category VARCHAR NOT NULL,
    leveling_speed FLOAT NOT NULL,
    base_exp SMALLINT NOT NULL,
    catch_rate SMALLINT NOT NULL,
);