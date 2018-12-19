#!/usr/bin/env python3

from flask import Flask, render_template
app = Flask(__name__)

@app.route("/")
def hello():
    # return "Hello, world"
    return render_template('index.html')

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=5050)
