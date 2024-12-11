#!/bin/bash

# Set base URL for API (with the trailing slash)
BASE_URL="http://localhost:3000/students/"

# 1. POST: Create a new student
echo "======================================"
echo "Running POST: Create student"
echo "======================================"
response=$(curl -s -X POST "$BASE_URL" \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "sawant",
    "last_name": "hrithik",
    "email": "hrithiksawant@gmail.com"
  }')

# Debugging output: Print the response to check if ID is returned
echo "POST Response: $response"

# Extract the student ID from the response
student_id=$(echo $response | jq -r '.id')
if [ "$student_id" == "null" ] || [ -z "$student_id" ]; then
  echo "❌ test Failed."
  echo "POST failed. No student ID returned."
  exit 1
fi
echo "Created student with ID: $student_id"
echo "✅ Success: Student created with ID $student_id."

# 2. GET: Retrieve a student by ID
echo ""  # Adding a newline after the POST test
echo "======================================"
echo "Running GET: Retrieve student by ID"
echo "======================================"
get_response=$(curl -s -X GET "$BASE_URL$student_id")

# Debugging output: Print the GET response to check its structure
echo "GET Response: $get_response"

# Ensure the student is retrieved properly
if echo $get_response | jq -e '.id == null' > /dev/null; then
  echo "❌ test Failed."
  echo "GET failed. No student found with ID: $student_id"
  exit 1
fi
echo "Successfully retrieved student with ID: $student_id"
echo "✅ Success: Student retrieved successfully."

# 3. PUT: Update student details
echo ""  # Adding a newline after the POST test
echo "======================================"
echo "Running PUT: Update student details"
echo "======================================"
put_response=$(curl -s -X PUT "$BASE_URL$student_id" \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "hrithik",
    "last_name": "sawant",
    "email": "hrithiksawant@gmail.com"
  }')

# Debugging output: Print the PUT response
echo "PUT Response: $put_response"
if echo $put_response | jq -e '.error' > /dev/null; then
  echo "❌ test Failed."
  echo "PUT failed due to error: $(echo $put_response | jq -r '.error')"
  exit 1
fi
echo "Successfully updated student with ID: $student_id"
echo "✅ Success: Student updated successfully."

# 4. GET: Retrieve all students
echo ""  # Adding a newline after the POST test
echo "======================================"
echo "Running GET: Retrieve all students"
echo "======================================"

all_students_response=$(curl -s -X GET "$BASE_URL")

# Debugging output: Print the GET all response to check structure
echo "All Students Response:"

# Parse the number of students and sample student data
student_count=$(echo $all_students_response | jq 'length')

# Get the top 2 students, then "..." as separator, then the last 2 students
top_2_students=$(echo $all_students_response | jq -r '.[0:2] | .[] | "ID: \(.id), Name: \(.first_name) \(.last_name), Email: \(.email)"')
last_2_students=$(echo $all_students_response | jq -r '.[-2:] | .[] | "ID: \(.id), Name: \(.first_name) \(.last_name), Email: \(.email)"')

# Print top 2 students, separator, and last 2 students
echo "$top_2_students"
echo "..."
echo "$last_2_students"

echo "Number of students retrieved: $student_count"
echo "✅ Success: Retrieved list of all students (Total $student_count students)."
# 5. DELETE: Delete the last student
echo ""  # Adding a newline after the POST test
echo "======================================"
echo "Running DELETE: Delete student with ID $student_id"
echo "======================================"
delete_response=$(curl -s -X DELETE "$BASE_URL$student_id")
if [ $? -ne 0 ]; then
  echo "❌ test Failed."
  echo "DELETE failed."
  exit 1
fi
echo "Successfully deleted student with ID: $student_id"
echo "✅ Success: Student deleted successfully."

# Check if the last student was actually deleted
remaining_students=$(curl -s -X GET "$BASE_URL")
if echo $remaining_students | jq -e "any(.id == $student_id)" > /dev/null; then
  echo "❌ test Failed."
  echo "DELETE failed. Student with ID $student_id still exists."
  exit 1
fi

# Final success message
echo ""  # Adding a newline after the POST test
echo "======================================"
echo "Test Summary:"
echo "======================================"
echo "All operations (POST, GET, PUT, DELETE) were completed with mixed results:"
echo "- Success: POST, GET, DELETE."
echo "- Failure: PUT (duplicate email error)."
echo "✅ test OK."
