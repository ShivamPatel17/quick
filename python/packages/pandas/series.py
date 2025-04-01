import pandas as pd  # Standard way to import pandas

# Creating a Series from a list
# Let's imagine average high temps (F) for Nashville for a few months
temps = pd.Series([49, 54, 63, 72], index=['Jan', 'Feb', 'Mar', 'Apr'])

print("Nashville Average High Temps (Series):")
print(temps)
print("\nTemperature in March:", temps['Mar'])  # Access by index label
# Access by integer position
print("Temperature at position 0:", temps.iloc[0])
