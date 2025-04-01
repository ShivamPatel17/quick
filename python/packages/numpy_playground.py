import numpy as np  # Standard convention to import as np

# Create a NumPy array from a Python list
numpy_numbers = np.array([10, 20, 30, 40, 50])

# Create a NumPy array of floats
numpy_floats = np.array([3.14, 2.71, 1.618], dtype=float)  # Can specify dtype

# NumPy arrays allow fast element-wise operations
doubled_numbers = numpy_numbers * 2

# Create a 2-dimensional array (matrix)
matrix = np.array([[1, 2], [3, 4]])

print(f"\nNumPy integer array: {numpy_numbers}")
print(f"NumPy float array: {numpy_floats}")
print(f"Doubled numbers: {doubled_numbers}")
print(f"NumPy 2D array (matrix):\n{matrix}")

for dn in doubled_numbers:
    print(dn)
