<!-- <p align="center"><img width=12.5% src="https://github.com/anfederico/Clairvoyant/blob/master/media/Logo.png"></p>
<p align="center"><img width=60% src="https://github.com/anfederico/Clairvoyant/blob/master/media/Clairvoyant.png"></p>

![React](https://img.shields.io/badge/python-v3.6+-blue.svg)
[![Build Status](https://travis-ci.org/anfederico/Clairvoyant.svg?branch=master)](https://travis-ci.org/anfederico/Clairvoyant)
![Dependencies](https://img.shields.io/badge/dependencies-up%20to%20date-brightgreen.svg)
[![GitHub Issues](https://img.shields.io/github/issues/anfederico/Clairvoyant.svg)](https://github.com/anfederico/Clairvoyant/issues)
![Contributions welcome](https://img.shields.io/badge/contributions-welcome-orange.svg)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)
-->
## Overview
This is a project designed to help developers of all levels have a better understanding on how to build a complex web application with a unique set of languages and tools from scratch, without relying on toolchains.

What will you learn in the process:

- Go
- GORM (Go ORM[Object Relational Mapping])
- PostgreSQL
- GraphQL
- React
- Webpack

> Still under initial development. Coming out soon.

##### If you have a particular web application in mind that I'd like to build, please raise it an issue.


## Backend [Go]

###GraphQL
#### Query
		Queries:
		=======================
		Fetch tut with id = 2
		=======================
		{
	  		tutorial(id: 2) {
	    		title,
	    		comments {
	      			body
	    		}
	  		}
		}
		
		Mutation:
		=======================
		Create tut with title = 'Third Tut'
		=======================
		mutation M {
	  		newTut: createTutorial(title: "Second Tutorial") {
	    		title
	  		}
		}
