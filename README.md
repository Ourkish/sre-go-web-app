# sre-web-app

## Objective

- Run a provided sre-web-app application
- Implement a deployment system with Kubernetes (minikub, k3d ... ) and Helm.
- You can improve, propose and implement everything you want.
- Monitor the application.

## Prerequisites

- Linux workstation (I am using Debian12) with at least 4 Gb of ram.
  `````
  cat /etc/*release*
  `````
- Install the following packages [curl, Docker, Minikube, kubectl, Helm, Terraform]:
  `````
  sudo mkdir /home/debian/sre-web-app
  sudo cd /home/debian/sre-web-app
  git clone https://github.com/Ourkish/sre-go-web-app.git
  sudo chmod +x auto_install.sh
  sudo ./auto_install.sh
  `````
- get the web app
  The application is a Hello World web server with two path: / and /metrics.
  The binary and its sources are located to /root
  create a Dockerfile for the app and test it locally :
  
## Understanding the operation

- Minikube provides a local Kubernetes environment where you can deploy and manage containerized applications using Kubernetes resources.
- Docker containers can be built and pushed to container registries. Kubernetes can then pull these containers and deploy them as pods (the smallest deployable units in Kubernetes).
- kubectl is used to interact with the Minikube Kubernetes cluster. You can create pods, services, deployments, and other Kubernetes resources using kubectl commands.
- Helm simplifies the installation of complex applications in Kubernetes. Helm charts describe how to deploy applications, including the necessary Kubernetes resources. You can use Helm to install and manage applications in your Minikube cluster.
- Terraform can manage Kubernetes resources in Minikube as well as other infrastructure resources outside of Kubernetes. It uses Terraform providers to interact with the Kubernetes API and provision resources like namespaces, secrets, and ConfigMaps.

## In Practice

- Quick Start :
  Clone the project on your environement and run those commands (if not already done !) :
  `````
  git clone https://github.com/Ourkish/sre-web-app.git
  sudo chmod +x dep_install.sh
  sudo ./auto_install.sh
  sudo terraform apply
  `````
  
- Allow EST/WEST traffic on exposed ports :
  `````
  kubectl port-forward --address 57.128.112.35 service/grafana 3000:80
  kubectl port-forward --address 57.128.112.35 service/prometheus-server 9090:80
  kubectl port-forward --address 57.128.112.35 svc/sre-hiring-service 8080:8080
  `````

- Add a cron to check disk and to be notified by email if disk run's out of space :
  `````
  sudo vim /etc/crontab
  */30 * * * * /home/debian/sre-web-app/healthchecks_app/healthchecks
  `````

  you can configure your creds for SMTP service here :
  `````
  sudo vim /home/ubuntu/sre-web-app/HealthChecks.go
  recipient := "your.email@example.com"
  smtpServer := "smtp.example.com"
  smtpPort := "587"
  smtpUsername := "your.smtp.username"
  smtpPassword := "your.smtp.password"
  `````
## Potential Improvements

- Continuous Integration & Deployment (CI/CD) :
Implement a CI/CD pipeline using tools like Jenkins, GitLab CI.., This pipeline would automatically build the Docker container, push it to a registry, and deploy it to the Kubernetes cluster whenever 
there are updates to the main branch of the repository.

- Logging & Observability :
Integrate a centralized logging solution like Graylog or Loki to gather logs from the application and Kubernetes nodes.

- Scalability :
Optimize the application's Kubernetes deployment to use Horizontal Pod Autoscaling. This ensures the application scales out based on CPU or memory usage.

- Security Enhancements :
Implement Network Policies in Kubernetes to restrict the communication between pods.
Regularly scan Docker images for vulnerabilities.
Encrypt sensitive data using Kubernetes secrets.

- High Availability :
If moving beyond Minikube, consider setting up a multi-node Kubernetes cluster across different zones to ensure high availability.

- Backup & Disaster Recovery :
Set up regular backups of the Kubernetes cluster configuration and data. 
Document and periodically test a disaster recovery procedure.

## Showcase :

- Below are screenshots showcasing the results following the deployment and execution of the code.

  Deploy using Helm and Terraform :

  ![deploy](https://github.com/Ourkish/sre-web-app/assets/67292535/9b5cdf6d-1465-459c-8f42-d3e8b31b8ec5)

  App metrics :

  ![app_metrics](https://github.com/Ourkish/sre-web-app/assets/67292535/ba6dde2c-e924-48a6-806b-4a532617b31b)

  Application Monitoring :
  
  ![go_metrics](https://github.com/Ourkish/sre-web-app/assets/67292535/0ccf2924-62ec-4f6f-be02-bb7065244781)

  ![go_process](https://github.com/Ourkish/sre-web-app/assets/67292535/abeb9b48-7b5f-4813-b773-f76ae5ca802a)

  show pods after deployment :
  
  ![get_pods](https://github.com/Ourkish/sre-web-app/assets/67292535/88317871-704f-423e-aad2-044a46898682)

  Kubernetes Dashboard :
  
  ![kub_metrics](https://github.com/Ourkish/sre-web-app/assets/67292535/d1a7a269-31da-43b4-816f-187bec36391b)

  To destroy the cluster use this command :
  `````
  sudo terraform destroy
  `````
  you will be asked to confirme, if you agree to destroy you should see an output like this :
  
  ![destroy_cluster](https://github.com/Ourkish/sre-web-app/assets/67292535/f59a86f6-ebb8-4583-b93a-82ba823591c6)
