import array

# Create an array of signed integers ('i' is the type code)
integer_array = array.array("i", [10, 20, 30, 40, 50])

# Create an array of double-precision floats ('d' is the type code)
float_array = array.array("d", [3.14, 2.71, 1.618])


# Trying to add a different type will cause an error:
# integer_array.append("hello") # This would raise a TypeError

print(f"\nInteger array: {integer_array}")
print(f"Float array: {float_array}")
print(f"First integer: {integer_array[0]}")  # Accessing works like lists


for i in integer_array:
    print(i)
