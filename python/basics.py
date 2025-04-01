
print("hi")

place = "Nashville"
print(f"\nCharacters in '{place}':")
for character in place:
    print(character)

# An empty list
my_list = []

# A list of integers
numbers = [10, 20, 30, 40, 50]

# A list of strings (maybe some Nashville spots)
nashville_landmarks = [
    "Ryman Auditorium",
    "Grand Ole Opry",
    "Parthenon",
    "Country Music Hall of Fame",
]

# A list with mixed data types
mixed_data = ["Nashville", 615, True, 37.20]  # String, integer, boolean, float

# Accessing elements (uses 0-based indexing)
first_number = numbers[0]  # Gets 10
second_landmark = nashville_landmarks[1]  # Gets "Grand Ole Opry"

print(f"Numbers: {numbers}")
print(f"Landmarks: {nashville_landmarks}")
print(f"First number: {first_number}")
print(f"Second landmark: {second_landmark}")
print(f"Mixed data: {mixed_data}")
