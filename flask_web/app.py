#import main Flask class and request object
from flask import Flask, jsonify, request
import requests

# Flask App
app = Flask(__name__)

# URL declaration
reg_url = 'https://localhost:9000'

# Defining Routes
@app.route('/')
def hello_world():
    res = requests.get(reg_url, verify=False)
    res = res.json()
    return res['message'][::-1]
    
def main():
    app.run(debug=True, host='0.0.0.0')
if __name__ == '__main__':
    main()
