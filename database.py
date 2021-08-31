from sqlalchemy import create_engine, Column, Integer, String, Numeric, Float, Table, MetaData
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker
from sqlalchemy_utils import database_exists, create_database
import pandas as pd
from typing import Dict
from docker_library import is_container_running, run_container




run_container("postgres", "pokemon-postgres")
# look at docker_library for explanation of Docker
container_running = is_container_running('pokemon-postgres')




# if not container_running:
#     pass


# client.containers.run("postgres", detach=True, ports=[5432])


# Manual Command to run a docker container
# docker run --name pokemon-postgres -e POSTGRES_PASSWORD=pokemon -d -p 5432:5432 postgres


# SQL Lite
# engine = create_engine('sqlite:///:memory:')

# Postgres
engine = create_engine('postgresql://postgres:pokemon@localhost:5432/Pokemon')

if not database_exists(engine.url):
    create_database(engine.url)

# My SQL
# engine = create_engine('mysql+pymydsql://root@localhost/mydb')


base = declarative_base()

pokemon_csv = "./Pokemon.csv"

# explanation of csv reader
# https://www.delftstack.com/howto/python/python-csv-to-dictionary/
items = pd.read_csv(pokemon_csv, index_col=0, sep=",", encoding='cp1252')
# items = pd.read_csv(pokemon_csv, index_col=0, sep=",")

# https://pandas.pydata.org/docs/reference/api/pandas.DataFrame.transpose.html?highlight=transpose#pandas.DataFrame.transpose
# transpose flips the keys to be row 0 instead of column 0
pokemon_csv_dict = items.transpose().to_dict(orient='series')


class Pokemon(base):
    # A class using Declarative at a minimum needs a __tablename__ attribute,
    # and at least one Column which is part of a primary key
    __tablename__ = 'Pokedex'
    dexnum = Column(Integer, primary_key=True)
    # some databases require len of stings
    name = Column(String(11))
    type1 = Column(String(10))
    type2 = Column(String(10))
    stage = Column(String(12))
    evolve_level = Column(Integer)
    gender_ratio = Column(String(10))
    height = Column(Float(3))
    weight = Column(Float(3))
    description = Column(String(125))
    category = Column(String(20))
    lvl_speed = Column(Float(3))
    base_exp = Column(Integer)
    catch_rate = Column(Integer)

    # __init__ where are you?

    # Our User class, as defined using the Declarative system, : from Base in __init__.py
    # has been provided with a constructor (e.g. __init__() method)
    # which automatically accepts keyword names that match the columns we’ve mapped.
    # We are free to define any explicit __init__() method we prefer on our class,
    # which will override the default method provided by Declarative.

    # __repr__ :  special method used to represent a class’s objects as a string.
    def __repr__(self):
        return "<Pokemon(dexnum=%d, name='%s', type1='%s', type2='%s', evolve_level=%d," \
               "gender_ratio='%s', height=%f, weight=%f, category='%s'," \
               "lvl_speed=%f, base_exp=%d, catch_rate=%d" \
               "description='%s')>" % \
               (self.dexnum, self.name, self.type1, self.type2, self.evolve_level,
                self.gender_ratio, self.height, self.weight, self.category, self.lvl_speed,
                self.base_exp, self.catch_rate, self.description)


def createPokemon(dexnum: int, pokemon_csv_dict: Dict) -> Pokemon:
    return Pokemon(
        dexnum=dexnum,
        name=pokemon_csv_dict[dexnum]['Pokemon_Name'],
        type1=pokemon_csv_dict[dexnum]['Type1'],
        type2=pokemon_csv_dict[dexnum]['Type2'] if pokemon_csv_dict[dexnum]['Type2'] != "-" else "None",
        stage=pokemon_csv_dict[dexnum]['Stage'],
        evolve_level=int(pokemon_csv_dict[dexnum]['Evolve_Level']),
        gender_ratio=pokemon_csv_dict[dexnum]['Gender_Ratio'],
        height=float(pokemon_csv_dict[dexnum]['Height']),
        weight=float(pokemon_csv_dict[dexnum]['Weight']),
        description=pokemon_csv_dict[dexnum]['Description'],
        category=pokemon_csv_dict[dexnum]['Category'],
        lvl_speed=float(pokemon_csv_dict[dexnum]['Leveling_Speed']),
        base_exp=int(pokemon_csv_dict[dexnum]['Base_Exp']),
        catch_rate=int(pokemon_csv_dict[dexnum]['Catch_Rate']),
    )




metadata = MetaData()

Pokedex_table = Table('Pokedex', metadata,
    Column('dexnum', Integer, primary_key=True),
    Column('name', String(11), nullable=False),
    Column('type1', String(10), nullable=False),
    Column('type2', String(10), nullable=True),
    Column('stage', String(12), nullable=False),
    Column('evolve_level', Integer, nullable=False),
    Column('gender_ratio', String(5), nullable=False),
    Column('height', Float(3), nullable=False),
    Column('weight', Float(3), nullable=False),
    Column('description', String(125), nullable=False),
    Column('category', String(20), nullable=False),
    Column('lvl_speed', Float(3), nullable=False),
    Column('base_exp', Integer, nullable=False),
    Column('catch_rate', Integer, nullable=False)
                      )

try:
    Pokemon.__table__.drop(engine)
except Exception as e:
    print(e)


base.metadata.create_all(engine)

Session = sessionmaker(bind=engine)
session = Session()

for i in range(1, 152):
    created_pokemon = createPokemon(i, pokemon_csv_dict)
    session.add(created_pokemon)
    session.commit()

query = session.query(Pokemon)
# get grabs ID of 1
P1 = query.get(1)
print(P1)

# user_count = session.query(Pokemon).count()
# print(f"User Count Before Delete: {user_count}")
#
# # Id Of 1 no longer exists
# session.delete(P1)
# session.commit()
#
# user_count = session.query(Pokemon).count()
# print(f"User Count After Delete: {user_count}")
#
#
# query = session.query(Pokemon)
# P1 = query.get(2)
# print(P1)
