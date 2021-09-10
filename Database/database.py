from sqlalchemy import create_engine, Column, Integer, String, Float, MetaData, ForeignKey
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker, relationship, backref
import pandas as pd
from typing import Dict
from Docker.docker_library import is_container_running, run_container
import os


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
# Driver :// user : password @ hostname(uri) : port / Database
engine = create_engine('postgresql://postgres:pokemon@localhost:5432/Pokemon')

# print("HERE")
# abbreviated_database = 'postgresql://postgres:pokemon@localhost:5432/Pokemon'
# if not database_exists(engine.url):
#     create_database(abbreviated_database)
#     # conn = engine.connect()
#     # conn.execute("commit")
#     # conn.execute("create database Pokemon")
#     # conn.close()

# My SQL
# engine = create_engine('mysql+pymydsql://root@localhost/mydb')


base = declarative_base()


# grabs route correctly independent of OS routing:
#   Linux/Mac: ..//CSV//Pokemon.csv
#   Windows: ..\\CSV\\Pokemon.csv
pokemon_csv = os.path.join('..', 'CSV', "Pokemon.csv")
base_stats_csv = os.path.join('..', 'CSV', "Base_Stats.csv")



# explanation of csv reader
# https://www.delftstack.com/howto/python/python-csv-to-dictionary/
pokedex_items = pd.read_csv(pokemon_csv, index_col=0, sep=",", encoding='cp1252')
base_stats_items = pd.read_csv(base_stats_csv, index_col=0, sep=",", encoding='cp1252')
# items = pd.read_csv(pokemon_csv, index_col=0, sep=",")

# https://pandas.pydata.org/docs/reference/api/pandas.DataFrame.transpose.html?highlight=transpose#pandas.DataFrame.transpose
# transpose flips the keys to be row 0 instead of column 0
pokemon_csv_dict = pokedex_items.transpose().to_dict(orient='series')
base_stats_dict = base_stats_items.transpose().to_dict(orient='series')


class Pokemon(base):
    # A class using Declarative at a minimum needs a __tablename__ attribute,
    # and at least one Column which is part of a primary key
    __tablename__ = 'Pokedex'
    dexnum = Column(Integer, primary_key=True)
    # some databases require len of stings
    name = Column(String(11), nullable=False)
    type1 = Column(String(10), nullable=False)
    type2 = Column(String(10), nullable=True)
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

    # child = relationship("Child", back_populates="Parent", uselist=False)

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


class Base_Stats(base):
    __tablename__ = 'Stats'
    dexnum = Column(Integer, primary_key=True)
    hp = Column(Integer)
    attack = Column(Integer)
    defense = Column(Integer)
    special_attack = Column(Integer)
    special_defense = Column(Integer)
    speed = Column(Integer)
    # total = Column(Integer)
    # dexnum = Column(Integer, ForeignKey(Pokemon.dexnum), primary_key=True)

    # parent = relationship("Parent", backref=backref("child", uselist=False))

    def __repr__(self):
        return "<Base_Stats(dexnum=%d, hp=%d, attack=%d, defense=%d," \
               "special_attack=%d, special_defense=%d, speed=%d)>" % \
               (self.dexnum, self.hp, self.attack, self.defense, self.special_attack,
                self.special_defense, self.speed)


class User(base):
    __tablename__='Users'
    id = Column(Integer, primary_key=True)
    name = Column(String(50))
    email = Column(String(25))
    password = Column(String(25))
    telephone = Column(String(10))
    address = Column(String(250))
    city = Column(String(25))
    state = Column(String(25))
    zipcode = Column(Integer)

    def __repr__(self):
        return "<User(id=%d, name='%s', email='%s', password='%s', telephone='%s'," \
               "address='%s'," \
               "city='%s', state='%s', zipcode=%d)>" % \
               (self.id, self.name, self.email, self.password, self.telephone,
                self.address, self.city, self.state, self.zipcode)


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


def createBaseStats(dexnum: int, base_stats_csv_dict: Dict) -> Base_Stats:
    return Base_Stats(
        dexnum=dexnum,
        hp=int(base_stats_dict[dexnum]['HP']),
        attack=int(base_stats_dict[dexnum]['Attack']),
        defense=int(base_stats_dict[dexnum]['Defense']),
        special_attack=int(base_stats_dict[dexnum]['Sp. Atk']),
        special_defense=int(base_stats_dict[dexnum]['Sp. Def']),
        speed=int(base_stats_dict[dexnum]['Speed']),
                      )


metadata = MetaData()


try:
    print("Dropping Table on startup")
    Pokemon.__table__.drop(engine)
    Base_Stats.__table__.drop(engine)
    User.__table__.drop(engine)
except Exception as e:
    print(f"Failed to Drop: Table doesn't exists")


# # ERROR on first run
base.metadata.create_all(engine)

Session = sessionmaker(bind=engine)
session = Session()


for i in range(1, 152):
    created_pokemon = createPokemon(i, pokemon_csv_dict)
    session.add(created_pokemon)
    session.commit()
    created_base_stats = createBaseStats(i, base_stats_dict)
    session.add(created_base_stats)
    session.commit()

query = session.query(Pokemon)
# get grabs ID of 1
P1 = query.get(1)
print(P1)

query = session.query(Base_Stats)
B1 = query.get(1)
print(B1)

user_count = session.query(Pokemon).count()
print(f"Pokemon Count Before Delete: {user_count}")

base_count = session.query(Base_Stats).count()
print(f"Base Stat Count Before Delete: {base_count}")


# Id Of 1 no longer exists
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
