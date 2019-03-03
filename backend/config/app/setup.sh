#!/bin/sh

# pip3: Install from requirements.txt
pip3 install -r $PY_REQUIREMENTS_PATH

# Run Flask server
python3 $PY_FLASK_APP_PATH
