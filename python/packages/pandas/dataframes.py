import pandas as pd

# Creating a DataFrame from a dictionary
data = {
    'Neighborhood': ['Downtown', 'East Nashville', 'The Gulch', 'Germantown'],
    'Avg Home Price (k)': [550, 680, 750, 710],
    'Has Coffee Shop': [True, True, True, False]
}
nashville_df = pd.DataFrame(data)

# Setting a meaningful index (optional, but often useful)
# nashville_df = nashville_df.set_index('Neighborhood') # Example if you wanted Neighborhood as the index

print("\nNashville Neighborhood Data (DataFrame):")
print(nashville_df)
