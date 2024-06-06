### Prerequisites for running the application on your system

This project requires the following tools to be installed on your system:

1) Make

   i) Windows Users (Chocolatey):

   Install Chocolatey from https://chocolatey.org/.
   
   Open a terminal window and run:
   
        choco install make 
        
   ii) Mac users(Homebrew):
   
   Open Terminal and run:
   
        /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)" 
   
   Follow the on-screen instructions for Homebrew installation. Then, run:
   
        brew install make 

   iii) For Debian based distro run
   
         sudo apt-get update && sudo apt-get install make

   iv) For RPM based distro
   
         sudo yum update && sudo yum install make

2) Docker
   
    Visit the official Docker documentation: https://docs.docker.com/engine/install/
   
    Follow the installation instructions specific to your operating system (Windows, macOS, or Linux).

3) Install Mysql and configure it(The version used in the application is 8.0)

3) Create a .env file to inject the following environment variables into the container


        DB_USER=     Fill in the required details created as part of step3
        DB_PASSWORD=
        DB_HOST=
        DB_PORT=
        DB_NAME=
        MYSQL_ROOT_PASSWORD= 

5) Use make commands to run docker compose
### Why this Project?
#### This project is an imitation of the the milestones achieved during my internship

The projects is broken down into 10 Milestones each of which contains Learning Outcomes,Problem Statement, Expectations and Materials Required

<details>
<summary><b>1 - Create a simple REST API Webserver</b></summary>
  
### Learning Outcomes
  
Learnt about the best practices for REST APIs.

Learnt about the Twelve-Factor App methodology.

### Problem Statement

Create a student CRUD REST API using Golang and Gin


### Functional Requirement

Using the API we should be able to perform the following operations.

Add a new student.

Get all students.

Get a student with an ID.

Update existing student information.

Delete a student record.

### Expectations

Create a public repository on GitHub.

The repository should contain the following

        README.md file explaining the purpose of the repo, along with local setup instructions.
              
        Explicitly maintaining dependencies in a file ex (pom.xml, build.gradle, go.mod, requirements.txt, etc).
              
        Makefile to build and run the REST API locally.
              
        Ability to run DB schema migrations to create the student table.
              
        Config (such as database URL) should not be hard-coded in the code and should be passed through environment variables.
              
        Postman collection for the APIs.

API expectations
        Support API versioning (e.g., api/v1/<resource>).
        
        Using proper HTTP verbs for different operations.
        
        API should emit meaningful logs with appropriate log levels.
        
        API should have a /healthcheck endpoint.
        
        Unit tests for different endpoints.

### Further Reading

        The Twelve-Factor App
        
        Readme Driven Development
        
        Best Practices for REST API design

</details>

<details>
<summary><b>2 - Containerise REST API</b></summary>
<!-- Include details for Milestone 2 here -->
  
  ## Learning Outcomes

  Learnt how to Dockerise an application.

  Learnt about Mutli-stage Dockerfile.

  Learnt about Dockerfile best practices.

## Problem Statement
Create Dockerfile for the REST API.
## Expectations


    1) API should be run using the docker image.

    2) Dockerfile should have different stages to build and run the API.

    3) We should be able to inject environment variables while running the docker container at runtime.

    4) README.md should be updated with proper instructions to build the image and run the docker container.

    5) Similarly appropriate make targets should be added in the Makefile.

    6) The docker image should be properly tagged using semver tagging, use of latest tag is heavily discouraged.

    7) Appropriate measures should be taken to reduce docker image size. We want our images to have a small size footprint. 

## Sources

    Dockerfile Best Practices

    Advanced Dockerfile

    Hadolint

    Semantic versioning
</details>

<details>
<summary><b>3 - Setup one-click local development setup</b></summary>

## Problem Statement
  We want to simplify the process of setting up API on the local machine for development. The idea is to enable other   team members to run the API and its dependent services with the least amount of steps involved in getting this up and running. 
  
We won’t be assuming that other team members have the required tools already installed on their local. So we will be going one step further and providing them with simple bash functions to install the required tools.
## Expectations

  1)  API and its dependent services should be run using docker-compose.
    
  2)  README.md file should be updated with instructions
     
       i) To add pre-requisites for any existing tools that must already be installed (e.g., docker, make, etc)
      
  3)  When we run the make target to start the REST API docker container,
     
       i) Later it should invoke the docker compose command to start the API docker container and dependent services
      
 ## Source
 
 Docker compose docs

</details>




<details>
<summary><b>4 - Setup a CI Pipeline</b></summary>

#### Problem Statement:
Implementing a Continuous Integration (CI) pipeline is crucial for automating the build, test, and deployment processes of our application. The goal is to set up a CI pipeline that automatically builds and tests our codebase whenever changes are pushed to the repository.

#### Learning Outcomes:
- Understanding of CI/CD concepts and practices.
- Experience with popular CI tools such as Jenkins, GitLab CI, or GitHub Actions.
- Knowledge of configuring build scripts, running tests, and automating deployment tasks.

#### Expectations:
- Configure a CI pipeline to automatically build the project when changes are pushed to the repository.
- Include steps for running unit tests, integration tests, and any other relevant checks.
- Ensure that the CI pipeline integrates seamlessly with version control and notifies relevant stakeholders of build status.

#### Further Reading:
- [Continuous Integration, Delivery, and Deployment](https://www.atlassian.com/continuous-delivery/ci-vs-ci-vs-cd)
- [Introduction to GitLab CI/CD](https://docs.gitlab.com/ee/ci/)
- [GitHub Actions Documentation](https://docs.github.com/en/actions)

</details>

<details>
<summary><b>5 - Deploy REST API & its dependent services on Bare Metal</b></summary>

#### Problem Statement:
Deploying the REST API and its dependent services on bare metal servers provides full control over the infrastructure. The objective is to set up and configure the necessary servers to host the application and ensure its availability and reliability.

#### Learning Outcomes:
- Understanding of server provisioning and configuration management.
- Experience with deploying and managing applications on physical servers.
- Knowledge of networking, security, and monitoring practices for bare metal environments.

#### Expectations:
- Provision bare metal servers and install the required operating system and dependencies.
- Configure network settings, firewall rules, and security measures to protect the servers.
- Deploy the REST API and its dependent services, ensuring high availability and scalability.

</details>


<details>
<summary><b>6 - Setup Kubernetes cluster</b></summary>

#### Problem Statement:
Setting up a Kubernetes cluster provides a scalable and resilient platform for deploying containerized applications. The objective is to configure and deploy a Kubernetes cluster that can host our application and its dependent services.

#### Learning Outcomes:
- Understanding of Kubernetes architecture and components.
- Experience with provisioning and configuring Kubernetes clusters.
- Knowledge of deploying and managing applications in a Kubernetes environment.

#### Expectations:
- Provision a Kubernetes cluster on a cloud provider or on-premises infrastructure.
- Configure cluster networking, storage, and security settings.
- Deploy necessary Kubernetes resources such as pods, services, and ingresses for hosting the application.

</details>

<details>
<summary><b>7 - Deploy REST API & its dependent services in K8s</b></summary>

#### Problem Statement:
Deploying the REST API and its dependent services in a Kubernetes cluster leverages the benefits of container orchestration. The goal is to containerize the application components and deploy them as Kubernetes resources.

#### Learning Outcomes:
- Experience with containerization using Docker or other container runtimes.
- Understanding of Kubernetes deployment manifests and resource definitions.
- Knowledge of service discovery, load balancing, and scaling in Kubernetes.

#### Expectations:
- Containerize the REST API and its dependent services using Docker or other containerization tools.
- Write Kubernetes deployment and service manifests for each application component.
- Deploy the containerized applications to the Kubernetes cluster and verify their functionality.

</details>

<details>
<summary><b>8 - Deploy REST API & its dependent services using Helm Charts</b></summary>

#### Problem Statement:
Using Helm charts simplifies the deployment and management of complex applications in Kubernetes. The objective is to package the REST API and its dependent services into Helm charts and deploy them using Helm.

#### Learning Outcomes:
- Understanding of Helm chart structure and templating.
- Experience with creating custom Helm charts for applications.
- Knowledge of Helm commands for installing, upgrading, and managing releases.

#### Expectations:
- Organize the application components into Helm chart templates.
- Parameterize the Helm charts to allow customization during deployment.
- Deploy the Helm charts to the Kubernetes cluster using Helm and verify successful deployment.

</details>

<details>
<summary><b>9 - Setup one-click deployments using ArgoCD</b></summary>

#### Problem Statement:
ArgoCD provides a GitOps workflow for declarative continuous delivery of Kubernetes applications. The goal is to set up ArgoCD to automate deployments and maintain application configuration in sync with Git repositories.

#### Learning Outcomes:
- Understanding of GitOps principles and practices.
- Experience with setting up and configuring ArgoCD for Kubernetes clusters.
- Knowledge of managing application deployments and configuration with ArgoCD.

#### Expectations:
- Install and configure ArgoCD in the Kubernetes cluster.
- Define application manifests and sync policies in ArgoCD repositories.
- Automate application deployments by syncing changes from Git repositories to the cluster.

</details>

<details>
<summary><b>10 - Setup an observability stack</b></summary>

#### Problem Statement:
Monitoring and observability are essential for understanding the health and performance of our application. The objective is to set up an observability stack that includes logging, metrics, and tracing capabilities.

#### Learning Outcomes:
- Understanding of observability concepts and tools.
- Experience with setting up monitoring and logging solutions for distributed systems.
- Knowledge of analyzing and troubleshooting application issues using observability data.

#### Expectations:
- Deploy logging solutions such as Elasticsearch, Fluentd, and Kibana (EFK) or Loki and Grafana (Promtail).
- Set up metrics collection using Prometheus and visualize metrics with Grafana.
- Instrument the application code for distributed tracing using tools like Jaeger or Zipkin.

</details>

<details>
<summary><b>11 - Configure dashboards & alerts</b></summary>

#### Problem Statement:
Dashboards and alerts provide insights into the health and performance of our application and infrastructure. The goal is to configure dashboards and alerts based on key metrics and events.

#### Learning Outcomes:
- Understanding of monitoring dashboards and alerting systems.
- Experience with configuring dashboards and defining alerting rules.
- Knowledge of responding to alerts and troubleshooting issues proactively.

#### Expectations:
- Create dashboards in Grafana or similar tools to visualize important metrics and trends.
- Define alerting rules based on thresholds, anomalies, or specific events.
- Set up notification channels (e.g., email, Slack) to receive alerts and respond to incidents promptly.

</details>

