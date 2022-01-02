from os.path import exists, join
from os import remove
from Kaggle.kaggleApi import kaggleAPI
import pandas as pd
import zipfile
from sklearn.tree import DecisionTreeRegressor
from sklearn.metrics import mean_absolute_error
from sklearn.model_selection import train_test_split
from sklearn.ensemble import RandomForestRegressor

# filename = "MinimumWage.csv"
# dataset = "brandonconrady/us-minimum-wage-1938-2020"

filename = "melb_data.csv"
dataset = "dansbecker/melbourne-housing-snapshot"

fileroute = join("./", filename)
filezip = filename + ".zip"

file_exists = exists(fileroute)

if not file_exists:
    kaggleAPI(dataset, filename)
    if not exists(fileroute):
        with zipfile.ZipFile(filezip, 'r') as zipref:
            zipref.extractall("./")
        remove(filezip)


df = pd.read_csv(filename)
# print(df)


# Filter rows with missing price values
melbourne_data = df.dropna(axis=0)
# Choose target and features
y = melbourne_data.Price
melbourne_features = ['Rooms', 'Bathroom', 'Landsize', 'Lattitude', 'Longtitude']
X = melbourne_data[melbourne_features]

print(X.describe())

# split data into training and validation data, for both features and target
# The split is based on a random number generator. Supplying a numeric value to
# the random_state argument guarantees we get the same split every time we
# run this script.
train_X, val_X, train_y, val_y = train_test_split(X, y, random_state=0)


# Define model. Specify a number for random_state to ensure same results each run
melbourne_model = DecisionTreeRegressor(random_state=1)
# Fit model
melbourne_model.fit(train_X, train_y)


print("Making predictions for the following 5 houses:")
print(X.head())
print("The predictions are")
print(melbourne_model.predict(X.head()))

# get predicted prices on validation data
val_predictions = melbourne_model.predict(val_X)
print("average absolute error from decision tree")
print(mean_absolute_error(val_y, val_predictions))


def get_mae(max_leaf_nodes, train_X, val_X, train_y, val_y):
    model = DecisionTreeRegressor(max_leaf_nodes=max_leaf_nodes, random_state=0)
    model.fit(train_X, train_y)
    preds_val = model.predict(val_X)
    mae = mean_absolute_error(val_y, preds_val)
    return(mae)


# compare MAE with differing values of max_leaf_nodes
for max_leaf_nodes in [5, 50, 500, 5000]:
    my_mae = get_mae(max_leaf_nodes, train_X, val_X, train_y, val_y)
    print("Max leaf nodes: %d  \t\t Mean Absolute Error:  %d" %(max_leaf_nodes, my_mae))


forest_model = RandomForestRegressor(random_state=1)
forest_model.fit(train_X, train_y)
melb_preds = forest_model.predict(val_X)
print("average absolute error for forest model prediction")
print(mean_absolute_error(val_y, melb_preds))

