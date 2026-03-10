#!/usr/bin/env oak shell
# Example shell script for Oak virtual shell

# Print header
echo ================================
echo Oak Virtual Shell Demo Script
echo ================================
echo

# Show current directory
echo Current directory:
pwd
echo

# Create project structure
echo Creating project structure...
mkdir /myproject
cd /myproject
mkdir src
mkdir lib
mkdir test
mkdir docs

# Create some files
echo Creating files...
touch src/main.oak
touch README.md
touch .gitignore

# Show the structure
echo
echo Project structure created:
ls
echo

# Create a simple config file
echo Creating configuration...
cd /myproject
export PROJECT_NAME=MyOakProject
export VERSION=1.0.0
export AUTHOR=Developer

# Show environment
echo
echo Environment variables:
env
echo

# Create subdirectories
echo Creating source structure...
cd src
mkdir models
mkdir views
mkdir controllers
touch models/user.oak
touch views/index.oak
touch controllers/app.oak

echo
echo Source directory contents:
ls
echo

# Back to project root
cd /myproject

# Display final message
echo
echo ================================
echo Project setup complete!
echo ================================
echo
echo Next steps:
echo   1. Edit src/main.oak
echo   2. Add your code in src/
echo   3. Run: oak src/main.oak
echo
