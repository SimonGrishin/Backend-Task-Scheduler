

<!-- PROJECT LOGO -->
<br />
<div align="center">

  <h3 align="center">Backend-Task-Scheduler</h3>

  <p align="center">
    Task scheduling system using Golang, RESTful API, MongoDB, Docker, and Kubernetes.
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

[![Product Name Screen Shot][product-screenshot]](https://example.com)

This project was seperated into 5 exercises:

1. Configure task driven scheduler in Golang
2. Building a RESTful API with Gin
3. MongoDB: Integrating MongoDB with the Gin API
4. Docker: Containerizing the API with Docker
5. Kubernetes: Depolying the Containerized API on Kubernetes

The purpose of this project is to build a task scheduling system that builds upon itself to upgrade functionality.
This README will explain each step with visuals and Command-Line arguments to run the applications.
<br />
<br >



### Built With

Tools and Frameworks with links to learning resources

* [![Go][Go.js]][Go-url]
  * Learning: https://go.dev/tour/welcome/1
* [![Gin][Gin.js]][Gin-url]
  * Learning: https://go.dev/doc/tutorial/web-service-gin
* [![MongoDB][MongoDB.js]][MongoDB-url]
  * Learning: https://www.mongodb.com/docs/manual/tutorial/
* [![Docker][Docker.io]][Docker-url]
  * Learning: https://docs.docker.com/get-started/overview/
* [![Kub][Kub.dev]][Kub-url]
  * Learning: https://kubernetes.io/docs/tutorials/kubernetes-basics/

References to each tool and language is provided above.

<p align="right">(<a href="#readme-top">back to top</a>)</p>




<!-- GETTING STARTED -->
## Getting Started

The IDE used for this project is Microsoft's Visual Studio Code
It is able to run CLI and the extension manager is able to create dockerfiles, go files, YAML and JSON needed for configuration


### Prerequisites

you will need to download Go, VScode, Docker Desktop, and Postman
* Go: https://go.dev/doc/install
* MS VScode: https://code.visualstudio.com/download
* Docker: https://www.docker.com/products/docker-desktop/
* Postman: https://www.postman.com/downloads/

you will need to set up MongoDB demo account
* MongoDB: https://account.mongodb.com/account/login
* Instructions: https://www.mongodb.com/docs/atlas/tutorial/create-atlas-account/
  * after set up create a collection names `tasks`  inside the `task-scheduler` collection for default. Otherwise update the config.yaml for _databaseName_ and _connectionString_
 


### Installation

_These are the instructions for cloning the repository and installing any dependencies_

1. Clone the repo
   ```sh
   git clone https://github.com/SimonGrishin/Backend-Task-Scheduler.git
   ```

2. For any issues with required modules
   ```sh
   go get <module_name>@latest
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



#### Exercise 1: Configure task driven scheduler in Go

Develop a basic Go program that acts as a task scheduler. The program will read
tasks and their types from a YAML configuration file, and execute these tasks based on their
defined type. This exercise introduces basic Go syntax, enums, file operations, and working with
YAML configurations.

Running the task_scheduler.go:
```sh
go run task_scheduler.go
```

#### Exercise 2: Building a RESTful API with Gin

Expand upon the knowledge gained from the first exercise by developing a RESTful
API using the Gin framework. This API will manage a collection of tasks (similar to the ones in
the first exercise) with the ability to Create, Read, Update, and Delete (CRUD) tasks

Running the task_scheduler.go:
  ```sh
  go run restful_api.go
  ```

#### Exercise 3: MongoDB: Integrating MongoDB with Gin API

Enhance the task management API developed in the previous exercise by
integrating MongoDB for persistent data storage. This involves setting up a MongoDB database,
connecting it with the Go application, and modifying the API to perform CRUD operations on the
database.



<!-- USAGE EXAMPLES -->
## Usage

Use this space to show useful examples of how a project can be used. Additional screenshots, code examples and demos work well in this space. You may also link to more resources.

_For more examples, please refer to the [Documentation](https://example.com)_

<p align="right">(<a href="#readme-top">back to top</a>)</p>






<!-- CONTACT -->
## Contact

Simon Grishin - [@LinkedIn](https://linkedin/in/simongrishin) - simon.grishin@gmail.com

<p align="right">(<a href="#readme-top">back to top</a>)</p>





<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/othneildrew/Best-README-Template.svg?style=for-the-badge
[contributors-url]: https://github.com/othneildrew/Best-README-Template/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/othneildrew/Best-README-Template.svg?style=for-the-badge
[forks-url]: https://github.com/othneildrew/Best-README-Template/network/members
[stars-shield]: https://img.shields.io/github/stars/othneildrew/Best-README-Template.svg?style=for-the-badge
[stars-url]: https://github.com/othneildrew/Best-README-Template/stargazers
[issues-shield]: https://img.shields.io/github/issues/othneildrew/Best-README-Template.svg?style=for-the-badge
[issues-url]: https://github.com/othneildrew/Best-README-Template/issues
[license-shield]: https://img.shields.io/github/license/othneildrew/Best-README-Template.svg?style=for-the-badge
[license-url]: https://github.com/othneildrew/Best-README-Template/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/othneildrew
[product-screenshot]: images/screenshot.png
[Go.js]: https://img.shields.io/badge/Go-000020?style=for-the-badge&logo=go&logoColor=61DAFB
[Go-url]: https://go.dev
[Gin.js]: https://img.shields.io/badge/Gin-000000?style=for-the-badge&logo=Gin&logoColor=61DAFB
[Gin-url]: https://gin-gonic.com/
[MongoDB.js]: https://img.shields.io/badge/Mongodb-35495E?style=for-the-badge&logo=Mongodb&logoColor=4FC08D
[MongoDB-url]: https://www.mongodb.com/
[Docker.io]: https://img.shields.io/badge/Docker-384d54?style=for-the-badge&logo=Docker&logoColor=0db7ed
[Docker-url]: https://www.docker.com/
[Kub.dev]: https://img.shields.io/badge/kubernetes-a0b4d4?style=for-the-badge&logo=kubernetes&logoColor=3970e4
[Kub-url]: https://svelte.dev/
