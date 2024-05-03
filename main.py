import os
import requests
import bs4

from flask import Flask, Blueprint, render_template
from flask_cors import CORS
from dotenv import load_dotenv

from routes import main

app = None

if os.path.exists('.env'):
	load_dotenv()
else:
	print('No .env file found\nPlease create one.')
	exit()

if __name__ == '__main__':
	app = Flask(__name__)
	CORS(app)
	app.register_blueprint(main)
