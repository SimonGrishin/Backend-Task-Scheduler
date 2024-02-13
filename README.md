

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

you will need to download VScode, Docker Desktop
* MS VScode: https://code.visualstudio.com/download
* Docker: https://www.docker.com/products/docker-desktop/

you will need to set up MongoDB demo account
* MongoDB: https://account.mongodb.com/account/login
* Instructions: https://www.mongodb.com/docs/atlas/tutorial/create-atlas-account/
  * after set up create a collection names `tasks`  inside the `task-scheduler` collection for default. Otherwise update the config.yaml for _databaseName_ and _connectionString_
 


### Installation

_Below is an example of how you can instruct your audience on installing and setting up your app. This template doesn't rely on any external dependencies or services._

1. Get a free API Key at [https://example.com](https://example.com)
2. Clone the repo
   ```sh
   git clone https://github.com/your_username_/Project-Name.git
   ```
3. Install NPM packages
   ```sh
   npm install
   ```
4. Enter your API in `config.js`
   ```js
   const API_KEY = 'ENTER YOUR API';
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



#### Exercise 1: Configure task driven scheduler in Go

Develop a basic Go program that acts as a task scheduler. The program will read
tasks and their types from a YAML configuration file, and execute these tasks based on their
defined type. This exercise introduces basic Go syntax, enums, file operations, and working with
YAML configurations.

<!-- USAGE EXAMPLES -->
## Usage

Use this space to show useful examples of how a project can be used. Additional screenshots, code examples and demos work well in this space. You may also link to more resources.

_For more examples, please refer to the [Documentation](https://example.com)_

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [x] Add Changelog
- [x] Add back to top links
- [ ] Add Additional Templates w/ Examples
- [ ] Add "components" document to easily copy & paste sections of the readme
- [ ] Multi-language Support
    - [ ] Chinese
    - [ ] Spanish

See the [open issues](https://github.com/othneildrew/Best-README-Template/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Your Name - [@your_twitter](https://twitter.com/your_username) - email@example.com

Project Link: [https://github.com/your_username/repo_name](https://github.com/your_username/repo_name)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

Use this space to list resources you find helpful and would like to give credit to. I've included a few of my favorites to kick things off!

* [Choose an Open Source License](https://choosealicense.com)
* [GitHub Emoji Cheat Sheet](https://www.webpagefx.com/tools/emoji-cheat-sheet)
* [Malven's Flexbox Cheatsheet](https://flexbox.malven.co/)
* [Malven's Grid Cheatsheet](https://grid.malven.co/)
* [Img Shields](https://shields.io)
* [GitHub Pages](https://pages.github.com)
* [Font Awesome](https://fontawesome.com)
* [React Icons](https://react-icons.github.io/react-icons/search)

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
