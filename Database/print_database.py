from sqlalchemy import create_engine, inspect
from sqlalchemy.orm import sessionmaker, declarative_base


Base = declarative_base()

engine = create_engine('postgresql://postgres:pokemon@localhost:5432/Pokemon')

Session = sessionmaker(bind=engine)
inspector = inspect(engine)

schemas = inspector.get_schema_names()
for schema in schemas:
    print(f"schema: {schema}")
    for table_name in inspector.get_table_names(schema=schema):
        print(f"table name: {table_name}")
        for column in inspector.get_columns(table_name, schema=schema):
            print(f"Column: {column}")
